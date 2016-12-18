---
title: "Formula Go"
description: >-
  Build, upload, and create brew formula for golang application.
date: 2016-12-17
lastmod: 2016-12-17
slug: readme
---
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](./info/licenses/)
[![Build Status](https://travis-ci.org/jkawamoto/fgo.svg?branch=master)](https://travis-ci.org/jkawamoto/fgo)
[![Code Climate](https://codeclimate.com/github/jkawamoto/fgo/badges/gpa.svg)](https://codeclimate.com/github/jkawamoto/fgo)
[![Release](https://img.shields.io/badge/release-0.2.3-brightgreen.svg)](https://github.com/jkawamoto/fgo/releases/tag/v0.2.3)

Build, upload, and create brew formula for golang application.

### Requires
* make
* [goxc](https://github.com/laher/goxc)
* [ghr](https://github.com/tcnksm/ghr)


### Usage
~~~
fgo [global options] command [arguments...]

COMMANDS:
     init     create Makefile and other related directories.
     build    build binaries, upload them, an update brew formula.
     update   update only brew formula.
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --pkg NAME, -p NAME   overwrite directory NAME to store package files. (Default: pkg)
   --brew NAME, -b NAME  overwrite directory NAME to store homebrew formula. (Default: homebrew)
   --help, -h            show help
   --version, -v         print the version
~~~


### Installation
```sh
$ go get github.com/jkawamoto/fgo
```
or if you're a [Homebrew](http://brew.sh/) user,

```sh
$ brew tap jkawamoto/fog
$ brew install fgo
```

Compiled binaries are also available in
[Github](https://github.com/jkawamoto/fgo/releases).


### License
This software is released under the MIT License, see [LICENSE](./info/licenses/).
