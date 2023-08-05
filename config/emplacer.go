package config

import "github.com/charmbracelet/log"

func (c *Config) EmplaceDefaults() {
	log.Info("contract", "emplacing defaults...", true)
	for k, v := range c.Defaults {
		if k == "run-ctx" {
			log.Debug("contract", "[defaults] emplacing default -run-ctx with:", v)
			for i := range c.Services {
				if c.Services[i].RunContext == "" {
					c.Services[i].RunContext = v
					log.Info("contract", "[update] updated -run-ctx for:", c.Services[i].Github)
				} else {
					log.Warn(
						"contract",
						"[defaults] skipped emplacing default 'run-ctx', overriden by:",
						c.Services[i].RunContext,
					)
				}
			}
			log.Info("contract", "[update] updated -run-ctx by default emplacer with value:", v)
		} else if k == "cmd" {
			log.Debug("contract", "[defaults] emplacing default -cmd with:", v)
			for i := range c.Services {
				if c.Services[i].Cmd == "" {
					c.Services[i].Cmd = v
					log.Info("contract", "[update] updated -cmd for:", c.Services[i].Github)
				} else {
					log.Warn("contract", "[defaults] skipped emplacing default 'cmd', overriden by:", c.Services[i].Cmd)
				}
			}
		} else if k == "sync" {
			log.Debug("contract", "[defaults] emplacing default -sync with:", v)
			for i := range c.Services {
				if c.Services[i].Sync == "" && c.Services[i].RunContext != "cargo" {
					c.Services[i].Sync = v
					log.Info("contract", "[update] updated -sync for:", c.Services[i].Github)
				} else {
					log.Warn(
						"contract",
						"[defaults] skipped emplacing defaults 'sync', overriden by:",
						c.Services[i].Github,
					)
				}
			}
		}
	}
	log.Debug("contract", "[finished] finished emplacing defaults in config:", c.Services)
}
