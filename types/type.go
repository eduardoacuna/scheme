package types

import (
	"io"

	"github.com/eduardoacuna/scheme/errors"
)

// Object is the interface implemented by all scheme values
type Object interface{}

// Immediate is the type of immediate unique values
type Immediate int

const (
	nullObject Immediate = iota
	trueObject
	falseObject
	eofObject
	undefinedObject
	unspecifiedObject
)

// Null returns the scheme null value
func Null() Immediate {
	return nullObject
}

// True returns the scheme true value
func True() Immediate {
	return trueObject
}

// False returns the scheme false value
func False() Immediate {
	return falseObject
}

// Boolean constructs either a true or false value
func Boolean(b bool) Immediate {
	if b {
		return trueObject
	}
	return falseObject
}

// EOF returns the scheme eof value
func EOF() Immediate {
	return eofObject
}

// Undefined returns the scheme undefined value
func Undefined() Immediate {
	return undefinedObject
}

// Unspecified returns the scheme unspecified value
func Unspecified() Immediate {
	return unspecifiedObject
}

// Fixnum is the type of immediate integer values
type Fixnum int64

// NewFixnum constructs a fixnum value
func NewFixnum(value int64) Fixnum {
	return Fixnum(value)
}

// Character is the type of immediate character values
type Character rune

// NewCharacter constructs a character value
func NewCharacter(value rune) Character {
	return Character(value)
}

// Flonum is the type of immediate floating point values
type Flonum float64

// NewFlonum constructs a flonum value
func NewFlonum(value float64) Flonum {
	return Flonum(value)
}

// Pair is the type of cons cells
type Pair struct {
	Car Object
	Cdr Object
}

// NewPair constructs a Pair reference
func NewPair(car, cdr Object) (*Pair, error) {
	return &Pair{
		Car: car,
		Cdr: cdr,
	}, nil
}

// Car returns the first component of a pair
func Car(cons *Pair) (Object, error) {
	if cons == nil {
		return nil, errors.NewError(errors.NilError, "given a nil reference", "cons:", cons)
	}
	return cons.Car, nil
}

// Cdr returns the second component of a pair
func Cdr(cons *Pair) (Object, error) {
	if cons == nil {
		return nil, errors.NewError(errors.NilError, "given a nil reference", "cons:", cons)
	}
	return cons.Cdr, nil
}

// Symbol is the type of symbol values
type Symbol struct {
	Name string
}

// symbolTable is the global mapping of strings to symbol values
var symbolTable = map[string]*Symbol{}

// GetSymbol takes a string and returns it's corresponding symbol value
func GetSymbol(name string) *Symbol {
	sym, ok := symbolTable[name]
	if !ok {
		sym = &Symbol{
			Name: name,
		}
		symbolTable[name] = sym
	}
	return sym
}

// SymbolName returns the name of a symbol
func SymbolName(sym *Symbol) (string, error) {
	if sym == nil {
		return "", errors.NewError(errors.NilError, "given a nil reference", "sym:", sym)
	}
	return sym.Name, nil
}

// InputPort is the type of input port values
type InputPort struct {
	Reader io.Reader
}

// NewInputPort constructs an InputPort reference
func NewInputPort(reader io.Reader) *InputPort {
	return &InputPort{
		Reader: reader,
	}
}

// OutputPort is the type of output port values
type OutputPort struct {
	Writer io.Writer
}

// NewOutputPort constructs an OutputPort reference
func NewOutputPort(writer io.Writer) *OutputPort {
	return &OutputPort{
		Writer: writer,
	}
}

// String is the type of string values
type String struct {
	Elements []Character
	Length   int
}

// NewString constructs a String reference
func NewString(length int, value Character) (*String, error) {
	if length < 0 {
		return nil, errors.NewError(errors.ValueError, "given a length < 0", "length:", length)
	}
	elms := make([]Character, length)
	for i := range elms {
		elms[i] = value
	}
	return &String{
		Elements: elms,
		Length:   length,
	}, nil
}

// StringRef returns the character at a string position
func StringRef(str *String, i int) (Character, error) {
	if str == nil {
		return NewCharacter(0), errors.NewError(errors.NilError, "given a nil reference", "str:", str)
	}
	if i < 0 || i >= len(str.Elements) {
		return NewCharacter(0), errors.NewError(errors.OutOfBoundsError, "given a bad string index", "i:", i)
	}
	return str.Elements[i], nil
}

// StringSet assigns a character at a string position
func StringSet(str *String, i int, c Character) error {
	if str == nil {
		return errors.NewError(errors.NilError, "given a nil reference", "str:", str)
	}
	if i < 0 || i >= len(str.Elements) {
		return errors.NewError(errors.OutOfBoundsError, "given a bad string index", "i:", i)
	}
	str.Elements[i] = c
	return nil
}

// Vector is the type of vector values
type Vector struct {
	Elements []Object
	Length   int
}

// NewVector constructs a Vector reference
func NewVector(length int, value Object) (*Vector, error) {
	if length < 0 {
		return nil, errors.NewError(errors.ValueError, "given a length < 0", "length:", length)
	}
	elms := make([]Object, length)
	for i := range elms {
		elms[i] = value
	}
	return &Vector{
		Elements: elms,
		Length:   length,
	}, nil
}

// VectorRef returns the object at a vector position
func VectorRef(vec *Vector, i int) (Object, error) {
	if vec == nil {
		return nil, errors.NewError(errors.NilError, "given a nil reference", "vec:", vec)
	}
	if i < 0 || i >= len(vec.Elements) {
		return nil, errors.NewError(errors.OutOfBoundsError, "given a bad vector index", "i:", i)
	}
	return vec.Elements[i], nil
}

// VectorSet assigns a object at a string position
func VectorSet(vec *Vector, i int, x Object) error {
	if vec == nil {
		return errors.NewError(errors.NilError, "given a nil reference", "vec:", vec)
	}
	if i < 0 || i >= len(vec.Elements) {
		return errors.NewError(errors.OutOfBoundsError, "given a bad vector index", "i:", i)
	}
	vec.Elements[i] = x
	return nil
}

// ByteVector is the type of byte-vector values
type ByteVector struct {
	Elements []byte
	Length   int
}

// NewByteVector constructs a ByteVector reference
func NewByteVector(length int, value Fixnum) (*ByteVector, error) {
	if length < 0 {
		return nil, errors.NewError(errors.ValueError, "given a length < 0", "length:", length)
	}
	if int64(value) < 0 || int64(value) > 255 {
		return nil, errors.NewError(errors.ValueError, "given a number not in [0, 255]", "value:", value)
	}
	elms := make([]byte, length)
	for i := range elms {
		elms[i] = byte(value)
	}
	return &ByteVector{
		Elements: elms,
		Length:   length,
	}, nil
}

// ByteVectorRef returns the byte at a string position
func ByteVectorRef(bv *ByteVector, i int) (Fixnum, error) {
	if bv == nil {
		return NewFixnum(0), errors.NewError(errors.NilError, "given a nil reference", "bv:", bv)
	}
	if i < 0 || i >= len(bv.Elements) {
		return NewFixnum(0), errors.NewError(errors.OutOfBoundsError, "given a bad byte-vector index", "i:", i)
	}
	return NewFixnum(int64(bv.Elements[i])), nil
}

// ByteVectorSet assigns a byte at a string position
func ByteVectorSet(bv *ByteVector, i int, n Fixnum) error {
	if bv == nil {
		return errors.NewError(errors.NilError, "given a nil reference", "bv:", bv)
	}
	if i < 0 || i >= len(bv.Elements) {
		return errors.NewError(errors.OutOfBoundsError, "given a bad byte-vector index", "i:", i)
	}
	if int64(n) < 0 || int64(n) > 255 {
		return errors.NewError(errors.ValueError, "given a number not in [0, 255]", "n:", n)
	}
	bv.Elements[i] = byte(n)
	return nil
}
