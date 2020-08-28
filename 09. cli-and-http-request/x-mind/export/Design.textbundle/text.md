# Design

## Package

### clihelper

- doSomething()
- CliBinder

	- Properties

		- IOutputLogger
		- IErrorLogger
		- ICmdInput

			- CMD

	- What can do

		- output

			- Bind(writer, stdIn)
			- in = cmd.Std
			- Bind(IOutputLogger, in)

		- error

			- Bind(writer, stdIn)

- ICliBinderParameters

	- cmd, Output, Error

### httphelper

- poster Interface

	- write method(data datastruct, dataStore []datastruct) 

		- data dataStruct

			- getInput(data)

				- isEmpty(data)

					- Verify input

				- Take input

		- append to array (data, dataStore)
		- write to file

			- Append to private file

- getter interface

	- get method(data []datastruct, searchItem string) output string

		- Read []dataStruct from private file / crete abstraction
		- SearchItem string

			- Parse SearchItem string
			- Filter criteria
			- search by filterItem
			- return item/s

## What to achieve

### Mutliple CLI call

### Bind with CLI output

## main

### cliBinder := NewCliBinder(ICliBinderParameters) 

### cliBinder.RunAsync()

### cliBinder.RunSync()

## Actual Design

### cmdrunner

- CliBindingProperties (Struct)

	- WritersCollection

		- []OutputLoggers

			- []Writer

		- []ErrorOutputLoggers

			- []Writer

		- WriterConfiguration

			- IsJsonLoggingEnabled

				- bool

			- IsHttpEnabled

				- bool

			- IsWriteToFile

				- bool

			- WritingToFileLocation

				- string

			- sqliteDb

	- CmdRunningInfo

		- Title
		- Description
		- ID

			- UUID / GUID

		- ParentId

			- nil

		- SessionId
		- BufferSize

	- cmd

- RunAsync
- RunSync

### Use Case of clihelper

- Code Samples

	- argumentsProvided := flag.String("run", "Run,Title,Message 1,Message 2,Stream Delay,Run Times", "Write the argument")
	flag.Parse()
	// argStr := string(*argumentsProvided)
	argsProvided := parseString(argumentsProvided)

	cmd := exec.Command("cli-streamer", "-args", argsProvided)

	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	// cmd.Wait() should be called only after we finish reading
	// from stdoutIn and stderrIn.
	// wg ensures that we finish
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
		log.Println(string(stdout))
		log.Println("Trying to write it out.")
		wg.Done()
	}()

	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)

	wg.Wait()

	err = cmd.Wait()
## Steps to Perform on Enhancements

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

## CLI Process 1

### CLI Process 2 (1)

### CLI Process 2 (2)

## User 1

### invokes

- CLI 1

	- Cli 2 (1)

		- Cli 3 (2)

	- Cli 3

## User 2

### invokes

- CLI 1

	- Cli 2 (1)

		- Cli 3 (2)

	- Cli 3

## Writer (interface)

### Write(config, line string)

### Eg.

- p> line
- {...}

*XMind: ZEN - Trial Version*