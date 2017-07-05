require 'rbconfig'
class {{.Package | Title}} < Formula
  desc ""
  homepage "https://github.com/{{.UserName}}/{{.Package}}"
  version "{{"{{.Version}}"}}"

  if Hardware::CPU.is_64_bit?
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{"{{.Version}}"}}/{{"{{.Mac64.FileName}}"}}"
      sha256 "{{"{{.Mac64.Hash}}"}}"
    when /linux/
      :linux
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  else
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{"{{.Version}}"}}/{{"{{.Mac386.FileName}}"}}"
      sha256 "{{"{{.Mac386.Hash}}"}}"
    when /linux/
      :linux
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  end

  def install
    bin.install "{{.Package}}"
  end

  test do
    system "{{.Package}}"
  end

end
