<div align="left">
    <h1 >
        skateboard
        <img src="./assets/skateboard.svg" align="left"/>
    </h1>
</div>

>> Warning! This is _alpha software_ still under rapid development. Check back later for a v1 release.

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-critical.svg)](https://www.gnu.org/licenses/gpl-3.0) <img src="https://img.shields.io/badge/Fennel-1.3.1-green.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAQAAAC1+jfqAAABPklEQVQoz02RO0tDQRCF95rExNII4gNsBGsbHxj9A/4KO61sBBs7f4JE8QWiCLZa2SiIjdjcwk4UDBcUEVSCDzSafJ7Z+8BddtmdOXPmzIxzjoCc7hH2iPjWjvQalSVH4Lw7oMA6tlrZ3aJKERdDCpzxxidNOV6p6/7VhhOKzhbbXDHjY+/oYZAnvZpKBVVzV3inl0UPCPVv58YDLM0PFcchmzLPy2Axl4SZFkuz66gJ5ejnUd9GIpBMbM0AXV7JJA/e2PL06Wo4CStbzTqdLEh5/R9HAhiTs007sKpVR5iJ9CmOWPG9sDRlcQwwnEQnIqd4YUjOks6yTBd0cO854jJl3uKWkmeY5pkluvkSwM5q3Mm8oq7p80pMxUbSk1NjjYeVZ0c9OGCWOc6TYteMNZ2nFTnBMR/eGbHPeDruPzYIcSxR7bvBAAAAAElFTkSuQmCC"/> <img src="https://img.shields.io/badge/Lua-5.4.6-purple.svg?logo=lua"/> ![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/phasewalk1/skateboard)


# Your application, on wheels
```scheme
(local trucks (require :trucks))

(do
  (trucks.mkconfig!)

  (trucks.system! {:panic "unwind"
    :verbose "true"})

  (trucks.defaults! {:run-ctx "npm"
    :cmd "run devstart"
    :sync "npm install"})

  (trucks.service! :fe {:github "mattg1243/sb-frontend"})
  (trucks.service! :courier {:github "phasewalk1/courier"
                     :run-ctx "cargo"
                     :cmd "watch-x run"})

  ;; (trucks.mallgrab! (trucks.contract!))

  (local contract (trucks.contract!))
  contract)
```

Working in a small team but want to onboard new developers without having them clone and run every service? *skateboard* provides an automated approach to bundling your system by using a single contract file, and provides bindings for generating bootstrap scripts to help you automate onboarding and focus on shipping.
## Features
- [trucks](https://github.com/phasewalk1/skateboard/blob/master/contracts/trucks.contract.fnl)
    - Give your application wheels by writing a trucks contract that defines your system and its components. Contracts are written in Lisp syntax using skateboard's [trucks library](https://github.com/phasewalk1/skateboard/blob/master/trucks/trucks.fnl) in [Fennel](https://fennel-lang.org)
    - _skateboard_ embeds any necessary dependencies so you can get started right away with `skateboard install`
- Sync service repositories
- Run the system and all it’s services with goroutines in a single shell with unified logs
  - A user in possession of a valid _trucks contract_ can spin up the application in a single command, `skateboard up`

## Installation

### Install the skateboard binary with go

```bash
 go install github.com/phasewalk1/skateboard@latest
```

Try to run `skateboard -h` to ensure that `$GOPATH/bin` is in path, usually at `~/go/bin/` or `/usr/local/go/bin`.

### Install trucks
This will create a home for skateboard at `$HOME/.skateboard`, and will sync and build any dependencies needed to work with trucks. 
```bash
 skateboard install

...
2023/08/02 14:34:01 DEBU copying trucks into $HOME/.skateboard
2023/08/02 14:34:01 DEBU pwd: /home/kat/.skateboard/skateboard="missing value"
2023/08/02 14:34:01 DEBU trucks.install.trucks executing:=/usr/bin/make
2023/08/02 14:34:01 DEBU mkdir -p include/
fennel --compile trucks/config.fnl > include/config.lua
mkdir -p include/
fennel --compile trucks/util.fnl > include/util.lua
mkdir -p include/
fennel --compile trucks/trucks.fnl > include/trucks.lua

2023/08/02 14:34:01 DEBU trucks.install.trucks executing:="/usr/bin/sh -c cp -r include/*.lua ../include"
2023/08/02 14:34:01 DEBU
2023/08/02 14:34:01 DEBU trucks.install.trucks executing:="/usr/bin/cp -r trucks .."
2023/08/02 14:34:01 DEBU
2023/08/02 14:34:01 DEBU trucks.install.trucks executing:="/usr/bin/rm ../trucks/installer.go"
2023/08/02 14:34:01 DEBU
2023/08/02 14:34:01 DEBU removing skateboard directory
2023/08/02 14:34:01 DEBU trucks.install.trucks executing:="/usr/bin/rm -rf skateboard"
2023/08/02 14:34:01 DEBU
2023/08/02 14:34:01 INFO created home at: %s /home/kat/.skateboard="missing value"

```

## Usage / Examples
### Create a new contract
You can scaffold a new contract by running
```bash
skateboard new my-contract
```
This will create a directory at `my-contract/` and initialize it as a git repository; it also scaffolds an example contract for you at `trucks.contract.fnl` that looks exactly like the example contract above.

### Share your wheels
Once you've defined your application in a trucks contract, you can share the contract with anyone who has _skateboard installed_. Once they have your contract, they can run your application on wheels by navigating to the directory the contract is in and running
```bash
skateboard up -n
```
>> _The `-n` flag only needs to be passed the first time you run `up` on a contract. This tells skateboard to clone new copies of the services before attempting to run them; But it can also be used if you have existing copies and want to start with a clean slate, pass `--force -n` to force clone new copies._
