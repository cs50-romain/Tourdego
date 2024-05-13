// This package represents an interactive shell
package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cs50-romain/tourdego/pkg/color"
	"cs50-romain/tourdego/pkg/esc"
	"golang.org/x/term"
)

// Represents a shell
type Shell struct {
	prompt		string
	history		*os.File
	autocomplete	bool
	RootCmd		*Cmd
	RawMode		bool
	//Remember the previous command - Will be an option that can be set or unset
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
		//os.Exit(0)
		return nil
	})
	s.AddCommand("quit", quitCommand)
	
	s.AddCommand("exit", &Cmd{
		Name: "exit",
		Help: "exit the program",
		Handler: func(...string) error {
			fmt.Println("Goodbye!")
			s.cleanUp()
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

// Raw mode should be enabled if autocompletion is required or if user wants raw mode.
func WithRawMode() {

}

// Starts a Shell
func (s *Shell) Start() error {
	if s.RawMode == true || s.autocomplete == true {
		fmt.Println("Running raw mode...")

		return s.RunRawMode()	
	} else {
		return s.RunCookedMode()
	}
}

func (s *Shell) RunRawMode() error {
	// save current commands to json file.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	buf := make([]byte, 0)
	fmt.Print(s.prompt)
	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)

		if b[0] == esc.CTRLC() {
			fmt.Printf("\n%s\n", esc.MoveCursorLeft(1000))
			TermWrite([]byte(esc.MoveCursorLeft(1000)), "")
			return nil
		} else if b[0] == esc.TAB() {
			os.Stdout.Write([]byte("tabbed\n"))
			fmt.Printf("%s%s ", esc.MoveCursorLeft(1000), s.prompt)
		} else if b[0] == esc.ENTER() {
			TermWrite(esc.NEWLINE(), "")
			s.parseCommand(string(buf))
			
			if string(buf) == "quit" || string(buf) == "exit" {
				fmt.Printf("%s", esc.MoveCursorLeft(1000))
				TermWrite([]byte(esc.MoveCursorLeft(1000)), "")
				return nil
			}

			TermWrite(esc.NEWLINE(), s.prompt + " ")
			buf = make([]byte, 0)
		} else {
			os.Stdout.Write(b)
			buf = append(buf, b...)
		}
		// Expand buffer size if current buffer too small
	}
}

func parseByte() {
	buf := make([]byte, 32)
	b := make([]byte, 1)
	os.Stdin.Read(b)

	if b[0] == esc.CTRLC() {
		fmt.Printf("\n%sGoodbye...\n", esc.MoveCursorLeft(1000))
		//fmt.Print(cursor_left)
		TermWrite([]byte(esc.MoveCursorLeft(1000)), "")
		return
	} else if b[0] == esc.TAB() {
		os.Stdout.Write([]byte("tabbed\n"))
		fmt.Printf("%s> ", esc.MoveCursorLeft(1000))
	} else if b[0] == esc.ENTER() {
		// Parse the command
		TermWrite(esc.NEWLINE(), "")
		os.Stdout.Write(buf[2:])
		//os.Stdout.Write([]byte("\u001b[E> "))
		TermWrite(esc.NEWLINE(), "> ")
		buf = make([]byte, 32)
	} else {
		os.Stdout.Write(b)
		buf = append(buf, b...)
	}
}

func TermWrite(ESC_CODE []byte, extra string) {
	line := fmt.Sprintf("%s%s%s", esc.ESCAPE(), ESC_CODE, extra)
	os.Stdout.Write([]byte(line))
}

func (s *Shell) RunCookedMode() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Remember to print in color by calling SetPromptColor()
		fmt.Print(s.prompt)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		// This code can probably be moved to a function
		lineAfterRemovingEOL := strings.Trim(userInput, "\n")
		s.parseCommand(lineAfterRemovingEOL)
	}
}

func (s *Shell) parseCommand(input string) error {
	var err error
	userInputSplit := strings.Split(input, " ")

	commandName := userInputSplit[0]
	subcommands := userInputSplit[1:]
	command, ok := s.commands[commandName]
	
	if !ok {
		// Try the lastCommand's NextCmd's handler
		//err = s.lastCommand.NextCmd.Handler(userInputSplit...)
		//if err != nil {
		fmt.Printf("%s%s%s\n",color.Red, "Invalid command", color.Reset)
		//}
		return nil
	}
	// if history is enabled, we want to add to history file.

	err = command.Handler(subcommands...)
	if err != nil {
		return err
	}
	// Make the command the last command executed
	s.lastCommand = command
	return nil
}

func (s *Shell) cleanUp() {
	if s.RawMode == true {
		fmt.Printf("%s\n", esc.MoveCursorLeft(1000))
		TermWrite([]byte(esc.MoveCursorLeft(1000)), "")
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
