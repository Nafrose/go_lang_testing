package main

func main() {
	wc := WriterConfiguration{ShouldWriteToFile: true, WriteToFilelocation: "people.list.json"}
	alex := PersonType{Id: 1, FirstName: "Alex", LastName: "Hales"}
	alex.Writer(wc)
}
