package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
	"vezdecode/pkg"
)

var mutex = &sync.Mutex{}
var listener = NewListener()
var id = 0

type Listener struct {
	queue       []pkg.Task
	currentTask time.Time
}

func NewListener() Listener {
	return Listener{
		queue: []pkg.Task{},
	}
}

func addTask(timeDuration string, listener *Listener) {
	mutex.Lock()
	listener.queue = append(listener.queue, pkg.NewTask(id, timeDuration))
	mutex.Unlock()
	id++
}

func runTask(task pkg.Task) {
	fmt.Println("Task running")
	time.Sleep(task.Duration)
	fmt.Println("Task is done")
}

func takeTask(listener *Listener) {
	mutex.Lock()
	task := listener.queue[0]
	listener.queue = listener.queue[1:]
	mutex.Unlock()
	listener.currentTask = time.Now()
	runTask(task)
}

func getDuration(listener *Listener) string {
	var sum float64 = 0

	for _, t := range listener.queue {
		sum += t.Duration.Seconds()
	}

	return strconv.FormatFloat(sum, 'f', 6, 64)
}

func getMod(listener Listener) string {
	var sum float64 = 0

	for _, t := range listener.queue {
		sum += t.Duration.Seconds()
	}

	sum += time.Since(listener.currentTask).Seconds()

	return strconv.FormatFloat(sum, 'f', 6, 64)
}

func runServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", root)
	mux.HandleFunc("/add", add)
	mux.HandleFunc("/schedule", shedule)
	mux.HandleFunc("/time", timeDuration)

	log.Fatal(http.ListenAndServe(":5000", mux))
}

func runListener(listener *Listener) {
	for {
		if len(listener.queue) > 0 {
			takeTask(listener)
		}
	}
}

func root(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Привет от Небинарного дерева!"))
}

func add(w http.ResponseWriter, r *http.Request) {
	timeDur := r.URL.Query().Get("duration")
	addTask(timeDur, &listener)
	w.Write([]byte(timeDur))
}

func shedule(w http.ResponseWriter, _ *http.Request) {
	duration := getDuration(&listener)
	w.Write([]byte(duration))
}

func timeDuration(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(getMod(listener)))
}

func ddos(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		_, err := http.Post("http://localhost:5000/add?duration=0h0m20s", "text/plain", nil)

		if err != nil {
			log.Fatal(err)
		}
	}

	resp, _ := http.Get("http://localhost:5000/time")

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	fmt.Println("")

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("=== Start serving ===")
	wg.Add(1)
	go runServer()
	wg.Add(1)
	go runListener(&listener)

	ddosArr := [5]func(wg *sync.WaitGroup){ddos, ddos, ddos, ddos, ddos}

	for _, t := range ddosArr {
		wg.Add(1)
		go t(&wg)
	}

	wg.Wait()
}
