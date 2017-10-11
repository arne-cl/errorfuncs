package errorfuncs

func LogFatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogFatalfIfErr(err error, errorMessage string, arguments ...interface{}) {
	if err != nil {
		log.Fatalf(errorMessage, arguments)
	}
}

