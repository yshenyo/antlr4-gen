package com.thinice.bbc.core.parser.antlr.postgresql;

import com.thinice.bbc.core.parser.antlr.parser.ParserColumnInfo;
import com.thinice.bbc.core.parser.dto.info.ColumnInfo;
import lombok.Data;

@Data
public class PostgresqlColumnListener extends PostgreSQLParserBaseListener {
    private ParserColumnInfo info;

    public PostgresqlColumnListener(String content) {
        info = new ParserColumnInfo();
        info.setContent(content);
    }


    @Override
    public void enterColumnDef(PostgreSQLParser.ColumnDefContext ctx) {
        ColumnInfo columnInfo = new ColumnInfo();
        columnInfo.setOwner("");
        columnInfo.setWhere(false);
        columnInfo.setColumn(info.getContent(ctx.colid()));
        if (null != ctx.typename()) {
            columnInfo.setColumn_type(info.getContent(ctx.typename()));
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
    public void enterColumnref(PostgreSQLParser.ColumnrefContext ctx) {
        if (info.getSubQueryStart() > -1) {
            return;
        }
        ColumnInfo columnInfo = new ColumnInfo();
        columnInfo.setOwner("");
        columnInfo.setColumn_type("");
        if (info.getSubQueryStart() > -1 && ctx.start.getStartIndex() >= info.getSubQueryStart() && ctx.stop.getStopIndex() <= info.getSubQueryStop()) {
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
    public void enterColumnlist(PostgreSQLParser.ColumnlistContext ctx) {
        ctx.columnElem().forEach(col -> {
            ColumnInfo columnInfo = new ColumnInfo();
            columnInfo.setOwner("");
            columnInfo.setColumn_type("");
            columnInfo.setWhere(false);
            columnInfo.setColumn(info.getContent(col));
            if (null != info.getTableInfos() && info.getTableInfos().size() > 0) {
                columnInfo.setTable_name(info.getTableInfos().get(0).getTable_name());
            } else {
                columnInfo.setTable_name("UNKNOWN");
            }
            columnInfo.setFull_name(columnInfo.getTable_name() + "." + columnInfo.getColumn());
            info.getColumnInfos().add(columnInfo);
        });
    }
    @Override
    public void enterInsert_column_list(PostgreSQLParser.Insert_column_listContext ctx) {
        ctx.insert_column_item().forEach(col -> {
            ColumnInfo columnInfo = new ColumnInfo();
            columnInfo.setOwner("");
            columnInfo.setColumn_type("");
            columnInfo.setWhere(false);
            columnInfo.setColumn(info.getContent(col.colid()));
            if (null != info.getTableInfos() && info.getTableInfos().size() > 0) {
                columnInfo.setTable_name(info.getTableInfos().get(0).getTable_name());
            } else {
                columnInfo.setTable_name("UNKNOWN");
            }
            columnInfo.setFull_name(columnInfo.getTable_name() + "." + columnInfo.getColumn());
            info.getColumnInfos().add(columnInfo);
        });
    }
    @Override
    public void enterSelect_no_parens(PostgreSQLParser.Select_no_parensContext ctx) {
        info.enterSubQuery(ctx);
    }

    @Override
    public void exitSelect_no_parens(PostgreSQLParser.Select_no_parensContext ctx) {
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterWhere_clause(PostgreSQLParser.Where_clauseContext ctx) {
        info.enterWhere(ctx);
    }

    @Override
    public void exitWhere_clause(PostgreSQLParser.Where_clauseContext ctx) {
        info.exitWhere(ctx);
    }




}
