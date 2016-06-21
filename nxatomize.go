package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"

	bp "github.com/nexustix/boilerplate"
	nrc "github.com/nexustix/nxReplicatorCommon"
)

func main() {
	version := "V.0-1-0"
	fmt.Printf("<-> NxAtomize Version: %s\n", version)

	args := os.Args

	usr, err := user.Current()
	bp.FailError(err)
	workingDir := usr.HomeDir
	//fmt.Printf("<-> Working dir: %s\n", workingDir)
	fmt.Printf("<-> Home dir: %s\n", workingDir)

	atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", "atoms")
	providerDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", "providers")

	providers := nrc.ProviderList{Dir: providerDir, Filename: "providers.json"}
	providers.LoadEntries()

	var providerPath string
	providerPath = providers.GetEntry(bp.StringAtIndex(1, args)).Path

	atomManager := nrc.AtomManager{WorkingDir: atomDir}

	//switch provider management
	switch bp.StringAtIndex(1, args) {
	case "add provider":

		providers.AddEntry(bp.StringAtIndex(2, args), bp.StringAtIndex(3, args))
		providers.SaveEntries()
		os.Exit(0)
	case "del provider":
		providers.RemoveEntry(bp.StringAtIndex(2, args))
		providers.SaveEntries()
	}

	providerID := bp.StringAtIndex(1, args)
	providerAction := bp.StringAtIndex(2, args)
	providerQuerry := bp.StringAtIndex(3, args)

	//switch atom management
	switch providerAction {
	case "search":
		//atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("atoms", providerID))
		//providerDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("providers", bp.StringAtIndex(1, args)))
		fmt.Printf("search\n")

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
				nrc.OutputToAtomsAndAdd("curse", string(out), &atomManager)

			} else {
				fmt.Printf("<!> ERROR command of '%s' empty\n", providerID)
			}
		} else {
			fmt.Printf("<!> ERROR provider '%s' not found\n", providerID)
		}

		//fmt.Println(providers.GetEntry(bp.StringAtIndex(1, args)).Path, "search", "mystcraft")
		//exec.Command(providers.GetEntry(bp.StringAtIndex(1, args)).Path, "search", "mystcraft")

	case "deps":
		//atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("atoms", bp.StringAtIndex(1, args)))
		//providerDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("providers", bp.StringAtIndex(1, args)))
		fmt.Printf("deps\n")

	case "downinfo":
		//atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("atoms", bp.StringAtIndex(1, args)))
		//providerDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("providers", bp.StringAtIndex(1, args)))
		fmt.Printf("downinfo\n")

	}
	//fmt.Printf("it works !\n")
}
