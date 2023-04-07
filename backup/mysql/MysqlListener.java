package com.thinice.bbc.core.parser.antlr.mysql;

import com.thinice.bbc.core.parser.antlr.parser.ParserInfo;
import com.thinice.bbc.core.parser.dto.info.ConditionInfo;
import com.thinice.bbc.core.parser.dto.info.TableInfo;
import lombok.Data;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.tree.ParseTreeWalker;

import java.util.*;

@Data
public class MysqlListener extends MySqlParserBaseListener {

    private ParserInfo info;
    private int statementCount = 0;

    public MysqlListener(String content) {
        info = new ParserInfo();
        info.setContent(content);
    }


    @Override
    public void enterCallStatement(MySqlParser.CallStatementContext ctx) {
        info.getFunctions().add(info.getContent(ctx).replaceFirst("(?i)CALL", "").trim());
    }
    @Override
    public void enterSpecificFunctionCall(MySqlParser.SpecificFunctionCallContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }


    @Override
    public void enterAggregateFunctionCall(MySqlParser.AggregateFunctionCallContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }


    @Override
    public void enterNonAggregateFunctionCall(MySqlParser.NonAggregateFunctionCallContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }


    @Override
    public void enterScalarFunctionCall(MySqlParser.ScalarFunctionCallContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }


    @Override
    public void enterUdfFunctionCall(MySqlParser.UdfFunctionCallContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }

    @Override
    public void enterPasswordFunctionCall(MySqlParser.PasswordFunctionCallContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }

    @Override
    public void enterTableName(MySqlParser.TableNameContext ctx) {
        TableInfo tableInfo = new TableInfo();
        tableInfo.setTable_name(info.getContent(ctx));
        tableInfo.setType(info.getSqlType());
        if (ctx.start.getStartIndex() >= info.getSubQueryStart() && ctx.stop.getStopIndex() <= info.getSubQueryStop()) {
            tableInfo.setType("Select");
        }
        tableInfo.setName(tableInfo.getTable_name().toUpperCase());
        tableInfo.setStart(ctx.start.getStartIndex());
        tableInfo.setStop(ctx.stop.getStartIndex());
        tableInfo.setAlias("");
        if (ctx.getParent() instanceof MySqlParser.AtomTableItemContext) {
            MySqlParser.UidContext uid = ((MySqlParser.AtomTableItemContext) ctx.getParent()).uid();
            if (null != uid) {
                tableInfo.setAlias(info.getContent(uid).toUpperCase());
            }
        }
        info.getTableInfos().add(tableInfo);
        info.getTables().add(tableInfo);
    }
    @Override
    public void enterSingleDeleteStatement(MySqlParser.SingleDeleteStatementContext ctx) {
        if (statementCount == 0 && null != ctx.expression()) {
            info.setConditions(getConditions(ctx.expression()));
        }
        statementCount++;
    }
    @Override
    public void enterMultipleDeleteStatement(MySqlParser.MultipleDeleteStatementContext ctx) {
        if (statementCount == 0 && null != ctx.expression()) {
            info.setConditions(getConditions(ctx.expression()));
        }
        statementCount++;
    }
    @Override
    public void enterHandlerReadIndexStatement(MySqlParser.HandlerReadIndexStatementContext ctx) {
        if (statementCount == 0 && null != ctx.expression()) {
            info.setConditions(getConditions(ctx.expression()));
        }
        statementCount++;
    }
    @Override
    public void enterHandlerReadStatement(MySqlParser.HandlerReadStatementContext ctx) {
        if (statementCount == 0 && null != ctx.expression()) {
            info.setConditions(getConditions(ctx.expression()));
        }
        statementCount++;
    }
    @Override
    public void enterSingleUpdateStatement(MySqlParser.SingleUpdateStatementContext ctx) {
        if (statementCount == 0 && null != ctx.expression()) {
            info.setConditions(getConditions(ctx.expression()));
        }
        statementCount++;
    }
    @Override
    public void enterMultipleUpdateStatement(MySqlParser.MultipleUpdateStatementContext ctx) {
        if (statementCount == 0 && null != ctx.expression()) {
            info.setConditions(getConditions(ctx.expression()));
        }
        statementCount++;
    }
    @Override
    public void enterSimpleSelect(MySqlParser.SimpleSelectContext ctx) {
        MySqlParser.FromClauseContext from = ctx.querySpecification().fromClause();
        if (null != from && statementCount == 0 && null != from.expression()) {
            info.setConditions(getConditions(from.expression()));
        }
        statementCount++;
        info.subQueryCount++;
        if (info.subQueryCount > 0) {
            info.enterSubQuery(ctx);
        }
    }
    @Override
    public void exitSimpleSelect(MySqlParser.SimpleSelectContext ctx) {
        if (info.subQueryCount > 0) {
            getColumns(ctx);
            Iterator<TableInfo> it = info.getTables().iterator();
            while (it.hasNext()) {
                TableInfo tableInfo = it.next();
                if (tableInfo.getStart() >= ctx.start.getStartIndex() && tableInfo.getStop() <= ctx.stop.getStopIndex()) {
                    it.remove();
                }
            }
        }
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterParenthesisSelect(MySqlParser.ParenthesisSelectContext ctx) {
        info.subQueryCount++;
        if (info.subQueryCount > 0) {
            info.enterSubQuery(ctx);
        }
    }
    @Override
    public void exitParenthesisSelect(MySqlParser.ParenthesisSelectContext ctx) {
        if (info.subQueryCount > 0) {
            getColumns(ctx);
            Iterator<TableInfo> it = info.getTables().iterator();
            while (it.hasNext()) {
                TableInfo tableInfo = it.next();
                if (tableInfo.getStart() >= ctx.start.getStartIndex() && tableInfo.getStop() <= ctx.stop.getStopIndex()) {
                    it.remove();
                }
            }
        }
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterUnionSelect(MySqlParser.UnionSelectContext ctx) {
        MySqlParser.FromClauseContext from = ctx.querySpecification().fromClause();
        if (null != from && statementCount == 0 && null != from.expression()) {
            info.setConditions(getConditions(from.expression()));
        }
        statementCount++;
        info.subQueryCount++;
        if (info.subQueryCount > 0) {
            info.enterSubQuery(ctx);
        }
    }
    @Override
    public void exitUnionSelect(MySqlParser.UnionSelectContext ctx) {
        if (info.subQueryCount > 0) {
            getColumns(ctx);
            Iterator<TableInfo> it = info.getTables().iterator();
            while (it.hasNext()) {
                TableInfo tableInfo = it.next();
                if (tableInfo.getStart() >= ctx.start.getStartIndex() && tableInfo.getStop() <= ctx.stop.getStopIndex()) {
                    it.remove();
                }
            }
        }
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterUnionParenthesisSelect(MySqlParser.UnionParenthesisSelectContext ctx) {
        info.subQueryCount++;
        if (info.subQueryCount > 0) {
            info.enterSubQuery(ctx);
        }
    }
    @Override
    public void exitUnionParenthesisSelect(MySqlParser.UnionParenthesisSelectContext ctx) {
        if (info.subQueryCount > 0) {
            getColumns(ctx);
            Iterator<TableInfo> it = info.getTables().iterator();
            while (it.hasNext()) {
                TableInfo tableInfo = it.next();
                if (tableInfo.getStart() >= ctx.start.getStartIndex() && tableInfo.getStop() <= ctx.stop.getStopIndex()) {
                    it.remove();
                }
            }
        }
        info.exitSubQuery(ctx);
    }

    @Override
    public void enterWithLateralStatement(MySqlParser.WithLateralStatementContext ctx) {
        MySqlParser.FromClauseContext from = ctx.querySpecificationNointo().fromClause();
        if (null != from && statementCount == 0 && null != from.expression()) {
            info.setConditions(getConditions(from.expression()));
        }
        statementCount++;
        info.subQueryCount++;
        if (info.subQueryCount > 0) {
            info.enterSubQuery(ctx);
        }
    }
    @Override
    public void exitWithLateralStatement(MySqlParser.WithLateralStatementContext ctx) {
        if (info.subQueryCount > 0) {
            getColumns(ctx);
            Iterator<TableInfo> it = info.getTables().iterator();
            while (it.hasNext()) {
                TableInfo tableInfo = it.next();
                if (tableInfo.getStart() >= ctx.start.getStartIndex() && tableInfo.getStop() <= ctx.stop.getStopIndex()) {
                    it.remove();
                }
            }
        }
        info.exitSubQuery(ctx);
    }


    @Override
    public void enterDdlStatement(MySqlParser.DdlStatementContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("Statement", "");
        info.setStatement(statement);
        info.setSqlType(statement);
    }
    @Override
    public void exitDdlStatement(MySqlParser.DdlStatementContext ctx) {
        getColumns(ctx);
    }
    @Override
    public void enterDmlStatement(MySqlParser.DmlStatementContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("Statement", "");
        info.setStatement(statement);
        info.setSqlType(statement);
        if (null != ctx.selectStatement()) {
            info.subQueryCount--;
        }
    }
    @Override
    public void exitDmlStatement(MySqlParser.DmlStatementContext ctx) {
        getColumns(ctx);
    }
    @Override
    public void enterTransactionStatement(MySqlParser.TransactionStatementContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("Statement", "");
        info.setStatement(statement);
        info.setSqlType(statement);
    }
    @Override
    public void exitTransactionStatement(MySqlParser.TransactionStatementContext ctx) {
        getColumns(ctx);
    }
    @Override
    public void enterReplicationStatement(MySqlParser.ReplicationStatementContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("Statement", "");
        info.setStatement(statement);
        info.setSqlType(statement);
    }
    @Override
    public void exitReplicationStatement(MySqlParser.ReplicationStatementContext ctx) {
        getColumns(ctx);
    }
    @Override
    public void enterPreparedStatement(MySqlParser.PreparedStatementContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("Statement", "");
        info.setStatement(statement);
        info.setSqlType(statement);
    }
    @Override
    public void exitPreparedStatement(MySqlParser.PreparedStatementContext ctx) {
        getColumns(ctx);
    }
    @Override
    public void enterAdministrationStatement(MySqlParser.AdministrationStatementContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("Statement", "");
        info.setStatement(statement);
        info.setSqlType(statement);
    }
    @Override
    public void exitAdministrationStatement(MySqlParser.AdministrationStatementContext ctx) {
        getColumns(ctx);
    }
    @Override
    public void enterUtilityStatement(MySqlParser.UtilityStatementContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("Statement", "");
        info.setStatement(statement);
        info.setSqlType(statement);
    }
    @Override
    public void exitUtilityStatement(MySqlParser.UtilityStatementContext ctx) {
        getColumns(ctx);
    }
    @Override
    public void enterSimpleDescribeStatement(MySqlParser.SimpleDescribeStatementContext ctx) {
        String sql = info.getContent(ctx).trim();
        String statement;
        if (sql.toUpperCase().startsWith("DESCRIBE")) {
            statement = "Describe";
        } else if (sql.toUpperCase().startsWith("DESC")) {
            statement = "Desc";
        } else {
            statement = "Explain";
        }
        info.setStatement(statement);
        info.setSqlType(statement);
    }
    @Override
    public void enterFullDescribeStatement(MySqlParser.FullDescribeStatementContext ctx) {
        String statement;
        String sql = info.getContent(ctx).trim();
        if (sql.toUpperCase().startsWith("DESCRIBE")) {
            statement = "Describe";
        } else if (sql.toUpperCase().startsWith("DESC")) {
            statement = "Desc";
        } else {
            statement = "Explain";
        }
        info.setStatement(statement);
        info.setSqlType(statement);
    }

    @Override
    public void enterDescribeStatements(MySqlParser.DescribeStatementsContext ctx) {
        String sqlType = "Explain";
        if (null != ctx.selectStatement()) {
            sqlType = "Select";
            info.subQueryCount--;
        } else if (null != ctx.deleteStatement()) {
            sqlType = "Delete";
        } else if (null != ctx.updateStatement()) {
            sqlType = "Update";
        } else if (null != ctx.insertStatement()) {
            sqlType = "Insert";
        } else if (null != ctx.replaceStatement()) {
            sqlType = "Replace";
        }
        info.setSqlType(sqlType);
    }

    private List<ConditionInfo> getConditions(MySqlParser.ExpressionContext ctx) {
        MysqlExpressionListener listener = new MysqlExpressionListener(info.getContent());
        ParseTreeWalker.DEFAULT.walk(listener, ctx);
        return listener.getInfo().getConditions();
    }

    private void getColumns(ParserRuleContext ctx) {
        List<TableInfo> tables = new ArrayList<>();
        Map<String, TableInfo> tableNameMap = new HashMap<>();
        Map<String, TableInfo> tableAliasMap = new HashMap<>();
        for (TableInfo tableInfo : info.getTables()) {
            if (tableInfo.getStart() >= ctx.start.getStartIndex() && tableInfo.getStop() <= ctx.stop.getStartIndex() && !tableNameMap.containsKey(tableInfo.getName())) {
                tables.add(tableInfo);
                tableNameMap.put(tableInfo.getName(), tableInfo);
                if (null != tableInfo.getAlias() && !tableInfo.getAlias().equals("")) {
                    tableAliasMap.put(tableInfo.getAlias(), tableInfo);
                }
            }
        }
        MysqlColumnListener listener = new MysqlColumnListener(info.getContent(), ctx.start.getStartIndex(), ctx.stop.getStopIndex());
        listener.getInfo().setTableInfos(tables);
        listener.getInfo().setTableNameMap(tableNameMap);
        listener.getInfo().setTableAliasMap(tableAliasMap);
        ParseTreeWalker.DEFAULT.walk(listener, ctx);
        if (null != listener.getInfo().getColumnInfos()) {
            info.getColumnInfos().addAll(listener.getInfo().getColumnInfos());
        }
    }
}
