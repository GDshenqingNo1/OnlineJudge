package problem

type Group struct{}

func (g *Group) Problem() *SProblem {
	return &insProblem
}
