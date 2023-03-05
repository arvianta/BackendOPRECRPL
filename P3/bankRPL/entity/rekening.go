package entity

type Rekening struct {
	ID        uint64   `gorm:"primaryKey" json:"id"`
	Number    string   `json:"number" binding:"required"`
	Balance   uint64   `json:"balance" binding:"required"`
	NasabahID uint64   `gorm:"foreignKey:ID;references:NasabahID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"nasabah_id" binding:"required"`
	Nasabah   *Nasabah `json:"rekening,omitempty"`
}
