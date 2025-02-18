package maketemplate

import "bytes"

type K8STemplateData struct {
	Values map[string]interface{}
}

func NewK8sTemplateData() *ArgoTemplateData {
	return &ArgoTemplateData{}
}

func (ad *K8STemplateData) Load() (error, bytes.Buffer) {

	return nil, bytes.Buffer{}
}
