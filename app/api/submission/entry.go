package submission

type Group struct{}

func (g *Group) Submission() *SubmitApi {
	return &insSubmission
}
