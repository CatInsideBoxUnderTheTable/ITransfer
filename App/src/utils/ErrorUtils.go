package utils

func PanicIfErr(val any, panicMessage string) {
	if val != nil {
		panic(panicMessage)
	}
}
