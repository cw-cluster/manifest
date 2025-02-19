package makeManifest

type templatePath struct {
	basePath    string
	overlayPath string
}

func NewTempldatePathConstructor(inputPath string) *templatePath {
	return &templatePath{
		basePath:    inputPath + "base/",
		overlayPath: inputPath + "overlay/",
	}
}
