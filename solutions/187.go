package solutions

const (
    None = 0
    Seen = 1
    Duplicated = 2
)

func findRepeatedDnaSequences(s string) []string {
    var repeated []string
    var current uint32

    seen := make([]byte, 1<<20)
    sequence := map[byte]uint32 {
        'A': 0,
        'C': 1,
        'G': 2,
        'T': 3,
    }

    for i := 0; i < len(s); i++ {
        current = ((current << 2) & 0xFFFFF) | sequence[s[i]]

        if i < 9 {
            continue
        }

        if seen[current] == None {
            seen[current] = Seen
        } else if seen[current] == Seen {
            repeated = append(repeated, s[i - 9: i + 1])

            seen[current] = Duplicated
        }
    }

    return repeated
}
