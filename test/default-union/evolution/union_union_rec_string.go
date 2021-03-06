// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"io"
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionUnionRecStringTypeEnum int
const (

	 UnionUnionRecStringTypeEnumUnionRec UnionUnionRecStringTypeEnum = 0

	 UnionUnionRecStringTypeEnumString UnionUnionRecStringTypeEnum = 1

)

type UnionUnionRecString struct {

	UnionRec *UnionRec

	String string

	UnionType UnionUnionRecStringTypeEnum
}

func writeUnionUnionRecString(r *UnionUnionRecString, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionUnionRecStringTypeEnumUnionRec:
		return writeUnionRec(r.UnionRec, w)
        
	case UnionUnionRecStringTypeEnumString:
		return vm.WriteString(r.String, w)
        
	}
	return fmt.Errorf("invalid value for *UnionUnionRecString")
}

func NewUnionUnionRecString() *UnionUnionRecString {
	return &UnionUnionRecString{}
}

func (_ *UnionUnionRecString) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetString(v string) { panic("Unsupported operation") }
func (r *UnionUnionRecString) SetLong(v int64) { 
	r.UnionType = (UnionUnionRecStringTypeEnum)(v)
}
func (r *UnionUnionRecString) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		r.UnionRec = NewUnionRec()
		
		
		return r.UnionRec
		
	
	case 1:
		
		
		return (*types.String)(&r.String)
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionUnionRecString) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionUnionRecString) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionUnionRecString) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionUnionRecString) Finalize()  { }
