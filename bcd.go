package bcd

func IntToBcd(value int) int {
	return (((value / 10) % 10) << 4) | (value % 10)
}

func BcdToInt(value int) int {
	return (int)((value>>4)*10 + (value & 0x0F))
}
