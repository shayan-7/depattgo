package behavioral

type Memento interface {
	Restore()
}

type Snapshot struct {
	state      string
	originator *Originator
}

func newSnapshot(s string, o *Originator) *Snapshot {
	return &Snapshot{s, o}
}

func (s *Snapshot) Restore() {
	s.originator.SetState(s.state)
}

type IOriginator interface {
	Save() Memento
}

type Originator struct {
	state string
}

func (o *Originator) Save() Memento {
	return newSnapshot(o.state, o)
}

func (o *Originator) SetState(s string) {
	o.state = s
}

type Command struct {
	mementoArray []Memento
}

func (c *Command) Undo() {
	if len(c.mementoArray) < 0 {
		return
	}
	lastIndex := len(c.mementoArray) - 1
	m := c.mementoArray[lastIndex]
	m.Restore()
	c.mementoArray = c.mementoArray[:lastIndex]
}

func (c *Command) Append(m Memento) {
	c.mementoArray = append(c.mementoArray, m)
}
