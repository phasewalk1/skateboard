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
	"os"
	"path/filepath"

	"github.com/phasewalk1/skateboard/bootstrap"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
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

func init() {
	rootCmd.AddCommand(checkCmd)
	home, err := bootstrap.SkateboardPath()
	if err != nil {
		home = "~/.skateboard"
	}
	checkCmd.Flags().StringVarP(&home, "home", "H", home, "path to skateboard home directory")
}
