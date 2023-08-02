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
	"github.com/phasewalk1/skateboard/config"

	"github.com/spf13/cobra"
	"github.com/yuin/gopher-lua"
)

var trucksCmd = &cobra.Command{
	Use:   "trucks",
	Short: "Use trucks to write your config in Fennel/Lua",
	Long:  `This command checks a contract for validity`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trucks called")
		L := lua.NewState()
		defer L.Close()

		if err := L.DoFile("include/trucks.contract.lua"); err != nil {
			panic(err)
		}

		luaConfig := L.Get(-1)
		fmt.Println(luaConfig.String())

		if _, ok := luaConfig.(*lua.LTable); !ok {
			fmt.Println("the returned value is not a valid config table")
			return
		}
		var cfg config.Config
		ToTable(luaConfig.(*lua.LTable), &cfg)
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
	rootCmd.AddCommand(trucksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trucksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trucksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
