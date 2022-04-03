package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func startCheck(url string) {
	table.SetHeader([]string{"Host", "Vulnerability Possibility"})
	if !strings.Contains(url, "://") {
		url = fmt.Sprintf("http://%s", url)
	}
	if !strings.HasSuffix(url, "/") {
		url = fmt.Sprintf("%s/", url)
	}
	isvulnerable := heuristicCheck(url)

	var addnew jsondata
	if !isvulnerable {
		var tabledata = []string{url, "NO"}
		addnew.Host = url
		addnew.IsVulnerable = false
		table.Append(tabledata)
	} else {
		var tabledata = []string{url, "YES"}
		addnew.Host = url
		addnew.IsVulnerable = true
		table.Append(tabledata)
	}

	if exploitMode {
		iscomplete, path := runExploit(url)
		addnew.ExploitCompleted = iscomplete
		addnew.PayloadPath = path
	}

	finalData = append(finalData, addnew)
	fmt.Println("")
}

func main() {
	fmt.Println(banner, "\n\n", lackOfArt, "\n ")
	flag.StringVar(&url, "url", "", "Specify a single target URL.")
	flag.StringVar(&targetfile, "file", "", "Specify a file containing a list of target URLs.")
	flag.BoolVar(&exploitMode, "exploit", false, "Turns on exploitation mode.")
	flag.StringVar(&outfile, "outfile", "hunt4spring.json", "Output file name to store results.")
	flag.Parse()

	if len(url) < 1 && len(targetfile) < 1 {
		log.Fatalln("You need to specify a URL or a file containing URLs.")
	}
	if len(url) > 0 {
		startCheck(url)

	} else if len(targetfile) > 0 {
		file, err := os.Open(targetfile)
		if err != nil {
			log.Fatalln("failed to open file named", targetfile)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			startCheck(strings.TrimSpace(scanner.Text()))
		}
	}

	table.Render()

	if !serializeJSON(outfile) {
		log.Fatalln("Could not serialize output to JSON file.")
	}
}
