package flag

import "github.com/bcowtech/config/internal/flag"

func Process(target interface{}) error {
	return flag.Process(target)
}
