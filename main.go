package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // reader
	fmt.Print("input: ")
	input, _ := reader.ReadString('\n')   // get input
	input = strings.TrimSpace(input)      // remove \n from string
	arrInput := strings.Split(input, ",") // split string to array

	res := 1 // initiate res as one
	for _, str := range arrInput {
		int, _ := strconv.Atoi(str)
		res = res * int
	}

	fmt.Println(res)
}
