package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{
    ReadBufferSize:  10240, //1024 = 1kb
	WriteBufferSize: 10240,
    CheckOrigin:func(r *http.Request) bool {
        return true
    },
}

type Message struct {
    Message string `json:"message"`
}

func main() {
    fs := http.FileServer(http.Dir("public"))
    http.Handle("/", fs)

    http.HandleFunc("/ws", handleConnections)

    go handleMessages()

    log.Println("http server started on :8000")
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

//注册成为 websocket
func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()
    clients[ws] = true

    //不断的从页面上获取数据 然后广播发送出去
    for {
        //将从页面上接收数据改为不接收 直接发送
        //var msg Message
        //err := ws.ReadJSON(&msg)
        //if err != nil {
        //  log.Printf("error: %v", err)
        //  delete(clients, ws)
        //  break
        //}

    }
}

//广播发送至页面
func handleMessages() {
    for {
        time.Sleep(time.Second * 3)
        msg := Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
        
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("client.WriteJSON error: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}