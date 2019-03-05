package types

type CalcService struct{}

type Args struct {
	Operand1, Operand2 int
}

type Reply struct {
	Result int
}

func (cs *CalcService) Sum(args Args, reply *Reply) error {
	reply.Result = args.Operand1 + args.Operand2
	return nil
}
