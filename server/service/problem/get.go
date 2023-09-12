package problem

import (
	"main/model"
	"main/server/dao"
)

func GetProblem(id int64) (*model.Problem, error) {
	return dao.GetProblem(id)
}
