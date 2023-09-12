package problem

import "main/internal/gateway/dao"

func DeleteProblem(id int64) error {
	return dao.DeleteProblem(id)
}
