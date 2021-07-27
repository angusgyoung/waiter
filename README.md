# Waiter

[![Tests](https://github.com/angusgyoung/waiter/actions/workflows/test.yml/badge.svg)](https://github.com/angusgyoung/waiter/actions/workflows/test.yml)

[![Container build](https://github.com/angusgyoung/waiter/actions/workflows/build-container.yml/badge.svg)](https://github.com/angusgyoung/waiter/actions/workflows/build-container.yml)

[![CodeQL](https://github.com/angusgyoung/waiter/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/angusgyoung/waiter/actions/workflows/codeql-analysis.yml)

Waiter is a cli for automating tasks, with support for retries, batching and templating.

## Features

## Usage

```
$ waiter help

Waiter is a task runner, that exeutes simple or complex tasks, with support for batching, retries and schedules.

Usage:
  waiter [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  task        Manage tasks

Flags:
      --config string   config file (default is $HOME/.waiter.yaml)
  -h, --help            help for waiter
      --level string    log level

Use "waiter [command] --help" for more information about a command.
```