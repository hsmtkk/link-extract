package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/hsmtkk/link-extract/download"
	"github.com/hsmtkk/link-extract/html"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:  "link-extract url suffix",
	Args: cobra.ExactArgs(2),
	Run:  run,
}

func run(cmd *cobra.Command, args []string) {
	url := args[0]
	suffix := args[1]
	htmlString, err := download.Download(url)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(htmlString)
	links, err := html.Parse(htmlString)
	if err != nil {
		log.Fatal(err)
	}
	for _, link := range links {
		if strings.HasSuffix(link, suffix) {
			fmt.Println(link)
		}
	}
}
