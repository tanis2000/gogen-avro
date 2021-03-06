// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     avro_java_string.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

func NewEventWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewEvent()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type EventReader struct {
	r io.Reader
	p *vm.Program
}

func NewEventReader(r io.Reader) (*EventReader, error){
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &EventReader {
		r: containerReader,
		p: deser,
	}, nil
}

func (r EventReader) Read() (*Event, error) {
	t := NewEvent()
        err := vm.Eval(r.r, r.p, t)
	return t, err
}
