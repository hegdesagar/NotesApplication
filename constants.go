/*
* Constants file for note applications.
* @Author : Sagar_R
*/
package main

const (

	//Preferences
	_CHAR_SET 			= "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_PATH 				= "notes"
	_SPECIAL_CHARACTERS = "[$#@!%^&*(){}<>?~`/+\\=;:\"]"
	
	//Text Fomatting 
	_SEPARATOR 			= "----------------------------------------------------"
	_TAB_SPACE 			= "	"
	_COLON 				= " : "
	_NEW_LINE			= "\n"
	_UNDERSCORE			= "_"

	//Instruction constants
	_READ_INSTRUCTION 	= "READ"
	_WRITE_INSTRUCTION 	= "WRITE"
	_REMOVE_INSTRUCTION = "REMOVE"

	_TEXT_FILE_EXTENSION = ".txt"

	//Messages
	_NOTE_ENTER_MESSAGE  = "Enter Content Here -> "

	//INFO messages
	_NOTE_CREATE_SUCCESS 		= "INFO : Note created Successfully"
	_NOTE_NOT_FOUND 			= "INFO : No notes found for the pattern provided"
	_NOTE_NOT_FOUND_DELETE		= "INFO : Note not found, "
	_NOTE_DELETED_SUCCESSFULLY	= "INFO : Note deleted successfully,"

	//Error Messages
	_INVALID_ARGUMENT_MESSAGE			 = " ERROR : Invalid commandline parameters!!"
	_INVALID_ARGUMENT_ADDITIONAL_MESSAGE = "ERROR : note <Instruction: Write/Read/Remove> <Subject> \n"
	_INVALID_CHARACTERS_MESSAGE			 = "ERROR : Invalid Characters,Do not use special characters for subject "
	_INVALID_INSTRUCTION_MESSAGE		 = "ERROR : Invalid Instruction,use appropriate instruction, Write/Read/Remove"
	_NOTE_INPUT_ERROR					 = "ERROR : An error occured while reading input. Please try again"
)

