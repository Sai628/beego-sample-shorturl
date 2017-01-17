package models

import (
    "crypto/md5"
    "fmt"
    "io"
)

const (
    DOMAIN = "http://sai628.com/"
)

var (
    globalnum int
)

func init() {
    globalnum = 100000000
}

func GetMD5(lurl string) string {
    h := md5.New()
    salt1 := "salt4shorturl"
    io.WriteString(h, lurl + salt1)
    urlmd5 := fmt.Sprintf("%x", h.Sum(nil))
    return urlmd5
}

func Generate() (tiny string) {
    globalnum++
    num := globalnum
    fmt.Println(num)

    alpha := merge(getRange(48, 57), getRange(65, 90))  // [0-9] + [A-Z]
    alpha = merge(alpha, getRange(97, 122))  // alpha + [a-z]
    if num < 62 {
        tiny = string(alpha[num])
    } else {
        var runes []rune
        runes = append(runes, alpha[num%62])  // len(alpha) = 62
        num = num / 62
        for num >= 1 {
            if num < 62 {
                runes = append(runes, alpha[num - 1])
            } else {
                runes = append(runes, alpha[num%62])
            }
            num = num / 62
        }

        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        tiny = string(runes)
    }
    return DOMAIN + tiny
}

func getRange(start, end rune) (ran []rune) {
    for i := start; i < end; i++ {
        ran = append(ran, i)
    }
    return ran
}

func merge(a, b []rune) []rune {
    c := make([]rune, len(a) + len(b))
    copy(c, a)
    copy(c[len(a):], b)
    return c
}
