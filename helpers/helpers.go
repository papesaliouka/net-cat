package helpers

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"os"
	"strings"
)

func GetPort() string {
	port := "8989"
	args := os.Args[1:]
	if len(args) == 0 {
		port = "8989"
	} else if len(args) == 1 {
		port = args[0]
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}
	return port
}

func GetName(conn net.Conn)  {
	fmt.Fprintf(conn, "[ENTER YOUR NAME]:")
	name, err := bufio.NewReader(conn).ReadString('\n')


	if err !=nil{
		nameChannel <-"anonymous"
	}
	name = strings.Trim(name, "\r\n")

	nameChannel <-name

}

func GetDate() string {
	time := time.Now()
	formatedTime := time.Format("2006-01-02 15:04:05")
	return formatedTime
}
