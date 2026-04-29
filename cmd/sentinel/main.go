/*
©Anil Kumar | 2026
main.go

Entry point for the sentinel CLI
*/

package main

import (
	"github.com/Anilokumar/linux-persistence-scanner/internal/cli"
	_ "github.com/Anilokumar/linux-persistence-scanner/internal/scanner"
)

func main() {
	cli.Execute()
}
