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

(local mod {})

(local util (require :util))
(local config (require :config))

(tset mod :mkconfig! (. config :mkconfig!))
(tset mod :system! (. config :system!))
(tset mod :contract! (. config :contract!))
(tset mod :defaults! (. config :defaults!))
(tset mod :service! (. config :service!))
(tset mod :mallgrab! (. util :mallgrab!))

mod
