package builders

import . "github.com/nafrose/exploring/clirunner/clihelper/structs"

type CliBindingPropertiesBuilder struct {
	CliBindingProperties *CliBindingProperties
}

func NewCliBindingPropertiesBuilder() *CliBindingPropertiesBuilder {
	return &CliBindingPropertiesBuilder{&CliBindingProperties{}}
}

func (cliBindingPropertiesBuilder *CliBindingPropertiesBuilder) Build() *CliBindingProperties {
	return cliBindingPropertiesBuilder.CliBindingProperties
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
