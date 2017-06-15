# Formula Go
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Build Status](https://travis-ci.org/jkawamoto/fgo.svg?branch=master)](https://travis-ci.org/jkawamoto/fgo)
[![Code Climate](https://codeclimate.com/github/jkawamoto/fgo/badges/gpa.svg)](https://codeclimate.com/github/jkawamoto/fgo)
[![Release](https://img.shields.io/badge/release-0.2.4-brightgreen.svg)](https://github.com/jkawamoto/fgo/releases/tag/v0.2.4)

Formula Go helps you to build and upload your software written in
[Go](https://golang.org/), and prepare a [homebrew](http://brew.sh/) formula
for it.

Formula Go assumes your software is hosted in [GitHub](https://github.com/),
and the pre-compiled binaries are uploaded in the release page of it.

## Usage
~~~
fgo [global options] command [arguments...]

COMMANDS:
     init     create Makefile and other related directories.
     build    build binaries, upload them, and update brew formula.
     update   update only brew formula.
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --pkg NAME, -p NAME   overwrite directory NAME to store package files. (Default: pkg)
   --brew NAME, -b NAME  overwrite directory NAME to store homebrew formula. (Default: homebrew)
   --help, -h            show help
   --version, -v         print the version
~~~

### Initialization
After starting your project, initialization is required.
To initialize Formula Go, run `fgo init`.
It creates a Makefile, which will be used to compile your project,
and a template of homebrew formula. If a Makefile or a template of homebrew
formula already exist, fgo won't overwrite them.

To create the template of homebrew formula, a user name and a repository name
in GitHub is required. By default, fgo checks your git configuration to get
those information but you can given them by the arguments.

If your git configuration doesn't have both information and you don't give them
as the arguments, this command will skip to create the template. In this case,
you need to re-run init command after setting git configuration.

You can edit the Makefile and the template of homebrew formula, but build and
release targets are necessary to run build command.


### Build and upload
`fgo build [version]` runs build and release targets in the Makefile to build
your software and upload the binary files to GitHub. This command takes an
argument, version, which specifies the version to be created. If it is omitted,
"snapshot" will be used and uploading will be skipped.

This command also updates the homebrew formula. After finishing this command,
you need to push the updated formula.


### Update
`fgo update [version]` updates the homebrew formula for a given version.
build command updates the homebrew formula but sometimes you may need to
re-update it to a specific version. This command do that.


## Installation
```shell
$ brew tap jkawamoto/fgo
$ brew install fgo
```

Formula Go requires make, [goxc](https://github.com/laher/goxc), and
[ghr](https://github.com/tcnksm/ghr).
The above `brew install` command also installs those dependencies if necessary.


## License
This software is released under the MIT License, see [LICENSE](LICENSES.md).
