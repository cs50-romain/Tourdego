package main

import (
	shell "cs50-romain/tourdego/pkg"
	"fmt"

	"cs50-romain/tourdego/pkg/color"
)

func main() {
	shell := shell.NewShell(">")
	shell.SetPromptColor(color.Blue)
	shell.AddCommand("view", func(s ...string) error {
		fmt.Println("Viewing something in the distance...")
		return nil
	})

	Intro()
	fmt.Println(color.Reset)
	if err := shell.Start(); err != nil {
		panic(err)
	}
}

func Intro() {
	fmt.Printf("%s%s%s%s\n", color.Bold, color.Yellow, "This is an interactive shell.\nPress quit or exit to exit this shell. Please enjoy!", color.Reset)
}
