package db2

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bbkdevadmin/bbm/antlr4/common"
	"github.com/zjbobingtech/bbm-lib/bbcV2"
	"reflect"
	"strings"
)

type Db2Listener struct {
	*BaseDB2zSQLParserListener

	*common.ParserInfo
}

var Db2StatementExcludes = []string{
	"Statement",
	"Context",
	"*db2.",
}

func NewDb2Listener(content string) *Db2Listener {
	return &Db2Listener{
		ParserInfo: common.NewParserInfo(content, make([]bbcV2.TableInfo, 0)),
	}
}

func (s *Db2Listener) EnterTableName(ctx *TableNameContext) {
	if strings.HasPrefix(strings.TrimSpace(s.GetContent(ctx.GetStop().GetStop(), len(s.Content)-1)), ".") {
		return
	}
	tableInfo := bbcV2.TableInfo{
		TableName: s.GetContentByCtx(ctx),
		Type:      s.SqlType,
		Name:      strings.ToUpper(s.GetContentByCtx(ctx)),
		Start:     ctx.GetStart().GetStart(),
		Stop:      ctx.GetStop().GetStop(),
	}
	if s.InSubQuery(ctx) {
		tableInfo.Type = "Select"
	}
	alias := ctx.GetParent().(*antlr.BaseParserRuleContext).GetTypedRuleContext(reflect.TypeOf(CorrelationNameContext{}), 0)
	if alias != nil {
		tableInfo.Alias = s.GetContentByCtx(alias.(*CorrelationNameContext))
	}
	s.TableInfos = append(s.TableInfos, tableInfo)
	s.Tables = append(s.Tables, tableInfo)
}

func (s *Db2Listener) EnterBasicPredicate(ctx *BasicPredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition: s.GetContentByCtx(ctx),
		Operator:  strings.ToUpper(s.GetContentByCtx(ctx.ComparisonOperator())),
	}
	if len(ctx.AllExpression()) > 0 {
		conditionInfo.Left = s.GetContentByCtx(ctx.Expression(0))
		conditionInfo.Right = s.GetContentByCtx(ctx.Expression(1))
	} else {
		conditionInfo.Left = s.GetContentByCtx(ctx.RowValueExpression(0))
		conditionInfo.Right = s.GetContentByCtx(ctx.RowValueExpression(1))
	}
	conditionInfo.ColumnName = conditionInfo.Left
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterQuantifiedPredicate(ctx *QuantifiedPredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition: s.GetContentByCtx(ctx),
	}
	if ctx.ComparisonOperator() != nil {
		conditionInfo.Operator = strings.ToUpper(s.GetContentByCtx(ctx.ComparisonOperator()))
		conditionInfo.Left = s.GetContentByCtx(ctx.Expression())
		conditionInfo.Right = strings.TrimSpace(s.GetContent(ctx.ComparisonOperator().GetStop().GetStop()+1, ctx.GetStop().GetStop()))
	} else if ctx.EQ() != nil {
		conditionInfo.Operator = strings.ToUpper(ctx.EQ().GetText())
		conditionInfo.Left = s.GetContentByCtx(ctx.RowValueExpression())
		conditionInfo.Right = strings.TrimSpace(s.GetContent(ctx.EQ().GetSymbol().GetStop()+1, ctx.GetStop().GetStop()))
	} else {
		conditionInfo.Operator = strings.ToUpper(ctx.NE().GetText())
		conditionInfo.Left = s.GetContentByCtx(ctx.RowValueExpression())
		conditionInfo.Right = strings.TrimSpace(s.GetContent(ctx.NE().GetSymbol().GetStop()+1, ctx.GetStop().GetStop()))
	}
	conditionInfo.ColumnName = conditionInfo.Left
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterArrayExistsPredicate(ctx *ArrayExistsPredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition: s.GetContentByCtx(ctx),
		Operator:  "ARRAY_EXISTS",
	}
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterBetweenPredicate(ctx *BetweenPredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition:  s.GetContentByCtx(ctx),
		Operator:   "BETWEEN",
		Left:       s.GetContentByCtx(ctx.Expression(0)),
		Right:      s.GetContent(ctx.Expression(1).GetStop().GetStop(), ctx.GetStop().GetStop()),
		ColumnName: s.GetContentByCtx(ctx.Expression(0)),
	}
	if ctx.NOT() != nil {
		conditionInfo.Operator = "NOT BETWEEN"
	}
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterDistinctPredicate(ctx *DistinctPredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition:  s.GetContentByCtx(ctx),
		Operator:   "DISTINCT",
		Left:       s.GetContentByCtx(ctx.Expression(0)),
		Right:      s.GetContentByCtx(ctx.Expression(1)),
		ColumnName: s.GetContentByCtx(ctx.Expression(0)),
	}
	if ctx.NOT() != nil {
		conditionInfo.Operator = "NOT DISTINCT"
	}
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterExistsPredicate(ctx *ExistsPredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition: s.GetContentByCtx(ctx),
		Operator:  "EXISTS",
	}
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterInPredicate(ctx *InPredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition:  s.GetContentByCtx(ctx),
		Operator:   "IN",
		Left:       s.GetContentByCtx(ctx.Expression(0)),
		Right:      s.GetContent(ctx.IN().GetSymbol().GetStop()+1, ctx.GetStop().GetStop()),
		ColumnName: s.GetContentByCtx(ctx.Expression(0)),
	}
	if ctx.NOT() != nil {
		conditionInfo.Operator = "NOT IN"
	}
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterLikePredicate(ctx *LikePredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition:  s.GetContentByCtx(ctx),
		Operator:   "LIKE",
		Left:       s.GetContentByCtx(ctx.Expression(0)),
		Right:      s.GetContent(ctx.Expression(1).GetStop().GetStop(), ctx.GetStop().GetStop()),
		ColumnName: s.GetContentByCtx(ctx.Expression(0)),
	}
	if ctx.NOT() != nil {
		conditionInfo.Operator = "NOT LIKE"
	}
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterNullPredicate(ctx *NullPredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition:  s.GetContentByCtx(ctx),
		Operator:   "NULL",
		Left:       s.GetContentByCtx(ctx.Expression()),
		ColumnName: s.GetContentByCtx(ctx.Expression()),
	}
	if ctx.NOT() != nil {
		conditionInfo.Operator = "NOT NULL"
	}
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterXmlExistsPredicate(ctx *XmlExistsPredicateContext) {
	if s.SubQueryStart > -1 {
		return
	}
	conditionInfo := bbcV2.Condition{
		Condition: s.GetContentByCtx(ctx),
		Operator:  "XMLEXISTS",
	}
	s.Conditions = append(s.Conditions, conditionInfo)
}

func (s *Db2Listener) EnterSqlStatement(ctx *SqlStatementContext) {
	var i int
	if ctx.EXEC_SQL() != nil {
		i++
	}
	statement := s.GetStatement(ctx.GetChild(0), Db2StatementExcludes)
	if strings.ToUpper(statement) == "query" {
		statement = "select"
		s.SubQueryCount--
	}
	s.SqlType = statement
	s.Statement = statement
}

func (s *Db2Listener) ExitSqlStatement(ctx *SqlStatementContext) {
	s.GetColumns(ctx)
}

func (s *Db2Listener) EnterSubSelect(ctx *SubSelectContext) {
	s.SubQueryCount++
	if s.SubQueryCount > 0 {
		s.EnterSubQuery(ctx)
	}
}

func (s *Db2Listener) EnterExplainPlanClause(ctx *ExplainPlanClauseContext) {
	sqlType := "Explain"
	if ctx.Query() != nil {
		sqlType = "Select"
		s.SubQueryCount--
	} else if ctx.SearchedDelete() != nil {
		sqlType = "Delete"
	} else if ctx.SearchedUpdate() != nil {
		sqlType = "Update"
	} else if ctx.InsertStatement() != nil {
		sqlType = "Insert"
	} else if ctx.MergeStatement() != nil {
		sqlType = "Merge"
	}
	s.SqlType = sqlType
}

func (s *Db2Listener) ExitSubSelect(ctx *SubSelectContext) {
	if s.SubQueryCount > 0 {
		s.GetColumns(ctx)
		tables := make([]bbcV2.TableInfo, 0)
		for _, tableInfo := range s.Tables {
			if !(tableInfo.Start >= ctx.GetStart().GetStart() && tableInfo.Stop <= ctx.GetStop().GetStop()) {
				tables = append(tables, tableInfo)
			}
		}
		s.Tables = tables
	}
	s.ExitSubQuery(ctx)
}

func (s *Db2Listener) EnterFunctionInvocation(ctx *FunctionInvocationContext) {
	s.Functions = append(s.Functions, s.GetContentByCtx(ctx))
}

func (s *Db2Listener) GetColumns(ctx antlr.ParserRuleContext) {
	tables := make([]bbcV2.TableInfo, 0)
	tableNameMap := make(map[string]bbcV2.TableInfo)
	tableAliasMap := make(map[string]bbcV2.TableInfo)
	for _, tableInfo := range s.Tables {
		if _, ok := tableNameMap[tableInfo.Name]; tableInfo.Start >= ctx.GetStart().GetStart() && tableInfo.Stop <= ctx.GetStop().GetStop() && !ok {
			tables = append(tables, tableInfo)
			tableNameMap[tableInfo.TableName] = tableInfo
			if tableInfo.Alias != "" {
				tableAliasMap[tableInfo.Alias] = tableInfo
			}
		}
	}
	listener := NewDb2ColumnListener(s.Content, tables, tableNameMap, tableAliasMap, ctx.GetStart().GetStart(), ctx.GetStop().GetStop())
	antlr.ParseTreeWalkerDefault.Walk(listener, ctx)
	if len(listener.ColumnInfos) > 0 {
		s.ColumnInfos = append(s.ColumnInfos, listener.ColumnInfos...)
	}
}
