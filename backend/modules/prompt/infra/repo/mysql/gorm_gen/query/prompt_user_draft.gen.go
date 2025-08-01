// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/coze-dev/coze-loop/backend/modules/prompt/infra/repo/mysql/gorm_gen/model"
)

func newPromptUserDraft(db *gorm.DB, opts ...gen.DOOption) promptUserDraft {
	_promptUserDraft := promptUserDraft{}

	_promptUserDraft.promptUserDraftDo.UseDB(db, opts...)
	_promptUserDraft.promptUserDraftDo.UseModel(&model.PromptUserDraft{})

	tableName := _promptUserDraft.promptUserDraftDo.TableName()
	_promptUserDraft.ALL = field.NewAsterisk(tableName)
	_promptUserDraft.ID = field.NewInt64(tableName, "id")
	_promptUserDraft.SpaceID = field.NewInt64(tableName, "space_id")
	_promptUserDraft.PromptID = field.NewInt64(tableName, "prompt_id")
	_promptUserDraft.UserID = field.NewString(tableName, "user_id")
	_promptUserDraft.TemplateType = field.NewString(tableName, "template_type")
	_promptUserDraft.Messages = field.NewString(tableName, "messages")
	_promptUserDraft.ModelConfig = field.NewString(tableName, "model_config")
	_promptUserDraft.VariableDefs = field.NewString(tableName, "variable_defs")
	_promptUserDraft.Tools = field.NewString(tableName, "tools")
	_promptUserDraft.ToolCallConfig = field.NewString(tableName, "tool_call_config")
	_promptUserDraft.BaseVersion = field.NewString(tableName, "base_version")
	_promptUserDraft.IsDraftEdited = field.NewInt32(tableName, "is_draft_edited")
	_promptUserDraft.CreatedAt = field.NewTime(tableName, "created_at")
	_promptUserDraft.UpdatedAt = field.NewTime(tableName, "updated_at")
	_promptUserDraft.DeletedAt = field.NewField(tableName, "deleted_at")

	_promptUserDraft.fillFieldMap()

	return _promptUserDraft
}

// promptUserDraft Draft表
type promptUserDraft struct {
	promptUserDraftDo promptUserDraftDo

	ALL            field.Asterisk
	ID             field.Int64  // 主键ID
	SpaceID        field.Int64  // 空间ID
	PromptID       field.Int64  // Prompt ID
	UserID         field.String // 用户ID
	TemplateType   field.String // 模版类型
	Messages       field.String // 托管消息列表
	ModelConfig    field.String // 模型配置
	VariableDefs   field.String // 变量定义
	Tools          field.String // tools
	ToolCallConfig field.String // tool调用配置
	BaseVersion    field.String // 草稿关联版本
	IsDraftEdited  field.Int32  // 草稿内容是否基于BaseVersion有变更
	CreatedAt      field.Time   // 创建时间
	UpdatedAt      field.Time   // 更新时间
	DeletedAt      field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (p promptUserDraft) Table(newTableName string) *promptUserDraft {
	p.promptUserDraftDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p promptUserDraft) As(alias string) *promptUserDraft {
	p.promptUserDraftDo.DO = *(p.promptUserDraftDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *promptUserDraft) updateTableName(table string) *promptUserDraft {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.SpaceID = field.NewInt64(table, "space_id")
	p.PromptID = field.NewInt64(table, "prompt_id")
	p.UserID = field.NewString(table, "user_id")
	p.TemplateType = field.NewString(table, "template_type")
	p.Messages = field.NewString(table, "messages")
	p.ModelConfig = field.NewString(table, "model_config")
	p.VariableDefs = field.NewString(table, "variable_defs")
	p.Tools = field.NewString(table, "tools")
	p.ToolCallConfig = field.NewString(table, "tool_call_config")
	p.BaseVersion = field.NewString(table, "base_version")
	p.IsDraftEdited = field.NewInt32(table, "is_draft_edited")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *promptUserDraft) WithContext(ctx context.Context) *promptUserDraftDo {
	return p.promptUserDraftDo.WithContext(ctx)
}

func (p promptUserDraft) TableName() string { return p.promptUserDraftDo.TableName() }

func (p promptUserDraft) Alias() string { return p.promptUserDraftDo.Alias() }

func (p promptUserDraft) Columns(cols ...field.Expr) gen.Columns {
	return p.promptUserDraftDo.Columns(cols...)
}

func (p *promptUserDraft) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *promptUserDraft) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 15)
	p.fieldMap["id"] = p.ID
	p.fieldMap["space_id"] = p.SpaceID
	p.fieldMap["prompt_id"] = p.PromptID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["template_type"] = p.TemplateType
	p.fieldMap["messages"] = p.Messages
	p.fieldMap["model_config"] = p.ModelConfig
	p.fieldMap["variable_defs"] = p.VariableDefs
	p.fieldMap["tools"] = p.Tools
	p.fieldMap["tool_call_config"] = p.ToolCallConfig
	p.fieldMap["base_version"] = p.BaseVersion
	p.fieldMap["is_draft_edited"] = p.IsDraftEdited
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p promptUserDraft) clone(db *gorm.DB) promptUserDraft {
	p.promptUserDraftDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p promptUserDraft) replaceDB(db *gorm.DB) promptUserDraft {
	p.promptUserDraftDo.ReplaceDB(db)
	return p
}

type promptUserDraftDo struct{ gen.DO }

func (p promptUserDraftDo) Debug() *promptUserDraftDo {
	return p.withDO(p.DO.Debug())
}

func (p promptUserDraftDo) WithContext(ctx context.Context) *promptUserDraftDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p promptUserDraftDo) ReadDB() *promptUserDraftDo {
	return p.Clauses(dbresolver.Read)
}

func (p promptUserDraftDo) WriteDB() *promptUserDraftDo {
	return p.Clauses(dbresolver.Write)
}

func (p promptUserDraftDo) Session(config *gorm.Session) *promptUserDraftDo {
	return p.withDO(p.DO.Session(config))
}

func (p promptUserDraftDo) Clauses(conds ...clause.Expression) *promptUserDraftDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p promptUserDraftDo) Returning(value interface{}, columns ...string) *promptUserDraftDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p promptUserDraftDo) Not(conds ...gen.Condition) *promptUserDraftDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p promptUserDraftDo) Or(conds ...gen.Condition) *promptUserDraftDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p promptUserDraftDo) Select(conds ...field.Expr) *promptUserDraftDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p promptUserDraftDo) Where(conds ...gen.Condition) *promptUserDraftDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p promptUserDraftDo) Order(conds ...field.Expr) *promptUserDraftDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p promptUserDraftDo) Distinct(cols ...field.Expr) *promptUserDraftDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p promptUserDraftDo) Omit(cols ...field.Expr) *promptUserDraftDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p promptUserDraftDo) Join(table schema.Tabler, on ...field.Expr) *promptUserDraftDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p promptUserDraftDo) LeftJoin(table schema.Tabler, on ...field.Expr) *promptUserDraftDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p promptUserDraftDo) RightJoin(table schema.Tabler, on ...field.Expr) *promptUserDraftDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p promptUserDraftDo) Group(cols ...field.Expr) *promptUserDraftDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p promptUserDraftDo) Having(conds ...gen.Condition) *promptUserDraftDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p promptUserDraftDo) Limit(limit int) *promptUserDraftDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p promptUserDraftDo) Offset(offset int) *promptUserDraftDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p promptUserDraftDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *promptUserDraftDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p promptUserDraftDo) Unscoped() *promptUserDraftDo {
	return p.withDO(p.DO.Unscoped())
}

func (p promptUserDraftDo) Create(values ...*model.PromptUserDraft) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p promptUserDraftDo) CreateInBatches(values []*model.PromptUserDraft, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p promptUserDraftDo) Save(values ...*model.PromptUserDraft) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p promptUserDraftDo) First() (*model.PromptUserDraft, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromptUserDraft), nil
	}
}

func (p promptUserDraftDo) Take() (*model.PromptUserDraft, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromptUserDraft), nil
	}
}

func (p promptUserDraftDo) Last() (*model.PromptUserDraft, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromptUserDraft), nil
	}
}

func (p promptUserDraftDo) Find() ([]*model.PromptUserDraft, error) {
	result, err := p.DO.Find()
	return result.([]*model.PromptUserDraft), err
}

func (p promptUserDraftDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromptUserDraft, err error) {
	buf := make([]*model.PromptUserDraft, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p promptUserDraftDo) FindInBatches(result *[]*model.PromptUserDraft, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p promptUserDraftDo) Attrs(attrs ...field.AssignExpr) *promptUserDraftDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p promptUserDraftDo) Assign(attrs ...field.AssignExpr) *promptUserDraftDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p promptUserDraftDo) Joins(fields ...field.RelationField) *promptUserDraftDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p promptUserDraftDo) Preload(fields ...field.RelationField) *promptUserDraftDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p promptUserDraftDo) FirstOrInit() (*model.PromptUserDraft, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromptUserDraft), nil
	}
}

func (p promptUserDraftDo) FirstOrCreate() (*model.PromptUserDraft, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromptUserDraft), nil
	}
}

func (p promptUserDraftDo) FindByPage(offset int, limit int) (result []*model.PromptUserDraft, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p promptUserDraftDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p promptUserDraftDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p promptUserDraftDo) Delete(models ...*model.PromptUserDraft) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *promptUserDraftDo) withDO(do gen.Dao) *promptUserDraftDo {
	p.DO = *do.(*gen.DO)
	return p
}
