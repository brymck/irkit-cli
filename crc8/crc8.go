package crc8

const crc8Polynomial = 0x31

func Crc8FromByte(b byte, crc byte) byte {
	crc ^= b
	for i := 0; i < 8; i++ {
		if crc&0x80 != 0 {
			crc = (crc << 1) ^ crc8Polynomial
		} else {
			crc <<= 1
		}
	}
	return crc
}

func Crc8FromBoolean(b bool, crc byte) byte {
	if b {
		return Crc8FromByte(1, crc)
	}
	return Crc8FromByte(0, crc)
}

func Crc8FromBytes(data []byte, size int, crc byte) byte {
	for i := 0; i < size; i++ {
		if i < len(data) {
			crc = Crc8FromByte(data[i], crc)
		} else {
			crc = Crc8FromByte(0, crc)
		}
	}
	return crc
}

func Crc8FromString(data string, len int, crc byte) byte {
	return Crc8FromBytes([]byte(data), len, crc)
}
