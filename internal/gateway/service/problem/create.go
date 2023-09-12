package problem

import (
	"main/internal/data/model"
	"main/internal/gateway/dao"
)

func CreateProlem(problem *model.Problem) error {
	// 忽略ID字段
	problem.ID = 0

	return dao.InsertProblem(problem)
}
