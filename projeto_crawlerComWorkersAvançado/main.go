package main

import (
	"fmt"
	"sync"
)

var urls = []string {
	"https://golang.org",
	"https://noroteiro.com",
	"https://github.com",
	"https://google.com",
	"https://amazon.com.br",
	"https://go.dev",
	"https://vercel.com",
	"https://facebook.com",
	"https://youtube.com",
}


func main(){
	const workersCounts = 10
	jobs := make(chan string, len(urls))

	var wg sync.WaitGroup

	for i := 1; i<=workersCounts; i++{
		wg.Add(1)
		go Worker(i, jobs, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	wg.Wait()
	fmt.Println("Todas as tarefas foram concluídas.")
}