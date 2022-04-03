package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type jsondata struct {
	Host             string `json:"host"`
	IsVulnerable     bool   `json:"is_vulnerable"`
	ExploitCompleted bool   `json:"exploit_completed,omitempty"`
	PayloadPath      string `json:"payload_path,omitempty"`
}

var (
	url         string
	targetfile  string
	outfile     string
	exploitMode bool
	finalData   []jsondata
	banner      = `
 _    _             _   _  _   _____            _             
| |  | |           | | | || | / ____|          (_)            
| |__| |_   _ _ __ | |_| || || (___  _ __  _ __ _ _ __   __ _ 
|  __  | | | | '_ \| __|__   _\___ \| '_ \| '__| | '_ \ / _  |
| |  | | |_| | | | | |_   | | ____) | |_) | |  | | | | | (_| |
|_|  |_|\__,_|_| |_|\__|  |_||_____/| .__/|_|  |_|_| |_|\__, |
                                    | |                  __/ |
                                    |_|                 |___/ 
																							   
`
	lackOfArt = ` [+] Hunt4Spring by RedHunt Labs - A Modern Attack Surface (ASM) Management Company
  [+] Author: Umair Nehri (RHL Research Team)
  [+] Continuously Track Your Attack Surface using https://redhuntlabs.com/nvadr.`

	table = tablewriter.NewWriter(os.Stdout)
)
