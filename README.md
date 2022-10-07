# GODIRB

Godirb is a golang remake of dirb

## Usage

Configuration of godirb is mostly based off yaml files used as profiles,
in the `~/.godirb` directory you can have several `<profile>.yaml` files,
when using godirb you can use the `-p` flag to specify which to use.

Profiles allow you to save specific config, and avoid having to repeat several flags each run.

Example:

`ls ~/.godirb`

```
default.yaml
deep.yaml
```

and when running the command, you can pick which file to load,

`godirb -p deep`

### Installation

## Using the install script

You can install `godirb` using this command:

`curl -s https://raw.githubusercontent.com/matg94/godirb/main/install.sh | sh`

which will run [this script](https://raw.githubusercontent.com/matg94/godirb/main/install.sh) and download the binaries for you.
This will also download and setup the default config in `~/.godirb`.

You will need to move the binaries to path after that to use it in any directory, for example,
`sudo mv ./godirb /usr/local/bin/godirb`

## Using go

run `go install github.com/matg94/godirb@latest` to install the CLI.

Optionally copy & pase or download [the default config](https://raw.githubusercontent.com/matg94/godirb/main/default.yaml),
You can copy paste from [here](#yaml-configuration) as well,
and create a file: `~/.godirb/default.yaml` with its contents.

This sets reasonable defaults that you can get started overriding with the CLI flags.

You should make a copy of the file, and rename it to your profile before making changes.
`cp ~/.godirb/default.yaml ~/.godirb/newprofile.yaml`, so that you can always fall back to those defaults.

Make the changes you want to the new profile file, which you can then run using

`godirb -p newprofile -url http://localhost`

### Yaml Configuration

`default.yaml`
```
worker:
  wordlists:
    - common.txt
  ignore: 
    - 403
    - 401
  max_threads: 5
logging:
  stats: True
  success_logger:
    file: /dev/null
    json_dump: False
    live: True
```

The default settings contain a basic configuration, meant for quick use with a simple config to get started.

It runs in standard output mode, displaying only successes and showing stats, it ignores 403s and 401s, and uses the [common.txt wordlist]().

For example,

Command:
`godirb -url http://localhost:8080 -p default`

Output:
```
-------------------------------
Words Generated:  4614
-------------------------------
200 | /
200 | /config
200 | /docs
200 | /external
200 | /favicon.ico
200 | /index.php
200 | /php.ini
200 | /phpinfo.php
200 | /robots.txt
-------------------------------
Code  | Count
----- | -----
200   | 9
404   | 4601
403   | 4
-------------------------------
Time taken: 0.483981237 seconds
Total Hits: 4614
Final Rate: 9522 requests per second
-------------------------------
```

You can override profile config by specifying flags, for example:

Command:
`godirb -url http://localhost:8080 -limiter 2000 -silent`

Output:
```
-------------------------------
Words Generated:  4614
-------------------------------
-------------------------------
Code  | Count
----- | -----
200   | 9
404   | 4601
403   | 4
-------------------------------
Time taken: 2.306208387 seconds
Total Hits: 4614
Final Rate: 2000 requests per second
-------------------------------
```

### Flags

In general you should use profiles to set the desired config, as there is more control in the yaml.
However, you can use flags to quickly override some of the settings in the yaml.

In general, the profile will be loaded, and any flags will be added on top to edit the config when running.


`-url <url>` is required, the target url

`-p <profile-name>` is a profile name which needs a `<profile-name>.yaml` file in `~/.godirb`

`-conf <path-to-file>` is a path to a config yaml file, relative to current directory

`-limiter <int>` is the maximum requests per second allowed

`-threads <int>` number of threads to use, default is 10

`-cookie <cookie-string` string which will be passed as a cookie in requests

`-words <path-to-file>` path to a wordlist file that will be loaded

`-pipe` boolean flag that will output a json dump at the end to allow pipeing

`-stats` will display basic statistics about the run

`-out <path>` path to a json file to store the results

`-silent` silences all live outputs
``
