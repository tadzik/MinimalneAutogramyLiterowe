package main

import (
	"./autogramy"
	"bytes"
	"fmt"
	"math"
	"math/rand"
)

const (
	generations = 1000
	popCount    = 50
	rankCount   = 6
)

type Population struct {
	genomes [popCount]Genom
	best    []Genom
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
	for i := 0; i < rankCount; i++ {
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
	for i := 0; i < generations; i++ {
		//calculate and sum scores
		scoresSum := 0
		for j := range scores {
			scores[j].genom = population.genomes[j]
			toSentence(&population.genomes[j], sentence)
			scores[j].score = (int)(sentence.Score())
			scoresSum += scores[j].score
			// check if we wonna push the best element on the list of best genomes
			if len(population.best) == 0 {
				population.best = append(population.best, population.genomes[j])
			} else {
				newSentence := &autogramy.Sentence{}
				toSentence(&population.best[len(population.best)-1], newSentence)
				newScore := (int)(newSentence.Score())
				if scores[j].score < newScore {
					population.best = append(population.best, scores[j].genom)
				}
			}

		}
		for j := range population.genomes {
			spawnGenome(&population.genomes[j], &scores, scoresSum)
		}
		// replace old population with new population

		// best
		// sort
	}
}
func main() {
	rand.Seed(43)
	sen := &autogramy.Sentence{}
	var population Population
	for i := range population.genomes {
		randomizeGenom(&population.genomes[i])
	}
	runAlgorithm(&population)
	for i := range population.best {
		toSentence(&population.best[i], sen)
		fmt.Println(sen.String())
		fmt.Println((int)(sen.Score()))

	}
}
