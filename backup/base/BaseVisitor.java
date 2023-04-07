package com.thinice.bbc.core.parser.antlr.base;


import com.thinice.bbc.core.parser.dto.info.TableInfo;
import lombok.Data;

import java.util.ArrayList;
import java.util.List;

@Data
public class BaseVisitor extends SqlBaseBaseVisitor<String> {

    private String statement;

    private List<TableInfo> tableInfos = new ArrayList<>();

    public void parse(SqlBaseParser.StatementContext statementContext) {
        String stmt = statementContext.getClass().getSimpleName();
        switch (stmt) {
            case "StatementDefaultContext"://select or with
                SqlBaseParser.StatementDefaultContext defaultContext = (SqlBaseParser.StatementDefaultContext) statementContext;
                SqlBaseParser.WithContext with = defaultContext.query().with();
                if (with != null) {
                    statement = "With";
                } else {
                    statement = "Select";
                }
                break;
            default:
                statement = stmt.replaceAll("Context", "");
                break;
        }
    }

}
