package catalog

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

func ConvertToHTML(data []TestCatalog, fileName string) {
	tmpl := getHTMLTemplate()
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func getHTMLTemplate() *template.Template {
	var templates = []string{"templates/test_catalog_template.html"}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	return tmpl
}

func WriteToJson(testData []TestCatalog, fileName string) {
	jsonData, err := json.Marshal(testData)
	if err != nil {
		panic(err)
	}
	// Write JSON data to a file
	err = ioutil.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func ReadFromJson(fileName string) ([]TestCatalog, error) {
	jsonData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Convert JSON data to tests struct
	var tests []TestCatalog
	err = json.Unmarshal(jsonData, &tests)
	if err != nil {
		return nil, err
	}
	return tests, nil
}
