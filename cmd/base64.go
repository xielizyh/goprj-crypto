package cmd

import (
	"log"
	"main/internal/base64"
	"strings"

	cobra "github.com/spf13/cobra"
)

var base64Mode string
var base64Msg string

// base64Desc 长的帮助描述
var base64Desc = strings.Join([]string{
	"Base64编码转换",
}, "\n")

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64编码转换",
	Long:  base64Desc,
	Run: func(cmd *cobra.Command, args []string) {
		var value string
		var err error
		if base64Mode == "d" {
			value, err = base64.Decode(base64Msg)
		} else {
			value, err = base64.Encode(base64Msg)
		}
		if err != nil {
			log.Fatalln(err)
		} else {
			if base64Mode == "d" {
				log.Printf("输出base64解码后内容: %s", value)
			} else {
				log.Printf("输出base64编码后内容: %s", value)
			}
		}
	},
}

func init() {
	// 对命令行选项参数的解析和绑定
	base64Cmd.Flags().StringVarP(&base64Mode, "pattern", "p", "e", "请指定编码(e)或解码(d)")
	base64Cmd.Flags().StringVarP(&base64Msg, "msg", "m", "hello world", "请输入消息内容")
}
