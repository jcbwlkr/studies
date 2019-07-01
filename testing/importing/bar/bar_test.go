package bar

import (
	"fmt"
	"testing"

	"github.com/jcbwlkr/studies/testing/importing/foo_test"
)

func TestBar(t *testing.T) {
	fmt.Println(foo_test.Check())
}
