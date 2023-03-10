***************************************************************
Notes Application ReadMe
***************************************************************

*   Notes Application is developed using Golang(Go). The user can create new notes and would be 
    stored in the folder "notesfolder", which the user can change the location if 
    required by updating the preferences section in Constants.go file. The user 
    can also read the note(s) provided the note exists in the location for the pattern 
    provided. The user can delete the notes created by them 
    (i.e, if they are the owner of the note) by providing the four character string generated randomly (XXXX).

*   Below are the steps to follow to execute and build the code.
    
    * STEP 1:Build the source code using Go, if Go not installed,
	    follow https://go.dev/doc/install 
            To build the source code run the below command in terminal 
	    from the "NoteApp" directory.
            CMD-> go build note
    * STEP 2: We need to provide appropriate setUID permissions for other users to execute this application,
            execute the below command in terminal from the "NoteApp" directory.
            CMD-> sudo chmod u+s ./note  
    * STEP 3: Use the below command for execute instructions (note: the instructions 
            read/write/remove are not case sensitive)
            CMD->    CREATE : ./note Write noteName
            CMD->    READ   : ./note Read pattern   
            CMD->    DELETE : ./note Remove XXXX     

* Preferences : the user can change the preference for the application which are declared in the
                Constants.go file (preferences section)  
            	_CHAR_SET : The set of character used for generating the random string XXXX which 
                                would be append to the noteName.
	        _PATH 	  :	The folder location where the notes are stored.
	        _SPECIAL_CHARACTERS : Set of special characters which are not allowed for noteName.  

* Notes/Summary:* User inputs are validated
                * Users will not be allowed to delete notes which they are not the owner, althought
                  they would be able to read the notes.      
                * I did try to setUid programatically using go's syscal.Setuid(UID), but the i 
                  wasn't able to successfully implement this.
                * I have styled the comments (class and method level) like java, using annotations. 
                  Although, annotation are not supported in go.
                * The part of code to create random string is referred from 
                   https://www.calhoun.io/creating-random-strings-in-go/  
                * The application doesn't use third-party components.
                * Separation of concerns is take care off, File access is all part of FileUtil.go.
