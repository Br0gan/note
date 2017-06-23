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
	fmt.Println(`
NAME
note -- note taking utility

SYNOPSIS
note [note_dir note_file] [--find] [--list] [--get]

DESCRIPTION
Note taking utitlity to optimize your time with taking notes and looking up past notes.
Note will make a folder of the first argument which should be general topic of the note
and will open up the vim text editor of the second argument which should be a more
specific name for your note. Once you save and quit the note will be saved in your NOTES_PATH.

Example: taking a note about multiplication for math..

note math multiplication

note will be saved as ~/${NOTES_PATH}/math/multiplication

FLAGS
--find
Prints the file path, line number, and search results of the string you are trying to find.
i.e: note --find "some text to find"

--list
--get
`)
}

func checkFlags() {
	flag := os.Args[1]
	switch {
	case strings.Contains(flag, "--help") || strings.Contains(flag, "-h"):
		helpMenu()
		os.Exit(0)
	case strings.Contains(flag, "--find") || strings.Contains(flag, "-f"):
		searchNotes()
	case strings.Contains(flag, "--list") || strings.Contains(flag, "-l"):
		listNotes()
	case strings.Contains(flag, "--get") || strings.Contains(flag, "-g"):
		getNotes()
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

func searchNotes() {
	c := make(chan *note)

	out, err := exec.Command("fgrep", "-rHni", os.Args[2], notesPath).Output()
	if err != nil {
		fmt.Println("search failed: ", err)
	}

	go func() {
		slice := strings.Split(string(out), "\n")
		for i := 0; i < len(slice)-1; i++ {
			output := strings.TrimPrefix(slice[i], notesPath)
			n := &note{
				Path:    strings.Split(output, ":")[0],
				LineNum: strings.Split(output, ":")[1],
				Context: strings.Split(output, ":")[2],
			}
			c <- n
		}
		close(c)
	}()

	findOutput(c)
	os.Exit(0)
}

func findOutput(c chan *note) {
	for note := range c {
		fmt.Printf("%v found in: %v\non line: %v\ncontext: %v\n\n", os.Args[2], note.Path, note.LineNum, note.Context)
	}
}

func listNotes() {
	if len(os.Args) == 3 {
		listNote()
	}

	dir, err := exec.Command("ls", notesPath).Output()
	if err != nil {
		fmt.Println("list failed: ", err)
	}
	fmt.Printf("Notebooks:\n%v", string(dir))
	os.Exit(0)
}

func listNote() {
	dir := os.Args[2]

	file, err := exec.Command("ls", notesPath+"/"+dir).Output()
	if err != nil {
		fmt.Println("list failed: ", err)
	}

	fmt.Printf("Notes in Notebook %v:\n%v", dir, string(file))
	os.Exit(0)
}

func getNotes() {
}

type note struct {
	Path    string
	LineNum string
	Context string
}
