![Sonar banner](./img/github-banner.png)

# Sonar - the Docker utility [![CI Status](https://circleci.com/gh/felicianotech/sonar.svg?style=shield)](https://app.circleci.com/pipelines/github/felicianotech/sonar) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/felicianotech/sonar/trunk/LICENSE) [![CircleCI Orb Version](https://badges.circleci.com/orbs/hubci/sonar.svg)][orb-page]

Sonar is the Docker and Docker Hub utility that you've been missing.
It can display information on Docker images, tags, and layers including the packages installed in those images.

Docker Hub metrics such as stars and pulls can be read while tasks such as updating the readme, summary, or even deleting tags from Docker Hub can be done with Sonar.


## Table of Contents

- [Install Sonar](#install-sonar)
  - [Linux](#linux)
  - [macOS](#macos)
  - [Windows](#windows)
  - [Continuous Integration (CI) Systems](#continuous-integration-ci-systems)
- [Configuring](#configuring)
- [Features](#features)


## Install Sonar

### Linux

There are a few ways you can install Sonar on a Linux amd64 or arm64 system.

#### Linux Snap (recommended)
Sonar can be installed via snap for users of Ubuntu (and Ubuntu flavors), `Pop!_OS`, Debian, Fedora, and more.
It's an easy installation that auto updates, allows to you install multiple versions of Sonar, or even follow the development track.
You can install Sonar via snap by running in a terminal:

```bash
sudo snap install sonar
sudo snap connect sonar:docker docker
```

The latter command is needed if you use commands such as `sonar packages list`.
This is only available in Sonar edge right now.
It'll be available in the v0.14 release.

If you don't have the `snap` command available, you might be able to find instructions for your distro [here](https://docs.snapcraft.io/core/install).

You can also install via the Snap Store by clicking the following button:

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/sonar)

#### Debian Package (.deb)
You can install Sonar into an Apt based computer by download the `.deb` file to the desired system.

For graphical systems, you can download it from the [GitHub Releases page][gh-releases].
Many distros allow you to double-click the file to install.
Via terminal, you can do the following:

```bash
wget https://github.com/felicianotech/sonar/releases/download/v0.12.0/sonar_0.12.0_amd64.deb
sudo dpkg -i sonar_0.12.0_amd64.deb
```

`0.12.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

#### Binary Install
You can download and run the raw Sonar binary from the [GitHub Releases page][gh-releases] if you don't want to use any package manager.
Simply download the tarball for your OS and architecture and extract the binary to somewhere in your `PATH`.
Here's one way to do this with `curl` and `tar`:

```bash
dlURL="https://github.com/felicianotech/sonar/releases/download/v0.12.0/sonar-v0.12.0-linux-amd64.tar.gz"
curl -sSL $dlURL | sudo tar -xz -C /usr/local/bin sonar
```

`0.12.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

### macOS

There are two ways you can install Sonar on a macOS (amd64) system.
Support for M1 macs (the arm64 chip) is coming in early 20201.

#### Brew (recommended)

Installing Sonar via brew is a simple one-liner:

```bash
brew install sonar
```

#### Binary Install
You can download and run the raw Sonar binary from the [GitHub Releases page][gh-releases] if you don't want to use Brew.
Simply download the tarball for your OS and architecture and extract the binary to somewhere in your `PATH`.
Here's one way to do this with `curl` and `tar`:

```bash
dlURL="https://github.com/felicianotech/sonar/releases/download/v0.12.0/sonar-v0.12.0-macos-amd64.tar.gz"
curl -sSL $dlURL | sudo tar -xz -C /usr/local/bin sonar
```

`0.12.0` may need to be replaced with your desired version.

### Windows

Sonar supports Windows 10 by downloading and installing the binary.
Chocolately support is likely coming in the future.
If there's a Windows package manager you'd like support for (including Chocolately), please open and Issue and ask for it.

#### Binary Install (exe)
You can download and run the Sonar executable from the [GitHub Releases page][gh-releases].
Simply download the zip for architecture and extract the exe.

### Continuous Integration (CI) Systems

Sonar can be installed in a CI environment pretty much the same way you'd install it on your own computer.
There is 1st-party support for some CI platforms in order to make the process easier.

#### CircleCI
Sonar can be installed in a CircleCI build by using the [Sonar CircleCI Orb][orb-page].
Please visit that link for more instructions.



#### GitHub Actions
Coming soon, probably.
Open an Issue to request it and demonstrate demand.


## Configuring

Most commands don't need authentication.
Some, such as `sonar set description`, require Docker / Docker Hub credentials.
These can be set via the environment variables `DOCKER_USER` and `DOCKER_PASS`.

You can also set the password specifically with the global `--password` flag.
Please be careful with this, you don't want a password in your shell history.
The suggestion is to pass an environment variable (that isn't already `DOCKER_PASS`) by doing something like this:

```bash
sonar set readme my/image --password=$THE_PASSWORD_ENVAR ./file.md
```


## Features

*Lists* - List images belonging to a namespace and tags belonging to an image.

*Metadata* - Set your image's Docker Hub description from the command-line.

*Packages* - List the Apt/deb, RPM, and/or Pip packages installed in an image.

*Tag Deletion* - Delete tags on Docker Hub individually or in bulk.

Run `sonar help` to see all commands available.


## License

This repository is licensed under the MIT license.
The license can be found [here](./LICENSE).



[gh-releases]: https://github.com/felicianotech/sonar/releases
[orb-page]: https://circleci.com/developer/orbs/orb/hubci/sonar
