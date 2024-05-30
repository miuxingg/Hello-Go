package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchFn(task string, ch chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	for i := 0; i <= 1e5; i++ {
		// Do something
		fmt.Printf("%v: %v\n", task, i)
	}
	ch <- "\nThis task doneeee " + task
	wg.Done()
}

func forFn(count *int) {

}
func main() {
	nowTime := time.Now()
	ch := make(chan string)
	wg := new(sync.WaitGroup)
	go fetchFn("1", ch, wg)
	go fetchFn("2", ch, wg)
	go fetchFn("3", ch, wg)
	// time.Sleep(10 * time.Second)
	x, y, z := <-ch, <-ch, <-ch
	fmt.Println(x, y, z)
	close(ch)
	wg.Wait()

	fmt.Println("\nDone: ", time.Since(nowTime))

	// lock := new(sync.Mutex)
	// count := 0

	// for i := 1; i <= 5; i++ {
	// 	go func() {
	// 		for j := 1; j <= 5000; j++ {
	// 			lock.Lock()
	// 			count++
	// 			fmt.Println(count)
	// 			lock.Unlock()
	// 		}
	// 	}()
	// }
	// time.Sleep(time.Second * 2)
}
