package define

// 一级属性
const (
	Att_Begin       = 0 // 32位属性开始
	Att_SecondBegin = 0 // 二级属性开始

	Att_Atk            = 0  // 0 攻击力最终值
	Att_Armor          = 1  // 1 护甲最终值
	Att_SelfDmgInc     = 2  // 2 自身伤害加成
	Att_EnemyWoundInc  = 3  // 3 敌方受伤加成
	Att_SelfDmgDec     = 4  // 4 自身伤害消减
	Att_EnemyWoundDec  = 5  // 5 敌方受伤消减
	Att_Crit           = 6  // 6 暴击值
	Att_CritInc        = 7  // 7 暴击倍数加层
	Att_Tenacity       = 8  // 8 韧性
	Att_Heal           = 9  // 9 治疗强度最终值
	Att_RealDmg        = 10 // 10 真实伤害
	Att_MoveSpeed      = 11 // 11 移动速度最终值
	Att_AtbSpeed       = 12 // 12 时间槽速度最终值
	Att_EffectHit      = 13 // 13 技能效果命中
	Att_EffectResist   = 14 // 14 技能效果抵抗
	Att_MaxHP          = 15 // 15 生命值上限最终值
	Att_CurHP          = 16 // 16 当前生命值
	Att_MaxMP          = 17 // 17 蓝量上限
	Att_CurMP          = 18 // 18 当前蓝量
	Att_GenMP          = 19 // 19 mp恢复值
	Att_MaxRage        = 20 // 20 怒气值上限
	Att_GenRagePercent = 21 // 21 回怒百分比
	Att_InitRage       = 22 // 22 初始怒气
	Att_Hit            = 23 // 23 命中
	Att_Dodge          = 24 // 24 闪避
	Att_MoveScope      = 25 // 25 移动范围

	Att_DmgTypeBegin = 26 // 26 各系伤害加成begin
	Att_DmgPhysics   = 26 // 26 物理系伤害加成
	Att_DmgEarth     = 27 // 27 地系伤害加成
	Att_DmgWater     = 28 // 28 水系伤害加成
	Att_DmgFire      = 29 // 29 火系伤害加成
	Att_DmgWind      = 30 // 30 风系伤害加成
	Att_DmgThunder   = 31 // 31 雷系伤害加成
	Att_DmgTime      = 32 // 32 时系伤害加成
	Att_DmgSpace     = 33 // 33 空系伤害加成
	Att_DmgSteel     = 34 // 34 钢系伤害加成
	Att_DmgDeath     = 35 // 35 灭系伤害加成
	Att_DmgTypeEnd   = 36 // 36 各系伤害加成end

	Att_ResTypeBegin = 36 // 36 各系伤害抗性
	Att_ResPhysics   = 36 // 36 物理系伤害抗性
	Att_ResEarth     = 37 // 37 地系伤害抗性
	Att_ResWater     = 38 // 38 水系伤害抗性
	Att_ResFire      = 39 // 39 火系伤害抗性
	Att_ResWind      = 40 // 40 风系伤害抗性
	Att_ResThunder   = 41 // 41 雷系伤害抗性
	Att_ResTime      = 42 // 42 时系伤害抗性
	Att_ResSpace     = 43 // 43 空系伤害抗性
	Att_ResSteel     = 44 // 44 钢系伤害抗性
	Att_ResDeath     = 45 // 45 灭系伤害抗性
	Att_ResTypeEnd   = 46 // 46 各系伤害抗性

	Att_SecondEnd = 46 // 二级属性结束
	Att_End       = 46 // 32位属性结束
)

// 基础属性枚举
const (
	Att_Base_Begin    = 0
	Att_AtkBase       = 0 // 0 攻击力基础值
	Att_ArmorBase     = 1 // 1 护甲基础值
	Att_HealBase      = 2 // 2 治疗强度基础值
	Att_MoveSpeedBase = 3 // 3 移动速度基础值
	Att_AtbSpeedBase  = 4 // 4 时间槽速度基础值
	Att_MaxHPBase     = 5 // 5 生命值上限基础值
	Att_Base_End      = 6
)

// 百分比属性枚举
const (
	Att_Percent_Begin    = 0
	Att_AtkPercent       = 0 // 0 攻击力百分比
	Att_ArmorPercent     = 1 // 1 护甲百分比
	Att_HealPercent      = 2 // 2 治疗强度百分比
	Att_MoveSpeedPercent = 3 // 3 移动速度百分比
	Att_AtbSpeedPercent  = 4 // 4 时间槽速度百分比
	Att_MaxHPPercent     = 5 // 5 生命值上限百分比
	Att_Percent_End      = 6
)

const (
	AttFinalNum   = Att_End - Att_Begin // 32位属性数量
	AttBaseNum    = Att_Base_End - Att_Base_Begin
	AttPercentNum = Att_Percent_End - Att_Percent_Begin
)

// combat legacy: do not edit!
const (
	Att_Block  = 9  // 格挡
	Att_Broken = 10 // 破击
	Att_DmgDec = 14 // 伤害减免
)

// 属性名
var AttNames = [Att_End]string{
	Att_Atk:            "攻击力",
	Att_SelfDmgInc:     "自身伤害加成",
	Att_EnemyWoundInc:  "敌方受伤加成",
	Att_SelfDmgDec:     "自身伤害消减",
	Att_EnemyWoundDec:  "敌方受伤消减",
	Att_Crit:           "暴击值",
	Att_CritInc:        "暴击倍数加层",
	Att_Tenacity:       "韧性",
	Att_Heal:           "治疗强度",
	Att_RealDmg:        "真实伤害",
	Att_MoveSpeed:      "战场移动速度",
	Att_AtbSpeed:       "时间槽速度",
	Att_EffectHit:      "技能效果命中",
	Att_EffectResist:   "技能效果抵抗",
	Att_MaxHP:          "生命值上限",
	Att_CurHP:          "当前生命值",
	Att_MaxMP:          "蓝量上限",
	Att_CurMP:          "当前蓝量",
	Att_GenMP:          "mp恢复值",
	Att_MaxRage:        "怒气值",
	Att_GenRagePercent: "怒气增长提高百分比",
	Att_InitRage:       "初始怒气",
	Att_Hit:            "命中",
	Att_Dodge:          "闪避",
	Att_MoveScope:      "移动范围",

	Att_DmgPhysics: "物理系伤害加成",
	Att_DmgEarth:   "地系伤害加成",
	Att_DmgWater:   "水系伤害加成",
	Att_DmgFire:    "火系伤害加成",
	Att_DmgWind:    "风系伤害加成",
	Att_DmgThunder: "雷系伤害加成",
	Att_DmgTime:    "时系伤害加成",
	Att_DmgSpace:   "空系伤害加成",
	Att_DmgSteel:   "钢系伤害加成",
	Att_DmgDeath:   "灭系伤害加成",

	Att_ResPhysics: "物理系伤害抗性",
	Att_ResEarth:   "地系伤害抗性",
	Att_ResWater:   "水系伤害抗性",
	Att_ResFire:    "火系伤害抗性",
	Att_ResWind:    "风系伤害抗性",
	Att_ResThunder: "雷系伤害抗性",
	Att_ResTime:    "时系伤害抗性",
	Att_ResSpace:   "空系伤害抗性",
	Att_ResSteel:   "钢系伤害抗性",
	Att_ResDeath:   "灭系伤害抗性",
}

// 基础属性名
var AttBaseNames = [Att_Base_End]string{
	Att_AtkBase:       "攻击力",
	Att_ArmorBase:     "护甲",
	Att_HealBase:      "治疗强度",
	Att_MoveSpeedBase: "战场移动速度",
	Att_AtbSpeedBase:  "时间槽速度",
	Att_MaxHPBase:     "生命值上限",
}

// 百分比属性名
var AttPercentNames = [Att_Percent_End]string{
	Att_AtkPercent:       "攻击力百分比",
	Att_ArmorPercent:     "护甲百分比",
	Att_HealPercent:      "治疗强度百分比",
	Att_MoveSpeedPercent: "战场移动速度百分比",
	Att_AtbSpeedPercent:  "时间槽速度百分比",
	Att_MaxHPPercent:     "生命值上限百分比",
}
