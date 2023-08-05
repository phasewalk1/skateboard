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
	"path/filepath"

	"github.com/phasewalk1/skateboard/bootstrap"
	"github.com/phasewalk1/skateboard/util"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var force bool

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync the services based on the contract",
	Long: `This command will sync services based on the 'skateboard.toml' contract.
It loads the service map from the contract and performs operations based on it.`,
	Run: func(cmd *cobra.Command, args []string) {
		contractPath := filepath.Join(".", "skateboard.toml")
		sysConfig, err := bootstrap.LoadTrucksContract(contractPath)
		if err != nil {
			log.Fatalf("Error loading workload contract from file: %v", err)
		}
		util.SyncRemotes(sysConfig.Services, force)
	},
}

func init() {
	syncCmd.Flags().BoolVarP(&force, "force", "f", false, "Force sync with new clones")
	rootCmd.AddCommand(syncCmd)
}
