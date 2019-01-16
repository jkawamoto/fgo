# Formula Go
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://jkawamoto.github.io/fgo/info/licenses/)
[![Build Status](https://travis-ci.org/jkawamoto/fgo.svg?branch=master)](https://travis-ci.org/jkawamoto/fgo)
[![wercker status](https://app.wercker.com/status/9ab256a7b25d4d3980ed7821136b4177/s/master "wercker status")](https://app.wercker.com/project/byKey/9ab256a7b25d4d3980ed7821136b4177)
[![go report](https://goreportcard.com/badge/github.com/jkawamoto/fgo)](https://goreportcard.com/report/github.com/jkawamoto/fgo)
[![Release](https://img.shields.io/badge/release-0.3.4-brightgreen.svg)](https://github.com/jkawamoto/fgo/releases/tag/v0.3.4)

[![fgo](https://jkawamoto.github.io/fgo/img/small-banner.png)](https://jkawamoto.github.io/fgo/)

Formula Go helps you to build and upload your software written in
[Go](https://golang.org/);
and then prepare [Homebrew](http://brew.sh/) and
[Linuxbrew](http://linuxbrew.sh/) formulae and for it.

Formula Go assumes your software is hosted in [GitHub](https://github.com/),
and the pre-compiled binaries are uploaded in the release page of it.

## Usage
### Step 1: Initialization for your project
Run the following command in the top directory of your project
(same directory as `.git` exists):

```shell
$ fgo init
```

It creates `homebrew` directory and generates a template of Homebrew
formulae there.

Formula Go uses `make` to compile and upload your program.
This initialization step creates a `Makefile` for this purpose as well as
a template of Homebrew formulae.

To create the template, GitHub's user name and repository name are required.
By default, Formula Go checks your git configuration to obtain
those information but you can give them via arguments
(run `fgo init -h` for more information).

If your git configuration doesn't have both information and you don't
give them as arguments, this command will skip to create the template.
In this case, you need to re-run init command after setting git
configuration.

You can edit the Makefile and the template of Homebrew formulae.
Make sure `build` and `release` targets in the Makefile are necessary for
Formula Go.

### Step 2: Build and upload your program
When you're ready for releasing your program as a certain version, run

```shell
$ fgo build [version]
```

It starts to build your program to `pkg` directory and uploads built binary
files to GitHub.
After uploading binary files, it updates a Homebrew formula in `homebrew`
directory.

For debugging, you may need to build your program but omit uploading.
In this case, omit version name in the above command.
Formula Go builds your program in `pkg/snapshot` directory without uploading.

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

### Update the brew formula without build
`fgo update [version]` only updates the Homebrew formula for a given version.

build command updates the Homebrew formula but sometimes you may need to
re-update it for a specific version. This command do that.


### Commands and options
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
