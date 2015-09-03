package sessionstores

import (
	"testing"

	"github.com/samora/ussd-go/Godeps/_workspace/src/github.com/stretchr/testify/suite"
)

func TestRedis(t *testing.T) {
	store := NewRedis("localhost:6379")
	suite.Run(t, NewStoreSuite(store))
}
