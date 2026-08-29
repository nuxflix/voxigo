package main

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// versionCmd is `voxigo version`.
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the voxigo module version",
		Run: func(cmd *cobra.Command, _ []string) {
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), moduleVersion())
		},
	}
}

// moduleVersion reports the module version from build info, or "devel" when
// the binary was built from a working tree that is not a published module.
func moduleVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok || info == nil {
		return "devel"
	}
	if v := info.Main.Version; v != "" && v != "(devel)" {
		return v
	}
	return "devel"
}
