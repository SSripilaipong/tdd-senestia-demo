package do_shopping

type IProductRepo interface {
	FindById(id int) (ProductDetailRepr, error)
}
