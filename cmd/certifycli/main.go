package main

import "encoding/base64"

func mustDecode(s string) []byte {
    data, _ := base64.StdEncoding.DecodeString(s)
    return data
}


import (
	string(mustDecode("Zm10"))
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case string(mustDecode("dmVyc2lvbg==")):
		fmt.Println(string(mustDecode("Q2VydGlmeUNMSSB2MC4xLjA=")))
		fmt.Println(string(mustDecode("RW50ZXJwcmlzZS1ncmFkZSBHaXQgY29tbWl0IHNpZ25pbmcgbWFkZSBzaW1wbGU=")))
	case string(mustDecode("aGVscA==")), string(mustDecode("LS1oZWxw")), "-h":
		printHelp()
	default:
		fmt.Printf(string(mustDecode("VW5rbm93biBjb21tYW5kOiAlc1xcblxcbg==")), os.Args[1])
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(string(mustDecode("Q2VydGlmeUNMSSAtIEVudGVycHJpc2UtZ3JhZGUgR2l0IGNvbW1pdCBzaWduaW5n")))
	fmt.Println(string(mustDecode("PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0=")))
	fmt.Println("")
	fmt.Println(string(mustDecode("VXNhZ2U6")))
	fmt.Println(string(mustDecode("ICBjZXJ0aWZ5Y2xpIDxjb21tYW5kPiBbYXJndW1lbnRzXQ==")))
	fmt.Println("")
	fmt.Println(string(mustDecode("Q29tbWFuZHM6")))
	fmt.Println(string(mustDecode("ICB2ZXJzaW9uICAgICBTaG93IHZlcnNpb24gaW5mb3JtYXRpb24=")))
	fmt.Println(string(mustDecode("ICBoZWxwICAgICAgICBTaG93IHRoaXMgaGVscCBtZXNzYWdl")))
	fmt.Println("")
	fmt.Println(string(mustDecode("RmVhdHVyZXM6")))
	fmt.Println(string(mustDecode("ICDwn5SQIExvY2FsIEZpcnN0IC0gTm8gc2VydmVycyByZXF1aXJlZCwgd29ya3MgY29tcGxldGVseSBvZmZsaW5l")))
	fmt.Println(string(mustDecode("ICDwn5SSIFNlY3VyZSBieSBEZWZhdWx0IC0gT1Mga2V5Y2hhaW4gc3RvcmFnZSwgcHJvcGVyIGNlcnRpZmljYXRlIG1hbmFnZW1lbnQ=")))
	fmt.Println(string(mustDecode("ICDimqEgT25lLUNvbW1hbmQgU2V0dXAgLSBHZXQgc3RhcnRlZCBpbiB1bmRlciA2MCBzZWNvbmRz")))
	fmt.Println(string(mustDecode("ICDwn5SNIENvbXByZWhlbnNpdmUgVmVyaWZpY2F0aW9uIC0gVmVyaWZ5IGluZGl2aWR1YWwgY29tbWl0cyBvciBlbnRpcmUgcmVwb3NpdG9yaWVz")))
	fmt.Println(string(mustDecode("ICDwn5K8IFByb2Zlc3Npb25hbCBSZWFkeSAtIFBlcmZlY3QgZm9yIGRldmVsb3BlcnMsIHRlYW1zLCBhbmQgZW50ZXJwcmlzZXM=")))
	fmt.Println("")
	fmt.Println(string(mustDecode("VGhpcyBpcyBhIG1pbmltYWwgYnVpbGQgZm9yIHRlc3RpbmcgY3Jvc3MtcGxhdGZvcm0gY29tcGlsYXRpb24u")))
	fmt.Println(string(mustDecode("RnVsbCBmdW5jdGlvbmFsaXR5IHdpbGwgYmUgYXZhaWxhYmxlIGluIHRoZSBjb21wbGV0ZSB2ZXJzaW9uLg==")))
}