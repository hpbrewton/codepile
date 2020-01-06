package main

import (
	"log"
	"gonum.org/v1/gonum/mat"
)

// type VPTree struct {
// 	Ndim uint
// 	[] mat.Vector
// }

type Weighted struct {
	angle float64
	value interface{}
}

type KVP struct {
	v mat.Vector
	o interface{}
}

func Create(kos chan KVP, done chan bool) {
	for kvp := range kos {
		log.Println(kvp.v, kvp.o)
	}
	done <- true
}

// finds the min(n, |VPTree|) vectors closest to k in the vtree,
// and returns those vectors paired with their angles
// there is no change to the VPTree
// func (vtree *VPTree) Nearest(k mat.Vector, n uint) []Weighted {
// 	return []Weighted{}
// }

func main() {
	kos := make(chan KVP, 128)
	done := make(chan bool)
	go Create(kos, done)
	kos <- KVP{mat.NewVecDense(3, []float64{1, 2, 3}), 5}
	kos <- KVP{mat.NewVecDense(3, []float64{1, 2, 3}), 5}
	kos <- KVP{mat.NewVecDense(3, []float64{1, 2, 3}), 5}
	kos <- KVP{mat.NewVecDense(3, []float64{1, 2, 3}), 5}
	close(kos)
	<-done
}