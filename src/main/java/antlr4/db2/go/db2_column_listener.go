package db2

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bbkdevadmin/bbm/antlr4/common"
	"github.com/zjbobingtech/bbm-lib/bbcV2"
	"reflect"
	"strings"
)

type Db2ColumnListener struct {
	*BaseDB2zSQLParserListener

	*common.ParserColumnInfo
}

func NewDb2ColumnListener(content string, tables []bbcV2.TableInfo, tableNameMap, tableAliasMap map[string]bbcV2.TableInfo, start, stop int) *Db2ColumnListener {
	return &Db2ColumnListener{
		ParserColumnInfo: common.NewParserColumnInfo(content, tables, tableNameMap, tableAliasMap, start, stop),
	}
}

func (s *Db2ColumnListener) EnterSelectColumns(ctx *SelectColumnsContext) {
	if s.SubQueryStart > -1 || ctx.SPLAT() == nil {
		return
	}
	columnInfo := bbcV2.Column{
		Column:    "*",
		TableName: "UNKNOWN",
	}
	if ctx.TableName() != nil {
		columnInfo.TableName = s.GetContentByCtx(ctx.TableName())
	} else if len(s.TableInfos) == 1 {
		columnInfo.TableName = s.TableInfos[0].TableName
	}
	s.ColumnInfos = append(s.ColumnInfos, columnInfo)
}

func (s *Db2ColumnListener) EnterColumnName(ctx *ColumnNameContext) {
	if s.SubQueryStart > -1 {
		return
	}
	columnInfo := bbcV2.Column{
		TableName: "UNKNOWN",
	}
	if s.WhereStart > -1 && ctx.GetStart().GetStart() >= s.WhereStart && ctx.GetStop().GetStop() <= s.WhereStop {
		columnInfo.Where = true
	}
	dataType := ctx.GetParent().(*antlr.BaseParserRuleContext).GetTypedRuleContext(reflect.TypeOf(DataTypeContext{}), 0)
	if dataType != nil {
		columnInfo.ColumnType = s.GetContentByCtx(dataType.(*DataTypeContext))
	}
	cs := strings.Split(s.GetContentByCtx(ctx), ".")
	var table string
	switch len(cs) {
	case 2:
		table = strings.ToUpper(cs[0])
		columnInfo.Column = cs[1]
		if info, ok := s.TableNameMap[table]; table != "" && ok {
			columnInfo.TableName = info.TableName
		} else if info, ok := s.TableAliasMap[table]; table != "" && ok {
			columnInfo.TableName = info.TableName
		} else {
			columnInfo.TableName = table
		}
	case 3:
		table = strings.ToUpper(cs[0] + "." + cs[1])
		columnInfo.Column = cs[2]
		if info, ok := s.TableNameMap[table]; table != "" && ok {
			columnInfo.TableName = info.TableName
		} else if info, ok := s.TableAliasMap[table]; table != "" && ok {
			columnInfo.TableName = info.TableName
		} else {
			columnInfo.TableName = table
		}
	default:
		if len(s.TableInfos) == 1 {
			columnInfo.TableName = s.TableInfos[0].TableName
		}
		columnInfo.Column = s.GetContentByCtx(ctx)
	}
	columnInfo.FullName = columnInfo.TableName + "." + columnInfo.Column
	s.ColumnInfos = append(s.ColumnInfos, columnInfo)
}

func (s *Db2ColumnListener) EnterSubSelect(ctx *SubSelectContext) {
	s.EnterSubQuery(ctx)
}

func (s *Db2ColumnListener) ExitSubSelect(ctx *SubSelectContext) {
	s.ExitSubQuery(ctx)
}

func (s *Db2ColumnListener) EnterSearchCondition(ctx *SearchConditionContext) {
	s.EnterWhere(ctx)
}

func (s *Db2ColumnListener) ExitSearchCondition(ctx *SearchConditionContext) {
	s.ExitWhere(ctx)
}
