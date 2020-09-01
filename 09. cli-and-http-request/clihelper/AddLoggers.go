package clihelper

func (o *OutputLoggers) Add(writer Writer) *OutputLoggers {
	result := make([]Writer, 0, 10)
	result = append(result, o)
	return result
}
