package clihelper

func WriteUsingErrortLoggers(cliBindingProperties CliBindingProperties, line string) {
	if err != nil {
		for _, errorLogger = cliBindingProperties.WritersCollection.ErrorLoggers {
			errorLogger.Writer(cliBindingProperties.WritersCollection.WriterConfiguration, line)
		return out, err
	}
}
