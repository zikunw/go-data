# go-data

This Go package provides generic data types and fast serialization/deserialization method, without any code generation
or reflect library overhead.

## Features
1. Provides generic tuple types.
2. Provides batch data structure for the tuple types.
3. Provides fast runtime serialization/deserialization for both tuples and batches.

## Example Usage

Encoding/decoding tuple.
``` go
t := dt.T3(42, float32(42), "Hello world")
fmt.Println(t)

enc := dt.CreateEncoderFunc(&t)
buf := make([]byte, 100)
enc(&t, buf)

newTuple := dt.Tuple3[int, float32, string]{}
dec := dt.CreateDecoderFunc(&newTuple)
dec(&newTuple, buf)
fmt.Printf("newTuple: %v\n", newTuple) // Out: {42 42 "Hello world."}
```

Encoding/decoding batch.
``` go
batch := NewBatchOf(&Tuple3[int, int, string]{}, 5)
batch.Data[0] = &T3(0, 42, "Hello world")
batch.Data[1] = &T3(1, 42, "Hello world")
batch.Data[2] = &T3(2, 42, "Hello world")
batch.Data[3] = &T3(3, 42, "Hello world")
batch.Data[4] = &T3(4, 42, "Hello world")

buf := make([]byte, 500)
n1 := batch.Encode(buf)

target_batch := NewBatchOf(&Tuple3[int, int, string]{}, 5)
n2, err := target_batch.Decode(buf)
```

## Motivation

Recently, I’ve been working on a distributed framework that requires different services to pass around user-defined data types. Specifically, we want to flexibly define tuples (multiple values) and entries (key-value pairs) and efficiently serialize and deserialize them. Each key or value mentioned here is a primitive type available in Go.

Existing libraries and approaches have some limitations:

1. Many libraries provide generic structs like Tuple2[A, B], but there is no efficient way to serialize them.
2. Current serialization libraries rely on one of the following: manual handling, code generation, or the reflect library. Each of these approaches is either too labor-intensive or compromises performance.

Since we know the generic tuples and entries consist of primitive types, we can create encoder and decoder functions during initialization. This approach avoids the constant overhead of using the reflect library while achieving performance comparable to manual handling and code generation.
