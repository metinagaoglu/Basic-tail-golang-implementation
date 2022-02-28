package main

import(
	"os"
)

func main() {

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	text := "[Info] - Some new log\n"

	if _, err = file.WriteString(text); err != nil {
		panic(err)
	}
}
