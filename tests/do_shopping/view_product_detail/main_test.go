package view_product_detail

import (
	"github.com/google/go-cmp/cmp"
	"tdd-senestia-demo/domain/do_shopping"
	"tdd-senestia-demo/domain/do_shopping/usecase"
	"tdd-senestia-demo/tests/do_shopping/mock"
	"testing"
)

func Test_should_find_product_by_id_from_repository(t *testing.T) {
	repo := &mock.ProductRepoMock{}
	u := usecase.NewDoShoppingUsecase(repo)
	_, _ = u.ViewProductDetail(1234)

	Assert(t, repo.FindById_id == 1234)
}

func Test_should_return_product_detail_response(t *testing.T) {
	repo := &mock.ProductRepoMock{FindById_Return: do_shopping.ProductDetailRepr{
		Id:    1234,
		Name:  "Shane",
		Price: 999.0,
	}}
	u := usecase.NewDoShoppingUsecase(repo)
	result, _ := u.ViewProductDetail(1234)

	expected := do_shopping.ProductDetailResponse{
		Id:    1234,
		Name:  "Shane",
		Price: 999.0,
	}
	Assert(t, cmp.Equal(result, expected))
}

func Assert(t *testing.T, cond bool) {
	if !cond {
		t.Error("assertion failed")
	}
}
