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

  ;; (trucks.mallgrab! (trucks.contract!))

  (local contract (trucks.contract!))
  contract)
