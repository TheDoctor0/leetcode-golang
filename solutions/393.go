package solutions

func validUtf8(data []int) bool {
    var startBit uint8 = 7
    currentBits := 0

    for _, value := range data {
        if currentBits == 0 {
            mask := 1<<startBit

            for value & mask > 0 {
                startBit--
                mask = 1<<startBit
                currentBits++

                if currentBits > 4 {
                    return false
                }
            }

            if currentBits > 1 {
                currentBits--
            }

            startBit = 7
        } else if currentBits > 0 && value & (1<<7) > 0 && value & (1<<6) == 0 {
            currentBits--
        } else {
            return false
        }
    }

    if currentBits > 0 {
        return false
    }

    return true
}
