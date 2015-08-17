package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/inotify"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s <file path>\n", os.Args[0])
		os.Exit(1)
	}
	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Watch(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case ev := <-watcher.Event:
			log.Print(ev.String())
		case err := <-watcher.Error:
			log.Fatal(err)
		}
	}
}
