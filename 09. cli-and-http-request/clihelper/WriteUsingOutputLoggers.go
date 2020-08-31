package clihelper

func WriteUsingOutputLoggers(c CliBindingProperties, d []byte) {
	_, err := w.Write(d)
	if err != nil {
		return out, err
	}
}
