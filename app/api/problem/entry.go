package problem

type Group struct{}

func (g *Group) Problem() *ProblemApi {
	return &insProblem
}
