package main

import "github.com/gorilla/websocket"

// client represents single chatting client
type client struct {
	socket *websocket.Conn
	// send is a channel for sending message
	send chan []byte
	// room is the room for this client to chat
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}
