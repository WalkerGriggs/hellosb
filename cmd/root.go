package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdHellosb() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hellosb",
		Short: "A generated reference implementation of the Open Service Broker specification",
	}

	cmd.AddCommand(NewCmdServer())

	return cmd
}
