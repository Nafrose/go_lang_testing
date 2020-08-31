package main

func main() {
	writerConfig := WriterCollection.
		OutputLoggers.AddLoggers().
		ErrorOutputLoggers().Add().
		SetConfig(WriterConfiguration{ShouldWriteToFile: true, WriteToFilelocation: "people.list.json"})

	alex := PersonType{Id: 1, FirstName: "Alex", LastName: "Hales"}
	alex.Writer(wc)
}
