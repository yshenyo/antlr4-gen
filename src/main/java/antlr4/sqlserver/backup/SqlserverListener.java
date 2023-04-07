package com.thinice.bbc.core.parser.antlr.sqlserver;

import com.thinice.bbc.core.parser.antlr.oracle.OracleColumnListener;
import com.thinice.bbc.core.parser.antlr.oracle.PlSqlParser;
import com.thinice.bbc.core.parser.antlr.oracle.PlSqlParserBaseListener;
import com.thinice.bbc.core.parser.antlr.parser.ParserInfo;
import com.thinice.bbc.core.parser.dto.info.ConditionInfo;
import com.thinice.bbc.core.parser.dto.info.TableInfo;
import com.thinice.bbc.core.util.Tools;
import lombok.Data;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.tree.ParseTreeWalker;

import java.util.*;

@Data
public class SqlserverListener extends TSqlParserBaseListener {

    private ParserInfo info;
    private int statementCount = 0;

    public SqlserverListener(String content) {
        info = new ParserInfo();
        info.setContent(content);
    }
    //statement
    @Override
    public void enterBatch_level_statement(TSqlParser.Batch_level_statementContext ctx) {
        String operator = "Create";
        if (info.getContent(ctx).toUpperCase().startsWith("ALTER")) {
            operator = "Alter";
        }
        String statement;
        if (null != ctx.create_or_alter_function()) {
            statement = operator + "_Function";
        } else if (null != ctx.create_or_alter_procedure()) {
            statement = operator + "_Procedure";
        } else if (null != ctx.create_or_alter_trigger()) {
            statement = operator + "_Trigger";
        } else {
            statement = "Create_View";
        }
        info.setStatement(statement);
        info.setSqlType(statement);
    }
    @Override
    public void exitBatch_level_statement(TSqlParser.Batch_level_statementContext ctx) {
        getColumns(ctx);
    }

    @Override
    public void enterSql_clauses(TSqlParser.Sql_clausesContext ctx) {
        if (info.getContent(ctx).trim().equalsIgnoreCase(";")) {
            return;
        }
        statementCount++;
        String statement = ctx.getChild(0).getChild(0).getClass().getSimpleName()
                .replaceAll("_statement", "")
                .replaceAll("_standalone", "")
                .replaceAll("Context", "");
        if (statement.equalsIgnoreCase("TRANSACTION")) {
            statement = getInfo().getContent(ctx).trim().split(" ")[0] + "_Transaction";
        } else if (statement.equalsIgnoreCase("CONVERSATION")) {
            statement = ctx.getChild(0).getChild(0).getChild(0).getClass().getSimpleName()
                    .replaceAll("Context", "");
        } else if (null != ctx.dbcc_clause()) {
            statement = "Dbcc";
        }
        if (info.getStatement().equalsIgnoreCase("UNKNOWN")) {
            info.setStatement(statement);
        }
        info.setSqlType(statement);
    }
    @Override
    public void exitSql_clauses(TSqlParser.Sql_clausesContext ctx) {
        statementCount--;
        if (statementCount <= 1) {
            getColumns(ctx);
            info.setSqlType(info.getStatement());
        }
    }

    //tables
    @Override public void enterFull_table_name(TSqlParser.Full_table_nameContext ctx) {
        TableInfo tableInfo = new TableInfo();
        tableInfo.setTable_name(info.getContent(ctx));
        tableInfo.setType(Tools.LineToHump(info.getSqlType()));
        if (ctx.start.getStartIndex() >= info.getSubQueryStart() && ctx.stop.getStopIndex() <= info.getSubQueryStop()) {
            tableInfo.setType("Select");
        }
        tableInfo.setName(tableInfo.getTable_name().toUpperCase());
        TSqlParser.As_table_aliasContext alias = ctx.getParent().getRuleContext(TSqlParser.As_table_aliasContext.class, 0);
        tableInfo.setAlias("");
        if (null != alias) {
            tableInfo.setAlias(info.getContent(alias));
        }
        tableInfo.setStart(ctx.start.getStartIndex());
        tableInfo.setStop(ctx.stop.getStartIndex());
        info.getTableInfos().add(tableInfo);
        info.getTables().add(tableInfo);
    }
    @Override public void enterTable_name(TSqlParser.Table_nameContext ctx) {
        TableInfo tableInfo = new TableInfo();
        tableInfo.setTable_name(info.getContent(ctx));
        tableInfo.setType(Tools.LineToHump(info.getSqlType()));
        if (ctx.start.getStartIndex() >= info.getSubQueryStart() && ctx.stop.getStopIndex() <= info.getSubQueryStop()) {
            tableInfo.setType("Select");
        }
        tableInfo.setName(tableInfo.getTable_name().toUpperCase());
        tableInfo.setAlias("");
        tableInfo.setStart(ctx.start.getStartIndex());
        tableInfo.setStop(ctx.stop.getStartIndex());
        info.getTableInfos().add(tableInfo);
        info.getTables().add(tableInfo);
    }
    @Override
    public void enterSubquery(TSqlParser.SubqueryContext ctx) {
        info.enterSubQuery(ctx);
    }
    @Override
    public void exitSubquery(TSqlParser.SubqueryContext ctx) {
        getColumns(ctx.select_statement());
        Iterator<TableInfo> it = info.getTables().iterator();
        while (it.hasNext()) {
            TableInfo tableInfo = it.next();
            if (tableInfo.getStart() >= ctx.start.getStartIndex() && tableInfo.getStop() <= ctx.stop.getStopIndex()) {
                it.remove();
            }
        }
        info.exitSubQuery(ctx);
    }
    //conditions
    @Override
    public void enterPredicate(TSqlParser.PredicateContext ctx) {
        if (info.getSubQueryStart() > -1) {
            return;
        }
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setCondition(info.getContent(ctx));
        if (null != ctx.EXISTS()) {
            conditionInfo.setOperator("EXISTS");
            conditionInfo.setLeft(info.getContent(ctx));
            conditionInfo.setRight(info.getContent(ctx.subquery()));
        } else if (null != ctx.freetext_predicate()) {
            conditionInfo.setLeft(info.getContent(ctx));
        } else if (null != ctx.comparison_operator()) {
            conditionInfo.setLeft(info.getContent(ctx.expression(0)));
            conditionInfo.setOperator(info.getContent(ctx.comparison_operator()));
            switch (ctx.expression().size()) {
                case 1:
                    conditionInfo.setRight(info.getContent(ctx.subquery()));
                    break;
                case 2:
                    conditionInfo.setRight(info.getContent(ctx.expression(1)));
                    break;
            }
        } else if (null != ctx.BETWEEN()) {
            conditionInfo.setOperator("BETWEEN");
            conditionInfo.setLeft(info.getContent(ctx.expression(0)));
            conditionInfo.setRight(info.getContent(ctx.expression(1).start.getStartIndex(), ctx.stop.getStopIndex()));
        } else if (null != ctx.MULT_ASSIGN()) {
            conditionInfo.setOperator("*=");
            conditionInfo.setLeft(info.getContent(ctx.expression(0)));
            conditionInfo.setRight(info.getContent(ctx.expression(1)));
        } else if (null != ctx.IN()) {
            conditionInfo.setOperator("IN");
            conditionInfo.setLeft(info.getContent(ctx.expression(0)));
            if (null != ctx.subquery()) {
                conditionInfo.setRight(info.getContent(ctx.subquery()));
            } else {
                conditionInfo.setRight(info.getContent(ctx.expression_list_()));
            }
        } else if (null != ctx.LIKE()) {
            conditionInfo.setOperator("LIKE");
            conditionInfo.setLeft(info.getContent(ctx.expression(0)));
            conditionInfo.setRight(info.getContent(ctx.expression(1).start.getStartIndex(), ctx.stop.getStopIndex()));
        } else if (null != ctx.IS()) {
            conditionInfo.setOperator("IS");
            conditionInfo.setLeft(info.getContent(ctx.expression(0)));
            conditionInfo.setRight(info.getContent(ctx.null_notnull()));
        }
        if (null != ctx.NOT(0)) {
            conditionInfo.setOperator("NOT " + conditionInfo.getOperator());
        }
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        info.getConditions().add(conditionInfo);
    }
    //functions
//    @Override
//    public void enterExecute_body(TSqlParser.Execute_bodyContext ctx) {
//        info.getFunctions().add(info.getContent(ctx));
//    }
//    @Override
//    public void enterExecute_body_batch(TSqlParser.Execute_body_batchContext ctx) {
//        info.getFunctions().add(info.getContent(ctx));
//    }
    @Override
    public void enterRanking_windowed_function(TSqlParser.Ranking_windowed_functionContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }
    @Override
    public void enterAggregate_windowed_function(TSqlParser.Aggregate_windowed_functionContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }
    @Override
    public void enterAnalytic_windowed_function(TSqlParser.Analytic_windowed_functionContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }
    @Override
    public void enterBUILT_IN_FUNC(TSqlParser.BUILT_IN_FUNCContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }
    @Override
    public void enterSCALAR_FUNCTION(TSqlParser.SCALAR_FUNCTIONContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }
    @Override
    public void enterFREE_TEXT(TSqlParser.FREE_TEXTContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }
    @Override
    public void enterPARTITION_FUNC(TSqlParser.PARTITION_FUNCContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }
    @Override
    public void enterHIERARCHYID_METHOD(TSqlParser.HIERARCHYID_METHODContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
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
        SqlserverColumnListener listener = new SqlserverColumnListener(info.getContent());
        listener.getInfo().setTableInfos(tables);
        listener.getInfo().setTableNameMap(tableNameMap);
        listener.getInfo().setTableAliasMap(tableAliasMap);
        ParseTreeWalker.DEFAULT.walk(listener, ctx);
        if (null != listener.getInfo().getColumnInfos()) {
            info.getColumnInfos().addAll(listener.getInfo().getColumnInfos());
        }
    }
}
