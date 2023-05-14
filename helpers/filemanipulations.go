package helpers

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func WriteLog(data string) {
	file, err := os.OpenFile("db/log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Print(err)
	}
}

func ReadLogo(conn net.Conn) {
	file, err := os.Open("db/ubuntu.txt")
	if err != nil {
		fmt.Println("Error reading the ubuntu text file", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintln(conn, string(scanner.Text()))
	}
}

func ReadHistory(conn net.Conn) {
	historychat, err := os.Open("db/chathistory.txt")
	if err != nil {
		fmt.Println("Error reading the ubuntu text historychat", err)
	}
	defer historychat.Close()
	scanner1 := bufio.NewScanner(historychat)
	for scanner1.Scan() {
		line := scanner1.Text()
		if len(line) == 0{
			continue
		}
		fmt.Fprintln(conn, line)
	}
}

func WriteHistory(data string) {
	file, err := os.OpenFile("db/chathistory.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Print(err)
	}

}


