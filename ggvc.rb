# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Ggvc < Formula
  desc "Script for generating semantic version based on the version in gradle.properties and git branch"
  homepage "https://github.com/titenkov/gradle-git-version-calculator"
  url "https://github.com/titenkov/gradle-git-version-calculator/archive/v0.0.1.tar.gz"
  sha256 "8f01873a98531832905fac4d47e351fd4e13a9c47994fb2f2c7ea7cecb7b1ad8"
  # depends_on "cmake" => :build

  def install
    bin.install "ggvc"
  end
end