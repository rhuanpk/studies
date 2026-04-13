package main

import (
	"dev/pkg/errer/sentinels"
	"dev/pkg/logging"
	"errors"
	"log"
)

func repository() error {
	err := errors.New("externalerror")
	// return logging.Repository("repositoryfunc", []string{"repositorymessage"}, []error{sentinel.ErrNoData, sentinel.ErrInternal}, err)
	return logging.Repository("repositoryfunc", nil, []error{sentinels.ErrNoData, sentinels.ErrInternal}, err)
}

func service() error {
	err := repository()
	// return logging.Service("servicefunc", []string{"servicemessage"}, []error{sentinel.ErrCheck, sentinel.ErrValidation}, err)
	return logging.Service("", []string{"servicemessage"}, nil, errors.New("internal error"), err)
}

func controller() {
	err := service()
	// log.Println(logging.Controller("controllerfunc", []string{"controllermessage"}, nil, err))
	log.Println(logging.Controller("controllerfunc", nil, nil, err))
}

func handler() {
	controller()
}

func main() {
	handler()
}
