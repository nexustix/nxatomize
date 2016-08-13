package main

import (
	"fmt"
	"os"
	"os/user"

	bp "github.com/nexustix/boilerplate"
	nrc "github.com/nxReplicator/nxReplicatorCommon"
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

	//var providerPath string
	//providerPath = providers.GetEntry(bp.StringAtIndex(1, args)).Path

	atomManager := nrc.AtomManager{WorkingDir: atomDir}

	//switch provider management
	switch bp.StringAtIndex(1, args) {
	case "add provider":

		providers.AddEntry(bp.StringAtIndex(2, args), bp.StringAtIndex(3, args))
		providers.SaveEntries()
		return
		//os.Exit(0)
	case "del provider":
		providers.RemoveEntry(bp.StringAtIndex(2, args))
		providers.SaveEntries()
		return
	}

	//providerID := bp.StringAtIndex(1, args)
	providerAction := bp.StringAtIndex(1, args)
	//providerQuerry := bp.StringAtIndex(3, args)

	//switch atom management
	switch providerAction {
	case "search":
		//fmt.Printf("search\n")
		doSearch(args, providers, &atomManager)

	case "depsearch":
		//atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("atoms", bp.StringAtIndex(1, args)))
		//providerDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("providers", bp.StringAtIndex(1, args)))
		//fmt.Printf("depsearch\n")
		doDepsearch(args, providers, &atomManager)

	case "downinfo":
		//atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("atoms", bp.StringAtIndex(1, args)))
		//providerDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("providers", bp.StringAtIndex(1, args)))
		//fmt.Printf("downinfo\n")
		doDownInfo(args, providers, &atomManager)

	}
	//fmt.Printf("it works !\n")
}
