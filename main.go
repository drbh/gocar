package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

type TomlConfig struct {
	Package packageS          `toml:"package"`
	Bin     binary            `toml:"bin"`
	Deps    map[string]string `toml:"deps"`
	Gets    map[string]string `toml:"gets"`
}

type packageS struct {
	Name    string
	Version string
	Authors []string
}

type binary struct {
	Loc string
}

func createDirectory(directoryPath string) {
	//choose your permissions well
	pathErr := os.MkdirAll(directoryPath, 0777)

	//check if you need to panic, fallback or report
	if pathErr != nil {
		fmt.Println(pathErr)
	}
}
func createFiles(filePath, content string) {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"

	app.Action = func(c *cli.Context) error {

		switch c.Args().Get(0) {
		case "new":

			name := c.Args().Get(1)
			createDirectory("./" + name + "/src")

			fmt.Println("✨ Created" + name)

			createFiles("./"+name+"/src/main.go", `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
`)
			createFiles("./"+name+"/gocar.toml", `[package]
name = "`+name+`"
version = "0.1.0"
authors = [""]

[bin]
loc = "src/main.go"

[deps]

[gets]
`)
			createFiles("./"+name+"/.gitignore", ``)

			createFiles("./"+name+"/README.md", `# `+name+`

**generated with gocar**
`)
			fmt.Println("✅ Completed template process")

		case "run":

			var config TomlConfig
			if _, err := toml.DecodeFile("gocar.toml", &config); err != nil {
				fmt.Println(err)
				return nil
			}
			out, err := exec.Command("go", "run", "./"+config.Bin.Loc).Output()
			if err != nil {
				fmt.Printf("%s", err)
			}
			output := string(out[:])
			fmt.Println(output)

			// fmt.Println(config)

		case "add":

			fmt.Println("should add line to .toml")

		default:
			fmt.Println("It's after noon")
		}
		return nil

	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
