package main

import (
	"fmt"
	"os/exec"

	bp "github.com/nexustix/boilerplate"
	nrc "github.com/nexustix/nxReplicatorCommon"
)

func doDownInfo(args []string, providers nrc.ProviderList) {

	var providerPath string
	providerPath = providers.GetEntry(bp.StringAtIndex(1, args)).Path

	providerID := bp.StringAtIndex(1, args)
	providerAction := bp.StringAtIndex(2, args)
	providerQuerry := bp.StringAtIndex(3, args)

	if providers.HasEntry(providerID) {
		if providerPath != "" {
			fmt.Printf("</> EXEC >%s %s %s<\n", providerPath, providerAction, providerQuerry)
			providerCommand := exec.Command(providerPath, providerAction, providerQuerry)
			out, err := providerCommand.Output()
			bp.FailError(err)
			fmt.Println(string(out))

		} else {
			fmt.Printf("<!> ERROR command of '%s' empty\n", providerID)
		}
	} else {
		fmt.Printf("<!> ERROR provider '%s' not found\n", providerID)
	}
}
