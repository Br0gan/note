package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var notesPath string = os.Getenv("NOTES_PATH")

func main() {
	if len(os.Args) == 1 {
		fmt.Println("wrong use of note")
		helpMenu()
		os.Exit(1)
	}

	checkFlags()
}

func helpMenu() {
	fmt.Println("this is the help menu")
}

func checkFlags() {
	flag := os.Args[1]
	switch flag {
	case "--help":
		helpMenu()
		os.Exit(0)
	case "--find":
		searchNotes()
	default:
		noteHandler()
	}
}

func noteHandler() {
	if err := os.MkdirAll(notesPath+"/"+os.Args[1], os.ModePerm); err != nil {
		fmt.Println("script failed: ", err)
	}

	filePath := notesPath + "/" + os.Args[1] + "/"
	cmd := exec.Command("vim", filePath+os.Args[2])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("script failed: ", err)
	}
}

func checkIfHelp(f string) {
	if f == "--help" {
		helpMenu()
		os.Exit(0)
	}
}

func checkIfFind(f string) {
	if f == "--find" {
		searchNotes()
	}
}

func searchNotes() {
	ns := []note{}
	out, err := exec.Command("fgrep", "-rHni", os.Args[2], notesPath).Output()
	if err != nil {
		fmt.Println("search failed: ", err)
	}

	slice := strings.Split(string(out), "\n")
	for i := 0; i < len(slice)-1; i++ {
		n := &note{}
		output := strings.TrimPrefix(slice[i], notesPath)
		n.Path = strings.Split(output, ":")[0]
		n.LineNum = strings.Split(output, ":")[1]
		n.Context = strings.Split(output, ":")[2]
		ns = append(ns, *n)
	}
	findOutput(ns)
	os.Exit(0)
}

func findOutput(s []note) {
	for _, note := range s {
		fmt.Printf("%v found in: %v\non line: %v\ncontext: %v\n", os.Args[2], note.Path, note.LineNum, note.Context)
	}
}

type note struct {
	Path    string
	LineNum string
	Context string
}
