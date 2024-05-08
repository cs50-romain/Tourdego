// This package represents an interactive shell
package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cs50-romain/tourdego/pkg/color"
)

// Represents a shell
type Shell struct {
	prompt		string
	history		*os.File
	autocomplete	bool
	RootCmd		*Cmd
	//Remember the previous command
	lastCommand	*Cmd
	commands	map[string]*Cmd
}

// Creates a new shell
func NewShell(prompt string) *Shell {
	s := &Shell{
		prompt: prompt,
		autocomplete: false,
		commands: make(map[string]*Cmd),
	}

	helpCommand := &Cmd{
		Name: "help",
		Help: "Displays help of each command. help <command|optional>",
	}
	helpCommand.Handler = func(subCommandsSpecified ...string) error {
		if len(subCommandsSpecified) == 0 {
			for _, command := range helpCommand.subCommands {
				fmt.Printf("  -%s: %s\n", command.Name ,command.Help)
			}
		} else {
			for _, subcommandName := range subCommandsSpecified {
				command, ok := s.commands[subcommandName] 
				if !ok {
					continue
				}

				fmt.Printf("  -%s: %s\n", command.Name ,command.Help)
			}
		}
		return nil
	}

	s.AddCommand("help", helpCommand)

	quitCommand := NewCmd("quit", "quit the program")
	quitCommand.HandlerMethod(func(subcommands ...string) error {
		fmt.Println("Goodbye!")
		os.Exit(0)
		return nil
	})
	s.AddCommand("quit", quitCommand)
	
	s.AddCommand("exit", &Cmd{
		Name: "exit",
		Help: "exit the program",
		Handler: func(...string) error {
			fmt.Println("Goodbye!")
			os.Exit(0)
			return nil		
		},
	})

	return s
}

func WithExitCommands() {

}

func WithHelp() {

}

// Starts a Shell
func (s *Shell) Start() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Remember to print in color by calling SetPromptColor()
		fmt.Print(s.prompt)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		commandName := strings.Trim(userInput, "\n")
		// More parsing needs done to separate the rootcommand from the subcommands/options
		command, ok := s.commands[commandName]

		

		if !ok {
			fmt.Printf("%s%s%s\n",color.Red, "Invalid command", color.Reset)
			continue
		}
		// if history is enabled, we want to add to history file.


		err = command.Handler()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Make the command the last command executed
		s.lastCommand = command
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

func (s *Shell) AddCommand(commandName string, command *Cmd) {
	s.commands[commandName] = command
	s.commands["help"].AddSubCommands(command)
}
