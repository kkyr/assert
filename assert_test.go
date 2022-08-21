package assert_test

import (
	"github.com/kkyr/assert"
	"testing"
)

func TestAsserter_Equal(t *testing.T) {
	assert.Field(t, "engine").Require().Equal(4, 4)
	assert.Field(t, "yo").Nil(nil)
	assert.Field(t, "speed").Equal(5, 5)
}
