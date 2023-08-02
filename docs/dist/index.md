<div align="left">
    <h1 >
        skateboard 
        <img src="./assets/skateboard.svg" align="left"/>
    </h1>
</div>

>> Warning! This is _alpha software_ still under rapid development. Check back later for a v1 release.

[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)

Working in a small team but want to onboard new developers without having them clone and run every service? *skateboard* provides an automated approach to bundling your system by using a single contract file, and provides bindings for generating bootstrap scripts to help you automate onboarding and focus on shipping.
## Features
- Define all system configuration in `skateboard.toml` (or `yml` if you prefer), no manual docker coding necessary
- Sync service repositories
- Run the system and all it’s services with goroutines in a single shell with unified logs 

## Installation

Install skateboard with go

```bash
 go install github.com/phasewalk1/skateboard@latest
```
    
## Usage/Examples
## Setup
Before you can get started, we need to make the system portable. We do this by providing a single `skateboard.toml` (or `.yml`) file that tells _skateboard_ how to bundle your application and its components. The common practice for storing this _manifest_ would be to create a new repository so it can be checked into version control. Or if it makes sense for you to embed the contract into a specific component's repository, you can do that too. In the example below, we're bundling an app called `sb`, and creating a `sb-bootstrap` repository for maintaining the `skateboard` contract. Steps that should be taken for embedding a contract into an existing repository are denoted with 🛠️.

#### Scaffolding a `skateboard` contract
> 🚀 Creating a new Contract Workspace

Let's create a new github repository for sharing our applications `skateboard` contract with new developers. Assuming our application is called `sb`, we can run
```bash
skateboard new sb-bootstrap
```
This will create a working _git repository_ in `sb-bootstrap/` and an example `skateboard.toml` at the root. If you would like to use YAML instead of the default (TOML), pass the `-y` flag,
```bash
# this will scaffold `skateboard.yml` instead of `skateboard.toml`
skateboard new sb-bootstrap -y
```

> 🛠️ Embedding a Contract

If you want to embed skateboard into an existing repository using TOML, run the following command from the project's root     
```bash
skateboard init
```
or
```bash
skateboard init -y
```
for YAML. This will only scaffold an example `skateboard` contract.

#### An example `skateboard` contract
After scaffolding a contract using either `skateboard new` or `skateboard init`, you'll have an example contract in the projects root. If you generated a toml contract, it would look something like this:
```toml
[[service]]
  name = "user"
  github = "<repo-owner>/<repo-name>"
  run-ctx = "npm"
  cmd = "run devstart"

[[service]]
  name = "using-docker"
  github = "<repo-owner>/<repo-name>"
  profile = "<path-to-dockerfile>"
  env-bootstrap = "<path-to-dev-env-file>"
```

### Writing `skateboard` Contracts
`skateboard` calls its manifest file a _contract_. This tells `skateboard` how to bundle your application and serve it with `skateboard up`. Below is the go definition for a service.
```go
type Service struct {
    // The unique identifier for the service
	Name         string `toml:"name"`
    // The remote repository <owner>/<repo>
	Github       string `toml:"github"`
    // [Optional] The run context. Use this if not using a profile.
    // Some examples:
    //    - 'npm'
    //    - 'cargo'
    //    - 'go'
	RunContext   string `toml:"run-ctx"`
    // [Optional] Command to be called by the run context. Examples:
    //    - 'run devstart'
    //    - 'run --release'
    //    - 'run cmd/main.go'
	Cmd          string `toml:"cmd"`
    // [Optional] Path to the dockerfile
	Profile      string `toml:"profile,omitempty"`
    // [Optional] Path to a .env file for development environments
	EnvBootstrap string `toml:"env-bootstrap,omitempty"`
}
```

#### Service Modes
A service can be written to either use a run context and command, or to bundle and serve a container. These two different modes are called,

1. _Capsule mode_: 
  * Set the `profile` key to a valid path to a dockerfile (from the root of the services directory).
  * Optionally, set the `env-bootstrap` field to a path that points to a `.env` or equivalent file that can be passed to the dockerfile if required.
  * Don't set the `run-ctx` or `cmd` fields.
2. _Custom mode_.
  * Use the `run-ctx` field to specify a run context.
    * This can be any binary that can run the service
  * Set the `cmd` field to a command that should run to serve the service
    * An example with `run-ctx = "npm"`,
      * `cmd = "run devstart"`

Picking up where we left off with our example contract, let's add our application's service to it. Our application will use three services, a frontend, a user service, and a messaging service. Two of our services are written in Typescript, so we'll use `npm` as the run context for those, and one is written in Rust; we'll use _capsule mode_ for that one by passing a profile and path to a file with necessary environment variables (if the container requires one).

```toml
[[service]]
  name = "user"
  github = "mattg1243/sb-user-service"
  run-ctx = "npm"
  cmd = "run devstart"

[[service]]
  name = "frontend"
  github = "mattg1243/sb-frontend"
  run-ctx = "npm"
  cmd = "start"

[[service]]
  name = "courier"
  github = "phasewalk1/courier"
  profile = "docker/Dockerfile"
  env-bootstrap = "docker/environment/.env.dev"
```

> 🚀
>
> If you're following along the example and or are using a unique repository for storing your skateboard contracts make sure to add the service repository paths to your `.gitignore` so they are not embedded into your contract repository and instead, synced down everytime.

### Syncing From a Contract
Congratulations! Your entire application is ready to go now. Let's test it out.

#### Using `skateboard sync` to Bootstrap Repositories
We can pull down all the service repositories by using the `sync` command.

```bash
skateboard sync
2023/07/31 20:29:18 Synced 'user' with remote repository mattg1243/sb-user-service':
Cloning into 'sb-user-service'...

2023/07/31 20:29:18 Synced 'frontend' with remote repository mattg1243/sb-frontend':
Cloning into 'sb-frontend'...

2023/07/31 20:29:20 Synced 'courier' with remote repository phasewalk1/courier':
Cloning into 'courier'...

2023/07/31 20:29:20 Finished syncing components against contract
```
#### `sync` forcefully
If the directories already exist, or you'd like to sync _forcefully_ pass the `-f` or `--force` flag to `sync`,
```bash
skateboard sync -f
```
> ⚠️ This will remove the service repositories if they exist at runtime and pull down a new copy from the remote.

### Running Your Application With `skateboard up`
Now that we've synced the contract and have working copies of our service sources, we can run them all with `skateboard up`. New users who've just obtained a copy of a working contract can navigate to the directory in which the contract lives and run,

> ⚠️ Make sure to populate any necessary `.env` files in the service directories if required before running the next command.

```bash
skateboard up
```
to start the services.

If you're following from the example and have already synced the repositories, you can pass `--no-sync` or `-n` to skip syncing sources,
```bash
skateboard up --no-sync
```
