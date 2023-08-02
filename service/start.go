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

package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/phasewalk1/skateboard/config"
)

var Work []*exec.Cmd
var WorkMutex sync.Mutex

func StartSvc(
	ctx context.Context,
	svc config.Service,
) (*exec.Cmd, error) {
	repoName := strings.Split(svc.Github, "/")[1]
	fmt.Println("repoName:", repoName)

	if svc.RunContext != "" {
		cmdArgs := append([]string{svc.RunContext}, strings.Split(svc.Cmd, " ")...)
		fmt.Printf("Got cmdArgs: %v\n", cmdArgs)

		cmd := exec.CommandContext(ctx, cmdArgs[0], cmdArgs[1:]...)
		cmd.Dir = repoName
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Add the command to the cmds slice
		WorkMutex.Lock()
		Work = append(Work, cmd)
		WorkMutex.Unlock()

		fmt.Printf("Starting: %v in directory: %v\n", cmd.Args, cmd.Dir)
		err := cmd.Start()
		if err != nil {
			fmt.Printf("Failed to start service: %v\n", err)
			WorkMutex.Lock()
			defer WorkMutex.Unlock()
			for i, c := range Work {
				if c == cmd {
					Work = append(Work[:i], Work[i+1:]...)
					break
				}
			}
			return nil, fmt.Errorf("Failed to start service: %v", err)
		}
		return cmd, nil
	}

	fmt.Println("svc.RunContext is empty")
	return nil, nil
}
