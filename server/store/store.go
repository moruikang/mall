package store

var client Factory

// Factory defines the mall platform storage interface.
type Factory interface {
	Categorys() CategoryStore
	Products() ProductStore
	Orders() OrderStore
	Close() error
}
