package internal

import "tdd-senestia-demo/domain/do_shopping"

type DoShoppingUsecase struct {
	Repo do_shopping.IProductRepo
}

func (u *DoShoppingUsecase) ViewProductDetail(id int) (do_shopping.ProductDetailResponse, error) {
	repr, _ := u.Repo.FindById(id)
	response := ResponseFromRepr(repr)
	return response, nil
}
