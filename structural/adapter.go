package structural

type port struct {
	speed int
	name  string
}

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c Client) GetPortSpeed(comp computer) int {
	return comp.toLightningPort().speed
}

type computer interface {
	// The toLightningPort returned value could be a
	// seperate LightningPort struct, instead of port struct
	// that doesn't denote explicitly which type of port we are dealing with.
	// but for the sake of simplicity, only one common type is defined.
	toLightningPort() port
}

type Mac struct {
	port
}

func NewMac() *Mac {
	return &Mac{
		port{200, "lightning"},
	}
}

func (m *Mac) toLightningPort() port {
	return m.port
}

type Windows struct {
	port
}

func NewWindows() *Windows {
	return &Windows{
		port{100, "windows"},
	}
}

func (w *Windows) toUSBport() port {
	return w.port
}

type WindowsAdapter struct {
	*Windows
}

func NewWindowsAdapter(w *Windows) *WindowsAdapter {
	return &WindowsAdapter{w}
}

func (wa *WindowsAdapter) toLightningPort() port {
	p := wa.toUSBport()
	return port{speed: p.speed * 2, name: "lightning"}
}
