package cmd

import (
	"log"
	"main/internal/hash"
	"strings"

	cobra "github.com/spf13/cobra"
)

var hashAlg string
var message string

// desc 长的帮助描述
var desc = strings.Join([]string{
	"哈希命令支持的算法如下：",
	"sha256",
	"sha512",
}, "\n")

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "计算消息摘要",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		hash := hash.Compute(hashAlg, message)
		log.Printf("输出消息摘要：%x", hash)
	},
}

func init() {
	// 对命令行选项参数的解析和绑定
	hashCmd.Flags().StringVarP(&hashAlg, "alg", "a", "sha256", "请输入哈希算法")
	hashCmd.Flags().StringVarP(&message, "msg", "m", "hello world", "请输入消息内容")
}
