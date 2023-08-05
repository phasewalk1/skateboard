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
	"path/filepath"

	"github.com/phasewalk1/skateboard/bootstrap"

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
		path := filepath.Join(".", "trucks.contract.fnl")
		if _, err := os.Stat(path); err != nil {
			path = filepath.Join(".", "trucks.contract.toml")
		}
		if _, err := bootstrap.LoadTrucksContract(path); err != nil {
			log.Fatal("check", "contract", err)
			return
		}
		return
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
