package entities

//Restaurant entity
type Restaurant struct {
	ID         uint   `gorm:"primarkey"`
	Name       string `gorm:"unique"`
	Address    string
	Number     int
	CP         string `gorm:"size:5"`
	RFC        string `gorm:"size:13"`
	Logo       []byte
	Categories []Category `gorm:"constraint:,OnDelete:CASCADE;"`
	Clients    []Client   `gorm:"constraint:,OnDelete:CASCADE;"`
	Employees  []Employee `gorm:"constraint:,OnDelete:CASCADE;"`
}
