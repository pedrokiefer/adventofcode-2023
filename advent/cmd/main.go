package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

//go:embed templates/main_template.tpl
var mainFile string

//go:embed templates/test_template.tpl
var testFile string

var day = flag.Int("day", 0, "day to generate")
var name = flag.String("name", "", "name of the struct")

func main() {
	flag.Parse()
	if *day == 0 {
		log.Fatal("must provide a day")
	}
	if *name == "" {
		log.Fatal("must provide a name")
	}

	values := make(map[string]interface{})
	values["Day"] = *day
	values["Name"] = *name

	mainTpl, err := template.New("main").Parse(mainFile)
	if err != nil {
		log.Fatal(err)
	}
	testTpl, err := template.New("test").Parse(testFile)
	if err != nil {
		log.Fatal(err)
	}

	main := processTemplate(mainTpl, values)
	test := processTemplate(testTpl, values)

	os.Mkdir(fmt.Sprintf("day%d", *day), 0755)
	mainFile, err := os.Create(fmt.Sprintf("day%d/main.go", *day))
	if err != nil {
		log.Fatal(err)
	}
	defer mainFile.Close()
	mainFile.WriteString(main)

	testFile, err := os.Create(fmt.Sprintf("day%d/day%d_test.go", *day, *day))
	if err != nil {
		log.Fatal(err)
	}
	defer testFile.Close()
	testFile.WriteString(test)

}

func processTemplate(t *template.Template, vars map[string]interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}
