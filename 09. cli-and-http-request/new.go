package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type PersonType struct {
	Id          int       `json:"ID"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Alias       string    `json:"alias"`
	DateOfBirth time.Time `json:"dob"`
	Profession  string    `json:"profession"`
	Parent      string    `json:"parent"`
}

var persons = make([]PersonType, 0, 10)
var peopleArray = make([]PersonType, 0, 10)

type cmdrunner struct{}

type CliBinderProperties struct {
	writersCollection []WriterCollection
	// CmdRunningInfo    cmdRunningInfo
	// cmd               cmd
}

// type CmdRunningInfo struct {
// 	Title, Description                  string
// 	ID, ParentId, SessionId, BufferSize int
// }

type WriterCollection struct {
	OutputLoggers []OutputLogger
	// ErrorOutputLoggers  []ErrorOutputLogger
	WriterConfiguration WriterConfiguration
}

type WriterConfiguration struct {
	IsJsonLoggingEnabled bool
	IsHttpEnabled        bool
	ShouldWriteToFile    bool
	WriteToFilelocation  string
}

type OutputLogger struct {
}

// type ErrorOutputLogger struct {
// 	ErrorWriter []Writer
// }

type Writer interface {
	Writer(WriterConfiguration)
}

func (p PersonType) Writer(wC WriterConfiguration) {
	peopleArray = appendToPersonsArray(p, wC)
	if wC.ShouldWriteToFile == true {
		appendToFile(peopleArray, wC.WriteToFilelocation)
	}

}

func getFileContentAsBytes(wC WriterConfiguration) []byte {
	peopleJSON, _ := ioutil.ReadFile(wC.WriteToFilelocation)
	return peopleJSON
}

func getArrayFromFileContent(wC WriterConfiguration) []PersonType {
	array := getFileContentAsBytes(wC)
	_ = json.Unmarshal(array, &peopleArray)
	return peopleArray
}

func getMaxId(wC WriterConfiguration) int {
	peopleArray = getArrayFromFileContent(wC)
	maxId := 0
	len := len(peopleArray)

	if peopleArray == nil {
		maxId = 1
	} else {
		for i := 0; i < len; i++ {
			if peopleArray[i].Id > maxId {
				maxId = peopleArray[i].Id
			}
		}

	}

	return maxId
}

func appendToPersonsArray(inputFromCmdInputProcess PersonType, wC WriterConfiguration) []PersonType {
	maxId := getMaxId(wC)
	if inputFromCmdInputProcess.Id != 0 {
		inputFromCmdInputProcess.Id = maxId + 1
		peopleArray = append(peopleArray, inputFromCmdInputProcess)
	} else {
		peopleArray = peopleArray
	}
	fmt.Println(inputFromCmdInputProcess)
	fmt.Println(peopleArray)
	return peopleArray
}

func appendToFile(peopleArray []PersonType, fileName string) {
	var mutex sync.Mutex
	file, _ := json.MarshalIndent(peopleArray, "", " ")
	mutex.Lock()
	err := ioutil.WriteFile(fileName, file, 0644)

	if err != nil {
		panic(err)
	}

	mutex.Unlock()
}

func main() {
	wc := WriterConfiguration{ShouldWriteToFile: true, WriteToFilelocation: "people.list.json"}
	alex := PersonType{Id: 1, FirstName: "Alex", LastName: "Hales"}
	alex.Writer(wc)
}
