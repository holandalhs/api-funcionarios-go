package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/holandalhs/api-funcionarios-go/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/funcionarios", getFuncionarios)
	e.POST("/funcionarios", createFuncionario)
	e.GET("/funcionarios/:id", getFuncionario)
	e.PUT("/funcionarios/:id", updateFuncionario)
	e.DELETE("/funcionarios/:id", deleteFuncionario)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

// Handler
func getFuncionarios(c echo.Context) error {
	return c.String(http.StatusOK, "Listando todos os funcionários.")
}

func createFuncionario(c echo.Context) error {
	funcionario := db.Funcionario{}
	if err := c.Bind(&funcionario); err != nil {
		return err
	}
	if err := db.AddFuncionario(funcionario); err != nil {
		return c.String(http.StatusInternalServerError, "Erro interno do servidor")
	}
	return c.String(http.StatusOK, "Criando funcionário.")
}

func getFuncionario(c echo.Context) error {
	id := c.Param("id")
	getFunc := fmt.Sprintf("Pegando o id do funcionário %s", id)
	/* return c.String(http.StatusOK, "Listando funcionário.") */
	return c.String(http.StatusOK, getFunc)
}

func updateFuncionario(c echo.Context) error {
	id := c.Param("id")
	upFunc := fmt.Sprintf("Atualizando o id do funcionário %s", id)
	/* return c.String(http.StatusOK, "Atualizando funcionário.") */
	return c.String(http.StatusOK, upFunc)
}

func deleteFuncionario(c echo.Context) error {
	id := c.Param("id")
	delFunc := fmt.Sprintf("Excluindo o id do funcionário %s", id)
	/* return c.String(http.StatusOK, "Deletando funcionário.") */
	return c.String(http.StatusOK, delFunc)
}
