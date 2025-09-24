package main

import (
	"fmt"
	"log"
	"net"

	"github.com/RedHat007-rgb/day-phase-1-24th/packages/golib/ws"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)
func BinanceConnection(url string) *websocket.Conn{
	 conn:=ws.ConnectToWs(url)
	return conn
}


func main(){
	lis,err:=net.Listen("tcp",":50051")
	if err!=nil{
		log.Fatalln("cannot establish connection",err)
	}
	// webby:=BinanceConnection("wss://stream.binance.com:9443/ws/btcusdc@ticker")
	// _,msg,err:=webby.ReadMessage()
	// if err!=nil{
	// 	log.Fatalln("reading error of messages",err)
	// }
	fmt.Println(string(msg))
	grpcServer:=grpc.NewServer()
	log.Println("go service running on 50051")
	if err:=grpcServer.Serve(lis);err!=nil{
		log.Fatalf("failed to serve %v",err)
	}
	

}
