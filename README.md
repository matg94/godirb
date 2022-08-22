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