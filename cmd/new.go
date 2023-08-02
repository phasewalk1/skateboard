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
	"github.com/phasewalk1/skateboard/bootstrap"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new <path>",
	Short: "Scaffold a new contract repository",
	Long:  `Creates a working git repository at path and scaffolds an example skateboard contract`,
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		trucks, _ := cmd.Flags().GetBool("trucks")
		if !trucks {
			ymode, _ := cmd.Flags().GetBool("yaml")
			if ymode == false {
				log.Debug("mode: toml")
			} else {
				log.Debug("mode: yaml")
			}
		} else {
			log.Debug("mode: trucks (default)")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("bootstrapping", args[0])
		trucks, _ := cmd.Flags().GetBool("trucks")
		force, _ := cmd.Flags().GetBool("force")
		if trucks {
			bootstrap.BootstrapDir(args[0], false, false, force)
		} else {
			ymode, _ := cmd.Flags().GetBool("yaml")
			tmode, _ := cmd.Flags().GetBool("toml")
			if ymode && tmode {
				log.Fatalf("Can't toggle both yaml and toml")
				return
			} else {
				bootstrap.BootstrapDir(args[0], ymode, tmode, force)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().Bool("trucks", true, "Use Trucks for config")
	newCmd.Flags().BoolP("yaml", "y", false, "Use YAML for config")
	newCmd.Flags().BoolP("toml", "t", false, "Use TOML for config")
	newCmd.Flags().BoolP("force", "f", false, "Force create the new workspace")
}
