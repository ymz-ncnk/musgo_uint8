package musgo_uint8

import "testing"

func TestMusgoUint8(t *testing.T) {
	var n uint8 = 5
	buf := make([]byte, Uint8(n).SizeMUS())
	Uint8(n).MarshalMUS(buf)

	var an uint8
	(*Uint8)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
