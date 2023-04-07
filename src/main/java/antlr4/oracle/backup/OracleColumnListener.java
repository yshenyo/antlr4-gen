package com.thinice.bbc.core.parser.antlr.oracle;

import com.thinice.bbc.core.parser.antlr.parser.ParserColumnInfo;
import com.thinice.bbc.core.parser.dto.info.ColumnInfo;
import lombok.Data;

import java.util.List;

@Data
public class OracleColumnListener extends PlSqlParserBaseListener {

    private ParserColumnInfo info;

    public OracleColumnListener(String content) {
        info = new ParserColumnInfo();
        info.setContent(content);
    }

    @Override
    public void enterSubquery(PlSqlParser.SubqueryContext ctx) {
        info.enterSubQuery(ctx);
    }
    @Override
    public void exitSubquery(PlSqlParser.SubqueryContext ctx) {
        info.exitSubQuery(ctx);
    }
    @Override
    public void enterWhere_clause(PlSqlParser.Where_clauseContext ctx) {
        info.enterWhere(ctx);
    }
    @Override
    public void exitWhere_clause(PlSqlParser.Where_clauseContext ctx) {
        info.exitWhere(ctx);
    }
    @Override
    public void enterSelect_list_elements(PlSqlParser.Select_list_elementsContext ctx) {
        if (info.getSubQueryStart() > -1 || null == ctx.ASTERISK()) {
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
    public void enterSelected_list(PlSqlParser.Selected_listContext ctx) {
        if (info.getSubQueryStart() > -1 || null == ctx.ASTERISK()) {
            return;
        }
        ColumnInfo columnInfo = new ColumnInfo();
        columnInfo.setOwner("");
        columnInfo.setColumn_type("");
        columnInfo.setWhere(false);
        if (null != info.getTableInfos() && info.getTableInfos().size() == 1) {
            columnInfo.setTable_name(info.getTableInfos().get(0).getTable_name());
        } else {
            columnInfo.setTable_name("UNKNOWN");
        }
        columnInfo.setColumn("*");
        columnInfo.setFull_name(columnInfo.getTable_name() + "." + columnInfo.getColumn());
        info.getColumnInfos().add(columnInfo);
    }

    @Override
    public void enterVariable_name(PlSqlParser.Variable_nameContext ctx) {
        if (info.getSubQueryStart() == -1 && null == ctx.bind_variable()) {
            ColumnInfo columnInfo = new ColumnInfo();
            columnInfo.setOwner("");
            if (info.getWhereStart() > -1 && ctx.start.getStartIndex() >= info.getWhereStart() && ctx.stop.getStopIndex() <= info.getWhereStop()) {
                columnInfo.setWhere(true);
            } else {
                columnInfo.setWhere(false);
            }
            if (ctx.id_expression().size() == 1) {
                if (null != info.getTableInfos() && info.getTableInfos().size() == 1) {
                    columnInfo.setTable_name(info.getTableInfos().get(0).getTable_name());
                } else {
                    columnInfo.setTable_name("UNKNOWN");
                }
                columnInfo.setColumn(info.getContent(ctx.id_expression(0)));
            } else {
                List<PlSqlParser.Id_expressionContext> ids = ctx.id_expression();
                String table = info.getContent(ids.get(ids.size() - 2)).toUpperCase();
                columnInfo.setColumn(info.getContent(ids.get(ids.size() - 1)));
                if (info.getTableNameMap().containsKey(table)) {
                    columnInfo.setTable_name(info.getTableNameMap().get(table).getTable_name());
                } else if (info.getTableAliasMap().containsKey(table)) {
                    columnInfo.setTable_name(info.getTableAliasMap().get(table).getTable_name());
                } else {
                    columnInfo.setTable_name(table);
                }
            }
            columnInfo.setColumn_type("");
            columnInfo.setFull_name(columnInfo.getTable_name() + "." + columnInfo.getColumn());
            info.getColumnInfos().add(columnInfo);
        }
    }

    @Override
    public void enterColumn_definition(PlSqlParser.Column_definitionContext ctx) {
        ColumnInfo columnInfo = new ColumnInfo();
        columnInfo.setOwner("");
        columnInfo.setWhere(false);
        columnInfo.setColumn(info.getContent(ctx.column_name()));
        if (null != ctx.datatype()) {
            if (null != ctx.datatype().native_datatype_element()) {
                columnInfo.setColumn_type(info.getContent(ctx.datatype().native_datatype_element()));
            } else {
                columnInfo.setColumn_type(info.getContent(ctx.datatype()));
            }
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
    public void enterColumn_list(PlSqlParser.Column_listContext ctx) {
        ctx.column_name().forEach(col -> {
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

}
