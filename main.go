package main

import (
	"fmt"
	"net"
	"io/ioutil"
	"os"
	"bufio"
)

func main() {
	scores := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	initials := [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	nums, err := ioutil.ReadFile("/tmp/dat")
	if err != nil {
                panic(err)
        }

	fmt.Println(string(nums))

	f, err := os.Open("scores")
	if err != nil {
		panic(err)
	}

	

	net.Listen("tcp", ":9156")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
		
	}
}
