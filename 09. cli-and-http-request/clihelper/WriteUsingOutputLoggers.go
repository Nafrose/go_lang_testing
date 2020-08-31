package clihelper

func WriteUsingOutputLoggers(c CliBindingProperties, d []byte) {
	_, err := w.Writer(c.WritersCollection.WriterConfiguration.WriteToFilelocation, d)
	if err != nil {
		return out, err
	}
}
