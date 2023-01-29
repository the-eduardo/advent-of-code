package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("bgvyzdsv")
	fmt.Printf("%x", md5.Sum(data))
	for i := 0; i < 100000000; i++ {
		data = []byte(fmt.Sprintf("bgvyzdsv%d", i))
		if fmt.Sprintf("%x", md5.Sum(data))[:6] == "000000" {
			fmt.Printf("\nbgvyzdsv%d", i)
			// show the md5 hash
			fmt.Printf("\n%x", md5.Sum(data))
			break
		}
	}
}
