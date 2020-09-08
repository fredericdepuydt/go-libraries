package tcp

import "net"
import "fmt"
import "bufio"
import "strings"

func handleConnection(conn net.Conn){
	message, _ := bufio.NewReader(conn).ReadString('\n')
    // output message received
    fmt.Print("Message Received:", string(message))
    // sample process for string received
    newmessage := strings.ToUpper(message)
    // send new string back to client
    conn.Write([]byte(newmessage + "\n"))
}