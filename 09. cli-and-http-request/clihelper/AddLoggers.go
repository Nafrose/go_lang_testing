package clihelper

import . "github.com/nafrose/exploring/clirunner/clihelper/structs"

type CliBindingPropertiesBuilder struct {
	cliBindingProperties *CliBindingProperties
}

func NewCliBindingPropertiesBuilder() *CliBindingPropertiesBuilder {
	return &CliBindingPropertiesBuilder{&CliBindingProperties{}}
}

func (cli *CliBindingPropertiesBuilder) Build() *CliBindingProperties {
	return cli.cliBindingProperties
}
//
//func (cli *CliBindingPropertiesBuilder) WritersCollectionBuild() *WritersCollectionBuilder {
//	return &WritersCollectionBuilder{*it}
//}
//
//func (cli *CliBindingPropertiesBuilder) CmdRunningInfoBuild() *CmdRunningInfoBuilder {
//	return &CmdRunningInfoBuilder{*it}
//}
//
//func (cli *CliBindingPropertiesBuilder) CmdBuild() *CmdBuilder {
//	return &CmdBuilder{*it}
//}

//func (o *OutputLoggers) Add(writer Writer) *OutputLoggers {
//	result := make([]Writer, 0, 10)
//	result = append(result, o)
//	return result
//}
