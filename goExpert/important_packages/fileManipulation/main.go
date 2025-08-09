package main

import (
	"bufio"
	"os"
)

func main() {
	/**
	* Write files
	 */
	// This file will be created on the root of the folder where the command go run is running
	f, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	// size, err := f.WriteString("meu arquivinho")
	size, err := f.Write([]byte("oioiouooi fdsafa fdafwef fadsfasd"))

	if err != nil {
		panic(err)
	}

	println(size)

	// Necessário sempre
	f.Close()

	fileRead, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	println(fileRead)
	// Need a cast or parse because on the file there is only bytes
	println(string(fileRead))

	println("=============================")

	// Read chunks

	file, err := os.Open("file.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	buffer := make([]byte, 3)

	for {
		chunk, err := reader.Read(buffer)

		if err != nil {
			break
		}

		println(string(buffer[:chunk]))
	}

	println("=============================")
	println(string(buffer))
	file.Close()

	err = os.Remove("file.txt")

	if err != nil {
		panic(err)
	}
}
