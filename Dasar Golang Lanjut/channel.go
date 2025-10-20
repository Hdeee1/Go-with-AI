package main

import "fmt"

func main() {
	// membuat channel yang hanya bisa menampung tipe data string   
	messages := make(chan string) 

	// goroutine ini mengirim pesan ke channel
	go func() {
		messages <- "Ping!" //Kirim data "Ping!" ke channel 
	}()

	msg := <-messages //menunggu dan menerima data dari channel 
	fmt.Println(msg) //output
}