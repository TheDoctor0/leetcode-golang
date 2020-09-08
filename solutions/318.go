package solutions

type MaskPair struct {
    mask int
    length int
}

func maxProduct(words []string) int {
    masks := make([]MaskPair, len(words))

    for i := 0; i < len(words); i++ {
        masks[i] = MaskPair{mask: getBitMask(words[i]), length: len(words[i])}
    }

    result := 0

    for i := 0; i < len(masks); i++ {
        for j := i + 1; j < len(masks); j++ {
            if masks[i].mask & masks[j].mask != 0 {
                continue
            }

            current := masks[i].length * masks[j].length

            if current > result {
                result = current
            }
        }
    }

    return result
}

func getBitMask(s string) int {
    result := 0

    for i := 0; i < len(s); i++ {
        result |= 1 << int(s[i] - 'a')
    }

    return result
}
