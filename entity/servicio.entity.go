package entity

func (Service) TableName() string {
	return "servicio"
}

type Service struct {
	ID      uint   `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Nombre  string `gorm:"type:varchar(100);column:nombre" json:"nombre"`
	Prefijo string `gorm:"type:varchar;not null;column:prefijo" json:"prefijo"`
	Puerto  uint64 `gorm:"type:uint8;column:puerto" json:"puerto"`
}
