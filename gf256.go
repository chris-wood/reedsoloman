type Field struct {
	log [256] byte
	exp [512] byte
}

func (f *Field) Add(x byte, y byte) byte {
	return x ^ y
}

func (f *Field) Exp(e int) byte {
	if e < 0 {
		return 0
	}
	return f.exp[e % 255]
}

func (f *Field) Log(e byte) int {
	if e < 0 {
		return -1
	}
	return int(f.log[e])
}

func (f *Field) Mul(x byte, y byte) byte {
	if x == 0 || y == 0 {
		return 0
	}
	return f.exp[int(f.log[x]) + int(f.log[y]) % 255]
}

// x = alpha^y
// x^(-1) = (alpha^y)^(-1) = alpha^(-y mod 256) = alpha^(255 - y)
func (f *Field) Inv(x byte) byte {
	if x == 0 {
		return 0
	}
	return f.exp[255 - f.log[x]]
}

func mul(x int, y int, poly int) int {
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

func CreateField(poly, alpha int) *Field {
	var f Field
	x := 1
	for i := 0; i < 255; i++ {
		f.exp[i] = byte(x)
		f.exp[i + 255] = byte(x)
		f.log[x] = byte(i)
		x = mul(x, alpha, poly)
	}
	f.log[0] = 255
	return &f
}
