package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type Funcionario struct {
	gorm.Model        /* para incluir os campos/estruturas abaixo na struct */
	Nome       string `json:"nome"`
	CPF        string `json:"cpf"`
	Email      string `json:"email"`
	Idade      int    `json:"idade"`
	Ativo      bool   `json:"registro"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("funcionario.db"), &gorm.Config{})
	/* substituir gorm.db pelo nome da struct que irei usar: funcionario */
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Funcionario{}) //para gerenciamento do banco

	return db
}

func AddFuncionario(funcionario Funcionario) {
	db := Init()

	/* 	funcionario := Funcionario{
		Nome:  "Rock",
		CPF:   "68554454",
		Email: "rock@gmail.com",
		Idade: 3,
		Ativo: true,
	} */

	/* result := db.Create(&funcionario)
	if result.Error != nil {
		fmt.Println("Erro ao cadastrar funcionário")
	} */

	if result := db.Create(&funcionario); result.Error != nil {
		fmt.Println("Erro ao cadastrar funcionário")
	}

	fmt.Println("Funcionário cadastrado!")
}
