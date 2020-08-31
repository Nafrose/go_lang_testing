package clihelper

import "io"

func AttachErrorLoggers(c CliBindingProperties, r StdInParameter) ([]byte, error) {
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
