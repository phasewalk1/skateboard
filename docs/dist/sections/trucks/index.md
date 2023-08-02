# trucks

## Writing a trucks contract

### Loading the trucks module

```scheme
(local trucks (require :trucks))
```

### Defining a contract
#### Setup trucks
```scheme
;; setup an empty configuration table
(do
  (trucks.mkconfig!)
```
> `do` block opens here...

#### Set system wide configuration
```scheme
;; set system-wide configs
(trucks.system! {:panic "unwind"
  :verbose "true"})
```

#### Set service level defaults
```scheme
;; set defaults for all services
(trucks.defaults! {:run-ctx "npm"
  :cmd "run devstart"
  :sync "npm install"})
```

#### Define services
```scheme
;; add services
(trucks.service! :fe {:github "mattg1243/sb-frontend"})
(trucks.service! :courier {:github "phasewalk1/courier"
                   :run-ctx "cargo"
                   :cmd "watch-x run"
                   :no-sync "true"})
```

#### Export the contract
```scheme
;; optionally - log the contract before exporting it
;; (trucks.mallgrab! (trucks.contract!))

;; export the contract
(local contract (trucks.contract!))
contract)
```
> `do` block closes here...
