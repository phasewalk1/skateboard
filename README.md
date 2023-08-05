# skateboard

>> Warning! This is _alpha software_ still under rapid development.

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-critical.svg)](https://www.gnu.org/licenses/gpl-3.0) <img src="https://img.shields.io/badge/Fennel-1.3.1-green.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAQAAAC1+jfqAAABPklEQVQoz02RO0tDQRCF95rExNII4gNsBGsbHxj9A/4KO61sBBs7f4JE8QWiCLZa2SiIjdjcwk4UDBcUEVSCDzSafJ7Z+8BddtmdOXPmzIxzjoCc7hH2iPjWjvQalSVH4Lw7oMA6tlrZ3aJKERdDCpzxxidNOV6p6/7VhhOKzhbbXDHjY+/oYZAnvZpKBVVzV3inl0UPCPVv58YDLM0PFcchmzLPy2Axl4SZFkuz66gJ5ejnUd9GIpBMbM0AXV7JJA/e2PL06Wo4CStbzTqdLEh5/R9HAhiTs007sKpVR5iJ9CmOWPG9sDRlcQwwnEQnIqd4YUjOks6yTBd0cO854jJl3uKWkmeY5pkluvkSwM5q3Mm8oq7p80pMxUbSk1NjjYeVZ0c9OGCWOc6TYteMNZ2nFTnBMR/eGbHPeDruPzYIcSxR7bvBAAAAAElFTkSuQmCC"/> <img src="https://img.shields.io/badge/Lua-5.4.6-purple.svg?logo=lua"/> ![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/phasewalk1/skateboard)


# Your application, on wheels
```fennel
(local trucks (require :trucks))

(do
  (trucks.mkconfig!)

  (trucks.system! {:panic "unwind"
    :verbose "true"})

  ;; (trucks.defaults! {:run-ctx "npm"
  ;;  :cmd "run devstart"
  ;;  :sync "npm install"})

  (trucks.service! :docs {:github "phasewalk1/phasewalk1.github.io"
                     :run-ctx "hugo"
                     :cmd "serve -D"
                     :sync "git submodule update --init --recursive"})
  (trucks.service! :courier {:github "phasewalk1/courier"
                     :run-ctx "cargo"
                     :cmd "watch -x run"})

  (local contract (trucks.contract!))
  contract)
```

# skateboard is not virtualization software
For smaller teams who are deploying applications with only a handful of services, _skateboard_ gives them wheels to share their system without relying on virtualization. This brings a few core advantages over what I would call _container overkill_.
1. _skateboard can sync and launch individual services much faster than containerized composers._
2. _skateboard has a much smaller footprint; **no containers means no impactful storage footprint**._
3. _no background daemon and no os-level virtualization means less RAM usage._
4. _trucks allows developers to be more expressive when defining their system than yaml._

Of course -- because _skateboard_ is **not virtualization software** -- this means     
1. _It doesn't provide the isolation between processes and system resources like CPU, memory, and I/O offered by full-fledged containerization solutions._
2. _It doesn't provide the same level of system security as software like Docker, which leverages namespaces and cgroups to isolate resources._
3. _It's not designed for deployment across multiple machines; instead, it excels in managing multi-service applications on a single host._
4. _It doesn't have an inbuilt dependency management system. Instead, it assumes that your services are managed by an appropriate language-specific package manager, such as [npm](https://npmjs.com),  [Cargo](https://doc.rust-lang.org/cargo/), [pip](https://pypi.org/project/pip/), etc. Skateboard expects these tools to handle the heavy lifting of dependency management. In a trucks contract, the `sync` field specifies the command skateboard will run to ensure all dependencies are met, like npm install or cargo build, before it tries to run your services_

In other words, skateboard is designed for simplicity and speed in situations where the overhead and complexity of virtualization are unnecessary. By focusing on sharing and running multi-service applications on a single host, it streamlines the development and deployment process for small teams building simpler applications.

## Features
- [trucks](https://github.com/phasewalk1/skateboard/blob/master/contracts/trucks.contract.fnl)
    - Give your application wheels by writing a trucks contract that defines your system and its components. Contracts are written in Lisp syntax using skateboard's [trucks library](https://github.com/phasewalk1/skateboard/blob/master/trucks/trucks.fnl) in [Fennel](https://fennel-lang.org)
    - _skateboard_ embeds necessary Trucks / Fennel dependencies so you can get started right away
      + `skateboard install`
- Sync service repositories
- Run the system and all it’s services with goroutines in a single shell with unified logs
  - A user in possession of a valid _trucks contract_ can spin up the application in a single command
    + `skateboard up`

# Installation
>> ⚠️
>> _skateboard assumes you have a Lua runtime installed_. If you don't already or you fear your runtime may be out of date, see https://lua.org/download.html.

## Install the skateboard binary with go

```bash
 go install github.com/phasewalk1/skateboard@latest
```

Try to run `skateboard -h` to ensure that `$GOPATH/bin` is in path, usually at `~/go/bin/` or `/usr/local/go/bin`.

## Install trucks
This will create a home for skateboard at `$HOME/.skateboard`, and will sync and build any dependencies needed to work with trucks. 
```bash
 skateboard install
```
# Usage / Examples
## Create a new contract
You can scaffold a new contract by running
```bash
skateboard new my-contract
```
This will create a directory at `my-contract/` and initialize it as a git repository; it also scaffolds an example contract for you at `trucks.contract.fnl` that looks exactly like the example contract above.

## Share your wheels
Once you've defined your application in a trucks contract, you can share the contract with anyone who has _skateboard installed_. Once they have your contract, they can run your application on wheels by navigating to the directory the contract is in and running
```bash
skateboard up -n
```
>> _The `-n` (or `--new-clone`) flag only needs to be passed the first time you run `up` on a contract. This tells skateboard to clone new copies of the services before attempting to run them; But it can also be used if you have existing copies and want to start with a clean slate, pass `--force -n` to force clone new copies. You can also run `up -x` (or `up --no-sync`) to skip running any _sync_ operations defined on the services, i.e., to skip running `npm install` if you already have the necessary node modules from a previous `up` invocation._
