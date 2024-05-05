package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	now := time.Now()
	userID := 10
	resChan := make(chan string, 3)

	wg := &sync.WaitGroup{}

	go fetchUserData(userID, resChan, wg)
	go fetchUserRecommend(userID, resChan, wg)
	go fetchUserlikes(userID, resChan, wg)
	wg.Add(3)
	wg.Wait()
	close(resChan)
	for resp := range resChan {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(userID int, resChan chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	resChan <- "Data"
	wg.Done()
}

func fetchUserRecommend(userID int, resChan chan string, wg *sync.WaitGroup) {
	time.Sleep(120)
	resChan <- "Recommendation"
	wg.Done()
}

func fetchUserlikes(userID int, resChan chan string, wg *sync.WaitGroup){
	time.Sleep(50)
	resChan <- "Likes"
	wg.Done()
}

