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
	"os/exec"
	"sync"

	"github.com/phasewalk1/skateboard/config"
)

func RunSvc(
	wg *sync.WaitGroup,
	ctx context.Context,
	cfg *config.Config,
	cmd *exec.Cmd,
	svc config.Service,
	errChan chan<- error,
	cancel context.CancelFunc,
) {
	defer wg.Done()

	if svc.RunContext != "" {
		// Watch for cancellation
		go func() {
			select {
			case <-ctx.Done():
				if cmd.Process != nil {
					fmt.Print("Error!:", cmd.Args)
					cmd.Process.Kill()
				}
			}
		}()

		err := cmd.Wait()
		if err != nil {
			fmt.Printf("Service %s exited with error: %v\n", svc.Name, err)
			errChan <- fmt.Errorf("Service exited with error: %v", err)
			cancel()
		}
	}
}
