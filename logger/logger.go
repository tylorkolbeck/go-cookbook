package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Log() {
	content := "Hello from Go!"

	file, err := os.Create("./logs/log.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Wrote to log with %v characters\n", length)
	defer file.Close()
	defer readFile("./logs/log.txt")

}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file: ", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
