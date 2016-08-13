package main

import (
	"fmt"
	"os/exec"

	bp "github.com/nexustix/boilerplate"
	nrc "github.com/nxReplicator/nxReplicatorCommon"
)

func doDownInfo(args []string, providers nrc.ProviderList, atomManager *nrc.AtomManager) {

	var providerPath string
	providerPath = providers.GetEntry(bp.StringAtIndex(2, args)).Path

	providerID := bp.StringAtIndex(2, args)
	providerAction := bp.StringAtIndex(1, args)
	providerQuerry := bp.StringAtIndex(3, args)

	if providerID == "custom" {
		if atomManager.HasEntry("custom", providerQuerry) {
			tmpAtom := atomManager.GetEntry("custom", providerQuerry)
			fmt.Printf("%s|%s\n", tmpAtom.URL, tmpAtom.Filename)
		}
	} else {

		if providers.HasEntry(providerID) {
			if providerPath != "" {
				fmt.Printf("</> EXEC >%s %s %s<\n", providerPath, providerAction, providerQuerry)
				providerCommand := exec.Command(providerPath, providerAction, providerQuerry)
				out, err := providerCommand.Output()
				bp.FailError(err)
				fmt.Print(string(out))

			} else {
				fmt.Printf("<!> ERROR command of '%s' empty\n", providerID)
			}
		} else {
			fmt.Printf("<!> ERROR provider '%s' not found\n", providerID)
		}
	}
}
