// Code generated from DB2zSQLParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package db2 // DB2zSQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by DB2zSQLParser.
type DB2zSQLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DB2zSQLParser#startRule.
	VisitStartRule(ctx *StartRuleContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlStatement.
	VisitSqlStatement(ctx *SqlStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#cursorName.
	VisitCursorName(ctx *CursorNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#statementName.
	VisitStatementName(ctx *StatementNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#descriptorName.
	VisitDescriptorName(ctx *DescriptorNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#holdability.
	VisitHoldability(ctx *HoldabilityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#returnability.
	VisitReturnability(ctx *ReturnabilityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rowsetPositioning.
	VisitRowsetPositioning(ctx *RowsetPositioningContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#notNullPhrase.
	VisitNotNullPhrase(ctx *NotNullPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#allocateCursorStatement.
	VisitAllocateCursorStatement(ctx *AllocateCursorStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rsLocatorVariable.
	VisitRsLocatorVariable(ctx *RsLocatorVariableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterDatabaseStatement.
	VisitAlterDatabaseStatement(ctx *AlterDatabaseStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterFunctionStatement.
	VisitAlterFunctionStatement(ctx *AlterFunctionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterFunctionStatementExternal.
	VisitAlterFunctionStatementExternal(ctx *AlterFunctionStatementExternalContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterFunctionStatementCompiledSqlScalar.
	VisitAlterFunctionStatementCompiledSqlScalar(ctx *AlterFunctionStatementCompiledSqlScalarContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterWhichFunction1.
	VisitAlterWhichFunction1(ctx *AlterWhichFunction1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterWhichFunction2.
	VisitAlterWhichFunction2(ctx *AlterWhichFunction2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionCompiledSqlScalarRoutineSpecification.
	VisitFunctionCompiledSqlScalarRoutineSpecification(ctx *FunctionCompiledSqlScalarRoutineSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterFunctionStatementInlineSqlScalar.
	VisitAlterFunctionStatementInlineSqlScalar(ctx *AlterFunctionStatementInlineSqlScalarContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterFunctionStatementSqlTable.
	VisitAlterFunctionStatementSqlTable(ctx *AlterFunctionStatementSqlTableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionReturnsClause.
	VisitFunctionReturnsClause(ctx *FunctionReturnsClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionDesignator1.
	VisitFunctionDesignator1(ctx *FunctionDesignator1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionDesignator2.
	VisitFunctionDesignator2(ctx *FunctionDesignator2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterIndexStatement.
	VisitAlterIndexStatement(ctx *AlterIndexStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterMaskStatement.
	VisitAlterMaskStatement(ctx *AlterMaskStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterPermissionStatement.
	VisitAlterPermissionStatement(ctx *AlterPermissionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterProcedureStatement.
	VisitAlterProcedureStatement(ctx *AlterProcedureStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterProcedureSQLPLStatement.
	VisitAlterProcedureSQLPLStatement(ctx *AlterProcedureSQLPLStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterWhichProcedureSQLPL1.
	VisitAlterWhichProcedureSQLPL1(ctx *AlterWhichProcedureSQLPL1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterWhichProcedureSQLPL2.
	VisitAlterWhichProcedureSQLPL2(ctx *AlterWhichProcedureSQLPL2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#applicationCompatibilityPhrase.
	VisitApplicationCompatibilityPhrase(ctx *ApplicationCompatibilityPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterSequenceStatement.
	VisitAlterSequenceStatement(ctx *AlterSequenceStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterStogroupStatement.
	VisitAlterStogroupStatement(ctx *AlterStogroupStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTableStatement.
	VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTablespaceStatement.
	VisitAlterTablespaceStatement(ctx *AlterTablespaceStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTriggerStatement.
	VisitAlterTriggerStatement(ctx *AlterTriggerStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTrustedContextStatement.
	VisitAlterTrustedContextStatement(ctx *AlterTrustedContextStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterViewStatement.
	VisitAlterViewStatement(ctx *AlterViewStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#associateLocatorsStatement.
	VisitAssociateLocatorsStatement(ctx *AssociateLocatorsStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#beginDeclareSectionStatement.
	VisitBeginDeclareSectionStatement(ctx *BeginDeclareSectionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#callStatement.
	VisitCallStatement(ctx *CallStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#closeStatement.
	VisitCloseStatement(ctx *CloseStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#commentStatement.
	VisitCommentStatement(ctx *CommentStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#commitStatement.
	VisitCommitStatement(ctx *CommitStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#connectStatement.
	VisitConnectStatement(ctx *ConnectStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createAliasStatement.
	VisitCreateAliasStatement(ctx *CreateAliasStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createAuxiliaryTableStatement.
	VisitCreateAuxiliaryTableStatement(ctx *CreateAuxiliaryTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createDatabaseStatement.
	VisitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatement.
	VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementExternalScalar.
	VisitCreateFunctionStatementExternalScalar(ctx *CreateFunctionStatementExternalScalarContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementExternalScalarReturnsPhrase.
	VisitCreateFunctionStatementExternalScalarReturnsPhrase(ctx *CreateFunctionStatementExternalScalarReturnsPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementExternalTable.
	VisitCreateFunctionStatementExternalTable(ctx *CreateFunctionStatementExternalTableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementExternalTableReturnsPhrase.
	VisitCreateFunctionStatementExternalTableReturnsPhrase(ctx *CreateFunctionStatementExternalTableReturnsPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#externalTableFunctionColumn.
	VisitExternalTableFunctionColumn(ctx *ExternalTableFunctionColumnContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementSourced.
	VisitCreateFunctionStatementSourced(ctx *CreateFunctionStatementSourcedContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementSourcedReturnsPhrase.
	VisitCreateFunctionStatementSourcedReturnsPhrase(ctx *CreateFunctionStatementSourcedReturnsPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementSourcedSourcePhrase.
	VisitCreateFunctionStatementSourcedSourcePhrase(ctx *CreateFunctionStatementSourcedSourcePhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementInlineSqlScalar.
	VisitCreateFunctionStatementInlineSqlScalar(ctx *CreateFunctionStatementInlineSqlScalarContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementCompiledSqlScalar.
	VisitCreateFunctionStatementCompiledSqlScalar(ctx *CreateFunctionStatementCompiledSqlScalarContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementSqlTable.
	VisitCreateFunctionStatementSqlTable(ctx *CreateFunctionStatementSqlTableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createGlobalTemporaryTableStatement.
	VisitCreateGlobalTemporaryTableStatement(ctx *CreateGlobalTemporaryTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createIndexStatement.
	VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createLobTablespaceStatement.
	VisitCreateLobTablespaceStatement(ctx *CreateLobTablespaceStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createMaskStatement.
	VisitCreateMaskStatement(ctx *CreateMaskStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createPermissionStatement.
	VisitCreatePermissionStatement(ctx *CreatePermissionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createProcedureSQLPLStatement.
	VisitCreateProcedureSQLPLStatement(ctx *CreateProcedureSQLPLStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlRoutineBody.
	VisitSqlRoutineBody(ctx *SqlRoutineBodyContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#obfuscatedStatementText.
	VisitObfuscatedStatementText(ctx *ObfuscatedStatementTextContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#probablySQLPL.
	VisitProbablySQLPL(ctx *ProbablySQLPLContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createProcedureStatement.
	VisitCreateProcedureStatement(ctx *CreateProcedureStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createRoleStatement.
	VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createSequenceStatement.
	VisitCreateSequenceStatement(ctx *CreateSequenceStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createStogroupStatement.
	VisitCreateStogroupStatement(ctx *CreateStogroupStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTableStatement.
	VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTableOptions.
	VisitCreateTableOptions(ctx *CreateTableOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTablespaceStatement.
	VisitCreateTablespaceStatement(ctx *CreateTablespaceStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTriggerStatement.
	VisitCreateTriggerStatement(ctx *CreateTriggerStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTrustedContextStatement.
	VisitCreateTrustedContextStatement(ctx *CreateTrustedContextStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTypeArrayStatement.
	VisitCreateTypeArrayStatement(ctx *CreateTypeArrayStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTypeDistinctStatement.
	VisitCreateTypeDistinctStatement(ctx *CreateTypeDistinctStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createVariableStatement.
	VisitCreateVariableStatement(ctx *CreateVariableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createViewStatement.
	VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#declareCursorStatement.
	VisitDeclareCursorStatement(ctx *DeclareCursorStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#declareGlobalTemporaryTableStatement.
	VisitDeclareGlobalTemporaryTableStatement(ctx *DeclareGlobalTemporaryTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#declareTableStatement.
	VisitDeclareTableStatement(ctx *DeclareTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#declareStatementStatement.
	VisitDeclareStatementStatement(ctx *DeclareStatementStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#declareVariableStatement.
	VisitDeclareVariableStatement(ctx *DeclareVariableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#describeStatement.
	VisitDescribeStatement(ctx *DescribeStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#describeCursorStatement.
	VisitDescribeCursorStatement(ctx *DescribeCursorStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#describeInputStatement.
	VisitDescribeInputStatement(ctx *DescribeInputStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#describeOutputStatement.
	VisitDescribeOutputStatement(ctx *DescribeOutputStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#describeProcedureStatement.
	VisitDescribeProcedureStatement(ctx *DescribeProcedureStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#describeTableStatement.
	VisitDescribeTableStatement(ctx *DescribeTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropStatement.
	VisitDropStatement(ctx *DropStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#endDeclareSectionStatement.
	VisitEndDeclareSectionStatement(ctx *EndDeclareSectionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#exchangeStatement.
	VisitExchangeStatement(ctx *ExchangeStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#executeImmediateStatement.
	VisitExecuteImmediateStatement(ctx *ExecuteImmediateStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#explainStatement.
	VisitExplainStatement(ctx *ExplainStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#fetchStatement.
	VisitFetchStatement(ctx *FetchStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#freeLocatorStatement.
	VisitFreeLocatorStatement(ctx *FreeLocatorStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#getDiagnosticsStatement.
	VisitGetDiagnosticsStatement(ctx *GetDiagnosticsStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantStatement.
	VisitGrantStatement(ctx *GrantStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#holdLocatorStatement.
	VisitHoldLocatorStatement(ctx *HoldLocatorStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#includeStatement.
	VisitIncludeStatement(ctx *IncludeStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#labelStatement.
	VisitLabelStatement(ctx *LabelStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#lockTableStatement.
	VisitLockTableStatement(ctx *LockTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#mergeStatement.
	VisitMergeStatement(ctx *MergeStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#openStatement.
	VisitOpenStatement(ctx *OpenStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#prepareStatement.
	VisitPrepareStatement(ctx *PrepareStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#refreshTableStatement.
	VisitRefreshTableStatement(ctx *RefreshTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#releaseConnectionStatement.
	VisitReleaseConnectionStatement(ctx *ReleaseConnectionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#releaseSavepointStatement.
	VisitReleaseSavepointStatement(ctx *ReleaseSavepointStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#renameStatement.
	VisitRenameStatement(ctx *RenameStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeStatement.
	VisitRevokeStatement(ctx *RevokeStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rollbackStatement.
	VisitRollbackStatement(ctx *RollbackStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#savepointStatement.
	VisitSavepointStatement(ctx *SavepointStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#setAssignmentStatement.
	VisitSetAssignmentStatement(ctx *SetAssignmentStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#setConnectionStatement.
	VisitSetConnectionStatement(ctx *SetConnectionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#setEncryptionPasswordStatement.
	VisitSetEncryptionPasswordStatement(ctx *SetEncryptionPasswordStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#setPathStatement.
	VisitSetPathStatement(ctx *SetPathStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#setSchemaStatement.
	VisitSetSchemaStatement(ctx *SetSchemaStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#setSessionTimezoneStatement.
	VisitSetSessionTimezoneStatement(ctx *SetSessionTimezoneStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#setSpecialRegisterStatement.
	VisitSetSpecialRegisterStatement(ctx *SetSpecialRegisterStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#signalStatement.
	VisitSignalStatement(ctx *SignalStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#transferOwnershipStatement.
	VisitTransferOwnershipStatement(ctx *TransferOwnershipStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#truncateStatement.
	VisitTruncateStatement(ctx *TruncateStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesStatement.
	VisitValuesStatement(ctx *ValuesStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesIntoStatement.
	VisitValuesIntoStatement(ctx *ValuesIntoStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#wheneverStatement.
	VisitWheneverStatement(ctx *WheneverStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesIntoTargetVariable.
	VisitValuesIntoTargetVariable(ctx *ValuesIntoTargetVariableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#ownedObject.
	VisitOwnedObject(ctx *OwnedObjectContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#newOwner.
	VisitNewOwner(ctx *NewOwnerContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantCollectionStatement.
	VisitGrantCollectionStatement(ctx *GrantCollectionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantDatabaseStatement.
	VisitGrantDatabaseStatement(ctx *GrantDatabaseStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantFunctionOrProcedureStatement.
	VisitGrantFunctionOrProcedureStatement(ctx *GrantFunctionOrProcedureStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantPackageStatement.
	VisitGrantPackageStatement(ctx *GrantPackageStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantPlanStatement.
	VisitGrantPlanStatement(ctx *GrantPlanStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantSchemaStatement.
	VisitGrantSchemaStatement(ctx *GrantSchemaStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantSequenceStatement.
	VisitGrantSequenceStatement(ctx *GrantSequenceStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantSystemStatement.
	VisitGrantSystemStatement(ctx *GrantSystemStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantTableStatement.
	VisitGrantTableStatement(ctx *GrantTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantTypeOrJarStatement.
	VisitGrantTypeOrJarStatement(ctx *GrantTypeOrJarStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantVariableStatement.
	VisitGrantVariableStatement(ctx *GrantVariableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantUseOfStatement.
	VisitGrantUseOfStatement(ctx *GrantUseOfStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeCollectionStatement.
	VisitRevokeCollectionStatement(ctx *RevokeCollectionStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeDatabaseStatement.
	VisitRevokeDatabaseStatement(ctx *RevokeDatabaseStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeFunctionOrProcedureStatement.
	VisitRevokeFunctionOrProcedureStatement(ctx *RevokeFunctionOrProcedureStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokePackageStatement.
	VisitRevokePackageStatement(ctx *RevokePackageStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokePlanStatement.
	VisitRevokePlanStatement(ctx *RevokePlanStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeSchemaStatement.
	VisitRevokeSchemaStatement(ctx *RevokeSchemaStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeSequenceStatement.
	VisitRevokeSequenceStatement(ctx *RevokeSequenceStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeSystemStatement.
	VisitRevokeSystemStatement(ctx *RevokeSystemStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeTableStatement.
	VisitRevokeTableStatement(ctx *RevokeTableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeTypeOrJarStatement.
	VisitRevokeTypeOrJarStatement(ctx *RevokeTypeOrJarStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeVariableStatement.
	VisitRevokeVariableStatement(ctx *RevokeVariableStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeUseOfStatement.
	VisitRevokeUseOfStatement(ctx *RevokeUseOfStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantUseOfTarget.
	VisitGrantUseOfTarget(ctx *GrantUseOfTargetContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantVariableAuthority.
	VisitGrantVariableAuthority(ctx *GrantVariableAuthorityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantTableAuthority.
	VisitGrantTableAuthority(ctx *GrantTableAuthorityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantSystemAuthority.
	VisitGrantSystemAuthority(ctx *GrantSystemAuthorityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantSequenceAuthority.
	VisitGrantSequenceAuthority(ctx *GrantSequenceAuthorityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantSchemaAuthority.
	VisitGrantSchemaAuthority(ctx *GrantSchemaAuthorityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantPlanAuthority.
	VisitGrantPlanAuthority(ctx *GrantPlanAuthorityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantPackageAuthority.
	VisitGrantPackageAuthority(ctx *GrantPackageAuthorityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#packageSpecification.
	VisitPackageSpecification(ctx *PackageSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionSpecification.
	VisitFunctionSpecification(ctx *FunctionSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantee.
	VisitGrantee(ctx *GranteeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#withGrantOption.
	VisitWithGrantOption(ctx *WithGrantOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeByOption.
	VisitRevokeByOption(ctx *RevokeByOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#revokeDependentPrivilegesOption.
	VisitRevokeDependentPrivilegesOption(ctx *RevokeDependentPrivilegesOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#grantDatabaseAuthority.
	VisitGrantDatabaseAuthority(ctx *GrantDatabaseAuthorityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#statementInformation.
	VisitStatementInformation(ctx *StatementInformationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#statementInformationVariableEquate.
	VisitStatementInformationVariableEquate(ctx *StatementInformationVariableEquateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#statementInformationItemName.
	VisitStatementInformationItemName(ctx *StatementInformationItemNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#conditionInformation.
	VisitConditionInformation(ctx *ConditionInformationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#conditionInformationVariableEquate.
	VisitConditionInformationVariableEquate(ctx *ConditionInformationVariableEquateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#conditionInformationItemName.
	VisitConditionInformationItemName(ctx *ConditionInformationItemNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#connectionInformationItemName.
	VisitConnectionInformationItemName(ctx *ConnectionInformationItemNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#combinedInformation.
	VisitCombinedInformation(ctx *CombinedInformationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#combinedInformationOption.
	VisitCombinedInformationOption(ctx *CombinedInformationOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#fetchOrientation.
	VisitFetchOrientation(ctx *FetchOrientationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rowPositioned.
	VisitRowPositioned(ctx *RowPositionedContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rowsetPositioned.
	VisitRowsetPositioned(ctx *RowsetPositionedContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#singleRowFetch.
	VisitSingleRowFetch(ctx *SingleRowFetchContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#fetchTargetVariable.
	VisitFetchTargetVariable(ctx *FetchTargetVariableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#multipleRowFetch.
	VisitMultipleRowFetch(ctx *MultipleRowFetchContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#multipleRowFetchForClause.
	VisitMultipleRowFetchForClause(ctx *MultipleRowFetchForClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#multipleRowFetchIntoClause.
	VisitMultipleRowFetchIntoClause(ctx *MultipleRowFetchIntoClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#explainPlanClause.
	VisitExplainPlanClause(ctx *ExplainPlanClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#explainStmtcacheClause.
	VisitExplainStmtcacheClause(ctx *ExplainStmtcacheClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#explainPackageClause.
	VisitExplainPackageClause(ctx *ExplainPackageClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#explainStabilizedDynamicQueryClause.
	VisitExplainStabilizedDynamicQueryClause(ctx *ExplainStabilizedDynamicQueryClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#packageScopeSpecification.
	VisitPackageScopeSpecification(ctx *PackageScopeSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#collectionName.
	VisitCollectionName(ctx *CollectionNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#packageScopePackageName.
	VisitPackageScopePackageName(ctx *PackageScopePackageNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#versionName.
	VisitVersionName(ctx *VersionNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sourceRowData.
	VisitSourceRowData(ctx *SourceRowDataContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aliasDesignation.
	VisitAliasDesignation(ctx *AliasDesignationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropDatabaseClause.
	VisitDropDatabaseClause(ctx *DropDatabaseClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropFunctionClause.
	VisitDropFunctionClause(ctx *DropFunctionClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropIndexClause.
	VisitDropIndexClause(ctx *DropIndexClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropMaskClause.
	VisitDropMaskClause(ctx *DropMaskClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropPackageClause.
	VisitDropPackageClause(ctx *DropPackageClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropPermissionClause.
	VisitDropPermissionClause(ctx *DropPermissionClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropProcedureClause.
	VisitDropProcedureClause(ctx *DropProcedureClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropRoleClause.
	VisitDropRoleClause(ctx *DropRoleClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropSequenceClause.
	VisitDropSequenceClause(ctx *DropSequenceClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropStogroupClause.
	VisitDropStogroupClause(ctx *DropStogroupClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropSynonymClause.
	VisitDropSynonymClause(ctx *DropSynonymClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropTableClause.
	VisitDropTableClause(ctx *DropTableClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropTablespaceClause.
	VisitDropTablespaceClause(ctx *DropTablespaceClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropTriggerClause.
	VisitDropTriggerClause(ctx *DropTriggerClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropTrustedContextClause.
	VisitDropTrustedContextClause(ctx *DropTrustedContextClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropTypeClause.
	VisitDropTypeClause(ctx *DropTypeClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropVariableClause.
	VisitDropVariableClause(ctx *DropVariableClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropViewClause.
	VisitDropViewClause(ctx *DropViewClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#packageDesignator.
	VisitPackageDesignator(ctx *PackageDesignatorContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#describeUsingOption.
	VisitDescribeUsingOption(ctx *DescribeUsingOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#declareGlobalTemporaryTableLikeClause.
	VisitDeclareGlobalTemporaryTableLikeClause(ctx *DeclareGlobalTemporaryTableLikeClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#onCommitClause.
	VisitOnCommitClause(ctx *OnCommitClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#loggedWithRollbackClause.
	VisitLoggedWithRollbackClause(ctx *LoggedWithRollbackClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createViewCheckOptionClause.
	VisitCreateViewCheckOptionClause(ctx *CreateViewCheckOptionClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trustedContextDefaultRoleClause.
	VisitTrustedContextDefaultRoleClause(ctx *TrustedContextDefaultRoleClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trustedContextEnableDisableClause.
	VisitTrustedContextEnableDisableClause(ctx *TrustedContextEnableDisableClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trustedContextDefaultSecurityLabelClause.
	VisitTrustedContextDefaultSecurityLabelClause(ctx *TrustedContextDefaultSecurityLabelClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trustedContextAttributesClause.
	VisitTrustedContextAttributesClause(ctx *TrustedContextAttributesClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trustedContextWithUseForClause.
	VisitTrustedContextWithUseForClause(ctx *TrustedContextWithUseForClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trustedContextAttribute1.
	VisitTrustedContextAttribute1(ctx *TrustedContextAttribute1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trustedContextAttribute2.
	VisitTrustedContextAttribute2(ctx *TrustedContextAttribute2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trustedContextUseFor.
	VisitTrustedContextUseFor(ctx *TrustedContextUseForContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#userOptions.
	VisitUserOptions(ctx *UserOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#triggerDefinition.
	VisitTriggerDefinition(ctx *TriggerDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#triggerActivationTime.
	VisitTriggerActivationTime(ctx *TriggerActivationTimeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#triggerEvent.
	VisitTriggerEvent(ctx *TriggerEventContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#triggerGranularity.
	VisitTriggerGranularity(ctx *TriggerGranularityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#triggeredAction.
	VisitTriggeredAction(ctx *TriggeredActionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlTriggerBody.
	VisitSqlTriggerBody(ctx *SqlTriggerBodyContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#triggeredSqlStatement.
	VisitTriggeredSqlStatement(ctx *TriggeredSqlStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#triggerDefinitionOption.
	VisitTriggerDefinitionOption(ctx *TriggerDefinitionOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTableInClause.
	VisitCreateTableInClause(ctx *CreateTableInClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#customVolatileClause.
	VisitCustomVolatileClause(ctx *CustomVolatileClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTableColumnDefinition.
	VisitCreateTableColumnDefinition(ctx *CreateTableColumnDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#editprocClause.
	VisitEditprocClause(ctx *EditprocClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#validprocClause.
	VisitValidprocClause(ctx *ValidprocClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#auditClause.
	VisitAuditClause(ctx *AuditClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#obidClause.
	VisitObidClause(ctx *ObidClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dataCaptureClause.
	VisitDataCaptureClause(ctx *DataCaptureClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#restrictOnDropClause.
	VisitRestrictOnDropClause(ctx *RestrictOnDropClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#ccsidClause1.
	VisitCcsidClause1(ctx *CcsidClause1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#ccsidClause2.
	VisitCcsidClause2(ctx *CcsidClause2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#cardinalityClause.
	VisitCardinalityClause(ctx *CardinalityClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#appendClause.
	VisitAppendClause(ctx *AppendClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#memberClause.
	VisitMemberClause(ctx *MemberClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trackmodClause.
	VisitTrackmodClause(ctx *TrackmodClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#pagenumClause.
	VisitPagenumClause(ctx *PagenumClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#fieldprocClause.
	VisitFieldprocClause(ctx *FieldprocClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asSecurityLabelClause.
	VisitAsSecurityLabelClause(ctx *AsSecurityLabelClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#implicitlyHiddenClause.
	VisitImplicitlyHiddenClause(ctx *ImplicitlyHiddenClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#inlineLengthClause.
	VisitInlineLengthClause(ctx *InlineLengthClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#copyOptions.
	VisitCopyOptions(ctx *CopyOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#copyOptionIdentity.
	VisitCopyOptionIdentity(ctx *CopyOptionIdentityContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#copyOptionRowChangeTimestamp.
	VisitCopyOptionRowChangeTimestamp(ctx *CopyOptionRowChangeTimestampContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#copyOptionColumnDefaults.
	VisitCopyOptionColumnDefaults(ctx *CopyOptionColumnDefaultsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#copyOptionXmlTypeModifiers.
	VisitCopyOptionXmlTypeModifiers(ctx *CopyOptionXmlTypeModifiersContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asResultTable.
	VisitAsResultTable(ctx *AsResultTableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#declareGlobalTemporaryTableAsResultTable.
	VisitDeclareGlobalTemporaryTableAsResultTable(ctx *DeclareGlobalTemporaryTableAsResultTableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTableMaterializedQueryDefinition.
	VisitCreateTableMaterializedQueryDefinition(ctx *CreateTableMaterializedQueryDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTableColumnConstraint.
	VisitCreateTableColumnConstraint(ctx *CreateTableColumnConstraintContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#organizationClause.
	VisitOrganizationClause(ctx *OrganizationClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createGlobalTemporaryTableColumnDefinition.
	VisitCreateGlobalTemporaryTableColumnDefinition(ctx *CreateGlobalTemporaryTableColumnDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#declareGlobalTemporaryTableColumnDefinition.
	VisitDeclareGlobalTemporaryTableColumnDefinition(ctx *DeclareGlobalTemporaryTableColumnDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterDeclaration1.
	VisitParameterDeclaration1(ctx *ParameterDeclaration1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterDeclaration2.
	VisitParameterDeclaration2(ctx *ParameterDeclaration2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterDeclaration3.
	VisitParameterDeclaration3(ctx *ParameterDeclaration3Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementExternalScalarOptions.
	VisitCreateFunctionStatementExternalScalarOptions(ctx *CreateFunctionStatementExternalScalarOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#externalNameOption1.
	VisitExternalNameOption1(ctx *ExternalNameOption1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#externalNameOption2.
	VisitExternalNameOption2(ctx *ExternalNameOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dynamicResultSetOption.
	VisitDynamicResultSetOption(ctx *DynamicResultSetOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#languageOption1.
	VisitLanguageOption1(ctx *LanguageOption1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#languageOption2.
	VisitLanguageOption2(ctx *LanguageOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#languageOption3.
	VisitLanguageOption3(ctx *LanguageOption3Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#languageOption4.
	VisitLanguageOption4(ctx *LanguageOption4Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#languageOption5.
	VisitLanguageOption5(ctx *LanguageOption5Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterStyleOption1.
	VisitParameterStyleOption1(ctx *ParameterStyleOption1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterStyleOption2.
	VisitParameterStyleOption2(ctx *ParameterStyleOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterStyleOption3.
	VisitParameterStyleOption3(ctx *ParameterStyleOption3Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#deterministicOption.
	VisitDeterministicOption(ctx *DeterministicOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#fencedOption.
	VisitFencedOption(ctx *FencedOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nullInputOption1.
	VisitNullInputOption1(ctx *NullInputOption1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nullInputOption2.
	VisitNullInputOption2(ctx *NullInputOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nullInputOption3.
	VisitNullInputOption3(ctx *NullInputOption3Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nullInputOption4.
	VisitNullInputOption4(ctx *NullInputOption4Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#debugOption.
	VisitDebugOption(ctx *DebugOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlDataOption1.
	VisitSqlDataOption1(ctx *SqlDataOption1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlDataOption2.
	VisitSqlDataOption2(ctx *SqlDataOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlDataOption3.
	VisitSqlDataOption3(ctx *SqlDataOption3Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlDataOption4.
	VisitSqlDataOption4(ctx *SqlDataOption4Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#externalActionOption.
	VisitExternalActionOption(ctx *ExternalActionOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#packagePathOption.
	VisitPackagePathOption(ctx *PackagePathOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#scratchpadOption.
	VisitScratchpadOption(ctx *ScratchpadOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#finalCallOption.
	VisitFinalCallOption(ctx *FinalCallOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parallelOption1.
	VisitParallelOption1(ctx *ParallelOption1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parallelOption2.
	VisitParallelOption2(ctx *ParallelOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dbinfoOption.
	VisitDbinfoOption(ctx *DbinfoOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#cardinalityOption.
	VisitCardinalityOption(ctx *CardinalityOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#collectionIdOption.
	VisitCollectionIdOption(ctx *CollectionIdOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#wlmEnvironmentOption1.
	VisitWlmEnvironmentOption1(ctx *WlmEnvironmentOption1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#wlmEnvironmentOption2.
	VisitWlmEnvironmentOption2(ctx *WlmEnvironmentOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#wlmEnvironmentOption3.
	VisitWlmEnvironmentOption3(ctx *WlmEnvironmentOption3Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asuTimeOption.
	VisitAsuTimeOption(ctx *AsuTimeOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#stayResidentOption.
	VisitStayResidentOption(ctx *StayResidentOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#programTypeOption.
	VisitProgramTypeOption(ctx *ProgramTypeOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#securityOption.
	VisitSecurityOption(ctx *SecurityOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#stopAfterFailureOption.
	VisitStopAfterFailureOption(ctx *StopAfterFailureOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#runOptionsOption.
	VisitRunOptionsOption(ctx *RunOptionsOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#commitOnReturnOption.
	VisitCommitOnReturnOption(ctx *CommitOnReturnOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#specialRegistersOption.
	VisitSpecialRegistersOption(ctx *SpecialRegistersOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#specialRegistersOption2.
	VisitSpecialRegistersOption2(ctx *SpecialRegistersOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dispatchOption.
	VisitDispatchOption(ctx *DispatchOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#securedOption.
	VisitSecuredOption(ctx *SecuredOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#specificNameOption1.
	VisitSpecificNameOption1(ctx *SpecificNameOption1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#specificNameOption2.
	VisitSpecificNameOption2(ctx *SpecificNameOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterOption1.
	VisitParameterOption1(ctx *ParameterOption1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterOption2.
	VisitParameterOption2(ctx *ParameterOption2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementExternalTableOptions.
	VisitCreateFunctionStatementExternalTableOptions(ctx *CreateFunctionStatementExternalTableOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementSourcedOptions.
	VisitCreateFunctionStatementSourcedOptions(ctx *CreateFunctionStatementSourcedOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#inlineSqlScalarFunctionDefinition.
	VisitInlineSqlScalarFunctionDefinition(ctx *InlineSqlScalarFunctionDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementInlineSqlScalarOptions.
	VisitCreateFunctionStatementInlineSqlScalarOptions(ctx *CreateFunctionStatementInlineSqlScalarOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#compiledSqlScalarFunctionDefinition.
	VisitCompiledSqlScalarFunctionDefinition(ctx *CompiledSqlScalarFunctionDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementCompiledSqlScalarOptions.
	VisitCreateFunctionStatementCompiledSqlScalarOptions(ctx *CreateFunctionStatementCompiledSqlScalarOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlTableFunctionDefinition.
	VisitSqlTableFunctionDefinition(ctx *SqlTableFunctionDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createFunctionStatementSqlTableOptions.
	VisitCreateFunctionStatementSqlTableOptions(ctx *CreateFunctionStatementSqlTableOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sequenceAlias.
	VisitSequenceAlias(ctx *SequenceAliasContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#tableAlias.
	VisitTableAlias(ctx *TableAliasContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#authorization.
	VisitAuthorization(ctx *AuthorizationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#searchedDelete.
	VisitSearchedDelete(ctx *SearchedDeleteContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#positionedDelete.
	VisitPositionedDelete(ctx *PositionedDeleteContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#searchedUpdate.
	VisitSearchedUpdate(ctx *SearchedUpdateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#positionedUpdate.
	VisitPositionedUpdate(ctx *PositionedUpdateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sourceValues.
	VisitSourceValues(ctx *SourceValuesContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesSingleRow.
	VisitValuesSingleRow(ctx *ValuesSingleRowContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesMultipleRow.
	VisitValuesMultipleRow(ctx *ValuesMultipleRowContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#matchingCondition.
	VisitMatchingCondition(ctx *MatchingConditionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#modificationOperation.
	VisitModificationOperation(ctx *ModificationOperationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#assignmentClause.
	VisitAssignmentClause(ctx *AssignmentClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#setAssignmentClause.
	VisitSetAssignmentClause(ctx *SetAssignmentClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#setAssignmentTargetVariable.
	VisitSetAssignmentTargetVariable(ctx *SetAssignmentTargetVariableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#updateOperation.
	VisitUpdateOperation(ctx *UpdateOperationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#deleteOperation.
	VisitDeleteOperation(ctx *DeleteOperationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#insertOperation.
	VisitInsertOperation(ctx *InsertOperationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#signalInformation.
	VisitSignalInformation(ctx *SignalInformationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesList1.
	VisitValuesList1(ctx *ValuesList1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesList2.
	VisitValuesList2(ctx *ValuesList2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesList3.
	VisitValuesList3(ctx *ValuesList3Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesList4.
	VisitValuesList4(ctx *ValuesList4Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#includeColumns.
	VisitIncludeColumns(ctx *IncludeColumnsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#multipleRowInsert.
	VisitMultipleRowInsert(ctx *MultipleRowInsertContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#regenerateClause.
	VisitRegenerateClause(ctx *RegenerateClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterIndexOptions.
	VisitAlterIndexOptions(ctx *AlterIndexOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#bufferpoolOption.
	VisitBufferpoolOption(ctx *BufferpoolOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#closeOption.
	VisitCloseOption(ctx *CloseOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#copyOption.
	VisitCopyOption(ctx *CopyOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dssizeOption.
	VisitDssizeOption(ctx *DssizeOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#piecesizeOption.
	VisitPiecesizeOption(ctx *PiecesizeOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#clusterOption.
	VisitClusterOption(ctx *ClusterOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#paddedOption.
	VisitPaddedOption(ctx *PaddedOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#compressOption.
	VisitCompressOption(ctx *CompressOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#defineOption.
	VisitDefineOption(ctx *DefineOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#locksizeOption.
	VisitLocksizeOption(ctx *LocksizeOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#lockmaxOption.
	VisitLockmaxOption(ctx *LockmaxOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#enableDisableOption.
	VisitEnableDisableOption(ctx *EnableDisableOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#loggedOption.
	VisitLoggedOption(ctx *LoggedOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#notAtomicPhrase.
	VisitNotAtomicPhrase(ctx *NotAtomicPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterIndexPartitionOptions.
	VisitAlterIndexPartitionOptions(ctx *AlterIndexPartitionOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#usingSpecification1.
	VisitUsingSpecification1(ctx *UsingSpecification1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#freeSpecification.
	VisitFreeSpecification(ctx *FreeSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#gbpcacheSpecification.
	VisitGbpcacheSpecification(ctx *GbpcacheSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitionElement.
	VisitPartitionElement(ctx *PartitionElementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#applCompatValue.
	VisitApplCompatValue(ctx *ApplCompatValueContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionLevel.
	VisitFunctionLevel(ctx *FunctionLevelContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionParameterType.
	VisitFunctionParameterType(ctx *FunctionParameterTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionDataType.
	VisitFunctionDataType(ctx *FunctionDataTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionBuiltInType.
	VisitFunctionBuiltInType(ctx *FunctionBuiltInTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#procedureBuiltinType.
	VisitProcedureBuiltinType(ctx *ProcedureBuiltinTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTypeArrayBuiltinType.
	VisitCreateTypeArrayBuiltinType(ctx *CreateTypeArrayBuiltinTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTypeArrayBuiltinType2.
	VisitCreateTypeArrayBuiltinType2(ctx *CreateTypeArrayBuiltinType2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createVariableBuiltInType.
	VisitCreateVariableBuiltInType(ctx *CreateVariableBuiltInTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sourceDataType.
	VisitSourceDataType(ctx *SourceDataTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionExternalOptionList.
	VisitFunctionExternalOptionList(ctx *FunctionExternalOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionCompiledSqlScalarOptionList.
	VisitFunctionCompiledSqlScalarOptionList(ctx *FunctionCompiledSqlScalarOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionInlineSqlScalarOptionList.
	VisitFunctionInlineSqlScalarOptionList(ctx *FunctionInlineSqlScalarOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionSqlTableOptionList.
	VisitFunctionSqlTableOptionList(ctx *FunctionSqlTableOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#procedureOptionList.
	VisitProcedureOptionList(ctx *ProcedureOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createProcedureOptionList.
	VisitCreateProcedureOptionList(ctx *CreateProcedureOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#procedureSQLPLOptionList.
	VisitProcedureSQLPLOptionList(ctx *ProcedureSQLPLOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#versionOption.
	VisitVersionOption(ctx *VersionOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#commitOnReturnOptionSQLPL.
	VisitCommitOnReturnOptionSQLPL(ctx *CommitOnReturnOptionSQLPLContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#schemaQualifierOption.
	VisitSchemaQualifierOption(ctx *SchemaQualifierOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#currentDataOption.
	VisitCurrentDataOption(ctx *CurrentDataOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#immediateWriteOption.
	VisitImmediateWriteOption(ctx *ImmediateWriteOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#explainOption.
	VisitExplainOption(ctx *ExplainOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#reoptOption.
	VisitReoptOption(ctx *ReoptOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#packageOwnerOption.
	VisitPackageOwnerOption(ctx *PackageOwnerOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#deferPrepareOption.
	VisitDeferPrepareOption(ctx *DeferPrepareOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#degreeOption.
	VisitDegreeOption(ctx *DegreeOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dynamicRulesOption.
	VisitDynamicRulesOption(ctx *DynamicRulesOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#concurrentAccessOption.
	VisitConcurrentAccessOption(ctx *ConcurrentAccessOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#applicationEncodingOption.
	VisitApplicationEncodingOption(ctx *ApplicationEncodingOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#isolationLevelOption.
	VisitIsolationLevelOption(ctx *IsolationLevelOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#keepDynamicOption.
	VisitKeepDynamicOption(ctx *KeepDynamicOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#opthintOption.
	VisitOpthintOption(ctx *OpthintOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlPathOption.
	VisitSqlPathOption(ctx *SqlPathOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlPathOptionItem.
	VisitSqlPathOptionItem(ctx *SqlPathOptionItemContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#queryAccelerationOption.
	VisitQueryAccelerationOption(ctx *QueryAccelerationOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#queryAccelerationOptionItem.
	VisitQueryAccelerationOptionItem(ctx *QueryAccelerationOptionItemContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#getAccelArchiveOption.
	VisitGetAccelArchiveOption(ctx *GetAccelArchiveOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#accelerationOption.
	VisitAccelerationOption(ctx *AccelerationOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#releaseAtOption.
	VisitReleaseAtOption(ctx *ReleaseAtOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#businessTimeSensitiveOption.
	VisitBusinessTimeSensitiveOption(ctx *BusinessTimeSensitiveOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#systemTimeSensitiveOption.
	VisitSystemTimeSensitiveOption(ctx *SystemTimeSensitiveOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#archiveSensitiveOption.
	VisitArchiveSensitiveOption(ctx *ArchiveSensitiveOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#applcompatOption.
	VisitApplcompatOption(ctx *ApplcompatOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#validateOption.
	VisitValidateOption(ctx *ValidateOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#roundingOption.
	VisitRoundingOption(ctx *RoundingOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#roundingOptionItem.
	VisitRoundingOptionItem(ctx *RoundingOptionItemContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dateFormatOption.
	VisitDateFormatOption(ctx *DateFormatOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dateTimeFormatOptionItem.
	VisitDateTimeFormatOptionItem(ctx *DateTimeFormatOptionItemContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#timeFormatOption.
	VisitTimeFormatOption(ctx *TimeFormatOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#decimalOption.
	VisitDecimalOption(ctx *DecimalOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#forUpdateOption.
	VisitForUpdateOption(ctx *ForUpdateOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#concentrateStatementsOption.
	VisitConcentrateStatementsOption(ctx *ConcentrateStatementsOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#acceleratorOption.
	VisitAcceleratorOption(ctx *AcceleratorOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#procedureDataType.
	VisitProcedureDataType(ctx *ProcedureDataTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterSequenceOptionList.
	VisitAlterSequenceOptionList(ctx *AlterSequenceOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createSequenceOptionList.
	VisitCreateSequenceOptionList(ctx *CreateSequenceOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asTypeOption.
	VisitAsTypeOption(ctx *AsTypeOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#startOption.
	VisitStartOption(ctx *StartOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#restartOption.
	VisitRestartOption(ctx *RestartOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#incrementOption.
	VisitIncrementOption(ctx *IncrementOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#minvalueOption.
	VisitMinvalueOption(ctx *MinvalueOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#maxvalueOption.
	VisitMaxvalueOption(ctx *MaxvalueOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#cycleOption.
	VisitCycleOption(ctx *CycleOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#cacheOption.
	VisitCacheOption(ctx *CacheOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#orderOption.
	VisitOrderOption(ctx *OrderOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#keyLabelOption.
	VisitKeyLabelOption(ctx *KeyLabelOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dataclasOption.
	VisitDataclasOption(ctx *DataclasOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#mgmtclasOption.
	VisitMgmtclasOption(ctx *MgmtclasOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#storclasOption.
	VisitStorclasOption(ctx *StorclasOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterStogroupOptionList.
	VisitAlterStogroupOptionList(ctx *AlterStogroupOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTableOptionList.
	VisitAlterTableOptionList(ctx *AlterTableOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTablespaceOptionList.
	VisitAlterTablespaceOptionList(ctx *AlterTablespaceOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createTablespaceOptionList.
	VisitCreateTablespaceOptionList(ctx *CreateTablespaceOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#trustedContextOptionList.
	VisitTrustedContextOptionList(ctx *TrustedContextOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#databaseOptionList.
	VisitDatabaseOptionList(ctx *DatabaseOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createIndexOptionList.
	VisitCreateIndexOptionList(ctx *CreateIndexOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#createLobTablespaceOptionList.
	VisitCreateLobTablespaceOptionList(ctx *CreateLobTablespaceOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#inDatabaseOption.
	VisitInDatabaseOption(ctx *InDatabaseOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#segsizeOption.
	VisitSegsizeOption(ctx *SegsizeOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#numpartsOption.
	VisitNumpartsOption(ctx *NumpartsOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitionByGrowthSpecification.
	VisitPartitionByGrowthSpecification(ctx *PartitionByGrowthSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitionByRangeSpecification.
	VisitPartitionByRangeSpecification(ctx *PartitionByRangeSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitionByRangePartitionPhrase.
	VisitPartitionByRangePartitionPhrase(ctx *PartitionByRangePartitionPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#insertAlgorithmOption.
	VisitInsertAlgorithmOption(ctx *InsertAlgorithmOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#maxrowsOption.
	VisitMaxrowsOption(ctx *MaxrowsOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#maxpartitionsOption.
	VisitMaxpartitionsOption(ctx *MaxpartitionsOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#usingSpecification2.
	VisitUsingSpecification2(ctx *UsingSpecification2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#stogroupOptions.
	VisitStogroupOptions(ctx *StogroupOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlIndexSpecification.
	VisitXmlIndexSpecification(ctx *XmlIndexSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlPatternClause.
	VisitXmlPatternClause(ctx *XmlPatternClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterAttributesOptions.
	VisitAlterAttributesOptions(ctx *AlterAttributesOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#addAttributesOptions.
	VisitAddAttributesOptions(ctx *AddAttributesOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dropAttributesOptions.
	VisitDropAttributesOptions(ctx *DropAttributesOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#includeColumnPhrase.
	VisitIncludeColumnPhrase(ctx *IncludeColumnPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#userClause.
	VisitUserClause(ctx *UserClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#userClauseAddOptions.
	VisitUserClauseAddOptions(ctx *UserClauseAddOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#userClauseReplaceOptions.
	VisitUserClauseReplaceOptions(ctx *UserClauseReplaceOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#userClauseDropOptions.
	VisitUserClauseDropOptions(ctx *UserClauseDropOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#useOptions.
	VisitUseOptions(ctx *UseOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterPartitionClause.
	VisitAlterPartitionClause(ctx *AlterPartitionClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#usingBlock.
	VisitUsingBlock(ctx *UsingBlockContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#freeBlock.
	VisitFreeBlock(ctx *FreeBlockContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#moveTableClause.
	VisitMoveTableClause(ctx *MoveTableClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#gbpcacheBlock.
	VisitGbpcacheBlock(ctx *GbpcacheBlockContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aliasDesignator.
	VisitAliasDesignator(ctx *AliasDesignatorContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#multipleColumnList.
	VisitMultipleColumnList(ctx *MultipleColumnListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionDesignator.
	VisitFunctionDesignator(ctx *FunctionDesignatorContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterType.
	VisitParameterType(ctx *ParameterTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTableColumnDefinitionOptionList1.
	VisitAlterTableColumnDefinitionOptionList1(ctx *AlterTableColumnDefinitionOptionList1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTableColumnDefinitionOptionList2.
	VisitAlterTableColumnDefinitionOptionList2(ctx *AlterTableColumnDefinitionOptionList2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#columnConstraint.
	VisitColumnConstraint(ctx *ColumnConstraintContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#generatedClause.
	VisitGeneratedClause(ctx *GeneratedClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#generatedClause2.
	VisitGeneratedClause2(ctx *GeneratedClause2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asIdentityClause.
	VisitAsIdentityClause(ctx *AsIdentityClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asIdentityClauseOptionList.
	VisitAsIdentityClauseOptionList(ctx *AsIdentityClauseOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asRowChangeTimestampClause.
	VisitAsRowChangeTimestampClause(ctx *AsRowChangeTimestampClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asRowTransactionStartIDClause.
	VisitAsRowTransactionStartIDClause(ctx *AsRowTransactionStartIDClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asRowTransactionTimestampClause.
	VisitAsRowTransactionTimestampClause(ctx *AsRowTransactionTimestampClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#asGeneratedExpressionClause.
	VisitAsGeneratedExpressionClause(ctx *AsGeneratedExpressionClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nonDeterministicExpression.
	VisitNonDeterministicExpression(ctx *NonDeterministicExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nonDeterministicExpressionSessionVariable.
	VisitNonDeterministicExpressionSessionVariable(ctx *NonDeterministicExpressionSessionVariableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#columnAlteration.
	VisitColumnAlteration(ctx *ColumnAlterationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#columnAlterationOptionList.
	VisitColumnAlterationOptionList(ctx *ColumnAlterationOptionListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alteredDataType.
	VisitAlteredDataType(ctx *AlteredDataTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#builtInType.
	VisitBuiltInType(ctx *BuiltInTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sequenceDataType.
	VisitSequenceDataType(ctx *SequenceDataTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sequenceBuiltInType.
	VisitSequenceBuiltInType(ctx *SequenceBuiltInTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlDataType.
	VisitSqlDataType(ctx *SqlDataTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlTypeModifier.
	VisitXmlTypeModifier(ctx *XmlTypeModifierContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlSchemaSpecification.
	VisitXmlSchemaSpecification(ctx *XmlSchemaSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlElementName.
	VisitXmlElementName(ctx *XmlElementNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#piName.
	VisitPiName(ctx *PiNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#registeredXmlSchemaName.
	VisitRegisteredXmlSchemaName(ctx *RegisteredXmlSchemaNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#targetNamespace.
	VisitTargetNamespace(ctx *TargetNamespaceContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#schemaLocation.
	VisitSchemaLocation(ctx *SchemaLocationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#identityAlteration.
	VisitIdentityAlteration(ctx *IdentityAlterationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#uniqueConstraint.
	VisitUniqueConstraint(ctx *UniqueConstraintContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#referentialConstraint.
	VisitReferentialConstraint(ctx *ReferentialConstraintContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#referencesClause.
	VisitReferencesClause(ctx *ReferencesClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#checkConstraint.
	VisitCheckConstraint(ctx *CheckConstraintContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitioningClause.
	VisitPartitioningClause(ctx *PartitioningClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitionExpression.
	VisitPartitionExpression(ctx *PartitionExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitionLimitKey.
	VisitPartitionLimitKey(ctx *PartitionLimitKeyContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitioningPhrase.
	VisitPartitioningPhrase(ctx *PartitioningPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitionHashSpace.
	VisitPartitionHashSpace(ctx *PartitionHashSpaceContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterHashOrganization.
	VisitAlterHashOrganization(ctx *AlterHashOrganizationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitioningClauseElement.
	VisitPartitioningClauseElement(ctx *PartitioningClauseElementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#partitionClause.
	VisitPartitionClause(ctx *PartitionClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rotatePartitionClause.
	VisitRotatePartitionClause(ctx *RotatePartitionClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#extraRowOption.
	VisitExtraRowOption(ctx *ExtraRowOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#materializedQueryDefinition.
	VisitMaterializedQueryDefinition(ctx *MaterializedQueryDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#materializedQueryAlteration.
	VisitMaterializedQueryAlteration(ctx *MaterializedQueryAlterationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#refreshableTableOptions.
	VisitRefreshableTableOptions(ctx *RefreshableTableOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dataInitiallyDeferredPhrase.
	VisitDataInitiallyDeferredPhrase(ctx *DataInitiallyDeferredPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#refreshDeferredPhrase.
	VisitRefreshDeferredPhrase(ctx *RefreshDeferredPhraseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#refreshableTableOptionsList.
	VisitRefreshableTableOptionsList(ctx *RefreshableTableOptionsListContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#materializedQueryTableAlteration.
	VisitMaterializedQueryTableAlteration(ctx *MaterializedQueryTableAlterationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#periodDefinition.
	VisitPeriodDefinition(ctx *PeriodDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTableColumnDefinition.
	VisitAlterTableColumnDefinition(ctx *AlterTableColumnDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#externalProgramName.
	VisitExternalProgramName(ctx *ExternalProgramNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#packagePath.
	VisitPackagePath(ctx *PackagePathContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#collectionID.
	VisitCollectionID(ctx *CollectionIDContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#runTimeOptions.
	VisitRunTimeOptions(ctx *RunTimeOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#keyExpression.
	VisitKeyExpression(ctx *KeyExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rowChangeExpression.
	VisitRowChangeExpression(ctx *RowChangeExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sequenceReference.
	VisitSequenceReference(ctx *SequenceReferenceContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionInvocation.
	VisitFunctionInvocation(ctx *FunctionInvocationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#scalarFunctionInvocation.
	VisitScalarFunctionInvocation(ctx *ScalarFunctionInvocationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aggregateFunctionInvocation.
	VisitAggregateFunctionInvocation(ctx *AggregateFunctionInvocationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#regressionFunctionInvocation.
	VisitRegressionFunctionInvocation(ctx *RegressionFunctionInvocationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#externalFunctionInvocation.
	VisitExternalFunctionInvocation(ctx *ExternalFunctionInvocationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#labeledDuration.
	VisitLabeledDuration(ctx *LabeledDurationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#durationSuffix.
	VisitDurationSuffix(ctx *DurationSuffixContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlCastSpecification.
	VisitXmlCastSpecification(ctx *XmlCastSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlParseSpecification.
	VisitXmlParseSpecification(ctx *XmlParseSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayElementSpecification.
	VisitArrayElementSpecification(ctx *ArrayElementSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayIndex.
	VisitArrayIndex(ctx *ArrayIndexContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayConstructor.
	VisitArrayConstructor(ctx *ArrayConstructorContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#olapSpecification.
	VisitOlapSpecification(ctx *OlapSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#orderedOlapSpecification.
	VisitOrderedOlapSpecification(ctx *OrderedOlapSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#olapSpecificationFunction.
	VisitOlapSpecificationFunction(ctx *OlapSpecificationFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#lagFunction.
	VisitLagFunction(ctx *LagFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#leadFunction.
	VisitLeadFunction(ctx *LeadFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#respectNullsClause.
	VisitRespectNullsClause(ctx *RespectNullsClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#windowPartitionClause.
	VisitWindowPartitionClause(ctx *WindowPartitionClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#windowOrderClause.
	VisitWindowOrderClause(ctx *WindowOrderClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#windowOrderClauseQualifier.
	VisitWindowOrderClauseQualifier(ctx *WindowOrderClauseQualifierContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#numberingSpecification.
	VisitNumberingSpecification(ctx *NumberingSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aggregationSpecification.
	VisitAggregationSpecification(ctx *AggregationSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aggregateFunction.
	VisitAggregateFunction(ctx *AggregateFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#regressionFunction.
	VisitRegressionFunction(ctx *RegressionFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#olapColumnFunction.
	VisitOlapColumnFunction(ctx *OlapColumnFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#firstValueFunction.
	VisitFirstValueFunction(ctx *FirstValueFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#lastValueFunction.
	VisitLastValueFunction(ctx *LastValueFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nthValueFunction.
	VisitNthValueFunction(ctx *NthValueFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#ratioToReportFunction.
	VisitRatioToReportFunction(ctx *RatioToReportFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#listaggFunction.
	VisitListaggFunction(ctx *ListaggFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayaggFunction.
	VisitArrayaggFunction(ctx *ArrayaggFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayaggOrdinaryFunction.
	VisitArrayaggOrdinaryFunction(ctx *ArrayaggOrdinaryFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayaggAssociativeFunction.
	VisitArrayaggAssociativeFunction(ctx *ArrayaggAssociativeFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#correlationFunction.
	VisitCorrelationFunction(ctx *CorrelationFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#covarianceFunction.
	VisitCovarianceFunction(ctx *CovarianceFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#covarianceSampFunction.
	VisitCovarianceSampFunction(ctx *CovarianceSampFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#cumeDistFunction.
	VisitCumeDistFunction(ctx *CumeDistFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#percentileContFunction.
	VisitPercentileContFunction(ctx *PercentileContFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#percentileDiscFunction.
	VisitPercentileDiscFunction(ctx *PercentileDiscFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#percentRankFunction.
	VisitPercentRankFunction(ctx *PercentRankFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlaggFunction.
	VisitXmlaggFunction(ctx *XmlaggFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlaggOrderByClause.
	VisitXmlaggOrderByClause(ctx *XmlaggOrderByClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlaggOrderByOption.
	VisitXmlaggOrderByOption(ctx *XmlaggOrderByOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aggregateOrderByClause.
	VisitAggregateOrderByClause(ctx *AggregateOrderByClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aggregateOrderByOption.
	VisitAggregateOrderByOption(ctx *AggregateOrderByOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#windowAggregationGroupClause.
	VisitWindowAggregationGroupClause(ctx *WindowAggregationGroupClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#groupStart.
	VisitGroupStart(ctx *GroupStartContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#groupBetween.
	VisitGroupBetween(ctx *GroupBetweenContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#groupEnd.
	VisitGroupEnd(ctx *GroupEndContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#groupBound1.
	VisitGroupBound1(ctx *GroupBound1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#groupBound2.
	VisitGroupBound2(ctx *GroupBound2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#unboundedPreceding.
	VisitUnboundedPreceding(ctx *UnboundedPrecedingContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#unboundedFollowing.
	VisitUnboundedFollowing(ctx *UnboundedFollowingContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#boundedPreceding.
	VisitBoundedPreceding(ctx *BoundedPrecedingContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#boundedFollowing.
	VisitBoundedFollowing(ctx *BoundedFollowingContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#currentRow.
	VisitCurrentRow(ctx *CurrentRowContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#scalarFunction.
	VisitScalarFunction(ctx *ScalarFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#tableFunction.
	VisitTableFunction(ctx *TableFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#specialRegister.
	VisitSpecialRegister(ctx *SpecialRegisterContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiAnalogyFunction.
	VisitAiAnalogyFunction(ctx *AiAnalogyFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiFunctionExpression.
	VisitAiFunctionExpression(ctx *AiFunctionExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiAnalogyFunctionSource.
	VisitAiAnalogyFunctionSource(ctx *AiAnalogyFunctionSourceContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiAnalogyFunctionTarget.
	VisitAiAnalogyFunctionTarget(ctx *AiAnalogyFunctionTargetContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiAnalogyFunctionSource1.
	VisitAiAnalogyFunctionSource1(ctx *AiAnalogyFunctionSource1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiAnalogyFunctionSource2.
	VisitAiAnalogyFunctionSource2(ctx *AiAnalogyFunctionSource2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiAnalogyFunctionTarget1.
	VisitAiAnalogyFunctionTarget1(ctx *AiAnalogyFunctionTarget1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiAnalogyFunctionTarget2.
	VisitAiAnalogyFunctionTarget2(ctx *AiAnalogyFunctionTarget2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiSemanticClusterFunction.
	VisitAiSemanticClusterFunction(ctx *AiSemanticClusterFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiSemanticClusterMemberExpression.
	VisitAiSemanticClusterMemberExpression(ctx *AiSemanticClusterMemberExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiSemanticClusterClusteringExpression.
	VisitAiSemanticClusterClusteringExpression(ctx *AiSemanticClusterClusteringExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiSimilarityFunction.
	VisitAiSimilarityFunction(ctx *AiSimilarityFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiSimilarityExpression.
	VisitAiSimilarityExpression(ctx *AiSimilarityExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiSimilarityExpression1.
	VisitAiSimilarityExpression1(ctx *AiSimilarityExpression1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aiSimilarityExpression2.
	VisitAiSimilarityExpression2(ctx *AiSimilarityExpression2Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlelementFunction.
	VisitXmlelementFunction(ctx *XmlelementFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlforestFunction.
	VisitXmlforestFunction(ctx *XmlforestFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlmodifyFunction.
	VisitXmlmodifyFunction(ctx *XmlmodifyFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlpiFunction.
	VisitXmlpiFunction(ctx *XmlpiFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlqueryFunction.
	VisitXmlqueryFunction(ctx *XmlqueryFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlattributesFunction.
	VisitXmlattributesFunction(ctx *XmlattributesFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlserializeFunction.
	VisitXmlserializeFunction(ctx *XmlserializeFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlnamespaceFunction.
	VisitXmlnamespaceFunction(ctx *XmlnamespaceFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlnamespaceOption.
	VisitXmlnamespaceOption(ctx *XmlnamespaceOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlserializeFunctionOptions.
	VisitXmlserializeFunctionOptions(ctx *XmlserializeFunctionOptionsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlFunctionOptionClause.
	VisitXmlFunctionOptionClause(ctx *XmlFunctionOptionClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlFunctionOption.
	VisitXmlFunctionOption(ctx *XmlFunctionOptionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#elementContentExpression.
	VisitElementContentExpression(ctx *ElementContentExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xqueryExpressionConstant.
	VisitXqueryExpressionConstant(ctx *XqueryExpressionConstantContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xqueryArgument.
	VisitXqueryArgument(ctx *XqueryArgumentContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmltableFunctionSpecification.
	VisitXmltableFunctionSpecification(ctx *XmltableFunctionSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rowXqueryExpressionConstant.
	VisitRowXqueryExpressionConstant(ctx *RowXqueryExpressionConstantContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rowXqueryArgument.
	VisitRowXqueryArgument(ctx *RowXqueryArgumentContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xqueryContextItemExpression.
	VisitXqueryContextItemExpression(ctx *XqueryContextItemExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xqueryVariableExpression.
	VisitXqueryVariableExpression(ctx *XqueryVariableExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlTableRegularColumnDefinition.
	VisitXmlTableRegularColumnDefinition(ctx *XmlTableRegularColumnDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#defaultClause.
	VisitDefaultClause(ctx *DefaultClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#defaultClause1.
	VisitDefaultClause1(ctx *DefaultClause1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#defaultClauseAllowables.
	VisitDefaultClauseAllowables(ctx *DefaultClauseAllowablesContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#distinctTypeCastFunctionName.
	VisitDistinctTypeCastFunctionName(ctx *DistinctTypeCastFunctionNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#columnXqueryExpressionConstant.
	VisitColumnXqueryExpressionConstant(ctx *ColumnXqueryExpressionConstantContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlTableOrdinalityColumnDefinition.
	VisitXmlTableOrdinalityColumnDefinition(ctx *XmlTableOrdinalityColumnDefinitionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlnamespacesDeclaration.
	VisitXmlnamespacesDeclaration(ctx *XmlnamespacesDeclarationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlnamespacesFunctionSpecification.
	VisitXmlnamespacesFunctionSpecification(ctx *XmlnamespacesFunctionSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlnamespacesFunctionArguments.
	VisitXmlnamespacesFunctionArguments(ctx *XmlnamespacesFunctionArgumentsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#namespaceUri.
	VisitNamespaceUri(ctx *NamespaceUriContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#namespacePrefix.
	VisitNamespacePrefix(ctx *NamespacePrefixContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#timeZoneSpecificExpression.
	VisitTimeZoneSpecificExpression(ctx *TimeZoneSpecificExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#timeZoneExpressionSubset.
	VisitTimeZoneExpressionSubset(ctx *TimeZoneExpressionSubsetContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#caseExpression.
	VisitCaseExpression(ctx *CaseExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#resultExpression.
	VisitResultExpression(ctx *ResultExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#searchedWhenClause.
	VisitSearchedWhenClause(ctx *SearchedWhenClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#simpleWhenClause.
	VisitSimpleWhenClause(ctx *SimpleWhenClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#searchCondition.
	VisitSearchCondition(ctx *SearchConditionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#checkCondition.
	VisitCheckCondition(ctx *CheckConditionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#basicPredicate.
	VisitBasicPredicate(ctx *BasicPredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#rowValueExpression.
	VisitRowValueExpression(ctx *RowValueExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#quantifiedPredicate.
	VisitQuantifiedPredicate(ctx *QuantifiedPredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayExistsPredicate.
	VisitArrayExistsPredicate(ctx *ArrayExistsPredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#betweenPredicate.
	VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#distinctPredicate.
	VisitDistinctPredicate(ctx *DistinctPredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#existsPredicate.
	VisitExistsPredicate(ctx *ExistsPredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#inPredicate.
	VisitInPredicate(ctx *InPredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nullPredicate.
	VisitNullPredicate(ctx *NullPredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmlExistsPredicate.
	VisitXmlExistsPredicate(ctx *XmlExistsPredicateContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayExpression.
	VisitArrayExpression(ctx *ArrayExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#castSpecification.
	VisitCastSpecification(ctx *CastSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterMarker.
	VisitParameterMarker(ctx *ParameterMarkerContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#castDataType.
	VisitCastDataType(ctx *CastDataTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#castBuiltInType.
	VisitCastBuiltInType(ctx *CastBuiltInTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#integerInParens.
	VisitIntegerInParens(ctx *IntegerInParensContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#length.
	VisitLength(ctx *LengthContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#ccsidQualifier.
	VisitCcsidQualifier(ctx *CcsidQualifierContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#forDataQualifier.
	VisitForDataQualifier(ctx *ForDataQualifierContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#distinctTypeName.
	VisitDistinctTypeName(ctx *DistinctTypeNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#ccsidValue.
	VisitCcsidValue(ctx *CcsidValueContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sourceColumnName.
	VisitSourceColumnName(ctx *SourceColumnNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#targetColumnName.
	VisitTargetColumnName(ctx *TargetColumnNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#buildColumnName.
	VisitBuildColumnName(ctx *BuildColumnNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#beginColumnName.
	VisitBeginColumnName(ctx *BeginColumnNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#endColumnName.
	VisitEndColumnName(ctx *EndColumnNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#correlationName.
	VisitCorrelationName(ctx *CorrelationNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#locationName.
	VisitLocationName(ctx *LocationNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#schemaName.
	VisitSchemaName(ctx *SchemaNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#alterTableName.
	VisitAlterTableName(ctx *AlterTableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#auxTableName.
	VisitAuxTableName(ctx *AuxTableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#historyTableName.
	VisitHistoryTableName(ctx *HistoryTableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#cloneTableName.
	VisitCloneTableName(ctx *CloneTableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#archiveTableName.
	VisitArchiveTableName(ctx *ArchiveTableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#viewName.
	VisitViewName(ctx *ViewNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#programName.
	VisitProgramName(ctx *ProgramNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#packageName.
	VisitPackageName(ctx *PackageNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#planName.
	VisitPlanName(ctx *PlanNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#variableName.
	VisitVariableName(ctx *VariableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#arrayTypeName.
	VisitArrayTypeName(ctx *ArrayTypeNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#jarName.
	VisitJarName(ctx *JarNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#savepointName.
	VisitSavepointName(ctx *SavepointNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#aliasName.
	VisitAliasName(ctx *AliasNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#constraintName.
	VisitConstraintName(ctx *ConstraintNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#routineVersionID.
	VisitRoutineVersionID(ctx *RoutineVersionIDContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#versionID.
	VisitVersionID(ctx *VersionIDContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#maskName.
	VisitMaskName(ctx *MaskNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#permissionName.
	VisitPermissionName(ctx *PermissionNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#procedureName.
	VisitProcedureName(ctx *ProcedureNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sequenceName.
	VisitSequenceName(ctx *SequenceNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#memberName.
	VisitMemberName(ctx *MemberNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#databaseName.
	VisitDatabaseName(ctx *DatabaseNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#tablespaceName.
	VisitTablespaceName(ctx *TablespaceNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#acceleratorName.
	VisitAcceleratorName(ctx *AcceleratorNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#catalogName.
	VisitCatalogName(ctx *CatalogNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#triggerName.
	VisitTriggerName(ctx *TriggerNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#contextName.
	VisitContextName(ctx *ContextNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#authorizationName.
	VisitAuthorizationName(ctx *AuthorizationNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#profileName.
	VisitProfileName(ctx *ProfileNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#roleName.
	VisitRoleName(ctx *RoleNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#seclabelName.
	VisitSeclabelName(ctx *SeclabelNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#parameterName.
	VisitParameterName(ctx *ParameterNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#addressValue.
	VisitAddressValue(ctx *AddressValueContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#jobnameValue.
	VisitJobnameValue(ctx *JobnameValueContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#servauthValue.
	VisitServauthValue(ctx *ServauthValueContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#encryptionValue.
	VisitEncryptionValue(ctx *EncryptionValueContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#bpName.
	VisitBpName(ctx *BpNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#stogroupName.
	VisitStogroupName(ctx *StogroupNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dcName.
	VisitDcName(ctx *DcNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#mcName.
	VisitMcName(ctx *McNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#scName.
	VisitScName(ctx *ScNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#volumeID.
	VisitVolumeID(ctx *VolumeIDContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#keyLabelName.
	VisitKeyLabelName(ctx *KeyLabelNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#specificName.
	VisitSpecificName(ctx *SpecificNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#hostLabel.
	VisitHostLabel(ctx *HostLabelContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#hostVariable.
	VisitHostVariable(ctx *HostVariableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#hostIdentifier.
	VisitHostIdentifier(ctx *HostIdentifierContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#hostStructure.
	VisitHostStructure(ctx *HostStructureContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nullIndicator.
	VisitNullIndicator(ctx *NullIndicatorContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nullIndicatorStructure.
	VisitNullIndicatorStructure(ctx *NullIndicatorStructureContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#globalVariableName.
	VisitGlobalVariableName(ctx *GlobalVariableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlParameterName.
	VisitSqlParameterName(ctx *SqlParameterNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlVariableName.
	VisitSqlVariableName(ctx *SqlVariableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#transitionVariableName.
	VisitTransitionVariableName(ctx *TransitionVariableNameContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#synonym.
	VisitSynonym(ctx *SynonymContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#intoClause.
	VisitIntoClause(ctx *IntoClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#correlationClause.
	VisitCorrelationClause(ctx *CorrelationClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#tableReference.
	VisitTableReference(ctx *TableReferenceContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#singleTableReference.
	VisitSingleTableReference(ctx *SingleTableReferenceContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#periodSpecification.
	VisitPeriodSpecification(ctx *PeriodSpecificationContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#periodClause.
	VisitPeriodClause(ctx *PeriodClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#nestedTableExpression.
	VisitNestedTableExpression(ctx *NestedTableExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#dataChangeTableReference.
	VisitDataChangeTableReference(ctx *DataChangeTableReferenceContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#tableFunctionReference.
	VisitTableFunctionReference(ctx *TableFunctionReferenceContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#tableUdfCardinalityClause.
	VisitTableUdfCardinalityClause(ctx *TableUdfCardinalityClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#typedCorrelationClause.
	VisitTypedCorrelationClause(ctx *TypedCorrelationClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#tableLocatorReference.
	VisitTableLocatorReference(ctx *TableLocatorReferenceContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#xmltableExpression.
	VisitXmltableExpression(ctx *XmltableExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#collectionDerivedTable.
	VisitCollectionDerivedTable(ctx *CollectionDerivedTableContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#joinCondition.
	VisitJoinCondition(ctx *JoinConditionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#fullJoinExpression.
	VisitFullJoinExpression(ctx *FullJoinExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#castFunction.
	VisitCastFunction(ctx *CastFunctionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#ordinaryArrayExpression.
	VisitOrdinaryArrayExpression(ctx *OrdinaryArrayExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#associativeArrayExpression.
	VisitAssociativeArrayExpression(ctx *AssociativeArrayExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#groupingExpression.
	VisitGroupingExpression(ctx *GroupingExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#groupingSets.
	VisitGroupingSets(ctx *GroupingSetsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#groupingSetsGroup.
	VisitGroupingSetsGroup(ctx *GroupingSetsGroupContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#superGroups.
	VisitSuperGroups(ctx *SuperGroupsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#selectColumns.
	VisitSelectColumns(ctx *SelectColumnsContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#unpackedRow.
	VisitUnpackedRow(ctx *UnpackedRowContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#subSelect.
	VisitSubSelect(ctx *SubSelectContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#selectIntoStatement.
	VisitSelectIntoStatement(ctx *SelectIntoStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#commonTableExpression.
	VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#updateClause.
	VisitUpdateClause(ctx *UpdateClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#readOnlyClause.
	VisitReadOnlyClause(ctx *ReadOnlyClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#optimizeClause.
	VisitOptimizeClause(ctx *OptimizeClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#isolationClause.
	VisitIsolationClause(ctx *IsolationClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#lockClause.
	VisitLockClause(ctx *LockClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#skipLockedDataClause.
	VisitSkipLockedDataClause(ctx *SkipLockedDataClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#querynoClause.
	VisitQuerynoClause(ctx *QuerynoClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#scalarFullSelect.
	VisitScalarFullSelect(ctx *ScalarFullSelectContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#fullSelect.
	VisitFullSelect(ctx *FullSelectContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#valuesClause.
	VisitValuesClause(ctx *ValuesClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sortKey.
	VisitSortKey(ctx *SortKeyContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#offsetClause.
	VisitOffsetClause(ctx *OffsetClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#fetchClause.
	VisitFetchClause(ctx *FetchClauseContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#identifier1.
	VisitIdentifier1(ctx *Identifier1Context) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlidentifier.
	VisitSqlidentifier(ctx *SqlidentifierContext) interface{}

	// Visit a parse tree produced by DB2zSQLParser#sqlKeyword.
	VisitSqlKeyword(ctx *SqlKeywordContext) interface{}
}
