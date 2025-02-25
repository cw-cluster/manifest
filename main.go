package main

import (
	"manifest/cmd"
	"manifest/pkg/makeManifest"
)

func main() {
	makeManifest.MakeManifest("values.yaml", "dev")
	cmd.Execute()
}
