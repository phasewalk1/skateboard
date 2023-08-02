<div align="left">
    <h1 >
        skateboard 
        <img src="./docs/dist/assets/skateboard.svg" align="left"/>
    </h1>
</div>

<br/>

>> Warning! This is _alpha software_ still under rapid development. Check back later for a v1 release.

[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)

# Your application, on wheels
```fennel
(local trucks (require :trucks))

(do
  (trucks.system! :unwind true)
  (trucks.system! :verbose true)

  (trucks.defaults! {:run-ctx "npm"
    :cmd "run devstart"
    :sync "npm install"})

  (trucks.service! :fe {:github "mattg1243/sb-frontend"})
  (trucks.service! :courier {:github "phasewalk1/courier"
                     :run-ctx "cargo"
                     :cmd "watch -x run"})

  ;; (trucks.mallgrab! (trucks.config!))

  (trucks.contract!)
```

Working in a small team but want to onboard new developers without having them clone and run every service? *skateboard* provides an automated approach to bundling your system by using a single contract file, and provides bindings for generating bootstrap scripts to help you automate onboarding and focus on shipping.
## Features
- Pick your posion
    - Define your application and it's components in [Fennel](https://fennel-lang.org), [Lua](https://lua.org), [TOML](https://toml.io), or [YAML](https://yaml.org).
- Sync service repositories
- Run the system and all it’s services with goroutines in a single shell with unified logs 

## Installation

Install skateboard with go

```bash
 go install github.com/phasewalk1/skateboard@latest
```
