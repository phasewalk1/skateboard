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
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Run services manually",
	Run: func(cmd *cobra.Command, args []string) {
		prof, _ := cmd.Flags().GetString("profile")

		if prof == "default" {
			log.Println("[profile]: @default")
		} else {
			log.Println("[profile]: @{}", prof)
		}

		build, _ := cmd.Flags().GetBool("build")
		log.Println(profile(prof))
		log.Println("Build: ", build)

		dc := append([]string{"up", "--build"}, args...)
		dcCmd := exec.Command("docker-compose", dc...)

		dcCmd.Stdout = os.Stdout
		dcCmd.Stderr = os.Stderr
		log.Println("build called")

		err := dcCmd.Run()
		if err != nil {
			log.Fatal("Error running docker compose:", err)
		}
	},
}

func profile(prof string) string {
	return fmt.Sprintf("docker-compose.%s.yml", prof)
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	serviceCmd.Flags().StringP("profile", "p", "default", "Use a container profile")
	serviceCmd.Flags().BoolP("build", "b", false, "Build the containers before running them")
}
