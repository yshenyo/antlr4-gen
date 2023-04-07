package com.thinice.bbc.core.parser.antlr.mysql;

import com.thinice.bbc.core.parser.antlr.parser.ParserInfo;
import com.thinice.bbc.core.parser.dto.info.ConditionInfo;
import lombok.Data;

import java.util.List;

@Data
public class MysqlExpressionListener extends MySqlParserBaseListener {

    private ParserInfo info;

    public MysqlExpressionListener(String content) {
        info = new ParserInfo();
        info.setContent(content);
    }

    @Override
    public void enterInPredicate(MySqlParser.InPredicateContext ctx) {
        if (null != ctx.IN()) {
            String operator = "IN";
            if (null != ctx.NOT()) {
                operator = "NOT IN";
            }
            ConditionInfo conditionInfo = new ConditionInfo();
            conditionInfo.setOperator(operator);
            conditionInfo.setLeft(info.getContent(ctx.predicate()));
            if (null != ctx.selectStatement()) {
                conditionInfo.setRight("(" + info.getContent(ctx.selectStatement()) + ")");
            } else if (null != ctx.expressions()) {
                conditionInfo.setRight("(" + info.getContent(ctx.expressions()) + ")");
            }
            conditionInfo.setColumn_name(conditionInfo.getLeft());
            conditionInfo.setCondition(info.getContent(ctx));
            info.getConditions().add(conditionInfo);
        }
    }
    @Override
    public void enterIsNullPredicate(MySqlParser.IsNullPredicateContext ctx) {
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setOperator("IS");
        conditionInfo.setLeft(info.getContent(ctx.predicate()));
        conditionInfo.setRight(info.getContent(ctx.nullNotnull()));
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        conditionInfo.setCondition(info.getContent(ctx));
        info.getConditions().add(conditionInfo);
    }
    @Override
    public void enterBinaryComparisonPredicate(MySqlParser.BinaryComparisonPredicateContext ctx) {
        List<MySqlParser.PredicateContext> predicateContexts = ctx.predicate();
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setOperator(info.getContent(ctx.comparisonOperator()));
        conditionInfo.setLeft(info.getContent(predicateContexts.get(0)));
        conditionInfo.setRight(info.getContent(predicateContexts.get(1)));
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        conditionInfo.setCondition(info.getContent(ctx));
        info.getConditions().add(conditionInfo);
    }
    @Override
    public void enterSubqueryComparisonPredicate(MySqlParser.SubqueryComparisonPredicateContext ctx) {
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setOperator(info.getContent(ctx.comparisonOperator()));
        conditionInfo.setLeft(info.getContent(ctx.predicate()));
        conditionInfo.setRight(info.getContent(ctx.comparisonOperator().stop.getStopIndex() + 1, ctx.stop.getStopIndex()));
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        conditionInfo.setCondition(info.getContent(ctx));
        info.getConditions().add(conditionInfo);
    }
    @Override
    public void enterBetweenPredicate(MySqlParser.BetweenPredicateContext ctx) {
        String operator = "BETWEEN";
        if (null != ctx.NOT()) {
            operator = "NOT BETWEEN";
        }
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setOperator(operator);
        conditionInfo.setLeft(info.getContent(ctx.predicate(0)));
        conditionInfo.setRight(info.getContent(ctx.predicate(1).start.getStartIndex(), ctx.stop.getStopIndex()));
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        conditionInfo.setCondition(info.getContent(ctx));
        info.getConditions().add(conditionInfo);
    }
    @Override
    public void enterSoundsLikePredicate(MySqlParser.SoundsLikePredicateContext ctx) {
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setOperator("SOUNDS LIKE");
        conditionInfo.setLeft(info.getContent(ctx.predicate(0)));
        conditionInfo.setRight(info.getContent(ctx.predicate(1)));
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        conditionInfo.setCondition(info.getContent(ctx));
        info.getConditions().add(conditionInfo);
    }
    @Override
    public void enterLikePredicate(MySqlParser.LikePredicateContext ctx) {
        String operator = "LIKE";
        if (null != ctx.NOT()) {
            operator = "NOT LIKE";
        }
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setOperator(operator);
        conditionInfo.setLeft(info.getContent(ctx.predicate(0)));
        conditionInfo.setRight(info.getContent(ctx.predicate(1).start.getStartIndex(), ctx.stop.getStopIndex()));
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        conditionInfo.setCondition(info.getContent(ctx));
        info.getConditions().add(conditionInfo);
    }
    @Override
    public void enterRegexpPredicate(MySqlParser.RegexpPredicateContext ctx) {
        String operator = "";
        if (null != ctx.NOT()) {
            operator = "NOT ";
        }
        if (null == ctx.REGEXP()) {
            operator += "RLIKE";
        } else {
            operator += "REGEXP";
        }
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setOperator(operator);
        conditionInfo.setLeft(info.getContent(ctx.predicate(0)));
        conditionInfo.setRight(info.getContent(ctx.predicate(1)));
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        conditionInfo.setCondition(info.getContent(ctx));
        info.getConditions().add(conditionInfo);
    }
    @Override
    public void enterExpressionAtomPredicate(MySqlParser.ExpressionAtomPredicateContext ctx) {

    }
    @Override
    public void enterJsonMemberOfPredicate(MySqlParser.JsonMemberOfPredicateContext ctx) {
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setOperator("MEMBER OF");
        conditionInfo.setLeft(info.getContent(ctx.predicate(0)));
        conditionInfo.setRight(info.getContent(ctx.predicate(1)));
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        conditionInfo.setCondition(info.getContent(ctx));
        info.getConditions().add(conditionInfo);
    }


}
