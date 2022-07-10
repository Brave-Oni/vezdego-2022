package pkg

import (
	"fmt"
	"sync"
	"time"
)

//import "sync"
//
//type Listener interface {
//	Run(id int)
//	Run(id int, wg *sync.WaitGroup)
//}
//

type Task struct {
	id       int
	Duration time.Duration
}

func NewTask(id int, duration string) (t Task) {
	timeDuration, err := time.ParseDuration(duration)

	if err != nil {
		panic("Incorrect duration")
	}

	return Task{
		id:       id,
		Duration: timeDuration,
	}
}

func (t Task) Run(wg ...*sync.WaitGroup) {
	fmt.Printf("Task %d is running\n", t.id)
	time.Sleep(t.Duration)
	fmt.Printf("Task %d done !\n", t.id)
}
