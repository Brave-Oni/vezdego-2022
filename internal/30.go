package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"vezdecode/pkg"
)

func main() {
	var wg sync.WaitGroup

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter number of goroutines: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	value, err := strconv.Atoi(text)

	if err != nil {
		log.Fatal(err)
	}

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

	channel := make(chan int, value)

	for scanner.Scan() {
		task := pkg.NewTask(id, scanner.Text())
		id++
		wg.Add(1)
		channel <- 1
		go task.Run(&wg)
	}

	wg.Wait()
}
