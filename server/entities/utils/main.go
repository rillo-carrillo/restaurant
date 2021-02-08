package eutils

import (
	"testorm/entities"

	"gorm.io/gorm"
)

//CreateEntities func creates the entities and relationships for the tables.
func CreateEntities(DB *gorm.DB) error {
	if err := DB.SetupJoinTable(&entities.Order{}, "Products", &entities.OrderProduct{}); err != nil {
		return err
	}
	DB.AutoMigrate(&entities.Restaurant{}, &entities.Category{}, &entities.Role{}, &entities.Employee{}, &entities.Client{}, &entities.Product{}, &entities.Modifier{}, &entities.Order{})
	return nil
}

//InitializeDatabase func create set of entries on the DB for testing.
func InitializeDatabase(DB *gorm.DB) error {
	manager := entities.Role{
		Code: 1,
		Name: "Manager",
	}
	chef := entities.Role{
		Code: 3,
		Name: "Chef",
	}
	servicio := entities.Role{
		Code: 6,
		Name: "Mesero",
	}
	restaurant := entities.Restaurant{
		Name:    "El Meson de Priscila",
		Address: "Calle Tezozomoc 4493 int 1, Zapopan, Jalisco",
		CP:      45050,
		Number:  4493,
		RFC:     "AAAA000000XXX",
		Categories: []entities.Category{
			entities.Category{
				Name: "Bebidas",
				Products: []entities.Product{
					entities.Product{
						Name:  "Coca Cola",
						Price: 16,
						Modifiers: []entities.Modifier{
							entities.Modifier{
								Description: "Lata 355ml",
							},
							entities.Modifier{
								Description: "Botella Vidrio 500ml",
								Price:       4,
							},
						},
					},
				},
			},
			entities.Category{
				Name: "Desayunos",
				Products: []entities.Product{
					entities.Product{
						Name:     "Huevos Al gusto",
						Price:    60,
						Comments: "Termino medio.",
						Modifiers: []entities.Modifier{
							entities.Modifier{
								Description: "Sin salsa",
							},
						},
					},
					entities.Product{
						Name:     "Chilaquiles",
						Price:    60,
						Comments: "Termino medio.",
						Modifiers: []entities.Modifier{
							entities.Modifier{
								Description: "Verdes",
							},
							entities.Modifier{
								Description: "Rojos",
							},
							entities.Modifier{
								Description: "Pollo",
								Price:       10,
							},
							entities.Modifier{
								Description: "Queso",
								Price:       10,
							},
						},
					},
				},
			},
			entities.Category{
				Name: "Infantil",
				Products: []entities.Product{
					entities.Product{
						Name:  "Hot Cakes",
						Price: 35,
						Modifiers: []entities.Modifier{
							entities.Modifier{
								Description: "Chocomilk",
							},
							entities.Modifier{
								Description: "Jugo",
							},
							entities.Modifier{
								Description: "Fruta",
							},
						},
					},
				},
			},
		},
		Clients: []entities.Client{
			entities.Client{
				Name:           "Test",
				FirstLastname:  "Test",
				SecondLastname: "Test",
			},
		},
	}
	if err := DB.Create(&manager); err.Error != nil {
		return err.Error
	}
	if err := DB.Create(&chef); err.Error != nil {
		return err.Error
	}
	if err := DB.Create(&servicio); err.Error != nil {
		return err.Error
	}
	if err := DB.Create(&restaurant); err.Error != nil {
		return err.Error
	}
	emp := entities.Employee{
		Name:   "Mesero 1",
		RoleID: 6,
	}
	DB.Model(&restaurant).Association("Employees").Append(&emp)
	order := entities.Order{}
	DB.Model(&emp).Association("Order").Append(&order)
	coca := entities.Product{}
	DB.First(&coca, "id=?", 1)
	coca2 := entities.Product{}
	DB.First(&coca2, "id=?", 1)
	mod1 := entities.Modifier{}
	DB.First(&mod1, "id=?",1)
	mod2 := entities.Modifier{}
	DB.First(&mod2, "id=?",2)
	mod2:=
	DB.Model(&coca).Association("Modifiers").
	huevos := entities.Product{}
	DB.First(&huevos, "id=?", 2)
	DB.Model(&order).Association("Products").Append(&coca, &huevos)
	return nil
}
