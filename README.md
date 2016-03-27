<p align="center">
  <img src="http://i.imgur.com/JPtZVfU.png" alt="Liberty CLI"/>
</p>

*liberty - the state of being free within society from oppressive restrictions imposed by authority on one's behaviour or political views*


### Note
This project is currently a draft, it's not ready for use yet and depends heavily on those willing to contribute. If you're interested in helping, please don't hesitate.


### What is Liberty
Liberty is a simple, friendly and powerful dependency manager for Javascript projects that has no single point of failure. We do this by avoiding a central list of packages and we go straight to the source. We encourage developers to think carefully about what dependencies they want in their projects. If you need a package and you want to ensure it still exists when the original author removes the code, you can fork your own project and use that.

Liberty is proud to offer:
- A simple, convenient way to manage dependencies
- No single point of failure
- No namespace conflicts


### Installation
We do not have a public binary for Liberty at the moment. You are welcome to download the source and run `./build.sh`. However, this will fail if you don't have the required dependencies. We are looking into having a script to download any pre-requisites you'll need to build the source.

### Getting Started

#### init
To integrate your project with Liberty, you will need to have a liberty file. This is represented as `liberty.json` at the root of your project directory. You can run `liberty init` to generate this file for you. It will ask you a few questions about your project. If you answer these questions, you can `liberty publish` your package later. This will allow other users to search for your package and find basic details about it, whilst not being dependent on our servers.

#### install
Once you have created your liberty file and populated it with dependencies, you can execute `liberty install` which will fetch the dependencies and place them in the `_libs` directory. If you have included a binary in your dependencies, this will be installed into a location in your $PATH that will allow you to access the binary via the shell.

#### update
After a while, you'll need to update Liberty. If you execute `liberty update` you will be upgraded to the latest stable version of liberty.
