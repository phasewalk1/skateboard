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

(lambda print-service! [service]
  (do
    (each [k v (pairs service)]
      (print (.. "k: " k " v: " v)))))

(lambda print-services! [services]
  (do
  (print "print-services!:")
    (each [k v (pairs services)]
      (print-service! v))))

(lambda print-defaults! [defaults]
  (do
    (print "print-defaults!")
    (each [k v (pairs defaults)]
      (print (.. "defaults_k: " k))
      (print (.. "defaults_v: " v)))))

(lambda mallgrab! [trucks]
  (do
    (print "mallgrab!")
    (print-services! (. trucks :services))
    (print-defaults! (. trucks :defaults))))

{:print-services! print-services!
 :print-defaults! print-defaults!
 :mallgrab! mallgrab!}
