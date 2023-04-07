package com.thinice.bbc.core.parser.antlr.hive.v4;

import com.thinice.bbc.core.parser.antlr.parser.ParserInfo;
import com.thinice.bbc.core.parser.dto.info.TableInfo;
import lombok.Data;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.tree.ParseTreeWalker;

import java.util.*;

@Data
public class HiveV4Listener extends HiveParserBaseListener {

    private ParserInfo info;

    public HiveV4Listener(String content) {
        info = new ParserInfo();
        info.setContent(content);
    }

    @Override
    public void enterFunction_(HiveParser.Function_Context ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }

    @Override
    public void enterSearchCondition(HiveParser.SearchConditionContext ctx) {
        HiveV4ConditionListener listener = new HiveV4ConditionListener(info.getContent());
        ParseTreeWalker.DEFAULT.walk(listener, ctx);
        info.setConditions(listener.getInfo().getConditions());
    }

    @Override
    public void enterTableName(HiveParser.TableNameContext ctx) {
        TableInfo tableInfo = new TableInfo();
        tableInfo.setTable_name(info.getContent(ctx));
        tableInfo.setType(info.getSqlType());
        if (ctx.start.getStartIndex() >= info.getSubQueryStart() && ctx.stop.getStopIndex() <= info.getSubQueryStop()) {
            tableInfo.setType("Select");
        }
        tableInfo.setName(tableInfo.getTable_name().toUpperCase());
        tableInfo.setStart(ctx.start.getStartIndex());
        tableInfo.setStop(ctx.stop.getStartIndex());
        HiveParser.Id_Context alias = ((HiveParser.TableSourceContext) ctx.getParent()).id_();
        if (null != alias) {
            tableInfo.setAlias(info.getContent(alias));
        }
        info.getTableInfos().add(tableInfo);
        info.getTables().add(tableInfo);
    }
    @Override
    public void enterSelectStatement(HiveParser.SelectStatementContext ctx) {
        info.subQueryCount++;
        if (info.subQueryCount > 0) {
            info.enterSubQuery(ctx);
        }
    }
    @Override
    public void exitSelectStatement(HiveParser.SelectStatementContext ctx) {
        info.exitSubQuery(ctx);
    }

    @Override
    public void enterExecStatement(HiveParser.ExecStatementContext ctx) {
        String stmt = ctx.getChild(0).getClass().getSimpleName();
        switch (stmt) {
            case "QueryStatementExpressionContext":
                HiveParser.WithClauseContext withClauseContext = ctx.queryStatementExpression().withClause();
                if (withClauseContext != null) {
                    stmt = "With";
                } else {
                    stmt = "Select";
                    info.subQueryCount--;
                }
                break;
            case "DdlStatementContext":
                HiveParser.DdlStatementContext ddlStatementContext = ctx.ddlStatement();
                stmt = ddlStatementContext.getChild(0).getClass().getSimpleName();
                switch (stmt) {
                    case "ResourcePlanDdlStatementsContext":
                        HiveParser.ResourcePlanDdlStatementsContext resourcePlanDdlStatementsContext = ddlStatementContext.resourcePlanDdlStatements();
                        stmt = resourcePlanDdlStatementsContext.getChild(0).getClass().getSimpleName();
                }
                break;
            case "SqlTransactionStatementContext":
                HiveParser.SqlTransactionStatementContext sqlTransactionStatementContext = ctx.sqlTransactionStatement();
                stmt = sqlTransactionStatementContext.getChild(0).getClass().getSimpleName();
                break;
        }
        stmt = stmt.replaceAll("Statement", "").replaceAll("Context", "");
        if (null != info.getStatement() && info.getStatement().equals("Explain")) {
            info.setStatement(stmt);
        }
        info.setSqlType(stmt);
    }
    @Override
    public void exitExecStatement(HiveParser.ExecStatementContext ctx) {
        getColumns(ctx);
    }
    @Override
    public void exitSubQueryExpression(HiveParser.SubQueryExpressionContext ctx) {
        getColumns(ctx);
        Iterator<TableInfo> it = info.getTables().iterator();
        while (it.hasNext()) {
            TableInfo tableInfo = it.next();
            if (tableInfo.getStart() >= ctx.start.getStartIndex() && tableInfo.getStop() <= ctx.stop.getStopIndex()) {
                it.remove();
            }
        }
    }

    @Override
    public void enterExplainStatement(HiveParser.ExplainStatementContext ctx) {
        info.setStatement("Explain");
    }
    @Override
    public void exitExplainStatement(HiveParser.ExplainStatementContext ctx) {
        if (null == ctx.execStatement()) {
            getColumns(ctx);
        }
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
        HiveV4ColumnListener listener = new HiveV4ColumnListener(info.getContent());
        listener.getInfo().setTableInfos(tables);
        listener.getInfo().setTableNameMap(tableNameMap);
        listener.getInfo().setTableAliasMap(tableAliasMap);
        ParseTreeWalker.DEFAULT.walk(listener, ctx);
        if (null != listener.getInfo().getColumnInfos()) {
            info.getColumnInfos().addAll(listener.getInfo().getColumnInfos());
        }
    }
}
