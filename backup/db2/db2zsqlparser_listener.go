// Code generated from DB2zSQLParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package db2 // DB2zSQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// DB2zSQLParserListener is a complete listener for a parse tree produced by DB2zSQLParser.
type DB2zSQLParserListener interface {
	antlr.ParseTreeListener

	// EnterStartRule is called when entering the startRule production.
	EnterStartRule(c *StartRuleContext)

	// EnterSqlStatement is called when entering the sqlStatement production.
	EnterSqlStatement(c *SqlStatementContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterCursorName is called when entering the cursorName production.
	EnterCursorName(c *CursorNameContext)

	// EnterStatementName is called when entering the statementName production.
	EnterStatementName(c *StatementNameContext)

	// EnterDescriptorName is called when entering the descriptorName production.
	EnterDescriptorName(c *DescriptorNameContext)

	// EnterHoldability is called when entering the holdability production.
	EnterHoldability(c *HoldabilityContext)

	// EnterReturnability is called when entering the returnability production.
	EnterReturnability(c *ReturnabilityContext)

	// EnterRowsetPositioning is called when entering the rowsetPositioning production.
	EnterRowsetPositioning(c *RowsetPositioningContext)

	// EnterNotNullPhrase is called when entering the notNullPhrase production.
	EnterNotNullPhrase(c *NotNullPhraseContext)

	// EnterAllocateCursorStatement is called when entering the allocateCursorStatement production.
	EnterAllocateCursorStatement(c *AllocateCursorStatementContext)

	// EnterRsLocatorVariable is called when entering the rsLocatorVariable production.
	EnterRsLocatorVariable(c *RsLocatorVariableContext)

	// EnterAlterDatabaseStatement is called when entering the alterDatabaseStatement production.
	EnterAlterDatabaseStatement(c *AlterDatabaseStatementContext)

	// EnterAlterFunctionStatement is called when entering the alterFunctionStatement production.
	EnterAlterFunctionStatement(c *AlterFunctionStatementContext)

	// EnterAlterFunctionStatementExternal is called when entering the alterFunctionStatementExternal production.
	EnterAlterFunctionStatementExternal(c *AlterFunctionStatementExternalContext)

	// EnterAlterFunctionStatementCompiledSqlScalar is called when entering the alterFunctionStatementCompiledSqlScalar production.
	EnterAlterFunctionStatementCompiledSqlScalar(c *AlterFunctionStatementCompiledSqlScalarContext)

	// EnterAlterWhichFunction1 is called when entering the alterWhichFunction1 production.
	EnterAlterWhichFunction1(c *AlterWhichFunction1Context)

	// EnterAlterWhichFunction2 is called when entering the alterWhichFunction2 production.
	EnterAlterWhichFunction2(c *AlterWhichFunction2Context)

	// EnterFunctionCompiledSqlScalarRoutineSpecification is called when entering the functionCompiledSqlScalarRoutineSpecification production.
	EnterFunctionCompiledSqlScalarRoutineSpecification(c *FunctionCompiledSqlScalarRoutineSpecificationContext)

	// EnterAlterFunctionStatementInlineSqlScalar is called when entering the alterFunctionStatementInlineSqlScalar production.
	EnterAlterFunctionStatementInlineSqlScalar(c *AlterFunctionStatementInlineSqlScalarContext)

	// EnterAlterFunctionStatementSqlTable is called when entering the alterFunctionStatementSqlTable production.
	EnterAlterFunctionStatementSqlTable(c *AlterFunctionStatementSqlTableContext)

	// EnterFunctionReturnsClause is called when entering the functionReturnsClause production.
	EnterFunctionReturnsClause(c *FunctionReturnsClauseContext)

	// EnterFunctionDesignator1 is called when entering the functionDesignator1 production.
	EnterFunctionDesignator1(c *FunctionDesignator1Context)

	// EnterFunctionDesignator2 is called when entering the functionDesignator2 production.
	EnterFunctionDesignator2(c *FunctionDesignator2Context)

	// EnterAlterIndexStatement is called when entering the alterIndexStatement production.
	EnterAlterIndexStatement(c *AlterIndexStatementContext)

	// EnterAlterMaskStatement is called when entering the alterMaskStatement production.
	EnterAlterMaskStatement(c *AlterMaskStatementContext)

	// EnterAlterPermissionStatement is called when entering the alterPermissionStatement production.
	EnterAlterPermissionStatement(c *AlterPermissionStatementContext)

	// EnterAlterProcedureStatement is called when entering the alterProcedureStatement production.
	EnterAlterProcedureStatement(c *AlterProcedureStatementContext)

	// EnterAlterProcedureSQLPLStatement is called when entering the alterProcedureSQLPLStatement production.
	EnterAlterProcedureSQLPLStatement(c *AlterProcedureSQLPLStatementContext)

	// EnterAlterWhichProcedureSQLPL1 is called when entering the alterWhichProcedureSQLPL1 production.
	EnterAlterWhichProcedureSQLPL1(c *AlterWhichProcedureSQLPL1Context)

	// EnterAlterWhichProcedureSQLPL2 is called when entering the alterWhichProcedureSQLPL2 production.
	EnterAlterWhichProcedureSQLPL2(c *AlterWhichProcedureSQLPL2Context)

	// EnterApplicationCompatibilityPhrase is called when entering the applicationCompatibilityPhrase production.
	EnterApplicationCompatibilityPhrase(c *ApplicationCompatibilityPhraseContext)

	// EnterAlterSequenceStatement is called when entering the alterSequenceStatement production.
	EnterAlterSequenceStatement(c *AlterSequenceStatementContext)

	// EnterAlterStogroupStatement is called when entering the alterStogroupStatement production.
	EnterAlterStogroupStatement(c *AlterStogroupStatementContext)

	// EnterAlterTableStatement is called when entering the alterTableStatement production.
	EnterAlterTableStatement(c *AlterTableStatementContext)

	// EnterAlterTablespaceStatement is called when entering the alterTablespaceStatement production.
	EnterAlterTablespaceStatement(c *AlterTablespaceStatementContext)

	// EnterAlterTriggerStatement is called when entering the alterTriggerStatement production.
	EnterAlterTriggerStatement(c *AlterTriggerStatementContext)

	// EnterAlterTrustedContextStatement is called when entering the alterTrustedContextStatement production.
	EnterAlterTrustedContextStatement(c *AlterTrustedContextStatementContext)

	// EnterAlterViewStatement is called when entering the alterViewStatement production.
	EnterAlterViewStatement(c *AlterViewStatementContext)

	// EnterAssociateLocatorsStatement is called when entering the associateLocatorsStatement production.
	EnterAssociateLocatorsStatement(c *AssociateLocatorsStatementContext)

	// EnterBeginDeclareSectionStatement is called when entering the beginDeclareSectionStatement production.
	EnterBeginDeclareSectionStatement(c *BeginDeclareSectionStatementContext)

	// EnterCallStatement is called when entering the callStatement production.
	EnterCallStatement(c *CallStatementContext)

	// EnterCloseStatement is called when entering the closeStatement production.
	EnterCloseStatement(c *CloseStatementContext)

	// EnterCommentStatement is called when entering the commentStatement production.
	EnterCommentStatement(c *CommentStatementContext)

	// EnterCommitStatement is called when entering the commitStatement production.
	EnterCommitStatement(c *CommitStatementContext)

	// EnterConnectStatement is called when entering the connectStatement production.
	EnterConnectStatement(c *ConnectStatementContext)

	// EnterCreateAliasStatement is called when entering the createAliasStatement production.
	EnterCreateAliasStatement(c *CreateAliasStatementContext)

	// EnterCreateAuxiliaryTableStatement is called when entering the createAuxiliaryTableStatement production.
	EnterCreateAuxiliaryTableStatement(c *CreateAuxiliaryTableStatementContext)

	// EnterCreateDatabaseStatement is called when entering the createDatabaseStatement production.
	EnterCreateDatabaseStatement(c *CreateDatabaseStatementContext)

	// EnterCreateFunctionStatement is called when entering the createFunctionStatement production.
	EnterCreateFunctionStatement(c *CreateFunctionStatementContext)

	// EnterCreateFunctionStatementExternalScalar is called when entering the createFunctionStatementExternalScalar production.
	EnterCreateFunctionStatementExternalScalar(c *CreateFunctionStatementExternalScalarContext)

	// EnterCreateFunctionStatementExternalScalarReturnsPhrase is called when entering the createFunctionStatementExternalScalarReturnsPhrase production.
	EnterCreateFunctionStatementExternalScalarReturnsPhrase(c *CreateFunctionStatementExternalScalarReturnsPhraseContext)

	// EnterCreateFunctionStatementExternalTable is called when entering the createFunctionStatementExternalTable production.
	EnterCreateFunctionStatementExternalTable(c *CreateFunctionStatementExternalTableContext)

	// EnterCreateFunctionStatementExternalTableReturnsPhrase is called when entering the createFunctionStatementExternalTableReturnsPhrase production.
	EnterCreateFunctionStatementExternalTableReturnsPhrase(c *CreateFunctionStatementExternalTableReturnsPhraseContext)

	// EnterExternalTableFunctionColumn is called when entering the externalTableFunctionColumn production.
	EnterExternalTableFunctionColumn(c *ExternalTableFunctionColumnContext)

	// EnterCreateFunctionStatementSourced is called when entering the createFunctionStatementSourced production.
	EnterCreateFunctionStatementSourced(c *CreateFunctionStatementSourcedContext)

	// EnterCreateFunctionStatementSourcedReturnsPhrase is called when entering the createFunctionStatementSourcedReturnsPhrase production.
	EnterCreateFunctionStatementSourcedReturnsPhrase(c *CreateFunctionStatementSourcedReturnsPhraseContext)

	// EnterCreateFunctionStatementSourcedSourcePhrase is called when entering the createFunctionStatementSourcedSourcePhrase production.
	EnterCreateFunctionStatementSourcedSourcePhrase(c *CreateFunctionStatementSourcedSourcePhraseContext)

	// EnterCreateFunctionStatementInlineSqlScalar is called when entering the createFunctionStatementInlineSqlScalar production.
	EnterCreateFunctionStatementInlineSqlScalar(c *CreateFunctionStatementInlineSqlScalarContext)

	// EnterCreateFunctionStatementCompiledSqlScalar is called when entering the createFunctionStatementCompiledSqlScalar production.
	EnterCreateFunctionStatementCompiledSqlScalar(c *CreateFunctionStatementCompiledSqlScalarContext)

	// EnterCreateFunctionStatementSqlTable is called when entering the createFunctionStatementSqlTable production.
	EnterCreateFunctionStatementSqlTable(c *CreateFunctionStatementSqlTableContext)

	// EnterCreateGlobalTemporaryTableStatement is called when entering the createGlobalTemporaryTableStatement production.
	EnterCreateGlobalTemporaryTableStatement(c *CreateGlobalTemporaryTableStatementContext)

	// EnterCreateIndexStatement is called when entering the createIndexStatement production.
	EnterCreateIndexStatement(c *CreateIndexStatementContext)

	// EnterCreateLobTablespaceStatement is called when entering the createLobTablespaceStatement production.
	EnterCreateLobTablespaceStatement(c *CreateLobTablespaceStatementContext)

	// EnterCreateMaskStatement is called when entering the createMaskStatement production.
	EnterCreateMaskStatement(c *CreateMaskStatementContext)

	// EnterCreatePermissionStatement is called when entering the createPermissionStatement production.
	EnterCreatePermissionStatement(c *CreatePermissionStatementContext)

	// EnterCreateProcedureSQLPLStatement is called when entering the createProcedureSQLPLStatement production.
	EnterCreateProcedureSQLPLStatement(c *CreateProcedureSQLPLStatementContext)

	// EnterSqlRoutineBody is called when entering the sqlRoutineBody production.
	EnterSqlRoutineBody(c *SqlRoutineBodyContext)

	// EnterObfuscatedStatementText is called when entering the obfuscatedStatementText production.
	EnterObfuscatedStatementText(c *ObfuscatedStatementTextContext)

	// EnterProbablySQLPL is called when entering the probablySQLPL production.
	EnterProbablySQLPL(c *ProbablySQLPLContext)

	// EnterCreateProcedureStatement is called when entering the createProcedureStatement production.
	EnterCreateProcedureStatement(c *CreateProcedureStatementContext)

	// EnterCreateRoleStatement is called when entering the createRoleStatement production.
	EnterCreateRoleStatement(c *CreateRoleStatementContext)

	// EnterCreateSequenceStatement is called when entering the createSequenceStatement production.
	EnterCreateSequenceStatement(c *CreateSequenceStatementContext)

	// EnterCreateStogroupStatement is called when entering the createStogroupStatement production.
	EnterCreateStogroupStatement(c *CreateStogroupStatementContext)

	// EnterCreateTableStatement is called when entering the createTableStatement production.
	EnterCreateTableStatement(c *CreateTableStatementContext)

	// EnterCreateTableOptions is called when entering the createTableOptions production.
	EnterCreateTableOptions(c *CreateTableOptionsContext)

	// EnterCreateTablespaceStatement is called when entering the createTablespaceStatement production.
	EnterCreateTablespaceStatement(c *CreateTablespaceStatementContext)

	// EnterCreateTriggerStatement is called when entering the createTriggerStatement production.
	EnterCreateTriggerStatement(c *CreateTriggerStatementContext)

	// EnterCreateTrustedContextStatement is called when entering the createTrustedContextStatement production.
	EnterCreateTrustedContextStatement(c *CreateTrustedContextStatementContext)

	// EnterCreateTypeArrayStatement is called when entering the createTypeArrayStatement production.
	EnterCreateTypeArrayStatement(c *CreateTypeArrayStatementContext)

	// EnterCreateTypeDistinctStatement is called when entering the createTypeDistinctStatement production.
	EnterCreateTypeDistinctStatement(c *CreateTypeDistinctStatementContext)

	// EnterCreateVariableStatement is called when entering the createVariableStatement production.
	EnterCreateVariableStatement(c *CreateVariableStatementContext)

	// EnterCreateViewStatement is called when entering the createViewStatement production.
	EnterCreateViewStatement(c *CreateViewStatementContext)

	// EnterDeclareCursorStatement is called when entering the declareCursorStatement production.
	EnterDeclareCursorStatement(c *DeclareCursorStatementContext)

	// EnterDeclareGlobalTemporaryTableStatement is called when entering the declareGlobalTemporaryTableStatement production.
	EnterDeclareGlobalTemporaryTableStatement(c *DeclareGlobalTemporaryTableStatementContext)

	// EnterDeclareTableStatement is called when entering the declareTableStatement production.
	EnterDeclareTableStatement(c *DeclareTableStatementContext)

	// EnterDeclareStatementStatement is called when entering the declareStatementStatement production.
	EnterDeclareStatementStatement(c *DeclareStatementStatementContext)

	// EnterDeclareVariableStatement is called when entering the declareVariableStatement production.
	EnterDeclareVariableStatement(c *DeclareVariableStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterDescribeStatement is called when entering the describeStatement production.
	EnterDescribeStatement(c *DescribeStatementContext)

	// EnterDescribeCursorStatement is called when entering the describeCursorStatement production.
	EnterDescribeCursorStatement(c *DescribeCursorStatementContext)

	// EnterDescribeInputStatement is called when entering the describeInputStatement production.
	EnterDescribeInputStatement(c *DescribeInputStatementContext)

	// EnterDescribeOutputStatement is called when entering the describeOutputStatement production.
	EnterDescribeOutputStatement(c *DescribeOutputStatementContext)

	// EnterDescribeProcedureStatement is called when entering the describeProcedureStatement production.
	EnterDescribeProcedureStatement(c *DescribeProcedureStatementContext)

	// EnterDescribeTableStatement is called when entering the describeTableStatement production.
	EnterDescribeTableStatement(c *DescribeTableStatementContext)

	// EnterDropStatement is called when entering the dropStatement production.
	EnterDropStatement(c *DropStatementContext)

	// EnterEndDeclareSectionStatement is called when entering the endDeclareSectionStatement production.
	EnterEndDeclareSectionStatement(c *EndDeclareSectionStatementContext)

	// EnterExchangeStatement is called when entering the exchangeStatement production.
	EnterExchangeStatement(c *ExchangeStatementContext)

	// EnterExecuteStatement is called when entering the executeStatement production.
	EnterExecuteStatement(c *ExecuteStatementContext)

	// EnterExecuteImmediateStatement is called when entering the executeImmediateStatement production.
	EnterExecuteImmediateStatement(c *ExecuteImmediateStatementContext)

	// EnterExplainStatement is called when entering the explainStatement production.
	EnterExplainStatement(c *ExplainStatementContext)

	// EnterFetchStatement is called when entering the fetchStatement production.
	EnterFetchStatement(c *FetchStatementContext)

	// EnterFreeLocatorStatement is called when entering the freeLocatorStatement production.
	EnterFreeLocatorStatement(c *FreeLocatorStatementContext)

	// EnterGetDiagnosticsStatement is called when entering the getDiagnosticsStatement production.
	EnterGetDiagnosticsStatement(c *GetDiagnosticsStatementContext)

	// EnterGrantStatement is called when entering the grantStatement production.
	EnterGrantStatement(c *GrantStatementContext)

	// EnterHoldLocatorStatement is called when entering the holdLocatorStatement production.
	EnterHoldLocatorStatement(c *HoldLocatorStatementContext)

	// EnterIncludeStatement is called when entering the includeStatement production.
	EnterIncludeStatement(c *IncludeStatementContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterLabelStatement is called when entering the labelStatement production.
	EnterLabelStatement(c *LabelStatementContext)

	// EnterLockTableStatement is called when entering the lockTableStatement production.
	EnterLockTableStatement(c *LockTableStatementContext)

	// EnterMergeStatement is called when entering the mergeStatement production.
	EnterMergeStatement(c *MergeStatementContext)

	// EnterOpenStatement is called when entering the openStatement production.
	EnterOpenStatement(c *OpenStatementContext)

	// EnterPrepareStatement is called when entering the prepareStatement production.
	EnterPrepareStatement(c *PrepareStatementContext)

	// EnterRefreshTableStatement is called when entering the refreshTableStatement production.
	EnterRefreshTableStatement(c *RefreshTableStatementContext)

	// EnterReleaseConnectionStatement is called when entering the releaseConnectionStatement production.
	EnterReleaseConnectionStatement(c *ReleaseConnectionStatementContext)

	// EnterReleaseSavepointStatement is called when entering the releaseSavepointStatement production.
	EnterReleaseSavepointStatement(c *ReleaseSavepointStatementContext)

	// EnterRenameStatement is called when entering the renameStatement production.
	EnterRenameStatement(c *RenameStatementContext)

	// EnterRevokeStatement is called when entering the revokeStatement production.
	EnterRevokeStatement(c *RevokeStatementContext)

	// EnterRollbackStatement is called when entering the rollbackStatement production.
	EnterRollbackStatement(c *RollbackStatementContext)

	// EnterSavepointStatement is called when entering the savepointStatement production.
	EnterSavepointStatement(c *SavepointStatementContext)

	// EnterSetAssignmentStatement is called when entering the setAssignmentStatement production.
	EnterSetAssignmentStatement(c *SetAssignmentStatementContext)

	// EnterSetConnectionStatement is called when entering the setConnectionStatement production.
	EnterSetConnectionStatement(c *SetConnectionStatementContext)

	// EnterSetEncryptionPasswordStatement is called when entering the setEncryptionPasswordStatement production.
	EnterSetEncryptionPasswordStatement(c *SetEncryptionPasswordStatementContext)

	// EnterSetPathStatement is called when entering the setPathStatement production.
	EnterSetPathStatement(c *SetPathStatementContext)

	// EnterSetSchemaStatement is called when entering the setSchemaStatement production.
	EnterSetSchemaStatement(c *SetSchemaStatementContext)

	// EnterSetSessionTimezoneStatement is called when entering the setSessionTimezoneStatement production.
	EnterSetSessionTimezoneStatement(c *SetSessionTimezoneStatementContext)

	// EnterSetSpecialRegisterStatement is called when entering the setSpecialRegisterStatement production.
	EnterSetSpecialRegisterStatement(c *SetSpecialRegisterStatementContext)

	// EnterSignalStatement is called when entering the signalStatement production.
	EnterSignalStatement(c *SignalStatementContext)

	// EnterTransferOwnershipStatement is called when entering the transferOwnershipStatement production.
	EnterTransferOwnershipStatement(c *TransferOwnershipStatementContext)

	// EnterTruncateStatement is called when entering the truncateStatement production.
	EnterTruncateStatement(c *TruncateStatementContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterValuesStatement is called when entering the valuesStatement production.
	EnterValuesStatement(c *ValuesStatementContext)

	// EnterValuesIntoStatement is called when entering the valuesIntoStatement production.
	EnterValuesIntoStatement(c *ValuesIntoStatementContext)

	// EnterWheneverStatement is called when entering the wheneverStatement production.
	EnterWheneverStatement(c *WheneverStatementContext)

	// EnterValuesIntoTargetVariable is called when entering the valuesIntoTargetVariable production.
	EnterValuesIntoTargetVariable(c *ValuesIntoTargetVariableContext)

	// EnterOwnedObject is called when entering the ownedObject production.
	EnterOwnedObject(c *OwnedObjectContext)

	// EnterNewOwner is called when entering the newOwner production.
	EnterNewOwner(c *NewOwnerContext)

	// EnterGrantCollectionStatement is called when entering the grantCollectionStatement production.
	EnterGrantCollectionStatement(c *GrantCollectionStatementContext)

	// EnterGrantDatabaseStatement is called when entering the grantDatabaseStatement production.
	EnterGrantDatabaseStatement(c *GrantDatabaseStatementContext)

	// EnterGrantFunctionOrProcedureStatement is called when entering the grantFunctionOrProcedureStatement production.
	EnterGrantFunctionOrProcedureStatement(c *GrantFunctionOrProcedureStatementContext)

	// EnterGrantPackageStatement is called when entering the grantPackageStatement production.
	EnterGrantPackageStatement(c *GrantPackageStatementContext)

	// EnterGrantPlanStatement is called when entering the grantPlanStatement production.
	EnterGrantPlanStatement(c *GrantPlanStatementContext)

	// EnterGrantSchemaStatement is called when entering the grantSchemaStatement production.
	EnterGrantSchemaStatement(c *GrantSchemaStatementContext)

	// EnterGrantSequenceStatement is called when entering the grantSequenceStatement production.
	EnterGrantSequenceStatement(c *GrantSequenceStatementContext)

	// EnterGrantSystemStatement is called when entering the grantSystemStatement production.
	EnterGrantSystemStatement(c *GrantSystemStatementContext)

	// EnterGrantTableStatement is called when entering the grantTableStatement production.
	EnterGrantTableStatement(c *GrantTableStatementContext)

	// EnterGrantTypeOrJarStatement is called when entering the grantTypeOrJarStatement production.
	EnterGrantTypeOrJarStatement(c *GrantTypeOrJarStatementContext)

	// EnterGrantVariableStatement is called when entering the grantVariableStatement production.
	EnterGrantVariableStatement(c *GrantVariableStatementContext)

	// EnterGrantUseOfStatement is called when entering the grantUseOfStatement production.
	EnterGrantUseOfStatement(c *GrantUseOfStatementContext)

	// EnterRevokeCollectionStatement is called when entering the revokeCollectionStatement production.
	EnterRevokeCollectionStatement(c *RevokeCollectionStatementContext)

	// EnterRevokeDatabaseStatement is called when entering the revokeDatabaseStatement production.
	EnterRevokeDatabaseStatement(c *RevokeDatabaseStatementContext)

	// EnterRevokeFunctionOrProcedureStatement is called when entering the revokeFunctionOrProcedureStatement production.
	EnterRevokeFunctionOrProcedureStatement(c *RevokeFunctionOrProcedureStatementContext)

	// EnterRevokePackageStatement is called when entering the revokePackageStatement production.
	EnterRevokePackageStatement(c *RevokePackageStatementContext)

	// EnterRevokePlanStatement is called when entering the revokePlanStatement production.
	EnterRevokePlanStatement(c *RevokePlanStatementContext)

	// EnterRevokeSchemaStatement is called when entering the revokeSchemaStatement production.
	EnterRevokeSchemaStatement(c *RevokeSchemaStatementContext)

	// EnterRevokeSequenceStatement is called when entering the revokeSequenceStatement production.
	EnterRevokeSequenceStatement(c *RevokeSequenceStatementContext)

	// EnterRevokeSystemStatement is called when entering the revokeSystemStatement production.
	EnterRevokeSystemStatement(c *RevokeSystemStatementContext)

	// EnterRevokeTableStatement is called when entering the revokeTableStatement production.
	EnterRevokeTableStatement(c *RevokeTableStatementContext)

	// EnterRevokeTypeOrJarStatement is called when entering the revokeTypeOrJarStatement production.
	EnterRevokeTypeOrJarStatement(c *RevokeTypeOrJarStatementContext)

	// EnterRevokeVariableStatement is called when entering the revokeVariableStatement production.
	EnterRevokeVariableStatement(c *RevokeVariableStatementContext)

	// EnterRevokeUseOfStatement is called when entering the revokeUseOfStatement production.
	EnterRevokeUseOfStatement(c *RevokeUseOfStatementContext)

	// EnterGrantUseOfTarget is called when entering the grantUseOfTarget production.
	EnterGrantUseOfTarget(c *GrantUseOfTargetContext)

	// EnterGrantVariableAuthority is called when entering the grantVariableAuthority production.
	EnterGrantVariableAuthority(c *GrantVariableAuthorityContext)

	// EnterGrantTableAuthority is called when entering the grantTableAuthority production.
	EnterGrantTableAuthority(c *GrantTableAuthorityContext)

	// EnterGrantSystemAuthority is called when entering the grantSystemAuthority production.
	EnterGrantSystemAuthority(c *GrantSystemAuthorityContext)

	// EnterGrantSequenceAuthority is called when entering the grantSequenceAuthority production.
	EnterGrantSequenceAuthority(c *GrantSequenceAuthorityContext)

	// EnterGrantSchemaAuthority is called when entering the grantSchemaAuthority production.
	EnterGrantSchemaAuthority(c *GrantSchemaAuthorityContext)

	// EnterGrantPlanAuthority is called when entering the grantPlanAuthority production.
	EnterGrantPlanAuthority(c *GrantPlanAuthorityContext)

	// EnterGrantPackageAuthority is called when entering the grantPackageAuthority production.
	EnterGrantPackageAuthority(c *GrantPackageAuthorityContext)

	// EnterPackageSpecification is called when entering the packageSpecification production.
	EnterPackageSpecification(c *PackageSpecificationContext)

	// EnterFunctionSpecification is called when entering the functionSpecification production.
	EnterFunctionSpecification(c *FunctionSpecificationContext)

	// EnterGrantee is called when entering the grantee production.
	EnterGrantee(c *GranteeContext)

	// EnterWithGrantOption is called when entering the withGrantOption production.
	EnterWithGrantOption(c *WithGrantOptionContext)

	// EnterRevokeByOption is called when entering the revokeByOption production.
	EnterRevokeByOption(c *RevokeByOptionContext)

	// EnterRevokeDependentPrivilegesOption is called when entering the revokeDependentPrivilegesOption production.
	EnterRevokeDependentPrivilegesOption(c *RevokeDependentPrivilegesOptionContext)

	// EnterGrantDatabaseAuthority is called when entering the grantDatabaseAuthority production.
	EnterGrantDatabaseAuthority(c *GrantDatabaseAuthorityContext)

	// EnterStatementInformation is called when entering the statementInformation production.
	EnterStatementInformation(c *StatementInformationContext)

	// EnterStatementInformationVariableEquate is called when entering the statementInformationVariableEquate production.
	EnterStatementInformationVariableEquate(c *StatementInformationVariableEquateContext)

	// EnterStatementInformationItemName is called when entering the statementInformationItemName production.
	EnterStatementInformationItemName(c *StatementInformationItemNameContext)

	// EnterConditionInformation is called when entering the conditionInformation production.
	EnterConditionInformation(c *ConditionInformationContext)

	// EnterConditionInformationVariableEquate is called when entering the conditionInformationVariableEquate production.
	EnterConditionInformationVariableEquate(c *ConditionInformationVariableEquateContext)

	// EnterConditionInformationItemName is called when entering the conditionInformationItemName production.
	EnterConditionInformationItemName(c *ConditionInformationItemNameContext)

	// EnterConnectionInformationItemName is called when entering the connectionInformationItemName production.
	EnterConnectionInformationItemName(c *ConnectionInformationItemNameContext)

	// EnterCombinedInformation is called when entering the combinedInformation production.
	EnterCombinedInformation(c *CombinedInformationContext)

	// EnterCombinedInformationOption is called when entering the combinedInformationOption production.
	EnterCombinedInformationOption(c *CombinedInformationOptionContext)

	// EnterFetchOrientation is called when entering the fetchOrientation production.
	EnterFetchOrientation(c *FetchOrientationContext)

	// EnterRowPositioned is called when entering the rowPositioned production.
	EnterRowPositioned(c *RowPositionedContext)

	// EnterRowsetPositioned is called when entering the rowsetPositioned production.
	EnterRowsetPositioned(c *RowsetPositionedContext)

	// EnterSingleRowFetch is called when entering the singleRowFetch production.
	EnterSingleRowFetch(c *SingleRowFetchContext)

	// EnterFetchTargetVariable is called when entering the fetchTargetVariable production.
	EnterFetchTargetVariable(c *FetchTargetVariableContext)

	// EnterMultipleRowFetch is called when entering the multipleRowFetch production.
	EnterMultipleRowFetch(c *MultipleRowFetchContext)

	// EnterMultipleRowFetchForClause is called when entering the multipleRowFetchForClause production.
	EnterMultipleRowFetchForClause(c *MultipleRowFetchForClauseContext)

	// EnterMultipleRowFetchIntoClause is called when entering the multipleRowFetchIntoClause production.
	EnterMultipleRowFetchIntoClause(c *MultipleRowFetchIntoClauseContext)

	// EnterExplainPlanClause is called when entering the explainPlanClause production.
	EnterExplainPlanClause(c *ExplainPlanClauseContext)

	// EnterExplainStmtcacheClause is called when entering the explainStmtcacheClause production.
	EnterExplainStmtcacheClause(c *ExplainStmtcacheClauseContext)

	// EnterExplainPackageClause is called when entering the explainPackageClause production.
	EnterExplainPackageClause(c *ExplainPackageClauseContext)

	// EnterExplainStabilizedDynamicQueryClause is called when entering the explainStabilizedDynamicQueryClause production.
	EnterExplainStabilizedDynamicQueryClause(c *ExplainStabilizedDynamicQueryClauseContext)

	// EnterPackageScopeSpecification is called when entering the packageScopeSpecification production.
	EnterPackageScopeSpecification(c *PackageScopeSpecificationContext)

	// EnterCollectionName is called when entering the collectionName production.
	EnterCollectionName(c *CollectionNameContext)

	// EnterPackageScopePackageName is called when entering the packageScopePackageName production.
	EnterPackageScopePackageName(c *PackageScopePackageNameContext)

	// EnterVersionName is called when entering the versionName production.
	EnterVersionName(c *VersionNameContext)

	// EnterSourceRowData is called when entering the sourceRowData production.
	EnterSourceRowData(c *SourceRowDataContext)

	// EnterAliasDesignation is called when entering the aliasDesignation production.
	EnterAliasDesignation(c *AliasDesignationContext)

	// EnterDropDatabaseClause is called when entering the dropDatabaseClause production.
	EnterDropDatabaseClause(c *DropDatabaseClauseContext)

	// EnterDropFunctionClause is called when entering the dropFunctionClause production.
	EnterDropFunctionClause(c *DropFunctionClauseContext)

	// EnterDropIndexClause is called when entering the dropIndexClause production.
	EnterDropIndexClause(c *DropIndexClauseContext)

	// EnterDropMaskClause is called when entering the dropMaskClause production.
	EnterDropMaskClause(c *DropMaskClauseContext)

	// EnterDropPackageClause is called when entering the dropPackageClause production.
	EnterDropPackageClause(c *DropPackageClauseContext)

	// EnterDropPermissionClause is called when entering the dropPermissionClause production.
	EnterDropPermissionClause(c *DropPermissionClauseContext)

	// EnterDropProcedureClause is called when entering the dropProcedureClause production.
	EnterDropProcedureClause(c *DropProcedureClauseContext)

	// EnterDropRoleClause is called when entering the dropRoleClause production.
	EnterDropRoleClause(c *DropRoleClauseContext)

	// EnterDropSequenceClause is called when entering the dropSequenceClause production.
	EnterDropSequenceClause(c *DropSequenceClauseContext)

	// EnterDropStogroupClause is called when entering the dropStogroupClause production.
	EnterDropStogroupClause(c *DropStogroupClauseContext)

	// EnterDropSynonymClause is called when entering the dropSynonymClause production.
	EnterDropSynonymClause(c *DropSynonymClauseContext)

	// EnterDropTableClause is called when entering the dropTableClause production.
	EnterDropTableClause(c *DropTableClauseContext)

	// EnterDropTablespaceClause is called when entering the dropTablespaceClause production.
	EnterDropTablespaceClause(c *DropTablespaceClauseContext)

	// EnterDropTriggerClause is called when entering the dropTriggerClause production.
	EnterDropTriggerClause(c *DropTriggerClauseContext)

	// EnterDropTrustedContextClause is called when entering the dropTrustedContextClause production.
	EnterDropTrustedContextClause(c *DropTrustedContextClauseContext)

	// EnterDropTypeClause is called when entering the dropTypeClause production.
	EnterDropTypeClause(c *DropTypeClauseContext)

	// EnterDropVariableClause is called when entering the dropVariableClause production.
	EnterDropVariableClause(c *DropVariableClauseContext)

	// EnterDropViewClause is called when entering the dropViewClause production.
	EnterDropViewClause(c *DropViewClauseContext)

	// EnterPackageDesignator is called when entering the packageDesignator production.
	EnterPackageDesignator(c *PackageDesignatorContext)

	// EnterDescribeUsingOption is called when entering the describeUsingOption production.
	EnterDescribeUsingOption(c *DescribeUsingOptionContext)

	// EnterDeclareGlobalTemporaryTableLikeClause is called when entering the declareGlobalTemporaryTableLikeClause production.
	EnterDeclareGlobalTemporaryTableLikeClause(c *DeclareGlobalTemporaryTableLikeClauseContext)

	// EnterOnCommitClause is called when entering the onCommitClause production.
	EnterOnCommitClause(c *OnCommitClauseContext)

	// EnterLoggedWithRollbackClause is called when entering the loggedWithRollbackClause production.
	EnterLoggedWithRollbackClause(c *LoggedWithRollbackClauseContext)

	// EnterCreateViewCheckOptionClause is called when entering the createViewCheckOptionClause production.
	EnterCreateViewCheckOptionClause(c *CreateViewCheckOptionClauseContext)

	// EnterTrustedContextDefaultRoleClause is called when entering the trustedContextDefaultRoleClause production.
	EnterTrustedContextDefaultRoleClause(c *TrustedContextDefaultRoleClauseContext)

	// EnterTrustedContextEnableDisableClause is called when entering the trustedContextEnableDisableClause production.
	EnterTrustedContextEnableDisableClause(c *TrustedContextEnableDisableClauseContext)

	// EnterTrustedContextDefaultSecurityLabelClause is called when entering the trustedContextDefaultSecurityLabelClause production.
	EnterTrustedContextDefaultSecurityLabelClause(c *TrustedContextDefaultSecurityLabelClauseContext)

	// EnterTrustedContextAttributesClause is called when entering the trustedContextAttributesClause production.
	EnterTrustedContextAttributesClause(c *TrustedContextAttributesClauseContext)

	// EnterTrustedContextWithUseForClause is called when entering the trustedContextWithUseForClause production.
	EnterTrustedContextWithUseForClause(c *TrustedContextWithUseForClauseContext)

	// EnterTrustedContextAttribute1 is called when entering the trustedContextAttribute1 production.
	EnterTrustedContextAttribute1(c *TrustedContextAttribute1Context)

	// EnterTrustedContextAttribute2 is called when entering the trustedContextAttribute2 production.
	EnterTrustedContextAttribute2(c *TrustedContextAttribute2Context)

	// EnterTrustedContextUseFor is called when entering the trustedContextUseFor production.
	EnterTrustedContextUseFor(c *TrustedContextUseForContext)

	// EnterUserOptions is called when entering the userOptions production.
	EnterUserOptions(c *UserOptionsContext)

	// EnterTriggerDefinition is called when entering the triggerDefinition production.
	EnterTriggerDefinition(c *TriggerDefinitionContext)

	// EnterTriggerActivationTime is called when entering the triggerActivationTime production.
	EnterTriggerActivationTime(c *TriggerActivationTimeContext)

	// EnterTriggerEvent is called when entering the triggerEvent production.
	EnterTriggerEvent(c *TriggerEventContext)

	// EnterTriggerGranularity is called when entering the triggerGranularity production.
	EnterTriggerGranularity(c *TriggerGranularityContext)

	// EnterTriggeredAction is called when entering the triggeredAction production.
	EnterTriggeredAction(c *TriggeredActionContext)

	// EnterSqlTriggerBody is called when entering the sqlTriggerBody production.
	EnterSqlTriggerBody(c *SqlTriggerBodyContext)

	// EnterTriggeredSqlStatement is called when entering the triggeredSqlStatement production.
	EnterTriggeredSqlStatement(c *TriggeredSqlStatementContext)

	// EnterTriggerDefinitionOption is called when entering the triggerDefinitionOption production.
	EnterTriggerDefinitionOption(c *TriggerDefinitionOptionContext)

	// EnterCreateTableInClause is called when entering the createTableInClause production.
	EnterCreateTableInClause(c *CreateTableInClauseContext)

	// EnterCustomVolatileClause is called when entering the customVolatileClause production.
	EnterCustomVolatileClause(c *CustomVolatileClauseContext)

	// EnterCreateTableColumnDefinition is called when entering the createTableColumnDefinition production.
	EnterCreateTableColumnDefinition(c *CreateTableColumnDefinitionContext)

	// EnterEditprocClause is called when entering the editprocClause production.
	EnterEditprocClause(c *EditprocClauseContext)

	// EnterValidprocClause is called when entering the validprocClause production.
	EnterValidprocClause(c *ValidprocClauseContext)

	// EnterAuditClause is called when entering the auditClause production.
	EnterAuditClause(c *AuditClauseContext)

	// EnterObidClause is called when entering the obidClause production.
	EnterObidClause(c *ObidClauseContext)

	// EnterDataCaptureClause is called when entering the dataCaptureClause production.
	EnterDataCaptureClause(c *DataCaptureClauseContext)

	// EnterRestrictOnDropClause is called when entering the restrictOnDropClause production.
	EnterRestrictOnDropClause(c *RestrictOnDropClauseContext)

	// EnterCcsidClause1 is called when entering the ccsidClause1 production.
	EnterCcsidClause1(c *CcsidClause1Context)

	// EnterCcsidClause2 is called when entering the ccsidClause2 production.
	EnterCcsidClause2(c *CcsidClause2Context)

	// EnterCardinalityClause is called when entering the cardinalityClause production.
	EnterCardinalityClause(c *CardinalityClauseContext)

	// EnterAppendClause is called when entering the appendClause production.
	EnterAppendClause(c *AppendClauseContext)

	// EnterMemberClause is called when entering the memberClause production.
	EnterMemberClause(c *MemberClauseContext)

	// EnterTrackmodClause is called when entering the trackmodClause production.
	EnterTrackmodClause(c *TrackmodClauseContext)

	// EnterPagenumClause is called when entering the pagenumClause production.
	EnterPagenumClause(c *PagenumClauseContext)

	// EnterFieldprocClause is called when entering the fieldprocClause production.
	EnterFieldprocClause(c *FieldprocClauseContext)

	// EnterAsSecurityLabelClause is called when entering the asSecurityLabelClause production.
	EnterAsSecurityLabelClause(c *AsSecurityLabelClauseContext)

	// EnterImplicitlyHiddenClause is called when entering the implicitlyHiddenClause production.
	EnterImplicitlyHiddenClause(c *ImplicitlyHiddenClauseContext)

	// EnterInlineLengthClause is called when entering the inlineLengthClause production.
	EnterInlineLengthClause(c *InlineLengthClauseContext)

	// EnterCopyOptions is called when entering the copyOptions production.
	EnterCopyOptions(c *CopyOptionsContext)

	// EnterCopyOptionIdentity is called when entering the copyOptionIdentity production.
	EnterCopyOptionIdentity(c *CopyOptionIdentityContext)

	// EnterCopyOptionRowChangeTimestamp is called when entering the copyOptionRowChangeTimestamp production.
	EnterCopyOptionRowChangeTimestamp(c *CopyOptionRowChangeTimestampContext)

	// EnterCopyOptionColumnDefaults is called when entering the copyOptionColumnDefaults production.
	EnterCopyOptionColumnDefaults(c *CopyOptionColumnDefaultsContext)

	// EnterCopyOptionXmlTypeModifiers is called when entering the copyOptionXmlTypeModifiers production.
	EnterCopyOptionXmlTypeModifiers(c *CopyOptionXmlTypeModifiersContext)

	// EnterAsResultTable is called when entering the asResultTable production.
	EnterAsResultTable(c *AsResultTableContext)

	// EnterDeclareGlobalTemporaryTableAsResultTable is called when entering the declareGlobalTemporaryTableAsResultTable production.
	EnterDeclareGlobalTemporaryTableAsResultTable(c *DeclareGlobalTemporaryTableAsResultTableContext)

	// EnterCreateTableMaterializedQueryDefinition is called when entering the createTableMaterializedQueryDefinition production.
	EnterCreateTableMaterializedQueryDefinition(c *CreateTableMaterializedQueryDefinitionContext)

	// EnterCreateTableColumnConstraint is called when entering the createTableColumnConstraint production.
	EnterCreateTableColumnConstraint(c *CreateTableColumnConstraintContext)

	// EnterOrganizationClause is called when entering the organizationClause production.
	EnterOrganizationClause(c *OrganizationClauseContext)

	// EnterCreateGlobalTemporaryTableColumnDefinition is called when entering the createGlobalTemporaryTableColumnDefinition production.
	EnterCreateGlobalTemporaryTableColumnDefinition(c *CreateGlobalTemporaryTableColumnDefinitionContext)

	// EnterDeclareGlobalTemporaryTableColumnDefinition is called when entering the declareGlobalTemporaryTableColumnDefinition production.
	EnterDeclareGlobalTemporaryTableColumnDefinition(c *DeclareGlobalTemporaryTableColumnDefinitionContext)

	// EnterParameterDeclaration1 is called when entering the parameterDeclaration1 production.
	EnterParameterDeclaration1(c *ParameterDeclaration1Context)

	// EnterParameterDeclaration2 is called when entering the parameterDeclaration2 production.
	EnterParameterDeclaration2(c *ParameterDeclaration2Context)

	// EnterParameterDeclaration3 is called when entering the parameterDeclaration3 production.
	EnterParameterDeclaration3(c *ParameterDeclaration3Context)

	// EnterCreateFunctionStatementExternalScalarOptions is called when entering the createFunctionStatementExternalScalarOptions production.
	EnterCreateFunctionStatementExternalScalarOptions(c *CreateFunctionStatementExternalScalarOptionsContext)

	// EnterExternalNameOption1 is called when entering the externalNameOption1 production.
	EnterExternalNameOption1(c *ExternalNameOption1Context)

	// EnterExternalNameOption2 is called when entering the externalNameOption2 production.
	EnterExternalNameOption2(c *ExternalNameOption2Context)

	// EnterDynamicResultSetOption is called when entering the dynamicResultSetOption production.
	EnterDynamicResultSetOption(c *DynamicResultSetOptionContext)

	// EnterLanguageOption1 is called when entering the languageOption1 production.
	EnterLanguageOption1(c *LanguageOption1Context)

	// EnterLanguageOption2 is called when entering the languageOption2 production.
	EnterLanguageOption2(c *LanguageOption2Context)

	// EnterLanguageOption3 is called when entering the languageOption3 production.
	EnterLanguageOption3(c *LanguageOption3Context)

	// EnterLanguageOption4 is called when entering the languageOption4 production.
	EnterLanguageOption4(c *LanguageOption4Context)

	// EnterLanguageOption5 is called when entering the languageOption5 production.
	EnterLanguageOption5(c *LanguageOption5Context)

	// EnterParameterStyleOption1 is called when entering the parameterStyleOption1 production.
	EnterParameterStyleOption1(c *ParameterStyleOption1Context)

	// EnterParameterStyleOption2 is called when entering the parameterStyleOption2 production.
	EnterParameterStyleOption2(c *ParameterStyleOption2Context)

	// EnterParameterStyleOption3 is called when entering the parameterStyleOption3 production.
	EnterParameterStyleOption3(c *ParameterStyleOption3Context)

	// EnterDeterministicOption is called when entering the deterministicOption production.
	EnterDeterministicOption(c *DeterministicOptionContext)

	// EnterFencedOption is called when entering the fencedOption production.
	EnterFencedOption(c *FencedOptionContext)

	// EnterNullInputOption1 is called when entering the nullInputOption1 production.
	EnterNullInputOption1(c *NullInputOption1Context)

	// EnterNullInputOption2 is called when entering the nullInputOption2 production.
	EnterNullInputOption2(c *NullInputOption2Context)

	// EnterNullInputOption3 is called when entering the nullInputOption3 production.
	EnterNullInputOption3(c *NullInputOption3Context)

	// EnterNullInputOption4 is called when entering the nullInputOption4 production.
	EnterNullInputOption4(c *NullInputOption4Context)

	// EnterDebugOption is called when entering the debugOption production.
	EnterDebugOption(c *DebugOptionContext)

	// EnterSqlDataOption1 is called when entering the sqlDataOption1 production.
	EnterSqlDataOption1(c *SqlDataOption1Context)

	// EnterSqlDataOption2 is called when entering the sqlDataOption2 production.
	EnterSqlDataOption2(c *SqlDataOption2Context)

	// EnterSqlDataOption3 is called when entering the sqlDataOption3 production.
	EnterSqlDataOption3(c *SqlDataOption3Context)

	// EnterSqlDataOption4 is called when entering the sqlDataOption4 production.
	EnterSqlDataOption4(c *SqlDataOption4Context)

	// EnterExternalActionOption is called when entering the externalActionOption production.
	EnterExternalActionOption(c *ExternalActionOptionContext)

	// EnterPackagePathOption is called when entering the packagePathOption production.
	EnterPackagePathOption(c *PackagePathOptionContext)

	// EnterScratchpadOption is called when entering the scratchpadOption production.
	EnterScratchpadOption(c *ScratchpadOptionContext)

	// EnterFinalCallOption is called when entering the finalCallOption production.
	EnterFinalCallOption(c *FinalCallOptionContext)

	// EnterParallelOption1 is called when entering the parallelOption1 production.
	EnterParallelOption1(c *ParallelOption1Context)

	// EnterParallelOption2 is called when entering the parallelOption2 production.
	EnterParallelOption2(c *ParallelOption2Context)

	// EnterDbinfoOption is called when entering the dbinfoOption production.
	EnterDbinfoOption(c *DbinfoOptionContext)

	// EnterCardinalityOption is called when entering the cardinalityOption production.
	EnterCardinalityOption(c *CardinalityOptionContext)

	// EnterCollectionIdOption is called when entering the collectionIdOption production.
	EnterCollectionIdOption(c *CollectionIdOptionContext)

	// EnterWlmEnvironmentOption1 is called when entering the wlmEnvironmentOption1 production.
	EnterWlmEnvironmentOption1(c *WlmEnvironmentOption1Context)

	// EnterWlmEnvironmentOption2 is called when entering the wlmEnvironmentOption2 production.
	EnterWlmEnvironmentOption2(c *WlmEnvironmentOption2Context)

	// EnterWlmEnvironmentOption3 is called when entering the wlmEnvironmentOption3 production.
	EnterWlmEnvironmentOption3(c *WlmEnvironmentOption3Context)

	// EnterAsuTimeOption is called when entering the asuTimeOption production.
	EnterAsuTimeOption(c *AsuTimeOptionContext)

	// EnterStayResidentOption is called when entering the stayResidentOption production.
	EnterStayResidentOption(c *StayResidentOptionContext)

	// EnterProgramTypeOption is called when entering the programTypeOption production.
	EnterProgramTypeOption(c *ProgramTypeOptionContext)

	// EnterSecurityOption is called when entering the securityOption production.
	EnterSecurityOption(c *SecurityOptionContext)

	// EnterStopAfterFailureOption is called when entering the stopAfterFailureOption production.
	EnterStopAfterFailureOption(c *StopAfterFailureOptionContext)

	// EnterRunOptionsOption is called when entering the runOptionsOption production.
	EnterRunOptionsOption(c *RunOptionsOptionContext)

	// EnterCommitOnReturnOption is called when entering the commitOnReturnOption production.
	EnterCommitOnReturnOption(c *CommitOnReturnOptionContext)

	// EnterSpecialRegistersOption is called when entering the specialRegistersOption production.
	EnterSpecialRegistersOption(c *SpecialRegistersOptionContext)

	// EnterSpecialRegistersOption2 is called when entering the specialRegistersOption2 production.
	EnterSpecialRegistersOption2(c *SpecialRegistersOption2Context)

	// EnterDispatchOption is called when entering the dispatchOption production.
	EnterDispatchOption(c *DispatchOptionContext)

	// EnterSecuredOption is called when entering the securedOption production.
	EnterSecuredOption(c *SecuredOptionContext)

	// EnterSpecificNameOption1 is called when entering the specificNameOption1 production.
	EnterSpecificNameOption1(c *SpecificNameOption1Context)

	// EnterSpecificNameOption2 is called when entering the specificNameOption2 production.
	EnterSpecificNameOption2(c *SpecificNameOption2Context)

	// EnterParameterOption1 is called when entering the parameterOption1 production.
	EnterParameterOption1(c *ParameterOption1Context)

	// EnterParameterOption2 is called when entering the parameterOption2 production.
	EnterParameterOption2(c *ParameterOption2Context)

	// EnterCreateFunctionStatementExternalTableOptions is called when entering the createFunctionStatementExternalTableOptions production.
	EnterCreateFunctionStatementExternalTableOptions(c *CreateFunctionStatementExternalTableOptionsContext)

	// EnterCreateFunctionStatementSourcedOptions is called when entering the createFunctionStatementSourcedOptions production.
	EnterCreateFunctionStatementSourcedOptions(c *CreateFunctionStatementSourcedOptionsContext)

	// EnterInlineSqlScalarFunctionDefinition is called when entering the inlineSqlScalarFunctionDefinition production.
	EnterInlineSqlScalarFunctionDefinition(c *InlineSqlScalarFunctionDefinitionContext)

	// EnterCreateFunctionStatementInlineSqlScalarOptions is called when entering the createFunctionStatementInlineSqlScalarOptions production.
	EnterCreateFunctionStatementInlineSqlScalarOptions(c *CreateFunctionStatementInlineSqlScalarOptionsContext)

	// EnterCompiledSqlScalarFunctionDefinition is called when entering the compiledSqlScalarFunctionDefinition production.
	EnterCompiledSqlScalarFunctionDefinition(c *CompiledSqlScalarFunctionDefinitionContext)

	// EnterCreateFunctionStatementCompiledSqlScalarOptions is called when entering the createFunctionStatementCompiledSqlScalarOptions production.
	EnterCreateFunctionStatementCompiledSqlScalarOptions(c *CreateFunctionStatementCompiledSqlScalarOptionsContext)

	// EnterSqlTableFunctionDefinition is called when entering the sqlTableFunctionDefinition production.
	EnterSqlTableFunctionDefinition(c *SqlTableFunctionDefinitionContext)

	// EnterCreateFunctionStatementSqlTableOptions is called when entering the createFunctionStatementSqlTableOptions production.
	EnterCreateFunctionStatementSqlTableOptions(c *CreateFunctionStatementSqlTableOptionsContext)

	// EnterSequenceAlias is called when entering the sequenceAlias production.
	EnterSequenceAlias(c *SequenceAliasContext)

	// EnterTableAlias is called when entering the tableAlias production.
	EnterTableAlias(c *TableAliasContext)

	// EnterAuthorization is called when entering the authorization production.
	EnterAuthorization(c *AuthorizationContext)

	// EnterSearchedDelete is called when entering the searchedDelete production.
	EnterSearchedDelete(c *SearchedDeleteContext)

	// EnterPositionedDelete is called when entering the positionedDelete production.
	EnterPositionedDelete(c *PositionedDeleteContext)

	// EnterSearchedUpdate is called when entering the searchedUpdate production.
	EnterSearchedUpdate(c *SearchedUpdateContext)

	// EnterPositionedUpdate is called when entering the positionedUpdate production.
	EnterPositionedUpdate(c *PositionedUpdateContext)

	// EnterSourceValues is called when entering the sourceValues production.
	EnterSourceValues(c *SourceValuesContext)

	// EnterValuesSingleRow is called when entering the valuesSingleRow production.
	EnterValuesSingleRow(c *ValuesSingleRowContext)

	// EnterValuesMultipleRow is called when entering the valuesMultipleRow production.
	EnterValuesMultipleRow(c *ValuesMultipleRowContext)

	// EnterMatchingCondition is called when entering the matchingCondition production.
	EnterMatchingCondition(c *MatchingConditionContext)

	// EnterModificationOperation is called when entering the modificationOperation production.
	EnterModificationOperation(c *ModificationOperationContext)

	// EnterAssignmentClause is called when entering the assignmentClause production.
	EnterAssignmentClause(c *AssignmentClauseContext)

	// EnterSetAssignmentClause is called when entering the setAssignmentClause production.
	EnterSetAssignmentClause(c *SetAssignmentClauseContext)

	// EnterSetAssignmentTargetVariable is called when entering the setAssignmentTargetVariable production.
	EnterSetAssignmentTargetVariable(c *SetAssignmentTargetVariableContext)

	// EnterUpdateOperation is called when entering the updateOperation production.
	EnterUpdateOperation(c *UpdateOperationContext)

	// EnterDeleteOperation is called when entering the deleteOperation production.
	EnterDeleteOperation(c *DeleteOperationContext)

	// EnterInsertOperation is called when entering the insertOperation production.
	EnterInsertOperation(c *InsertOperationContext)

	// EnterSignalInformation is called when entering the signalInformation production.
	EnterSignalInformation(c *SignalInformationContext)

	// EnterValuesList1 is called when entering the valuesList1 production.
	EnterValuesList1(c *ValuesList1Context)

	// EnterValuesList2 is called when entering the valuesList2 production.
	EnterValuesList2(c *ValuesList2Context)

	// EnterValuesList3 is called when entering the valuesList3 production.
	EnterValuesList3(c *ValuesList3Context)

	// EnterValuesList4 is called when entering the valuesList4 production.
	EnterValuesList4(c *ValuesList4Context)

	// EnterIncludeColumns is called when entering the includeColumns production.
	EnterIncludeColumns(c *IncludeColumnsContext)

	// EnterMultipleRowInsert is called when entering the multipleRowInsert production.
	EnterMultipleRowInsert(c *MultipleRowInsertContext)

	// EnterRegenerateClause is called when entering the regenerateClause production.
	EnterRegenerateClause(c *RegenerateClauseContext)

	// EnterAlterIndexOptions is called when entering the alterIndexOptions production.
	EnterAlterIndexOptions(c *AlterIndexOptionsContext)

	// EnterBufferpoolOption is called when entering the bufferpoolOption production.
	EnterBufferpoolOption(c *BufferpoolOptionContext)

	// EnterCloseOption is called when entering the closeOption production.
	EnterCloseOption(c *CloseOptionContext)

	// EnterCopyOption is called when entering the copyOption production.
	EnterCopyOption(c *CopyOptionContext)

	// EnterDssizeOption is called when entering the dssizeOption production.
	EnterDssizeOption(c *DssizeOptionContext)

	// EnterPiecesizeOption is called when entering the piecesizeOption production.
	EnterPiecesizeOption(c *PiecesizeOptionContext)

	// EnterClusterOption is called when entering the clusterOption production.
	EnterClusterOption(c *ClusterOptionContext)

	// EnterPaddedOption is called when entering the paddedOption production.
	EnterPaddedOption(c *PaddedOptionContext)

	// EnterCompressOption is called when entering the compressOption production.
	EnterCompressOption(c *CompressOptionContext)

	// EnterDefineOption is called when entering the defineOption production.
	EnterDefineOption(c *DefineOptionContext)

	// EnterLocksizeOption is called when entering the locksizeOption production.
	EnterLocksizeOption(c *LocksizeOptionContext)

	// EnterLockmaxOption is called when entering the lockmaxOption production.
	EnterLockmaxOption(c *LockmaxOptionContext)

	// EnterEnableDisableOption is called when entering the enableDisableOption production.
	EnterEnableDisableOption(c *EnableDisableOptionContext)

	// EnterLoggedOption is called when entering the loggedOption production.
	EnterLoggedOption(c *LoggedOptionContext)

	// EnterNotAtomicPhrase is called when entering the notAtomicPhrase production.
	EnterNotAtomicPhrase(c *NotAtomicPhraseContext)

	// EnterAlterIndexPartitionOptions is called when entering the alterIndexPartitionOptions production.
	EnterAlterIndexPartitionOptions(c *AlterIndexPartitionOptionsContext)

	// EnterUsingSpecification1 is called when entering the usingSpecification1 production.
	EnterUsingSpecification1(c *UsingSpecification1Context)

	// EnterFreeSpecification is called when entering the freeSpecification production.
	EnterFreeSpecification(c *FreeSpecificationContext)

	// EnterGbpcacheSpecification is called when entering the gbpcacheSpecification production.
	EnterGbpcacheSpecification(c *GbpcacheSpecificationContext)

	// EnterPartitionElement is called when entering the partitionElement production.
	EnterPartitionElement(c *PartitionElementContext)

	// EnterApplCompatValue is called when entering the applCompatValue production.
	EnterApplCompatValue(c *ApplCompatValueContext)

	// EnterFunctionLevel is called when entering the functionLevel production.
	EnterFunctionLevel(c *FunctionLevelContext)

	// EnterFunctionParameterType is called when entering the functionParameterType production.
	EnterFunctionParameterType(c *FunctionParameterTypeContext)

	// EnterFunctionDataType is called when entering the functionDataType production.
	EnterFunctionDataType(c *FunctionDataTypeContext)

	// EnterFunctionBuiltInType is called when entering the functionBuiltInType production.
	EnterFunctionBuiltInType(c *FunctionBuiltInTypeContext)

	// EnterProcedureBuiltinType is called when entering the procedureBuiltinType production.
	EnterProcedureBuiltinType(c *ProcedureBuiltinTypeContext)

	// EnterCreateTypeArrayBuiltinType is called when entering the createTypeArrayBuiltinType production.
	EnterCreateTypeArrayBuiltinType(c *CreateTypeArrayBuiltinTypeContext)

	// EnterCreateTypeArrayBuiltinType2 is called when entering the createTypeArrayBuiltinType2 production.
	EnterCreateTypeArrayBuiltinType2(c *CreateTypeArrayBuiltinType2Context)

	// EnterCreateVariableBuiltInType is called when entering the createVariableBuiltInType production.
	EnterCreateVariableBuiltInType(c *CreateVariableBuiltInTypeContext)

	// EnterSourceDataType is called when entering the sourceDataType production.
	EnterSourceDataType(c *SourceDataTypeContext)

	// EnterFunctionExternalOptionList is called when entering the functionExternalOptionList production.
	EnterFunctionExternalOptionList(c *FunctionExternalOptionListContext)

	// EnterFunctionCompiledSqlScalarOptionList is called when entering the functionCompiledSqlScalarOptionList production.
	EnterFunctionCompiledSqlScalarOptionList(c *FunctionCompiledSqlScalarOptionListContext)

	// EnterFunctionInlineSqlScalarOptionList is called when entering the functionInlineSqlScalarOptionList production.
	EnterFunctionInlineSqlScalarOptionList(c *FunctionInlineSqlScalarOptionListContext)

	// EnterFunctionSqlTableOptionList is called when entering the functionSqlTableOptionList production.
	EnterFunctionSqlTableOptionList(c *FunctionSqlTableOptionListContext)

	// EnterProcedureOptionList is called when entering the procedureOptionList production.
	EnterProcedureOptionList(c *ProcedureOptionListContext)

	// EnterCreateProcedureOptionList is called when entering the createProcedureOptionList production.
	EnterCreateProcedureOptionList(c *CreateProcedureOptionListContext)

	// EnterProcedureSQLPLOptionList is called when entering the procedureSQLPLOptionList production.
	EnterProcedureSQLPLOptionList(c *ProcedureSQLPLOptionListContext)

	// EnterVersionOption is called when entering the versionOption production.
	EnterVersionOption(c *VersionOptionContext)

	// EnterCommitOnReturnOptionSQLPL is called when entering the commitOnReturnOptionSQLPL production.
	EnterCommitOnReturnOptionSQLPL(c *CommitOnReturnOptionSQLPLContext)

	// EnterSchemaQualifierOption is called when entering the schemaQualifierOption production.
	EnterSchemaQualifierOption(c *SchemaQualifierOptionContext)

	// EnterCurrentDataOption is called when entering the currentDataOption production.
	EnterCurrentDataOption(c *CurrentDataOptionContext)

	// EnterImmediateWriteOption is called when entering the immediateWriteOption production.
	EnterImmediateWriteOption(c *ImmediateWriteOptionContext)

	// EnterExplainOption is called when entering the explainOption production.
	EnterExplainOption(c *ExplainOptionContext)

	// EnterReoptOption is called when entering the reoptOption production.
	EnterReoptOption(c *ReoptOptionContext)

	// EnterPackageOwnerOption is called when entering the packageOwnerOption production.
	EnterPackageOwnerOption(c *PackageOwnerOptionContext)

	// EnterDeferPrepareOption is called when entering the deferPrepareOption production.
	EnterDeferPrepareOption(c *DeferPrepareOptionContext)

	// EnterDegreeOption is called when entering the degreeOption production.
	EnterDegreeOption(c *DegreeOptionContext)

	// EnterDynamicRulesOption is called when entering the dynamicRulesOption production.
	EnterDynamicRulesOption(c *DynamicRulesOptionContext)

	// EnterConcurrentAccessOption is called when entering the concurrentAccessOption production.
	EnterConcurrentAccessOption(c *ConcurrentAccessOptionContext)

	// EnterApplicationEncodingOption is called when entering the applicationEncodingOption production.
	EnterApplicationEncodingOption(c *ApplicationEncodingOptionContext)

	// EnterIsolationLevelOption is called when entering the isolationLevelOption production.
	EnterIsolationLevelOption(c *IsolationLevelOptionContext)

	// EnterKeepDynamicOption is called when entering the keepDynamicOption production.
	EnterKeepDynamicOption(c *KeepDynamicOptionContext)

	// EnterOpthintOption is called when entering the opthintOption production.
	EnterOpthintOption(c *OpthintOptionContext)

	// EnterSqlPathOption is called when entering the sqlPathOption production.
	EnterSqlPathOption(c *SqlPathOptionContext)

	// EnterSqlPathOptionItem is called when entering the sqlPathOptionItem production.
	EnterSqlPathOptionItem(c *SqlPathOptionItemContext)

	// EnterQueryAccelerationOption is called when entering the queryAccelerationOption production.
	EnterQueryAccelerationOption(c *QueryAccelerationOptionContext)

	// EnterQueryAccelerationOptionItem is called when entering the queryAccelerationOptionItem production.
	EnterQueryAccelerationOptionItem(c *QueryAccelerationOptionItemContext)

	// EnterGetAccelArchiveOption is called when entering the getAccelArchiveOption production.
	EnterGetAccelArchiveOption(c *GetAccelArchiveOptionContext)

	// EnterAccelerationOption is called when entering the accelerationOption production.
	EnterAccelerationOption(c *AccelerationOptionContext)

	// EnterReleaseAtOption is called when entering the releaseAtOption production.
	EnterReleaseAtOption(c *ReleaseAtOptionContext)

	// EnterBusinessTimeSensitiveOption is called when entering the businessTimeSensitiveOption production.
	EnterBusinessTimeSensitiveOption(c *BusinessTimeSensitiveOptionContext)

	// EnterSystemTimeSensitiveOption is called when entering the systemTimeSensitiveOption production.
	EnterSystemTimeSensitiveOption(c *SystemTimeSensitiveOptionContext)

	// EnterArchiveSensitiveOption is called when entering the archiveSensitiveOption production.
	EnterArchiveSensitiveOption(c *ArchiveSensitiveOptionContext)

	// EnterApplcompatOption is called when entering the applcompatOption production.
	EnterApplcompatOption(c *ApplcompatOptionContext)

	// EnterValidateOption is called when entering the validateOption production.
	EnterValidateOption(c *ValidateOptionContext)

	// EnterRoundingOption is called when entering the roundingOption production.
	EnterRoundingOption(c *RoundingOptionContext)

	// EnterRoundingOptionItem is called when entering the roundingOptionItem production.
	EnterRoundingOptionItem(c *RoundingOptionItemContext)

	// EnterDateFormatOption is called when entering the dateFormatOption production.
	EnterDateFormatOption(c *DateFormatOptionContext)

	// EnterDateTimeFormatOptionItem is called when entering the dateTimeFormatOptionItem production.
	EnterDateTimeFormatOptionItem(c *DateTimeFormatOptionItemContext)

	// EnterTimeFormatOption is called when entering the timeFormatOption production.
	EnterTimeFormatOption(c *TimeFormatOptionContext)

	// EnterDecimalOption is called when entering the decimalOption production.
	EnterDecimalOption(c *DecimalOptionContext)

	// EnterForUpdateOption is called when entering the forUpdateOption production.
	EnterForUpdateOption(c *ForUpdateOptionContext)

	// EnterConcentrateStatementsOption is called when entering the concentrateStatementsOption production.
	EnterConcentrateStatementsOption(c *ConcentrateStatementsOptionContext)

	// EnterAcceleratorOption is called when entering the acceleratorOption production.
	EnterAcceleratorOption(c *AcceleratorOptionContext)

	// EnterProcedureDataType is called when entering the procedureDataType production.
	EnterProcedureDataType(c *ProcedureDataTypeContext)

	// EnterAlterSequenceOptionList is called when entering the alterSequenceOptionList production.
	EnterAlterSequenceOptionList(c *AlterSequenceOptionListContext)

	// EnterCreateSequenceOptionList is called when entering the createSequenceOptionList production.
	EnterCreateSequenceOptionList(c *CreateSequenceOptionListContext)

	// EnterAsTypeOption is called when entering the asTypeOption production.
	EnterAsTypeOption(c *AsTypeOptionContext)

	// EnterStartOption is called when entering the startOption production.
	EnterStartOption(c *StartOptionContext)

	// EnterRestartOption is called when entering the restartOption production.
	EnterRestartOption(c *RestartOptionContext)

	// EnterIncrementOption is called when entering the incrementOption production.
	EnterIncrementOption(c *IncrementOptionContext)

	// EnterMinvalueOption is called when entering the minvalueOption production.
	EnterMinvalueOption(c *MinvalueOptionContext)

	// EnterMaxvalueOption is called when entering the maxvalueOption production.
	EnterMaxvalueOption(c *MaxvalueOptionContext)

	// EnterCycleOption is called when entering the cycleOption production.
	EnterCycleOption(c *CycleOptionContext)

	// EnterCacheOption is called when entering the cacheOption production.
	EnterCacheOption(c *CacheOptionContext)

	// EnterOrderOption is called when entering the orderOption production.
	EnterOrderOption(c *OrderOptionContext)

	// EnterKeyLabelOption is called when entering the keyLabelOption production.
	EnterKeyLabelOption(c *KeyLabelOptionContext)

	// EnterDataclasOption is called when entering the dataclasOption production.
	EnterDataclasOption(c *DataclasOptionContext)

	// EnterMgmtclasOption is called when entering the mgmtclasOption production.
	EnterMgmtclasOption(c *MgmtclasOptionContext)

	// EnterStorclasOption is called when entering the storclasOption production.
	EnterStorclasOption(c *StorclasOptionContext)

	// EnterAlterStogroupOptionList is called when entering the alterStogroupOptionList production.
	EnterAlterStogroupOptionList(c *AlterStogroupOptionListContext)

	// EnterAlterTableOptionList is called when entering the alterTableOptionList production.
	EnterAlterTableOptionList(c *AlterTableOptionListContext)

	// EnterAlterTablespaceOptionList is called when entering the alterTablespaceOptionList production.
	EnterAlterTablespaceOptionList(c *AlterTablespaceOptionListContext)

	// EnterCreateTablespaceOptionList is called when entering the createTablespaceOptionList production.
	EnterCreateTablespaceOptionList(c *CreateTablespaceOptionListContext)

	// EnterTrustedContextOptionList is called when entering the trustedContextOptionList production.
	EnterTrustedContextOptionList(c *TrustedContextOptionListContext)

	// EnterDatabaseOptionList is called when entering the databaseOptionList production.
	EnterDatabaseOptionList(c *DatabaseOptionListContext)

	// EnterCreateIndexOptionList is called when entering the createIndexOptionList production.
	EnterCreateIndexOptionList(c *CreateIndexOptionListContext)

	// EnterCreateLobTablespaceOptionList is called when entering the createLobTablespaceOptionList production.
	EnterCreateLobTablespaceOptionList(c *CreateLobTablespaceOptionListContext)

	// EnterInDatabaseOption is called when entering the inDatabaseOption production.
	EnterInDatabaseOption(c *InDatabaseOptionContext)

	// EnterSegsizeOption is called when entering the segsizeOption production.
	EnterSegsizeOption(c *SegsizeOptionContext)

	// EnterNumpartsOption is called when entering the numpartsOption production.
	EnterNumpartsOption(c *NumpartsOptionContext)

	// EnterPartitionByGrowthSpecification is called when entering the partitionByGrowthSpecification production.
	EnterPartitionByGrowthSpecification(c *PartitionByGrowthSpecificationContext)

	// EnterPartitionByRangeSpecification is called when entering the partitionByRangeSpecification production.
	EnterPartitionByRangeSpecification(c *PartitionByRangeSpecificationContext)

	// EnterPartitionByRangePartitionPhrase is called when entering the partitionByRangePartitionPhrase production.
	EnterPartitionByRangePartitionPhrase(c *PartitionByRangePartitionPhraseContext)

	// EnterInsertAlgorithmOption is called when entering the insertAlgorithmOption production.
	EnterInsertAlgorithmOption(c *InsertAlgorithmOptionContext)

	// EnterMaxrowsOption is called when entering the maxrowsOption production.
	EnterMaxrowsOption(c *MaxrowsOptionContext)

	// EnterMaxpartitionsOption is called when entering the maxpartitionsOption production.
	EnterMaxpartitionsOption(c *MaxpartitionsOptionContext)

	// EnterUsingSpecification2 is called when entering the usingSpecification2 production.
	EnterUsingSpecification2(c *UsingSpecification2Context)

	// EnterStogroupOptions is called when entering the stogroupOptions production.
	EnterStogroupOptions(c *StogroupOptionsContext)

	// EnterXmlIndexSpecification is called when entering the xmlIndexSpecification production.
	EnterXmlIndexSpecification(c *XmlIndexSpecificationContext)

	// EnterXmlPatternClause is called when entering the xmlPatternClause production.
	EnterXmlPatternClause(c *XmlPatternClauseContext)

	// EnterAlterAttributesOptions is called when entering the alterAttributesOptions production.
	EnterAlterAttributesOptions(c *AlterAttributesOptionsContext)

	// EnterAddAttributesOptions is called when entering the addAttributesOptions production.
	EnterAddAttributesOptions(c *AddAttributesOptionsContext)

	// EnterDropAttributesOptions is called when entering the dropAttributesOptions production.
	EnterDropAttributesOptions(c *DropAttributesOptionsContext)

	// EnterIncludeColumnPhrase is called when entering the includeColumnPhrase production.
	EnterIncludeColumnPhrase(c *IncludeColumnPhraseContext)

	// EnterUserClause is called when entering the userClause production.
	EnterUserClause(c *UserClauseContext)

	// EnterUserClauseAddOptions is called when entering the userClauseAddOptions production.
	EnterUserClauseAddOptions(c *UserClauseAddOptionsContext)

	// EnterUserClauseReplaceOptions is called when entering the userClauseReplaceOptions production.
	EnterUserClauseReplaceOptions(c *UserClauseReplaceOptionsContext)

	// EnterUserClauseDropOptions is called when entering the userClauseDropOptions production.
	EnterUserClauseDropOptions(c *UserClauseDropOptionsContext)

	// EnterUseOptions is called when entering the useOptions production.
	EnterUseOptions(c *UseOptionsContext)

	// EnterAlterPartitionClause is called when entering the alterPartitionClause production.
	EnterAlterPartitionClause(c *AlterPartitionClauseContext)

	// EnterUsingBlock is called when entering the usingBlock production.
	EnterUsingBlock(c *UsingBlockContext)

	// EnterFreeBlock is called when entering the freeBlock production.
	EnterFreeBlock(c *FreeBlockContext)

	// EnterMoveTableClause is called when entering the moveTableClause production.
	EnterMoveTableClause(c *MoveTableClauseContext)

	// EnterGbpcacheBlock is called when entering the gbpcacheBlock production.
	EnterGbpcacheBlock(c *GbpcacheBlockContext)

	// EnterAliasDesignator is called when entering the aliasDesignator production.
	EnterAliasDesignator(c *AliasDesignatorContext)

	// EnterMultipleColumnList is called when entering the multipleColumnList production.
	EnterMultipleColumnList(c *MultipleColumnListContext)

	// EnterFunctionDesignator is called when entering the functionDesignator production.
	EnterFunctionDesignator(c *FunctionDesignatorContext)

	// EnterParameterType is called when entering the parameterType production.
	EnterParameterType(c *ParameterTypeContext)

	// EnterAlterTableColumnDefinitionOptionList1 is called when entering the alterTableColumnDefinitionOptionList1 production.
	EnterAlterTableColumnDefinitionOptionList1(c *AlterTableColumnDefinitionOptionList1Context)

	// EnterAlterTableColumnDefinitionOptionList2 is called when entering the alterTableColumnDefinitionOptionList2 production.
	EnterAlterTableColumnDefinitionOptionList2(c *AlterTableColumnDefinitionOptionList2Context)

	// EnterColumnConstraint is called when entering the columnConstraint production.
	EnterColumnConstraint(c *ColumnConstraintContext)

	// EnterGeneratedClause is called when entering the generatedClause production.
	EnterGeneratedClause(c *GeneratedClauseContext)

	// EnterGeneratedClause2 is called when entering the generatedClause2 production.
	EnterGeneratedClause2(c *GeneratedClause2Context)

	// EnterAsIdentityClause is called when entering the asIdentityClause production.
	EnterAsIdentityClause(c *AsIdentityClauseContext)

	// EnterAsIdentityClauseOptionList is called when entering the asIdentityClauseOptionList production.
	EnterAsIdentityClauseOptionList(c *AsIdentityClauseOptionListContext)

	// EnterAsRowChangeTimestampClause is called when entering the asRowChangeTimestampClause production.
	EnterAsRowChangeTimestampClause(c *AsRowChangeTimestampClauseContext)

	// EnterAsRowTransactionStartIDClause is called when entering the asRowTransactionStartIDClause production.
	EnterAsRowTransactionStartIDClause(c *AsRowTransactionStartIDClauseContext)

	// EnterAsRowTransactionTimestampClause is called when entering the asRowTransactionTimestampClause production.
	EnterAsRowTransactionTimestampClause(c *AsRowTransactionTimestampClauseContext)

	// EnterAsGeneratedExpressionClause is called when entering the asGeneratedExpressionClause production.
	EnterAsGeneratedExpressionClause(c *AsGeneratedExpressionClauseContext)

	// EnterNonDeterministicExpression is called when entering the nonDeterministicExpression production.
	EnterNonDeterministicExpression(c *NonDeterministicExpressionContext)

	// EnterNonDeterministicExpressionSessionVariable is called when entering the nonDeterministicExpressionSessionVariable production.
	EnterNonDeterministicExpressionSessionVariable(c *NonDeterministicExpressionSessionVariableContext)

	// EnterColumnAlteration is called when entering the columnAlteration production.
	EnterColumnAlteration(c *ColumnAlterationContext)

	// EnterColumnAlterationOptionList is called when entering the columnAlterationOptionList production.
	EnterColumnAlterationOptionList(c *ColumnAlterationOptionListContext)

	// EnterAlteredDataType is called when entering the alteredDataType production.
	EnterAlteredDataType(c *AlteredDataTypeContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterBuiltInType is called when entering the builtInType production.
	EnterBuiltInType(c *BuiltInTypeContext)

	// EnterSequenceDataType is called when entering the sequenceDataType production.
	EnterSequenceDataType(c *SequenceDataTypeContext)

	// EnterSequenceBuiltInType is called when entering the sequenceBuiltInType production.
	EnterSequenceBuiltInType(c *SequenceBuiltInTypeContext)

	// EnterSqlDataType is called when entering the sqlDataType production.
	EnterSqlDataType(c *SqlDataTypeContext)

	// EnterXmlTypeModifier is called when entering the xmlTypeModifier production.
	EnterXmlTypeModifier(c *XmlTypeModifierContext)

	// EnterXmlSchemaSpecification is called when entering the xmlSchemaSpecification production.
	EnterXmlSchemaSpecification(c *XmlSchemaSpecificationContext)

	// EnterXmlElementName is called when entering the xmlElementName production.
	EnterXmlElementName(c *XmlElementNameContext)

	// EnterPiName is called when entering the piName production.
	EnterPiName(c *PiNameContext)

	// EnterRegisteredXmlSchemaName is called when entering the registeredXmlSchemaName production.
	EnterRegisteredXmlSchemaName(c *RegisteredXmlSchemaNameContext)

	// EnterTargetNamespace is called when entering the targetNamespace production.
	EnterTargetNamespace(c *TargetNamespaceContext)

	// EnterSchemaLocation is called when entering the schemaLocation production.
	EnterSchemaLocation(c *SchemaLocationContext)

	// EnterIdentityAlteration is called when entering the identityAlteration production.
	EnterIdentityAlteration(c *IdentityAlterationContext)

	// EnterUniqueConstraint is called when entering the uniqueConstraint production.
	EnterUniqueConstraint(c *UniqueConstraintContext)

	// EnterReferentialConstraint is called when entering the referentialConstraint production.
	EnterReferentialConstraint(c *ReferentialConstraintContext)

	// EnterReferencesClause is called when entering the referencesClause production.
	EnterReferencesClause(c *ReferencesClauseContext)

	// EnterCheckConstraint is called when entering the checkConstraint production.
	EnterCheckConstraint(c *CheckConstraintContext)

	// EnterPartitioningClause is called when entering the partitioningClause production.
	EnterPartitioningClause(c *PartitioningClauseContext)

	// EnterPartitionExpression is called when entering the partitionExpression production.
	EnterPartitionExpression(c *PartitionExpressionContext)

	// EnterPartitionLimitKey is called when entering the partitionLimitKey production.
	EnterPartitionLimitKey(c *PartitionLimitKeyContext)

	// EnterPartitioningPhrase is called when entering the partitioningPhrase production.
	EnterPartitioningPhrase(c *PartitioningPhraseContext)

	// EnterPartitionHashSpace is called when entering the partitionHashSpace production.
	EnterPartitionHashSpace(c *PartitionHashSpaceContext)

	// EnterAlterHashOrganization is called when entering the alterHashOrganization production.
	EnterAlterHashOrganization(c *AlterHashOrganizationContext)

	// EnterPartitioningClauseElement is called when entering the partitioningClauseElement production.
	EnterPartitioningClauseElement(c *PartitioningClauseElementContext)

	// EnterPartitionClause is called when entering the partitionClause production.
	EnterPartitionClause(c *PartitionClauseContext)

	// EnterRotatePartitionClause is called when entering the rotatePartitionClause production.
	EnterRotatePartitionClause(c *RotatePartitionClauseContext)

	// EnterExtraRowOption is called when entering the extraRowOption production.
	EnterExtraRowOption(c *ExtraRowOptionContext)

	// EnterMaterializedQueryDefinition is called when entering the materializedQueryDefinition production.
	EnterMaterializedQueryDefinition(c *MaterializedQueryDefinitionContext)

	// EnterMaterializedQueryAlteration is called when entering the materializedQueryAlteration production.
	EnterMaterializedQueryAlteration(c *MaterializedQueryAlterationContext)

	// EnterRefreshableTableOptions is called when entering the refreshableTableOptions production.
	EnterRefreshableTableOptions(c *RefreshableTableOptionsContext)

	// EnterDataInitiallyDeferredPhrase is called when entering the dataInitiallyDeferredPhrase production.
	EnterDataInitiallyDeferredPhrase(c *DataInitiallyDeferredPhraseContext)

	// EnterRefreshDeferredPhrase is called when entering the refreshDeferredPhrase production.
	EnterRefreshDeferredPhrase(c *RefreshDeferredPhraseContext)

	// EnterRefreshableTableOptionsList is called when entering the refreshableTableOptionsList production.
	EnterRefreshableTableOptionsList(c *RefreshableTableOptionsListContext)

	// EnterMaterializedQueryTableAlteration is called when entering the materializedQueryTableAlteration production.
	EnterMaterializedQueryTableAlteration(c *MaterializedQueryTableAlterationContext)

	// EnterPeriodDefinition is called when entering the periodDefinition production.
	EnterPeriodDefinition(c *PeriodDefinitionContext)

	// EnterAlterTableColumnDefinition is called when entering the alterTableColumnDefinition production.
	EnterAlterTableColumnDefinition(c *AlterTableColumnDefinitionContext)

	// EnterExternalProgramName is called when entering the externalProgramName production.
	EnterExternalProgramName(c *ExternalProgramNameContext)

	// EnterPackagePath is called when entering the packagePath production.
	EnterPackagePath(c *PackagePathContext)

	// EnterCollectionID is called when entering the collectionID production.
	EnterCollectionID(c *CollectionIDContext)

	// EnterRunTimeOptions is called when entering the runTimeOptions production.
	EnterRunTimeOptions(c *RunTimeOptionsContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterKeyExpression is called when entering the keyExpression production.
	EnterKeyExpression(c *KeyExpressionContext)

	// EnterRowChangeExpression is called when entering the rowChangeExpression production.
	EnterRowChangeExpression(c *RowChangeExpressionContext)

	// EnterSequenceReference is called when entering the sequenceReference production.
	EnterSequenceReference(c *SequenceReferenceContext)

	// EnterFunctionInvocation is called when entering the functionInvocation production.
	EnterFunctionInvocation(c *FunctionInvocationContext)

	// EnterScalarFunctionInvocation is called when entering the scalarFunctionInvocation production.
	EnterScalarFunctionInvocation(c *ScalarFunctionInvocationContext)

	// EnterAggregateFunctionInvocation is called when entering the aggregateFunctionInvocation production.
	EnterAggregateFunctionInvocation(c *AggregateFunctionInvocationContext)

	// EnterRegressionFunctionInvocation is called when entering the regressionFunctionInvocation production.
	EnterRegressionFunctionInvocation(c *RegressionFunctionInvocationContext)

	// EnterExternalFunctionInvocation is called when entering the externalFunctionInvocation production.
	EnterExternalFunctionInvocation(c *ExternalFunctionInvocationContext)

	// EnterLabeledDuration is called when entering the labeledDuration production.
	EnterLabeledDuration(c *LabeledDurationContext)

	// EnterDurationSuffix is called when entering the durationSuffix production.
	EnterDurationSuffix(c *DurationSuffixContext)

	// EnterXmlCastSpecification is called when entering the xmlCastSpecification production.
	EnterXmlCastSpecification(c *XmlCastSpecificationContext)

	// EnterXmlParseSpecification is called when entering the xmlParseSpecification production.
	EnterXmlParseSpecification(c *XmlParseSpecificationContext)

	// EnterArrayElementSpecification is called when entering the arrayElementSpecification production.
	EnterArrayElementSpecification(c *ArrayElementSpecificationContext)

	// EnterArrayIndex is called when entering the arrayIndex production.
	EnterArrayIndex(c *ArrayIndexContext)

	// EnterArrayConstructor is called when entering the arrayConstructor production.
	EnterArrayConstructor(c *ArrayConstructorContext)

	// EnterOlapSpecification is called when entering the olapSpecification production.
	EnterOlapSpecification(c *OlapSpecificationContext)

	// EnterOrderedOlapSpecification is called when entering the orderedOlapSpecification production.
	EnterOrderedOlapSpecification(c *OrderedOlapSpecificationContext)

	// EnterOlapSpecificationFunction is called when entering the olapSpecificationFunction production.
	EnterOlapSpecificationFunction(c *OlapSpecificationFunctionContext)

	// EnterLagFunction is called when entering the lagFunction production.
	EnterLagFunction(c *LagFunctionContext)

	// EnterLeadFunction is called when entering the leadFunction production.
	EnterLeadFunction(c *LeadFunctionContext)

	// EnterRespectNullsClause is called when entering the respectNullsClause production.
	EnterRespectNullsClause(c *RespectNullsClauseContext)

	// EnterWindowPartitionClause is called when entering the windowPartitionClause production.
	EnterWindowPartitionClause(c *WindowPartitionClauseContext)

	// EnterWindowOrderClause is called when entering the windowOrderClause production.
	EnterWindowOrderClause(c *WindowOrderClauseContext)

	// EnterWindowOrderClauseQualifier is called when entering the windowOrderClauseQualifier production.
	EnterWindowOrderClauseQualifier(c *WindowOrderClauseQualifierContext)

	// EnterNumberingSpecification is called when entering the numberingSpecification production.
	EnterNumberingSpecification(c *NumberingSpecificationContext)

	// EnterAggregationSpecification is called when entering the aggregationSpecification production.
	EnterAggregationSpecification(c *AggregationSpecificationContext)

	// EnterAggregateFunction is called when entering the aggregateFunction production.
	EnterAggregateFunction(c *AggregateFunctionContext)

	// EnterRegressionFunction is called when entering the regressionFunction production.
	EnterRegressionFunction(c *RegressionFunctionContext)

	// EnterOlapColumnFunction is called when entering the olapColumnFunction production.
	EnterOlapColumnFunction(c *OlapColumnFunctionContext)

	// EnterFirstValueFunction is called when entering the firstValueFunction production.
	EnterFirstValueFunction(c *FirstValueFunctionContext)

	// EnterLastValueFunction is called when entering the lastValueFunction production.
	EnterLastValueFunction(c *LastValueFunctionContext)

	// EnterNthValueFunction is called when entering the nthValueFunction production.
	EnterNthValueFunction(c *NthValueFunctionContext)

	// EnterRatioToReportFunction is called when entering the ratioToReportFunction production.
	EnterRatioToReportFunction(c *RatioToReportFunctionContext)

	// EnterListaggFunction is called when entering the listaggFunction production.
	EnterListaggFunction(c *ListaggFunctionContext)

	// EnterArrayaggFunction is called when entering the arrayaggFunction production.
	EnterArrayaggFunction(c *ArrayaggFunctionContext)

	// EnterArrayaggOrdinaryFunction is called when entering the arrayaggOrdinaryFunction production.
	EnterArrayaggOrdinaryFunction(c *ArrayaggOrdinaryFunctionContext)

	// EnterArrayaggAssociativeFunction is called when entering the arrayaggAssociativeFunction production.
	EnterArrayaggAssociativeFunction(c *ArrayaggAssociativeFunctionContext)

	// EnterCorrelationFunction is called when entering the correlationFunction production.
	EnterCorrelationFunction(c *CorrelationFunctionContext)

	// EnterCovarianceFunction is called when entering the covarianceFunction production.
	EnterCovarianceFunction(c *CovarianceFunctionContext)

	// EnterCovarianceSampFunction is called when entering the covarianceSampFunction production.
	EnterCovarianceSampFunction(c *CovarianceSampFunctionContext)

	// EnterCumeDistFunction is called when entering the cumeDistFunction production.
	EnterCumeDistFunction(c *CumeDistFunctionContext)

	// EnterPercentileContFunction is called when entering the percentileContFunction production.
	EnterPercentileContFunction(c *PercentileContFunctionContext)

	// EnterPercentileDiscFunction is called when entering the percentileDiscFunction production.
	EnterPercentileDiscFunction(c *PercentileDiscFunctionContext)

	// EnterPercentRankFunction is called when entering the percentRankFunction production.
	EnterPercentRankFunction(c *PercentRankFunctionContext)

	// EnterXmlaggFunction is called when entering the xmlaggFunction production.
	EnterXmlaggFunction(c *XmlaggFunctionContext)

	// EnterXmlaggOrderByClause is called when entering the xmlaggOrderByClause production.
	EnterXmlaggOrderByClause(c *XmlaggOrderByClauseContext)

	// EnterXmlaggOrderByOption is called when entering the xmlaggOrderByOption production.
	EnterXmlaggOrderByOption(c *XmlaggOrderByOptionContext)

	// EnterAggregateOrderByClause is called when entering the aggregateOrderByClause production.
	EnterAggregateOrderByClause(c *AggregateOrderByClauseContext)

	// EnterAggregateOrderByOption is called when entering the aggregateOrderByOption production.
	EnterAggregateOrderByOption(c *AggregateOrderByOptionContext)

	// EnterWindowAggregationGroupClause is called when entering the windowAggregationGroupClause production.
	EnterWindowAggregationGroupClause(c *WindowAggregationGroupClauseContext)

	// EnterGroupStart is called when entering the groupStart production.
	EnterGroupStart(c *GroupStartContext)

	// EnterGroupBetween is called when entering the groupBetween production.
	EnterGroupBetween(c *GroupBetweenContext)

	// EnterGroupEnd is called when entering the groupEnd production.
	EnterGroupEnd(c *GroupEndContext)

	// EnterGroupBound1 is called when entering the groupBound1 production.
	EnterGroupBound1(c *GroupBound1Context)

	// EnterGroupBound2 is called when entering the groupBound2 production.
	EnterGroupBound2(c *GroupBound2Context)

	// EnterUnboundedPreceding is called when entering the unboundedPreceding production.
	EnterUnboundedPreceding(c *UnboundedPrecedingContext)

	// EnterUnboundedFollowing is called when entering the unboundedFollowing production.
	EnterUnboundedFollowing(c *UnboundedFollowingContext)

	// EnterBoundedPreceding is called when entering the boundedPreceding production.
	EnterBoundedPreceding(c *BoundedPrecedingContext)

	// EnterBoundedFollowing is called when entering the boundedFollowing production.
	EnterBoundedFollowing(c *BoundedFollowingContext)

	// EnterCurrentRow is called when entering the currentRow production.
	EnterCurrentRow(c *CurrentRowContext)

	// EnterScalarFunction is called when entering the scalarFunction production.
	EnterScalarFunction(c *ScalarFunctionContext)

	// EnterTableFunction is called when entering the tableFunction production.
	EnterTableFunction(c *TableFunctionContext)

	// EnterSpecialRegister is called when entering the specialRegister production.
	EnterSpecialRegister(c *SpecialRegisterContext)

	// EnterAiAnalogyFunction is called when entering the aiAnalogyFunction production.
	EnterAiAnalogyFunction(c *AiAnalogyFunctionContext)

	// EnterAiFunctionExpression is called when entering the aiFunctionExpression production.
	EnterAiFunctionExpression(c *AiFunctionExpressionContext)

	// EnterAiAnalogyFunctionSource is called when entering the aiAnalogyFunctionSource production.
	EnterAiAnalogyFunctionSource(c *AiAnalogyFunctionSourceContext)

	// EnterAiAnalogyFunctionTarget is called when entering the aiAnalogyFunctionTarget production.
	EnterAiAnalogyFunctionTarget(c *AiAnalogyFunctionTargetContext)

	// EnterAiAnalogyFunctionSource1 is called when entering the aiAnalogyFunctionSource1 production.
	EnterAiAnalogyFunctionSource1(c *AiAnalogyFunctionSource1Context)

	// EnterAiAnalogyFunctionSource2 is called when entering the aiAnalogyFunctionSource2 production.
	EnterAiAnalogyFunctionSource2(c *AiAnalogyFunctionSource2Context)

	// EnterAiAnalogyFunctionTarget1 is called when entering the aiAnalogyFunctionTarget1 production.
	EnterAiAnalogyFunctionTarget1(c *AiAnalogyFunctionTarget1Context)

	// EnterAiAnalogyFunctionTarget2 is called when entering the aiAnalogyFunctionTarget2 production.
	EnterAiAnalogyFunctionTarget2(c *AiAnalogyFunctionTarget2Context)

	// EnterAiSemanticClusterFunction is called when entering the aiSemanticClusterFunction production.
	EnterAiSemanticClusterFunction(c *AiSemanticClusterFunctionContext)

	// EnterAiSemanticClusterMemberExpression is called when entering the aiSemanticClusterMemberExpression production.
	EnterAiSemanticClusterMemberExpression(c *AiSemanticClusterMemberExpressionContext)

	// EnterAiSemanticClusterClusteringExpression is called when entering the aiSemanticClusterClusteringExpression production.
	EnterAiSemanticClusterClusteringExpression(c *AiSemanticClusterClusteringExpressionContext)

	// EnterAiSimilarityFunction is called when entering the aiSimilarityFunction production.
	EnterAiSimilarityFunction(c *AiSimilarityFunctionContext)

	// EnterAiSimilarityExpression is called when entering the aiSimilarityExpression production.
	EnterAiSimilarityExpression(c *AiSimilarityExpressionContext)

	// EnterAiSimilarityExpression1 is called when entering the aiSimilarityExpression1 production.
	EnterAiSimilarityExpression1(c *AiSimilarityExpression1Context)

	// EnterAiSimilarityExpression2 is called when entering the aiSimilarityExpression2 production.
	EnterAiSimilarityExpression2(c *AiSimilarityExpression2Context)

	// EnterXmlelementFunction is called when entering the xmlelementFunction production.
	EnterXmlelementFunction(c *XmlelementFunctionContext)

	// EnterXmlforestFunction is called when entering the xmlforestFunction production.
	EnterXmlforestFunction(c *XmlforestFunctionContext)

	// EnterXmlmodifyFunction is called when entering the xmlmodifyFunction production.
	EnterXmlmodifyFunction(c *XmlmodifyFunctionContext)

	// EnterXmlpiFunction is called when entering the xmlpiFunction production.
	EnterXmlpiFunction(c *XmlpiFunctionContext)

	// EnterXmlqueryFunction is called when entering the xmlqueryFunction production.
	EnterXmlqueryFunction(c *XmlqueryFunctionContext)

	// EnterXmlattributesFunction is called when entering the xmlattributesFunction production.
	EnterXmlattributesFunction(c *XmlattributesFunctionContext)

	// EnterXmlserializeFunction is called when entering the xmlserializeFunction production.
	EnterXmlserializeFunction(c *XmlserializeFunctionContext)

	// EnterXmlnamespaceFunction is called when entering the xmlnamespaceFunction production.
	EnterXmlnamespaceFunction(c *XmlnamespaceFunctionContext)

	// EnterXmlnamespaceOption is called when entering the xmlnamespaceOption production.
	EnterXmlnamespaceOption(c *XmlnamespaceOptionContext)

	// EnterXmlserializeFunctionOptions is called when entering the xmlserializeFunctionOptions production.
	EnterXmlserializeFunctionOptions(c *XmlserializeFunctionOptionsContext)

	// EnterXmlFunctionOptionClause is called when entering the xmlFunctionOptionClause production.
	EnterXmlFunctionOptionClause(c *XmlFunctionOptionClauseContext)

	// EnterXmlFunctionOption is called when entering the xmlFunctionOption production.
	EnterXmlFunctionOption(c *XmlFunctionOptionContext)

	// EnterElementContentExpression is called when entering the elementContentExpression production.
	EnterElementContentExpression(c *ElementContentExpressionContext)

	// EnterXqueryExpressionConstant is called when entering the xqueryExpressionConstant production.
	EnterXqueryExpressionConstant(c *XqueryExpressionConstantContext)

	// EnterXqueryArgument is called when entering the xqueryArgument production.
	EnterXqueryArgument(c *XqueryArgumentContext)

	// EnterXmltableFunctionSpecification is called when entering the xmltableFunctionSpecification production.
	EnterXmltableFunctionSpecification(c *XmltableFunctionSpecificationContext)

	// EnterRowXqueryExpressionConstant is called when entering the rowXqueryExpressionConstant production.
	EnterRowXqueryExpressionConstant(c *RowXqueryExpressionConstantContext)

	// EnterRowXqueryArgument is called when entering the rowXqueryArgument production.
	EnterRowXqueryArgument(c *RowXqueryArgumentContext)

	// EnterXqueryContextItemExpression is called when entering the xqueryContextItemExpression production.
	EnterXqueryContextItemExpression(c *XqueryContextItemExpressionContext)

	// EnterXqueryVariableExpression is called when entering the xqueryVariableExpression production.
	EnterXqueryVariableExpression(c *XqueryVariableExpressionContext)

	// EnterXmlTableRegularColumnDefinition is called when entering the xmlTableRegularColumnDefinition production.
	EnterXmlTableRegularColumnDefinition(c *XmlTableRegularColumnDefinitionContext)

	// EnterDefaultClause is called when entering the defaultClause production.
	EnterDefaultClause(c *DefaultClauseContext)

	// EnterDefaultClause1 is called when entering the defaultClause1 production.
	EnterDefaultClause1(c *DefaultClause1Context)

	// EnterDefaultClauseAllowables is called when entering the defaultClauseAllowables production.
	EnterDefaultClauseAllowables(c *DefaultClauseAllowablesContext)

	// EnterDistinctTypeCastFunctionName is called when entering the distinctTypeCastFunctionName production.
	EnterDistinctTypeCastFunctionName(c *DistinctTypeCastFunctionNameContext)

	// EnterColumnXqueryExpressionConstant is called when entering the columnXqueryExpressionConstant production.
	EnterColumnXqueryExpressionConstant(c *ColumnXqueryExpressionConstantContext)

	// EnterXmlTableOrdinalityColumnDefinition is called when entering the xmlTableOrdinalityColumnDefinition production.
	EnterXmlTableOrdinalityColumnDefinition(c *XmlTableOrdinalityColumnDefinitionContext)

	// EnterXmlnamespacesDeclaration is called when entering the xmlnamespacesDeclaration production.
	EnterXmlnamespacesDeclaration(c *XmlnamespacesDeclarationContext)

	// EnterXmlnamespacesFunctionSpecification is called when entering the xmlnamespacesFunctionSpecification production.
	EnterXmlnamespacesFunctionSpecification(c *XmlnamespacesFunctionSpecificationContext)

	// EnterXmlnamespacesFunctionArguments is called when entering the xmlnamespacesFunctionArguments production.
	EnterXmlnamespacesFunctionArguments(c *XmlnamespacesFunctionArgumentsContext)

	// EnterNamespaceUri is called when entering the namespaceUri production.
	EnterNamespaceUri(c *NamespaceUriContext)

	// EnterNamespacePrefix is called when entering the namespacePrefix production.
	EnterNamespacePrefix(c *NamespacePrefixContext)

	// EnterTimeZoneSpecificExpression is called when entering the timeZoneSpecificExpression production.
	EnterTimeZoneSpecificExpression(c *TimeZoneSpecificExpressionContext)

	// EnterTimeZoneExpressionSubset is called when entering the timeZoneExpressionSubset production.
	EnterTimeZoneExpressionSubset(c *TimeZoneExpressionSubsetContext)

	// EnterCaseExpression is called when entering the caseExpression production.
	EnterCaseExpression(c *CaseExpressionContext)

	// EnterResultExpression is called when entering the resultExpression production.
	EnterResultExpression(c *ResultExpressionContext)

	// EnterSearchedWhenClause is called when entering the searchedWhenClause production.
	EnterSearchedWhenClause(c *SearchedWhenClauseContext)

	// EnterSimpleWhenClause is called when entering the simpleWhenClause production.
	EnterSimpleWhenClause(c *SimpleWhenClauseContext)

	// EnterSearchCondition is called when entering the searchCondition production.
	EnterSearchCondition(c *SearchConditionContext)

	// EnterCheckCondition is called when entering the checkCondition production.
	EnterCheckCondition(c *CheckConditionContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterBasicPredicate is called when entering the basicPredicate production.
	EnterBasicPredicate(c *BasicPredicateContext)

	// EnterRowValueExpression is called when entering the rowValueExpression production.
	EnterRowValueExpression(c *RowValueExpressionContext)

	// EnterQuantifiedPredicate is called when entering the quantifiedPredicate production.
	EnterQuantifiedPredicate(c *QuantifiedPredicateContext)

	// EnterArrayExistsPredicate is called when entering the arrayExistsPredicate production.
	EnterArrayExistsPredicate(c *ArrayExistsPredicateContext)

	// EnterBetweenPredicate is called when entering the betweenPredicate production.
	EnterBetweenPredicate(c *BetweenPredicateContext)

	// EnterDistinctPredicate is called when entering the distinctPredicate production.
	EnterDistinctPredicate(c *DistinctPredicateContext)

	// EnterExistsPredicate is called when entering the existsPredicate production.
	EnterExistsPredicate(c *ExistsPredicateContext)

	// EnterInPredicate is called when entering the inPredicate production.
	EnterInPredicate(c *InPredicateContext)

	// EnterLikePredicate is called when entering the likePredicate production.
	EnterLikePredicate(c *LikePredicateContext)

	// EnterNullPredicate is called when entering the nullPredicate production.
	EnterNullPredicate(c *NullPredicateContext)

	// EnterXmlExistsPredicate is called when entering the xmlExistsPredicate production.
	EnterXmlExistsPredicate(c *XmlExistsPredicateContext)

	// EnterArrayExpression is called when entering the arrayExpression production.
	EnterArrayExpression(c *ArrayExpressionContext)

	// EnterCastSpecification is called when entering the castSpecification production.
	EnterCastSpecification(c *CastSpecificationContext)

	// EnterParameterMarker is called when entering the parameterMarker production.
	EnterParameterMarker(c *ParameterMarkerContext)

	// EnterCastDataType is called when entering the castDataType production.
	EnterCastDataType(c *CastDataTypeContext)

	// EnterCastBuiltInType is called when entering the castBuiltInType production.
	EnterCastBuiltInType(c *CastBuiltInTypeContext)

	// EnterIntegerInParens is called when entering the integerInParens production.
	EnterIntegerInParens(c *IntegerInParensContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// EnterCcsidQualifier is called when entering the ccsidQualifier production.
	EnterCcsidQualifier(c *CcsidQualifierContext)

	// EnterForDataQualifier is called when entering the forDataQualifier production.
	EnterForDataQualifier(c *ForDataQualifierContext)

	// EnterDistinctTypeName is called when entering the distinctTypeName production.
	EnterDistinctTypeName(c *DistinctTypeNameContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterCcsidValue is called when entering the ccsidValue production.
	EnterCcsidValue(c *CcsidValueContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterSourceColumnName is called when entering the sourceColumnName production.
	EnterSourceColumnName(c *SourceColumnNameContext)

	// EnterTargetColumnName is called when entering the targetColumnName production.
	EnterTargetColumnName(c *TargetColumnNameContext)

	// EnterBuildColumnName is called when entering the buildColumnName production.
	EnterBuildColumnName(c *BuildColumnNameContext)

	// EnterBeginColumnName is called when entering the beginColumnName production.
	EnterBeginColumnName(c *BeginColumnNameContext)

	// EnterEndColumnName is called when entering the endColumnName production.
	EnterEndColumnName(c *EndColumnNameContext)

	// EnterCorrelationName is called when entering the correlationName production.
	EnterCorrelationName(c *CorrelationNameContext)

	// EnterLocationName is called when entering the locationName production.
	EnterLocationName(c *LocationNameContext)

	// EnterSchemaName is called when entering the schemaName production.
	EnterSchemaName(c *SchemaNameContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterAlterTableName is called when entering the alterTableName production.
	EnterAlterTableName(c *AlterTableNameContext)

	// EnterAuxTableName is called when entering the auxTableName production.
	EnterAuxTableName(c *AuxTableNameContext)

	// EnterHistoryTableName is called when entering the historyTableName production.
	EnterHistoryTableName(c *HistoryTableNameContext)

	// EnterCloneTableName is called when entering the cloneTableName production.
	EnterCloneTableName(c *CloneTableNameContext)

	// EnterArchiveTableName is called when entering the archiveTableName production.
	EnterArchiveTableName(c *ArchiveTableNameContext)

	// EnterViewName is called when entering the viewName production.
	EnterViewName(c *ViewNameContext)

	// EnterProgramName is called when entering the programName production.
	EnterProgramName(c *ProgramNameContext)

	// EnterPackageName is called when entering the packageName production.
	EnterPackageName(c *PackageNameContext)

	// EnterPlanName is called when entering the planName production.
	EnterPlanName(c *PlanNameContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterVariableName is called when entering the variableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterArrayTypeName is called when entering the arrayTypeName production.
	EnterArrayTypeName(c *ArrayTypeNameContext)

	// EnterJarName is called when entering the jarName production.
	EnterJarName(c *JarNameContext)

	// EnterSavepointName is called when entering the savepointName production.
	EnterSavepointName(c *SavepointNameContext)

	// EnterAliasName is called when entering the aliasName production.
	EnterAliasName(c *AliasNameContext)

	// EnterConstraintName is called when entering the constraintName production.
	EnterConstraintName(c *ConstraintNameContext)

	// EnterRoutineVersionID is called when entering the routineVersionID production.
	EnterRoutineVersionID(c *RoutineVersionIDContext)

	// EnterVersionID is called when entering the versionID production.
	EnterVersionID(c *VersionIDContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterMaskName is called when entering the maskName production.
	EnterMaskName(c *MaskNameContext)

	// EnterPermissionName is called when entering the permissionName production.
	EnterPermissionName(c *PermissionNameContext)

	// EnterProcedureName is called when entering the procedureName production.
	EnterProcedureName(c *ProcedureNameContext)

	// EnterSequenceName is called when entering the sequenceName production.
	EnterSequenceName(c *SequenceNameContext)

	// EnterMemberName is called when entering the memberName production.
	EnterMemberName(c *MemberNameContext)

	// EnterDatabaseName is called when entering the databaseName production.
	EnterDatabaseName(c *DatabaseNameContext)

	// EnterTablespaceName is called when entering the tablespaceName production.
	EnterTablespaceName(c *TablespaceNameContext)

	// EnterAcceleratorName is called when entering the acceleratorName production.
	EnterAcceleratorName(c *AcceleratorNameContext)

	// EnterCatalogName is called when entering the catalogName production.
	EnterCatalogName(c *CatalogNameContext)

	// EnterTriggerName is called when entering the triggerName production.
	EnterTriggerName(c *TriggerNameContext)

	// EnterContextName is called when entering the contextName production.
	EnterContextName(c *ContextNameContext)

	// EnterAuthorizationName is called when entering the authorizationName production.
	EnterAuthorizationName(c *AuthorizationNameContext)

	// EnterProfileName is called when entering the profileName production.
	EnterProfileName(c *ProfileNameContext)

	// EnterRoleName is called when entering the roleName production.
	EnterRoleName(c *RoleNameContext)

	// EnterSeclabelName is called when entering the seclabelName production.
	EnterSeclabelName(c *SeclabelNameContext)

	// EnterParameterName is called when entering the parameterName production.
	EnterParameterName(c *ParameterNameContext)

	// EnterAddressValue is called when entering the addressValue production.
	EnterAddressValue(c *AddressValueContext)

	// EnterJobnameValue is called when entering the jobnameValue production.
	EnterJobnameValue(c *JobnameValueContext)

	// EnterServauthValue is called when entering the servauthValue production.
	EnterServauthValue(c *ServauthValueContext)

	// EnterEncryptionValue is called when entering the encryptionValue production.
	EnterEncryptionValue(c *EncryptionValueContext)

	// EnterBpName is called when entering the bpName production.
	EnterBpName(c *BpNameContext)

	// EnterStogroupName is called when entering the stogroupName production.
	EnterStogroupName(c *StogroupNameContext)

	// EnterDcName is called when entering the dcName production.
	EnterDcName(c *DcNameContext)

	// EnterMcName is called when entering the mcName production.
	EnterMcName(c *McNameContext)

	// EnterScName is called when entering the scName production.
	EnterScName(c *ScNameContext)

	// EnterVolumeID is called when entering the volumeID production.
	EnterVolumeID(c *VolumeIDContext)

	// EnterKeyLabelName is called when entering the keyLabelName production.
	EnterKeyLabelName(c *KeyLabelNameContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterSpecificName is called when entering the specificName production.
	EnterSpecificName(c *SpecificNameContext)

	// EnterHostLabel is called when entering the hostLabel production.
	EnterHostLabel(c *HostLabelContext)

	// EnterHostVariable is called when entering the hostVariable production.
	EnterHostVariable(c *HostVariableContext)

	// EnterHostIdentifier is called when entering the hostIdentifier production.
	EnterHostIdentifier(c *HostIdentifierContext)

	// EnterHostStructure is called when entering the hostStructure production.
	EnterHostStructure(c *HostStructureContext)

	// EnterNullIndicator is called when entering the nullIndicator production.
	EnterNullIndicator(c *NullIndicatorContext)

	// EnterNullIndicatorStructure is called when entering the nullIndicatorStructure production.
	EnterNullIndicatorStructure(c *NullIndicatorStructureContext)

	// EnterGlobalVariableName is called when entering the globalVariableName production.
	EnterGlobalVariableName(c *GlobalVariableNameContext)

	// EnterSqlParameterName is called when entering the sqlParameterName production.
	EnterSqlParameterName(c *SqlParameterNameContext)

	// EnterSqlVariableName is called when entering the sqlVariableName production.
	EnterSqlVariableName(c *SqlVariableNameContext)

	// EnterTransitionVariableName is called when entering the transitionVariableName production.
	EnterTransitionVariableName(c *TransitionVariableNameContext)

	// EnterSynonym is called when entering the synonym production.
	EnterSynonym(c *SynonymContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterIntoClause is called when entering the intoClause production.
	EnterIntoClause(c *IntoClauseContext)

	// EnterCorrelationClause is called when entering the correlationClause production.
	EnterCorrelationClause(c *CorrelationClauseContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterTableReference is called when entering the tableReference production.
	EnterTableReference(c *TableReferenceContext)

	// EnterSingleTableReference is called when entering the singleTableReference production.
	EnterSingleTableReference(c *SingleTableReferenceContext)

	// EnterPeriodSpecification is called when entering the periodSpecification production.
	EnterPeriodSpecification(c *PeriodSpecificationContext)

	// EnterPeriodClause is called when entering the periodClause production.
	EnterPeriodClause(c *PeriodClauseContext)

	// EnterNestedTableExpression is called when entering the nestedTableExpression production.
	EnterNestedTableExpression(c *NestedTableExpressionContext)

	// EnterDataChangeTableReference is called when entering the dataChangeTableReference production.
	EnterDataChangeTableReference(c *DataChangeTableReferenceContext)

	// EnterTableFunctionReference is called when entering the tableFunctionReference production.
	EnterTableFunctionReference(c *TableFunctionReferenceContext)

	// EnterTableUdfCardinalityClause is called when entering the tableUdfCardinalityClause production.
	EnterTableUdfCardinalityClause(c *TableUdfCardinalityClauseContext)

	// EnterTypedCorrelationClause is called when entering the typedCorrelationClause production.
	EnterTypedCorrelationClause(c *TypedCorrelationClauseContext)

	// EnterTableLocatorReference is called when entering the tableLocatorReference production.
	EnterTableLocatorReference(c *TableLocatorReferenceContext)

	// EnterXmltableExpression is called when entering the xmltableExpression production.
	EnterXmltableExpression(c *XmltableExpressionContext)

	// EnterCollectionDerivedTable is called when entering the collectionDerivedTable production.
	EnterCollectionDerivedTable(c *CollectionDerivedTableContext)

	// EnterJoinCondition is called when entering the joinCondition production.
	EnterJoinCondition(c *JoinConditionContext)

	// EnterFullJoinExpression is called when entering the fullJoinExpression production.
	EnterFullJoinExpression(c *FullJoinExpressionContext)

	// EnterCastFunction is called when entering the castFunction production.
	EnterCastFunction(c *CastFunctionContext)

	// EnterOrdinaryArrayExpression is called when entering the ordinaryArrayExpression production.
	EnterOrdinaryArrayExpression(c *OrdinaryArrayExpressionContext)

	// EnterAssociativeArrayExpression is called when entering the associativeArrayExpression production.
	EnterAssociativeArrayExpression(c *AssociativeArrayExpressionContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterGroupingExpression is called when entering the groupingExpression production.
	EnterGroupingExpression(c *GroupingExpressionContext)

	// EnterGroupingSets is called when entering the groupingSets production.
	EnterGroupingSets(c *GroupingSetsContext)

	// EnterGroupingSetsGroup is called when entering the groupingSetsGroup production.
	EnterGroupingSetsGroup(c *GroupingSetsGroupContext)

	// EnterSuperGroups is called when entering the superGroups production.
	EnterSuperGroups(c *SuperGroupsContext)

	// EnterSelectColumns is called when entering the selectColumns production.
	EnterSelectColumns(c *SelectColumnsContext)

	// EnterUnpackedRow is called when entering the unpackedRow production.
	EnterUnpackedRow(c *UnpackedRowContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterSubSelect is called when entering the subSelect production.
	EnterSubSelect(c *SubSelectContext)

	// EnterSelectIntoStatement is called when entering the selectIntoStatement production.
	EnterSelectIntoStatement(c *SelectIntoStatementContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterCommonTableExpression is called when entering the commonTableExpression production.
	EnterCommonTableExpression(c *CommonTableExpressionContext)

	// EnterUpdateClause is called when entering the updateClause production.
	EnterUpdateClause(c *UpdateClauseContext)

	// EnterReadOnlyClause is called when entering the readOnlyClause production.
	EnterReadOnlyClause(c *ReadOnlyClauseContext)

	// EnterOptimizeClause is called when entering the optimizeClause production.
	EnterOptimizeClause(c *OptimizeClauseContext)

	// EnterIsolationClause is called when entering the isolationClause production.
	EnterIsolationClause(c *IsolationClauseContext)

	// EnterLockClause is called when entering the lockClause production.
	EnterLockClause(c *LockClauseContext)

	// EnterSkipLockedDataClause is called when entering the skipLockedDataClause production.
	EnterSkipLockedDataClause(c *SkipLockedDataClauseContext)

	// EnterQuerynoClause is called when entering the querynoClause production.
	EnterQuerynoClause(c *QuerynoClauseContext)

	// EnterScalarFullSelect is called when entering the scalarFullSelect production.
	EnterScalarFullSelect(c *ScalarFullSelectContext)

	// EnterFullSelect is called when entering the fullSelect production.
	EnterFullSelect(c *FullSelectContext)

	// EnterValuesClause is called when entering the valuesClause production.
	EnterValuesClause(c *ValuesClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterSortKey is called when entering the sortKey production.
	EnterSortKey(c *SortKeyContext)

	// EnterOffsetClause is called when entering the offsetClause production.
	EnterOffsetClause(c *OffsetClauseContext)

	// EnterFetchClause is called when entering the fetchClause production.
	EnterFetchClause(c *FetchClauseContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifier1 is called when entering the identifier1 production.
	EnterIdentifier1(c *Identifier1Context)

	// EnterSqlidentifier is called when entering the sqlidentifier production.
	EnterSqlidentifier(c *SqlidentifierContext)

	// EnterSqlKeyword is called when entering the sqlKeyword production.
	EnterSqlKeyword(c *SqlKeywordContext)

	// ExitStartRule is called when exiting the startRule production.
	ExitStartRule(c *StartRuleContext)

	// ExitSqlStatement is called when exiting the sqlStatement production.
	ExitSqlStatement(c *SqlStatementContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitCursorName is called when exiting the cursorName production.
	ExitCursorName(c *CursorNameContext)

	// ExitStatementName is called when exiting the statementName production.
	ExitStatementName(c *StatementNameContext)

	// ExitDescriptorName is called when exiting the descriptorName production.
	ExitDescriptorName(c *DescriptorNameContext)

	// ExitHoldability is called when exiting the holdability production.
	ExitHoldability(c *HoldabilityContext)

	// ExitReturnability is called when exiting the returnability production.
	ExitReturnability(c *ReturnabilityContext)

	// ExitRowsetPositioning is called when exiting the rowsetPositioning production.
	ExitRowsetPositioning(c *RowsetPositioningContext)

	// ExitNotNullPhrase is called when exiting the notNullPhrase production.
	ExitNotNullPhrase(c *NotNullPhraseContext)

	// ExitAllocateCursorStatement is called when exiting the allocateCursorStatement production.
	ExitAllocateCursorStatement(c *AllocateCursorStatementContext)

	// ExitRsLocatorVariable is called when exiting the rsLocatorVariable production.
	ExitRsLocatorVariable(c *RsLocatorVariableContext)

	// ExitAlterDatabaseStatement is called when exiting the alterDatabaseStatement production.
	ExitAlterDatabaseStatement(c *AlterDatabaseStatementContext)

	// ExitAlterFunctionStatement is called when exiting the alterFunctionStatement production.
	ExitAlterFunctionStatement(c *AlterFunctionStatementContext)

	// ExitAlterFunctionStatementExternal is called when exiting the alterFunctionStatementExternal production.
	ExitAlterFunctionStatementExternal(c *AlterFunctionStatementExternalContext)

	// ExitAlterFunctionStatementCompiledSqlScalar is called when exiting the alterFunctionStatementCompiledSqlScalar production.
	ExitAlterFunctionStatementCompiledSqlScalar(c *AlterFunctionStatementCompiledSqlScalarContext)

	// ExitAlterWhichFunction1 is called when exiting the alterWhichFunction1 production.
	ExitAlterWhichFunction1(c *AlterWhichFunction1Context)

	// ExitAlterWhichFunction2 is called when exiting the alterWhichFunction2 production.
	ExitAlterWhichFunction2(c *AlterWhichFunction2Context)

	// ExitFunctionCompiledSqlScalarRoutineSpecification is called when exiting the functionCompiledSqlScalarRoutineSpecification production.
	ExitFunctionCompiledSqlScalarRoutineSpecification(c *FunctionCompiledSqlScalarRoutineSpecificationContext)

	// ExitAlterFunctionStatementInlineSqlScalar is called when exiting the alterFunctionStatementInlineSqlScalar production.
	ExitAlterFunctionStatementInlineSqlScalar(c *AlterFunctionStatementInlineSqlScalarContext)

	// ExitAlterFunctionStatementSqlTable is called when exiting the alterFunctionStatementSqlTable production.
	ExitAlterFunctionStatementSqlTable(c *AlterFunctionStatementSqlTableContext)

	// ExitFunctionReturnsClause is called when exiting the functionReturnsClause production.
	ExitFunctionReturnsClause(c *FunctionReturnsClauseContext)

	// ExitFunctionDesignator1 is called when exiting the functionDesignator1 production.
	ExitFunctionDesignator1(c *FunctionDesignator1Context)

	// ExitFunctionDesignator2 is called when exiting the functionDesignator2 production.
	ExitFunctionDesignator2(c *FunctionDesignator2Context)

	// ExitAlterIndexStatement is called when exiting the alterIndexStatement production.
	ExitAlterIndexStatement(c *AlterIndexStatementContext)

	// ExitAlterMaskStatement is called when exiting the alterMaskStatement production.
	ExitAlterMaskStatement(c *AlterMaskStatementContext)

	// ExitAlterPermissionStatement is called when exiting the alterPermissionStatement production.
	ExitAlterPermissionStatement(c *AlterPermissionStatementContext)

	// ExitAlterProcedureStatement is called when exiting the alterProcedureStatement production.
	ExitAlterProcedureStatement(c *AlterProcedureStatementContext)

	// ExitAlterProcedureSQLPLStatement is called when exiting the alterProcedureSQLPLStatement production.
	ExitAlterProcedureSQLPLStatement(c *AlterProcedureSQLPLStatementContext)

	// ExitAlterWhichProcedureSQLPL1 is called when exiting the alterWhichProcedureSQLPL1 production.
	ExitAlterWhichProcedureSQLPL1(c *AlterWhichProcedureSQLPL1Context)

	// ExitAlterWhichProcedureSQLPL2 is called when exiting the alterWhichProcedureSQLPL2 production.
	ExitAlterWhichProcedureSQLPL2(c *AlterWhichProcedureSQLPL2Context)

	// ExitApplicationCompatibilityPhrase is called when exiting the applicationCompatibilityPhrase production.
	ExitApplicationCompatibilityPhrase(c *ApplicationCompatibilityPhraseContext)

	// ExitAlterSequenceStatement is called when exiting the alterSequenceStatement production.
	ExitAlterSequenceStatement(c *AlterSequenceStatementContext)

	// ExitAlterStogroupStatement is called when exiting the alterStogroupStatement production.
	ExitAlterStogroupStatement(c *AlterStogroupStatementContext)

	// ExitAlterTableStatement is called when exiting the alterTableStatement production.
	ExitAlterTableStatement(c *AlterTableStatementContext)

	// ExitAlterTablespaceStatement is called when exiting the alterTablespaceStatement production.
	ExitAlterTablespaceStatement(c *AlterTablespaceStatementContext)

	// ExitAlterTriggerStatement is called when exiting the alterTriggerStatement production.
	ExitAlterTriggerStatement(c *AlterTriggerStatementContext)

	// ExitAlterTrustedContextStatement is called when exiting the alterTrustedContextStatement production.
	ExitAlterTrustedContextStatement(c *AlterTrustedContextStatementContext)

	// ExitAlterViewStatement is called when exiting the alterViewStatement production.
	ExitAlterViewStatement(c *AlterViewStatementContext)

	// ExitAssociateLocatorsStatement is called when exiting the associateLocatorsStatement production.
	ExitAssociateLocatorsStatement(c *AssociateLocatorsStatementContext)

	// ExitBeginDeclareSectionStatement is called when exiting the beginDeclareSectionStatement production.
	ExitBeginDeclareSectionStatement(c *BeginDeclareSectionStatementContext)

	// ExitCallStatement is called when exiting the callStatement production.
	ExitCallStatement(c *CallStatementContext)

	// ExitCloseStatement is called when exiting the closeStatement production.
	ExitCloseStatement(c *CloseStatementContext)

	// ExitCommentStatement is called when exiting the commentStatement production.
	ExitCommentStatement(c *CommentStatementContext)

	// ExitCommitStatement is called when exiting the commitStatement production.
	ExitCommitStatement(c *CommitStatementContext)

	// ExitConnectStatement is called when exiting the connectStatement production.
	ExitConnectStatement(c *ConnectStatementContext)

	// ExitCreateAliasStatement is called when exiting the createAliasStatement production.
	ExitCreateAliasStatement(c *CreateAliasStatementContext)

	// ExitCreateAuxiliaryTableStatement is called when exiting the createAuxiliaryTableStatement production.
	ExitCreateAuxiliaryTableStatement(c *CreateAuxiliaryTableStatementContext)

	// ExitCreateDatabaseStatement is called when exiting the createDatabaseStatement production.
	ExitCreateDatabaseStatement(c *CreateDatabaseStatementContext)

	// ExitCreateFunctionStatement is called when exiting the createFunctionStatement production.
	ExitCreateFunctionStatement(c *CreateFunctionStatementContext)

	// ExitCreateFunctionStatementExternalScalar is called when exiting the createFunctionStatementExternalScalar production.
	ExitCreateFunctionStatementExternalScalar(c *CreateFunctionStatementExternalScalarContext)

	// ExitCreateFunctionStatementExternalScalarReturnsPhrase is called when exiting the createFunctionStatementExternalScalarReturnsPhrase production.
	ExitCreateFunctionStatementExternalScalarReturnsPhrase(c *CreateFunctionStatementExternalScalarReturnsPhraseContext)

	// ExitCreateFunctionStatementExternalTable is called when exiting the createFunctionStatementExternalTable production.
	ExitCreateFunctionStatementExternalTable(c *CreateFunctionStatementExternalTableContext)

	// ExitCreateFunctionStatementExternalTableReturnsPhrase is called when exiting the createFunctionStatementExternalTableReturnsPhrase production.
	ExitCreateFunctionStatementExternalTableReturnsPhrase(c *CreateFunctionStatementExternalTableReturnsPhraseContext)

	// ExitExternalTableFunctionColumn is called when exiting the externalTableFunctionColumn production.
	ExitExternalTableFunctionColumn(c *ExternalTableFunctionColumnContext)

	// ExitCreateFunctionStatementSourced is called when exiting the createFunctionStatementSourced production.
	ExitCreateFunctionStatementSourced(c *CreateFunctionStatementSourcedContext)

	// ExitCreateFunctionStatementSourcedReturnsPhrase is called when exiting the createFunctionStatementSourcedReturnsPhrase production.
	ExitCreateFunctionStatementSourcedReturnsPhrase(c *CreateFunctionStatementSourcedReturnsPhraseContext)

	// ExitCreateFunctionStatementSourcedSourcePhrase is called when exiting the createFunctionStatementSourcedSourcePhrase production.
	ExitCreateFunctionStatementSourcedSourcePhrase(c *CreateFunctionStatementSourcedSourcePhraseContext)

	// ExitCreateFunctionStatementInlineSqlScalar is called when exiting the createFunctionStatementInlineSqlScalar production.
	ExitCreateFunctionStatementInlineSqlScalar(c *CreateFunctionStatementInlineSqlScalarContext)

	// ExitCreateFunctionStatementCompiledSqlScalar is called when exiting the createFunctionStatementCompiledSqlScalar production.
	ExitCreateFunctionStatementCompiledSqlScalar(c *CreateFunctionStatementCompiledSqlScalarContext)

	// ExitCreateFunctionStatementSqlTable is called when exiting the createFunctionStatementSqlTable production.
	ExitCreateFunctionStatementSqlTable(c *CreateFunctionStatementSqlTableContext)

	// ExitCreateGlobalTemporaryTableStatement is called when exiting the createGlobalTemporaryTableStatement production.
	ExitCreateGlobalTemporaryTableStatement(c *CreateGlobalTemporaryTableStatementContext)

	// ExitCreateIndexStatement is called when exiting the createIndexStatement production.
	ExitCreateIndexStatement(c *CreateIndexStatementContext)

	// ExitCreateLobTablespaceStatement is called when exiting the createLobTablespaceStatement production.
	ExitCreateLobTablespaceStatement(c *CreateLobTablespaceStatementContext)

	// ExitCreateMaskStatement is called when exiting the createMaskStatement production.
	ExitCreateMaskStatement(c *CreateMaskStatementContext)

	// ExitCreatePermissionStatement is called when exiting the createPermissionStatement production.
	ExitCreatePermissionStatement(c *CreatePermissionStatementContext)

	// ExitCreateProcedureSQLPLStatement is called when exiting the createProcedureSQLPLStatement production.
	ExitCreateProcedureSQLPLStatement(c *CreateProcedureSQLPLStatementContext)

	// ExitSqlRoutineBody is called when exiting the sqlRoutineBody production.
	ExitSqlRoutineBody(c *SqlRoutineBodyContext)

	// ExitObfuscatedStatementText is called when exiting the obfuscatedStatementText production.
	ExitObfuscatedStatementText(c *ObfuscatedStatementTextContext)

	// ExitProbablySQLPL is called when exiting the probablySQLPL production.
	ExitProbablySQLPL(c *ProbablySQLPLContext)

	// ExitCreateProcedureStatement is called when exiting the createProcedureStatement production.
	ExitCreateProcedureStatement(c *CreateProcedureStatementContext)

	// ExitCreateRoleStatement is called when exiting the createRoleStatement production.
	ExitCreateRoleStatement(c *CreateRoleStatementContext)

	// ExitCreateSequenceStatement is called when exiting the createSequenceStatement production.
	ExitCreateSequenceStatement(c *CreateSequenceStatementContext)

	// ExitCreateStogroupStatement is called when exiting the createStogroupStatement production.
	ExitCreateStogroupStatement(c *CreateStogroupStatementContext)

	// ExitCreateTableStatement is called when exiting the createTableStatement production.
	ExitCreateTableStatement(c *CreateTableStatementContext)

	// ExitCreateTableOptions is called when exiting the createTableOptions production.
	ExitCreateTableOptions(c *CreateTableOptionsContext)

	// ExitCreateTablespaceStatement is called when exiting the createTablespaceStatement production.
	ExitCreateTablespaceStatement(c *CreateTablespaceStatementContext)

	// ExitCreateTriggerStatement is called when exiting the createTriggerStatement production.
	ExitCreateTriggerStatement(c *CreateTriggerStatementContext)

	// ExitCreateTrustedContextStatement is called when exiting the createTrustedContextStatement production.
	ExitCreateTrustedContextStatement(c *CreateTrustedContextStatementContext)

	// ExitCreateTypeArrayStatement is called when exiting the createTypeArrayStatement production.
	ExitCreateTypeArrayStatement(c *CreateTypeArrayStatementContext)

	// ExitCreateTypeDistinctStatement is called when exiting the createTypeDistinctStatement production.
	ExitCreateTypeDistinctStatement(c *CreateTypeDistinctStatementContext)

	// ExitCreateVariableStatement is called when exiting the createVariableStatement production.
	ExitCreateVariableStatement(c *CreateVariableStatementContext)

	// ExitCreateViewStatement is called when exiting the createViewStatement production.
	ExitCreateViewStatement(c *CreateViewStatementContext)

	// ExitDeclareCursorStatement is called when exiting the declareCursorStatement production.
	ExitDeclareCursorStatement(c *DeclareCursorStatementContext)

	// ExitDeclareGlobalTemporaryTableStatement is called when exiting the declareGlobalTemporaryTableStatement production.
	ExitDeclareGlobalTemporaryTableStatement(c *DeclareGlobalTemporaryTableStatementContext)

	// ExitDeclareTableStatement is called when exiting the declareTableStatement production.
	ExitDeclareTableStatement(c *DeclareTableStatementContext)

	// ExitDeclareStatementStatement is called when exiting the declareStatementStatement production.
	ExitDeclareStatementStatement(c *DeclareStatementStatementContext)

	// ExitDeclareVariableStatement is called when exiting the declareVariableStatement production.
	ExitDeclareVariableStatement(c *DeclareVariableStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitDescribeStatement is called when exiting the describeStatement production.
	ExitDescribeStatement(c *DescribeStatementContext)

	// ExitDescribeCursorStatement is called when exiting the describeCursorStatement production.
	ExitDescribeCursorStatement(c *DescribeCursorStatementContext)

	// ExitDescribeInputStatement is called when exiting the describeInputStatement production.
	ExitDescribeInputStatement(c *DescribeInputStatementContext)

	// ExitDescribeOutputStatement is called when exiting the describeOutputStatement production.
	ExitDescribeOutputStatement(c *DescribeOutputStatementContext)

	// ExitDescribeProcedureStatement is called when exiting the describeProcedureStatement production.
	ExitDescribeProcedureStatement(c *DescribeProcedureStatementContext)

	// ExitDescribeTableStatement is called when exiting the describeTableStatement production.
	ExitDescribeTableStatement(c *DescribeTableStatementContext)

	// ExitDropStatement is called when exiting the dropStatement production.
	ExitDropStatement(c *DropStatementContext)

	// ExitEndDeclareSectionStatement is called when exiting the endDeclareSectionStatement production.
	ExitEndDeclareSectionStatement(c *EndDeclareSectionStatementContext)

	// ExitExchangeStatement is called when exiting the exchangeStatement production.
	ExitExchangeStatement(c *ExchangeStatementContext)

	// ExitExecuteStatement is called when exiting the executeStatement production.
	ExitExecuteStatement(c *ExecuteStatementContext)

	// ExitExecuteImmediateStatement is called when exiting the executeImmediateStatement production.
	ExitExecuteImmediateStatement(c *ExecuteImmediateStatementContext)

	// ExitExplainStatement is called when exiting the explainStatement production.
	ExitExplainStatement(c *ExplainStatementContext)

	// ExitFetchStatement is called when exiting the fetchStatement production.
	ExitFetchStatement(c *FetchStatementContext)

	// ExitFreeLocatorStatement is called when exiting the freeLocatorStatement production.
	ExitFreeLocatorStatement(c *FreeLocatorStatementContext)

	// ExitGetDiagnosticsStatement is called when exiting the getDiagnosticsStatement production.
	ExitGetDiagnosticsStatement(c *GetDiagnosticsStatementContext)

	// ExitGrantStatement is called when exiting the grantStatement production.
	ExitGrantStatement(c *GrantStatementContext)

	// ExitHoldLocatorStatement is called when exiting the holdLocatorStatement production.
	ExitHoldLocatorStatement(c *HoldLocatorStatementContext)

	// ExitIncludeStatement is called when exiting the includeStatement production.
	ExitIncludeStatement(c *IncludeStatementContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitLabelStatement is called when exiting the labelStatement production.
	ExitLabelStatement(c *LabelStatementContext)

	// ExitLockTableStatement is called when exiting the lockTableStatement production.
	ExitLockTableStatement(c *LockTableStatementContext)

	// ExitMergeStatement is called when exiting the mergeStatement production.
	ExitMergeStatement(c *MergeStatementContext)

	// ExitOpenStatement is called when exiting the openStatement production.
	ExitOpenStatement(c *OpenStatementContext)

	// ExitPrepareStatement is called when exiting the prepareStatement production.
	ExitPrepareStatement(c *PrepareStatementContext)

	// ExitRefreshTableStatement is called when exiting the refreshTableStatement production.
	ExitRefreshTableStatement(c *RefreshTableStatementContext)

	// ExitReleaseConnectionStatement is called when exiting the releaseConnectionStatement production.
	ExitReleaseConnectionStatement(c *ReleaseConnectionStatementContext)

	// ExitReleaseSavepointStatement is called when exiting the releaseSavepointStatement production.
	ExitReleaseSavepointStatement(c *ReleaseSavepointStatementContext)

	// ExitRenameStatement is called when exiting the renameStatement production.
	ExitRenameStatement(c *RenameStatementContext)

	// ExitRevokeStatement is called when exiting the revokeStatement production.
	ExitRevokeStatement(c *RevokeStatementContext)

	// ExitRollbackStatement is called when exiting the rollbackStatement production.
	ExitRollbackStatement(c *RollbackStatementContext)

	// ExitSavepointStatement is called when exiting the savepointStatement production.
	ExitSavepointStatement(c *SavepointStatementContext)

	// ExitSetAssignmentStatement is called when exiting the setAssignmentStatement production.
	ExitSetAssignmentStatement(c *SetAssignmentStatementContext)

	// ExitSetConnectionStatement is called when exiting the setConnectionStatement production.
	ExitSetConnectionStatement(c *SetConnectionStatementContext)

	// ExitSetEncryptionPasswordStatement is called when exiting the setEncryptionPasswordStatement production.
	ExitSetEncryptionPasswordStatement(c *SetEncryptionPasswordStatementContext)

	// ExitSetPathStatement is called when exiting the setPathStatement production.
	ExitSetPathStatement(c *SetPathStatementContext)

	// ExitSetSchemaStatement is called when exiting the setSchemaStatement production.
	ExitSetSchemaStatement(c *SetSchemaStatementContext)

	// ExitSetSessionTimezoneStatement is called when exiting the setSessionTimezoneStatement production.
	ExitSetSessionTimezoneStatement(c *SetSessionTimezoneStatementContext)

	// ExitSetSpecialRegisterStatement is called when exiting the setSpecialRegisterStatement production.
	ExitSetSpecialRegisterStatement(c *SetSpecialRegisterStatementContext)

	// ExitSignalStatement is called when exiting the signalStatement production.
	ExitSignalStatement(c *SignalStatementContext)

	// ExitTransferOwnershipStatement is called when exiting the transferOwnershipStatement production.
	ExitTransferOwnershipStatement(c *TransferOwnershipStatementContext)

	// ExitTruncateStatement is called when exiting the truncateStatement production.
	ExitTruncateStatement(c *TruncateStatementContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitValuesStatement is called when exiting the valuesStatement production.
	ExitValuesStatement(c *ValuesStatementContext)

	// ExitValuesIntoStatement is called when exiting the valuesIntoStatement production.
	ExitValuesIntoStatement(c *ValuesIntoStatementContext)

	// ExitWheneverStatement is called when exiting the wheneverStatement production.
	ExitWheneverStatement(c *WheneverStatementContext)

	// ExitValuesIntoTargetVariable is called when exiting the valuesIntoTargetVariable production.
	ExitValuesIntoTargetVariable(c *ValuesIntoTargetVariableContext)

	// ExitOwnedObject is called when exiting the ownedObject production.
	ExitOwnedObject(c *OwnedObjectContext)

	// ExitNewOwner is called when exiting the newOwner production.
	ExitNewOwner(c *NewOwnerContext)

	// ExitGrantCollectionStatement is called when exiting the grantCollectionStatement production.
	ExitGrantCollectionStatement(c *GrantCollectionStatementContext)

	// ExitGrantDatabaseStatement is called when exiting the grantDatabaseStatement production.
	ExitGrantDatabaseStatement(c *GrantDatabaseStatementContext)

	// ExitGrantFunctionOrProcedureStatement is called when exiting the grantFunctionOrProcedureStatement production.
	ExitGrantFunctionOrProcedureStatement(c *GrantFunctionOrProcedureStatementContext)

	// ExitGrantPackageStatement is called when exiting the grantPackageStatement production.
	ExitGrantPackageStatement(c *GrantPackageStatementContext)

	// ExitGrantPlanStatement is called when exiting the grantPlanStatement production.
	ExitGrantPlanStatement(c *GrantPlanStatementContext)

	// ExitGrantSchemaStatement is called when exiting the grantSchemaStatement production.
	ExitGrantSchemaStatement(c *GrantSchemaStatementContext)

	// ExitGrantSequenceStatement is called when exiting the grantSequenceStatement production.
	ExitGrantSequenceStatement(c *GrantSequenceStatementContext)

	// ExitGrantSystemStatement is called when exiting the grantSystemStatement production.
	ExitGrantSystemStatement(c *GrantSystemStatementContext)

	// ExitGrantTableStatement is called when exiting the grantTableStatement production.
	ExitGrantTableStatement(c *GrantTableStatementContext)

	// ExitGrantTypeOrJarStatement is called when exiting the grantTypeOrJarStatement production.
	ExitGrantTypeOrJarStatement(c *GrantTypeOrJarStatementContext)

	// ExitGrantVariableStatement is called when exiting the grantVariableStatement production.
	ExitGrantVariableStatement(c *GrantVariableStatementContext)

	// ExitGrantUseOfStatement is called when exiting the grantUseOfStatement production.
	ExitGrantUseOfStatement(c *GrantUseOfStatementContext)

	// ExitRevokeCollectionStatement is called when exiting the revokeCollectionStatement production.
	ExitRevokeCollectionStatement(c *RevokeCollectionStatementContext)

	// ExitRevokeDatabaseStatement is called when exiting the revokeDatabaseStatement production.
	ExitRevokeDatabaseStatement(c *RevokeDatabaseStatementContext)

	// ExitRevokeFunctionOrProcedureStatement is called when exiting the revokeFunctionOrProcedureStatement production.
	ExitRevokeFunctionOrProcedureStatement(c *RevokeFunctionOrProcedureStatementContext)

	// ExitRevokePackageStatement is called when exiting the revokePackageStatement production.
	ExitRevokePackageStatement(c *RevokePackageStatementContext)

	// ExitRevokePlanStatement is called when exiting the revokePlanStatement production.
	ExitRevokePlanStatement(c *RevokePlanStatementContext)

	// ExitRevokeSchemaStatement is called when exiting the revokeSchemaStatement production.
	ExitRevokeSchemaStatement(c *RevokeSchemaStatementContext)

	// ExitRevokeSequenceStatement is called when exiting the revokeSequenceStatement production.
	ExitRevokeSequenceStatement(c *RevokeSequenceStatementContext)

	// ExitRevokeSystemStatement is called when exiting the revokeSystemStatement production.
	ExitRevokeSystemStatement(c *RevokeSystemStatementContext)

	// ExitRevokeTableStatement is called when exiting the revokeTableStatement production.
	ExitRevokeTableStatement(c *RevokeTableStatementContext)

	// ExitRevokeTypeOrJarStatement is called when exiting the revokeTypeOrJarStatement production.
	ExitRevokeTypeOrJarStatement(c *RevokeTypeOrJarStatementContext)

	// ExitRevokeVariableStatement is called when exiting the revokeVariableStatement production.
	ExitRevokeVariableStatement(c *RevokeVariableStatementContext)

	// ExitRevokeUseOfStatement is called when exiting the revokeUseOfStatement production.
	ExitRevokeUseOfStatement(c *RevokeUseOfStatementContext)

	// ExitGrantUseOfTarget is called when exiting the grantUseOfTarget production.
	ExitGrantUseOfTarget(c *GrantUseOfTargetContext)

	// ExitGrantVariableAuthority is called when exiting the grantVariableAuthority production.
	ExitGrantVariableAuthority(c *GrantVariableAuthorityContext)

	// ExitGrantTableAuthority is called when exiting the grantTableAuthority production.
	ExitGrantTableAuthority(c *GrantTableAuthorityContext)

	// ExitGrantSystemAuthority is called when exiting the grantSystemAuthority production.
	ExitGrantSystemAuthority(c *GrantSystemAuthorityContext)

	// ExitGrantSequenceAuthority is called when exiting the grantSequenceAuthority production.
	ExitGrantSequenceAuthority(c *GrantSequenceAuthorityContext)

	// ExitGrantSchemaAuthority is called when exiting the grantSchemaAuthority production.
	ExitGrantSchemaAuthority(c *GrantSchemaAuthorityContext)

	// ExitGrantPlanAuthority is called when exiting the grantPlanAuthority production.
	ExitGrantPlanAuthority(c *GrantPlanAuthorityContext)

	// ExitGrantPackageAuthority is called when exiting the grantPackageAuthority production.
	ExitGrantPackageAuthority(c *GrantPackageAuthorityContext)

	// ExitPackageSpecification is called when exiting the packageSpecification production.
	ExitPackageSpecification(c *PackageSpecificationContext)

	// ExitFunctionSpecification is called when exiting the functionSpecification production.
	ExitFunctionSpecification(c *FunctionSpecificationContext)

	// ExitGrantee is called when exiting the grantee production.
	ExitGrantee(c *GranteeContext)

	// ExitWithGrantOption is called when exiting the withGrantOption production.
	ExitWithGrantOption(c *WithGrantOptionContext)

	// ExitRevokeByOption is called when exiting the revokeByOption production.
	ExitRevokeByOption(c *RevokeByOptionContext)

	// ExitRevokeDependentPrivilegesOption is called when exiting the revokeDependentPrivilegesOption production.
	ExitRevokeDependentPrivilegesOption(c *RevokeDependentPrivilegesOptionContext)

	// ExitGrantDatabaseAuthority is called when exiting the grantDatabaseAuthority production.
	ExitGrantDatabaseAuthority(c *GrantDatabaseAuthorityContext)

	// ExitStatementInformation is called when exiting the statementInformation production.
	ExitStatementInformation(c *StatementInformationContext)

	// ExitStatementInformationVariableEquate is called when exiting the statementInformationVariableEquate production.
	ExitStatementInformationVariableEquate(c *StatementInformationVariableEquateContext)

	// ExitStatementInformationItemName is called when exiting the statementInformationItemName production.
	ExitStatementInformationItemName(c *StatementInformationItemNameContext)

	// ExitConditionInformation is called when exiting the conditionInformation production.
	ExitConditionInformation(c *ConditionInformationContext)

	// ExitConditionInformationVariableEquate is called when exiting the conditionInformationVariableEquate production.
	ExitConditionInformationVariableEquate(c *ConditionInformationVariableEquateContext)

	// ExitConditionInformationItemName is called when exiting the conditionInformationItemName production.
	ExitConditionInformationItemName(c *ConditionInformationItemNameContext)

	// ExitConnectionInformationItemName is called when exiting the connectionInformationItemName production.
	ExitConnectionInformationItemName(c *ConnectionInformationItemNameContext)

	// ExitCombinedInformation is called when exiting the combinedInformation production.
	ExitCombinedInformation(c *CombinedInformationContext)

	// ExitCombinedInformationOption is called when exiting the combinedInformationOption production.
	ExitCombinedInformationOption(c *CombinedInformationOptionContext)

	// ExitFetchOrientation is called when exiting the fetchOrientation production.
	ExitFetchOrientation(c *FetchOrientationContext)

	// ExitRowPositioned is called when exiting the rowPositioned production.
	ExitRowPositioned(c *RowPositionedContext)

	// ExitRowsetPositioned is called when exiting the rowsetPositioned production.
	ExitRowsetPositioned(c *RowsetPositionedContext)

	// ExitSingleRowFetch is called when exiting the singleRowFetch production.
	ExitSingleRowFetch(c *SingleRowFetchContext)

	// ExitFetchTargetVariable is called when exiting the fetchTargetVariable production.
	ExitFetchTargetVariable(c *FetchTargetVariableContext)

	// ExitMultipleRowFetch is called when exiting the multipleRowFetch production.
	ExitMultipleRowFetch(c *MultipleRowFetchContext)

	// ExitMultipleRowFetchForClause is called when exiting the multipleRowFetchForClause production.
	ExitMultipleRowFetchForClause(c *MultipleRowFetchForClauseContext)

	// ExitMultipleRowFetchIntoClause is called when exiting the multipleRowFetchIntoClause production.
	ExitMultipleRowFetchIntoClause(c *MultipleRowFetchIntoClauseContext)

	// ExitExplainPlanClause is called when exiting the explainPlanClause production.
	ExitExplainPlanClause(c *ExplainPlanClauseContext)

	// ExitExplainStmtcacheClause is called when exiting the explainStmtcacheClause production.
	ExitExplainStmtcacheClause(c *ExplainStmtcacheClauseContext)

	// ExitExplainPackageClause is called when exiting the explainPackageClause production.
	ExitExplainPackageClause(c *ExplainPackageClauseContext)

	// ExitExplainStabilizedDynamicQueryClause is called when exiting the explainStabilizedDynamicQueryClause production.
	ExitExplainStabilizedDynamicQueryClause(c *ExplainStabilizedDynamicQueryClauseContext)

	// ExitPackageScopeSpecification is called when exiting the packageScopeSpecification production.
	ExitPackageScopeSpecification(c *PackageScopeSpecificationContext)

	// ExitCollectionName is called when exiting the collectionName production.
	ExitCollectionName(c *CollectionNameContext)

	// ExitPackageScopePackageName is called when exiting the packageScopePackageName production.
	ExitPackageScopePackageName(c *PackageScopePackageNameContext)

	// ExitVersionName is called when exiting the versionName production.
	ExitVersionName(c *VersionNameContext)

	// ExitSourceRowData is called when exiting the sourceRowData production.
	ExitSourceRowData(c *SourceRowDataContext)

	// ExitAliasDesignation is called when exiting the aliasDesignation production.
	ExitAliasDesignation(c *AliasDesignationContext)

	// ExitDropDatabaseClause is called when exiting the dropDatabaseClause production.
	ExitDropDatabaseClause(c *DropDatabaseClauseContext)

	// ExitDropFunctionClause is called when exiting the dropFunctionClause production.
	ExitDropFunctionClause(c *DropFunctionClauseContext)

	// ExitDropIndexClause is called when exiting the dropIndexClause production.
	ExitDropIndexClause(c *DropIndexClauseContext)

	// ExitDropMaskClause is called when exiting the dropMaskClause production.
	ExitDropMaskClause(c *DropMaskClauseContext)

	// ExitDropPackageClause is called when exiting the dropPackageClause production.
	ExitDropPackageClause(c *DropPackageClauseContext)

	// ExitDropPermissionClause is called when exiting the dropPermissionClause production.
	ExitDropPermissionClause(c *DropPermissionClauseContext)

	// ExitDropProcedureClause is called when exiting the dropProcedureClause production.
	ExitDropProcedureClause(c *DropProcedureClauseContext)

	// ExitDropRoleClause is called when exiting the dropRoleClause production.
	ExitDropRoleClause(c *DropRoleClauseContext)

	// ExitDropSequenceClause is called when exiting the dropSequenceClause production.
	ExitDropSequenceClause(c *DropSequenceClauseContext)

	// ExitDropStogroupClause is called when exiting the dropStogroupClause production.
	ExitDropStogroupClause(c *DropStogroupClauseContext)

	// ExitDropSynonymClause is called when exiting the dropSynonymClause production.
	ExitDropSynonymClause(c *DropSynonymClauseContext)

	// ExitDropTableClause is called when exiting the dropTableClause production.
	ExitDropTableClause(c *DropTableClauseContext)

	// ExitDropTablespaceClause is called when exiting the dropTablespaceClause production.
	ExitDropTablespaceClause(c *DropTablespaceClauseContext)

	// ExitDropTriggerClause is called when exiting the dropTriggerClause production.
	ExitDropTriggerClause(c *DropTriggerClauseContext)

	// ExitDropTrustedContextClause is called when exiting the dropTrustedContextClause production.
	ExitDropTrustedContextClause(c *DropTrustedContextClauseContext)

	// ExitDropTypeClause is called when exiting the dropTypeClause production.
	ExitDropTypeClause(c *DropTypeClauseContext)

	// ExitDropVariableClause is called when exiting the dropVariableClause production.
	ExitDropVariableClause(c *DropVariableClauseContext)

	// ExitDropViewClause is called when exiting the dropViewClause production.
	ExitDropViewClause(c *DropViewClauseContext)

	// ExitPackageDesignator is called when exiting the packageDesignator production.
	ExitPackageDesignator(c *PackageDesignatorContext)

	// ExitDescribeUsingOption is called when exiting the describeUsingOption production.
	ExitDescribeUsingOption(c *DescribeUsingOptionContext)

	// ExitDeclareGlobalTemporaryTableLikeClause is called when exiting the declareGlobalTemporaryTableLikeClause production.
	ExitDeclareGlobalTemporaryTableLikeClause(c *DeclareGlobalTemporaryTableLikeClauseContext)

	// ExitOnCommitClause is called when exiting the onCommitClause production.
	ExitOnCommitClause(c *OnCommitClauseContext)

	// ExitLoggedWithRollbackClause is called when exiting the loggedWithRollbackClause production.
	ExitLoggedWithRollbackClause(c *LoggedWithRollbackClauseContext)

	// ExitCreateViewCheckOptionClause is called when exiting the createViewCheckOptionClause production.
	ExitCreateViewCheckOptionClause(c *CreateViewCheckOptionClauseContext)

	// ExitTrustedContextDefaultRoleClause is called when exiting the trustedContextDefaultRoleClause production.
	ExitTrustedContextDefaultRoleClause(c *TrustedContextDefaultRoleClauseContext)

	// ExitTrustedContextEnableDisableClause is called when exiting the trustedContextEnableDisableClause production.
	ExitTrustedContextEnableDisableClause(c *TrustedContextEnableDisableClauseContext)

	// ExitTrustedContextDefaultSecurityLabelClause is called when exiting the trustedContextDefaultSecurityLabelClause production.
	ExitTrustedContextDefaultSecurityLabelClause(c *TrustedContextDefaultSecurityLabelClauseContext)

	// ExitTrustedContextAttributesClause is called when exiting the trustedContextAttributesClause production.
	ExitTrustedContextAttributesClause(c *TrustedContextAttributesClauseContext)

	// ExitTrustedContextWithUseForClause is called when exiting the trustedContextWithUseForClause production.
	ExitTrustedContextWithUseForClause(c *TrustedContextWithUseForClauseContext)

	// ExitTrustedContextAttribute1 is called when exiting the trustedContextAttribute1 production.
	ExitTrustedContextAttribute1(c *TrustedContextAttribute1Context)

	// ExitTrustedContextAttribute2 is called when exiting the trustedContextAttribute2 production.
	ExitTrustedContextAttribute2(c *TrustedContextAttribute2Context)

	// ExitTrustedContextUseFor is called when exiting the trustedContextUseFor production.
	ExitTrustedContextUseFor(c *TrustedContextUseForContext)

	// ExitUserOptions is called when exiting the userOptions production.
	ExitUserOptions(c *UserOptionsContext)

	// ExitTriggerDefinition is called when exiting the triggerDefinition production.
	ExitTriggerDefinition(c *TriggerDefinitionContext)

	// ExitTriggerActivationTime is called when exiting the triggerActivationTime production.
	ExitTriggerActivationTime(c *TriggerActivationTimeContext)

	// ExitTriggerEvent is called when exiting the triggerEvent production.
	ExitTriggerEvent(c *TriggerEventContext)

	// ExitTriggerGranularity is called when exiting the triggerGranularity production.
	ExitTriggerGranularity(c *TriggerGranularityContext)

	// ExitTriggeredAction is called when exiting the triggeredAction production.
	ExitTriggeredAction(c *TriggeredActionContext)

	// ExitSqlTriggerBody is called when exiting the sqlTriggerBody production.
	ExitSqlTriggerBody(c *SqlTriggerBodyContext)

	// ExitTriggeredSqlStatement is called when exiting the triggeredSqlStatement production.
	ExitTriggeredSqlStatement(c *TriggeredSqlStatementContext)

	// ExitTriggerDefinitionOption is called when exiting the triggerDefinitionOption production.
	ExitTriggerDefinitionOption(c *TriggerDefinitionOptionContext)

	// ExitCreateTableInClause is called when exiting the createTableInClause production.
	ExitCreateTableInClause(c *CreateTableInClauseContext)

	// ExitCustomVolatileClause is called when exiting the customVolatileClause production.
	ExitCustomVolatileClause(c *CustomVolatileClauseContext)

	// ExitCreateTableColumnDefinition is called when exiting the createTableColumnDefinition production.
	ExitCreateTableColumnDefinition(c *CreateTableColumnDefinitionContext)

	// ExitEditprocClause is called when exiting the editprocClause production.
	ExitEditprocClause(c *EditprocClauseContext)

	// ExitValidprocClause is called when exiting the validprocClause production.
	ExitValidprocClause(c *ValidprocClauseContext)

	// ExitAuditClause is called when exiting the auditClause production.
	ExitAuditClause(c *AuditClauseContext)

	// ExitObidClause is called when exiting the obidClause production.
	ExitObidClause(c *ObidClauseContext)

	// ExitDataCaptureClause is called when exiting the dataCaptureClause production.
	ExitDataCaptureClause(c *DataCaptureClauseContext)

	// ExitRestrictOnDropClause is called when exiting the restrictOnDropClause production.
	ExitRestrictOnDropClause(c *RestrictOnDropClauseContext)

	// ExitCcsidClause1 is called when exiting the ccsidClause1 production.
	ExitCcsidClause1(c *CcsidClause1Context)

	// ExitCcsidClause2 is called when exiting the ccsidClause2 production.
	ExitCcsidClause2(c *CcsidClause2Context)

	// ExitCardinalityClause is called when exiting the cardinalityClause production.
	ExitCardinalityClause(c *CardinalityClauseContext)

	// ExitAppendClause is called when exiting the appendClause production.
	ExitAppendClause(c *AppendClauseContext)

	// ExitMemberClause is called when exiting the memberClause production.
	ExitMemberClause(c *MemberClauseContext)

	// ExitTrackmodClause is called when exiting the trackmodClause production.
	ExitTrackmodClause(c *TrackmodClauseContext)

	// ExitPagenumClause is called when exiting the pagenumClause production.
	ExitPagenumClause(c *PagenumClauseContext)

	// ExitFieldprocClause is called when exiting the fieldprocClause production.
	ExitFieldprocClause(c *FieldprocClauseContext)

	// ExitAsSecurityLabelClause is called when exiting the asSecurityLabelClause production.
	ExitAsSecurityLabelClause(c *AsSecurityLabelClauseContext)

	// ExitImplicitlyHiddenClause is called when exiting the implicitlyHiddenClause production.
	ExitImplicitlyHiddenClause(c *ImplicitlyHiddenClauseContext)

	// ExitInlineLengthClause is called when exiting the inlineLengthClause production.
	ExitInlineLengthClause(c *InlineLengthClauseContext)

	// ExitCopyOptions is called when exiting the copyOptions production.
	ExitCopyOptions(c *CopyOptionsContext)

	// ExitCopyOptionIdentity is called when exiting the copyOptionIdentity production.
	ExitCopyOptionIdentity(c *CopyOptionIdentityContext)

	// ExitCopyOptionRowChangeTimestamp is called when exiting the copyOptionRowChangeTimestamp production.
	ExitCopyOptionRowChangeTimestamp(c *CopyOptionRowChangeTimestampContext)

	// ExitCopyOptionColumnDefaults is called when exiting the copyOptionColumnDefaults production.
	ExitCopyOptionColumnDefaults(c *CopyOptionColumnDefaultsContext)

	// ExitCopyOptionXmlTypeModifiers is called when exiting the copyOptionXmlTypeModifiers production.
	ExitCopyOptionXmlTypeModifiers(c *CopyOptionXmlTypeModifiersContext)

	// ExitAsResultTable is called when exiting the asResultTable production.
	ExitAsResultTable(c *AsResultTableContext)

	// ExitDeclareGlobalTemporaryTableAsResultTable is called when exiting the declareGlobalTemporaryTableAsResultTable production.
	ExitDeclareGlobalTemporaryTableAsResultTable(c *DeclareGlobalTemporaryTableAsResultTableContext)

	// ExitCreateTableMaterializedQueryDefinition is called when exiting the createTableMaterializedQueryDefinition production.
	ExitCreateTableMaterializedQueryDefinition(c *CreateTableMaterializedQueryDefinitionContext)

	// ExitCreateTableColumnConstraint is called when exiting the createTableColumnConstraint production.
	ExitCreateTableColumnConstraint(c *CreateTableColumnConstraintContext)

	// ExitOrganizationClause is called when exiting the organizationClause production.
	ExitOrganizationClause(c *OrganizationClauseContext)

	// ExitCreateGlobalTemporaryTableColumnDefinition is called when exiting the createGlobalTemporaryTableColumnDefinition production.
	ExitCreateGlobalTemporaryTableColumnDefinition(c *CreateGlobalTemporaryTableColumnDefinitionContext)

	// ExitDeclareGlobalTemporaryTableColumnDefinition is called when exiting the declareGlobalTemporaryTableColumnDefinition production.
	ExitDeclareGlobalTemporaryTableColumnDefinition(c *DeclareGlobalTemporaryTableColumnDefinitionContext)

	// ExitParameterDeclaration1 is called when exiting the parameterDeclaration1 production.
	ExitParameterDeclaration1(c *ParameterDeclaration1Context)

	// ExitParameterDeclaration2 is called when exiting the parameterDeclaration2 production.
	ExitParameterDeclaration2(c *ParameterDeclaration2Context)

	// ExitParameterDeclaration3 is called when exiting the parameterDeclaration3 production.
	ExitParameterDeclaration3(c *ParameterDeclaration3Context)

	// ExitCreateFunctionStatementExternalScalarOptions is called when exiting the createFunctionStatementExternalScalarOptions production.
	ExitCreateFunctionStatementExternalScalarOptions(c *CreateFunctionStatementExternalScalarOptionsContext)

	// ExitExternalNameOption1 is called when exiting the externalNameOption1 production.
	ExitExternalNameOption1(c *ExternalNameOption1Context)

	// ExitExternalNameOption2 is called when exiting the externalNameOption2 production.
	ExitExternalNameOption2(c *ExternalNameOption2Context)

	// ExitDynamicResultSetOption is called when exiting the dynamicResultSetOption production.
	ExitDynamicResultSetOption(c *DynamicResultSetOptionContext)

	// ExitLanguageOption1 is called when exiting the languageOption1 production.
	ExitLanguageOption1(c *LanguageOption1Context)

	// ExitLanguageOption2 is called when exiting the languageOption2 production.
	ExitLanguageOption2(c *LanguageOption2Context)

	// ExitLanguageOption3 is called when exiting the languageOption3 production.
	ExitLanguageOption3(c *LanguageOption3Context)

	// ExitLanguageOption4 is called when exiting the languageOption4 production.
	ExitLanguageOption4(c *LanguageOption4Context)

	// ExitLanguageOption5 is called when exiting the languageOption5 production.
	ExitLanguageOption5(c *LanguageOption5Context)

	// ExitParameterStyleOption1 is called when exiting the parameterStyleOption1 production.
	ExitParameterStyleOption1(c *ParameterStyleOption1Context)

	// ExitParameterStyleOption2 is called when exiting the parameterStyleOption2 production.
	ExitParameterStyleOption2(c *ParameterStyleOption2Context)

	// ExitParameterStyleOption3 is called when exiting the parameterStyleOption3 production.
	ExitParameterStyleOption3(c *ParameterStyleOption3Context)

	// ExitDeterministicOption is called when exiting the deterministicOption production.
	ExitDeterministicOption(c *DeterministicOptionContext)

	// ExitFencedOption is called when exiting the fencedOption production.
	ExitFencedOption(c *FencedOptionContext)

	// ExitNullInputOption1 is called when exiting the nullInputOption1 production.
	ExitNullInputOption1(c *NullInputOption1Context)

	// ExitNullInputOption2 is called when exiting the nullInputOption2 production.
	ExitNullInputOption2(c *NullInputOption2Context)

	// ExitNullInputOption3 is called when exiting the nullInputOption3 production.
	ExitNullInputOption3(c *NullInputOption3Context)

	// ExitNullInputOption4 is called when exiting the nullInputOption4 production.
	ExitNullInputOption4(c *NullInputOption4Context)

	// ExitDebugOption is called when exiting the debugOption production.
	ExitDebugOption(c *DebugOptionContext)

	// ExitSqlDataOption1 is called when exiting the sqlDataOption1 production.
	ExitSqlDataOption1(c *SqlDataOption1Context)

	// ExitSqlDataOption2 is called when exiting the sqlDataOption2 production.
	ExitSqlDataOption2(c *SqlDataOption2Context)

	// ExitSqlDataOption3 is called when exiting the sqlDataOption3 production.
	ExitSqlDataOption3(c *SqlDataOption3Context)

	// ExitSqlDataOption4 is called when exiting the sqlDataOption4 production.
	ExitSqlDataOption4(c *SqlDataOption4Context)

	// ExitExternalActionOption is called when exiting the externalActionOption production.
	ExitExternalActionOption(c *ExternalActionOptionContext)

	// ExitPackagePathOption is called when exiting the packagePathOption production.
	ExitPackagePathOption(c *PackagePathOptionContext)

	// ExitScratchpadOption is called when exiting the scratchpadOption production.
	ExitScratchpadOption(c *ScratchpadOptionContext)

	// ExitFinalCallOption is called when exiting the finalCallOption production.
	ExitFinalCallOption(c *FinalCallOptionContext)

	// ExitParallelOption1 is called when exiting the parallelOption1 production.
	ExitParallelOption1(c *ParallelOption1Context)

	// ExitParallelOption2 is called when exiting the parallelOption2 production.
	ExitParallelOption2(c *ParallelOption2Context)

	// ExitDbinfoOption is called when exiting the dbinfoOption production.
	ExitDbinfoOption(c *DbinfoOptionContext)

	// ExitCardinalityOption is called when exiting the cardinalityOption production.
	ExitCardinalityOption(c *CardinalityOptionContext)

	// ExitCollectionIdOption is called when exiting the collectionIdOption production.
	ExitCollectionIdOption(c *CollectionIdOptionContext)

	// ExitWlmEnvironmentOption1 is called when exiting the wlmEnvironmentOption1 production.
	ExitWlmEnvironmentOption1(c *WlmEnvironmentOption1Context)

	// ExitWlmEnvironmentOption2 is called when exiting the wlmEnvironmentOption2 production.
	ExitWlmEnvironmentOption2(c *WlmEnvironmentOption2Context)

	// ExitWlmEnvironmentOption3 is called when exiting the wlmEnvironmentOption3 production.
	ExitWlmEnvironmentOption3(c *WlmEnvironmentOption3Context)

	// ExitAsuTimeOption is called when exiting the asuTimeOption production.
	ExitAsuTimeOption(c *AsuTimeOptionContext)

	// ExitStayResidentOption is called when exiting the stayResidentOption production.
	ExitStayResidentOption(c *StayResidentOptionContext)

	// ExitProgramTypeOption is called when exiting the programTypeOption production.
	ExitProgramTypeOption(c *ProgramTypeOptionContext)

	// ExitSecurityOption is called when exiting the securityOption production.
	ExitSecurityOption(c *SecurityOptionContext)

	// ExitStopAfterFailureOption is called when exiting the stopAfterFailureOption production.
	ExitStopAfterFailureOption(c *StopAfterFailureOptionContext)

	// ExitRunOptionsOption is called when exiting the runOptionsOption production.
	ExitRunOptionsOption(c *RunOptionsOptionContext)

	// ExitCommitOnReturnOption is called when exiting the commitOnReturnOption production.
	ExitCommitOnReturnOption(c *CommitOnReturnOptionContext)

	// ExitSpecialRegistersOption is called when exiting the specialRegistersOption production.
	ExitSpecialRegistersOption(c *SpecialRegistersOptionContext)

	// ExitSpecialRegistersOption2 is called when exiting the specialRegistersOption2 production.
	ExitSpecialRegistersOption2(c *SpecialRegistersOption2Context)

	// ExitDispatchOption is called when exiting the dispatchOption production.
	ExitDispatchOption(c *DispatchOptionContext)

	// ExitSecuredOption is called when exiting the securedOption production.
	ExitSecuredOption(c *SecuredOptionContext)

	// ExitSpecificNameOption1 is called when exiting the specificNameOption1 production.
	ExitSpecificNameOption1(c *SpecificNameOption1Context)

	// ExitSpecificNameOption2 is called when exiting the specificNameOption2 production.
	ExitSpecificNameOption2(c *SpecificNameOption2Context)

	// ExitParameterOption1 is called when exiting the parameterOption1 production.
	ExitParameterOption1(c *ParameterOption1Context)

	// ExitParameterOption2 is called when exiting the parameterOption2 production.
	ExitParameterOption2(c *ParameterOption2Context)

	// ExitCreateFunctionStatementExternalTableOptions is called when exiting the createFunctionStatementExternalTableOptions production.
	ExitCreateFunctionStatementExternalTableOptions(c *CreateFunctionStatementExternalTableOptionsContext)

	// ExitCreateFunctionStatementSourcedOptions is called when exiting the createFunctionStatementSourcedOptions production.
	ExitCreateFunctionStatementSourcedOptions(c *CreateFunctionStatementSourcedOptionsContext)

	// ExitInlineSqlScalarFunctionDefinition is called when exiting the inlineSqlScalarFunctionDefinition production.
	ExitInlineSqlScalarFunctionDefinition(c *InlineSqlScalarFunctionDefinitionContext)

	// ExitCreateFunctionStatementInlineSqlScalarOptions is called when exiting the createFunctionStatementInlineSqlScalarOptions production.
	ExitCreateFunctionStatementInlineSqlScalarOptions(c *CreateFunctionStatementInlineSqlScalarOptionsContext)

	// ExitCompiledSqlScalarFunctionDefinition is called when exiting the compiledSqlScalarFunctionDefinition production.
	ExitCompiledSqlScalarFunctionDefinition(c *CompiledSqlScalarFunctionDefinitionContext)

	// ExitCreateFunctionStatementCompiledSqlScalarOptions is called when exiting the createFunctionStatementCompiledSqlScalarOptions production.
	ExitCreateFunctionStatementCompiledSqlScalarOptions(c *CreateFunctionStatementCompiledSqlScalarOptionsContext)

	// ExitSqlTableFunctionDefinition is called when exiting the sqlTableFunctionDefinition production.
	ExitSqlTableFunctionDefinition(c *SqlTableFunctionDefinitionContext)

	// ExitCreateFunctionStatementSqlTableOptions is called when exiting the createFunctionStatementSqlTableOptions production.
	ExitCreateFunctionStatementSqlTableOptions(c *CreateFunctionStatementSqlTableOptionsContext)

	// ExitSequenceAlias is called when exiting the sequenceAlias production.
	ExitSequenceAlias(c *SequenceAliasContext)

	// ExitTableAlias is called when exiting the tableAlias production.
	ExitTableAlias(c *TableAliasContext)

	// ExitAuthorization is called when exiting the authorization production.
	ExitAuthorization(c *AuthorizationContext)

	// ExitSearchedDelete is called when exiting the searchedDelete production.
	ExitSearchedDelete(c *SearchedDeleteContext)

	// ExitPositionedDelete is called when exiting the positionedDelete production.
	ExitPositionedDelete(c *PositionedDeleteContext)

	// ExitSearchedUpdate is called when exiting the searchedUpdate production.
	ExitSearchedUpdate(c *SearchedUpdateContext)

	// ExitPositionedUpdate is called when exiting the positionedUpdate production.
	ExitPositionedUpdate(c *PositionedUpdateContext)

	// ExitSourceValues is called when exiting the sourceValues production.
	ExitSourceValues(c *SourceValuesContext)

	// ExitValuesSingleRow is called when exiting the valuesSingleRow production.
	ExitValuesSingleRow(c *ValuesSingleRowContext)

	// ExitValuesMultipleRow is called when exiting the valuesMultipleRow production.
	ExitValuesMultipleRow(c *ValuesMultipleRowContext)

	// ExitMatchingCondition is called when exiting the matchingCondition production.
	ExitMatchingCondition(c *MatchingConditionContext)

	// ExitModificationOperation is called when exiting the modificationOperation production.
	ExitModificationOperation(c *ModificationOperationContext)

	// ExitAssignmentClause is called when exiting the assignmentClause production.
	ExitAssignmentClause(c *AssignmentClauseContext)

	// ExitSetAssignmentClause is called when exiting the setAssignmentClause production.
	ExitSetAssignmentClause(c *SetAssignmentClauseContext)

	// ExitSetAssignmentTargetVariable is called when exiting the setAssignmentTargetVariable production.
	ExitSetAssignmentTargetVariable(c *SetAssignmentTargetVariableContext)

	// ExitUpdateOperation is called when exiting the updateOperation production.
	ExitUpdateOperation(c *UpdateOperationContext)

	// ExitDeleteOperation is called when exiting the deleteOperation production.
	ExitDeleteOperation(c *DeleteOperationContext)

	// ExitInsertOperation is called when exiting the insertOperation production.
	ExitInsertOperation(c *InsertOperationContext)

	// ExitSignalInformation is called when exiting the signalInformation production.
	ExitSignalInformation(c *SignalInformationContext)

	// ExitValuesList1 is called when exiting the valuesList1 production.
	ExitValuesList1(c *ValuesList1Context)

	// ExitValuesList2 is called when exiting the valuesList2 production.
	ExitValuesList2(c *ValuesList2Context)

	// ExitValuesList3 is called when exiting the valuesList3 production.
	ExitValuesList3(c *ValuesList3Context)

	// ExitValuesList4 is called when exiting the valuesList4 production.
	ExitValuesList4(c *ValuesList4Context)

	// ExitIncludeColumns is called when exiting the includeColumns production.
	ExitIncludeColumns(c *IncludeColumnsContext)

	// ExitMultipleRowInsert is called when exiting the multipleRowInsert production.
	ExitMultipleRowInsert(c *MultipleRowInsertContext)

	// ExitRegenerateClause is called when exiting the regenerateClause production.
	ExitRegenerateClause(c *RegenerateClauseContext)

	// ExitAlterIndexOptions is called when exiting the alterIndexOptions production.
	ExitAlterIndexOptions(c *AlterIndexOptionsContext)

	// ExitBufferpoolOption is called when exiting the bufferpoolOption production.
	ExitBufferpoolOption(c *BufferpoolOptionContext)

	// ExitCloseOption is called when exiting the closeOption production.
	ExitCloseOption(c *CloseOptionContext)

	// ExitCopyOption is called when exiting the copyOption production.
	ExitCopyOption(c *CopyOptionContext)

	// ExitDssizeOption is called when exiting the dssizeOption production.
	ExitDssizeOption(c *DssizeOptionContext)

	// ExitPiecesizeOption is called when exiting the piecesizeOption production.
	ExitPiecesizeOption(c *PiecesizeOptionContext)

	// ExitClusterOption is called when exiting the clusterOption production.
	ExitClusterOption(c *ClusterOptionContext)

	// ExitPaddedOption is called when exiting the paddedOption production.
	ExitPaddedOption(c *PaddedOptionContext)

	// ExitCompressOption is called when exiting the compressOption production.
	ExitCompressOption(c *CompressOptionContext)

	// ExitDefineOption is called when exiting the defineOption production.
	ExitDefineOption(c *DefineOptionContext)

	// ExitLocksizeOption is called when exiting the locksizeOption production.
	ExitLocksizeOption(c *LocksizeOptionContext)

	// ExitLockmaxOption is called when exiting the lockmaxOption production.
	ExitLockmaxOption(c *LockmaxOptionContext)

	// ExitEnableDisableOption is called when exiting the enableDisableOption production.
	ExitEnableDisableOption(c *EnableDisableOptionContext)

	// ExitLoggedOption is called when exiting the loggedOption production.
	ExitLoggedOption(c *LoggedOptionContext)

	// ExitNotAtomicPhrase is called when exiting the notAtomicPhrase production.
	ExitNotAtomicPhrase(c *NotAtomicPhraseContext)

	// ExitAlterIndexPartitionOptions is called when exiting the alterIndexPartitionOptions production.
	ExitAlterIndexPartitionOptions(c *AlterIndexPartitionOptionsContext)

	// ExitUsingSpecification1 is called when exiting the usingSpecification1 production.
	ExitUsingSpecification1(c *UsingSpecification1Context)

	// ExitFreeSpecification is called when exiting the freeSpecification production.
	ExitFreeSpecification(c *FreeSpecificationContext)

	// ExitGbpcacheSpecification is called when exiting the gbpcacheSpecification production.
	ExitGbpcacheSpecification(c *GbpcacheSpecificationContext)

	// ExitPartitionElement is called when exiting the partitionElement production.
	ExitPartitionElement(c *PartitionElementContext)

	// ExitApplCompatValue is called when exiting the applCompatValue production.
	ExitApplCompatValue(c *ApplCompatValueContext)

	// ExitFunctionLevel is called when exiting the functionLevel production.
	ExitFunctionLevel(c *FunctionLevelContext)

	// ExitFunctionParameterType is called when exiting the functionParameterType production.
	ExitFunctionParameterType(c *FunctionParameterTypeContext)

	// ExitFunctionDataType is called when exiting the functionDataType production.
	ExitFunctionDataType(c *FunctionDataTypeContext)

	// ExitFunctionBuiltInType is called when exiting the functionBuiltInType production.
	ExitFunctionBuiltInType(c *FunctionBuiltInTypeContext)

	// ExitProcedureBuiltinType is called when exiting the procedureBuiltinType production.
	ExitProcedureBuiltinType(c *ProcedureBuiltinTypeContext)

	// ExitCreateTypeArrayBuiltinType is called when exiting the createTypeArrayBuiltinType production.
	ExitCreateTypeArrayBuiltinType(c *CreateTypeArrayBuiltinTypeContext)

	// ExitCreateTypeArrayBuiltinType2 is called when exiting the createTypeArrayBuiltinType2 production.
	ExitCreateTypeArrayBuiltinType2(c *CreateTypeArrayBuiltinType2Context)

	// ExitCreateVariableBuiltInType is called when exiting the createVariableBuiltInType production.
	ExitCreateVariableBuiltInType(c *CreateVariableBuiltInTypeContext)

	// ExitSourceDataType is called when exiting the sourceDataType production.
	ExitSourceDataType(c *SourceDataTypeContext)

	// ExitFunctionExternalOptionList is called when exiting the functionExternalOptionList production.
	ExitFunctionExternalOptionList(c *FunctionExternalOptionListContext)

	// ExitFunctionCompiledSqlScalarOptionList is called when exiting the functionCompiledSqlScalarOptionList production.
	ExitFunctionCompiledSqlScalarOptionList(c *FunctionCompiledSqlScalarOptionListContext)

	// ExitFunctionInlineSqlScalarOptionList is called when exiting the functionInlineSqlScalarOptionList production.
	ExitFunctionInlineSqlScalarOptionList(c *FunctionInlineSqlScalarOptionListContext)

	// ExitFunctionSqlTableOptionList is called when exiting the functionSqlTableOptionList production.
	ExitFunctionSqlTableOptionList(c *FunctionSqlTableOptionListContext)

	// ExitProcedureOptionList is called when exiting the procedureOptionList production.
	ExitProcedureOptionList(c *ProcedureOptionListContext)

	// ExitCreateProcedureOptionList is called when exiting the createProcedureOptionList production.
	ExitCreateProcedureOptionList(c *CreateProcedureOptionListContext)

	// ExitProcedureSQLPLOptionList is called when exiting the procedureSQLPLOptionList production.
	ExitProcedureSQLPLOptionList(c *ProcedureSQLPLOptionListContext)

	// ExitVersionOption is called when exiting the versionOption production.
	ExitVersionOption(c *VersionOptionContext)

	// ExitCommitOnReturnOptionSQLPL is called when exiting the commitOnReturnOptionSQLPL production.
	ExitCommitOnReturnOptionSQLPL(c *CommitOnReturnOptionSQLPLContext)

	// ExitSchemaQualifierOption is called when exiting the schemaQualifierOption production.
	ExitSchemaQualifierOption(c *SchemaQualifierOptionContext)

	// ExitCurrentDataOption is called when exiting the currentDataOption production.
	ExitCurrentDataOption(c *CurrentDataOptionContext)

	// ExitImmediateWriteOption is called when exiting the immediateWriteOption production.
	ExitImmediateWriteOption(c *ImmediateWriteOptionContext)

	// ExitExplainOption is called when exiting the explainOption production.
	ExitExplainOption(c *ExplainOptionContext)

	// ExitReoptOption is called when exiting the reoptOption production.
	ExitReoptOption(c *ReoptOptionContext)

	// ExitPackageOwnerOption is called when exiting the packageOwnerOption production.
	ExitPackageOwnerOption(c *PackageOwnerOptionContext)

	// ExitDeferPrepareOption is called when exiting the deferPrepareOption production.
	ExitDeferPrepareOption(c *DeferPrepareOptionContext)

	// ExitDegreeOption is called when exiting the degreeOption production.
	ExitDegreeOption(c *DegreeOptionContext)

	// ExitDynamicRulesOption is called when exiting the dynamicRulesOption production.
	ExitDynamicRulesOption(c *DynamicRulesOptionContext)

	// ExitConcurrentAccessOption is called when exiting the concurrentAccessOption production.
	ExitConcurrentAccessOption(c *ConcurrentAccessOptionContext)

	// ExitApplicationEncodingOption is called when exiting the applicationEncodingOption production.
	ExitApplicationEncodingOption(c *ApplicationEncodingOptionContext)

	// ExitIsolationLevelOption is called when exiting the isolationLevelOption production.
	ExitIsolationLevelOption(c *IsolationLevelOptionContext)

	// ExitKeepDynamicOption is called when exiting the keepDynamicOption production.
	ExitKeepDynamicOption(c *KeepDynamicOptionContext)

	// ExitOpthintOption is called when exiting the opthintOption production.
	ExitOpthintOption(c *OpthintOptionContext)

	// ExitSqlPathOption is called when exiting the sqlPathOption production.
	ExitSqlPathOption(c *SqlPathOptionContext)

	// ExitSqlPathOptionItem is called when exiting the sqlPathOptionItem production.
	ExitSqlPathOptionItem(c *SqlPathOptionItemContext)

	// ExitQueryAccelerationOption is called when exiting the queryAccelerationOption production.
	ExitQueryAccelerationOption(c *QueryAccelerationOptionContext)

	// ExitQueryAccelerationOptionItem is called when exiting the queryAccelerationOptionItem production.
	ExitQueryAccelerationOptionItem(c *QueryAccelerationOptionItemContext)

	// ExitGetAccelArchiveOption is called when exiting the getAccelArchiveOption production.
	ExitGetAccelArchiveOption(c *GetAccelArchiveOptionContext)

	// ExitAccelerationOption is called when exiting the accelerationOption production.
	ExitAccelerationOption(c *AccelerationOptionContext)

	// ExitReleaseAtOption is called when exiting the releaseAtOption production.
	ExitReleaseAtOption(c *ReleaseAtOptionContext)

	// ExitBusinessTimeSensitiveOption is called when exiting the businessTimeSensitiveOption production.
	ExitBusinessTimeSensitiveOption(c *BusinessTimeSensitiveOptionContext)

	// ExitSystemTimeSensitiveOption is called when exiting the systemTimeSensitiveOption production.
	ExitSystemTimeSensitiveOption(c *SystemTimeSensitiveOptionContext)

	// ExitArchiveSensitiveOption is called when exiting the archiveSensitiveOption production.
	ExitArchiveSensitiveOption(c *ArchiveSensitiveOptionContext)

	// ExitApplcompatOption is called when exiting the applcompatOption production.
	ExitApplcompatOption(c *ApplcompatOptionContext)

	// ExitValidateOption is called when exiting the validateOption production.
	ExitValidateOption(c *ValidateOptionContext)

	// ExitRoundingOption is called when exiting the roundingOption production.
	ExitRoundingOption(c *RoundingOptionContext)

	// ExitRoundingOptionItem is called when exiting the roundingOptionItem production.
	ExitRoundingOptionItem(c *RoundingOptionItemContext)

	// ExitDateFormatOption is called when exiting the dateFormatOption production.
	ExitDateFormatOption(c *DateFormatOptionContext)

	// ExitDateTimeFormatOptionItem is called when exiting the dateTimeFormatOptionItem production.
	ExitDateTimeFormatOptionItem(c *DateTimeFormatOptionItemContext)

	// ExitTimeFormatOption is called when exiting the timeFormatOption production.
	ExitTimeFormatOption(c *TimeFormatOptionContext)

	// ExitDecimalOption is called when exiting the decimalOption production.
	ExitDecimalOption(c *DecimalOptionContext)

	// ExitForUpdateOption is called when exiting the forUpdateOption production.
	ExitForUpdateOption(c *ForUpdateOptionContext)

	// ExitConcentrateStatementsOption is called when exiting the concentrateStatementsOption production.
	ExitConcentrateStatementsOption(c *ConcentrateStatementsOptionContext)

	// ExitAcceleratorOption is called when exiting the acceleratorOption production.
	ExitAcceleratorOption(c *AcceleratorOptionContext)

	// ExitProcedureDataType is called when exiting the procedureDataType production.
	ExitProcedureDataType(c *ProcedureDataTypeContext)

	// ExitAlterSequenceOptionList is called when exiting the alterSequenceOptionList production.
	ExitAlterSequenceOptionList(c *AlterSequenceOptionListContext)

	// ExitCreateSequenceOptionList is called when exiting the createSequenceOptionList production.
	ExitCreateSequenceOptionList(c *CreateSequenceOptionListContext)

	// ExitAsTypeOption is called when exiting the asTypeOption production.
	ExitAsTypeOption(c *AsTypeOptionContext)

	// ExitStartOption is called when exiting the startOption production.
	ExitStartOption(c *StartOptionContext)

	// ExitRestartOption is called when exiting the restartOption production.
	ExitRestartOption(c *RestartOptionContext)

	// ExitIncrementOption is called when exiting the incrementOption production.
	ExitIncrementOption(c *IncrementOptionContext)

	// ExitMinvalueOption is called when exiting the minvalueOption production.
	ExitMinvalueOption(c *MinvalueOptionContext)

	// ExitMaxvalueOption is called when exiting the maxvalueOption production.
	ExitMaxvalueOption(c *MaxvalueOptionContext)

	// ExitCycleOption is called when exiting the cycleOption production.
	ExitCycleOption(c *CycleOptionContext)

	// ExitCacheOption is called when exiting the cacheOption production.
	ExitCacheOption(c *CacheOptionContext)

	// ExitOrderOption is called when exiting the orderOption production.
	ExitOrderOption(c *OrderOptionContext)

	// ExitKeyLabelOption is called when exiting the keyLabelOption production.
	ExitKeyLabelOption(c *KeyLabelOptionContext)

	// ExitDataclasOption is called when exiting the dataclasOption production.
	ExitDataclasOption(c *DataclasOptionContext)

	// ExitMgmtclasOption is called when exiting the mgmtclasOption production.
	ExitMgmtclasOption(c *MgmtclasOptionContext)

	// ExitStorclasOption is called when exiting the storclasOption production.
	ExitStorclasOption(c *StorclasOptionContext)

	// ExitAlterStogroupOptionList is called when exiting the alterStogroupOptionList production.
	ExitAlterStogroupOptionList(c *AlterStogroupOptionListContext)

	// ExitAlterTableOptionList is called when exiting the alterTableOptionList production.
	ExitAlterTableOptionList(c *AlterTableOptionListContext)

	// ExitAlterTablespaceOptionList is called when exiting the alterTablespaceOptionList production.
	ExitAlterTablespaceOptionList(c *AlterTablespaceOptionListContext)

	// ExitCreateTablespaceOptionList is called when exiting the createTablespaceOptionList production.
	ExitCreateTablespaceOptionList(c *CreateTablespaceOptionListContext)

	// ExitTrustedContextOptionList is called when exiting the trustedContextOptionList production.
	ExitTrustedContextOptionList(c *TrustedContextOptionListContext)

	// ExitDatabaseOptionList is called when exiting the databaseOptionList production.
	ExitDatabaseOptionList(c *DatabaseOptionListContext)

	// ExitCreateIndexOptionList is called when exiting the createIndexOptionList production.
	ExitCreateIndexOptionList(c *CreateIndexOptionListContext)

	// ExitCreateLobTablespaceOptionList is called when exiting the createLobTablespaceOptionList production.
	ExitCreateLobTablespaceOptionList(c *CreateLobTablespaceOptionListContext)

	// ExitInDatabaseOption is called when exiting the inDatabaseOption production.
	ExitInDatabaseOption(c *InDatabaseOptionContext)

	// ExitSegsizeOption is called when exiting the segsizeOption production.
	ExitSegsizeOption(c *SegsizeOptionContext)

	// ExitNumpartsOption is called when exiting the numpartsOption production.
	ExitNumpartsOption(c *NumpartsOptionContext)

	// ExitPartitionByGrowthSpecification is called when exiting the partitionByGrowthSpecification production.
	ExitPartitionByGrowthSpecification(c *PartitionByGrowthSpecificationContext)

	// ExitPartitionByRangeSpecification is called when exiting the partitionByRangeSpecification production.
	ExitPartitionByRangeSpecification(c *PartitionByRangeSpecificationContext)

	// ExitPartitionByRangePartitionPhrase is called when exiting the partitionByRangePartitionPhrase production.
	ExitPartitionByRangePartitionPhrase(c *PartitionByRangePartitionPhraseContext)

	// ExitInsertAlgorithmOption is called when exiting the insertAlgorithmOption production.
	ExitInsertAlgorithmOption(c *InsertAlgorithmOptionContext)

	// ExitMaxrowsOption is called when exiting the maxrowsOption production.
	ExitMaxrowsOption(c *MaxrowsOptionContext)

	// ExitMaxpartitionsOption is called when exiting the maxpartitionsOption production.
	ExitMaxpartitionsOption(c *MaxpartitionsOptionContext)

	// ExitUsingSpecification2 is called when exiting the usingSpecification2 production.
	ExitUsingSpecification2(c *UsingSpecification2Context)

	// ExitStogroupOptions is called when exiting the stogroupOptions production.
	ExitStogroupOptions(c *StogroupOptionsContext)

	// ExitXmlIndexSpecification is called when exiting the xmlIndexSpecification production.
	ExitXmlIndexSpecification(c *XmlIndexSpecificationContext)

	// ExitXmlPatternClause is called when exiting the xmlPatternClause production.
	ExitXmlPatternClause(c *XmlPatternClauseContext)

	// ExitAlterAttributesOptions is called when exiting the alterAttributesOptions production.
	ExitAlterAttributesOptions(c *AlterAttributesOptionsContext)

	// ExitAddAttributesOptions is called when exiting the addAttributesOptions production.
	ExitAddAttributesOptions(c *AddAttributesOptionsContext)

	// ExitDropAttributesOptions is called when exiting the dropAttributesOptions production.
	ExitDropAttributesOptions(c *DropAttributesOptionsContext)

	// ExitIncludeColumnPhrase is called when exiting the includeColumnPhrase production.
	ExitIncludeColumnPhrase(c *IncludeColumnPhraseContext)

	// ExitUserClause is called when exiting the userClause production.
	ExitUserClause(c *UserClauseContext)

	// ExitUserClauseAddOptions is called when exiting the userClauseAddOptions production.
	ExitUserClauseAddOptions(c *UserClauseAddOptionsContext)

	// ExitUserClauseReplaceOptions is called when exiting the userClauseReplaceOptions production.
	ExitUserClauseReplaceOptions(c *UserClauseReplaceOptionsContext)

	// ExitUserClauseDropOptions is called when exiting the userClauseDropOptions production.
	ExitUserClauseDropOptions(c *UserClauseDropOptionsContext)

	// ExitUseOptions is called when exiting the useOptions production.
	ExitUseOptions(c *UseOptionsContext)

	// ExitAlterPartitionClause is called when exiting the alterPartitionClause production.
	ExitAlterPartitionClause(c *AlterPartitionClauseContext)

	// ExitUsingBlock is called when exiting the usingBlock production.
	ExitUsingBlock(c *UsingBlockContext)

	// ExitFreeBlock is called when exiting the freeBlock production.
	ExitFreeBlock(c *FreeBlockContext)

	// ExitMoveTableClause is called when exiting the moveTableClause production.
	ExitMoveTableClause(c *MoveTableClauseContext)

	// ExitGbpcacheBlock is called when exiting the gbpcacheBlock production.
	ExitGbpcacheBlock(c *GbpcacheBlockContext)

	// ExitAliasDesignator is called when exiting the aliasDesignator production.
	ExitAliasDesignator(c *AliasDesignatorContext)

	// ExitMultipleColumnList is called when exiting the multipleColumnList production.
	ExitMultipleColumnList(c *MultipleColumnListContext)

	// ExitFunctionDesignator is called when exiting the functionDesignator production.
	ExitFunctionDesignator(c *FunctionDesignatorContext)

	// ExitParameterType is called when exiting the parameterType production.
	ExitParameterType(c *ParameterTypeContext)

	// ExitAlterTableColumnDefinitionOptionList1 is called when exiting the alterTableColumnDefinitionOptionList1 production.
	ExitAlterTableColumnDefinitionOptionList1(c *AlterTableColumnDefinitionOptionList1Context)

	// ExitAlterTableColumnDefinitionOptionList2 is called when exiting the alterTableColumnDefinitionOptionList2 production.
	ExitAlterTableColumnDefinitionOptionList2(c *AlterTableColumnDefinitionOptionList2Context)

	// ExitColumnConstraint is called when exiting the columnConstraint production.
	ExitColumnConstraint(c *ColumnConstraintContext)

	// ExitGeneratedClause is called when exiting the generatedClause production.
	ExitGeneratedClause(c *GeneratedClauseContext)

	// ExitGeneratedClause2 is called when exiting the generatedClause2 production.
	ExitGeneratedClause2(c *GeneratedClause2Context)

	// ExitAsIdentityClause is called when exiting the asIdentityClause production.
	ExitAsIdentityClause(c *AsIdentityClauseContext)

	// ExitAsIdentityClauseOptionList is called when exiting the asIdentityClauseOptionList production.
	ExitAsIdentityClauseOptionList(c *AsIdentityClauseOptionListContext)

	// ExitAsRowChangeTimestampClause is called when exiting the asRowChangeTimestampClause production.
	ExitAsRowChangeTimestampClause(c *AsRowChangeTimestampClauseContext)

	// ExitAsRowTransactionStartIDClause is called when exiting the asRowTransactionStartIDClause production.
	ExitAsRowTransactionStartIDClause(c *AsRowTransactionStartIDClauseContext)

	// ExitAsRowTransactionTimestampClause is called when exiting the asRowTransactionTimestampClause production.
	ExitAsRowTransactionTimestampClause(c *AsRowTransactionTimestampClauseContext)

	// ExitAsGeneratedExpressionClause is called when exiting the asGeneratedExpressionClause production.
	ExitAsGeneratedExpressionClause(c *AsGeneratedExpressionClauseContext)

	// ExitNonDeterministicExpression is called when exiting the nonDeterministicExpression production.
	ExitNonDeterministicExpression(c *NonDeterministicExpressionContext)

	// ExitNonDeterministicExpressionSessionVariable is called when exiting the nonDeterministicExpressionSessionVariable production.
	ExitNonDeterministicExpressionSessionVariable(c *NonDeterministicExpressionSessionVariableContext)

	// ExitColumnAlteration is called when exiting the columnAlteration production.
	ExitColumnAlteration(c *ColumnAlterationContext)

	// ExitColumnAlterationOptionList is called when exiting the columnAlterationOptionList production.
	ExitColumnAlterationOptionList(c *ColumnAlterationOptionListContext)

	// ExitAlteredDataType is called when exiting the alteredDataType production.
	ExitAlteredDataType(c *AlteredDataTypeContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitBuiltInType is called when exiting the builtInType production.
	ExitBuiltInType(c *BuiltInTypeContext)

	// ExitSequenceDataType is called when exiting the sequenceDataType production.
	ExitSequenceDataType(c *SequenceDataTypeContext)

	// ExitSequenceBuiltInType is called when exiting the sequenceBuiltInType production.
	ExitSequenceBuiltInType(c *SequenceBuiltInTypeContext)

	// ExitSqlDataType is called when exiting the sqlDataType production.
	ExitSqlDataType(c *SqlDataTypeContext)

	// ExitXmlTypeModifier is called when exiting the xmlTypeModifier production.
	ExitXmlTypeModifier(c *XmlTypeModifierContext)

	// ExitXmlSchemaSpecification is called when exiting the xmlSchemaSpecification production.
	ExitXmlSchemaSpecification(c *XmlSchemaSpecificationContext)

	// ExitXmlElementName is called when exiting the xmlElementName production.
	ExitXmlElementName(c *XmlElementNameContext)

	// ExitPiName is called when exiting the piName production.
	ExitPiName(c *PiNameContext)

	// ExitRegisteredXmlSchemaName is called when exiting the registeredXmlSchemaName production.
	ExitRegisteredXmlSchemaName(c *RegisteredXmlSchemaNameContext)

	// ExitTargetNamespace is called when exiting the targetNamespace production.
	ExitTargetNamespace(c *TargetNamespaceContext)

	// ExitSchemaLocation is called when exiting the schemaLocation production.
	ExitSchemaLocation(c *SchemaLocationContext)

	// ExitIdentityAlteration is called when exiting the identityAlteration production.
	ExitIdentityAlteration(c *IdentityAlterationContext)

	// ExitUniqueConstraint is called when exiting the uniqueConstraint production.
	ExitUniqueConstraint(c *UniqueConstraintContext)

	// ExitReferentialConstraint is called when exiting the referentialConstraint production.
	ExitReferentialConstraint(c *ReferentialConstraintContext)

	// ExitReferencesClause is called when exiting the referencesClause production.
	ExitReferencesClause(c *ReferencesClauseContext)

	// ExitCheckConstraint is called when exiting the checkConstraint production.
	ExitCheckConstraint(c *CheckConstraintContext)

	// ExitPartitioningClause is called when exiting the partitioningClause production.
	ExitPartitioningClause(c *PartitioningClauseContext)

	// ExitPartitionExpression is called when exiting the partitionExpression production.
	ExitPartitionExpression(c *PartitionExpressionContext)

	// ExitPartitionLimitKey is called when exiting the partitionLimitKey production.
	ExitPartitionLimitKey(c *PartitionLimitKeyContext)

	// ExitPartitioningPhrase is called when exiting the partitioningPhrase production.
	ExitPartitioningPhrase(c *PartitioningPhraseContext)

	// ExitPartitionHashSpace is called when exiting the partitionHashSpace production.
	ExitPartitionHashSpace(c *PartitionHashSpaceContext)

	// ExitAlterHashOrganization is called when exiting the alterHashOrganization production.
	ExitAlterHashOrganization(c *AlterHashOrganizationContext)

	// ExitPartitioningClauseElement is called when exiting the partitioningClauseElement production.
	ExitPartitioningClauseElement(c *PartitioningClauseElementContext)

	// ExitPartitionClause is called when exiting the partitionClause production.
	ExitPartitionClause(c *PartitionClauseContext)

	// ExitRotatePartitionClause is called when exiting the rotatePartitionClause production.
	ExitRotatePartitionClause(c *RotatePartitionClauseContext)

	// ExitExtraRowOption is called when exiting the extraRowOption production.
	ExitExtraRowOption(c *ExtraRowOptionContext)

	// ExitMaterializedQueryDefinition is called when exiting the materializedQueryDefinition production.
	ExitMaterializedQueryDefinition(c *MaterializedQueryDefinitionContext)

	// ExitMaterializedQueryAlteration is called when exiting the materializedQueryAlteration production.
	ExitMaterializedQueryAlteration(c *MaterializedQueryAlterationContext)

	// ExitRefreshableTableOptions is called when exiting the refreshableTableOptions production.
	ExitRefreshableTableOptions(c *RefreshableTableOptionsContext)

	// ExitDataInitiallyDeferredPhrase is called when exiting the dataInitiallyDeferredPhrase production.
	ExitDataInitiallyDeferredPhrase(c *DataInitiallyDeferredPhraseContext)

	// ExitRefreshDeferredPhrase is called when exiting the refreshDeferredPhrase production.
	ExitRefreshDeferredPhrase(c *RefreshDeferredPhraseContext)

	// ExitRefreshableTableOptionsList is called when exiting the refreshableTableOptionsList production.
	ExitRefreshableTableOptionsList(c *RefreshableTableOptionsListContext)

	// ExitMaterializedQueryTableAlteration is called when exiting the materializedQueryTableAlteration production.
	ExitMaterializedQueryTableAlteration(c *MaterializedQueryTableAlterationContext)

	// ExitPeriodDefinition is called when exiting the periodDefinition production.
	ExitPeriodDefinition(c *PeriodDefinitionContext)

	// ExitAlterTableColumnDefinition is called when exiting the alterTableColumnDefinition production.
	ExitAlterTableColumnDefinition(c *AlterTableColumnDefinitionContext)

	// ExitExternalProgramName is called when exiting the externalProgramName production.
	ExitExternalProgramName(c *ExternalProgramNameContext)

	// ExitPackagePath is called when exiting the packagePath production.
	ExitPackagePath(c *PackagePathContext)

	// ExitCollectionID is called when exiting the collectionID production.
	ExitCollectionID(c *CollectionIDContext)

	// ExitRunTimeOptions is called when exiting the runTimeOptions production.
	ExitRunTimeOptions(c *RunTimeOptionsContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitKeyExpression is called when exiting the keyExpression production.
	ExitKeyExpression(c *KeyExpressionContext)

	// ExitRowChangeExpression is called when exiting the rowChangeExpression production.
	ExitRowChangeExpression(c *RowChangeExpressionContext)

	// ExitSequenceReference is called when exiting the sequenceReference production.
	ExitSequenceReference(c *SequenceReferenceContext)

	// ExitFunctionInvocation is called when exiting the functionInvocation production.
	ExitFunctionInvocation(c *FunctionInvocationContext)

	// ExitScalarFunctionInvocation is called when exiting the scalarFunctionInvocation production.
	ExitScalarFunctionInvocation(c *ScalarFunctionInvocationContext)

	// ExitAggregateFunctionInvocation is called when exiting the aggregateFunctionInvocation production.
	ExitAggregateFunctionInvocation(c *AggregateFunctionInvocationContext)

	// ExitRegressionFunctionInvocation is called when exiting the regressionFunctionInvocation production.
	ExitRegressionFunctionInvocation(c *RegressionFunctionInvocationContext)

	// ExitExternalFunctionInvocation is called when exiting the externalFunctionInvocation production.
	ExitExternalFunctionInvocation(c *ExternalFunctionInvocationContext)

	// ExitLabeledDuration is called when exiting the labeledDuration production.
	ExitLabeledDuration(c *LabeledDurationContext)

	// ExitDurationSuffix is called when exiting the durationSuffix production.
	ExitDurationSuffix(c *DurationSuffixContext)

	// ExitXmlCastSpecification is called when exiting the xmlCastSpecification production.
	ExitXmlCastSpecification(c *XmlCastSpecificationContext)

	// ExitXmlParseSpecification is called when exiting the xmlParseSpecification production.
	ExitXmlParseSpecification(c *XmlParseSpecificationContext)

	// ExitArrayElementSpecification is called when exiting the arrayElementSpecification production.
	ExitArrayElementSpecification(c *ArrayElementSpecificationContext)

	// ExitArrayIndex is called when exiting the arrayIndex production.
	ExitArrayIndex(c *ArrayIndexContext)

	// ExitArrayConstructor is called when exiting the arrayConstructor production.
	ExitArrayConstructor(c *ArrayConstructorContext)

	// ExitOlapSpecification is called when exiting the olapSpecification production.
	ExitOlapSpecification(c *OlapSpecificationContext)

	// ExitOrderedOlapSpecification is called when exiting the orderedOlapSpecification production.
	ExitOrderedOlapSpecification(c *OrderedOlapSpecificationContext)

	// ExitOlapSpecificationFunction is called when exiting the olapSpecificationFunction production.
	ExitOlapSpecificationFunction(c *OlapSpecificationFunctionContext)

	// ExitLagFunction is called when exiting the lagFunction production.
	ExitLagFunction(c *LagFunctionContext)

	// ExitLeadFunction is called when exiting the leadFunction production.
	ExitLeadFunction(c *LeadFunctionContext)

	// ExitRespectNullsClause is called when exiting the respectNullsClause production.
	ExitRespectNullsClause(c *RespectNullsClauseContext)

	// ExitWindowPartitionClause is called when exiting the windowPartitionClause production.
	ExitWindowPartitionClause(c *WindowPartitionClauseContext)

	// ExitWindowOrderClause is called when exiting the windowOrderClause production.
	ExitWindowOrderClause(c *WindowOrderClauseContext)

	// ExitWindowOrderClauseQualifier is called when exiting the windowOrderClauseQualifier production.
	ExitWindowOrderClauseQualifier(c *WindowOrderClauseQualifierContext)

	// ExitNumberingSpecification is called when exiting the numberingSpecification production.
	ExitNumberingSpecification(c *NumberingSpecificationContext)

	// ExitAggregationSpecification is called when exiting the aggregationSpecification production.
	ExitAggregationSpecification(c *AggregationSpecificationContext)

	// ExitAggregateFunction is called when exiting the aggregateFunction production.
	ExitAggregateFunction(c *AggregateFunctionContext)

	// ExitRegressionFunction is called when exiting the regressionFunction production.
	ExitRegressionFunction(c *RegressionFunctionContext)

	// ExitOlapColumnFunction is called when exiting the olapColumnFunction production.
	ExitOlapColumnFunction(c *OlapColumnFunctionContext)

	// ExitFirstValueFunction is called when exiting the firstValueFunction production.
	ExitFirstValueFunction(c *FirstValueFunctionContext)

	// ExitLastValueFunction is called when exiting the lastValueFunction production.
	ExitLastValueFunction(c *LastValueFunctionContext)

	// ExitNthValueFunction is called when exiting the nthValueFunction production.
	ExitNthValueFunction(c *NthValueFunctionContext)

	// ExitRatioToReportFunction is called when exiting the ratioToReportFunction production.
	ExitRatioToReportFunction(c *RatioToReportFunctionContext)

	// ExitListaggFunction is called when exiting the listaggFunction production.
	ExitListaggFunction(c *ListaggFunctionContext)

	// ExitArrayaggFunction is called when exiting the arrayaggFunction production.
	ExitArrayaggFunction(c *ArrayaggFunctionContext)

	// ExitArrayaggOrdinaryFunction is called when exiting the arrayaggOrdinaryFunction production.
	ExitArrayaggOrdinaryFunction(c *ArrayaggOrdinaryFunctionContext)

	// ExitArrayaggAssociativeFunction is called when exiting the arrayaggAssociativeFunction production.
	ExitArrayaggAssociativeFunction(c *ArrayaggAssociativeFunctionContext)

	// ExitCorrelationFunction is called when exiting the correlationFunction production.
	ExitCorrelationFunction(c *CorrelationFunctionContext)

	// ExitCovarianceFunction is called when exiting the covarianceFunction production.
	ExitCovarianceFunction(c *CovarianceFunctionContext)

	// ExitCovarianceSampFunction is called when exiting the covarianceSampFunction production.
	ExitCovarianceSampFunction(c *CovarianceSampFunctionContext)

	// ExitCumeDistFunction is called when exiting the cumeDistFunction production.
	ExitCumeDistFunction(c *CumeDistFunctionContext)

	// ExitPercentileContFunction is called when exiting the percentileContFunction production.
	ExitPercentileContFunction(c *PercentileContFunctionContext)

	// ExitPercentileDiscFunction is called when exiting the percentileDiscFunction production.
	ExitPercentileDiscFunction(c *PercentileDiscFunctionContext)

	// ExitPercentRankFunction is called when exiting the percentRankFunction production.
	ExitPercentRankFunction(c *PercentRankFunctionContext)

	// ExitXmlaggFunction is called when exiting the xmlaggFunction production.
	ExitXmlaggFunction(c *XmlaggFunctionContext)

	// ExitXmlaggOrderByClause is called when exiting the xmlaggOrderByClause production.
	ExitXmlaggOrderByClause(c *XmlaggOrderByClauseContext)

	// ExitXmlaggOrderByOption is called when exiting the xmlaggOrderByOption production.
	ExitXmlaggOrderByOption(c *XmlaggOrderByOptionContext)

	// ExitAggregateOrderByClause is called when exiting the aggregateOrderByClause production.
	ExitAggregateOrderByClause(c *AggregateOrderByClauseContext)

	// ExitAggregateOrderByOption is called when exiting the aggregateOrderByOption production.
	ExitAggregateOrderByOption(c *AggregateOrderByOptionContext)

	// ExitWindowAggregationGroupClause is called when exiting the windowAggregationGroupClause production.
	ExitWindowAggregationGroupClause(c *WindowAggregationGroupClauseContext)

	// ExitGroupStart is called when exiting the groupStart production.
	ExitGroupStart(c *GroupStartContext)

	// ExitGroupBetween is called when exiting the groupBetween production.
	ExitGroupBetween(c *GroupBetweenContext)

	// ExitGroupEnd is called when exiting the groupEnd production.
	ExitGroupEnd(c *GroupEndContext)

	// ExitGroupBound1 is called when exiting the groupBound1 production.
	ExitGroupBound1(c *GroupBound1Context)

	// ExitGroupBound2 is called when exiting the groupBound2 production.
	ExitGroupBound2(c *GroupBound2Context)

	// ExitUnboundedPreceding is called when exiting the unboundedPreceding production.
	ExitUnboundedPreceding(c *UnboundedPrecedingContext)

	// ExitUnboundedFollowing is called when exiting the unboundedFollowing production.
	ExitUnboundedFollowing(c *UnboundedFollowingContext)

	// ExitBoundedPreceding is called when exiting the boundedPreceding production.
	ExitBoundedPreceding(c *BoundedPrecedingContext)

	// ExitBoundedFollowing is called when exiting the boundedFollowing production.
	ExitBoundedFollowing(c *BoundedFollowingContext)

	// ExitCurrentRow is called when exiting the currentRow production.
	ExitCurrentRow(c *CurrentRowContext)

	// ExitScalarFunction is called when exiting the scalarFunction production.
	ExitScalarFunction(c *ScalarFunctionContext)

	// ExitTableFunction is called when exiting the tableFunction production.
	ExitTableFunction(c *TableFunctionContext)

	// ExitSpecialRegister is called when exiting the specialRegister production.
	ExitSpecialRegister(c *SpecialRegisterContext)

	// ExitAiAnalogyFunction is called when exiting the aiAnalogyFunction production.
	ExitAiAnalogyFunction(c *AiAnalogyFunctionContext)

	// ExitAiFunctionExpression is called when exiting the aiFunctionExpression production.
	ExitAiFunctionExpression(c *AiFunctionExpressionContext)

	// ExitAiAnalogyFunctionSource is called when exiting the aiAnalogyFunctionSource production.
	ExitAiAnalogyFunctionSource(c *AiAnalogyFunctionSourceContext)

	// ExitAiAnalogyFunctionTarget is called when exiting the aiAnalogyFunctionTarget production.
	ExitAiAnalogyFunctionTarget(c *AiAnalogyFunctionTargetContext)

	// ExitAiAnalogyFunctionSource1 is called when exiting the aiAnalogyFunctionSource1 production.
	ExitAiAnalogyFunctionSource1(c *AiAnalogyFunctionSource1Context)

	// ExitAiAnalogyFunctionSource2 is called when exiting the aiAnalogyFunctionSource2 production.
	ExitAiAnalogyFunctionSource2(c *AiAnalogyFunctionSource2Context)

	// ExitAiAnalogyFunctionTarget1 is called when exiting the aiAnalogyFunctionTarget1 production.
	ExitAiAnalogyFunctionTarget1(c *AiAnalogyFunctionTarget1Context)

	// ExitAiAnalogyFunctionTarget2 is called when exiting the aiAnalogyFunctionTarget2 production.
	ExitAiAnalogyFunctionTarget2(c *AiAnalogyFunctionTarget2Context)

	// ExitAiSemanticClusterFunction is called when exiting the aiSemanticClusterFunction production.
	ExitAiSemanticClusterFunction(c *AiSemanticClusterFunctionContext)

	// ExitAiSemanticClusterMemberExpression is called when exiting the aiSemanticClusterMemberExpression production.
	ExitAiSemanticClusterMemberExpression(c *AiSemanticClusterMemberExpressionContext)

	// ExitAiSemanticClusterClusteringExpression is called when exiting the aiSemanticClusterClusteringExpression production.
	ExitAiSemanticClusterClusteringExpression(c *AiSemanticClusterClusteringExpressionContext)

	// ExitAiSimilarityFunction is called when exiting the aiSimilarityFunction production.
	ExitAiSimilarityFunction(c *AiSimilarityFunctionContext)

	// ExitAiSimilarityExpression is called when exiting the aiSimilarityExpression production.
	ExitAiSimilarityExpression(c *AiSimilarityExpressionContext)

	// ExitAiSimilarityExpression1 is called when exiting the aiSimilarityExpression1 production.
	ExitAiSimilarityExpression1(c *AiSimilarityExpression1Context)

	// ExitAiSimilarityExpression2 is called when exiting the aiSimilarityExpression2 production.
	ExitAiSimilarityExpression2(c *AiSimilarityExpression2Context)

	// ExitXmlelementFunction is called when exiting the xmlelementFunction production.
	ExitXmlelementFunction(c *XmlelementFunctionContext)

	// ExitXmlforestFunction is called when exiting the xmlforestFunction production.
	ExitXmlforestFunction(c *XmlforestFunctionContext)

	// ExitXmlmodifyFunction is called when exiting the xmlmodifyFunction production.
	ExitXmlmodifyFunction(c *XmlmodifyFunctionContext)

	// ExitXmlpiFunction is called when exiting the xmlpiFunction production.
	ExitXmlpiFunction(c *XmlpiFunctionContext)

	// ExitXmlqueryFunction is called when exiting the xmlqueryFunction production.
	ExitXmlqueryFunction(c *XmlqueryFunctionContext)

	// ExitXmlattributesFunction is called when exiting the xmlattributesFunction production.
	ExitXmlattributesFunction(c *XmlattributesFunctionContext)

	// ExitXmlserializeFunction is called when exiting the xmlserializeFunction production.
	ExitXmlserializeFunction(c *XmlserializeFunctionContext)

	// ExitXmlnamespaceFunction is called when exiting the xmlnamespaceFunction production.
	ExitXmlnamespaceFunction(c *XmlnamespaceFunctionContext)

	// ExitXmlnamespaceOption is called when exiting the xmlnamespaceOption production.
	ExitXmlnamespaceOption(c *XmlnamespaceOptionContext)

	// ExitXmlserializeFunctionOptions is called when exiting the xmlserializeFunctionOptions production.
	ExitXmlserializeFunctionOptions(c *XmlserializeFunctionOptionsContext)

	// ExitXmlFunctionOptionClause is called when exiting the xmlFunctionOptionClause production.
	ExitXmlFunctionOptionClause(c *XmlFunctionOptionClauseContext)

	// ExitXmlFunctionOption is called when exiting the xmlFunctionOption production.
	ExitXmlFunctionOption(c *XmlFunctionOptionContext)

	// ExitElementContentExpression is called when exiting the elementContentExpression production.
	ExitElementContentExpression(c *ElementContentExpressionContext)

	// ExitXqueryExpressionConstant is called when exiting the xqueryExpressionConstant production.
	ExitXqueryExpressionConstant(c *XqueryExpressionConstantContext)

	// ExitXqueryArgument is called when exiting the xqueryArgument production.
	ExitXqueryArgument(c *XqueryArgumentContext)

	// ExitXmltableFunctionSpecification is called when exiting the xmltableFunctionSpecification production.
	ExitXmltableFunctionSpecification(c *XmltableFunctionSpecificationContext)

	// ExitRowXqueryExpressionConstant is called when exiting the rowXqueryExpressionConstant production.
	ExitRowXqueryExpressionConstant(c *RowXqueryExpressionConstantContext)

	// ExitRowXqueryArgument is called when exiting the rowXqueryArgument production.
	ExitRowXqueryArgument(c *RowXqueryArgumentContext)

	// ExitXqueryContextItemExpression is called when exiting the xqueryContextItemExpression production.
	ExitXqueryContextItemExpression(c *XqueryContextItemExpressionContext)

	// ExitXqueryVariableExpression is called when exiting the xqueryVariableExpression production.
	ExitXqueryVariableExpression(c *XqueryVariableExpressionContext)

	// ExitXmlTableRegularColumnDefinition is called when exiting the xmlTableRegularColumnDefinition production.
	ExitXmlTableRegularColumnDefinition(c *XmlTableRegularColumnDefinitionContext)

	// ExitDefaultClause is called when exiting the defaultClause production.
	ExitDefaultClause(c *DefaultClauseContext)

	// ExitDefaultClause1 is called when exiting the defaultClause1 production.
	ExitDefaultClause1(c *DefaultClause1Context)

	// ExitDefaultClauseAllowables is called when exiting the defaultClauseAllowables production.
	ExitDefaultClauseAllowables(c *DefaultClauseAllowablesContext)

	// ExitDistinctTypeCastFunctionName is called when exiting the distinctTypeCastFunctionName production.
	ExitDistinctTypeCastFunctionName(c *DistinctTypeCastFunctionNameContext)

	// ExitColumnXqueryExpressionConstant is called when exiting the columnXqueryExpressionConstant production.
	ExitColumnXqueryExpressionConstant(c *ColumnXqueryExpressionConstantContext)

	// ExitXmlTableOrdinalityColumnDefinition is called when exiting the xmlTableOrdinalityColumnDefinition production.
	ExitXmlTableOrdinalityColumnDefinition(c *XmlTableOrdinalityColumnDefinitionContext)

	// ExitXmlnamespacesDeclaration is called when exiting the xmlnamespacesDeclaration production.
	ExitXmlnamespacesDeclaration(c *XmlnamespacesDeclarationContext)

	// ExitXmlnamespacesFunctionSpecification is called when exiting the xmlnamespacesFunctionSpecification production.
	ExitXmlnamespacesFunctionSpecification(c *XmlnamespacesFunctionSpecificationContext)

	// ExitXmlnamespacesFunctionArguments is called when exiting the xmlnamespacesFunctionArguments production.
	ExitXmlnamespacesFunctionArguments(c *XmlnamespacesFunctionArgumentsContext)

	// ExitNamespaceUri is called when exiting the namespaceUri production.
	ExitNamespaceUri(c *NamespaceUriContext)

	// ExitNamespacePrefix is called when exiting the namespacePrefix production.
	ExitNamespacePrefix(c *NamespacePrefixContext)

	// ExitTimeZoneSpecificExpression is called when exiting the timeZoneSpecificExpression production.
	ExitTimeZoneSpecificExpression(c *TimeZoneSpecificExpressionContext)

	// ExitTimeZoneExpressionSubset is called when exiting the timeZoneExpressionSubset production.
	ExitTimeZoneExpressionSubset(c *TimeZoneExpressionSubsetContext)

	// ExitCaseExpression is called when exiting the caseExpression production.
	ExitCaseExpression(c *CaseExpressionContext)

	// ExitResultExpression is called when exiting the resultExpression production.
	ExitResultExpression(c *ResultExpressionContext)

	// ExitSearchedWhenClause is called when exiting the searchedWhenClause production.
	ExitSearchedWhenClause(c *SearchedWhenClauseContext)

	// ExitSimpleWhenClause is called when exiting the simpleWhenClause production.
	ExitSimpleWhenClause(c *SimpleWhenClauseContext)

	// ExitSearchCondition is called when exiting the searchCondition production.
	ExitSearchCondition(c *SearchConditionContext)

	// ExitCheckCondition is called when exiting the checkCondition production.
	ExitCheckCondition(c *CheckConditionContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitBasicPredicate is called when exiting the basicPredicate production.
	ExitBasicPredicate(c *BasicPredicateContext)

	// ExitRowValueExpression is called when exiting the rowValueExpression production.
	ExitRowValueExpression(c *RowValueExpressionContext)

	// ExitQuantifiedPredicate is called when exiting the quantifiedPredicate production.
	ExitQuantifiedPredicate(c *QuantifiedPredicateContext)

	// ExitArrayExistsPredicate is called when exiting the arrayExistsPredicate production.
	ExitArrayExistsPredicate(c *ArrayExistsPredicateContext)

	// ExitBetweenPredicate is called when exiting the betweenPredicate production.
	ExitBetweenPredicate(c *BetweenPredicateContext)

	// ExitDistinctPredicate is called when exiting the distinctPredicate production.
	ExitDistinctPredicate(c *DistinctPredicateContext)

	// ExitExistsPredicate is called when exiting the existsPredicate production.
	ExitExistsPredicate(c *ExistsPredicateContext)

	// ExitInPredicate is called when exiting the inPredicate production.
	ExitInPredicate(c *InPredicateContext)

	// ExitLikePredicate is called when exiting the likePredicate production.
	ExitLikePredicate(c *LikePredicateContext)

	// ExitNullPredicate is called when exiting the nullPredicate production.
	ExitNullPredicate(c *NullPredicateContext)

	// ExitXmlExistsPredicate is called when exiting the xmlExistsPredicate production.
	ExitXmlExistsPredicate(c *XmlExistsPredicateContext)

	// ExitArrayExpression is called when exiting the arrayExpression production.
	ExitArrayExpression(c *ArrayExpressionContext)

	// ExitCastSpecification is called when exiting the castSpecification production.
	ExitCastSpecification(c *CastSpecificationContext)

	// ExitParameterMarker is called when exiting the parameterMarker production.
	ExitParameterMarker(c *ParameterMarkerContext)

	// ExitCastDataType is called when exiting the castDataType production.
	ExitCastDataType(c *CastDataTypeContext)

	// ExitCastBuiltInType is called when exiting the castBuiltInType production.
	ExitCastBuiltInType(c *CastBuiltInTypeContext)

	// ExitIntegerInParens is called when exiting the integerInParens production.
	ExitIntegerInParens(c *IntegerInParensContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)

	// ExitCcsidQualifier is called when exiting the ccsidQualifier production.
	ExitCcsidQualifier(c *CcsidQualifierContext)

	// ExitForDataQualifier is called when exiting the forDataQualifier production.
	ExitForDataQualifier(c *ForDataQualifierContext)

	// ExitDistinctTypeName is called when exiting the distinctTypeName production.
	ExitDistinctTypeName(c *DistinctTypeNameContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitCcsidValue is called when exiting the ccsidValue production.
	ExitCcsidValue(c *CcsidValueContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitSourceColumnName is called when exiting the sourceColumnName production.
	ExitSourceColumnName(c *SourceColumnNameContext)

	// ExitTargetColumnName is called when exiting the targetColumnName production.
	ExitTargetColumnName(c *TargetColumnNameContext)

	// ExitBuildColumnName is called when exiting the buildColumnName production.
	ExitBuildColumnName(c *BuildColumnNameContext)

	// ExitBeginColumnName is called when exiting the beginColumnName production.
	ExitBeginColumnName(c *BeginColumnNameContext)

	// ExitEndColumnName is called when exiting the endColumnName production.
	ExitEndColumnName(c *EndColumnNameContext)

	// ExitCorrelationName is called when exiting the correlationName production.
	ExitCorrelationName(c *CorrelationNameContext)

	// ExitLocationName is called when exiting the locationName production.
	ExitLocationName(c *LocationNameContext)

	// ExitSchemaName is called when exiting the schemaName production.
	ExitSchemaName(c *SchemaNameContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitAlterTableName is called when exiting the alterTableName production.
	ExitAlterTableName(c *AlterTableNameContext)

	// ExitAuxTableName is called when exiting the auxTableName production.
	ExitAuxTableName(c *AuxTableNameContext)

	// ExitHistoryTableName is called when exiting the historyTableName production.
	ExitHistoryTableName(c *HistoryTableNameContext)

	// ExitCloneTableName is called when exiting the cloneTableName production.
	ExitCloneTableName(c *CloneTableNameContext)

	// ExitArchiveTableName is called when exiting the archiveTableName production.
	ExitArchiveTableName(c *ArchiveTableNameContext)

	// ExitViewName is called when exiting the viewName production.
	ExitViewName(c *ViewNameContext)

	// ExitProgramName is called when exiting the programName production.
	ExitProgramName(c *ProgramNameContext)

	// ExitPackageName is called when exiting the packageName production.
	ExitPackageName(c *PackageNameContext)

	// ExitPlanName is called when exiting the planName production.
	ExitPlanName(c *PlanNameContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitVariableName is called when exiting the variableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitArrayTypeName is called when exiting the arrayTypeName production.
	ExitArrayTypeName(c *ArrayTypeNameContext)

	// ExitJarName is called when exiting the jarName production.
	ExitJarName(c *JarNameContext)

	// ExitSavepointName is called when exiting the savepointName production.
	ExitSavepointName(c *SavepointNameContext)

	// ExitAliasName is called when exiting the aliasName production.
	ExitAliasName(c *AliasNameContext)

	// ExitConstraintName is called when exiting the constraintName production.
	ExitConstraintName(c *ConstraintNameContext)

	// ExitRoutineVersionID is called when exiting the routineVersionID production.
	ExitRoutineVersionID(c *RoutineVersionIDContext)

	// ExitVersionID is called when exiting the versionID production.
	ExitVersionID(c *VersionIDContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitMaskName is called when exiting the maskName production.
	ExitMaskName(c *MaskNameContext)

	// ExitPermissionName is called when exiting the permissionName production.
	ExitPermissionName(c *PermissionNameContext)

	// ExitProcedureName is called when exiting the procedureName production.
	ExitProcedureName(c *ProcedureNameContext)

	// ExitSequenceName is called when exiting the sequenceName production.
	ExitSequenceName(c *SequenceNameContext)

	// ExitMemberName is called when exiting the memberName production.
	ExitMemberName(c *MemberNameContext)

	// ExitDatabaseName is called when exiting the databaseName production.
	ExitDatabaseName(c *DatabaseNameContext)

	// ExitTablespaceName is called when exiting the tablespaceName production.
	ExitTablespaceName(c *TablespaceNameContext)

	// ExitAcceleratorName is called when exiting the acceleratorName production.
	ExitAcceleratorName(c *AcceleratorNameContext)

	// ExitCatalogName is called when exiting the catalogName production.
	ExitCatalogName(c *CatalogNameContext)

	// ExitTriggerName is called when exiting the triggerName production.
	ExitTriggerName(c *TriggerNameContext)

	// ExitContextName is called when exiting the contextName production.
	ExitContextName(c *ContextNameContext)

	// ExitAuthorizationName is called when exiting the authorizationName production.
	ExitAuthorizationName(c *AuthorizationNameContext)

	// ExitProfileName is called when exiting the profileName production.
	ExitProfileName(c *ProfileNameContext)

	// ExitRoleName is called when exiting the roleName production.
	ExitRoleName(c *RoleNameContext)

	// ExitSeclabelName is called when exiting the seclabelName production.
	ExitSeclabelName(c *SeclabelNameContext)

	// ExitParameterName is called when exiting the parameterName production.
	ExitParameterName(c *ParameterNameContext)

	// ExitAddressValue is called when exiting the addressValue production.
	ExitAddressValue(c *AddressValueContext)

	// ExitJobnameValue is called when exiting the jobnameValue production.
	ExitJobnameValue(c *JobnameValueContext)

	// ExitServauthValue is called when exiting the servauthValue production.
	ExitServauthValue(c *ServauthValueContext)

	// ExitEncryptionValue is called when exiting the encryptionValue production.
	ExitEncryptionValue(c *EncryptionValueContext)

	// ExitBpName is called when exiting the bpName production.
	ExitBpName(c *BpNameContext)

	// ExitStogroupName is called when exiting the stogroupName production.
	ExitStogroupName(c *StogroupNameContext)

	// ExitDcName is called when exiting the dcName production.
	ExitDcName(c *DcNameContext)

	// ExitMcName is called when exiting the mcName production.
	ExitMcName(c *McNameContext)

	// ExitScName is called when exiting the scName production.
	ExitScName(c *ScNameContext)

	// ExitVolumeID is called when exiting the volumeID production.
	ExitVolumeID(c *VolumeIDContext)

	// ExitKeyLabelName is called when exiting the keyLabelName production.
	ExitKeyLabelName(c *KeyLabelNameContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitSpecificName is called when exiting the specificName production.
	ExitSpecificName(c *SpecificNameContext)

	// ExitHostLabel is called when exiting the hostLabel production.
	ExitHostLabel(c *HostLabelContext)

	// ExitHostVariable is called when exiting the hostVariable production.
	ExitHostVariable(c *HostVariableContext)

	// ExitHostIdentifier is called when exiting the hostIdentifier production.
	ExitHostIdentifier(c *HostIdentifierContext)

	// ExitHostStructure is called when exiting the hostStructure production.
	ExitHostStructure(c *HostStructureContext)

	// ExitNullIndicator is called when exiting the nullIndicator production.
	ExitNullIndicator(c *NullIndicatorContext)

	// ExitNullIndicatorStructure is called when exiting the nullIndicatorStructure production.
	ExitNullIndicatorStructure(c *NullIndicatorStructureContext)

	// ExitGlobalVariableName is called when exiting the globalVariableName production.
	ExitGlobalVariableName(c *GlobalVariableNameContext)

	// ExitSqlParameterName is called when exiting the sqlParameterName production.
	ExitSqlParameterName(c *SqlParameterNameContext)

	// ExitSqlVariableName is called when exiting the sqlVariableName production.
	ExitSqlVariableName(c *SqlVariableNameContext)

	// ExitTransitionVariableName is called when exiting the transitionVariableName production.
	ExitTransitionVariableName(c *TransitionVariableNameContext)

	// ExitSynonym is called when exiting the synonym production.
	ExitSynonym(c *SynonymContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitIntoClause is called when exiting the intoClause production.
	ExitIntoClause(c *IntoClauseContext)

	// ExitCorrelationClause is called when exiting the correlationClause production.
	ExitCorrelationClause(c *CorrelationClauseContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitTableReference is called when exiting the tableReference production.
	ExitTableReference(c *TableReferenceContext)

	// ExitSingleTableReference is called when exiting the singleTableReference production.
	ExitSingleTableReference(c *SingleTableReferenceContext)

	// ExitPeriodSpecification is called when exiting the periodSpecification production.
	ExitPeriodSpecification(c *PeriodSpecificationContext)

	// ExitPeriodClause is called when exiting the periodClause production.
	ExitPeriodClause(c *PeriodClauseContext)

	// ExitNestedTableExpression is called when exiting the nestedTableExpression production.
	ExitNestedTableExpression(c *NestedTableExpressionContext)

	// ExitDataChangeTableReference is called when exiting the dataChangeTableReference production.
	ExitDataChangeTableReference(c *DataChangeTableReferenceContext)

	// ExitTableFunctionReference is called when exiting the tableFunctionReference production.
	ExitTableFunctionReference(c *TableFunctionReferenceContext)

	// ExitTableUdfCardinalityClause is called when exiting the tableUdfCardinalityClause production.
	ExitTableUdfCardinalityClause(c *TableUdfCardinalityClauseContext)

	// ExitTypedCorrelationClause is called when exiting the typedCorrelationClause production.
	ExitTypedCorrelationClause(c *TypedCorrelationClauseContext)

	// ExitTableLocatorReference is called when exiting the tableLocatorReference production.
	ExitTableLocatorReference(c *TableLocatorReferenceContext)

	// ExitXmltableExpression is called when exiting the xmltableExpression production.
	ExitXmltableExpression(c *XmltableExpressionContext)

	// ExitCollectionDerivedTable is called when exiting the collectionDerivedTable production.
	ExitCollectionDerivedTable(c *CollectionDerivedTableContext)

	// ExitJoinCondition is called when exiting the joinCondition production.
	ExitJoinCondition(c *JoinConditionContext)

	// ExitFullJoinExpression is called when exiting the fullJoinExpression production.
	ExitFullJoinExpression(c *FullJoinExpressionContext)

	// ExitCastFunction is called when exiting the castFunction production.
	ExitCastFunction(c *CastFunctionContext)

	// ExitOrdinaryArrayExpression is called when exiting the ordinaryArrayExpression production.
	ExitOrdinaryArrayExpression(c *OrdinaryArrayExpressionContext)

	// ExitAssociativeArrayExpression is called when exiting the associativeArrayExpression production.
	ExitAssociativeArrayExpression(c *AssociativeArrayExpressionContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitGroupingExpression is called when exiting the groupingExpression production.
	ExitGroupingExpression(c *GroupingExpressionContext)

	// ExitGroupingSets is called when exiting the groupingSets production.
	ExitGroupingSets(c *GroupingSetsContext)

	// ExitGroupingSetsGroup is called when exiting the groupingSetsGroup production.
	ExitGroupingSetsGroup(c *GroupingSetsGroupContext)

	// ExitSuperGroups is called when exiting the superGroups production.
	ExitSuperGroups(c *SuperGroupsContext)

	// ExitSelectColumns is called when exiting the selectColumns production.
	ExitSelectColumns(c *SelectColumnsContext)

	// ExitUnpackedRow is called when exiting the unpackedRow production.
	ExitUnpackedRow(c *UnpackedRowContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitSubSelect is called when exiting the subSelect production.
	ExitSubSelect(c *SubSelectContext)

	// ExitSelectIntoStatement is called when exiting the selectIntoStatement production.
	ExitSelectIntoStatement(c *SelectIntoStatementContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitCommonTableExpression is called when exiting the commonTableExpression production.
	ExitCommonTableExpression(c *CommonTableExpressionContext)

	// ExitUpdateClause is called when exiting the updateClause production.
	ExitUpdateClause(c *UpdateClauseContext)

	// ExitReadOnlyClause is called when exiting the readOnlyClause production.
	ExitReadOnlyClause(c *ReadOnlyClauseContext)

	// ExitOptimizeClause is called when exiting the optimizeClause production.
	ExitOptimizeClause(c *OptimizeClauseContext)

	// ExitIsolationClause is called when exiting the isolationClause production.
	ExitIsolationClause(c *IsolationClauseContext)

	// ExitLockClause is called when exiting the lockClause production.
	ExitLockClause(c *LockClauseContext)

	// ExitSkipLockedDataClause is called when exiting the skipLockedDataClause production.
	ExitSkipLockedDataClause(c *SkipLockedDataClauseContext)

	// ExitQuerynoClause is called when exiting the querynoClause production.
	ExitQuerynoClause(c *QuerynoClauseContext)

	// ExitScalarFullSelect is called when exiting the scalarFullSelect production.
	ExitScalarFullSelect(c *ScalarFullSelectContext)

	// ExitFullSelect is called when exiting the fullSelect production.
	ExitFullSelect(c *FullSelectContext)

	// ExitValuesClause is called when exiting the valuesClause production.
	ExitValuesClause(c *ValuesClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitSortKey is called when exiting the sortKey production.
	ExitSortKey(c *SortKeyContext)

	// ExitOffsetClause is called when exiting the offsetClause production.
	ExitOffsetClause(c *OffsetClauseContext)

	// ExitFetchClause is called when exiting the fetchClause production.
	ExitFetchClause(c *FetchClauseContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifier1 is called when exiting the identifier1 production.
	ExitIdentifier1(c *Identifier1Context)

	// ExitSqlidentifier is called when exiting the sqlidentifier production.
	ExitSqlidentifier(c *SqlidentifierContext)

	// ExitSqlKeyword is called when exiting the sqlKeyword production.
	ExitSqlKeyword(c *SqlKeywordContext)
}
