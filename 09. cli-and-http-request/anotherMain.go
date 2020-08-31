package main

func main() {
	wC := WriterConfiguration{ShouldWriteToFile: true, WriteToFilelocation: "people.list.json"})

	alex := PersonType{Id: 1, FirstName: "Alex", LastName: "Hales"}
	alex.Writer(wc)
}
