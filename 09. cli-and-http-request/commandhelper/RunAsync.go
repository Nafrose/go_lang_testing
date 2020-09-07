package commandhelper

import "log"

// import . "github.com/nafrose/exploring/clirunner/commandhelper/structs"

func RunAsync(c *CliBindingProperties) *StdInParameter {
	if c.CmdRunningInfo.IsAsync {
		stdInParameter := NewStdInParameters(c.Cmd)
		err := c.Cmd.Start()
		if err != nil {
			log.Fatalf("cmd.Start() failed with '%s'\n", err)
		}

		return stdInParameter
	}

	return nil
}
