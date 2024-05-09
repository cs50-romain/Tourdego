package main

import (
	shell "cs50-romain/tourdego/pkg"
	"fmt"

	"cs50-romain/tourdego/pkg/color"
)

func main() {
	sh := shell.NewShell(">")
	sh.SetPromptColor(color.Blue)
	selectCmd := &shell.Cmd{
		Name: "select",
		Help: "select an option",
		Handler: func(s ...string) error {
			fmt.Println("Selecting option", s)
			return nil
		},
	}
	sh.AddCommand("select", selectCmd)

	sh.AddCommand("view", &shell.Cmd{
		Name: "view",
		Help: "echo the user input",
		Handler: func (...string) error {
			fmt.Println("Viewing something in the distance")
			return nil
		},
		NextCmd: selectCmd,
	})

	Intro()
	fmt.Println(color.Reset)
	if err := sh.Start(); err != nil {
		panic(err)
	}
}

func Intro() {
	fmt.Printf("%s%s%s%s\n", color.Bold, color.Yellow, "This is an interactive shell.\nPress quit or exit to exit this shell. Please enjoy!", color.Reset)
}
