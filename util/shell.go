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
	"os/exec"

	"github.com/charmbracelet/log"
)

func ExecWithFatal(cmd *exec.Cmd, scope string, msg string) {
	out, err := cmd.CombinedOutput()
	log.Debug(scope, "executing:", cmd.String())
	if err != nil {
		log.Warn(string(out))
		log.Fatal(msg, err)
		panic(err)
	}
	log.Debug(string(out))
}
