package cmd

import "github.com/spf13/cobra"

// 根命令
var rootCmd = &cobra.Command{}

// Execute 执行根命令
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 注册Hash命令
	rootCmd.AddCommand(hashCmd)
	// 注册Mac命令
	rootCmd.AddCommand(macCmd)
}
