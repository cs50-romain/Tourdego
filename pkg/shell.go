package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cs50-romain/tourdego/pkg/color"
)

type Shell struct {
	prompt		string
	history		*os.File
	autocomplete	bool
	commands	map[string]func(...string) error
}

func NewShell(prompt string) *Shell {
	s := &Shell{
		prompt: prompt,
		autocomplete: false,
		commands: make(map[string]func(...string) error),
	}

	s.AddCommand("quit", func(...string) error {
		fmt.Println("Goodbye!")
		os.Exit(0)
		return nil
	})

	return s
}

func (s *Shell) Start() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Remember to print in color by calling SetPromptColor()
		fmt.Print(s.prompt)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		command := strings.Trim(userInput, "\n")
		handler, ok := s.commands[command]

		if !ok {
			fmt.Printf("%s%s%s\n",color.Red, "Invalid command", color.Reset)
			continue
		}
		// if history is enabled, we want to add to history file.

		err = handler()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

// Return a string with the color escape character and reset.
func (s *Shell) SetPromptColor(setColor string) {
	s.prompt = fmt.Sprintf("%s%s%s", setColor, s.prompt, color.Reset)
}

func (s *Shell) SetPromptBold(isBold bool) {
	if isBold {
		s.prompt = fmt.Sprintf("%s%s", color.Bold, s.prompt)
	}
}

func (s *Shell) AddCommand(command string, handler func(...string) error) {
	s.commands[command] = handler 
}
