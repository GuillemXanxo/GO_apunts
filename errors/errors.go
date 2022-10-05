package main

import (
	"Apunts_Go/errors/errortypes"
	"apunts_go/errors/errortypes"
	"fmt"
	"io/ioutil"
	"log"
)

//An error is an interface
//This means we can create our own error types as structs that implement error interface

//Ways to work with errors:

// 1- Propagate the error to the caller
func Read1(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	} else {
		return string(file), nil
	}
}

// 2- Exit the program --> reserved only for main and for very important errors
func Read2(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

// 3- Log the error and continue on
func Read3(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error in Read3:", err)
	}
	return string(file)
}

// 4- Personalized errors:

type errKind int

const (
	_ errKind = iota
	noFile
	noDirectory
) //LACKS MORE EXPLANATION

//5- Error types:
//We have created a package errors with file errortypes.go with a struct BadResponseErr
func Read5(filename string) ([]byte, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errortypes.WrapNotFoundErr(err, "Error while reading file")
	}
	return file, nil
}

//The function calling Read5() would look like
func FakeMain() {
	fileBytes, err := Read5("file.txt")
	if errortypes.IsNotFoundErr(err) { // --> We check wether this error is of type notFoundErr declared in package errortypes
		log.Fatal(err)
	}
	fmt.Println(string(fileBytes))
}

/*WHY WORKING SO COMPLICATED?
we are defining error by behaviour. All errors that mean we could not reach our data will be notFounErr.
example: we are reading a file from our pc and uploading it to an s3 bucket.
we could create errors for not finding the data and errors for not having connection. We can then decide the code to react differently according to the type of error
*/
