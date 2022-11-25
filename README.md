# `hashtool`
### Hashing tool

A simple tool for hasing strings and files.

## Build and Install
1. [Install](https://go.dev/doc/install) golang
> ```bash
> $ go version
> go version go1.17 xxx
> ```

2. [Clone](https://github.com/pangduckwai/hashtool) the repository from GitHub
> ```bash
> $ git clone https://github.com/pangduckwai/hashtool.git
> ```

3. Build the executable `hashtool`
> ```bash
> $ cd .../hashtool
> $ go build
> $ sudo ln -s hashtool /usr/local/bin/hashtool # for example
> ```

## Usage
```
Usage:
 hashtool [version | help]
   {-a ALGR | --algorithm=ALGR}
   {-i FILE | --in=FILE}
   {-b SIZE | --buffer=SIZE}
   {-v | --verbose}
```

- Commands
  - omitting means to do hashing
  - `version`
    - display the current version
  - `help`
    - display the help message

- Options
  - `-a algorithm` | `--algorithm=algorithm`
    - hashing algorithm to use, support `md5`, `sha1` and **_`sha256`_**
  - `-i filename` | `--in=filename`
    - name of the input file, omitting means input from `stdin`
  - `-b size` | `--buffer=size`
    - the buffer size used to read large inputs (**_`1024KB`_**)
  - `-v | --verbose`
    - display detail operation messages during processing if specified

## Changelog
### v0.1.0
- first usable version
