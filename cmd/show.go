package cmd

import (
    "fmt"                      
    "github.com/spf13/cobra"   
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current block list",
	Long: `Display all websites currently in your focus block list.

Example:
  focus show`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Load from actual storage
		
		
		blockedSites := []string{
			"youtube.com",
			"facebook.com",
			"twitter.com",
		}
		
		if len(blockedSites) == 0 {
			fmt.Println("✓ Block list is empty")
			return
		}
		
		fmt.Println("✓ Blocked websites:")
		for i, site := range blockedSites {
			fmt.Printf("  %d. %s\n", i+1, site)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}