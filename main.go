package main

import (
	"fmt"
	"github.com/ksmanoj/go-test-catalog/catalog"
	"os"
)

func main() {
	var currentDir string
	if len(os.Args) > 1 {
		currentDir = os.Args[1]
	} else {
		currentDir, _ = os.Getwd()
	}

	var testCatalog []catalog.TestCatalog
	testCatalog = catalog.Generate(currentDir)

	// this json data is currently not required. can be use in future use cases to build a dashboard
	catalog.WriteToJson(testCatalog, "test_catalog.json")

	// this step is not required if we are hosting the catalog
	catalog.ConvertToHTML(testCatalog, "output.html")
	fmt.Println("Refer to generated output.html and test_catalog.json")

	// hosts the test catalog locally
	catalog.HostCatalog(testCatalog)
}
