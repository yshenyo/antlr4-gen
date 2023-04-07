package com.thinice.bbc.core.parser.antlr.sqlserver;

import com.thinice.bbc.core.parser.antlr.oracle.PlSqlParser;
import com.thinice.bbc.core.parser.antlr.oracle.PlSqlParserBaseListener;
import com.thinice.bbc.core.parser.antlr.parser.ParserColumnInfo;
import com.thinice.bbc.core.parser.dto.info.ColumnInfo;
import lombok.Data;

import java.util.List;

@Data
public class SqlserverColumnListener extends TSqlParserBaseListener {

    private ParserColumnInfo info;

    public SqlserverColumnListener(String content) {
        info = new ParserColumnInfo();
        info.setContent(content);
    }

    //columns
    @Override
    public void enterColumn_definition(TSqlParser.Column_definitionContext ctx) {
        if (info.getSubQueryStart() > -1) {
            return;
        }
        ColumnInfo columnInfo = new ColumnInfo();
        columnInfo.setOwner("");
        columnInfo.setWhere(false);
        columnInfo.setColumn(info.getContent(ctx.id_()));
        if (null != ctx.data_type()) {
            columnInfo.setColumn_type(info.getContent(ctx.data_type()));
        } else {
            columnInfo.setColumn_type("");
        }
        String table;
        if (null != info.getTableInfos() && info.getTableInfos().size() == 1) {
            table = info.getTableInfos().get(0).getTable_name();
        } else {
            table = "UNKNOWN";
        }
        columnInfo.setTable_name(table);
        columnInfo.setFull_name(table + "." + columnInfo.getColumn());
        info.getColumnInfos().add(columnInfo);
    }
    @Override
    public void enterFull_column_name(TSqlParser.Full_column_nameContext ctx) {
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
        columnInfo.setColumn(info.getContent(ctx.id_()));
        String table;
        if (null != ctx.full_table_name()) {
            table = info.getContent(ctx.full_table_name()).toUpperCase();
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
    public void enterInsert_column_name_list(TSqlParser.Insert_column_name_listContext ctx) {
        if (info.getSubQueryStart() > -1) {
            return;
        }
        String table;
        if (null != info.getTableInfos() && info.getTableInfos().size() == 1) {
            table = info.getTableInfos().get(0).getTable_name();
        } else {
            table = "UNKNOWN";
        }
        ctx.insert_column_id().forEach(col -> {
            ColumnInfo columnInfo = new ColumnInfo();
            columnInfo.setOwner("");
            columnInfo.setColumn_type("");
            columnInfo.setWhere(false);
            String[] cs = info.getContent(col).split("\\.");
            columnInfo.setColumn(cs[cs.length - 1]);
            columnInfo.setTable_name(table);
            columnInfo.setFull_name(table + "." + columnInfo.getColumn());
            info.getColumnInfos().add(columnInfo);
        });
    }
    @Override
    public void enterColumn_name_list_with_order(TSqlParser.Column_name_list_with_orderContext ctx) {
        if (info.getSubQueryStart() > -1) {
            return;
        }
        String table;
        if (null != info.getTableInfos() && info.getTableInfos().size() == 1) {
            table = info.getTableInfos().get(0).getTable_name();
        } else {
            table = "UNKNOWN";
        }
        ctx.id_().forEach(col -> {
            ColumnInfo columnInfo = new ColumnInfo();
            columnInfo.setOwner("");
            columnInfo.setColumn_type("");
            columnInfo.setWhere(false);
            columnInfo.setColumn(info.getContent(col));
            columnInfo.setTable_name(table);
            columnInfo.setFull_name(table + "." + columnInfo.getColumn());
            info.getColumnInfos().add(columnInfo);
        });
    }
    @Override
    public void enterColumn_name_list(TSqlParser.Column_name_listContext ctx) {
        if (info.getSubQueryStart() > -1) {
            return;
        }
        String table;
        if (null != info.getTableInfos() && info.getTableInfos().size() == 1) {
            table = info.getTableInfos().get(0).getTable_name();
        } else {
            table = "UNKNOWN";
        }
        ctx.id_().forEach(col -> {
            ColumnInfo columnInfo = new ColumnInfo();
            columnInfo.setOwner("");
            columnInfo.setColumn_type("");
            columnInfo.setWhere(false);
            columnInfo.setColumn(info.getContent(col));
            columnInfo.setTable_name(table);
            columnInfo.setFull_name(table + "." + columnInfo.getColumn());
            info.getColumnInfos().add(columnInfo);
        });
    }
    @Override
    public void enterSubquery(TSqlParser.SubqueryContext ctx) {
        info.enterSubQuery(ctx);
    }
    @Override
    public void exitSubquery(TSqlParser.SubqueryContext ctx) {
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterSearch_condition(TSqlParser.Search_conditionContext ctx) {
        info.enterWhere(ctx);
    }
    @Override
    public void exitSearch_condition(TSqlParser.Search_conditionContext ctx) {
        info.exitWhere(ctx);
    }
}
