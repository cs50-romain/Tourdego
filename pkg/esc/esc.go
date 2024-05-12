package esc

import "fmt"

/// To get autocomplete working, terminal will need to be in raw mode. That allows us to read/listen to each character being inputted. But requires more work to read and write.
// Read: Read each character and add to buffer. If character is \t or \n, something will need to done. Until then keep reading and adding to buffer.
// When a \n is read, we parse the whole buffer (command + options) and return new line. Display prompt
// When a \t is read, we parse the command. Look for commands something similar to the command. Create an queue of those commands in order. Then user ASCII escape code to move cursor up, then left, then replace line with the prompt and the most similar work.
// Cache the original command. If the next input is a \t, then look for the second most similar word and do the escape codes again.

const (
	cursor_up = "\033[1A"
	cursor_left = "\u001b[1000D"
	ctrlc = 3
	enter = 13
	tab = 9
)

var newline = []byte("\u001b[E")
var escape = []byte("\u001b")

func NEWLINE() []byte {
	return newline	
}

func ESCAPE() []byte {
	return escape
}

func CTRLC() byte {
	return ctrlc
}

func ENTER() byte {
	return enter
}

func TAB() byte {
	return tab
}

func MoveCursorFarLeft() string {
	return fmt.Sprintf("\u001b[1000D")
}

func MoveCursorLeft(amount int) string {
	return fmt.Sprintf("\u001b[%dD", amount)
}

func MoveCursorUp(amount int) string {
	return fmt.Sprintf("\u001b[%dA", amount)
}
