/*
 * Notes Application : User can create, read and delete notes
 * Create	: go run note WRITE noteName
 * Read		: go run note READ SUBSTRING
 * DELETE	: go run note remove XXXX ( where XXXX is the random string generated)
 * @author: Sagar_R
*/
package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"regexp"
)

/*
 * Main entry point of the application.
 * @method	: main
 * @param	: empty
 * @return	: void
*/
func main(){
	//Take the user input and validate the arguments
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal(_INVALID_ARGUMENT_MESSAGE)
		log.Fatal(_INVALID_ARGUMENT_ADDITIONAL_MESSAGE, os.Args[0])
		return //exit the process
	}

	instruction := os.Args[1]	 //fetch the instruction from the argument
	noteName := os.Args[2]		//fetch the nam of the note from the argument

	//validate the name of the note provided; Do not allow these special characters $#@!%^&*(){}<>?~`/+=;:\""
	match, regErr := regexp.MatchString(_SPECIAL_CHARACTERS, noteName)
    if regErr == nil && match == true {
		log.Fatal(_INVALID_CHARACTERS_MESSAGE)
    } 
	
	//Create directory "notes" if it doesnt exits	
	if !createDirNotes() { 
		return
	}
	//Swith based on the instruction
	switch strings.ToUpper(instruction) {
		case _WRITE_INSTRUCTION 	: writeInstruction(noteName)
		case _READ_INSTRUCTION 		: readInstruction(noteName)
		case _REMOVE_INSTRUCTION 	: deleteInstruction(noteName)
		default : log.Fatal(_INVALID_INSTRUCTION_MESSAGE)
	}
}

/*
 * Method to print the notes on the console.
 * @method	: readInstruction
 * @param	: noteNameSubstr SUBSTRING of type string
 * @return	: void
*/
func readInstruction(noteNameSubstr string){
	//find all the note files with the matching substring
	var files = findFiles(noteNameSubstr)
	if len(files)==0 {
		log.Println(_NOTE_NOT_FOUND)
		return
	}
	for _, file := range files {
		content := readFile(file.Name())
		//Print the contents of the note
		fmt.Println(_SEPARATOR)
		fmt.Println(file.Name() , _COLON)
		fmt.Println(string(content))
		fmt.Print(_NEW_LINE)
	}
}

/*
 * Method to delete the note. It takes a random generated string XXXX to identify the note to be deleted.
 * @method	: deleteInstruction
 * @param	: subjectPattern 	random string XXXX (generated during creation)
 * @return	: void
*/
func deleteInstruction(subjectPattern string){
	var files = findFiles(_UNDERSCORE + subjectPattern)
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
	if removeFile(files[0].Name()) {
		log.Println(_NOTE_DELETED_SUCCESSFULLY, files[0].Name())
	}
}

/*
 * Method to create a note. 
 * @method	: writeInstruction
 * @param	: noteName 	name of the note provided by the user
 * @return	: none
*/
func writeInstruction(noteName string) {
	//Creating Subject; a random string of 4 characters
	subjectName := getString(4)
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
	if noteCreate(fileName , input) {
		fmt.Println(_NOTE_CREATE_SUCCESS,subjectName) //prints success message when the note is created
	}
}