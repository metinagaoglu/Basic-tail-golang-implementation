package main

import(
	"io"
	"fmt"
	"bufio"
	"time"
	"os"
)

func main() {

	file, _ := os.Open("log.txt")
	followFile(file)

}

func followFile(file io.Reader) error {
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}

		fmt.Print(string(line))
		if err == io.EOF {
			time.Sleep(time.Second)
		}
	}
}