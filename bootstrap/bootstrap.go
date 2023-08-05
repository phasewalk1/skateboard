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
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/phasewalk1/skateboard/config"
)

func BootstrapDir(path string, ymode bool, tmode bool, force bool) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if !force {
			log.Fatal("Error: Directory already exists:", path)
			return err
		}
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatal("Error creating directory:", err)
		return err
	}

	cmd := exec.Command("git", "init", path)
	if err := cmd.Run(); err != nil {
		log.Fatal("Error initializing Git repository:", err)
		return err
	}

	if !ymode && !tmode {
		err := bootstrapTrucksContract(path)
		if err != nil {
			log.Fatal("Failed to bootstrap trucks contract", err)
			return err
		}
	}

	err := bootstrapContract(path, ymode)
	if err != nil {
		log.Fatal("Failed to bootstrap contract file", err)
		return err
	}

	log.Info("New contract repository created at", path)
	return nil
}

func bootstrapContract(path string, ymode bool) error {
	if ymode == false {
		err := bootstrapToml(path)
		if err != nil {
			log.Fatal("Failed to boostrap contract (TOML)")
			return err
		}
	} else {
		err := boostrapYaml(path)
		if err != nil {
			log.Fatal("Failed to boostrap contract (YAML)")
		}
		log.Fatal("mode: YAML (unimplemented)")
		return nil
	}
	return nil
}

func bootstrapTrucksContract(path string) error {
	home, err := SkateboardPath()
	if err != nil {
		log.Fatal("skateboard", "home", err)
		return err
	}
	templPath := filepath.Join(home, "templates/trucks.contract.fnl")
	trucksBytes, err := ioutil.ReadFile(templPath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	file, err := os.Create(filepath.Join(path, "trucks.contract.fnl"))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(string(trucksBytes))
	if err != nil {
		return err
	}
	log.Info("example contract written!", "(trucks)", "trucks.contract.fnl")
	return nil
}

func boostrapYaml(path string) error {
	log.Info("bootstrapYaml called!")
	return nil
}

func bootstrapToml(path string) error {
	defaultServices := config.DefaultSvcMap()

	var buf strings.Builder
	for _, svc := range defaultServices {
		buf.WriteString("[[service]]\n")

		v := reflect.ValueOf(svc)
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			value := v.Field(i).Interface()

			// Skip fields with zero values
			if reflect.DeepEqual(value, reflect.Zero(field.Type).Interface()) {
				continue
			}

			// Write field to TOML
			tomlKey := strings.Split(field.Tag.Get("toml"), ",")[0] // Only take the first part of the tag
			if tomlKey != "-" {
				buf.WriteString(fmt.Sprintf("  %s = %q\n", tomlKey, value))
			}
		}

		buf.WriteString("\n")
	}

	tomlFilePath := filepath.Join(path, "skateboard.toml")
	tomlFile, err := os.Create(tomlFilePath)
	if err != nil {
		log.Fatal("Error creating contract file:", err)
		return err
	}
	defer tomlFile.Close()

	if _, err := tomlFile.WriteString(buf.String()); err != nil {
		log.Fatal("Error writing contract content:", err)
		return err
	}
	log.Info("example contract written!", "(toml)", "trucks.contract.toml")
	return nil
}
