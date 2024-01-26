package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	records, err := readData("file.csv")
	if err != nil {
		log.Fatalf("Could not read csv %v", err)
	}

	output := display(concatenate(swapValues(records)))

	for val := range output {
		fmt.Println(val)
	}
}

func readData(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	out := make(chan []string)

	go func() {
		cr := csv.NewReader(f)
		cr.FieldsPerRecord = 3

		for {
			record, err := cr.Read()
			if err == io.EOF {
				close(out)
				break
			}

			out <- record
		}
	}()

	return out, nil
}

func swapValues(data <-chan []string) <-chan []string {
	oc := make(chan []string)

	go func(data <-chan []string) {
		defer close(oc)

		for row := range data {
			row[0], row[1], row[2] = row[1], row[2], row[0]
			oc <- row
		}
	}(data)

	return oc
}

func concatenate(data <-chan []string) <-chan []string {
	out := make(chan []string)

	go func(data <-chan []string) {
		defer close(out)
		for arr := range data {
			for i := 0; i < len(arr); i++ {
				arr[i] = "9" + arr[i] + "0"
			}
			out <- arr
		}
	}(data)

	return out
}

func display(data <-chan []string) <-chan string {
	out := make(chan string)

	go func(data <-chan []string) {
		defer close(out)
		for arr := range data {
			for i := 0; i < len(arr); i++ {
				out <- arr[i]
			}
		}
	}(data)

	return out
}
