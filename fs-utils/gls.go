package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

func traverse(dirs []os.DirEntry, basePath string, output chan string, wg *sync.WaitGroup, limit int) {
	defer wg.Done()
	if limit == 0 {
		return
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			output <- dir.Name()
		} else {
			output <- dir.Name() + "/"
			path := basePath + "/" + dir.Name()
			dirs, _ := os.ReadDir(path)
			wg.Add(1)
			go traverse(dirs, path, output, wg, limit-1)
		}
	}
}
func main() {
	limit := flag.Int("depth", 200000, "Defines the depth of recursive search to perform")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: program [path] --depth=N")
		os.Exit(1)
	}
	path := flag.Args()[0]
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	output := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go traverse(dirs, path, output, &wg, *limit)
	go func() {
		wg.Wait()
		close(output)
	}()
	for p := range output {
		print(p, "\t")
	}

}
