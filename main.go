package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/MarkKremer/microphone"
	"github.com/gopxl/beep/wav"
	"github.com/gordonklaus/portaudio"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Feed the device.")
		return
	}

	fmt.Println("Recording. Press Ctrl-C to stop.")

	portaudio.Initialize()
	defer portaudio.Terminate()

	stream, format, err := microphone.OpenDefaultStream(44100, 1)
	if err != nil {
		log.Fatal(err)
	}

	filename := os.Args[1]

	if !strings.HasSuffix(filename, ".wav") {
		filename += ".wav"
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	go func() {
		<-sig
		stream.Stop()
		stream.Close()
	}()

	stream.Start()
	defer stream.Close()

	var ti []byte
	go wav.Encode(ti, stream, format)

	time.Sleep(10 * time.Second)

}
