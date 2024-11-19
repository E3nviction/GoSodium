package types

type Element struct {
	Tag        string            `json:"tag"`
	Attributes map[string]string `json:"attributes"`
	Children   []Element         `json:"children"`
}
