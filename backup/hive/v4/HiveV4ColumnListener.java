package com.thinice.bbc.core.parser.antlr.hive.v4;

import com.thinice.bbc.core.parser.antlr.parser.ParserColumnInfo;
import com.thinice.bbc.core.parser.dto.info.ColumnInfo;
import lombok.Data;

@Data
public class HiveV4ColumnListener extends HiveParserBaseListener {
    private ParserColumnInfo info;

    public HiveV4ColumnListener(String content) {
        info = new ParserColumnInfo();
        info.setContent(content);
    }
    @Override
    public void enterSubQueryExpression(HiveParser.SubQueryExpressionContext ctx) {
        info.enterSubQuery(ctx);
    }
    @Override
    public void exitSubQueryExpression(HiveParser.SubQueryExpressionContext ctx) {
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterTableAllColumns(HiveParser.TableAllColumnsContext ctx) {
        if (info.getSubQueryStart() > -1) {
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
    public void enterPrecedenceFieldExpression(HiveParser.PrecedenceFieldExpressionContext ctx) {
        if (info.getSubQueryStart() > -1) {
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
    public void enterColumnNameTypeConstraint(HiveParser.ColumnNameTypeConstraintContext ctx) {
        if (info.getSubQueryStart() > -1) {
            return;
        }
        ColumnInfo columnInfo = new ColumnInfo();
        columnInfo.setOwner("");
        columnInfo.setColumn_type("");
        columnInfo.setWhere(false);
        columnInfo.setColumn(info.getContent(ctx.id_()));
        if (null != ctx.colType()) {
            columnInfo.setColumn_type(info.getContent(ctx.colType()));
        } else {
            columnInfo.setColumn_type("");
        }
        if (null != info.getTableInfos() && info.getTableInfos().size() > 0) {
            columnInfo.setTable_name(info.getTableInfos().get(0).getTable_name());
        } else {
            columnInfo.setTable_name("UNKNOWN");
        }
        columnInfo.setFull_name(columnInfo.getTable_name() + "." + columnInfo.getColumn());
        info.getColumnInfos().add(columnInfo);
    }

    @Override
    public void enterWhereClause(HiveParser.WhereClauseContext ctx) {
        info.enterWhere(ctx);
    }

    @Override
    public void exitWhereClause(HiveParser.WhereClauseContext ctx) {
        info.exitWhere(ctx);
    }
}
