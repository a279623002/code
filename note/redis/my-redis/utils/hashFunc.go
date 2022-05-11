package utils

//通用的哈希函数库有下面这些混合了加法和一位操作的字符串哈希算法

//Robert Sedgwicks
func RSHash(str string) uint64 {
	a := 63689
	b := 378551

	hash := uint64(0)

	for i := 0; i < len(str); i++ {
		hash = hash * uint64(a) + uint64(str[i])
		a *= b
	}
	return (hash % 0xFFFFFFFF) & 0xFFFFFFFF
}

//Justin Sobel
func JSHash(str string) uint64 {
	hash := uint64(1315423911)

	for i := 0; i < len(str); i++ {
		hash ^= ((hash << 5) + uint64(str[i]) + hash >> 2)
	}
	return (hash % 0xFFFFFFFF) & 0xFFFFFFFF
}

//PJW 彼得 J 温伯格
func PJWHash(str string) uint64 {
	BitsInUnsignedInt := (uint64)(4 * 8)
	ThreeQuarters := (uint64)((BitsInUnsignedInt * 3) / 4)
	OneEighth := (uint64)(BitsInUnsignedInt / 8)
	HighBits := (uint64)(0xFFFFFFFF) << (BitsInUnsignedInt - OneEighth)
	hash := uint64(0)
	test := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash << OneEighth) + uint64(str[i])
		if test = hash & HighBits; test != 0 {
			hash = ((hash ^ (test >> ThreeQuarters)) & (^HighBits))
		}
	}
	return (hash % 0xFFFFFFFF) & 0xFFFFFFFF
}

// Brian Kernighan 和 Dennis Ritchie
func BKDRHash(str string) uint64 {
	seed := uint64(131) // 31 131 1313 13131 131313 etc..
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint64(str[i])
	}
	return (hash % 0xFFFFFFFF) & 0xFFFFFFFF
}

// 开源的 SDBM
func SDBMHash(str string) uint64 {
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = uint64(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return (hash % 0xFFFFFFFF) & 0xFFFFFFFF
}

// Daniel J.Bernstein
func DJBHash(str string) uint64 {
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + uint64(str[i])
	}
	return (hash % 0xFFFFFFFF) & 0xFFFFFFFF
}

// Knuth
func DEKHash(str string) uint64 {
	hash := uint64(len(str))
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) ^ (hash >> 27)) ^ uint64(str[i])
	}
	return (hash % 0xFFFFFFFF) & 0xFFFFFFFF
}

// Arash Partow
func APHash(str string) uint64 {
	hash := uint64(0xAAAAAAAA)
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= ((hash << 7) ^ uint64(str[i])*(hash>>3))
		} else {
			hash ^= (^((hash << 11) + uint64(str[i]) ^ (hash >> 5)))
		}
	}
	return (hash % 0xFFFFFFFF) & 0xFFFFFFFF
}

func GetHash(fn, str string) uint64 {
	hash := uint64(0)
	switch fn {
	case "APHash":
		hash = APHash(str)
	case "DEKHash":
		hash = DEKHash(str)
	case "DJBHash":
		hash = DJBHash(str)
	case "SDBMHash":
		hash = SDBMHash(str)
	case "BKDRHash":
		hash = BKDRHash(str)
	case "PJWHash":
		hash = PJWHash(str)
	case "JSHash":
		hash = JSHash(str)
	case "RSHash":
		hash = RSHash(str)
	default:
		hash = APHash(str)
	}
	return hash
}