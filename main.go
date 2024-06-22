package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Preparing quiz...")
	f := openFile()
	defer f.Close()
	records := readContent(f)

	fmt.Println("Ready")
	fmt.Print("Enter any key to start...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for _, r := range records {
		fmt.Println(r)
		time.Sleep(time.Second)
	}

}

func readContent(f *os.File) [][]string {
	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error reading records: ", err)
	}
	return records
}

func openFile() *os.File {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln(err)
	}
	return f
}
