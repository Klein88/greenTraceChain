package Model

type UserKey struct {
	Id            int    `gorm:"column:Id" json:"Id"`
	Username      string `gorm:"column:Username"`
	PrivateKeyPEM string `gorm:"column:PrivateKey"`
	PublicKeyPEM  string `gorm:"column:PublicKey"`
}
