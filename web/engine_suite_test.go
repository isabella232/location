package web_test

import (
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWeb(t *testing.T) {
	_ = godotenv.Load("../.env.test")
	RegisterFailHandler(Fail)
	RunSpecs(t, "Web Suite")
}
