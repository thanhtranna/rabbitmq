package main

import (
	"fmt"

	"rabbitmq/buff"
)

func main() {
	// create a circular buffer of size 100 in recent mode (searches the buffer
	// from the most recent to oldest element)
	buffer, err := buff.Init(100, buff.Recent)
	if err != nil {
		// error will only occur if size is less than 1, or if incorrect mode is
		// provided
		panic(err)
	}

	data := []byte("hello")

	// check if data is in buffer
	fmt.Printf("%s in buffer :: %t\n", data, buffer.Test(data))

	testData := []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
	buffer.Add(testData)

	// add data to buffer
	buffer.Add(data)
	fmt.Printf("%s in buffer :: %t\n", data, buffer.Test(data))

	buffer.Add([]byte("hello2"))
	buffer.Add([]byte("hello3"))

	// get the most recent and oldest elements
	fmt.Printf("most recent :: %s\n", buffer.GetRecent())
	fmt.Printf("oldest :: %s\n", buffer.GetOldest())

	// reset buffer
	buffer.Reset()
	fmt.Printf("%s in buffer :: %t\n", data, buffer.Test(data))
}
