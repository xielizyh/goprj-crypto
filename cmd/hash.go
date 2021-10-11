package cmd

import (
	"log"
	"main/internal/hash"
	"strings"

	cobra "github.com/spf13/cobra"
)

var hashAlg string
var hashMsg string
var hashFile string

// hashDesc 长的帮助描述
var hashDesc = strings.Join([]string{
	"哈希命令支持的算法如下：",
	"sha256",
	"sha512",
	"sm3",
}, "\n")

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "计算消息摘要",
	Long:  hashDesc,
	Run: func(cmd *cobra.Command, args []string) {
		var value []byte
		var err error
		if hashFile == "" {
			value, err = hash.Compute(hashAlg, hashMsg)
		} else {
			value, err = hash.ComputeFile(hashAlg, hashFile)
		}
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Printf("输出消息摘要：%x", value)
		}
	},
}

func init() {
	// 对命令行选项参数的解析和绑定
	hashCmd.Flags().StringVarP(&hashAlg, "alg", "a", "sha256", "请输入哈希算法")
	hashCmd.Flags().StringVarP(&hashMsg, "msg", "m", "hello world", "请输入消息内容")
	hashCmd.Flags().StringVarP(&hashFile, "file", "f", "", "请指定消息文件名")
}
