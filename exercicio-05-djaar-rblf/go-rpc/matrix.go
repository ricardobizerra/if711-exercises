package goRpc

import (
	"exercicio-05-djaar-rblf/matrix"
	"exercicio-05-djaar-rblf/shared"
)

type MatrixService struct{}

func (s *MatrixService) Multiply(req shared.Request, res *shared.Reply) error {
	response := matrix.Multiply(req.A, req.B)

	res.R = response

	return nil
}
