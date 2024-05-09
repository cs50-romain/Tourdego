package shell

type Cmd struct {
	Name	string
	options []string
	Help	string
	// Function to execute for the command
	Handler	func(...string) error
	subCommands []*Cmd
	NextCmd *Cmd
}

func NewCmd(name string, help string) *Cmd {
	return &Cmd{
		Name: name,
		Help: help,
	}
}

func (c *Cmd) AddSubCommands(userSpecifiedSubCommands ...*Cmd) {
	c.subCommands = append(c.subCommands, userSpecifiedSubCommands...)
}

func WithOptions() {
	
}

func AddChild(options ...Cmd) error {
	return nil	
}

func (c *Cmd) HandlerMethod(userSpecifiedHandler func(...string) error) {
	c.Handler = userSpecifiedHandler
}
