package com.thinice.bbc.core.parser.antlr.hive.v4;

import com.thinice.bbc.core.parser.antlr.parser.ParserInfo;
import com.thinice.bbc.core.parser.dto.info.ConditionInfo;
import lombok.Data;
@Data
public class HiveV4ConditionListener extends HiveParserBaseListener {
    private ParserInfo info;
    private int precedencStart = -1;
    private int precedencStop = -1;

    public HiveV4ConditionListener(String content) {
        info = new ParserInfo();
        info.setContent(content);
    }


    @Override
    public void enterPrecedenceEqualExpression(HiveParser.PrecedenceEqualExpressionContext ctx) {
        ConditionInfo conditionInfo = new ConditionInfo();
        conditionInfo.setLeft(info.getContent(ctx.precedenceSimilarExpression(0)));
        conditionInfo.setColumn_name(conditionInfo.getLeft());
        if (null != ctx.precedenceEqualOperator(0)) {
            HiveParser.PrecedenceEqualOperatorContext operator = ctx.precedenceEqualOperator(0);
            conditionInfo.setOperator(info.getContent(operator).toUpperCase());
            conditionInfo.setRight(info.getContent(operator.stop.getStopIndex()+1, ctx.stop.getStopIndex()));
        } else if (null != ctx.precedenceDistinctOperator(0)) {
            HiveParser.PrecedenceDistinctOperatorContext operator = ctx.precedenceDistinctOperator(0);
            conditionInfo.setOperator(info.getContent(operator).toUpperCase());
            conditionInfo.setRight(info.getContent(operator.stop.getStopIndex()+1, ctx.stop.getStopIndex()));
        }
        conditionInfo.setCondition(info.getContent(ctx));
        info.getConditions().add(conditionInfo);
    }
    @Override
    public void enterPrecedenceSimilarExpressionMain(HiveParser.PrecedenceSimilarExpressionMainContext ctx) {
        if (precedencStart < 0 && ctx.getChildCount() > 1) {
            ConditionInfo conditionInfo = new ConditionInfo();
            conditionInfo.setLeft(info.getContent(ctx.precedenceBitwiseOrExpression()));
            conditionInfo.setColumn_name(conditionInfo.getLeft());
            HiveParser.PrecedenceSimilarExpressionPartContext part = ctx.precedenceSimilarExpressionPart();
            if (null != part) {
                if (null != part.precedenceSimilarOperator()) {
                    HiveParser.PrecedenceSimilarOperatorContext operator = part.precedenceSimilarOperator();
                    conditionInfo.setOperator(info.getContent(operator).toUpperCase());
                    conditionInfo.setRight(info.getContent(operator.stop.getStopIndex()+1, ctx.stop.getStopIndex()));
                } else if (null != part.KW_NOT()) {
                    HiveParser.PrecedenceSimilarExpressionPartNotContext not = part.precedenceSimilarExpressionPartNot();
                    if (null != not) {
                        if (null != not.precedenceRegexpOperator()) {
                            conditionInfo.setOperator("NOT " + info.getContent(not.precedenceRegexpOperator()).toUpperCase());
                            conditionInfo.setRight(info.getContent(not.precedenceBitwiseOrExpression()));
                        } else {
                            HiveParser.PrecedenceSimilarExpressionAtomContext atom = not.precedenceSimilarExpressionAtom();
                            if (null != atom) {
                                if (null != atom.KW_IN()) {
                                    String operator = "NOT IN";
                                    conditionInfo.setOperator(operator);
                                    conditionInfo.setRight(info.getContent(atom.precedenceSimilarExpressionIn()));
                                } else if (null != atom.KW_BETWEEN()) {
                                    String operator = "NOT BETWEEN";
                                    conditionInfo.setOperator(operator);
                                    conditionInfo.setRight(info.getContent(atom.precedenceBitwiseOrExpression(0).start.getStartIndex(), atom.stop.getStopIndex()));
                                } else if (null != atom.KW_LIKE()) {
                                    String operator = "NOT LIKE";
                                    conditionInfo.setOperator(operator);
                                    conditionInfo.setRight(info.getContent(atom.expressionsInParenthesis()));
                                } else if (null != atom.precedenceSimilarExpressionQuantifierPredicate()) {
                                    HiveParser.PrecedenceSimilarExpressionQuantifierPredicateContext predicate = atom.precedenceSimilarExpressionQuantifierPredicate();
                                    conditionInfo.setOperator("NOT " + info.getContent(predicate.subQuerySelectorOperator()).toUpperCase());
                                    conditionInfo.setRight(info.getContent(predicate.quantifierType().start.getStartIndex(), predicate.stop.getStopIndex()));
                                }
                            }
                        }
                    }
                } else {
                    HiveParser.PrecedenceSimilarExpressionAtomContext atom = part.precedenceSimilarExpressionAtom();
                    if (null != atom) {
                        if (null != atom.KW_IN()) {
                            String operator = "IN";
                            conditionInfo.setOperator(operator);
                            conditionInfo.setRight(info.getContent(atom.precedenceSimilarExpressionIn()));
                        } else if (null != atom.KW_BETWEEN()) {
                            String operator = "BETWEEN";
                            conditionInfo.setOperator(operator);
                            conditionInfo.setRight(info.getContent(atom.precedenceBitwiseOrExpression(0).start.getStartIndex(), atom.stop.getStopIndex()));
                        } else if (null != atom.KW_LIKE()) {
                            String operator = "LIKE";
                            conditionInfo.setOperator(operator);
                            conditionInfo.setRight(info.getContent(atom.expressionsInParenthesis()));
                        } else if (null != atom.precedenceSimilarExpressionQuantifierPredicate()) {
                            HiveParser.PrecedenceSimilarExpressionQuantifierPredicateContext predicate = atom.precedenceSimilarExpressionQuantifierPredicate();
                            conditionInfo.setOperator(info.getContent(predicate.subQuerySelectorOperator()).toUpperCase());
                            conditionInfo.setRight(info.getContent(predicate.quantifierType().start.getStartIndex(), predicate.stop.getStopIndex()));
                        }
                    }
                }
            }
            conditionInfo.setCondition(info.getContent(ctx));
            info.getConditions().add(conditionInfo);
            precedencStart = ctx.start.getStartIndex();
            precedencStop = ctx.stop.getStopIndex();
        }
    }
    @Override
    public void exitPrecedenceSimilarExpressionMain(HiveParser.PrecedenceSimilarExpressionMainContext ctx) {
        if (precedencStart == ctx.start.getStartIndex() && precedencStop == ctx.stop.getStopIndex()) {
            precedencStart = -1;
            precedencStop = -1;
        }
    }


}
