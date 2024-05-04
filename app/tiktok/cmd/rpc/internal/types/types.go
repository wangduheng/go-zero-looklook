package types

type RedPacket struct {
	Name      string `json:"name"`
	JoinCount string `json:"joinCount"`
	Time      string `json:"time"`
	// TotalReds     string   `json:"totalReds"`
	JoinCondition []string `json:"joinCondition"` // 假设这是一个字符串数组
	CT            string   `json:"ct"`            // 假设日期时间以字符串形式存储
}
