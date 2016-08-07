package main

import (
	"fmt"
	"os/exec"

	bp "github.com/nexustix/boilerplate"
	nrc "github.com/nexustix/nxReplicatorCommon"
)

func doSearch(args []string, providers nrc.ProviderList, atomManager *nrc.AtomManager) {
	//atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("atoms", providerID))
	//providerDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("providers", bp.StringAtIndex(1, args)))

	var providerPath string
	providerPath = providers.GetEntry(bp.StringAtIndex(1, args)).Path

	providerID := bp.StringAtIndex(1, args)
	providerAction := bp.StringAtIndex(2, args)
	providerQuerry := bp.StringAtIndex(3, args)

	//providers.GetEntry(bp.StringAtIndex(1, args)
	if providers.HasEntry(providerID) {
		if providerPath != "" {
			fmt.Printf("</> EXEC >%s %s %s<\n", providerPath, providerAction, providerQuerry)

			//FIXME potentially dangerous if one is careless
			providerCommand := exec.Command(providerPath, providerAction, providerQuerry)
			out, err := providerCommand.Output()
			bp.FailError(err)
			fmt.Println(string(out))

			//atomManager.SetEntry("curse", nrc.StringToAtom(string(out)))
			//nrc.OutputToAtomsAndAdd("curse", string(out), &atomManager, true)
			//XXX
			nrc.OutputToAtomsAndAdd(providerID, string(out), atomManager, true)
			///*
			tmpAtoms := nrc.OutputToAtoms(string(out), true)

			for _, v := range tmpAtoms {
				fmt.Printf("<-> '%s'\n", v.ID)

				tmpAtom := atomManager.GetEntry(providerID, v.ID)
				if tmpAtom.DoDepCheck {
					fmt.Printf("</> EXEC >%s %s %s %s<\n", "nxatomize", providerID, "depsearch", v.ID)
					//FIXME potentially dangerous if one is careless
					providerCommand := exec.Command("nxatomize", providerID, "depsearch", v.ID)
					output, err := providerCommand.Output()
					if bp.GotError(err) {
						fmt.Printf("<!> ERROR getting deps of: '%s'\n", v.ID)
					} else {
						fmt.Printf("%s", output)
						fmt.Printf("<-> done getting deps of: '%s'\n", v.ID)
					}
				} else {
					fmt.Printf("<-> deps of '%s' already indexed\n", v.ID)
				}
				fmt.Println("#####")
			}
			//*/

		} else {
			fmt.Printf("<!> ERROR command of '%s' empty\n", providerID)
		}

	} else {
		fmt.Printf("<!> ERROR provider '%s' not found\n", providerID)
	}

	//fmt.Println(providers.GetEntry(bp.StringAtIndex(1, args)).Path, "search", "mystcraft")
	//exec.Command(providers.GetEntry(bp.StringAtIndex(1, args)).Path, "search", "mystcraft")

}
