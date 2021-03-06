// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type ComplexUnionTestRecord struct {

	
	
		UnionField *UnionNullArrayIntMapIntNestedUnionRecord
	

}

func NewComplexUnionTestRecord() (*ComplexUnionTestRecord) {
	return &ComplexUnionTestRecord{}
}

func DeserializeComplexUnionTestRecord(r io.Reader) (*ComplexUnionTestRecord, error) {
	t := NewComplexUnionTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func DeserializeComplexUnionTestRecordFromSchema(r io.Reader, schema string) (*ComplexUnionTestRecord, error) {
	t := NewComplexUnionTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func writeComplexUnionTestRecord(r *ComplexUnionTestRecord, w io.Writer) error {
	var err error
	
	err = writeUnionNullArrayIntMapIntNestedUnionRecord( r.UnionField, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *ComplexUnionTestRecord) Serialize(w io.Writer) error {
	return writeComplexUnionTestRecord(r, w)
}

func (r *ComplexUnionTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"UnionField\",\"type\":[\"null\",{\"items\":\"int\",\"type\":\"array\"},{\"type\":\"map\",\"values\":\"int\"},{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"}],\"name\":\"NestedUnionRecord\",\"type\":\"record\"}]}],\"name\":\"ComplexUnionTestRecord\",\"type\":\"record\"}"
}

func (r *ComplexUnionTestRecord) SchemaName() string {
	return "ComplexUnionTestRecord"
}

func (_ *ComplexUnionTestRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *ComplexUnionTestRecord) SetInt(v int32) { panic("Unsupported operation") }
func (_ *ComplexUnionTestRecord) SetLong(v int64) { panic("Unsupported operation") }
func (_ *ComplexUnionTestRecord) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *ComplexUnionTestRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *ComplexUnionTestRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *ComplexUnionTestRecord) SetString(v string) { panic("Unsupported operation") }
func (_ *ComplexUnionTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ComplexUnionTestRecord) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.UnionField = NewUnionNullArrayIntMapIntNestedUnionRecord()
	
		
		
			return r.UnionField
		
	
	}
	panic("Unknown field index")
}

func (r *ComplexUnionTestRecord) SetDefault(i int) {
	switch (i) {
	
        
	
	}
	panic("Unknown field index")
}

func (_ *ComplexUnionTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ComplexUnionTestRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *ComplexUnionTestRecord) Finalize() { }
