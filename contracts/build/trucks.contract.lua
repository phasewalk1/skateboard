local trucks = require("trucks")
trucks["mkconfig!"]()
trucks["system!"]({panic = "unwind", verbose = "true"})
trucks["defaults!"]({["run-ctx"] = "npm", cmd = "run devstart", sync = "npm install"})
trucks["service!"]("fe", {github = "mattg1243/sb-frontend"})
trucks["service!"]("courier", {github = "phasewalk1/courier", ["run-ctx"] = "cargo", cmd = "watch-x run"})
local contract = trucks["contract!"]()
return contract
