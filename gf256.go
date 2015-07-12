type Field struct {
	
}

func (f *Field) Add(x byte, y byte) byte {
	return x ^ y
}

func (f *Field) Mul(x int, y int, poly int) int {
    z := 0
    for x > 0 {
        if (x & 1) != 0 {
            z ^= y
        }
        x >>= 1
        y <<= 1
        if (y & 0x100) != 0 {
            y ^= poly
        }
    }
    return z
}


