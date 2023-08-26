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

/*
 * Return the magnitude
 */
func mag(r, g, b, rSmol, gSmol, bSmol uint8) int {
	tempR := int(r) - int(rSmol)
	tempG := int(g) - int(gSmol)
	tempB := int(b) - int(bSmol)
	tempR *= tempR
	tempG *= tempG
	tempB *= tempB

	return tempR + tempG + tempB
}

/*
 * Return the magnitude
 */
func mag2(r, g, b, rSmol, gSmol, bSmol uint8) int {
	tempR := (float64(r) - float64(rSmol))*0.3	
	tempG := (float64(g) - float64(gSmol))*0.59
	tempB := (float64(b) - float64(bSmol))*0.11
	tempR *= tempR
	tempG *= tempG
	tempB *= tempB

	return int(tempR + tempG + tempB)
}
