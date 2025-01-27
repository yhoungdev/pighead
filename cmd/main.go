package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	dirToWatch := "./"
	err = watcher.Add(dirToWatch)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Watching for file changes...")

	runApp()

	for {
		select {
		case event := <-watcher.Events:
			fmt.Println("Event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("Modified file:", event.Name)

				runApp()
			}
		case err := <-watcher.Errors:
			fmt.Println("Error:", err)
		}
	}
}

func runApp() {
	fmt.Println("Restarting application...")
	cmd := exec.Command("go", "run", "./cmd/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running app:", err)
	}
}
