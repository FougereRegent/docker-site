package helper

func MemSet(data []byte, value byte) {
	for i := 0; i < len(data); i++ {
		data[i] = value
	}
}
