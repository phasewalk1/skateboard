/*
Copyright (C) 2023 Ethan Gallucci

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/phasewalk1/skateboard/bootstrap"
	"github.com/phasewalk1/skateboard/config"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

var home string

// TODO:
// - [x] Impl `check` for trucks
// - [ ] Impl `check` for yaml
// - [ ] Impl `check` for toml
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks a contract for validity",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("running check command")

		log.Debug("starting a new Lua state")
		L := lua.NewState()
		log.Debug("defer: closing Lua state")
		defer L.Close()

		home, _ := cmd.Flags().GetString("home")
		log.Debug("active-home: ", home)
		homeIncld := fmt.Sprintf("%s/include", home)
		log.Debug("active-include: ", homeIncld)
		extendLuaPath := fmt.Sprintf("package.path = package.path .. ';%s/?.lua'", homeIncld)
		log.Debug("extending Lua package.path: ", extendLuaPath)
		if err := L.DoString(extendLuaPath); err != nil {
			log.Fatal("failed to extend Lua package.path: ", err)
			return
		}
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal("failed to get current working directory: ", err)
			return
		}
		log.Debug("current working directory: ", pwd)
		execFnl := "config = require 'fennel'.dofile('trucks.contract.fnl'); return config"
		if err := L.DoString(execFnl); err != nil {
			log.Fatal("failed to execute fennel.dofile('trucks.contract.fnl'): ", err)
			return
		}

		luaConfig := L.Get(-1)
		fmt.Println(luaConfig.String())

		if _, ok := luaConfig.(*lua.LTable); !ok {
			log.Warn("fennel")
			log.Debug("returned value: ", luaConfig.String())
			log.Fatal("failed to load trucks.contract.lua: ", err)
			return
		}
		var cfg config.Config
		ToTable(luaConfig.(*lua.LTable), &cfg)
        log.Info("trucks contract is valid ✓")
		cfg.Show()
	},
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

func init() {
	rootCmd.AddCommand(checkCmd)
	home, err := bootstrap.SkateboardPath()
	if err != nil {
		home = "~/.skateboard"
	}
	checkCmd.Flags().StringVarP(&home, "home", "H", home, "path to skateboard home directory")
}
