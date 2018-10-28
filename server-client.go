//测试 net包

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listenner, err := net.Listen("tcp", "localhost:5000") //net包监听5000端口
	if err != nil {
		log.Fatal("监听端口的时候发生了错误，错误发生在mai.go中的main函数中，代码是:", err)
	}
	for { // 对于每一个输入进行监听
		conn, err := listenner.Accept()
		if err != nil {
			log.Fatal("accept函数发生错误main.go main", err)
		}
		go serverHandle(conn)

	}

}

func serverHandle(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal("在conn读取buf的时候发生了错误 main.go serverHandle", err)
		}
		fmt.Println(string(buf[:n]))
	}

}




package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Fatal("net.Dial client.go main中发生错误：", err)
	}
	reader := bufio.NewReader(os.Stdin)
	ClientName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("readString发生错误", err)
	}
	Client := strings.Trim(ClientName, "\n")
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := reader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		// fmt.Printf("input:--%s--", input)
		// fmt.Printf("trimmedInput:--%s--", trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(Client + " says: " + trimmedInput))
	}

}
