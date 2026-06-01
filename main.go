package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{} // zero memory allocation
}

func (s *Server) StartAndListen() {
	// you can name your for loop
running:
	for {
		//block until someone is sending a message to a channel
		select {
		case msg := <-s.msgch:
			fmt.Println("received message from", msg)
		case <-s.quitch:
			fmt.Println("the server is doing graceful shutdown")
			// logic for the graceful shutdown
			break running
		default:
		}
	}
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Joe Biden",
		Payload: payload,
	}
	msgch <- msg
}

func graceFullQuitServer(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.StartAndListen()
	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "Hello world")
	}()
	go func() {
		time.Sleep(4 * time.Second)
		graceFullQuitServer(s.quitch)
	}()
	select {}
}

// func main() {
// 	now := time.Now()
// 	userID := 10
// 	bufferForSameTimeWrite := 5
// 	respch := make(chan string, bufferForSameTimeWrite)
// 	wg := &sync.WaitGroup{}

// 	go fetchUserData(userID, respch, wg)
// 	wg.Add(1)
// 	go fetchUserRecommendations(userID, respch, wg)
// 	wg.Add(1)
// 	go fetchUserLikes(userID, respch, wg)
// 	wg.Add(1)
// 	wg.Wait()

// 	close(respch)
// 	for resp := range respch {
// 		fmt.Println(resp)
// 	}

// 	fmt.Println(time.Since(now))
// }

// func fetchUserData(userID int, respsch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)
// 	respsch <- "user data"
// 	wg.Done()
// }

// func fetchUserRecommendations(userID int, respsch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)
// 	respsch <- "user recommendations"
// 	wg.Done()
// }

// func fetchUserLikes(userID int, respsch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(50 * time.Millisecond)
// 	respsch <- "user likes"
// 	wg.Done()
// }
