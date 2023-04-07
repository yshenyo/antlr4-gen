package com.thinice.bbc.core.parser.antlr.informix;

import com.thinice.bbc.core.parser.antlr.parser.ParserColumnInfo;
import lombok.Data;

@Data
public class InformixColumnListener extends InformixSQLParserBaseListener {

    private ParserColumnInfo info;

    public InformixColumnListener(String content, int start, int stop) {
        info = new ParserColumnInfo();
        info.setContent(content);
        info.setStart(start);
        info.setStop(stop);
    }

}
