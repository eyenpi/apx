package cmd

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2023
	Description: Apx is a wrapper around multiple package managers to install packages and run commands inside a managed container.
*/

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/apx/core"
	"github.com/vanilla-os/orchid/cmdr"
)

func NewAutoRemoveCommand() *cmdr.Command {
	cmd := cmdr.NewCommand(
		"autoremove",
		apx.Trans("autoremove.long"),
		apx.Trans("autoremove.short"),
		autoRemove,
	).WithBoolFlag(
		cmdr.NewBoolFlag(
			"all",
			"a",
			apx.Trans("apx.allFlag"),
			false,
		))
	return cmd
}

func autoRemove(cmd *cobra.Command, args []string) error {
	if cmd.Flag("nix").Changed {
		return errors.New(apx.Trans("apx.notForNix"))

	}
	if cmd.Flag("all").Changed {
		if err := core.ApplyForAll("autoremove", []string{}); err != nil {
			return err
		}
		return nil
	}

	command := append([]string{}, container.GetPkgCommand("autoremove")...)
	command = append(command, args...)

	container.Run(command...)

	return nil
}
