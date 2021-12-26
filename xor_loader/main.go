package main

import (
	"fmt"
	"io/ioutil"
)

func xor(data []byte,key []byte) *[]byte{
	var encdata = make([]byte,len(data))
	for i,j :=0,0;i<len(data);i++{
		encdata[i] = data[i] ^ key[j]
		j = (j+1)%len(key)
	}
	return &encdata
}

func main(){
	filedata,_:= ioutil.ReadFile("home_payload.bin")
	key  := []byte("wulala")

	fmt.Println(key)


	/* ====== 分割线 ====== */
	encdata := *xor(filedata,key)
	ioutil.WriteFile("encpayload.bin",encdata,0666)

	filedata,_= ioutil.ReadFile("encpayload.bin")
	decdata := *xor(filedata,key)
	ioutil.WriteFile("decpayload.bin",decdata,0666)
}
