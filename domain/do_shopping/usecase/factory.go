package usecase

import (
	"tdd-senestia-demo/domain/do_shopping"
	"tdd-senestia-demo/domain/do_shopping/internal"
)

func NewDoShoppingUsecase(repo do_shopping.IProductRepo) internal.DoShoppingUsecase {
	return internal.DoShoppingUsecase{Repo: repo}
}
