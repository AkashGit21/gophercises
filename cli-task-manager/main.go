package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AkashGit21/gophercises/cli-task-manager/cmd"
	"github.com/AkashGit21/gophercises/cli-task-manager/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))

	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
