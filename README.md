# Note
Note Taking Utility

## Description
Note taking utitlity to optimize your time with taking/searching/deleting notes.
Note will make a folder or notebook which should be a general topic of the note
and will open up the vim text editor so you can immediately start taking your note.
Once you save and quit the note will be saved in your NOTES_PATH.

## Usage
First make sure you have go installed on your machine.
You can follow instructions here if you do not: [GO_Download](https://golang.org/dl/)

You will first need to create an environmental variable for the path where you would like your
notebooks and notes to live.  
In your `.bashrc` or `.profile` add: `export NOTES_PATH="[[ path to notes ]]"`  
where `[[ path to notes ]]` is where you would like this content to live.  
i.e: `/Users/lenguti/notes`

Once cloned you can move into your note directory and run: `make build`  
This will compile the script and place it in your `$PATH`
