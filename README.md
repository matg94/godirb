# GODIRB

Godirb is a golang remake of dirb

## Usage

Configuration of godirb is mostly based off yaml files used as profiles,
in the `~/.godirb` directory you can have several `<profile>.yaml` files,
when using godirb you can use the `-p` flag to specify which to use.

Example:

`ls ~/.godirb`

```
default.yaml
deep.yaml
```

and when running the command, you can pick which file to load,

`godirb -p deep`

### Yaml Configuration

`default.yaml`
```
worker:
  limiter:
    enabled: False
    requests_per_second: 30000
  wordlists:
    - common.txt
  append_only: False
  append: 
    - .html
  ignore: 
    - 403
    - 401
  max_threads: 3
requests:
  cookie: abcdefg
  headers:
  - header: Authorization
    content: Bearer 123
  - header: Content-Type
    content: application/json
logging:
  stats: True
  debug_logger:
    file: /dev/null
    json_dump: False
    live: False
  success_logger:
    file: /dev/null
    json_dump: False
    live: True
  error_logger:
    file: /dev/null
    json_dump: False
    live: False
```

The default settings contain a basic configuration, meant to demonstrate the
godirb features available.

### Installation

run `go install github.com/matg94/godirb@latest` to install the CLI.

then copy & pase or download https://raw.githubusercontent.com/matg94/godirb/main/default.yaml,
and have a `~/.godirb/default.yaml` with its contents.

You should make a copy of the file, and rename it to your profile before making changes.
`cp ~/.godirb/default.yaml ~/.godirb/newprofile.yaml`.

Make the changes you want to the new profile file, which you can then run using

`godirb -p newprofile -url http://localhost`

### Flags

In general you should use profiles to set the desired config, as there is more control in the yaml.
However, you can use flags to quickly override some of the settings in the yaml.

In general, the profile will be loaded, and any flags will be added on top to edit the config when running.

`-url <url>` is required, the target url
`-p <profile-name>` is a profile name which needs a `<profile-name>.yaml` file in `~/.godirb`
`-conf <path-to-file>` is a path to a config yaml file, relative to pwd
`-limiter <int>` is the maximum requests per second allowed
`-threads <int>` number of threads to use, default is 10
`-cookie <cookie-string` string which will be passed as a cookie in requests
`-words <path-to-file>` path to a wordlist file that will be loaded
`-pipe` boolean flag that will output a json dump at the end to allow pipeing
`-stats` will display basic statistics about the run
`-out <path>` path to a json file to store the results
`-silent` silences all live outputs
``
