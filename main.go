package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(">>", line)
	}

}
