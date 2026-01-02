package cmd

import (
    "fmt"                      // Uses fmt.Printf
    "github.com/spf13/cobra"   // Uses cobra.Command
)

var addCmd = &cobra.Command{
	Use:   "add [website]",
	Short: "Add a website to the block list",
	Long: `Add one or more websites to your focus block list.
These websites will be blocked during focus sessions.

Example:
  focus add youtube.com
  focus add youtube.com facebook.com twitter.com`,
	Args: cobra.MinimumNArgs(1), // Require at least one website
	Run: func(cmd *cobra.Command, args []string) {
		// Implement the add website logic here
		for _, site := range args {
			fmt.Printf("✓ Added %s to block list\n", site)
		}
		
		// TODO: Save to config/storage
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}