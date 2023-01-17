package submission

type Group struct{}

func (g *Group) Submission() *SSubmission {
	return &insSubmission
}
