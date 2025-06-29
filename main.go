package main

import (
	"Github-User-Activity/utils"
	"flag"
	"fmt"
	"os"
)

func main() {
	var showEmails = flag.Bool("show-emails", false, "Show author emails in output")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: Github-User-Activity [--show-emails] <username>\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	userName := flag.Arg(0)
	fmt.Println(`
 ██████╗ ██╗  ██╗    ███████╗████████╗ █████╗ ██╗     ██╗  ██╗███████╗██████╗ 
██╔════╝ ██║  ██║    ██╔════╝╚══██╔══╝██╔══██╗██║     ██║ ██╔╝██╔════╝██╔══██╗
██║  ███╗███████║    ███████╗   ██║   ███████║██║     █████╔╝ █████╗  ██████╔╝
██║   ██║██╔══██║    ╚════██║   ██║   ██╔══██║██║     ██╔═██╗ ██╔══╝  ██╔══██╗
╚██████╔╝██║  ██║    ███████║   ██║   ██║  ██║███████╗██║  ██╗███████╗██║  ██║
 ╚═════╝ ╚═╝  ╚═╝    ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
                                                                              
	`)
	utils.ProcessJsonArray(userName, *showEmails)
}
