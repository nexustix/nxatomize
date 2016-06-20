package main

import (
	"fmt"
	"os/user"

	bp "github.com/nexustix/boilerplate"
	nrc "github.com/nexustix/nxReplicatorCommon"
)

func main() {
	version := "V.0-1-0"
	fmt.Printf("<-> NxAtomize Version: %s\n", version)

	//args := os.Args

	usr, err := user.Current()
	bp.FailError(err)
	workingDir := usr.HomeDir
	//fmt.Printf("<-> Working dir: %s\n", workingDir)
	fmt.Printf("<-> Home dir: %s\n", workingDir)

	//atomDir :=
	nrc.InitWorkFolder(workingDir, ".nxreplicator", "atoms")
	//providerDir :=
	nrc.InitWorkFolder(workingDir, ".nxreplicator", "providers")

	/*
		switch args[0] {
		case "condition":

		}
	*/
	fmt.Printf("it works !\n")
}
