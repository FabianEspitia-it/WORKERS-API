package models


import(
	"gorm.io/gorm"
)

type Worker struct {
    gorm.Model
    WorkerName      string `gorm:"type:varchar(100);not null" json:"name"`
    Email     string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
    Password  string `gorm:"type:varchar(100);not null;unique_index" json:"password"`
    RolID     uint   `gorm:"not null" json:"rol_id"`
    CountryID uint   `gorm:"not null" json:"country_id"`
	LeaderID  uint   `gorm:"not null" json:"leader_id"`
}
