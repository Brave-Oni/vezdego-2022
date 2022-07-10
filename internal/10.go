package internal

import (
	"bufio"
	"log"
	"os"
	"vezdecode/pkg"
)

func main() {
	file, err := os.Open("example.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	id := 1

	for scanner.Scan() {
		task := pkg.NewTask(id, scanner.Text())
		id++
		task.Run()
	}
}
