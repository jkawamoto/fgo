class {{.Package | Title}} < Formula
  desc ""
  homepage "https://github.com/{{.UserName}}/{{.Package}}"
  version "{{"{{.Version}}"}}"

  if Hardware::CPU.is_64_bit?
    url "https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{"{{.Version}}"}}/{{"{{.Mac64.FileName}}"}}"
    sha256 "{{"{{.Mac64.Hash}}"}}"
  else
    url "https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{"{{.Version}}"}}/{{"{{.Mac386.FileName}}"}}"
    sha256 "{{"{{.Mac386.Hash}}"}}"
  end

  def install
    bin.install "{{.Package}}"
  end

  test do
    system "{{.Package}}"
  end

end
