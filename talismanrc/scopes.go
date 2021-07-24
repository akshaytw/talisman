package talismanrc

var knownScopes = map[string][]string{
	"node":   {"package-lock.json", "node_modules/"},
	"go":     {"makefile", "go.mod", "go.sum", "Gopkg.toml", "Gopkg.lock", "glide.yaml", "glide.lock", "vendor/"},
	"images": {"*.jpeg", "*.jpg", "*.png", "*.tiff", "*.bmp"},
	"unity":  {"ProjectSettings/ProjectSettings.asset"},
}
