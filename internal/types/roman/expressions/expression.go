package expressions

type Expression interface {
	Solve(c Context) (Context, error)
}
