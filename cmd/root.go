package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "focus",
	Short: "Focus Guardian - Block distracting websites",
	Long: `Focus Guardian helps you stay productive by blocking 
distracting websites during focus sessions.

Example:
  focus add youtube.com 
  focus remove facebook.com
  focus remove all
  focus show`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

