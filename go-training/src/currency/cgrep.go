package main

import (
	"runtime"
	"os"
	"fmt"
	"path/filepath"
	"regexp"
	"log"
	"bufio"
	"bytes"
	"io"
)

type Job struct {
	filename string
	results  chan<- Result
}

func (job Job) Do(lineRx *regexp.Regexp) {
	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error:%s\n", err)
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := reader.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n\r")
		if lineRx.Match(line) {

			job.results <- Result{job.filename, lino, string(line)}
		}
		if err != nil {

			if err != io.EOF {
				log.Printf("error:%d:%s\n", lino, err)
			}
			break

		}
	}
}

type Result struct {
	filename string
	lino     int
	line     string
}

var workers int = runtime.NumCPU()

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("usage : %s <regexp> <file> \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatalf("invaild regexp:%s\n", err)
	} else {
		grep(lineRx, commandLineFiles(os.Args[2:]))
	}
}


//regexp的实现线程安全的,可以在多个goroutine中共享
func grep(lineRx *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, minimum(1000, len(filenames)))
	done := make(chan struct{}, workers)

	go addJobs(jobs, filenames, results)
	for i := 0; i < workers; i++ {  //启动和CPU核心数一样多的goroutine处理job
		go doJobs(done, lineRx, jobs)
	}

	go awaitCompletion(done, results) // Executes in its own goroutine
	processResults(results)           // Blocks until the work is done

}


//往jobs通道中增加工作
func addJobs(jobs chan<- Job, filenames []string, results chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job{filename, results}
	}
	close(jobs)
}

func doJobs(done chan<- struct{}, lineRx *regexp.Regexp, jobs <-chan Job) {
	for job := range jobs {
		job.Do(lineRx)
	}
	done <- struct{}{}
}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

func minimum(x int, ys ...int) int {
	for _, y := range ys {
		if y < x {
			x = y
		}
	}

	return x
}

func awaitCompletion(done <-chan struct{}, results chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(results)
}

func processResults(results <-chan Result) {
	for result := range results {
		fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
	}
}