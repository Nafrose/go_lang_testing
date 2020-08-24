package cliHelper

import (
	"crypto/tls"
	_ "flag"
	_ "fmt"
	"io"
	_ "io"
	_ "io/ioutil"
	_ "log"
	_ "os"
	_ "os/exec"
	_ "sync"
)

func CreateProcessConfig(
	properties ProcessConfigProperties,
	writers WritersCollection) ProcessConfig {
	return ProcessConfig{
		Writers:    writers,
		Properties: properties,
	}
}

func BindOutputWithProcessConfigWriters(writer Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 0, processConfig.Properties.BufferSize)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			writer.Write(string(out))
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
