package structural

type Device interface {
	isEnabled() bool
	enable()
	disable()
	getVolume() int
	setVolume(int)
	getChannel() int
	setChannel(int)
}

type Remote struct {
	Device
}

func NewRemote(d Device) *Remote {
	return &Remote{d}
}

func (r *Remote) TogglePower() {
	if r.isEnabled() {
		r.disable()
		return
	}
	r.enable()
}

func (r *Remote) VolumeUp() {
	r.setVolume(r.getVolume() + 10)
}

func (r *Remote) VolumeDown() {
	r.setVolume(r.getVolume() - 10)
}

func (r *Remote) ChannelUp() {
	r.setChannel(r.getChannel() + 1)
}

func (r *Remote) ChannelDown() {
	r.setChannel(r.getChannel() - 1)
}

type TV struct {
	power   bool
	volume  int
	channel int
}

func NewTV() *TV {
	return &TV{power: false, volume: 100, channel: 1}
}

func (tv *TV) isEnabled() bool {
	return tv.power
}

func (tv *TV) enable() {
	tv.power = true
}

func (tv *TV) disable() {
	tv.power = false
}

func (tv *TV) getVolume() int {
	return tv.volume
}

func (tv *TV) setVolume(v int) {
	tv.volume = v
}

func (tv *TV) getChannel() int {
	return tv.channel
}

func (tv *TV) setChannel(v int) {
	tv.channel = v
}

type Radio struct {
	power   bool
	volume  int
	channel int
}

func NewRadio() *Radio {
	return &Radio{power: false, volume: 100, channel: 1}
}

func (r *Radio) isEnabled() bool {
	return r.power
}

func (r *Radio) enable() {
	r.power = true
}

func (r *Radio) disable() {
	r.power = false
}

func (r *Radio) getVolume() int {
	return r.volume
}

func (r *Radio) setVolume(v int) {
	r.volume = v
}

func (r *Radio) getChannel() int {
	return r.channel
}

func (r *Radio) setChannel(v int) {
	r.channel = v
}
