local trucks = require("trucks")
do
  trucks["mkconfig!"]()
end
trucks["defaults!"]({["run-ctx"] = "npm", cmd = "run devstart", sync = "npm install"})
local services = {{name = "user", github = "mattg1243/sb-user-service"}, {name = "courier", github = "phasewalk1/courier"}}
trucks["service!"](services)
return trucks["getconfig!"]()
