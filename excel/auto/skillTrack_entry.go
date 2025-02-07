package auto

import (
	"fmt"
	"strings"

	"github.com/east-eden/server/excel"
	"github.com/east-eden/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

var skillTrackEntries *SkillTrackEntries //SkillTrack.csv全局变量

// SkillTrack.csv属性表
type SkillTrackEntry struct {
	Id               int32             `json:"Id,omitempty"`               // 主键
	TrackID          int32             `json:"TrackID,omitempty"`          // 多主键之一
	TrackType        int32             `json:"TrackType,omitempty"`        //track类型
	StartTime        decimal.Decimal   `json:"StartTime,omitempty"`        //开始时间
	DurationTime     decimal.Decimal   `json:"DurationTime,omitempty"`     //持续时间
	EffectIndex      int32             `json:"EffectIndex,omitempty"`      //第几个effect
	AnimName         string            `json:"AnimName,omitempty"`         //动作名称
	FxName           string            `json:"FxName,omitempty"`           //特效名称
	FxSlot           string            `json:"FxSlot,omitempty"`           //特效挂载点
	MoveFxTarget     int32             `json:"MoveFxTarget,omitempty"`     //移动特效目标
	MoveUnit         int32             `json:"MoveUnit,omitempty"`         //移动单位
	Position         []decimal.Decimal `json:"Position,omitempty"`         //单位坐标
	RotateUnit       int32             `json:"RotateUnit,omitempty"`       //移动单位
	Rotation         []decimal.Decimal `json:"Rotation,omitempty"`         //单位旋转
	CameraName       string            `json:"CameraName,omitempty"`       //相机名称
	CameraLookat     int32             `json:"CameraLookat,omitempty"`     //相机朝向
	CameraLookatSlot string            `json:"CameraLookatSlot,omitempty"` //
	CameraFollow     int32             `json:"CameraFollow,omitempty"`     //相机跟随
	CameraFollowSlot string            `json:"CameraFollowSlot,omitempty"` //
	HitAnimName      string            `json:"HitAnimName,omitempty"`      //受击动作
	HitFxName        string            `json:"HitFxName,omitempty"`        //受击特效
	HitFxSlot        string            `json:"HitFxSlot,omitempty"`        //受击特效插槽
	HitStopTime      decimal.Decimal   `json:"HitStopTime,omitempty"`      //受击动作
	Weight           int32             `json:"Weight,omitempty"`           //伤害权重
	FollowDistance   decimal.Decimal   `json:"FollowDistance,omitempty"`   //追击移动
	RetreatDistance  decimal.Decimal   `json:"RetreatDistance,omitempty"`  //击退距离
}

// SkillTrack.csv属性表集合
type SkillTrackEntries struct {
	Rows map[string]*SkillTrackEntry `json:"Rows,omitempty"` //
}

func init() {
	excel.AddEntryLoader("SkillTrack.csv", (*SkillTrackEntries)(nil))
}

func (e *SkillTrackEntries) Load(excelFileRaw *excel.ExcelFileRaw) error {

	skillTrackEntries = &SkillTrackEntries{
		Rows: make(map[string]*SkillTrackEntry, 100),
	}

	for _, v := range excelFileRaw.CellData {
		entry := &SkillTrackEntry{}
		err := mapstructure.Decode(v, entry)
		if !utils.ErrCheck(err, "decode excel data to struct failed", v) {
			return err
		}

		key := fmt.Sprintf("%d+%d", entry.Id, entry.TrackID)
		skillTrackEntries.Rows[key] = entry
	}

	log.Info().Str("excel_file", excelFileRaw.Filename).Msg("excel load success")
	return nil

}

func GetSkillTrackEntry(keys ...int32) (*SkillTrackEntry, bool) {
	keyName := make([]string, 0, len(keys))
	for _, key := range keys {
		keyName = append(keyName, cast.ToString(key))
	}

	finalKey := strings.Join(keyName, "+")
	entry, ok := skillTrackEntries.Rows[finalKey]
	return entry, ok
}

func GetSkillTrackSize() int32 {
	return int32(len(skillTrackEntries.Rows))
}

func GetSkillTrackRows() map[string]*SkillTrackEntry {
	return skillTrackEntries.Rows
}
