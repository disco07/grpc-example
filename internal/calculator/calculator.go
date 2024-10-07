package calculator

import (
	"context"
	"github.com/disco07/gRPC/generated/calulator"
	"google.golang.org/grpc/metadata"
	"log"
)

type Calculator struct {
	calculator.CalculatorServiceServer
}

func NewCalculator() calculator.CalculatorServiceServer {
	return &Calculator{}
}

func (c *Calculator) Add(ctx context.Context, req *calculator.AddRequest) (*calculator.AddResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Aucune métadonnée trouvée dans le contexte")
	} else {
		log.Printf("Métadonnées reçues : %v", md)
	}

	result := req.GetA() + req.GetB()

	return &calculator.AddResponse{Result: result}, nil
}
