package generator

type JSONSchema struct {
	Definitions map[string]Definition `json:"$defs"`
}

type Definition struct {
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
}

type Property struct {
	Type   string `json:"type"`
	Format string `json:"format,omitempty"`
	Ref    string `json:"$ref,omitempty"`
}
