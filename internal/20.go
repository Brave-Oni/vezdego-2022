package internal

//
//import (
//	"bufio"
//	"fmt"
//	"log"
//	"os"
//	"sync"
//	"time"
//	"vezdecode/pkg"
//)
//
//func (t Task) Run(wg *sync.WaitGroup) {
//	defer wg.Done()
//	fmt.Printf("Task %d is running\n", t.id)
//	time.Sleep(t.duration)
//	fmt.Printf("Task %d done !\n", t.id)
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	file, err := os.Open("example.txt")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer func(file *os.File) {
//		err := file.Close()
//		if err != nil {
//
//		}
//	}(file)
//
//	scanner := bufio.NewScanner(file)
//	scanner.Split(bufio.ScanWords)
//
//	id := 1
//
//	for scanner.Scan() {
//		task := pkg.NewTask(id, scanner.Text())
//		id++
//		wg.Add(1)
//		go task.Run(&wg)
//	}
//
//	wg.Wait()
//}
