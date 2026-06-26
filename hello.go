package main

import "fmt"

func main(){
	keterangan :="lulus"
	nilai:= 100

	if nilai<=10 {
		keterangan = "tidak lulus"
	}
	fmt.Println(keterangan + " yeyy!!")
}