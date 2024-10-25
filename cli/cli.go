package cli

import (
	"flag"
	"fmt"
	"os"
)

var Preference_StrongPassword bool = false

func HandleFlags() {
	h := flag.Bool("h", false, "Display help")
	help := flag.Bool("help", false, "Display help")

	s := flag.Bool("s", false, "Generate a password with 71 bits of entropy")
	strong := flag.Bool("strong", false, "Generate a password with 71 bits of entropy")

	v := flag.Bool("v", false, "Print version")
	version := flag.Bool("version", false, "Print version")

	flag.Parse()

	if *s || *strong {
		Preference_StrongPassword = true
	}

	if *v || *version {
		fmt.Println("corepass, version 1.1.1")
		os.Exit(0)
	}

	if *h || *help {
		fmt.Println(`Usage: corepass [flags]

A tool to generate secure and memorably chunkable passwords
flags:
    -h, --help      Display help
    -s, --strong    Generate a password with 71 bits of entropy`)
		os.Exit(0)
	}
}
