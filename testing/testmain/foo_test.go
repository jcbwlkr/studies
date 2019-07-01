package testmain

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Start TestMain")
	res := m.Run()

	log.Println("End TestMain")
	os.Exit(res)
}

func TestHelloWorld(t *testing.T) {
	t.Log("Hi")
}

func TestGoodbyeWorld(t *testing.T) {
	t.Log("Bye")
}
