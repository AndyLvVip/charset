package main

import (
	"parser"
	"fetcher"
)

func main() {
	//fetcher.Download(fetcher.BigSmallBank)
	parser.Parse(fetcher.BigSmallBank)
}
