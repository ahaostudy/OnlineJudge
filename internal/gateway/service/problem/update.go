package problem

import "main/internal/gateway/dao"

func UpdateProblem(id int64, problem map[string]any) error {
	delete(problem, "id")
	delete(problem, "author_id")
	return dao.UpdateProblem(id, problem)
}
