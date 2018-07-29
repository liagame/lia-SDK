package cmd

import (
	"fmt"
	"github.com/liagame/lia-cli/internal"
	"github.com/liagame/lia-cli/internal/config"
	"github.com/liagame/lia-cli/internal/settings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var resetTrackingId bool
var analyticsOptIn bool
var analyticsOptOut bool

var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Views the user's settings",
	Long:  `Views the user's settings.'`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.UpdateIfTime(true)

		if resetTrackingId {
			oldTrackingId := viper.GetString("trackingId")
			newTrackingId := settings.GenerateTrackingId()
			viper.Set("trackingId", newTrackingId)
			viper.WriteConfig() // writes to the lia settings file
			fmt.Printf("Replacing old trackingId: %s with: %s\n", oldTrackingId, newTrackingId)
		}

		if analyticsOptIn {
			viper.Set("analyticsAllow", true)
			viper.Set("analyticsAllowedVersion", config.VERSION)
			viper.WriteConfig() // writes to the lia settings file
			fmt.Printf("Opting in to anonymous analytics usage reporting.\n")
		}

		if analyticsOptOut {
			viper.Set("analyticsAllow", false)
			viper.Set("analyticsAllowedVersion", config.VERSION)
			viper.WriteConfig() // writes to the lia settings file
			fmt.Printf("Opting out of anonymous analytics usage reporting.\n")
		}

		fmt.Printf("TrackingId: %s\n", viper.GetString("trackingId"))

		var analyticsAllowed string
		switch viper.Get("analyticsAllow") {
		case true:
			analyticsAllowed = "true"
		case false:
			analyticsAllowed = "false"
		case nil:
			analyticsAllowed = "not set"
		}

		fmt.Printf("Allow analytics: %s\n", analyticsAllowed)
		fmt.Printf("Allow analytics for version: %s\n", viper.GetString("analyticsAllowedVersion"))

	},
}

func init() {
	rootCmd.AddCommand(settingsCmd)

	settingsCmd.Flags().BoolVarP(&resetTrackingId, "reset-tracking-id", "t", false, "Reset anonymous tracking ID")
	settingsCmd.Flags().BoolVarP(&analyticsOptIn, "analytics-opt-in", "i", false, "Opt-in for anonymous analytics usage report")
	settingsCmd.Flags().BoolVarP(&analyticsOptOut, "analytics-opt-out", "o", false, "Opt-out from anonymous analytics usage report")
}