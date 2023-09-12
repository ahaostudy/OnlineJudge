package problem

import (
	"main/internal/data/model"
	"main/internal/gateway/dao"
)

func GetProblemList(page, count int) ([]*model.Problem, error) {
	return dao.GetProblemListLimit((page-1)*count, count)
}
