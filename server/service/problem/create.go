package problem

import (
	"main/model"
	"main/server/dao"
)

func CreateProlem(problem *model.Problem) error {
	// 忽略ID字段
	problem.ID = 0

	return dao.InsertProblem(problem)
}
