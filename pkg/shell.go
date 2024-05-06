package shell

import (
	"bufio"
	"fmt"
	"os"

	"cs50-romain/tourdego/pkg/color"
)

type Shell struct {
	prompt		string
	history		*os.File
	autocomplete	bool
	commands	map[string]func() error
}

func NewShell(prompt string) *Shell {
	return &Shell{
		prompt: prompt,
		autocomplete: false,
		commands: make(map[string]func() error),
	}
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
		fmt.Println(userInput)
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

func (s *Shell) AddCommand(command string, option ...string) {
	s.commands[command] = func () error {
		return nil
	}
}
