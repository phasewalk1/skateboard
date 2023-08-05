package bootstrap

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/phasewalk1/skateboard/config"

	"github.com/charmbracelet/log"
	lua "github.com/yuin/gopher-lua"
)

func LoadTrucksContract(path string) (config.Config, error) {
	log.Debug("starting a new Lua state")

	L := lua.NewState()
	log.Debug("defer:", "closing Lua state")
	defer L.Close()

	home, err := SkateboardPath()
	if err != nil {
		home = "~/.skateboard/"
	}

	log.Debug("home", "active", home)
	homeIncl := fmt.Sprintf("%s/include", home)
	log.Debug("include", "active", home)

	extendLuaPath := fmt.Sprintf("package.path = package.path .. ';%s/?.lua'", homeIncl)
	log.Debug("lua", "extending lua path", extendLuaPath)
	if err := L.DoString(extendLuaPath); err != nil {
		log.Fatal("failed to extend Lua package path:", err)
		return config.Config{}, err
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to get pwd:", err)
		return config.Config{}, err
	}
	log.Debug("env", "pwd", pwd)

	execFnl := "config = require 'fennel'.dofile('trucks.contract.fnl'); return config"
	if err := L.DoString(execFnl); err != nil {
		log.Fatal("failed to execute 'fennel.dofile':", err)
		return config.Config{}, err
	}

	luaConfig := L.Get(-1)
	log.Debug("config", "luaConfig", luaConfig.String())

	if _, ok := luaConfig.(*lua.LTable); !ok {
		log.Warn("config", "execution", ok)
		log.Debug("config", "return", luaConfig.String())
		log.Fatal("config", "fail", "failed to load trucks.contract.fnl config table")
		return config.Config{}, fmt.Errorf("failed to load trucks.contract.fnl")
	}

	var cfg config.Config
	ToTable(luaConfig.(*lua.LTable), &cfg)
	log.Info("trucks contract is valid ✓")
	log.Debug("config", "system", cfg.System)
	log.Debug("config", "services", cfg.Services)
	log.Debug("config", "defaults", cfg.Defaults)

	return cfg, nil
}

func luaValueToGoValue(value lua.LValue) interface{} {
	switch v := value.(type) {
	case *lua.LTable:
		return luaTableToGoMap(v)
	case lua.LNumber:
		return float64(v)
	case lua.LString:
		return string(v)
	case lua.LBool:
		return bool(v)
	default:
		return nil
	}
}

func luaTableToGoMap(table *lua.LTable) map[string]interface{} {
	goMap := make(map[string]interface{})
	table.ForEach(func(key lua.LValue, value lua.LValue) {
		switch key.String() {
		case "services":
			services := make([]map[string]interface{}, 0)
			value.(*lua.LTable).ForEach(func(_ lua.LValue, serviceValue lua.LValue) {
				service := luaTableToGoMap(serviceValue.(*lua.LTable))
				services = append(services, service)
			})
			goMap[key.String()] = services
		default:
			goMap[key.String()] = luaValueToGoValue(value)
		}
	})
	return goMap
}

func ToTable(table *lua.LTable, output interface{}) {
	goMap := luaTableToGoMap(table)

	jsonData, err := json.Marshal(goMap)
	if err != nil {
		fmt.Println("Error converting map to JSON:", err)
		return
	}

	err = json.Unmarshal(jsonData, output)
	if err != nil {
		fmt.Println("Error unmarshaling JSON into struct:", err)
	}
}
