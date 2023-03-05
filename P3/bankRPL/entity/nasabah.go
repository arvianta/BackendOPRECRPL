package entity

type Nasabah struct {
	ID             uint64     `gorm:"primaryKey" json:"id"`
	Nama           string     `json:"nama" binding:"required"`
	Tempat_lahir   string     `json:"tempat_lahir" binding:"required"`
	Tanggal_lahir  string     `json:"tanggal_lahir" binding:"required"`
	Tempat_tinggal string     `json:"tempat_tinggal" binding:"required"`
	Pekerjaan      string     `json:"pekerjaan" binding:"required"`
	Rekening       []Rekening `json:"rekening,omitempty" gorm:"foreignKey:NasabahID"`
}
