package mq

import (
	"encoding/json"
	"main/config"
	"main/internal/data/model"
	"path/filepath"
)

// GenerateJudgeMQMsg 生成判题服务消息
func GenerateJudgeMQMsg(judgeID string, codePath string, langID int64, problem *model.Problem) ([]byte, error) {
	req := judgeRequest{
		JudgeID:  judgeID,
		Codepath: filepath.Join(config.ConfJudge.File.CodePath, codePath),
		LangID:   int(langID),
		Problem:  problem,
	}
	msg, err := json.Marshal(req)
	return msg, err
}

// GenerateContestSubmitMQMsg 生成提交服务消息
func GenerateContestSubmitMQMsg(contestID, problemID, userID int64) ([]byte, error) {
	req := contestSubmitRequest{
		ContestID: contestID,
		ProblemID: problemID,
		UserID:    userID,
	}
	msg, err := json.Marshal(req)
	return msg, err
}
