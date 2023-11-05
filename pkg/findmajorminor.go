package Myls

func Major(dev uint64) uint32 {

    major := uint32((dev & 0x00000000000fff00) >> 8)

    major |= uint32((dev & 0xfffff00000000000) >> 32)

    return major

}

func Minor(dev uint64) uint32 {

    minor := uint32((dev & 0x00000000000000ff) >> 0)

    minor |= uint32((dev & 0x00000ffffff00000) >> 12)

    return minor

}