# go-data

This Go package provides generic data types and fast serialization/deserialization method, without any code generation
or reflect library overhead.

## Usage

``` go
package main

import (
	"fmt"

  dt "github.com/zikunw/go-data"
)

func main() {
	t := dt.Tuple3[int, float32, string]{V1: 42, V2: 42, V3: "Hello world."}
	fmt.Println(t)

	enc := dt.CreateEncoderFunc(&t)
	buf := make([]byte, 100)
	enc(&t, buf)

	newTuple := dt.Tuple3[int, float32, string]{}
	dec := dt.CreateDecoderFunc(&newTuple)
	dec(&newTuple, buf)
	fmt.Printf("newTuple: %v\n", newTuple) // Out: {42 42 "Hello world."}
}
```

## Motivation

Recently, I’ve been working on a distributed framework that requires different services to pass around user-defined data types. Specifically, we want to flexibly define tuples (multiple values) and entries (key-value pairs) and efficiently serialize and deserialize them. Each key or value mentioned here is a primitive type available in Go.

Existing libraries and approaches have some limitations:

1. Many libraries provide generic structs like Tuple2[A, B], but there is no efficient way to serialize them.
2. Current serialization libraries rely on one of the following: manual handling, code generation, or the reflect library. Each of these approaches is either too labor-intensive or compromises performance.

Since we know the generic tuples and entries consist of primitive types, we can create encoder and decoder functions during initialization. This approach avoids the constant overhead of using the reflect library while achieving performance comparable to manual handling and code generation.