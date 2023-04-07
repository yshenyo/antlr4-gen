package com.thinice.bbc.core.parser.antlr.oracle;

import com.thinice.bbc.core.parser.antlr.parser.ParserInfo;
import com.thinice.bbc.core.parser.dto.info.ConditionInfo;
import com.thinice.bbc.core.parser.dto.info.TableInfo;
import com.thinice.bbc.core.util.Tools;
import lombok.Data;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.tree.ParseTreeWalker;

import java.util.*;

@Data
public class OracleListener extends PlSqlParserBaseListener {

    private ParserInfo info;

    public OracleListener(String content) {
        info = new ParserInfo();
        info.setContent(content);
    }

    @Override
    public void enterCall_statement(PlSqlParser.Call_statementContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }

    @Override
    public void enterSubquery(PlSqlParser.SubqueryContext ctx) {
        info.subQueryCount++;
        if (info.subQueryCount > 0) {
            info.enterSubQuery(ctx);
        }
    }

    @Override
    public void exitSubquery(PlSqlParser.SubqueryContext ctx) {
        if (info.subQueryCount > 0) {
            getColumns(ctx.subquery_basic_elements());
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
    public void enterTableview_name(PlSqlParser.Tableview_nameContext ctx) {
        if (info.getSubQueryStart() == -1 && !info.getSqlType().equalsIgnoreCase("SELECT")) {
            TableInfo tableInfo = new TableInfo();
            tableInfo.setTable_name(info.getContent(ctx));
            tableInfo.setType(Tools.LineToHump(info.getSqlType()));
            if (ctx.start.getStartIndex() >= info.getSubQueryStart() && ctx.stop.getStopIndex() <= info.getSubQueryStop()) {
                tableInfo.setType("Select");
            }
            tableInfo.setName(tableInfo.getTable_name().toUpperCase());
            tableInfo.setAlias("");
            tableInfo.setStart(ctx.start.getStartIndex());
            tableInfo.setStop(ctx.stop.getStopIndex());
            info.getTableInfos().add(tableInfo);
            info.getTables().add(tableInfo);
        }
    }
    @Override
    public void enterTable_name(PlSqlParser.Table_nameContext ctx) {
        if (info.getSubQueryStart() == -1 && !info.getSqlType().equalsIgnoreCase("SELECT")) {
            TableInfo tableInfo = new TableInfo();
            tableInfo.setTable_name(info.getContent(ctx));
            tableInfo.setType(Tools.LineToHump(info.getSqlType()));
            if (ctx.start.getStartIndex() >= info.getSubQueryStart() && ctx.stop.getStopIndex() <= info.getSubQueryStop()) {
                tableInfo.setType("Select");
            }
            tableInfo.setName(tableInfo.getTable_name().toUpperCase());
            tableInfo.setAlias("");
            tableInfo.setStart(ctx.start.getStartIndex());
            tableInfo.setStop(ctx.stop.getStopIndex());
            info.getTableInfos().add(tableInfo);
            info.getTables().add(tableInfo);
        }
    }

    @Override
    public void enterTable_ref_aux(PlSqlParser.Table_ref_auxContext ctx) {
        TableInfo tableInfo = new TableInfo();
        tableInfo.setTable_name(info.getContent(ctx.table_ref_aux_internal()));
        tableInfo.setType(Tools.LineToHump(info.getSqlType()));
        if (ctx.start.getStartIndex() >= info.getSubQueryStart() && ctx.stop.getStopIndex() <= info.getSubQueryStop()) {
            tableInfo.setType("Select");
        }
        tableInfo.setName(tableInfo.getTable_name().toUpperCase());
        if (null != ctx.table_alias()) {
            tableInfo.setAlias(info.getContent(ctx.table_alias()).toUpperCase());
        } else {
            tableInfo.setAlias("");
        }
        tableInfo.setStart(ctx.table_ref_aux_internal().start.getStartIndex());
        tableInfo.setStop(ctx.table_ref_aux_internal().stop.getStopIndex());
        info.getTableInfos().add(tableInfo);
        info.getTables().add(tableInfo);
    }


    @Override
    public void enterRelational_expression(PlSqlParser.Relational_expressionContext ctx) {
        if (info.getSubQueryStart() > -1) {
            return;
        }
        PlSqlParser.Relational_expressionContext left = ctx.relational_expression(0);
        if (null != left) {
            PlSqlParser.Relational_operatorContext operator = ctx.relational_operator();
            ConditionInfo conditionInfo = new ConditionInfo();
            conditionInfo.setCondition(info.getContent(ctx));
            conditionInfo.setOperator(info.getContent(operator).toUpperCase());
            conditionInfo.setColumn_name(info.getContent(left));
            conditionInfo.setLeft(info.getContent(left));
            conditionInfo.setRight(info.getContent(ctx.relational_expression(1)));
            info.getConditions().add(conditionInfo);
        }
    }

    @Override
    public void enterCompound_expression(PlSqlParser.Compound_expressionContext ctx) {
        if (info.getSubQueryStart() > -1) {
            return;
        }
        if (ctx.getChildCount() > 1) {
            ConditionInfo conditionInfo = new ConditionInfo();
            PlSqlParser.ConcatenationContext left = ctx.concatenation(0);
            conditionInfo.setCondition(info.getContent(ctx));
            conditionInfo.setColumn_name(info.getContent(left));
            conditionInfo.setLeft(info.getContent(left));
            conditionInfo.setOperator(getCompoundOperator(ctx));
            ParserRuleContext right = ctx.in_elements();
            if (null != right) {
                conditionInfo.setRight(info.getContent(right.start.getStartIndex(), ctx.stop.getStopIndex()));
            } else {
                right = ctx.between_elements();
                if (null != right) {
                    conditionInfo.setRight(info.getContent(right.start.getStartIndex(), ctx.stop.getStopIndex()));
                } else {
                    right = ctx.concatenation(1);
                    if (null != right) {
                        conditionInfo.setRight(info.getContent(right.start.getStartIndex(), ctx.stop.getStopIndex()));
                    }
                }
            }
            info.getConditions().add(conditionInfo);
        }
    }


    @Override
    public void enterUnit_statement(PlSqlParser.Unit_statementContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("_statement", "");
        info.setSqlType(statement);
        info.setStatement(statement);
    }
    @Override
    public void exitUnit_statement(PlSqlParser.Unit_statementContext ctx) {
        if (null != ctx.data_manipulation_language_statements() && null != ctx.data_manipulation_language_statements().select_statement()) {
            getColumns(ctx.data_manipulation_language_statements().select_statement().select_only_statement().subquery().subquery_basic_elements());
        } else {
            getColumns(ctx);
        }
    }

    @Override
    public void enterTransaction_control_statements(PlSqlParser.Transaction_control_statementsContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("_statement", "");
        info.setSqlType(statement);
        info.setStatement(statement);
    }
//    @Override
//    public void exitTransaction_control_statements(PlSqlParser.Transaction_control_statementsContext ctx) {
//        getColumns(ctx);
//    }

    @Override
    public void enterData_manipulation_language_statements(PlSqlParser.Data_manipulation_language_statementsContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("Context", "").replaceAll("_statement", "");
        info.setSqlType(statement);
        info.setStatement(statement);
        if (null != ctx.select_statement()) {
            info.subQueryCount--;
        }
    }
//    @Override
//    public void exitData_manipulation_language_statements(PlSqlParser.Data_manipulation_language_statementsContext ctx) {
//        if (null == ctx.select_statement()) {
//            getColumns(ctx);
//        }
//    }


    @Override
    public void enterExplain_statement(PlSqlParser.Explain_statementContext ctx) {
        String sqlType = "Explain";
        if (null != ctx.select_statement()) {
            sqlType = "Select";
            info.subQueryCount--;
        } else if (null != ctx.delete_statement()) {
            sqlType = "Delete";
        } else if (null != ctx.insert_statement()) {
            sqlType = "Insert";
        } else if (null != ctx.update_statement()) {
            sqlType = "Update";
        } else if (null != ctx.merge_statement()) {
            sqlType = "Merge";
        }
        info.setSqlType(sqlType);
    }

    private String getCompoundOperator(PlSqlParser.Compound_expressionContext ctx) {
        String operator = "";
        if (null != ctx.NOT()) {
            operator += "NOT ";
        }
        if (null != ctx.IN()) {
            operator += "IN";
        } else if (null != ctx.BETWEEN()) {
            operator += "BETWEEN";
        } else if (null != ctx.LIKE()) {
            operator += "LIKE";
        } else if (null != ctx.LIKEC()) {
            operator += "LIKEC";
        } else if (null != ctx.LIKE2()) {
            operator += "LIKE2";
        } else if (null != ctx.LIKE4()) {
            operator += "LIKE4";
        }
        return operator;
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
        OracleColumnListener listener = new OracleColumnListener(info.getContent());
        listener.getInfo().setTableInfos(tables);
        listener.getInfo().setTableNameMap(tableNameMap);
        listener.getInfo().setTableAliasMap(tableAliasMap);
        ParseTreeWalker.DEFAULT.walk(listener, ctx);
        if (null != listener.getInfo().getColumnInfos()) {
            info.getColumnInfos().addAll(listener.getInfo().getColumnInfos());
        }
    }
}
