package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Worker(id int, jobs <-chan string, wg *sync.WaitGroup){
	defer wg.Done()

	for url := range jobs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("[Worker %d] Erro: %s\n", id, err)
			continue
		}
		fmt.Printf("[Worker %d] %s -> %d\n", id, url, resp.StatusCode)
		resp.Body.Close()
		time.Sleep(500*time.Millisecond)
	}
}