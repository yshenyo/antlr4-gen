package com.thinice.bbc.core.parser.antlr.parser;

import com.thinice.bbc.core.parser.antlr.mysql.MySqlParser;
import lombok.Data;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.tree.TerminalNode;

@Data
public class MysqlParserColumnInfo extends ParserColumnInfo {
    protected boolean inDefinition = false;
    @Override
    public void enterSubQuery(ParserRuleContext ctx) {
        if (subQueryStart == -1 && !(ctx.start.getStartIndex() == start && ctx.stop.getStopIndex() == stop)) {
            subQueryStart = ctx.start.getStartIndex();
            subQueryStop = ctx.stop.getStopIndex();
        }
    }
    @Override
    public void enterWhere(ParserRuleContext ctx) {
        TerminalNode where = ctx.getParent().getToken(MySqlParser.WHERE, 0);
        int start = ctx.start.getStartIndex();
        if (null != where && whereStart == -1 && where.getSymbol().getStopIndex() < start - 1 && getContent(where.getSymbol().getStopIndex() + 1, start - 1).trim().equals("")) {
            whereStart = start;
            whereStop = ctx.stop.getStopIndex();
        }
    }

}
