package cmd

import (
	"fmt"

	"github.com/c4pt0r/tiup/pkg/utils"
	"github.com/c4pt0r/tiup/pkg/version"
	"github.com/spf13/cobra"
)

type showCmd struct {
	*baseCmd
}

func newShowCmd() *showCmd {
	cmdShow := &showCmd{
		newBaseCmd(&cobra.Command{
			Use:   "show",
			Short: "Display global information",
			RunE: func(cmd *cobra.Command, args []string) error {
				fmt.Printf("TiUP Version: %s\n", version.NewTiUPVersion())
				fmt.Printf("TiUP Build: %s\n", version.NewTiUPBuildInfo())
				profileDir, err := utils.GetOrCreateProfileDir()
				if err != nil {
					return err
				}
				fmt.Printf("TiUP home: %s\n", profileDir)

				return showCurrentVersion()
			},
		}),
	}

	return cmdShow
}
