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
