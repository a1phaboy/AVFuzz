package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	FileName string
	EncryptKey string
	OutputFile string
	Stdin bool
}

func SetOptions() *Options{
	options := &Options{}
	flag.StringVar(&options.FileName,"f","","raw文件")
	flag.StringVar(&options.EncryptKey,"k","a1phaboy","设置加密密钥")
	flag.StringVar(&options.OutputFile,"o","xor.exe","生成文件的名字，路径默认在当前文件夹")
	flag.Parse()
	options.Stdin = hasStdin()
	Showbanner()
	if len(options.FileName) == 0  {
		flag.Usage()
		os.Exit(0)
	}
	if options.FileName != "" && !FileExists(options.FileName) {
		ErrorLog(fmt.Sprintf("[-] 文件 %s 不存在!\n", options.FileName))
		os.Exit(0)
	}

	return options
}
func hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	}
	return true
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
