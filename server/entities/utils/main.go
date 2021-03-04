package eutils

import (
	"github.com/rillo-carrillo/restaurant/server/entities"
	orm "github.com/rillo-carrillo/restaurant/server/orm/entities"
	"github.com/rillo-carrillo/restaurant/server/utils"
	"gorm.io/gorm"
)

//CreateEntities func creates the entities and relationships for the tables.
func CreateEntities(DB *gorm.DB) error {

	DB.AutoMigrate(&orm.User{}, &entities.Restaurant{}, &entities.Category{}, &entities.Role{}, &entities.Employee{}, &entities.Client{}, &entities.Product{}, &entities.Modifier{}, &entities.Order{}, &entities.OrderProduct{}, &entities.OrderDetail{})

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
		CP:      "45050",
		Number:  4493,
		RFC:     "AAAA000000XXX",
		Categories: []entities.Category{
			entities.Category{
				Name: "Bebidas",
				Products: []entities.Product{
					entities.Product{
						Name: "Coca Cola",
						Modifiers: []entities.Modifier{
							entities.Modifier{
								Description: "Lata 355ml",
								Price:       16,
							},
							entities.Modifier{
								Description: "Botella Vidrio 500ml",
								Price:       20,
							},
						},
					},
				},
			},
			entities.Category{
				Name: "Desayunos",
				Products: []entities.Product{
					entities.Product{
						Name: "Huevos Al gusto",

						Modifiers: []entities.Modifier{
							entities.Modifier{
								Description: "Sin salsa",
								Price:       60,
							},
						},
					},
					entities.Product{
						Name: "Chilaquiles",
						Modifiers: []entities.Modifier{
							entities.Modifier{
								Description: "Verdes",
								Price:       60,
							},
							entities.Modifier{
								Description: "Rojos",
								Price:       60,
							},
							entities.Modifier{
								Description: "Combinados",
								Price:       60,
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
						Name: "Hot Cakes",
						Modifiers: []entities.Modifier{
							entities.Modifier{
								Description: "Chocomilk",
								Price:       35,
							},
							entities.Modifier{
								Description: "Jugo",
								Price:       35,
							},
							entities.Modifier{
								Description: "Fruta",
								Price:       35,
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
	pass, _ := utils.GeneratePassword("admin")
	adm := entities.Employee{
		Name:     "Admin",
		RoleID:   1,
		Username: "admin",
		Password: pass,
	}
	pass, _ = utils.GeneratePassword("mesero1")
	emp := entities.Employee{
		Name:     "Mesero 1",
		RoleID:   6,
		Username: "mesero1",
		Password: pass,
	}
	DB.Model(&restaurant).Association("Employees").Append(&adm)
	DB.Model(&restaurant).Association("Employees").Append(&emp)
	order := entities.Order{}
	DB.Model(&emp).Association("Order").Append(&order)
	coca := entities.Product{}
	huevos := entities.Product{}
	DB.First(&huevos, "id=?", 2)
	DB.Model(&order).Association("Products").Append(&coca, &huevos)
	return nil
}
