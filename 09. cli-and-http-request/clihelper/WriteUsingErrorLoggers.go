package clihelper

func WriteUsingErrortLoggers(cliBindingProperties CliBindingProperties, line string) {
	_, err := w.Write(line)
	if err != nil {
		for _, errorLogger = cliBindingProperties.WritersCollection.ErrorLoggers {
			errorLogger.Writer(cliBindingProperties.WritersCollection.WriterConfiguration, line)
		return out, err
	}
}
