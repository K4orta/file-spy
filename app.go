package main

import (
	"github.com/howeyc/fsnotify"
	"io"
	"log"
	"os"
	"path"
)

func main() {
	watchDir := os.Args[1]
	moveDir := os.Args[2]
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	log.Println("File Spy is running...")
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsCreate() || ev.IsModify() {
					log.Println("Moving a file...")
					_, filename := path.Split(ev.Name)
					efv := copyFile(ev.Name, moveDir+"/"+filename)
					if efv != nil {
						log.Println(efv)
					}
				}

			}
		}
	}()
	err = watcher.Watch(watchDir)
	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}

func copyFile(filein string, fileout string) (err error) {
	src, err := os.Open(filein)
	defer src.Close()
	dest, err := os.Create(fileout)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, src)
	if err == nil {
		si, err := os.Stat(filein)
		if err != nil {
			err = os.Chmod(fileout, si.Mode())
		}
	}

	return
}
