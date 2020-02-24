package main

import (
	"fmt"
)

func main()  {
	ch := make(chan int)
	go genearate(ch)
	for{
		prime := <- ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func genearate(ch chan int)  {
	for i := 2;; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int)  {
	i := <- in
	for{
		if i%prime != 0 {
			out <- i 
		}
	}
}