# CLI Rewrite using DesignPatterns

## Links

- Interfaces:
  - [Implement interface in GO](https://bit.ly/31M3VYd)
- Parameters by Reference or Value:
  - [go - Can functions be passed as parameters? - Stack Overflow](https://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters)
  - [Pass function as an argument to another function in Go (Golang) – Welcome To Golang By Example](https://golangbyexample.com/func-as-func-argument-go/)
- HttpRouter:
  - [julienschmidt/httprouter: A high performance HTTP request router that scales well](https://github.com/julienschmidt/httprouter)
  - [httprouter - GoDoc](https://godoc.org/github.com/julienschmidt/httprouter)
  - [go - How to stop http.ListenAndServe() - Stack Overflow](https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve)
  - [express - Golang httpRouter returns the last response when used with the slice of functions - Stack Overflow](https://stackoverflow.com/questions/39905623/golang-httprouter-returns-the-last-response-when-used-with-the-slice-of-function)
  - [go - why goroutine block main func in this http server? - Stack Overflow](https://stackoverflow.com/questions/43861055/why-goroutine-block-main-func-in-this-http-server)
  - [go - How to pass a httprouter.Handle to a Prometheus http.HandleFunc - Stack Overflow](https://stackoverflow.com/questions/55737480/how-to-pass-a-httprouter-handle-to-a-prometheus-http-handlefunc)
  - [express - Golang httpRouter returns the last response when used with the slice of functions - Stack Overflow](https://stackoverflow.com/questions/39905623/golang-httprouter-returns-the-last-response-when-used-with-the-slice-of-function)
  - Shutdown
    - [Graceful shutdown · Issue #253 · julienschmidt/httprouter](https://github.com/julienschmidt/httprouter/issues/253)
    - [go - How to stop http.ListenAndServe() - Stack Overflow](https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve)

## Use golang httprouter library

- Create HTTP endpoints which can serve REST Content:
  - Create data-struct as person.sample.json
  - create 5 person using the same struct. Use cli command to take input from user and then add it to the JSON list (Hint use slice peopleSlice := make([]PersonType, length: 0, capacity: 10) and then peopleSlice = append(peopleSlice, newPerson) eg. Slice Example)

## CLI options

- cli.go -run create-person-cli: Will prompt to the user asking for first name, lastname and so on from the json and then app will take that input and create a struct and push it to json list and the finally writes all json back to file people.list.json day folder\data\people.list.json
- cli.go -run create-person-json: take a json input from cli and does the same as above. [Use mutex while writing to file.]
- cli.go -get people: will result all the person in the json file in json format in the console as json.
- cli.go -get people/5: will return a single person by search through id:5. 
- cli.go -get "people:5,people:1": will return a single person by search through id:5 and id:1.
- cli.go -get people/firstname/alim: will return a single person by search through firstName:alim do a contains search.
- cli.go -get people-plain/firstname: will all the first name in the system with new line separator. Note: We don't need to create dynamic search for other fields such as lastname (not JSON).
- cli.go -get people-plain/firstname-and-dateofbirth: will all the first name and date of birth in the system with 2 new line separator (not JSON).
- cli.go -get people-plain/id-and-firstname: will all the id and first in the system with 2 new line separator.
- cli.go -throw "message": will throw an error with given message.

### (Person, create a person Struct, use the CLI to delegate) with end points as follows (Hint for http endpoint create you can help from aukgit - Go modules training)

- localhost:8080/people Http-GET : will return list of all the person in the system as JSON array
- localhost:8080/people/create Http-POST : will create and store person to the json file.
- localhost:8080/people/5 Http-GET : will get a person as JSON with id : 5.
- localhost:8080/firstName/alim Http-GET : will get a person as JSON with first name contains alim.

## Steps to perform for enhancements

```notes
CliBindingPropertiesBuilder.
  OutputLoggers().Add(outputLogger).
  ErrorLoggers().Add(errorLogger).
  CmdRunningConfig().Add(config)
  

writerConfig := WritersConfiguration.
  OutputLogger().Add(...).
  ErrorOutputLogger().Add(...).
  SetConfig(WriterConfiguration{IsJson...}).
  Build()

cliBindingProperties := CliBindingPropertiesBuilder.
  SetWriterConfig(writerConfig).
  SetExecutor(cmd)  

clirunner.RunAsync(cliBindingProperties)
	-> 
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	-> stdoutIn, stderrIn -> stdInParameter(output, error)
	-> delegation (..) -> CliBind(cliBindingProperties, stdInParameter)
			-> CliBind(..) execute lines 107-116
				-> Line 110 stdout, errStdout = attachLoggers(cliBindingProperties, stdInParameter)
					-> attachLoggers(....)
							-> 
	attachLoggers(cliBindingProperties, r stdInParameter) 
						
	var out []byte
	buf := make([]byte, 1, cliBindingProperties.Config.BufferSize)
	for {
		n, err := r.StdOutIn.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			WriteUsingOutputLoggers(&cliBindingProperties, d) //
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}

	attachErrorLoggers(cliBindingProperties, r stdInParameter) 
						
	var out []byte
	buf := make([]byte, 1, cliBindingProperties.Config.BufferSize)
	for {
		n, err := r.StdOutErrorIn.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			WriteUsingErrorLoggers(&cliBindingProperties, d) //
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
```

## Module Initialize

```bash
go mod init github.com/YourName/YourModuleName
go get new-package
go get go.uber.org/zap
go mod tidy

```