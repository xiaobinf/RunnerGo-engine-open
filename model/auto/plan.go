package auto

import "github.com/Runner-Go-Team/RunnerGo-engine-open/model"

type Plan struct {
	PlanId        string               `json:"plan_id" bson:"plan_id"`     // 计划id
	PlanName      string               `json:"plan_name" bson:"plan_name"` // 计划名称
	ReportId      string               `json:"report_id" bson:"report_id"` // 报告名称
	Partition     int32                `json:"partition"`
	TeamId        string               `json:"team_id" bson:"team_id"`             // 团队id
	ReportName    string               `json:"report_name" bson:"report_name"`     // 报告名称
	MachineNum    int64                `json:"machine_num" bson:"machine_num"`     // 使用的机器数量
	ConfigTask    *ConfigTask          `json:"config_task" bson:"config_task"`     // 任务配置
	Variable      []*model.PlanKv      `json:"variable" bson:"variable"`           // 全局变量
	Scenes        []model.Scene        `json:"scenes" bson:"scenes"`               // 场景
	Configuration *model.Configuration `json:"configuration" bson:"configuration"` // 场景配置
}
