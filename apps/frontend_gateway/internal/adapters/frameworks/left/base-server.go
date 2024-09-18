package left

type BaseServer interface {
	ListenAndServe() error
}