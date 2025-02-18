package maketemplate

import "bytes"

type TemplateData interface {
	Load() (error, bytes.Buffer)
}
