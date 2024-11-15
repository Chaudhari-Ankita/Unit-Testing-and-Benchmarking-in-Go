# Unit-Testing-and-Benchmarking-in-Go

In-class Lab Assignment 4

## Simple Calculator

This is a simple calculator file performing two major options: division and square root calculation.

## How to run the program

To run the program, you can use the following command in your terminal:

```go
go run main.go
```

## How to test the programs

To test the programs, you can use the following command in your terminal:

```go
go test -v
```

## How to run the benchmark program

To run the benchmark program, you can use the following command in your terminal:

```go
go test -bench=.
```

## Division Operation

### Description

Returns the quotient of `a` divided by `b`. Throws an error if `b` is 0.

### Error Handling

In the testing file, we handle edge cases such as division by zero, which returns nil. Additionally, if the error message does not match the expected message, an error is thrown.

## Square Root Operation
