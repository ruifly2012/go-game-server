package define

import (
	pbGlobal "github.com/east-eden/server/proto/global"
)

const (
	RankId_Begin             = iota
	RankId_LocalPlayerLevel  // 本服玩家等级榜
	RankId_GlobalPlayerLevel // 全服玩家等级榜
	RankId_End
)

// 排行榜复合键值
type RankKey struct {
	ObjId  int64 `json:"obj_id" bson:"obj_id"`   // 排行榜对象id -- 玩家id或工会id
	RankId int32 `json:"rank_id" bson:"rank_id"` // 排行榜id
}

// 排行榜元数据
type RankMetadata struct {
	RankKey `json:"_id" bson:"_id"` // 排行榜key
	ObjName string                  `json:"name" bson:"name"`   // 排行数据名字 -- 玩家名字
	Score   float64                 `json:"score" bson:"score"` // 排行榜得分
	Date    int64                   `json:"date" bson:"date"`   // 分数更新时间
}

func (r *RankMetadata) FromPB(pb *pbGlobal.RankMetadata) {
	r.ObjId = pb.GetObjId()
	r.ObjName = pb.GetObjName()
	r.Score = pb.GetScore()
	r.Date = pb.GetDate()
}

func (r *RankMetadata) ToPB() *pbGlobal.RankMetadata {
	pb := &pbGlobal.RankMetadata{
		ObjId:   r.ObjId,
		ObjName: r.ObjName,
		Score:   r.Score,
		Date:    r.Date,
	}
	return pb
}
