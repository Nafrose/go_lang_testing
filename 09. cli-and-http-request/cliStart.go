package main

type cmdrunner struct{}
type PersonType struct {
	Id          int       `json:"ID"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Alias       string    `json:"alias"`
	DateOfBirth time.Time `json:"dob"`
	Profession  string    `json:"profession"`
	Parent      string    `json:"parent"`
}
type CliBinderProperties struct {
	writersCollection []WriterCollection
	CmdRunningInfo    CmdRunningInfo
	cmd               string
}

cmd := ""

type CmdRunningInfo struct {
	Title, Description                  string
	ID, ParentId, SessionId, BufferSize int
}

type WritersCollection struct {
	OutputLoggers       []OutputLogger
	ErrorOutputLoggers  []ErrorOutputLogger
	WriterConfiguration WriterConfiguration
}

type WriterConfiguration struct {
	IsJsonLoggingEnabled bool
	IsHttpEnabled        bool
	ShouldWriteToFile    bool
	WriteToFilelocation  string
}

type Writer interface {
	Writer(WriterConfiguration)
}

type Logger struct {
	OutputLogger OutputLogger
	ErrorOutputLogger ErrorOutputLogger
}
type Add interface {
	add()
}
type OutputLogger struct {
	output PersonType
}

func (o *OutputLogger) Writer (wC WriterConfiguration) []*OutputLogger {
	result := make([]PersonType,0,10)
	arrayInJSON, _ := ioutil.ReadFile(wC.WriteToFilelocation)
	err = json.Unmarshal(arrayInJSON, &result)

	if err != nil {
		fmt.Println(err)
	}
	maxId := getMaxId(result)
	if OutputLogger.Id != 0 {
		OutputLoggger.Id = maxId + 1
		result = append(result, OutputLogger)
	} else {
		result = result
	}
	writeToFile(result)
}

func getMaxId(array []personType) int {
	maxId := 0
	len := len(array)

	if array == nil {
		maxId = 1
	} else {
		for i := 0; i < len; i++ {
			if array[i].Id > maxId {
				maxId = array[i].Id
			}
		}

	}

	return maxId
}

func writeToFile(persons []PersonType) {
	var mutex sync.Mutex
	file, _ := json.MarshalIndent(persons, "", " ")
	mutex.Lock()
	err := ioutil.WriteFile("people.list.json", file, 0644)

	if err != nil {
		panic(err)
	}

	mutex.Unlock()
}

type ErrorOutputLogger struct {
	ErrorWriter []Writer
}

func clirunner.RunAsync(CliBinderProperties) {
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
}
func attachLoggers(CliBindingProperties, r stdInParameter) {
	var out []byte
	buff := make([]byte, 1, CliBinderProperties.CmdRunningInfo.BufferSize)
	for {
		n, err := r.StdOutIn.Read(buf[:])
		if n>0 {
			d :=buf[:n]
			out = append(out, d)
			WriteUsingOutputLoggers(&cliBinderProperties, d)
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func attachErrorLoggers(CliBindingProperties, r stdInParameter) {
	var out []byte
	buff := make([]byte, 1, CliBinderProperties.CmdRunningInfo.BufferSize)
	for {
		n, err := r.StdErrIn.Read(buf[:])
		if n>0 {
			d :=buf[:n]
			out = append(out, d)
			WriteUsingErrorLoggers(&cliBinderProperties, d)
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func main () {
	wc := WriterConfiguration{ShouldWriteToFile: true, WriteToFilelocation: "people.list.json"}
	alex := PersonType{Id: 1, FirstName: "Alex", LastName: "Hales"}
	alex.Writer(wc)
}