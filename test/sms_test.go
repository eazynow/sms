package z80

import (
	smslib "github.com/eazynow/sms/segamastersystem"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"testing"
)

func BenchmarkRendering(b *testing.B) {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err.Error())
	}

	screen := smslib.NewSDL2xScreen(false)

	displayLoop := smslib.NewSDLLoop(screen)
	go displayLoop.Run()

	sms := smslib.NewSMS(displayLoop)

	sms.LoadROM("../roms/blockhead.sms")

	numOfGeneratedFrames := 100
	generatedFrames := make([]smslib.DisplayData, numOfGeneratedFrames)

	for i := 0; i < numOfGeneratedFrames; i++ {
		generatedFrames = append(generatedFrames, *sms.RenderFrame())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, frame := range generatedFrames {
			displayLoop.Display() <- &frame
		}
	}
}

func BenchmarkCPU(b *testing.B) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err.Error())
	}

	screen := smslib.NewSDL2xScreen(false)

	displayLoop := smslib.NewSDLLoop(screen)
	go displayLoop.Run()

	sms := smslib.NewSMS(displayLoop)

	sms.LoadROM("../roms/blockhead.sms")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sms.RenderFrame()
	}
}
