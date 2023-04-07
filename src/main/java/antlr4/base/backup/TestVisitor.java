package com.thinice.bbc.core.parser.antlr.base;

import java.util.List;

public class TestVisitor extends SqlBaseBaseVisitor {
    public void visitCreateTable(SqlBaseParser.StatementContext ctx) {
        String tableName = ctx.getRuleContext(SqlBaseParser.QualifiedNameContext.class, 0).getText();
        System.out.println("TableName: " + tableName);
        List<SqlBaseParser.TableElementContext> tableElements = ctx.getRuleContexts(SqlBaseParser.TableElementContext.class);
        for (SqlBaseParser.TableElementContext tableElement : tableElements) {
            SqlBaseParser.ColumnDefinitionContext columnDefinition = tableElement.columnDefinition();
            String columnName = columnDefinition.getRuleContexts(SqlBaseParser.IdentifierContext.class).get(0).getText();
            String columnType = columnDefinition.getRuleContexts(SqlBaseParser.TypeContext.class).get(0).getText();
            System.out.println("Column: " + columnName + " Type: " + columnType);
        }
    }

}
