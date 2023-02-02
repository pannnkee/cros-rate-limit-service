package dto

// DecisionReq 决策入参
type DecisionReq struct {
	Service  string                 `json:"service"`  // 服务名
	Key      string                 `json:"key"`      // 决策流key
	Version  string                 `json:"version"`  // 决策流版本
	Features map[string]interface{} `json:"features"` // 特征
}
