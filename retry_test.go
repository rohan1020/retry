package retry_test

import (
	"errors"
	"fmt"
	"github.com/rohan1020/retry"
	"testing"
)

func TestRetrying(t *testing.T) {

	fmt.Println("Retry Test:")

	i := 0

	retry.Do(func() error {

		switch i {

		case 3:
			return nil

		default:
			i = i + 1
			return errors.New("Simulated Error")

		}

		return nil

	})

}
