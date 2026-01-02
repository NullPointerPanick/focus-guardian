package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
    Use:   "remove [website]",
    Short: "Remove websites from the block list",
    Long:  `Remove one or more websites from your focus block list.`,
    Args:  cobra.MinimumNArgs(1),
    
	Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 1 && args[0] == "all" {
			removeAllCmd.Run(cmd, args)
            return
        }
        for _, site := range args {
            // Implement the remove website logic here
			
			fmt.Printf("✓ Removed %s from block list\n", site)
        }
    },
}

var removeAllCmd = &cobra.Command{
    Use:   "all",
    Short: "Remove all websites",
    Args:  cobra.NoArgs,
    
	Run: func(cmd *cobra.Command, args []string) {
        // Implement the remove all websites logic here
		
		fmt.Println("✓ Removed all websites from block list")
    },
}

func init() {
    rootCmd.AddCommand(removeCmd)
    removeCmd.AddCommand(removeAllCmd)
}