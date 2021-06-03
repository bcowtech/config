package resource

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"

	proto "github.com/bcowtech/structproto"
	"github.com/bcowtech/structproto/valuebinder"
)

var (
	typeOfByteArray = reflect.TypeOf([]byte{})
)

type ResourceBinder struct {
	BaseDir string
}

func (p *ResourceBinder) Init(context *proto.StructProtoContext) error {
	return nil
}

func (p *ResourceBinder) Bind(field proto.FieldInfo, rv reflect.Value) error {
	filename := path.Join(p.BaseDir, field.Name())

	fileinfo, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if fileinfo.Mode().IsRegular() {
		buffer, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		switch rv.Type() {
		case typeOfByteArray:
			rv.Set(reflect.ValueOf(buffer))
			return nil
		}
		return valuebinder.BytesArgsBinder(rv).Bind(buffer)
	}
	return nil
}

func (p *ResourceBinder) Deinit(context *proto.StructProtoContext) error {
	return nil
}
