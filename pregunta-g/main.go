package main

import (
	"fmt"
	"sync"
)

type Message struct {
	Text string
}

type Mailbox struct {
	ch chan Message
}

func (m *Mailbox) Send(msg Message, wg *sync.WaitGroup) {
	defer wg.Done()
	m.ch <- msg
	fmt.Println("enviado")
}

func main() {
	box := Mailbox{ch: make(chan Message)}
	var wg sync.WaitGroup
	wg.Add(1)
	go box.Send(Message{Text: "hola"}, &wg)
	msg := <-box.ch
	fmt.Println(msg.Text)
	wg.Wait()
}
