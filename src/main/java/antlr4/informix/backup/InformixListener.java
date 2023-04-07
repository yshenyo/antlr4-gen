package com.thinice.bbc.core.parser.antlr.informix;

import com.thinice.bbc.core.parser.antlr.parser.ParserInfo;
import com.thinice.bbc.core.parser.dto.info.TableInfo;
import lombok.Data;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.tree.ParseTreeWalker;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Data
public class InformixListener extends InformixSQLParserBaseListener {

    private ParserInfo info;

    public InformixListener(String content) {
        info = new ParserInfo();
        info.setContent(content);
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
        InformixColumnListener listener = new InformixColumnListener(info.getContent(), ctx.start.getStartIndex(), ctx.stop.getStopIndex());
        listener.getInfo().setTableInfos(tables);
        listener.getInfo().setTableNameMap(tableNameMap);
        listener.getInfo().setTableAliasMap(tableAliasMap);
        ParseTreeWalker.DEFAULT.walk(listener, ctx);
        if (null != listener.getInfo().getColumnInfos()) {
            info.getColumnInfos().addAll(listener.getInfo().getColumnInfos());
        }
    }
}
