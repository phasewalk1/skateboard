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

package util

import (
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/phasewalk1/skateboard/config"
)

func getSvcName(svc config.Service) string {
	parts := strings.Split(svc.Name, ".")
	return parts[len(parts)-1]
}

func SyncRemotes(svcSlice []config.Service, force bool) error {
	var wg sync.WaitGroup
	for _, svc := range svcSlice {
		wg.Add(1)
		go func(svc config.Service) {
			defer wg.Done()
			svcName := getSvcName(svc)
			gitUrl := "https://github.com/" + svc.Github
			repoName := strings.Split(svc.Github, "/")[1]

			if force {
				err := os.RemoveAll(repoName)
				if err != nil {
					log.Fatalf("Failed to remove existing directory '%s': '%s'", repoName, err)
				}
			}

			cmd := exec.Command("git", "clone", gitUrl)

			out, err := cmd.CombinedOutput()

			if err != nil {
				log.Warn("Failed to clone '" + gitUrl + "': " + err.Error())
			} else {
				log.Info("Synced '" + svcName + "' with remote repository " + svc.Github + "':\n" + string(out))
			}

			if svc.Sync != "" {
				CallServiceSync(&svc, repoName)
			}
		}(svc)
	}

	wg.Wait()
	log.Print("Finished syncing components against contract")
	return nil
}

func CallServiceSync(s *config.Service, dir string) error {
	if s.Sync != "" {
		customSyncParts := strings.Split(s.Sync, " ")
		cmd := exec.Command(customSyncParts[0], customSyncParts[1:]...)
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			log.Fatal("service!", "sync", err)
		} else {
			err = cmd.Wait()
			if err != nil {
				log.Fatal("service!", "sync", err)
			} else {
				log.Info("service!", "sync-return", err)
			}
		}
	}
	return nil
}
