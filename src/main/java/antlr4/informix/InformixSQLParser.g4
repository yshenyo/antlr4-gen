/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2023 by 455741807@qq.com
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 * Project : InformixSQLParser; an ANTLR4 grammar for Informix database SQL language,
 * see as https://www.ibm.com/docs/en/informix-servers/14.10?topic=programming-guide-sql-syntax
 */

parser grammar InformixSQLParser;

options {
    tokenVocab = InformixSQLLexer;
}

sqlScript
    : unitStatement* EOF
    ;

unitStatement
    : (allocateCollection
    | allocateDescriptor
    | allocateRow
    | alterAccessMethod
    | alterFragment
    | alterFunction
    | alterIndex
    | alterProcedure
    | alterRoutine
    | alterSecurityLabelComponent
    | alterSequence
    | alterTable
    | alterTrustedContext
    | alterUser
    | beginWork
    | closeStatement
    | closeDatabase
    | commitWork
    | connectStatement
    | createAccessMethod
    | createAggregate
    | createCast
    | createDatabase
    | createDefaultUser
    | createDistinctType
    | createExternalTable
    | createFunction
    | createFunctionFrom
    | createIndex
    | createOpaqueType
    | createOpclass
    | createProcedure
    | createProcedureFrom
    | createRole
    | createRoutineFrom
    | createRowType
    | createSchema
    | createSecurityLabel
    | createSecurityLabelComponent
    | createSecurityPolicy
    | createSequence
    | createSynonym
    | createTable
    | createTempTable
    | createTrigger
    | createTrustedContext
    | createUser
    | createView
    | createXadatasource
    | createXadatasourceType
    | databaseStatement
    | deallocateCollection
    | deallocateDescriptor
    | deallocateRow
    | declareStatement
    | deleteStatement
    | describeStatement
    | describeInput
    | disconnectStatement
    | dropAccessMethod
    | dropAggregate
    | dropCast
    | dropDatabase
    | dropFunction
    | dropIndex
    | dropOpclass
    | dropProcedure
    | dropRole
    | dropRoutine
    | dropRowType
    | dropSecurity
    | dropSequence
    | dropSynonym
    | dropTable
    | dropTrigger
    | dropTrustedContext
    | dropType
    | dropUser
    | dropView
    | dropXadatasource
    | dropXadatasourceType
    | executeStatement
    | executeFunction
    | executeImmediate
    | executeProcedure
    | fetchStatement
    | flushStatement
    | freeStatement
    | getDescriptor
    | getDiagnostics
    | grantStatement
    | grantFragment
    | infoStatement
    | insertStatement
    | loadStatement
    | lockTable
    | mergeStatement
    | openStatement
    | outputStatement
    | prepareStatement
    | putStatement
    | releaseSavepoint
    | renameColumn
    | renameConstraint
    | renameDatabase
    | renameIndex
    | renameSecurity
    | renameSequence
    | renameTable
    | renameTrustedContext
    | renameUser
    | revokeStatement
    | revokeFragment
    | rollbackWork
    | saveExternalDirectives
    | savepointStatement
    | selectStatement
    | setAutofree
    | setCollation
    | setConnection
    | setConstraints
    | setDatabaseObjectMode
    | setDataskip
    | setDebugFile
    | setDeferredPrepare
    | setDescriptor
    | setEncryptionPassword
    | setEnvironment
    | setExplain
    | setIndexes
    | setIsolation
    | setLockMode
    | setLog
    | setOptimization
    | setPdqpriority
    | setRole
    | setSessionAuthorization
    | setStatementCache
    | setTransaction
    | setTransactionMode
    | setTriggers
    | setUserPassword
    | startViolationsTable
    | stopViolationsTable
    | truncateStatement
    | unloadStatement
    | unlockTable
    | updateStatement
    | updateStatistics
    | wheneverStatement
    ) SCOL
    ;

splStatement
    : (labelStatement
    | callStatement
    | caseStatement
    | continueStatement
    | defineStatement
    | exitStatement
    | forStatement
    | foreachStatement
    | gotoStatement
    | ifStatement
    | letStatement
    | loopStatement
    | onException
    | raiseException
    | returnStatement
    | systemStatement
    | traceStatement
    | whileStatement
    ) SCOL
    ;

labelStatement
    : LT2 identifier GT2
    ;

callStatement
    : CALL (procedureName OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR
    | functionName OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR RETURNING identifier (COMMA identifier)*
    | routineName (RETURNING identifier (COMMA identifier)*)?
    )
    ;

caseStatement
    : CASE (elseClause
    | (WHEN (constExprssion | NULL) THEN statementBlock)+ elseClause?
    ) END CASE
    ;

elseClause
    : ELSE statementBlock
    ;

continueStatement
    : CONTINUE (FOR | FOREACH | LOOP | WHILE)? SCOL?
    ;

exitStatement
    : EXIT (FOREACH | (FOR | LOOP | WHILE)? (identifier? WHEN condition)? ) SCOL?
    ;

forStatement
    : labelStatement? FOR identifier (IN OPEN_PAR (range | expression) (COMMA (range | expression))* CLOSE_PAR
    | ASSIGN range
    )
    (statementBlock END FOR identifier?
    | LOOP statementBlock END LOOP identifier?
    )
    SCOL?
    ;

range
    : left=expression TO right=expression (STEP expression)?
    ;

foreachStatement
    : FOREACH (WITH HOLD | identifier (WITH HOLD)? FOR | routineCall (INTO identifier (COMMA identifier)*)?)? selectStatement?
    statementBlock END FOREACH SCOL?
    ;

routineCall
    : EXECUTE (PROCEDURE procedureName | FUNCTION functionName)  OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR
    ;

gotoStatement
    : GOTO identifier SCOL?
    ;

ifStatement
    : IF condition THEN ifStatementList? (ELIF condition THEN ifStatementList?)+ (ELSE ifStatementList)? END IF SCOL?
    ;

ifStatementList
    : BEGIN statementBlock END
    | ifSubsetSplStatement
    | ifSubsetSqlStatement
    ;

ifSubsetSplStatement
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

ifSubsetSqlStatement
    : allocateDescriptor
    | closeDatabase
    | connectStatement
    | createDatabase
    | createProcedure
    | databaseStatement
    | deallocateDescriptor
    | describeStatement
    | disconnectStatement
    | dropDatabase
    | executeStatement
    | flushStatement
    | getDescriptor
    | getDiagnostics
    | infoStatement
    | loadStatement
    | outputStatement
    | putStatement
    | renameDatabase
    | setAutofree
//    | setConnection
//    | setDescriptor
    | unloadStatement
    | wheneverStatement
    ;

letStatement
    : LET identifier (COMMA identifier)*
    (functionName OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR
    | expression
    | OPEN_PAR selectStatement (COMMA selectStatement)* CLOSE_PAR
    )
    (COMMA (functionName OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR
    | expression
    | OPEN_PAR selectStatement (COMMA selectStatement)* CLOSE_PAR
    ))*
    ;

loopStatement
    : labelStatement? (WHILE condition | FOR identifier (IN OPEN_PAR (range | expression) CLOSE_PAR | ASSIGN range))?
    LOOP statementBlock END LOOP labelStatement? SCOL?
    ;

raiseException
    : RAISE EXCEPTION expression (COMMA expression (COMMA (identifier | expression))?) SCOL?
    ;

returnStatement
    : RETURN (expression (COMMA expression)* (WITH RESUME)?)? SCOL?
    ;

systemStatement
    : SYSTEM (identifier | expression) SCOL?
    ;

traceStatement
    : TRACE (ON | OFF | PROCEDURE | expression) SCOL?
    ;

whileStatement
    : labelStatement? WHILE condition (statementBlock END WHILE identifier? | LOOP statementBlock END LOOP identifier?) SCOL?
    ;

allocateCollection
    : ALLOCATE COLLECTION SYM1 identifier
    ;

allocateDescriptor
    : ALLOCATE DESCRIPTOR anyName (WITH MAX (numeric | identifier))?
    ;

allocateRow
    : ALLOCATE ROW SYM1 identifier
    ;

alterAccessMethod
    : ALTER ACCESS_METHOD identifier ((MODIFY | ADD) purposeOptions | DROP identifier)
    (COMMA (MODIFY | ADD) purposeOptions| DROP identifier)*
    ;

purposeOptions
    : identifier (ASSIGN (numeric | quotedString | identifier))?
    ;

alterFragment
    : ALTER FRAGMENT (
    ON TABLE tableName (attachClause | detachClause | initClause | addClause | dropClause | modifyClause)
    | ON INDEX indexName (initClause | addClause | dropClause | modifyClause)
    )
    ;

attachClause
    : ATTACH tableName asClause? (COMMA tableName asClause?)*
    ;

asClause
    : AS (PARTITION identifier)? (
    (expression | listExpression) ((AFTER | BEFORE)? identifier)?
    | rangeIntervalExpression
    | REMAINDER
    )
    ;

listExpression
    : VALUES (OPEN_PAR (constExprssion (COMMA constExprssion)*)? CLOSE_PAR | IS? NULL)
    ;

rangeIntervalExpression
    : VALUES (LT constExprssion | IS NULL)
    ;

detachClause
    : DETACH PARTITION? identifier tableName
    ;

initClause
    : INIT (WITH ROWIDS) identifier (fragmentClauseTables | (PARTITION identifier)? IN identifier | fragmentClauseIndexes)
    ;

fragmentClauseTables
    : (FRAGMENT | PARTITION) BY ( ROUND ROBIN (IN identifier (COMMA identifier)* | PARTITION identifier IN identifier (COMMA PARTITION identifier IN identifier)*)
    | EXPRESSION fragmentList
    | RANGE OPEN_PAR identifier CLOSE_PAR intervalFragmentClause
    | LIST OPEN_PAR identifier CLOSE_PAR listFragmentClause
    )
    ;

fragmentList
    : (PARTITION identifier)? (OPEN_PAR expression CLOSE_PAR | REMAINDER) IN identifier
    (COMMA (PARTITION identifier)? (OPEN_PAR expression CLOSE_PAR | REMAINDER) IN identifier)*
    ;

intervalFragmentClause
    : INTERVAL OPEN_PAR identifier? CLOSE_PAR rollingWindowClause? ((STORE IN)? (identifier (COMMA identifier)* | createFunction))?
    PARTITION identifier (VALUES LT constExprssion | VALUES IS NULL) IN identifier
    (COMMA PARTITION identifier (VALUES LT constExprssion | VALUES IS NULL) IN identifier)*
    ;

rollingWindowClause
    : ROLLING OPEN_PAR numeric FRAGMENTS CLOSE_PAR (DETACH | DISCARD)
    | (ROLLING OPEN_PAR numeric FRAGMENTS)? LIMIT TO numeric IDENTIFIER (DETACH | DISCARD) (INTERVAL FIRST | ANY | INTERVAL ONLY)?
    ;

listFragmentClause
    : PARTITION identifier listExpression IN identifier (COMMA PARTITION identifier listExpression IN identifier)*
    (COMMA PARTITION identifier REMAINDER IN identifier)?
    ;

fragmentClauseIndexes
    : (FRAGMENT | PARTITION) BY (RANGE OPEN_PAR expression CLOSE_PAR intervalFragmentClause
    | LIST OPEN_PAR expression CLOSE_PAR listFragmentClause
    | EXPRESSION expressionFragmentClause)
    ;

expressionFragmentClause
    : PARTITION identifier expression IN identifier (COMMA PARTITION identifier expression IN identifier)*
    (COMMA PARTITION identifier REMAINDER IN identifier)?
    ;

addClause
    : ADD (((PARTITION identifier)? REMAINDER IN)? (PARTITION identifier IN)? identifier
    | (PARTITION identifier)? addExpression IN identifier ((BEFORE | AFTER) identifier)?
    | INTERVAL STORE? IN identifier (COMMA identifier)*
    )
    ;

addExpression
    : VALUES (LT constExprssion | IS? NULL | constExprssion (COMMA constExprssion)*)
    | expression
    ;

dropClause
    : DROP (PARTITION? identifier | INTERVAL STORE? IN identifier (COMMA identifier)*)
    ;

modifyClause
    : MODIFY (PARTITION? identifier TO PARTITION identifier fragmentExpression? (IN identifier)?
    | INTERVAL ((ENABLED | DISABLED)? STORE? IN (identifier (COMMA identifier)* | createFunction OPEN_PAR CLOSE_PAR) | ENABLED | DISABLED | TRANSITION TO constExprssion | rollingWindowClause)
    )
    ;

fragmentExpression
    : expression
    | VALUES (LT constExprssion | IS? NULL | constExprssion (COMMA constExprssion)*)
    | REMAINDER
    ;

alterFunction
    : ALTER (FUNCTION functionName OPEN_PAR dataType (COMMA dataType)* CLOSE_PAR
    | SPECIFIC FUNCTION identifier
    ) WITH OPEN_PAR ((ADD | MODIFY | DROP) routineModifier | MODIFY EXTERNAL NAME ASSIGN sharedObjectFilename=identifier) CLOSE_PAR
    ;

routineModifier
    : addingOrModifyingRoutineModifier
    | droppingRoutineModifier
    ;

addingOrModifyingRoutineModifier
    :
    | NOT? VARIANT
    | NEGATOR ASSIGN identifier
    | CLASS ASSIGN identifier
    | ITERATOR
    | PARALLELIZABLE
    | HANDLESNULLS
    | INTERNAL
    | PERCALL_COST ASSIGN numeric
    | COSTFUNC ASSIGN identifier
    | SELFUNC ASSIGN identifier
    | SELCONST ASSIGN numeric
    | STACK ASSIGN numeric
    ;

droppingRoutineModifier
    : NOT? VARIANT
    | NEGATOR
    | CLASS
    | ITERATOR
    | PARALLELIZABLE
    | HANDLESNULLS
    | INTERNAL
    | COSTFUNC
    | PERCALL_COST
    | SELFUNC
    | SELCONST
    | STACK
    ;

alterIndex
    : ALTER INDEX indexName TO NOT? CLUSTER
    ;

alterProcedure
    : ALTER (PROCEDURE procedureName OPEN_PAR dataType (COMMA dataType)* CLOSE_PAR
    | SPECIFIC PROCEDURE procedureName
    ) WITH OPEN_PAR ((ADD | MODIFY | DROP) routineModifier | MODIFY EXTERNAL NAME ASSIGN sharedObjectFilename=identifier) CLOSE_PAR
    ;

alterRoutine
    : ALTER (ROUTINE routineName OPEN_PAR dataType (COMMA dataType)* CLOSE_PAR
    | SPECIFIC ROUTINE routineName
    ) WITH OPEN_PAR ((ADD | MODIFY | DROP) routineModifier | MODIFY EXTERNAL NAME ASSIGN sharedObjectFilename=identifier) CLOSE_PAR
    ;

alterSecurityLabelComponent
    : ALTER SECURITY LABEL COMPONENT identifier ADD (ARRAY SYM2 quotedString (COMMA quotedString)* (BEFORE | AFTER) quotedString (COMMA quotedString (COMMA quotedString)* (BEFORE | AFTER) quotedString)* SYM3
    | SET SYM4 quotedString (COMMA quotedString)* SYM5
    | TREE OPEN_PAR quotedString UNDER quotedString (COMMA quotedString UNDER quotedString)* CLOSE_PAR
    )
    ;

alterSequence
    : ALTER SEQUENCE identifier (
    | CYCLE | NOCYCLE
    | CACHE numeric | NOCACHE
    | ORDER | NOORDER
    | INCREMENT BY? numeric
    | RESTART WITH? numeric
    | MINVALUE numeric | NOMINVALUE
    | MAXVALUE numeric | NOMAXVALUE
    )
    ;

alterTable
    : ALTER TABLE tableName (basicTableOption | loggingTypeOption | addTypeClause | statisticsOptions)
    ;

basicTableOption
    : addColumnClause
    //| addAuditClause
    | addConstraintClause
    | addOrDropSpecializedColumns
//    | dropAuditClause
    | dropConstraintClause
    | dropColumnClause
    | lockModeClause
    | modifyClause
    | modifyExtentSizeClause
    | modifyNextSizeClause
    | putClause
    | securityPolicyClause
    ;

addColumnClause
    : ADD newColumnClause (COMMA newColumnClause)*
    ;

newColumnClause
    : newColumn=columnName dataType (defaultClause | singleColumnConstraint)* (BEFORE columnName)? (COLUMN? SECURED WITH identifier)?
    ;

defaultClause
    : DEFAULT (identifier | expression | NULL | USER | CURRENT_USER | (CURRENT | SYSDATE) datetimeField? | TODAY)
    ;

singleColumnConstraint
    : ((NULL | NOT NULL | UNIQUE | DISTINCT | PRIMARY KEY | referencesClause | checkClause)
    constraintDefinition?
    )+
    ;

referencesClause
    : REFERENCES tableName (OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR)? (ON DELETE CASCADE)?
    ;

checkClause
    : CHECK OPEN_PAR condition CLOSE_PAR
    ;

constraintDefinition
    : (CONSTRAINT identifier | ON DELETE CASCADE)? (DISABLED | (ENABLED | FILTERING (WITHOUT ERROR | WITH ERROR)?) NOVALIDATE?)
    | (CONSTRAINT identifier | ON DELETE CASCADE) (DISABLED | (ENABLED | FILTERING (WITHOUT ERROR | WITH ERROR)?) NOVALIDATE?)?
    ;

addConstraintClause
    : ADD CONSTRAINT OPEN_PAR? multipleColumnConstraint+ CLOSE_PAR?
    ;

multipleColumnConstraint
    : (
    ((NOT NULL | NULL | UNIQUE | DISTINCT | PRIMARY KEY | referencesClause) OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR
    | checkClause
    | foreignKeyDefinition
    ) constraintDefinition?
    | foreignKeyDefinition (CONSTRAINT identifier)? INDEX DISABLED
    )
    ;

foreignKeyDefinition
    : FOREIGN KEY OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR referencesClause
    ;

addOrDropSpecializedColumns
    : (ADD | DROP) (CRCOLS | ERKEY | REPLCHECK | ROWIDS | VERCOLS)
    ;

dropConstraintClause
    : DROP CONSTRAINT OPEN_PAR? identifier (COMMA identifier)* CLOSE_PAR?
    ;

dropColumnClause
    : DROP (columnName | OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR)
    ;

lockModeClause
    : LOCK MODE OPEN_PAR (PAGE | ROW) CLOSE_PAR
    ;

modifyExtentSizeClause
    : MODIFY EXTENT SIZE expression
    ;

modifyNextSizeClause
    : MODIFY NEXT SIZE expression
    ;

putClause
    : PUT identifier IN OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR
    (OPEN_PAR (COMMA? (EXTENT SIZE numeric | NO LOG | LOG | HIGH INTEG | MODERATE INTEG | NO? KEEP ACCESS TIME))* CLOSE_PAR)?
    ;

securityPolicyClause
    : ADD SECURITY POLICY identifier
    | DROP SECURITY POLICY
    ;

loggingTypeOption
    : TYPE OPEN_PAR (STANDARD | RAW) CLOSE_PAR
    ;

addTypeClause
    : ADD TYPE rowType=identifier
    ;

statisticsOptions
    : (STATCHANGE (AUTO | numeric)) (STATLEVEL (FRAGMENT | TABLE | AUTO))?
    | (STATCHANGE (AUTO | numeric))? (STATLEVEL (FRAGMENT | TABLE | AUTO))
    ;

alterTrustedContext
    : ALTER TRUSTED CONTEXT identifier ( ALTER (SYSTEM AUTHID identifier | ATTRIBUTES OPEN_PAR ADDRESS identifier (COMMA ADDRESS identifier)* CLOSE_PAR | NO DEFAULT ROLE | DEFAULT ROLE identifier | DISABLE | ENABLE)+
    | (ADD | DROP) ATTRIBUTES OPEN_PAR ADDRESS identifier (COMMA ADDRESS identifier)* CLOSE_PAR
    | (ADD | REPLACE) USE FOR authorizedUserClause (COMMA authorizedUserClause)*
    | DROP USE FOR (userName | PUBLIC) (COMMA (userName | PUBLIC))*
    )+
    ;

authorizedUserClause
    : (userName (ROLE roleName)? | PUBLIC) (WITH AUTHENTICATION | WITHOUT AUTHENTICATION)?
    ;

alterUser
    : ALTER (DEFAULT USER | USER userName (LOCK | ACCOUNT UNLOCK)?)
    (COMMA? ((ADD | MODIFY) (PASSWORD quotedString | UID numeric | GROUP OPEN_PAR (identifier | numeric) (COMMA (identifier | numeric))* CLOSE_PAR
    | USER identifier | AUTHORIZATION OPEN_PAR (DBSA | DBSSO | AAO | BARGROUP) (COMMA (DBSA | DBSSO | AAO | BARGROUP))* CLOSE_PAR | HOME quotedString)
    | DROP (PASSWORD | UID | GROUP OPEN_PAR (identifier | numeric) (COMMA (identifier | numeric))* CLOSE_PAR
    | USER | AUTHORIZATION OPEN_PAR (DBSA | DBSSO | AAO | BARGROUP) (COMMA (DBSA | DBSSO | AAO | BARGROUP))* CLOSE_PAR | HOME)
    )
    )+
    ;

beginWork
    : BEGIN WORK? (WITHOUT REPLICATION)?
    ;

closeStatement
    : CLOSE cursorId=identifier
    ;

closeDatabase
    : CLOSE DATABASE
    ;

commitWork
    : COMMIT WORK?
    ;

connectStatement
    : CONNECT TO (databaseEnvironment (userAuthenticationClause TRUSTED?)? | DEFAULT) (WITH CONCURRENT TRANSACTION)?
    ;

databaseEnvironment
    : identifier (AS identifier)?
    ;

userAuthenticationClause
    : USER quotedString USING quotedString
    ;

createAccessMethod
    : CREATE (SECONDARY | PRIMARY) ACCESS_METHOD (IF NOT EXISTS)? identifier OPEN_PAR purposeOptions (COMMA purposeOptions)* CLOSE_PAR
    ;

createAggregate
    : CREATE AGGREGATE (IF NOT EXISTS)? identifier WITH OPEN_PAR modifiers (COMMA modifiers)* CLOSE_PAR
    ;

modifiers
    : INIT ASSIGN identifier
    | ITER ASSIGN identifier
    | COMBINE ASSIGN identifier
    | FINAL ASSIGN identifier
    | HANDLESNULLS
    ;

createCast
    : CREATE (EXPLICIT | IMPLICIT)? CAST (IF NOT EXISTS)? OPEN_PAR dataType AS dataType (WITH FUNCTION functionName)? CLOSE_PAR
    ;

createDatabase
    : CREATE DATABASE (IF NOT EXISTS)? identifier (IN identifier)? (WITH (BUFFERED? LOG | LOG MODE ANSI))? (NLSCASE SENSITIVE | NLSCASE INSENSITIVE)?
    ;

createDistinctType
    : CREATE DISTINCT TYPE (IF NOT EXISTS)? dataType AS dataType
    ;

createDefaultUser
    : CREATE DEFAULT USER WITH properties
    ;

createExternalTable
    : CREATE EXTERNAL TABLE (IF NOT EXISTS)? tableName columnDeifinition USING OPEN_PAR tableOptions? datafilesClause tableOptions? CLOSE_PAR
    ;

columnDeifinition
    : SAMEAS identifier
    | columnName dataType otherOptionalClause (COMMA columnName dataType otherOptionalClause)*
    ;

dataType
    : builtInDataTypes
    | complexDataType
    | dataTypeName
    ;

otherOptionalClause
    : (EXTERNAL CHAR OPEN_PAR numeric CLOSE_PAR (NULL STRING_LITERAL | NOT NULL)?
    )?
    ;

tableOptions
    : (FORMAT SYM6 (DELIMITED | INFORMIX | FIXED) SYM6
    | DEFAULT | loadingModeOption | DBDATE STRING_LITERAL | DBMONEY STRING_LITERAL | DELIMITER STRING_LITERAL | RECORDEND STRING_LITERAL
    | MAXERRORS numeric | REJECTFILE STRING_LITERAL | ESCAPE (ON | OFF)? | (NUMROWS | SIZE) numeric
    )
    ;

loadingModeOption
    : EXPRESS | DELUXE
    ;

datafilesClause
    : DATAFILES OPEN_PAR STRING_LITERAL (COMMA STRING_LITERAL)* CLOSE_PAR
    ;

createFunction
    : CREATE DBA? FUNCTION (IF NOT EXISTS)? functionName OPEN_PAR routineParameterList? CLOSE_PAR
    (referencingClause FOR tableName)? returnClause (SPECIFIC identifier)? (WITH OPEN_PAR routineModifier (COMMA routineModifier)* CLOSE_PAR)? SCOL?
    (statementBlock | externalRoutineReference) END FUNCTION (DOCUMENT quotedString (COMMA quotedString)*)? (WITH LISTING IN STRING_LITERAL)?
    ;

routineParameterList
    : (IN | OUT | INOUT)? parameter (COMMA (IN | OUT | INOUT)? parameter)*
    ;

parameter
    : identifier? ((dataType | LINK columnName) (DEFAULT numeric)?
    | REFERENCES (BYTE | TEXT) (DEFAULT NULL)?
    )
    ;

referencingClause
    : REFERENCING ((NEW | OLD) AS? identifier)+ FOR tableName
    ;

returnClause
    : (RETURNING | RETURNS) (dataTypeName | REFERENCE (BYTE | TEXT)) (AS identifier)?
    (COMMA (RETURNING | RETURNS) (dataTypeName | REFERENCE (BYTE | TEXT)) (AS identifier)?)*
    ;

statementBlock
    : defineStatement* onException* statementBlockClause*
    ;

statementBlockClause
    : executeFunction
    | executeProcedure
    | subsetSplStatement
    | subsetSqlStatement
    | BEGIN statementBlock END
    ;

defineStatement
    : DEFINE (
    | identifier (COMMA identifier)* (dataTypeName | REFERENCES (BYTE | TEXT) | LIKE columnName | PROCEDURE | BLOB | CLOB | subsetComplexDataTypes)
    ) SCOL?
    ;

onException
    : ON EXCEPTION (IN OPEN_PAR numeric (COMMA numeric)* CLOSE_PAR)?
    (SET identifier (COMMA identifier)?)? statementBlock END EXCEPTION (WITH RESUME)?
    SCOL?
    ;

subsetComplexDataTypes
    : COLLECTION
    | (SET | MULTISET | LIST) OPEN_PAR (dataTypeName | (SET | MULTISET | LIST) dataTypeName NOT NULL) NOT NULL CLOSE_PAR
    | row=identifier
    | ROW (OPEN_PAR field=identifier dataTypeName (COMMA field=identifier dataTypeName)* CLOSE_PAR)
    ;

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

subsetSqlStatement
    : closeDatabase
    | connectStatement
    | createDatabase
    | createFunction
    | createFunctionFrom
    | createProcedure
    | createProcedureFrom
    | createRoutineFrom
    | databaseStatement
    | disconnectStatement
    | dropDatabase
    | executeStatement
    | flushStatement
    | infoStatement
    | loadStatement
    | outputStatement
    | putStatement
    | renameDatabase
    | setAutofree
//    | setConnection
    | unloadStatement
    | updateStatistics
    ;

externalRoutineReference
    : EXTERNAL NAME quotedString LANGUAGE (C | JAVA) (PARAMETER STYLE INFORMIX?)? (NOT? VARIANT)?
    ;

createFunctionFrom
    : CREATE FUNCTION FROM (IF NOT EXISTS)? identifier
    ;

createIndex
    : CREATE indexTypeOptions INDEX (IF NOT EXISTS)? indexName ON tableName indexKeySpecs indexOptions?
    ;

indexTypeOptions
    : (DISTINCT | UNIQUE)? CLUSTER?
    ;

indexKeySpecs
    : OPEN_PAR (columnName | functionName OPEN_PAR func_col=columnName (COMMA func_col=columnName | op_class=identifier | bsonFieldSpecification)* CLOSE_PAR) (ASC | DESC)? CLOSE_PAR
    ;

bsonFieldSpecification
    : (BSON_GET | bsonValueFunction) OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR USING BSON
    ;

bsonValueFunction
    : BSON_VALUE_BIGINT | BSON_VALUE_BOOLEAN | BSON_VALUE_DATE | BSON_VALUE_DOUBLE | BSON_VALUE_INT | BSON_VALUE_LVARCHAR | BSON_VALUE_TIMESTAMP | BSON_VALUE_VARCHAR
    ;

indexOptions
    : usingAccessMethodClause filefactorOption? storageOptions? indexModes? hashOnClause? ONLINE? COMPRESSED? extentSizeOptions?
    | usingAccessMethodClause? filefactorOption storageOptions? indexModes? hashOnClause? ONLINE? COMPRESSED? extentSizeOptions?
    | usingAccessMethodClause? filefactorOption? storageOptions indexModes? hashOnClause? ONLINE? COMPRESSED? extentSizeOptions?
    | usingAccessMethodClause? filefactorOption? storageOptions? indexModes hashOnClause? ONLINE? COMPRESSED? extentSizeOptions?
    | usingAccessMethodClause? filefactorOption? storageOptions? indexModes? hashOnClause ONLINE? COMPRESSED? extentSizeOptions?
    | usingAccessMethodClause? filefactorOption? storageOptions? indexModes? hashOnClause? ONLINE COMPRESSED? extentSizeOptions?
    | usingAccessMethodClause? filefactorOption? storageOptions? indexModes? hashOnClause? ONLINE? COMPRESSED extentSizeOptions?
    | usingAccessMethodClause? filefactorOption? storageOptions? indexModes? hashOnClause? ONLINE? COMPRESSED? extentSizeOptions
    ;

usingAccessMethodClause
    : USING identifier OPEN_PAR quotedString ASSIGN (quotedString | numeric) (COMMA quotedString ASSIGN (quotedString | numeric))* CLOSE_PAR
    ;

filefactorOption
    : FILLFACTOR numeric
    ;

storageOptions
    : IN (identifier | TABLE)
    | fragmentClauseIndexes
    ;

indexModes
    : ENABLED | DISABLED | FILTERING (WITHOUT ERROR | WITH ERROR)?
    ;

hashOnClause
    : HASH ON OPEN_PAR identifier (COMMA identifier)? CLOSE_PAR WITH numeric BUCKETS
    ;

extentSizeOptions
    : (EXTENT SIZE expression) (NEXT SIZE expression)?
    | (EXTENT SIZE expression)? (NEXT SIZE expression)
    ;

createOpaqueType
    : CREATE OPAQUE TYPE (IF NOT EXISTS)? identifier OPEN_PAR INTERNALLENGTH ASSIGN (numeric | VARIABLE) (COMMA opaqueTypeModifier)* CLOSE_PAR
    ;

opaqueTypeModifier
    : MAXLEN ASSIGN numeric
    | CANNOTHASH
    | PASSEDBYVALUE
    | ALIGNMENT ASSIGN numeric
    ;

createOpclass
    : CREATE OPCLASS (IF NOT EXISTS)? identifier FOR identifier STRATEGIES OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR
    SUPPORT OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR
    ;

createProcedure
    : CREATE DBA? PROCEDURE (IF NOT EXISTS)? procedureName OPEN_PAR routineParameterList? CLOSE_PAR (referencingClause FOR tableName)? returnClause? (SPECIFIC identifier)?
    (WITH OPEN_PAR routineModifier (COMMA routineModifier)* CLOSE_PAR)? SCOL? (statementBlock | externalRoutineReference) END PROCEDURE
    (DOCUMENT quotedString (COMMA quotedString)*)? (WITH LISTING IN STRING_LITERAL)?
    ;

createProcedureFrom
    : CREATE PROCEDURE FROM (IF NOT EXISTS)? identifier
    ;

//https://www.ibm.com/docs/en/informix-servers/14.10?topic=statements-create-role-statement
createRole
    : CREATE ROLE (IF NOT EXISTS)? roleName
    ;

createRoutineFrom
    : CREATE ROUTINE (IF NOT EXISTS)? identifier
    ;

createRowType
    : CREATE ROW TYPE (IF NOT EXISTS)? identifier (OPEN_PAR fieldDefinition (COMMA fieldDefinition)* CLOSE_PAR
    | (OPEN_PAR fieldDefinition (COMMA fieldDefinition)* CLOSE_PAR)? UNDER dataType
    )
    ;

fieldDefinition
    : identifier dataTypeName (NOT NULL)?
    ;

createSchema
    : CREATE SCHEMA AUTHORIZATION userName (createTable | createView | grantStatement | createIndex | createSynonym | createTrigger | createSequence | createRowType
    | createOpaqueType | createDistinctType | createCast)+ SCOL?
    ;

createSecurityLabel
    : CREATE SECURITY LABEL (IF NOT EXISTS)? identifier COMPONENT identifier quotedString (COMMA quotedString)* (COMMA COMPONENT identifier quotedString (COMMA quotedString)*)*
    ;

createSecurityLabelComponent
    : CREATE SECURITY LABEL COMPONENT (IF NOT EXISTS)? identifier (ARRAY SYM2 quotedString (COMMA quotedString)* SYM3
    | SET SYM4 quotedString (COMMA quotedString)* SYM5
    | TREE OPEN_PAR quotedString ROOT (COMMA quotedString UNDER quotedString)* CLOSE_PAR
    )
    ;

createSecurityPolicy
    : CREATE SECURITY POLICY (IF NOT EXISTS)? identifier COMPONENTS identifier (COMMA identifier)* (WITH IDSLBACRULES)?
    (RESTRICT NOT AUTHORIZED WRITE SECURITY LABEL | OVERRIDE NOT AUTHORIZED WRITE SECURITY LABEL)?
    ;

createSequence
    : CREATE SEQUENCE (IF NOT EXISTS)? identifier (INCREMENT BY? numeric
    | START WITH? numeric | NOMAXVALUE | MAXVALUE numeric | NOMINVALUE | MINVALUE numeric
    | NOCYCLE | CYCLE | CACHE numeric | NOCACHE | ORDER | NOORDER
    )*
    ;

createSynonym
    : CREATE (PUBLIC | PRIVATE)? SYNONYM (IF NOT EXISTS)? identifier FOR tableName
    ;

createTable
    : CREATE (STANDARD | RAW)? TABLE (IF NOT EXISTS)? tableName (columnDeifinition (COMMA columnDeifinition)*
    (COMMA (multipleColumnConstraintForTable | columnDeifinition))* | ofTypeClause)
    (WITH (AUDIT | CRCOLS | ERKEY | REPLCHECK | VERCOLS))? (COMMA (WITH (AUDIT | CRCOLS | ERKEY | REPLCHECK | VERCOLS)))*
    securityPolicyClause? storageOptions? (LOCK MODE (PAGE | ROW))? usingAccessMethodClause? statisticsOptions?
    ;

multipleColumnConstraintForTable
    : (NOT? NULL (UNIQUE | DISTINCT | PRIMARY KEY | referencesClause) OPEN_PAR identifier (COMMA identifier)* referencesClause CLOSE_PAR
    | FOREIGN KEY OPEN_PAR identifier (COMMA identifier)* referencesClause CLOSE_PAR | checkClause) constraintDefinition
    ;

ofTypeClause
    : OF TYPE identifier (OPEN_PAR multipleColumnConstraintForTable CLOSE_PAR)? (UNDER identifier)?
    ;

createTempTable
    : CREATE TEMP TABLE (IF NOT EXISTS)? tableName OPEN_PAR columnDeifinition (COMMA columnDeifinition)* (COMMA (multipleColumnConstraintForTable | columnDeifinition))* CLOSE_PAR
    (WITH NO LOG)? storageOptions? lockModeClause? usingAccessMethodClause?
    ;

createTrigger
    : CREATE TRIGGER (IF NOT EXISTS)? identifier (triggerOnTable | INSTEAD OF triggerOnView) (ENABLED | DISABLED)?
    ;

triggerOnTable
    : (DELETE | SELECT (OF columnName (COMMA columnName)*)?) ON tableName deleteAndSelectClause
    | UPDATE (OF columnName (COMMA columnName)*)? ON tableName updateClause
    | INSERT ON tableName (REFERENCING NEW AS? identifier correlatedTableAction | actionClause)
    ;

deleteAndSelectClause
    : actionClause
    | REFERENCING ((OLD | NEW) AS? identifier)+ correlatedTableAction
    ;

updateClause
    : actionClause
    | REFERENCING ((OLD | NEW) AS? identifier)+ (REFERENCING NEW AS? identifier)? correlatedTableAction
    ;

actionClause
    : (BEFORE triggeredAction)? (FOR EACH ROW triggeredAction)? (AFTER triggeredAction)?
    ;

correlatedTableAction
    : (BEFORE triggeredAction)? FOR EACH ROW triggeredAction (AFTER triggeredAction)?
    ;

triggeredAction
    : (WHEN OPEN_PAR condition CLOSE_PAR)? OPEN_PAR (insertStatement | updateStatement | deleteStatement | executeProcedure | executeFunction)
    (COMMA (insertStatement | updateStatement | deleteStatement | executeProcedure | executeFunction))* CLOSE_PAR
    (COMMA (WHEN OPEN_PAR condition CLOSE_PAR)? OPEN_PAR (insertStatement | updateStatement | deleteStatement | executeProcedure | executeFunction)
    (COMMA (insertStatement | updateStatement | deleteStatement | executeProcedure | executeFunction))* CLOSE_PAR)*
    ;

triggerOnView
    : INSERT ON viewName (REFERENCING NEW AS? identifier)?
    | DELETE ON viewName (REFERENCING OLD AS? identifier)?
    | UPDATE ON viewName ((REFERENCING OLD AS? identifier (NEW AS? identifier)?)? | (REFERENCING NEW AS? identifier (OLD AS? identifier)?)?)
    ;

createTrustedContext
    : CREATE TRUSTED CONTEXT identifier (USER | BASED UPON CONNECTION USING SYSTEM AUTHID) identifier
    (ATTRIBUTES OPEN_PAR ADDRESS quotedString (COMMA ADDRESS quotedString)* CLOSE_PAR
    | WITH USE FOR authorizedUserClause (COMMA authorizedUserClause)*
    | NO DEFAULT ROLE | DEFAULT ROLE identifier | DISABLE | ENABLE
    )*
    ;

createUser
    : CREATE (DEFAULT USER WITH (ACCOUNT UNLOCK | ACCOUNT LOCK)? properties | USER userName (PASSWORD quotedString)? (ACCOUNT UNLOCK | ACCOUNT LOCK)? properties?)
    ;

properties
    : PROPERTIES (UID numeric GROUP OPEN_PAR (numeric | identifier) (COMMA (numeric | identifier))* CLOSE_PAR
    | USER identifier (GROUP OPEN_PAR (numeric | identifier) (COMMA (numeric | identifier))* CLOSE_PAR)?)
    (HOME quotedString)? (AUTHORIZATION OPEN_PAR (DBSA | DBSSO | AAO | BARGROUP) (COMMA (DBSA | DBSSO | AAO | BARGROUP))* CLOSE_PAR)?
    ;

createView
    : CREATE VIEW (IF NOT EXISTS)? viewName identifier? (OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR | OF TYPE dataType)? AS
    selectStatement (WITH CHECK OPTION)?
    ;

createXadatasource
    : CREATE XADATASOURCE (IF NOT EXISTS)? identifier USING identifier
    ;

createXadatasourceType
    : CREATE XADATASOURCE TYPE (IF NOT EXISTS)? identifier OPEN_PAR purposeOptions (COMMA purposeOptions)* CLOSE_PAR
    ;

databaseStatement
    : DATABASE databaseName=anyName EXCLUSIVE?
    ;

deallocateCollection
    : DEALLOCATE COLLECTION SYM1 identifier
    ;

deallocateDescriptor
    : DEALLOCATE DESCRIPTOR (identifier | quotedString)
    ;

deallocateRow
    : DEALLOCATE ROW SYM1 identifier
    ;

declareStatement
    : DECLARE identifier (CURSOR (WITH HOLD)? (FOR insertStatement | forUpdateOrForReadOnlySelectOptions | otherSelectOrFunctionOptions)
    | SCROLL CURSOR (WITH HOLD)? otherSelectOrFunctionOptions
    //|  CURSOR FOR
    )
    ;

forUpdateOrForReadOnlySelectOptions
    : selectStatement (FOR READ ONLY | FOR UPDATE (OF identifier (COMMA identifier)*))
    ;

otherSelectOrFunctionOptions
    : FOR (selectStatement | identifier | executeProcedure | executeFunction)
    ;

deleteStatement
    : DELETE optimizerDirectives? FROM? (tableName (AS? alias)? | ONLY OPEN_PAR tableName CLOSE_PAR | collectionDerivedTable)
    (WHERE condition | WHERE CURRENT OF identifier)?
    ;

optimizerDirectives //TODO: add more
    : (SYM9 | SYM4 | SYM10)  (SYM4 | SYM11)?
    ;

describeStatement
    : DESCRIBE OUTPUT? identifier (USING SQL DESCRIPTOR (identifier | quotedString)
    | INTO SQL DESCRIPTOR (identifier | quotedString) //TODO: sqlda_pointer
    )
    ;

describeInput
    : DESCRIBE INPUT identifier (USING SQL DESCRIPTOR (identifier | quotedString)
    | INTO SQL DESCRIPTOR (identifier | quotedString) //TODO: sqlda_pointer
    )
    ;

disconnectStatement
    : DISCONNECT (CURRENT | ALL | DEFAULT | identifier | quotedString)
    ;

dropAccessMethod
    : DROP ACCESS_METHOD (IF EXISTS)? accessMethodName=identifier RESTRICT
    ;

dropAggregate
    : DROP AGGREGATE (IF EXISTS)? aggregateName=identifier
    ;

dropCast
    : DROP CAST (IF EXISTS)? OPEN_PAR dataType AS dataType CLOSE_PAR
    ;

dropDatabase
    : DROP DATABASE (IF EXISTS)? databaseName=identifier
    ;

dropFunction
    : DROP FUNCTION (IF EXISTS)? (functionName (OPEN_PAR dataType (COMMA dataType)* CLOSE_PAR)?
    | SPECIFIC FUNCTION (IF EXISTS)? identifier)
    ;

dropIndex
    : DROP INDEX (IF EXISTS)? indexName ONLINE?
    ;

dropOpclass
    : DROP OPCLASS (IF EXISTS)? identifier RESTRICT
    ;

dropProcedure
    : DROP (PROCEDURE (IF EXISTS)? procedureName (OPEN_PAR dataType (COMMA dataType)* CLOSE_PAR)?
    | SPECIFIC PROCEDURE (IF EXISTS)? identifier
    )
    ;

//https://www.ibm.com/docs/en/informix-servers/14.10?topic=statements-drop-role-statement
dropRole
    : DROP ROLE (IF EXISTS)? roleName
    ;

dropRoutine
    : DROP (ROUTINE (IF EXISTS)? routineName (OPEN_PAR dataType (COMMA dataType)* CLOSE_PAR)?
    | SPECIFIC ROUTINE (IF EXISTS)? identifier
    )
    ;

dropRowType
    : DROP ROW TYPE (IF EXISTS)? dataType RESTRICT
    ;

dropSecurity
    : DROP SECURITY (IF EXISTS)? (LABEL (IF EXISTS)? identifier RESTRICT?
    | LABEL COMPONENT (IF EXISTS)? identifier RESTRICT?
    | POLICY (IF EXISTS)? identifier (CASCADE | RESTRICT)
    )
    ;

dropSequence
    : DROP SEQUENCE (IF EXISTS)? identifier
    ;

dropSynonym
    : DROP SYNONYM (IF EXISTS)? synonymName=identifier
    ;

//https://www.ibm.com/docs/en/informix-servers/14.10?topic=statements-drop-table-statement
dropTable
    : DROP TABLE (IF EXISTS)? tableName (CASCADE | RESTRICT)?
    ;

dropTrigger
    : DROP TRIGGER (IF EXISTS)? triggerName=identifier
    ;

dropTrustedContext
    : DROP TRUSTED CONTEXT contextName=anyName
    ;

//https://www.ibm.com/docs/en/informix-servers/14.10?topic=statements-drop-type-statement
dropType
    : DROP TYPE (IF EXISTS)? dataTypeName RESTRICT
    ;

//https://www.ibm.com/docs/en/informix-servers/14.10?topic=statements-drop-user-statement-unix-linux
dropUser
    : DROP USER userName
    ;

//https://www.ibm.com/docs/en/informix-servers/14.10?topic=statements-drop-view-statement
dropView
    : DROP VIEW (IF EXISTS)? viewName (CASCADE | RESTRICT)?
    ;

dropXadatasource
    : DROP XADATASOURCE (IF EXISTS)? xaSourceName=identifier RESTRICT
    ;

dropXadatasourceType
    : DROP XADATASOURCE TYPE (IF EXISTS)? xaSourceName=identifier RESTRICT
    ;

executeStatement
    : EXECUTE intoClause? usingClause?
    ;

intoClause
    : INTO (identifier ((SYM1 | INDICATOR) identifier)?
    | SQL DESCRIPTOR (identifier | quotedString)
    //| DESCRIPTOR sqlda_pointer TODO: sqlda_pointer
    )
    ;

usingClause
    : USING (identifier ((SYM1 | INDICATOR) identifier)?
    | SQL DESCRIPTOR (identifier | quotedString)
    //| DESCRIPTOR sqlda_pointer TODO: sqlda_pointer
    )
    ;

executeFunction
    : EXECUTE FUNCTION functionName (OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR)? intoClause? (WITH TRIGGER REFERENCES)?
    ;

argument
    : (identifier ASSIGN) (expression | NULL | OPEN_PAR selectStatement CLOSE_PAR)
    ;

executeImmediate
    : EXECUTE IMMEDIATE (quotedString | identifier | expression)
    ;

executeProcedure
    : EXECUTE PROCEDURE procedureName (OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR)? (INTO identifier (COMMA identifier)*)? (WITH TRIGGER REFERENCES)?
    ;

fetchStatement
    : FETCH (NEXT | PRIOR | FIRST | LAST | CURRENT | ABSOLUTE PLUS? (numeric | identifier) | RELATIVE (numeric | identifier))? identifier
    (USING SQL DESCRIPTOR (identifier | quotedString)
    | INTO (identifier | identifier ((INDICATOR | SYM1) identifier)?) (COMMA (identifier | identifier ((INDICATOR | SYM1) identifier)?))*)?
    ;

flushStatement
    : FLUSH identifier
    ;

freeStatement
    : FREE identifier
    ;

getDescriptor
    : GET DESCRIPTOR (identifier | quotedString) (identifier ASSIGN COUNT | VALUE (identifier | numeric) describedItemInformation (COMMA describedItemInformation)*)
    ;

describedItemInformation
    : identifier ASSIGN (DATA | IDATA | ILENGTH | INDICATOR | ITYPE | LENGTH | NAME | NULLABLE| PRECISION | SCALE
    | TYPE | EXTYPEID | EXTYPELENGTH | EXTYPENAME | EXTYPEOWNERLENGTH | EXTYPEOWNERNAME | SOURCEID | SOURCETYPE)
    ;

getDiagnostics
    : GET DIAGNOSTICS (statementClause | exceptionClause)
    ;

statementClause
    : identifier ASSIGN (identifier | numeric)
    ;

exceptionClause
    : EXCEPTION (identifier | numeric) identifier ASSIGN (identifier | numeric) (COMMA identifier ASSIGN (identifier | numeric))*
    ;

grantStatement
    : GRANT ((CONNECT | RESOURCE | DBA | DEFAULT ROLE rolename) TO (PUBLIC | userName (COMMA userName)*)
    | rolename toOptions | securityAdministrationOptions | accessToPropertiesClause
    | (tableLevelPrivileges | routineLevelPrivileges | languageLevelPrivileges | typeLevelPrivileges | sequenceLevelPrivileges) toOptions
    )
    ;

rolename
    : identifier
    | STRING_LITERAL
    ;

toOptions
    : TO (userName (COMMA userName)* (WITH GRANT OPTION)? | (userName | PUBLIC) (COMMA (userName | PUBLIC))*) (AS STRING_LITERAL)?
    ;

securityAdministrationOptions
    : dbsecadmClause
    | exemptionClause
    | securityLabelClause
    | setsessionauthClause
    ;

dbsecadmClause
    : DBSECADM TO USER identifier (COMMA USER identifier)*
    ;

exemptionClause
    : EXEMPTION ON RULE (IDSLBACREADARRAY | IDSLBACREADTREE | IDSLBACREADSET
    | IDSLBACWRITEARRAY (WRITEDOWN | WRITEUP)? | IDSLBACWRITESET | IDSLBACWRITETREE | ALL)
    FOR policy=identifier TO USER identifier (COMMA USER identifier)*
    ;

securityLabelClause
    : SECURITY LABEL policy=identifier DOT label=identifier TO USER identifier (COMMA USER identifier)*
    (FOR ALL ACCESS | FOR READ ACCESS | FOR WRITE ACCESS)?
    ;

setsessionauthClause
    : SETSESSIONAUTH ON (PUBLIC | USER identifier (COMMA USER identifier)*) TO (USER | ROLE) identifier (COMMA (USER | ROLE) identifier)*
    ;

accessToPropertiesClause
    : ACCESS TO (PUBLIC | USER identifier (COMMA USER identifier)*) PROPERTIES
    (UID numeric GROUP OPEN_PAR (numeric | identifier) (COMMA (numeric | identifier))* CLOSE_PAR
    | USER identifier (GROUP OPEN_PAR (numeric | identifier) (COMMA (numeric | identifier))* CLOSE_PAR)?)
    (HOME quotedString)? (COMMA AUTHORIZATION OPEN_PAR quotedString (COMMA quotedString)* CLOSE_PAR)?
    ;

tableLevelPrivileges
    : (ALL PRIVILEGES?
    | (INSERT | DELETE | ALTER | INDEX | UNDER | (UPDATE | SELECT | REFERENCES) (OPEN_PAR (identifier (COMMA identifier)*)? CLOSE_PAR)?
    (COMMA (INSERT | DELETE | ALTER | INDEX | UNDER | (UPDATE | SELECT | REFERENCES) (OPEN_PAR (identifier (COMMA identifier)*)? CLOSE_PAR)?))*
    )
    ) ON tableName
    ;

routineLevelPrivileges
    : EXECUTE ON (spl_routine=identifier
    | (PROCEDURE | FUNCTION | ROUTINE) routine=identifier OPEN_PAR routineParameterList? CLOSE_PAR
    | SPECIFIC (PROCEDURE | FUNCTION | ROUTINE) identifier
    )
    ;

languageLevelPrivileges
    : USAGE ON LANGUAGE (SPL | C | JAVA)
    ;

typeLevelPrivileges
    : USAGE ON TYPE dataTypeName
    | UNDER ON TYPE dataTypeName
    ;

sequenceLevelPrivileges
    : (ALL | (SELECT | ALTER) (COMMA (SELECT | ALTER))*) ON identifier
    ;

grantFragment
    : GRANT FRAGMENT fragmentLevelPrivileges ON tableName OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR
    TO (PUBLIC | STRING_LITERAL (COMMA STRING_LITERAL)*) (WITH GRANT OPTION)? AS STRING_LITERAL
    ;

fragmentLevelPrivileges
    : ALL | (INSERT | DELETE | UPDATE) (COMMA (INSERT | DELETE | UPDATE))*
    ;

infoStatement
    : INFO (TABLES |
    (COLUMNS | INDEXES | STATUS | PRIVILEGES | ACCESS | FRAGMENTS | REFERENCES) FOR tableName
    )
    ;

insertStatement
    : INSERT (INTO tableName (OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR)? (valuesClause | executeRoutineClause | selectStatement)
    | (AT position=numeric)? INTO collectionDerivedTable fieldOptions
    )
    ;

executeRoutineClause
    : EXECUTE (PROCEDURE | FUNCTION) identifier OPEN_PAR argument (COMMA argument)* CLOSE_PAR
    ;

collectionDerivedTable
    : TABLE OPEN_PAR (expression CLOSE_PAR (AS? alias)? (OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR)?
    | identifier CLOSE_PAR
    )
    ;

fieldOptions
    : (field=identifier (COMMA field=identifier)*)? (valuesClause | selectStatement)
    ;

valuesClause
    : VALUES OPEN_PAR (identifier (SYM1 identifier | SYM12 identifier)?
    | NULL | USER | quotedString | numeric | constExprssion | columnExpression
    | literalCollection | literalRow | expression
    ) CLOSE_PAR
    ;

loadStatement
    : LOAD FROM quotedString (DELIMITER quotedString)? INSERT INTO tableName (OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR)?
    ;

lockTable
    : LOCK TABLE tableName IN (EXCLUSIVE | SHARE) MODE
    ;

mergeStatement
    : MERGE (optimizerDirectives (COMMA optimizerDirectives)*)? INTO tableName (AS alias)?
    USING (tableName | selectStatement) (AS alias)? (OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR)? ON condition
    (WHEN NOT MATCHED THEN INSERT (OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR)? valuesClause
    | WHEN MATCHED THEN DELETE | WHEN MATCHED THEN UPDATE setClause
    )
    ;

setClause
    : SET (singleColumnFormat | multipleColumnFormat)
    ;

singleColumnFormat
    : column=identifier ASSIGN (expression | NULL | OPEN_PAR selectStatement CLOSE_PAR)
    (COMMA column=identifier ASSIGN (expression | NULL | OPEN_PAR selectStatement CLOSE_PAR))*
    ;

multipleColumnFormat
    : (OPEN_PAR column=identifier (COMMA column=identifier)* CLOSE_PAR | STAR) ASSIGN OPEN_PAR
    (expression | NULL | OPEN_PAR selectStatement CLOSE_PAR) (COMMA (expression | NULL | OPEN_PAR selectStatement CLOSE_PAR))*
    CLOSE_PAR
    ;

openStatement
    : OPEN identifier (USING identifier (COMMA identifier)*)? (WITH REOPTIMIZATION)?
    ;

outputStatement
    : OUTPUT TO PIPE? identifier (WITHOUT HEADINGS)? selectStatement
    ;

prepareStatement
    : PREPARE identifier FROM (expression | identifier | quotedString)
    ;

putStatement
    : PUT identifier (FROM identifier (INDICATOR identifier | SYM1 identifier | SYM12 identifier)? (COMMA identifier (INDICATOR identifier | SYM1 identifier | SYM12 identifier)?)*
    | USING SQL DESCRIPTOR quotedString)?
    ;

releaseSavepoint
    : RELEASE SAVEPOINT savepointName=identifier
    ;

renameColumn
    : RENAME COLUMN oldColumn=columnName TO newColumn=columnName
    ;

renameConstraint
    : RENAME CONSTRAINT oldConstraint=identifier TO newConstraint=identifier
    ;

renameDatabase
    : RENAME DATABASE oldDatabase=identifier TO newDatabase=identifier
    ;

renameIndex
    : RENAME INDEX oldIndex=indexName TO newIndex=indexName
    ;

renameSecurity
    : RENAME SECURITY (POLICY | LABEL identifier | LABEL COMPONENT) old_name=identifier TO new_name=identifier
    ;

renameSequence
    : RENAME SEQUENCE oldSequence=identifier TO newSequence=identifier
    ;

renameTable
    : RENAME TABLE oldTableName=tableName TO newTableName=tableName
    ;

renameTrustedContext
    : RENAME TRUSTED CONTEXT oldTrustedContextName=identifier TO newTrustedContextName=identifier
    ;

renameUser
    : RENAME USER oldUserName=identifier TO newUserName=identifier
    ;

revokeStatement
    : REVOKE
    ;

revokeFragment
    :
    ;

rollbackWork
    : ROLLBACK WORK? (TO SAVEPOINT savepoint=identifier?)?
    ;

saveExternalDirectives
    : SAVE EXTERNAL DIRECTIVES optimizerDirectives (COMMA optimizerDirectives)* (ACTIVE | INACTIVE | TEST ONLY)
    FOR selectStatement
    ;

savepointStatement
    : SAVEPOINT savepoint=identifier UNIQUE?
    ;

selectStatement
    : SELECT selectOptions ((UNION ALL? | INTERSECT | MINUS | EXCEPT) SELECT selectOptions)*
    orderByClause? limitClause? (FOR READ ONLY | FOR UPDATE (OF column=identifier (COMMA column=identifier)?)?)?
    intoTableClause?
    ;

selectOptions
    : optimizerDirectives? projectionClause intoClause? fromClause gridClause? whereClause? hierarchicalClause? groupByClause? havingClause?
    ;

projectionClause
    : (SKIP_ numeric)? ((FIRST | LIMIT) numeric)? (DISTINCT | ALL | UNIQUE)? selectList (COMMA selectList)*
    ;

selectList
    : (expression | columnName) (AS? displayLabel=identifier)?
    | (tableName DOT)? (columnName (AS? alias)? | STAR)
//    | external=identifier DOT STAR
    | OPEN_PAR (collectionSubquery | selectStatement) CLOSE_PAR
    ;

collectionSubquery
    : MULTISET OPEN_PAR (SELECT ITEM)? selectStatement CLOSE_PAR
    ;

fromClause
    : FROM (OPEN_PAR? (ansiTables ansiJoins?)+ CLOSE_PAR?
    | (ansiTables | otherTables) (COMMA outerJoinClause)+
    | lateralDerivedTable
    )
    ;

ansiTables
    : (tableName | ONLY? OPEN_PAR tableName CLOSE_PAR) (AS? alias)?
    | collectionDerivedTable
    | iterator
    ;

iterator
    : TABLE OPEN_PAR (FUNCTION | PROCEDURE)? identifier OPEN_PAR routineParameterList? CLOSE_PAR CLOSE_PAR
    (AS? tableName OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR)?
    ;

ansiJoins
    : ((INNER | (LEFT | RIGHT | FULL) OUTER?)? JOIN ansiTables onCLause
    | CROSS JOIN ansiTables
    ) (AS? alias)?
    ;

onCLause
    : ON (columnName relationalOperator columnName | function | condition
    | OPEN_PAR selectStatement CLOSE_PAR | OPEN_PAR collectionSubquery CLOSE_PAR)
    ((AND | OR) (columnName relationalOperator columnName | function | condition
    | OPEN_PAR selectStatement CLOSE_PAR | OPEN_PAR collectionSubquery CLOSE_PAR))*
    ;

otherTables
    : external=identifier (AS? alias (OPEN_PAR derivedColumn=columnName (COMMA derivedColumn=columnName)* CLOSE_PAR)?)?
    | OPEN_PAR collectionSubquery CLOSE_PAR
    | OPEN_PAR selectStatement CLOSE_PAR
    | (tableName | ONLY? OPEN_PAR tableName CLOSE_PAR) AS? alias OPEN_PAR derivedColumn=columnName (COMMA derivedColumn=columnName)* CLOSE_PAR
    ;

outerJoinClause
    : OUTER (ansiTables | otherTables
    | OPEN_PAR ((ansiTables | otherTables) (COMMA (ansiTables | otherTables))* (COMMA outerJoinClause)+
    | outerJoinClause (COMMA outerJoinClause)* (COMMA (ansiTables | otherTables))+)
    (COMMA ((ansiTables | otherTables) (COMMA (ansiTables | otherTables))* (COMMA outerJoinClause)+
    | outerJoinClause (COMMA outerJoinClause)* (COMMA (ansiTables | otherTables))+) )*
    CLOSE_PAR
    )
    ;

lateralDerivedTable
    : LATERAL OPEN_PAR selectStatement CLOSE_PAR AS? alias (OPEN_PAR columnAlias=alias (COMMA columnAlias=alias)* CLOSE_PAR)?
    ;

gridClause
    : GRID ALL? identifier
    ;

whereClause
    : WHERE (condition | (columnName relationalOperator columnName) | function
    | OPEN_PAR selectStatement CLOSE_PAR | OPEN_PAR collectionSubquery CLOSE_PAR | identifier)
    (logicalOperator (condition | (columnName relationalOperator columnName) | function
    | OPEN_PAR selectStatement CLOSE_PAR | OPEN_PAR collectionSubquery CLOSE_PAR | identifier))*
    ;

hierarchicalClause
    : startWithClause? connectByClause
    ;

startWithClause
    : START WITH condition
    ;

connectByClause
    : CONNECT BY NOCYCLE? condition
    ;

groupByClause
    : GROUP BY (tableName DOT expression | columnName | colAlias=alias | numeric)
    (COMMA (tableName DOT expression | columnName | colAlias=alias | numeric))*
    ;

havingClause
    : HAVING condition
    ;

orderByClause
    : ORDER SIBLINGS? BY ((columnName | expression) substring?
    | selectNumber=numeric | displayLabel=identifier | caseExpression | ROWID //| olapWindowExpressions TODO
    ) (ASC | DESC)? (NULLS FIRST | NULLS LAST)?
    ;

substring
    : SYM2 numeric COMMA numeric SYM3
    ;

limitClause
    : LIMIT numeric
    ;

intoTableClause
    : INTO TEMP tableName (WITH NO LOG)?
    | intoExternalClause
    | INTO (RAW | STANDARD)? tableName storageOptions? (LOCK MODE (PAGE | ROW))?
    ;

intoExternalClause
    : INTO EXTERNAL tableName USING OPEN_PAR (tableOptions COMMA)? datafilesClause tableOptions? CLOSE_PAR
    ;

setAutofree
    : SET AUTOFREE (ENABLED | DISABLED)? (FOR (cursorId=identifier | cursorIdVar=anyName))?
    ;

setCollation
    : SET (COLLATION quotedString | NO COLLATION)
    ;

setConnection
    : SET CONNECTION (
    (identifier | quotedString | DEFAULT | databaseEnvironment) DORMANT?
    | CURRENT DORMANT
    )
    ;

setConstraints
    : SET CONSTRAINTS
    ((ALL | constraintName=identifier (COMMA constraintName=identifier)*) (DEFERRED | IMMEDIATE)?
    | (FOR tableName | constraintName=identifier (COMMA constraintName=identifier)*) constraintMode
    )
    ;

constraintMode
    : DISABLED
    | (ENABLED
    | FILTERING (WITHOUT ERROR | WITH ERROR)
    ) NOVALIDATE?
    ;

setDatabaseObjectMode
    : SET (objectListFormat| tableFormat)
    ;

objectListFormat
    : CONSTRAINTS constraintName=identifier (COMMA constraintName=identifier)* constraintsAndUniqueIndexesMode
    | INDEXES indexName (COMMA indexName)* (constraintsAndUniqueIndexesMode | triggersAndDuplicateIndexesMode)
    | TRIGGERS triggerName=identifier (COMMA triggerName=identifier)* triggersAndDuplicateIndexesMode
    ;

constraintsAndUniqueIndexesMode
    : DISABLED
    | (ENABLED
    | FILTERING (WITHOUT ERROR | WITH ERROR)
    ) NOVALIDATE?
    ;

triggersAndDuplicateIndexesMode
    : DISABLED
    | ENABLED
    ;

tableFormat
    : (CONSTRAINTS | INDEXES | TRIGGERS) (COMMA (CONSTRAINTS | INDEXES | TRIGGERS))*
    FOR tableName
    (constraintsAndUniqueIndexesMode | triggersAndDuplicateIndexesMode)
    ;

setDataskip
    : SET DATASKIP (ON (identifier (COMMA identifier)*)?
    | OFF
    | DEFAULT
    )
    ;

setDebugFile
    : SET DEBUG FILE TO (identifier | expression | quotedString) (WITH APPEND)?
    ;

setDeferredPrepare
    : SET DEFERRED_PREPARE (ENABLED | DISABLED)?
    ;

setDescriptor
    : SET DESCRIPTOR (identifier | quotedString)
    (COUNT ASSIGN (identifier | numeric)
    | VALUE (identifier | numeric) itemDescriptor (COMMA itemDescriptor)*
    )
    ;

itemDescriptor
    : (TYPE | LENGTH | PRECISION | SCALE | NULLABLE | INDICATOR | ITYPE | ILENGTH) ASSIGN (identifier | numeric)
    | (DATA | IDATA) ASSIGN (identifier | quotedString | literalDatetime | literalInterval | numeric)
    | (NAME | EXTYPENAME | EXTYPEOWNERNAME) ASSIGN (identifier | quotedString)
    | (SOURCEID | SOURCETYPE | EXTYPEID | EXTYPELENGTH | EXTYPEOWNERLENGTH) ASSIGN (identifier | numeric)
    ;

setEncryptionPassword
    : SET ENCRYPTION PASSWORD SYM6 expression SYM6 (WITH HINT SYM6 expression SYM6)?
    ;

setEnvironment
    : SET ENVIRONMENT (
    | (AUTOLOCATE | IFX_SESSION_LIMIT_LOCKS) quotedString
    | (AUTO_READAHEAD | IFX_BATCHEDREAD_INDEX | IFX_BATCHEDREAD_TABLE) quotedString (COMMA quotedString)?
    | (EXTDIRECTIVES | IFX_AUTO_REPREPARE | NOVALIDATE | USTLOW_SAMPLE) (ON | OFF | quotedString | DEFAULT)
    | (AUTO_STAT_MODE | BOUND_IMPL_PDQ | IMPLICIT_PDQ) (ON | OFF | quotedString)
    | CLUSTER_TXN_SCOPE (quotedString | DEFAULT)
    | DEFAULTESCCHAR quotedString
    | FORCE_DDL_EXEC (ON | OFF | quotedString)
    | GRID_NODE_SKIP (ON | OFF | DEFAULT)
    | HDR_TXN_SCOPE quotedString
    | (INFORMIXCONRETRY | INFORMIXCONTIME) quotedString?
    | (OPTCOMPIND | STATCHANGE) (DEFAULT | quotedString)
    | (RETAINUPDATELOCKS | USELASTCOMMITTED) quotedString
    | (SELECT_GRID | SELECT_GRID_ALL) (quotedString | DEFAULT)
    | USE_DWA sessionEnvironmentOptions
    | USE_SHARDING (ON | OFF)
    )
    ;

sessionEnvironmentOptions
    : SESSION numeric
    ((ACCELERATE | DEBUG | FALLBACK | UNIQUECHECK) (ON | OFF)?
    | DEBUG FILE identifier?
    | PROBE (STOP | START)?
    )
    | PROBE CLEANUP
    ;

setExplain
    : SET EXPLAIN (ON AVOID_EXECUTE? | OFF | FILE TO (identifier | expression | quotedString))
    ;

setIndexes
    : SET INDEXES (indexName (COMMA indexName)* | FOR tableName)
    (DISABLED | ENABLED | FILTERING (WITHOUT ERROR | WITH ERROR))
    ;

setIsolation
    : SET ISOLATION TO?
    ( REPEATABLE READ
    | (COMMITTED READ (LAST COMMITTED)?
    | CURSOR STABILITY
    | DIRTY READ (WITH WARNING)?
    ) (RETAIN UPDATE LOCKS)?
    )
    ;

setLockMode
    : SET LOCK MODE TO (WAIT numeric | NOWAIT)
    ;

setLog
    : SET BUFFERED? LOG
    ;

setOptimization
    : SET OPTIMIZATION (HIGH | LOW | FIRST_ROWS | ALL_ROWS | environmentOptions)
    ;

environmentOptions
    : ENVIRONMENT (STAR_JOIN SYM6 (ENABLED | DISABLED | FORCED) SYM6
    | (FACT | AVOID_FACT | NON_DIM) (SYM6 tableName (COMMA tableName)* SYM6 | DEFAULT | SYM7))
    ;

setPdqpriority
    : SET PDQPRIORITY (DEFAULT | numeric | LOW | OFF | HIGH)
    ;

setRole
    : SET ROLE (NULL | NONE | roleName | quotedString | DEFAULT)
    ;

setSessionAuthorization
    : SET SESSION AUTHORIZATION TO quotedString (USING quotedString)?
    ;

setStatementCache
    : SET STATEMENT CACHE (ON | OFF)
    ;

setTransaction
    : SET TRANSACTION (READ WRITE | READ ONLY | ISOLATION LEVEL (READ COMMITTED | REPEATABLE READ | SERIALIZABLE | READ UNCOMMITTED))
    ;

setTransactionMode
    : SET CONSTRAINTS (ALL | (identifier (COMMA identifier)*)) (DEFERRED | IMMEDIATE)
    ;

setTriggers
    : SET TRIGGERS (trigger=identifier (COMMA trigger=identifier)* | FOR tableName) (ENABLED | DISABLED)?
    ;

setUserPassword
    : SET USER PASSWORD OLD old_password=quotedString NEW new_password=quotedString
    ;

startViolationsTable
    : START VIOLATIONS TABLE FOR tableName (USING identifier COMMA identifier)? (MAX ROWS numeric)?
    ;

stopViolationsTable
    : STOP VIOLATIONS TABLE FOR tableName
    ;

truncateStatement
    : TRUNCATE TABLE? tableName (DROP STORAGE | REUSE STORAGE KEEP STATISTICS)?
    ;

unloadStatement
    : UNLOAD TO filename=quotedString (DELIMITER delimiter=quotedString)? (selectStatement | identifier)
    ;

unlockTable
    : UNLOCK TABLE tableName
    ;

updateStatement
    : UPDATE (optimizerDirectives? updateTarget setClause whereOptions
    | collectionDerivedTable setClause (WHERE CURRENT OF identifier)?
    )
    ;

updateStatistics
    : UPDATE STATISTICS ((LOW? tableAndColumnScope (DROP DISTRIBUTIONS ONLY?)?
    | (MEDIUM | HIGH) tableAndColumnScope resolutionClause?) (FORCE | AUTO)
    | routineStatistics
    )
    ;

tableAndColumnScope
    : (FOR TABLE ((tableName | ONLY OPEN_PAR tableName CLOSE_PAR) (OPEN_PAR columnName (COMMA columnName)* CLOSE_PAR)?)?)?
    ;

resolutionClause
    : resolutionClauseMedium
    | resolutionClauseHigh
    ;

resolutionClauseMedium
    : (SAMPLING SIZE numeric) (RESOLUTION numeric numeric?)? (DISTRIBUTIONS ONLY)?
    | (SAMPLING SIZE numeric)? (RESOLUTION numeric numeric?) (DISTRIBUTIONS ONLY)?
    | (SAMPLING SIZE numeric)? (RESOLUTION numeric numeric?)? (DISTRIBUTIONS ONLY)
    ;

resolutionClauseHigh
    : (RESOLUTION numeric) (DISTRIBUTIONS ONLY)?
    | (RESOLUTION numeric)? (DISTRIBUTIONS ONLY)
    ;

routineStatistics
    : FOR ((ROUTINE | PROCEDURE | FUNCTION) (identifier (OPEN_PAR routineParameterList? CLOSE_PAR)?)?
    | SPECIFIC (ROUTINE | PROCEDURE | FUNCTION) identifier
    )
    ;

updateTarget
    : tableName (AS? alias)?
    | ONLY OPEN_PAR tableName CLOSE_PAR
    ;

whereOptions
    : (WHERE condition
    | WHERE CURRENT OF identifier
    )?
    ;

wheneverStatement
    : WHENEVER (SQLERROR | ERROR | NOT FOUND | SQLWARNING) (CONTINUE | STOP | CALL routineName | (GOTO | GO TO) SYM1? identifier)
    ;

constExprssion
    : quotedString
    | numeric
    | USER
    | CURRENT_USER
    | CURRENT_ROLE
    | DEFAULT_ROLE
    | anyName
    | TODAY
    | (CURRENT | SYSDATE) numeric?
    | literalDatetime
    | literalInterval
    | numeric UNITS timeUnit
    | identifier DOT (CURRVAL | NEXTVAL)
    | literalCollection
    | literalRow
    ;

literalDatetime
    : DATETIME OPEN_PAR IDENTIFIER CLOSE_PAR datetimeField
    ;

literalInterval
    : INTERVAL OPEN_PAR IDENTIFIER CLOSE_PAR intervalField
    ;

literalCollection
    : (SET | MULTISET | LIST) SYM4 literalData SYM5
    ;

literalData
    : (elementLiteralValue | nestedQuotationMarks literalCollection nestedQuotationMarks) (COMMA (elementLiteralValue | nestedQuotationMarks literalCollection nestedQuotationMarks))*
    ;

elementLiteralValue
    : quotedString
    | literalDatetime
    | numeric
    | literalInterval
    | literalRow
    ;

nestedQuotationMarks
    : QUOTA+
    ;

literalRow
    : SYM6 ROW OPEN_PAR fieldLiteralValue (COMMA fieldLiteralValue)* CLOSE_PAR SYM6
    | ROW OPEN_PAR fieldLiteralValue (COMMA fieldLiteralValue)* CLOSE_PAR
    ;

fieldLiteralValue
    : quotedString
    | numeric
    | USER
    | literalDatetime
    | literalInterval
    | literalCollection
    | ROW OPEN_PAR fieldLiteralValue CLOSE_PAR
    ;

expression
    : (PLUS | MINUS)? (castExpression | columnExpression | conditionalExpression | constExprssion | constructorExpression
    | function | identifier | aggregateExpression | NULL | OPEN_PAR expression CLOSE_PAR)
    (binaryOperator
    (PLUS | MINUS)? (castExpression | columnExpression | conditionalExpression | constExprssion | constructorExpression
    | function | identifier | aggregateExpression | NULL | OPEN_PAR expression CLOSE_PAR)
    )*
    ;

binaryOperator
    : PLUS
    | MINUS
    | STAR
    | DIV
    | PIPE2
    ;

castExpression
    : CAST OPEN_PAR expression (SYM8 dataType)* AS dataType CLOSE_PAR
//    | expression (SYM8 dataType)+
    ;

columnExpression
    : tableName DOT (columnName (SYM2 numeric COMMA numeric SYM3)? | ROWID | rowColumn=columnName (DOT STAR | (DOT identifier)+)?)
//    | expression (DOT STAR | (DOT identifier)+)
    ;

conditionalExpression
    : caseExpression
    | function
    ;

caseExpression
    : genericCaseExpression
    | linearCaseExpression
    ;

genericCaseExpression
    : CASE (WHEN condition THEN (expression | NULL))+ (ELSE (expression | NULL))? END
    ;

condition
    : NOT? (comparisionCondition | selectStatement | function)
    (logicalOperator NOT? (comparisionCondition | selectStatement | function))*
    ;

logicalOperator
    : AND | OR
    ;

comparisionCondition
    : expression relationalOperator expression
    | expression NOT? BETWEEN expression AND expression
    | inCondition
    | (identifier | expression) IS NOT? NULL
    | DELETING | INSERTING | SELECTING | UPDATING
    | identifier NOT? (LIKE | MATCHES) identifier (ESCAPE quotedString)?
    ;

inCondition
    : expression NOT? IN (OPEN_PAR (numeric | literalDatetime | identifier | literalInterval | USER | CURRENT_USER | TODAY | CURRENT datetimeField? | literalRow) CLOSE_PAR
    | identifier
    | OPEN_PAR literalCollection (COMMA literalCollection)* CLOSE_PAR
    | literalCollection
    )
    ;

constructorExpression
    : ROW OPEN_PAR expression (COMMA expression)* CLOSE_PAR
    | collectionConstructors
    ;

collectionConstructors
    : (SET | MULTISET | LIST) SYM4 (expression (COMMA expression)*)? SYM5
    ;

aggregateExpression
    : COUNT OPEN_PAR STAR CLOSE_PAR
    | ( AVG | COUNT | MAX | MIN | SUM | RANGE | STDEV | VARIANCE) (OPEN_PAR (aggregateScope | ALL? expression) CLOSE_PAR)?
    ;

aggregateScope
    : (ALL | DISTINCT | UNIQUE)? identifier
    ;

relationalOperator
    : LT
    | LT_EQ
    | GT
    | GT_EQ
    | ASSIGN
    | EQ
    | NOT_EQ1
    | NOT_EQ2
    ;

linearCaseExpression
    : CASE expression (WHEN expression THEN (expression | NULL))+ (ELSE (expression | NULL))? END
    ;

function
    : functionName OPEN_PAR
    ((identifier ASSIGN)? expression (COMMA (identifier ASSIGN)? expression)*)?
    (COMMA variableDeclaration)?
    CLOSE_PAR
    ;

variableDeclaration
    : identifier SYM13 (builtInDataTypes | identifier | complexDataType)
    ;

builtInDataTypes
    : BOOLEAN
    | JSON
    | BSON
    | characterDataType
    | numericDataType
    | largeObjectDataType
    | timeDataType
    | IDSSECURITYLABEL
    ;

numericDataType
    : exactNumericDataType
    | approximateNumericDataType
    ;

exactNumericDataType
    : BIGINT
    | INT
    | INTEGER
    | INT8
    | SMALLINT
    | (BIGSERIAL | SERIAL | SERIAL8) (OPEN_PAR numeric CLOSE_PAR)?
    | (MONEY | DECIMAL | DEC | NUMERIC) (OPEN_PAR numeric (COMMA numeric)? CLOSE_PAR)
    ;

approximateNumericDataType
    : (DECIMAL | DEC | NUMERIC) (OPEN_PAR numeric CLOSE_PAR)?
    | (FLOAT | DOUBLE PRECISION) (OPEN_PAR numeric CLOSE_PAR)?
    | SMALLFLOAT
    | REAL
    ;

characterDataType
    : (CHAR | CHARACTER | NCHAR) OPEN_PAR numeric CLOSE_PAR
    | (NVARCHAR | VARCHAR | CHARACTER VARYING) OPEN_PAR numeric (COMMA numeric)? CLOSE_PAR
    | LVARCHAR OPEN_PAR numeric CLOSE_PAR
    ;

largeObjectDataType
    : BLOB
    | CLOB
    | (TEXT | BYTE) (IN (TABLE | identifier))?
    ;

timeDataType
    : DATE
    | INTERVAL intervalField
    | DATETIME datetimeField
    ;

complexDataType
    : createRowType
    | collectionDataTypes
    ;

collectionDataTypes
    : COLLECTION
    | (SET | MULTISET | LIST) OPEN_PAR (dataTypeName | OPEN_PAR dataTypeName NOT NULL CLOSE_PAR) NOT NULL CLOSE_PAR
    ;

dataTypeName
    : identifier
    ;

roleName
    : anyName
    ;

alias
    : identifier
    ;

tableName
    : identifier
    ;

columnName
    : identifier
    ;

routineName
    : identifier
    ;

procedureName
    : identifier
    ;

functionName
    : identifier
    ;

viewName
    : identifier
    ;

indexName
    : identifier
    ;

userName
    : anyName
    ;

anyName
    : IDENTIFIER
    | keyword
    | STRING_LITERAL
    | OPEN_PAR anyName CLOSE_PAR
;

identifier
    : anyName (DOT anyName)*
    ;

numeric
    : NUMERIC_LITERAL
    ;

quotedString
    : QUOTED_STRING
    ;

datetimeField
    : YEAR TO (YEAR | MONTH | DAY | HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | MONTH TO (MONTH | DAY | HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | DAY TO (DAY | HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | HOUR TO (HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | MINUTE TO (MINUTE | SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | SECOND TO (SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | FRACTION TO FRACTION (OPEN_PAR numeric CLOSE_PAR)?
    ;

intervalField
    : DAY (OPEN_PAR numeric CLOSE_PAR)? TO (DAY | HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | HOUR (OPEN_PAR numeric CLOSE_PAR)? TO (HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | MINUTE (OPEN_PAR numeric CLOSE_PAR)? TO (MINUTE | SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | SECOND (OPEN_PAR numeric CLOSE_PAR)? TO (SECOND | FRACTION (OPEN_PAR numeric CLOSE_PAR)?)
    | FRACTION TO FRACTION (OPEN_PAR numeric CLOSE_PAR)?
    | YEAR (OPEN_PAR numeric CLOSE_PAR)? TO (YEAR | MONTH)
    | MONTH (OPEN_PAR numeric CLOSE_PAR)? TO MONTH
    ;

timeUnit
    : YEAR
    | MONTH
    | DAY
    | HOUR
    | MINUTE
    | SECOND
    | FRACTION
    ;

keyword
    : AAO
    | ABORT
    | ABSOLUTE
    | ACCELERATE
    | ACCESS
    | ACCESS_METHOD
    | ACCOUNT
    | ACTION
    | ACTIVE
    | ADD
    | ADDRESS
    | AFTER
    | AGGREGATE
    | ALIGNMENT
    | ALL
    | ALLOCATE
    | ALL_ROWS
    | ALTER
    | ANALYZE
    | AND
    | ANSI
    | ANY
    | APPEND
    | ARRAY
    | AS
    | ASC
    | AT
    | ATTACH
    | ATTRIBUTES
    | AUDIT
    | AUTHENTICATION
    | AUTHID
    | AUTHORIZATION
    | AUTHORIZED
    | AUTO
    | AUTOFREE
    | AUTOINCREMENT
    | AUTOLOCATE
    | AUTO_READAHEAD
    | AUTO_STAT_MODE
    | AVG
    | AVOID_EXECUTE
    | AVOID_FACT
    | BARGROUP
    | BASED
    | BEFORE
    | BEGIN
    | BETWEEN
    | BIGINT
    | BIGSERIAL
    | BLOB
    | BLOBDIR
    | BUFFERED
    | BOUND_IMPL_PDQ
    | BOOLEAN
    | BSON
    | BSON_GET
    | BSON_VALUE_BIGINT
    | BSON_VALUE_BOOLEAN
    | BSON_VALUE_DATE
    | BSON_VALUE_DOUBLE
    | BSON_VALUE_INT
    | BSON_VALUE_LVARCHAR
    | BSON_VALUE_OBJECTID
    | BSON_VALUE_TIMESTAMP
    | BSON_VALUE_VARCHAR
    | BUCKETS
    | BUFFERED
    | BY
    | BYTE
    | C
    | CACHE
    | CALL
    | CANNOTHASH
    | CASCADE
    | CASE
    | CAST
    | CHAR
    | CHARACTER
    | CHECK
    | CLASS
    | CLEANUP
    | CLOB
    | CLOBDIR
    | CLOSE
    | CLUSTER
    | CLUSTER_TXN_SCOPE
    | COLLATE
    | COLLATION
    | COLLECTION
    | COLUMN
    | COLUMNS
    | COMBINE
    | COMMIT
    | COMMITTED
    | COMPONENT
    | COMPRESSED
    | CONCURRENT
    | CONFLICT
    | CONNECT
    | CONNECTION
    | CONSTRAINT
    | CONSTRAINTS
    | CONTEXT
    | COSTFUNC
    | CONTINUE
    | COUNT
    | CRCOLS
    | CREATE
    | CROSS
    | CURRENT_DATE
    | CURRENT_ROLE
    | CURRENT_TIME
    | CURRENT_TIMESTAMP
    | CURRENT_USER
    | CURRVAL
    | CURSOR
    | CYCLE
    | DATA
    | DATABASE
    | DATAFILES
    | DATASKIP
    | DATE
    | DATETIME
    | DAY
    | DBDATE
    | DBMONEY
    | DBSA
    | DBSECADM
    | DBSSO
    | DEALLOCATE
    | DEBUG
    | DEC
    | DECIMAL
    | DECLARE
    | DEFAULT
    | DEFAULTESCCHAR
    | DEFAULT_ROLE
    | DEFERRABLE
    | DEFERRED
    | DEFERRED_PREPARE
    | DEFINE
    | DELETE
    | DELETING
    | DELIMITED
    | DELIMITER
    | DELUXE
    | DESC
    | DESCRIBE
    | DESCRIPTOR
    | DETACH
    | DIAGNOSTICS
    | DIRECTIVES
    | DIRTY
    | DISABLE
    | DISABLED
    | DISCARD
    | DISCONNECT
    | DISK
    | DISTINCT
    | DISTRIBUTIONS
    | DOCUMENT
    | DORMANT
    | DOUBLE
    | DROP
    | EACH
    | ELIF
    | ELSE
    | ENABLE
    | ENABLED
    | ENCRYPTION
    | END
    | ENVIRONMENT
    | ERKEY
    | ERROR
    | ESCAPE
    | EXCEPT
    | EXCEPTION
    | EXCLUSIVE
    | EXECUTE
    | EXEMPTION
    | EXISTS
    | EXIT
    | EXPLAIN
    | EXPLICIT
    | EXPRESS
    | EXPRESSION
    | EXTDIRECTIVES
    | EXTENT
    | EXTERNAL
    | EXTYPEID
    | EXTYPELENGTH
    | EXTYPENAME
    | EXTYPEOWNERLENGTH
    | EXTYPEOWNERNAME
    | FACT
    | FAIL
    | FALLBACK
    | FETCH
    | FILE
    | FILLFACTOR
    | FILTERING
    | FINAL
    | FIRST_ROWS
    | FIXED
    | FLOAT
    | FLUSH
    | FORCE
    | FORCED
    | FORCE_DDL_EXEC
    | FOUND
    | FOR
    | FOREACH
    | FOREIGN
    | FORMAT
    | FRACTION
    | FRAGMENT
    | FRAGMENTS
    | FREE
    | FROM
    | FULL
    | FUNCTION
    | GET
    | GLOB
    | GO
    | GOTO
    | GRANT
    | GRID
    | GRID_NODE_SKIP
    | GROUP
    | HANDLESNULLS
    | HASH
    | HAVING
    | HDR_TXN_SCOPE
    | HEADINGS
    | HIGH
    | HINT
    | HOLD
    | HOME
    | HOUR
    | IDATA
    | IDSLBACREADARRAY
    | IDSLBACREADTREE
    | IDSLBACREADSET
    | IDSLBACRULES
    | IDSLBACWRITEARRAY
    | IDSLBACWRITESET
    | IDSLBACWRITETREE
    | IDSSECURITYLABEL
    | IF
    | IFX_AUTO_REPREPARE
    | IFX_BATCHEDREAD_INDEX
    | IFX_BATCHEDREAD_TABLE
    | IFX_SESSION_LIMIT_LOCKS
    | IGNORE
    | ILENGTH
    | IMMEDIATE
    | IMPLICIT
    | IMPLICIT_PDQ
    | IN
    | INACTIVE
    | INCREMENT
    | INDEX
    | INDEXED
    | INDEXES
    | INDICATOR
    | INFO
    | INFORMIX
    | INFORMIXCONRETRY
    | INFORMIXCONTIME
    | INIT
    | INITIALLY
    | INNER
    | INOUT
    | INPUT
    | INSENSITIVE
    | INSERT
    | INSERTING
    | INSTEAD
    | INT
    | INT8
    | INTEGER
    | INTEG
    | INTERNAL
    | INTERNALLENGTH
    | INTERVAL
    | INTERSECT
    | INTO
    | ISOLATION
    | ITER
    | ITERATOR
    | ITEM
    | ITYPE
    | IS
    | ISNULL
    | JAVA
    | JOIN
    | JSON
    | KEEP
    | KEY
    | LABEL
    | LATERAL
    | LANGUAGE
    | LEFT
    | LENGTH
    | LET
    | LEVEL
    | LIKE
    | LIMIT
    | LINK
    | LIST
    | LISTING
    | LOAD
    | LOCK
    | LOCKS
    | LOG
    | LOOP
    | LOW
    | LVARCHAR
    | MATCH
    | MATCHED
    | MATCHES
    | MAX
    | MAXLEN
    | MAXERRORS
    | MAXVALUE
    | MEDIUM
    | MERGE
    | MIN
    | MINUTE
    | MINVALUE
    | MODE
    | MODERATE
    | MODIFY
    | MONEY
    | MONTH
    | MULTISET
    | NAME
    | NATURAL
    | NCHAR
    | NEGATOR
    | NEXT
    | NEXTVAL
    | NEW
    | NLSCASE
    | NO
    | NOCACHE
    | NOCYCLE
    | NOMAXVALUE
    | NOMINVALUE
    | NON_DIM
    | NONE
    | NOORDER
    | NOT
    | NOTNULL
    | NOVALIDATE
    | NOWAIT
    | NULL
    | NULLABLE
    | NUMERIC
    | NUMROWS
    | NVARCHAR
    | OF
    | OFF
    | OFFSET
    | OLD
    | ON
    | ONLINE
    | ONLY
    | OPAQUE
    | OPCLASS
    | OPEN
    | OPTCOMPIND
    | OPTIMIZATION
    | OPTION
    | OR
    | ORDER
    | OUT
    | OUTER
    | OUTPUT
    | OVERRIDE
    | PAGE
    | PARALLELIZABLE
    | PARAMETER
    | PASSEDBYVALUE
    | PASSWORD
    | PDQPRIORITY
    | PERCALL_COST
    | PLAN
    | POLICY
    | PRAGMA
    | PRECISION
    | PREPARE
    | PRIMARY
    | PRIOR
    | PRIVATE
    | PRIVILEGE
    | PRIVILEGES
    | PROBE
    | PROCEDURE
    | PROPERTIES
    | PUBLIC
    | PUT
    | QUERY
    | RAISE
    | RAW
    | READ
    | REAL
    | RECORDEND
    | RECURSIVE
    | REFERENCING
    | REFERENCE
    | REFERENCES
    | REGEXP
    | REINDEX
    | REJECTFILE
    | RELATIVE
    | RELEASE
    | REMAINDER
    | RENAME
    | REOPTIMIZATION
    | REPEATABLE
    | REPLACE
    | REPLCHECK
    | REPLICATION
    | RESTART
    | RESTRICT
    | RESOURCE
    | RESOLUTION
    | RESUME
    | RETAIN
    | RETAINUPDATELOCKS
    | RETURN
    | RETURNING
    | RETURNS
    | REUSE
    | REVOKE
    | RIGHT
    | ROBIN
    | ROLLBACK
    | ROOT
    | ROUND
    | ROUTINE
    | ROW
    | ROWID
    | ROWIDS
    | ROLE
    | ROLLING
    | ROWS
    | RULE
    | SAMEAS
    | SAMPLING
    | SAVE
    | SAVEPOINT
    | SCALE
    | SCROLL
    | SECOND
    | SECONDARY
    | SECURED
    | SECURITY
    | SELCONST
    | SELECT
    | SELECT_GRID
    | SELECT_GRID_ALL
    | SELECTING
    | SELFUNC
    | SENSITIVE
    | SEQUENCE
    | SERIALIZABLE
    | SESSION
    | SET
    | SETSESSIONAUTH
    | SCHEMA
    | SHARE
    | SIBLINGS
    | SIZE
    | SKIP_
    | SMALLFLOAT
    | SMALLINT
    | SOURCEID
    | SOURCETYPE
    | SPECIFIC
    | SPL
    | SQL
    | SQLERROR
    | SQLWARNING
    | STABILITY
    | STACK
    | STANDARD
    | STAR_JOIN
    | START
    | STATCHANGE
    | STATEMENT
    | STATISTICS
    | STATLEVEL
    | STATUS
    | STDEV
    | STEP
    | STOP
    | STORAGE
    | STORE
    | STRATEGIES
    | STYLE
    | SUM
    | SUPPORT
    | SYNONYM
    | SYSDATE
    | SYSTEM
    | TABLE
    | TABLES
    | TRACE
    | TEMP
    | TEMPORARY
    | TEST
    | TEXT
    | THEN
    | TIME
    | TO
    | TODAY
    | TRANSACTION
    | TRANSITION
    | TREE
    | TRIGGER
    | TRIGGERS
    | TRUNCATE
    | TRUSTED
    | TYPE
    | UID
    | UNCOMMITTED
    | UNDER
    | UNION
    | UNIQUE
    | UNIQUECHECK
    | UNITS
    | UNLOAD
    | UNLOCK
    | UPDATE
    | UPDATING
    | UPON
    | USAGE
    | USE
    | USE_DWA
    | USELASTCOMMITTED
    | USER
    | USE_SHARDING
    | USING
    | USTLOW_SAMPLE
    | VACUUM
    | VALUE
    | VALUES
    | VARCHAR
    | VARIABLE
    | VARIANCE
    | VARIANT
    | VARYING
    | VERCOLS
    | VIEW
    | VIOLATIONS
    | VIRTUAL
    | WAIT
    | WARNING
    | WHEN
    | WHENEVER
    | WHERE
    | WHILE
    | WITH
    | WITHOUT
    | WORK
    | WRITE
    | WRITEDOWN
    | WRITEUP
    | XADATASOURCE
    | YEAR
    | FIRST_VALUE
    | OVER
    | PARTITION
    | RANGE
    | PRECEDING
    | UNBOUNDED
    | CURRENT
    | FOLLOWING
    | CUME_DIST
    | DENSE_RANK
    | LAG
    | LAST_VALUE
    | LEAD
    | NTH_VALUE
    | NTILE
    | PERCENT_RANK
    | RANK
    | ROW_NUMBER
    | GENERATED
    | ALWAYS
    | STORED
    | TRUE
    | FALSE
    | WINDOW
    | NULLS
    | FIRST
    | LAST
    | FILTER
    | GROUPS
    | EXCLUDE
;