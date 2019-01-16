## 0.3.4 (2019-01-16)
### Fixed
- include vendor repositories 
- switch to go-assets from go-bindata


## 0.3.3 (2018-07-02)
### Fixed
- Version numbers in the download links


## 0.3.2 (2018-07-02)
### Update 
- Users can specify the command name instead of the repository name
- Support to deploy on Travis CI (Fixes #8)


## 0.3.1 (2017-07-07)
### Update
- `init` command takes `--desc` flag to specify description in brew formulae

### Fixed
- Output messages in Windows


## 0.3.0 (2017-07-06)
### Update
- Configuration file isn't used any more
- `init` command asks when overwriting existing files
- Supports [linuxbrew](http://linuxbrew.sh/)


## 0.2.5 (2017-06-16)
### Update
- `build` command takes option flags for [`ghr`](http://tcnksm.github.io/ghr/)
- `build` command parses `CHANGELOG.md` and creates a release note


## 0.2.4 (2016-12-21)
### Updated
- `init` command takes repository name to create a homebrew formula without git configurations.


## 0.2.3 (2016-08-12)
### Fixed
- Deprecated Hardware.is_64_bit? in formula.rb.


## 0.2.2 (2016-07-23)
### Fixed
- `Asset/formula.rb` has wrong links,
- Update command doesn't check global flags.


## 0.2.1 (2016-07-23)
### Fixed
- Problems of parsing global options.


## 0.2.0 (2016-07-21)
### New features
- `fgo update` command only updates homebrew formula,
  while `fgo build` command updates homebrew formula after building packages.
- `fgo init` command saves `--dest` and `--brew` options
  so that other commands use those given options.


## 0.1.1 (2016-07-21)
### Update
`fgo init` won't stop even if Makefile and/or a formula template exist,
so that uses can re-run this command without deleting generated files.


## 0.1.0 (2016-07-20)
Initial release
