package main

import (
    "./autogramy"
    "bytes"
    "fmt"
    "math"
    "math/rand"
)
const (
    generations=10
    popCount=10
)
type Population struct {
    genomes [popCount] Genom
    best []Genom
}

type Genom [26] float64

type GenomScore struct {
    genom *Genom
    score int
}

func toSentence(genom *Genom, sentence *autogramy.Sentence) {
    for i, v := range genom {
        sentence[i]=(int)(math.Floor(v*100))
        if sentence[i]==100 {
            sentence[i]--
        }
    }       
}

func randomizeGenom(genom *Genom)  {
    for i := range genom {
        genom[i]=rand.Float64()
    }      
}

func byteIdx(b byte) int {
    arr := []byte{ b }
    b = bytes.ToLower(arr)[0]

    return int(b) - int('a')
}
func runAlgorithm(population * Population) {
    scores:= [popCount] GenomScore{}
    sentence := &autogramy.Sentence{}
    for i:=0;i<generations;i++ {
        for j := range scores {
            scores[j].genom=&population.genomes[j]
            toSentence(&population.genomes[j],sentence)
            scores[j].score=(int)(sentence.Score())
            
            // check if we wonna push the best element on the list of best genomes
            if len(population.best) == 0 {
                population.best = append(population.best, population.genomes[j])
            } else {
                newSentence := &autogramy.Sentence{}
                toSentence(&population.genomes[j],newSentence)
                newScore:=(int)(sentence.Score())
                if (scores[j].score<newScore) {
                    population.best=append(population.best, *scores[j].genom)
                }
                    
            }
       
        }           
        //best
        // sort
    }
}
func main() {
    rand.Seed(43)
    sen := &autogramy.Sentence{}
    var population Population;
    for i := range population.genomes {
        randomizeGenom(&population.genomes[i])
    }
    runAlgorithm(&population);
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
    //fmt.Println(sen.String())
    for i := range population.best {
        toSentence(&population.best[i],sen)
        fmt.Println(sen.String())
        fmt.Println((int)(sen.Score()))
        
    }
    //fmt.Println("Score is", sen.Score())
}
