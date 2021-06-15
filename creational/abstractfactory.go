package creational

import (
	"errors"
	"fmt"
)

type Category int

const (
	VueCategory Category = iota
	ReactCategory
)

var ErrFactoryCategory = errors.New("error factory category")

type factoryCreator func() IFactory

type IButton interface {
	show() string
}

type VueButton struct{}

func NewVueButton() *VueButton {
	return &VueButton{}
}

func (vb *VueButton) show() string {
	return "I'm showing vue button"
}

type ReactButton struct{}

func NewReactButton() *ReactButton {
	return &ReactButton{}
}

func (rb *ReactButton) show() string {
	return "I'm showing react button"
}

type IWindow interface {
	render() string
}

type Window struct {
	height int
	width  int
}

func NewWindow(h, w int) *Window {
	return &Window{h, w}
}

type VueWindow struct {
	Window
}

func NewVueWindow(h, w int) *VueWindow {
	return &VueWindow{
		Window: *NewWindow(h, w),
	}
}

func (vw *VueWindow) render() string {
	return fmt.Sprintf(
		"Rendering Vue window by height: %d and width: %d",
		vw.height,
		vw.width,
	)
}

type ReactWindow struct {
	Window
}

func NewReactWindow(h, w int) *ReactWindow {
	return &ReactWindow{
		Window: *NewWindow(h, w),
	}
}

func (rw *ReactWindow) render() string {
	return fmt.Sprintf(
		"Rendering React window by height: %d and width: %d",
		rw.height,
		rw.width,
	)
}

type IFactory interface {
	createButton() IButton
	createWindow(height, width int) IWindow
}

type VueFactory struct{}

var NewVueFactory = factoryCreator(func() IFactory {
	return &VueFactory{}
})

func (vf *VueFactory) createButton() IButton {
	return NewVueButton()
}

func (vf *VueFactory) createWindow(h, w int) IWindow {
	return NewVueWindow(h, w)
}

type ReactFactory struct{}

var NewReactFactory = factoryCreator(func() IFactory {
	return &ReactFactory{}
})

func (rf *ReactFactory) createButton() IButton {
	return NewReactButton()
}

func (rf *ReactFactory) createWindow(h, w int) IWindow {
	return NewReactWindow(h, w)
}

var factoryMap = map[Category]factoryCreator{
	VueCategory:   NewVueFactory,
	ReactCategory: NewReactFactory,
}

func GetFactory(c Category) (IFactory, error) {
	f, ok := factoryMap[c]
	if !ok {
		return nil, ErrFactoryCategory
	}
	return f(), nil
}
