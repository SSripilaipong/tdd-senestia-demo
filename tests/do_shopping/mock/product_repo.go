package mock

import "tdd-senestia-demo/domain/do_shopping"

type ProductRepoMock struct {
	FindById_id     int
	FindById_Return do_shopping.ProductDetailRepr
}

func (p *ProductRepoMock) FindById(id int) (do_shopping.ProductDetailRepr, error) {
	p.FindById_id = id
	return p.FindById_Return, nil
}
