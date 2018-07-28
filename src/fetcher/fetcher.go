package fetcher

import (
	"net/http"
	"base"
	"os"
	"io"
)

type BankInfo struct {
	Url      string
	FileName string
}

var BigSmallBank = &BankInfo{Url:"https://ebank.cgbchina.com.cn/corporbank/superEbankNoDownload.jsp?pms=true", FileName:"BigSmallBank.txt"}
var SuperBank = &BankInfo{Url:"https://ebank.cgbchina.com.cn/corporbank/superEbankNoDownload.jsp", FileName:"SuperBank.txt"}

func Download(bankInfo *BankInfo) {
	resp, err := http.Get(bankInfo.Url)
	base.CheckErr(err)
	defer resp.Body.Close()
	file, err := os.Create(bankInfo.FileName)
	base.CheckErr(err)
	defer file.Close()

	data := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(data)
		if nil != err {
			if io.EOF == err {
				break
			}
			panic(err)
		}

		file.Write(data[0:n])
	}
}
