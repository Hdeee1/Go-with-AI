package main

import (
	"fmt"
	"time"
)

func say(text string) {
	for i := 0; i < 3; i++ {
		fmt.Println(text)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// menjalankan fungsi say sebagai Goroutime (async)
	go say("Hello world, I'm Goroutine!")
	
	// menjalankan fungsi say secara normal
	say("Hello world, I'm only func")
	
	// beri waktu agar goroutine sempat berjalan sebelum program selesai
	time.Sleep(500 * time.Millisecond)
}