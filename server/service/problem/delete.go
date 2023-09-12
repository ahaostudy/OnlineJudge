package problem

import "main/server/dao"

func DeleteProblem(id int64) error {
	return dao.DeleteProblem(id)
}
