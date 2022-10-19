package main

import(
	"io"
	"fmt"
	"bufio"
	"os"
	"syscall"
	"bytes"
	"encoding/binary"
)

func main() {
	followFile("log.txt")
}

func followFile(filename string) error {
	file, _ := os.Open(filename)

	inotifyFd, _ := syscall.InotifyInit()
	syscall.InotifyAddWatch(inotifyFd, filename, syscall.IN_MODIFY)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}

		fmt.Print(string(line))
		if err == io.EOF {
			continue
		}

		if err = waitForChange(inotifyFd); err != nil {
			return err
		}
	}
}

// file descriptor
func waitForChange(fd int) error {
	for {
		var buf [syscall.SizeofInotifyEvent]byte
		_ , err := syscall.Read(fd, buf[:])
		if err != nil {
			return err
		}

		read := bytes.NewReader(buf[:])
		var event = syscall.InotifyEvent{}

		binary.Read(read, binary.LittleEndian, &event)
		if event.Mask&syscall.IN_MODIFY == syscall.IN_MODIFY {
			return nil
		}
	}
}