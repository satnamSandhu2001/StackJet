package database

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/satnamSandhu2001/stackjet/pkg"
	"github.com/satnamSandhu2001/stackjet/pkg/colors"
)

func RunInitSQL() error {
	conn := Connect()
	defer conn.Close()

	_, err := conn.Exec(string(InitSQL))
	if err != nil {
		return fmt.Errorf("init.sql execution failed: %w", err)
	}

	// Insert default admin
	email := "admin@stackjet.com"
	password := "admin123"

	hashed, err := pkg.GenerateHash(password)
	if err != nil {
		return err
	}
	_, err = conn.Exec(`
		INSERT INTO users (email, password, role)
		VALUES (?, ?, 'superadmin')
	`, email, hashed)

	if err != nil {
		return fmt.Errorf("admin insert failed: %w", err)
	}

	fmt.Println(colors.SecondaryBold("🔐 Admin user created with the following credentials:"))
	fmt.Printf(colors.Secondary("      email:    %s\n"), email)
	fmt.Printf(colors.Secondary("      password: %s\n"), password)

	return nil
}
