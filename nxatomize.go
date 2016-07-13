package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

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
				//nrc.OutputToAtomsAndAdd("curse", string(out), &atomManager, true)
				nrc.OutputToAtomsAndAdd(providerID, string(out), &atomManager, true)
				///*
				tmpAtoms := nrc.OutputToAtoms(string(out), true)

				for _, v := range tmpAtoms {
					fmt.Printf("<-> '%s'\n", v.ID)

					tmpAtom := atomManager.GetEntry(providerID, v.ID)
					if tmpAtom.DoDepCheck {
						fmt.Printf("</> EXEC >%s %s %s %s<\n", "nxatomize", providerID, "depsearch", v.ID)
						//FIXME potentially dangerous if one is careless
						providerCommand := exec.Command("nxatomize", providerID, "depsearch", v.ID)
						_, err := providerCommand.Output()
						if bp.GotError(err) {
							fmt.Printf("<!> ERROR getting deps of: '%s'\n", v.ID)
						} else {
							fmt.Printf("<-> done getting deps of: '%s'\n", v.ID)
						}
					} else {
						fmt.Printf("<-> deps of '%s' already indexed\n", v.ID)
					}
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

	case "depsearch":
		//atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("atoms", bp.StringAtIndex(1, args)))
		//providerDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("providers", bp.StringAtIndex(1, args)))
		fmt.Printf("depsearch\n")

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
					fmt.Printf(">%v<\n", tmpAtom)

					fmt.Println("<~> >" + string(out) + "<")
				}
			} else {
				fmt.Printf("<!> ERROR command of '%s' empty\n", providerID)
			}
		} else {
			fmt.Printf("<!> ERROR provider '%s' not found\n", providerID)
		}

	case "downinfo":
		//atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("atoms", bp.StringAtIndex(1, args)))
		//providerDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", path.Join("providers", bp.StringAtIndex(1, args)))
		fmt.Printf("downinfo\n")

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
	//fmt.Printf("it works !\n")
}
