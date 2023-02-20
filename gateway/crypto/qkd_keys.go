package crypto

type Keys struct {
	Key_id  string `json:"key_id"`
	Key_tmp string `json:"key"`
	Key     []byte
}

type RequestObj struct {
	Keys []Keys `json:"Keys"`
}
