package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func doAuth() error {
	// migrations
	dbType := cel.DB.Type
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := cel.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := cel.RootPath + "/migrations/" + fileName + ".down.sql"

	err := copyFileFromTemplate("templates/migrations/auth-tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens cascade;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// copy files
	err = copyFileFromTemplate("templates/data/user.go.txt", cel.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/token.go.txt", cel.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/remember_token.go.txt", cel.RootPath+"/data/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy middleware
	err = copyFileFromTemplate("templates/middleware/auth.go.txt", cel.RootPath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/auth-token.go.txt", cel.RootPath+"/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/remember.go.txt", cel.RootPath+"/middleware/remember.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy handlers
	err = copyFileFromTemplate("templates/handlers/auth-handlers.go.txt", cel.RootPath+"/handlers/auth-handlers.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy views
	err = copyFileFromTemplate("templates/mailer/password-reset.html.tmpl", cel.RootPath+"/mail/mailer/password-reset.html.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/mailer/password-reset.text.tmpl", cel.RootPath+"/mail/mailer/password-reset.text.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/password-reset.text.tmpl", cel.RootPath+"/mail/mailer/password-reset.text.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/login.jet", cel.RootPath+"/views/login.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/forgot.jet", cel.RootPath+"/views/forgot.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/reset-password.jet", cel.RootPath+"/views/reset-password.jet")
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow("  - users, tokens and remember_tokens migrations created and executed")
	color.Yellow("  - user and token models created")
	color.Yellow("  - auth middleware created")
	color.Yellow("")
	color.Yellow("  - Don't forget to add user and token models in data/models.go, and to add appropriate middleware to your routes!")

	return nil
}
