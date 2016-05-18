package sms

import (
	"github.com/remogatto/application"
	"github.com/veandco/go-sdl2/sdl"
)

var keyMap = map[interface{}]int{
	sdl.SCANCODE_UP:    1, // Arrow keys
	sdl.SCANCODE_DOWN:  2,
	sdl.SCANCODE_LEFT:  4,
	sdl.SCANCODE_RIGHT: 8,
	sdl.SCANCODE_Z:     16, // Z and X for fire
	sdl.SCANCODE_X:     32,
	sdl.SCANCODE_R:     1 << 12, // R for reset button
}

type inputLoop struct {
	sms              *SMS
	pause, terminate chan int
}

func NewInputLoop(sms *SMS) *inputLoop {
	return &inputLoop{
		sms:       sms,
		pause:     make(chan int),
		terminate: make(chan int),
	}
}

func (l *inputLoop) Pause() chan int {
	return l.pause
}

func (l *inputLoop) Terminate() chan int {
	return l.terminate
}

func (l *inputLoop) Run() {
	var event sdl.Event

	running := true

	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				application.Debugf("[%d ms] KeyDown\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)

				if val, ok := keyMap[t.Keysym.Scancode]; ok {
					l.sms.Command <- CmdJoypadEvent{keyMap[val], JOYPAD_DOWN}
				} else if t.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
					running = false
				}
			case *sdl.KeyUpEvent:
				l.sms.Command <- CmdJoypadEvent{keyMap["up"], JOYPAD_UP}
			}
		}
	}

	sdl.Quit()

	application.Exit()
}
