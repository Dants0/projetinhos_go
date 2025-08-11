package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/schollz/progressbar/v3"
	"strings"
)

func DownloadFile(url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("falha ao baixar o arquivo: %s", resp.Status)
	}

	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer out.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		fmt.Sprintf("Baixando %s", filepath.Base(filename)),
	)

	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	return err
}

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		filename := filepath.Base(url)

		err := DownloadFile(url, filename)

		if err != nil {
			fmt.Printf("[Worker %d], erro: %v\n ", id, err)
		} else {
			fmt.Printf("[Worker %d], Download concluído: %s\n ", id, filename)
		}
	}

}

func readFiles(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer file.Close() // Ensure the file is closed

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

func main() {

	urls, err := readFiles("urls.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Urls %s", urls)

	const workerCount = 2
	jobs := make(chan string, len(urls))
	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	wg.Wait()

	fmt.Println("Todos os dowloads foram concluídos!")
}
