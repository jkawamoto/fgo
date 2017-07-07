# Formula Go
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://jkawamoto.github.io/fgo/info/licenses/)
[![Build Status](https://travis-ci.org/jkawamoto/fgo.svg?branch=master)](https://travis-ci.org/jkawamoto/fgo)
[![wercker status](https://app.wercker.com/status/9ab256a7b25d4d3980ed7821136b4177/s/master "wercker status")](https://app.wercker.com/project/byKey/9ab256a7b25d4d3980ed7821136b4177)
[![go report](https://goreportcard.com/badge/github.com/jkawamoto/fgo)](https://goreportcard.com/report/github.com/jkawamoto/fgo)
[![Release](https://img.shields.io/badge/release-0.3.1-brightgreen.svg)](https://github.com/jkawamoto/fgo/releases/tag/v0.3.1)

[![fgo](https://jkawamoto.github.io/fgo/img/small-banner.png)](https://jkawamoto.github.io/fgo/)

Formula Go helps you to build and upload your software written in
[Go](https://golang.org/);
and then prepare [homebrew](http://brew.sh/) and [linuxbrew](http://linuxbrew.sh/) formulae and for it.

Formula Go assumes your software is hosted in [GitHub](https://github.com/),
and the pre-compiled binaries are uploaded in the release page of it.

## Usage
~~~
fgo [global options] command [arguments...]

COMMANDS:
    init     create Makefile and other related directories.
    build    build binaries, upload them, and update the brew formula.
    update   update the brew formula.
    help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
    --pkg NAME, -p NAME   directory NAME to store package files
                          (default: "pkg")
    --brew NAME, -b NAME  directory NAME to store homebrew formula
                          (default: "homebrew")
    --help, -h            show help
    --version, -v         print the version
~~~

### Initialization
After starting your project, initialization is required.
To initialize Formula Go, run `fgo init`.
It creates a Makefile, which will be used to compile your project,
and a homebrew formula template.

To create the homebrew formula template,
GitHub's user name and repository name are required.
By default, Formula Go checks your git configuration to get
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

This command also updates your homebrew formula. After finishing this command,
you need to push the updated formula.

This command takes the following option flags:

#### -t, --token
This flag specifies a GitHub API token used to upload binaries and create a
release page. If this flag is not given but `GITHUB_TOKEN` environment
variable has been set, the environment variable will be used.

If neither this flag nor `GITHUB_TOKEN` are given, `github.token` variable
in your `.gitconfig` will be used.

#### -b, --body
This flag specifies a release note for the given version.

If this flag isn't given but your `CHANGELOG.md` contains a release note
associated with that version, the release note will be used.

#### -p, --process
This flag specifies how many goroutines will be used to upload binaries.
The default number is as same as the number of CPUs.

#### --delete
If this flag is set and there is a release for the given version, the
release will be deleted and a new release for the given version will be
created.

#### --draft
If this flag is set, the new release won't be published and will be kept as
a draft.

#### --pre
If this flag is set, the new release will be marked as a prerelease.


### Update
`fgo update [version]` updates the homebrew formula for a given version.
build command updates the homebrew formula but sometimes you may need to
re-update it to a specific version. This command do that.


## Installation
Formula Go is available in [homebrew](http://brew.sh/) and
[linuxbrew](http://linuxbrew.sh/).

```shell
$ brew tap jkawamoto/fgo
$ brew install fgo
```

Formula Go requires make, [goxc](https://github.com/laher/goxc), and
[ghr](https://github.com/tcnksm/ghr).
The above `brew install` command also installs those dependencies
if necessary.


## License
This software is released under the MIT License, see [LICENSE](https://jkawamoto.github.io/fgo/info/licenses/).
