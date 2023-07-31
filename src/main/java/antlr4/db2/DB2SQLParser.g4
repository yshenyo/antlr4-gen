/*
Copyright (C) 2021 - 2023 Craig Schneiderwent.  All rights reserved.
I accept no liability for damages of any kind resulting from the use
of this software.  Use at your own risk.

This software may be modified and distributed under the terms
of the MIT license. See the LICENSE file for details.

Rules for Db2 for z/OS SQL statements that can be embedded in an
application program are included here.  Version 12 documentation
served as original source material.

The ALTER FUNCTION variations (external), (inlined SQL scalar), and
(SQL table) are all variations on each other and are contained in
the alterFunctionStatement rule.

The rule signalStatement is not a full implementation of the syntax
of the SIGNAL statement, but a subset that is possible to embed in
an application program.

The rule trustedContextOptionList does not strictly match its
syntax diagram, for reasons documented with the rule.

The rule executeImmediateStatement only partially implements the
syntax of the EXECUTE IMMEDIATE	statement in order to avoid
implementing a grammar for the entire PL/I language here.  The same
is true of the prepareStatement rule.

This grammar does not include SQL/PL or the following SQL statements.

ALTER TRIGGER (advanced)
CREATE TRIGGER (advanced)

I do not know if the SQL statement 
ALTER PROCEDURE (SQL - external)
will work with this grammar.  It's been deprecated but the syntax
is similar enough to a regular external stored procedure that it
might work.

While SQL/PL is not included in this grammar, its presence is
tolerated to some extent in the CREATE and ALTER of native SQL
stored procedures and functions.

*/


parser grammar DB2SQLParser;

options {tokenVocab=DB2SQLLexer;}

startRule : sqlStatement* | EOF ;

/*
The order of the releaseSavepointStatement and
releaseConnectionStatement rules is significant. In order
for savepoints to be recognized correctly the former
must come before the latter.
*/
sqlStatement
	: EXEC_SQL?
	(
	query
	| allocateCursorStatement
	| alterAuditPolicyStatement
	| alterBufferpoolStatement
	| alterDatabasePartitionStatement
	| alterDatabaseStatement
    | alterEventMonitorStatement
	| alterFunctionStatement
	| alterHistogramTemplateStatement
	| alterIndexStatement
	| alterMaskStatement
	| alterMethodStatement
	| alterPackageStatement
	| alterPermissionStatement
	| alterProcedureStatement
	| alterSequenceStatement
	| alterStogroupStatement
	| alterTableStatement
	| alterTablespaceStatement
	| alterTriggerStatement
	| alterTrustedContextStatement
	| alterViewStatement
	| associateLocatorsStatement
	| beginDeclareSectionStatement
	| callStatement
	| closeStatement
	| commitStatement
	| commentStatement
	| connectStatement
	| createAliasStatement
	| createDatabaseStatement
	| createFunctionStatement
	| createGlobalTemporaryTableStatement
	| createIndexStatement
    | createMaskStatement
	| createPermissionStatement
	| createProcedureStatement
	| createRoleStatement
	| createRoleStatement
	| createSequenceStatement
	| createTableStatement
	| createTablespaceStatement
	| createTriggerStatement
	| createTrustedContextStatement
	| createTypeArrayStatement
	| createTypeDistinctStatement
	| createVariableStatement
	| createViewStatement
	| declareCursorStatement
	| declareGlobalTemporaryTableStatement
	| deleteStatement
	| describeStatement
	| disconnectStatement
	| dropStatement
	| endDeclareSectionStatement
	| executeStatement
	| executeImmediateStatement
	| explainStatement
	| fetchStatement
	| freeLocatorStatement
	| getDiagnosticsStatement
	| grantStatement
	| includeStatement
	| insertStatement
	| lockTableStatement
	| mergeStatement
	| openStatement
	| prepareStatement
	| releaseSavepointStatement
	| releaseConnectionStatement
	| renameStatement
    | reorgTableStatement
    | reorgIndexStatement
    | reorgchkStatement
	| revokeStatement
	| rollbackStatement
	| savepointStatement
	| setConnectionStatement
	| setEncryptionPasswordStatement
	| setPathStatement
	| setSchemaStatement
	| setSessionTimezoneStatement
	| setSpecialRegisterStatement
	| setAssignmentStatement
	| signalStatement
	| transferOwnershipStatement
	| truncateStatement
	| updateStatement
	| valuesIntoStatement
	| wheneverStatement
	)
	(SEMICOLON+ | (END_EXEC DOT?) | EOF)
	;

query
	: (
	subSelect
	| fullSelect
	| selectStatement
	| selectIntoStatement
	)
	;

cursorName
	: identifier
	;

statementName
	: identifier
	;

descriptorName
	: hostVariable
	;

holdability
	: ((WITHOUT HOLD) | (WITH HOLD))
	;

returnability
	: ((WITHOUT RETURN) | (WITH RETURN ((TO CALLER) | (TO CLIENT))?))
	;

allocateCursorStatement
	: (ALLOCATE cursorName CURSOR FOR RESULT SET rsLocatorVariable)
	;

rsLocatorVariable
	: hostVariable
	;

alterAuditPolicyStatement
    : ALTER AUDIT POLICY identifier (
    CATEGORIES (ALL | AUDIT | CHECKING | CONTEXT | EXECUTE (WITHOUT DATA | WITH DATA)? | OBJMAINT | SECMAINT | SYSADMIN | VALIDATE)
    STATUS (BOTH | FAILURE | SUCCESS | NONE)
    (COMMA CATEGORIES (ALL | AUDIT | CHECKING | CONTEXT | EXECUTE (WITHOUT DATA | WITH DATA)? | OBJMAINT | SECMAINT | SYSADMIN | VALIDATE)
    STATUS (BOTH | FAILURE | SUCCESS | NONE))*
    | ERROR TYPE (NORMAL | AUDIT)
    )+
    ;

alterBufferpoolStatement
    : ALTER BUFFERPOOL identifier ((IMMEDIATE | DEFERRED)? (MEMBER numeric) SIZE (numeric | numeric? AUTOMATIC)
    | ADD DATABASE PARTITION GROUP identifier
    | NUMBLOCKPAGES numeric (BLOCKSIZE numeric)?
    | BLOCKSIZE numeric
    )
    ;

alterDatabasePartitionStatement
    : ALTER DATABASE PARTITION GROUP identifier
    (ADD (DBPARTITIONNUM | DBPARTITIONNUMS) partitionsClause partitionOptions?
    |DROP (DBPARTITIONNUM | DBPARTITIONNUMS) partitionsClause
    )
    ;

partitionsClause
    : LPAREN numeric (TO numeric)? RPAREN
    ;

partitionOptions
    : LIKE DBPARTITIONNUM numeric
    | WITHOUT TABLESPACES
    ;

alterDatabaseStatement
	: ALTER DATABASE databaseName
	((ADD | DROP) STORAGE ON NONNUMERICLITERAL (COMMA NONNUMERICLITERAL)*)+
	;

alterEventMonitorStatement
    : ALTER EVENT MONITOR identifier (ADD LOGICAL GROUP identifier (LPAREN targetTableOptions RPAREN)?)+
    ;

targetTableOptions
    : (TABLE tableName
    | IN identifier
    | PCTDEACTIVATE numeric
    )+
    ;

alterFunctionStatement
	: ALTER functionDesignator
	(EXTERNAL NAME identifier
	| NOT? FENCED
	| NOT? SECURED
	| NOT? THREADSAFE
	)+
	;

functionDesignator
	: (
	((FUNCTION functionName (LPAREN functionParameterType (COMMA functionParameterType)* RPAREN)?) 
	| (SPECIFIC FUNCTION specificName))
	)
	;

alterHistogramTemplateStatement
    : ALTER HISTOGRAM TEMPLATE identifier HIGH BIN VALUE numeric
    ;

alterIndexStatement
	:
	ALTER INDEX indexName COMPRESS (NO | YES)
	;

alterMaskStatement
	: (
	ALTER MASK maskName (ENABLE | DISABLE )
	)
	;

alterMethodStatement
    : ALTER methodDesignator EXTERNAL NAME identifier
    ;

methodDesignator
    : METHOD identifier (LPAREN (dataType (COMMA dataType)*)? RPAREN)?
    | SPECIFIC METHOD specificName
    ;

alterPackageStatement
    : ALTER PACKAGE packageName (VERSION? versionID)? (ACCESS PLAN REUSE (YES | NO)
    | OPTIMIZATION PROFILE (NONE | identifier)
    | KEEP DYNAMIC (YES | NO)
    )+
    ;

alterPermissionStatement
	: (
	ALTER PERMISSION permissionName (ENABLE | DISABLE)
	)
	;

alterProcedureStatement
    : ALTER PROCEDURE procedureDesignator
	(EXTERNAL NAME identifier
	| NOT? FENCED
	| NOT? SECURED
	| NOT? THREADSAFE
	| NEW SAVEPOINT LEVEL
	| NO? EXTERNAL ACTION
	| ALTER PARAMETER parameterAlteration
	)+
	;

procedureDesignator
	: (
	((PROCEDURE procedureName (LPAREN functionParameterType (COMMA functionParameterType)* RPAREN)?)
	| (SPECIFIC PROCEDURE specificName))
	)
	;

parameterAlteration
    : identifier SET DATA TYPE dataType
    ;

alterSequenceStatement
	: (
	ALTER SEQUENCE sequenceName alterSequenceOptionList+
	)
	;

alterStogroupStatement
	: (
	ALTER STOGROUP stogroupName
	(ADD NONNUMERICLITERAL (COMMA NONNUMERICLITERAL)*
	| DROP NONNUMERICLITERAL (COMMA NONNUMERICLITERAL)*
	| OVERHEAD INTEGERLITERAL
	| DEVICE READ RATE INTEGERLITERAL
	| DATA TAG (INTEGERLITERAL | NONE)
	| SET AS DEFAULT
	)+
	)
	;

alterTableStatement
	: (
	ALTER TABLE tableName (alterTableOptionList+
	| ADD PARTITION addPartition
	| ATTACH PARTITION attachPartition
	| DETACH PARTITION identifier INTO tableName
	| ADD SECURITY POLICY identifier
	| DROP SECURITY POLICY
	| ADD VERSIONING USE HISTORY TABLE tableName
	| DROP VERSIONING
	)
	)
	;

addPartition
    : identifier? boundarySpec (IN tablespaceName)?
    (INDEX IN tablespaceName (LONG IN tablespaceName)?)?
    ;

boundarySpec
    : startingClause endingClause
    | endingClause
    ;

startingClause
    : STARTING FROM? (LPAREN (literal | MINVALUE | MAXVALUE) (COMMA (literal | MINVALUE | MAXVALUE))* RPAREN
    | literal | MINVALUE | MAXVALUE)
    ;

endingClause
    : ENDING AT? (LPAREN (literal | MINVALUE | MAXVALUE) (COMMA (literal | MINVALUE | MAXVALUE))* RPAREN
    | literal | MINVALUE | MAXVALUE) (INCLUSIVE | EXCLUSIVE)?
    ;

attachPartition
    : identifier? boundarySpec FROM tableName (BUILD MISSING INDEXES | REQUIRE MATCHING INDEXES)?
    ;

/*
2022-11-04 Changed subrule operator for alterPartitionClause from
? (0 or 1 allowed) to * (0 or more allowed) per documentation.
Noticed whilst working on issue 125.
*/
alterTablespaceStatement
	: (
	ALTER TABLESPACE (databaseName DOT)? tablespaceName 
	alterTablespaceOptionList* 
	alterPartitionClause*
	moveTableClause?
	)
	;

alterTriggerStatement
	: (
	ALTER TRIGGER (schemaName DOT) triggerName NOT? SECURED
	)
	;

alterTrustedContextStatement
	: (
	ALTER TRUSTED CONTEXT contextName trustedContextOptionList+
	)
	;

alterViewStatement
	: (
	ALTER VIEW viewName
	(
	ALTER COLUMN? columnName ADD SCOPE identifier
	| (ENABLE | DISABLE) QUERY OPTIMIZATION
	)
	)
	;

associateLocatorsStatement
	: (
	ASSOCIATE (RESULT SET)? (LOCATOR | LOCATORS) 
	LPAREN rsLocatorVariable (COMMA rsLocatorVariable)* RPAREN
	WITH PROCEDURE procedureName
	)
	;

beginDeclareSectionStatement
	: (BEGIN DECLARE SECTION)
	;

callStatement
	: (
	CALL procedureName
	(LPAREN argument (COMMA argument)* RPAREN)?
	)
	;

argument
    : identifier? (expression | DEFAULT | NULL)
    ;

closeStatement
	: (CLOSE cursorName (WITH RELEASE)?)
	;

commentStatement
	: (
	COMMENT ON ((
	(aliasDesignator
	| AUDIT POLICY identifier
	| (COLUMN tableName DOT columnName)
	| CONSTRAINT tableName DOT constraintName
	| DATABASE PARTITION GROUP identifier
	| functionDesignator
	| FUNCTION MAPPING identifier
	| HISTOGRAM TEMPLATE identifier
	| (INDEX indexName)
	| MASK identifier
	| MODULE identifier
	| NICKNAME identifier
	| (PACKAGE packageDesignator)
	| PERMISSION identifier
	| procedureDesignator
	| (ROLE roleName)
	| SCHEMA identifier
	| SECURITY LABEL identifier
	| SECURITY LABEL COMPONENT identifier
	| SECURITY POLICY identifier
	| (SEQUENCE sequenceName)
	| SERVER identifier
	| (TABLE tableName)
	| TABLESPACE identifier
	| THRESHOLD identifier
	| (TRIGGER triggerName )
	| (TRUSTED CONTEXT contextName)
	| (TYPE typeName)
	| TYPE MAPPING identifier
	| (VARIABLE variableName)
	| WORK ACTION SET identifier
	| WORK CLASS SET identifier
	| WORKLOAD identifier
	| WRAPPER identifier
	| XSROBJECT identifier)
	IS NONNUMERICLITERAL)
	| multipleColumnList)
	)
	;

commitStatement
	: (COMMIT WORK?)
	;

connectStatement
	: (
	CONNECT (
	(TO (locationName | hostVariable) lockBlock? authorization?)
	| RESET
	| authorization)?
	)
	;

lockBlock
    : IN SHARE MODE_
    | IN EXCLUSIVE MODE_ (ON SINGLE MEMBER)?
    ;

createAliasStatement
	: (
	CREATE (OR REPLACE)? PUBLIC? ALIAS (sequenceAlias | tableAlias | moduleAlias)
	)
	;

moduleAlias
    : identifier FOR MODULE identifier
    ;

createDatabaseStatement
	: (
	CREATE DATABASE databaseName databaseOptionList*
	)
	;

createFunctionStatement
	: (
	createFunctionStatementExternalScalar
	| createFunctionStatementExternalTable
	| createFunctionStatementSourced
	| createFunctionStatementInlineSqlScalar
	| createFunctionStatementCompiledSqlScalar
	| createFunctionStatementSqlTable
	)
	;

createFunctionStatementExternalScalar
	: (
	CREATE (OR REPLACE)? FUNCTION functionName
	LPAREN (parameterDeclaration1 (COMMA parameterDeclaration1)*)? RPAREN
	createFunctionStatementExternalScalarReturnsPhrase
	createFunctionStatementExternalScalarOptions+
	)
	;

createFunctionStatementExternalScalarReturnsPhrase
	: (
	RETURNS 
		((functionDataType (AS LOCATOR)?)
		| (functionDataType CAST FROM functionDataType (AS LOCATOR)?))
	)
	;

createFunctionStatementExternalTable
	: (
	CREATE (OR REPLACE)? FUNCTION functionName
	LPAREN (parameterDeclaration1 (COMMA parameterDeclaration1)*)? RPAREN
	createFunctionStatementExternalTableReturnsPhrase
	createFunctionStatementExternalTableOptions+
	)
	;

createFunctionStatementExternalTableReturnsPhrase
	:(
	RETURNS 
		((TABLE LPAREN externalTableFunctionColumn (COMMA externalTableFunctionColumn)* RPAREN)
		| (GENERIC TABLE))
	)
	;

externalTableFunctionColumn
	: (
	columnName functionDataType (AS LOCATOR)?
	)
	;

createFunctionStatementSourced
	: (
	CREATE (OR REPLACE)? FUNCTION functionName
	LPAREN (parameterDeclaration1 (COMMA parameterDeclaration1)*)? RPAREN
	createFunctionStatementSourcedReturnsPhrase
	createFunctionStatementSourcedOptions*
	createFunctionStatementSourcedSourcePhrase
	)
	;

createFunctionStatementSourcedReturnsPhrase
	: (RETURNS functionDataType (AS LOCATOR)?)
	;

createFunctionStatementSourcedSourcePhrase
	: (
	SOURCE 
		((functionName LPAREN parameterDeclaration1 (COMMA parameterDeclaration1)* RPAREN)
		| specificNameOption2)
	)
	;

createFunctionStatementInlineSqlScalar
	: (
	CREATE (OR REPLACE)? FUNCTION functionName
	LPAREN ((parameterDeclaration1 (COMMA parameterDeclaration1)*)?) RPAREN
	((WRAPPED obfuscatedStatementText) | inlineSqlScalarFunctionDefinition)
	)
	;

createFunctionStatementCompiledSqlScalar
	: (
	CREATE (OR REPLACE)? FUNCTION functionName
	LPAREN ((parameterDeclaration2 (COMMA parameterDeclaration2)*)?) RPAREN
	((WRAPPED obfuscatedStatementText) | compiledSqlScalarFunctionDefinition)
	)
	;

createFunctionStatementSqlTable
	: (
	CREATE (OR REPLACE)? FUNCTION functionName
	LPAREN ((parameterDeclaration2 (COMMA parameterDeclaration2)*)?) RPAREN
	((WRAPPED obfuscatedStatementText) | sqlTableFunctionDefinition)
	)
	;

createGlobalTemporaryTableStatement
	: (
	CREATE GLOBAL TEMPORARY TABLE tableName
	((LPAREN createGlobalTemporaryTableColumnDefinition (COMMA createGlobalTemporaryTableColumnDefinition)* RPAREN)
	| (LIKE tableName copyOptions?)
	| AS LPAREN fullSelect RPAREN WITH NO DATA copyOptions?)
    (ON COMMIT DELETE ROWS | ON COMMIT PRESERVE ROWS)?
    (NOT LOGGED (ON ROLLBACK DELETE ROWS)? | NOT LOGGED ON ROLLBACK PRESERVE ROWS | LOGGED)?
    (IN identifier)? distributionClause?
	)
	;

distributionClause
    : DISTRIBUTE BY HASH LPAREN columnName (COMMA columnName)* RPAREN
    ;

createIndexStatement
	: (
	CREATE (UNIQUE (WHERE NOT NULL)?)? INDEX indexName ON
		((tableName LPAREN 
		(columnName | keyExpression) (ASC | DESC | RANDOM)?
		(COMMA (columnName | keyExpression) (ASC | DESC | RANDOM)?)*
		(COMMA BUSINESS_TIME (WITH | WITHOUT) OVERLAPS)?
		RPAREN)
		| (auxTableName))
	createIndexOptionList*
	)
	;

createMaskStatement
	: (
	CREATE (OR REPLACE)? MASK maskName ON tableName (AS? correlationName)?
	FOR COLUMN columnName RETURN caseExpression enableDisableOption?
	)
	;

createPermissionStatement
	: (
	CREATE (OR REPLACE)? PERMISSION permissionName ON tableName (AS? correlationName)?
	FOR ROWS WHERE searchCondition ENFORCED FOR ALL ACCESS enableDisableOption?
	)
	;

/*
Note that...

	((WRAPPED obfuscatedStatementText) | sqlRoutineBody)

...must be in that order.  Reversing the order makes WRAPPED 
match sqlRoutineBody because WRAPPED is a sqlKeyword.  This
is incorrect and misleading.
*/

subsetSplStatement
    : callStatement
    | caseStatement
    | continueStatement
    | exitStatement
    | forStatement
    | foreachStatement
    | gotoStatement
    | ifStatement
    | letStatement
    | loopStatement
    | raiseException
    | returnStatement
    | systemStatement
    | traceStatement
    | whileStatement
    ;

caseStatement
    : CASE (elseClause
    | (WHEN (expression | NULL) THEN statementBlock)+ elseClause?
    ) END CASE
    ;

elseClause
    : ELSE statementBlock
    ;

continueStatement
    : CONTINUE (FOR | FOREACH | LOOP | WHILE)? SEMICOLON?
    ;

exitStatement
    : EXIT (FOREACH | (FOR | LOOP | WHILE)? (identifier? WHEN searchCondition)? ) SEMICOLON?
    ;

forStatement
    : FOR identifier (IN LPAREN (range | expression) (COMMA (range | expression))* RPAREN
    | EQ range
    )
    (statementBlock END FOR identifier?
    | LOOP statementBlock END LOOP identifier?
    )
    SEMICOLON?
    ;

foreachStatement
    : FOREACH (WITH HOLD | identifier (WITH HOLD)? FOR | routineCall (INTO identifier (COMMA identifier)*)?)? selectStatement?
    statementBlock END FOREACH SEMICOLON?
    ;

routineCall
    : EXECUTE (PROCEDURE procedureName | FUNCTION functionName)  LPAREN (argument (COMMA argument)*)? RPAREN
    ;

gotoStatement
    : GOTO identifier SEMICOLON?
    ;

ifStatement
    : IF searchCondition THEN ifStatementList? (ELIF searchCondition THEN ifStatementList?)+ (ELSE ifStatementList)? END IF SEMICOLON?
    ;

ifStatementList
    : BEGIN statementBlock END
    | sqlStatement
    | subsetSplStatement
    ;

letStatement
    : LET identifier (COMMA identifier)*
    (functionName LPAREN (argument (COMMA argument)*)? RPAREN
    | expression
    | LPAREN selectStatement (COMMA selectStatement)* RPAREN
    )
    (COMMA (functionName LPAREN (argument (COMMA argument)*)? RPAREN
    | expression
    | LPAREN selectStatement (COMMA selectStatement)* RPAREN
    ))*
    ;

loopStatement
    : (WHILE searchCondition | FOR identifier (IN LPAREN (range | expression) RPAREN | EQ range))?
    LOOP statementBlock END LOOP SEMICOLON?
    ;

range
    : left=expression TO right=expression (STEP expression)?
    ;

raiseException
    : RAISE EXCEPTION expression (COMMA expression (COMMA (identifier | expression))?) SEMICOLON?
    ;

returnStatement
    : RETURN (expression (COMMA expression)* (WITH RESUME)?)? SEMICOLON?
    ;

systemStatement
    : SYSTEM (identifier | expression) SEMICOLON?
    ;

traceStatement
    : TRACE (ON | OFF | PROCEDURE | expression) SEMICOLON?
    ;

whileStatement
    : WHILE searchCondition (statementBlock END WHILE identifier? | LOOP statementBlock END LOOP identifier?) SEMICOLON?
    ;

statementBlock
    : defineStatement* statementBlockClause*
    ;

statementBlockClause
    : executeFunction
    | executeProcedure
    | subsetSplStatement
    | sqlStatement
    | BEGIN statementBlock END
    ;

executeFunction
    : EXECUTE FUNCTION functionName (LPAREN (argument (COMMA argument)*)? RPAREN)? intoClause? (WITH TRIGGER REFERENCES)?
    ;

executeProcedure
    : EXECUTE PROCEDURE procedureName (LPAREN (argument (COMMA argument)*)? RPAREN)? (INTO identifier (COMMA identifier)*)? (WITH TRIGGER REFERENCES)?
    ;

defineStatement
    : DEFINE (
    | identifier (COMMA identifier)* (dataType | REFERENCES (BYTE | TEXT) | LIKE columnName | PROCEDURE | BLOB | CLOB)
    ) SEMICOLON?
    ;

sqlRoutineBody
	: probablySQLPL+
	;

obfuscatedStatementText
	: probablySQLPL+
	;

probablySQLPL
	: (identifier 
	| literal
	| DOT
	| COMMA 
	| COLON
	| SPLAT
	| SEMICOLON
	| LPAREN 
	| RPAREN
	| EQ
	)
	;

createProcedureStatement
	: (
	createProcededureStatementExternal
	| createProcededureStatementSourced
	| createProcededureStatementSQL
	)
	;

createProcededureStatementExternal
    :
	CREATE (OR REPLACE)? PROCEDURE procedureName
	(LPAREN (parameterDeclaration3 (COMMA parameterDeclaration3)*)? RPAREN)?
	(createProcedureOptionList* languageOption5? createProcedureOptionList*)
    ;

createProcededureStatementSourced
    :
	CREATE (OR REPLACE)? PROCEDURE procedureName
	sourceProcedureClause
	(createProcedureOptionList* languageOption5? createProcedureOptionList*)
    ;

sourceProcedureClause
    : SOURCE procedureName
    (LPAREN RPAREN | NUMBER OF PARAMETERS numeric)?
    (UNIQUE ID uniqueID)? FOR SERVER identifier
    ;

createProcededureStatementSQL
    :
	CREATE (OR REPLACE)? PROCEDURE procedureName
	(LPAREN (parameterDeclaration3 (COMMA parameterDeclaration3)*)? RPAREN)?
	(createProcedureOptionList* languageOption1? createProcedureOptionList*)
	statementBlock
    ;

createRoleStatement
	: (
	CREATE ROLE roleName
	)
	;

createSequenceStatement
	: (
	CREATE SEQUENCE sequenceName createSequenceOptionList+
	)
	;

/*
Moved options to createTableOptions rule so they could
be in any order.  Suggested by Martijn Rutte 2023-01-10.
*/
createTableStatement
	: (
	CREATE TABLE tableName
		(
			(LPAREN
				(createTableColumnDefinition
				| periodDefinition
				| uniqueConstraint
				| referentialConstraint
				| checkConstraint)
				(COMMA
				(createTableColumnDefinition
				| periodDefinition
				| uniqueConstraint
				| referentialConstraint
				| checkConstraint))*
			RPAREN)
			| (LIKE tableName copyOptions?)
			| (asResultTable copyOptions?)
			| createTableMaterializedQueryDefinition
		)
	createTableOptions*
	)
	;

createTableOptions
	: (
	createTableInClause
	| partitioningClause
	| organizationClause
	| editprocClause
	| validprocClause
	| auditClause
	| obidClause
	| dataCaptureClause
	| restrictOnDropClause
	| ccsidClause1
	| cardinalityClause
	| loggedOption
	| compressOption
	| appendClause
	| dssizeOption
	| bufferpoolOption
	| memberClause
	| trackmodClause
	| pagenumClause
	| keyLabelOption
	)
	;

createTablespaceStatement
	: (
	CREATE TABLESPACE tablespaceName 
	createTablespaceOptionList* 
	)
	;

/*
The syntax accepts WRAPPED obfuscatedStatementText, but not
in a static SQL context so it is not supported here.
*/
createTriggerStatement
	: (
	CREATE (OR REPLACE)? TRIGGER triggerName triggerDefinition
	)
	;

createTrustedContextStatement
	: (
	CREATE (OR REPLACE)? TRUSTED CONTEXT contextName
	BASED UPON CONNECTION USING SYSTEM AUTHID authorizationName
	(trustedContextDefaultRoleClause
	| trustedContextEnableDisableClause
	| trustedContextDefaultSecurityLabelClause
	| trustedContextAttributesClause
	| trustedContextWithUseForClause)+
	)
	;

createTypeArrayStatement
	: (
	CREATE (OR REPLACE)? TYPE arrayTypeName AS createTypeArrayBuiltinType
	ARRAY OPENSQBRACKET (INTEGERLITERAL | createTypeArrayBuiltinType2) CLOSESQBRACKET
	)
	;

createTypeDistinctStatement
	: (
	CREATE (OR REPLACE)? TYPE distinctTypeName AS sourceDataType (INLINE LENGTH INTEGERLITERAL)?
	)
	;

createVariableStatement
	: (
	CREATE (OR REPLACE)? VARIABLE variableName
	(createVariableBuiltInType | arrayTypeName) 
	((DEFAULT | CONSTANT) (NULL | INTEGERLITERAL | NONNUMERICLITERAL | specialRegister | LPAREN expression RPAREN))?
	)
	;

/*
The...

	(LPAREN columnName (COMMA columnName)* RPAREN)

...has been made optional, as it should be per the documentation.

Noted by Martijn Rutte 2023-01-10.

*/
createViewStatement
	: (
	CREATE (OR REPLACE)? VIEW viewName
	(LPAREN columnName (COMMA columnName)* RPAREN)? AS
	(WITH commonTableExpression (COMMA commonTableExpression)*)?
	fullSelect
	createViewCheckOptionClause? (WITH NO? ROW MOVEMENT)?
	)
	;

declareCursorStatement
	: (
	DECLARE cursorName
    (ASENSITIVE | INSENSITIVE)?
	CURSOR (holdability | returnability)* FOR (selectStatement | statementName)
	)
	;

declareGlobalTemporaryTableStatement
	: (
	DECLARE GLOBAL TEMPORARY TABLE tableName
		((LPAREN declareGlobalTemporaryTableColumnDefinition 
			(COMMA declareGlobalTemporaryTableColumnDefinition)* RPAREN)
		| declareGlobalTemporaryTableLikeClause
		| declareGlobalTemporaryTableAsResultTable)
		(
		onCommitClause
		| loggedWithRollbackClause
		| WITH REPLACE
		| IN identifier
		| distributionClause)*
	)
	;

deleteStatement
	: (searchedDelete | positionedDelete)
	;

describeStatement
	: (
	describeCursorStatement
	| describeInputStatement
	| describeOutputStatement
	| describeProcedureStatement
	| describeTableStatement
	)
	;

describeCursorStatement
	: (
	DESCRIBE CURSOR (cursorName | hostVariable) INTO descriptorName
	)
	;

describeInputStatement
	: (
	DESCRIBE INPUT statementName INTO descriptorName
	)
	;

describeOutputStatement
	: (
	DESCRIBE OUTPUT? statementName INTO descriptorName describeUsingOption?
	)
	;

describeProcedureStatement
	: (
	DESCRIBE PROCEDURE (procedureName | hostVariable) INTO descriptorName
	)
	;

describeTableStatement
	: (
	DESCRIBE TABLE hostVariable INTO descriptorName describeUsingOption?
	)
	;

disconnectStatement
    : DISCONNECT (CURRENT | ALL SQL? | identifier | hostVariable)
    ;

dropStatement
	: (
	DROP
		(aliasDesignation
		| AUDIT POLICY identifier
		| BUFFERPOOL identifier
		| DATABASE PARTITION GROUP identifier
		| EVENT MONITOR identifier
		| functionDesignator RESTRICT?
		| FUNCTION MAPPING identifier
		| HISTOGRAM TEMPLATE identifier
		| INDEX indexName
		| INDEX EXTENSION identifier RESTRICT
		| MASK maskName
		| methodDesignator RESTRICT?
		| MODULE identifier
		| NICKNAME identifier
		| PACKAGE identifier
		| PERMISSION identifier
		| procedureDesignator RESTRICT?
		| ROLE roleName
		| SCHEMA schemaName RESTRICT
		| SECURITY LABEL identifier
		| SECURITY LABEL COMPONENT identifier
		| SECURITY POLICY identifier
		| SEQUENCE sequenceName
		| SERVER identifier
		| STOGROUP identifier
		| TABLE tableName
		| TABLE HIERARCHY identifier
		| (TABLESPACE | TABLESPACES) tablespaceName (COMMA TABLESPACE)*
		| (TRANSFORM | TRANSFORMS) (ALL | identifier) FOR identifier
		| THRESHOLD identifier
		| TRIGGER triggerName
		| TRUSTED CONTEXT identifier
		| TYPE typeName RESTRICT?
		| TYPE MAPPING identifier
		| USAGE LIST identifier
		| USER MAPPING FOR (identifier | USER) SERVER identifier
		| VARIABLE variableName
		| VIEW viewName
		| VIEW HIERARCHY identifier
		| WORK ACTION SET identifier
		| WORK CLASS SET identifier
		| WORKLOAD identifier
		| WRAPPER identifier
		| XSROBJECT identifier
		)
	)
	;

endDeclareSectionStatement
	: (END DECLARE SECTION)
	;

executeStatement
	: (
	EXECUTE statementName 
		((USING (variable | arrayElementSpecification) (COMMA (variable | arrayElementSpecification))*)
		| (USING DESCRIPTOR descriptorName)
		| sourceRowData)
	)
	;

/*
This rule _almost_ covers the syntax.  The NONNUMERICLITERAL 
token corresponds to what the documentation refers to as 
string-expression, which "is any PL/I expression that yields 
a string."  I'm not going to implement a grammar for the 
entire PL/I language here.

*/
executeImmediateStatement
	: (EXECUTE IMMEDIATE expression)
	;

explainStatement
	: (
	EXPLAIN
		(explainPlanClause
		)
	)
	;

/*
The syntax diagram in the documentation lists fetchOrientation
as required.  It is optional here because one of its variants is 
optional.
*/
fetchStatement
	: (
	FETCH FROM? cursorName singleRowFetch
	)
	;

freeLocatorStatement
	: (FREE LOCATOR hostVariable (COMMA hostVariable)*)
	;

getDiagnosticsStatement
	: (
	GET (CURRENT | STACKED)? DIAGNOSTICS
	(statementInformation | conditionInformation | combinedInformation)
	)
	;

grantStatement
	: (
	grantCollectionStatement
	| grantDatabaseStatement
	| grantFunctionOrProcedureStatement	
	| grantPackageStatement
	| grantPlanStatement
	| grantSchemaStatement
	| grantSequenceStatement
	| grantSystemStatement
	| grantTableStatement
	| grantTypeOrJarStatement
	| grantVariableStatement
	| grantUseOfStatement
	)
	;

includeStatement
	: (INCLUDE memberName)
	;

insertStatement
	: (
	INSERT INTO (tableName | LPAREN fullSelect RPAREN) (LPAREN columnName (COMMA columnName)* RPAREN)?
	includeColumns?
	((VALUES (valuesList1 |
		(LPAREN valuesList1 (COMMA valuesList1)* RPAREN)))
	| ((WITH commonTableExpression (COMMA commonTableExpression)*)?
		fullSelect isolationClause? querynoClause?)
	| multipleRowInsert)
	(WITH (RR | RS | CS | UR))?
	)
	;

lockTableStatement
	: (
	LOCK TABLE tableName IN (SHARE | EXCLUSIVE) MODE_
	)
	;

mergeStatement
	: (
	MERGE INTO tableName correlationClause? includeColumns?
	USING ((LPAREN* tableReference RPAREN*) | sourceValues) ON searchCondition
	(WHEN matchingCondition THEN (modificationOperation | signalStatement))+ (ELSE IGNORE)?
	(WITH (RR | RS | CS | UR))?
	)
	;

openStatement
	: (
	OPEN cursorName 
	((USING variable (COMMA variable)*)
	| (USING DESCRIPTOR descriptorName))?
	)
	;

/*
This is incomplete, as the FROM string-expression phrase is
not implemented because it is specific to PL/I, string-expression
can be any valid PL/I string expression, and I'm not going to 
implement the entirety of the PL/I language here.
*/
prepareStatement
	: (
	PREPARE statementName 
	(OUTPUT? INTO descriptorName)?
	(INPUT INTO descriptorName)?
	FROM variable
	)
	;

refreshTableStatement
	: (
	REFRESH TABLE tableName (QUERYNO INTEGERLITERAL)?
	)
	;

releaseConnectionStatement
	: (
	RELEASE (CURRENT | (ALL SQL?) | locationName | hostVariable)
	)
	;

releaseSavepointStatement
	: (RELEASE TO? SAVEPOINT savepointName)
	;

renameStatement
	: (
	RENAME ((TABLE? tableName TO tableName) | (INDEX indexName TO indexName))
	)
	;

reorgTableStatement
    : (
    REORG TABLE tableName tableClause tablePartitioningClause? databasePartitionClause?
    )
    ;

tableClause
    : CLASSIC? classicOptions?
    | INPLACE? (inplaceOptions | STOP | PAUSE)?
    | RECLAIM EXTENTS (ALLOW WRITE ACCESS | ALLOW READ ACCESS | ALLOW NO ACCESS)?
    ;

classicOptions
    : (ALLOW READ ACCESS | ALLOW NO ACCESS) (USE identifier)? (INDEX identifier)? INDEXSCAN?
    (LONGLOBDATA (USE identifier)?)? (KEEPDICTIONARY | RESETDICTIONARY)? (START | RESUME)?
    ;

inplaceOptions
    : (ALLOW WRITE ACCESS | ALLOW READ ACCESS)
    (FULL (INDEX identifier)? (TRUNCATE TABLE | NOTRUNCATE TABLE)? | CLEANUP OVERFLOWS)
    ;

tablePartitioningClause
    : ON DATA PARTITION partitionName=identifier
    ;

databasePartitionClause
    : (DBPARTITIONNUM | DBPARTITIONNUMS) partitionSelectionClause
    | ALL DBPARTITIONNUMS (EXCEPT (DBPARTITIONNUM | DBPARTITIONNUMS) partitionSelectionClause)?
    ;

partitionSelectionClause
    : numeric (TO numeric)? (COMMA numeric (TO numeric)?)*
    ;

reorgIndexStatement
    : (
    REORG (INDEX indexName (FOR TABLE tableName)? | INDEXES ALL FOR TABLE tableName)?
    indexClause tablePartitioningClause? databasePartitionClause?
    )
    ;

indexClause
    : (ALLOW NO ACCESS | ALLOW READ ACCESS | ALLOW WRITE ACCESS)?
    (REBUILD | CLEANUP (ALL | PAGES) (RECLAIM EXTENTS)? | (CLEANUP (ALL | PAGES))? RECLAIM EXTENTS)?
    ;

reorgchkStatement
    : REORGCHK (UPDATE STATISTICS | CURRENT STATISTICS)? (ON TABLE USER | ON (SCHEMA identifier | TABLE (USER | SYSTEM | ALL | identifier)))?
    ;

revokeStatement
	: (
	revokeCollectionStatement
	| revokeDatabaseStatement
	| revokeFunctionOrProcedureStatement	
	| revokePackageStatement
	| revokePlanStatement
	| revokeSchemaStatement
	| revokeSequenceStatement
	| revokeSystemStatement
	| revokeTableStatement
	| revokeTypeOrJarStatement
	| revokeVariableStatement
	| revokeUseOfStatement
	)
	;

rollbackStatement
	: (ROLLBACK WORK? (TO SAVEPOINT savepointName?)?)
	;

savepointStatement
	: (
	SAVEPOINT savepointName UNIQUE?
	((ON ROLLBACK RETAIN CURSORS) | (ON ROLLBACK RETAIN LOCKS))+
	)
	;

setAssignmentStatement
	: (
	SET setAssignmentClause
	)
	;

setConnectionStatement
	: (SET CONNECTION (locationName | hostVariable))
	;

/*
The SET statement for this particular special register does not
conform to the pattern set by all the others.  There's always one.
*/
setEncryptionPasswordStatement
	: (
	SET ENCRYPTION_PASSWORD EQ? (NONNUMERICLITERAL | variable)
	(WITH HINT EQ? (NONNUMERICLITERAL | variable))?
	)
	;

/*
PATH is a token.  CURRENT PATH is a token.  Sometimes the former
is a synonym for the latter.  And sometimes not.
*/
setPathStatement
	: (SET (PATH | CURRENT_PATH) EQ? expression (COMMA? expression)*)
	;

/*
SCHEMA is a token.  CURRENT SCHEMA is a token, as is CURRENT_SCHEMA which
is considered equivalent.  Sometimes the former is a synonym for either of
the latter.  And sometimes not.
*/
setSchemaStatement
	: (SET (SCHEMA | CURRENT_SCHEMA) EQ? expression)
	;

/*
SESSIONTIMEZONE, SESSION TIME ZONE, SESSION TIMEZONE, TIME ZONE,
and TIMEZONE are all synonyms for each other in this context.
Changing the lexer rule for SESSION_TIME_ZONE to match all of
the above broke other things that expect TIME and ZONE to be
separate tokens.  So here we are.
*/
setSessionTimezoneStatement
	: (SET (SESSION_TIME_ZONE | (TIME ZONE) | TIMEZONE) EQ? (variable | NONNUMERICLITERAL))
	;

setSpecialRegisterStatement
	: (SET specialRegister EQ? (expression | NULL) (COMMA? expression)*)
	;

signalStatement
	: (
	SIGNAL SQLSTATE VALUE? NONNUMERICLITERAL signalInformation?
	)
	;

transferOwnershipStatement
	: (
	TRANSFER OWNERSHIP OF ownedObject TO newOwner REVOKE PRIVILEGES
	)
	;

truncateStatement
	: (
	TRUNCATE TABLE? tableName ((DROP | REUSE) STORAGE)?
	((IGNORE | (RESTRICT WHEN)) DELETE TRIGGERS)? (CONTINUE IDENTITY)?
	IMMEDIATE?
	)
	;

updateStatement
	: (searchedUpdate | positionedUpdate)
	;

valuesStatement
	: (
	VALUES (expression | (LPAREN expression (COMMA expression)* RPAREN))
	)
	;

valuesIntoStatement
	: (
	VALUES 
	(expression | NULL | (LPAREN (expression | NULL) (COMMA (expression | NULL))* RPAREN))
	INTO ((valuesIntoTargetVariable (COMMA valuesIntoTargetVariable)*) | arrayElementSpecification)
	)
	;

wheneverStatement
	: (
	WHENEVER ((NOT FOUND) | SQLERROR | SQLWARNING)
	(CONTINUE | ((GOTO | (GO TO)) COLON? hostLabel))
	)
	;

/*
These happen to be the same right now.  Taking advantage of that,
but insulating myself against future changes too.
*/
valuesIntoTargetVariable
	: (setAssignmentTargetVariable)
	;

ownedObject
	:( aliasDesignator
       	| AUDIT POLICY identifier
       	| (COLUMN tableName DOT columnName)
       	| CONSTRAINT tableName DOT constraintName
       	| DATABASE PARTITION GROUP identifier
       	| functionDesignator
       	| FUNCTION MAPPING identifier
       	| HISTOGRAM TEMPLATE identifier
       	| (INDEX indexName)
       	| MASK identifier
       	| MODULE identifier
       	| NICKNAME identifier
       	| (PACKAGE packageDesignator)
       	| PERMISSION identifier
       	| procedureDesignator
       	| (ROLE roleName)
       	| SCHEMA identifier
       	| SECURITY LABEL identifier
       	| SECURITY LABEL COMPONENT identifier
       	| SECURITY POLICY identifier
       	| (SEQUENCE sequenceName)
       	| SERVER identifier
       	| (TABLE tableName)
       	| TABLESPACE identifier
       	| THRESHOLD identifier
       	| (TRIGGER triggerName )
       	| (TRUSTED CONTEXT contextName)
       	| (TYPE typeName)
       	| TYPE MAPPING identifier
       	| (VARIABLE variableName)
       	| WORK ACTION SET identifier
       	| WORK CLASS SET identifier
       	| WORKLOAD identifier
       	| WRAPPER identifier
       	| XSROBJECT identifier
      	)
	;

newOwner
	: (
	(ROLE roleName)
	| (USER authorizationName)
	| (SESSION_USER)
	)
	;

grantCollectionStatement
	: (
	GRANT (CREATE | PACKADM) (ON | IN) COLLECTION
	((collectionID (COMMA collectionID)*) | SPLAT) TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantDatabaseStatement
	: (
	GRANT grantDatabaseAuthority (COMMA grantDatabaseAuthority)*
	ON DATABASE databaseName (COMMA databaseName)* TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantFunctionOrProcedureStatement
	: (
	GRANT EXECUTE ON
		((FUNCTION functionSpecification (COMMA functionSpecification)*)
		| (FUNCTION SPLAT)
		| (SPECIFIC FUNCTION specificName (COMMA specificName)*)
		| (PROCEDURE ((procedureName (COMMA procedureName)*) | SPLAT)))
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantPackageStatement
	: (
	GRANT (ALL | (grantPackageAuthority (COMMA grantPackageAuthority)*))
	ON PACKAGE packageSpecification (COMMA packageSpecification)*
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantPlanStatement
	: (
	GRANT grantPlanAuthority (COMMA grantPlanAuthority)*
	ON PLAN planName (COMMA planName)*
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantSchemaStatement
	: (
	GRANT grantSchemaAuthority (COMMA grantSchemaAuthority)*
	ON SCHEMA (SPLAT | (schemaName (COMMA schemaName)*))
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantSequenceStatement
	: (
	GRANT grantSequenceAuthority (COMMA grantSequenceAuthority)*
	ON SEQUENCE sequenceName (COMMA sequenceName)*
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantSystemStatement
	: (
	GRANT grantSystemAuthority (COMMA grantSystemAuthority)*
	(ON SYSTEM)?
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantTableStatement
	: (
	GRANT grantTableAuthority (COMMA grantTableAuthority)*
	ON TABLE? tableName
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantTypeOrJarStatement
	: (
	GRANT USAGE ON
	(((DATA | DISTINCT)? TYPE typeName (COMMA typeName)*)
	| (JAR jarName (COMMA jarName)*))
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantVariableStatement
	: (
	GRANT grantVariableAuthority (COMMA grantVariableAuthority)*
	ON VARIABLE variableName (COMMA variableName)*
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

grantUseOfStatement
	: (
	GRANT USE OF grantUseOfTarget
	TO
	grantee (COMMA grantee)*
	withGrantOption?
	)
	;

revokeCollectionStatement
	: (
	REVOKE (CREATE | PACKADM) (ON | IN) COLLECTION
	((collectionID (COMMA collectionID)*) | SPLAT) FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	)
	;

revokeDatabaseStatement
	: (
	REVOKE grantDatabaseAuthority (COMMA grantDatabaseAuthority)*
	ON DATABASE databaseName (COMMA databaseName)* FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	)
	;

revokeFunctionOrProcedureStatement
	: (
	REVOKE EXECUTE ON
		((FUNCTION functionSpecification (COMMA functionSpecification)*)
		| (FUNCTION SPLAT)
		| (SPECIFIC FUNCTION specificName (COMMA specificName)*)
		| (PROCEDURE ((procedureName (COMMA procedureName)*) | SPLAT)))
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	RESTRICT?
	)
	;

revokePackageStatement
	: (
	REVOKE (ALL | (grantPackageAuthority (COMMA grantPackageAuthority)*))
	ON PACKAGE packageSpecification (COMMA packageSpecification)*
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	)
	;

revokePlanStatement
	: (
	REVOKE grantPlanAuthority (COMMA grantPlanAuthority)*
	ON PLAN planName (COMMA planName)*
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	)
	;

revokeSchemaStatement
	: (
	REVOKE grantSchemaAuthority (COMMA grantSchemaAuthority)*
	ON SCHEMA (SPLAT | (schemaName (COMMA schemaName)*))
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	)
	;

revokeSequenceStatement
	: (
	REVOKE grantSequenceAuthority (COMMA grantSequenceAuthority)*
	ON SEQUENCE sequenceName (COMMA sequenceName)*
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	RESTRICT?
	)
	;

revokeSystemStatement
	: (
	REVOKE grantSystemAuthority (COMMA grantSystemAuthority)*
	(ON SYSTEM)?
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	)
	;

revokeTableStatement
	: (
	REVOKE grantTableAuthority (COMMA grantTableAuthority)*
	ON TABLE? tableName
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	)
	;

revokeTypeOrJarStatement
	: (
	REVOKE USAGE ON
	(((DATA | DISTINCT)? TYPE typeName (COMMA typeName)*)
	| (JAR jarName (COMMA jarName)*))
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	RESTRICT?
	)
	;

/*
The syntax diagram in the documentation says the RESTRICT option 
comes before the revokeDependentPrivilegesOption in this
particular case.
*/
revokeVariableStatement
	: (
	REVOKE grantVariableAuthority (COMMA grantVariableAuthority)*
	ON VARIABLE variableName (COMMA variableName)*
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	RESTRICT?
	revokeDependentPrivilegesOption?
	)
	;

revokeUseOfStatement
	: (
	REVOKE USE OF grantUseOfTarget
	FROM
	grantee (COMMA grantee)*
	revokeByOption?
	revokeDependentPrivilegesOption?
	)
	;

grantUseOfTarget
	: (
	(BUFFERPOOL bpName (COMMA bpName)*)
	| (ALL BUFFERPOOLS)
	| (STOGROUP stogroupName (COMMA stogroupName)*)
	| (TABLESPACE (databaseName DOT)? tablespaceName (COMMA (databaseName DOT)? tablespaceName)*)
	)
	;

grantVariableAuthority
	: (
	(ALL PRIVILEGES?)
	| READ
	| WRITE
	)
	;

grantTableAuthority
	: (
	(ALL PRIVILEGES?)
	| ALTER
	| DELETE
	| INDEX
	| INSERT
	| (REFERENCES (LPAREN columnName (COMMA columnName)* RPAREN)?)
	| SELECT
	| TRIGGER
	| UNLOAD
	| (UPDATE (LPAREN columnName (COMMA columnName)* RPAREN)?)
	)
	;

grantSystemAuthority
	: (
	ACCESSCTRL
	| ARCHIVE
	| BINDADD
	| BINDAGENT
	| BSDS
	| CREATEALIAS
	| CREATEDBA
	| CREATEDBC
	| CREATESG
	| CREATETMTAB
	| CREATE_SECURE_OBJECT
	| DATAACCESS
	| (DBADM ((WITH | WITHOUT) ACCESSCTRL)? ((WITH | WITHOUT) DATAACCESS)?)
	| DEBUGSESSION
	| DISPLAY
	| EXPLAIN
	| MONITOR1
	| MONITOR2
	| RECOVER
	| SQLADM
	| STOPALL
	| STOSPACE
	| SYSADM
	| SYSCTRL
	| SYSOPR
	| TRACE
	)
	;

grantSequenceAuthority
	: (
	ALTER
	| USAGE
	)
	;

grantSchemaAuthority
	: (
	ALTERIN
	| CREATEIN
	| DROPIN
	)
	;

grantPlanAuthority
	: (
	BIND
	| EXECUTE
	)
	;

grantPackageAuthority
	: (
	BIND
	| COPY
	| EXECUTE
	| RUN
	)
	;

packageSpecification
	: (collectionID DOT (packageName | SPLAT))
	;

functionSpecification
	: (
	functionName (LPAREN functionParameterType (COMMA functionParameterType)* RPAREN)?
	)
	;

grantee
	: (authorizationName | (ROLE roleName) | PUBLIC)
	;

withGrantOption
	: (WITH GRANT OPTION)
	;

revokeByOption
	: (
	BY 
		(ALL 
		| ((authorizationName | (ROLE roleName)) 
			(COMMA (authorizationName | (ROLE roleName)))*))
	)
	;

revokeDependentPrivilegesOption
	: (NOT? INCLUDING DEPENDENT PRIVILEGES)
	;

grantDatabaseAuthority
	: (
	DBADM
	| DBCTRL
	| DBMAINT
	| CREATETAB
	| CREATETS
	| DISPLAYDB
	| DROP
	| IMAGCOPY
	| LOAD
	| RECOVERDB
	| REORG
	| REPAIR
	| STARTDB
	| STATS
	| STOPDB
	)
	;

statementInformation
	: (
	(statementInformationVariableEquate (COMMA statementInformationVariableEquate)*)
	| (variable EQ DB2_GET_DIAGNOSTICS_DIAGNOSTICS)
	| (variable EQ DB2_SQL_NESTING_LEVEL)
	)
	;

statementInformationVariableEquate
	: (variable EQ statementInformationItemName)
	;

statementInformationItemName
	: (
	DB2_LAST_ROW
	| DB2_NUMBER_PARAMETER_MARKERS
	| DB2_NUMBER_RESULT_SETS
	| DB2_NUMBER_ROWS
	| DB2_RETURN_STATUS
	| DB2_SQL_ATTR_CURSOR_HOLD
	| DB2_SQL_ATTR_CURSOR_ROWSET
	| DB2_SQL_ATTR_CURSOR_SCROLLABLE
	| DB2_SQL_ATTR_CURSOR_SENSITIVITY
	| DB2_SQL_ATTR_CURSOR_TYPE
	| MORE_
	| NUMBER
	| ROW_COUNT
	)
	;

conditionInformation
	: (
	CONDITION (variable | INTEGERLITERAL) 
	conditionInformationVariableEquate (COMMA conditionInformationVariableEquate)*
	)
	;

conditionInformationVariableEquate
	: (variable EQ (conditionInformationItemName | connectionInformationItemName))
	;

conditionInformationItemName
	: (
	CATALOG_NAME
	| CONDITION_NUMBER
	| CURSOR_NAME
	| DB2_ERROR_CODE1
	| DB2_ERROR_CODE2
	| DB2_ERROR_CODE3
	| DB2_ERROR_CODE4
	| DB2_INTERNAL_ERROR_POINTER
	| DB2_LINE_NUMBER
	| DB2_MESSAGE_ID
	| DB2_MODULE_DETECTING_ERROR
	| DB2_ORDINAL_TOKEN_n
	| DB2_REASON_CODE
	| DB2_RETURNED_SQLCODE
	| DB2_ROW_NUMBER
	| DB2_SQLERRD_SET
	| DB2_SQLERRD1
	| DB2_SQLERRD2
	| DB2_SQLERRD3
	| DB2_SQLERRD4
	| DB2_SQLERRD5
	| DB2_SQLERRD6
	| DB2_TOKEN_COUNT
	| MESSAGE_TEXT
	| RETURNED_SQLSTATE
	| SERVER_NAME
	)
	;

connectionInformationItemName
	: (
	DB2_AUTHENTICATION_TYPE
	| DB2_AUTHORIZATION_ID
	| DB2_CONNECTION_STATE
	| DB2_CONNECTION_STATUS
	| DB2_ENCRYPTION_TYPE
	| DB2_SERVER_CLASS_NAME
	| DB2_PRODUCT_ID
	)
	;

combinedInformation
	: (
	variable EQ ALL
	combinedInformationOption (COMMA combinedInformationOption)*
	)
	;

combinedInformationOption
	: (
	STATEMENT
	| ((CONDITION | CONNECTION) (variable | INTEGERLITERAL)?)
	)
	;

singleRowFetch
	: (
	(INTO ((fetchTargetVariable (COMMA fetchTargetVariable)*) | arrayElementSpecification))
	| ((INTO | USING) DESCRIPTOR descriptorName)
	)
	;
/*
The target variable in this clause could be any of {global-variable-name,
host-variable-name, SQL-parameter-name, SQL-variable-name, 
transition-variable-name} all of which conform to the variableName rule
save for host-variable-name; thus we confine the rule to just those two.
*/

fetchTargetVariable
	: (variable | hostVariable)
	;

/*
The syntax diagram in the documentation shows that both of these
clauses are optional.  In ANTLR terms, this means an error of the
form "rule can match the empty string."  So my interpretation is
to have a rule that allows one, the other, or both, and then the
entire rule is optional in the fetchStatement rule. 
*/

explainPlanClause
	: (
	(PLAN SECTION? | ALL) ((FOR | WITH)? SNAPSHOT)? (WITH REOPT ONCE)? (SET QUERYNO EQ INTEGERLITERAL)? (SET QUERYTAG EQ NONNUMERICLITERAL)? FOR
	(query | insertStatement | mergeStatement | searchedDelete | searchedUpdate)
	)
	;

/*
For purposes of this grammar there is no difference between a
host-variable and a host-variable-array.  The first option in
the sourceRowData rule allows either, but only for former is 
coded as its rule will also match the latter.
*/
sourceRowData
	: (
	((USING hostVariable (COMMA hostVariable)*)
	| (USING DESCRIPTOR descriptorName))
	(FOR (INTEGERLITERAL | hostVariable) ROWS)?
	)
	;

aliasDesignation
	: (PUBLIC? ALIAS aliasName (FOR (TABLE | SEQUENCE))?)
	;

packageDesignator
	: (collectionID DOT packageName (VERSION? versionID)?)
	;

describeUsingOption
	: (USING (NAMES | LABELS | ANY | BOTH))
	;

declareGlobalTemporaryTableLikeClause
	: (LIKE tableName copyOptions?)
	;

onCommitClause
	: (ON COMMIT ((DELETE ROWS) | (PRESERVE ROWS) | (DROP TABLE)))
	;

loggedWithRollbackClause
	: (
	LOGGED
	| (NOT LOGGED (ON ROLLBACK (DELETE | PRESERVE) ROWS)?)
	)
	;

createViewCheckOptionClause
	: (WITH (CASCADED | LOCAL)? CHECK OPTION)
	;

trustedContextDefaultRoleClause
	: (
	(NO DEFAULT ROLE)
	| (DEFAULT ROLE roleName ((WITHOUT ROLE AS OBJECT OWNER) | (WITH ROLE AS OBJECT OWNER AND QUALIFIER))?)
	)
	;

trustedContextEnableDisableClause
	: (DISABLE | ENABLE)
	;

trustedContextDefaultSecurityLabelClause
	: ((NO DEFAULT SECURITY LABEL) | (DEFAULT SECURITY LABEL seclabelName))
	;

trustedContextAttributesClause
	: (
	ATTRIBUTES 
	LPAREN 
	((trustedContextAttribute1 (COMMA trustedContextAttribute1)*)
	| (trustedContextAttribute2 (COMMA trustedContextAttribute2)*))
	RPAREN
	)
	;

trustedContextWithUseForClause
	: (WITH USE FOR trustedContextUseFor (COMMA trustedContextUseFor)*)
	;

trustedContextAttribute1
	: (
	(ADDRESS addressValue)
	| (ENCRYPTION encryptionValue)
	| (SERVAUTH servauthValue)
	)
	;

trustedContextAttribute2
	: (
	(JOBNAME jobnameValue)
	)
	;

trustedContextUseFor
	: (
	(authorizationName userOptions*)
	| (EXTERNAL SECURITY PROFILE profileName userOptions*)
	| (PUBLIC (WITH | WITHOUT) AUTHENTICATION)
	)
	;

userOptions
	: (
	(ROLE roleName)
	| (SECURITY LABEL seclabelName)
	| ((WITH | WITHOUT) AUTHENTICATION)
	)
	;

triggerDefinition
	: (
	triggerActivationTime triggerEvent ON tableName
	(REFERENCING
		((OLD | NEW | OLD_TABLE | NEW_TABLE | (OLD TABLE) | (NEW TABLE)) AS? correlationName)+)?
	triggerGranularity triggerDefinitionOption? triggeredAction
	)
	;

triggerActivationTime
	: (
	((NO CASCADE)? BEFORE)
	| AFTER
	| (INSTEAD OF)
	)
	;

triggerEvent
	: (
	INSERT
	| DELETE
	| (UPDATE (OF columnName (COMMA columnName)*)?)
	)
	;

triggerGranularity
	: (
	(FOR EACH STATEMENT)
	| (FOR EACH ROW)
	)
	;

triggeredAction
	: (
	(WHEN LPAREN searchCondition RPAREN)? sqlTriggerBody
	)
	;

sqlTriggerBody
	: (
	triggeredSqlStatement
	| (BEGIN ATOMIC (triggeredSqlStatement SEMICOLON)+ END)
	)
	;

triggeredSqlStatement
	: (
	callStatement
	| searchedDelete
	| ((commonTableExpression)? fullSelect)
	| insertStatement
	| mergeStatement
	| refreshTableStatement
	| setAssignmentStatement
	| signalStatement
	| truncateStatement
	| searchedUpdate
	| valuesStatement
	)
	;

triggerDefinitionOption
	: (
	(NOT SECURED)
	| SECURED
	)
	;

createTableInClause
	: (
	(customVolatileClause? IN (databaseName DOT)? tablespaceName customVolatileClause?)
	| (customVolatileClause? IN DATABASE databaseName customVolatileClause?)
	| (customVolatileClause? IN ACCELERATOR acceleratorName customVolatileClause?)
	)
	;

/*
Optional CARDINALITY added 2023-01-10.
*/
customVolatileClause
	: (NOT? VOLATILE CARDINALITY?)
	;

createTableColumnDefinition
	: (
	columnName dataType?
	(NOT NULL)?
	generatedClause?
	createTableColumnConstraint?
	defaultClause?
	fieldprocClause?
	asSecurityLabelClause?
	implicitlyHiddenClause?
	inlineLengthClause?
	)
	;

editprocClause
	: (EDITPROC programName ((WITH | WITHOUT) ROW ATTRIBUTES)?)
	;

/*
NULL is only valid for ALTER TABLE.
*/
validprocClause
	: (VALIDPROC (programName | NULL))
	;

auditClause
	: (AUDIT (NONE | CHANGES | ALL))
	;

obidClause
	: (OBID INTEGERLITERAL)
	;

dataCaptureClause
	: (DATA CAPTURE (NONE | CHANGES (INCLUDE LONGVAR COLUMNS)?))
	;

restrictOnDropClause
	: (WITH RESTRICT ON DROP)
	;

ccsidClause1
	: (CCSID (ASCII | EBCDIC | UNICODE))
	;

ccsidClause2
	: (CCSID INTEGERLITERAL)
	;

cardinalityClause
	: (NOT? VOLATILE CARDINALITY?)
	;

appendClause
	: (APPEND (ON | OFF))
	;

memberClause
	: (MEMBER CLUSTER)
	;

trackmodClause
	: (TRACKMOD (YES | NO))
	;

pagenumClause
	: (PAGENUM (RELATIVE | ABSOLUTE))
	;

fieldprocClause
	: (FIELDPROC programName LPAREN literal (COMMA literal)* RPAREN)
	;

asSecurityLabelClause
	: (AS SECURITY LABEL)
	;

implicitlyHiddenClause
	: (IMPLICITLY HIDDEN_)
	;

inlineLengthClause
	: (INLINE LENGTH INTEGERLITERAL)
	;

copyOptions
	: (
		(
		copyOptionIdentity
		| copyOptionRowChangeTimestamp
		| copyOptionColumnDefaults
		| copyOptionXmlTypeModifiers
		)+
	)
	;

copyOptionIdentity
	: ((EXCLUDING | INCLUDING) IDENTITY (COLUMN ATTRIBUTES)?)
	;

copyOptionRowChangeTimestamp
	: ((EXCLUDING | INCLUDING) ROW CHANGE TIMESTAMP (COLUMN ATTRIBUTES)?)
	;

copyOptionColumnDefaults
	: (
		((EXCLUDING | INCLUDING) COLUMN? DEFAULTS)
		| (USING TYPE DEFAULTS)
	)
	;

copyOptionXmlTypeModifiers
	: (EXCLUDING XML TYPE MODIFIERS)
	;

asResultTable
	: (
	LPAREN (columnName (COMMA columnName)*)? RPAREN AS
	LPAREN fullSelect RPAREN
	WITH NO DATA
	)
	;

declareGlobalTemporaryTableAsResultTable
	: (
	AS LPAREN fullSelect RPAREN
	WITH NO DATA copyOptions?
	)
	;

createTableMaterializedQueryDefinition
	: (
	(LPAREN columnName (COMMA columnName)* RPAREN)? 
	AS materializedQueryDefinition
	)
	;

createTableColumnConstraint
	: (
	(CONSTRAINT constraintName)?
		(
			(PRIMARY KEY)
			| UNIQUE
			| referencesClause
			| (CHECK LPAREN checkCondition RPAREN)
		)
	)
	;

/*
Deprecated as of Db2 12.
*/
organizationClause
	: (
	ORGANIZE BY HASH UNIQUE
	LPAREN columnName (COMMA columnName)* RPAREN
	(HASH SPACE INTEGERLITERAL? sqlidentifier)?
	)
	;

createGlobalTemporaryTableColumnDefinition
	: (
	columnName dataType (NOT NULL)?
	)
	;

declareGlobalTemporaryTableColumnDefinition
	: (
	columnName dataType 
	(defaultClause1
	| generatedClause2
	| (NOT NULL))* 
	)
	;

parameterDeclaration1
	: (
	parameterName? ((functionDataType defaultProcedureClause?) | (TABLE LIKE tableName defaultProcedureClause?))
	)
	;

parameterDeclaration2
	: (
	parameterName functionDataType
	)
	;

parameterDeclaration3
	: (
	(IN | OUT | INOUT)? parameterName? procedureDataType defaultProcedureClause?
	)
	;

defaultProcedureClause
    : DEFAULT (NULL | literal | identifier | LPAREN expression RPAREN)
    ;

createFunctionStatementExternalScalarOptions
	: (
	externalNameOption1
	| languageOption3
	| parameterStyleOption2
	| deterministicOption
	| fencedOption
	| nullInputOption1
	| sqlDataOption3
	| externalActionOption
	| packagePathOption
	| scratchpadOption
	| finalCallOption
	| parallelOption2
	| dbinfoOption
	| cardinalityOption
	| collectionIdOption
	| wlmEnvironmentOption1
	| asuTimeOption
	| stayResidentOption
	| programTypeOption
	| securityOption
	| stopAfterFailureOption
	| runOptionsOption
	| specialRegistersOption
	| dispatchOption
	| securedOption
	| specificNameOption1
	| parameterOption1
	)
	;

externalNameOption1
	: (EXTERNAL (NAME (externalProgramName | identifier))?)
	;

dynamicResultSetOption
	: (DYNAMIC? RESULT (SET |SETS) INTEGERLITERAL)
	;

languageOption1
	: (LANGUAGE SQL)
	;

languageOption2
	: (LANGUAGE (ASSEMBLE | C_ | COBOL | PLI))
	;

languageOption3
	: (LANGUAGE (ASSEMBLE | C_ | COBOL | JAVA | PLI))
	;

languageOption5
	: (LANGUAGE (ASSEMBLE | C_ | COBOL | JAVA | PLI | REXX))
	;

parameterStyleOption1
	: (PARAMETER STYLE SQL)
	;

parameterStyleOption2
	: (PARAMETER STYLE (SQL | JAVA))
	;

parameterStyleOption3
	: (PARAMETER STYLE (SQL | DB2SQL | (STANDARD CALL) | GENERAL | (SIMPLE CALL) | ((GENERAL | (SIMPLE CALL)) WITH NULLS) | JAVA))
	;

deterministicOption
	: ((NOT? DETERMINISTIC) | (NOT? VARIANT))
	;

fencedOption
	: (FENCED)
	;

nullInputOption1
	: ((RETURNS NULL ON NULL INPUT) | (CALLED ON NULL INPUT) | (NULL CALL))
	;

nullInputOption2
	: ((CALLED ON NULL INPUT) | (NULL CALL))
	;

debugOption
	: ((DISALLOW | ALLOW | DISABLE) DEBUG MODE_)
	;

sqlDataOption1
	: ((READS SQL DATA) | (CONTAINS SQL))
	;

sqlDataOption2
	: ((READS SQL DATA) | (CONTAINS SQL) | (NO SQL))
	;

sqlDataOption3
	: ((MODIFIES SQL DATA) | (READS SQL DATA) | (CONTAINS SQL) | (NO SQL))
	;

sqlDataOption4
	: ((MODIFIES SQL DATA) | (READS SQL DATA) | (CONTAINS SQL))
	;

externalActionOption
	: (NO? EXTERNAL ACTION)
	;

packagePathOption
	: ((PACKAGE PATH packagePath) | (NO PACKAGE PATH))
	;

scratchpadOption
	: ((NO SCRATCHPAD) | (SCRATCHPAD INTEGERLITERAL))
	;

finalCallOption
	: (NO? FINAL CALL)
	;

parallelOption1
	: (DISALLOW PARALLEL)
	;

parallelOption2
	: ((ALLOW | DISALLOW) PARALLEL)
	;

dbinfoOption
	: (NO? DBINFO)
	;

cardinalityOption
	: (CARDINALITY INTEGERLITERAL)
	;

collectionIdOption
	: ((NO COLLID) | (COLLID collectionID))
	;

wlmEnvironmentOption1
	: (WLM ENVIRONMENT (identifier | (LPAREN identifier RPAREN)))
	;

wlmEnvironmentOption2
	: (WLM ENVIRONMENT (identifier | (LPAREN identifier COMMA SPLAT RPAREN)))
	;

wlmEnvironmentOption3
	: (WLM ENVIRONMENT FOR DEBUG MODE_ identifier)
	;

asuTimeOption
	: (ASUTIME ((NO LIMIT) | (LIMIT INTEGERLITERAL)))
	;

stayResidentOption
	: (STAY RESIDENT (NO | YES))
	;

programTypeOption
	: (PROGRAM TYPE (SUB | MAIN))
	;

securityOption
	: (SECURITY (DB2 | USER | DEFINER))
	;

stopAfterFailureOption
	: ((STOP AFTER SYSTEM DEFAULT FAILURES) | (STOP AFTER INTEGERLITERAL FAILURES) | (CONTINUE AFTER FAILURE))
	;

runOptionsOption
	: (RUN OPTIONS runTimeOptions)
	;

commitOnReturnOption
	: (COMMIT ON RETURN (YES | NO))
	;

specialRegistersOption
	: ((INHERIT | DEFAULT) SPECIAL REGISTERS)
	;

dispatchOption
	: (STATIC DISPATCH)
	;

securedOption
	: (NOT? SECURED)
	;

specificNameOption1
	: (SPECIFIC specificName?)
	;

specificNameOption2
	: (SPECIFIC specificName)
	;

parameterOption1
	: (PARAMETER 
		(ccsidClause1
		| (VARCHAR (NULTERM | STRUCTURE)))+)
	;

parameterOption2
	: (PARAMETER ccsidClause1)
	;

createFunctionStatementExternalTableOptions
	: (
	externalNameOption1
	| languageOption2
	| parameterStyleOption1
	| deterministicOption
	| fencedOption
	| nullInputOption1
	| sqlDataOption2
	| externalActionOption
	| packagePathOption
	| scratchpadOption
	| finalCallOption
	| parallelOption1
	| dbinfoOption
	| cardinalityOption
	| collectionIdOption
	| wlmEnvironmentOption1
	| asuTimeOption
	| stayResidentOption
	| programTypeOption
	| securityOption
	| stopAfterFailureOption
	| runOptionsOption
	| specialRegistersOption
	| dispatchOption
	| securedOption
	| specificNameOption1
	| parameterOption1
	)
	;

createFunctionStatementSourcedOptions
	: (
	specificNameOption2
	| parameterOption2
	)
	;

inlineSqlScalarFunctionDefinition
	: (
	RETURNS functionDataType languageOption1?
	createFunctionStatementInlineSqlScalarOptions+
	statementBlock
	)
	;

createFunctionStatementInlineSqlScalarOptions
	: (
	deterministicOption
	| nullInputOption1
	| sqlDataOption1
	| externalActionOption
	| dispatchOption
	| securedOption
	| specificNameOption1
	| parameterOption2
	)
	;

compiledSqlScalarFunctionDefinition
	: (
	RETURNS functionDataType versionOption?
	createFunctionStatementCompiledSqlScalarOptions+
	statementBlock
	)
	;

createFunctionStatementCompiledSqlScalarOptions
	: (
	languageOption1
	| specificNameOption1
	| deterministicOption
	| externalActionOption
	| sqlDataOption4
	| nullInputOption1
	| dispatchOption
	| parallelOption2
	| debugOption
	| parameterOption2
	| parameterStyleOption1
	| schemaQualifierOption
	| packageOwnerOption
	| asuTimeOption
	| specialRegistersOption
	| wlmEnvironmentOption3
	| currentDataOption
	| degreeOption
	| concurrentAccessOption
	| dynamicRulesOption
	| applicationEncodingOption
	| explainOption
	| immediateWriteOption
	| isolationLevelOption
	| opthintOption
	| queryAccelerationOption
	| getAccelArchiveOption
	| accelerationOption
	| acceleratorOption
	| sqlPathOption
	| reoptOption
	| validateOption
	| roundingOption
	| dateFormatOption
	| decimalOption
	| forUpdateOption
	| timeFormatOption
	| securedOption
	| businessTimeSensitiveOption
	| systemTimeSensitiveOption
	| archiveSensitiveOption
	| applcompatOption
	| concentrateStatementsOption
	)
	;

sqlTableFunctionDefinition
	: (
	RETURNS TABLE LPAREN functionDataType (COMMA functionDataType)* RPAREN
	createFunctionStatementSqlTableOptions+
	statementBlock
	)
	;

createFunctionStatementSqlTableOptions
	: (
	languageOption1
	| specificNameOption1
	| deterministicOption
	| externalActionOption
	| sqlDataOption1
	| nullInputOption1
	| specialRegistersOption
	| dispatchOption
	| cardinalityOption
	| parameterOption2
	| securedOption
	)
	;

sequenceAlias
	: (
	aliasName FOR SEQUENCE sequenceName
	)
	;

tableAlias
	: (
	aliasName FOR TABLE? tableName
	)
	;

authorization
	: (USER hostVariable USING hostVariable)
	;

/*
Alternate syntax allows omission of FROM.  (WHERE searchCondition) should
be optional, i.e. DELETE T1; is a valid statement.
Noted by Martijn Rutte 2023-01-09.
*/
searchedDelete
	: (
	DELETE FROM? (tableName periodClause? | ONLY LPAREN tableName RPAREN | LPAREN fullSelect RPAREN) AS?
	correlationClause? includeColumns? (assignmentClause)? (WHERE searchCondition)?
	(WITH (RR | RS | CS | UR))?
	)
	;

/*
Alternate syntax allows omission of FROM.  Noted by Martijn Rutte 2023-01-09.
*/
positionedDelete
	: (
	DELETE FROM? (tableName | ONLY LPAREN tableName RPAREN)
	AS? correlationClause? WHERE CURRENT OF cursorName
	)
	;

searchedUpdate
	: (
	UPDATE (tableName periodClause? | ONLY LPAREN tableName RPAREN | LPAREN fullSelect RPAREN)
	AS? correlationClause? includeColumns?
	SET assignmentClause (WHERE searchCondition)?
	(WITH (RR | RS | CS | UR))?
	)
	;

positionedUpdate
	: (
	UPDATE (tableName | ONLY LPAREN tableName RPAREN)
	AS? correlationClause?
	SET assignmentClause
	WHERE CURRENT OF cursorName
	)
	;

sourceValues
	: (
	LPAREN VALUES 
	(valuesSingleRow | valuesMultipleRow) 
	RPAREN 
	AS? correlationName 
	LPAREN columnName (COMMA columnName)* RPAREN
	)
	;

valuesSingleRow
	: (
	valuesList3 | (LPAREN valuesList3 (COMMA valuesList3)* RPAREN)
	)
	;

valuesMultipleRow
	: (
	valuesList4 | (LPAREN valuesList4 (COMMA valuesList4)* RPAREN)
	FOR (hostVariable | INTEGERLITERAL) ROWS
	)
	;

matchingCondition
	: (
	NOT? MATCHED (AND searchCondition)?
	)
	;

modificationOperation
	: (updateOperation | deleteOperation | insertOperation)
	;

assignmentClause
	: (
	(columnName EQ valuesList1 (COMMA columnName EQ valuesList1)*)
	| (LPAREN columnName (COMMA columnName)* 
		RPAREN EQ 
		LPAREN ((valuesList1 (COMMA valuesList1)*) | fullSelect)
		RPAREN)
	)
	;

setAssignmentClause
	: (
	(arrayElementSpecification EQ (expression | NULL))
	| (setAssignmentTargetVariable EQ valuesList1 (COMMA setAssignmentTargetVariable EQ valuesList1)*)
	| (LPAREN setAssignmentTargetVariable (COMMA setAssignmentTargetVariable)* 
		RPAREN EQ 
		LPAREN 
			(((valuesList1 (COMMA valuesList1)*) | fullSelect)
			| subSelect
			| (VALUES valuesList1)
			| (VALUES LPAREN valuesList1 (COMMA valuesList1)* RPAREN))
		RPAREN)
	)
	;

setAssignmentTargetVariable
	: (
	globalVariableName
	| hostVariable
	| sqlParameterName
	| sqlVariableName
	| transitionVariableName
	)
	;

updateOperation
	: (
	UPDATE SET assignmentClause (COMMA assignmentClause)*
	)
	;

deleteOperation
	: (DELETE)
	;

insertOperation
	: (
	INSERT LPAREN columnName (COMMA columnName)* RPAREN
	VALUES (valuesList1 |
		(LPAREN valuesList1 (COMMA valuesList1)* RPAREN))
	)
	;

signalInformation
	: (
	(SET MESSAGE_TEXT EQ expression (operator expression)*)
	| (LPAREN NONNUMERICLITERAL RPAREN)
	)
	;

/*
Changed...

	(expression (operator expression)*)

...to...

	(expression ((operator expression) | INTEGERLITERAL)*)
	
...on 2023-01-03 to deal with hitherto unseen syntax...

	INSERT INTO WEATHER VALUES
		(NEWCW.CITY, 
		1.8*NEWCW.TEMPC+32)

...where +32 was lexed as a single token, breaking the
parser rule valuesList1.  I don't have access to a DB2
instance to test if this is a problem with the test case
or not, so here we are.  If you're writing an interpreter
you'll have to deal with this modified rule's handling
of an optional operator.
*/
valuesList1
	: (
	(expression ((operator expression) | INTEGERLITERAL)*)
	| DEFAULT 
	| NULL
	)
	;

valuesList2
	: (expression | hostVariable | DEFAULT | NULL)
	;

valuesList3
	: (expression | NULL)
	;

valuesList4
	: (expression | hostVariable | NULL)
	;

includeColumns
	: (INCLUDE LPAREN columnName dataType (COMMA columnName dataType)* RPAREN)
	;

multipleRowInsert
	: (
	VALUES (valuesList2 | (LPAREN valuesList2 (COMMA valuesList2)* RPAREN))
	(FOR (hostVariable | INTEGERLITERAL) ROWS)?
	(ATOMIC | notAtomicPhrase)
	)
	;


/*
2022-11-04 Issue 125 Changed usingSpecification1 to usingBlock for consistency.
*/
bufferpoolOption
	: (BUFFERPOOL bpName)
	;

closeOption
	: (CLOSE (YES | NO))
	;

copyOption
	: (COPY (YES | NO))
	;

dssizeOption
	: (DSSIZE INTEGERLITERAL? sqlidentifier)
	;

/*
The optional INTEGERLITERAL has been added because it turns
out that both...

PIECESIZE 4K

...and...

PIECESIZE 4 K

...are syntactically correct.  Documentation correction
request submitted to IBM on 05-Aug-2022.

IBM responded to the documentation correction request on
17-Aug-2022, noting that in _any_ place the syntax
<integer>[K|M|G] was allowed, whitespace was allowed
between the <integer> and the [K|M|G].  IBM indicated they
would be updating their documentation to reflect this.

All such changes tagged with #KMG
*/
piecesizeOption
	: (PIECESIZE INTEGERLITERAL? sqlidentifier)
	;

clusterOption
	: (NOT? CLUSTER)
	;

paddedOption
	: (NOT? PADDED)
	;

compressOption
	: (COMPRESS ((YES (FIXEDLENGTH | HUFFMAN)?) | NO))
	;

defineOption
	: (DEFINE (YES | NO))
	;

locksizeOption
	: (LOCKSIZE (ANY | TABLESPACE | TABLE | PAGE | ROW | LOB))
	;

lockmaxOption
	: (LOCKMAX (SYSTEM | INTEGERLITERAL))
	;

enableDisableOption
	: (ENABLE | DISABLE)
	;

/*
Although the latter two options are "supported as alternatives,
they are not the preferred syntax."
*/
loggedOption
	: ((NOT? LOGGED) | (LOG NO) | (LOG YES))
	;

notAtomicPhrase
	: (NOT ATOMIC CONTINUE ON SQLEXCEPTION)
	;

/*
2022-11-04 Issue 125 Changed usingSpecification1 to usingBlock for consistency.
*/

usingSpecification1
	: (
	(USING ((VCAT catalogName) | (STOGROUP stogroupName)))
	| (PRIQTY INTEGERLITERAL)
	| (SECQTY INTEGERLITERAL)
	| (ERASE (YES | NO))
	)
	;

freeSpecification
	: (
	(FREEPAGE INTEGERLITERAL)
	| (PCTFREE INTEGERLITERAL)
	)
	;

gbpcacheSpecification
	: (
	GBPCACHE (CHANGED | ALL | SYSTEM | NONE)
	)
	;

partitionElement
	: (
	PARTITION INTEGERLITERAL
	(ENDING AT? LPAREN 
		(literal | MAXVALUE | MINVALUE) (COMMA (literal | MAXVALUE | MINVALUE))* 
	RPAREN INCLUSIVE?)?
	)
	;

functionParameterType
	: (functionDataType (AS LOCATOR)?)
	;

functionDataType
	: (functionBuiltInType | distinctTypeName)
	;

functionBuiltInType
	: (
	SMALLINT
	| INTEGER
	| INT
	| BIGINT
	| ((DECIMAL | DEC | NUMERIC) (integerInParens | (LPAREN RPAREN))?)
	| (DECFLOAT (integerInParens | (LPAREN RPAREN))?)
	| (FLOAT (integerInParens | (LPAREN RPAREN))?)
	| REAL
	| (DOUBLE PRECISION?)
	| ((((CHARACTER | CHAR) VARYING? ) | VARCHAR) (length | (LPAREN RPAREN))? ccsidClause1? forDataQualifier?)
	| ((((CHARACTER | CHAR) LARGE OBJECT) | CLOB) (length | (LPAREN RPAREN))? ccsidClause1? forDataQualifier?)
	| ((GRAPHIC | VARGRAPHIC | DBCLOB) (length | (LPAREN RPAREN))? ccsidClause1?)
	| (BINARY (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY VARYING?) | VARBINARY) (integerInParens | (LPAREN RPAREN)))
	| (((BINARY LARGE OBJECT) | BLOB) length?)
	| DATE
	| TIME
	| (TIMESTAMP integerInParens? ((WITH | WITHOUT) TIME ZONE)?)
	| ROWID
	| XML
	)
	;

procedureBuiltinType
	: (
	SMALLINT
	| INTEGER
	| INT
	| BIGINT
	| ((DECIMAL | DEC | NUMERIC) (integerInParens | (LPAREN RPAREN))?)
	| (DECFLOAT (integerInParens | (LPAREN RPAREN))?)
	| (FLOAT (integerInParens | (LPAREN RPAREN))?)
	| REAL
	| (DOUBLE PRECISION?)
	| ((((CHARACTER | CHAR) VARYING? ) | VARCHAR) (length | (LPAREN RPAREN))? ccsidClause1? forDataQualifier?)
	| ((((CHARACTER | CHAR) LARGE OBJECT) | CLOB) (length | (LPAREN RPAREN))? ccsidClause1? forDataQualifier?)
	| ((GRAPHIC | VARGRAPHIC | DBCLOB) (length | (LPAREN RPAREN))? ccsidClause1?)
	| (BINARY (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY VARYING?) | VARBINARY) (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY LARGE OBJECT) | BLOB) length?)
	| DATE
	| TIME
	| (TIMESTAMP integerInParens? ((WITH | WITHOUT) TIME ZONE)?)
	| ROWID
	)
	;

createTypeArrayBuiltinType
	: (
	SMALLINT
	| INTEGER
	| INT
	| BIGINT
	| ((DECIMAL | DEC | NUMERIC) (integerInParens | (LPAREN RPAREN)))
	| (DECFLOAT (integerInParens | (LPAREN RPAREN)))
	| (FLOAT (integerInParens | (LPAREN RPAREN)))
	| REAL
	| (DOUBLE PRECISION?)
	| ((((CHARACTER | CHAR) VARYING? ) | VARCHAR) length? ccsidClause1? forDataQualifier?)
	| ((((CHARACTER | CHAR) LARGE OBJECT) | CLOB) length? ccsidClause1? forDataQualifier?)
	| ((GRAPHIC | VARGRAPHIC | DBCLOB) (length | (LPAREN RPAREN))? ccsidClause1?)
	| (BINARY (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY VARYING?) | VARBINARY) (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY LARGE OBJECT) | BLOB) length?)
	| DATE
	| TIME
	| (TIMESTAMP integerInParens? ((WITH | WITHOUT) TIME ZONE)?)
	)
	;

createTypeArrayBuiltinType2
	: (
	INTEGER
	| INT
	| ((((CHARACTER | CHAR) VARYING? ) | VARCHAR) length? ccsidClause1? forDataQualifier?)
	)
	;

createVariableBuiltInType
	: (
	SMALLINT
	| INTEGER
	| INT
	| BIGINT
	| ((DECIMAL | DEC | NUMERIC) (integerInParens | (LPAREN RPAREN)))
	| (DECFLOAT (integerInParens | (LPAREN RPAREN)))
	| (FLOAT (integerInParens | (LPAREN RPAREN)))
	| REAL
	| (DOUBLE PRECISION?)
	| ((((CHARACTER | CHAR) VARYING? ) | VARCHAR) length? forDataQualifier?)
	| ((((CHARACTER | CHAR) LARGE OBJECT) | CLOB) length? forDataQualifier?)
	| ((GRAPHIC | VARGRAPHIC | DBCLOB) length?)
	| (BINARY (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY VARYING?) | VARBINARY) (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY LARGE OBJECT) | BLOB) length?)
	| DATE
	| TIME
	| (TIMESTAMP integerInParens? ((WITH | WITHOUT) TIME ZONE)?)
	)
	;

sourceDataType
	: procedureBuiltinType
	;

/*
Removed languageOption5 from createProcedureOptionList so a
distinction could be made between external (mandatory languageOption5) 
and native SQL (optional languageOption1) stored procedures.
*/
createProcedureOptionList
	: (
	specificNameOption2
	| dynamicResultSetOption
	| parameterOption1
	| externalNameOption1
	| sqlDataOption3
	| parameterStyleOption3
	| deterministicOption
	| packagePathOption
	| fencedOption
	| dbinfoOption
	| collectionIdOption
	| wlmEnvironmentOption2
	| asuTimeOption
	| stayResidentOption
	| programTypeOption
	| securityOption
	| runOptionsOption
	| commitOnReturnOption
	| specialRegistersOption
	| nullInputOption2
	| stopAfterFailureOption
	| debugOption
	)
	;

versionOption
	: (VERSION (identifier | literal))
	;

schemaQualifierOption
	: (QUALIFIER schemaName)
	;

currentDataOption
	: (CURRENT DATA (YES | NO))
	;

immediateWriteOption
	: ((WITH | WITHOUT) IMMEDIATE WRITE)
	;

explainOption
	: ((WITH | WITHOUT) EXPLAIN)
	;

reoptOption
	: (REOPT (NONE | ALWAYS | ONCE))
	;

packageOwnerOption
	: (PACKAGE OWNER authorizationName (AS (ROLE | USER))?)
	;

degreeOption
	: (DEGREE (ANY | INTEGERLITERAL))
	;

dynamicRulesOption
	: (DYNAMICRULES
		( RUN
		| BIND
		| DEFINEBIND
		| DEFINERUN
		| INVOKEBIND
		| INVOKERUN
		)
	)
	;

concurrentAccessOption
	: (CONCURRENT ACCESS RESOLUTION
		((USE CURRENTLY COMMITTED) | (WAIT FOR OUTCOME))
	)
	;

applicationEncodingOption
	: (APPLICATION ENCODING SCHEME (ASCII | EBCDIC | UNICODE))
	;

isolationLevelOption
	: (ISOLATION LEVEL (CS | RS | RR | UR))
	;

opthintOption
	: (OPTHINT NONNUMERICLITERAL)
	;

sqlPathOption
	: (SQL PATH sqlPathOptionItem (COMMA sqlPathOptionItem)*)
	;
	
sqlPathOptionItem
	: ((SYSTEM PATH) | (SESSION USER) | USER | SQLIDENTIFIER)
	;

queryAccelerationOption
	: (QUERY ACCELERATION queryAccelerationOptionItem)
	;

queryAccelerationOptionItem
	: (NONE
	| ENABLE
	| (ENABLE WITH FAILBACK)
	| ELIGIBLE
	| ALL
	)
	;

getAccelArchiveOption
	: (GET_ACCEL_ARCHIVE (NO | YES))
	;

accelerationOption
	: (ACCELERATION WAITFORDATA NUMERICLITERAL)
	;

businessTimeSensitiveOption
	: (BUSINESS_TIME SENSITIVE (YES | NO))
	;

systemTimeSensitiveOption
	: (SYSTEM_TIME SENSITIVE (YES | NO))
	;

archiveSensitiveOption
	: (ARCHIVE SENSITIVE (YES | NO))
	;

applcompatOption
	: (APPLCOMPAT APPLCOMPAT_LEVEL)
	;

validateOption
	: (VALIDATE (RUN | BIND))
	;

roundingOption
	: (ROUNDING roundingOptionItem)
	;

roundingOptionItem
	: ( DEC_ROUND_CEILING
	| DEC_ROUND_DOWN
	| DEC_ROUND_FLOOR
	| DEC_ROUND_HALF_DOWN
	| DEC_ROUND_HALF_EVEN
	| DEC_ROUND_HALF_UP
	| DEC_ROUND_UP
	)
	;

dateFormatOption
	: (DATE FORMAT dateTimeFormatOptionItem)
	;

dateTimeFormatOptionItem
	: (ISO | EUR | USA | JIS | LOCAL)
	;

timeFormatOption
	: (TIME FORMAT dateTimeFormatOptionItem)
	;

decimalOption
	: (DECIMAL LPAREN INTEGERLITERAL (COMMA INTEGERLITERAL)? RPAREN)
	;

forUpdateOption
	: (FOR UPDATE CLAUSE (REQUIRED | OPTIONAL))
	;

concentrateStatementsOption
	: (CONCENTRATE STATEMENTS (OFF | (WITH LITERALS)))
	;

acceleratorOption
	: (ACCELERATOR acceleratorName)
	;

procedureDataType
	: (procedureBuiltinType | distinctTypeName)
	;

alterSequenceOptionList
	: (
	restartOption
	| incrementOption
	| minvalueOption
	| maxvalueOption
	| cycleOption
	| cacheOption
	| orderOption
	)
	;

createSequenceOptionList
	: (
	asTypeOption
	| startOption
	| incrementOption
	| minvalueOption
	| maxvalueOption
	| cycleOption
	| cacheOption
	| orderOption
	)
	;

asTypeOption
	: (AS sequenceDataType)
	;

startOption
	: (START WITH INTEGERLITERAL)
	;

restartOption
	: (RESTART (WITH INTEGERLITERAL)?)
	;

incrementOption
	: (INCREMENT BY INTEGERLITERAL)
	;

minvalueOption
	: ((NO MINVALUE) | (MINVALUE INTEGERLITERAL))
	;

maxvalueOption
	: ((NO MAXVALUE) | (MAXVALUE INTEGERLITERAL))
	;

cycleOption
	: (NO? CYCLE)
	;

cacheOption
	: ((NO CACHE) | (CACHE INTEGERLITERAL))
	;

orderOption
	: (NO? ORDER)
	;

keyLabelOption
	: ((NO KEY LABEL) | (KEY LABEL keyLabelName))
	;

/*
Added..

	| (referentialConstraint)

...as it seems to be allowed on its own.  Noted by Martijn Rutte 2023-01-10.
*/
alterTableOptionList
	: (
	(ADD COLUMN? alterTableColumnDefinition)
	| (ALTER COLUMN? columnAlteration)
	| (RENAME COLUMN sourceColumnName TO targetColumnName)
	| (DROP COLUMN? columnName (CASCADE | RESTRICT)?)
	| (ADD periodDefinition)
	| DROP PERIOD identifier
	| (ADD (uniqueConstraint | referentialConstraint | checkConstraint | distributionClause))
	| (referentialConstraint)
	| (DROP ((PRIMARY KEY) | ((UNIQUE | (FOREIGN KEY) | CHECK | CONSTRAINT) constraintName)))
	| (ADD partitioningClause)
	| ADD (CHECK | (FOREIGN KEY)) constraintName constraintAlteration
	| DROP DISTRIBUTION
	| (ADD ((MATERIALIZED QUERY) | QUERY)? materializedQueryDefinition)
	| (DROP MATERIALIZED? QUERY)
	| dataCaptureClause
	| ACTIVATE NOT LOGGED INITIALLY (WITH EMPTY TABLE)?
	| PCTFREE INTEGERLITERAL
	| LOCKSIZE (ROW | TABLE | BLOCKINSERT)
	| (ADD RESTRICT ON DROP)
	| (DROP RESTRICT ON DROP)
	| ((ACTIVATE | DEACTIVATE) ROW ACCESS CONTROL)
	| ((ACTIVATE | DEACTIVATE) COLUMN ACCESS CONTROL)
	| appendClause
	| NO? VOLATILE CARDINALITY?
	| COMPRESS (YES (ADAPTIVE | STATIC)? | NO)
	| (ACTIVATE | DEACTIVATE) VALUE COMPRESSION
	| LOG INDEX BUILD (NULL | ON | OFF)
	)
	;

constraintAttributes
    : ((ENABLE | DISABLE) QUERY OPTIMIZATION
    | ENFORCED
    | NOT ENFORCED (NOT? TRUSTED)?
    )+
    ;

constraintAlteration
    : ((ENABLE | DISABLE) QUERY OPTIMIZATION
    | ENFORCED
    | NOT ENFORCED (NOT? TRUSTED)?
    )+
    ;

alterTablespaceOptionList
	: (
	bufferpoolOption
	| ccsidClause2
	| closeOption
	| compressOption
	| (DROP PENDING CHANGES)
	| dssizeOption
	| insertAlgorithmOption
	| lockmaxOption
	| locksizeOption
	| loggedOption
	| maxrowsOption
	| maxpartitionsOption
	| (MEMBER CLUSTER (YES | NO))
	| segsizeOption
	| trackmodClause
	| (usingBlock)
	| (freeBlock)
	| (gbpcacheBlock)
	| (PAGENUM RELATIVE)
	)
	;

/*
2022-11-01 Maarten van Haasteren noticed this rule should
reference usingSpecification2 instead of usingBlock.
2023-01-10 Martijn Rutte noticed this rule should allow
pagenumClause on its own and not just as part of
partitionByRangeSpecification.
*/
createTablespaceOptionList
	: (
	inDatabaseOption
	| bufferpoolOption
	| partitionByGrowthSpecification
	| partitionByRangeSpecification
	| segsizeOption
	| ccsidClause1
	| closeOption
	| compressOption
	| defineOption
	| freeBlock
	| gbpcacheBlock
	| insertAlgorithmOption
	| lockmaxOption
	| locksizeOption
	| loggedOption
	| maxrowsOption
	| maxpartitionsOption
	| memberClause
	| trackmodClause
	| pagenumClause
	| usingSpecification2
	)
	;

/*
This rule does not strictly follow the syntax diagram, as the diagram
is at odds with at least one example and arguably with the narrative.

More specifically, it is unclear where the ALTER keyword is required,
and so this rule makes it optional where its use seems to me to be
ambiguously documented.
*/
trustedContextOptionList
	: (
	(ALTER SYSTEM AUTHID authorizationName)
	| (ALTER NO DEFAULT ROLE)
	| (ALTER DEFAULT ROLE roleName 
		((WITHOUT ROLE AS OBJECT OWNER) | (WITH ROLE AS OBJECT OWNER AND QUALIFIER))?)
	| (ALTER? ENABLE)
	| (ALTER? DISABLE)
	| (ALTER? NO DEFAULT SECURITY LABEL)
	| (ALTER? DEFAULT SECURITY LABEL seclabelName)
	| (ALTER ATTRIBUTES LPAREN alterAttributesOptions (COMMA alterAttributesOptions)* RPAREN)
	| (ADD ATTRIBUTES LPAREN addAttributesOptions (COMMA addAttributesOptions)* RPAREN)
	| (DROP ATTRIBUTES LPAREN dropAttributesOptions (COMMA dropAttributesOptions)* RPAREN)
	)
	;

databaseOptionList
	: (
	bufferpoolOption
	| (INDEXBP bpName)
	| (AS WORKFILE (FOR memberName)?)
	| (STOGROUP ( SYSDEFLT | stogroupName)?)
	| ccsidClause1
	)
	;

createIndexOptionList
	: (
	(xmlIndexSpecification)
	| includeColumnPhrase
	| clusterOption
	| (PARTITIONED)
	| paddedOption
	| compressOption
	| usingSpecification2
	| freeSpecification
	| gbpcacheSpecification
	| defineOption
	| ((INCLUDE | EXCLUDE) NULL KEYS)
	| (PARTITION BY RANGE? LPAREN
		partitionElement (usingSpecification2 | freeSpecification | gbpcacheSpecification | dssizeOption)*
		(COMMA partitionElement (usingSpecification2 | freeSpecification | gbpcacheSpecification | dssizeOption)*)* RPAREN)
	| bufferpoolOption
	| closeOption
	| (DEFER (NO | YES))
	| dssizeOption
	| piecesizeOption
	| copyOption
	)
	;

inDatabaseOption
	: (IN databaseName)
	;

segsizeOption
	: (SEGSIZE INTEGERLITERAL)
	;

numpartsOption
	: (NUMPARTS INTEGERLITERAL)
	;

partitionByGrowthSpecification
	: (
	(maxpartitionsOption numpartsOption?)
	| dssizeOption
	)
	;

partitionByRangeSpecification
	: (
	numpartsOption
		(
			(LPAREN 
			(partitionByRangePartitionPhrase (COMMA partitionByRangePartitionPhrase)*)*
			RPAREN)
		|	pagenumClause
		|	dssizeOption
		)*
	)
	;

/*
2022-11-02 Maarten van Haasteren noticed this rule should
reference usingSpecification2 instead of usingBlock.
*/
partitionByRangePartitionPhrase
	: (
	(PARTITION | PART) INTEGERLITERAL
		(
		usingSpecification2
		| freeBlock
		| gbpcacheBlock
		| compressOption
		| trackmodClause
		| dssizeOption
		)*
	)
	;

insertAlgorithmOption
	: (INSERT ALGORITHM INTEGERLITERAL)
	;

maxrowsOption
	: (MAXROWS INTEGERLITERAL)
	;

maxpartitionsOption
	: (MAXPARTITIONS INTEGERLITERAL)
	;

usingSpecification2
	: (
	USING
		((STOGROUP stogroupName stogroupOptions*)
		| (VCAT catalogName))
	)
	;

/*
Refactored these options out of usingSpecification2 per
suggestion by Martijn Rutte 2022-12-06.  More concise
processing via SonarQube.
*/
stogroupOptions
	: (
	(PRIQTY INTEGERLITERAL)
	| (SECQTY INTEGERLITERAL)
	| (ERASE (NO | YES))
	)
	;

xmlIndexSpecification
	: (
	GENERATE (KEY | KEYS) USING XMLPATTERN xmlPatternClause AS SQL sqlDataType
	)
	;

/*
An xmlPatternClause has nontrivial syntax, but for purposes of this
grammar it is simply a quoted string.  It probably warrants a grammar
of its own.
*/
xmlPatternClause
	: NONNUMERICLITERAL
	;

alterAttributesOptions
	: (
	(ADDRESS addressValue)
	| (ENCRYPTION encryptionValue)
	| (SERVAUTH servauthValue)
	| (JOBNAME jobnameValue)
	)
	;

addAttributesOptions
	: (
	(ADDRESS addressValue)
	| (SERVAUTH servauthValue)
	| (JOBNAME jobnameValue)
	)
	;

dropAttributesOptions
	: (
	(ADDRESS addressValue?)
	| (SERVAUTH servauthValue?)
	| (JOBNAME jobnameValue?)
	)
	;

includeColumnPhrase
	: (INCLUDE LPAREN columnName (COMMA columnName)* RPAREN)
	;

alterPartitionClause
	: (
	((ALTER? PARTITION) | PART) INTEGERLITERAL
		( (usingBlock)
		| (freeBlock)
		| (gbpcacheBlock)
		| compressOption
		| dssizeOption
		| trackmodClause
		)+
	)
	;

usingBlock
	: (
	usingSpecification1+
	)
	;

/*
Removed '+' modifier per suggestion by Martijn Rutte 2022-12-14
as its absence makes rule processing easier.
*/
freeBlock
	: (
	((FREEPAGE INTEGERLITERAL)
	| (PCTFREE INTEGERLITERAL)
	| (PCTFREE (INTEGERLITERAL? FOR UPDATE INTEGERLITERAL)?)
	)
	)
	;

moveTableClause
	: (
	MOVE TABLE tableName TO TABLESPACE (databaseName DOT)? tablespaceName
	)
	;

gbpcacheBlock
	: (
	GBPCACHE (CHANGED | ALL | SYSTEM | NONE)
	)
	;

aliasDesignator
	: (
	PUBLIC? ALIAS aliasName FOR (TABLE | SEQUENCE)
	)
	;

multipleColumnList
	: (
	tableName LPAREN
	columnName IS NONNUMERICLITERAL
	(COMMA columnName IS NONNUMERICLITERAL)*
	RPAREN
	)
	;

alterTableColumnDefinitionOptionList1
	: (
	(defaultClause1)
	| (NOT NULL)
	| (columnConstraint)
	| (generatedClause)
	| implicitlyHiddenClause
    | COMPRESS SYSTEM DEFAULT
    | COLUMN? SECURED WITH identifier
    | NOT HIDDEN_
    | SCOPE tableName
    | lobOptions
	)
	;

lobOptions
    : (
     NOT? LOGGED
    | NOT? COMPACT
    )+
    ;

/*
Changed to refer to defaultClause instead of defaultClause2
as these are identical.
*/
alterTableColumnDefinitionOptionList2
	: (
	(defaultClause)
	| (NOT NULL)
	| (columnConstraint)
	| (generatedClause)
	| implicitlyHiddenClause
    | COMPRESS SYSTEM DEFAULT
    | COLUMN? SECURED WITH identifier
    | NOT HIDDEN_
    | SCOPE tableName
    | lobOptions
	)
	;

columnConstraint
	: (referencesClause | checkConstraint)
	;

/*
Made...

	(asIdentityClause | asRowChangeTimestampClause)

...optional as it was not by by mistake.  Noted by Martijn Rutte 2023-01-10.
*/
generatedClause
	: (
	(GENERATED (ALWAYS | (BY DEFAULT))? (asIdentityClause | asRowChangeTimestampClause)?)
	| (GENERATED ALWAYS?
		(asRowTransactionStartIDClause 
		| asRowTransactionTimestampClause 
		| asGeneratedExpressionClause))
	)
	;

generatedClause2
	: (GENERATED (ALWAYS | (BY DEFAULT)) asIdentityClause?)
	;

asIdentityClause
	: (
	AS IDENTITY (LPAREN 
	asIdentityClauseOptionList (COMMA? asIdentityClauseOptionList)* 
	RPAREN)?
	)
	;

asIdentityClauseOptionList
	: (
	startOption
	| incrementOption
	| minvalueOption
	| maxvalueOption
	| cycleOption
	| cacheOption
	| orderOption
	)
	;

asRowChangeTimestampClause
	: (FOR EACH ROW ON UPDATE AS ROW CHANGE TIMESTAMP)
	;

asRowTransactionStartIDClause
	: (AS TRANSACTION START ID)
	;

asRowTransactionTimestampClause
	: (AS ROW (BEGIN | START | END))
	;

asGeneratedExpressionClause
	: (AS LPAREN nonDeterministicExpression RPAREN)
	;

nonDeterministicExpression
	: (
	(DATA CHANGE OPERATION)
	| specialRegister
	| nonDeterministicExpressionSessionVariable
	)
	;

nonDeterministicExpressionSessionVariable
	: (
	(SYSIBM DOT PACKAGE_NAME)
	| (SYSIBM DOT PACKAGE_SCHEMA)
	| (SYSIBM DOT PACKAGE_VERSION)
	)
	;

columnAlteration
	: (columnName columnAlterationOptionList+)
	;

columnAlterationOptionList
	: (
	(SET DATA TYPE alteredDataType (INLINE LENGTH INTEGERLITERAL)?)
	| (SET defaultClause)
	| (SET INLINE LENGTH INTEGERLITERAL)
	| (SET GENERATED (ALWAYS | (BY DEFAULT)) identityAlteration?)
	| (identityAlteration)
	| (SET GENERATED ALWAYS? (asRowTransactionTimestampClause | asRowTransactionStartIDClause))
	| (DROP DEFAULT)
	)
	;

/*
In the IBM documentation, alteredDataType differs from dataType in that
it is a proper subset thereof.  The dataType rule includes a provision 
for CCSID on CHAR, VARCHAR, CLOB, GRAPHIC, VARGRAPHIC, and DBCLOB types
which is absent from the alteredDataType rule.  For purposes of this
grammar, a difference which makes no difference is no difference.
*/
alteredDataType
	: dataType
	;

/*
The difference between dataType and castDataType is in the coding
of the CCSID and FOR ... DATA qualifiers.  Sneaky.
*/
dataType
	: (builtInType | distinctTypeName)
	;

builtInType
	: (
	SMALLINT
	| INTEGER
	| INT
	| BIGINT
	| ((DECIMAL | DEC | NUMERIC) (integerInParens | (LPAREN RPAREN))?)
	| (DECFLOAT (integerInParens | (LPAREN RPAREN))?)
	| (FLOAT (integerInParens | (LPAREN RPAREN))?)
	| REAL
	| (DOUBLE PRECISION?)
	| ((((CHARACTER | CHAR) VARYING? ) | VARCHAR) (length | (LPAREN RPAREN))? (forDataQualifier | ccsidClause2)?)
	| ((((CHARACTER | CHAR) LARGE OBJECT) | CLOB) (length | (LPAREN RPAREN))? (forDataQualifier | ccsidClause2)?)
	| ((GRAPHIC | VARGRAPHIC | DBCLOB) (length | (LPAREN RPAREN))? ccsidClause2?)
	| (BINARY (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY VARYING?) | VARBINARY) (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY LARGE OBJECT) | BLOB) length?)
	| DATE
	| TIME
	| (TIMESTAMP integerInParens? ((WITH | WITHOUT) TIME ZONE)?)
	| ROWID
	| (XML (LPAREN xmlTypeModifier RPAREN)?)
	)
	;

sequenceDataType
	: (sequenceBuiltInType | distinctTypeName)
	;

sequenceBuiltInType
	: (
	SMALLINT
	| INTEGER
	| INT
	| BIGINT
	| ((DECIMAL | DEC | NUMERIC) integerInParens?)
	)
	;

sqlDataType
	: (
	(VARCHAR LPAREN INTEGERLITERAL RPAREN)
	| (DECFLOAT (LPAREN INTEGERLITERAL RPAREN)?)
	| DATE
	| (TIMESTAMP (LPAREN INTEGERLITERAL RPAREN)?)
	)
	;

xmlTypeModifier
	: (
	XMLSCHEMA 
	xmlSchemaSpecification (ELEMENT xmlElementName)?
	(COMMA xmlSchemaSpecification (ELEMENT xmlElementName)?)*
	)
	;

xmlSchemaSpecification
	: (
	(ID registeredXmlSchemaName)
	| (((URL targetNamespace) | (NO NAMESPACE)) (LOCATION schemaLocation)?)
	)
	;

/*
Documentation is a bit sketchy on details for the following
four items.  Examples would be nice.
*/
xmlElementName
	: (identifier | literal)
	;

piName
	: (NONNUMERICLITERAL)
	;

registeredXmlSchemaName
	: (
	SYSXSR DOT sqlidentifier
	)
	;

targetNamespace
	: (NONNUMERICLITERAL)
	;

schemaLocation
	: (NONNUMERICLITERAL)
	;

identityAlteration
	: (
	(RESTART (WITH INTEGERLITERAL)?)
	| (SET incrementOption)
	| (SET minvalueOption)
	| (SET maxvalueOption)
	| (SET cycleOption)
	| (SET cacheOption)
	| (SET orderOption)
	)
	;

uniqueConstraint
	: (
	(CONSTRAINT constraintName)? 
	((PRIMARY KEY) | UNIQUE) 
	LPAREN
	columnName (COMMA columnName)* 
	(COMMA BUSINESS_TIME WITHOUT OVERLAPS)? 
	RPAREN
	constraintAttributes?
	)
	;

referentialConstraint
	: (
	((CONSTRAINT constraintName FOREIGN KEY) | (FOREIGN KEY constraintName?))
	LPAREN
	columnName (PERIOD BUSINESS_TIME)? (COMMA columnName (PERIOD BUSINESS_TIME)?)* 
	RPAREN
	referencesClause
	)
	;

referencesClause
	: (
	REFERENCES tableName 
	(LPAREN
	columnName (PERIOD BUSINESS_TIME)? (COMMA columnName (PERIOD BUSINESS_TIME)?)* 
	RPAREN)?
	(ON DELETE (RESTRICT | (NO ACTION) | CASCADE | (SET NULL)))? 
	(NOT? ENFORCED)?
	(ENABLE QUERY OPTIMIZATION)?	
	)
	;

checkConstraint
	: (
	(CONSTRAINT constraintName)? CHECK LPAREN checkCondition RPAREN
	)
	;

partitioningClause
	: (
	PARTITION BY 
		((RANGE? 
		LPAREN 
		partitionExpression (COMMA partitionExpression)* 
		RPAREN
		LPAREN
		partitioningClauseElement (COMMA partitioningClauseElement)*
		RPAREN)
		| (SIZE (EVERY INTEGERLITERAL? sqlidentifier)?))
	)
	;

/*
Made...

	(ASC | DESC)

...optional because it wasn't and it should be.  Found by Martijn Rutte 2022-01-09.
*/
partitionExpression
	: (
	columnName (NULLS LAST)? (ASC | DESC)?
	)
	;

/*
Changed...

	(INTEGERLITERAL | MAXVALUE | MINVALUE)

...to...

	(literal | MAXVALUE | MINVALUE)

...because the literal must match the data type of the column
it references.  Found by Martijn Rutte 2022-01-09.
*/

/*
The partitionHashSpace rule can be before the INCLUSIVE token in
the create table statement, or after the INCLUSIVE token in the
alter table statement.  Also, it's deprecated as of Db2 12.
*/

partitioningClauseElement
	: (
	(PARTITION partitionName=identifier)? boundarySpec partitionTablespaceOptions?
	| boundarySpec EVERY (
	LPAREN literal durationSuffix? RPAREN
	| literal durationSuffix?
	)
	)
	;

partitionTablespaceOptions
    : (IN tablespaceName)
      (INDEX IN tablespaceName (LONG IN tablespaceName)?)?
    |
     (IN tablespaceName)?
           (INDEX IN tablespaceName (LONG IN tablespaceName)?)
    ;

materializedQueryDefinition
	: (
	LPAREN fullSelect RPAREN refreshableTableOptions
	)
	;

refreshableTableOptions
	: (dataInitiallyDeferredPhrase refreshDeferredPhrase refreshableTableOptionsList*)
	;

dataInitiallyDeferredPhrase
	: (DATA INITIALLY DEFERRED)
	;

refreshDeferredPhrase
	: (REFRESH DEFERRED)
	;

refreshableTableOptionsList
	: (
	(MAINTAINED BY (SYSTEM | USER))
	| (enableDisableOption QUERY OPTIMIZATION)
	)
	;

periodDefinition
	: (
	PERIOD FOR?
	((SYSTEM_TIME LPAREN beginColumnName COMMA endColumnName RPAREN)
	| (BUSINESS_TIME LPAREN beginColumnName COMMA endColumnName (EXCLUSIVE | INCLUSIVE) RPAREN))
	)
	;

alterTableColumnDefinition
	: (
	(columnName builtInType alterTableColumnDefinitionOptionList1*)
	| (columnName distinctTypeName alterTableColumnDefinitionOptionList2*)
	)
	;

externalProgramName
	: (identifier | NONNUMERICLITERAL)
	;

packagePath
	: (
	collectionID
	| SESSION_USER
	| USER
	| (CURRENT PACKAGE PATH)
	| (CURRENT PATH)
	| hostVariable
	| NONNUMERICLITERAL
	)
	;

collectionID
	: identifier
	;

runTimeOptions
	: NONNUMERICLITERAL
	;

comparisonOperator
	: (EQ | GT | LT | GE | LE | NE)
	;

operator
	: (SPLAT | PLUS | MINUS | SLASH | CONCAT | CONCATOP)
	;

expression
	: (
	functionInvocation
	| LPAREN expression RPAREN
	| labeledDuration
	| literal
	| specialRegister
	| columnName
	| hostVariable
	| scalarFullSelect
	| timeZoneSpecificExpression
	| caseExpression
	| castSpecification
	| xmlCastSpecification
	| xmlParseSpecification
	| arrayElementSpecification
	| arrayConstructor
	| olapSpecification
	| rowChangeExpression
	| sequenceReference
	| ((operator | INTEGERLITERAL) expression)
	| ((functionInvocation
		| LPAREN expression RPAREN
		| labeledDuration
		| literal
		| specialRegister
		| columnName
		| hostVariable
		| scalarFullSelect
		| timeZoneSpecificExpression
		| caseExpression
		| castSpecification
		| xmlCastSpecification
		| xmlParseSpecification
		| arrayElementSpecification
		| arrayConstructor
		| olapSpecification
		| rowChangeExpression
		| sequenceReference)
		((operator | INTEGERLITERAL) expression)*)
	)
	;

keyExpression
	: (expression)
	;

rowChangeExpression
	: ROW CHANGE (TIMESTAMP | TOKEN) FOR tableName
	;

sequenceReference
	: (NEXT | PREVIOUS) VALUE FOR tableName
	;

functionInvocation
	: (
	scalarFunctionInvocation
	| aggregateFunctionInvocation
	| regressionFunctionInvocation
	| externalFunctionInvocation
	)
	;

scalarFunctionInvocation
	: (
	xmlattributesFunction
	| xmlelementFunction
	| xmlforestFunction
	| xmlmodifyFunction
	| xmlnamespaceFunction
	| xmlpiFunction
	| xmlqueryFunction
	| xmlserializeFunction
	| aiAnalogyFunction
	| aiSemanticClusterFunction
	| aiSimilarityFunction
	| ((schemaName DOT)? scalarFunction LPAREN (expression (COMMA expression)*)? RPAREN (AS NONNUMERICLITERAL)?)
	)
	;

aggregateFunctionInvocation
	: (
	arrayaggFunction
	| correlationFunction
	| covarianceFunction
	| covarianceSampFunction
	| cumeDistFunction
	| listaggFunction
	| percentileContFunction
	| percentileDiscFunction
	| percentRankFunction
	| xmlaggFunction
	| ((schemaName DOT)? aggregateFunction LPAREN DISTINCT? (expression | SPLAT) RPAREN)
	)
	;

regressionFunctionInvocation
	: ((schemaName DOT)? regressionFunction
	LPAREN
	expression COMMA expression
	RPAREN)
	;

externalFunctionInvocation
	: ((schemaName DOT)? sqlidentifier
	LPAREN
	expression (COMMA expression)*
	RPAREN)
	;

labeledDuration
	: (
	(functionInvocation
	| (LPAREN expression RPAREN)
	| INTEGERLITERAL
	| columnName
	| variable)
	durationSuffix
	)
	;

durationSuffix
	: (
	YEAR
	| YEARS
	| MONTH
	| MONTHS
	| DAY
	| DAYS
	| HOUR
	| HOURS
	| MINUTE
	| MINUTES
	| SECOND
	| SECONDS
	| MICROSECOND
	| MICROSECONDS
	)
	;

/*
Corrected to add LPAREN and RPAREN per Martijn Rutte's discovery
that XMLCAST invocations did not parse correctly.  2023-01-24
*/
xmlCastSpecification
	: XMLCAST LPAREN (expression | NULL | parameterMarker) AS dataType RPAREN
	;

xmlParseSpecification
	: XMLPARSE LPAREN DOCUMENT expression ((STRIP | PRESERVE) WHITESPACE)? RPAREN
	;

arrayElementSpecification
	: arrayExpression OPENSQBRACKET arrayIndex CLOSESQBRACKET
	;

arrayIndex
	: expression (operator? expression)*
	;

arrayConstructor
	: ARRAY
	OPENSQBRACKET
	(
	QUESTIONMARK
	| fullSelect
	| ((expression | NULL) (COMMA (expression | NULL))*)
	)
	CLOSESQBRACKET
	;
	
olapSpecification
	: orderedOlapSpecification
	| numberingSpecification
	| aggregationSpecification
	;

orderedOlapSpecification
	: olapSpecificationFunction
	OVER LPAREN windowPartitionClause? windowOrderClause RPAREN
	;

olapSpecificationFunction
	: (
	(CUME_DIST LPAREN RPAREN)
	| (PERCENT_RANK LPAREN RPAREN)
	| (RANK LPAREN RPAREN)
	| (DENSE_RANK LPAREN RPAREN)
	| (NTILE LPAREN expression RPAREN)
	| lagFunction
	| leadFunction
	)
	;

lagFunction
	: LAG LPAREN expression 
	(
	COMMA INTEGERLITERAL 
		(COMMA expression 
			(COMMA ((RESPECT NULLS) | (IGNORE NULLS)))?)? RPAREN
	)
	;

leadFunction
	: LEAD LPAREN expression 
	(
	COMMA INTEGERLITERAL 
		(COMMA expression 
			(COMMA respectNullsClause)?)? RPAREN
	)
	;

respectNullsClause
	: ((RESPECT NULLS) | (IGNORE NULLS))
	;

windowPartitionClause
	: (PARTITION BY expression (COMMA expression)*)
	;

windowOrderClause
	: ORDER BY expression windowOrderClauseQualifier? (COMMA expression windowOrderClauseQualifier?)*
	;

windowOrderClauseQualifier
	: (ASC | DESC) (NULLS (FIRST | LAST))?
	;

numberingSpecification
	: ROW_NUMBER LPAREN RPAREN OVER LPAREN windowPartitionClause? windowOrderClause? RPAREN
	;

aggregationSpecification
	: (aggregateFunctionInvocation | olapColumnFunction) OVER LPAREN windowPartitionClause?
	((RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING)
	| (windowOrderClause ((RANGE BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) | windowAggregationGroupClause)?))?
	RPAREN
	;

aggregateFunction
	: (
	ARRAY_AGG
	| AVG
	| CORR
	| CORRELATION
	| COUNT
	| COUNT_BIG
	| COVAR_POP
	| COVARIANCE
	| COVAR
	| COVAR_SAMP
	| COVARIANCE_SAMP
	| CUME_DIST
	| GROUPING
	| LISTAGG
	| MAX
	| MEDIAN
	| MIN
	| PERCENTILE_CONT
	| PERCENTILE_DISC
	| PERCENT_RANK
	| STDDEV_POP
	| STDDEV
	| STDDEV_SAMP
	| SUM
	| VAR_POP
	| VARIANCE
	| VAR
	| VAR_SAMP
	| VARIANCE_SAMP
	| XMLAGG
	)
	;

regressionFunction
	: (
	REGR_AVGX
	| REGR_AVGY
	| REGR_COUNT
	| REGR_INTERCEPT
	| REGR_ICPT
	| REGR_R2
	| REGR_SLOPE
	| REGR_SXX
	| REGR_SXY
	| REGR_SYY
	)
	;

olapColumnFunction
	: (
	firstValueFunction
	| lastValueFunction
	| nthValueFunction
	| ratioToReportFunction
	)
	;

firstValueFunction
	: (FIRST_VALUE LPAREN expression (COMMA respectNullsClause)? RPAREN)
	;

lastValueFunction
	: (LAST_VALUE LPAREN expression (COMMA respectNullsClause)? RPAREN)
	;

nthValueFunction
	: (NTH_VALUE LPAREN expression COMMA INTEGERLITERAL RPAREN)
	;

ratioToReportFunction
	: (RATIO_TO_REPORT LPAREN expression RPAREN)
	;

listaggFunction
	: (
	LISTAGG LPAREN (ALL | DISTINCT)? expression (COMMA NONNUMERICLITERAL)? RPAREN
	(WITHIN GROUP LPAREN 
		ORDER BY sortKey (ASC | DESC)? (COMMA sortKey (ASC | DESC)?)* 
	RPAREN)?
	)
	;

arrayaggFunction
	: (arrayaggOrdinaryFunction | arrayaggAssociativeFunction)
	;

arrayaggOrdinaryFunction
	: (
	ARRAY_AGG LPAREN expression
	(ORDER BY sortKey (ASC | DESC)? (COMMA sortKey (ASC | DESC)?)*)?
	RPAREN
	)
	;

arrayaggAssociativeFunction
	: (
	ARRAY_AGG LPAREN expression (COMMA expression)? RPAREN
	)
	;

correlationFunction
	: ((CORR | CORRELATION) LPAREN expression COMMA expression RPAREN)
	;

covarianceFunction
	: ((COVAR_POP | COVARIANCE | COVAR) LPAREN expression COMMA expression RPAREN)
	;

covarianceSampFunction
	: ((COVAR_SAMP | COVARIANCE_SAMP) LPAREN expression COMMA expression RPAREN)
	;

cumeDistFunction
	: (
	CUME_DIST LPAREN
	expression (COMMA expression)*
	RPAREN 
	WITHIN GROUP LPAREN aggregateOrderByClause RPAREN
	)
	;

percentileContFunction
	: (
	PERCENTILE_CONT LPAREN expression RPAREN
	WITHIN GROUP LPAREN ORDER BY expression (ASC | DESC)? RPAREN
	)
	;

percentileDiscFunction
	: (
	PERCENTILE_DISC LPAREN expression RPAREN
	WITHIN GROUP LPAREN ORDER BY expression (ASC | DESC)? RPAREN
	)
	;

percentRankFunction
	: (
	PERCENT_RANK LPAREN
	expression (COMMA expression)*
	RPAREN 
	WITHIN GROUP LPAREN aggregateOrderByClause RPAREN
	)
	;

xmlaggFunction
	: (XMLAGG LPAREN expression xmlaggOrderByClause? RPAREN)
	;

xmlaggOrderByClause
	: (ORDER BY xmlaggOrderByOption (COMMA xmlaggOrderByOption)*)
	;

xmlaggOrderByOption
	: (expression (ASC | DESC)?)
	;

aggregateOrderByClause
	: (ORDER BY aggregateOrderByOption (COMMA aggregateOrderByOption)*)
	;

aggregateOrderByOption
	: (sortKey (ASC | DESC)? (NULLS (LAST | FIRST))?)
	;

windowAggregationGroupClause
	: (ROWS | RANGE) (groupStart | groupBetween | groupEnd)
	;

groupStart
	: (unboundedPreceding | boundedPreceding | currentRow)
	;

groupBetween
	: BETWEEN groupBound1 AND groupBound2
	;

groupEnd
	: (unboundedFollowing | boundedFollowing)
	;

groupBound1
	: (unboundedPreceding | boundedPreceding | boundedFollowing | currentRow)
	;

groupBound2
	: (unboundedFollowing | boundedPreceding | boundedFollowing | currentRow)
	;

unboundedPreceding
	: UNBOUNDED PRECEDING
	;

unboundedFollowing
	: UNBOUNDED FOLLOWING
	;

boundedPreceding
	: INTEGERLITERAL PRECEDING
	;

boundedFollowing
	: INTEGERLITERAL FOLLOWING
	;

currentRow
	: CURRENT ROW
	;

scalarFunction
	: (
	ABS
	| ABSVAL
	| ACOS
	| ADD_DAYS
	| ADD_MONTHS
	| AI_ANALOGY
	| AI_SEMANTIC_CLUSTER
	| AI_SIMILARITY
	| APPLCOMPAT_LEVEL
	| ARRAY_DELETE
	| ARRAY_FIRST
	| ARRAY_LAST
	| ARRAY_NEXT
	| ARRAY_PRIOR
	| ARRAY_TRIM
	| ASCII
	| ASCII_CHR
	| ASCIISTR
	| ASCII_STR
	| ASIN
	| ATTACH
	| ATAN
	| ATAN2
	| ATANH
	| AUTOMATIC
	| BIGINT
	| BINARY
	| BITAND
	| BITANDNOT
	| BITNOT
	| BITOR
	| BITXOR
	| BLOB
	| BLOCKSIZE
	| BTRIM
	| CARDINALITY
	| CATEGORIES
	| CCSID_ENCODING
	| CEIL
	| CEILING
	| CHAR
	| CHAR9
	| CHARACTER_LENGTH
	| CHAR_LENGTH
	| CHR
	| CLOB
	| COALESCE
	| COLLATION_KEY
	| COMPARE_DECFLOAT
	| CONCAT
	| CONTAINS
	| COS
	| COSH
	| DATE
	| DAY
	| DAYOFMONTH
	| DAYOFWEEK
	| DAYOFWEEK_ISO
	| DAYOFYEAR
	| DAYS
	| DAYS_BETWEEN
	| DBCLOB
	| DEC
	| DECFLOAT
	| DECFLOAT_FORMAT
	| DECFLOAT_SORTKEY
	| DECIMAL
	| DECODE
	| DECRYPT_BINARY
	| DECRYPT_BIT
	| DECRYPT_CHAR
	| DECRYPT_DATAKEY_BIGINT
	| DECRYPT_DATAKEY_BIT
	| DECRYPT_DATAKEY_CLOB
	| DECRYPT_DATAKEY_DBCLOB
	| DECRYPT_DATAKEY_DECIMAL
	| DECRYPT_DATAKEY_INTEGER
	| DECRYPT_DATAKEY_VARCHAR
	| DECRYPT_DATAKEY_VARGRAPHIC
	| DECRYPT_DB
	| DEGREES
	| DIFFERENCE
	| DIGITS
	| DOUBLE
	| DOUBLE_PRECISION
	| DSN_XMLVALIDATE
	| EBCDIC_CHR
	| EBCDIC_STR
	| ENCRYPT_DATAKEY
	| ENCRYPT_TDES
	| ERROR
	| EXP
	| EXTRACT
	| FLOAT
	| FLOOR
	| GENERATE_UNIQUE
	| GENERATE_UNIQUE_BINARY
	| GETHINT
	| GETVARIABLE
	| GRAPHIC
	| GREATEST
	| HASH
	| HASH_CRC32
	| HASH_MD5
	| HASH_SHA1
	| HASH_SHA256
	| HEX
	| HOUR
	| IDENTITY_VAL_LOCAL
	| IFNULL
	| INSERT
	| INSTR
	| INT
	| INTEGER
	| JULIAN_DAY
	| LAST_DAY
	| LCASE
	| LEAST
	| LEFT
	| LENGTH
	| LN
	| LOCATE
	| LOCATE_IN_STRING
	| LOG10
	| LOWER
	| LPAD
	| LTRIM
	| MAX_CARDINALITY
	| MICROSECOND
	| MIDNIGHT_SECONDS
	| MINUTE
	| MOD
	| MONTH
	| MONTHS_BETWEEN
	| MQREAD
	| MQREADCLOB
	| MQRECEIVE
	| MQRECEIVECLOB
	| MQSEND
	| MULTIPLY_ALT
	| NEXT_DAY
	| NEXT_MONTH
	| NORMAL
	| NORMALIZE_DECFLOAT
	| NORMALIZE_STRING
	| NULLIF
	| NUMBLOCKPAGES
	| NVL
	| OVERLAY
	| PACK
	| POLICY
	| POSITION
	| POSSTR
	| POW
	| POWER
	| QUANTIZE
	| QUARTER
	| RADIANS
	| RAISE_ERROR
	| RAND
	| RANDOM
	| REAL
	| REGEXP_COUNT
	| REGEXP_INSTR
	| REGEXP_LIKE
	| REGEXP_REPLACE
	| REGEXP_SUBSTR
	| REPEAT
	| REPLACE
	| RID
	| RIGHT
	| ROUND
	| ROUND_TIMESTAMP
	| ROWID
	| RPAD
	| RTRIM
	| SCORE
	| SECOND
	| SIGN
	| SIN
	| SINH
	| SMALLINT
	| SOAPHTTPC
	| SOAPHTTPNC
	| SOAPHTTPNV
	| SOAPHTTPV
	| SOUNDEX
	| SPACE
	| SQRT
	| STATUS
	| STRIP
	| STRLEFT
	| STRPOS
	| STRRIGHT
	| SUBSTR
	| SUBSTRING
	| SUCCESS
	| TAN
	| TANH
	| TIME
	| TIMESTAMP
	| TIMESTAMPADD
	| TIMESTAMPDIFF
	| TIMESTAMP_FORMAT
	| TIMESTAMP_ISO
	| TIMESTAMP_TZ
	| TO_CHAR
	| TO_CLOB
	| TO_DATE
	| TO_NUMBER
	| TOTALORDER
	| TO_TIMESTAMP
	| TRANSLATE
	| TRIM
	| TRIM_ARRAY
	| TRUNC
	| TRUNCATE
	| TRUNC_TIMESTAMP
	| UCASE
	| UNICODE
	| UNICODE_STR
	| UNISTR
	| UPPER
	| VALUE
	| VARBINARY
	| VARCHAR
	| VARCHAR9
	| VARCHAR_BIT_FORMAT
	| VARCHAR_FORMAT
	| VARGRAPHIC
	| VERIFY_GROUP_FOR_USER
	| VERIFY_ROLE_FOR_USER
	| VERIFY_TRUSTED_CONTEXT_ROLE_FOR_USER
	| WEEK
	| WEEK_ISO
	| WRAP
	| XMLATTRIBUTES
	| XMLCOMMENT
	| XMLCONCAT
	| XMLDOCUMENT
	| XMLELEMENT
	| XMLFOREST
	| XMLMODIFY
	| XMLNAMESPACES
	| XMLPARSE
	| XMLPI
	| XMLQUERY
	| XMLSERIALIZE
	| XMLTEXT
	| XMLXSROBJECTID
	| XSLTRANSFORM
	| YEAR
	)
	;

tableFunction
	: (
	ADMIN_TASK_LIST
	| ADMIN_TASK_OUTPUT
	| ADMIN_TASK_STATUS
	| BLOCKING_THREADS
	| MQREADALL
	| MQREADALLCLOB
	| MQRECEIVEALL
	| MQRECEIVEALLCLOB
	| XMLTABLE
	)
	;
	
specialRegister
	: (
	CURRENT_ACCELERATOR
	| CURRENT_APPLICATION_COMPATIBILITY
	| CURRENT_APPLICATION_ENCODING_SCHEME
	| CURRENT_CLIENT_ACCTNG
	| CURRENT_CLIENT_APPLNAME
	| CURRENT_CLIENT_CORR_TOKEN
	| CURRENT_CLIENT_USERID
	| CURRENT_CLIENT_WRKSTNNAME
	| CURRENT_DATE
	| CURRENT_DEBUG_MODE
	| CURRENT_DECFLOAT_ROUNDING_MODE
	| CURRENT_DEGREE
	| CURRENT_EXPLAIN_MODE
	| CURRENT_GET_ACCEL_ARCHIVE
	| CURRENT_LOCALE_LC_CTYPE
	| CURRENT_MAINTAINED_TABLE_TYPES_FOR_OPTIMIZATION
	| CURRENT_MEMBER
	| CURRENT_OPTIMIZATION_HINT
	| CURRENT_PACKAGE_PATH
	| CURRENT_PACKAGESET
	| CURRENT_PATH
	| CURRENT_PRECISION
	| CURRENT_QUERY_ACCELERATION
	| CURRENT_QUERY_ACCELERATION_WAITFORDATA
	| CURRENT_REFRESH_AGE
	| CURRENT_ROUTINE_VERSION
	| CURRENT_RULES
	| CURRENT_SCHEMA
	| CURRENT_SERVER
	| CURRENT_SQLID
	| CURRENT_TEMPORAL_BUSINESS_TIME
	| CURRENT_TEMPORAL_SYSTEM_TIME
	| CURRENT_TIME
	| CURRENT_TIMESTAMP
	| CURRENT_TIME_ZONE
	| ENCRYPTION_PASSWORD
	| SESSION_TIME_ZONE
	| SESSION_USER
	| USER
	)
	;

aiAnalogyFunction
	: (
	AI_ANALOGY LPAREN
	aiAnalogyFunctionSource1 COMMA
	aiAnalogyFunctionTarget1 COMMA
	aiAnalogyFunctionSource2 COMMA
	aiAnalogyFunctionTarget2 
	RPAREN
	)
	;

aiFunctionExpression
	: (expression (USING (MODEL COLUMN)? columnName)?)
	;

aiAnalogyFunctionSource
	: aiFunctionExpression
	;

aiAnalogyFunctionTarget
	: aiFunctionExpression
	;

aiAnalogyFunctionSource1
	: aiAnalogyFunctionSource
	;

aiAnalogyFunctionSource2
	: aiAnalogyFunctionSource
	;

aiAnalogyFunctionTarget1
	: aiAnalogyFunctionTarget
	;

aiAnalogyFunctionTarget2
	: aiAnalogyFunctionTarget
	;

aiSemanticClusterFunction
	: (
	AI_SEMANTIC_CLUSTER LPAREN
	aiSemanticClusterMemberExpression COMMA
	aiSemanticClusterClusteringExpression
	(COMMA aiSemanticClusterClusteringExpression)*
	RPAREN
	)
	;

aiSemanticClusterMemberExpression
	: aiFunctionExpression
	;

aiSemanticClusterClusteringExpression
	: expression
	;

aiSimilarityFunction
	: (
	AI_SIMILARITY LPAREN
	aiSimilarityExpression1 COMMA
	aiSimilarityExpression2
	RPAREN
	)
	;

aiSimilarityExpression
	: aiFunctionExpression
	;

aiSimilarityExpression1
	: aiSimilarityExpression
	;

aiSimilarityExpression2
	: aiSimilarityExpression
	;

xmlelementFunction
	: (
	XMLELEMENT LPAREN NAME xmlElementName
	(COMMA xmlnamespacesDeclaration)?
	(COMMA xmlattributesFunction)?
	(COMMA expression)*
	xmlFunctionOptionClause?
	RPAREN
	)
	;

xmlforestFunction
	: (
	XMLFOREST LPAREN xmlnamespaceFunction?
	elementContentExpression (AS xmlElementName)?
	(COMMA elementContentExpression (AS xmlElementName)?)*
	xmlFunctionOptionClause?
	RPAREN
	)
	;

xmlmodifyFunction
	: (
	XMLMODIFY LPAREN expression
	(COMMA expression AS (NONNUMERICLITERAL | identifier))*
	RPAREN
	)
	;

xmlpiFunction
	: (
	XMLPI LPAREN NAME piName (COMMA expression)? RPAREN
	)
	;

xmlqueryFunction
	: (
	XMLQUERY LPAREN xqueryExpressionConstant
	(PASSING (BY REF)? xqueryArgument (COMMA xqueryArgument)*)?
	(RETURNING SEQUENCE (BY REF)?)? (EMPTY ON EMPTY)?
	RPAREN
	)
	;

xmlattributesFunction
	: (
	XMLATTRIBUTES LPAREN 
	expression AS NONNUMERICLITERAL (COMMA expression AS NONNUMERICLITERAL)*
	RPAREN
	)
	;

xmlserializeFunction
	: (
	XMLSERIALIZE LPAREN CONTENT? expression AS dataType
	xmlserializeFunctionOptions*
	RPAREN
	)
	;

xmlnamespaceFunction
	: (
	XMLNAMESPACES LPAREN xmlnamespaceOption (COMMA xmlnamespaceOption)* RPAREN
	)
	;

xmlnamespaceOption
	: (
	(namespaceUri AS namespacePrefix)
	| (DEFAULT namespaceUri)
	| (NO DEFAULT)
	)
	;

xmlserializeFunctionOptions
	: (
	(VERSION NONNUMERICLITERAL)
	| ((EXCLUDING | INCLUDING) XMLDECLARATION)
	)
	;

xmlFunctionOptionClause
	: (
	OPTION xmlFunctionOption+
	)
	;

xmlFunctionOption
	: (
	((EMPTY | NULL) ON NULL)
	| (XMLBINARY USING? (BASE64 | HEX))
	)
	;

elementContentExpression
	: (expression)
	;

xqueryExpressionConstant
	: (NONNUMERICLITERAL)
	;

xqueryArgument
	: (xqueryContextItemExpression | (xqueryVariableExpression AS (identifier | NONNUMERICLITERAL)))
	;

xmltableFunctionSpecification
	: (
	XMLTABLE
	LPAREN
	(xmlnamespacesDeclaration COMMA)?
	rowXqueryExpressionConstant
	(PASSING (BY REF)? rowXqueryArgument (COMMA rowXqueryArgument)*)?
	(COLUMNS (xmlTableRegularColumnDefinition | xmlTableOrdinalityColumnDefinition)
		(COMMA (xmlTableRegularColumnDefinition | xmlTableOrdinalityColumnDefinition))*)?
	RPAREN
	)
	;

rowXqueryExpressionConstant
	: (NONNUMERICLITERAL)
	;

rowXqueryArgument
	: (
	(xqueryContextItemExpression | (xqueryVariableExpression AS (identifier | NONNUMERICLITERAL)))
	)
	;

xqueryContextItemExpression
	: (expression)
	;

xqueryVariableExpression
	: (expression)
	;

xmlTableRegularColumnDefinition
	: (
	columnName
	dataType
	(defaultClause | (PATH columnXqueryExpressionConstant))?
	)
	;

/*
Made...

	(defaultClauseAllowables
	| (distinctTypeCastFunctionName LPAREN defaultClauseAllowables RPAREN))

...optional because it wasn't by my mistake.  Noted by Martijn Rutte 2023-01-09.
*/
defaultClause
	: (
	WITH? DEFAULT
	(defaultClauseAllowables
	| (distinctTypeCastFunctionName LPAREN defaultClauseAllowables RPAREN))?
	)
	;

defaultClause1
	: (
	WITH? DEFAULT defaultClauseAllowables?
	)
	;

defaultClauseAllowables
	: (
	literal
	| SESSION_USER
	| USER
	| CURRENT_SQLID
	| NULL
	)
	;

distinctTypeCastFunctionName
	: (identifier DOT identifier)
	;

/*
castFunction
	: (
	castSpecification
	| scalarFunctionInvocation
	| charFunctionSpecification
	| clobFunctionSpecification
	| dbclobFunctionSpecification
	| graphicFunctionSpecification1
	| graphicFunctionSpecification2
	| vargraphicFunctionSpecification1
	| vargraphicFunctionSpecification2
	)
	;

expressionAndCodeUnitsArguments
	: (
	expression 
	(COMMA INTEGERLITERAL (COMMA (CODEUNITS16 | CODEUNITS32 | OCTETS))?)?
	)
	;

clobFunctionSpecification
	: (
	CLOB
	LPAREN
	expressionAndCodeUnitsArguments
	RPAREN
	)
	;

dbclobFunctionSpecification
	: (
	DBCLOB
	LPAREN
	expressionAndCodeUnitsArguments
	RPAREN
	)
	;

graphicFunctionSpecification1
	: (
	GRAPHIC
	LPAREN
	expressionAndCodeUnitsArguments
	RPAREN
	)
	;

graphicFunctionSpecification2
	: (
	GRAPHIC
	LPAREN
	expression 
	NONNUMERICLITERAL?
	RPAREN
	)
	;

vargraphicFunctionSpecification1
	: (
	GRAPHIC
	LPAREN
	expressionAndCodeUnitsArguments
	RPAREN
	)
	;

vargraphicFunctionSpecification2
	: (
	GRAPHIC
	LPAREN
	expression 
	NONNUMERICLITERAL?
	RPAREN
	)
	;
*/

columnXqueryExpressionConstant
	: (NONNUMERICLITERAL)
	;

xmlTableOrdinalityColumnDefinition
	: (columnName FOR ORDINALITY)
	;

xmlnamespacesDeclaration
	: (
	xmlnamespacesFunctionSpecification
	(COMMA xmlnamespacesFunctionSpecification)*
	)
	;

xmlnamespacesFunctionSpecification
	: (
	XMLNAMESPACES
	LPAREN
	xmlnamespacesFunctionArguments
	(COMMA xmlnamespacesFunctionArguments)*
	RPAREN
	)
	;

xmlnamespacesFunctionArguments
	: (
	((namespaceUri AS namespacePrefix)
	| (DEFAULT namespaceUri)
	| (NO DEFAULT))
	)
	;

namespaceUri
	: NONNUMERICLITERAL
	;

namespacePrefix
	: NONNUMERICLITERAL
	;

timeZoneSpecificExpression
	: timeZoneExpressionSubset
	((AT LOCAL) | (AT TIME ZONE timeZoneExpressionSubset))
	;

timeZoneExpressionSubset
	: (
	functionInvocation
	| literal
	| columnName
	| hostVariable
	| specialRegister
	| scalarFullSelect
	| caseExpression
	| castSpecification
	)
	;

caseExpression
	: CASE
	(searchedWhenClause+ | simpleWhenClause)
	((ELSE NULL) | (ELSE resultExpression))?
	END
	;

resultExpression
	: expression
	;

searchedWhenClause
	: (
	WHEN searchCondition THEN (resultExpression | NULL)
	)
	;

simpleWhenClause
	: (
	expression
	(WHEN expression THEN (resultExpression | NULL))+
	)
	;

searchCondition
	: NOT?
	((predicate (SELECTIVITY NUMERICLITERAL)?) | (LPAREN searchCondition RPAREN))
	((AND | OR) NOT? (predicate | (LPAREN searchCondition RPAREN)))*
	;

checkCondition
	: (searchCondition)
	;

predicate
	: basicPredicate
	| quantifiedPredicate
	| arrayExistsPredicate
	| betweenPredicate
	| distinctPredicate
	| existsPredicate
	| inPredicate
	| likePredicate
	| nullPredicate
	| xmlExistsPredicate
	;

basicPredicate
	: ((expression comparisonOperator expression)
	| (rowValueExpression comparisonOperator rowValueExpression))
	;

rowValueExpression
	: LPAREN expression
	(COMMA expression)*
	RPAREN
	;

quantifiedPredicate
	: ((expression comparisonOperator (SOME | ANY | ALL) LPAREN fullSelect RPAREN)
	| (rowValueExpression EQ (SOME | ANY) LPAREN fullSelect RPAREN)
	| (rowValueExpression NE ALL LPAREN fullSelect RPAREN))
	;

arrayExistsPredicate
	: ARRAY_EXISTS
	LPAREN
	arrayExpression
	INTEGERLITERAL
	RPAREN
	;

betweenPredicate
	: expression NOT? BETWEEN expression AND expression
	;

distinctPredicate
	: expression IS NOT? DISTINCT FROM expression
	;

existsPredicate
	: EXISTS LPAREN fullSelect RPAREN
	;

inPredicate
	: expression NOT? IN (
	(LPAREN fullSelect RPAREN)
	| (LPAREN expression (COMMA expression)* RPAREN)
	)
	;

likePredicate
	: expression NOT? LIKE expression (ESCAPE expression)?
	;

nullPredicate
	: expression ((IS NOT? NULL) | ISNULL | NOTNULL)
	;

xmlExistsPredicate
	: XMLEXISTS
	LPAREN
	NONNUMERICLITERAL
	(PASSING (BY REF)? expression (COMMA expression)*)?
	RPAREN
	;

arrayExpression
	: variable
	| castSpecification
	;

castSpecification
	: CAST
	LPAREN
	(expression | NULL | parameterMarker)
	AS
	castDataType
	RPAREN
	;

parameterMarker
	: QUESTIONMARK
	;

castDataType
	: (
	castBuiltInType
	| distinctTypeName
	| arrayType
	)
	;

castBuiltInType
	: (
	SMALLINT
	| INTEGER
	| INT
	| BIGINT
	| ((DECIMAL | DEC | NUMERIC) (integerInParens | (LPAREN RPAREN)))
	| (DECFLOAT (integerInParens | (LPAREN RPAREN)))
	| (FLOAT (integerInParens | (LPAREN RPAREN)))
	| REAL
	| (DOUBLE PRECISION?)
	| ((((CHARACTER | CHAR) VARYING? ) | VARCHAR) (length | (LPAREN RPAREN))? ccsidQualifier?)
	| ((((CHARACTER | CHAR) LARGE OBJECT) | CLOB) (length | (LPAREN RPAREN))? ccsidQualifier?)
	| ((GRAPHIC | VARGRAPHIC | DBCLOB) (length | (LPAREN RPAREN))? ccsidQualifier?)
	| (BINARY (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY VARYING?) | VARBINARY) (integerInParens | (LPAREN RPAREN))?)
	| (((BINARY LARGE OBJECT) | BLOB) length?)
	| DATE
	| TIME
	| (TIMESTAMP integerInParens? ((WITH | WITHOUT) TIME ZONE)?)
	| ROWID
	| XML
	)
	;

integerInParens
	: (LPAREN INTEGERLITERAL (COMMA INTEGERLITERAL)? RPAREN)
	;

/*
It turns out the lengthQualifier of K or M or G gets lexed
as being part of the INTEGERLITERAL and becomes an SQLIDENTIFIER.

It also turns out that...

1024
1K
1 K

...are all syntactically correct.  So we can see an INTEGERLITERAL,
an INTEGERLITERAL followed by an sqlidentifier, or an sqlidentifier.
*/
length
	: (
	LPAREN
	(INTEGERLITERAL | (INTEGERLITERAL sqlidentifier) | sqlidentifier)
	(CODEUNITS16 | CODEUNITS32 | OCTETS)?
	RPAREN
	)
	;

ccsidQualifier
	: (
	CCSID
	(((ASCII | EBCDIC | UNICODE) forDataQualifier?) | INTEGERLITERAL)
	)
	;

forDataQualifier
	: (FOR (SBCS | MIXED | BIT) DATA)
	;

distinctTypeName
	: (schemaName DOT)? identifier
	;

arrayType
	: identifier
	;

literal
	: NUMERICLITERAL
	| NONNUMERICLITERAL
	| INTEGERLITERAL
	;

numeric
    : NUMERICLITERAL
    | INTEGERLITERAL
    ;

columnName
	: (((correlationName | tableName) DOT)? identifier1)
	;

sourceColumnName
	: columnName
	;

targetColumnName
	: columnName
	;

buildColumnName
	: identifier
	;

beginColumnName
	: identifier
	;

endColumnName
	: identifier
	;

correlationName
	: identifier
	;

locationName
	: sqlidentifier
	;

schemaName
	: identifier1
	;

tableName
	: (((locationName DOT schemaName DOT) | (schemaName DOT))? identifier)
	;

auxTableName
	: (((locationName DOT schemaName DOT) | (schemaName DOT))? identifier)
	;

viewName
	: ((locationName DOT schemaName DOT) | (schemaName DOT))? identifier correlationName?
	;

programName
	: identifier
	;

packageName
	: identifier
	;

planName
	: identifier
	;

typeName
	: ((schemaName DOT)? identifier)
	;

variableName
	: ((schemaName DOT)? identifier)
	;

arrayTypeName
	: ((schemaName DOT)? identifier)
	;

jarName
	: ((schemaName DOT)? identifier)
	;

savepointName
	: identifier
	;

aliasName
	: identifier
	;

constraintName
	: identifier
	;

uniqueID
	: identifier
	;

versionID
	: (identifier | NUMERICLITERAL | INTEGERLITERAL) (DOT? (identifier | NUMERICLITERAL | INTEGERLITERAL))*
	;

indexName
	: (schemaName DOT)? identifier
	;

maskName
	: (schemaName DOT)? identifier
	;

permissionName
	: (schemaName DOT)? identifier
	;

procedureName
	: ((locationName DOT schemaName DOT) | (schemaName DOT))? identifier
	;

sequenceName
	: (schemaName DOT)? identifier
	;

memberName
	: identifier
	;

databaseName
	: identifier
	;

tablespaceName
	: identifier
	;

acceleratorName
	: identifier
	;

catalogName
	: identifier
	;

triggerName
	: identifier
	;

contextName
	: identifier
	;

authorizationName
	: identifier
	;

profileName
	: identifier
	;

roleName
	: identifier
	;

seclabelName
	: identifier
	;

parameterName
	: identifier
	;

addressValue
	: NONNUMERICLITERAL
	;

jobnameValue
	: NONNUMERICLITERAL
	;

servauthValue
	: NONNUMERICLITERAL
	;

encryptionValue
	: NONNUMERICLITERAL
	;

bpName
	: identifier
	;

stogroupName
	: identifier
	;

keyLabelName
	: (identifier | NONNUMERICLITERAL)
	;

functionName
	: (schemaName DOT)? identifier
	;

specificName
	: (schemaName DOT)? identifier
	;

hostLabel
	: (INTEGERLITERAL | identifier) (INTEGERLITERAL | (MINUS identifier))*
	;

hostVariable
	: COLON (hostStructure DOT)? hostIdentifier (INDICATOR? COLON (nullIndicatorStructure DOT)? nullIndicator)?
	;

hostIdentifier
	: (INTEGERLITERAL | identifier) (INTEGERLITERAL | (MINUS identifier))*
	;

hostStructure
	: (INTEGERLITERAL | identifier) (INTEGERLITERAL | (MINUS identifier))*
	;

nullIndicator
	: (INTEGERLITERAL | identifier) (INTEGERLITERAL | (MINUS identifier))*
	;

nullIndicatorStructure
	: (INTEGERLITERAL | identifier) (INTEGERLITERAL | (MINUS identifier))*
	;

globalVariableName
	: ((schemaName DOT)? identifier1)
	;

sqlParameterName
	: ((schemaName DOT)? identifier1)
	;

sqlVariableName
	: ((schemaName DOT)? identifier1)
	;

transitionVariableName
	: columnName
	;

/*
Trigger variables, global variables, SQL variables, all
these conform to the pattern (schemaName DOT)? identifier.
*/
variable
	: ((schemaName DOT)? identifier)
	| hostVariable
	;

intoClause
	: INTO
	(variable | arrayElementSpecification)
	(COMMA variable)*
	;

correlationClause
	: AS?
	correlationName
	(LPAREN
	buildColumnName
	(COMMA buildColumnName)*
	RPAREN)?
	;

/*	
fromClause
	: FROM
	tableName correlationClause?
	(COMMA tableName correlationClause?)*
	;
*/

fromClause
	: (
	FROM
	((LPAREN* tableReference RPAREN*) | collectionDerivedTable)
	(COMMA ((LPAREN* tableReference RPAREN*) | collectionDerivedTable))*
	)
	;

tableReference
	: (
	singleTableReference
	| nestedTableExpression
	| dataChangeTableReference
	| tableFunctionReference
	| tableLocatorReference
	| xmltableExpression
	| collectionDerivedTable
/*
The following is brought to you by the ANTLR 4.9.2 message
"The following sets of rules are mutually left-recursive [tableReference, joinedTable]"
*/

	| ((singleTableReference 
		| nestedTableExpression 
		| tableFunctionReference 
		| tableLocatorReference 
		| xmltableExpression 
		| collectionDerivedTable 
		| (LPAREN+ tableReference RPAREN+)
		| ((singleTableReference 
			| nestedTableExpression 
			| tableFunctionReference 
			| tableLocatorReference 
			| xmltableExpression 
			| (LPAREN+ tableReference RPAREN+)
			| collectionDerivedTable)
				(INNER | ((LEFT | RIGHT | FULL) OUTER?))? JOIN
						tableReference ON joinCondition))
		((INNER | ((LEFT | RIGHT | FULL) OUTER?))? JOIN
					tableReference ON joinCondition)+)
	| ((singleTableReference 
		| nestedTableExpression 
		| tableFunctionReference 
		| tableLocatorReference 
		| xmltableExpression 
		| collectionDerivedTable 
		| (LPAREN+ tableReference RPAREN+)
		| ((singleTableReference 
			| nestedTableExpression 
			| tableFunctionReference 
			| tableLocatorReference 
			| xmltableExpression 
			| (LPAREN+ tableReference RPAREN+)
			| collectionDerivedTable)
				(INNER | ((LEFT | RIGHT | FULL) OUTER?))? JOIN
					tableReference ON joinCondition)) CROSS JOIN tableReference)
	| (LPAREN+ tableReference RPAREN+)
	)
	;

singleTableReference
	: (tableName AS? correlationName? periodSpecification* correlationClause?)
	;

periodSpecification
	: (
	FOR (SYSTEM_TIME | BUSINESS_TIME) 
	((AS OF expression) | (FROM expression TO expression) | (BETWEEN expression AND expression))
	)
	;

periodClause
	: (
	FOR PORTION OF BUSINESS_TIME 
	((FROM expression TO expression) | (BETWEEN expression AND expression))
	)
	;

nestedTableExpression
	: (TABLE? LPAREN fullSelect RPAREN correlationClause?)
	;

/**/
dataChangeTableReference
	: (
	(FINAL TABLE LPAREN insertStatement RPAREN correlationClause?)
	| ((FINAL | OLD) TABLE searchedUpdate)
	| (OLD TABLE searchedDelete)
	| (FINAL TABLE mergeStatement)
	)
	;

/**/

tableFunctionReference
	: (
	TABLE LPAREN 
	(schemaName DOT)?
	(scalarFunction | aggregateFunction | regressionFunction | tableFunction | identifier)
	LPAREN
	((expression | (TABLE tableName)) (COMMA (expression | (TABLE tableName)))*)?
	RPAREN
	tableUdfCardinalityClause?
	RPAREN
	(correlationClause | typedCorrelationClause)?
	)
	;

tableUdfCardinalityClause
	: (
	CARDINALITY MULTIPLIER? (INTEGERLITERAL | NUMERICLITERAL)
	)
	;

typedCorrelationClause
	: (
	AS? correlationName LPAREN columnName dataType (COMMA columnName dataType)* RPAREN
	)
	;

tableLocatorReference
	: (
	TABLE
	LPAREN
	identifier
	LIKE
	tableName
	RPAREN
	correlationName?
	)
	;

xmltableExpression
	: (xmltableFunctionSpecification correlationClause?)
	;

/*
correlationClause
	: (AS? correlationName (LPAREN columnName (COMMA columnName)* RPAREN)?)
	;
*/

collectionDerivedTable
	: (
	UNNEST
	LPAREN
	(ordinaryArrayExpression | associativeArrayExpression)
	(COMMA (ordinaryArrayExpression | associativeArrayExpression))*
	RPAREN
	(WITH ORDINALITY)?
	correlationClause?
	)
	;

/* moved to the interior of tableReference due to mutual left-recursion error
joinedTable
	: (
	(tableReference
	(INNER | ((LEFT | RIGHT | FULL) OUTER?))
	JOIN
	tableReference ON joinCondition)
	| (tableReference CROSS JOIN tableReference)
	| (LPAREN joinedTable RPAREN)
	)
	;
*/

joinCondition
	: (
	searchCondition
	| (fullJoinExpression EQ fullJoinExpression)
	)
	;

fullJoinExpression
	: (
	(columnName
	| castFunction
	| (COALESCE LPAREN (columnName | castFunction) (COMMA (columnName | castFunction))* RPAREN))
	)
	;

castFunction
	: (castSpecification)
	;

ordinaryArrayExpression
	: (expression)
	;

associativeArrayExpression
	: (expression)
	;

whereClause
	: WHERE searchCondition
	;

groupByClause
	: GROUP BY
	(groupingExpression | groupingSets | superGroups)
	;

havingClause
	: HAVING searchCondition
	;

groupingExpression
	: (expression (COMMA expression)*)
	;

groupingSets
	: GROUPING SETS groupingSetsGroup
	;

groupingSetsGroup
	: LPAREN 
	(groupingSetsGroup | groupingExpression | superGroups) 
	(COMMA (groupingSetsGroup | groupingExpression | superGroups))* 
	RPAREN
	;

superGroups
	: (
	((ROLLUP | CUBE) LPAREN groupingExpression RPAREN)
	| (LPAREN RPAREN)
	)
	;

selectColumns
	: (
	(expression ((operator expression) | INTEGERLITERAL)* (AS? (buildColumnName | NONNUMERICLITERAL))?)
	| (tableName DOT SPLAT)
	| unpackedRow
	)
	;

unpackedRow
	: (
	UNPACK LPAREN expression RPAREN DOT SPLAT AS 
	LPAREN 
	columnName dataType (COMMA columnName dataType)* 
	RPAREN
	)
	;

selectClause
	: (
	SELECT
	(ALL | DISTINCT)?
	(SPLAT | (selectColumns (COMMA selectColumns)*))
	)
	;

subSelect
	: (
	selectClause
	fromClause
	whereClause?
	groupByClause?
	havingClause?
	orderByClause?
	offsetClause?
	fetchClause?
	)
	;

selectIntoStatement
	: (
	(WITH commonTableExpression (COMMA commonTableExpression)*)?
	selectClause
	intoClause
	fromClause
	whereClause?
	groupByClause?
	havingClause?
	orderByClause?
	offsetClause?
	fetchClause?
	(isolationClause | skipLockedDataClause)?
	querynoClause?
	)
	;

selectStatement
	: (
	(WITH commonTableExpression (COMMA commonTableExpression)*)?
	fullSelect
	(
	updateClause
	| readOnlyClause
	| optimizeClause
	| isolationClause 
	| skipLockedDataClause
	| querynoClause
	)*
	)
	;

/*
Made...

	(LPAREN columnName (COMMA columnName)* RPAREN)

...optional, as it should be per the documentation.  Noted
by Martijn Rutte 2023-01-10.
*/
commonTableExpression
	: tableName 
	(LPAREN columnName (COMMA columnName)* RPAREN)?
	AS LPAREN fullSelect RPAREN
	;

updateClause
	: (FOR UPDATE (OF columnName (COMMA columnName)*)?)
	;

readOnlyClause
	: (FOR READ ONLY)
	;

optimizeClause
	: OPTIMIZE FOR INTEGERLITERAL (ROW | ROWS)
	;

isolationClause
	: WITH 
	(
	(RR lockClause?)
	| (RS lockClause?)
	| CS
	| UR
	)
	;

lockClause
	: (USE AND KEEP (EXCLUSIVE | UPDATE | SHARE) LOCKS)
	;

skipLockedDataClause
	: (SKIP_ LOCKED DATA)
	;

querynoClause
	: (QUERYNO INTEGERLITERAL)
	;

scalarFullSelect
	: LPAREN
	fullSelect
	RPAREN
	;

fullSelect
	: ((LPAREN fullSelect RPAREN) | subSelect | valuesClause)
	((UNION | EXCEPT | INTERSECT) (DISTINCT | ALL)? (subSelect | (LPAREN fullSelect RPAREN)))*
	orderByClause?
	offsetClause?
	fetchClause?
	;

valuesClause
	: VALUES
	(sequenceReference
	| (LPAREN sequenceReference (COMMA sequenceReference)* RPAREN))
	;

orderByClause
	: ORDER BY 
	(
	(sortKey (ASC | DESC)? (COMMA sortKey (ASC | DESC)?)*)
	| (INPUT SEQUENCE)
	| (ORDER OF tableName)
	)
	;

sortKey
	: (columnName | INTEGERLITERAL | expression)
	;

offsetClause
	: OFFSET INTEGERLITERAL (ROW | ROWS)
	;

fetchClause
	: FETCH (FIRST | NEXT) INTEGERLITERAL? (ROW | ROWS) ONLY
	;

identifier
	: sqlidentifier
	| sqlKeyword
	| specialRegister
	| scalarFunction
	| aggregateFunction
	| regressionFunction
	| tableFunction
	;

identifier1
	: sqlidentifier
	| sqlKeyword
	| scalarFunction
	| aggregateFunction
	| regressionFunction
	| tableFunction
	;

sqlidentifier
	: (SQLIDENTIFIER | NONNUMERICLITERAL)
	;
	
sqlKeyword
	: (
	  ABSOLUTE
	| ACCELERATION
	| ACCELERATOR
	| ACCESS
	| ACCESSCTRL
	| ACTION
	| ACTIVATE
	| ACTIVE
	| ADD
	| ADDRESS
	| AFTER
	| ALGORITHM
	| ALIAS
	| ALL
	| ALLOCATE
	| ALLOW
	| ALTER
	| ALTERAND
	| ALTERIN
	| ALWAYS
	| AND
	| ANY
	| APPEND
	| APPLCOMPAT
	| APPLICATION
	| ARCHIVE
	| ARRAY
	| ARRAY_AGG
	| ARRAY_EXISTS
	| AS
	| ASC
	| ASCII
	| ASENSITIVE
	| ASSEMBLE
	| ASSOCIATE
	| ASUTIME
	| AT
	| ATOMIC
	| ATTRIBUTES
	| AUDIT
	| AUTHENTICATION
	| AUTHID
	| AUTONOMOUS
	| AUX
	| AUXILIARY
	| AVG
	| BASE64
	| BASED
	| BEFORE
	| BEGIN
	| BETWEEN
	| BIGINT
	| BIN
	| BINARY
	| BIND
	| BINDADD
	| BINDAGENT
	| BIT
	| BLOB
	| BOTH
	| BSDS
	| BUFFERPOOL
	| BUFFERPOOLS
	| BUSINESS
	| BUSINESS_TIME
	| BY
	| C_
	| CACHE
	| CALL
	| CALLED
	| CALLER
	| CAPTURE
	| CASCADE
	| CASCADED
	| CASE
	| CAST
	| CATALOG_NAME
	| CCSID
	| CHANGE
	| CHANGED
	| CHANGES
	| CHAR
	| CHARACTER
	| CHECK
	| CHECKING
	| CLASS
	| CLAUSE
	| CLIENT
	| CLOB
	| CLONE
	| CLOSE
	| CLUSTER
	| COBOL
	| CODEUNITS16
	| CODEUNITS32
	| COLLECTION
	| COLLID
	| COLUMN
	| COLUMNS
	| COMMENT
	| COMMIT
	| COMMITTED
	| COMPATIBILITY
	| COMPONENT
	| COMPRESS
	| CONCAT
	| CONCENTRATE
	| CONCURRENT
	| CONDITION
	| CONDITION_NUMBER
	| CONNECT
	| CONNECTION
	| CONSTANT
	| CONSTRAINT
	| CONTAINS
	| CONTENT
	| CONTEXT
	| CONTINUE
	| CONTROL
	| COPY
	| CORR
	| CORRELATION
	| COVAR
	| COVARIANCE
	| COVARIANCE_SAMP
	| COVAR_POP
	| COVAR_SAMP
	| CREATE
	| CREATEALIAS
	| CREATEDBA
	| CREATEDBC
	| CREATEIN
	| CREATE_SECURE_OBJECT
	| CREATESG
	| CREATETAB
	| CREATETMTAB
	| CREATETS
	| CROSS
	| CS
	| CUBE
	| CUME_DIST
	| CURRENT
	| CURRENT_DATE
	| CURRENTLY
	| CURRENT_PATH
	| CURRENT_SCHEMA
	| CURRENT_SERVER
	| CURRENT_TIME
	| CURRENT_TIMESTAMP
	| CURRENT_TIME_ZONE
	| CURRVAL
	| CURSOR
	| CURSOR_NAME
	| CURSORS
	| CYCLE
	| DATA
	| DATAACCESS
	| DATABASE
	| DATACLAS
	| DATE
	| DAY
	| DAYS
	| DB2
	| DB2_AUTHENTICATION_TYPE
	| DB2_AUTHORIZATION_ID
	| DB2_CONNECTION_STATE
	| DB2_CONNECTION_STATUS
	| DB2_ENCRYPTION_TYPE
	| DB2_ERROR_CODE1
	| DB2_ERROR_CODE2
	| DB2_ERROR_CODE3
	| DB2_ERROR_CODE4
	| DB2_GET_DIAGNOSTICS_DIAGNOSTICS
	| DB2_INTERNAL_ERROR_POINTER
	| DB2_LAST_ROW
	| DB2_LINE_NUMBER
	| DB2_MESSAGE_ID
	| DB2_MODULE_DETECTING_ERROR
	| DB2_NUMBER_PARAMETER_MARKERS
	| DB2_NUMBER_RESULT_SETS
	| DB2_NUMBER_ROWS
	| DB2_ORDINAL_TOKEN_
	| DB2_PRODUCT_ID
	| DB2_REASON_CODE
	| DB2_RETURNED_SQLCODE
	| DB2_RETURN_STATUS
	| DB2_ROW_NUMBER
	| DB2_SERVER_CLASS_NAME
	| DB2SQL
	| DB2_SQL_ATTR_CURSOR_HOLD
	| DB2_SQL_ATTR_CURSOR_ROWSET
	| DB2_SQL_ATTR_CURSOR_SCROLLABLE
	| DB2_SQL_ATTR_CURSOR_SENSITIVITY
	| DB2_SQL_ATTR_CURSOR_TYPE
	| DB2_SQLERRD1
	| DB2_SQLERRD2
	| DB2_SQLERRD3
	| DB2_SQLERRD4
	| DB2_SQLERRD5
	| DB2_SQLERRD6
	| DB2_SQLERRD_SET
	| DB2_SQL_NESTING_LEVEL
	| DB2_TOKEN_COUNT
	| DBADM
	| DBCLOB
	| DBCTRL
	| DBINFO
	| DBMAINT
	| DBPARTITIONNUM
	| DBPARTITIONNUMS
	| DEACTIVATE
	| DEALLOCATE
	| DEBUG
	| DEBUGSESSION
	| DEC
	| DECFLOAT
	| DECIMAL
	| DECLARE
	| DEC_ROUND_CEILING
	| DEC_ROUND_DOWN
	| DEC_ROUND_FLOOR
	| DEC_ROUND_HALF_DOWN
	| DEC_ROUND_HALF_EVEN
	| DEC_ROUND_HALF_UP
	| DEC_ROUND_UP
	| DEFAULT
	| DEFAULTS
	| DEFER
	| DEFERRED
	| DEFINE
	| DEFINEBIND
	| DEFINER
	| DEFINERUN
	| DEGREE
	| DELETE
	| DENSE_RANK
	| DEPENDENT
	| DESC
	| DESCRIBE
	| DESCRIPTOR
	| DETERMINISTIC
	| DEVICE
	| DIAGNOSTICS
	| DISABLE
	| DISALLOW
	| DISCONNECT
	| DISPATCH
	| DISPLAY
	| DISPLAYDB
	| DISTINCT
	| DISTRIBUTE
	| DO
	| DOCUMENT
	| DOUBLE
	| DROP
	| DROPIN
	| DSSIZE
	| DYNAMIC
	| DYNAMICRULES
	| EACH
	| EBCDIC
	| EDITPROC
	| ELEMENT
	| ELIGIBLE
	| ELSE
	| ELSEIF
	| EMPTY
	| ENABLE
	| ENCODING
	| ENCRYPTION
	| END
	| END_EXEC
	| ENDING
	| ENFORCED
	| EVENT
	| ENVIRONMENT
	| ERASE
	| ESCAPE
	| EUR
	| EVERY
	| EXCEPT
	| EXCEPTION
	| EXCHANGE
	| EXCLUDE
	| EXCLUDING
	| EXCLUSIVE
	| EXEC_SQL
	| EXECUTE
	| EXISTS
	| EXIT
	| EXPLAIN
	| EXTERNAL
	| EXTRA
	| FAILBACK
	| FAILURE
	| FAILURES
	| FENCED
	| FETCH
	| FIELDPROC
	| FINAL
	| FIRST
	| FIRST_VALUE
	| FIXEDLENGTH
	| FLOAT
	| FOLLOWING
	| FOR
	| FOREIGN
	| FORMAT
	| FOUND
	| FREE
	| FREEPAGE
	| FROM
	| FULL
	| FUNCTION
	| GBPCACHE
	| GENERAL
	| GENERATE
	| GENERATED
	| GENERIC
	| GET
	| GET_ACCEL_ARCHIVE
	| GLOBAL
	| GO
	| GOTO
	| GRANT
	| GRAPHIC
	| GROUP
	| GROUPING
	| HANDLER
	| HAVING
	| HIGH
	| HISTOGRAM
	| HIDDEN_
	| HINT
	| HISTORY
	| HOLD
	| HOUR
	| HOURS
	| HUFFMAN
	| ID
	| IDENTITY
	| IF
	| IGNORE
	| IMAGCOPY
	| IMMEDIATE
	| IMPLICITLY
	| IN
	| INCLUDE
	| INCLUDING
	| INCLUSIVE
	| INCREMENT
	| INDEX
	| INDEXBP
	| INDICATOR
	| INHERIT
	| INITIALLY
	| INLINE
	| INNER
	| INOUT
	| INPUT
	| INSENSITIVE
	| INSERT
	| INSTEAD
	| INT
	| INTEGER
	| INTERSECT
	| INTO
	| INVOKEBIND
	| INVOKERUN
	| IS
	| ISNULL
	| ISO
	| ISOBID
	| ISOLATION
	| ITERATE
	| JAR
	| JAVA
	| JIS
	| JOBNAME
	| JOIN
	| KEEP
	| KEY
	| KEYS
	| LABEL
	| LABELS
	| LAG
	| LANGUAGE
	| LARGE
	| LAST
	| LAST_VALUE
	| LC_CTYPE
	| LEAD
	| LEAVE
	| LEFT
	| LEVEL
	| LIKE
	| LIMIT
	| LISTAGG
	| LITERALS
	| LOAD
	| LOB
	| LOCAL
	| LOCALE
	| LOCATION
	| LOCATOR
	| LOCATORS
	| LOCK
	| LOCKED
	| LOCKMAX
	| LOCKS
	| LOCKSIZE
	| LOG
	| LOGGED
	| LOGICAL
	| LONG
	| LOOP
	| MAIN
	| MAINTAINED
	| MAPPING
	| MASK
	| MATCHED
	| MATERIALIZED
	| MAXPARTITIONS
	| MAXROWS
	| MAXVALUE
	| MEDIAN
	| MEMBER
	| MERGE
	| MESSAGE_TEXT
	| MGMTCLAS
	| MICROSECOND
	| MICROSECONDS
	| MINUTE
	| MINUTES
	| MINVALUE
	| MIXED
	| MODE_
	| MODEL
	| MODIFIERS
	| MODIFIES
	| MODULE
	| MONITOR
	| MONITOR1
	| MONITOR2
	| MONTH
	| MONTHS
	| MORE_
	| MOVE
	| MOVEMENT
	| MULTIPLIER
	| NAME
	| NAMES
	| NAMESPACE
	| NEW
	| NEW_TABLE
	| NEXT
	| NEXTVAL
	| NICKNAME
	| NO
	| NODEFER
	| NONE
	| NOT
	| NOTNULL
	| NTH_VALUE
	| NTILE
	| NULL
	| NULLS
	| NULTERM
	| NUMBER
	| NUMERIC
	| NUMPARTS
	| OBID
	| OBJECT
	| OBJMAINT
	| OCTETS
	| OF
	| OFFSET
	| OLD
	| OLD_TABLE
	| ON
	| ONCE
	| ONLY
	| OPEN
	| OPERATION
	| OPTHINT
	| OPTIMIZATION
	| OPTIMIZE
	| OPTION
	| OPTIONAL
	| OPTIONS
	| OR
	| ORDER
	| ORDINALITY
	| ORGANIZATION
	| ORGANIZE
	| OUT
	| OUTCOME
	| OUTER
	| OUTPUT
	| OVER
	| OVERHEAD
	| OVERLAPS
	| OVERRIDING
	| OWNER
	| OWNERSHIP
	| PACKADM
	| PACKAGE
	| PACKAGE_NAME
	| PACKAGE_SCHEMA
	| PACKAGE_VERSION
	| PADDED
	| PAGE
	| PAGENUM
	| PARALLEL
	| PARAMETER
	| PART
	| PARTITION
	| PARTITIONED
	| PARTITIONING
	| PASSING
	| PASSWORD
	| PATH
	| PCTDEACTIVATE
	| PCTFREE
	| PENDING
	| PERCENTILE_CONT
	| PERCENTILE_DISC
	| PERCENT_RANK
	| PERIOD
	| PERMISSION
	| PIECESIZE
	| PLAN
	| PLI
	| PORTION
	| POSITIONING
	| PRECEDING
	| PRECISION
	| PREPARE
	| PRESERVE
	| PREVIOUS
	| PREVVAL
	| PRIMARY
	| PRIOR
	| PRIQTY
	| PRIVILEGES
	| PROCEDURE
	| PROFILE
	| PROGRAM
	| PSID
	| PUBLIC
	| QUALIFIER
	| QUERY
	| QUERYNO
	| QUERYTAG
	| RANGE
	| RANK
	| RATE
	| RATIO_TO_REPORT
	| READ
	| READS
	| REAL
	| RECOVER
	| RECOVERDB
	| REF
	| REFERENCE
	| REFERENCES
	| REFERENCING
	| REFRESH
	| REGENERATE
	| REGISTERS
	| REGR_AVGX
	| REGR_AVGY
	| REGR_COUNT
	| REGR_ICPT
	| REGR_INTERCEPT
	| REGR_R2
	| REGR_SLOPE
	| REGR_SXX
	| REGR_SXY
	| REGR_SYY
	| RELATIVE
	| RELEASE
	| REMOVE
	| RENAME
	| REOPT
	| REORG
	| REPAIR
	| REPEAT
	| REQUIRED
	| RESET
	| RESIDENT
	| RESIGNAL
	| RESOLUTION
	| RESPECT
	| RESTART
	| RESTRICT
	| RESULT
	| RESULT_SET_LOCATOR
	| RETAIN
	| RETURN
	| RETURNED_SQLSTATE
	| RETURNING
	| RETURNS
	| REUSE
	| REVOKE
	| REXX
	| RIGHT
	| ROLE
	| ROLLBACK
	| ROLLUP
	| ROTATE
	| ROUND_CEILING
	| ROUND_DOWN
	| ROUND_FLOOR
	| ROUND_HALF_DOWN
	| ROUND_HALF_EVEN
	| ROUND_HALF_UP
	| ROUNDING
	| ROUND_UP
	| ROW
	| ROW_COUNT
	| ROWID
	| ROW_NUMBER
	| ROWS
	| ROWSET
	| RR
	| RS
	| RUN
	| SAVEPOINT
	| SBCS
	| SCHEMA
	| SCHEME
	| SCOPE
	| SCRATCHPAD
	| SCROLL
	| SECOND
	| SECONDS
	| SECQTY
	| SECTION
	| SECURED
	| SECURITY
	| SEGSIZE
	| SELECT
	| SELECTIVITY
	| SECMAINT
	| SENSITIVE
	| SEQUENCE
	| SERVAUTH
	| SERVER
	| SERVER_NAME
	| SESSION
	| SESSION_USER
	| SET
	| SETS
	| SHARE
	| SIGNAL
	| SIMPLE
	| SINGLE
	| SIZE
	| SNAPSHOT
	| SKIP_
	| SMALLINT
	| SOME
	| SOURCE
	| SPECIAL
	| SPECIFIC
	| SQL
	| SQLADM
	| SQLERROR
	| SQLEXCEPTION
	| SQLID
	| SQLSTATE
	| SQLWARNING
	| STABILIZED
	| STACKED
	| STANDARD
	| START
	| STARTDB
	| STARTING
	| STATEMENT
	| STATEMENTS
	| STATIC
	| STATS
	| STAY
	| STDDEV
	| STDDEV_POP
	| STDDEV_SAMP
	| STMTCACHE
	| STMTID
	| STMTTOKEN
	| STOGROUP
	| STOP
	| STOPALL
	| STOPDB
	| STORAGE
	| STORCLAS
	| STORES
	| STOSPACE
	| STRUCTURE
	| STYLE
	| SUB
	| SUM
	| SUMMARY
	| SYNONYM
	| SYSADM
	| SYSADMIN
	| SYSCTRL
	| SYSDATE
	| SYSDEFLT
	| SYSIBM
	| SYSOPR
	| SYSTEM
	| SYSTEM_TIME
	| SYSTIMESTAMP
	| SYSXSR
	| TABLE
	| TABLESPACE
	| TABLESPACES
	| TAG
	| TEMPLATE
	| TEMPORARY
	| THEN
	| THREADSAFE
	| THRESHOLD
	| TIME
	| TIMESTAMP
	| TIMEZONE
	| TO
	| TOKEN
	| TRACE
	| TRACKMOD
	| TRANSACTION
	| TRANSFER
	| TRANSFORM
	| TRANSFORMS
	| TRIGGER
	| TRIGGERS
	| TRUNCATE
	| TRUSTED
	| TYPE
	| UNBOUNDED
	| UNDO
	| UNICODE
	| UNION
	| UNIQUE
	| UNLOAD
	| UNNEST
	| UNPACK
	| UNTIL
	| UPDATE
	| UPON
	| UR
	| URL
	| USA
	| USAGE
	| USE
	| USER
	| USING
	| VALIDATE
	| VALIDPROC
	| VALUE
	| VALUES
	| VAR
	| VARBINARY
	| VARCHAR
	| VARGRAPHIC
	| VARIABLE
	| VARIANCE
	| VARIANCE_SAMP
	| VARIANT
	| VAR_POP
	| VAR_SAMP
	| VARYING
	| VCAT
	| VERSION
	| VERSIONING
	| VERSIONS
	| VIEW
	| VOLATILE
	| VOLUMES
	| WAIT
	| WAITFORDATA
	| WHEN
	| WHENEVER
	| WHERE
	| WHILE
	| WITH
	| WITHIN
	| WITHOUT
	| WLM
	| WORK
	| WORKFILE
	| WORKLOAD
	| WRAPPED
	| WRAPPER
	| WRITE
	| XML
	| XMLAGG
	| XMLBINARY
	| XMLCAST
	| XMLDECLARATION
	| XMLEXISTS
	| XMLNAMESPACES
	| XMLPATTERN
	| XMLSCHEMA
	| XSROBJECT
	| YEAR
	| YEARS
	| YES
	| ZONE
	| BUILD
	| MISSING
    | INDEXES
    | REQUIRE
    | MATCHING
    | DETACH
    | DISTRIBUTION
    | LONGVAR
    | BLOCKINSERT
    | ADAPTIVE
    | COMPRESSION
    | EXTENSION
    | HIERARCHY
    | LIST
    | COMPACT
    | CLASSIC
    | INPLACE
    | RECLAIM
    | EXTENTS
    | PAUSE
    | INDEXSCAN
    | RESETDICTIONARY
    | KEEPDICTIONARY
    | LONGLOBDATA
    | CLEANUP
    | OVERFLOWS
    | NOTRUNCATE
    | RESUME
    | REORGCHK
    | STATISTICS
	)
	;


