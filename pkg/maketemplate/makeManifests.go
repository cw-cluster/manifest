package maketemplate

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"os"

	"gopkg.in/yaml.v2"
)

type templateData struct {
	Values map[string]interface{}
}

func (t *templateData) renderBaseManifest(path string, manifest string) {
	tmplData, err := os.ReadFile(path + manifest)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	tmpl, err := template.New("manifest").Parse(string(tmplData))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	var rendered bytes.Buffer
	err = tmpl.Execute(&rendered, t)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	_, err = os.Stat("/base")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("base", 0744)
		} else {
			fmt.Println(err)
		}
	}

	err = os.WriteFile("base/"+manifest, rendered.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error Create Yaml:", err)
	}
}

func (t *templateData) renderOverlayManifest(path string, manifest string) {
	tmplData, err := os.ReadFile(path + manifest)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	tmpl, err := template.New("manifest").Parse(string(tmplData))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	var rendered bytes.Buffer
	err = tmpl.Execute(&rendered, t)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	_, err = os.Stat("/overlay")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("overlay", 0744)
		} else {
			fmt.Println(err)
		}
	}

	err = os.WriteFile("overlay/dev"+manifest, rendered.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error Create Yaml:", err)
	}
}

func getFileList(s string) []fs.DirEntry {
	flist, err := os.ReadDir(s)
	if err != nil {
		fmt.Println(err)
	}
	return flist
}

func MakeManifest(v string) {
	// Load values.yaml
	valuesData, err := os.ReadFile(v)
	if err != nil {
		fmt.Println("Error reading values.yaml:", err)
		return
	}

	var valuesMap map[string]interface{}
	err = yaml.Unmarshal(valuesData, &valuesMap)
	if err != nil {
		fmt.Println("Error parsing values.yaml:", err)
		return
	}

	data := templateData{
		Values: valuesMap,
	}

	path := data.Values["template"].(string)
	basePath := path + "base/"
	overlayPath := path + "overlay/"
	baselist := getFileList(basePath)
	overlaylist := getFileList(overlayPath)
	for _, i := range baselist {
		data.renderBaseManifest(basePath, i.Name())
	}
	for _, i := range overlaylist {
		data.renderOverlayManifest(overlayPath, i.Name())
	}
}
