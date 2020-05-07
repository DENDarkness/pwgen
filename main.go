package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getChar(i int) []byte {
	var r []byte
	if i == 0 {
		for i := 33; i < 126; i++ {
			r = append(r, byte(i))
		}
	}
	if i == 1 {
		for i := 48; i < 57; i++ {
			r = append(r, byte(i))
		}
	}
	if i == 2 {
		for i := 48; i < 57; i++ {
			r = append(r, byte(i))
		}
		for i := 65; i < 90; i++ {
			r = append(r, byte(i))
		}
	}
	if i == 3 {
		for i := 48; i < 57; i++ {
			r = append(r, byte(i))
		}
		for i := 65; i < 90; i++ {
			r = append(r, byte(i))
		}
		for i := 97; i < 122; i++ {
			r = append(r, byte(i))
		}
	}
	return r
}

func genPass(i int, c []byte) string {
	rand.Seed(time.Now().UnixNano())
	var r []byte
	for a := 0; a < i; a++ {
		r = append(r, c[rand.Intn(len(c))])
	}
	return string(r)
}

func main() {

	fmt.Print(genPass(24, getChar(0)))
}
