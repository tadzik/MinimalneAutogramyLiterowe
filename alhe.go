package main

import (
    "bytes"
    "fmt"
    "math"
)

func byteIdx(b byte) int {
    arr := []byte{ b }
    b = bytes.ToLower(arr)[0]

    return int(b) - int('a')
}

func idxByte(i int) byte {
    arr := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
                  'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
    return arr[i]
}

func numString(n int) string {
    if n > 99 { panic("FUCKUP") }
    arr := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
                    "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
                    "seventeen", "eighteen", "nineteen", "twenty"}
    if n <= 20 {
        return arr[n]
    }

    arr2 := []string{"", "", "twenty", "thirty", "fourty", "fifty",
                     "sixty", "seventy", "eighty", "ninety"}
    prefix := ""
    for i := 9; i >= 0; i-- {
        if n > i * 10 {
            n -= i * 10
            prefix = arr2[i]
            break
        }
    }
    if n == 0 {
        return prefix
    }
    return fmt.Sprintf("%s-%s", prefix, arr[n])
}

type Sentence [26]int

func (me *Sentence) writePart(buf *bytes.Buffer, n int) {
    buf.WriteString(" ")
    buf.WriteString(numString(me[n]))
    buf.WriteString(" ")
    buf.WriteString(string(idxByte(n)))
    buf.WriteString("'s,")
}

func (me *Sentence) String() string {
    buf := bytes.NewBuffer([]byte("this sentence contains only"))
    for i, v := range me {
        if v == 0 {
            continue
        }
        me.writePart(buf, i)
    }
    buf.WriteString(" and")
    return buf.String()
}

func isImportant(b byte) bool {
    return int(b) >= int('a') && int(b) <= int('z')
}

func (me *Sentence) score() float64 {
    var counts [26]int

    // counts for "this sentence contains only (...) and"
    counts[byteIdx('a')] = 2
    counts[byteIdx('c')] = 2
    counts[byteIdx('d')] = 1
    counts[byteIdx('e')] = 3
    counts[byteIdx('h')] = 1
    counts[byteIdx('i')] = 2
    counts[byteIdx('l')] = 1
    counts[byteIdx('n')] = 6
    counts[byteIdx('o')] = 2
    counts[byteIdx('s')] = 3
    counts[byteIdx('t')] = 3
    counts[byteIdx('y')] = 1

    for k, v := range me {
        if v == 0 {
            continue
        }
        counts[k]++
        counts[byteIdx('s')]++
        repr := numString(v)
        for _, v := range []byte(repr) {
            if isImportant(v) {
                counts[byteIdx(v)]++
            }
        }
    }

    diffSum := 0.0
    //charSum := 0.0
    for i := 0; i < 26; i++ {
        //charSum += float64(me[i])
        if int(math.Abs(float64(counts[i] - me[i]))) > 0 {
            fmt.Println("There were supposed to be", me[i],
                        string(idxByte(i)), "but there are", counts[i])
            diffSum += math.Abs(float64(counts[i] - me[i]))
        }
    }

    return diffSum// + charSum * 0.001
}

func main() {
    sen := &Sentence{}
    sen[byteIdx('a')] = 3
    sen[byteIdx('c')] = 3
    sen[byteIdx('d')] = 2
    sen[byteIdx('e')] = 25
    sen[byteIdx('f')] = 9
    sen[byteIdx('g')] = 4
    sen[byteIdx('h')] = 8
    sen[byteIdx('i')] = 12
    sen[byteIdx('l')] = 3
    sen[byteIdx('n')] = 15
    sen[byteIdx('o')] = 9
    sen[byteIdx('r')] = 8
    sen[byteIdx('s')] = 24
    sen[byteIdx('t')] = 18
    sen[byteIdx('u')] = 5
    sen[byteIdx('v')] = 4
    sen[byteIdx('w')] = 6
    sen[byteIdx('x')] = 2
    sen[byteIdx('y')] = 4
    fmt.Println(sen.String())
    fmt.Println("Score is", sen.score())
}
