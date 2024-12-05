package main

import (
	"dev/pkg"
	"errors"
	"log"
)

func repository() error {
	err := errors.New("genericerror")
	return pkg.BuildErrorLog(pkg.RepositoryLayer, "repositoryfunc", "repositorymessage", pkg.ErrNoData, pkg.ErrInternal, err)
}

func service() error {
	// err := errors.New("genericerror")
	err := repository()
	return pkg.BuildErrorLog(pkg.ServiceLayer, "servicefunc", "servicemessage", pkg.ErrCheck, pkg.ErrValidation, err)
}

func controller() {
	// err := errors.New("genericerror")
	err := service()
	log.Println(pkg.BuildErrorLog(pkg.ControllerLayer, "controllerfunc", "controllermessage", err))
}

func handler() {
	controller()
}

func main() {
	handler()
}
