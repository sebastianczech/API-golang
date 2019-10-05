package info

type InfoApi struct {
	Name    string `json:"Name"`
	Version string `json:"Version"`
}

func NewInfoApi() *InfoApi {
	return &InfoApi{
		Name:    "",
		Version: "",
	}
}
