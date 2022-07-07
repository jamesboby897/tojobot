package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

var buffer = make([][]byte, 0)

func main() {

	err := loadSound()
	if err != nil {
		fmt.Println("Error loading sound: ", err)
		return
	}

}
func loadSound() error {

	file, err := os.Open("test.dca")
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}
	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
	fmt.Println(buffer)
	return nil
}
