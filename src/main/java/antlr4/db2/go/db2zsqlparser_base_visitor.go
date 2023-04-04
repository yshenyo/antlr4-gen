// Code generated from DB2zSQLParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package db2 // DB2zSQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseDB2zSQLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDB2zSQLParserVisitor) VisitStartRule(ctx *StartRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlStatement(ctx *SqlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCursorName(ctx *CursorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStatementName(ctx *StatementNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDescriptorName(ctx *DescriptorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitHoldability(ctx *HoldabilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitReturnability(ctx *ReturnabilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRowsetPositioning(ctx *RowsetPositioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNotNullPhrase(ctx *NotNullPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAllocateCursorStatement(ctx *AllocateCursorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRsLocatorVariable(ctx *RsLocatorVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterDatabaseStatement(ctx *AlterDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterFunctionStatement(ctx *AlterFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterFunctionStatementExternal(ctx *AlterFunctionStatementExternalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterFunctionStatementCompiledSqlScalar(ctx *AlterFunctionStatementCompiledSqlScalarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterWhichFunction1(ctx *AlterWhichFunction1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterWhichFunction2(ctx *AlterWhichFunction2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionCompiledSqlScalarRoutineSpecification(ctx *FunctionCompiledSqlScalarRoutineSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterFunctionStatementInlineSqlScalar(ctx *AlterFunctionStatementInlineSqlScalarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterFunctionStatementSqlTable(ctx *AlterFunctionStatementSqlTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionReturnsClause(ctx *FunctionReturnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionDesignator1(ctx *FunctionDesignator1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionDesignator2(ctx *FunctionDesignator2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterIndexStatement(ctx *AlterIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterMaskStatement(ctx *AlterMaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterPermissionStatement(ctx *AlterPermissionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterProcedureStatement(ctx *AlterProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterProcedureSQLPLStatement(ctx *AlterProcedureSQLPLStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterWhichProcedureSQLPL1(ctx *AlterWhichProcedureSQLPL1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterWhichProcedureSQLPL2(ctx *AlterWhichProcedureSQLPL2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitApplicationCompatibilityPhrase(ctx *ApplicationCompatibilityPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterSequenceStatement(ctx *AlterSequenceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterStogroupStatement(ctx *AlterStogroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTablespaceStatement(ctx *AlterTablespaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTriggerStatement(ctx *AlterTriggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTrustedContextStatement(ctx *AlterTrustedContextStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterViewStatement(ctx *AlterViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAssociateLocatorsStatement(ctx *AssociateLocatorsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBeginDeclareSectionStatement(ctx *BeginDeclareSectionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCallStatement(ctx *CallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCloseStatement(ctx *CloseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCommentStatement(ctx *CommentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCommitStatement(ctx *CommitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitConnectStatement(ctx *ConnectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateAliasStatement(ctx *CreateAliasStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateAuxiliaryTableStatement(ctx *CreateAuxiliaryTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementExternalScalar(ctx *CreateFunctionStatementExternalScalarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementExternalScalarReturnsPhrase(ctx *CreateFunctionStatementExternalScalarReturnsPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementExternalTable(ctx *CreateFunctionStatementExternalTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementExternalTableReturnsPhrase(ctx *CreateFunctionStatementExternalTableReturnsPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExternalTableFunctionColumn(ctx *ExternalTableFunctionColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementSourced(ctx *CreateFunctionStatementSourcedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementSourcedReturnsPhrase(ctx *CreateFunctionStatementSourcedReturnsPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementSourcedSourcePhrase(ctx *CreateFunctionStatementSourcedSourcePhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementInlineSqlScalar(ctx *CreateFunctionStatementInlineSqlScalarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementCompiledSqlScalar(ctx *CreateFunctionStatementCompiledSqlScalarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementSqlTable(ctx *CreateFunctionStatementSqlTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateGlobalTemporaryTableStatement(ctx *CreateGlobalTemporaryTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateLobTablespaceStatement(ctx *CreateLobTablespaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateMaskStatement(ctx *CreateMaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreatePermissionStatement(ctx *CreatePermissionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateProcedureSQLPLStatement(ctx *CreateProcedureSQLPLStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlRoutineBody(ctx *SqlRoutineBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitObfuscatedStatementText(ctx *ObfuscatedStatementTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitProbablySQLPL(ctx *ProbablySQLPLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateProcedureStatement(ctx *CreateProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateSequenceStatement(ctx *CreateSequenceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateStogroupStatement(ctx *CreateStogroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTableOptions(ctx *CreateTableOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTablespaceStatement(ctx *CreateTablespaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTriggerStatement(ctx *CreateTriggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTrustedContextStatement(ctx *CreateTrustedContextStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTypeArrayStatement(ctx *CreateTypeArrayStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTypeDistinctStatement(ctx *CreateTypeDistinctStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateVariableStatement(ctx *CreateVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeclareCursorStatement(ctx *DeclareCursorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeclareGlobalTemporaryTableStatement(ctx *DeclareGlobalTemporaryTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeclareTableStatement(ctx *DeclareTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeclareStatementStatement(ctx *DeclareStatementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeclareVariableStatement(ctx *DeclareVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDescribeStatement(ctx *DescribeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDescribeCursorStatement(ctx *DescribeCursorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDescribeInputStatement(ctx *DescribeInputStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDescribeOutputStatement(ctx *DescribeOutputStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDescribeProcedureStatement(ctx *DescribeProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDescribeTableStatement(ctx *DescribeTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropStatement(ctx *DropStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitEndDeclareSectionStatement(ctx *EndDeclareSectionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExchangeStatement(ctx *ExchangeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExecuteImmediateStatement(ctx *ExecuteImmediateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExplainStatement(ctx *ExplainStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFetchStatement(ctx *FetchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFreeLocatorStatement(ctx *FreeLocatorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGetDiagnosticsStatement(ctx *GetDiagnosticsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantStatement(ctx *GrantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitHoldLocatorStatement(ctx *HoldLocatorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIncludeStatement(ctx *IncludeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLabelStatement(ctx *LabelStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLockTableStatement(ctx *LockTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMergeStatement(ctx *MergeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOpenStatement(ctx *OpenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPrepareStatement(ctx *PrepareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRefreshTableStatement(ctx *RefreshTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitReleaseConnectionStatement(ctx *ReleaseConnectionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitReleaseSavepointStatement(ctx *ReleaseSavepointStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRenameStatement(ctx *RenameStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeStatement(ctx *RevokeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRollbackStatement(ctx *RollbackStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSavepointStatement(ctx *SavepointStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSetAssignmentStatement(ctx *SetAssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSetConnectionStatement(ctx *SetConnectionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSetEncryptionPasswordStatement(ctx *SetEncryptionPasswordStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSetPathStatement(ctx *SetPathStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSetSchemaStatement(ctx *SetSchemaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSetSessionTimezoneStatement(ctx *SetSessionTimezoneStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSetSpecialRegisterStatement(ctx *SetSpecialRegisterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSignalStatement(ctx *SignalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTransferOwnershipStatement(ctx *TransferOwnershipStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTruncateStatement(ctx *TruncateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesStatement(ctx *ValuesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesIntoStatement(ctx *ValuesIntoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWheneverStatement(ctx *WheneverStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesIntoTargetVariable(ctx *ValuesIntoTargetVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOwnedObject(ctx *OwnedObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNewOwner(ctx *NewOwnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantCollectionStatement(ctx *GrantCollectionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantDatabaseStatement(ctx *GrantDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantFunctionOrProcedureStatement(ctx *GrantFunctionOrProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantPackageStatement(ctx *GrantPackageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantPlanStatement(ctx *GrantPlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantSchemaStatement(ctx *GrantSchemaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantSequenceStatement(ctx *GrantSequenceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantSystemStatement(ctx *GrantSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantTableStatement(ctx *GrantTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantTypeOrJarStatement(ctx *GrantTypeOrJarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantVariableStatement(ctx *GrantVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantUseOfStatement(ctx *GrantUseOfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeCollectionStatement(ctx *RevokeCollectionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeDatabaseStatement(ctx *RevokeDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeFunctionOrProcedureStatement(ctx *RevokeFunctionOrProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokePackageStatement(ctx *RevokePackageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokePlanStatement(ctx *RevokePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeSchemaStatement(ctx *RevokeSchemaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeSequenceStatement(ctx *RevokeSequenceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeSystemStatement(ctx *RevokeSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeTableStatement(ctx *RevokeTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeTypeOrJarStatement(ctx *RevokeTypeOrJarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeVariableStatement(ctx *RevokeVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeUseOfStatement(ctx *RevokeUseOfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantUseOfTarget(ctx *GrantUseOfTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantVariableAuthority(ctx *GrantVariableAuthorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantTableAuthority(ctx *GrantTableAuthorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantSystemAuthority(ctx *GrantSystemAuthorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantSequenceAuthority(ctx *GrantSequenceAuthorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantSchemaAuthority(ctx *GrantSchemaAuthorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantPlanAuthority(ctx *GrantPlanAuthorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantPackageAuthority(ctx *GrantPackageAuthorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPackageSpecification(ctx *PackageSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionSpecification(ctx *FunctionSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantee(ctx *GranteeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWithGrantOption(ctx *WithGrantOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeByOption(ctx *RevokeByOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRevokeDependentPrivilegesOption(ctx *RevokeDependentPrivilegesOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGrantDatabaseAuthority(ctx *GrantDatabaseAuthorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStatementInformation(ctx *StatementInformationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStatementInformationVariableEquate(ctx *StatementInformationVariableEquateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStatementInformationItemName(ctx *StatementInformationItemNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitConditionInformation(ctx *ConditionInformationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitConditionInformationVariableEquate(ctx *ConditionInformationVariableEquateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitConditionInformationItemName(ctx *ConditionInformationItemNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitConnectionInformationItemName(ctx *ConnectionInformationItemNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCombinedInformation(ctx *CombinedInformationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCombinedInformationOption(ctx *CombinedInformationOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFetchOrientation(ctx *FetchOrientationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRowPositioned(ctx *RowPositionedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRowsetPositioned(ctx *RowsetPositionedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSingleRowFetch(ctx *SingleRowFetchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFetchTargetVariable(ctx *FetchTargetVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMultipleRowFetch(ctx *MultipleRowFetchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMultipleRowFetchForClause(ctx *MultipleRowFetchForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMultipleRowFetchIntoClause(ctx *MultipleRowFetchIntoClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExplainPlanClause(ctx *ExplainPlanClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExplainStmtcacheClause(ctx *ExplainStmtcacheClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExplainPackageClause(ctx *ExplainPackageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExplainStabilizedDynamicQueryClause(ctx *ExplainStabilizedDynamicQueryClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPackageScopeSpecification(ctx *PackageScopeSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCollectionName(ctx *CollectionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPackageScopePackageName(ctx *PackageScopePackageNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitVersionName(ctx *VersionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSourceRowData(ctx *SourceRowDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAliasDesignation(ctx *AliasDesignationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropDatabaseClause(ctx *DropDatabaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropFunctionClause(ctx *DropFunctionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropIndexClause(ctx *DropIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropMaskClause(ctx *DropMaskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropPackageClause(ctx *DropPackageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropPermissionClause(ctx *DropPermissionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropProcedureClause(ctx *DropProcedureClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropRoleClause(ctx *DropRoleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropSequenceClause(ctx *DropSequenceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropStogroupClause(ctx *DropStogroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropSynonymClause(ctx *DropSynonymClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropTableClause(ctx *DropTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropTablespaceClause(ctx *DropTablespaceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropTriggerClause(ctx *DropTriggerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropTrustedContextClause(ctx *DropTrustedContextClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropTypeClause(ctx *DropTypeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropVariableClause(ctx *DropVariableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropViewClause(ctx *DropViewClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPackageDesignator(ctx *PackageDesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDescribeUsingOption(ctx *DescribeUsingOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeclareGlobalTemporaryTableLikeClause(ctx *DeclareGlobalTemporaryTableLikeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOnCommitClause(ctx *OnCommitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLoggedWithRollbackClause(ctx *LoggedWithRollbackClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateViewCheckOptionClause(ctx *CreateViewCheckOptionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrustedContextDefaultRoleClause(ctx *TrustedContextDefaultRoleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrustedContextEnableDisableClause(ctx *TrustedContextEnableDisableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrustedContextDefaultSecurityLabelClause(ctx *TrustedContextDefaultSecurityLabelClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrustedContextAttributesClause(ctx *TrustedContextAttributesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrustedContextWithUseForClause(ctx *TrustedContextWithUseForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrustedContextAttribute1(ctx *TrustedContextAttribute1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrustedContextAttribute2(ctx *TrustedContextAttribute2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrustedContextUseFor(ctx *TrustedContextUseForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUserOptions(ctx *UserOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTriggerDefinition(ctx *TriggerDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTriggerActivationTime(ctx *TriggerActivationTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTriggerEvent(ctx *TriggerEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTriggerGranularity(ctx *TriggerGranularityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTriggeredAction(ctx *TriggeredActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlTriggerBody(ctx *SqlTriggerBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTriggeredSqlStatement(ctx *TriggeredSqlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTriggerDefinitionOption(ctx *TriggerDefinitionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTableInClause(ctx *CreateTableInClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCustomVolatileClause(ctx *CustomVolatileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTableColumnDefinition(ctx *CreateTableColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitEditprocClause(ctx *EditprocClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValidprocClause(ctx *ValidprocClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAuditClause(ctx *AuditClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitObidClause(ctx *ObidClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDataCaptureClause(ctx *DataCaptureClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRestrictOnDropClause(ctx *RestrictOnDropClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCcsidClause1(ctx *CcsidClause1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCcsidClause2(ctx *CcsidClause2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCardinalityClause(ctx *CardinalityClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAppendClause(ctx *AppendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMemberClause(ctx *MemberClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrackmodClause(ctx *TrackmodClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPagenumClause(ctx *PagenumClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFieldprocClause(ctx *FieldprocClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsSecurityLabelClause(ctx *AsSecurityLabelClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitImplicitlyHiddenClause(ctx *ImplicitlyHiddenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitInlineLengthClause(ctx *InlineLengthClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCopyOptions(ctx *CopyOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCopyOptionIdentity(ctx *CopyOptionIdentityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCopyOptionRowChangeTimestamp(ctx *CopyOptionRowChangeTimestampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCopyOptionColumnDefaults(ctx *CopyOptionColumnDefaultsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCopyOptionXmlTypeModifiers(ctx *CopyOptionXmlTypeModifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsResultTable(ctx *AsResultTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeclareGlobalTemporaryTableAsResultTable(ctx *DeclareGlobalTemporaryTableAsResultTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTableMaterializedQueryDefinition(ctx *CreateTableMaterializedQueryDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTableColumnConstraint(ctx *CreateTableColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOrganizationClause(ctx *OrganizationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateGlobalTemporaryTableColumnDefinition(ctx *CreateGlobalTemporaryTableColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeclareGlobalTemporaryTableColumnDefinition(ctx *DeclareGlobalTemporaryTableColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterDeclaration1(ctx *ParameterDeclaration1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterDeclaration2(ctx *ParameterDeclaration2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterDeclaration3(ctx *ParameterDeclaration3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementExternalScalarOptions(ctx *CreateFunctionStatementExternalScalarOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExternalNameOption1(ctx *ExternalNameOption1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExternalNameOption2(ctx *ExternalNameOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDynamicResultSetOption(ctx *DynamicResultSetOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLanguageOption1(ctx *LanguageOption1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLanguageOption2(ctx *LanguageOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLanguageOption3(ctx *LanguageOption3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLanguageOption4(ctx *LanguageOption4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLanguageOption5(ctx *LanguageOption5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterStyleOption1(ctx *ParameterStyleOption1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterStyleOption2(ctx *ParameterStyleOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterStyleOption3(ctx *ParameterStyleOption3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeterministicOption(ctx *DeterministicOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFencedOption(ctx *FencedOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNullInputOption1(ctx *NullInputOption1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNullInputOption2(ctx *NullInputOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNullInputOption3(ctx *NullInputOption3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNullInputOption4(ctx *NullInputOption4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDebugOption(ctx *DebugOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlDataOption1(ctx *SqlDataOption1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlDataOption2(ctx *SqlDataOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlDataOption3(ctx *SqlDataOption3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlDataOption4(ctx *SqlDataOption4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExternalActionOption(ctx *ExternalActionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPackagePathOption(ctx *PackagePathOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitScratchpadOption(ctx *ScratchpadOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFinalCallOption(ctx *FinalCallOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParallelOption1(ctx *ParallelOption1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParallelOption2(ctx *ParallelOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDbinfoOption(ctx *DbinfoOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCardinalityOption(ctx *CardinalityOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCollectionIdOption(ctx *CollectionIdOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWlmEnvironmentOption1(ctx *WlmEnvironmentOption1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWlmEnvironmentOption2(ctx *WlmEnvironmentOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWlmEnvironmentOption3(ctx *WlmEnvironmentOption3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsuTimeOption(ctx *AsuTimeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStayResidentOption(ctx *StayResidentOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitProgramTypeOption(ctx *ProgramTypeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSecurityOption(ctx *SecurityOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStopAfterFailureOption(ctx *StopAfterFailureOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRunOptionsOption(ctx *RunOptionsOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCommitOnReturnOption(ctx *CommitOnReturnOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSpecialRegistersOption(ctx *SpecialRegistersOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSpecialRegistersOption2(ctx *SpecialRegistersOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDispatchOption(ctx *DispatchOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSecuredOption(ctx *SecuredOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSpecificNameOption1(ctx *SpecificNameOption1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSpecificNameOption2(ctx *SpecificNameOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterOption1(ctx *ParameterOption1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterOption2(ctx *ParameterOption2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementExternalTableOptions(ctx *CreateFunctionStatementExternalTableOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementSourcedOptions(ctx *CreateFunctionStatementSourcedOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitInlineSqlScalarFunctionDefinition(ctx *InlineSqlScalarFunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementInlineSqlScalarOptions(ctx *CreateFunctionStatementInlineSqlScalarOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCompiledSqlScalarFunctionDefinition(ctx *CompiledSqlScalarFunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementCompiledSqlScalarOptions(ctx *CreateFunctionStatementCompiledSqlScalarOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlTableFunctionDefinition(ctx *SqlTableFunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateFunctionStatementSqlTableOptions(ctx *CreateFunctionStatementSqlTableOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSequenceAlias(ctx *SequenceAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTableAlias(ctx *TableAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAuthorization(ctx *AuthorizationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSearchedDelete(ctx *SearchedDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPositionedDelete(ctx *PositionedDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSearchedUpdate(ctx *SearchedUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPositionedUpdate(ctx *PositionedUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSourceValues(ctx *SourceValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesSingleRow(ctx *ValuesSingleRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesMultipleRow(ctx *ValuesMultipleRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMatchingCondition(ctx *MatchingConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitModificationOperation(ctx *ModificationOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAssignmentClause(ctx *AssignmentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSetAssignmentClause(ctx *SetAssignmentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSetAssignmentTargetVariable(ctx *SetAssignmentTargetVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUpdateOperation(ctx *UpdateOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeleteOperation(ctx *DeleteOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitInsertOperation(ctx *InsertOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSignalInformation(ctx *SignalInformationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesList1(ctx *ValuesList1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesList2(ctx *ValuesList2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesList3(ctx *ValuesList3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesList4(ctx *ValuesList4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIncludeColumns(ctx *IncludeColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMultipleRowInsert(ctx *MultipleRowInsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRegenerateClause(ctx *RegenerateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterIndexOptions(ctx *AlterIndexOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBufferpoolOption(ctx *BufferpoolOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCloseOption(ctx *CloseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCopyOption(ctx *CopyOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDssizeOption(ctx *DssizeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPiecesizeOption(ctx *PiecesizeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitClusterOption(ctx *ClusterOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPaddedOption(ctx *PaddedOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCompressOption(ctx *CompressOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDefineOption(ctx *DefineOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLocksizeOption(ctx *LocksizeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLockmaxOption(ctx *LockmaxOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitEnableDisableOption(ctx *EnableDisableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLoggedOption(ctx *LoggedOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNotAtomicPhrase(ctx *NotAtomicPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterIndexPartitionOptions(ctx *AlterIndexPartitionOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUsingSpecification1(ctx *UsingSpecification1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFreeSpecification(ctx *FreeSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGbpcacheSpecification(ctx *GbpcacheSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitionElement(ctx *PartitionElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitApplCompatValue(ctx *ApplCompatValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionLevel(ctx *FunctionLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionParameterType(ctx *FunctionParameterTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionDataType(ctx *FunctionDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionBuiltInType(ctx *FunctionBuiltInTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitProcedureBuiltinType(ctx *ProcedureBuiltinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTypeArrayBuiltinType(ctx *CreateTypeArrayBuiltinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTypeArrayBuiltinType2(ctx *CreateTypeArrayBuiltinType2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateVariableBuiltInType(ctx *CreateVariableBuiltInTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSourceDataType(ctx *SourceDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionExternalOptionList(ctx *FunctionExternalOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionCompiledSqlScalarOptionList(ctx *FunctionCompiledSqlScalarOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionInlineSqlScalarOptionList(ctx *FunctionInlineSqlScalarOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionSqlTableOptionList(ctx *FunctionSqlTableOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitProcedureOptionList(ctx *ProcedureOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateProcedureOptionList(ctx *CreateProcedureOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitProcedureSQLPLOptionList(ctx *ProcedureSQLPLOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitVersionOption(ctx *VersionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCommitOnReturnOptionSQLPL(ctx *CommitOnReturnOptionSQLPLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSchemaQualifierOption(ctx *SchemaQualifierOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCurrentDataOption(ctx *CurrentDataOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitImmediateWriteOption(ctx *ImmediateWriteOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExplainOption(ctx *ExplainOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitReoptOption(ctx *ReoptOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPackageOwnerOption(ctx *PackageOwnerOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDeferPrepareOption(ctx *DeferPrepareOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDegreeOption(ctx *DegreeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDynamicRulesOption(ctx *DynamicRulesOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitConcurrentAccessOption(ctx *ConcurrentAccessOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitApplicationEncodingOption(ctx *ApplicationEncodingOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIsolationLevelOption(ctx *IsolationLevelOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitKeepDynamicOption(ctx *KeepDynamicOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOpthintOption(ctx *OpthintOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlPathOption(ctx *SqlPathOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlPathOptionItem(ctx *SqlPathOptionItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitQueryAccelerationOption(ctx *QueryAccelerationOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitQueryAccelerationOptionItem(ctx *QueryAccelerationOptionItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGetAccelArchiveOption(ctx *GetAccelArchiveOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAccelerationOption(ctx *AccelerationOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitReleaseAtOption(ctx *ReleaseAtOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBusinessTimeSensitiveOption(ctx *BusinessTimeSensitiveOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSystemTimeSensitiveOption(ctx *SystemTimeSensitiveOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArchiveSensitiveOption(ctx *ArchiveSensitiveOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitApplcompatOption(ctx *ApplcompatOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValidateOption(ctx *ValidateOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRoundingOption(ctx *RoundingOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRoundingOptionItem(ctx *RoundingOptionItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDateFormatOption(ctx *DateFormatOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDateTimeFormatOptionItem(ctx *DateTimeFormatOptionItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTimeFormatOption(ctx *TimeFormatOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDecimalOption(ctx *DecimalOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitForUpdateOption(ctx *ForUpdateOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitConcentrateStatementsOption(ctx *ConcentrateStatementsOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAcceleratorOption(ctx *AcceleratorOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitProcedureDataType(ctx *ProcedureDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterSequenceOptionList(ctx *AlterSequenceOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateSequenceOptionList(ctx *CreateSequenceOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsTypeOption(ctx *AsTypeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStartOption(ctx *StartOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRestartOption(ctx *RestartOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIncrementOption(ctx *IncrementOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMinvalueOption(ctx *MinvalueOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMaxvalueOption(ctx *MaxvalueOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCycleOption(ctx *CycleOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCacheOption(ctx *CacheOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOrderOption(ctx *OrderOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitKeyLabelOption(ctx *KeyLabelOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDataclasOption(ctx *DataclasOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMgmtclasOption(ctx *MgmtclasOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStorclasOption(ctx *StorclasOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterStogroupOptionList(ctx *AlterStogroupOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTableOptionList(ctx *AlterTableOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTablespaceOptionList(ctx *AlterTablespaceOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateTablespaceOptionList(ctx *CreateTablespaceOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTrustedContextOptionList(ctx *TrustedContextOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDatabaseOptionList(ctx *DatabaseOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateIndexOptionList(ctx *CreateIndexOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCreateLobTablespaceOptionList(ctx *CreateLobTablespaceOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitInDatabaseOption(ctx *InDatabaseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSegsizeOption(ctx *SegsizeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNumpartsOption(ctx *NumpartsOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitionByGrowthSpecification(ctx *PartitionByGrowthSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitionByRangeSpecification(ctx *PartitionByRangeSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitionByRangePartitionPhrase(ctx *PartitionByRangePartitionPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitInsertAlgorithmOption(ctx *InsertAlgorithmOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMaxrowsOption(ctx *MaxrowsOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMaxpartitionsOption(ctx *MaxpartitionsOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUsingSpecification2(ctx *UsingSpecification2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStogroupOptions(ctx *StogroupOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlIndexSpecification(ctx *XmlIndexSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlPatternClause(ctx *XmlPatternClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterAttributesOptions(ctx *AlterAttributesOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAddAttributesOptions(ctx *AddAttributesOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDropAttributesOptions(ctx *DropAttributesOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIncludeColumnPhrase(ctx *IncludeColumnPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUserClause(ctx *UserClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUserClauseAddOptions(ctx *UserClauseAddOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUserClauseReplaceOptions(ctx *UserClauseReplaceOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUserClauseDropOptions(ctx *UserClauseDropOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUseOptions(ctx *UseOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterPartitionClause(ctx *AlterPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUsingBlock(ctx *UsingBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFreeBlock(ctx *FreeBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMoveTableClause(ctx *MoveTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGbpcacheBlock(ctx *GbpcacheBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAliasDesignator(ctx *AliasDesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMultipleColumnList(ctx *MultipleColumnListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionDesignator(ctx *FunctionDesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterType(ctx *ParameterTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTableColumnDefinitionOptionList1(ctx *AlterTableColumnDefinitionOptionList1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTableColumnDefinitionOptionList2(ctx *AlterTableColumnDefinitionOptionList2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitColumnConstraint(ctx *ColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGeneratedClause(ctx *GeneratedClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGeneratedClause2(ctx *GeneratedClause2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsIdentityClause(ctx *AsIdentityClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsIdentityClauseOptionList(ctx *AsIdentityClauseOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsRowChangeTimestampClause(ctx *AsRowChangeTimestampClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsRowTransactionStartIDClause(ctx *AsRowTransactionStartIDClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsRowTransactionTimestampClause(ctx *AsRowTransactionTimestampClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAsGeneratedExpressionClause(ctx *AsGeneratedExpressionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNonDeterministicExpression(ctx *NonDeterministicExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNonDeterministicExpressionSessionVariable(ctx *NonDeterministicExpressionSessionVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitColumnAlteration(ctx *ColumnAlterationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitColumnAlterationOptionList(ctx *ColumnAlterationOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlteredDataType(ctx *AlteredDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBuiltInType(ctx *BuiltInTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSequenceDataType(ctx *SequenceDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSequenceBuiltInType(ctx *SequenceBuiltInTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlDataType(ctx *SqlDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlTypeModifier(ctx *XmlTypeModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlSchemaSpecification(ctx *XmlSchemaSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlElementName(ctx *XmlElementNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPiName(ctx *PiNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRegisteredXmlSchemaName(ctx *RegisteredXmlSchemaNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTargetNamespace(ctx *TargetNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSchemaLocation(ctx *SchemaLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIdentityAlteration(ctx *IdentityAlterationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUniqueConstraint(ctx *UniqueConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitReferentialConstraint(ctx *ReferentialConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitReferencesClause(ctx *ReferencesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCheckConstraint(ctx *CheckConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitioningClause(ctx *PartitioningClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitionExpression(ctx *PartitionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitionLimitKey(ctx *PartitionLimitKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitioningPhrase(ctx *PartitioningPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitionHashSpace(ctx *PartitionHashSpaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterHashOrganization(ctx *AlterHashOrganizationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitioningClauseElement(ctx *PartitioningClauseElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPartitionClause(ctx *PartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRotatePartitionClause(ctx *RotatePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExtraRowOption(ctx *ExtraRowOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMaterializedQueryDefinition(ctx *MaterializedQueryDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMaterializedQueryAlteration(ctx *MaterializedQueryAlterationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRefreshableTableOptions(ctx *RefreshableTableOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDataInitiallyDeferredPhrase(ctx *DataInitiallyDeferredPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRefreshDeferredPhrase(ctx *RefreshDeferredPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRefreshableTableOptionsList(ctx *RefreshableTableOptionsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMaterializedQueryTableAlteration(ctx *MaterializedQueryTableAlterationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPeriodDefinition(ctx *PeriodDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTableColumnDefinition(ctx *AlterTableColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExternalProgramName(ctx *ExternalProgramNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPackagePath(ctx *PackagePathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCollectionID(ctx *CollectionIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRunTimeOptions(ctx *RunTimeOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitKeyExpression(ctx *KeyExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRowChangeExpression(ctx *RowChangeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSequenceReference(ctx *SequenceReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionInvocation(ctx *FunctionInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitScalarFunctionInvocation(ctx *ScalarFunctionInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAggregateFunctionInvocation(ctx *AggregateFunctionInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRegressionFunctionInvocation(ctx *RegressionFunctionInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExternalFunctionInvocation(ctx *ExternalFunctionInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLabeledDuration(ctx *LabeledDurationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDurationSuffix(ctx *DurationSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlCastSpecification(ctx *XmlCastSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlParseSpecification(ctx *XmlParseSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayElementSpecification(ctx *ArrayElementSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayIndex(ctx *ArrayIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayConstructor(ctx *ArrayConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOlapSpecification(ctx *OlapSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOrderedOlapSpecification(ctx *OrderedOlapSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOlapSpecificationFunction(ctx *OlapSpecificationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLagFunction(ctx *LagFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLeadFunction(ctx *LeadFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRespectNullsClause(ctx *RespectNullsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWindowPartitionClause(ctx *WindowPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWindowOrderClause(ctx *WindowOrderClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWindowOrderClauseQualifier(ctx *WindowOrderClauseQualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNumberingSpecification(ctx *NumberingSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAggregationSpecification(ctx *AggregationSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAggregateFunction(ctx *AggregateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRegressionFunction(ctx *RegressionFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOlapColumnFunction(ctx *OlapColumnFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFirstValueFunction(ctx *FirstValueFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLastValueFunction(ctx *LastValueFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNthValueFunction(ctx *NthValueFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRatioToReportFunction(ctx *RatioToReportFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitListaggFunction(ctx *ListaggFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayaggFunction(ctx *ArrayaggFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayaggOrdinaryFunction(ctx *ArrayaggOrdinaryFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayaggAssociativeFunction(ctx *ArrayaggAssociativeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCorrelationFunction(ctx *CorrelationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCovarianceFunction(ctx *CovarianceFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCovarianceSampFunction(ctx *CovarianceSampFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCumeDistFunction(ctx *CumeDistFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPercentileContFunction(ctx *PercentileContFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPercentileDiscFunction(ctx *PercentileDiscFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPercentRankFunction(ctx *PercentRankFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlaggFunction(ctx *XmlaggFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlaggOrderByClause(ctx *XmlaggOrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlaggOrderByOption(ctx *XmlaggOrderByOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAggregateOrderByClause(ctx *AggregateOrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAggregateOrderByOption(ctx *AggregateOrderByOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWindowAggregationGroupClause(ctx *WindowAggregationGroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGroupStart(ctx *GroupStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGroupBetween(ctx *GroupBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGroupEnd(ctx *GroupEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGroupBound1(ctx *GroupBound1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGroupBound2(ctx *GroupBound2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUnboundedPreceding(ctx *UnboundedPrecedingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUnboundedFollowing(ctx *UnboundedFollowingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBoundedPreceding(ctx *BoundedPrecedingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBoundedFollowing(ctx *BoundedFollowingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCurrentRow(ctx *CurrentRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitScalarFunction(ctx *ScalarFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTableFunction(ctx *TableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSpecialRegister(ctx *SpecialRegisterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiAnalogyFunction(ctx *AiAnalogyFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiFunctionExpression(ctx *AiFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiAnalogyFunctionSource(ctx *AiAnalogyFunctionSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiAnalogyFunctionTarget(ctx *AiAnalogyFunctionTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiAnalogyFunctionSource1(ctx *AiAnalogyFunctionSource1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiAnalogyFunctionSource2(ctx *AiAnalogyFunctionSource2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiAnalogyFunctionTarget1(ctx *AiAnalogyFunctionTarget1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiAnalogyFunctionTarget2(ctx *AiAnalogyFunctionTarget2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiSemanticClusterFunction(ctx *AiSemanticClusterFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiSemanticClusterMemberExpression(ctx *AiSemanticClusterMemberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiSemanticClusterClusteringExpression(ctx *AiSemanticClusterClusteringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiSimilarityFunction(ctx *AiSimilarityFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiSimilarityExpression(ctx *AiSimilarityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiSimilarityExpression1(ctx *AiSimilarityExpression1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAiSimilarityExpression2(ctx *AiSimilarityExpression2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlelementFunction(ctx *XmlelementFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlforestFunction(ctx *XmlforestFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlmodifyFunction(ctx *XmlmodifyFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlpiFunction(ctx *XmlpiFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlqueryFunction(ctx *XmlqueryFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlattributesFunction(ctx *XmlattributesFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlserializeFunction(ctx *XmlserializeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlnamespaceFunction(ctx *XmlnamespaceFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlnamespaceOption(ctx *XmlnamespaceOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlserializeFunctionOptions(ctx *XmlserializeFunctionOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlFunctionOptionClause(ctx *XmlFunctionOptionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlFunctionOption(ctx *XmlFunctionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitElementContentExpression(ctx *ElementContentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXqueryExpressionConstant(ctx *XqueryExpressionConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXqueryArgument(ctx *XqueryArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmltableFunctionSpecification(ctx *XmltableFunctionSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRowXqueryExpressionConstant(ctx *RowXqueryExpressionConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRowXqueryArgument(ctx *RowXqueryArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXqueryContextItemExpression(ctx *XqueryContextItemExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXqueryVariableExpression(ctx *XqueryVariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlTableRegularColumnDefinition(ctx *XmlTableRegularColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDefaultClause(ctx *DefaultClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDefaultClause1(ctx *DefaultClause1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDefaultClauseAllowables(ctx *DefaultClauseAllowablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDistinctTypeCastFunctionName(ctx *DistinctTypeCastFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitColumnXqueryExpressionConstant(ctx *ColumnXqueryExpressionConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlTableOrdinalityColumnDefinition(ctx *XmlTableOrdinalityColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlnamespacesDeclaration(ctx *XmlnamespacesDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlnamespacesFunctionSpecification(ctx *XmlnamespacesFunctionSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlnamespacesFunctionArguments(ctx *XmlnamespacesFunctionArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNamespaceUri(ctx *NamespaceUriContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNamespacePrefix(ctx *NamespacePrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTimeZoneSpecificExpression(ctx *TimeZoneSpecificExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTimeZoneExpressionSubset(ctx *TimeZoneExpressionSubsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCaseExpression(ctx *CaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitResultExpression(ctx *ResultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSearchedWhenClause(ctx *SearchedWhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSimpleWhenClause(ctx *SimpleWhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSearchCondition(ctx *SearchConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCheckCondition(ctx *CheckConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBasicPredicate(ctx *BasicPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRowValueExpression(ctx *RowValueExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitQuantifiedPredicate(ctx *QuantifiedPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayExistsPredicate(ctx *ArrayExistsPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDistinctPredicate(ctx *DistinctPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitExistsPredicate(ctx *ExistsPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitInPredicate(ctx *InPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNullPredicate(ctx *NullPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmlExistsPredicate(ctx *XmlExistsPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayExpression(ctx *ArrayExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCastSpecification(ctx *CastSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterMarker(ctx *ParameterMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCastDataType(ctx *CastDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCastBuiltInType(ctx *CastBuiltInTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIntegerInParens(ctx *IntegerInParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLength(ctx *LengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCcsidQualifier(ctx *CcsidQualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitForDataQualifier(ctx *ForDataQualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDistinctTypeName(ctx *DistinctTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCcsidValue(ctx *CcsidValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSourceColumnName(ctx *SourceColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTargetColumnName(ctx *TargetColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBuildColumnName(ctx *BuildColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBeginColumnName(ctx *BeginColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitEndColumnName(ctx *EndColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCorrelationName(ctx *CorrelationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLocationName(ctx *LocationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSchemaName(ctx *SchemaNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAlterTableName(ctx *AlterTableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAuxTableName(ctx *AuxTableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitHistoryTableName(ctx *HistoryTableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCloneTableName(ctx *CloneTableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArchiveTableName(ctx *ArchiveTableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitViewName(ctx *ViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitProgramName(ctx *ProgramNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPackageName(ctx *PackageNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPlanName(ctx *PlanNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitVariableName(ctx *VariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitArrayTypeName(ctx *ArrayTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitJarName(ctx *JarNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSavepointName(ctx *SavepointNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAliasName(ctx *AliasNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitConstraintName(ctx *ConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRoutineVersionID(ctx *RoutineVersionIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitVersionID(ctx *VersionIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMaskName(ctx *MaskNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPermissionName(ctx *PermissionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitProcedureName(ctx *ProcedureNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSequenceName(ctx *SequenceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMemberName(ctx *MemberNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDatabaseName(ctx *DatabaseNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTablespaceName(ctx *TablespaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAcceleratorName(ctx *AcceleratorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCatalogName(ctx *CatalogNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTriggerName(ctx *TriggerNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitContextName(ctx *ContextNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAuthorizationName(ctx *AuthorizationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitProfileName(ctx *ProfileNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitRoleName(ctx *RoleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSeclabelName(ctx *SeclabelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitParameterName(ctx *ParameterNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAddressValue(ctx *AddressValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitJobnameValue(ctx *JobnameValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitServauthValue(ctx *ServauthValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitEncryptionValue(ctx *EncryptionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitBpName(ctx *BpNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitStogroupName(ctx *StogroupNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDcName(ctx *DcNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitMcName(ctx *McNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitScName(ctx *ScNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitVolumeID(ctx *VolumeIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitKeyLabelName(ctx *KeyLabelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSpecificName(ctx *SpecificNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitHostLabel(ctx *HostLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitHostVariable(ctx *HostVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitHostIdentifier(ctx *HostIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitHostStructure(ctx *HostStructureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNullIndicator(ctx *NullIndicatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNullIndicatorStructure(ctx *NullIndicatorStructureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGlobalVariableName(ctx *GlobalVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlParameterName(ctx *SqlParameterNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlVariableName(ctx *SqlVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTransitionVariableName(ctx *TransitionVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSynonym(ctx *SynonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIntoClause(ctx *IntoClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCorrelationClause(ctx *CorrelationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTableReference(ctx *TableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSingleTableReference(ctx *SingleTableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPeriodSpecification(ctx *PeriodSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitPeriodClause(ctx *PeriodClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitNestedTableExpression(ctx *NestedTableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitDataChangeTableReference(ctx *DataChangeTableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTableFunctionReference(ctx *TableFunctionReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTableUdfCardinalityClause(ctx *TableUdfCardinalityClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTypedCorrelationClause(ctx *TypedCorrelationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitTableLocatorReference(ctx *TableLocatorReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitXmltableExpression(ctx *XmltableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCollectionDerivedTable(ctx *CollectionDerivedTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitJoinCondition(ctx *JoinConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFullJoinExpression(ctx *FullJoinExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCastFunction(ctx *CastFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOrdinaryArrayExpression(ctx *OrdinaryArrayExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitAssociativeArrayExpression(ctx *AssociativeArrayExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGroupingExpression(ctx *GroupingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGroupingSets(ctx *GroupingSetsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitGroupingSetsGroup(ctx *GroupingSetsGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSuperGroups(ctx *SuperGroupsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSelectColumns(ctx *SelectColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUnpackedRow(ctx *UnpackedRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSubSelect(ctx *SubSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSelectIntoStatement(ctx *SelectIntoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitUpdateClause(ctx *UpdateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitReadOnlyClause(ctx *ReadOnlyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOptimizeClause(ctx *OptimizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIsolationClause(ctx *IsolationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitLockClause(ctx *LockClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSkipLockedDataClause(ctx *SkipLockedDataClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitQuerynoClause(ctx *QuerynoClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitScalarFullSelect(ctx *ScalarFullSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFullSelect(ctx *FullSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitValuesClause(ctx *ValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSortKey(ctx *SortKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitOffsetClause(ctx *OffsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitFetchClause(ctx *FetchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitIdentifier1(ctx *Identifier1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlidentifier(ctx *SqlidentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDB2zSQLParserVisitor) VisitSqlKeyword(ctx *SqlKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}
