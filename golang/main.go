package main

import "time"

func main(){
	t := time.Tick(time.Second)
	select {
	case <- t:
		fmt.Println("1 second elapsed")
	}
}