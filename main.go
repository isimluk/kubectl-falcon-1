package main

import (
	"os"

	"github.com/crowdstrike/kubectl-falcon/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	root := cmd.NewCmdFalcon(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
