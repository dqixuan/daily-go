package model

import "time"

type activity struct {
	Id              int64     `db:"I_ID"`               // 自增主键
	Title           string    `db:"CH_TITLE"`           // 排期标题
	StartAt         time.Time `db:"D_START_AT"`         // 开始时间
	EndAt           time.Time `db:"D_END_AT"`           // 结束时间
	Type            int32     `db:"I_TYPE"`             // 类型 1活动 2公告
	Status          int32     `db:"I_STATUS"`           // 状态 1启用 2未启用 3已删除
	DisplayPosition int32     `db:"I_DISPLAY_POSITION"` // 显示位置 1弹框 2首页 3发现页
	Creator         string    `db:"CH_CREATOR"`         // 创建者
	NoticeTitle     string    `db:"CH_NOTICE_TITLE"`    // 公告标题
	NoticeContent   string    `db:"CH_NOTICE_CONTENT"`  // 公告内容
	Image           string    `db:"CH_IMAGE"`           // 图片链接
	JumpType        int32     `db:"I_JUMP_TYPE"`        // 跳转类型 1H5 2语音房 3营地 4动态
	TargetId        int64     `db:"I_TARGET_ID"`        // 跳转类型2、3、4，目标ID
	Ios             int32     `db:"I_IOS"`              // 跳转H5，1选择ios平台
	IosUrl          string    `db:"CH_IOS_URL"`         // iOS对应链接
	Android         int32     `db:"I_ANDROID"`          // 1选择Android平台
	AndroidUrl      string    `db:"CH_ANDROID_URL"`     // Android对应链接
	ToNew           int32     `db:"I_TO_NEW"`           // 1向新用户曝光
	CreatedAt       time.Time `db:"D_CREATED_AT"`       // 创建时间
	UpdatedAt       time.Time `db:"D_UPDATED_AT"`       // 更新时间
}
