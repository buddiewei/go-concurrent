package go_concurrent

type FuncWithErr[R any] func(...any) (R, error)

type FuncWithReturn[R any] func(...any) R

type ConFuncAggregatorFunc func(cf *conFunc)

type ConFuncAggregator struct {
}
