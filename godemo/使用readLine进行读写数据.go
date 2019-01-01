//读写数据，记住，写入数据的时候要flush一下。不然不会被写入的。
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	testI := 0
	inputFile, _ := os.Open("./test.txt")
	outputFile, _ := os.OpenFile("./储存.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		outputString := string(inputString) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
		testI++
	}
	outputWriter.Flush()
	fmt.Println("Conversion done", testI)
}
