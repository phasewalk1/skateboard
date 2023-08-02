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

package config

import (
	"io/ioutil"

    "github.com/charmbracelet/log"
	"github.com/BurntSushi/toml"
)

type System struct {
	Panic string `toml:"panic"`
}

type Config struct {
	System   System            `toml:"system"`
	Defaults map[string]string `toml:"defaults"`
	Services []Service         `toml:"service"`
}

func (c *Config) Show() {
    log.Debug("trucks.system!")
	c.ShowSystem()
	c.ShowDefaults()
    log.Debug("trucks.service!")
	c.ShowServices()
}

func (c *Config) ShowSystem() {
	log.Debug("trucks.system!", "panic:", c.System.Panic)
}

func (c *Config) ShowDefaults() {
	for setting, val := range c.Defaults {
        log.Debug("trucks.defaults!", setting, val)
	}
}

func (c *Config) ShowServices() {
	for svc := range c.Services {
        log.Debug("trucks.service!", c.Services[svc])
	}
}

type ServiceMap struct {
	Services map[string]map[string]string `json:"services"`
	Defaults map[string]string            `json:"defaults"`
}

func (c *Config) PanicIsUnwind() bool {
	return c.System.Panic == "unwind" || c.System.Panic == ""
}

func (c *Config) PanicIsKeep() bool {
	return c.System.Panic == "keep"
}

type Service struct {
	Name         string `toml:"name"`
	Github       string `toml:"github"`
	RunContext   string `toml:"run-ctx"`
	Cmd          string `toml:"cmd"`
	Profile      string `toml:"profile,omitempty"`
	EnvBootstrap string `toml:"env-bootstrap,omitempty"`
	Sync         string `toml:"sync"`
}

type Services struct {
	ServiceList []Service `toml:"service"`
}

func SvcToMap(s *Service) map[string]interface{} {
	return map[string]interface{}{
		"github":        s.Github,
		"run-ctx":       s.RunContext,
		"cmd":           s.Cmd,
		"profile":       s.Profile,
		"env-bootstrap": s.EnvBootstrap,
	}
}

func DefaultSvcMap() []Service {
	services := []Service{
		{
			Name:       "user",
			Github:     "<repo-owner>/<repo-name>",
			RunContext: "npm",
			Cmd:        "run devstart",
			Sync:       "npm install",
		},
		{
			Name:         "using-docker",
			Github:       "<repo-owner>/<repo-name>",
			Profile:      "<path-to-dockerfile>",
			EnvBootstrap: "<path-to-dev-env-file>",
		},
	}
	return services
}

func LoadWorkloadContractFromFile(filePath string) (*Config, error) {
	var config Config

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if _, err := toml.Decode(string(data), &config); err != nil {
		return nil, err
	}

	return &config, nil
}
