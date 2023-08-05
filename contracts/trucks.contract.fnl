(local trucks (require :trucks))

(do
  (trucks.mkconfig!)

  (trucks.system! {:panic "unwind"
    :verbose "true"})

  (trucks.defaults! {:run-ctx "npm"
    :cmd "run devstart"
    :sync "npm install"})

  (trucks.service! :user {:github "mattg1243/sb-user-service"})
  (trucks.service! :fe {:github "mattg1243/sb-frontend"
                     :cmd "start"})
  (trucks.service! :courier {:github "phasewalk1/courier"
                     :run-ctx "cargo"
                     :cmd "watch -x run"})

  ;; (trucks.mallgrab! (trucks.contract!))

  (local contract (trucks.contract!))
  contract)
