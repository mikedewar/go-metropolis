package metropolisHastings

import (
	"math"
	"math/rand"
)

func (mh MetropolisHastings) Run(initialState, float64, out chan float64) {
	y := initialState
	for {
		x := mh.Sample(y)
		if !mh.Reject(x, y) {
			out <- x
			y = x
		}
	}
}

type MetropolisHastings struct {
	Q func(float64) float64 // proposal distribution
	P func(float64) float64 // function to estimate probability at a point x
}

func (mh MetropolisHastings) Propose(x float64) float64 {

}

func (mh MetropolisHastings) Reject(x1, x2) bool {
	α := mh.P(x1) / mh.P(x2)
	if α >= 1 || rand.Float64() < α {
		return false
	}
	return true
}

func NewMetropolisHastings(Q ProposalDistrbution, P Distrbution, x0 float64) MetropolisHastings {
	return MetropolisHastings(Q, P, x0)
}

func GaussianSample(μ, σ float64) float64 {
	return NormFloat64()*math.Pow(σ, 2) + μ
}

func GaussianEval(μ, σ float64) float64 {
	return 1 / (σ * math.Sqrt(2*math.Pi)) * math.Exp(-0.5*math.Pow((x-μ)/σ, **2))
}
