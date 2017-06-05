package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImmediate(t *testing.T) {
	assert.True(t, reflect.TypeOf(Null()).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, Null(), Null(), "they should be equal")
	assert.NotEqual(t, Null(), True(), "they shouldn't be equal")
	assert.NotEqual(t, Null(), False(), "they shouldn't be equal")
	assert.NotEqual(t, Null(), EOF(), "they shouldn't be equal")
	assert.NotEqual(t, Null(), Undefined(), "they shouldn't be equal")
	assert.NotEqual(t, Null(), Unspecified(), "they shouldn't be equal")

	assert.True(t, reflect.TypeOf(True()).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, True(), True(), "they should be equal")
	assert.NotEqual(t, True(), False(), "they shouldn't be equal")
	assert.NotEqual(t, True(), EOF(), "they shouldn't be equal")
	assert.NotEqual(t, True(), Undefined(), "they shouldn't be equal")
	assert.NotEqual(t, True(), Unspecified(), "they shouldn't be equal")

	assert.True(t, reflect.TypeOf(False()).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, False(), False(), "they should be equal")
	assert.NotEqual(t, False(), EOF(), "they shouldn't be equal")
	assert.NotEqual(t, False(), Undefined(), "they shouldn't be equal")
	assert.NotEqual(t, False(), Unspecified(), "they shouldn't be equal")

	assert.Equal(t, True(), Boolean(true), "they should be equal")
	assert.Equal(t, False(), Boolean(false), "they should be equal")

	assert.True(t, reflect.TypeOf(EOF()).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, EOF(), EOF(), "they should be equal")
	assert.NotEqual(t, EOF(), Undefined(), "they shouldn't be equal")
	assert.NotEqual(t, EOF(), Unspecified(), "they shouldn't be equal")

	assert.True(t, reflect.TypeOf(Undefined()).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, Undefined(), Undefined(), "they should be equal")
	assert.NotEqual(t, Undefined(), Unspecified(), "they shouldn't be equal")

	assert.True(t, reflect.TypeOf(Unspecified()).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, Unspecified(), Unspecified(), "they should be equal")
}

func TestFixnum(t *testing.T) {
	assert.True(t, reflect.TypeOf(NewFixnum(120)).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, NewFixnum(666), NewFixnum(666), "they should be equal")
	assert.NotEqual(t, NewFixnum(123), 123, "they shouldn't be equal")
}

func TestCharacter(t *testing.T) {
	assert.True(t, reflect.TypeOf(NewCharacter('x')).Size() <= 8, "byte width should be at most a word")
	assert.True(t, reflect.TypeOf(NewCharacter('巨')).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, NewCharacter('流'), NewCharacter('流'), "they should be equal")
	assert.NotEqual(t, NewCharacter('F'), 'F', "they shouldn't be equal")
}

func TestFlonum(t *testing.T) {
	assert.True(t, reflect.TypeOf(NewFlonum(0.1)).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, NewFlonum(666.333), NewFlonum(666.333), "they should be equal")
	assert.NotEqual(t, NewFlonum(-321.0), -321.0, "they shouldn't be equal")
}

func TestPair(t *testing.T) {
	cons1, err1 := NewPair(NewFixnum(1), NewFixnum(2))
	cons2, err2 := NewPair(NewFixnum(1), NewFixnum(2))
	cons3, err3 := NewPair(cons1, cons2)

	assert.NoError(t, err1, "it shouln't be an error")
	assert.NoError(t, err2, "it shouln't be an error")
	assert.NoError(t, err3, "it shouln't be an error")

	cons3car, err := Car(cons3)
	assert.NoError(t, err, "it shouldn't be an error")
	cons3cdr, err := Cdr(cons3)
	assert.NoError(t, err, "it shouldn't be an error")

	assert.True(t, reflect.TypeOf(cons1).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, cons1, cons2, "they should be equal")
	assert.False(t, cons1 == cons2, "they shouldn't be the same")
	assert.Equal(t, cons1, cons3car, "they should be equal")
	assert.Equal(t, cons2, cons3cdr, "they should be equal")

	_, err = Car(nil)
	assert.Error(t, err, "it should be an error")
	_, err = Cdr(nil)
	assert.Error(t, err, "it should be an error")
}

func TestSymbol(t *testing.T) {
	sym := GetSymbol("foo")

	assert.True(t, reflect.TypeOf(sym).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, sym, GetSymbol("foo"), "they should be equal")
	assert.True(t, sym == GetSymbol("foo"), "they should be the same")
	assert.NotEqual(t, sym, "foo", "they shouldn't be equal")
	assert.NotEqual(t, sym, GetSymbol("FOO"), "they shouldn't be equal")

	name, err := SymbolName(sym)
	assert.Equal(t, "foo", name, "they should be equal")

	_, err = SymbolName(nil)
	assert.Error(t, err, "it should be an error")
}

func TestString(t *testing.T) {
	str, err := NewString(5, NewCharacter('x'))

	assert.NoError(t, err, "it shouldn't be an error")

	_, err = NewString(-4, NewCharacter('x'))
	assert.Error(t, err, "it should be an error")

	str2, err := NewString(5, NewCharacter('x'))
	assert.NoError(t, err, "it shouldn't be an error")

	assert.True(t, reflect.TypeOf(str).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, str, str2, "they should be equal")
	assert.False(t, str == str2, "they shouldn't be the same")
	assert.NotEqual(t, str, "xxxxx", "they shouldn't be equal")

	c0, err := StringRef(str, 0)
	assert.NoError(t, err, "it shouldn't be an error")
	c1, err := StringRef(str, 1)
	assert.NoError(t, err, "it shouldn't be an error")
	c2, err := StringRef(str, 2)
	assert.NoError(t, err, "it shouldn't be an error")
	c3, err := StringRef(str, 3)
	assert.NoError(t, err, "it shouldn't be an error")
	c4, err := StringRef(str, 4)
	assert.NoError(t, err, "it shouldn't be an error")

	assert.Equal(t, NewCharacter('x'), c0, "they should be equal")
	assert.Equal(t, NewCharacter('x'), c1, "they should be equal")
	assert.Equal(t, NewCharacter('x'), c2, "they should be equal")
	assert.Equal(t, NewCharacter('x'), c3, "they should be equal")
	assert.Equal(t, NewCharacter('x'), c4, "they should be equal")

	err = StringSet(str, 0, NewCharacter('h'))
	assert.NoError(t, err, "it shouldn't be an error")
	err = StringSet(str, 1, NewCharacter('e'))
	assert.NoError(t, err, "it shouldn't be an error")
	err = StringSet(str, 2, NewCharacter('l'))
	assert.NoError(t, err, "it shouldn't be an error")
	err = StringSet(str, 3, NewCharacter('l'))
	assert.NoError(t, err, "it shouldn't be an error")
	err = StringSet(str, 4, NewCharacter('o'))
	assert.NoError(t, err, "it shouldn't be an error")

	c0, err = StringRef(str, 0)
	assert.NoError(t, err, "it shouldn't be an error")
	c1, err = StringRef(str, 1)
	assert.NoError(t, err, "it shouldn't be an error")
	c2, err = StringRef(str, 2)
	assert.NoError(t, err, "it shouldn't be an error")
	c3, err = StringRef(str, 3)
	assert.NoError(t, err, "it shouldn't be an error")
	c4, err = StringRef(str, 4)
	assert.NoError(t, err, "it shouldn't be an error")

	assert.NotEqual(t, str, "hello", "they shouldn't be equal")
	assert.Equal(t, NewCharacter('h'), c0, "they should be equal")
	assert.Equal(t, NewCharacter('e'), c1, "they should be equal")
	assert.Equal(t, NewCharacter('l'), c2, "they should be equal")
	assert.Equal(t, NewCharacter('l'), c3, "they should be equal")
	assert.Equal(t, NewCharacter('o'), c4, "they should be equal")

	_, err = StringRef(nil, 0)
	assert.Error(t, err, "it should be an error")
	_, err = StringRef(str, -1)
	assert.Error(t, err, "it should be an error")
	_, err = StringRef(str, 5)
	assert.Error(t, err, "it should be an error")

	err = StringSet(nil, 0, 'x')
	assert.Error(t, err, "it should be an error")
	err = StringSet(str, -1, 'x')
	assert.Error(t, err, "it should be an error")
	err = StringSet(str, 5, 'x')
	assert.Error(t, err, "it should be an error")
}

func TestVector(t *testing.T) {
	vec, err := NewVector(5, Null())

	assert.NoError(t, err, "it shouldn't be an error")

	_, err = NewVector(-4, Null())
	assert.Error(t, err, "it should be an error")

	vec2, err := NewVector(5, Null())
	assert.NoError(t, err, "it shouldn't be an error")

	assert.True(t, reflect.TypeOf(vec).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, vec, vec2, "they should be equal")
	assert.False(t, vec == vec2, "they shouldn't be the same")

	x0, err := VectorRef(vec, 0)
	assert.NoError(t, err, "it shouldn't be an error")
	x1, err := VectorRef(vec, 1)
	assert.NoError(t, err, "it shouldn't be an error")
	x2, err := VectorRef(vec, 2)
	assert.NoError(t, err, "it shouldn't be an error")
	x3, err := VectorRef(vec, 3)
	assert.NoError(t, err, "it shouldn't be an error")
	x4, err := VectorRef(vec, 4)
	assert.NoError(t, err, "it shouldn't be an error")

	assert.Equal(t, Null(), x0, "they should be equal")
	assert.Equal(t, Null(), x1, "they should be equal")
	assert.Equal(t, Null(), x2, "they should be equal")
	assert.Equal(t, Null(), x3, "they should be equal")
	assert.Equal(t, Null(), x4, "they should be equal")

	err = VectorSet(vec, 0, True())
	assert.NoError(t, err, "it shouldn't be an error")
	err = VectorSet(vec, 1, False())
	assert.NoError(t, err, "it shouldn't be an error")
	err = VectorSet(vec, 2, NewFixnum(69))
	assert.NoError(t, err, "it shouldn't be an error")
	err = VectorSet(vec, 3, NewFlonum(44.2))
	assert.NoError(t, err, "it shouldn't be an error")
	err = VectorSet(vec, 4, NewCharacter('r'))
	assert.NoError(t, err, "it shouldn't be an error")

	x0, err = VectorRef(vec, 0)
	assert.NoError(t, err, "it shouldn't be an error")
	x1, err = VectorRef(vec, 1)
	assert.NoError(t, err, "it shouldn't be an error")
	x2, err = VectorRef(vec, 2)
	assert.NoError(t, err, "it shouldn't be an error")
	x3, err = VectorRef(vec, 3)
	assert.NoError(t, err, "it shouldn't be an error")
	x4, err = VectorRef(vec, 4)
	assert.NoError(t, err, "it shouldn't be an error")

	assert.Equal(t, True(), x0, "they should be equal")
	assert.Equal(t, False(), x1, "they should be equal")
	assert.Equal(t, NewFixnum(69), x2, "they should be equal")
	assert.Equal(t, NewFlonum(44.2), x3, "they should be equal")
	assert.Equal(t, NewCharacter('r'), x4, "they should be equal")

	_, err = VectorRef(nil, 0)
	assert.Error(t, err, "it should be an error")
	_, err = VectorRef(vec, -1)
	assert.Error(t, err, "it should be an error")
	_, err = VectorRef(vec, 5)
	assert.Error(t, err, "it should be an error")

	err = VectorSet(nil, 0, Null())
	assert.Error(t, err, "it should be an error")
	err = VectorSet(vec, -1, Null())
	assert.Error(t, err, "it should be an error")
	err = VectorSet(vec, 5, Null())
	assert.Error(t, err, "it should be an error")
}

func TestByteVector(t *testing.T) {
	bv, err := NewByteVector(5, NewFixnum(255))

	assert.NoError(t, err, "it shouldn't be an error")

	_, err = NewByteVector(-4, NewFixnum(255))
	assert.Error(t, err, "it should be an error")

	bv2, err := NewByteVector(5, NewFixnum(255))
	assert.NoError(t, err, "it shouldn't be an error")

	assert.True(t, reflect.TypeOf(bv).Size() <= 8, "byte width should be at most a word")
	assert.Equal(t, bv, bv2, "they should be equal")
	assert.False(t, bv == bv2, "they shouldn't be the same")
	assert.NotEqual(t, bv, "xxxxx", "they shouldn't be equal")

	b0, err := ByteVectorRef(bv, 0)
	assert.NoError(t, err, "it shouldn't be an error")
	b1, err := ByteVectorRef(bv, 1)
	assert.NoError(t, err, "it shouldn't be an error")
	b2, err := ByteVectorRef(bv, 2)
	assert.NoError(t, err, "it shouldn't be an error")
	b3, err := ByteVectorRef(bv, 3)
	assert.NoError(t, err, "it shouldn't be an error")
	b4, err := ByteVectorRef(bv, 4)
	assert.NoError(t, err, "it shouldn't be an error")

	assert.Equal(t, NewFixnum(255), b0, "they should be equal")
	assert.Equal(t, NewFixnum(255), b1, "they should be equal")
	assert.Equal(t, NewFixnum(255), b2, "they should be equal")
	assert.Equal(t, NewFixnum(255), b3, "they should be equal")
	assert.Equal(t, NewFixnum(255), b4, "they should be equal")

	err = ByteVectorSet(bv, 0, NewFixnum(1))
	assert.NoError(t, err, "it shouldn't be an error")
	err = ByteVectorSet(bv, 1, NewFixnum(2))
	assert.NoError(t, err, "it shouldn't be an error")
	err = ByteVectorSet(bv, 2, NewFixnum(4))
	assert.NoError(t, err, "it shouldn't be an error")
	err = ByteVectorSet(bv, 3, NewFixnum(4))
	assert.NoError(t, err, "it shouldn't be an error")
	err = ByteVectorSet(bv, 4, NewFixnum(8))
	assert.NoError(t, err, "it shouldn't be an error")

	b0, err = ByteVectorRef(bv, 0)
	assert.NoError(t, err, "it shouldn't be an error")
	b1, err = ByteVectorRef(bv, 1)
	assert.NoError(t, err, "it shouldn't be an error")
	b2, err = ByteVectorRef(bv, 2)
	assert.NoError(t, err, "it shouldn't be an error")
	b3, err = ByteVectorRef(bv, 3)
	assert.NoError(t, err, "it shouldn't be an error")
	b4, err = ByteVectorRef(bv, 4)
	assert.NoError(t, err, "it shouldn't be an error")

	assert.NotEqual(t, bv, "hello", "they shouldn't be equal")
	assert.Equal(t, NewFixnum(1), b0, "they should be equal")
	assert.Equal(t, NewFixnum(2), b1, "they should be equal")
	assert.Equal(t, NewFixnum(4), b2, "they should be equal")
	assert.Equal(t, NewFixnum(4), b3, "they should be equal")
	assert.Equal(t, NewFixnum(8), b4, "they should be equal")

	_, err = ByteVectorRef(nil, 0)
	assert.Error(t, err, "it should be an error")
	_, err = ByteVectorRef(bv, -1)
	assert.Error(t, err, "it should be an error")
	_, err = ByteVectorRef(bv, 5)
	assert.Error(t, err, "it should be an error")

	err = ByteVectorSet(nil, 0, 'x')
	assert.Error(t, err, "it should be an error")
	err = ByteVectorSet(bv, -1, 'x')
	assert.Error(t, err, "it should be an error")
	err = ByteVectorSet(bv, 5, 'x')
	assert.Error(t, err, "it should be an error")
}
