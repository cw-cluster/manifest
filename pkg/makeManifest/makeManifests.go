package makeManifest

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"os"

	"gopkg.in/yaml.v2"
)

type templateData struct {
	valuePath string `default:"values.yaml"`
	Values    map[string]interface{}
}

// Rendering Manifest
func (t *templateData) renderManifest(tmplData []byte) bytes.Buffer {
	var rendered bytes.Buffer

	tmpl, err := template.New("manifest").Parse(string(tmplData))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		panic(err)
	}

	err = tmpl.Execute(&rendered, t)
	if err != nil {
		fmt.Println("Error executing template:", err)
		panic(err)
	}

	// err = os.WriteFile("base/"+manifest, rendered.Bytes(), 0644)
	// if err != nil {
	// 	fmt.Println("Error Create Yaml:", err)
	// }
	return rendered
}

// Load, Unmarshal values.yaml and update to field valuesData
func (t *templateData) loadValuesData() {
	valuesData, err := os.ReadFile(t.valuePath)
	if err != nil {
		fmt.Println("Error reading values.yaml:", err)
		return
	}

	err = yaml.Unmarshal(valuesData, &t.Values)
	if err != nil {
		fmt.Println("Error parsing values.yaml:", err)
		return
	}
}

func getFileList(s string) []fs.DirEntry {
	flist, err := os.ReadDir(s)
	if err != nil {
		fmt.Println(err)
	}
	return flist
}

// Read Template Manifest File in PATH
func readTmpl(path string, manifest string) []byte {
	tmplData, err := os.ReadFile(path + manifest)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		panic(err)
	}
	return tmplData
}

func NewTemplateDataConstructor(inputPath string) *templateData {
	return &templateData{
		valuePath: inputPath,
	}
}

func MakeManifest(v string, b string) {
	t := NewTemplateDataConstructor(v)
	t.loadValuesData()
	p := NewTempldatePathConstructor(t.Values["template"].(string))

	_, err := os.Stat("base")
	if err != nil {
		if os.IsExist(err) {
			os.Mkdir("base", 0744)
			for _, i := range getFileList(p.basePath) {
				d := t.renderManifest(readTmpl(p.basePath, i.Name()))
				os.WriteFile("base/"+i.Name(), d.Bytes(), 0744)
			}
		}
	} else {
		fmt.Println("base already exists")
	}

	_, err = os.Stat("overlay/" + b)
	if err != nil {
		if os.IsExist(err) {
			os.MkdirAll("overlay/dev", 0744)
			for _, i := range getFileList(p.overlayPath) {
				d := t.renderManifest(readTmpl(p.overlayPath, i.Name()))
				os.WriteFile("overlay/"+b+"/"+i.Name(), d.Bytes(), 0744)
			}
		}
	} else {
		fmt.Println("overlay/" + b + " already exists")
	}

}
