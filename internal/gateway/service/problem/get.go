package problem

import (
	"main/internal/data/model"
	"main/internal/gateway/dao"
)

func GetProblem(id int64) (*model.Problem, error) {
	return dao.GetProblem(id)
}
