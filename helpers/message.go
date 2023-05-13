package helpers

import (
	"fmt"
	"net"
	"strings"
	"bufio"
)

func  ReadMessage(conn net.Conn,name string, connections *[]net.Conn , remoteAddrToName map[string]string) {
	if len(name)==0{
		return
	}
	for {
		cursor := fmt.Sprintf("[%s][%s]:",GetDate(),name)
		fmt.Fprint(conn,cursor)

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		broadcastMessage(msg,conn,*connections,name,remoteAddrToName)
	}
}

func broadcastMessage(message string, sender net.Conn, connections []net.Conn, name string, remoteAddrToName map[string]string) {
	if len(message)==0{
		return
	}
	var cursor string
	for _, conn := range connections {
			if conn != sender {
					cursor = fmt.Sprintf("\n[%s][%s]:", GetDate(), name)
					fmt.Fprintf(conn, "%s%s", cursor, message)
					remoteCursor:=fmt.Sprintf("[%s][%s]:", GetDate(), remoteAddrToName[conn.RemoteAddr().String()])
					fmt.Fprintf(conn,"\n%s",remoteCursor)
			}
	}
	WriteHistory(cursor + message)
	WriteLog(cursor+message+"\n")
}
