# MusGo_uint8
Provides serialization of the Golang's `uint8` type.

# How to use
Simply cast `uint8` to `musgo_uint8.Uint8`:
```go
	var n uint8 = 5
	// Marshal
	buf := make([]byte, musgo_uint8.Uint8(n).SizeMUS())
	musgo_uint8.Uint8(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_uint8.Uint8)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

