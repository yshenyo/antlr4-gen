package com.thinice.bbc.core.parser.antlr.mysql;

import com.thinice.bbc.core.parser.antlr.parser.MysqlParserColumnInfo;
import com.thinice.bbc.core.parser.dto.info.ColumnInfo;
import lombok.Data;

@Data
public class MysqlColumnListener extends MySqlParserBaseListener {
    private MysqlParserColumnInfo info;

    public MysqlColumnListener(String content, int start, int stop) {
        info = new MysqlParserColumnInfo();
        info.setContent(content);
        info.setStart(start);
        info.setStop(stop);
    }


    @Override
    public void enterSelectStarElement(MySqlParser.SelectStarElementContext ctx) {
        if (info.isInDefinition() || info.getSubQueryStart() > -1) {
            return;
        }
        ColumnInfo columnInfo = new ColumnInfo();
        columnInfo.setOwner("");
        columnInfo.setColumn_type("");
        columnInfo.setWhere(false);
        columnInfo.setColumn("*");
        String table = "";
        String fullColumn = info.getContent(ctx);
        if (fullColumn.contains(".")) {
            table = fullColumn.substring(0, fullColumn.lastIndexOf(".")).toUpperCase();
            if (!table.equals("") && info.getTableNameMap().containsKey(table)) {
                table = info.getTableNameMap().get(table).getTable_name();
            } else if (!table.equals("") && info.getTableAliasMap().containsKey(table)) {
                table = info.getTableAliasMap().get(table).getTable_name();
            } else {
                table = "UNKNOWN";
            }
        } else {
            if (null != info.getTableInfos() && info.getTableInfos().size() == 1) {
                table = info.getTableInfos().get(0).getTable_name();
            } else {
                table = "UNKNOWN";
            }
        }
        columnInfo.setTable_name(table);
        columnInfo.setFull_name(table + "." + columnInfo.getColumn());
        info.getColumnInfos().add(columnInfo);
    }
    @Override
    public void enterFullColumnName(MySqlParser.FullColumnNameContext ctx) {
        if (info.isInDefinition() || info.getSubQueryStart() > -1) {
            return;
        }
        ColumnInfo columnInfo = new ColumnInfo();
        columnInfo.setOwner("");
        columnInfo.setColumn_type("");
        if (info.getWhereStart() > -1 && ctx.start.getStartIndex() >= info.getWhereStart() && ctx.stop.getStopIndex() <= info.getWhereStop()) {
            columnInfo.setWhere(true);
        } else {
            columnInfo.setWhere(false);
        }

        String table = "";
        String fullColumn = info.getContent(ctx);
        if (fullColumn.contains(".")) {
            table = fullColumn.substring(0, fullColumn.lastIndexOf(".")).toUpperCase();
            if (!table.equals("") && info.getTableNameMap().containsKey(table)) {
                table = info.getTableNameMap().get(table).getTable_name();
            } else if (!table.equals("") && info.getTableAliasMap().containsKey(table)) {
                table = info.getTableAliasMap().get(table).getTable_name();
            } else {
                table = "UNKNOWN";
            }
            columnInfo.setColumn(fullColumn.substring(fullColumn.lastIndexOf(".") + 1));
        } else {
            if (null != info.getTableInfos() && info.getTableInfos().size() == 1) {
                table = info.getTableInfos().get(0).getTable_name();
            } else {
                table = "UNKNOWN";
            }
            columnInfo.setColumn(info.getContent(ctx));
        }
        columnInfo.setTable_name(table);
        columnInfo.setFull_name(table + "." + columnInfo.getColumn());
        info.getColumnInfos().add(columnInfo);
    }
    @Override
    public void enterCreateDefinitions(MySqlParser.CreateDefinitionsContext ctx) {
        info.setInDefinition(true);
    }

    @Override
    public void exitCreateDefinitions(MySqlParser.CreateDefinitionsContext ctx) {
        info.setInDefinition(false);
    }
    @Override
    public void enterColumnDeclaration(MySqlParser.ColumnDeclarationContext ctx) {
        ColumnInfo columnInfo = new ColumnInfo();
        columnInfo.setOwner("");
        columnInfo.setWhere(false);
        columnInfo.setColumn(info.getContent(ctx.fullColumnName()));
        columnInfo.setColumn_type(info.getContent(ctx.columnDefinition().dataType()));
        if (null != info.getTableInfos() && info.getTableInfos().size() > 0) {
            columnInfo.setTable_name(info.getTableInfos().get(0).getTable_name());
        } else {
            columnInfo.setTable_name("UNKNOWN");
        }
        columnInfo.setFull_name(columnInfo.getTable_name() + "." + columnInfo.getColumn());
        info.getColumnInfos().add(columnInfo);
    }
    @Override
    public void enterSimpleSelect(MySqlParser.SimpleSelectContext ctx) {
        info.enterSubQuery(ctx);
    }
    @Override
    public void exitSimpleSelect(MySqlParser.SimpleSelectContext ctx) {
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterParenthesisSelect(MySqlParser.ParenthesisSelectContext ctx) {
        info.enterSubQuery(ctx);
    }
    @Override
    public void exitParenthesisSelect(MySqlParser.ParenthesisSelectContext ctx) {
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterUnionSelect(MySqlParser.UnionSelectContext ctx) {
        info.enterSubQuery(ctx);
    }
    @Override
    public void exitUnionSelect(MySqlParser.UnionSelectContext ctx) {
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterUnionParenthesisSelect(MySqlParser.UnionParenthesisSelectContext ctx) {
        info.enterSubQuery(ctx);
    }
    @Override
    public void exitUnionParenthesisSelect(MySqlParser.UnionParenthesisSelectContext ctx) {
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterWithLateralStatement(MySqlParser.WithLateralStatementContext ctx) {
        info.enterSubQuery(ctx);
    }
    @Override
    public void exitWithLateralStatement(MySqlParser.WithLateralStatementContext ctx) {
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterIsExpression(MySqlParser.IsExpressionContext ctx) {
        info.enterWhere(ctx);
    }
    @Override
    public void exitIsExpression(MySqlParser.IsExpressionContext ctx) {
        info.exitWhere(ctx);
    }

    @Override
    public void enterNotExpression(MySqlParser.NotExpressionContext ctx) {
        info.enterWhere(ctx);
    }
    @Override
    public void exitNotExpression(MySqlParser.NotExpressionContext ctx) {
        info.exitWhere(ctx);
    }

    @Override
    public void enterLogicalExpression(MySqlParser.LogicalExpressionContext ctx) {
        info.enterWhere(ctx);
    }
    @Override
    public void exitLogicalExpression(MySqlParser.LogicalExpressionContext ctx) {
        info.exitWhere(ctx);
    }

    @Override
    public void enterPredicateExpression(MySqlParser.PredicateExpressionContext ctx) {
        info.enterWhere(ctx);
    }
    @Override
    public void exitPredicateExpression(MySqlParser.PredicateExpressionContext ctx) {
        info.exitWhere(ctx);
    }





}
