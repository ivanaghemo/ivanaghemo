package structs

type Point struct {
	X float32
	Y float32
}

type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
	Location Point    `json:"location"`
}

type TopSecretRequest struct {
	Satellites []Satellite `json:"satellites"`
}

type TopSecretResponse struct {
	Position Point  `json:"position"`
	Message  string `json:"message"`
}
