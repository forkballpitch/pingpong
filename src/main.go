package src

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main(){

	table := make(chan *Ball)
    go player("ping", table)
	go player("pong", table)
	table<-new(Ball)
	time.Sleep(1*time.Second)
	<-table
	panic("shoe me the stucks")
}


func player(name string , table chan *Ball){

	for{
		Ball :=<-table
		Ball.hits++
		fmt.Println(name,Ball.hits)
		time.Sleep(100 *time.Millisecond)
		table<-Ball

	}

}