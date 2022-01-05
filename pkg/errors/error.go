package errors

import "fmt"

var (
	MISSING_NAME = fmt.Errorf(
		"name of label should not be nil pr empty",
	)
)
