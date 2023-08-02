;; Copyright (C) 2023 Ethan Gallucci
;;
;;  This program is free software: you can redistribute it and/or modify
;;	it under the terms of the GNU General Public License as published by
;;	the Free Software Foundation, either version 3 of the License, or
;;	(at your option) any later version.
;;
;;	This program is distributed in the hope that it will be useful,
;;	but WITHOUT ANY WARRANTY; without even the implied warranty of
;;	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
;;	GNU General Public License for more details.

(lambda mkconfig! []
  (local config {})
  (tset config :system {})
  (tset config :defaults {})
  (tset config :services {})
  (tset _G :config config))

(lambda system! [system]
  (tset (. _G :config) :system system))

(lambda contract! []
  (. _G :config))

(lambda defaults! [defaults]
  (tset (. _G :config) :defaults defaults))

(lambda service! [name service]
  (tset service :name name)
  (table.insert (. _G :config :services) service))

(lambda services! [services]
  (each [_ v (ipairs services)]
    (print (.. "Adding service " (. v :name)))
    (table.insert (. _G :config :services) v)))

{:mkconfig! mkconfig!
 :system! system!
 :contract! contract!
 :defaults! defaults!
 :service! service!}
