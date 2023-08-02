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
	"context"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/phasewalk1/skateboard/bootstrap"
	"github.com/phasewalk1/skateboard/config"
	"github.com/phasewalk1/skateboard/util"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var noSync bool
var wg sync.WaitGroup
var startupWg sync.WaitGroup

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Spinup the application and all its services",
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		contractPath := filepath.Join(".", "skateboard.toml")

		sysConfig, err := config.LoadWorkloadContractFromFile(contractPath)
		log.Debug("sysConfig.System.Panic:", sysConfig.System.Panic)

		ctx, cancel := context.WithCancel(context.Background())
		if !noSync {
			if err != nil {
				log.Fatalf("Error loading workload contract from file: %v", err)
			}
			fmt.Println("Syncing service repositories...")
			util.SyncRemotes(sysConfig.Services, force)
		} else {
			fmt.Println("Skipped sync operation")
		}
		bootstrap.Upper(sysConfig, ctx, cancel)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	upCmd.Flags().BoolVarP(&noSync, "no-sync", "n", false, "Skip syncing repositories")
	upCmd.Flags().BoolVarP(&force, "force", "f", false, "Pass --force when syncing")
}
