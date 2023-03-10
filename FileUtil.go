/*
* Utility to handle the file operations of the application.
* @Author : Sagar_R
*/
package main

import (
	"log"
	"io/ioutil"
	"regexp"
	"os"
	"math/rand"
	"time"
	"path/filepath"
)

/*
 * Find all the note files matching the pattern and return
 * @method	: findFiles
 * @param	: pattern String	pattern to match
 * @return	: []os.FileInfo		List of files found for the pattern
*/
func findFiles(pattern string) []os.FileInfo {
	var fileInfoArray []os.FileInfo
	files, ioErr := ioutil.ReadDir(_PATH)
    if ioErr != nil {
        log.Fatal(ioErr)
    }

	for _, file := range files {
		//if strings.Contains(file.Name() , noteName) && !file.IsDir() {
		res1,e :=	regexp.MatchString(pattern , file.Name())
		if e != nil {
			log.Fatal(e)
		}
		if res1 && !file.IsDir() {
			//fmt.Println("Found :", file.Name())
			fileInfoArray = append(fileInfoArray, file)
		}
		
    }
	return fileInfoArray
}

/*
 * Read the contents of a file and return
 * @method	: readFile
 * @param	: fileName String	Name of the file to read
 * @return	: string			Return the contents of the file
*/
func readFile(fileName string) string {
	filePath := filepath.Join(_PATH, fileName)
	content, readErr := ioutil.ReadFile(filePath)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return string(content)
}

/*
 * Remove note file if found, else inform user.
 * @method	: removeFile
 * @param	: fileName String	Name of the file to read
 * @return	: bool				Success or Failure
*/
func removeFile(fileName string) bool {
	filePath := filepath.Join(_PATH, fileName)
	e := os.Remove(filePath)
    if e != nil {
        log.Fatal(e)
		return false
    } 
	return true
}

/*
 * Create parent directory for note files.
 * @method	: createDirNotes
 * @param	: none
 * @return	: bool				Success or Failure
*/
func createDirNotes() bool {
	dirErr := os.MkdirAll(_PATH, os.ModePerm)
	if dirErr != nil {
		log.Fatal(dirErr)
		return false
	}
	return true
}

/*
 * Create note file in the notes directory
 * @method	: noteCreate
 * @param	: fileName string	name of the note file to create
 * @param	: content string	body content of the file
 * @return	: bool				Success or Failure
*/
func noteCreate(fileName string, content string) bool {
	filePath := filepath.Join(_PATH, fileName)
	file, errCreate := os.Create(filePath)
	if errCreate != nil {
		log.Fatal(errCreate)
		return false
	}
	defer file.Close()
	if _, writeErr := file.WriteString(content + "\n");
	writeErr != nil {
		log.Fatal(writeErr)
		return false
	}
	return true
}

//refered from https://www.calhoun.io/creating-random-strings-in-go/ 
var stringRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
 
/*
 * Generate random string
 * @method	: getRandomString
 * @param	: length int		length of the string to generate
 * @param	: charset string	character set to be used to generate
 * @return	: string			random string generated
*/
func getRandomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
	  b[i] = charset[stringRand.Intn(len(charset))]
	}
	return string(b)
}
 
/*
 * get random string
 * @method	: getString
 * @param	: length int		length of the string to generate
 * @return	: string			random string
*/
func getString(length int) string {
	return getRandomString(length, _CHAR_SET)
}
