package cmd

import (
	"log"
	"main/internal/mac"
	"strings"

	cobra "github.com/spf13/cobra"
)

var macAlg string
var macKey string
var macMsg string
var macFile string

// desc 长的帮助描述
var macDesc = strings.Join([]string{
	"消息认证码命令支持的算法如下：",
	"hmac-sha256",
}, "\n")

var macCmd = &cobra.Command{
	Use:   "mac",
	Short: "计算消息认证码",
	Long:  macDesc,
	Run: func(cmd *cobra.Command, args []string) {
		var value []byte
		var err error
		if macFile == "" {
			value, err = mac.Compute(macAlg, macKey, macMsg)
		} else {
			value, err = mac.ComputeFile(macAlg, macKey, macFile)
		}
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Printf("输出消息认证码：%x", value)
		}
	},
}

func init() {
	// 对命令行选项参数的解析和绑定
	macCmd.Flags().StringVarP(&macAlg, "alg", "a", "hmac-sha256", "请输入算法")
	macCmd.Flags().StringVarP(&macKey, "key", "k", "", "请输入密钥(hex形式)")
	macCmd.Flags().StringVarP(&macMsg, "msg", "m", "hello world", "请输入消息内容")
	macCmd.Flags().StringVarP(&macFile, "file", "f", "", "请指定消息文件名")
}
