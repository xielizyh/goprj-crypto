package cmd

import "github.com/spf13/cobra"

// 根命令
var rootCmd = &cobra.Command{}

// Execute 执行根命令
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 注册Utf8命令
	rootCmd.AddCommand(utf8Cmd)
	// 注册Hash命令
	rootCmd.AddCommand(hashCmd)
	// 注册Mac命令
	rootCmd.AddCommand(macCmd)
	// 注册Enc命令
	rootCmd.AddCommand(encCmd)
}
