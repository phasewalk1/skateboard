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

package bootstrap

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/phasewalk1/skateboard/config"
	"github.com/phasewalk1/skateboard/service"
)

func Upper(cfg *config.Config, ctx context.Context, cancel context.CancelFunc) error {
	defer cancel()
	errChan := make(chan error, len(cfg.Services))
	var wg sync.WaitGroup

	for _, svc := range cfg.Services {
		wg.Add(1)
		go func(svc config.Service) {
			defer wg.Done()

			fmt.Printf("Starting service: %s\n", svc.Name)
			cmd, err := service.StartSvc(ctx, svc)
			if err != nil {
				errChan <- err
				return
			}

			service.RunSvc(&wg, ctx, cfg, cmd, svc, errChan, cancel)
			fmt.Printf("Finished service: %s\n", svc.Name)
		}(svc)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil && cfg.PanicIsUnwind() {
			log.Println("Error in service:", err)
			cancel()
			for _, cmd := range service.Work {
				log.Println("Halting!:", cmd.Args)
				if cmd.Process != nil {
					cmd.Process.Kill()
				}
			}
			return nil
		}
	}
	return nil
}
