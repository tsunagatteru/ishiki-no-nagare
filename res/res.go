package res

import "embed"

//go:embed templates
//go:embed static
var embedFS embed.FS

func GetEmbedFS() embed.FS {
	return embedFS
}
