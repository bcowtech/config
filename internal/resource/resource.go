package resource

import (
	"os"

	"github.com/bcowtech/structproto"
)

const (
	TagName = "resource"
)

func Process(baseDir string, target interface{}) error {
	baseDir = os.ExpandEnv(baseDir)
	if len(baseDir) > 0 {
		// exist path
		if _, err := os.Stat(baseDir); os.IsNotExist(err) {
			return nil
		}
	}

	prototype, err := structproto.Prototypify(target, &structproto.StructProtoOption{
		TagName:     TagName,
		TagResolver: ResourceTagResolver,
	})
	if err != nil {
		return err
	}

	return prototype.Bind(&ResourceBinder{
		BaseDir: baseDir,
	})
}
