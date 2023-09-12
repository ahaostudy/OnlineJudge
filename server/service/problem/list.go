package problem

import (
	"main/model"
	"main/server/dao"
)

func GetProblemList(page, count int) ([]*model.Problem, error) {
	return dao.GetProblemListLimit((page-1)*count, count)
}
