//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)


func producer(stream Stream, tch chan<- *Tweet)  {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tch)
			return
		}
		tch <- tweet
	}
}

func consumer(tch <-chan *Tweet) {
	for t := range tch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	tch := make(chan *Tweet)
	// Producer
	go producer(stream, tch)
	consumer(tch)
	fmt.Printf("Process took %s\n", time.Since(start))
}
