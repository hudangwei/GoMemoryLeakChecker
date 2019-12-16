package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"
)

//go:noinline
func CallTest() {
	fmt.Println("haha")
}

//go:noinline
func GenHeapProfile() {
	f, err := os.Create("1.hprof")
	if err != nil {
		return
	}
	p := pprof.Lookup("heap")
	p.WriteTo(f, 0)
	f.Close()
}

func main() {
	CallTest()
	// GenHeapProfile()
	Run(nil)
}

func Run(fn func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s\n", s.String())
		if fn != nil {
			fn()
		}
		time.Sleep(500 * time.Millisecond)
		return
	}
}
