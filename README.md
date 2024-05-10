# Tourdego

Tourdego is an interactive shell library written in golang. It's supposed to be small, give you options and you know fun to use.
As AI explains it very well, here is why it's named tourdego: 

**Tourde Go is a clever play on words that combines "tour de force" and "Go," the programming language. The name suggests that this small golang interactive shell is a powerful tool that can handle complex tasks with ease. The use of "Go" in the name also highlights the fact that this shell is specifically designed for use with the Go programming language. The addition of "Tour" adds a sense of adventure and excitement, suggesting that using this shell will be a fun and enjoyable experience.**

# Demo:
Here is a small example of how to use this shell:
```
package main

import (
	shell "cs50-romain/tourdego/pkg"
	"fmt"

	"cs50-romain/tourdego/pkg/color"
)

func main() {
	shell := shell.NewShell(">")
	shell.SetPromptColor(color.Blue)
	shell.AddCommand("view")

	Intro()
	fmt.Println(color.Reset)
	if err := shell.Start(); err != nil {
		panic(err)
	}
}

func Intro() {
	fmt.Printf("%s%s%s%s\n", color.Bold, color.Yellow, "This is an interactive shell.\nPress quit or exit to exit this shell. Please enjoy!", color.Reset)
}
```

# AutoCompletion & Command History
Autocompletion will be the last feature I add to this project for now. Command history will have to wait and I explain below.
Well this is a little harder to implement. Here's what I'm thinking: First the terminal would have to be in raw mode to capture a lot of things (up, down, enter, display the command history...). 
