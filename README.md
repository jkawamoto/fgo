Formula Go
============
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Build Status](https://travis-ci.org/jkawamoto/fgo.svg?branch=master)](https://travis-ci.org/jkawamoto/fgo)
[![Code Climate](https://codeclimate.com/github/jkawamoto/fgo/badges/gpa.svg)](https://codeclimate.com/github/jkawamoto/fgo)
[![Release](https://img.shields.io/badge/release-0.2.0-lightgrey.svg)](https://github.com/jkawamoto/fgo/releases/tag/v0.2.0)

Build, upload, and create brew formula for golang application.

Usage
------
~~~
commands [global options] command [arguments...]

COMMANDS:
     init    create Makefile and other related directories.
     build   build binaries, upload them, an update brew formula.
     update  update only brew formula.

GLOBAL OPTIONS:
   --dest NAME, -d NAME  overwrite directory NAME to store package files.
                         (default: "pkg")
   --brew NAME, -b NAME  overwrite directory NAME to store homebrew formula.
                         (default: "brew")
   --help, -h            show help
   --version, -v         print the version
~~~


License
=======
This software is released under the MIT License, see [LICENSE](LICENSE).
