package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golistic/gengrpc"
)

type config struct {
	module              string
	sourcePackage       string
	pathInSourcePackage string
	destPath            string
}

func main() {
	c, err := getConfig()
	if err != nil {
		exitWithErr(err)
	}

	source, err := gengrpc.NewGoPackage(c.sourcePackage, c.pathInSourcePackage)
	if err != nil {
		exitWithErr(err)
	}

	if err := gengrpc.Generate(c.module, source, c.destPath); err != nil {
		exitWithErr(err)
	}
}

func getConfig() (*config, error) {
	c := &config{}
	flag.StringVar(&c.module, "module", "",
		"Go package which will hold generated protocol buffer files.")
	flag.StringVar(&c.sourcePackage, "source", "",
		"Go package (without path) which provides the contract files (*.proto).")
	flag.StringVar(&c.pathInSourcePackage, "source-path", "",
		"Path in source package holding contract files (*.proto).")
	flag.StringVar(&c.destPath, "dest-path", ".",
		"Path where generate code will be written to.")

	flag.Parse()

	if c.module == "" {
		return nil, fmt.Errorf("-module flag is required")
	}
	c.module = strings.TrimSuffix(c.module, "/")

	if c.sourcePackage == "" {
		return nil, fmt.Errorf("-source flag is required")
	}
	c.sourcePackage = strings.TrimSuffix(c.sourcePackage, "/")

	if c.pathInSourcePackage != "" {
		c.pathInSourcePackage = strings.Trim(c.pathInSourcePackage, "/")
	}

	return c, nil
}

func exitWithErr(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
