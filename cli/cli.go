package cli

import (
	"flag"
	"fmt"
	"os"
)

var Preference_StrongPassword bool = false

func HandleFlags() {
	if len(os.Args) >= 2 {
		h := flag.Bool("h", false, "Display help")
		help := flag.Bool("help", false, "Display help")
		s := flag.Bool("s", false, "Generate stronger password")
		strong := flag.Bool("strong", false, "Generate stronger password")

		flag.Parse()

		if *s || *strong {
			Preference_StrongPassword = true
			return
		}

		if *h || *help {
			fmt.Println("Usage: core-password [flags]")
			fmt.Println("A tool to generate secure and memorably chunkable passwords")
			fmt.Println("flags:")
			fmt.Println("    -h, --help       Display help")
			fmt.Println("    -s, --strong    Generate a password with 71 bits of entropy")
			os.Exit(0)
		}
	}
}
