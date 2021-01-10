package mchv3

type EncryptData struct {
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	OriginalType   string `json:"original_type,omitempty"`
	Nonce          string `json:"nonce"`
}
