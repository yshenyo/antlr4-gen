package com.thinice.bbc.core.parser.antlr.informix;
// Generated from InformixSQLParser.g4 by ANTLR 4.12.0
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link InformixSQLParser}.
 */
public interface InformixSQLParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link InformixSQLParser#sqlScript}.
	 * @param ctx the parse tree
	 */
	void enterSqlScript(InformixSQLParser.SqlScriptContext ctx);
	/**
	 * Exit a parse tree produced by {@link InformixSQLParser#sqlScript}.
	 * @param ctx the parse tree
	 */
	void exitSqlScript(InformixSQLParser.SqlScriptContext ctx);
	/**
	 * Enter a parse tree produced by {@link InformixSQLParser#unitStatement}.
	 * @param ctx the parse tree
	 */
	void enterUnitStatement(InformixSQLParser.UnitStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link InformixSQLParser#unitStatement}.
	 * @param ctx the parse tree
	 */
	void exitUnitStatement(InformixSQLParser.UnitStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link InformixSQLParser#createRole}.
	 * @param ctx the parse tree
	 */
	void enterCreateRole(InformixSQLParser.CreateRoleContext ctx);
	/**
	 * Exit a parse tree produced by {@link InformixSQLParser#createRole}.
	 * @param ctx the parse tree
	 */
	void exitCreateRole(InformixSQLParser.CreateRoleContext ctx);
	/**
	 * Enter a parse tree produced by {@link InformixSQLParser#dropRole}.
	 * @param ctx the parse tree
	 */
	void enterDropRole(InformixSQLParser.DropRoleContext ctx);
	/**
	 * Exit a parse tree produced by {@link InformixSQLParser#dropRole}.
	 * @param ctx the parse tree
	 */
	void exitDropRole(InformixSQLParser.DropRoleContext ctx);
	/**
	 * Enter a parse tree produced by {@link InformixSQLParser#keyword}.
	 * @param ctx the parse tree
	 */
	void enterKeyword(InformixSQLParser.KeywordContext ctx);
	/**
	 * Exit a parse tree produced by {@link InformixSQLParser#keyword}.
	 * @param ctx the parse tree
	 */
	void exitKeyword(InformixSQLParser.KeywordContext ctx);
	/**
	 * Enter a parse tree produced by {@link InformixSQLParser#roleName}.
	 * @param ctx the parse tree
	 */
	void enterRoleName(InformixSQLParser.RoleNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link InformixSQLParser#roleName}.
	 * @param ctx the parse tree
	 */
	void exitRoleName(InformixSQLParser.RoleNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link InformixSQLParser#anyName}.
	 * @param ctx the parse tree
	 */
	void enterAnyName(InformixSQLParser.AnyNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link InformixSQLParser#anyName}.
	 * @param ctx the parse tree
	 */
	void exitAnyName(InformixSQLParser.AnyNameContext ctx);
}