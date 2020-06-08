package version

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"gopkg.in/yaml.v2"
)

const FlagLong = "long"

// Cmd prints out the application's version information
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the application's version",
		RunE: func(cmd *cobra.Command, args []string) error {
			verInfo := NewVersionInfo()

			if !viper.GetBool(FlagLong) {
				v := verInfo.Version
				if verInfo.NetworkType != "" {
					v = v + "-" + verInfo.NetworkType
				}
				fmt.Println(v)
				return nil
			}

			var bz []byte
			var err error

			switch viper.GetString(cli.OutputFlag) {
			case "json":
				bz, err = json.Marshal(verInfo)
			default:
				bz, err = yaml.Marshal(&verInfo)
			}

			if err != nil {
				return err
			}

			_, err = fmt.Println(string(bz))
			return err
		},
	}

	cmd.Flags().Bool(FlagLong, false, "Print long version information")
	return cmd
}
