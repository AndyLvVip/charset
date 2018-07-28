package parser

import (
	"fetcher"
	"os"
	"base"
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strings"
)

func Parse(bankInfo *fetcher.BankInfo) {
	file, err := os.Open(bankInfo.FileName)
	base.CheckErr(err)
	defer file.Close()
	r := simplifiedchinese.GBK.NewDecoder().Reader(file)
	reader := bufio.NewReader(r)
	line, err := reader.ReadString('\n')
	base.CheckErr(err)
	fmt.Println(strings.TrimSpace(line))

	line, err = reader.ReadString('\n')
	base.CheckErr(err)
	fmt.Println(strings.TrimSpace(line))
}
