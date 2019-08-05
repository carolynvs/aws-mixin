package main

import (
	"github.com/deislabs/porter-aws/pkg/aws"
	"github.com/spf13/cobra"
)

func buildUninstallCommand(m *aws.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Execute the uninstall functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
