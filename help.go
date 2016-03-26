package main

type Argument struct {
	Name  string
	Usage string
}

type Arguments []Argument

func GetHelpArguments() Arguments {
	var arguments = Arguments{
		Argument{
			Name: "liberty init",
			Usage: "Initializes your project with a liberty file. If you provide the tool with additional information,\nit will enable other developers to find your code using our website. However, we don't store anything more than we ask.",
		},
		Argument{
			Name: "liberty install",
			Usage: "Downloads and installs any dependencies in your liberty file.\nThe'lib' dependencies will be installed in a directory names _libs at the root of your project.",
		},
		Argument{
			Name: "liberty update",
			Usage: "Updates liberty to the latest stable version.",
		},
		Argument{
			Name: "liberty version",
			Usage: "Displays the version that is in the $PATH. Also displays the build architecture.",
		},
		Argument{
			Name: "liberty list",
			Usage: `Displays a list of the installed packages and their versions. Also displays whether they are 'binary' dependencies or 'lib' dependencies.`,
		},
	}
	return arguments
}