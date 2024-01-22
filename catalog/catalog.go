package catalog

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Generate(dir string) []TestCatalog {
	fset := token.NewFileSet()
	var testCatalog []TestCatalog
	var (
		scenario    = "#scenario"
		metadata    = "#metadata"
		description = "#description"
	)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//info.Name()
		if info.IsDir() || !strings.HasSuffix(path, "_test.go") {
			return nil
		}
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		ast.Inspect(file, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok || !strings.HasPrefix(fn.Name.Name, "Test") || fn.Doc == nil {
				return true
			}

			comment := fn.Doc.Text()
			placeholders := []string{"# description", "# Description", "# scenario", "# Scenario", "# metadata", "# Metadata", "# MetaData"}
			for _, placeholder := range placeholders {
				comment = strings.ReplaceAll(comment, placeholder, "#"+strings.ToLower(placeholder[2:]))
			}

			indices := map[string]int{
				description: strings.Index(strings.ToLower(comment), description),
				scenario:    strings.Index(strings.ToLower(comment), scenario),
				metadata:    strings.Index(strings.ToLower(comment), metadata),
			}

			if indices[description] != -1 && indices[scenario] != -1 && indices[metadata] != -1 {
				_data := func() TestMetaData {
					_tmd := TestMetaData{}
					_tmd.Dir = filepath.Dir(path)[len(dir)+1:]
					_tmd.File = filepath.Base(path)
					_tmd.Package = file.Name.Name
					return _tmd
				}
				//_testpackage := _data
				_testname := fn.Name.Name
				_description := comment[indices[description]+len(description) : indices[scenario]]
				_scenario := comment[indices[scenario]+len(scenario) : indices[metadata]]
				_metadata := comment[indices[metadata]+len(metadata):]
				end := strings.Index(_metadata, "\n\n")
				if end != -1 {
					_metadata = _metadata[:end]
				}
				_md := convertToMetadata(_metadata)
				testCatalog = append(testCatalog,
					TestCatalog{
						TestMetaData: _data(),
						TestName:     _testname,
						Description:  _description,
						Scenario:     _scenario,
						Types:        _md.Types,
						Components:   _md.Components,
					})
			}

			return true
		})

		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("catalog generated for", len(testCatalog), "tests")
	return testCatalog
}

func convertToMetadata(input string) CatalogMetaData {
	lines := strings.Split(input, "\n")
	md := CatalogMetaData{}

	trimSlice := func(s string) []string {
		parts := strings.Split(s, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		return parts
	}

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(strings.ToLower(parts[0]))
			value := strings.TrimSpace(strings.ToLower(parts[1]))

			switch key {
			case "testtype":
				md.Types = append(md.Types, trimSlice(value)...)
			case "component":
				md.Components = append(md.Components, trimSlice(value)...)
			}
		}
	}

	return md
}

func HostCatalog(testCatalog []TestCatalog) {
	tmpl := getHTMLTemplate()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, testCatalog)
	})
	fmt.Println("Access catalog at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type TestCatalog struct {
	TestMetaData TestMetaData `json:"metadata"`
	TestName     string       `json:"testName"`
	Description  string       `json:"description"`
	Scenario     string       `json:"scenario"`
	Types        []string     `json:"types"`
	Components   []string     `json:"components"`
}

type CatalogMetaData struct {
	Types      []string
	Components []string
}

type TestMetaData struct {
	Dir     string `json:"dir"`
	Package string `json:"package"`
	File    string `json:"file"`
}
