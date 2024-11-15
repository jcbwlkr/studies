package main

import (
	"errors"
	"time"

	"bitbucket.org/ardanlabs/log"
)

func main() {
	log.Init("my-test")

	id := time.Now().String()
	log.Startf(id, "i say main")

	foo(id)

	log.Complete(id)
}

func foo(id interface{}) {
	log.Startf(id, "i say foo")
	defer log.Complete(id)

	log.Info(id, "foo work 1")

	if err := bar(id); err != nil {
		log.Info(id, "bar failed... oh well")
	}

	log.Info(id, "foo work 2")
}

func bar(id interface{}) error {
	log.Startf(id, "i say bar")

	log.Info(id, "Do some bar work")

	if id != nil {
		err := errors.New("Something happened")
		log.CompleteErr(err, id)
		return err
	}

	log.Complete(id)
	return nil
}
