package models

type AboutInfo struct {
	Name      string   `json:"name"`
	Bio       string   `json:"bio"`
	Skills    []string `json:"skills"`
	Interests []string `json:"interests"`
}
