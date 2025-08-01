// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNamePromptDebugContext = "prompt_debug_context"

// PromptDebugContext 用户调试prompt上下文信息表
type PromptDebugContext struct {
	ID            int64                 `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`                                         // 主键ID
	PromptID      int64                 `gorm:"column:prompt_id;type:bigint(20);not null;uniqueIndex:uniq_prompt_id_user_id,priority:1;comment:prompt id" json:"prompt_id"`      // prompt id
	UserID        string                `gorm:"column:user_id;type:varchar(128);not null;uniqueIndex:uniq_prompt_id_user_id,priority:2;comment:user id" json:"user_id"`          // user id
	MockContexts  *string               `gorm:"column:mock_contexts;type:longtext;comment:上下文信息，json格式" json:"mock_contexts"`                                                    // 上下文信息，json格式
	MockVariables *string               `gorm:"column:mock_variables;type:longtext;comment:mock变量值，json格式" json:"mock_variables"`                                                // mock变量值，json格式
	MockTools     *string               `gorm:"column:mock_tools;type:longtext;comment:mock tool结果，json格式" json:"mock_tools"`                                                    // mock tool结果，json格式
	DebugConfig   *string               `gorm:"column:debug_config;type:text;comment:调试配置" json:"debug_config"`                                                                  // 调试配置
	CompareConfig *string               `gorm:"column:compare_config;type:longtext;comment:训练场配置" json:"compare_config"`                                                         // 训练场配置
	CreatedAt     time.Time             `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`                               // 创建时间
	UpdatedAt     time.Time             `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`                               // 更新时间
	DeletedAt     soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint(20);not null;column:deleted_at;not null;default:0;softDelete:milli;comment:删除时间" json:"deleted_at"` // 删除时间
}

// TableName PromptDebugContext's table name
func (*PromptDebugContext) TableName() string {
	return TableNamePromptDebugContext
}
