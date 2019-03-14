package concat

// Doit implements the concat command
func Doit(args []string) string {
	var s string // zero value == ""

	for _, a := range args {
		s += a
	}

	return s
}
