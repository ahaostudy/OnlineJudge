package mq

import (
	"encoding/json"
	"main/config"
	"main/internal/data/model"
	"path/filepath"
)

// GenerateJudgeMQMsg 生成判题服务消息
func GenerateJudgeMQMsg(judgeID string, codePath string, langID int64, problem *model.Problem) ([]byte, error) {
	req := PrivateJudgeRequest{
		JudgeID:  judgeID,
		Codepath: filepath.Join(config.ConfJudge.File.CodePath, codePath),
		LangID:   int(langID),
		Problem:  problem,
	}
	msg, err := json.Marshal(req)
	return msg, err
}

// GeneratePrivateJudgeMQMsg 生成判题服务消息
func GeneratePrivateJudgeMQMsg(judgeID string, codePath string, langID int64, problem *model.Problem) ([]byte, error) {
	req := PrivateJudgeRequest{
		JudgeID:  judgeID,
		Codepath: filepath.Join(config.ConfJudge.File.CodePath, codePath),
		LangID:   int(langID),
		Problem:  problem,
	}
	msg, err := json.Marshal(req)
	return msg, err
}

// GenerateContestSubmitMQMsg 生成比赛提交服务消息
func GenerateContestSubmitMQMsg(contestID, userID int64) ([]byte, error) {
	req := ContestSubmitRequest{
		ContestID: contestID,
		UserID:    userID,
	}
	msg, err := json.Marshal(req)
	return msg, err
}
