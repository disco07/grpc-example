package calculator

import (
	"context"
	"fmt"
	"github.com/disco07/gRPC/generated/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
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

func (c *Calculator) Modulo(in *calculator.ModuloRequest, stream grpc.ServerStreamingServer[calculator.ModuloResponse]) error {
	var k int64 = 2
	var i = in.GetA()

	for i > 1 {
		if i%k == 0 {
			err := stream.SendMsg(&calculator.ModuloResponse{Result: k})
			if err != nil {
				return err
			}

			i = i / k
		} else {
			k++
		}
	}

	return nil
}

func (c *Calculator) Average(stream grpc.ClientStreamingServer[calculator.AverageRequest, calculator.AverageResponse]) error {
	var res float32
	var sum, count int64
	for {
		req, err := stream.Recv()
		fmt.Println("Message received", req.GetA())

		if err != nil {
			if err == io.EOF {
				res = float32(sum) / float32(count)

				return stream.SendAndClose(&calculator.AverageResponse{Result: res})
			}

			return err
		}

		sum += req.GetA()
		count++
	}
}

func (c *Calculator) Max(biStream grpc.BidiStreamingServer[calculator.MaxRequest, calculator.MaxResponse]) error {
	var prev, curr int64

	for {
		req, err := biStream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		curr = req.GetA()

		if prev < curr {
			prev = curr

			err = biStream.SendMsg(&calculator.MaxResponse{Result: prev})
			if err != nil {
				return err
			}
		}
	}
}
