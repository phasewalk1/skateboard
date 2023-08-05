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
	"sync"

	"github.com/charmbracelet/log"
	"github.com/phasewalk1/skateboard/config"
	"github.com/phasewalk1/skateboard/service"
)

func Upper(cfg *config.Config, noSync bool, ctx context.Context, cancel context.CancelFunc) error {
	defer cancel()
	log.Debug("service!", "defer:", "cancel()")
	errChan := make(chan error, len(cfg.Services))
	log.Debug("service!", "channels", "made error chan")
	var wg sync.WaitGroup

	for _, svc := range cfg.Services {
		log.Debug("service!", "service:", svc)
		wg.Add(1)
		log.Debug("service!", "waitgroup:", "adding 1...")
		go func(svc config.Service) {
			log.Debug("service!", "defer:", "wg.Done()")

			log.Info("service!", "starting", true)
			cmd, err := service.StartSvc(ctx, svc, noSync)
			if err != nil {
				errChan <- err
				wg.Done()
				return
			}

			service.RunSvc(&wg, ctx, cfg, cmd, svc, errChan, cancel)
			log.Info("service!", "running", true)
		}(svc)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil && cfg.PanicIsUnwind() {
			log.Warn("service!", "error in service:", err)
			cancel()
			for _, cmd := range service.Work {
				log.Warn("service!", "halting!:", cmd.Args)
				if cmd.Process != nil {
					cmd.Process.Kill()
				}
			}
			return nil
		}
	}
	return nil
}
