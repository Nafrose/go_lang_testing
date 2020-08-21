package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
)

type personType struct {
	Id          int       `json:"ID"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Alias       string    `json:"alias"`
	DateOfBirth time.Time `json:"dob"`
	Profession  string    `json:"profession"`
	Parent      string    `json:"parent"`
}

var person personType
var persons = make([]personType, 0, 10)
var propertyArray = make([]string, 0, 10)

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func GetLine(promptMessage string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if !IsEmpty(promptMessage) {
		fmt.Print(promptMessage)
	}

	return reader.ReadString('\n')
}

func getFileContentAsBytes() []byte {
	peopleJSON, _ := ioutil.ReadFile("people.list.json")
	return peopleJSON
}

func getArrayFromFileContent(array []byte) []personType {
	var peopleArray = make([]personType, 0, 10)
	_ = json.Unmarshal(array, &peopleArray)
	return peopleArray
}

func getMaxId(peopleArray []personType) int {
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

func appendToPersonsArray(person personType, peopleArray []personType) []personType {
	maxId := getMaxId(peopleArray)
	if person.Id != 0 {
		person.Id = maxId + 1
		persons = append(peopleArray, person)
	} else {
		persons = peopleArray
	}
	return persons
}

func appendToFile() {
	var mutex sync.Mutex
	file, _ := json.MarshalIndent(persons, "", " ")
	mutex.Lock()
	err := ioutil.WriteFile("people.list.json", file, 0644)

	if err != nil {
		panic(err)
	}

	mutex.Unlock()
}

func getPeople() string {
	peopleJSON, _ := ioutil.ReadFile("people.list.json")
	allPeople := string(peopleJSON)
	return allPeople
}

func parseArgument(s string) []string {
	var argumentsAfterParsing []string
	var argumentsAfterRemovingSpace []string
	var noOfSearchFilters int

	if !strings.Contains(s, ":") {
		if strings.Contains(s, "/") {
			argumentsAfterParsing = strings.Split(s, "/")
		} else if !strings.Contains(s, "/") {
			argumentsAfterParsing = strings.Split(s, "/")
			argumentsAfterParsing = append(argumentsAfterParsing, "")
		}

	} else {
		argumentsWithoutComma := strings.Join(strings.Split(s, ","), " ")
		noOfSearchFilters = strings.Count(argumentsWithoutComma, ":")
		argumentsWithoutColon := strings.Join(strings.Split(argumentsWithoutComma, ":"), " ")
		argumentsAfterRemovingSpace = strings.Split(argumentsWithoutColon, " ")
		index := 0

		for i := 1; i < noOfSearchFilters+1; i++ {
			argumentsAfterParsing = append(argumentsAfterParsing, argumentsAfterRemovingSpace[index+i])
			if len(argumentsAfterParsing) < (noOfSearchFilters) {
				index++
			} else {
				break
			}
		}
	}

	return argumentsAfterParsing
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil {
		return true
	} else {
		return false
	}

}

func getIndividualPersonIndex(s string) int {
	var index int
	if isDigit(s) {
		idToSearchFor, _ := strconv.Atoi(s)
		for i := 0; i < len(persons); i++ {
			if idToSearchFor == persons[i].Id {
				index = i
			} else {
				index = 0 //negative value avoided for now
			}
		}
	} else {
		for i := 0; i < len(persons); i++ {
			if s == persons[i].FirstName {
				index = i
			} else {
				index = 0 //negative value avoided for now
			}
		}
	}

	return index
}

func getIndividualPersonInfo(s string) personType {
	i := getIndividualPersonIndex(s)
	return persons[i]

}

func getAllValuesForProperty(s string) {
	for i := 0; i < len(persons); i++ {
		fmt.Println(persons[i].FirstName)
	}
}

func getAllValuesForMultiProperty(s string) {
	for i := 0; i < len(persons); i++ {
		if s == "firstname-and-dateofbirth" {
			fmt.Printf("First name: %s  \n", persons[i].FirstName)
			fmt.Printf("Date of birth: %d  \n", persons[i].DateOfBirth)
		} else if s == "id-and-firstname" {
			fmt.Printf("Id: %d  \n", persons[i].Id)
			fmt.Printf("First name: %s  \n", persons[i].FirstName)
		}
	}

}

func getDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

// func getRawBodyToCliJson(s string) string {
// 	var p fastjson.Parser
// 	v, err := p.Parse(s)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return v
// }

func People(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cmd := exec.Command(getDir()+"/input.exe", "-get", "people")
	outputFromOtherFile, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Fprintf(w, "List of all people are - %s \n", outputFromOtherFile)
}
func CreatePeople(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "The person has been added %s \n", r.Body)
	fmt.Printf("%T %v ", r.Body, r.Body)
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	fmt.Fprintf(w, "The person has been added %s \n", buf)
	fmt.Printf("%T %v", buf, buf)

	fmt.Println(buf.String())
	// fmt.Fprintf(w, "The person to be been added %s \n", string(buf))
	// cmd := exec.Command(getDir()+"/input.exe", "-run", "create-person-json-http")

	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Stdin = strings.NewReader(buf.String() + "\n")
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	login := exec.Command(getDir()+"/input.exe", "-run", "create-person-json-http")

	buffer := bytes.Buffer{}
	buffer.Write([]byte(`{"Id":1,"FirstName":"Billy","LastName":"Ul Karim","Alias":"alim","DateOfBirth":"2018-09-22T19:42:31+07:00","Profession":"Mentoring","Parent":"null"}\n`))
	login.Stdin = &buffer

	login.Stdout = os.Stdout
	login.Stderr = os.Stderr

	err := login.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// fmt.Fprintf(w, "The CLI output runs %s \n", string(compiledCliPersonBytes))
}
func getPeopleById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "The person you are looking for is - %s \n", getIndividualPersonInfo(ps.ByName("Id")))
}

func getPeopleByFirstName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "The person you are looking for is - \n", getIndividualPersonInfo(ps.ByName("name")))
}

func isFlagPassed(name string) bool {
	found := false
	flag.VisitAll(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	methodOfChoicePtr := flag.String("run", " ", "Write the preferred method for input")
	processOfChoicePtr := flag.String("get", " ", "Write the required operation")
	throwErrorPtr := flag.String("throw", "message", "anything written as value")
	flag.Parse()
	methodOfChoice := *methodOfChoicePtr
	processOfChoice := *processOfChoicePtr
	throwError := *throwErrorPtr

	output := getFileContentAsBytes()
	peopleArray := (getArrayFromFileContent(output))

	if strings.Compare(methodOfChoice, "create-person-cli") == 0 {
		person.FirstName, _ = GetLine("What is your First name?")
		person.LastName, _ = GetLine("What is your Last name?")
		person.Alias, _ = GetLine("What is your Alias?")
		person.Profession, _ = GetLine("What is your profession?")
		person.Parent, _ = GetLine("Enter parents information")

		appendToPersonsArray(person, peopleArray)
		appendToFile()

	} else if strings.Compare(methodOfChoice, "create-person-json-http") == 0 {
		fmt.Printf("method of choice: %s\n", methodOfChoice)
		jsonInput, _ := GetLine("Enter JSON-from-http")
		fmt.Println(jsonInput)
		jsonInputByte := []byte(jsonInput)
		fmt.Println(jsonInputByte)
		err := json.Unmarshal(jsonInputByte, &person)
		if err != nil {
			fmt.Println(err)
		}

		appendToPersonsArray(person, peopleArray)
		appendToFile()

	} else if strings.Compare(methodOfChoice, "create-person-json") == 0 {
		jsonInput, _ := GetLine("Enter JSON")
		jsonInputByte := []byte(jsonInput)
		err := json.Unmarshal(jsonInputByte, &person)
		if err != nil {
			fmt.Println(err)
		}

		appendToPersonsArray(person, peopleArray)
		appendToFile()

	} else if strings.Compare(methodOfChoice, " ") == 0 {
		appendToPersonsArray(person, peopleArray)
	}

	additionalArguments := parseArgument(processOfChoice)

	if strings.Compare(additionalArguments[0], "people") == 0 && strings.Compare(additionalArguments[1], "") == 0 {
		fmt.Println(getPeople())

	} else if strings.Compare(additionalArguments[0], "people") == 0 && isDigit(additionalArguments[1]) {
		fmt.Println(getIndividualPersonInfo(additionalArguments[1]))

	} else if isDigit(additionalArguments[0]) && len(additionalArguments) > 1 {
		fmt.Println(len(additionalArguments))
		for i := 0; i < len(additionalArguments); i++ {
			fmt.Println(getIndividualPersonInfo(additionalArguments[i]))
		}

	} else if strings.Compare(additionalArguments[0], "people") == 0 && !isDigit(additionalArguments[1]) && len(additionalArguments) > 2 {
		// 	findFilterProperty(additionalArguments[1])
		fmt.Println(getIndividualPersonInfo(additionalArguments[2]))

	} else if strings.Compare(additionalArguments[0], "people-plain") == 0 && strings.Compare(additionalArguments[1], "firstname") == 0 && len(additionalArguments) > 2 {
		getAllValuesForProperty(additionalArguments[2])

	} else if strings.Compare(additionalArguments[0], "people-plain") == 0 && strings.Compare(additionalArguments[1], "firstname-and-dateofbirth") == 0 {
		getAllValuesForMultiProperty(additionalArguments[1])

	} else if strings.Compare(additionalArguments[0], "people-plain") == 0 && strings.Compare(additionalArguments[1], "id-and-firstname") == 0 {
		getAllValuesForMultiProperty(additionalArguments[1])

	}

	if strings.Compare(throwError, "") != 0 {
		fmt.Println(throwError)
	}

	/* httprouter package */

	if len(os.Args) < 2 {
		router := httprouter.New()
		router.GET("/people", People)
		router.POST("/people/create", CreatePeople)
		router.GET("/people/:Id", getPeopleById)
		router.GET("/firstname/:name", getPeopleByFirstName)

		log.Fatal(http.ListenAndServe(":8080", router))
	}

}
