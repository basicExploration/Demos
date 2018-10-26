package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("hello.go") //将读取的open传入newReader的读取流
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		inputString, readerError := reader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)//
		if readerError == io.EOF {
			return
		}
	}
}
