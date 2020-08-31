package clihelper

import (
	"io"
	"log"
	"sync"
)

type CmdRunner struct {
	CliBindingProperties CliBindingProperties
}

type Writer interface {
	Writer(wc WriterConfiguration, line string)
}

type stdInParameter struct {
	stdoutIn []byte
	stderrIn error
}

func RunAsync(c CliBindingProperties) stdInParameter {
	var stdout, stderr []byte
	var stdInParameter stdInParameter
	var errStdout, errStderr error
	stdInParameter.stdoutIn, _ = c.Cmd.StdoutPipe()
	stdInParameter.stderrIn, _ = c.Cmd.StderrPipe()
	err := c.Cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	return stdInParameter
}

func CliBind(c CliBindingProperties, s stdInParameter) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout := AttachLoggers(c, s)
		log.Println(string(stdout))
		log.Println("Trying to write it out.")
		wg.Done()
	}()
}
func AttachLoggers(c CliBindingProperties, r stdInParameter) ([]byte, error) {
	var out []byte
	buff := make([]byte, 1, c.CmdRunningInfo.BufferSize)
	for {
		n, err := r.stdoutIn.Read(buff[:])
		if n > 0 {
			d := buff[:n]
			out = append(out, d...)
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

func AttachErrorLoggers(c CliBindingProperties, r stdInParameter) ([]byte, error) {
	var out []byte
	buff := make([]byte, 1, c.CmdRunningInfo.BufferSize)
	for {
		n, err := r.stderrIn.Read(buff[:])
		if n > 0 {
			d := buff[:n]
			out = append(out, d...)
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

func WriteUsingOutputLoggers(c CliBindingProperties, d []byte) {
	_, err := w.Write(d)
	if err != nil {
		return out, err
	}
}

func WriteUsingErrortLoggers(c CliBindingProperties, d []byte) {
	_, err := w.Write(d)
	if err != nil {
		return out, err
	}
}

func main() {
	wc := WriterConfiguration{ShouldWriteToFile: true, WriteToFilelocation: "people.list.json"}
	alex := PersonType{Id: 1, FirstName: "Alex", LastName: "Hales"}
	alex.Writer(wc)
}
