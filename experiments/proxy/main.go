package main

import (
	"fmt"
	"log"
	"os"
)

type myReader struct{}

func (myReader *myReader) Read(b []byte) (int, error) {
	fmt.Println("in > ")
	return os.Stdin.Read(b)
}

type myWriter struct{}

func (myWriter *myWriter) Write(b []byte) (int, error) {
	fmt.Println("out < ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader myReader
		writer myWriter
	)
	input := make([]byte, 4096)
	s, err := reader.Read(input)
	checkError(err)

	fmt.Printf("\nRead %d bytes from stdin\n", s)

	s, err = writer.Write(input)
	checkError(err)

	fmt.Printf("\nWrite %d bytes to stdout\n", s)

}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
