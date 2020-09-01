package main

import (
	_ "cli-and-http-request\clihelper\Structs"
)

func main() {
	wC := WriterConfiguration{ShouldWriteToFile: true}
	cliBindingProp := CliBinderProperties{CmdRunningInfo.IsAsync: true, Cmd : "cli-streamer.exe"}
	// alex := PersonType{Id: 1, FirstName: "Alex", LastName: "Hales"}

	cmd := exec.Command(cliBindingProp.Cmd, "", "")
	RunAsync(cliBindingProp)
}
