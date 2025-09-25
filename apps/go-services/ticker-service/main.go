package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/RedHat007-rgb/day-phase-1-24th/packages/golib/redis"
	"github.com/RedHat007-rgb/day-phase-1-24th/packages/golib/ws"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)
func BinanceConnection(url string) *websocket.Conn{
	 conn:=ws.ConnectToWs(url)
	return conn
}




func main(){
	ctx:=context.Background()
	lis,err:=net.Listen("tcp",":50051")
	if err!=nil{
		log.Fatalln("cannot establish connection",err)
	}
	webby:=BinanceConnection("wss://stream.binance.com:9443/ws/btcusdc@ticker")
	_,msg,err:=webby.ReadMessage()
	if err!=nil{
		log.Fatalln("reading error of messages",err)
	}
	fmt.Println(string(msg))



	rdb:=redis.NewConnection()
	rdb.PublishToRedis(ctx,"message.btcusdt.binance",msg)
	
	
	 count,err:=rdb.AddSymbol(ctx,"active_symbols:btcusdt","userid1");
	 	if err!=nil{
			fmt.Println("error while adding.... %v",err)
		}else{
			log.Println("%d",count)
		}
	grpcServer:=grpc.NewServer()
	log.Println("go service running on 50051")
	if err:=grpcServer.Serve(lis);err!=nil{
		log.Fatalf("failed to serve %v",err)
	}
	

}
