package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets6c245014d82f700ec3bc61559707d49fe804e5af = "#\n# Makefile\n#\nVERSION = snapshot\nGHRFLAGS =\n.PHONY: build release\n\ndefault: build\n\nbuild:\n\tgoxc -d={{.Dest}} -pv=$(VERSION)\n\nrelease:\n\tghr {{if .UserName}} -u {{.UserName}} {{end}} $(GHRFLAGS) v$(VERSION) {{.Dest}}/$(VERSION)\n"
var _Assets0c0f86697a7d7a06cdd29a8aa8c625c62138aff5 = "require 'rbconfig'\nclass {{.Package | Title}} < Formula\n  desc \"{{.Description}}\"\n  homepage \"https://github.com/{{.UserName}}/{{.Package}}\"\n  version \"{{\"{{.Version}}\"}}\"\n\n  if Hardware::CPU.is_64_bit?\n    case RbConfig::CONFIG['host_os']\n    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/\n      :windows\n    when /darwin|mac os/\n      url \"https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{\"{{.Version}}\"}}/{{\"{{.Mac64.FileName}}\"}}\"\n      sha256 \"{{\"{{.Mac64.Hash}}\"}}\"\n    when /linux/\n      url \"https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{\"{{.Version}}\"}}/{{\"{{.Linux64.FileName}}\"}}\"\n      sha256 \"{{\"{{.Linux64.Hash}}\"}}\"\n    when /solaris|bsd/\n      :unix\n    else\n      :unknown\n    end\n  else\n    case RbConfig::CONFIG['host_os']\n    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/\n      :windows\n    when /darwin|mac os/\n      url \"https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{\"{{.Version}}\"}}/{{\"{{.Mac386.FileName}}\"}}\"\n      sha256 \"{{\"{{.Mac386.Hash}}\"}}\"\n    when /linux/\n      url \"https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{\"{{.Version}}\"}}/{{\"{{.Linux386.FileName}}\"}}\"\n      sha256 \"{{\"{{.Linux386.Hash}}\"}}\"\n    when /solaris|bsd/\n      :unix\n    else\n      :unknown\n    end\n  end\n\n  def install\n    bin.install \"{{.Package}}\"\n  end\n\n  test do\n    system \"{{.Package}}\"\n  end\n\nend\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"Makefile", "formula.rb"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1547619795, 1547619795096660261),
		Data:     nil,
	}, "/Makefile": &assets.File{
		Path:     "/Makefile",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1547616782, 1547616782860661995),
		Data:     []byte(_Assets6c245014d82f700ec3bc61559707d49fe804e5af),
	}, "/formula.rb": &assets.File{
		Path:     "/formula.rb",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1547616782, 1547616782860861788),
		Data:     []byte(_Assets0c0f86697a7d7a06cdd29a8aa8c625c62138aff5),
	}}, "")
