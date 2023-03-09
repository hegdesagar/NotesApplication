/*
Notes Application : User can create, read and delete notes
@author: Sagar_R
*/
package main

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
	"bufio"
	"strings"
	"io/ioutil"
	"regexp"
)

/*
	Main entry point of the application
*/
func main(){
	//Take the user input and validate the arguments
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal(_INVALID_ARGUMENT_MESSAGE)
		log.Fatal(_INVALID_ARGUMENT_ADDITIONAL_MESSAGE, os.Args[0])
		os.Exit(0) //exit the process
	}

	instruction := os.Args[1]	 //fetch the instruction from the argument
	noteName := os.Args[2]		//fetch the nam of the note from the argument

	//validate the name of the note provided; Do not allow these special characters $#@!%^&*(){}<>?~`/+=-;:\""
	match, err := regexp.MatchString(_SPECIAL_CHARACTERS, noteName)
    if err == nil && match == true {
		log.Fatal(_INVALID_CHARACTERS_MESSAGE)
    } 
	//Swith based on the instruction
	switch strings.ToUpper(instruction) {
		case _WRITE_INSTRUCTION 	: writeInstruction(noteName)
		case _READ_INSTRUCTION 		: readInstruction(noteName)
		case _REMOVE_INSTRUCTION 	: deleteInstruction(noteName)
		default : log.Fatal(_INVALID_INSTRUCTION_MESSAGE)
	}
}

//https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html 
func readInstruction(noteNameSubstr string){
	//find all the note files with the matching substring
	var files = findFiles(noteNameSubstr)
	if len(files)==0 {
		log.Println(_NOTE_NOT_FOUND)
		return
	}
	for _, file := range files {
		filePath := filepath.Join(_PATH, file.Name())
		content, readErr := ioutil.ReadFile(filePath)
		if readErr != nil {
			log.Fatal(readErr)
		}
		//Print the contents of the note
		fmt.Println(_SEPARATOR)
		fmt.Println(file.Name() , _COLON)
		fmt.Println(string(content))
		fmt.Print(_NEW_LINE)
	}
}

func deleteInstruction(subjectPattern string){
	var files = findFiles(subjectPattern)
	if len(files) == 0 {
		log.Println(_NOTE_NOT_FOUND)
		return
	}
	if len(files) > 1 {
		//Highly unlikely, but we are generating SUBJECT text randomly.
		//There might be a case where it generated similar text
		log.Println("More than one file found") 
		return
	}
	filePath := filepath.Join(_PATH, files[0].Name())
	e := os.Remove(filePath)
    if e != nil {
        log.Println(_NOTE_NOT_FOUND_DELETE, filePath)
   } 
   log.Println(_NOTE_DELETED_SUCCESSFULLY, files[0].Name())
}

func writeInstruction(noteName string) {
	//Create directory "notes" if it doesnt exits	
	err := os.MkdirAll(_PATH, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	//Creating Subject; a random string of 4 characters
	subjectName := String(4)
	// Take note content from user
	fmt.Print(_NOTE_ENTER_MESSAGE)
	reader := bufio.NewReader(os.Stdin)
	input, ReadErr := reader.ReadString('\n')
	if ReadErr != nil {
		log.Println(_NOTE_INPUT_ERROR, ReadErr)
		return
	}
	// convert CRLF to LF and trim the last character
	input = strings.Replace(input, "\n", "", -1)
	//create a note with the noteName_Subject
	fileName := noteName + _UNDERSCORE + subjectName + _TEXT_FILE_EXTENSION 
	filePath := filepath.Join(_PATH, fileName)
	file, errCreate := os.Create(filePath)
	if errCreate != nil {
		log.Fatal(errCreate)
	}
	defer file.Close()
	if _, writeErr := file.WriteString(input + "\n");
	writeErr != nil {
		log.Fatal(writeErr)
	}
	fmt.Println(_NOTE_CREATE_SUCCESS)
}