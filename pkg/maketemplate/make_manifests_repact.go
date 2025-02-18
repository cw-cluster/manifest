package maketemplate

import (
	"bytes"
	"fmt"
	"html/template"
	_ "io/fs"
	"os"

	_ "gopkg.in/yaml.v2"
)

type ManifestConstructor struct {
	InputPath    string
	OutputDir    string
	ManifestName string
	templateData *TemplateData
}

func NewManifestConstructor(inputFile, outputDir, manifestName string, templateData *TemplateData) *ManifestConstructor {
	return &ManifestConstructor{
		InputPath:    inputFile,
		OutputDir:    outputDir,
		templateData: templateData,
		ManifestName: manifestName,
	}
}

func (mc *ManifestConstructor) RenderManifest() error {
	tmplData, err := os.ReadFile(mc.InputPath + mc.ManifestName)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return err
	}

	tmpl, err := template.New("manifest").Parse(string(tmplData))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	var rendered bytes.Buffer
	err = tmpl.Execute(&rendered, mc.templateData)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return err
	}

	_, err = os.Stat(mc.OutputDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("base", 0744)
		} else {
			fmt.Println(err)
		}
	}

	err = os.WriteFile("base/"+mc.OutputDir, rendered.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error Create Yaml:", err)
	}

	return nil
}
