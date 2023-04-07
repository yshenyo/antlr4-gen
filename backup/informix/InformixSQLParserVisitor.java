package com.thinice.bbc.core.parser.antlr.informix;
// Generated from InformixSQLParser.g4 by ANTLR 4.12.0
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link InformixSQLParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface InformixSQLParserVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link InformixSQLParser#sqlScript}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSqlScript(InformixSQLParser.SqlScriptContext ctx);
	/**
	 * Visit a parse tree produced by {@link InformixSQLParser#unitStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnitStatement(InformixSQLParser.UnitStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link InformixSQLParser#createRole}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCreateRole(InformixSQLParser.CreateRoleContext ctx);
	/**
	 * Visit a parse tree produced by {@link InformixSQLParser#dropRole}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDropRole(InformixSQLParser.DropRoleContext ctx);
	/**
	 * Visit a parse tree produced by {@link InformixSQLParser#keyword}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKeyword(InformixSQLParser.KeywordContext ctx);
	/**
	 * Visit a parse tree produced by {@link InformixSQLParser#roleName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRoleName(InformixSQLParser.RoleNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link InformixSQLParser#anyName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAnyName(InformixSQLParser.AnyNameContext ctx);
}