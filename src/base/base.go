package base

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}
