package clihelper

func WriteUsingOutputLoggers(cliBindingProperties CliBindingProperties, line string) {
	_, err := w.Writer(line)
	if err != nil {
		for _, outputLogger = cliBindingProperties.WritersCollection.OutputLoggers {
			outputLogger.Writer(cliBindingProperties.WritersCollection.WriterConfiguration, line)
		return out, err
	}
}
