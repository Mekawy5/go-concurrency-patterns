package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	data := displayData(processData(generateData()))

	for str := range data {
		fmt.Println(str)
	}
}

func generateData() <-chan int64 {
	const fp = "integers.txt"
	out := make(chan int64)

	go func() {
		f, _ := os.Open(fp)
		defer close(out)
		defer f.Close()

		r := bufio.NewReader(f)

		for {
			line, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}

			integer, _ := strconv.ParseInt(string(line), 10, 0)
			out <- integer
		}
	}()

	return out
}

func processData(c <-chan int64) <-chan int64 {
	oc := make(chan int64)
	go func() {
		for n := range c {
			input := n * n
			oc <- input
		}
		close(oc)
	}()
	return oc
}

func displayData(c <-chan int64) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)

		for n := range c {
			str := strconv.Itoa(int(n))

			out <- fmt.Sprintf("Current number is: %s \n", str)
		}
	}()

	return out
}
