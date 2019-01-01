package main

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)
var sy sync.WaitGroup
func main() {
	data := []string{"9090", "9091", "9092"}
	for _, v := range data {
		go Wall(v)
	}
time.Sleep(time.Hour*1000000)
}
func Wall(port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
