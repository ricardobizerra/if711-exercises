package goRpc

import (
	"exercicio-06-djaar-rblf/matrix"
	"exercicio-06-djaar-rblf/shared"
)

type MatrixService struct{}

func (s *MatrixService) Multiply(req shared.Request, res *shared.Reply) error {
	response := matrix.Multiply(req.A, req.B)

	res.R = response

	return nil
}
