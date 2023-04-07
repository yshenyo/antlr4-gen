// Code generated from DB2zSQLParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // DB2zSQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseDB2zSQLParserListener is a complete listener for a parse tree produced by DB2zSQLParser.
type BaseDB2zSQLParserListener struct{}

var _ DB2zSQLParserListener = &BaseDB2zSQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDB2zSQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDB2zSQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDB2zSQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDB2zSQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStartRule is called when production startRule is entered.
func (s *BaseDB2zSQLParserListener) EnterStartRule(ctx *StartRuleContext) {}

// ExitStartRule is called when production startRule is exited.
func (s *BaseDB2zSQLParserListener) ExitStartRule(ctx *StartRuleContext) {}

// EnterSqlStatement is called when production sqlStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlStatement(ctx *SqlStatementContext) {}

// ExitSqlStatement is called when production sqlStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlStatement(ctx *SqlStatementContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseDB2zSQLParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseDB2zSQLParserListener) ExitQuery(ctx *QueryContext) {}

// EnterCursorName is called when production cursorName is entered.
func (s *BaseDB2zSQLParserListener) EnterCursorName(ctx *CursorNameContext) {}

// ExitCursorName is called when production cursorName is exited.
func (s *BaseDB2zSQLParserListener) ExitCursorName(ctx *CursorNameContext) {}

// EnterStatementName is called when production statementName is entered.
func (s *BaseDB2zSQLParserListener) EnterStatementName(ctx *StatementNameContext) {}

// ExitStatementName is called when production statementName is exited.
func (s *BaseDB2zSQLParserListener) ExitStatementName(ctx *StatementNameContext) {}

// EnterDescriptorName is called when production descriptorName is entered.
func (s *BaseDB2zSQLParserListener) EnterDescriptorName(ctx *DescriptorNameContext) {}

// ExitDescriptorName is called when production descriptorName is exited.
func (s *BaseDB2zSQLParserListener) ExitDescriptorName(ctx *DescriptorNameContext) {}

// EnterHoldability is called when production holdability is entered.
func (s *BaseDB2zSQLParserListener) EnterHoldability(ctx *HoldabilityContext) {}

// ExitHoldability is called when production holdability is exited.
func (s *BaseDB2zSQLParserListener) ExitHoldability(ctx *HoldabilityContext) {}

// EnterReturnability is called when production returnability is entered.
func (s *BaseDB2zSQLParserListener) EnterReturnability(ctx *ReturnabilityContext) {}

// ExitReturnability is called when production returnability is exited.
func (s *BaseDB2zSQLParserListener) ExitReturnability(ctx *ReturnabilityContext) {}

// EnterRowsetPositioning is called when production rowsetPositioning is entered.
func (s *BaseDB2zSQLParserListener) EnterRowsetPositioning(ctx *RowsetPositioningContext) {}

// ExitRowsetPositioning is called when production rowsetPositioning is exited.
func (s *BaseDB2zSQLParserListener) ExitRowsetPositioning(ctx *RowsetPositioningContext) {}

// EnterNotNullPhrase is called when production notNullPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterNotNullPhrase(ctx *NotNullPhraseContext) {}

// ExitNotNullPhrase is called when production notNullPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitNotNullPhrase(ctx *NotNullPhraseContext) {}

// EnterAllocateCursorStatement is called when production allocateCursorStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAllocateCursorStatement(ctx *AllocateCursorStatementContext) {
}

// ExitAllocateCursorStatement is called when production allocateCursorStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAllocateCursorStatement(ctx *AllocateCursorStatementContext) {
}

// EnterRsLocatorVariable is called when production rsLocatorVariable is entered.
func (s *BaseDB2zSQLParserListener) EnterRsLocatorVariable(ctx *RsLocatorVariableContext) {}

// ExitRsLocatorVariable is called when production rsLocatorVariable is exited.
func (s *BaseDB2zSQLParserListener) ExitRsLocatorVariable(ctx *RsLocatorVariableContext) {}

// EnterAlterDatabaseStatement is called when production alterDatabaseStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterDatabaseStatement(ctx *AlterDatabaseStatementContext) {}

// ExitAlterDatabaseStatement is called when production alterDatabaseStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterDatabaseStatement(ctx *AlterDatabaseStatementContext) {}

// EnterAlterFunctionStatement is called when production alterFunctionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterFunctionStatement(ctx *AlterFunctionStatementContext) {}

// ExitAlterFunctionStatement is called when production alterFunctionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterFunctionStatement(ctx *AlterFunctionStatementContext) {}

// EnterAlterFunctionStatementExternal is called when production alterFunctionStatementExternal is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterFunctionStatementExternal(ctx *AlterFunctionStatementExternalContext) {
}

// ExitAlterFunctionStatementExternal is called when production alterFunctionStatementExternal is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterFunctionStatementExternal(ctx *AlterFunctionStatementExternalContext) {
}

// EnterAlterFunctionStatementCompiledSqlScalar is called when production alterFunctionStatementCompiledSqlScalar is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterFunctionStatementCompiledSqlScalar(ctx *AlterFunctionStatementCompiledSqlScalarContext) {
}

// ExitAlterFunctionStatementCompiledSqlScalar is called when production alterFunctionStatementCompiledSqlScalar is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterFunctionStatementCompiledSqlScalar(ctx *AlterFunctionStatementCompiledSqlScalarContext) {
}

// EnterAlterWhichFunction1 is called when production alterWhichFunction1 is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterWhichFunction1(ctx *AlterWhichFunction1Context) {}

// ExitAlterWhichFunction1 is called when production alterWhichFunction1 is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterWhichFunction1(ctx *AlterWhichFunction1Context) {}

// EnterAlterWhichFunction2 is called when production alterWhichFunction2 is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterWhichFunction2(ctx *AlterWhichFunction2Context) {}

// ExitAlterWhichFunction2 is called when production alterWhichFunction2 is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterWhichFunction2(ctx *AlterWhichFunction2Context) {}

// EnterFunctionCompiledSqlScalarRoutineSpecification is called when production functionCompiledSqlScalarRoutineSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionCompiledSqlScalarRoutineSpecification(ctx *FunctionCompiledSqlScalarRoutineSpecificationContext) {
}

// ExitFunctionCompiledSqlScalarRoutineSpecification is called when production functionCompiledSqlScalarRoutineSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionCompiledSqlScalarRoutineSpecification(ctx *FunctionCompiledSqlScalarRoutineSpecificationContext) {
}

// EnterAlterFunctionStatementInlineSqlScalar is called when production alterFunctionStatementInlineSqlScalar is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterFunctionStatementInlineSqlScalar(ctx *AlterFunctionStatementInlineSqlScalarContext) {
}

// ExitAlterFunctionStatementInlineSqlScalar is called when production alterFunctionStatementInlineSqlScalar is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterFunctionStatementInlineSqlScalar(ctx *AlterFunctionStatementInlineSqlScalarContext) {
}

// EnterAlterFunctionStatementSqlTable is called when production alterFunctionStatementSqlTable is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterFunctionStatementSqlTable(ctx *AlterFunctionStatementSqlTableContext) {
}

// ExitAlterFunctionStatementSqlTable is called when production alterFunctionStatementSqlTable is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterFunctionStatementSqlTable(ctx *AlterFunctionStatementSqlTableContext) {
}

// EnterFunctionReturnsClause is called when production functionReturnsClause is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionReturnsClause(ctx *FunctionReturnsClauseContext) {}

// ExitFunctionReturnsClause is called when production functionReturnsClause is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionReturnsClause(ctx *FunctionReturnsClauseContext) {}

// EnterFunctionDesignator1 is called when production functionDesignator1 is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionDesignator1(ctx *FunctionDesignator1Context) {}

// ExitFunctionDesignator1 is called when production functionDesignator1 is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionDesignator1(ctx *FunctionDesignator1Context) {}

// EnterFunctionDesignator2 is called when production functionDesignator2 is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionDesignator2(ctx *FunctionDesignator2Context) {}

// ExitFunctionDesignator2 is called when production functionDesignator2 is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionDesignator2(ctx *FunctionDesignator2Context) {}

// EnterAlterIndexStatement is called when production alterIndexStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterIndexStatement(ctx *AlterIndexStatementContext) {}

// ExitAlterIndexStatement is called when production alterIndexStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterIndexStatement(ctx *AlterIndexStatementContext) {}

// EnterAlterMaskStatement is called when production alterMaskStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterMaskStatement(ctx *AlterMaskStatementContext) {}

// ExitAlterMaskStatement is called when production alterMaskStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterMaskStatement(ctx *AlterMaskStatementContext) {}

// EnterAlterPermissionStatement is called when production alterPermissionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterPermissionStatement(ctx *AlterPermissionStatementContext) {
}

// ExitAlterPermissionStatement is called when production alterPermissionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterPermissionStatement(ctx *AlterPermissionStatementContext) {
}

// EnterAlterProcedureStatement is called when production alterProcedureStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterProcedureStatement(ctx *AlterProcedureStatementContext) {
}

// ExitAlterProcedureStatement is called when production alterProcedureStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterProcedureStatement(ctx *AlterProcedureStatementContext) {
}

// EnterAlterProcedureSQLPLStatement is called when production alterProcedureSQLPLStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterProcedureSQLPLStatement(ctx *AlterProcedureSQLPLStatementContext) {
}

// ExitAlterProcedureSQLPLStatement is called when production alterProcedureSQLPLStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterProcedureSQLPLStatement(ctx *AlterProcedureSQLPLStatementContext) {
}

// EnterAlterWhichProcedureSQLPL1 is called when production alterWhichProcedureSQLPL1 is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterWhichProcedureSQLPL1(ctx *AlterWhichProcedureSQLPL1Context) {
}

// ExitAlterWhichProcedureSQLPL1 is called when production alterWhichProcedureSQLPL1 is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterWhichProcedureSQLPL1(ctx *AlterWhichProcedureSQLPL1Context) {
}

// EnterAlterWhichProcedureSQLPL2 is called when production alterWhichProcedureSQLPL2 is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterWhichProcedureSQLPL2(ctx *AlterWhichProcedureSQLPL2Context) {
}

// ExitAlterWhichProcedureSQLPL2 is called when production alterWhichProcedureSQLPL2 is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterWhichProcedureSQLPL2(ctx *AlterWhichProcedureSQLPL2Context) {
}

// EnterApplicationCompatibilityPhrase is called when production applicationCompatibilityPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterApplicationCompatibilityPhrase(ctx *ApplicationCompatibilityPhraseContext) {
}

// ExitApplicationCompatibilityPhrase is called when production applicationCompatibilityPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitApplicationCompatibilityPhrase(ctx *ApplicationCompatibilityPhraseContext) {
}

// EnterAlterSequenceStatement is called when production alterSequenceStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterSequenceStatement(ctx *AlterSequenceStatementContext) {}

// ExitAlterSequenceStatement is called when production alterSequenceStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterSequenceStatement(ctx *AlterSequenceStatementContext) {}

// EnterAlterStogroupStatement is called when production alterStogroupStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterStogroupStatement(ctx *AlterStogroupStatementContext) {}

// ExitAlterStogroupStatement is called when production alterStogroupStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterStogroupStatement(ctx *AlterStogroupStatementContext) {}

// EnterAlterTableStatement is called when production alterTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTableStatement(ctx *AlterTableStatementContext) {}

// ExitAlterTableStatement is called when production alterTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTableStatement(ctx *AlterTableStatementContext) {}

// EnterAlterTablespaceStatement is called when production alterTablespaceStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTablespaceStatement(ctx *AlterTablespaceStatementContext) {
}

// ExitAlterTablespaceStatement is called when production alterTablespaceStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTablespaceStatement(ctx *AlterTablespaceStatementContext) {
}

// EnterAlterTriggerStatement is called when production alterTriggerStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTriggerStatement(ctx *AlterTriggerStatementContext) {}

// ExitAlterTriggerStatement is called when production alterTriggerStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTriggerStatement(ctx *AlterTriggerStatementContext) {}

// EnterAlterTrustedContextStatement is called when production alterTrustedContextStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTrustedContextStatement(ctx *AlterTrustedContextStatementContext) {
}

// ExitAlterTrustedContextStatement is called when production alterTrustedContextStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTrustedContextStatement(ctx *AlterTrustedContextStatementContext) {
}

// EnterAlterViewStatement is called when production alterViewStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterViewStatement(ctx *AlterViewStatementContext) {}

// ExitAlterViewStatement is called when production alterViewStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterViewStatement(ctx *AlterViewStatementContext) {}

// EnterAssociateLocatorsStatement is called when production associateLocatorsStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterAssociateLocatorsStatement(ctx *AssociateLocatorsStatementContext) {
}

// ExitAssociateLocatorsStatement is called when production associateLocatorsStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitAssociateLocatorsStatement(ctx *AssociateLocatorsStatementContext) {
}

// EnterBeginDeclareSectionStatement is called when production beginDeclareSectionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterBeginDeclareSectionStatement(ctx *BeginDeclareSectionStatementContext) {
}

// ExitBeginDeclareSectionStatement is called when production beginDeclareSectionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitBeginDeclareSectionStatement(ctx *BeginDeclareSectionStatementContext) {
}

// EnterCallStatement is called when production callStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCallStatement(ctx *CallStatementContext) {}

// ExitCallStatement is called when production callStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCallStatement(ctx *CallStatementContext) {}

// EnterCloseStatement is called when production closeStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCloseStatement(ctx *CloseStatementContext) {}

// ExitCloseStatement is called when production closeStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCloseStatement(ctx *CloseStatementContext) {}

// EnterCommentStatement is called when production commentStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCommentStatement(ctx *CommentStatementContext) {}

// ExitCommentStatement is called when production commentStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCommentStatement(ctx *CommentStatementContext) {}

// EnterCommitStatement is called when production commitStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCommitStatement(ctx *CommitStatementContext) {}

// ExitCommitStatement is called when production commitStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCommitStatement(ctx *CommitStatementContext) {}

// EnterConnectStatement is called when production connectStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterConnectStatement(ctx *ConnectStatementContext) {}

// ExitConnectStatement is called when production connectStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitConnectStatement(ctx *ConnectStatementContext) {}

// EnterCreateAliasStatement is called when production createAliasStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateAliasStatement(ctx *CreateAliasStatementContext) {}

// ExitCreateAliasStatement is called when production createAliasStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateAliasStatement(ctx *CreateAliasStatementContext) {}

// EnterCreateAuxiliaryTableStatement is called when production createAuxiliaryTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateAuxiliaryTableStatement(ctx *CreateAuxiliaryTableStatementContext) {
}

// ExitCreateAuxiliaryTableStatement is called when production createAuxiliaryTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateAuxiliaryTableStatement(ctx *CreateAuxiliaryTableStatementContext) {
}

// EnterCreateDatabaseStatement is called when production createDatabaseStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) {
}

// ExitCreateDatabaseStatement is called when production createDatabaseStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) {
}

// EnterCreateFunctionStatement is called when production createFunctionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatement(ctx *CreateFunctionStatementContext) {
}

// ExitCreateFunctionStatement is called when production createFunctionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatement(ctx *CreateFunctionStatementContext) {
}

// EnterCreateFunctionStatementExternalScalar is called when production createFunctionStatementExternalScalar is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementExternalScalar(ctx *CreateFunctionStatementExternalScalarContext) {
}

// ExitCreateFunctionStatementExternalScalar is called when production createFunctionStatementExternalScalar is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementExternalScalar(ctx *CreateFunctionStatementExternalScalarContext) {
}

// EnterCreateFunctionStatementExternalScalarReturnsPhrase is called when production createFunctionStatementExternalScalarReturnsPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementExternalScalarReturnsPhrase(ctx *CreateFunctionStatementExternalScalarReturnsPhraseContext) {
}

// ExitCreateFunctionStatementExternalScalarReturnsPhrase is called when production createFunctionStatementExternalScalarReturnsPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementExternalScalarReturnsPhrase(ctx *CreateFunctionStatementExternalScalarReturnsPhraseContext) {
}

// EnterCreateFunctionStatementExternalTable is called when production createFunctionStatementExternalTable is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementExternalTable(ctx *CreateFunctionStatementExternalTableContext) {
}

// ExitCreateFunctionStatementExternalTable is called when production createFunctionStatementExternalTable is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementExternalTable(ctx *CreateFunctionStatementExternalTableContext) {
}

// EnterCreateFunctionStatementExternalTableReturnsPhrase is called when production createFunctionStatementExternalTableReturnsPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementExternalTableReturnsPhrase(ctx *CreateFunctionStatementExternalTableReturnsPhraseContext) {
}

// ExitCreateFunctionStatementExternalTableReturnsPhrase is called when production createFunctionStatementExternalTableReturnsPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementExternalTableReturnsPhrase(ctx *CreateFunctionStatementExternalTableReturnsPhraseContext) {
}

// EnterExternalTableFunctionColumn is called when production externalTableFunctionColumn is entered.
func (s *BaseDB2zSQLParserListener) EnterExternalTableFunctionColumn(ctx *ExternalTableFunctionColumnContext) {
}

// ExitExternalTableFunctionColumn is called when production externalTableFunctionColumn is exited.
func (s *BaseDB2zSQLParserListener) ExitExternalTableFunctionColumn(ctx *ExternalTableFunctionColumnContext) {
}

// EnterCreateFunctionStatementSourced is called when production createFunctionStatementSourced is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementSourced(ctx *CreateFunctionStatementSourcedContext) {
}

// ExitCreateFunctionStatementSourced is called when production createFunctionStatementSourced is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementSourced(ctx *CreateFunctionStatementSourcedContext) {
}

// EnterCreateFunctionStatementSourcedReturnsPhrase is called when production createFunctionStatementSourcedReturnsPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementSourcedReturnsPhrase(ctx *CreateFunctionStatementSourcedReturnsPhraseContext) {
}

// ExitCreateFunctionStatementSourcedReturnsPhrase is called when production createFunctionStatementSourcedReturnsPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementSourcedReturnsPhrase(ctx *CreateFunctionStatementSourcedReturnsPhraseContext) {
}

// EnterCreateFunctionStatementSourcedSourcePhrase is called when production createFunctionStatementSourcedSourcePhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementSourcedSourcePhrase(ctx *CreateFunctionStatementSourcedSourcePhraseContext) {
}

// ExitCreateFunctionStatementSourcedSourcePhrase is called when production createFunctionStatementSourcedSourcePhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementSourcedSourcePhrase(ctx *CreateFunctionStatementSourcedSourcePhraseContext) {
}

// EnterCreateFunctionStatementInlineSqlScalar is called when production createFunctionStatementInlineSqlScalar is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementInlineSqlScalar(ctx *CreateFunctionStatementInlineSqlScalarContext) {
}

// ExitCreateFunctionStatementInlineSqlScalar is called when production createFunctionStatementInlineSqlScalar is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementInlineSqlScalar(ctx *CreateFunctionStatementInlineSqlScalarContext) {
}

// EnterCreateFunctionStatementCompiledSqlScalar is called when production createFunctionStatementCompiledSqlScalar is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementCompiledSqlScalar(ctx *CreateFunctionStatementCompiledSqlScalarContext) {
}

// ExitCreateFunctionStatementCompiledSqlScalar is called when production createFunctionStatementCompiledSqlScalar is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementCompiledSqlScalar(ctx *CreateFunctionStatementCompiledSqlScalarContext) {
}

// EnterCreateFunctionStatementSqlTable is called when production createFunctionStatementSqlTable is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementSqlTable(ctx *CreateFunctionStatementSqlTableContext) {
}

// ExitCreateFunctionStatementSqlTable is called when production createFunctionStatementSqlTable is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementSqlTable(ctx *CreateFunctionStatementSqlTableContext) {
}

// EnterCreateGlobalTemporaryTableStatement is called when production createGlobalTemporaryTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateGlobalTemporaryTableStatement(ctx *CreateGlobalTemporaryTableStatementContext) {
}

// ExitCreateGlobalTemporaryTableStatement is called when production createGlobalTemporaryTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateGlobalTemporaryTableStatement(ctx *CreateGlobalTemporaryTableStatementContext) {
}

// EnterCreateIndexStatement is called when production createIndexStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// ExitCreateIndexStatement is called when production createIndexStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// EnterCreateLobTablespaceStatement is called when production createLobTablespaceStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateLobTablespaceStatement(ctx *CreateLobTablespaceStatementContext) {
}

// ExitCreateLobTablespaceStatement is called when production createLobTablespaceStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateLobTablespaceStatement(ctx *CreateLobTablespaceStatementContext) {
}

// EnterCreateMaskStatement is called when production createMaskStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateMaskStatement(ctx *CreateMaskStatementContext) {}

// ExitCreateMaskStatement is called when production createMaskStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateMaskStatement(ctx *CreateMaskStatementContext) {}

// EnterCreatePermissionStatement is called when production createPermissionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreatePermissionStatement(ctx *CreatePermissionStatementContext) {
}

// ExitCreatePermissionStatement is called when production createPermissionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreatePermissionStatement(ctx *CreatePermissionStatementContext) {
}

// EnterCreateProcedureSQLPLStatement is called when production createProcedureSQLPLStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateProcedureSQLPLStatement(ctx *CreateProcedureSQLPLStatementContext) {
}

// ExitCreateProcedureSQLPLStatement is called when production createProcedureSQLPLStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateProcedureSQLPLStatement(ctx *CreateProcedureSQLPLStatementContext) {
}

// EnterSqlRoutineBody is called when production sqlRoutineBody is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlRoutineBody(ctx *SqlRoutineBodyContext) {}

// ExitSqlRoutineBody is called when production sqlRoutineBody is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlRoutineBody(ctx *SqlRoutineBodyContext) {}

// EnterObfuscatedStatementText is called when production obfuscatedStatementText is entered.
func (s *BaseDB2zSQLParserListener) EnterObfuscatedStatementText(ctx *ObfuscatedStatementTextContext) {
}

// ExitObfuscatedStatementText is called when production obfuscatedStatementText is exited.
func (s *BaseDB2zSQLParserListener) ExitObfuscatedStatementText(ctx *ObfuscatedStatementTextContext) {
}

// EnterProbablySQLPL is called when production probablySQLPL is entered.
func (s *BaseDB2zSQLParserListener) EnterProbablySQLPL(ctx *ProbablySQLPLContext) {}

// ExitProbablySQLPL is called when production probablySQLPL is exited.
func (s *BaseDB2zSQLParserListener) ExitProbablySQLPL(ctx *ProbablySQLPLContext) {}

// EnterCreateProcedureStatement is called when production createProcedureStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateProcedureStatement(ctx *CreateProcedureStatementContext) {
}

// ExitCreateProcedureStatement is called when production createProcedureStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateProcedureStatement(ctx *CreateProcedureStatementContext) {
}

// EnterCreateRoleStatement is called when production createRoleStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// ExitCreateRoleStatement is called when production createRoleStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// EnterCreateSequenceStatement is called when production createSequenceStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateSequenceStatement(ctx *CreateSequenceStatementContext) {
}

// ExitCreateSequenceStatement is called when production createSequenceStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateSequenceStatement(ctx *CreateSequenceStatementContext) {
}

// EnterCreateStogroupStatement is called when production createStogroupStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateStogroupStatement(ctx *CreateStogroupStatementContext) {
}

// ExitCreateStogroupStatement is called when production createStogroupStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateStogroupStatement(ctx *CreateStogroupStatementContext) {
}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterCreateTableOptions is called when production createTableOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTableOptions(ctx *CreateTableOptionsContext) {}

// ExitCreateTableOptions is called when production createTableOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTableOptions(ctx *CreateTableOptionsContext) {}

// EnterCreateTablespaceStatement is called when production createTablespaceStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTablespaceStatement(ctx *CreateTablespaceStatementContext) {
}

// ExitCreateTablespaceStatement is called when production createTablespaceStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTablespaceStatement(ctx *CreateTablespaceStatementContext) {
}

// EnterCreateTriggerStatement is called when production createTriggerStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTriggerStatement(ctx *CreateTriggerStatementContext) {}

// ExitCreateTriggerStatement is called when production createTriggerStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTriggerStatement(ctx *CreateTriggerStatementContext) {}

// EnterCreateTrustedContextStatement is called when production createTrustedContextStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTrustedContextStatement(ctx *CreateTrustedContextStatementContext) {
}

// ExitCreateTrustedContextStatement is called when production createTrustedContextStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTrustedContextStatement(ctx *CreateTrustedContextStatementContext) {
}

// EnterCreateTypeArrayStatement is called when production createTypeArrayStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTypeArrayStatement(ctx *CreateTypeArrayStatementContext) {
}

// ExitCreateTypeArrayStatement is called when production createTypeArrayStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTypeArrayStatement(ctx *CreateTypeArrayStatementContext) {
}

// EnterCreateTypeDistinctStatement is called when production createTypeDistinctStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTypeDistinctStatement(ctx *CreateTypeDistinctStatementContext) {
}

// ExitCreateTypeDistinctStatement is called when production createTypeDistinctStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTypeDistinctStatement(ctx *CreateTypeDistinctStatementContext) {
}

// EnterCreateVariableStatement is called when production createVariableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateVariableStatement(ctx *CreateVariableStatementContext) {
}

// ExitCreateVariableStatement is called when production createVariableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateVariableStatement(ctx *CreateVariableStatementContext) {
}

// EnterCreateViewStatement is called when production createViewStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateViewStatement(ctx *CreateViewStatementContext) {}

// ExitCreateViewStatement is called when production createViewStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateViewStatement(ctx *CreateViewStatementContext) {}

// EnterDeclareCursorStatement is called when production declareCursorStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDeclareCursorStatement(ctx *DeclareCursorStatementContext) {}

// ExitDeclareCursorStatement is called when production declareCursorStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDeclareCursorStatement(ctx *DeclareCursorStatementContext) {}

// EnterDeclareGlobalTemporaryTableStatement is called when production declareGlobalTemporaryTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDeclareGlobalTemporaryTableStatement(ctx *DeclareGlobalTemporaryTableStatementContext) {
}

// ExitDeclareGlobalTemporaryTableStatement is called when production declareGlobalTemporaryTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDeclareGlobalTemporaryTableStatement(ctx *DeclareGlobalTemporaryTableStatementContext) {
}

// EnterDeclareTableStatement is called when production declareTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDeclareTableStatement(ctx *DeclareTableStatementContext) {}

// ExitDeclareTableStatement is called when production declareTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDeclareTableStatement(ctx *DeclareTableStatementContext) {}

// EnterDeclareStatementStatement is called when production declareStatementStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDeclareStatementStatement(ctx *DeclareStatementStatementContext) {
}

// ExitDeclareStatementStatement is called when production declareStatementStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDeclareStatementStatement(ctx *DeclareStatementStatementContext) {
}

// EnterDeclareVariableStatement is called when production declareVariableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDeclareVariableStatement(ctx *DeclareVariableStatementContext) {
}

// ExitDeclareVariableStatement is called when production declareVariableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDeclareVariableStatement(ctx *DeclareVariableStatementContext) {
}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterDescribeStatement is called when production describeStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDescribeStatement(ctx *DescribeStatementContext) {}

// ExitDescribeStatement is called when production describeStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDescribeStatement(ctx *DescribeStatementContext) {}

// EnterDescribeCursorStatement is called when production describeCursorStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDescribeCursorStatement(ctx *DescribeCursorStatementContext) {
}

// ExitDescribeCursorStatement is called when production describeCursorStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDescribeCursorStatement(ctx *DescribeCursorStatementContext) {
}

// EnterDescribeInputStatement is called when production describeInputStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDescribeInputStatement(ctx *DescribeInputStatementContext) {}

// ExitDescribeInputStatement is called when production describeInputStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDescribeInputStatement(ctx *DescribeInputStatementContext) {}

// EnterDescribeOutputStatement is called when production describeOutputStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDescribeOutputStatement(ctx *DescribeOutputStatementContext) {
}

// ExitDescribeOutputStatement is called when production describeOutputStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDescribeOutputStatement(ctx *DescribeOutputStatementContext) {
}

// EnterDescribeProcedureStatement is called when production describeProcedureStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDescribeProcedureStatement(ctx *DescribeProcedureStatementContext) {
}

// ExitDescribeProcedureStatement is called when production describeProcedureStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDescribeProcedureStatement(ctx *DescribeProcedureStatementContext) {
}

// EnterDescribeTableStatement is called when production describeTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDescribeTableStatement(ctx *DescribeTableStatementContext) {}

// ExitDescribeTableStatement is called when production describeTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDescribeTableStatement(ctx *DescribeTableStatementContext) {}

// EnterDropStatement is called when production dropStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterDropStatement(ctx *DropStatementContext) {}

// ExitDropStatement is called when production dropStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitDropStatement(ctx *DropStatementContext) {}

// EnterEndDeclareSectionStatement is called when production endDeclareSectionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterEndDeclareSectionStatement(ctx *EndDeclareSectionStatementContext) {
}

// ExitEndDeclareSectionStatement is called when production endDeclareSectionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitEndDeclareSectionStatement(ctx *EndDeclareSectionStatementContext) {
}

// EnterExchangeStatement is called when production exchangeStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterExchangeStatement(ctx *ExchangeStatementContext) {}

// ExitExchangeStatement is called when production exchangeStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitExchangeStatement(ctx *ExchangeStatementContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterExecuteStatement(ctx *ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitExecuteStatement(ctx *ExecuteStatementContext) {}

// EnterExecuteImmediateStatement is called when production executeImmediateStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterExecuteImmediateStatement(ctx *ExecuteImmediateStatementContext) {
}

// ExitExecuteImmediateStatement is called when production executeImmediateStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitExecuteImmediateStatement(ctx *ExecuteImmediateStatementContext) {
}

// EnterExplainStatement is called when production explainStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterExplainStatement(ctx *ExplainStatementContext) {}

// ExitExplainStatement is called when production explainStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitExplainStatement(ctx *ExplainStatementContext) {}

// EnterFetchStatement is called when production fetchStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterFetchStatement(ctx *FetchStatementContext) {}

// ExitFetchStatement is called when production fetchStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitFetchStatement(ctx *FetchStatementContext) {}

// EnterFreeLocatorStatement is called when production freeLocatorStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterFreeLocatorStatement(ctx *FreeLocatorStatementContext) {}

// ExitFreeLocatorStatement is called when production freeLocatorStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitFreeLocatorStatement(ctx *FreeLocatorStatementContext) {}

// EnterGetDiagnosticsStatement is called when production getDiagnosticsStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGetDiagnosticsStatement(ctx *GetDiagnosticsStatementContext) {
}

// ExitGetDiagnosticsStatement is called when production getDiagnosticsStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGetDiagnosticsStatement(ctx *GetDiagnosticsStatementContext) {
}

// EnterGrantStatement is called when production grantStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantStatement(ctx *GrantStatementContext) {}

// ExitGrantStatement is called when production grantStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantStatement(ctx *GrantStatementContext) {}

// EnterHoldLocatorStatement is called when production holdLocatorStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterHoldLocatorStatement(ctx *HoldLocatorStatementContext) {}

// ExitHoldLocatorStatement is called when production holdLocatorStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitHoldLocatorStatement(ctx *HoldLocatorStatementContext) {}

// EnterIncludeStatement is called when production includeStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterIncludeStatement(ctx *IncludeStatementContext) {}

// ExitIncludeStatement is called when production includeStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitIncludeStatement(ctx *IncludeStatementContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterLabelStatement is called when production labelStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterLabelStatement(ctx *LabelStatementContext) {}

// ExitLabelStatement is called when production labelStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitLabelStatement(ctx *LabelStatementContext) {}

// EnterLockTableStatement is called when production lockTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterLockTableStatement(ctx *LockTableStatementContext) {}

// ExitLockTableStatement is called when production lockTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitLockTableStatement(ctx *LockTableStatementContext) {}

// EnterMergeStatement is called when production mergeStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterMergeStatement(ctx *MergeStatementContext) {}

// ExitMergeStatement is called when production mergeStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitMergeStatement(ctx *MergeStatementContext) {}

// EnterOpenStatement is called when production openStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterOpenStatement(ctx *OpenStatementContext) {}

// ExitOpenStatement is called when production openStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitOpenStatement(ctx *OpenStatementContext) {}

// EnterPrepareStatement is called when production prepareStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterPrepareStatement(ctx *PrepareStatementContext) {}

// ExitPrepareStatement is called when production prepareStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitPrepareStatement(ctx *PrepareStatementContext) {}

// EnterRefreshTableStatement is called when production refreshTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// ExitRefreshTableStatement is called when production refreshTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// EnterReleaseConnectionStatement is called when production releaseConnectionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterReleaseConnectionStatement(ctx *ReleaseConnectionStatementContext) {
}

// ExitReleaseConnectionStatement is called when production releaseConnectionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitReleaseConnectionStatement(ctx *ReleaseConnectionStatementContext) {
}

// EnterReleaseSavepointStatement is called when production releaseSavepointStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterReleaseSavepointStatement(ctx *ReleaseSavepointStatementContext) {
}

// ExitReleaseSavepointStatement is called when production releaseSavepointStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitReleaseSavepointStatement(ctx *ReleaseSavepointStatementContext) {
}

// EnterRenameStatement is called when production renameStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRenameStatement(ctx *RenameStatementContext) {}

// ExitRenameStatement is called when production renameStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRenameStatement(ctx *RenameStatementContext) {}

// EnterRevokeStatement is called when production revokeStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeStatement(ctx *RevokeStatementContext) {}

// ExitRevokeStatement is called when production revokeStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeStatement(ctx *RevokeStatementContext) {}

// EnterRollbackStatement is called when production rollbackStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRollbackStatement(ctx *RollbackStatementContext) {}

// ExitRollbackStatement is called when production rollbackStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRollbackStatement(ctx *RollbackStatementContext) {}

// EnterSavepointStatement is called when production savepointStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSavepointStatement(ctx *SavepointStatementContext) {}

// ExitSavepointStatement is called when production savepointStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSavepointStatement(ctx *SavepointStatementContext) {}

// EnterSetAssignmentStatement is called when production setAssignmentStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSetAssignmentStatement(ctx *SetAssignmentStatementContext) {}

// ExitSetAssignmentStatement is called when production setAssignmentStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSetAssignmentStatement(ctx *SetAssignmentStatementContext) {}

// EnterSetConnectionStatement is called when production setConnectionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSetConnectionStatement(ctx *SetConnectionStatementContext) {}

// ExitSetConnectionStatement is called when production setConnectionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSetConnectionStatement(ctx *SetConnectionStatementContext) {}

// EnterSetEncryptionPasswordStatement is called when production setEncryptionPasswordStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSetEncryptionPasswordStatement(ctx *SetEncryptionPasswordStatementContext) {
}

// ExitSetEncryptionPasswordStatement is called when production setEncryptionPasswordStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSetEncryptionPasswordStatement(ctx *SetEncryptionPasswordStatementContext) {
}

// EnterSetPathStatement is called when production setPathStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSetPathStatement(ctx *SetPathStatementContext) {}

// ExitSetPathStatement is called when production setPathStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSetPathStatement(ctx *SetPathStatementContext) {}

// EnterSetSchemaStatement is called when production setSchemaStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSetSchemaStatement(ctx *SetSchemaStatementContext) {}

// ExitSetSchemaStatement is called when production setSchemaStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSetSchemaStatement(ctx *SetSchemaStatementContext) {}

// EnterSetSessionTimezoneStatement is called when production setSessionTimezoneStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSetSessionTimezoneStatement(ctx *SetSessionTimezoneStatementContext) {
}

// ExitSetSessionTimezoneStatement is called when production setSessionTimezoneStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSetSessionTimezoneStatement(ctx *SetSessionTimezoneStatementContext) {
}

// EnterSetSpecialRegisterStatement is called when production setSpecialRegisterStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSetSpecialRegisterStatement(ctx *SetSpecialRegisterStatementContext) {
}

// ExitSetSpecialRegisterStatement is called when production setSpecialRegisterStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSetSpecialRegisterStatement(ctx *SetSpecialRegisterStatementContext) {
}

// EnterSignalStatement is called when production signalStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSignalStatement(ctx *SignalStatementContext) {}

// ExitSignalStatement is called when production signalStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSignalStatement(ctx *SignalStatementContext) {}

// EnterTransferOwnershipStatement is called when production transferOwnershipStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterTransferOwnershipStatement(ctx *TransferOwnershipStatementContext) {
}

// ExitTransferOwnershipStatement is called when production transferOwnershipStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitTransferOwnershipStatement(ctx *TransferOwnershipStatementContext) {
}

// EnterTruncateStatement is called when production truncateStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterTruncateStatement(ctx *TruncateStatementContext) {}

// ExitTruncateStatement is called when production truncateStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitTruncateStatement(ctx *TruncateStatementContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterValuesStatement is called when production valuesStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesStatement(ctx *ValuesStatementContext) {}

// ExitValuesStatement is called when production valuesStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesStatement(ctx *ValuesStatementContext) {}

// EnterValuesIntoStatement is called when production valuesIntoStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesIntoStatement(ctx *ValuesIntoStatementContext) {}

// ExitValuesIntoStatement is called when production valuesIntoStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesIntoStatement(ctx *ValuesIntoStatementContext) {}

// EnterWheneverStatement is called when production wheneverStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterWheneverStatement(ctx *WheneverStatementContext) {}

// ExitWheneverStatement is called when production wheneverStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitWheneverStatement(ctx *WheneverStatementContext) {}

// EnterValuesIntoTargetVariable is called when production valuesIntoTargetVariable is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesIntoTargetVariable(ctx *ValuesIntoTargetVariableContext) {
}

// ExitValuesIntoTargetVariable is called when production valuesIntoTargetVariable is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesIntoTargetVariable(ctx *ValuesIntoTargetVariableContext) {
}

// EnterOwnedObject is called when production ownedObject is entered.
func (s *BaseDB2zSQLParserListener) EnterOwnedObject(ctx *OwnedObjectContext) {}

// ExitOwnedObject is called when production ownedObject is exited.
func (s *BaseDB2zSQLParserListener) ExitOwnedObject(ctx *OwnedObjectContext) {}

// EnterNewOwner is called when production newOwner is entered.
func (s *BaseDB2zSQLParserListener) EnterNewOwner(ctx *NewOwnerContext) {}

// ExitNewOwner is called when production newOwner is exited.
func (s *BaseDB2zSQLParserListener) ExitNewOwner(ctx *NewOwnerContext) {}

// EnterGrantCollectionStatement is called when production grantCollectionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantCollectionStatement(ctx *GrantCollectionStatementContext) {
}

// ExitGrantCollectionStatement is called when production grantCollectionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantCollectionStatement(ctx *GrantCollectionStatementContext) {
}

// EnterGrantDatabaseStatement is called when production grantDatabaseStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantDatabaseStatement(ctx *GrantDatabaseStatementContext) {}

// ExitGrantDatabaseStatement is called when production grantDatabaseStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantDatabaseStatement(ctx *GrantDatabaseStatementContext) {}

// EnterGrantFunctionOrProcedureStatement is called when production grantFunctionOrProcedureStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantFunctionOrProcedureStatement(ctx *GrantFunctionOrProcedureStatementContext) {
}

// ExitGrantFunctionOrProcedureStatement is called when production grantFunctionOrProcedureStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantFunctionOrProcedureStatement(ctx *GrantFunctionOrProcedureStatementContext) {
}

// EnterGrantPackageStatement is called when production grantPackageStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantPackageStatement(ctx *GrantPackageStatementContext) {}

// ExitGrantPackageStatement is called when production grantPackageStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantPackageStatement(ctx *GrantPackageStatementContext) {}

// EnterGrantPlanStatement is called when production grantPlanStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantPlanStatement(ctx *GrantPlanStatementContext) {}

// ExitGrantPlanStatement is called when production grantPlanStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantPlanStatement(ctx *GrantPlanStatementContext) {}

// EnterGrantSchemaStatement is called when production grantSchemaStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantSchemaStatement(ctx *GrantSchemaStatementContext) {}

// ExitGrantSchemaStatement is called when production grantSchemaStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantSchemaStatement(ctx *GrantSchemaStatementContext) {}

// EnterGrantSequenceStatement is called when production grantSequenceStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantSequenceStatement(ctx *GrantSequenceStatementContext) {}

// ExitGrantSequenceStatement is called when production grantSequenceStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantSequenceStatement(ctx *GrantSequenceStatementContext) {}

// EnterGrantSystemStatement is called when production grantSystemStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantSystemStatement(ctx *GrantSystemStatementContext) {}

// ExitGrantSystemStatement is called when production grantSystemStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantSystemStatement(ctx *GrantSystemStatementContext) {}

// EnterGrantTableStatement is called when production grantTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantTableStatement(ctx *GrantTableStatementContext) {}

// ExitGrantTableStatement is called when production grantTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantTableStatement(ctx *GrantTableStatementContext) {}

// EnterGrantTypeOrJarStatement is called when production grantTypeOrJarStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantTypeOrJarStatement(ctx *GrantTypeOrJarStatementContext) {
}

// ExitGrantTypeOrJarStatement is called when production grantTypeOrJarStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantTypeOrJarStatement(ctx *GrantTypeOrJarStatementContext) {
}

// EnterGrantVariableStatement is called when production grantVariableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantVariableStatement(ctx *GrantVariableStatementContext) {}

// ExitGrantVariableStatement is called when production grantVariableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantVariableStatement(ctx *GrantVariableStatementContext) {}

// EnterGrantUseOfStatement is called when production grantUseOfStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantUseOfStatement(ctx *GrantUseOfStatementContext) {}

// ExitGrantUseOfStatement is called when production grantUseOfStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantUseOfStatement(ctx *GrantUseOfStatementContext) {}

// EnterRevokeCollectionStatement is called when production revokeCollectionStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeCollectionStatement(ctx *RevokeCollectionStatementContext) {
}

// ExitRevokeCollectionStatement is called when production revokeCollectionStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeCollectionStatement(ctx *RevokeCollectionStatementContext) {
}

// EnterRevokeDatabaseStatement is called when production revokeDatabaseStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeDatabaseStatement(ctx *RevokeDatabaseStatementContext) {
}

// ExitRevokeDatabaseStatement is called when production revokeDatabaseStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeDatabaseStatement(ctx *RevokeDatabaseStatementContext) {
}

// EnterRevokeFunctionOrProcedureStatement is called when production revokeFunctionOrProcedureStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeFunctionOrProcedureStatement(ctx *RevokeFunctionOrProcedureStatementContext) {
}

// ExitRevokeFunctionOrProcedureStatement is called when production revokeFunctionOrProcedureStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeFunctionOrProcedureStatement(ctx *RevokeFunctionOrProcedureStatementContext) {
}

// EnterRevokePackageStatement is called when production revokePackageStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokePackageStatement(ctx *RevokePackageStatementContext) {}

// ExitRevokePackageStatement is called when production revokePackageStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokePackageStatement(ctx *RevokePackageStatementContext) {}

// EnterRevokePlanStatement is called when production revokePlanStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokePlanStatement(ctx *RevokePlanStatementContext) {}

// ExitRevokePlanStatement is called when production revokePlanStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokePlanStatement(ctx *RevokePlanStatementContext) {}

// EnterRevokeSchemaStatement is called when production revokeSchemaStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeSchemaStatement(ctx *RevokeSchemaStatementContext) {}

// ExitRevokeSchemaStatement is called when production revokeSchemaStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeSchemaStatement(ctx *RevokeSchemaStatementContext) {}

// EnterRevokeSequenceStatement is called when production revokeSequenceStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeSequenceStatement(ctx *RevokeSequenceStatementContext) {
}

// ExitRevokeSequenceStatement is called when production revokeSequenceStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeSequenceStatement(ctx *RevokeSequenceStatementContext) {
}

// EnterRevokeSystemStatement is called when production revokeSystemStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeSystemStatement(ctx *RevokeSystemStatementContext) {}

// ExitRevokeSystemStatement is called when production revokeSystemStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeSystemStatement(ctx *RevokeSystemStatementContext) {}

// EnterRevokeTableStatement is called when production revokeTableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeTableStatement(ctx *RevokeTableStatementContext) {}

// ExitRevokeTableStatement is called when production revokeTableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeTableStatement(ctx *RevokeTableStatementContext) {}

// EnterRevokeTypeOrJarStatement is called when production revokeTypeOrJarStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeTypeOrJarStatement(ctx *RevokeTypeOrJarStatementContext) {
}

// ExitRevokeTypeOrJarStatement is called when production revokeTypeOrJarStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeTypeOrJarStatement(ctx *RevokeTypeOrJarStatementContext) {
}

// EnterRevokeVariableStatement is called when production revokeVariableStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeVariableStatement(ctx *RevokeVariableStatementContext) {
}

// ExitRevokeVariableStatement is called when production revokeVariableStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeVariableStatement(ctx *RevokeVariableStatementContext) {
}

// EnterRevokeUseOfStatement is called when production revokeUseOfStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeUseOfStatement(ctx *RevokeUseOfStatementContext) {}

// ExitRevokeUseOfStatement is called when production revokeUseOfStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeUseOfStatement(ctx *RevokeUseOfStatementContext) {}

// EnterGrantUseOfTarget is called when production grantUseOfTarget is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantUseOfTarget(ctx *GrantUseOfTargetContext) {}

// ExitGrantUseOfTarget is called when production grantUseOfTarget is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantUseOfTarget(ctx *GrantUseOfTargetContext) {}

// EnterGrantVariableAuthority is called when production grantVariableAuthority is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantVariableAuthority(ctx *GrantVariableAuthorityContext) {}

// ExitGrantVariableAuthority is called when production grantVariableAuthority is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantVariableAuthority(ctx *GrantVariableAuthorityContext) {}

// EnterGrantTableAuthority is called when production grantTableAuthority is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantTableAuthority(ctx *GrantTableAuthorityContext) {}

// ExitGrantTableAuthority is called when production grantTableAuthority is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantTableAuthority(ctx *GrantTableAuthorityContext) {}

// EnterGrantSystemAuthority is called when production grantSystemAuthority is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantSystemAuthority(ctx *GrantSystemAuthorityContext) {}

// ExitGrantSystemAuthority is called when production grantSystemAuthority is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantSystemAuthority(ctx *GrantSystemAuthorityContext) {}

// EnterGrantSequenceAuthority is called when production grantSequenceAuthority is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantSequenceAuthority(ctx *GrantSequenceAuthorityContext) {}

// ExitGrantSequenceAuthority is called when production grantSequenceAuthority is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantSequenceAuthority(ctx *GrantSequenceAuthorityContext) {}

// EnterGrantSchemaAuthority is called when production grantSchemaAuthority is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantSchemaAuthority(ctx *GrantSchemaAuthorityContext) {}

// ExitGrantSchemaAuthority is called when production grantSchemaAuthority is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantSchemaAuthority(ctx *GrantSchemaAuthorityContext) {}

// EnterGrantPlanAuthority is called when production grantPlanAuthority is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantPlanAuthority(ctx *GrantPlanAuthorityContext) {}

// ExitGrantPlanAuthority is called when production grantPlanAuthority is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantPlanAuthority(ctx *GrantPlanAuthorityContext) {}

// EnterGrantPackageAuthority is called when production grantPackageAuthority is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantPackageAuthority(ctx *GrantPackageAuthorityContext) {}

// ExitGrantPackageAuthority is called when production grantPackageAuthority is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantPackageAuthority(ctx *GrantPackageAuthorityContext) {}

// EnterPackageSpecification is called when production packageSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterPackageSpecification(ctx *PackageSpecificationContext) {}

// ExitPackageSpecification is called when production packageSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitPackageSpecification(ctx *PackageSpecificationContext) {}

// EnterFunctionSpecification is called when production functionSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionSpecification(ctx *FunctionSpecificationContext) {}

// ExitFunctionSpecification is called when production functionSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionSpecification(ctx *FunctionSpecificationContext) {}

// EnterGrantee is called when production grantee is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantee(ctx *GranteeContext) {}

// ExitGrantee is called when production grantee is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantee(ctx *GranteeContext) {}

// EnterWithGrantOption is called when production withGrantOption is entered.
func (s *BaseDB2zSQLParserListener) EnterWithGrantOption(ctx *WithGrantOptionContext) {}

// ExitWithGrantOption is called when production withGrantOption is exited.
func (s *BaseDB2zSQLParserListener) ExitWithGrantOption(ctx *WithGrantOptionContext) {}

// EnterRevokeByOption is called when production revokeByOption is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeByOption(ctx *RevokeByOptionContext) {}

// ExitRevokeByOption is called when production revokeByOption is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeByOption(ctx *RevokeByOptionContext) {}

// EnterRevokeDependentPrivilegesOption is called when production revokeDependentPrivilegesOption is entered.
func (s *BaseDB2zSQLParserListener) EnterRevokeDependentPrivilegesOption(ctx *RevokeDependentPrivilegesOptionContext) {
}

// ExitRevokeDependentPrivilegesOption is called when production revokeDependentPrivilegesOption is exited.
func (s *BaseDB2zSQLParserListener) ExitRevokeDependentPrivilegesOption(ctx *RevokeDependentPrivilegesOptionContext) {
}

// EnterGrantDatabaseAuthority is called when production grantDatabaseAuthority is entered.
func (s *BaseDB2zSQLParserListener) EnterGrantDatabaseAuthority(ctx *GrantDatabaseAuthorityContext) {}

// ExitGrantDatabaseAuthority is called when production grantDatabaseAuthority is exited.
func (s *BaseDB2zSQLParserListener) ExitGrantDatabaseAuthority(ctx *GrantDatabaseAuthorityContext) {}

// EnterStatementInformation is called when production statementInformation is entered.
func (s *BaseDB2zSQLParserListener) EnterStatementInformation(ctx *StatementInformationContext) {}

// ExitStatementInformation is called when production statementInformation is exited.
func (s *BaseDB2zSQLParserListener) ExitStatementInformation(ctx *StatementInformationContext) {}

// EnterStatementInformationVariableEquate is called when production statementInformationVariableEquate is entered.
func (s *BaseDB2zSQLParserListener) EnterStatementInformationVariableEquate(ctx *StatementInformationVariableEquateContext) {
}

// ExitStatementInformationVariableEquate is called when production statementInformationVariableEquate is exited.
func (s *BaseDB2zSQLParserListener) ExitStatementInformationVariableEquate(ctx *StatementInformationVariableEquateContext) {
}

// EnterStatementInformationItemName is called when production statementInformationItemName is entered.
func (s *BaseDB2zSQLParserListener) EnterStatementInformationItemName(ctx *StatementInformationItemNameContext) {
}

// ExitStatementInformationItemName is called when production statementInformationItemName is exited.
func (s *BaseDB2zSQLParserListener) ExitStatementInformationItemName(ctx *StatementInformationItemNameContext) {
}

// EnterConditionInformation is called when production conditionInformation is entered.
func (s *BaseDB2zSQLParserListener) EnterConditionInformation(ctx *ConditionInformationContext) {}

// ExitConditionInformation is called when production conditionInformation is exited.
func (s *BaseDB2zSQLParserListener) ExitConditionInformation(ctx *ConditionInformationContext) {}

// EnterConditionInformationVariableEquate is called when production conditionInformationVariableEquate is entered.
func (s *BaseDB2zSQLParserListener) EnterConditionInformationVariableEquate(ctx *ConditionInformationVariableEquateContext) {
}

// ExitConditionInformationVariableEquate is called when production conditionInformationVariableEquate is exited.
func (s *BaseDB2zSQLParserListener) ExitConditionInformationVariableEquate(ctx *ConditionInformationVariableEquateContext) {
}

// EnterConditionInformationItemName is called when production conditionInformationItemName is entered.
func (s *BaseDB2zSQLParserListener) EnterConditionInformationItemName(ctx *ConditionInformationItemNameContext) {
}

// ExitConditionInformationItemName is called when production conditionInformationItemName is exited.
func (s *BaseDB2zSQLParserListener) ExitConditionInformationItemName(ctx *ConditionInformationItemNameContext) {
}

// EnterConnectionInformationItemName is called when production connectionInformationItemName is entered.
func (s *BaseDB2zSQLParserListener) EnterConnectionInformationItemName(ctx *ConnectionInformationItemNameContext) {
}

// ExitConnectionInformationItemName is called when production connectionInformationItemName is exited.
func (s *BaseDB2zSQLParserListener) ExitConnectionInformationItemName(ctx *ConnectionInformationItemNameContext) {
}

// EnterCombinedInformation is called when production combinedInformation is entered.
func (s *BaseDB2zSQLParserListener) EnterCombinedInformation(ctx *CombinedInformationContext) {}

// ExitCombinedInformation is called when production combinedInformation is exited.
func (s *BaseDB2zSQLParserListener) ExitCombinedInformation(ctx *CombinedInformationContext) {}

// EnterCombinedInformationOption is called when production combinedInformationOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCombinedInformationOption(ctx *CombinedInformationOptionContext) {
}

// ExitCombinedInformationOption is called when production combinedInformationOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCombinedInformationOption(ctx *CombinedInformationOptionContext) {
}

// EnterFetchOrientation is called when production fetchOrientation is entered.
func (s *BaseDB2zSQLParserListener) EnterFetchOrientation(ctx *FetchOrientationContext) {}

// ExitFetchOrientation is called when production fetchOrientation is exited.
func (s *BaseDB2zSQLParserListener) ExitFetchOrientation(ctx *FetchOrientationContext) {}

// EnterRowPositioned is called when production rowPositioned is entered.
func (s *BaseDB2zSQLParserListener) EnterRowPositioned(ctx *RowPositionedContext) {}

// ExitRowPositioned is called when production rowPositioned is exited.
func (s *BaseDB2zSQLParserListener) ExitRowPositioned(ctx *RowPositionedContext) {}

// EnterRowsetPositioned is called when production rowsetPositioned is entered.
func (s *BaseDB2zSQLParserListener) EnterRowsetPositioned(ctx *RowsetPositionedContext) {}

// ExitRowsetPositioned is called when production rowsetPositioned is exited.
func (s *BaseDB2zSQLParserListener) ExitRowsetPositioned(ctx *RowsetPositionedContext) {}

// EnterSingleRowFetch is called when production singleRowFetch is entered.
func (s *BaseDB2zSQLParserListener) EnterSingleRowFetch(ctx *SingleRowFetchContext) {}

// ExitSingleRowFetch is called when production singleRowFetch is exited.
func (s *BaseDB2zSQLParserListener) ExitSingleRowFetch(ctx *SingleRowFetchContext) {}

// EnterFetchTargetVariable is called when production fetchTargetVariable is entered.
func (s *BaseDB2zSQLParserListener) EnterFetchTargetVariable(ctx *FetchTargetVariableContext) {}

// ExitFetchTargetVariable is called when production fetchTargetVariable is exited.
func (s *BaseDB2zSQLParserListener) ExitFetchTargetVariable(ctx *FetchTargetVariableContext) {}

// EnterMultipleRowFetch is called when production multipleRowFetch is entered.
func (s *BaseDB2zSQLParserListener) EnterMultipleRowFetch(ctx *MultipleRowFetchContext) {}

// ExitMultipleRowFetch is called when production multipleRowFetch is exited.
func (s *BaseDB2zSQLParserListener) ExitMultipleRowFetch(ctx *MultipleRowFetchContext) {}

// EnterMultipleRowFetchForClause is called when production multipleRowFetchForClause is entered.
func (s *BaseDB2zSQLParserListener) EnterMultipleRowFetchForClause(ctx *MultipleRowFetchForClauseContext) {
}

// ExitMultipleRowFetchForClause is called when production multipleRowFetchForClause is exited.
func (s *BaseDB2zSQLParserListener) ExitMultipleRowFetchForClause(ctx *MultipleRowFetchForClauseContext) {
}

// EnterMultipleRowFetchIntoClause is called when production multipleRowFetchIntoClause is entered.
func (s *BaseDB2zSQLParserListener) EnterMultipleRowFetchIntoClause(ctx *MultipleRowFetchIntoClauseContext) {
}

// ExitMultipleRowFetchIntoClause is called when production multipleRowFetchIntoClause is exited.
func (s *BaseDB2zSQLParserListener) ExitMultipleRowFetchIntoClause(ctx *MultipleRowFetchIntoClauseContext) {
}

// EnterExplainPlanClause is called when production explainPlanClause is entered.
func (s *BaseDB2zSQLParserListener) EnterExplainPlanClause(ctx *ExplainPlanClauseContext) {}

// ExitExplainPlanClause is called when production explainPlanClause is exited.
func (s *BaseDB2zSQLParserListener) ExitExplainPlanClause(ctx *ExplainPlanClauseContext) {}

// EnterExplainStmtcacheClause is called when production explainStmtcacheClause is entered.
func (s *BaseDB2zSQLParserListener) EnterExplainStmtcacheClause(ctx *ExplainStmtcacheClauseContext) {}

// ExitExplainStmtcacheClause is called when production explainStmtcacheClause is exited.
func (s *BaseDB2zSQLParserListener) ExitExplainStmtcacheClause(ctx *ExplainStmtcacheClauseContext) {}

// EnterExplainPackageClause is called when production explainPackageClause is entered.
func (s *BaseDB2zSQLParserListener) EnterExplainPackageClause(ctx *ExplainPackageClauseContext) {}

// ExitExplainPackageClause is called when production explainPackageClause is exited.
func (s *BaseDB2zSQLParserListener) ExitExplainPackageClause(ctx *ExplainPackageClauseContext) {}

// EnterExplainStabilizedDynamicQueryClause is called when production explainStabilizedDynamicQueryClause is entered.
func (s *BaseDB2zSQLParserListener) EnterExplainStabilizedDynamicQueryClause(ctx *ExplainStabilizedDynamicQueryClauseContext) {
}

// ExitExplainStabilizedDynamicQueryClause is called when production explainStabilizedDynamicQueryClause is exited.
func (s *BaseDB2zSQLParserListener) ExitExplainStabilizedDynamicQueryClause(ctx *ExplainStabilizedDynamicQueryClauseContext) {
}

// EnterPackageScopeSpecification is called when production packageScopeSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterPackageScopeSpecification(ctx *PackageScopeSpecificationContext) {
}

// ExitPackageScopeSpecification is called when production packageScopeSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitPackageScopeSpecification(ctx *PackageScopeSpecificationContext) {
}

// EnterCollectionName is called when production collectionName is entered.
func (s *BaseDB2zSQLParserListener) EnterCollectionName(ctx *CollectionNameContext) {}

// ExitCollectionName is called when production collectionName is exited.
func (s *BaseDB2zSQLParserListener) ExitCollectionName(ctx *CollectionNameContext) {}

// EnterPackageScopePackageName is called when production packageScopePackageName is entered.
func (s *BaseDB2zSQLParserListener) EnterPackageScopePackageName(ctx *PackageScopePackageNameContext) {
}

// ExitPackageScopePackageName is called when production packageScopePackageName is exited.
func (s *BaseDB2zSQLParserListener) ExitPackageScopePackageName(ctx *PackageScopePackageNameContext) {
}

// EnterVersionName is called when production versionName is entered.
func (s *BaseDB2zSQLParserListener) EnterVersionName(ctx *VersionNameContext) {}

// ExitVersionName is called when production versionName is exited.
func (s *BaseDB2zSQLParserListener) ExitVersionName(ctx *VersionNameContext) {}

// EnterSourceRowData is called when production sourceRowData is entered.
func (s *BaseDB2zSQLParserListener) EnterSourceRowData(ctx *SourceRowDataContext) {}

// ExitSourceRowData is called when production sourceRowData is exited.
func (s *BaseDB2zSQLParserListener) ExitSourceRowData(ctx *SourceRowDataContext) {}

// EnterAliasDesignation is called when production aliasDesignation is entered.
func (s *BaseDB2zSQLParserListener) EnterAliasDesignation(ctx *AliasDesignationContext) {}

// ExitAliasDesignation is called when production aliasDesignation is exited.
func (s *BaseDB2zSQLParserListener) ExitAliasDesignation(ctx *AliasDesignationContext) {}

// EnterDropDatabaseClause is called when production dropDatabaseClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropDatabaseClause(ctx *DropDatabaseClauseContext) {}

// ExitDropDatabaseClause is called when production dropDatabaseClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropDatabaseClause(ctx *DropDatabaseClauseContext) {}

// EnterDropFunctionClause is called when production dropFunctionClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropFunctionClause(ctx *DropFunctionClauseContext) {}

// ExitDropFunctionClause is called when production dropFunctionClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropFunctionClause(ctx *DropFunctionClauseContext) {}

// EnterDropIndexClause is called when production dropIndexClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropIndexClause(ctx *DropIndexClauseContext) {}

// ExitDropIndexClause is called when production dropIndexClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropIndexClause(ctx *DropIndexClauseContext) {}

// EnterDropMaskClause is called when production dropMaskClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropMaskClause(ctx *DropMaskClauseContext) {}

// ExitDropMaskClause is called when production dropMaskClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropMaskClause(ctx *DropMaskClauseContext) {}

// EnterDropPackageClause is called when production dropPackageClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropPackageClause(ctx *DropPackageClauseContext) {}

// ExitDropPackageClause is called when production dropPackageClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropPackageClause(ctx *DropPackageClauseContext) {}

// EnterDropPermissionClause is called when production dropPermissionClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropPermissionClause(ctx *DropPermissionClauseContext) {}

// ExitDropPermissionClause is called when production dropPermissionClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropPermissionClause(ctx *DropPermissionClauseContext) {}

// EnterDropProcedureClause is called when production dropProcedureClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropProcedureClause(ctx *DropProcedureClauseContext) {}

// ExitDropProcedureClause is called when production dropProcedureClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropProcedureClause(ctx *DropProcedureClauseContext) {}

// EnterDropRoleClause is called when production dropRoleClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropRoleClause(ctx *DropRoleClauseContext) {}

// ExitDropRoleClause is called when production dropRoleClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropRoleClause(ctx *DropRoleClauseContext) {}

// EnterDropSequenceClause is called when production dropSequenceClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropSequenceClause(ctx *DropSequenceClauseContext) {}

// ExitDropSequenceClause is called when production dropSequenceClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropSequenceClause(ctx *DropSequenceClauseContext) {}

// EnterDropStogroupClause is called when production dropStogroupClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropStogroupClause(ctx *DropStogroupClauseContext) {}

// ExitDropStogroupClause is called when production dropStogroupClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropStogroupClause(ctx *DropStogroupClauseContext) {}

// EnterDropSynonymClause is called when production dropSynonymClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropSynonymClause(ctx *DropSynonymClauseContext) {}

// ExitDropSynonymClause is called when production dropSynonymClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropSynonymClause(ctx *DropSynonymClauseContext) {}

// EnterDropTableClause is called when production dropTableClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropTableClause(ctx *DropTableClauseContext) {}

// ExitDropTableClause is called when production dropTableClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropTableClause(ctx *DropTableClauseContext) {}

// EnterDropTablespaceClause is called when production dropTablespaceClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropTablespaceClause(ctx *DropTablespaceClauseContext) {}

// ExitDropTablespaceClause is called when production dropTablespaceClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropTablespaceClause(ctx *DropTablespaceClauseContext) {}

// EnterDropTriggerClause is called when production dropTriggerClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropTriggerClause(ctx *DropTriggerClauseContext) {}

// ExitDropTriggerClause is called when production dropTriggerClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropTriggerClause(ctx *DropTriggerClauseContext) {}

// EnterDropTrustedContextClause is called when production dropTrustedContextClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropTrustedContextClause(ctx *DropTrustedContextClauseContext) {
}

// ExitDropTrustedContextClause is called when production dropTrustedContextClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropTrustedContextClause(ctx *DropTrustedContextClauseContext) {
}

// EnterDropTypeClause is called when production dropTypeClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropTypeClause(ctx *DropTypeClauseContext) {}

// ExitDropTypeClause is called when production dropTypeClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropTypeClause(ctx *DropTypeClauseContext) {}

// EnterDropVariableClause is called when production dropVariableClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropVariableClause(ctx *DropVariableClauseContext) {}

// ExitDropVariableClause is called when production dropVariableClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropVariableClause(ctx *DropVariableClauseContext) {}

// EnterDropViewClause is called when production dropViewClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDropViewClause(ctx *DropViewClauseContext) {}

// ExitDropViewClause is called when production dropViewClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDropViewClause(ctx *DropViewClauseContext) {}

// EnterPackageDesignator is called when production packageDesignator is entered.
func (s *BaseDB2zSQLParserListener) EnterPackageDesignator(ctx *PackageDesignatorContext) {}

// ExitPackageDesignator is called when production packageDesignator is exited.
func (s *BaseDB2zSQLParserListener) ExitPackageDesignator(ctx *PackageDesignatorContext) {}

// EnterDescribeUsingOption is called when production describeUsingOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDescribeUsingOption(ctx *DescribeUsingOptionContext) {}

// ExitDescribeUsingOption is called when production describeUsingOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDescribeUsingOption(ctx *DescribeUsingOptionContext) {}

// EnterDeclareGlobalTemporaryTableLikeClause is called when production declareGlobalTemporaryTableLikeClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDeclareGlobalTemporaryTableLikeClause(ctx *DeclareGlobalTemporaryTableLikeClauseContext) {
}

// ExitDeclareGlobalTemporaryTableLikeClause is called when production declareGlobalTemporaryTableLikeClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDeclareGlobalTemporaryTableLikeClause(ctx *DeclareGlobalTemporaryTableLikeClauseContext) {
}

// EnterOnCommitClause is called when production onCommitClause is entered.
func (s *BaseDB2zSQLParserListener) EnterOnCommitClause(ctx *OnCommitClauseContext) {}

// ExitOnCommitClause is called when production onCommitClause is exited.
func (s *BaseDB2zSQLParserListener) ExitOnCommitClause(ctx *OnCommitClauseContext) {}

// EnterLoggedWithRollbackClause is called when production loggedWithRollbackClause is entered.
func (s *BaseDB2zSQLParserListener) EnterLoggedWithRollbackClause(ctx *LoggedWithRollbackClauseContext) {
}

// ExitLoggedWithRollbackClause is called when production loggedWithRollbackClause is exited.
func (s *BaseDB2zSQLParserListener) ExitLoggedWithRollbackClause(ctx *LoggedWithRollbackClauseContext) {
}

// EnterCreateViewCheckOptionClause is called when production createViewCheckOptionClause is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateViewCheckOptionClause(ctx *CreateViewCheckOptionClauseContext) {
}

// ExitCreateViewCheckOptionClause is called when production createViewCheckOptionClause is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateViewCheckOptionClause(ctx *CreateViewCheckOptionClauseContext) {
}

// EnterTrustedContextDefaultRoleClause is called when production trustedContextDefaultRoleClause is entered.
func (s *BaseDB2zSQLParserListener) EnterTrustedContextDefaultRoleClause(ctx *TrustedContextDefaultRoleClauseContext) {
}

// ExitTrustedContextDefaultRoleClause is called when production trustedContextDefaultRoleClause is exited.
func (s *BaseDB2zSQLParserListener) ExitTrustedContextDefaultRoleClause(ctx *TrustedContextDefaultRoleClauseContext) {
}

// EnterTrustedContextEnableDisableClause is called when production trustedContextEnableDisableClause is entered.
func (s *BaseDB2zSQLParserListener) EnterTrustedContextEnableDisableClause(ctx *TrustedContextEnableDisableClauseContext) {
}

// ExitTrustedContextEnableDisableClause is called when production trustedContextEnableDisableClause is exited.
func (s *BaseDB2zSQLParserListener) ExitTrustedContextEnableDisableClause(ctx *TrustedContextEnableDisableClauseContext) {
}

// EnterTrustedContextDefaultSecurityLabelClause is called when production trustedContextDefaultSecurityLabelClause is entered.
func (s *BaseDB2zSQLParserListener) EnterTrustedContextDefaultSecurityLabelClause(ctx *TrustedContextDefaultSecurityLabelClauseContext) {
}

// ExitTrustedContextDefaultSecurityLabelClause is called when production trustedContextDefaultSecurityLabelClause is exited.
func (s *BaseDB2zSQLParserListener) ExitTrustedContextDefaultSecurityLabelClause(ctx *TrustedContextDefaultSecurityLabelClauseContext) {
}

// EnterTrustedContextAttributesClause is called when production trustedContextAttributesClause is entered.
func (s *BaseDB2zSQLParserListener) EnterTrustedContextAttributesClause(ctx *TrustedContextAttributesClauseContext) {
}

// ExitTrustedContextAttributesClause is called when production trustedContextAttributesClause is exited.
func (s *BaseDB2zSQLParserListener) ExitTrustedContextAttributesClause(ctx *TrustedContextAttributesClauseContext) {
}

// EnterTrustedContextWithUseForClause is called when production trustedContextWithUseForClause is entered.
func (s *BaseDB2zSQLParserListener) EnterTrustedContextWithUseForClause(ctx *TrustedContextWithUseForClauseContext) {
}

// ExitTrustedContextWithUseForClause is called when production trustedContextWithUseForClause is exited.
func (s *BaseDB2zSQLParserListener) ExitTrustedContextWithUseForClause(ctx *TrustedContextWithUseForClauseContext) {
}

// EnterTrustedContextAttribute1 is called when production trustedContextAttribute1 is entered.
func (s *BaseDB2zSQLParserListener) EnterTrustedContextAttribute1(ctx *TrustedContextAttribute1Context) {
}

// ExitTrustedContextAttribute1 is called when production trustedContextAttribute1 is exited.
func (s *BaseDB2zSQLParserListener) ExitTrustedContextAttribute1(ctx *TrustedContextAttribute1Context) {
}

// EnterTrustedContextAttribute2 is called when production trustedContextAttribute2 is entered.
func (s *BaseDB2zSQLParserListener) EnterTrustedContextAttribute2(ctx *TrustedContextAttribute2Context) {
}

// ExitTrustedContextAttribute2 is called when production trustedContextAttribute2 is exited.
func (s *BaseDB2zSQLParserListener) ExitTrustedContextAttribute2(ctx *TrustedContextAttribute2Context) {
}

// EnterTrustedContextUseFor is called when production trustedContextUseFor is entered.
func (s *BaseDB2zSQLParserListener) EnterTrustedContextUseFor(ctx *TrustedContextUseForContext) {}

// ExitTrustedContextUseFor is called when production trustedContextUseFor is exited.
func (s *BaseDB2zSQLParserListener) ExitTrustedContextUseFor(ctx *TrustedContextUseForContext) {}

// EnterUserOptions is called when production userOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterUserOptions(ctx *UserOptionsContext) {}

// ExitUserOptions is called when production userOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitUserOptions(ctx *UserOptionsContext) {}

// EnterTriggerDefinition is called when production triggerDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterTriggerDefinition(ctx *TriggerDefinitionContext) {}

// ExitTriggerDefinition is called when production triggerDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitTriggerDefinition(ctx *TriggerDefinitionContext) {}

// EnterTriggerActivationTime is called when production triggerActivationTime is entered.
func (s *BaseDB2zSQLParserListener) EnterTriggerActivationTime(ctx *TriggerActivationTimeContext) {}

// ExitTriggerActivationTime is called when production triggerActivationTime is exited.
func (s *BaseDB2zSQLParserListener) ExitTriggerActivationTime(ctx *TriggerActivationTimeContext) {}

// EnterTriggerEvent is called when production triggerEvent is entered.
func (s *BaseDB2zSQLParserListener) EnterTriggerEvent(ctx *TriggerEventContext) {}

// ExitTriggerEvent is called when production triggerEvent is exited.
func (s *BaseDB2zSQLParserListener) ExitTriggerEvent(ctx *TriggerEventContext) {}

// EnterTriggerGranularity is called when production triggerGranularity is entered.
func (s *BaseDB2zSQLParserListener) EnterTriggerGranularity(ctx *TriggerGranularityContext) {}

// ExitTriggerGranularity is called when production triggerGranularity is exited.
func (s *BaseDB2zSQLParserListener) ExitTriggerGranularity(ctx *TriggerGranularityContext) {}

// EnterTriggeredAction is called when production triggeredAction is entered.
func (s *BaseDB2zSQLParserListener) EnterTriggeredAction(ctx *TriggeredActionContext) {}

// ExitTriggeredAction is called when production triggeredAction is exited.
func (s *BaseDB2zSQLParserListener) ExitTriggeredAction(ctx *TriggeredActionContext) {}

// EnterSqlTriggerBody is called when production sqlTriggerBody is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlTriggerBody(ctx *SqlTriggerBodyContext) {}

// ExitSqlTriggerBody is called when production sqlTriggerBody is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlTriggerBody(ctx *SqlTriggerBodyContext) {}

// EnterTriggeredSqlStatement is called when production triggeredSqlStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterTriggeredSqlStatement(ctx *TriggeredSqlStatementContext) {}

// ExitTriggeredSqlStatement is called when production triggeredSqlStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitTriggeredSqlStatement(ctx *TriggeredSqlStatementContext) {}

// EnterTriggerDefinitionOption is called when production triggerDefinitionOption is entered.
func (s *BaseDB2zSQLParserListener) EnterTriggerDefinitionOption(ctx *TriggerDefinitionOptionContext) {
}

// ExitTriggerDefinitionOption is called when production triggerDefinitionOption is exited.
func (s *BaseDB2zSQLParserListener) ExitTriggerDefinitionOption(ctx *TriggerDefinitionOptionContext) {
}

// EnterCreateTableInClause is called when production createTableInClause is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTableInClause(ctx *CreateTableInClauseContext) {}

// ExitCreateTableInClause is called when production createTableInClause is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTableInClause(ctx *CreateTableInClauseContext) {}

// EnterCustomVolatileClause is called when production customVolatileClause is entered.
func (s *BaseDB2zSQLParserListener) EnterCustomVolatileClause(ctx *CustomVolatileClauseContext) {}

// ExitCustomVolatileClause is called when production customVolatileClause is exited.
func (s *BaseDB2zSQLParserListener) ExitCustomVolatileClause(ctx *CustomVolatileClauseContext) {}

// EnterCreateTableColumnDefinition is called when production createTableColumnDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTableColumnDefinition(ctx *CreateTableColumnDefinitionContext) {
}

// ExitCreateTableColumnDefinition is called when production createTableColumnDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTableColumnDefinition(ctx *CreateTableColumnDefinitionContext) {
}

// EnterEditprocClause is called when production editprocClause is entered.
func (s *BaseDB2zSQLParserListener) EnterEditprocClause(ctx *EditprocClauseContext) {}

// ExitEditprocClause is called when production editprocClause is exited.
func (s *BaseDB2zSQLParserListener) ExitEditprocClause(ctx *EditprocClauseContext) {}

// EnterValidprocClause is called when production validprocClause is entered.
func (s *BaseDB2zSQLParserListener) EnterValidprocClause(ctx *ValidprocClauseContext) {}

// ExitValidprocClause is called when production validprocClause is exited.
func (s *BaseDB2zSQLParserListener) ExitValidprocClause(ctx *ValidprocClauseContext) {}

// EnterAuditClause is called when production auditClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAuditClause(ctx *AuditClauseContext) {}

// ExitAuditClause is called when production auditClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAuditClause(ctx *AuditClauseContext) {}

// EnterObidClause is called when production obidClause is entered.
func (s *BaseDB2zSQLParserListener) EnterObidClause(ctx *ObidClauseContext) {}

// ExitObidClause is called when production obidClause is exited.
func (s *BaseDB2zSQLParserListener) ExitObidClause(ctx *ObidClauseContext) {}

// EnterDataCaptureClause is called when production dataCaptureClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDataCaptureClause(ctx *DataCaptureClauseContext) {}

// ExitDataCaptureClause is called when production dataCaptureClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDataCaptureClause(ctx *DataCaptureClauseContext) {}

// EnterRestrictOnDropClause is called when production restrictOnDropClause is entered.
func (s *BaseDB2zSQLParserListener) EnterRestrictOnDropClause(ctx *RestrictOnDropClauseContext) {}

// ExitRestrictOnDropClause is called when production restrictOnDropClause is exited.
func (s *BaseDB2zSQLParserListener) ExitRestrictOnDropClause(ctx *RestrictOnDropClauseContext) {}

// EnterCcsidClause1 is called when production ccsidClause1 is entered.
func (s *BaseDB2zSQLParserListener) EnterCcsidClause1(ctx *CcsidClause1Context) {}

// ExitCcsidClause1 is called when production ccsidClause1 is exited.
func (s *BaseDB2zSQLParserListener) ExitCcsidClause1(ctx *CcsidClause1Context) {}

// EnterCcsidClause2 is called when production ccsidClause2 is entered.
func (s *BaseDB2zSQLParserListener) EnterCcsidClause2(ctx *CcsidClause2Context) {}

// ExitCcsidClause2 is called when production ccsidClause2 is exited.
func (s *BaseDB2zSQLParserListener) ExitCcsidClause2(ctx *CcsidClause2Context) {}

// EnterCardinalityClause is called when production cardinalityClause is entered.
func (s *BaseDB2zSQLParserListener) EnterCardinalityClause(ctx *CardinalityClauseContext) {}

// ExitCardinalityClause is called when production cardinalityClause is exited.
func (s *BaseDB2zSQLParserListener) ExitCardinalityClause(ctx *CardinalityClauseContext) {}

// EnterAppendClause is called when production appendClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAppendClause(ctx *AppendClauseContext) {}

// ExitAppendClause is called when production appendClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAppendClause(ctx *AppendClauseContext) {}

// EnterMemberClause is called when production memberClause is entered.
func (s *BaseDB2zSQLParserListener) EnterMemberClause(ctx *MemberClauseContext) {}

// ExitMemberClause is called when production memberClause is exited.
func (s *BaseDB2zSQLParserListener) ExitMemberClause(ctx *MemberClauseContext) {}

// EnterTrackmodClause is called when production trackmodClause is entered.
func (s *BaseDB2zSQLParserListener) EnterTrackmodClause(ctx *TrackmodClauseContext) {}

// ExitTrackmodClause is called when production trackmodClause is exited.
func (s *BaseDB2zSQLParserListener) ExitTrackmodClause(ctx *TrackmodClauseContext) {}

// EnterPagenumClause is called when production pagenumClause is entered.
func (s *BaseDB2zSQLParserListener) EnterPagenumClause(ctx *PagenumClauseContext) {}

// ExitPagenumClause is called when production pagenumClause is exited.
func (s *BaseDB2zSQLParserListener) ExitPagenumClause(ctx *PagenumClauseContext) {}

// EnterFieldprocClause is called when production fieldprocClause is entered.
func (s *BaseDB2zSQLParserListener) EnterFieldprocClause(ctx *FieldprocClauseContext) {}

// ExitFieldprocClause is called when production fieldprocClause is exited.
func (s *BaseDB2zSQLParserListener) ExitFieldprocClause(ctx *FieldprocClauseContext) {}

// EnterAsSecurityLabelClause is called when production asSecurityLabelClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAsSecurityLabelClause(ctx *AsSecurityLabelClauseContext) {}

// ExitAsSecurityLabelClause is called when production asSecurityLabelClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAsSecurityLabelClause(ctx *AsSecurityLabelClauseContext) {}

// EnterImplicitlyHiddenClause is called when production implicitlyHiddenClause is entered.
func (s *BaseDB2zSQLParserListener) EnterImplicitlyHiddenClause(ctx *ImplicitlyHiddenClauseContext) {}

// ExitImplicitlyHiddenClause is called when production implicitlyHiddenClause is exited.
func (s *BaseDB2zSQLParserListener) ExitImplicitlyHiddenClause(ctx *ImplicitlyHiddenClauseContext) {}

// EnterInlineLengthClause is called when production inlineLengthClause is entered.
func (s *BaseDB2zSQLParserListener) EnterInlineLengthClause(ctx *InlineLengthClauseContext) {}

// ExitInlineLengthClause is called when production inlineLengthClause is exited.
func (s *BaseDB2zSQLParserListener) ExitInlineLengthClause(ctx *InlineLengthClauseContext) {}

// EnterCopyOptions is called when production copyOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterCopyOptions(ctx *CopyOptionsContext) {}

// ExitCopyOptions is called when production copyOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitCopyOptions(ctx *CopyOptionsContext) {}

// EnterCopyOptionIdentity is called when production copyOptionIdentity is entered.
func (s *BaseDB2zSQLParserListener) EnterCopyOptionIdentity(ctx *CopyOptionIdentityContext) {}

// ExitCopyOptionIdentity is called when production copyOptionIdentity is exited.
func (s *BaseDB2zSQLParserListener) ExitCopyOptionIdentity(ctx *CopyOptionIdentityContext) {}

// EnterCopyOptionRowChangeTimestamp is called when production copyOptionRowChangeTimestamp is entered.
func (s *BaseDB2zSQLParserListener) EnterCopyOptionRowChangeTimestamp(ctx *CopyOptionRowChangeTimestampContext) {
}

// ExitCopyOptionRowChangeTimestamp is called when production copyOptionRowChangeTimestamp is exited.
func (s *BaseDB2zSQLParserListener) ExitCopyOptionRowChangeTimestamp(ctx *CopyOptionRowChangeTimestampContext) {
}

// EnterCopyOptionColumnDefaults is called when production copyOptionColumnDefaults is entered.
func (s *BaseDB2zSQLParserListener) EnterCopyOptionColumnDefaults(ctx *CopyOptionColumnDefaultsContext) {
}

// ExitCopyOptionColumnDefaults is called when production copyOptionColumnDefaults is exited.
func (s *BaseDB2zSQLParserListener) ExitCopyOptionColumnDefaults(ctx *CopyOptionColumnDefaultsContext) {
}

// EnterCopyOptionXmlTypeModifiers is called when production copyOptionXmlTypeModifiers is entered.
func (s *BaseDB2zSQLParserListener) EnterCopyOptionXmlTypeModifiers(ctx *CopyOptionXmlTypeModifiersContext) {
}

// ExitCopyOptionXmlTypeModifiers is called when production copyOptionXmlTypeModifiers is exited.
func (s *BaseDB2zSQLParserListener) ExitCopyOptionXmlTypeModifiers(ctx *CopyOptionXmlTypeModifiersContext) {
}

// EnterAsResultTable is called when production asResultTable is entered.
func (s *BaseDB2zSQLParserListener) EnterAsResultTable(ctx *AsResultTableContext) {}

// ExitAsResultTable is called when production asResultTable is exited.
func (s *BaseDB2zSQLParserListener) ExitAsResultTable(ctx *AsResultTableContext) {}

// EnterDeclareGlobalTemporaryTableAsResultTable is called when production declareGlobalTemporaryTableAsResultTable is entered.
func (s *BaseDB2zSQLParserListener) EnterDeclareGlobalTemporaryTableAsResultTable(ctx *DeclareGlobalTemporaryTableAsResultTableContext) {
}

// ExitDeclareGlobalTemporaryTableAsResultTable is called when production declareGlobalTemporaryTableAsResultTable is exited.
func (s *BaseDB2zSQLParserListener) ExitDeclareGlobalTemporaryTableAsResultTable(ctx *DeclareGlobalTemporaryTableAsResultTableContext) {
}

// EnterCreateTableMaterializedQueryDefinition is called when production createTableMaterializedQueryDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTableMaterializedQueryDefinition(ctx *CreateTableMaterializedQueryDefinitionContext) {
}

// ExitCreateTableMaterializedQueryDefinition is called when production createTableMaterializedQueryDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTableMaterializedQueryDefinition(ctx *CreateTableMaterializedQueryDefinitionContext) {
}

// EnterCreateTableColumnConstraint is called when production createTableColumnConstraint is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTableColumnConstraint(ctx *CreateTableColumnConstraintContext) {
}

// ExitCreateTableColumnConstraint is called when production createTableColumnConstraint is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTableColumnConstraint(ctx *CreateTableColumnConstraintContext) {
}

// EnterOrganizationClause is called when production organizationClause is entered.
func (s *BaseDB2zSQLParserListener) EnterOrganizationClause(ctx *OrganizationClauseContext) {}

// ExitOrganizationClause is called when production organizationClause is exited.
func (s *BaseDB2zSQLParserListener) ExitOrganizationClause(ctx *OrganizationClauseContext) {}

// EnterCreateGlobalTemporaryTableColumnDefinition is called when production createGlobalTemporaryTableColumnDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateGlobalTemporaryTableColumnDefinition(ctx *CreateGlobalTemporaryTableColumnDefinitionContext) {
}

// ExitCreateGlobalTemporaryTableColumnDefinition is called when production createGlobalTemporaryTableColumnDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateGlobalTemporaryTableColumnDefinition(ctx *CreateGlobalTemporaryTableColumnDefinitionContext) {
}

// EnterDeclareGlobalTemporaryTableColumnDefinition is called when production declareGlobalTemporaryTableColumnDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterDeclareGlobalTemporaryTableColumnDefinition(ctx *DeclareGlobalTemporaryTableColumnDefinitionContext) {
}

// ExitDeclareGlobalTemporaryTableColumnDefinition is called when production declareGlobalTemporaryTableColumnDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitDeclareGlobalTemporaryTableColumnDefinition(ctx *DeclareGlobalTemporaryTableColumnDefinitionContext) {
}

// EnterParameterDeclaration1 is called when production parameterDeclaration1 is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterDeclaration1(ctx *ParameterDeclaration1Context) {}

// ExitParameterDeclaration1 is called when production parameterDeclaration1 is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterDeclaration1(ctx *ParameterDeclaration1Context) {}

// EnterParameterDeclaration2 is called when production parameterDeclaration2 is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterDeclaration2(ctx *ParameterDeclaration2Context) {}

// ExitParameterDeclaration2 is called when production parameterDeclaration2 is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterDeclaration2(ctx *ParameterDeclaration2Context) {}

// EnterParameterDeclaration3 is called when production parameterDeclaration3 is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterDeclaration3(ctx *ParameterDeclaration3Context) {}

// ExitParameterDeclaration3 is called when production parameterDeclaration3 is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterDeclaration3(ctx *ParameterDeclaration3Context) {}

// EnterCreateFunctionStatementExternalScalarOptions is called when production createFunctionStatementExternalScalarOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementExternalScalarOptions(ctx *CreateFunctionStatementExternalScalarOptionsContext) {
}

// ExitCreateFunctionStatementExternalScalarOptions is called when production createFunctionStatementExternalScalarOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementExternalScalarOptions(ctx *CreateFunctionStatementExternalScalarOptionsContext) {
}

// EnterExternalNameOption1 is called when production externalNameOption1 is entered.
func (s *BaseDB2zSQLParserListener) EnterExternalNameOption1(ctx *ExternalNameOption1Context) {}

// ExitExternalNameOption1 is called when production externalNameOption1 is exited.
func (s *BaseDB2zSQLParserListener) ExitExternalNameOption1(ctx *ExternalNameOption1Context) {}

// EnterExternalNameOption2 is called when production externalNameOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterExternalNameOption2(ctx *ExternalNameOption2Context) {}

// ExitExternalNameOption2 is called when production externalNameOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitExternalNameOption2(ctx *ExternalNameOption2Context) {}

// EnterDynamicResultSetOption is called when production dynamicResultSetOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDynamicResultSetOption(ctx *DynamicResultSetOptionContext) {}

// ExitDynamicResultSetOption is called when production dynamicResultSetOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDynamicResultSetOption(ctx *DynamicResultSetOptionContext) {}

// EnterLanguageOption1 is called when production languageOption1 is entered.
func (s *BaseDB2zSQLParserListener) EnterLanguageOption1(ctx *LanguageOption1Context) {}

// ExitLanguageOption1 is called when production languageOption1 is exited.
func (s *BaseDB2zSQLParserListener) ExitLanguageOption1(ctx *LanguageOption1Context) {}

// EnterLanguageOption2 is called when production languageOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterLanguageOption2(ctx *LanguageOption2Context) {}

// ExitLanguageOption2 is called when production languageOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitLanguageOption2(ctx *LanguageOption2Context) {}

// EnterLanguageOption3 is called when production languageOption3 is entered.
func (s *BaseDB2zSQLParserListener) EnterLanguageOption3(ctx *LanguageOption3Context) {}

// ExitLanguageOption3 is called when production languageOption3 is exited.
func (s *BaseDB2zSQLParserListener) ExitLanguageOption3(ctx *LanguageOption3Context) {}

// EnterLanguageOption4 is called when production languageOption4 is entered.
func (s *BaseDB2zSQLParserListener) EnterLanguageOption4(ctx *LanguageOption4Context) {}

// ExitLanguageOption4 is called when production languageOption4 is exited.
func (s *BaseDB2zSQLParserListener) ExitLanguageOption4(ctx *LanguageOption4Context) {}

// EnterLanguageOption5 is called when production languageOption5 is entered.
func (s *BaseDB2zSQLParserListener) EnterLanguageOption5(ctx *LanguageOption5Context) {}

// ExitLanguageOption5 is called when production languageOption5 is exited.
func (s *BaseDB2zSQLParserListener) ExitLanguageOption5(ctx *LanguageOption5Context) {}

// EnterParameterStyleOption1 is called when production parameterStyleOption1 is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterStyleOption1(ctx *ParameterStyleOption1Context) {}

// ExitParameterStyleOption1 is called when production parameterStyleOption1 is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterStyleOption1(ctx *ParameterStyleOption1Context) {}

// EnterParameterStyleOption2 is called when production parameterStyleOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterStyleOption2(ctx *ParameterStyleOption2Context) {}

// ExitParameterStyleOption2 is called when production parameterStyleOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterStyleOption2(ctx *ParameterStyleOption2Context) {}

// EnterParameterStyleOption3 is called when production parameterStyleOption3 is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterStyleOption3(ctx *ParameterStyleOption3Context) {}

// ExitParameterStyleOption3 is called when production parameterStyleOption3 is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterStyleOption3(ctx *ParameterStyleOption3Context) {}

// EnterDeterministicOption is called when production deterministicOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDeterministicOption(ctx *DeterministicOptionContext) {}

// ExitDeterministicOption is called when production deterministicOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDeterministicOption(ctx *DeterministicOptionContext) {}

// EnterFencedOption is called when production fencedOption is entered.
func (s *BaseDB2zSQLParserListener) EnterFencedOption(ctx *FencedOptionContext) {}

// ExitFencedOption is called when production fencedOption is exited.
func (s *BaseDB2zSQLParserListener) ExitFencedOption(ctx *FencedOptionContext) {}

// EnterNullInputOption1 is called when production nullInputOption1 is entered.
func (s *BaseDB2zSQLParserListener) EnterNullInputOption1(ctx *NullInputOption1Context) {}

// ExitNullInputOption1 is called when production nullInputOption1 is exited.
func (s *BaseDB2zSQLParserListener) ExitNullInputOption1(ctx *NullInputOption1Context) {}

// EnterNullInputOption2 is called when production nullInputOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterNullInputOption2(ctx *NullInputOption2Context) {}

// ExitNullInputOption2 is called when production nullInputOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitNullInputOption2(ctx *NullInputOption2Context) {}

// EnterNullInputOption3 is called when production nullInputOption3 is entered.
func (s *BaseDB2zSQLParserListener) EnterNullInputOption3(ctx *NullInputOption3Context) {}

// ExitNullInputOption3 is called when production nullInputOption3 is exited.
func (s *BaseDB2zSQLParserListener) ExitNullInputOption3(ctx *NullInputOption3Context) {}

// EnterNullInputOption4 is called when production nullInputOption4 is entered.
func (s *BaseDB2zSQLParserListener) EnterNullInputOption4(ctx *NullInputOption4Context) {}

// ExitNullInputOption4 is called when production nullInputOption4 is exited.
func (s *BaseDB2zSQLParserListener) ExitNullInputOption4(ctx *NullInputOption4Context) {}

// EnterDebugOption is called when production debugOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDebugOption(ctx *DebugOptionContext) {}

// ExitDebugOption is called when production debugOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDebugOption(ctx *DebugOptionContext) {}

// EnterSqlDataOption1 is called when production sqlDataOption1 is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlDataOption1(ctx *SqlDataOption1Context) {}

// ExitSqlDataOption1 is called when production sqlDataOption1 is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlDataOption1(ctx *SqlDataOption1Context) {}

// EnterSqlDataOption2 is called when production sqlDataOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlDataOption2(ctx *SqlDataOption2Context) {}

// ExitSqlDataOption2 is called when production sqlDataOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlDataOption2(ctx *SqlDataOption2Context) {}

// EnterSqlDataOption3 is called when production sqlDataOption3 is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlDataOption3(ctx *SqlDataOption3Context) {}

// ExitSqlDataOption3 is called when production sqlDataOption3 is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlDataOption3(ctx *SqlDataOption3Context) {}

// EnterSqlDataOption4 is called when production sqlDataOption4 is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlDataOption4(ctx *SqlDataOption4Context) {}

// ExitSqlDataOption4 is called when production sqlDataOption4 is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlDataOption4(ctx *SqlDataOption4Context) {}

// EnterExternalActionOption is called when production externalActionOption is entered.
func (s *BaseDB2zSQLParserListener) EnterExternalActionOption(ctx *ExternalActionOptionContext) {}

// ExitExternalActionOption is called when production externalActionOption is exited.
func (s *BaseDB2zSQLParserListener) ExitExternalActionOption(ctx *ExternalActionOptionContext) {}

// EnterPackagePathOption is called when production packagePathOption is entered.
func (s *BaseDB2zSQLParserListener) EnterPackagePathOption(ctx *PackagePathOptionContext) {}

// ExitPackagePathOption is called when production packagePathOption is exited.
func (s *BaseDB2zSQLParserListener) ExitPackagePathOption(ctx *PackagePathOptionContext) {}

// EnterScratchpadOption is called when production scratchpadOption is entered.
func (s *BaseDB2zSQLParserListener) EnterScratchpadOption(ctx *ScratchpadOptionContext) {}

// ExitScratchpadOption is called when production scratchpadOption is exited.
func (s *BaseDB2zSQLParserListener) ExitScratchpadOption(ctx *ScratchpadOptionContext) {}

// EnterFinalCallOption is called when production finalCallOption is entered.
func (s *BaseDB2zSQLParserListener) EnterFinalCallOption(ctx *FinalCallOptionContext) {}

// ExitFinalCallOption is called when production finalCallOption is exited.
func (s *BaseDB2zSQLParserListener) ExitFinalCallOption(ctx *FinalCallOptionContext) {}

// EnterParallelOption1 is called when production parallelOption1 is entered.
func (s *BaseDB2zSQLParserListener) EnterParallelOption1(ctx *ParallelOption1Context) {}

// ExitParallelOption1 is called when production parallelOption1 is exited.
func (s *BaseDB2zSQLParserListener) ExitParallelOption1(ctx *ParallelOption1Context) {}

// EnterParallelOption2 is called when production parallelOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterParallelOption2(ctx *ParallelOption2Context) {}

// ExitParallelOption2 is called when production parallelOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitParallelOption2(ctx *ParallelOption2Context) {}

// EnterDbinfoOption is called when production dbinfoOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDbinfoOption(ctx *DbinfoOptionContext) {}

// ExitDbinfoOption is called when production dbinfoOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDbinfoOption(ctx *DbinfoOptionContext) {}

// EnterCardinalityOption is called when production cardinalityOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCardinalityOption(ctx *CardinalityOptionContext) {}

// ExitCardinalityOption is called when production cardinalityOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCardinalityOption(ctx *CardinalityOptionContext) {}

// EnterCollectionIdOption is called when production collectionIdOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCollectionIdOption(ctx *CollectionIdOptionContext) {}

// ExitCollectionIdOption is called when production collectionIdOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCollectionIdOption(ctx *CollectionIdOptionContext) {}

// EnterWlmEnvironmentOption1 is called when production wlmEnvironmentOption1 is entered.
func (s *BaseDB2zSQLParserListener) EnterWlmEnvironmentOption1(ctx *WlmEnvironmentOption1Context) {}

// ExitWlmEnvironmentOption1 is called when production wlmEnvironmentOption1 is exited.
func (s *BaseDB2zSQLParserListener) ExitWlmEnvironmentOption1(ctx *WlmEnvironmentOption1Context) {}

// EnterWlmEnvironmentOption2 is called when production wlmEnvironmentOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterWlmEnvironmentOption2(ctx *WlmEnvironmentOption2Context) {}

// ExitWlmEnvironmentOption2 is called when production wlmEnvironmentOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitWlmEnvironmentOption2(ctx *WlmEnvironmentOption2Context) {}

// EnterWlmEnvironmentOption3 is called when production wlmEnvironmentOption3 is entered.
func (s *BaseDB2zSQLParserListener) EnterWlmEnvironmentOption3(ctx *WlmEnvironmentOption3Context) {}

// ExitWlmEnvironmentOption3 is called when production wlmEnvironmentOption3 is exited.
func (s *BaseDB2zSQLParserListener) ExitWlmEnvironmentOption3(ctx *WlmEnvironmentOption3Context) {}

// EnterAsuTimeOption is called when production asuTimeOption is entered.
func (s *BaseDB2zSQLParserListener) EnterAsuTimeOption(ctx *AsuTimeOptionContext) {}

// ExitAsuTimeOption is called when production asuTimeOption is exited.
func (s *BaseDB2zSQLParserListener) ExitAsuTimeOption(ctx *AsuTimeOptionContext) {}

// EnterStayResidentOption is called when production stayResidentOption is entered.
func (s *BaseDB2zSQLParserListener) EnterStayResidentOption(ctx *StayResidentOptionContext) {}

// ExitStayResidentOption is called when production stayResidentOption is exited.
func (s *BaseDB2zSQLParserListener) ExitStayResidentOption(ctx *StayResidentOptionContext) {}

// EnterProgramTypeOption is called when production programTypeOption is entered.
func (s *BaseDB2zSQLParserListener) EnterProgramTypeOption(ctx *ProgramTypeOptionContext) {}

// ExitProgramTypeOption is called when production programTypeOption is exited.
func (s *BaseDB2zSQLParserListener) ExitProgramTypeOption(ctx *ProgramTypeOptionContext) {}

// EnterSecurityOption is called when production securityOption is entered.
func (s *BaseDB2zSQLParserListener) EnterSecurityOption(ctx *SecurityOptionContext) {}

// ExitSecurityOption is called when production securityOption is exited.
func (s *BaseDB2zSQLParserListener) ExitSecurityOption(ctx *SecurityOptionContext) {}

// EnterStopAfterFailureOption is called when production stopAfterFailureOption is entered.
func (s *BaseDB2zSQLParserListener) EnterStopAfterFailureOption(ctx *StopAfterFailureOptionContext) {}

// ExitStopAfterFailureOption is called when production stopAfterFailureOption is exited.
func (s *BaseDB2zSQLParserListener) ExitStopAfterFailureOption(ctx *StopAfterFailureOptionContext) {}

// EnterRunOptionsOption is called when production runOptionsOption is entered.
func (s *BaseDB2zSQLParserListener) EnterRunOptionsOption(ctx *RunOptionsOptionContext) {}

// ExitRunOptionsOption is called when production runOptionsOption is exited.
func (s *BaseDB2zSQLParserListener) ExitRunOptionsOption(ctx *RunOptionsOptionContext) {}

// EnterCommitOnReturnOption is called when production commitOnReturnOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCommitOnReturnOption(ctx *CommitOnReturnOptionContext) {}

// ExitCommitOnReturnOption is called when production commitOnReturnOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCommitOnReturnOption(ctx *CommitOnReturnOptionContext) {}

// EnterSpecialRegistersOption is called when production specialRegistersOption is entered.
func (s *BaseDB2zSQLParserListener) EnterSpecialRegistersOption(ctx *SpecialRegistersOptionContext) {}

// ExitSpecialRegistersOption is called when production specialRegistersOption is exited.
func (s *BaseDB2zSQLParserListener) ExitSpecialRegistersOption(ctx *SpecialRegistersOptionContext) {}

// EnterSpecialRegistersOption2 is called when production specialRegistersOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterSpecialRegistersOption2(ctx *SpecialRegistersOption2Context) {
}

// ExitSpecialRegistersOption2 is called when production specialRegistersOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitSpecialRegistersOption2(ctx *SpecialRegistersOption2Context) {
}

// EnterDispatchOption is called when production dispatchOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDispatchOption(ctx *DispatchOptionContext) {}

// ExitDispatchOption is called when production dispatchOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDispatchOption(ctx *DispatchOptionContext) {}

// EnterSecuredOption is called when production securedOption is entered.
func (s *BaseDB2zSQLParserListener) EnterSecuredOption(ctx *SecuredOptionContext) {}

// ExitSecuredOption is called when production securedOption is exited.
func (s *BaseDB2zSQLParserListener) ExitSecuredOption(ctx *SecuredOptionContext) {}

// EnterSpecificNameOption1 is called when production specificNameOption1 is entered.
func (s *BaseDB2zSQLParserListener) EnterSpecificNameOption1(ctx *SpecificNameOption1Context) {}

// ExitSpecificNameOption1 is called when production specificNameOption1 is exited.
func (s *BaseDB2zSQLParserListener) ExitSpecificNameOption1(ctx *SpecificNameOption1Context) {}

// EnterSpecificNameOption2 is called when production specificNameOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterSpecificNameOption2(ctx *SpecificNameOption2Context) {}

// ExitSpecificNameOption2 is called when production specificNameOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitSpecificNameOption2(ctx *SpecificNameOption2Context) {}

// EnterParameterOption1 is called when production parameterOption1 is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterOption1(ctx *ParameterOption1Context) {}

// ExitParameterOption1 is called when production parameterOption1 is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterOption1(ctx *ParameterOption1Context) {}

// EnterParameterOption2 is called when production parameterOption2 is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterOption2(ctx *ParameterOption2Context) {}

// ExitParameterOption2 is called when production parameterOption2 is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterOption2(ctx *ParameterOption2Context) {}

// EnterCreateFunctionStatementExternalTableOptions is called when production createFunctionStatementExternalTableOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementExternalTableOptions(ctx *CreateFunctionStatementExternalTableOptionsContext) {
}

// ExitCreateFunctionStatementExternalTableOptions is called when production createFunctionStatementExternalTableOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementExternalTableOptions(ctx *CreateFunctionStatementExternalTableOptionsContext) {
}

// EnterCreateFunctionStatementSourcedOptions is called when production createFunctionStatementSourcedOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementSourcedOptions(ctx *CreateFunctionStatementSourcedOptionsContext) {
}

// ExitCreateFunctionStatementSourcedOptions is called when production createFunctionStatementSourcedOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementSourcedOptions(ctx *CreateFunctionStatementSourcedOptionsContext) {
}

// EnterInlineSqlScalarFunctionDefinition is called when production inlineSqlScalarFunctionDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterInlineSqlScalarFunctionDefinition(ctx *InlineSqlScalarFunctionDefinitionContext) {
}

// ExitInlineSqlScalarFunctionDefinition is called when production inlineSqlScalarFunctionDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitInlineSqlScalarFunctionDefinition(ctx *InlineSqlScalarFunctionDefinitionContext) {
}

// EnterCreateFunctionStatementInlineSqlScalarOptions is called when production createFunctionStatementInlineSqlScalarOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementInlineSqlScalarOptions(ctx *CreateFunctionStatementInlineSqlScalarOptionsContext) {
}

// ExitCreateFunctionStatementInlineSqlScalarOptions is called when production createFunctionStatementInlineSqlScalarOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementInlineSqlScalarOptions(ctx *CreateFunctionStatementInlineSqlScalarOptionsContext) {
}

// EnterCompiledSqlScalarFunctionDefinition is called when production compiledSqlScalarFunctionDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterCompiledSqlScalarFunctionDefinition(ctx *CompiledSqlScalarFunctionDefinitionContext) {
}

// ExitCompiledSqlScalarFunctionDefinition is called when production compiledSqlScalarFunctionDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitCompiledSqlScalarFunctionDefinition(ctx *CompiledSqlScalarFunctionDefinitionContext) {
}

// EnterCreateFunctionStatementCompiledSqlScalarOptions is called when production createFunctionStatementCompiledSqlScalarOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementCompiledSqlScalarOptions(ctx *CreateFunctionStatementCompiledSqlScalarOptionsContext) {
}

// ExitCreateFunctionStatementCompiledSqlScalarOptions is called when production createFunctionStatementCompiledSqlScalarOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementCompiledSqlScalarOptions(ctx *CreateFunctionStatementCompiledSqlScalarOptionsContext) {
}

// EnterSqlTableFunctionDefinition is called when production sqlTableFunctionDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlTableFunctionDefinition(ctx *SqlTableFunctionDefinitionContext) {
}

// ExitSqlTableFunctionDefinition is called when production sqlTableFunctionDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlTableFunctionDefinition(ctx *SqlTableFunctionDefinitionContext) {
}

// EnterCreateFunctionStatementSqlTableOptions is called when production createFunctionStatementSqlTableOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateFunctionStatementSqlTableOptions(ctx *CreateFunctionStatementSqlTableOptionsContext) {
}

// ExitCreateFunctionStatementSqlTableOptions is called when production createFunctionStatementSqlTableOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateFunctionStatementSqlTableOptions(ctx *CreateFunctionStatementSqlTableOptionsContext) {
}

// EnterSequenceAlias is called when production sequenceAlias is entered.
func (s *BaseDB2zSQLParserListener) EnterSequenceAlias(ctx *SequenceAliasContext) {}

// ExitSequenceAlias is called when production sequenceAlias is exited.
func (s *BaseDB2zSQLParserListener) ExitSequenceAlias(ctx *SequenceAliasContext) {}

// EnterTableAlias is called when production tableAlias is entered.
func (s *BaseDB2zSQLParserListener) EnterTableAlias(ctx *TableAliasContext) {}

// ExitTableAlias is called when production tableAlias is exited.
func (s *BaseDB2zSQLParserListener) ExitTableAlias(ctx *TableAliasContext) {}

// EnterAuthorization is called when production authorization is entered.
func (s *BaseDB2zSQLParserListener) EnterAuthorization(ctx *AuthorizationContext) {}

// ExitAuthorization is called when production authorization is exited.
func (s *BaseDB2zSQLParserListener) ExitAuthorization(ctx *AuthorizationContext) {}

// EnterSearchedDelete is called when production searchedDelete is entered.
func (s *BaseDB2zSQLParserListener) EnterSearchedDelete(ctx *SearchedDeleteContext) {}

// ExitSearchedDelete is called when production searchedDelete is exited.
func (s *BaseDB2zSQLParserListener) ExitSearchedDelete(ctx *SearchedDeleteContext) {}

// EnterPositionedDelete is called when production positionedDelete is entered.
func (s *BaseDB2zSQLParserListener) EnterPositionedDelete(ctx *PositionedDeleteContext) {}

// ExitPositionedDelete is called when production positionedDelete is exited.
func (s *BaseDB2zSQLParserListener) ExitPositionedDelete(ctx *PositionedDeleteContext) {}

// EnterSearchedUpdate is called when production searchedUpdate is entered.
func (s *BaseDB2zSQLParserListener) EnterSearchedUpdate(ctx *SearchedUpdateContext) {}

// ExitSearchedUpdate is called when production searchedUpdate is exited.
func (s *BaseDB2zSQLParserListener) ExitSearchedUpdate(ctx *SearchedUpdateContext) {}

// EnterPositionedUpdate is called when production positionedUpdate is entered.
func (s *BaseDB2zSQLParserListener) EnterPositionedUpdate(ctx *PositionedUpdateContext) {}

// ExitPositionedUpdate is called when production positionedUpdate is exited.
func (s *BaseDB2zSQLParserListener) ExitPositionedUpdate(ctx *PositionedUpdateContext) {}

// EnterSourceValues is called when production sourceValues is entered.
func (s *BaseDB2zSQLParserListener) EnterSourceValues(ctx *SourceValuesContext) {}

// ExitSourceValues is called when production sourceValues is exited.
func (s *BaseDB2zSQLParserListener) ExitSourceValues(ctx *SourceValuesContext) {}

// EnterValuesSingleRow is called when production valuesSingleRow is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesSingleRow(ctx *ValuesSingleRowContext) {}

// ExitValuesSingleRow is called when production valuesSingleRow is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesSingleRow(ctx *ValuesSingleRowContext) {}

// EnterValuesMultipleRow is called when production valuesMultipleRow is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesMultipleRow(ctx *ValuesMultipleRowContext) {}

// ExitValuesMultipleRow is called when production valuesMultipleRow is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesMultipleRow(ctx *ValuesMultipleRowContext) {}

// EnterMatchingCondition is called when production matchingCondition is entered.
func (s *BaseDB2zSQLParserListener) EnterMatchingCondition(ctx *MatchingConditionContext) {}

// ExitMatchingCondition is called when production matchingCondition is exited.
func (s *BaseDB2zSQLParserListener) ExitMatchingCondition(ctx *MatchingConditionContext) {}

// EnterModificationOperation is called when production modificationOperation is entered.
func (s *BaseDB2zSQLParserListener) EnterModificationOperation(ctx *ModificationOperationContext) {}

// ExitModificationOperation is called when production modificationOperation is exited.
func (s *BaseDB2zSQLParserListener) ExitModificationOperation(ctx *ModificationOperationContext) {}

// EnterAssignmentClause is called when production assignmentClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAssignmentClause(ctx *AssignmentClauseContext) {}

// ExitAssignmentClause is called when production assignmentClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAssignmentClause(ctx *AssignmentClauseContext) {}

// EnterSetAssignmentClause is called when production setAssignmentClause is entered.
func (s *BaseDB2zSQLParserListener) EnterSetAssignmentClause(ctx *SetAssignmentClauseContext) {}

// ExitSetAssignmentClause is called when production setAssignmentClause is exited.
func (s *BaseDB2zSQLParserListener) ExitSetAssignmentClause(ctx *SetAssignmentClauseContext) {}

// EnterSetAssignmentTargetVariable is called when production setAssignmentTargetVariable is entered.
func (s *BaseDB2zSQLParserListener) EnterSetAssignmentTargetVariable(ctx *SetAssignmentTargetVariableContext) {
}

// ExitSetAssignmentTargetVariable is called when production setAssignmentTargetVariable is exited.
func (s *BaseDB2zSQLParserListener) ExitSetAssignmentTargetVariable(ctx *SetAssignmentTargetVariableContext) {
}

// EnterUpdateOperation is called when production updateOperation is entered.
func (s *BaseDB2zSQLParserListener) EnterUpdateOperation(ctx *UpdateOperationContext) {}

// ExitUpdateOperation is called when production updateOperation is exited.
func (s *BaseDB2zSQLParserListener) ExitUpdateOperation(ctx *UpdateOperationContext) {}

// EnterDeleteOperation is called when production deleteOperation is entered.
func (s *BaseDB2zSQLParserListener) EnterDeleteOperation(ctx *DeleteOperationContext) {}

// ExitDeleteOperation is called when production deleteOperation is exited.
func (s *BaseDB2zSQLParserListener) ExitDeleteOperation(ctx *DeleteOperationContext) {}

// EnterInsertOperation is called when production insertOperation is entered.
func (s *BaseDB2zSQLParserListener) EnterInsertOperation(ctx *InsertOperationContext) {}

// ExitInsertOperation is called when production insertOperation is exited.
func (s *BaseDB2zSQLParserListener) ExitInsertOperation(ctx *InsertOperationContext) {}

// EnterSignalInformation is called when production signalInformation is entered.
func (s *BaseDB2zSQLParserListener) EnterSignalInformation(ctx *SignalInformationContext) {}

// ExitSignalInformation is called when production signalInformation is exited.
func (s *BaseDB2zSQLParserListener) ExitSignalInformation(ctx *SignalInformationContext) {}

// EnterValuesList1 is called when production valuesList1 is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesList1(ctx *ValuesList1Context) {}

// ExitValuesList1 is called when production valuesList1 is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesList1(ctx *ValuesList1Context) {}

// EnterValuesList2 is called when production valuesList2 is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesList2(ctx *ValuesList2Context) {}

// ExitValuesList2 is called when production valuesList2 is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesList2(ctx *ValuesList2Context) {}

// EnterValuesList3 is called when production valuesList3 is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesList3(ctx *ValuesList3Context) {}

// ExitValuesList3 is called when production valuesList3 is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesList3(ctx *ValuesList3Context) {}

// EnterValuesList4 is called when production valuesList4 is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesList4(ctx *ValuesList4Context) {}

// ExitValuesList4 is called when production valuesList4 is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesList4(ctx *ValuesList4Context) {}

// EnterIncludeColumns is called when production includeColumns is entered.
func (s *BaseDB2zSQLParserListener) EnterIncludeColumns(ctx *IncludeColumnsContext) {}

// ExitIncludeColumns is called when production includeColumns is exited.
func (s *BaseDB2zSQLParserListener) ExitIncludeColumns(ctx *IncludeColumnsContext) {}

// EnterMultipleRowInsert is called when production multipleRowInsert is entered.
func (s *BaseDB2zSQLParserListener) EnterMultipleRowInsert(ctx *MultipleRowInsertContext) {}

// ExitMultipleRowInsert is called when production multipleRowInsert is exited.
func (s *BaseDB2zSQLParserListener) ExitMultipleRowInsert(ctx *MultipleRowInsertContext) {}

// EnterRegenerateClause is called when production regenerateClause is entered.
func (s *BaseDB2zSQLParserListener) EnterRegenerateClause(ctx *RegenerateClauseContext) {}

// ExitRegenerateClause is called when production regenerateClause is exited.
func (s *BaseDB2zSQLParserListener) ExitRegenerateClause(ctx *RegenerateClauseContext) {}

// EnterAlterIndexOptions is called when production alterIndexOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterIndexOptions(ctx *AlterIndexOptionsContext) {}

// ExitAlterIndexOptions is called when production alterIndexOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterIndexOptions(ctx *AlterIndexOptionsContext) {}

// EnterBufferpoolOption is called when production bufferpoolOption is entered.
func (s *BaseDB2zSQLParserListener) EnterBufferpoolOption(ctx *BufferpoolOptionContext) {}

// ExitBufferpoolOption is called when production bufferpoolOption is exited.
func (s *BaseDB2zSQLParserListener) ExitBufferpoolOption(ctx *BufferpoolOptionContext) {}

// EnterCloseOption is called when production closeOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCloseOption(ctx *CloseOptionContext) {}

// ExitCloseOption is called when production closeOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCloseOption(ctx *CloseOptionContext) {}

// EnterCopyOption is called when production copyOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCopyOption(ctx *CopyOptionContext) {}

// ExitCopyOption is called when production copyOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCopyOption(ctx *CopyOptionContext) {}

// EnterDssizeOption is called when production dssizeOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDssizeOption(ctx *DssizeOptionContext) {}

// ExitDssizeOption is called when production dssizeOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDssizeOption(ctx *DssizeOptionContext) {}

// EnterPiecesizeOption is called when production piecesizeOption is entered.
func (s *BaseDB2zSQLParserListener) EnterPiecesizeOption(ctx *PiecesizeOptionContext) {}

// ExitPiecesizeOption is called when production piecesizeOption is exited.
func (s *BaseDB2zSQLParserListener) ExitPiecesizeOption(ctx *PiecesizeOptionContext) {}

// EnterClusterOption is called when production clusterOption is entered.
func (s *BaseDB2zSQLParserListener) EnterClusterOption(ctx *ClusterOptionContext) {}

// ExitClusterOption is called when production clusterOption is exited.
func (s *BaseDB2zSQLParserListener) ExitClusterOption(ctx *ClusterOptionContext) {}

// EnterPaddedOption is called when production paddedOption is entered.
func (s *BaseDB2zSQLParserListener) EnterPaddedOption(ctx *PaddedOptionContext) {}

// ExitPaddedOption is called when production paddedOption is exited.
func (s *BaseDB2zSQLParserListener) ExitPaddedOption(ctx *PaddedOptionContext) {}

// EnterCompressOption is called when production compressOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCompressOption(ctx *CompressOptionContext) {}

// ExitCompressOption is called when production compressOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCompressOption(ctx *CompressOptionContext) {}

// EnterDefineOption is called when production defineOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDefineOption(ctx *DefineOptionContext) {}

// ExitDefineOption is called when production defineOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDefineOption(ctx *DefineOptionContext) {}

// EnterLocksizeOption is called when production locksizeOption is entered.
func (s *BaseDB2zSQLParserListener) EnterLocksizeOption(ctx *LocksizeOptionContext) {}

// ExitLocksizeOption is called when production locksizeOption is exited.
func (s *BaseDB2zSQLParserListener) ExitLocksizeOption(ctx *LocksizeOptionContext) {}

// EnterLockmaxOption is called when production lockmaxOption is entered.
func (s *BaseDB2zSQLParserListener) EnterLockmaxOption(ctx *LockmaxOptionContext) {}

// ExitLockmaxOption is called when production lockmaxOption is exited.
func (s *BaseDB2zSQLParserListener) ExitLockmaxOption(ctx *LockmaxOptionContext) {}

// EnterEnableDisableOption is called when production enableDisableOption is entered.
func (s *BaseDB2zSQLParserListener) EnterEnableDisableOption(ctx *EnableDisableOptionContext) {}

// ExitEnableDisableOption is called when production enableDisableOption is exited.
func (s *BaseDB2zSQLParserListener) ExitEnableDisableOption(ctx *EnableDisableOptionContext) {}

// EnterLoggedOption is called when production loggedOption is entered.
func (s *BaseDB2zSQLParserListener) EnterLoggedOption(ctx *LoggedOptionContext) {}

// ExitLoggedOption is called when production loggedOption is exited.
func (s *BaseDB2zSQLParserListener) ExitLoggedOption(ctx *LoggedOptionContext) {}

// EnterNotAtomicPhrase is called when production notAtomicPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterNotAtomicPhrase(ctx *NotAtomicPhraseContext) {}

// ExitNotAtomicPhrase is called when production notAtomicPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitNotAtomicPhrase(ctx *NotAtomicPhraseContext) {}

// EnterAlterIndexPartitionOptions is called when production alterIndexPartitionOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterIndexPartitionOptions(ctx *AlterIndexPartitionOptionsContext) {
}

// ExitAlterIndexPartitionOptions is called when production alterIndexPartitionOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterIndexPartitionOptions(ctx *AlterIndexPartitionOptionsContext) {
}

// EnterUsingSpecification1 is called when production usingSpecification1 is entered.
func (s *BaseDB2zSQLParserListener) EnterUsingSpecification1(ctx *UsingSpecification1Context) {}

// ExitUsingSpecification1 is called when production usingSpecification1 is exited.
func (s *BaseDB2zSQLParserListener) ExitUsingSpecification1(ctx *UsingSpecification1Context) {}

// EnterFreeSpecification is called when production freeSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterFreeSpecification(ctx *FreeSpecificationContext) {}

// ExitFreeSpecification is called when production freeSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitFreeSpecification(ctx *FreeSpecificationContext) {}

// EnterGbpcacheSpecification is called when production gbpcacheSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterGbpcacheSpecification(ctx *GbpcacheSpecificationContext) {}

// ExitGbpcacheSpecification is called when production gbpcacheSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitGbpcacheSpecification(ctx *GbpcacheSpecificationContext) {}

// EnterPartitionElement is called when production partitionElement is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitionElement(ctx *PartitionElementContext) {}

// ExitPartitionElement is called when production partitionElement is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitionElement(ctx *PartitionElementContext) {}

// EnterApplCompatValue is called when production applCompatValue is entered.
func (s *BaseDB2zSQLParserListener) EnterApplCompatValue(ctx *ApplCompatValueContext) {}

// ExitApplCompatValue is called when production applCompatValue is exited.
func (s *BaseDB2zSQLParserListener) ExitApplCompatValue(ctx *ApplCompatValueContext) {}

// EnterFunctionLevel is called when production functionLevel is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionLevel(ctx *FunctionLevelContext) {}

// ExitFunctionLevel is called when production functionLevel is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionLevel(ctx *FunctionLevelContext) {}

// EnterFunctionParameterType is called when production functionParameterType is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionParameterType(ctx *FunctionParameterTypeContext) {}

// ExitFunctionParameterType is called when production functionParameterType is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionParameterType(ctx *FunctionParameterTypeContext) {}

// EnterFunctionDataType is called when production functionDataType is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionDataType(ctx *FunctionDataTypeContext) {}

// ExitFunctionDataType is called when production functionDataType is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionDataType(ctx *FunctionDataTypeContext) {}

// EnterFunctionBuiltInType is called when production functionBuiltInType is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionBuiltInType(ctx *FunctionBuiltInTypeContext) {}

// ExitFunctionBuiltInType is called when production functionBuiltInType is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionBuiltInType(ctx *FunctionBuiltInTypeContext) {}

// EnterProcedureBuiltinType is called when production procedureBuiltinType is entered.
func (s *BaseDB2zSQLParserListener) EnterProcedureBuiltinType(ctx *ProcedureBuiltinTypeContext) {}

// ExitProcedureBuiltinType is called when production procedureBuiltinType is exited.
func (s *BaseDB2zSQLParserListener) ExitProcedureBuiltinType(ctx *ProcedureBuiltinTypeContext) {}

// EnterCreateTypeArrayBuiltinType is called when production createTypeArrayBuiltinType is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTypeArrayBuiltinType(ctx *CreateTypeArrayBuiltinTypeContext) {
}

// ExitCreateTypeArrayBuiltinType is called when production createTypeArrayBuiltinType is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTypeArrayBuiltinType(ctx *CreateTypeArrayBuiltinTypeContext) {
}

// EnterCreateTypeArrayBuiltinType2 is called when production createTypeArrayBuiltinType2 is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTypeArrayBuiltinType2(ctx *CreateTypeArrayBuiltinType2Context) {
}

// ExitCreateTypeArrayBuiltinType2 is called when production createTypeArrayBuiltinType2 is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTypeArrayBuiltinType2(ctx *CreateTypeArrayBuiltinType2Context) {
}

// EnterCreateVariableBuiltInType is called when production createVariableBuiltInType is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateVariableBuiltInType(ctx *CreateVariableBuiltInTypeContext) {
}

// ExitCreateVariableBuiltInType is called when production createVariableBuiltInType is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateVariableBuiltInType(ctx *CreateVariableBuiltInTypeContext) {
}

// EnterSourceDataType is called when production sourceDataType is entered.
func (s *BaseDB2zSQLParserListener) EnterSourceDataType(ctx *SourceDataTypeContext) {}

// ExitSourceDataType is called when production sourceDataType is exited.
func (s *BaseDB2zSQLParserListener) ExitSourceDataType(ctx *SourceDataTypeContext) {}

// EnterFunctionExternalOptionList is called when production functionExternalOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionExternalOptionList(ctx *FunctionExternalOptionListContext) {
}

// ExitFunctionExternalOptionList is called when production functionExternalOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionExternalOptionList(ctx *FunctionExternalOptionListContext) {
}

// EnterFunctionCompiledSqlScalarOptionList is called when production functionCompiledSqlScalarOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionCompiledSqlScalarOptionList(ctx *FunctionCompiledSqlScalarOptionListContext) {
}

// ExitFunctionCompiledSqlScalarOptionList is called when production functionCompiledSqlScalarOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionCompiledSqlScalarOptionList(ctx *FunctionCompiledSqlScalarOptionListContext) {
}

// EnterFunctionInlineSqlScalarOptionList is called when production functionInlineSqlScalarOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionInlineSqlScalarOptionList(ctx *FunctionInlineSqlScalarOptionListContext) {
}

// ExitFunctionInlineSqlScalarOptionList is called when production functionInlineSqlScalarOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionInlineSqlScalarOptionList(ctx *FunctionInlineSqlScalarOptionListContext) {
}

// EnterFunctionSqlTableOptionList is called when production functionSqlTableOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionSqlTableOptionList(ctx *FunctionSqlTableOptionListContext) {
}

// ExitFunctionSqlTableOptionList is called when production functionSqlTableOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionSqlTableOptionList(ctx *FunctionSqlTableOptionListContext) {
}

// EnterProcedureOptionList is called when production procedureOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterProcedureOptionList(ctx *ProcedureOptionListContext) {}

// ExitProcedureOptionList is called when production procedureOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitProcedureOptionList(ctx *ProcedureOptionListContext) {}

// EnterCreateProcedureOptionList is called when production createProcedureOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateProcedureOptionList(ctx *CreateProcedureOptionListContext) {
}

// ExitCreateProcedureOptionList is called when production createProcedureOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateProcedureOptionList(ctx *CreateProcedureOptionListContext) {
}

// EnterProcedureSQLPLOptionList is called when production procedureSQLPLOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterProcedureSQLPLOptionList(ctx *ProcedureSQLPLOptionListContext) {
}

// ExitProcedureSQLPLOptionList is called when production procedureSQLPLOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitProcedureSQLPLOptionList(ctx *ProcedureSQLPLOptionListContext) {
}

// EnterVersionOption is called when production versionOption is entered.
func (s *BaseDB2zSQLParserListener) EnterVersionOption(ctx *VersionOptionContext) {}

// ExitVersionOption is called when production versionOption is exited.
func (s *BaseDB2zSQLParserListener) ExitVersionOption(ctx *VersionOptionContext) {}

// EnterCommitOnReturnOptionSQLPL is called when production commitOnReturnOptionSQLPL is entered.
func (s *BaseDB2zSQLParserListener) EnterCommitOnReturnOptionSQLPL(ctx *CommitOnReturnOptionSQLPLContext) {
}

// ExitCommitOnReturnOptionSQLPL is called when production commitOnReturnOptionSQLPL is exited.
func (s *BaseDB2zSQLParserListener) ExitCommitOnReturnOptionSQLPL(ctx *CommitOnReturnOptionSQLPLContext) {
}

// EnterSchemaQualifierOption is called when production schemaQualifierOption is entered.
func (s *BaseDB2zSQLParserListener) EnterSchemaQualifierOption(ctx *SchemaQualifierOptionContext) {}

// ExitSchemaQualifierOption is called when production schemaQualifierOption is exited.
func (s *BaseDB2zSQLParserListener) ExitSchemaQualifierOption(ctx *SchemaQualifierOptionContext) {}

// EnterCurrentDataOption is called when production currentDataOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCurrentDataOption(ctx *CurrentDataOptionContext) {}

// ExitCurrentDataOption is called when production currentDataOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCurrentDataOption(ctx *CurrentDataOptionContext) {}

// EnterImmediateWriteOption is called when production immediateWriteOption is entered.
func (s *BaseDB2zSQLParserListener) EnterImmediateWriteOption(ctx *ImmediateWriteOptionContext) {}

// ExitImmediateWriteOption is called when production immediateWriteOption is exited.
func (s *BaseDB2zSQLParserListener) ExitImmediateWriteOption(ctx *ImmediateWriteOptionContext) {}

// EnterExplainOption is called when production explainOption is entered.
func (s *BaseDB2zSQLParserListener) EnterExplainOption(ctx *ExplainOptionContext) {}

// ExitExplainOption is called when production explainOption is exited.
func (s *BaseDB2zSQLParserListener) ExitExplainOption(ctx *ExplainOptionContext) {}

// EnterReoptOption is called when production reoptOption is entered.
func (s *BaseDB2zSQLParserListener) EnterReoptOption(ctx *ReoptOptionContext) {}

// ExitReoptOption is called when production reoptOption is exited.
func (s *BaseDB2zSQLParserListener) ExitReoptOption(ctx *ReoptOptionContext) {}

// EnterPackageOwnerOption is called when production packageOwnerOption is entered.
func (s *BaseDB2zSQLParserListener) EnterPackageOwnerOption(ctx *PackageOwnerOptionContext) {}

// ExitPackageOwnerOption is called when production packageOwnerOption is exited.
func (s *BaseDB2zSQLParserListener) ExitPackageOwnerOption(ctx *PackageOwnerOptionContext) {}

// EnterDeferPrepareOption is called when production deferPrepareOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDeferPrepareOption(ctx *DeferPrepareOptionContext) {}

// ExitDeferPrepareOption is called when production deferPrepareOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDeferPrepareOption(ctx *DeferPrepareOptionContext) {}

// EnterDegreeOption is called when production degreeOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDegreeOption(ctx *DegreeOptionContext) {}

// ExitDegreeOption is called when production degreeOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDegreeOption(ctx *DegreeOptionContext) {}

// EnterDynamicRulesOption is called when production dynamicRulesOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDynamicRulesOption(ctx *DynamicRulesOptionContext) {}

// ExitDynamicRulesOption is called when production dynamicRulesOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDynamicRulesOption(ctx *DynamicRulesOptionContext) {}

// EnterConcurrentAccessOption is called when production concurrentAccessOption is entered.
func (s *BaseDB2zSQLParserListener) EnterConcurrentAccessOption(ctx *ConcurrentAccessOptionContext) {}

// ExitConcurrentAccessOption is called when production concurrentAccessOption is exited.
func (s *BaseDB2zSQLParserListener) ExitConcurrentAccessOption(ctx *ConcurrentAccessOptionContext) {}

// EnterApplicationEncodingOption is called when production applicationEncodingOption is entered.
func (s *BaseDB2zSQLParserListener) EnterApplicationEncodingOption(ctx *ApplicationEncodingOptionContext) {
}

// ExitApplicationEncodingOption is called when production applicationEncodingOption is exited.
func (s *BaseDB2zSQLParserListener) ExitApplicationEncodingOption(ctx *ApplicationEncodingOptionContext) {
}

// EnterIsolationLevelOption is called when production isolationLevelOption is entered.
func (s *BaseDB2zSQLParserListener) EnterIsolationLevelOption(ctx *IsolationLevelOptionContext) {}

// ExitIsolationLevelOption is called when production isolationLevelOption is exited.
func (s *BaseDB2zSQLParserListener) ExitIsolationLevelOption(ctx *IsolationLevelOptionContext) {}

// EnterKeepDynamicOption is called when production keepDynamicOption is entered.
func (s *BaseDB2zSQLParserListener) EnterKeepDynamicOption(ctx *KeepDynamicOptionContext) {}

// ExitKeepDynamicOption is called when production keepDynamicOption is exited.
func (s *BaseDB2zSQLParserListener) ExitKeepDynamicOption(ctx *KeepDynamicOptionContext) {}

// EnterOpthintOption is called when production opthintOption is entered.
func (s *BaseDB2zSQLParserListener) EnterOpthintOption(ctx *OpthintOptionContext) {}

// ExitOpthintOption is called when production opthintOption is exited.
func (s *BaseDB2zSQLParserListener) ExitOpthintOption(ctx *OpthintOptionContext) {}

// EnterSqlPathOption is called when production sqlPathOption is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlPathOption(ctx *SqlPathOptionContext) {}

// ExitSqlPathOption is called when production sqlPathOption is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlPathOption(ctx *SqlPathOptionContext) {}

// EnterSqlPathOptionItem is called when production sqlPathOptionItem is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlPathOptionItem(ctx *SqlPathOptionItemContext) {}

// ExitSqlPathOptionItem is called when production sqlPathOptionItem is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlPathOptionItem(ctx *SqlPathOptionItemContext) {}

// EnterQueryAccelerationOption is called when production queryAccelerationOption is entered.
func (s *BaseDB2zSQLParserListener) EnterQueryAccelerationOption(ctx *QueryAccelerationOptionContext) {
}

// ExitQueryAccelerationOption is called when production queryAccelerationOption is exited.
func (s *BaseDB2zSQLParserListener) ExitQueryAccelerationOption(ctx *QueryAccelerationOptionContext) {
}

// EnterQueryAccelerationOptionItem is called when production queryAccelerationOptionItem is entered.
func (s *BaseDB2zSQLParserListener) EnterQueryAccelerationOptionItem(ctx *QueryAccelerationOptionItemContext) {
}

// ExitQueryAccelerationOptionItem is called when production queryAccelerationOptionItem is exited.
func (s *BaseDB2zSQLParserListener) ExitQueryAccelerationOptionItem(ctx *QueryAccelerationOptionItemContext) {
}

// EnterGetAccelArchiveOption is called when production getAccelArchiveOption is entered.
func (s *BaseDB2zSQLParserListener) EnterGetAccelArchiveOption(ctx *GetAccelArchiveOptionContext) {}

// ExitGetAccelArchiveOption is called when production getAccelArchiveOption is exited.
func (s *BaseDB2zSQLParserListener) ExitGetAccelArchiveOption(ctx *GetAccelArchiveOptionContext) {}

// EnterAccelerationOption is called when production accelerationOption is entered.
func (s *BaseDB2zSQLParserListener) EnterAccelerationOption(ctx *AccelerationOptionContext) {}

// ExitAccelerationOption is called when production accelerationOption is exited.
func (s *BaseDB2zSQLParserListener) ExitAccelerationOption(ctx *AccelerationOptionContext) {}

// EnterReleaseAtOption is called when production releaseAtOption is entered.
func (s *BaseDB2zSQLParserListener) EnterReleaseAtOption(ctx *ReleaseAtOptionContext) {}

// ExitReleaseAtOption is called when production releaseAtOption is exited.
func (s *BaseDB2zSQLParserListener) ExitReleaseAtOption(ctx *ReleaseAtOptionContext) {}

// EnterBusinessTimeSensitiveOption is called when production businessTimeSensitiveOption is entered.
func (s *BaseDB2zSQLParserListener) EnterBusinessTimeSensitiveOption(ctx *BusinessTimeSensitiveOptionContext) {
}

// ExitBusinessTimeSensitiveOption is called when production businessTimeSensitiveOption is exited.
func (s *BaseDB2zSQLParserListener) ExitBusinessTimeSensitiveOption(ctx *BusinessTimeSensitiveOptionContext) {
}

// EnterSystemTimeSensitiveOption is called when production systemTimeSensitiveOption is entered.
func (s *BaseDB2zSQLParserListener) EnterSystemTimeSensitiveOption(ctx *SystemTimeSensitiveOptionContext) {
}

// ExitSystemTimeSensitiveOption is called when production systemTimeSensitiveOption is exited.
func (s *BaseDB2zSQLParserListener) ExitSystemTimeSensitiveOption(ctx *SystemTimeSensitiveOptionContext) {
}

// EnterArchiveSensitiveOption is called when production archiveSensitiveOption is entered.
func (s *BaseDB2zSQLParserListener) EnterArchiveSensitiveOption(ctx *ArchiveSensitiveOptionContext) {}

// ExitArchiveSensitiveOption is called when production archiveSensitiveOption is exited.
func (s *BaseDB2zSQLParserListener) ExitArchiveSensitiveOption(ctx *ArchiveSensitiveOptionContext) {}

// EnterApplcompatOption is called when production applcompatOption is entered.
func (s *BaseDB2zSQLParserListener) EnterApplcompatOption(ctx *ApplcompatOptionContext) {}

// ExitApplcompatOption is called when production applcompatOption is exited.
func (s *BaseDB2zSQLParserListener) ExitApplcompatOption(ctx *ApplcompatOptionContext) {}

// EnterValidateOption is called when production validateOption is entered.
func (s *BaseDB2zSQLParserListener) EnterValidateOption(ctx *ValidateOptionContext) {}

// ExitValidateOption is called when production validateOption is exited.
func (s *BaseDB2zSQLParserListener) ExitValidateOption(ctx *ValidateOptionContext) {}

// EnterRoundingOption is called when production roundingOption is entered.
func (s *BaseDB2zSQLParserListener) EnterRoundingOption(ctx *RoundingOptionContext) {}

// ExitRoundingOption is called when production roundingOption is exited.
func (s *BaseDB2zSQLParserListener) ExitRoundingOption(ctx *RoundingOptionContext) {}

// EnterRoundingOptionItem is called when production roundingOptionItem is entered.
func (s *BaseDB2zSQLParserListener) EnterRoundingOptionItem(ctx *RoundingOptionItemContext) {}

// ExitRoundingOptionItem is called when production roundingOptionItem is exited.
func (s *BaseDB2zSQLParserListener) ExitRoundingOptionItem(ctx *RoundingOptionItemContext) {}

// EnterDateFormatOption is called when production dateFormatOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDateFormatOption(ctx *DateFormatOptionContext) {}

// ExitDateFormatOption is called when production dateFormatOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDateFormatOption(ctx *DateFormatOptionContext) {}

// EnterDateTimeFormatOptionItem is called when production dateTimeFormatOptionItem is entered.
func (s *BaseDB2zSQLParserListener) EnterDateTimeFormatOptionItem(ctx *DateTimeFormatOptionItemContext) {
}

// ExitDateTimeFormatOptionItem is called when production dateTimeFormatOptionItem is exited.
func (s *BaseDB2zSQLParserListener) ExitDateTimeFormatOptionItem(ctx *DateTimeFormatOptionItemContext) {
}

// EnterTimeFormatOption is called when production timeFormatOption is entered.
func (s *BaseDB2zSQLParserListener) EnterTimeFormatOption(ctx *TimeFormatOptionContext) {}

// ExitTimeFormatOption is called when production timeFormatOption is exited.
func (s *BaseDB2zSQLParserListener) ExitTimeFormatOption(ctx *TimeFormatOptionContext) {}

// EnterDecimalOption is called when production decimalOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDecimalOption(ctx *DecimalOptionContext) {}

// ExitDecimalOption is called when production decimalOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDecimalOption(ctx *DecimalOptionContext) {}

// EnterForUpdateOption is called when production forUpdateOption is entered.
func (s *BaseDB2zSQLParserListener) EnterForUpdateOption(ctx *ForUpdateOptionContext) {}

// ExitForUpdateOption is called when production forUpdateOption is exited.
func (s *BaseDB2zSQLParserListener) ExitForUpdateOption(ctx *ForUpdateOptionContext) {}

// EnterConcentrateStatementsOption is called when production concentrateStatementsOption is entered.
func (s *BaseDB2zSQLParserListener) EnterConcentrateStatementsOption(ctx *ConcentrateStatementsOptionContext) {
}

// ExitConcentrateStatementsOption is called when production concentrateStatementsOption is exited.
func (s *BaseDB2zSQLParserListener) ExitConcentrateStatementsOption(ctx *ConcentrateStatementsOptionContext) {
}

// EnterAcceleratorOption is called when production acceleratorOption is entered.
func (s *BaseDB2zSQLParserListener) EnterAcceleratorOption(ctx *AcceleratorOptionContext) {}

// ExitAcceleratorOption is called when production acceleratorOption is exited.
func (s *BaseDB2zSQLParserListener) ExitAcceleratorOption(ctx *AcceleratorOptionContext) {}

// EnterProcedureDataType is called when production procedureDataType is entered.
func (s *BaseDB2zSQLParserListener) EnterProcedureDataType(ctx *ProcedureDataTypeContext) {}

// ExitProcedureDataType is called when production procedureDataType is exited.
func (s *BaseDB2zSQLParserListener) ExitProcedureDataType(ctx *ProcedureDataTypeContext) {}

// EnterAlterSequenceOptionList is called when production alterSequenceOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterSequenceOptionList(ctx *AlterSequenceOptionListContext) {
}

// ExitAlterSequenceOptionList is called when production alterSequenceOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterSequenceOptionList(ctx *AlterSequenceOptionListContext) {
}

// EnterCreateSequenceOptionList is called when production createSequenceOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateSequenceOptionList(ctx *CreateSequenceOptionListContext) {
}

// ExitCreateSequenceOptionList is called when production createSequenceOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateSequenceOptionList(ctx *CreateSequenceOptionListContext) {
}

// EnterAsTypeOption is called when production asTypeOption is entered.
func (s *BaseDB2zSQLParserListener) EnterAsTypeOption(ctx *AsTypeOptionContext) {}

// ExitAsTypeOption is called when production asTypeOption is exited.
func (s *BaseDB2zSQLParserListener) ExitAsTypeOption(ctx *AsTypeOptionContext) {}

// EnterStartOption is called when production startOption is entered.
func (s *BaseDB2zSQLParserListener) EnterStartOption(ctx *StartOptionContext) {}

// ExitStartOption is called when production startOption is exited.
func (s *BaseDB2zSQLParserListener) ExitStartOption(ctx *StartOptionContext) {}

// EnterRestartOption is called when production restartOption is entered.
func (s *BaseDB2zSQLParserListener) EnterRestartOption(ctx *RestartOptionContext) {}

// ExitRestartOption is called when production restartOption is exited.
func (s *BaseDB2zSQLParserListener) ExitRestartOption(ctx *RestartOptionContext) {}

// EnterIncrementOption is called when production incrementOption is entered.
func (s *BaseDB2zSQLParserListener) EnterIncrementOption(ctx *IncrementOptionContext) {}

// ExitIncrementOption is called when production incrementOption is exited.
func (s *BaseDB2zSQLParserListener) ExitIncrementOption(ctx *IncrementOptionContext) {}

// EnterMinvalueOption is called when production minvalueOption is entered.
func (s *BaseDB2zSQLParserListener) EnterMinvalueOption(ctx *MinvalueOptionContext) {}

// ExitMinvalueOption is called when production minvalueOption is exited.
func (s *BaseDB2zSQLParserListener) ExitMinvalueOption(ctx *MinvalueOptionContext) {}

// EnterMaxvalueOption is called when production maxvalueOption is entered.
func (s *BaseDB2zSQLParserListener) EnterMaxvalueOption(ctx *MaxvalueOptionContext) {}

// ExitMaxvalueOption is called when production maxvalueOption is exited.
func (s *BaseDB2zSQLParserListener) ExitMaxvalueOption(ctx *MaxvalueOptionContext) {}

// EnterCycleOption is called when production cycleOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCycleOption(ctx *CycleOptionContext) {}

// ExitCycleOption is called when production cycleOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCycleOption(ctx *CycleOptionContext) {}

// EnterCacheOption is called when production cacheOption is entered.
func (s *BaseDB2zSQLParserListener) EnterCacheOption(ctx *CacheOptionContext) {}

// ExitCacheOption is called when production cacheOption is exited.
func (s *BaseDB2zSQLParserListener) ExitCacheOption(ctx *CacheOptionContext) {}

// EnterOrderOption is called when production orderOption is entered.
func (s *BaseDB2zSQLParserListener) EnterOrderOption(ctx *OrderOptionContext) {}

// ExitOrderOption is called when production orderOption is exited.
func (s *BaseDB2zSQLParserListener) ExitOrderOption(ctx *OrderOptionContext) {}

// EnterKeyLabelOption is called when production keyLabelOption is entered.
func (s *BaseDB2zSQLParserListener) EnterKeyLabelOption(ctx *KeyLabelOptionContext) {}

// ExitKeyLabelOption is called when production keyLabelOption is exited.
func (s *BaseDB2zSQLParserListener) ExitKeyLabelOption(ctx *KeyLabelOptionContext) {}

// EnterDataclasOption is called when production dataclasOption is entered.
func (s *BaseDB2zSQLParserListener) EnterDataclasOption(ctx *DataclasOptionContext) {}

// ExitDataclasOption is called when production dataclasOption is exited.
func (s *BaseDB2zSQLParserListener) ExitDataclasOption(ctx *DataclasOptionContext) {}

// EnterMgmtclasOption is called when production mgmtclasOption is entered.
func (s *BaseDB2zSQLParserListener) EnterMgmtclasOption(ctx *MgmtclasOptionContext) {}

// ExitMgmtclasOption is called when production mgmtclasOption is exited.
func (s *BaseDB2zSQLParserListener) ExitMgmtclasOption(ctx *MgmtclasOptionContext) {}

// EnterStorclasOption is called when production storclasOption is entered.
func (s *BaseDB2zSQLParserListener) EnterStorclasOption(ctx *StorclasOptionContext) {}

// ExitStorclasOption is called when production storclasOption is exited.
func (s *BaseDB2zSQLParserListener) ExitStorclasOption(ctx *StorclasOptionContext) {}

// EnterAlterStogroupOptionList is called when production alterStogroupOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterStogroupOptionList(ctx *AlterStogroupOptionListContext) {
}

// ExitAlterStogroupOptionList is called when production alterStogroupOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterStogroupOptionList(ctx *AlterStogroupOptionListContext) {
}

// EnterAlterTableOptionList is called when production alterTableOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTableOptionList(ctx *AlterTableOptionListContext) {}

// ExitAlterTableOptionList is called when production alterTableOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTableOptionList(ctx *AlterTableOptionListContext) {}

// EnterAlterTablespaceOptionList is called when production alterTablespaceOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTablespaceOptionList(ctx *AlterTablespaceOptionListContext) {
}

// ExitAlterTablespaceOptionList is called when production alterTablespaceOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTablespaceOptionList(ctx *AlterTablespaceOptionListContext) {
}

// EnterCreateTablespaceOptionList is called when production createTablespaceOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateTablespaceOptionList(ctx *CreateTablespaceOptionListContext) {
}

// ExitCreateTablespaceOptionList is called when production createTablespaceOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateTablespaceOptionList(ctx *CreateTablespaceOptionListContext) {
}

// EnterTrustedContextOptionList is called when production trustedContextOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterTrustedContextOptionList(ctx *TrustedContextOptionListContext) {
}

// ExitTrustedContextOptionList is called when production trustedContextOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitTrustedContextOptionList(ctx *TrustedContextOptionListContext) {
}

// EnterDatabaseOptionList is called when production databaseOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterDatabaseOptionList(ctx *DatabaseOptionListContext) {}

// ExitDatabaseOptionList is called when production databaseOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitDatabaseOptionList(ctx *DatabaseOptionListContext) {}

// EnterCreateIndexOptionList is called when production createIndexOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateIndexOptionList(ctx *CreateIndexOptionListContext) {}

// ExitCreateIndexOptionList is called when production createIndexOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateIndexOptionList(ctx *CreateIndexOptionListContext) {}

// EnterCreateLobTablespaceOptionList is called when production createLobTablespaceOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterCreateLobTablespaceOptionList(ctx *CreateLobTablespaceOptionListContext) {
}

// ExitCreateLobTablespaceOptionList is called when production createLobTablespaceOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitCreateLobTablespaceOptionList(ctx *CreateLobTablespaceOptionListContext) {
}

// EnterInDatabaseOption is called when production inDatabaseOption is entered.
func (s *BaseDB2zSQLParserListener) EnterInDatabaseOption(ctx *InDatabaseOptionContext) {}

// ExitInDatabaseOption is called when production inDatabaseOption is exited.
func (s *BaseDB2zSQLParserListener) ExitInDatabaseOption(ctx *InDatabaseOptionContext) {}

// EnterSegsizeOption is called when production segsizeOption is entered.
func (s *BaseDB2zSQLParserListener) EnterSegsizeOption(ctx *SegsizeOptionContext) {}

// ExitSegsizeOption is called when production segsizeOption is exited.
func (s *BaseDB2zSQLParserListener) ExitSegsizeOption(ctx *SegsizeOptionContext) {}

// EnterNumpartsOption is called when production numpartsOption is entered.
func (s *BaseDB2zSQLParserListener) EnterNumpartsOption(ctx *NumpartsOptionContext) {}

// ExitNumpartsOption is called when production numpartsOption is exited.
func (s *BaseDB2zSQLParserListener) ExitNumpartsOption(ctx *NumpartsOptionContext) {}

// EnterPartitionByGrowthSpecification is called when production partitionByGrowthSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitionByGrowthSpecification(ctx *PartitionByGrowthSpecificationContext) {
}

// ExitPartitionByGrowthSpecification is called when production partitionByGrowthSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitionByGrowthSpecification(ctx *PartitionByGrowthSpecificationContext) {
}

// EnterPartitionByRangeSpecification is called when production partitionByRangeSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitionByRangeSpecification(ctx *PartitionByRangeSpecificationContext) {
}

// ExitPartitionByRangeSpecification is called when production partitionByRangeSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitionByRangeSpecification(ctx *PartitionByRangeSpecificationContext) {
}

// EnterPartitionByRangePartitionPhrase is called when production partitionByRangePartitionPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitionByRangePartitionPhrase(ctx *PartitionByRangePartitionPhraseContext) {
}

// ExitPartitionByRangePartitionPhrase is called when production partitionByRangePartitionPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitionByRangePartitionPhrase(ctx *PartitionByRangePartitionPhraseContext) {
}

// EnterInsertAlgorithmOption is called when production insertAlgorithmOption is entered.
func (s *BaseDB2zSQLParserListener) EnterInsertAlgorithmOption(ctx *InsertAlgorithmOptionContext) {}

// ExitInsertAlgorithmOption is called when production insertAlgorithmOption is exited.
func (s *BaseDB2zSQLParserListener) ExitInsertAlgorithmOption(ctx *InsertAlgorithmOptionContext) {}

// EnterMaxrowsOption is called when production maxrowsOption is entered.
func (s *BaseDB2zSQLParserListener) EnterMaxrowsOption(ctx *MaxrowsOptionContext) {}

// ExitMaxrowsOption is called when production maxrowsOption is exited.
func (s *BaseDB2zSQLParserListener) ExitMaxrowsOption(ctx *MaxrowsOptionContext) {}

// EnterMaxpartitionsOption is called when production maxpartitionsOption is entered.
func (s *BaseDB2zSQLParserListener) EnterMaxpartitionsOption(ctx *MaxpartitionsOptionContext) {}

// ExitMaxpartitionsOption is called when production maxpartitionsOption is exited.
func (s *BaseDB2zSQLParserListener) ExitMaxpartitionsOption(ctx *MaxpartitionsOptionContext) {}

// EnterUsingSpecification2 is called when production usingSpecification2 is entered.
func (s *BaseDB2zSQLParserListener) EnterUsingSpecification2(ctx *UsingSpecification2Context) {}

// ExitUsingSpecification2 is called when production usingSpecification2 is exited.
func (s *BaseDB2zSQLParserListener) ExitUsingSpecification2(ctx *UsingSpecification2Context) {}

// EnterStogroupOptions is called when production stogroupOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterStogroupOptions(ctx *StogroupOptionsContext) {}

// ExitStogroupOptions is called when production stogroupOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitStogroupOptions(ctx *StogroupOptionsContext) {}

// EnterXmlIndexSpecification is called when production xmlIndexSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlIndexSpecification(ctx *XmlIndexSpecificationContext) {}

// ExitXmlIndexSpecification is called when production xmlIndexSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlIndexSpecification(ctx *XmlIndexSpecificationContext) {}

// EnterXmlPatternClause is called when production xmlPatternClause is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlPatternClause(ctx *XmlPatternClauseContext) {}

// ExitXmlPatternClause is called when production xmlPatternClause is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlPatternClause(ctx *XmlPatternClauseContext) {}

// EnterAlterAttributesOptions is called when production alterAttributesOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterAttributesOptions(ctx *AlterAttributesOptionsContext) {}

// ExitAlterAttributesOptions is called when production alterAttributesOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterAttributesOptions(ctx *AlterAttributesOptionsContext) {}

// EnterAddAttributesOptions is called when production addAttributesOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterAddAttributesOptions(ctx *AddAttributesOptionsContext) {}

// ExitAddAttributesOptions is called when production addAttributesOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitAddAttributesOptions(ctx *AddAttributesOptionsContext) {}

// EnterDropAttributesOptions is called when production dropAttributesOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterDropAttributesOptions(ctx *DropAttributesOptionsContext) {}

// ExitDropAttributesOptions is called when production dropAttributesOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitDropAttributesOptions(ctx *DropAttributesOptionsContext) {}

// EnterIncludeColumnPhrase is called when production includeColumnPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterIncludeColumnPhrase(ctx *IncludeColumnPhraseContext) {}

// ExitIncludeColumnPhrase is called when production includeColumnPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitIncludeColumnPhrase(ctx *IncludeColumnPhraseContext) {}

// EnterUserClause is called when production userClause is entered.
func (s *BaseDB2zSQLParserListener) EnterUserClause(ctx *UserClauseContext) {}

// ExitUserClause is called when production userClause is exited.
func (s *BaseDB2zSQLParserListener) ExitUserClause(ctx *UserClauseContext) {}

// EnterUserClauseAddOptions is called when production userClauseAddOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterUserClauseAddOptions(ctx *UserClauseAddOptionsContext) {}

// ExitUserClauseAddOptions is called when production userClauseAddOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitUserClauseAddOptions(ctx *UserClauseAddOptionsContext) {}

// EnterUserClauseReplaceOptions is called when production userClauseReplaceOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterUserClauseReplaceOptions(ctx *UserClauseReplaceOptionsContext) {
}

// ExitUserClauseReplaceOptions is called when production userClauseReplaceOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitUserClauseReplaceOptions(ctx *UserClauseReplaceOptionsContext) {
}

// EnterUserClauseDropOptions is called when production userClauseDropOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterUserClauseDropOptions(ctx *UserClauseDropOptionsContext) {}

// ExitUserClauseDropOptions is called when production userClauseDropOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitUserClauseDropOptions(ctx *UserClauseDropOptionsContext) {}

// EnterUseOptions is called when production useOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterUseOptions(ctx *UseOptionsContext) {}

// ExitUseOptions is called when production useOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitUseOptions(ctx *UseOptionsContext) {}

// EnterAlterPartitionClause is called when production alterPartitionClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterPartitionClause(ctx *AlterPartitionClauseContext) {}

// ExitAlterPartitionClause is called when production alterPartitionClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterPartitionClause(ctx *AlterPartitionClauseContext) {}

// EnterUsingBlock is called when production usingBlock is entered.
func (s *BaseDB2zSQLParserListener) EnterUsingBlock(ctx *UsingBlockContext) {}

// ExitUsingBlock is called when production usingBlock is exited.
func (s *BaseDB2zSQLParserListener) ExitUsingBlock(ctx *UsingBlockContext) {}

// EnterFreeBlock is called when production freeBlock is entered.
func (s *BaseDB2zSQLParserListener) EnterFreeBlock(ctx *FreeBlockContext) {}

// ExitFreeBlock is called when production freeBlock is exited.
func (s *BaseDB2zSQLParserListener) ExitFreeBlock(ctx *FreeBlockContext) {}

// EnterMoveTableClause is called when production moveTableClause is entered.
func (s *BaseDB2zSQLParserListener) EnterMoveTableClause(ctx *MoveTableClauseContext) {}

// ExitMoveTableClause is called when production moveTableClause is exited.
func (s *BaseDB2zSQLParserListener) ExitMoveTableClause(ctx *MoveTableClauseContext) {}

// EnterGbpcacheBlock is called when production gbpcacheBlock is entered.
func (s *BaseDB2zSQLParserListener) EnterGbpcacheBlock(ctx *GbpcacheBlockContext) {}

// ExitGbpcacheBlock is called when production gbpcacheBlock is exited.
func (s *BaseDB2zSQLParserListener) ExitGbpcacheBlock(ctx *GbpcacheBlockContext) {}

// EnterAliasDesignator is called when production aliasDesignator is entered.
func (s *BaseDB2zSQLParserListener) EnterAliasDesignator(ctx *AliasDesignatorContext) {}

// ExitAliasDesignator is called when production aliasDesignator is exited.
func (s *BaseDB2zSQLParserListener) ExitAliasDesignator(ctx *AliasDesignatorContext) {}

// EnterMultipleColumnList is called when production multipleColumnList is entered.
func (s *BaseDB2zSQLParserListener) EnterMultipleColumnList(ctx *MultipleColumnListContext) {}

// ExitMultipleColumnList is called when production multipleColumnList is exited.
func (s *BaseDB2zSQLParserListener) ExitMultipleColumnList(ctx *MultipleColumnListContext) {}

// EnterFunctionDesignator is called when production functionDesignator is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionDesignator(ctx *FunctionDesignatorContext) {}

// ExitFunctionDesignator is called when production functionDesignator is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionDesignator(ctx *FunctionDesignatorContext) {}

// EnterParameterType is called when production parameterType is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterType(ctx *ParameterTypeContext) {}

// ExitParameterType is called when production parameterType is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterType(ctx *ParameterTypeContext) {}

// EnterAlterTableColumnDefinitionOptionList1 is called when production alterTableColumnDefinitionOptionList1 is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTableColumnDefinitionOptionList1(ctx *AlterTableColumnDefinitionOptionList1Context) {
}

// ExitAlterTableColumnDefinitionOptionList1 is called when production alterTableColumnDefinitionOptionList1 is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTableColumnDefinitionOptionList1(ctx *AlterTableColumnDefinitionOptionList1Context) {
}

// EnterAlterTableColumnDefinitionOptionList2 is called when production alterTableColumnDefinitionOptionList2 is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTableColumnDefinitionOptionList2(ctx *AlterTableColumnDefinitionOptionList2Context) {
}

// ExitAlterTableColumnDefinitionOptionList2 is called when production alterTableColumnDefinitionOptionList2 is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTableColumnDefinitionOptionList2(ctx *AlterTableColumnDefinitionOptionList2Context) {
}

// EnterColumnConstraint is called when production columnConstraint is entered.
func (s *BaseDB2zSQLParserListener) EnterColumnConstraint(ctx *ColumnConstraintContext) {}

// ExitColumnConstraint is called when production columnConstraint is exited.
func (s *BaseDB2zSQLParserListener) ExitColumnConstraint(ctx *ColumnConstraintContext) {}

// EnterGeneratedClause is called when production generatedClause is entered.
func (s *BaseDB2zSQLParserListener) EnterGeneratedClause(ctx *GeneratedClauseContext) {}

// ExitGeneratedClause is called when production generatedClause is exited.
func (s *BaseDB2zSQLParserListener) ExitGeneratedClause(ctx *GeneratedClauseContext) {}

// EnterGeneratedClause2 is called when production generatedClause2 is entered.
func (s *BaseDB2zSQLParserListener) EnterGeneratedClause2(ctx *GeneratedClause2Context) {}

// ExitGeneratedClause2 is called when production generatedClause2 is exited.
func (s *BaseDB2zSQLParserListener) ExitGeneratedClause2(ctx *GeneratedClause2Context) {}

// EnterAsIdentityClause is called when production asIdentityClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAsIdentityClause(ctx *AsIdentityClauseContext) {}

// ExitAsIdentityClause is called when production asIdentityClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAsIdentityClause(ctx *AsIdentityClauseContext) {}

// EnterAsIdentityClauseOptionList is called when production asIdentityClauseOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterAsIdentityClauseOptionList(ctx *AsIdentityClauseOptionListContext) {
}

// ExitAsIdentityClauseOptionList is called when production asIdentityClauseOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitAsIdentityClauseOptionList(ctx *AsIdentityClauseOptionListContext) {
}

// EnterAsRowChangeTimestampClause is called when production asRowChangeTimestampClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAsRowChangeTimestampClause(ctx *AsRowChangeTimestampClauseContext) {
}

// ExitAsRowChangeTimestampClause is called when production asRowChangeTimestampClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAsRowChangeTimestampClause(ctx *AsRowChangeTimestampClauseContext) {
}

// EnterAsRowTransactionStartIDClause is called when production asRowTransactionStartIDClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAsRowTransactionStartIDClause(ctx *AsRowTransactionStartIDClauseContext) {
}

// ExitAsRowTransactionStartIDClause is called when production asRowTransactionStartIDClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAsRowTransactionStartIDClause(ctx *AsRowTransactionStartIDClauseContext) {
}

// EnterAsRowTransactionTimestampClause is called when production asRowTransactionTimestampClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAsRowTransactionTimestampClause(ctx *AsRowTransactionTimestampClauseContext) {
}

// ExitAsRowTransactionTimestampClause is called when production asRowTransactionTimestampClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAsRowTransactionTimestampClause(ctx *AsRowTransactionTimestampClauseContext) {
}

// EnterAsGeneratedExpressionClause is called when production asGeneratedExpressionClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAsGeneratedExpressionClause(ctx *AsGeneratedExpressionClauseContext) {
}

// ExitAsGeneratedExpressionClause is called when production asGeneratedExpressionClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAsGeneratedExpressionClause(ctx *AsGeneratedExpressionClauseContext) {
}

// EnterNonDeterministicExpression is called when production nonDeterministicExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterNonDeterministicExpression(ctx *NonDeterministicExpressionContext) {
}

// ExitNonDeterministicExpression is called when production nonDeterministicExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitNonDeterministicExpression(ctx *NonDeterministicExpressionContext) {
}

// EnterNonDeterministicExpressionSessionVariable is called when production nonDeterministicExpressionSessionVariable is entered.
func (s *BaseDB2zSQLParserListener) EnterNonDeterministicExpressionSessionVariable(ctx *NonDeterministicExpressionSessionVariableContext) {
}

// ExitNonDeterministicExpressionSessionVariable is called when production nonDeterministicExpressionSessionVariable is exited.
func (s *BaseDB2zSQLParserListener) ExitNonDeterministicExpressionSessionVariable(ctx *NonDeterministicExpressionSessionVariableContext) {
}

// EnterColumnAlteration is called when production columnAlteration is entered.
func (s *BaseDB2zSQLParserListener) EnterColumnAlteration(ctx *ColumnAlterationContext) {}

// ExitColumnAlteration is called when production columnAlteration is exited.
func (s *BaseDB2zSQLParserListener) ExitColumnAlteration(ctx *ColumnAlterationContext) {}

// EnterColumnAlterationOptionList is called when production columnAlterationOptionList is entered.
func (s *BaseDB2zSQLParserListener) EnterColumnAlterationOptionList(ctx *ColumnAlterationOptionListContext) {
}

// ExitColumnAlterationOptionList is called when production columnAlterationOptionList is exited.
func (s *BaseDB2zSQLParserListener) ExitColumnAlterationOptionList(ctx *ColumnAlterationOptionListContext) {
}

// EnterAlteredDataType is called when production alteredDataType is entered.
func (s *BaseDB2zSQLParserListener) EnterAlteredDataType(ctx *AlteredDataTypeContext) {}

// ExitAlteredDataType is called when production alteredDataType is exited.
func (s *BaseDB2zSQLParserListener) ExitAlteredDataType(ctx *AlteredDataTypeContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseDB2zSQLParserListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseDB2zSQLParserListener) ExitDataType(ctx *DataTypeContext) {}

// EnterBuiltInType is called when production builtInType is entered.
func (s *BaseDB2zSQLParserListener) EnterBuiltInType(ctx *BuiltInTypeContext) {}

// ExitBuiltInType is called when production builtInType is exited.
func (s *BaseDB2zSQLParserListener) ExitBuiltInType(ctx *BuiltInTypeContext) {}

// EnterSequenceDataType is called when production sequenceDataType is entered.
func (s *BaseDB2zSQLParserListener) EnterSequenceDataType(ctx *SequenceDataTypeContext) {}

// ExitSequenceDataType is called when production sequenceDataType is exited.
func (s *BaseDB2zSQLParserListener) ExitSequenceDataType(ctx *SequenceDataTypeContext) {}

// EnterSequenceBuiltInType is called when production sequenceBuiltInType is entered.
func (s *BaseDB2zSQLParserListener) EnterSequenceBuiltInType(ctx *SequenceBuiltInTypeContext) {}

// ExitSequenceBuiltInType is called when production sequenceBuiltInType is exited.
func (s *BaseDB2zSQLParserListener) ExitSequenceBuiltInType(ctx *SequenceBuiltInTypeContext) {}

// EnterSqlDataType is called when production sqlDataType is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlDataType(ctx *SqlDataTypeContext) {}

// ExitSqlDataType is called when production sqlDataType is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlDataType(ctx *SqlDataTypeContext) {}

// EnterXmlTypeModifier is called when production xmlTypeModifier is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlTypeModifier(ctx *XmlTypeModifierContext) {}

// ExitXmlTypeModifier is called when production xmlTypeModifier is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlTypeModifier(ctx *XmlTypeModifierContext) {}

// EnterXmlSchemaSpecification is called when production xmlSchemaSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlSchemaSpecification(ctx *XmlSchemaSpecificationContext) {}

// ExitXmlSchemaSpecification is called when production xmlSchemaSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlSchemaSpecification(ctx *XmlSchemaSpecificationContext) {}

// EnterXmlElementName is called when production xmlElementName is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlElementName(ctx *XmlElementNameContext) {}

// ExitXmlElementName is called when production xmlElementName is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlElementName(ctx *XmlElementNameContext) {}

// EnterPiName is called when production piName is entered.
func (s *BaseDB2zSQLParserListener) EnterPiName(ctx *PiNameContext) {}

// ExitPiName is called when production piName is exited.
func (s *BaseDB2zSQLParserListener) ExitPiName(ctx *PiNameContext) {}

// EnterRegisteredXmlSchemaName is called when production registeredXmlSchemaName is entered.
func (s *BaseDB2zSQLParserListener) EnterRegisteredXmlSchemaName(ctx *RegisteredXmlSchemaNameContext) {
}

// ExitRegisteredXmlSchemaName is called when production registeredXmlSchemaName is exited.
func (s *BaseDB2zSQLParserListener) ExitRegisteredXmlSchemaName(ctx *RegisteredXmlSchemaNameContext) {
}

// EnterTargetNamespace is called when production targetNamespace is entered.
func (s *BaseDB2zSQLParserListener) EnterTargetNamespace(ctx *TargetNamespaceContext) {}

// ExitTargetNamespace is called when production targetNamespace is exited.
func (s *BaseDB2zSQLParserListener) ExitTargetNamespace(ctx *TargetNamespaceContext) {}

// EnterSchemaLocation is called when production schemaLocation is entered.
func (s *BaseDB2zSQLParserListener) EnterSchemaLocation(ctx *SchemaLocationContext) {}

// ExitSchemaLocation is called when production schemaLocation is exited.
func (s *BaseDB2zSQLParserListener) ExitSchemaLocation(ctx *SchemaLocationContext) {}

// EnterIdentityAlteration is called when production identityAlteration is entered.
func (s *BaseDB2zSQLParserListener) EnterIdentityAlteration(ctx *IdentityAlterationContext) {}

// ExitIdentityAlteration is called when production identityAlteration is exited.
func (s *BaseDB2zSQLParserListener) ExitIdentityAlteration(ctx *IdentityAlterationContext) {}

// EnterUniqueConstraint is called when production uniqueConstraint is entered.
func (s *BaseDB2zSQLParserListener) EnterUniqueConstraint(ctx *UniqueConstraintContext) {}

// ExitUniqueConstraint is called when production uniqueConstraint is exited.
func (s *BaseDB2zSQLParserListener) ExitUniqueConstraint(ctx *UniqueConstraintContext) {}

// EnterReferentialConstraint is called when production referentialConstraint is entered.
func (s *BaseDB2zSQLParserListener) EnterReferentialConstraint(ctx *ReferentialConstraintContext) {}

// ExitReferentialConstraint is called when production referentialConstraint is exited.
func (s *BaseDB2zSQLParserListener) ExitReferentialConstraint(ctx *ReferentialConstraintContext) {}

// EnterReferencesClause is called when production referencesClause is entered.
func (s *BaseDB2zSQLParserListener) EnterReferencesClause(ctx *ReferencesClauseContext) {}

// ExitReferencesClause is called when production referencesClause is exited.
func (s *BaseDB2zSQLParserListener) ExitReferencesClause(ctx *ReferencesClauseContext) {}

// EnterCheckConstraint is called when production checkConstraint is entered.
func (s *BaseDB2zSQLParserListener) EnterCheckConstraint(ctx *CheckConstraintContext) {}

// ExitCheckConstraint is called when production checkConstraint is exited.
func (s *BaseDB2zSQLParserListener) ExitCheckConstraint(ctx *CheckConstraintContext) {}

// EnterPartitioningClause is called when production partitioningClause is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitioningClause(ctx *PartitioningClauseContext) {}

// ExitPartitioningClause is called when production partitioningClause is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitioningClause(ctx *PartitioningClauseContext) {}

// EnterPartitionExpression is called when production partitionExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitionExpression(ctx *PartitionExpressionContext) {}

// ExitPartitionExpression is called when production partitionExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitionExpression(ctx *PartitionExpressionContext) {}

// EnterPartitionLimitKey is called when production partitionLimitKey is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitionLimitKey(ctx *PartitionLimitKeyContext) {}

// ExitPartitionLimitKey is called when production partitionLimitKey is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitionLimitKey(ctx *PartitionLimitKeyContext) {}

// EnterPartitioningPhrase is called when production partitioningPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitioningPhrase(ctx *PartitioningPhraseContext) {}

// ExitPartitioningPhrase is called when production partitioningPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitioningPhrase(ctx *PartitioningPhraseContext) {}

// EnterPartitionHashSpace is called when production partitionHashSpace is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitionHashSpace(ctx *PartitionHashSpaceContext) {}

// ExitPartitionHashSpace is called when production partitionHashSpace is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitionHashSpace(ctx *PartitionHashSpaceContext) {}

// EnterAlterHashOrganization is called when production alterHashOrganization is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterHashOrganization(ctx *AlterHashOrganizationContext) {}

// ExitAlterHashOrganization is called when production alterHashOrganization is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterHashOrganization(ctx *AlterHashOrganizationContext) {}

// EnterPartitioningClauseElement is called when production partitioningClauseElement is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitioningClauseElement(ctx *PartitioningClauseElementContext) {
}

// ExitPartitioningClauseElement is called when production partitioningClauseElement is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitioningClauseElement(ctx *PartitioningClauseElementContext) {
}

// EnterPartitionClause is called when production partitionClause is entered.
func (s *BaseDB2zSQLParserListener) EnterPartitionClause(ctx *PartitionClauseContext) {}

// ExitPartitionClause is called when production partitionClause is exited.
func (s *BaseDB2zSQLParserListener) ExitPartitionClause(ctx *PartitionClauseContext) {}

// EnterRotatePartitionClause is called when production rotatePartitionClause is entered.
func (s *BaseDB2zSQLParserListener) EnterRotatePartitionClause(ctx *RotatePartitionClauseContext) {}

// ExitRotatePartitionClause is called when production rotatePartitionClause is exited.
func (s *BaseDB2zSQLParserListener) ExitRotatePartitionClause(ctx *RotatePartitionClauseContext) {}

// EnterExtraRowOption is called when production extraRowOption is entered.
func (s *BaseDB2zSQLParserListener) EnterExtraRowOption(ctx *ExtraRowOptionContext) {}

// ExitExtraRowOption is called when production extraRowOption is exited.
func (s *BaseDB2zSQLParserListener) ExitExtraRowOption(ctx *ExtraRowOptionContext) {}

// EnterMaterializedQueryDefinition is called when production materializedQueryDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterMaterializedQueryDefinition(ctx *MaterializedQueryDefinitionContext) {
}

// ExitMaterializedQueryDefinition is called when production materializedQueryDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitMaterializedQueryDefinition(ctx *MaterializedQueryDefinitionContext) {
}

// EnterMaterializedQueryAlteration is called when production materializedQueryAlteration is entered.
func (s *BaseDB2zSQLParserListener) EnterMaterializedQueryAlteration(ctx *MaterializedQueryAlterationContext) {
}

// ExitMaterializedQueryAlteration is called when production materializedQueryAlteration is exited.
func (s *BaseDB2zSQLParserListener) ExitMaterializedQueryAlteration(ctx *MaterializedQueryAlterationContext) {
}

// EnterRefreshableTableOptions is called when production refreshableTableOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterRefreshableTableOptions(ctx *RefreshableTableOptionsContext) {
}

// ExitRefreshableTableOptions is called when production refreshableTableOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitRefreshableTableOptions(ctx *RefreshableTableOptionsContext) {
}

// EnterDataInitiallyDeferredPhrase is called when production dataInitiallyDeferredPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterDataInitiallyDeferredPhrase(ctx *DataInitiallyDeferredPhraseContext) {
}

// ExitDataInitiallyDeferredPhrase is called when production dataInitiallyDeferredPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitDataInitiallyDeferredPhrase(ctx *DataInitiallyDeferredPhraseContext) {
}

// EnterRefreshDeferredPhrase is called when production refreshDeferredPhrase is entered.
func (s *BaseDB2zSQLParserListener) EnterRefreshDeferredPhrase(ctx *RefreshDeferredPhraseContext) {}

// ExitRefreshDeferredPhrase is called when production refreshDeferredPhrase is exited.
func (s *BaseDB2zSQLParserListener) ExitRefreshDeferredPhrase(ctx *RefreshDeferredPhraseContext) {}

// EnterRefreshableTableOptionsList is called when production refreshableTableOptionsList is entered.
func (s *BaseDB2zSQLParserListener) EnterRefreshableTableOptionsList(ctx *RefreshableTableOptionsListContext) {
}

// ExitRefreshableTableOptionsList is called when production refreshableTableOptionsList is exited.
func (s *BaseDB2zSQLParserListener) ExitRefreshableTableOptionsList(ctx *RefreshableTableOptionsListContext) {
}

// EnterMaterializedQueryTableAlteration is called when production materializedQueryTableAlteration is entered.
func (s *BaseDB2zSQLParserListener) EnterMaterializedQueryTableAlteration(ctx *MaterializedQueryTableAlterationContext) {
}

// ExitMaterializedQueryTableAlteration is called when production materializedQueryTableAlteration is exited.
func (s *BaseDB2zSQLParserListener) ExitMaterializedQueryTableAlteration(ctx *MaterializedQueryTableAlterationContext) {
}

// EnterPeriodDefinition is called when production periodDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterPeriodDefinition(ctx *PeriodDefinitionContext) {}

// ExitPeriodDefinition is called when production periodDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitPeriodDefinition(ctx *PeriodDefinitionContext) {}

// EnterAlterTableColumnDefinition is called when production alterTableColumnDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTableColumnDefinition(ctx *AlterTableColumnDefinitionContext) {
}

// ExitAlterTableColumnDefinition is called when production alterTableColumnDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTableColumnDefinition(ctx *AlterTableColumnDefinitionContext) {
}

// EnterExternalProgramName is called when production externalProgramName is entered.
func (s *BaseDB2zSQLParserListener) EnterExternalProgramName(ctx *ExternalProgramNameContext) {}

// ExitExternalProgramName is called when production externalProgramName is exited.
func (s *BaseDB2zSQLParserListener) ExitExternalProgramName(ctx *ExternalProgramNameContext) {}

// EnterPackagePath is called when production packagePath is entered.
func (s *BaseDB2zSQLParserListener) EnterPackagePath(ctx *PackagePathContext) {}

// ExitPackagePath is called when production packagePath is exited.
func (s *BaseDB2zSQLParserListener) ExitPackagePath(ctx *PackagePathContext) {}

// EnterCollectionID is called when production collectionID is entered.
func (s *BaseDB2zSQLParserListener) EnterCollectionID(ctx *CollectionIDContext) {}

// ExitCollectionID is called when production collectionID is exited.
func (s *BaseDB2zSQLParserListener) ExitCollectionID(ctx *CollectionIDContext) {}

// EnterRunTimeOptions is called when production runTimeOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterRunTimeOptions(ctx *RunTimeOptionsContext) {}

// ExitRunTimeOptions is called when production runTimeOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitRunTimeOptions(ctx *RunTimeOptionsContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseDB2zSQLParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseDB2zSQLParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseDB2zSQLParserListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseDB2zSQLParserListener) ExitOperator(ctx *OperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDB2zSQLParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDB2zSQLParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterKeyExpression is called when production keyExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterKeyExpression(ctx *KeyExpressionContext) {}

// ExitKeyExpression is called when production keyExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitKeyExpression(ctx *KeyExpressionContext) {}

// EnterRowChangeExpression is called when production rowChangeExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterRowChangeExpression(ctx *RowChangeExpressionContext) {}

// ExitRowChangeExpression is called when production rowChangeExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitRowChangeExpression(ctx *RowChangeExpressionContext) {}

// EnterSequenceReference is called when production sequenceReference is entered.
func (s *BaseDB2zSQLParserListener) EnterSequenceReference(ctx *SequenceReferenceContext) {}

// ExitSequenceReference is called when production sequenceReference is exited.
func (s *BaseDB2zSQLParserListener) ExitSequenceReference(ctx *SequenceReferenceContext) {}

// EnterFunctionInvocation is called when production functionInvocation is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionInvocation(ctx *FunctionInvocationContext) {}

// ExitFunctionInvocation is called when production functionInvocation is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionInvocation(ctx *FunctionInvocationContext) {}

// EnterScalarFunctionInvocation is called when production scalarFunctionInvocation is entered.
func (s *BaseDB2zSQLParserListener) EnterScalarFunctionInvocation(ctx *ScalarFunctionInvocationContext) {
}

// ExitScalarFunctionInvocation is called when production scalarFunctionInvocation is exited.
func (s *BaseDB2zSQLParserListener) ExitScalarFunctionInvocation(ctx *ScalarFunctionInvocationContext) {
}

// EnterAggregateFunctionInvocation is called when production aggregateFunctionInvocation is entered.
func (s *BaseDB2zSQLParserListener) EnterAggregateFunctionInvocation(ctx *AggregateFunctionInvocationContext) {
}

// ExitAggregateFunctionInvocation is called when production aggregateFunctionInvocation is exited.
func (s *BaseDB2zSQLParserListener) ExitAggregateFunctionInvocation(ctx *AggregateFunctionInvocationContext) {
}

// EnterRegressionFunctionInvocation is called when production regressionFunctionInvocation is entered.
func (s *BaseDB2zSQLParserListener) EnterRegressionFunctionInvocation(ctx *RegressionFunctionInvocationContext) {
}

// ExitRegressionFunctionInvocation is called when production regressionFunctionInvocation is exited.
func (s *BaseDB2zSQLParserListener) ExitRegressionFunctionInvocation(ctx *RegressionFunctionInvocationContext) {
}

// EnterExternalFunctionInvocation is called when production externalFunctionInvocation is entered.
func (s *BaseDB2zSQLParserListener) EnterExternalFunctionInvocation(ctx *ExternalFunctionInvocationContext) {
}

// ExitExternalFunctionInvocation is called when production externalFunctionInvocation is exited.
func (s *BaseDB2zSQLParserListener) ExitExternalFunctionInvocation(ctx *ExternalFunctionInvocationContext) {
}

// EnterLabeledDuration is called when production labeledDuration is entered.
func (s *BaseDB2zSQLParserListener) EnterLabeledDuration(ctx *LabeledDurationContext) {}

// ExitLabeledDuration is called when production labeledDuration is exited.
func (s *BaseDB2zSQLParserListener) ExitLabeledDuration(ctx *LabeledDurationContext) {}

// EnterDurationSuffix is called when production durationSuffix is entered.
func (s *BaseDB2zSQLParserListener) EnterDurationSuffix(ctx *DurationSuffixContext) {}

// ExitDurationSuffix is called when production durationSuffix is exited.
func (s *BaseDB2zSQLParserListener) ExitDurationSuffix(ctx *DurationSuffixContext) {}

// EnterXmlCastSpecification is called when production xmlCastSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlCastSpecification(ctx *XmlCastSpecificationContext) {}

// ExitXmlCastSpecification is called when production xmlCastSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlCastSpecification(ctx *XmlCastSpecificationContext) {}

// EnterXmlParseSpecification is called when production xmlParseSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlParseSpecification(ctx *XmlParseSpecificationContext) {}

// ExitXmlParseSpecification is called when production xmlParseSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlParseSpecification(ctx *XmlParseSpecificationContext) {}

// EnterArrayElementSpecification is called when production arrayElementSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayElementSpecification(ctx *ArrayElementSpecificationContext) {
}

// ExitArrayElementSpecification is called when production arrayElementSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayElementSpecification(ctx *ArrayElementSpecificationContext) {
}

// EnterArrayIndex is called when production arrayIndex is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayIndex(ctx *ArrayIndexContext) {}

// ExitArrayIndex is called when production arrayIndex is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayIndex(ctx *ArrayIndexContext) {}

// EnterArrayConstructor is called when production arrayConstructor is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayConstructor(ctx *ArrayConstructorContext) {}

// ExitArrayConstructor is called when production arrayConstructor is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayConstructor(ctx *ArrayConstructorContext) {}

// EnterOlapSpecification is called when production olapSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterOlapSpecification(ctx *OlapSpecificationContext) {}

// ExitOlapSpecification is called when production olapSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitOlapSpecification(ctx *OlapSpecificationContext) {}

// EnterOrderedOlapSpecification is called when production orderedOlapSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterOrderedOlapSpecification(ctx *OrderedOlapSpecificationContext) {
}

// ExitOrderedOlapSpecification is called when production orderedOlapSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitOrderedOlapSpecification(ctx *OrderedOlapSpecificationContext) {
}

// EnterOlapSpecificationFunction is called when production olapSpecificationFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterOlapSpecificationFunction(ctx *OlapSpecificationFunctionContext) {
}

// ExitOlapSpecificationFunction is called when production olapSpecificationFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitOlapSpecificationFunction(ctx *OlapSpecificationFunctionContext) {
}

// EnterLagFunction is called when production lagFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterLagFunction(ctx *LagFunctionContext) {}

// ExitLagFunction is called when production lagFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitLagFunction(ctx *LagFunctionContext) {}

// EnterLeadFunction is called when production leadFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterLeadFunction(ctx *LeadFunctionContext) {}

// ExitLeadFunction is called when production leadFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitLeadFunction(ctx *LeadFunctionContext) {}

// EnterRespectNullsClause is called when production respectNullsClause is entered.
func (s *BaseDB2zSQLParserListener) EnterRespectNullsClause(ctx *RespectNullsClauseContext) {}

// ExitRespectNullsClause is called when production respectNullsClause is exited.
func (s *BaseDB2zSQLParserListener) ExitRespectNullsClause(ctx *RespectNullsClauseContext) {}

// EnterWindowPartitionClause is called when production windowPartitionClause is entered.
func (s *BaseDB2zSQLParserListener) EnterWindowPartitionClause(ctx *WindowPartitionClauseContext) {}

// ExitWindowPartitionClause is called when production windowPartitionClause is exited.
func (s *BaseDB2zSQLParserListener) ExitWindowPartitionClause(ctx *WindowPartitionClauseContext) {}

// EnterWindowOrderClause is called when production windowOrderClause is entered.
func (s *BaseDB2zSQLParserListener) EnterWindowOrderClause(ctx *WindowOrderClauseContext) {}

// ExitWindowOrderClause is called when production windowOrderClause is exited.
func (s *BaseDB2zSQLParserListener) ExitWindowOrderClause(ctx *WindowOrderClauseContext) {}

// EnterWindowOrderClauseQualifier is called when production windowOrderClauseQualifier is entered.
func (s *BaseDB2zSQLParserListener) EnterWindowOrderClauseQualifier(ctx *WindowOrderClauseQualifierContext) {
}

// ExitWindowOrderClauseQualifier is called when production windowOrderClauseQualifier is exited.
func (s *BaseDB2zSQLParserListener) ExitWindowOrderClauseQualifier(ctx *WindowOrderClauseQualifierContext) {
}

// EnterNumberingSpecification is called when production numberingSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterNumberingSpecification(ctx *NumberingSpecificationContext) {}

// ExitNumberingSpecification is called when production numberingSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitNumberingSpecification(ctx *NumberingSpecificationContext) {}

// EnterAggregationSpecification is called when production aggregationSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterAggregationSpecification(ctx *AggregationSpecificationContext) {
}

// ExitAggregationSpecification is called when production aggregationSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitAggregationSpecification(ctx *AggregationSpecificationContext) {
}

// EnterAggregateFunction is called when production aggregateFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterAggregateFunction(ctx *AggregateFunctionContext) {}

// ExitAggregateFunction is called when production aggregateFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitAggregateFunction(ctx *AggregateFunctionContext) {}

// EnterRegressionFunction is called when production regressionFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterRegressionFunction(ctx *RegressionFunctionContext) {}

// ExitRegressionFunction is called when production regressionFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitRegressionFunction(ctx *RegressionFunctionContext) {}

// EnterOlapColumnFunction is called when production olapColumnFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterOlapColumnFunction(ctx *OlapColumnFunctionContext) {}

// ExitOlapColumnFunction is called when production olapColumnFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitOlapColumnFunction(ctx *OlapColumnFunctionContext) {}

// EnterFirstValueFunction is called when production firstValueFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterFirstValueFunction(ctx *FirstValueFunctionContext) {}

// ExitFirstValueFunction is called when production firstValueFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitFirstValueFunction(ctx *FirstValueFunctionContext) {}

// EnterLastValueFunction is called when production lastValueFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterLastValueFunction(ctx *LastValueFunctionContext) {}

// ExitLastValueFunction is called when production lastValueFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitLastValueFunction(ctx *LastValueFunctionContext) {}

// EnterNthValueFunction is called when production nthValueFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterNthValueFunction(ctx *NthValueFunctionContext) {}

// ExitNthValueFunction is called when production nthValueFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitNthValueFunction(ctx *NthValueFunctionContext) {}

// EnterRatioToReportFunction is called when production ratioToReportFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterRatioToReportFunction(ctx *RatioToReportFunctionContext) {}

// ExitRatioToReportFunction is called when production ratioToReportFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitRatioToReportFunction(ctx *RatioToReportFunctionContext) {}

// EnterListaggFunction is called when production listaggFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterListaggFunction(ctx *ListaggFunctionContext) {}

// ExitListaggFunction is called when production listaggFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitListaggFunction(ctx *ListaggFunctionContext) {}

// EnterArrayaggFunction is called when production arrayaggFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayaggFunction(ctx *ArrayaggFunctionContext) {}

// ExitArrayaggFunction is called when production arrayaggFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayaggFunction(ctx *ArrayaggFunctionContext) {}

// EnterArrayaggOrdinaryFunction is called when production arrayaggOrdinaryFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayaggOrdinaryFunction(ctx *ArrayaggOrdinaryFunctionContext) {
}

// ExitArrayaggOrdinaryFunction is called when production arrayaggOrdinaryFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayaggOrdinaryFunction(ctx *ArrayaggOrdinaryFunctionContext) {
}

// EnterArrayaggAssociativeFunction is called when production arrayaggAssociativeFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayaggAssociativeFunction(ctx *ArrayaggAssociativeFunctionContext) {
}

// ExitArrayaggAssociativeFunction is called when production arrayaggAssociativeFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayaggAssociativeFunction(ctx *ArrayaggAssociativeFunctionContext) {
}

// EnterCorrelationFunction is called when production correlationFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterCorrelationFunction(ctx *CorrelationFunctionContext) {}

// ExitCorrelationFunction is called when production correlationFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitCorrelationFunction(ctx *CorrelationFunctionContext) {}

// EnterCovarianceFunction is called when production covarianceFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterCovarianceFunction(ctx *CovarianceFunctionContext) {}

// ExitCovarianceFunction is called when production covarianceFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitCovarianceFunction(ctx *CovarianceFunctionContext) {}

// EnterCovarianceSampFunction is called when production covarianceSampFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterCovarianceSampFunction(ctx *CovarianceSampFunctionContext) {}

// ExitCovarianceSampFunction is called when production covarianceSampFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitCovarianceSampFunction(ctx *CovarianceSampFunctionContext) {}

// EnterCumeDistFunction is called when production cumeDistFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterCumeDistFunction(ctx *CumeDistFunctionContext) {}

// ExitCumeDistFunction is called when production cumeDistFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitCumeDistFunction(ctx *CumeDistFunctionContext) {}

// EnterPercentileContFunction is called when production percentileContFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterPercentileContFunction(ctx *PercentileContFunctionContext) {}

// ExitPercentileContFunction is called when production percentileContFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitPercentileContFunction(ctx *PercentileContFunctionContext) {}

// EnterPercentileDiscFunction is called when production percentileDiscFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterPercentileDiscFunction(ctx *PercentileDiscFunctionContext) {}

// ExitPercentileDiscFunction is called when production percentileDiscFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitPercentileDiscFunction(ctx *PercentileDiscFunctionContext) {}

// EnterPercentRankFunction is called when production percentRankFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterPercentRankFunction(ctx *PercentRankFunctionContext) {}

// ExitPercentRankFunction is called when production percentRankFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitPercentRankFunction(ctx *PercentRankFunctionContext) {}

// EnterXmlaggFunction is called when production xmlaggFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlaggFunction(ctx *XmlaggFunctionContext) {}

// ExitXmlaggFunction is called when production xmlaggFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlaggFunction(ctx *XmlaggFunctionContext) {}

// EnterXmlaggOrderByClause is called when production xmlaggOrderByClause is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlaggOrderByClause(ctx *XmlaggOrderByClauseContext) {}

// ExitXmlaggOrderByClause is called when production xmlaggOrderByClause is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlaggOrderByClause(ctx *XmlaggOrderByClauseContext) {}

// EnterXmlaggOrderByOption is called when production xmlaggOrderByOption is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlaggOrderByOption(ctx *XmlaggOrderByOptionContext) {}

// ExitXmlaggOrderByOption is called when production xmlaggOrderByOption is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlaggOrderByOption(ctx *XmlaggOrderByOptionContext) {}

// EnterAggregateOrderByClause is called when production aggregateOrderByClause is entered.
func (s *BaseDB2zSQLParserListener) EnterAggregateOrderByClause(ctx *AggregateOrderByClauseContext) {}

// ExitAggregateOrderByClause is called when production aggregateOrderByClause is exited.
func (s *BaseDB2zSQLParserListener) ExitAggregateOrderByClause(ctx *AggregateOrderByClauseContext) {}

// EnterAggregateOrderByOption is called when production aggregateOrderByOption is entered.
func (s *BaseDB2zSQLParserListener) EnterAggregateOrderByOption(ctx *AggregateOrderByOptionContext) {}

// ExitAggregateOrderByOption is called when production aggregateOrderByOption is exited.
func (s *BaseDB2zSQLParserListener) ExitAggregateOrderByOption(ctx *AggregateOrderByOptionContext) {}

// EnterWindowAggregationGroupClause is called when production windowAggregationGroupClause is entered.
func (s *BaseDB2zSQLParserListener) EnterWindowAggregationGroupClause(ctx *WindowAggregationGroupClauseContext) {
}

// ExitWindowAggregationGroupClause is called when production windowAggregationGroupClause is exited.
func (s *BaseDB2zSQLParserListener) ExitWindowAggregationGroupClause(ctx *WindowAggregationGroupClauseContext) {
}

// EnterGroupStart is called when production groupStart is entered.
func (s *BaseDB2zSQLParserListener) EnterGroupStart(ctx *GroupStartContext) {}

// ExitGroupStart is called when production groupStart is exited.
func (s *BaseDB2zSQLParserListener) ExitGroupStart(ctx *GroupStartContext) {}

// EnterGroupBetween is called when production groupBetween is entered.
func (s *BaseDB2zSQLParserListener) EnterGroupBetween(ctx *GroupBetweenContext) {}

// ExitGroupBetween is called when production groupBetween is exited.
func (s *BaseDB2zSQLParserListener) ExitGroupBetween(ctx *GroupBetweenContext) {}

// EnterGroupEnd is called when production groupEnd is entered.
func (s *BaseDB2zSQLParserListener) EnterGroupEnd(ctx *GroupEndContext) {}

// ExitGroupEnd is called when production groupEnd is exited.
func (s *BaseDB2zSQLParserListener) ExitGroupEnd(ctx *GroupEndContext) {}

// EnterGroupBound1 is called when production groupBound1 is entered.
func (s *BaseDB2zSQLParserListener) EnterGroupBound1(ctx *GroupBound1Context) {}

// ExitGroupBound1 is called when production groupBound1 is exited.
func (s *BaseDB2zSQLParserListener) ExitGroupBound1(ctx *GroupBound1Context) {}

// EnterGroupBound2 is called when production groupBound2 is entered.
func (s *BaseDB2zSQLParserListener) EnterGroupBound2(ctx *GroupBound2Context) {}

// ExitGroupBound2 is called when production groupBound2 is exited.
func (s *BaseDB2zSQLParserListener) ExitGroupBound2(ctx *GroupBound2Context) {}

// EnterUnboundedPreceding is called when production unboundedPreceding is entered.
func (s *BaseDB2zSQLParserListener) EnterUnboundedPreceding(ctx *UnboundedPrecedingContext) {}

// ExitUnboundedPreceding is called when production unboundedPreceding is exited.
func (s *BaseDB2zSQLParserListener) ExitUnboundedPreceding(ctx *UnboundedPrecedingContext) {}

// EnterUnboundedFollowing is called when production unboundedFollowing is entered.
func (s *BaseDB2zSQLParserListener) EnterUnboundedFollowing(ctx *UnboundedFollowingContext) {}

// ExitUnboundedFollowing is called when production unboundedFollowing is exited.
func (s *BaseDB2zSQLParserListener) ExitUnboundedFollowing(ctx *UnboundedFollowingContext) {}

// EnterBoundedPreceding is called when production boundedPreceding is entered.
func (s *BaseDB2zSQLParserListener) EnterBoundedPreceding(ctx *BoundedPrecedingContext) {}

// ExitBoundedPreceding is called when production boundedPreceding is exited.
func (s *BaseDB2zSQLParserListener) ExitBoundedPreceding(ctx *BoundedPrecedingContext) {}

// EnterBoundedFollowing is called when production boundedFollowing is entered.
func (s *BaseDB2zSQLParserListener) EnterBoundedFollowing(ctx *BoundedFollowingContext) {}

// ExitBoundedFollowing is called when production boundedFollowing is exited.
func (s *BaseDB2zSQLParserListener) ExitBoundedFollowing(ctx *BoundedFollowingContext) {}

// EnterCurrentRow is called when production currentRow is entered.
func (s *BaseDB2zSQLParserListener) EnterCurrentRow(ctx *CurrentRowContext) {}

// ExitCurrentRow is called when production currentRow is exited.
func (s *BaseDB2zSQLParserListener) ExitCurrentRow(ctx *CurrentRowContext) {}

// EnterScalarFunction is called when production scalarFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterScalarFunction(ctx *ScalarFunctionContext) {}

// ExitScalarFunction is called when production scalarFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitScalarFunction(ctx *ScalarFunctionContext) {}

// EnterTableFunction is called when production tableFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterTableFunction(ctx *TableFunctionContext) {}

// ExitTableFunction is called when production tableFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitTableFunction(ctx *TableFunctionContext) {}

// EnterSpecialRegister is called when production specialRegister is entered.
func (s *BaseDB2zSQLParserListener) EnterSpecialRegister(ctx *SpecialRegisterContext) {}

// ExitSpecialRegister is called when production specialRegister is exited.
func (s *BaseDB2zSQLParserListener) ExitSpecialRegister(ctx *SpecialRegisterContext) {}

// EnterAiAnalogyFunction is called when production aiAnalogyFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterAiAnalogyFunction(ctx *AiAnalogyFunctionContext) {}

// ExitAiAnalogyFunction is called when production aiAnalogyFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitAiAnalogyFunction(ctx *AiAnalogyFunctionContext) {}

// EnterAiFunctionExpression is called when production aiFunctionExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterAiFunctionExpression(ctx *AiFunctionExpressionContext) {}

// ExitAiFunctionExpression is called when production aiFunctionExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitAiFunctionExpression(ctx *AiFunctionExpressionContext) {}

// EnterAiAnalogyFunctionSource is called when production aiAnalogyFunctionSource is entered.
func (s *BaseDB2zSQLParserListener) EnterAiAnalogyFunctionSource(ctx *AiAnalogyFunctionSourceContext) {
}

// ExitAiAnalogyFunctionSource is called when production aiAnalogyFunctionSource is exited.
func (s *BaseDB2zSQLParserListener) ExitAiAnalogyFunctionSource(ctx *AiAnalogyFunctionSourceContext) {
}

// EnterAiAnalogyFunctionTarget is called when production aiAnalogyFunctionTarget is entered.
func (s *BaseDB2zSQLParserListener) EnterAiAnalogyFunctionTarget(ctx *AiAnalogyFunctionTargetContext) {
}

// ExitAiAnalogyFunctionTarget is called when production aiAnalogyFunctionTarget is exited.
func (s *BaseDB2zSQLParserListener) ExitAiAnalogyFunctionTarget(ctx *AiAnalogyFunctionTargetContext) {
}

// EnterAiAnalogyFunctionSource1 is called when production aiAnalogyFunctionSource1 is entered.
func (s *BaseDB2zSQLParserListener) EnterAiAnalogyFunctionSource1(ctx *AiAnalogyFunctionSource1Context) {
}

// ExitAiAnalogyFunctionSource1 is called when production aiAnalogyFunctionSource1 is exited.
func (s *BaseDB2zSQLParserListener) ExitAiAnalogyFunctionSource1(ctx *AiAnalogyFunctionSource1Context) {
}

// EnterAiAnalogyFunctionSource2 is called when production aiAnalogyFunctionSource2 is entered.
func (s *BaseDB2zSQLParserListener) EnterAiAnalogyFunctionSource2(ctx *AiAnalogyFunctionSource2Context) {
}

// ExitAiAnalogyFunctionSource2 is called when production aiAnalogyFunctionSource2 is exited.
func (s *BaseDB2zSQLParserListener) ExitAiAnalogyFunctionSource2(ctx *AiAnalogyFunctionSource2Context) {
}

// EnterAiAnalogyFunctionTarget1 is called when production aiAnalogyFunctionTarget1 is entered.
func (s *BaseDB2zSQLParserListener) EnterAiAnalogyFunctionTarget1(ctx *AiAnalogyFunctionTarget1Context) {
}

// ExitAiAnalogyFunctionTarget1 is called when production aiAnalogyFunctionTarget1 is exited.
func (s *BaseDB2zSQLParserListener) ExitAiAnalogyFunctionTarget1(ctx *AiAnalogyFunctionTarget1Context) {
}

// EnterAiAnalogyFunctionTarget2 is called when production aiAnalogyFunctionTarget2 is entered.
func (s *BaseDB2zSQLParserListener) EnterAiAnalogyFunctionTarget2(ctx *AiAnalogyFunctionTarget2Context) {
}

// ExitAiAnalogyFunctionTarget2 is called when production aiAnalogyFunctionTarget2 is exited.
func (s *BaseDB2zSQLParserListener) ExitAiAnalogyFunctionTarget2(ctx *AiAnalogyFunctionTarget2Context) {
}

// EnterAiSemanticClusterFunction is called when production aiSemanticClusterFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterAiSemanticClusterFunction(ctx *AiSemanticClusterFunctionContext) {
}

// ExitAiSemanticClusterFunction is called when production aiSemanticClusterFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitAiSemanticClusterFunction(ctx *AiSemanticClusterFunctionContext) {
}

// EnterAiSemanticClusterMemberExpression is called when production aiSemanticClusterMemberExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterAiSemanticClusterMemberExpression(ctx *AiSemanticClusterMemberExpressionContext) {
}

// ExitAiSemanticClusterMemberExpression is called when production aiSemanticClusterMemberExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitAiSemanticClusterMemberExpression(ctx *AiSemanticClusterMemberExpressionContext) {
}

// EnterAiSemanticClusterClusteringExpression is called when production aiSemanticClusterClusteringExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterAiSemanticClusterClusteringExpression(ctx *AiSemanticClusterClusteringExpressionContext) {
}

// ExitAiSemanticClusterClusteringExpression is called when production aiSemanticClusterClusteringExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitAiSemanticClusterClusteringExpression(ctx *AiSemanticClusterClusteringExpressionContext) {
}

// EnterAiSimilarityFunction is called when production aiSimilarityFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterAiSimilarityFunction(ctx *AiSimilarityFunctionContext) {}

// ExitAiSimilarityFunction is called when production aiSimilarityFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitAiSimilarityFunction(ctx *AiSimilarityFunctionContext) {}

// EnterAiSimilarityExpression is called when production aiSimilarityExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterAiSimilarityExpression(ctx *AiSimilarityExpressionContext) {}

// ExitAiSimilarityExpression is called when production aiSimilarityExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitAiSimilarityExpression(ctx *AiSimilarityExpressionContext) {}

// EnterAiSimilarityExpression1 is called when production aiSimilarityExpression1 is entered.
func (s *BaseDB2zSQLParserListener) EnterAiSimilarityExpression1(ctx *AiSimilarityExpression1Context) {
}

// ExitAiSimilarityExpression1 is called when production aiSimilarityExpression1 is exited.
func (s *BaseDB2zSQLParserListener) ExitAiSimilarityExpression1(ctx *AiSimilarityExpression1Context) {
}

// EnterAiSimilarityExpression2 is called when production aiSimilarityExpression2 is entered.
func (s *BaseDB2zSQLParserListener) EnterAiSimilarityExpression2(ctx *AiSimilarityExpression2Context) {
}

// ExitAiSimilarityExpression2 is called when production aiSimilarityExpression2 is exited.
func (s *BaseDB2zSQLParserListener) ExitAiSimilarityExpression2(ctx *AiSimilarityExpression2Context) {
}

// EnterXmlelementFunction is called when production xmlelementFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlelementFunction(ctx *XmlelementFunctionContext) {}

// ExitXmlelementFunction is called when production xmlelementFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlelementFunction(ctx *XmlelementFunctionContext) {}

// EnterXmlforestFunction is called when production xmlforestFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlforestFunction(ctx *XmlforestFunctionContext) {}

// ExitXmlforestFunction is called when production xmlforestFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlforestFunction(ctx *XmlforestFunctionContext) {}

// EnterXmlmodifyFunction is called when production xmlmodifyFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlmodifyFunction(ctx *XmlmodifyFunctionContext) {}

// ExitXmlmodifyFunction is called when production xmlmodifyFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlmodifyFunction(ctx *XmlmodifyFunctionContext) {}

// EnterXmlpiFunction is called when production xmlpiFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlpiFunction(ctx *XmlpiFunctionContext) {}

// ExitXmlpiFunction is called when production xmlpiFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlpiFunction(ctx *XmlpiFunctionContext) {}

// EnterXmlqueryFunction is called when production xmlqueryFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlqueryFunction(ctx *XmlqueryFunctionContext) {}

// ExitXmlqueryFunction is called when production xmlqueryFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlqueryFunction(ctx *XmlqueryFunctionContext) {}

// EnterXmlattributesFunction is called when production xmlattributesFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlattributesFunction(ctx *XmlattributesFunctionContext) {}

// ExitXmlattributesFunction is called when production xmlattributesFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlattributesFunction(ctx *XmlattributesFunctionContext) {}

// EnterXmlserializeFunction is called when production xmlserializeFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlserializeFunction(ctx *XmlserializeFunctionContext) {}

// ExitXmlserializeFunction is called when production xmlserializeFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlserializeFunction(ctx *XmlserializeFunctionContext) {}

// EnterXmlnamespaceFunction is called when production xmlnamespaceFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlnamespaceFunction(ctx *XmlnamespaceFunctionContext) {}

// ExitXmlnamespaceFunction is called when production xmlnamespaceFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlnamespaceFunction(ctx *XmlnamespaceFunctionContext) {}

// EnterXmlnamespaceOption is called when production xmlnamespaceOption is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlnamespaceOption(ctx *XmlnamespaceOptionContext) {}

// ExitXmlnamespaceOption is called when production xmlnamespaceOption is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlnamespaceOption(ctx *XmlnamespaceOptionContext) {}

// EnterXmlserializeFunctionOptions is called when production xmlserializeFunctionOptions is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlserializeFunctionOptions(ctx *XmlserializeFunctionOptionsContext) {
}

// ExitXmlserializeFunctionOptions is called when production xmlserializeFunctionOptions is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlserializeFunctionOptions(ctx *XmlserializeFunctionOptionsContext) {
}

// EnterXmlFunctionOptionClause is called when production xmlFunctionOptionClause is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlFunctionOptionClause(ctx *XmlFunctionOptionClauseContext) {
}

// ExitXmlFunctionOptionClause is called when production xmlFunctionOptionClause is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlFunctionOptionClause(ctx *XmlFunctionOptionClauseContext) {
}

// EnterXmlFunctionOption is called when production xmlFunctionOption is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlFunctionOption(ctx *XmlFunctionOptionContext) {}

// ExitXmlFunctionOption is called when production xmlFunctionOption is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlFunctionOption(ctx *XmlFunctionOptionContext) {}

// EnterElementContentExpression is called when production elementContentExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterElementContentExpression(ctx *ElementContentExpressionContext) {
}

// ExitElementContentExpression is called when production elementContentExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitElementContentExpression(ctx *ElementContentExpressionContext) {
}

// EnterXqueryExpressionConstant is called when production xqueryExpressionConstant is entered.
func (s *BaseDB2zSQLParserListener) EnterXqueryExpressionConstant(ctx *XqueryExpressionConstantContext) {
}

// ExitXqueryExpressionConstant is called when production xqueryExpressionConstant is exited.
func (s *BaseDB2zSQLParserListener) ExitXqueryExpressionConstant(ctx *XqueryExpressionConstantContext) {
}

// EnterXqueryArgument is called when production xqueryArgument is entered.
func (s *BaseDB2zSQLParserListener) EnterXqueryArgument(ctx *XqueryArgumentContext) {}

// ExitXqueryArgument is called when production xqueryArgument is exited.
func (s *BaseDB2zSQLParserListener) ExitXqueryArgument(ctx *XqueryArgumentContext) {}

// EnterXmltableFunctionSpecification is called when production xmltableFunctionSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterXmltableFunctionSpecification(ctx *XmltableFunctionSpecificationContext) {
}

// ExitXmltableFunctionSpecification is called when production xmltableFunctionSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitXmltableFunctionSpecification(ctx *XmltableFunctionSpecificationContext) {
}

// EnterRowXqueryExpressionConstant is called when production rowXqueryExpressionConstant is entered.
func (s *BaseDB2zSQLParserListener) EnterRowXqueryExpressionConstant(ctx *RowXqueryExpressionConstantContext) {
}

// ExitRowXqueryExpressionConstant is called when production rowXqueryExpressionConstant is exited.
func (s *BaseDB2zSQLParserListener) ExitRowXqueryExpressionConstant(ctx *RowXqueryExpressionConstantContext) {
}

// EnterRowXqueryArgument is called when production rowXqueryArgument is entered.
func (s *BaseDB2zSQLParserListener) EnterRowXqueryArgument(ctx *RowXqueryArgumentContext) {}

// ExitRowXqueryArgument is called when production rowXqueryArgument is exited.
func (s *BaseDB2zSQLParserListener) ExitRowXqueryArgument(ctx *RowXqueryArgumentContext) {}

// EnterXqueryContextItemExpression is called when production xqueryContextItemExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterXqueryContextItemExpression(ctx *XqueryContextItemExpressionContext) {
}

// ExitXqueryContextItemExpression is called when production xqueryContextItemExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitXqueryContextItemExpression(ctx *XqueryContextItemExpressionContext) {
}

// EnterXqueryVariableExpression is called when production xqueryVariableExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterXqueryVariableExpression(ctx *XqueryVariableExpressionContext) {
}

// ExitXqueryVariableExpression is called when production xqueryVariableExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitXqueryVariableExpression(ctx *XqueryVariableExpressionContext) {
}

// EnterXmlTableRegularColumnDefinition is called when production xmlTableRegularColumnDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlTableRegularColumnDefinition(ctx *XmlTableRegularColumnDefinitionContext) {
}

// ExitXmlTableRegularColumnDefinition is called when production xmlTableRegularColumnDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlTableRegularColumnDefinition(ctx *XmlTableRegularColumnDefinitionContext) {
}

// EnterDefaultClause is called when production defaultClause is entered.
func (s *BaseDB2zSQLParserListener) EnterDefaultClause(ctx *DefaultClauseContext) {}

// ExitDefaultClause is called when production defaultClause is exited.
func (s *BaseDB2zSQLParserListener) ExitDefaultClause(ctx *DefaultClauseContext) {}

// EnterDefaultClause1 is called when production defaultClause1 is entered.
func (s *BaseDB2zSQLParserListener) EnterDefaultClause1(ctx *DefaultClause1Context) {}

// ExitDefaultClause1 is called when production defaultClause1 is exited.
func (s *BaseDB2zSQLParserListener) ExitDefaultClause1(ctx *DefaultClause1Context) {}

// EnterDefaultClauseAllowables is called when production defaultClauseAllowables is entered.
func (s *BaseDB2zSQLParserListener) EnterDefaultClauseAllowables(ctx *DefaultClauseAllowablesContext) {
}

// ExitDefaultClauseAllowables is called when production defaultClauseAllowables is exited.
func (s *BaseDB2zSQLParserListener) ExitDefaultClauseAllowables(ctx *DefaultClauseAllowablesContext) {
}

// EnterDistinctTypeCastFunctionName is called when production distinctTypeCastFunctionName is entered.
func (s *BaseDB2zSQLParserListener) EnterDistinctTypeCastFunctionName(ctx *DistinctTypeCastFunctionNameContext) {
}

// ExitDistinctTypeCastFunctionName is called when production distinctTypeCastFunctionName is exited.
func (s *BaseDB2zSQLParserListener) ExitDistinctTypeCastFunctionName(ctx *DistinctTypeCastFunctionNameContext) {
}

// EnterColumnXqueryExpressionConstant is called when production columnXqueryExpressionConstant is entered.
func (s *BaseDB2zSQLParserListener) EnterColumnXqueryExpressionConstant(ctx *ColumnXqueryExpressionConstantContext) {
}

// ExitColumnXqueryExpressionConstant is called when production columnXqueryExpressionConstant is exited.
func (s *BaseDB2zSQLParserListener) ExitColumnXqueryExpressionConstant(ctx *ColumnXqueryExpressionConstantContext) {
}

// EnterXmlTableOrdinalityColumnDefinition is called when production xmlTableOrdinalityColumnDefinition is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlTableOrdinalityColumnDefinition(ctx *XmlTableOrdinalityColumnDefinitionContext) {
}

// ExitXmlTableOrdinalityColumnDefinition is called when production xmlTableOrdinalityColumnDefinition is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlTableOrdinalityColumnDefinition(ctx *XmlTableOrdinalityColumnDefinitionContext) {
}

// EnterXmlnamespacesDeclaration is called when production xmlnamespacesDeclaration is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlnamespacesDeclaration(ctx *XmlnamespacesDeclarationContext) {
}

// ExitXmlnamespacesDeclaration is called when production xmlnamespacesDeclaration is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlnamespacesDeclaration(ctx *XmlnamespacesDeclarationContext) {
}

// EnterXmlnamespacesFunctionSpecification is called when production xmlnamespacesFunctionSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlnamespacesFunctionSpecification(ctx *XmlnamespacesFunctionSpecificationContext) {
}

// ExitXmlnamespacesFunctionSpecification is called when production xmlnamespacesFunctionSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlnamespacesFunctionSpecification(ctx *XmlnamespacesFunctionSpecificationContext) {
}

// EnterXmlnamespacesFunctionArguments is called when production xmlnamespacesFunctionArguments is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlnamespacesFunctionArguments(ctx *XmlnamespacesFunctionArgumentsContext) {
}

// ExitXmlnamespacesFunctionArguments is called when production xmlnamespacesFunctionArguments is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlnamespacesFunctionArguments(ctx *XmlnamespacesFunctionArgumentsContext) {
}

// EnterNamespaceUri is called when production namespaceUri is entered.
func (s *BaseDB2zSQLParserListener) EnterNamespaceUri(ctx *NamespaceUriContext) {}

// ExitNamespaceUri is called when production namespaceUri is exited.
func (s *BaseDB2zSQLParserListener) ExitNamespaceUri(ctx *NamespaceUriContext) {}

// EnterNamespacePrefix is called when production namespacePrefix is entered.
func (s *BaseDB2zSQLParserListener) EnterNamespacePrefix(ctx *NamespacePrefixContext) {}

// ExitNamespacePrefix is called when production namespacePrefix is exited.
func (s *BaseDB2zSQLParserListener) ExitNamespacePrefix(ctx *NamespacePrefixContext) {}

// EnterTimeZoneSpecificExpression is called when production timeZoneSpecificExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterTimeZoneSpecificExpression(ctx *TimeZoneSpecificExpressionContext) {
}

// ExitTimeZoneSpecificExpression is called when production timeZoneSpecificExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitTimeZoneSpecificExpression(ctx *TimeZoneSpecificExpressionContext) {
}

// EnterTimeZoneExpressionSubset is called when production timeZoneExpressionSubset is entered.
func (s *BaseDB2zSQLParserListener) EnterTimeZoneExpressionSubset(ctx *TimeZoneExpressionSubsetContext) {
}

// ExitTimeZoneExpressionSubset is called when production timeZoneExpressionSubset is exited.
func (s *BaseDB2zSQLParserListener) ExitTimeZoneExpressionSubset(ctx *TimeZoneExpressionSubsetContext) {
}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterResultExpression is called when production resultExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterResultExpression(ctx *ResultExpressionContext) {}

// ExitResultExpression is called when production resultExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitResultExpression(ctx *ResultExpressionContext) {}

// EnterSearchedWhenClause is called when production searchedWhenClause is entered.
func (s *BaseDB2zSQLParserListener) EnterSearchedWhenClause(ctx *SearchedWhenClauseContext) {}

// ExitSearchedWhenClause is called when production searchedWhenClause is exited.
func (s *BaseDB2zSQLParserListener) ExitSearchedWhenClause(ctx *SearchedWhenClauseContext) {}

// EnterSimpleWhenClause is called when production simpleWhenClause is entered.
func (s *BaseDB2zSQLParserListener) EnterSimpleWhenClause(ctx *SimpleWhenClauseContext) {}

// ExitSimpleWhenClause is called when production simpleWhenClause is exited.
func (s *BaseDB2zSQLParserListener) ExitSimpleWhenClause(ctx *SimpleWhenClauseContext) {}

// EnterSearchCondition is called when production searchCondition is entered.
func (s *BaseDB2zSQLParserListener) EnterSearchCondition(ctx *SearchConditionContext) {}

// ExitSearchCondition is called when production searchCondition is exited.
func (s *BaseDB2zSQLParserListener) ExitSearchCondition(ctx *SearchConditionContext) {}

// EnterCheckCondition is called when production checkCondition is entered.
func (s *BaseDB2zSQLParserListener) EnterCheckCondition(ctx *CheckConditionContext) {}

// ExitCheckCondition is called when production checkCondition is exited.
func (s *BaseDB2zSQLParserListener) ExitCheckCondition(ctx *CheckConditionContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseDB2zSQLParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseDB2zSQLParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterBasicPredicate is called when production basicPredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterBasicPredicate(ctx *BasicPredicateContext) {}

// ExitBasicPredicate is called when production basicPredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitBasicPredicate(ctx *BasicPredicateContext) {}

// EnterRowValueExpression is called when production rowValueExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterRowValueExpression(ctx *RowValueExpressionContext) {}

// ExitRowValueExpression is called when production rowValueExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitRowValueExpression(ctx *RowValueExpressionContext) {}

// EnterQuantifiedPredicate is called when production quantifiedPredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterQuantifiedPredicate(ctx *QuantifiedPredicateContext) {}

// ExitQuantifiedPredicate is called when production quantifiedPredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitQuantifiedPredicate(ctx *QuantifiedPredicateContext) {}

// EnterArrayExistsPredicate is called when production arrayExistsPredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayExistsPredicate(ctx *ArrayExistsPredicateContext) {}

// ExitArrayExistsPredicate is called when production arrayExistsPredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayExistsPredicate(ctx *ArrayExistsPredicateContext) {}

// EnterBetweenPredicate is called when production betweenPredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterBetweenPredicate(ctx *BetweenPredicateContext) {}

// ExitBetweenPredicate is called when production betweenPredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {}

// EnterDistinctPredicate is called when production distinctPredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterDistinctPredicate(ctx *DistinctPredicateContext) {}

// ExitDistinctPredicate is called when production distinctPredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitDistinctPredicate(ctx *DistinctPredicateContext) {}

// EnterExistsPredicate is called when production existsPredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterExistsPredicate(ctx *ExistsPredicateContext) {}

// ExitExistsPredicate is called when production existsPredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitExistsPredicate(ctx *ExistsPredicateContext) {}

// EnterInPredicate is called when production inPredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterInPredicate(ctx *InPredicateContext) {}

// ExitInPredicate is called when production inPredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitInPredicate(ctx *InPredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterNullPredicate is called when production nullPredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterNullPredicate(ctx *NullPredicateContext) {}

// ExitNullPredicate is called when production nullPredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitNullPredicate(ctx *NullPredicateContext) {}

// EnterXmlExistsPredicate is called when production xmlExistsPredicate is entered.
func (s *BaseDB2zSQLParserListener) EnterXmlExistsPredicate(ctx *XmlExistsPredicateContext) {}

// ExitXmlExistsPredicate is called when production xmlExistsPredicate is exited.
func (s *BaseDB2zSQLParserListener) ExitXmlExistsPredicate(ctx *XmlExistsPredicateContext) {}

// EnterArrayExpression is called when production arrayExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayExpression(ctx *ArrayExpressionContext) {}

// ExitArrayExpression is called when production arrayExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayExpression(ctx *ArrayExpressionContext) {}

// EnterCastSpecification is called when production castSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterCastSpecification(ctx *CastSpecificationContext) {}

// ExitCastSpecification is called when production castSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitCastSpecification(ctx *CastSpecificationContext) {}

// EnterParameterMarker is called when production parameterMarker is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterMarker(ctx *ParameterMarkerContext) {}

// ExitParameterMarker is called when production parameterMarker is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterMarker(ctx *ParameterMarkerContext) {}

// EnterCastDataType is called when production castDataType is entered.
func (s *BaseDB2zSQLParserListener) EnterCastDataType(ctx *CastDataTypeContext) {}

// ExitCastDataType is called when production castDataType is exited.
func (s *BaseDB2zSQLParserListener) ExitCastDataType(ctx *CastDataTypeContext) {}

// EnterCastBuiltInType is called when production castBuiltInType is entered.
func (s *BaseDB2zSQLParserListener) EnterCastBuiltInType(ctx *CastBuiltInTypeContext) {}

// ExitCastBuiltInType is called when production castBuiltInType is exited.
func (s *BaseDB2zSQLParserListener) ExitCastBuiltInType(ctx *CastBuiltInTypeContext) {}

// EnterIntegerInParens is called when production integerInParens is entered.
func (s *BaseDB2zSQLParserListener) EnterIntegerInParens(ctx *IntegerInParensContext) {}

// ExitIntegerInParens is called when production integerInParens is exited.
func (s *BaseDB2zSQLParserListener) ExitIntegerInParens(ctx *IntegerInParensContext) {}

// EnterLength is called when production length is entered.
func (s *BaseDB2zSQLParserListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseDB2zSQLParserListener) ExitLength(ctx *LengthContext) {}

// EnterCcsidQualifier is called when production ccsidQualifier is entered.
func (s *BaseDB2zSQLParserListener) EnterCcsidQualifier(ctx *CcsidQualifierContext) {}

// ExitCcsidQualifier is called when production ccsidQualifier is exited.
func (s *BaseDB2zSQLParserListener) ExitCcsidQualifier(ctx *CcsidQualifierContext) {}

// EnterForDataQualifier is called when production forDataQualifier is entered.
func (s *BaseDB2zSQLParserListener) EnterForDataQualifier(ctx *ForDataQualifierContext) {}

// ExitForDataQualifier is called when production forDataQualifier is exited.
func (s *BaseDB2zSQLParserListener) ExitForDataQualifier(ctx *ForDataQualifierContext) {}

// EnterDistinctTypeName is called when production distinctTypeName is entered.
func (s *BaseDB2zSQLParserListener) EnterDistinctTypeName(ctx *DistinctTypeNameContext) {}

// ExitDistinctTypeName is called when production distinctTypeName is exited.
func (s *BaseDB2zSQLParserListener) ExitDistinctTypeName(ctx *DistinctTypeNameContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseDB2zSQLParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseDB2zSQLParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterCcsidValue is called when production ccsidValue is entered.
func (s *BaseDB2zSQLParserListener) EnterCcsidValue(ctx *CcsidValueContext) {}

// ExitCcsidValue is called when production ccsidValue is exited.
func (s *BaseDB2zSQLParserListener) ExitCcsidValue(ctx *CcsidValueContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseDB2zSQLParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseDB2zSQLParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterSourceColumnName is called when production sourceColumnName is entered.
func (s *BaseDB2zSQLParserListener) EnterSourceColumnName(ctx *SourceColumnNameContext) {}

// ExitSourceColumnName is called when production sourceColumnName is exited.
func (s *BaseDB2zSQLParserListener) ExitSourceColumnName(ctx *SourceColumnNameContext) {}

// EnterTargetColumnName is called when production targetColumnName is entered.
func (s *BaseDB2zSQLParserListener) EnterTargetColumnName(ctx *TargetColumnNameContext) {}

// ExitTargetColumnName is called when production targetColumnName is exited.
func (s *BaseDB2zSQLParserListener) ExitTargetColumnName(ctx *TargetColumnNameContext) {}

// EnterBuildColumnName is called when production buildColumnName is entered.
func (s *BaseDB2zSQLParserListener) EnterBuildColumnName(ctx *BuildColumnNameContext) {}

// ExitBuildColumnName is called when production buildColumnName is exited.
func (s *BaseDB2zSQLParserListener) ExitBuildColumnName(ctx *BuildColumnNameContext) {}

// EnterBeginColumnName is called when production beginColumnName is entered.
func (s *BaseDB2zSQLParserListener) EnterBeginColumnName(ctx *BeginColumnNameContext) {}

// ExitBeginColumnName is called when production beginColumnName is exited.
func (s *BaseDB2zSQLParserListener) ExitBeginColumnName(ctx *BeginColumnNameContext) {}

// EnterEndColumnName is called when production endColumnName is entered.
func (s *BaseDB2zSQLParserListener) EnterEndColumnName(ctx *EndColumnNameContext) {}

// ExitEndColumnName is called when production endColumnName is exited.
func (s *BaseDB2zSQLParserListener) ExitEndColumnName(ctx *EndColumnNameContext) {}

// EnterCorrelationName is called when production correlationName is entered.
func (s *BaseDB2zSQLParserListener) EnterCorrelationName(ctx *CorrelationNameContext) {}

// ExitCorrelationName is called when production correlationName is exited.
func (s *BaseDB2zSQLParserListener) ExitCorrelationName(ctx *CorrelationNameContext) {}

// EnterLocationName is called when production locationName is entered.
func (s *BaseDB2zSQLParserListener) EnterLocationName(ctx *LocationNameContext) {}

// ExitLocationName is called when production locationName is exited.
func (s *BaseDB2zSQLParserListener) ExitLocationName(ctx *LocationNameContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseDB2zSQLParserListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseDB2zSQLParserListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseDB2zSQLParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseDB2zSQLParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterAlterTableName is called when production alterTableName is entered.
func (s *BaseDB2zSQLParserListener) EnterAlterTableName(ctx *AlterTableNameContext) {}

// ExitAlterTableName is called when production alterTableName is exited.
func (s *BaseDB2zSQLParserListener) ExitAlterTableName(ctx *AlterTableNameContext) {}

// EnterAuxTableName is called when production auxTableName is entered.
func (s *BaseDB2zSQLParserListener) EnterAuxTableName(ctx *AuxTableNameContext) {}

// ExitAuxTableName is called when production auxTableName is exited.
func (s *BaseDB2zSQLParserListener) ExitAuxTableName(ctx *AuxTableNameContext) {}

// EnterHistoryTableName is called when production historyTableName is entered.
func (s *BaseDB2zSQLParserListener) EnterHistoryTableName(ctx *HistoryTableNameContext) {}

// ExitHistoryTableName is called when production historyTableName is exited.
func (s *BaseDB2zSQLParserListener) ExitHistoryTableName(ctx *HistoryTableNameContext) {}

// EnterCloneTableName is called when production cloneTableName is entered.
func (s *BaseDB2zSQLParserListener) EnterCloneTableName(ctx *CloneTableNameContext) {}

// ExitCloneTableName is called when production cloneTableName is exited.
func (s *BaseDB2zSQLParserListener) ExitCloneTableName(ctx *CloneTableNameContext) {}

// EnterArchiveTableName is called when production archiveTableName is entered.
func (s *BaseDB2zSQLParserListener) EnterArchiveTableName(ctx *ArchiveTableNameContext) {}

// ExitArchiveTableName is called when production archiveTableName is exited.
func (s *BaseDB2zSQLParserListener) ExitArchiveTableName(ctx *ArchiveTableNameContext) {}

// EnterViewName is called when production viewName is entered.
func (s *BaseDB2zSQLParserListener) EnterViewName(ctx *ViewNameContext) {}

// ExitViewName is called when production viewName is exited.
func (s *BaseDB2zSQLParserListener) ExitViewName(ctx *ViewNameContext) {}

// EnterProgramName is called when production programName is entered.
func (s *BaseDB2zSQLParserListener) EnterProgramName(ctx *ProgramNameContext) {}

// ExitProgramName is called when production programName is exited.
func (s *BaseDB2zSQLParserListener) ExitProgramName(ctx *ProgramNameContext) {}

// EnterPackageName is called when production packageName is entered.
func (s *BaseDB2zSQLParserListener) EnterPackageName(ctx *PackageNameContext) {}

// ExitPackageName is called when production packageName is exited.
func (s *BaseDB2zSQLParserListener) ExitPackageName(ctx *PackageNameContext) {}

// EnterPlanName is called when production planName is entered.
func (s *BaseDB2zSQLParserListener) EnterPlanName(ctx *PlanNameContext) {}

// ExitPlanName is called when production planName is exited.
func (s *BaseDB2zSQLParserListener) ExitPlanName(ctx *PlanNameContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseDB2zSQLParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseDB2zSQLParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BaseDB2zSQLParserListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BaseDB2zSQLParserListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterArrayTypeName is called when production arrayTypeName is entered.
func (s *BaseDB2zSQLParserListener) EnterArrayTypeName(ctx *ArrayTypeNameContext) {}

// ExitArrayTypeName is called when production arrayTypeName is exited.
func (s *BaseDB2zSQLParserListener) ExitArrayTypeName(ctx *ArrayTypeNameContext) {}

// EnterJarName is called when production jarName is entered.
func (s *BaseDB2zSQLParserListener) EnterJarName(ctx *JarNameContext) {}

// ExitJarName is called when production jarName is exited.
func (s *BaseDB2zSQLParserListener) ExitJarName(ctx *JarNameContext) {}

// EnterSavepointName is called when production savepointName is entered.
func (s *BaseDB2zSQLParserListener) EnterSavepointName(ctx *SavepointNameContext) {}

// ExitSavepointName is called when production savepointName is exited.
func (s *BaseDB2zSQLParserListener) ExitSavepointName(ctx *SavepointNameContext) {}

// EnterAliasName is called when production aliasName is entered.
func (s *BaseDB2zSQLParserListener) EnterAliasName(ctx *AliasNameContext) {}

// ExitAliasName is called when production aliasName is exited.
func (s *BaseDB2zSQLParserListener) ExitAliasName(ctx *AliasNameContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BaseDB2zSQLParserListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BaseDB2zSQLParserListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterRoutineVersionID is called when production routineVersionID is entered.
func (s *BaseDB2zSQLParserListener) EnterRoutineVersionID(ctx *RoutineVersionIDContext) {}

// ExitRoutineVersionID is called when production routineVersionID is exited.
func (s *BaseDB2zSQLParserListener) ExitRoutineVersionID(ctx *RoutineVersionIDContext) {}

// EnterVersionID is called when production versionID is entered.
func (s *BaseDB2zSQLParserListener) EnterVersionID(ctx *VersionIDContext) {}

// ExitVersionID is called when production versionID is exited.
func (s *BaseDB2zSQLParserListener) ExitVersionID(ctx *VersionIDContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseDB2zSQLParserListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseDB2zSQLParserListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterMaskName is called when production maskName is entered.
func (s *BaseDB2zSQLParserListener) EnterMaskName(ctx *MaskNameContext) {}

// ExitMaskName is called when production maskName is exited.
func (s *BaseDB2zSQLParserListener) ExitMaskName(ctx *MaskNameContext) {}

// EnterPermissionName is called when production permissionName is entered.
func (s *BaseDB2zSQLParserListener) EnterPermissionName(ctx *PermissionNameContext) {}

// ExitPermissionName is called when production permissionName is exited.
func (s *BaseDB2zSQLParserListener) ExitPermissionName(ctx *PermissionNameContext) {}

// EnterProcedureName is called when production procedureName is entered.
func (s *BaseDB2zSQLParserListener) EnterProcedureName(ctx *ProcedureNameContext) {}

// ExitProcedureName is called when production procedureName is exited.
func (s *BaseDB2zSQLParserListener) ExitProcedureName(ctx *ProcedureNameContext) {}

// EnterSequenceName is called when production sequenceName is entered.
func (s *BaseDB2zSQLParserListener) EnterSequenceName(ctx *SequenceNameContext) {}

// ExitSequenceName is called when production sequenceName is exited.
func (s *BaseDB2zSQLParserListener) ExitSequenceName(ctx *SequenceNameContext) {}

// EnterMemberName is called when production memberName is entered.
func (s *BaseDB2zSQLParserListener) EnterMemberName(ctx *MemberNameContext) {}

// ExitMemberName is called when production memberName is exited.
func (s *BaseDB2zSQLParserListener) ExitMemberName(ctx *MemberNameContext) {}

// EnterDatabaseName is called when production databaseName is entered.
func (s *BaseDB2zSQLParserListener) EnterDatabaseName(ctx *DatabaseNameContext) {}

// ExitDatabaseName is called when production databaseName is exited.
func (s *BaseDB2zSQLParserListener) ExitDatabaseName(ctx *DatabaseNameContext) {}

// EnterTablespaceName is called when production tablespaceName is entered.
func (s *BaseDB2zSQLParserListener) EnterTablespaceName(ctx *TablespaceNameContext) {}

// ExitTablespaceName is called when production tablespaceName is exited.
func (s *BaseDB2zSQLParserListener) ExitTablespaceName(ctx *TablespaceNameContext) {}

// EnterAcceleratorName is called when production acceleratorName is entered.
func (s *BaseDB2zSQLParserListener) EnterAcceleratorName(ctx *AcceleratorNameContext) {}

// ExitAcceleratorName is called when production acceleratorName is exited.
func (s *BaseDB2zSQLParserListener) ExitAcceleratorName(ctx *AcceleratorNameContext) {}

// EnterCatalogName is called when production catalogName is entered.
func (s *BaseDB2zSQLParserListener) EnterCatalogName(ctx *CatalogNameContext) {}

// ExitCatalogName is called when production catalogName is exited.
func (s *BaseDB2zSQLParserListener) ExitCatalogName(ctx *CatalogNameContext) {}

// EnterTriggerName is called when production triggerName is entered.
func (s *BaseDB2zSQLParserListener) EnterTriggerName(ctx *TriggerNameContext) {}

// ExitTriggerName is called when production triggerName is exited.
func (s *BaseDB2zSQLParserListener) ExitTriggerName(ctx *TriggerNameContext) {}

// EnterContextName is called when production contextName is entered.
func (s *BaseDB2zSQLParserListener) EnterContextName(ctx *ContextNameContext) {}

// ExitContextName is called when production contextName is exited.
func (s *BaseDB2zSQLParserListener) ExitContextName(ctx *ContextNameContext) {}

// EnterAuthorizationName is called when production authorizationName is entered.
func (s *BaseDB2zSQLParserListener) EnterAuthorizationName(ctx *AuthorizationNameContext) {}

// ExitAuthorizationName is called when production authorizationName is exited.
func (s *BaseDB2zSQLParserListener) ExitAuthorizationName(ctx *AuthorizationNameContext) {}

// EnterProfileName is called when production profileName is entered.
func (s *BaseDB2zSQLParserListener) EnterProfileName(ctx *ProfileNameContext) {}

// ExitProfileName is called when production profileName is exited.
func (s *BaseDB2zSQLParserListener) ExitProfileName(ctx *ProfileNameContext) {}

// EnterRoleName is called when production roleName is entered.
func (s *BaseDB2zSQLParserListener) EnterRoleName(ctx *RoleNameContext) {}

// ExitRoleName is called when production roleName is exited.
func (s *BaseDB2zSQLParserListener) ExitRoleName(ctx *RoleNameContext) {}

// EnterSeclabelName is called when production seclabelName is entered.
func (s *BaseDB2zSQLParserListener) EnterSeclabelName(ctx *SeclabelNameContext) {}

// ExitSeclabelName is called when production seclabelName is exited.
func (s *BaseDB2zSQLParserListener) ExitSeclabelName(ctx *SeclabelNameContext) {}

// EnterParameterName is called when production parameterName is entered.
func (s *BaseDB2zSQLParserListener) EnterParameterName(ctx *ParameterNameContext) {}

// ExitParameterName is called when production parameterName is exited.
func (s *BaseDB2zSQLParserListener) ExitParameterName(ctx *ParameterNameContext) {}

// EnterAddressValue is called when production addressValue is entered.
func (s *BaseDB2zSQLParserListener) EnterAddressValue(ctx *AddressValueContext) {}

// ExitAddressValue is called when production addressValue is exited.
func (s *BaseDB2zSQLParserListener) ExitAddressValue(ctx *AddressValueContext) {}

// EnterJobnameValue is called when production jobnameValue is entered.
func (s *BaseDB2zSQLParserListener) EnterJobnameValue(ctx *JobnameValueContext) {}

// ExitJobnameValue is called when production jobnameValue is exited.
func (s *BaseDB2zSQLParserListener) ExitJobnameValue(ctx *JobnameValueContext) {}

// EnterServauthValue is called when production servauthValue is entered.
func (s *BaseDB2zSQLParserListener) EnterServauthValue(ctx *ServauthValueContext) {}

// ExitServauthValue is called when production servauthValue is exited.
func (s *BaseDB2zSQLParserListener) ExitServauthValue(ctx *ServauthValueContext) {}

// EnterEncryptionValue is called when production encryptionValue is entered.
func (s *BaseDB2zSQLParserListener) EnterEncryptionValue(ctx *EncryptionValueContext) {}

// ExitEncryptionValue is called when production encryptionValue is exited.
func (s *BaseDB2zSQLParserListener) ExitEncryptionValue(ctx *EncryptionValueContext) {}

// EnterBpName is called when production bpName is entered.
func (s *BaseDB2zSQLParserListener) EnterBpName(ctx *BpNameContext) {}

// ExitBpName is called when production bpName is exited.
func (s *BaseDB2zSQLParserListener) ExitBpName(ctx *BpNameContext) {}

// EnterStogroupName is called when production stogroupName is entered.
func (s *BaseDB2zSQLParserListener) EnterStogroupName(ctx *StogroupNameContext) {}

// ExitStogroupName is called when production stogroupName is exited.
func (s *BaseDB2zSQLParserListener) ExitStogroupName(ctx *StogroupNameContext) {}

// EnterDcName is called when production dcName is entered.
func (s *BaseDB2zSQLParserListener) EnterDcName(ctx *DcNameContext) {}

// ExitDcName is called when production dcName is exited.
func (s *BaseDB2zSQLParserListener) ExitDcName(ctx *DcNameContext) {}

// EnterMcName is called when production mcName is entered.
func (s *BaseDB2zSQLParserListener) EnterMcName(ctx *McNameContext) {}

// ExitMcName is called when production mcName is exited.
func (s *BaseDB2zSQLParserListener) ExitMcName(ctx *McNameContext) {}

// EnterScName is called when production scName is entered.
func (s *BaseDB2zSQLParserListener) EnterScName(ctx *ScNameContext) {}

// ExitScName is called when production scName is exited.
func (s *BaseDB2zSQLParserListener) ExitScName(ctx *ScNameContext) {}

// EnterVolumeID is called when production volumeID is entered.
func (s *BaseDB2zSQLParserListener) EnterVolumeID(ctx *VolumeIDContext) {}

// ExitVolumeID is called when production volumeID is exited.
func (s *BaseDB2zSQLParserListener) ExitVolumeID(ctx *VolumeIDContext) {}

// EnterKeyLabelName is called when production keyLabelName is entered.
func (s *BaseDB2zSQLParserListener) EnterKeyLabelName(ctx *KeyLabelNameContext) {}

// ExitKeyLabelName is called when production keyLabelName is exited.
func (s *BaseDB2zSQLParserListener) ExitKeyLabelName(ctx *KeyLabelNameContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseDB2zSQLParserListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseDB2zSQLParserListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterSpecificName is called when production specificName is entered.
func (s *BaseDB2zSQLParserListener) EnterSpecificName(ctx *SpecificNameContext) {}

// ExitSpecificName is called when production specificName is exited.
func (s *BaseDB2zSQLParserListener) ExitSpecificName(ctx *SpecificNameContext) {}

// EnterHostLabel is called when production hostLabel is entered.
func (s *BaseDB2zSQLParserListener) EnterHostLabel(ctx *HostLabelContext) {}

// ExitHostLabel is called when production hostLabel is exited.
func (s *BaseDB2zSQLParserListener) ExitHostLabel(ctx *HostLabelContext) {}

// EnterHostVariable is called when production hostVariable is entered.
func (s *BaseDB2zSQLParserListener) EnterHostVariable(ctx *HostVariableContext) {}

// ExitHostVariable is called when production hostVariable is exited.
func (s *BaseDB2zSQLParserListener) ExitHostVariable(ctx *HostVariableContext) {}

// EnterHostIdentifier is called when production hostIdentifier is entered.
func (s *BaseDB2zSQLParserListener) EnterHostIdentifier(ctx *HostIdentifierContext) {}

// ExitHostIdentifier is called when production hostIdentifier is exited.
func (s *BaseDB2zSQLParserListener) ExitHostIdentifier(ctx *HostIdentifierContext) {}

// EnterHostStructure is called when production hostStructure is entered.
func (s *BaseDB2zSQLParserListener) EnterHostStructure(ctx *HostStructureContext) {}

// ExitHostStructure is called when production hostStructure is exited.
func (s *BaseDB2zSQLParserListener) ExitHostStructure(ctx *HostStructureContext) {}

// EnterNullIndicator is called when production nullIndicator is entered.
func (s *BaseDB2zSQLParserListener) EnterNullIndicator(ctx *NullIndicatorContext) {}

// ExitNullIndicator is called when production nullIndicator is exited.
func (s *BaseDB2zSQLParserListener) ExitNullIndicator(ctx *NullIndicatorContext) {}

// EnterNullIndicatorStructure is called when production nullIndicatorStructure is entered.
func (s *BaseDB2zSQLParserListener) EnterNullIndicatorStructure(ctx *NullIndicatorStructureContext) {}

// ExitNullIndicatorStructure is called when production nullIndicatorStructure is exited.
func (s *BaseDB2zSQLParserListener) ExitNullIndicatorStructure(ctx *NullIndicatorStructureContext) {}

// EnterGlobalVariableName is called when production globalVariableName is entered.
func (s *BaseDB2zSQLParserListener) EnterGlobalVariableName(ctx *GlobalVariableNameContext) {}

// ExitGlobalVariableName is called when production globalVariableName is exited.
func (s *BaseDB2zSQLParserListener) ExitGlobalVariableName(ctx *GlobalVariableNameContext) {}

// EnterSqlParameterName is called when production sqlParameterName is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlParameterName(ctx *SqlParameterNameContext) {}

// ExitSqlParameterName is called when production sqlParameterName is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlParameterName(ctx *SqlParameterNameContext) {}

// EnterSqlVariableName is called when production sqlVariableName is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlVariableName(ctx *SqlVariableNameContext) {}

// ExitSqlVariableName is called when production sqlVariableName is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlVariableName(ctx *SqlVariableNameContext) {}

// EnterTransitionVariableName is called when production transitionVariableName is entered.
func (s *BaseDB2zSQLParserListener) EnterTransitionVariableName(ctx *TransitionVariableNameContext) {}

// ExitTransitionVariableName is called when production transitionVariableName is exited.
func (s *BaseDB2zSQLParserListener) ExitTransitionVariableName(ctx *TransitionVariableNameContext) {}

// EnterSynonym is called when production synonym is entered.
func (s *BaseDB2zSQLParserListener) EnterSynonym(ctx *SynonymContext) {}

// ExitSynonym is called when production synonym is exited.
func (s *BaseDB2zSQLParserListener) ExitSynonym(ctx *SynonymContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseDB2zSQLParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseDB2zSQLParserListener) ExitVariable(ctx *VariableContext) {}

// EnterIntoClause is called when production intoClause is entered.
func (s *BaseDB2zSQLParserListener) EnterIntoClause(ctx *IntoClauseContext) {}

// ExitIntoClause is called when production intoClause is exited.
func (s *BaseDB2zSQLParserListener) ExitIntoClause(ctx *IntoClauseContext) {}

// EnterCorrelationClause is called when production correlationClause is entered.
func (s *BaseDB2zSQLParserListener) EnterCorrelationClause(ctx *CorrelationClauseContext) {}

// ExitCorrelationClause is called when production correlationClause is exited.
func (s *BaseDB2zSQLParserListener) ExitCorrelationClause(ctx *CorrelationClauseContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseDB2zSQLParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseDB2zSQLParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterTableReference is called when production tableReference is entered.
func (s *BaseDB2zSQLParserListener) EnterTableReference(ctx *TableReferenceContext) {}

// ExitTableReference is called when production tableReference is exited.
func (s *BaseDB2zSQLParserListener) ExitTableReference(ctx *TableReferenceContext) {}

// EnterSingleTableReference is called when production singleTableReference is entered.
func (s *BaseDB2zSQLParserListener) EnterSingleTableReference(ctx *SingleTableReferenceContext) {}

// ExitSingleTableReference is called when production singleTableReference is exited.
func (s *BaseDB2zSQLParserListener) ExitSingleTableReference(ctx *SingleTableReferenceContext) {}

// EnterPeriodSpecification is called when production periodSpecification is entered.
func (s *BaseDB2zSQLParserListener) EnterPeriodSpecification(ctx *PeriodSpecificationContext) {}

// ExitPeriodSpecification is called when production periodSpecification is exited.
func (s *BaseDB2zSQLParserListener) ExitPeriodSpecification(ctx *PeriodSpecificationContext) {}

// EnterPeriodClause is called when production periodClause is entered.
func (s *BaseDB2zSQLParserListener) EnterPeriodClause(ctx *PeriodClauseContext) {}

// ExitPeriodClause is called when production periodClause is exited.
func (s *BaseDB2zSQLParserListener) ExitPeriodClause(ctx *PeriodClauseContext) {}

// EnterNestedTableExpression is called when production nestedTableExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterNestedTableExpression(ctx *NestedTableExpressionContext) {}

// ExitNestedTableExpression is called when production nestedTableExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitNestedTableExpression(ctx *NestedTableExpressionContext) {}

// EnterDataChangeTableReference is called when production dataChangeTableReference is entered.
func (s *BaseDB2zSQLParserListener) EnterDataChangeTableReference(ctx *DataChangeTableReferenceContext) {
}

// ExitDataChangeTableReference is called when production dataChangeTableReference is exited.
func (s *BaseDB2zSQLParserListener) ExitDataChangeTableReference(ctx *DataChangeTableReferenceContext) {
}

// EnterTableFunctionReference is called when production tableFunctionReference is entered.
func (s *BaseDB2zSQLParserListener) EnterTableFunctionReference(ctx *TableFunctionReferenceContext) {}

// ExitTableFunctionReference is called when production tableFunctionReference is exited.
func (s *BaseDB2zSQLParserListener) ExitTableFunctionReference(ctx *TableFunctionReferenceContext) {}

// EnterTableUdfCardinalityClause is called when production tableUdfCardinalityClause is entered.
func (s *BaseDB2zSQLParserListener) EnterTableUdfCardinalityClause(ctx *TableUdfCardinalityClauseContext) {
}

// ExitTableUdfCardinalityClause is called when production tableUdfCardinalityClause is exited.
func (s *BaseDB2zSQLParserListener) ExitTableUdfCardinalityClause(ctx *TableUdfCardinalityClauseContext) {
}

// EnterTypedCorrelationClause is called when production typedCorrelationClause is entered.
func (s *BaseDB2zSQLParserListener) EnterTypedCorrelationClause(ctx *TypedCorrelationClauseContext) {}

// ExitTypedCorrelationClause is called when production typedCorrelationClause is exited.
func (s *BaseDB2zSQLParserListener) ExitTypedCorrelationClause(ctx *TypedCorrelationClauseContext) {}

// EnterTableLocatorReference is called when production tableLocatorReference is entered.
func (s *BaseDB2zSQLParserListener) EnterTableLocatorReference(ctx *TableLocatorReferenceContext) {}

// ExitTableLocatorReference is called when production tableLocatorReference is exited.
func (s *BaseDB2zSQLParserListener) ExitTableLocatorReference(ctx *TableLocatorReferenceContext) {}

// EnterXmltableExpression is called when production xmltableExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterXmltableExpression(ctx *XmltableExpressionContext) {}

// ExitXmltableExpression is called when production xmltableExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitXmltableExpression(ctx *XmltableExpressionContext) {}

// EnterCollectionDerivedTable is called when production collectionDerivedTable is entered.
func (s *BaseDB2zSQLParserListener) EnterCollectionDerivedTable(ctx *CollectionDerivedTableContext) {}

// ExitCollectionDerivedTable is called when production collectionDerivedTable is exited.
func (s *BaseDB2zSQLParserListener) ExitCollectionDerivedTable(ctx *CollectionDerivedTableContext) {}

// EnterJoinCondition is called when production joinCondition is entered.
func (s *BaseDB2zSQLParserListener) EnterJoinCondition(ctx *JoinConditionContext) {}

// ExitJoinCondition is called when production joinCondition is exited.
func (s *BaseDB2zSQLParserListener) ExitJoinCondition(ctx *JoinConditionContext) {}

// EnterFullJoinExpression is called when production fullJoinExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterFullJoinExpression(ctx *FullJoinExpressionContext) {}

// ExitFullJoinExpression is called when production fullJoinExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitFullJoinExpression(ctx *FullJoinExpressionContext) {}

// EnterCastFunction is called when production castFunction is entered.
func (s *BaseDB2zSQLParserListener) EnterCastFunction(ctx *CastFunctionContext) {}

// ExitCastFunction is called when production castFunction is exited.
func (s *BaseDB2zSQLParserListener) ExitCastFunction(ctx *CastFunctionContext) {}

// EnterOrdinaryArrayExpression is called when production ordinaryArrayExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterOrdinaryArrayExpression(ctx *OrdinaryArrayExpressionContext) {
}

// ExitOrdinaryArrayExpression is called when production ordinaryArrayExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitOrdinaryArrayExpression(ctx *OrdinaryArrayExpressionContext) {
}

// EnterAssociativeArrayExpression is called when production associativeArrayExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterAssociativeArrayExpression(ctx *AssociativeArrayExpressionContext) {
}

// ExitAssociativeArrayExpression is called when production associativeArrayExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitAssociativeArrayExpression(ctx *AssociativeArrayExpressionContext) {
}

// EnterComparison is called when production comparison is entered.
func (s *BaseDB2zSQLParserListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseDB2zSQLParserListener) ExitComparison(ctx *ComparisonContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseDB2zSQLParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseDB2zSQLParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseDB2zSQLParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseDB2zSQLParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseDB2zSQLParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseDB2zSQLParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterGroupingExpression is called when production groupingExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterGroupingExpression(ctx *GroupingExpressionContext) {}

// ExitGroupingExpression is called when production groupingExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitGroupingExpression(ctx *GroupingExpressionContext) {}

// EnterGroupingSets is called when production groupingSets is entered.
func (s *BaseDB2zSQLParserListener) EnterGroupingSets(ctx *GroupingSetsContext) {}

// ExitGroupingSets is called when production groupingSets is exited.
func (s *BaseDB2zSQLParserListener) ExitGroupingSets(ctx *GroupingSetsContext) {}

// EnterGroupingSetsGroup is called when production groupingSetsGroup is entered.
func (s *BaseDB2zSQLParserListener) EnterGroupingSetsGroup(ctx *GroupingSetsGroupContext) {}

// ExitGroupingSetsGroup is called when production groupingSetsGroup is exited.
func (s *BaseDB2zSQLParserListener) ExitGroupingSetsGroup(ctx *GroupingSetsGroupContext) {}

// EnterSuperGroups is called when production superGroups is entered.
func (s *BaseDB2zSQLParserListener) EnterSuperGroups(ctx *SuperGroupsContext) {}

// ExitSuperGroups is called when production superGroups is exited.
func (s *BaseDB2zSQLParserListener) ExitSuperGroups(ctx *SuperGroupsContext) {}

// EnterSelectColumns is called when production selectColumns is entered.
func (s *BaseDB2zSQLParserListener) EnterSelectColumns(ctx *SelectColumnsContext) {}

// ExitSelectColumns is called when production selectColumns is exited.
func (s *BaseDB2zSQLParserListener) ExitSelectColumns(ctx *SelectColumnsContext) {}

// EnterUnpackedRow is called when production unpackedRow is entered.
func (s *BaseDB2zSQLParserListener) EnterUnpackedRow(ctx *UnpackedRowContext) {}

// ExitUnpackedRow is called when production unpackedRow is exited.
func (s *BaseDB2zSQLParserListener) ExitUnpackedRow(ctx *UnpackedRowContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseDB2zSQLParserListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseDB2zSQLParserListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterSubSelect is called when production subSelect is entered.
func (s *BaseDB2zSQLParserListener) EnterSubSelect(ctx *SubSelectContext) {}

// ExitSubSelect is called when production subSelect is exited.
func (s *BaseDB2zSQLParserListener) ExitSubSelect(ctx *SubSelectContext) {}

// EnterSelectIntoStatement is called when production selectIntoStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSelectIntoStatement(ctx *SelectIntoStatementContext) {}

// ExitSelectIntoStatement is called when production selectIntoStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSelectIntoStatement(ctx *SelectIntoStatementContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseDB2zSQLParserListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseDB2zSQLParserListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterCommonTableExpression is called when production commonTableExpression is entered.
func (s *BaseDB2zSQLParserListener) EnterCommonTableExpression(ctx *CommonTableExpressionContext) {}

// ExitCommonTableExpression is called when production commonTableExpression is exited.
func (s *BaseDB2zSQLParserListener) ExitCommonTableExpression(ctx *CommonTableExpressionContext) {}

// EnterUpdateClause is called when production updateClause is entered.
func (s *BaseDB2zSQLParserListener) EnterUpdateClause(ctx *UpdateClauseContext) {}

// ExitUpdateClause is called when production updateClause is exited.
func (s *BaseDB2zSQLParserListener) ExitUpdateClause(ctx *UpdateClauseContext) {}

// EnterReadOnlyClause is called when production readOnlyClause is entered.
func (s *BaseDB2zSQLParserListener) EnterReadOnlyClause(ctx *ReadOnlyClauseContext) {}

// ExitReadOnlyClause is called when production readOnlyClause is exited.
func (s *BaseDB2zSQLParserListener) ExitReadOnlyClause(ctx *ReadOnlyClauseContext) {}

// EnterOptimizeClause is called when production optimizeClause is entered.
func (s *BaseDB2zSQLParserListener) EnterOptimizeClause(ctx *OptimizeClauseContext) {}

// ExitOptimizeClause is called when production optimizeClause is exited.
func (s *BaseDB2zSQLParserListener) ExitOptimizeClause(ctx *OptimizeClauseContext) {}

// EnterIsolationClause is called when production isolationClause is entered.
func (s *BaseDB2zSQLParserListener) EnterIsolationClause(ctx *IsolationClauseContext) {}

// ExitIsolationClause is called when production isolationClause is exited.
func (s *BaseDB2zSQLParserListener) ExitIsolationClause(ctx *IsolationClauseContext) {}

// EnterLockClause is called when production lockClause is entered.
func (s *BaseDB2zSQLParserListener) EnterLockClause(ctx *LockClauseContext) {}

// ExitLockClause is called when production lockClause is exited.
func (s *BaseDB2zSQLParserListener) ExitLockClause(ctx *LockClauseContext) {}

// EnterSkipLockedDataClause is called when production skipLockedDataClause is entered.
func (s *BaseDB2zSQLParserListener) EnterSkipLockedDataClause(ctx *SkipLockedDataClauseContext) {}

// ExitSkipLockedDataClause is called when production skipLockedDataClause is exited.
func (s *BaseDB2zSQLParserListener) ExitSkipLockedDataClause(ctx *SkipLockedDataClauseContext) {}

// EnterQuerynoClause is called when production querynoClause is entered.
func (s *BaseDB2zSQLParserListener) EnterQuerynoClause(ctx *QuerynoClauseContext) {}

// ExitQuerynoClause is called when production querynoClause is exited.
func (s *BaseDB2zSQLParserListener) ExitQuerynoClause(ctx *QuerynoClauseContext) {}

// EnterScalarFullSelect is called when production scalarFullSelect is entered.
func (s *BaseDB2zSQLParserListener) EnterScalarFullSelect(ctx *ScalarFullSelectContext) {}

// ExitScalarFullSelect is called when production scalarFullSelect is exited.
func (s *BaseDB2zSQLParserListener) ExitScalarFullSelect(ctx *ScalarFullSelectContext) {}

// EnterFullSelect is called when production fullSelect is entered.
func (s *BaseDB2zSQLParserListener) EnterFullSelect(ctx *FullSelectContext) {}

// ExitFullSelect is called when production fullSelect is exited.
func (s *BaseDB2zSQLParserListener) ExitFullSelect(ctx *FullSelectContext) {}

// EnterValuesClause is called when production valuesClause is entered.
func (s *BaseDB2zSQLParserListener) EnterValuesClause(ctx *ValuesClauseContext) {}

// ExitValuesClause is called when production valuesClause is exited.
func (s *BaseDB2zSQLParserListener) ExitValuesClause(ctx *ValuesClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseDB2zSQLParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseDB2zSQLParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterSortKey is called when production sortKey is entered.
func (s *BaseDB2zSQLParserListener) EnterSortKey(ctx *SortKeyContext) {}

// ExitSortKey is called when production sortKey is exited.
func (s *BaseDB2zSQLParserListener) ExitSortKey(ctx *SortKeyContext) {}

// EnterOffsetClause is called when production offsetClause is entered.
func (s *BaseDB2zSQLParserListener) EnterOffsetClause(ctx *OffsetClauseContext) {}

// ExitOffsetClause is called when production offsetClause is exited.
func (s *BaseDB2zSQLParserListener) ExitOffsetClause(ctx *OffsetClauseContext) {}

// EnterFetchClause is called when production fetchClause is entered.
func (s *BaseDB2zSQLParserListener) EnterFetchClause(ctx *FetchClauseContext) {}

// ExitFetchClause is called when production fetchClause is exited.
func (s *BaseDB2zSQLParserListener) ExitFetchClause(ctx *FetchClauseContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseDB2zSQLParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseDB2zSQLParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifier1 is called when production identifier1 is entered.
func (s *BaseDB2zSQLParserListener) EnterIdentifier1(ctx *Identifier1Context) {}

// ExitIdentifier1 is called when production identifier1 is exited.
func (s *BaseDB2zSQLParserListener) ExitIdentifier1(ctx *Identifier1Context) {}

// EnterSqlidentifier is called when production sqlidentifier is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlidentifier(ctx *SqlidentifierContext) {}

// ExitSqlidentifier is called when production sqlidentifier is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlidentifier(ctx *SqlidentifierContext) {}

// EnterSqlKeyword is called when production sqlKeyword is entered.
func (s *BaseDB2zSQLParserListener) EnterSqlKeyword(ctx *SqlKeywordContext) {}

// ExitSqlKeyword is called when production sqlKeyword is exited.
func (s *BaseDB2zSQLParserListener) ExitSqlKeyword(ctx *SqlKeywordContext) {}
