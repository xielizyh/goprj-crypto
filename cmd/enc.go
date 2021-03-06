package cmd

import (
	"log"
	"main/internal/enc"
	"strings"

	cobra "github.com/spf13/cobra"
)

var encAlg string
var encKey string
var encMsg string
var encInFile string
var encOutFile string
var encMode string

// desc 长的帮助描述
var encDesc = strings.Join([]string{
	"对称加解密命令支持的算法如下:",
	"aes-128-ecb",
	"aes-128-cbc",
	"aes-128-ctr",
	"sm4-ecb",
	"sm4-cbc",
	"sm4-ctr",
}, "\n")

var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "对称加解密",
	Long:  encDesc,
	Run: func(cmd *cobra.Command, args []string) {
		var value []byte
		var err error
		if encInFile == "" && encOutFile == "" {
			if encMode == "d" {
				value, err = enc.CipherDecrypt(encAlg, encKey, encMsg)
			} else {
				value, err = enc.CipherEncrypt(encAlg, encKey, encMsg)
			}
		} else {
			if encMode == "d" {
				err = enc.CipherDecryptFile(encAlg, encKey, encInFile, encOutFile)
			} else {
				err = enc.CipherEncryptFile(encAlg, encKey, encInFile, encOutFile)
			}
		}
		if err != nil {
			log.Fatalln(err)
		} else {
			if encMode == "d" {
				if encInFile == "" && encOutFile == "" {
					log.Printf("输出明文: %x", value)
				} else {
					log.Printf("解密成功: [%s->%s]", encInFile, encOutFile)
				}
			} else {
				if encInFile == "" && encOutFile == "" {
					log.Printf("输出密文: %x", value)
				} else {
					log.Printf("加密成功: [%s->%s]", encInFile, encOutFile)
				}
			}
		}
	},
}

func init() {
	// 对命令行选项参数的解析和绑定
	encCmd.Flags().StringVarP(&encAlg, "alg", "a", "", "请输入算法")
	encCmd.Flags().StringVarP(&encKey, "key", "k", "", "请输入密钥(hex形式)")
	encCmd.Flags().StringVarP(&encMsg, "msg", "m", "", "请输入消息内容(hex形式)")
	encCmd.Flags().StringVarP(&encInFile, "in", "i", "", "请指定输入消息文件名")
	encCmd.Flags().StringVarP(&encOutFile, "out", "o", "", "请指定输出消息文件名")
	encCmd.Flags().StringVarP(&encMode, "pattern", "p", "e", "请指定加密(e)或解密(d)")
}
