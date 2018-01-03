package remotecollaber

import (
	"fmt"
	"log"
	"os"

	//"syscall"

	"github.com/MarinX/keylogger"
	//"github.com/faiface/beep/wav"
)

type scratch struct {
}

// TODO look at using "/usr/bin/sndiod -a on
func readSource(dp string) *bytes.Buffer {
	byteSlice, err := os.Open(dp)
	if err != nil {
		log.Fatal("Device not found or loaded correctly")
	}
	sourceBuf := new(bytes.Buffer(byteSlice))
	sourceChan := make(chan []bytes)
	sourceChan <- sourceBuf
	return sourceBuf
}

func rawInput(input string) *bytes.Buffer {
	device, err := os.Open(input)
	stream := bytes.NewBuffer([]bytes{device})
	return stream
}

//Asks user to enter dev number from given list.
func getDevNum() int {
	devices, err := keylogger.NewDevices()
	if err != nil {
		fmt.Println("error= %s", err)
	}
	for _, devs := range devices {
		fmt.Println(devs)
	}
	var devNum int
	fmt.Println("Enter the number of your recording device:")
	fmt.Scanln("%s", &devNum)
	return devNum

}
