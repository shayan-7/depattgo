package behavioral

type command interface {
	Execute()
}

type Button struct {
	command
}

func (b *Button) Press() {
	b.command.Execute()
}

type OnCommand struct {
	*TV
}

func (c *OnCommand) Execute() {
	c.On()
}

type OffCommand struct {
	*TV
}

func (c *OffCommand) Execute() {
	c.Off()
}

type TV struct {
	isRunning bool
}

func (tv *TV) On() {
	tv.isRunning = true
}

func (tv *TV) Off() {
	tv.isRunning = false
}
