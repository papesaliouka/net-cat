package helpers

import (
	"fmt"
	"net"
)
var remoteAddrToName = map[string]string{}

var nameChannel = make(chan string,10)

func GetListener(port string) (net.Listener, bool) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Something went wrong : ", err)
		return nil, true
	}
	return listener, false
}

func HandleConnection(conn net.Conn, connections *[]net.Conn,i *int){
	ReadLogo(conn)
	GetName(conn)
	name := <-nameChannel
	remoteAddrToName[conn.RemoteAddr().String()]=name
	ReadHistory(conn)
	connexionNofication := fmt.Sprintf("%s has joined the chat",name)
	disconectNotication := fmt.Sprintf("%s has left the chat",name)
	Notify(connexionNofication,conn,*connections,name,remoteAddrToName)
	ReadMessage(conn,name,connections,remoteAddrToName)
	defer Notify(disconectNotication,conn,*connections,name,remoteAddrToName)
	*i--
}

func Notify(message string, sender net.Conn, connections []net.Conn, name string, remoteAddrToName map[string]string) {
	if len(message)==0||len(name)==0{
		return
	}
	for _, conn := range connections {
			if conn != sender {
					fmt.Fprintf(conn, "%s",message)
					remoteCursor:=fmt.Sprintf("[%s][%s]:", GetDate(), remoteAddrToName[conn.RemoteAddr().String()])
					fmt.Fprintf(conn,"\n%s",remoteCursor)
			}
	}
	WriteLog(message+"\n")
}