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
	boot "github.com/phasewalk1/skateboard/bootstrap"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install skateboard",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("checking for the existince of $HOME/.skateboard")
		force, _ := cmd.Flags().GetBool("force")
		sk8, err := boot.SkateboardPath()
		if err != nil {
			log.Fatal("couldn't get $HOME:", err)
			return
		}

		exis, err := boot.SkateboardExists(sk8)
		if err != nil {
			log.Fatal("couldn't check for $HOME/.skateboard:", err)
			return
		}
		if exis && !force {
			log.Warn("skateboard is already installed, use --force to reinstall")
			return
		}

		trucksToggle, _ := cmd.Flags().GetBool("trucks")
		noDeps, _ := cmd.Flags().GetBool("no-deps")
		log.Info("installing skateboard")
		scope := "skateboard.install"
		log.Debug(scope, "skateboard path:", sk8)
		log.Debug(scope, "trucks:", trucksToggle)
		log.Debug(scope, "force:", force)
		log.Debug(scope, "no-deps:", noDeps)
		err = boot.TryMakeSkateboard(sk8, force, trucksToggle, noDeps)
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().BoolP("trucks", "t", true, "install trucks")
	installCmd.Flags().BoolP("force", "f", false, "Force install skateboard even if it already exists")
	installCmd.Flags().BoolP("no-deps", "x", false, "Don't install dependencies")
}
