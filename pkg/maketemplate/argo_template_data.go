package maketemplate

import "bytes"

type ArgoTemplateData struct {
	Values map[string]interface{}
}

func NewArgoTemplateData() *ArgoTemplateData {
	return &ArgoTemplateData{}
}

func (ad *ArgoTemplateData) Load() (error, bytes.Buffer) {

	return nil, bytes.Buffer{}
}
