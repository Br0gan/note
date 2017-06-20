default:
	if [ -f tester ]; then make clean; fi
	make build 

build:
	go build -o note note.go
	mv note /usr/local/bin

clean:
	rm /usr/local/bin/note
