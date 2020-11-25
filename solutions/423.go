package solutions

import (
    "bytes"
)

func remove(str string, n int, count *[]int) {
    for _, s := range str {
        (*count)[s - 'a'] = (*count)[s - 'a'] - n
    }
}

func originalDigits(s string) string {
    counter := [26]int{}

    for _, value := range s {
        counter[value - 'a']++
    }

    numberCounter := [10]int{}
    numberCounter[0] = counter['z' - 'a'] // zero - z
    numberCounter[2] = counter['w' - 'a'] // two - w
    numberCounter[4] = counter['u' - 'a'] // four - u
    numberCounter[6] = counter['x' - 'a'] // six - x
    numberCounter[8] = counter['g' - 'a'] // eight - g

    numberCounter[3] = counter['h' - 'a'] - numberCounter[8] // eight & three - h
    numberCounter[5] = counter['f' - 'a'] - numberCounter[4] // four & five - f
    numberCounter[7] = counter['v' - 'a'] - numberCounter[5] // five & seven - value
    numberCounter[1] = counter['o' - 'a'] - numberCounter[0] - numberCounter[2] - numberCounter[4] // zero & one & two & four - o
    numberCounter[9] = counter['i' - 'a'] - numberCounter[5] - numberCounter[6] - numberCounter[8] // five & six & eight & nine - i

    result := bytes.Buffer{}

    for i := 0; i <= 9; i++ {
        for j := 0; j < numberCounter[i]; j++ {
            _ = result.WriteByte('0' + byte(i))
        }
    }
    return result.String()
}
