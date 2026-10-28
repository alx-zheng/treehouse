package cmd

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/kunchenguid/treehouse/v3/internal/updater"
	"github.com/spf13/cobra"
)

var runningFromNixStore = updater.RunningFromNixStore

// forkBuild names a locally patched build (set with -ldflags "-X .../cmd.forkBuild=<name>").
// Self-update would replace it with the stock release and silently drop the patches.
var forkBuild = ""

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update treehouse to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version == "dev" {
			fmt.Println("Skipping update: running a dev build")
			return nil
		}

		if forkBuild != "" {
			return fmt.Errorf("this is a patched build (%s); a self-update would drop the patches. Rebase the fork onto the new release and rebuild it instead", forkBuild)
		}

		if runningFromNixStore() {
			return errors.New("self-update is not available for Nix installs; upgrade through Nix")
		}

		fmt.Println("Checking for updates...")
		result, err := updater.CheckLatest(version)
		if err != nil {
			return fmt.Errorf("checking for updates: %w", err)
		}

		if !result.UpdateAvailable {
			fmt.Printf("treehouse is up to date (%s)\n", version)
			return nil
		}

		fmt.Printf("Updating %s → %s...\n", version, result.LatestVersion)
		if err := updater.Apply(result); err != nil {
			return fmt.Errorf("applying update: %w", err)
		}

		color.Green("Successfully updated treehouse %s → %s", version, result.LatestVersion)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
