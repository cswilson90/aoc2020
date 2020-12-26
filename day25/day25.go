package day25

const (
	RemainderNum  = 20201227
	SubjectNumber = 7
)

func GetEncryptionKey(publicKey1, publicKey2 int) int {
	var otherKey, loopSize int

	value := 1
	loops := 0
	for {
		if value == publicKey1 || value == publicKey2 {
			if value == publicKey1 {
				otherKey = publicKey2
			} else if value == publicKey2 {
				otherKey = publicKey1
			}
			loopSize = loops
			break
		}

		value = transform(value, SubjectNumber)
		loops++
	}

	encryptionKey := 1
	for i := 0; i < loopSize; i++ {
		encryptionKey = transform(encryptionKey, otherKey)
	}
	return encryptionKey
}

func transform(value, subjectNum int) int {
	return (value * subjectNum) % RemainderNum
}
