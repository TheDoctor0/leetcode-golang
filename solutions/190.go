package solutions

func reverseBits(num uint32) uint32 {
    var result uint32

    for i := 0; i < 32; i++ {
        result += num >> i << 31 >> i
    }

    return result
}
