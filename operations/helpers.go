package operations

func controlOverflow(a int64, b int64) uint8 {
	var ret int64

	if b > 0 {
		ret = a + b
		if ret > 255 {
			ret = 255
		}
	} else {
		ret = a + b
		if ret < 0 {
			ret = 0
		}
	}

	return uint8(ret)
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func toUint8(a, b, c, d uint32) (w, x, y, z uint8) {
	return uint8(a), uint8(b), uint8(c), uint8(d)
}
