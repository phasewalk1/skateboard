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
	"os"
	"path/filepath"
	"sync"

	"github.com/phasewalk1/skateboard/bootstrap"
	"github.com/phasewalk1/skateboard/util"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var wg sync.WaitGroup
var startupWg sync.WaitGroup

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Spinup the application and all its services",
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		path, _ := cmd.Flags().GetString("path")
		newClone, _ := cmd.Flags().GetBool("new-clone")
		noSync, _ := cmd.Flags().GetBool("no-sync")

		contractPath := filepath.Join(".", path)
		log.Debug("skateboard.up", "[load.contract]:", contractPath)

		ctx, cancel := context.WithCancel(context.Background())

		if filepath.Ext(contractPath) == ".fnl" {
			sysConfig, err := bootstrap.LoadTrucksContract(contractPath)
			sysConfig.EmplaceDefaults()
			if err != nil {
				log.Fatal("config", "fennel", err)
			}
			if newClone {
				log.Info("sync", "true", "syncing service repositories...")
				util.SyncRemotes(sysConfig.Services, force)
			} else {
				log.Info("sync", "false", "skipping sync operation")
			}
			log.Info("service!", "starting services...", true)
			bootstrap.Upper(&sysConfig, noSync, ctx, cancel)

			done := make(chan struct{})
			go func() {
				select {
				case <-done:
					log.Warn("all services have been stopped", "exiting!", 0)
					os.Exit(0)
				}
			}()

			go func() {
				wg.Wait()
				close(done)
			}()
		} else {
			log.Fatal("contract", "failed to read contract path:", contractPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// sync repositories w/ HEAD before running
	upCmd.Flags().BoolP(
		"new-clone",
		"n",
		false,
		"sync repositories w/ HEAD before running, use --force to override existing copies",
	)
	// rm existing copies before cloning a new copy
	upCmd.Flags().BoolP("force", "f", false, "Pass --force when syncing a new clone")
	// override default path to contract file
	upCmd.Flags().StringP("path", "p", "trucks.contract.fnl", "Path to the contract file")
	// do not run 'service.sync', i.e., skip operations like 'npm install' before running
	upCmd.Flags().BoolP("no-sync", "x", false, "skip running 'service.sync' on all services")
}
