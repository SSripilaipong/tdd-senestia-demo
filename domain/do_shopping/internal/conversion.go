package internal

import "tdd-senestia-demo/domain/do_shopping"

func ResponseFromRepr(repr do_shopping.ProductDetailRepr) do_shopping.ProductDetailResponse {
	return do_shopping.ProductDetailResponse{
		Id:    repr.Id,
		Name:  repr.Name,
		Price: repr.Price,
	}
}
