package util

import (
	"github.com/prometheus/common/model"
)

func init() {
	model.NameValidationScheme = model.UTF8Validation
}
