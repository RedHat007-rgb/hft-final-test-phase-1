package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)


func ConnectToWs(WsUrl string) *websocket.Conn{
	connection,_,err:=websocket.DefaultDialer.Dial(WsUrl,nil)
	if err!=nil{
		fmt.Println("webscoket dialing error",err)
	}
	return connection
}