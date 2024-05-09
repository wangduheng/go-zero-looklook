// Code generated by goctl. DO NOT EDIT.
package types

type TitokLive struct {
	OpenUrl    string `json:"open_url"`
	Info       string `json:"info"`
	Name       string `json:"name"`
	Id         int64  `db:"id"`
	CreateTime int64  `db:"create_time"`
	OpenTime int64    `db:"open_string"`
	IsNeedFork int64 

}

type SaveLiveReq struct {
	OpenUrl string `json:"open_url"`
	Info    string `json:"info"`
	Name    string `json:"name"`
}

type SaveLiveResp struct {
	Id int64 `json:"id"`
	Content string `json:"content"`
}

type FetchLiveResp struct {
	List []TitokLive `json:"list"`
}

type RedPacket struct {
	Name      string `json:"name"`
	JoinCount string `json:"joinCount"`
	Time      string `json:"time"`
	// TotalReds     string   `json:"totalReds"`
	JoinCondition []string `json:"joinCondition"` // 假设这是一个字符串数组
	CT            string   `json:"ct"`            // 假设日期时间以字符串形式存储
}


type ReceiveReq struct {
	Id int64 `json:"Id"`
}

