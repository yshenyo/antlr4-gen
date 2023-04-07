package com.thinice.bbc.core.parser.antlr.postgresql;

import com.thinice.bbc.core.parser.antlr.parser.ParserInfo;
import com.thinice.bbc.core.parser.dto.info.ConditionInfo;
import com.thinice.bbc.core.parser.dto.info.TableInfo;
import lombok.Data;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.tree.ParseTreeWalker;

import java.util.*;

@Data
public class PostgresqlListener extends PostgreSQLParserBaseListener {
    private ParserInfo info;
    private int inExprStart = -1;
    private int inExprStop = -1;

    public PostgresqlListener(String content) {
        info = new ParserInfo();
        info.setContent(content);
    }



    @Override
    public void enterFunc_application(PostgreSQLParser.Func_applicationContext ctx) {
        info.getFunctions().add(info.getContent(ctx));
    }
    @Override
    public void enterStmt(PostgreSQLParser.StmtContext ctx) {
        String statement = ctx.getChild(0).getClass().getSimpleName().replaceAll("stmt", "").replaceAll("Context", "");
        info.setStatement(statement);
        info.setSqlType(statement);
        if (info.getSqlType().equalsIgnoreCase("SELECT")) {
            info.subQueryCount--;
        }
    }
    @Override
    public void exitStmt(PostgreSQLParser.StmtContext ctx) {
        getColumns(ctx);
    }

    @Override
    public void enterSelect_clause(PostgreSQLParser.Select_clauseContext ctx) {
        info.subQueryCount++;
        if (info.subQueryCount > 0) {
            info.enterSubQuery(ctx);
        }
    }
    @Override
    public void exitSelect_clause(PostgreSQLParser.Select_clauseContext ctx) {
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
    public void enterRelation_expr(PostgreSQLParser.Relation_exprContext ctx) {
        TableInfo tableInfo = new TableInfo();
        tableInfo.setTable_name(info.getContent(ctx.qualified_name()));
        tableInfo.setType(info.getSqlType());
        if (ctx.start.getStartIndex() >= info.getSubQueryStart() && ctx.stop.getStopIndex() <= info.getSubQueryStop()) {
            tableInfo.setType("Select");
        }
        tableInfo.setName(tableInfo.getTable_name().toUpperCase());
        tableInfo.setStart(ctx.qualified_name().start.getStartIndex());
        tableInfo.setStop(ctx.qualified_name().stop.getStopIndex());
        tableInfo.setAlias("");
        ParserRuleContext parent = ctx.getParent();
        if (parent instanceof PostgreSQLParser.Table_refContext) {
            PostgreSQLParser.Table_refContext table_refContext = (PostgreSQLParser.Table_refContext) ctx.getParent();
            if (null != table_refContext.opt_alias_clause()) {
                tableInfo.setAlias(info.getContent(table_refContext.opt_alias_clause()));
            }
        } else if (parent instanceof PostgreSQLParser.Relation_expr_opt_aliasContext) {
            PostgreSQLParser.Relation_expr_opt_aliasContext relation_expr_opt_aliasContext = (PostgreSQLParser.Relation_expr_opt_aliasContext) ctx.getParent();
            if (null != relation_expr_opt_aliasContext.colid()) {
                tableInfo.setAlias(info.getContent(relation_expr_opt_aliasContext.colid()));
            }
        }
        info.getTableInfos().add(tableInfo);
        info.getTables().add(tableInfo);
    }

    @Override
    public void enterInsert_target(PostgreSQLParser.Insert_targetContext ctx) {
        TableInfo tableInfo = new TableInfo();
        tableInfo.setTable_name(info.getContent(ctx.qualified_name()));
        tableInfo.setType("Insert");
        tableInfo.setName(tableInfo.getTable_name().toUpperCase());
        tableInfo.setStart(ctx.qualified_name().start.getStartIndex());
        tableInfo.setStop(ctx.qualified_name().stop.getStopIndex());
        tableInfo.setAlias("");
        if (null != ctx.colid()) {
            tableInfo.setAlias(info.getContent(ctx.colid()));
        }
        info.getTableInfos().add(tableInfo);
        info.getTables().add(tableInfo);
    }

    @Override
    public void enterCreatestmt(PostgreSQLParser.CreatestmtContext ctx) {
        TableInfo tableInfo = new TableInfo();
        tableInfo.setTable_name(info.getContent(ctx.qualified_name(0)));
        tableInfo.setType("Create");
        tableInfo.setName(tableInfo.getTable_name().toUpperCase());
        tableInfo.setStart(ctx.qualified_name(0).start.getStartIndex());
        tableInfo.setStop(ctx.qualified_name(0).stop.getStopIndex());
        tableInfo.setAlias("");
        info.getTableInfos().add(tableInfo);
        info.getTables().add(tableInfo);
    }
    @Override
    public void enterA_expr_compare(PostgreSQLParser.A_expr_compareContext ctx) {
        if (inExprStart == -1 && inExprStop == -1) {
            List<PostgreSQLParser.A_expr_likeContext> likes = ctx.a_expr_like();
            ConditionInfo conditionInfo = new ConditionInfo();
            conditionInfo.setLeft(info.getContent(likes.get(0)));
            conditionInfo.setColumn_name(conditionInfo.getLeft());
            if (likes.size() > 1) {
                conditionInfo.setOperator(info.getContent(likes.get(0).stop.getStopIndex() + 1, likes.get(1).start.getStartIndex() - 1).trim().toUpperCase());
                conditionInfo.setRight(info.getContent(likes.get(1)));
            } else if (null != ctx.subquery_Op()) {
                conditionInfo.setOperator(info.getContent(ctx.subquery_Op()));
                conditionInfo.setRight(info.getContent(ctx.subquery_Op().stop.getStopIndex() + 1, ctx.stop.getStopIndex()));
            } else {
                return;
            }
            conditionInfo.setCondition(info.getContent(ctx));
            info.getConditions().add(conditionInfo);
            inExprStart = ctx.start.getStartIndex();
            inExprStop = ctx.stop.getStopIndex();
        }
    }
    @Override
    public void exitA_expr_compare(PostgreSQLParser.A_expr_compareContext ctx) {
        exitExpr(ctx);
    }

    @Override
    public void enterA_expr_is_not(PostgreSQLParser.A_expr_is_notContext ctx) {
        if (inExprStart == -1 && inExprStop == -1 && null != ctx.IS()) {
            ConditionInfo conditionInfo = new ConditionInfo();
            conditionInfo.setLeft(info.getContent(ctx.a_expr_compare()));
            conditionInfo.setColumn_name(conditionInfo.getLeft());
            String operator = "IS";
            String right = info.getContent(ctx.a_expr_compare().stop.getStopIndex() + 1, ctx.stop.getStopIndex()).replaceFirst("(?i)^IS", "").trim();
            if (null != ctx.NOT()) {
                operator = "IS NOT";
                right = right.replaceFirst("(?i)^NOT", "").trim();
            }
            conditionInfo.setRight(right);
            conditionInfo.setOperator(operator);
            conditionInfo.setCondition(info.getContent(ctx));
            info.getConditions().add(conditionInfo);
            inExprStart = ctx.start.getStartIndex();
            inExprStop = ctx.stop.getStopIndex();
        }
    }

    @Override
    public void exitA_expr_is_not(PostgreSQLParser.A_expr_is_notContext ctx) {
        exitExpr(ctx);
    }

    @Override
    public void enterA_expr_isnull(PostgreSQLParser.A_expr_isnullContext ctx) {
        if (inExprStart == -1 && inExprStop == -1 && (null != ctx.ISNULL() || null != ctx.NOTNULL())) {
            ConditionInfo conditionInfo = new ConditionInfo();
            conditionInfo.setLeft(info.getContent(ctx.a_expr_is_not()));
            conditionInfo.setColumn_name(conditionInfo.getLeft());
            String operator = "IS NULL";
            if (null != ctx.NOTNULL()) {
                operator = "NOT NULL";
            }
            conditionInfo.setOperator(operator);
            conditionInfo.setCondition(info.getContent(ctx));
            info.getConditions().add(conditionInfo);
            inExprStart = ctx.start.getStartIndex();
            inExprStop = ctx.stop.getStopIndex();
        }
    }
    @Override
    public void exitA_expr_isnull(PostgreSQLParser.A_expr_isnullContext ctx) {
        exitExpr(ctx);
    }

    @Override
    public void enterA_expr_in(PostgreSQLParser.A_expr_inContext ctx) {
        if (inExprStart == -1 && inExprStop == -1 && null != ctx.IN_P()) {
            ConditionInfo conditionInfo = new ConditionInfo();
            conditionInfo.setLeft(info.getContent(ctx.a_expr_unary_not()));
            conditionInfo.setColumn_name(conditionInfo.getLeft());
            String operator = "IN";
            if (null != ctx.NOT()) {
                operator = "NOT IN";
            }
            conditionInfo.setOperator(operator);
            conditionInfo.setRight(info.getContent(ctx.in_expr()));
            conditionInfo.setCondition(info.getContent(ctx));
            info.getConditions().add(conditionInfo);
            inExprStart = ctx.start.getStartIndex();
            inExprStop = ctx.stop.getStopIndex();
        }
    }
    @Override
    public void exitA_expr_in(PostgreSQLParser.A_expr_inContext ctx) {
        exitExpr(ctx);
    }

    @Override
    public void enterA_expr_between(PostgreSQLParser.A_expr_betweenContext ctx) {
        if (ctx.a_expr_in().size() > 1) {
            ConditionInfo conditionInfo = new ConditionInfo();
            conditionInfo.setLeft(info.getContent(ctx.a_expr_in(0)));
            conditionInfo.setColumn_name(conditionInfo.getLeft());
            String operator = "BETWEEN";
            if (null != ctx.NOT()) {
                operator = "NOT BETWEEN";
            }
            if (null != ctx.SYMMETRIC()) {
                operator += " SYMMETRIC";
            }
            conditionInfo.setOperator(operator);
            conditionInfo.setRight(info.getContent(ctx.a_expr_in(1).start.getStartIndex(), ctx.a_expr_in(2).stop.getStopIndex()));
            conditionInfo.setCondition(info.getContent(ctx));
            info.getConditions().add(conditionInfo);
            inExprStart = ctx.start.getStartIndex();
            inExprStop = ctx.stop.getStopIndex();
        }
    }
    @Override
    public void exitA_expr_between(PostgreSQLParser.A_expr_betweenContext ctx) {
        exitExpr(ctx);
    }
    @Override
    public void enterA_expr_like(PostgreSQLParser.A_expr_likeContext ctx) {
        if (inExprStart == -1 && inExprStop == -1 && ctx.a_expr_qual_op().size() > 1) {
            ConditionInfo conditionInfo = new ConditionInfo();
            conditionInfo.setLeft(info.getContent(ctx.a_expr_qual_op(0)));
            conditionInfo.setColumn_name(conditionInfo.getLeft());
            String operator = "";
            if (null != ctx.NOT()) {
                operator += "NOT ";
            }
            if (null != ctx.LIKE()) {
                operator += "LIKE";
            } else if (null != ctx.ILIKE()) {
                operator += "ILIKE";
            } else if (null != ctx.SIMILAR()) {
                operator += "SIMILAR TO";
            }
//            } else if (null != ctx.BETWEEN()) {
//                operator += "BETWEEN";
//            }
            conditionInfo.setOperator(operator);
            conditionInfo.setRight(info.getContent(ctx.a_expr_qual_op(1).start.getStartIndex(), ctx.stop.getStopIndex()));
            conditionInfo.setCondition(info.getContent(ctx));
            info.getConditions().add(conditionInfo);
            inExprStart = ctx.start.getStartIndex();
            inExprStop = ctx.stop.getStopIndex();
        }
    }
    @Override
    public void exitA_expr_like(PostgreSQLParser.A_expr_likeContext ctx) {
        exitExpr(ctx);
    }


    @Override
    public void enterExplainablestmt(PostgreSQLParser.ExplainablestmtContext ctx) {
        String sqlType = "Explain";
        if (null != ctx.createasstmt() || null != ctx.creatematviewstmt()) {
            sqlType = "Create";
        } else if (null != ctx.deletestmt()) {
            sqlType = "Delete";
        } else if (null != ctx.selectstmt()) {
            sqlType = "Select";
            info.subQueryCount--;
        } else if (null != ctx.insertstmt()) {
            sqlType = "Insert";
        } else if (null != ctx.updatestmt()) {
            sqlType = "Update";
        } else if (null != ctx.declarecursorstmt()) {
            sqlType = "Declare";
        } else if (null != ctx.refreshmatviewstmt()) {
            sqlType = "Refresh";
        } else if (null != ctx.executestmt()) {
            sqlType = "Execute";
        }
        info.setSqlType(sqlType);
    }


    private void exitExpr(ParserRuleContext ctx) {
        if (inExprStart != -1 && inExprStop != -1 && inExprStart == ctx.start.getStartIndex() && inExprStop == ctx.stop.getStopIndex()) {
            inExprStart = -1;
            inExprStop = -1;
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
        PostgresqlColumnListener listener = new PostgresqlColumnListener(info.getContent());
        listener.getInfo().setTableInfos(tables);
        listener.getInfo().setTableNameMap(tableNameMap);
        listener.getInfo().setTableAliasMap(tableAliasMap);
        ParseTreeWalker.DEFAULT.walk(listener, ctx);
        if (null != listener.getInfo().getColumnInfos()) {
            info.getColumnInfos().addAll(listener.getInfo().getColumnInfos());
        }
    }
}
