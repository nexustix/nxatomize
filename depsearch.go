package main

import (
	"fmt"
	"os/exec"
	"strings"

	bp "github.com/nexustix/boilerplate"
	nrc "github.com/nexustix/nxReplicatorCommon"
)

func doDepsearch(args []string, providers nrc.ProviderList, atomManager *nrc.AtomManager) {

	var providerPath string
	providerPath = providers.GetEntry(bp.StringAtIndex(2, args)).Path

	providerID := bp.StringAtIndex(2, args)
	providerAction := bp.StringAtIndex(1, args)
	providerQuerry := bp.StringAtIndex(3, args)

	if providers.HasEntry(providerID) {
		if providerPath != "" {
			fmt.Printf("<~> EXEC >%s %s %s<\n", providerPath, providerAction, providerQuerry)

			//FIXME potentially dangerous if one is careless
			providerCommand := exec.Command(providerPath, providerAction, providerQuerry)
			out, err := providerCommand.Output()
			//bp.FailError(err)
			if bp.GotError(err) {
				fmt.Println("<!> ERROR finding dependencies")
			} else {
				segments := strings.Split(strings.TrimSpace(string(out)), " ")
				tmpAtom := atomManager.GetEntry(providerID, providerQuerry)
				for _, v := range segments {
					if v != "" {
						tmpAtom.Dependencies = append(tmpAtom.Dependencies, v)
						fmt.Println("<~> FOUND >" + v + "<")
					}
				}
				tmpAtom.Dependencies = bp.EliminateDuplicates(tmpAtom.Dependencies)
				tmpAtom.DoDepCheck = false
				atomManager.SetEntry(providerID, tmpAtom)
				//fmt.Printf(">%v<\n", tmpAtom)

				//fmt.Println("<~> >" + string(out) + "<")
			}
		} else {
			fmt.Printf("<!> ERROR command of '%s' empty\n", providerID)
		}
	} else {
		fmt.Printf("<!> ERROR provider '%s' not found\n", providerID)
	}
}
