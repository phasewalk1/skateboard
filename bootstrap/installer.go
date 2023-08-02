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
	"os"
	"path/filepath"

	"github.com/phasewalk1/skateboard/trucks"

	"github.com/charmbracelet/log"
)

func SkateboardPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err

	}

	return filepath.Join(home, ".skateboard"), nil
}

func SkateboardExists(home string) (bool, error) {
	_, err := os.Stat(home)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func TryMakeSkateboard(sk8path string, force bool, trucksToggle bool, noDeps bool) error {
	exis, err := SkateboardExists(sk8path)
	if err != nil {
		return err
	}
	if exis && !force {
		return nil
	} else if exis && force {
		log.Warn("removing existing skateboard installation")
		os.RemoveAll(sk8path)
	}
	return mkskateboard(sk8path, trucksToggle, noDeps)
}

func mkskateboard(home string, trucksToggle bool, noDeps bool) error {
	err := os.Mkdir(home, 0755)
	if err != nil {
		log.Fatal("couldn't create $HOME/.skateboard:", err)
		return err
	}
	if trucksToggle {
		err := trucks.BootstrapTrucks(home, noDeps)
		if err != nil {
			log.Fatal("couldn't bootstrap trucks:", err)
			return err
		}
	}
	log.Info("created home at: %s", home)
	return nil
}
