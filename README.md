<p align="center">
  <img src="http://i.imgur.com/JPtZVfU.png" alt="Liberty CLI"/>
</p>

*liberty - the state of being free within society from oppressive restrictions imposed by authority on one's behaviour or political views*


### Note
This project is currently a draft, it's not ready for use yet and depends heavily on those willing to contribute. If you're interested in helping, please don't hesitate.

[![Build Status](https://travis-ci.org/liberty-org/cli.svg?branch=master)](https://travis-ci.org/liberty-org/cli)


### What is Liberty
Liberty is a simple, friendly and powerful dependency manager for Javascript projects that has no single point of failure. We do this by avoiding a central list of packages and we go straight to the source. We even support the IPFS as we strongly believe this is the direction that the web should evolve. We encourage developers to think carefully about what dependencies they want in their projects. If you need a package and you want to ensure it still exists when the original author removes the code, you can fork your own project and use that. It is ultimately an npm alternative.

Liberty is proud to offer:
- A simple, convenient way to manage dependencies
- IPFS support
- No single point of failure
- No namespace conflicts


### Installation
We do not have a public binary for Liberty at the moment. You are welcome to download the source and run `make`. However, this will fail if you don't have the required dependencies. We are looking into having the makefile download any pre-requisites you'll need to build the source.

### Getting Started

#### init
To integrate your project with Liberty, you will need to have a liberty file. This is represented as `liberty.json` at the root of your project directory. You can run `liberty init` to generate this file for you. It will ask you a few questions about your project. If you answer these questions, you can `liberty publish` your package later. This will allow other users to search for your package and find basic details about it, whilst not being dependent on our servers.

#### install
Once you have created your liberty file and populated it with dependencies, you can execute `liberty install` which will fetch the dependencies and place them in the `_libs` directory. If you have included a binary in your dependencies, this will be installed into a location in your $PATH that will allow you to access the binary via the shell.

#### update
After a while, you'll need to update Liberty. If you execute `liberty update` you will be upgraded to the latest stable version of liberty.

#### get
Use this to install single packages from the command line without first updating your liberty.json file

#### remove
Allows you to remove a lone package, specified by name via command line

#### help
Emit a summary of available subcommands and their purpose

#### set
Set a package to a specific version, denoted by a git tag
