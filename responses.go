package helm

import "time"

type AutoGenerated struct {
	Release Release `json:"release,omitempty"`

	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Release struct {
	Name string `json:"name"`
	Info struct {
		Status struct {
			Code string `json:"code"`
		} `json:"status"`
		FirstDeployed time.Time `json:"first_deployed"`
		LastDeployed  time.Time `json:"last_deployed"`
		Description   string    `json:"Description"`
	} `json:"info"`
	Chart struct {
		Metadata struct {
			Name        string `json:"name"`
			Version     string `json:"version"`
			Description string `json:"description"`
			APIVersion  string `json:"apiVersion"`
		} `json:"metadata"`
		Templates []struct {
			Name string `json:"name"`
			Data string `json:"data"`
		} `json:"templates"`
		Values struct {
			Raw string `json:"raw"`
		} `json:"values"`
		Files []struct {
			TypeURL string `json:"type_url"`
			Value   string `json:"value"`
		} `json:"files"`
	} `json:"chart"`
	Config struct {
	} `json:"config"`
	Manifest  string `json:"manifest"`
	Version   int    `json:"version"`
	Namespace string `json:"namespace"`
}
