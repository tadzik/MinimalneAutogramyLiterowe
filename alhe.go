package main

import (
	"./autogramy"
	"bytes"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

var profiling = flag.Bool("profiling", false, "perform profiling with runtime/pprof")
var numThreads = flag.Int("threads", 1, "how many cpu threads to use")
var generations = flag.Int("generations", 2000, "how many generations to calculate")
var rankCount = flag.Int("rankCount", 6, "Size of extra ranks for the ranking selection") 
//var popCount    = flag.Int("populations", 50, "how many populations to maintain")

const popCount = 50

type Population struct {
	genomes [popCount]Genom
	best    []GenomScore
}

type Genom [26]float64

type GenomScore struct {
	genom Genom
	score int
}

func toSentence(genom *Genom, sentence *autogramy.Sentence) {
	for i, v := range genom {
		sentence[i] = (int)(math.Floor(v * 100))
		if sentence[i] == 100 {
			sentence[i]--
		}
	}
}

func randomizeGenom(genom *Genom) {
	for i := range genom {
		genom[i] = rand.Float64()
	}
}

func byteIdx(b byte) int {
	arr := []byte{b}
	b = bytes.ToLower(arr)[0]

	return int(b) - int('a')
}

func findParents(scores *[popCount]GenomScore, scoresSum int) (*Genom, *Genom) {
	father := &scores[rand.Intn(len(scores))]
	mother := &scores[rand.Intn(len(scores))]
	if father.score < mother.score {
		father, mother = mother, father
	}
	for i := 0; i < *rankCount; i++ {
		candidate := &scores[rand.Intn(len(scores))]
		if candidate.score < mother.score {
			if candidate.score < father.score {
				father, mother = candidate, father
			} else {
				mother = candidate
			}
		}
	}
	return &father.genom, &mother.genom
	panic("Failed to choose parent")
}

func spawnGenome(genom *Genom, scores *[popCount]GenomScore, scoresSum int) {
	// random two parents.
	father, mother := findParents(scores, scoresSum)
	crossing1 := rand.Intn(len(genom))
	crossing2 := rand.Intn(len(genom))
	if crossing1 > crossing2 {
		crossing1, crossing2 = crossing2, crossing1
	}
	for i := range genom {
		// get genome from parents
		if i < crossing1 && i >= crossing2 {
			genom[i] = mother[i]
		} else {
			genom[i] = father[i]
		}
		// mutate
		genom[i] += rand.NormFloat64() / 20

		genom[i] = math.Min(genom[i], 1)
		genom[i] = math.Max(genom[i], 0)
	}

}

func runAlgorithm(population *Population) {
	scores := [popCount]GenomScore{}
	sentence := &autogramy.Sentence{}
	for i := 0; i < *generations; i++ {
		//calculate and sum scores
		scoresSum := 0
		var wg sync.WaitGroup
		for j := range scores {
			wg.Add(1)
			go func() {
				defer wg.Done()
				scores[j].genom = population.genomes[j]
				toSentence(&population.genomes[j], sentence)
				scores[j].score = (int)(sentence.Score())
				scoresSum += scores[j].score
				// check if we wonna push the best element on the list of best genomes
				if len(population.best) == 0 {
					population.best = append(population.best, scores[j])
                    
				} else {
					if scores[j].score < population.best[len(population.best)-1].score {
						population.best = append(population.best, scores[j])
					}
				}
			}()
		}
		wg.Wait()
		for j := range population.genomes {
			spawnGenome(&population.genomes[j], &scores, scoresSum)
		}
	}
}
func main() {
	flag.Parse()
	if *profiling {
		f, err := os.Create("profile")
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	runtime.GOMAXPROCS(*numThreads)
	fmt.Printf("Running on %d CPU threads\n", *numThreads)

	rand.Seed(43)
	sen := &autogramy.Sentence{}
	var population Population
	for i := range population.genomes {
		randomizeGenom(&population.genomes[i])
	}
	runAlgorithm(&population)
	for i := range population.best {
		toSentence(&population.best[i].genom, sen)
		fmt.Println(sen.String())
		fmt.Println((int)(sen.Score()))
	}
}
