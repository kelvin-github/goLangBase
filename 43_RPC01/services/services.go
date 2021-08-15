package services

import context "context"

type ProdService struct {
}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
    return &ProdResponse{ProdStock: "20"},nil
}