package behavioral

type Iterator interface {
	HasMore() bool
	GetNext() *User
}

type UserIterator struct {
	index int
	uc    *UserCollection
}

func (ui *UserIterator) HasMore() bool {
	return len(ui.uc.users) > ui.index
}

func (ui *UserIterator) GetNext() *User {
	users := ui.uc.users
	if !ui.HasMore() {
		return nil
	}
	u := users[ui.index]
	ui.index++
	return u
}

type Collection interface {
	GetIterator() Iterator
}

type UserCollection struct {
	users []*User
}

func (uc *UserCollection) GetIterator() Iterator {
	return &UserIterator{0, uc}
}

type User struct {
	Grade int
}
