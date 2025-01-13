package main

import (
	"fmt"

	"github.com/google/uuid"
)

// Generar ID único con la librería "github.com/google/uuid"
func main() {
	fmt.Println(uuid.New().String())
}
