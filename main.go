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

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/phasewalk1/skateboard/cmd"
	"github.com/phasewalk1/skateboard/service"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		for _, cmd := range service.Work {
			if cmd.Process != nil {
				// cmd.Process.Kill()
				syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM)
			}
		}
		os.Exit(1)
	}()
	cmd.Execute()
}
