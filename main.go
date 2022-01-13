package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	fileName := "./README.md"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Open File Error", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(">>", line)
	// }

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(">>", line)
	}

}
