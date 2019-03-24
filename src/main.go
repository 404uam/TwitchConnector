package main

import (
	"./server"
	"path/filepath"
)

func main() {
	absPath, _ := filepath.Abs("./server/settings.json")
	server.LoadSettings(absPath)
	server.RunWebServer()
}
