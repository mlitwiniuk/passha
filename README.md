# passha

Executes command in parallel on set of hosts

## Installation

```
$ go get github.com/mlitwiniuk/passha
$ go install github.com/mlitwiniuk/passha
```

## Running

Once installed, prepare config file (by default called `config.yml`) present in same directory as command executed.

Here's sample file:

```yml
cmd: ps -ax | grep sidekiq
hosts:
  - user@host
```

and run with

```
$ passha
```
