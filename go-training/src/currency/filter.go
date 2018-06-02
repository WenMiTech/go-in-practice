package main

import (
	"runtime"
	"log"
	"flag"
	"strings"
	"fmt"
	"path/filepath"
	"os"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	log.SetFlags(0)

	algorithm, minSize, maxSize, suffixes, files := handleCommandLine()

	if algorithm == 1 {
		sink((filterSize(minSize, maxSize, filterSuffixes(suffixes, source(files)))))
	} else {
		channel1 := source(files)
		channel2 := filterSuffixes(suffixes, channel1)
		channel3 := filterSize(minSize, maxSize, channel2)
		sink(channel3)
	}

}

func handleCommandLine() (algorithm int, minSize, maxSize int64, suffixes, files []string) {

	flag.IntVar(&algorithm, "algorithm", 1, "1 or 2")
	flag.Int64Var(&minSize, "min", -1, "min file size")
	flag.Int64Var(&maxSize, "max", -1, "file max size")

	// suffixes is a  point type var, it is actual an address,*suffixes is the content
	var suffixesOpt *string = flag.String("suffixes", "", "comma-separated list of file suffixes")

	flag.Parse()

	if algorithm != 1 && algorithm != 2 {
		algorithm = 1
	}
	if minSize > maxSize && maxSize != -1 {
		log.Fatalln("minSize must be < maxSize")
	}

	suffixes = []string{}

	if *suffixesOpt != "" {
		suffixes = strings.Split(*suffixesOpt, ",")
	}

	files = flag.Args()
	return algorithm, minSize, maxSize, suffixes, files
}

func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()

	return out
}

func sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {

		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}
			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}

		}

		close(out)
	}()

	return out
}

func filterSize(minSize, maxSize int64, in <-chan string) <-chan string {

	out := make(chan string, cap(in))

	go func() {
		for filename := range in {
			if minSize == -1 && maxSize == -1 {
				out <- filename // don't do a stat call it not needed
				continue
			}
			finfo, err := os.Stat(filename)
			if err != nil {
				continue // ignore files we can't process
			}
			size := finfo.Size()
			if (minSize == -1 || minSize > -1 && minSize <= size) &&
				(maxSize == -1 || maxSize > -1 && maxSize >= size) {
				out <- filename
			}
		}
		close(out)
	}()
	return out
}
