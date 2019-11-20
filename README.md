<div align="center">
	<p>
		<a href="https://circleci.com">
			<img alt="CircleCI" src="img/circleci-logo.svg" width="75" />
		</a>
	</p>
	<h1>Stubb - a Docker utility tool</h1>
</div>

[![Build Status](https://circleci.com/gh/CircleCI-Public/stubb.svg?style=shield)](https://circleci.com/gh/CircleCI-Public/stubb) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/CircleCI-Public/stubb/master/LICENSE)

***This project is brand new and in alpha right now. You may notice a lack of polish and/or drastic changes.***

`stubb` is a Docker utility tool useful for when you want to learn more about your Docker images.
You can list and compare Docker images, their tags, and pull metrics from Docker Hub.


## Table of Contents

- [Installing](#installing)
- Configuring
- Features


## Installing

### Debian Package (.deb) Instructions

Download the `.deb` file to the desired system.

For graphical systems, you can download it from the [GitHub Releases page][gh-releases].
Many distros allow you to double-click the file to install.
Via terminal, you can do the following:

```bash
wget https://github.com/CircleCI-Public/stubb/releases/download/v0.1.0/stubb_0.1.0_amd64.deb
sudo dpkg -i stubb_0.1.0_amd64.deb
```

`0.1.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

### Linux Snap

More instructions coming soon.

### macOS / Brew

More instructions coming soon.


## Configuring

Most commands don't need authentication.
Some, such as `stubb set description`, require Docker / Docker Hub credentials.
These can be set via the environment variables `DOCKER_USER` and `DOCKER_PASS`.


## Features

*Lists* - List images belonging to a namespace and tags belonging to an image.

*Metadata* - Set your image's Docker Hub description from the command-line.

Run `stubb help` to see all commands available.


## License

This repository is licensed under the MIT license.
The license can be found [here](./LICENSE).



[gh-releases]: https://github.com/CircleCI-Public/stubb/releases
