package usecase 

import (
	"math/rand"
	"time"
)


const (
    rs6Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
    rs6LetterIdxBits = 6
    rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
    rs6LetterIdxMax = 63 / rs6LetterIdxBits
)

func RandString(n int) string {
    randSrc := rand.NewSource(time.Now().UnixNano())
    b := make([]byte, n)
    cache, remain := randSrc.Int63(), rs6LetterIdxMax
    for i := n-1; i >= 0; {
        if remain == 0 {
            cache, remain = randSrc.Int63(), rs6LetterIdxMax
        }
        idx := int(cache & rs6LetterIdxMask)
        if idx < len(rs6Letters) {
            b[i] = rs6Letters[idx]
            i--
        }
        cache >>= rs6LetterIdxBits
        remain--
    }
    return string(b)
}