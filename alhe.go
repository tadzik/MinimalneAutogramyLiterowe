package main

import (
    "./autogramy"
    "bytes"
    "fmt"
    "math"
    "math/rand"
)
type Genom [26] float64

func toSentence(genom *Genom, sentence *autogramy.Sentence) {
    for i, v := range genom {
        sentence[i]=(int)(math.Floor(v*100))
        if sentence[i]==100 {
            sentence[i]--
        }
    }       
}

func getRandomGenom() *Genom {
    genom := &Genom{}
    for i := range genom {
        genom[i]=rand.Float64()
    }      
    return genom
}

func byteIdx(b byte) int {
    arr := []byte{ b }
    b = bytes.ToLower(arr)[0]

    return int(b) - int('a')
}

func main() {
    rand.Seed(13)
    sen := &autogramy.Sentence{}
    genom := getRandomGenom()     
    /*sen[byteIdx('a')] = 3
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
    sen[byteIdx('y')] = 4*/
    toSentence(genom,sen)
    fmt.Println(sen.String())
    fmt.Println("Score is", sen.Score())
}
