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
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var rootCmd = &cobra.Command{
	Use:   "skateboard",
	Short: "Pack your services once",
	Long:  `Onboarding as fast as light?`,

	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

var mkdocsCmd = &cobra.Command{
	Use:   "mkdocs",
	Short: "Generate documentation",
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd, "./docs/dist/cobra")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(".md documentation generated successfully in 'docs/cobra/'")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Bool("down", false, "Take down the whole thing")
	rootCmd.AddCommand(mkdocsCmd)
	log.SetLevel(log.DebugLevel)
}
