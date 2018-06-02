package main

import (
	"sync"
	"io"
)

func main() {

}

type Pool struct {
	m sync.Mutex
	resource chan io.Closer
	factory func()(io.Closer,error)
	closed bool
}