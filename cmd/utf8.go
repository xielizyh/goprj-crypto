package cmd

import (
	"log"
	"main/internal/utf8"
	"strings"

	cobra "github.com/spf13/cobra"
)

var utf8Form string
var utf8Msg string
var utf8Reverse bool

// utf8Desc 长的帮助描述
var utf8Desc = strings.Join([]string{
	"utf8编码转换支持的形式如下:",
	"base64",
	"hex",
}, "\n")

var utf8Cmd = &cobra.Command{
	Use:   "utf8",
	Short: "utf8编码转换",
	Long:  utf8Desc,
	Run: func(cmd *cobra.Command, args []string) {
		var value string
		var err error
		if utf8Reverse {
			value, err = utf8.ConvertReverse(utf8Form, utf8Msg)
		} else {
			value, err = utf8.Convert(utf8Form, utf8Msg)
		}
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Printf("输出转换结果: %s", value)
		}
	},
}

func init() {
	// 对命令行选项参数的解析和绑定
	utf8Cmd.Flags().StringVarP(&utf8Form, "form", "f", "hex", "请指定转换形式")
	utf8Cmd.Flags().BoolVarP(&utf8Reverse, "reverse", "r", false, "请指定是否反向转换")
	utf8Cmd.Flags().StringVarP(&utf8Msg, "msg", "m", "hello world", "请输入消息内容")
}
