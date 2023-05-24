package cmd

import (
	"flag"
	"log"
	"math"
)

type InputFlags struct {
	GetReltivePath         bool
	GetOnlyDir             bool
	OnlyTillLevel          int
	SortByLastModifiedTime bool
	GetOnlyPermissions     bool
	GetInXML               bool
	GetInJson              bool
	Root                   string
	NoIndentation          bool
}
//GetInput parses user provided flags and returns an struct of InputFlags
func GetInput() InputFlags {
	flagF := flag.Bool("f", false, "pass -w to get relative path")
	flagD := flag.Bool("d", false, "pass -d to get only directories")
	flagL := flag.Int("l", math.MaxInt64, "pass -l to get files under certain level")
	flagP := flag.Bool("p", false, "pass -p to get permissions")
	flagT := flag.Bool("t", false, "pass -t to sort files according to last modified date")
	flagX := flag.Bool("x", false, "pass -x to get output in XML form")
	flagJ := flag.Bool("j", false, "pass -j to get output in JSON from")
	flagI := flag.Bool("i", false, "pass -i to get output without indentation")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("Invalid arguments")

	}

	if *flagL < 1 {
		log.Fatal("Level cannot be less than 1")
	}

	inputFlags := InputFlags{
		GetReltivePath:         *flagF,
		GetOnlyDir:             *flagD,
		OnlyTillLevel:          *flagL,
		SortByLastModifiedTime: *flagT,
		GetOnlyPermissions:     *flagP,
		GetInXML:               *flagX,
		GetInJson:              *flagJ,
		Root:                   args[0],
		NoIndentation:          *flagI,
	}

	return inputFlags

}
