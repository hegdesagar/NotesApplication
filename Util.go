/*

@Author : Sagar_R
*/
package main

import (
	//"fmt"
	"log"
	"io/ioutil"
	"regexp"
	"os"
	"math/rand"
	"time"
)


func findFiles(pattern string) []os.FileInfo {
	var fileInfoArray []os.FileInfo
	files, err := ioutil.ReadDir(_PATH)
    if err != nil {
        log.Fatal(err)
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

//copied from https://www.calhoun.io/creating-random-strings-in-go/  :START
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
  
  func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
	  b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
  }
  
  func String(length int) string {
	return StringWithCharset(length, _CHAR_SET)
  }
  //END 
