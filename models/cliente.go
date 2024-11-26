package models

type Cliente struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Nome     string `json:"nome"`
	CPF      string `json:"cpf"`
	Zap      string `json:"zap"`
	Tel      string `json:"tel"`
	Endereco string `json:"endereco"`
	Numero   string `json:"numero"`
	Bairro   string `json:"bairro"`
	Cidade   string `json:"cidade"`
	UF       string `json:"uf"`
}
