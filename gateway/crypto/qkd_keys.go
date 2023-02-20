package crypto

type Keys struct {
	Key_id string `json:"key_id"`
	Key    string `json:"key"`
}

type RequestObj struct {
	Keys []Keys `json:"Keys"`
}
