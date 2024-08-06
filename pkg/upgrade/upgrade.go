package upgrade

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openshift/rosa/pkg/machinepool"
)

var (
	ErrInvalidMachinePoolIdentifier = fmt.Errorf("expected a valid identifier for the machine pool")
	ErrMissingMachinePoolIdentifier = fmt.Errorf("machine pool identifier is missing")
	machinepoolFlagName             = "machinepool"
)

func NewUpgradeArgsFunction(flagSource bool) func(cmd *cobra.Command, argv []string) error {
	return func(cmd *cobra.Command, argv []string) error {
		var machinepoolID string
		var err error

		if flagSource {
			flags := cmd.Flags()
			if !cmd.Flags().Changed(machinepoolFlagName) {
				return nil
			}
			if machinepoolID, err = flags.GetString(machinepoolFlagName); err != nil {
				return err
			}
		} else {
			if len(argv) != 1 {
				return ErrMissingMachinePoolIdentifier
			}
			machinepoolID = argv[0]
		}

		if !machinepool.MachinePoolKeyRE.MatchString(machinepoolID) {
			return ErrInvalidMachinePoolIdentifier
		}
		return nil
	}
}
