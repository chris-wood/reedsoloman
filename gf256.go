type Field struct {
	
}

func (f *Field) Add(x byte, y byte) byte {
	return x ^ y
}

func (f *Field) Mul(x byte, y byte) byte {
    z := 0
    for x > 0 {
        if (x & 1) != 0 {
            z += y
        }
        x >>= 1
        y <<= 1
    }
    return z
}


