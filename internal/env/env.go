package env

import (
	"os"
	"strings"

	"github.com/bcowtech/structproto"
	"github.com/bcowtech/structproto/valuebinder"
)

const (
	TagName = "env"
)

func Process(prefix string, target interface{}) error {
	if len(prefix) > 0 {
		prefix += "_"
	}

	prototype, err := structproto.Prototypify(target, &structproto.StructProtoOption{
		TagName:             TagName,
		ValueBinderProvider: valuebinder.BuildStringArgsBinder,
	})
	if err != nil {
		return err
	}

	var table structproto.NamedValues = make(structproto.NamedValues)
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		name, value := parts[0], parts[1]
		if strings.HasPrefix(name, prefix) {
			table[name[len(prefix):]] = value
		}
	}
	err = prototype.BindValues(table)
	if err != nil {
		return err
	}
	return nil
}
