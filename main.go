package main

import (
	"fmt"
)

func main() {
	t := Tuple3[int, float32, string]{V1: 42, V2: 42, V3: "Hello world."}
	fmt.Println(t)

	enc := CreateEncoderFunc(&t)
	buf := make([]byte, 100)
	n := enc(&t, buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("buf: %v\n", buf)

	newTuple := Tuple3[int, float32, string]{}
	dec := CreateDecoderFunc(&t)
	n, err := dec(&newTuple, buf)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("newTuple: %v\n", newTuple)
}
