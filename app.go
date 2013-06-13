package main

import (
	"github.com/howeyc/fsnotify"
	"log"
)

func main() {
	// fmt.Println("Hello World!")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event: ", ev)
			}
		}
	}()
	err = watcher.Watch("testDir")
	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}
