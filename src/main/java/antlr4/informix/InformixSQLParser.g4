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
    : (unitStatement | splStatement)* EOF
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
//    | setConnection
//    | setConstraints
//    | setDatabaseObjectMode
//    | setDataskip
//    | setDebugFile
//    | setDeferredPrepare
//    | setDescriptor
//    | setEncryption
//    | setEnvironment
//    | setExplain
//    | setIndexes
//    | setIsolation
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
    : CALL (procedure=identifier OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR
    | functionName=identifier OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR RETURNING identifier (COMMA identifier)*
    | routineName=identifier (RETURNING identifier (COMMA identifier)*)?
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
    : EXECUTE (PROCEDURE | FUNCTION) identifier OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR
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
    (functionName=identifier OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR
    | expression
    | OPEN_PAR selectStatement (COMMA selectStatement)* CLOSE_PAR
    )
    (COMMA (functionName=identifier OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR
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
    : ALLOCATE DESCRIPTOR anyName (WITH MAX (NUMERIC_LITERAL | identifier))?
    ;

allocateRow
    : ALLOCATE ROW SYM1 identifier
    ;

alterAccessMethod
    : ALTER ACCESS_METHOD identifier ((MODIFY | ADD) purposeOptions | DROP identifier)
    (COMMA (MODIFY | ADD) purposeOptions| DROP identifier)*
    ;

purposeOptions
    : identifier (ASSIGN (NUMERIC_LITERAL | QUOTED_STRING | identifier))?
    ;

alterFragment
    : ALTER FRAGMENT (
    ON TABLE identifier (attachClause | detachClause | initClause | addClause | dropClause | modifyClause)
    | ON INDEX identifier (initClause | addClause | dropClause | modifyClause)
    )
    ;

attachClause
    : ATTACH identifier asClause? (COMMA identifier asClause?)*
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
    : DETACH PARTITION? identifier identifier
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
    : ROLLING OPEN_PAR NUMERIC_LITERAL FRAGMENTS CLOSE_PAR (DETACH | DISCARD)
    | (ROLLING OPEN_PAR NUMERIC_LITERAL FRAGMENTS)? LIMIT TO NUMERIC_LITERAL IDENTIFIER (DETACH | DISCARD) (INTERVAL FIRST | ANY | INTERVAL ONLY)?
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
    : ALTER (FUNCTION identifier OPEN_PAR dataTypeName (COMMA dataTypeName)* CLOSE_PAR
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
    | PERCALL_COST ASSIGN NUMERIC_LITERAL
    | COSTFUNC ASSIGN identifier
    | SELFUNC ASSIGN identifier
    | SELCONST ASSIGN NUMERIC_LITERAL
    | STACK ASSIGN NUMERIC_LITERAL
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
    : ALTER INDEX identifier TO NOT? CLUSTER
    ;

alterProcedure
    : ALTER (PROCEDURE identifier OPEN_PAR dataTypeName (COMMA dataTypeName)* CLOSE_PAR
    | SPECIFIC PROCEDURE identifier
    ) WITH OPEN_PAR ((ADD | MODIFY | DROP) routineModifier | MODIFY EXTERNAL NAME ASSIGN sharedObjectFilename=identifier) CLOSE_PAR
    ;

alterRoutine
    : ALTER (ROUTINE identifier OPEN_PAR dataTypeName (COMMA dataTypeName)* CLOSE_PAR
    | SPECIFIC ROUTINE identifier
    ) WITH OPEN_PAR ((ADD | MODIFY | DROP) routineModifier | MODIFY EXTERNAL NAME ASSIGN sharedObjectFilename=identifier) CLOSE_PAR
    ;

alterSecurityLabelComponent
    : ALTER SECURITY LABEL COMPONENT identifier ADD (ARRAY SYM2 QUOTED_STRING (COMMA QUOTED_STRING)* (BEFORE | AFTER) QUOTED_STRING (COMMA QUOTED_STRING (COMMA QUOTED_STRING)* (BEFORE | AFTER) QUOTED_STRING)* SYM3
    | SET SYM4 QUOTED_STRING (COMMA QUOTED_STRING)* SYM5
    | TREE OPEN_PAR QUOTED_STRING UNDER QUOTED_STRING (COMMA QUOTED_STRING UNDER QUOTED_STRING)* CLOSE_PAR
    )
    ;

alterSequence
    : ALTER SEQUENCE identifier (
    | CYCLE | NOCYCLE
    | CACHE NUMERIC_LITERAL | NOCACHE
    | ORDER | NOORDER
    | INCREMENT BY? NUMERIC_LITERAL
    | RESTART WITH? NUMERIC_LITERAL
    | MINVALUE NUMERIC_LITERAL | NOMINVALUE
    | MAXVALUE NUMERIC_LITERAL | NOMAXVALUE
    )
    ;

alterTable
    : ALTER TABLE identifier (basicTableOption | loggingTypeOption | addTypeClause | statisticsOptions)
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
    : ADD newColumn (COMMA newColumn)*
    ;

newColumn
    : new_column=identifier dataTypeName (defaultClause | singleColumnConstraint)* (BEFORE identifier)? (COLUMN? SECURED WITH identifier)?
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
    ((NOT NULL | NULL | UNIQUE | DISTINCT | PRIMARY KEY | referencesClause) OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR
    | checkClause
    | foreignKeyDefinition
    ) constraintDefinition?
    | foreignKeyDefinition (CONSTRAINT identifier)? INDEX DISABLED
    )
    ;

foreignKeyDefinition
    : FOREIGN KEY OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR referencesClause
    ;

addOrDropSpecializedColumns
    : (ADD | DROP) (CRCOLS | ERKEY | REPLCHECK | ROWIDS | VERCOLS)
    ;

dropConstraintClause
    : DROP CONSTRAINT OPEN_PAR? identifier (COMMA identifier)* CLOSE_PAR?
    ;

dropColumnClause
    : DROP identifier (COMMA identifier)*
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
    (OPEN_PAR (COMMA? (EXTENT SIZE NUMERIC_LITERAL | NO LOG | LOG | HIGH INTEG | MODERATE INTEG | NO? KEEP ACCESS TIME))* CLOSE_PAR)?
    ;

securityPolicyClause
    : ADD SECURITY POLICY identifier
    | DROP SECURITY POLICY
    ;

loggingTypeOption
    : TYPE OPEN_PAR (STANDARD | RAW) CLOSE_PAR
    ;

addTypeClause
    : ADD TYPE identifier
    ;

statisticsOptions
    : (STATCHANGE (AUTO | NUMERIC_LITERAL)) (STATLEVEL (FRAGMENT | TABLE | AUTO))?
    | (STATCHANGE (AUTO | NUMERIC_LITERAL))? (STATLEVEL (FRAGMENT | TABLE | AUTO))
    ;

alterTrustedContext
    : ALTER TRUSTED CONTEXT identifier ( ALTER (SYSTEM AUTHID identifier | ATTRIBUTES OPEN_PAR ADDRESS identifier (COMMA ADDRESS identifier)* CLOSE_PAR | NO DEFAULT ROLE | DEFAULT ROLE identifier | DISABLE | ENABLE)+
    | (ADD | DROP) ATTRIBUTES OPEN_PAR ADDRESS identifier (COMMA ADDRESS identifier)* CLOSE_PAR
    | (ADD | REPLACE) USE FOR authorizedUserClause (COMMA authorizedUserClause)*
    | DROP USE FOR (identifier | PUBLIC) (COMMA (identifier | PUBLIC))*
    )+
    ;

authorizedUserClause
    : (identifier (ROLE identifier)? | PUBLIC) (WITH AUTHENTICATION | WITHOUT AUTHENTICATION)?
    ;

alterUser
    : ALTER (DEFAULT USER | USER identifier (LOCK | ACCOUNT UNLOCK)?)
    (COMMA? ((ADD | MODIFY) (PASSWORD QUOTED_STRING | UID NUMERIC_LITERAL | GROUP OPEN_PAR (identifier | NUMERIC_LITERAL) (COMMA (identifier | NUMERIC_LITERAL))* CLOSE_PAR
    | USER identifier | AUTHORIZATION OPEN_PAR (DBSA | DBSSO | AAO | BARGROUP) (COMMA (DBSA | DBSSO | AAO | BARGROUP))* CLOSE_PAR | HOME QUOTED_STRING)
    | DROP (PASSWORD | UID | GROUP OPEN_PAR (identifier | NUMERIC_LITERAL) (COMMA (identifier | NUMERIC_LITERAL))* CLOSE_PAR
    | USER | AUTHORIZATION OPEN_PAR (DBSA | DBSSO | AAO | BARGROUP) (COMMA (DBSA | DBSSO | AAO | BARGROUP))* CLOSE_PAR | HOME)
    )
    )+
    ;

beginWork
    : BEGIN WORK? (WITHOUT REPLICATION)?
    ;

closeStatement
    : CLOSE identifier
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
    : USER QUOTED_STRING USING QUOTED_STRING
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
    : CREATE (EXPLICIT | IMPLICIT)? CAST (IF NOT EXISTS)? OPEN_PAR dataTypeName AS dataTypeName (WITH FUNCTION identifier)? CLOSE_PAR
    ;

createDatabase
    : CREATE DATABASE (IF NOT EXISTS)? identifier (IN identifier)? (WITH (BUFFERED? LOG | LOG MODE ANSI))? (NLSCASE SENSITIVE | NLSCASE INSENSITIVE)?
    ;

createDistinctType
    : CREATE DISTINCT TYPE (IF NOT EXISTS)? dataTypeName AS dataTypeName
    ;

createDefaultUser
    : CREATE DEFAULT USER WITH properties
    ;

createExternalTable
    : CREATE EXTERNAL TABLE (IF NOT EXISTS)? tableName columnDeifinition USING OPEN_PAR tableOptions? datafilesClause tableOptions? CLOSE_PAR
    ;

columnDeifinition
    : SAMEAS identifier
    | identifier dataType otherOptionalClause (COMMA identifier dataType otherOptionalClause)*
    ;

dataType
    : builtInDataTypes
    | identifier DOT dataTypeName
    | complexDataType
    ;

otherOptionalClause
    : (EXTERNAL CHAR OPEN_PAR NUMERIC_LITERAL CLOSE_PAR (NULL STRING_LITERAL | NOT NULL)?
    )?
    ;

tableOptions
    : (FORMAT SYM6 (DELIMITED | INFORMIX | FIXED) SYM6
    | DEFAULT | loadingModeOption | DBDATE STRING_LITERAL | DBMONEY STRING_LITERAL | DELIMITER STRING_LITERAL | RECORDEND STRING_LITERAL
    | MAXERRORS NUMERIC_LITERAL | REJECTFILE STRING_LITERAL | ESCAPE (ON | OFF)? | (NUMROWS | SIZE) NUMERIC_LITERAL
    )
    ;

loadingModeOption
    : EXPRESS | DELUXE
    ;

datafilesClause
    : DATAFILES OPEN_PAR STRING_LITERAL (COMMA STRING_LITERAL)* CLOSE_PAR
    ;

createFunction
    : CREATE DBA? FUNCTION (IF NOT EXISTS)? identifier OPEN_PAR routineParameterList? CLOSE_PAR
    (referencingClause FOR tableName)? returnClause (SPECIFIC identifier)? (WITH OPEN_PAR routineModifier (COMMA routineModifier)* CLOSE_PAR)? SCOL?
    (statementBlock | externalRoutineReference) END FUNCTION (DOCUMENT QUOTED_STRING (COMMA QUOTED_STRING)*)? (WITH LISTING IN STRING_LITERAL)?
    ;

routineParameterList
    : (IN | OUT | INOUT)? parameter (COMMA (IN | OUT | INOUT)? parameter)*
    ;

parameter
    : identifier? ((dataTypeName | LINK identifier) (DEFAULT identifier)?
    | REFERENCES (BYTE | TEXT) (DEFAULT NULL)?
    )
    ;

referencingClause
    : REFERENCING ((NEW | OLD) AS? identifier)+ FOR identifier
    ;

returnClause
    : (RETURNING | RETURNS) (dataTypeName | REFERENCE (BYTE | TEXT)) (AS identifier)?
    (COMMA (RETURNING | RETURNS) (dataTypeName | REFERENCE (BYTE | TEXT)) (AS identifier)?)*
    ;

statementBlock
    : defineStatement* onException* (executeFunction | executeProcedure
    | subsetSplStatement | subsetSqlStatement | BEGIN statementBlock END
    )*
    ;

defineStatement
    : DEFINE (
    | identifier (COMMA identifier)* (dataTypeName | REFERENCES (BYTE | TEXT) | LIKE tableName DOT column=identifier | PROCEDURE | BLOB | CLOB | subsetComplexDataTypes)
    ) SCOL?
    ;

onException
    : ON EXCEPTION (IN OPEN_PAR NUMERIC_LITERAL (COMMA NUMERIC_LITERAL)* CLOSE_PAR)?
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
    : EXTERNAL NAME QUOTED_STRING LANGUAGE (C | JAVA) (PARAMETER STYLE INFORMIX?)? (NOT? VARIANT)?
    ;

createFunctionFrom
    : CREATE FUNCTION FROM (IF NOT EXISTS)? identifier
    ;

createIndex
    : CREATE indexTypeOptions INDEX (IF NOT EXISTS)? identifier ON tableName indexKeySpecs indexOptions?
    ;

indexTypeOptions
    : (DISTINCT | UNIQUE)? CLUSTER?
    ;

indexKeySpecs
    : OPEN_PAR (column=identifier | functionName=identifier OPEN_PAR func_col=identifier (COMMA func_col=identifier | op_class=identifier | bsonFieldSpecification)* CLOSE_PAR) (ASC | DESC)? CLOSE_PAR
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
    : USING identifier OPEN_PAR QUOTED_STRING ASSIGN (QUOTED_STRING | NUMERIC_LITERAL) (COMMA QUOTED_STRING ASSIGN (QUOTED_STRING | NUMERIC_LITERAL))* CLOSE_PAR
    ;

filefactorOption
    : FILLFACTOR NUMERIC_LITERAL
    ;

storageOptions
    : IN (identifier | TABLE)
    | fragmentClauseIndexes
    ;

indexModes
    : ENABLED | DISABLED | FILTERING (WITHOUT ERROR | WITH ERROR)?
    ;

hashOnClause
    : HASH ON OPEN_PAR identifier (COMMA identifier)? CLOSE_PAR WITH NUMERIC_LITERAL BUCKETS
    ;

extentSizeOptions
    : (EXTENT SIZE expression) (NEXT SIZE expression)?
    | (EXTENT SIZE expression)? (NEXT SIZE expression)
    ;

createOpaqueType
    : CREATE OPAQUE TYPE (IF NOT EXISTS)? identifier OPEN_PAR INTERNALLENGTH ASSIGN (NUMERIC_LITERAL | VARIABLE) (COMMA opaqueTypeModifier)* CLOSE_PAR
    ;

opaqueTypeModifier
    : MAXLEN ASSIGN NUMERIC_LITERAL
    | CANNOTHASH
    | PASSEDBYVALUE
    | ALIGNMENT ASSIGN NUMERIC_LITERAL
    ;

createOpclass
    : CREATE OPCLASS (IF NOT EXISTS)? identifier FOR identifier STRATEGIES OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR
    SUPPORT OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR
    ;

createProcedure
    : CREATE DBA? PROCEDURE (IF NOT EXISTS)? identifier OPEN_PAR routineParameterList? CLOSE_PAR (referencingClause FOR tableName)? returnClause? (SPECIFIC identifier)?
    (WITH OPEN_PAR routineModifier (COMMA routineModifier)* CLOSE_PAR)? SCOL? (statementBlock | externalRoutineReference) END PROCEDURE
    (DOCUMENT QUOTED_STRING (COMMA QUOTED_STRING)*)? (WITH LISTING IN STRING_LITERAL)?
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
    | (OPEN_PAR fieldDefinition (COMMA fieldDefinition)* CLOSE_PAR)? UNDER dataTypeName
    )
    ;

fieldDefinition
    : identifier dataTypeName (NOT NULL)?
    ;

createSchema
    : CREATE SCHEMA AUTHORIZATION identifier (createTable | createView | grantStatement | createIndex | createSynonym | createTrigger | createSequence | createRowType
    | createOpaqueType | createDistinctType | createCast)+ SCOL?
    ;

createSecurityLabel
    : CREATE SECURITY LABEL (IF NOT EXISTS)? identifier COMPONENT identifier QUOTED_STRING (COMMA QUOTED_STRING)* (COMMA COMPONENT identifier QUOTED_STRING (COMMA QUOTED_STRING)*)*
    ;

createSecurityLabelComponent
    : CREATE SECURITY LABEL COMPONENT (IF NOT EXISTS)? identifier (ARRAY SYM2 QUOTED_STRING (COMMA QUOTED_STRING)* SYM3
    | SET SYM4 QUOTED_STRING (COMMA QUOTED_STRING)* SYM5
    | TREE OPEN_PAR QUOTED_STRING ROOT (COMMA QUOTED_STRING UNDER QUOTED_STRING)* CLOSE_PAR
    )
    ;

createSecurityPolicy
    : CREATE SECURITY POLICY (IF NOT EXISTS)? identifier COMPONENTS identifier (COMMA identifier)* (WITH IDSLBACRULES)?
    (RESTRICT NOT AUTHORIZED WRITE SECURITY LABEL | OVERRIDE NOT AUTHORIZED WRITE SECURITY LABEL)?
    ;

createSequence
    : CREATE SEQUENCE (IF NOT EXISTS)? identifier (INCREMENT BY? NUMERIC_LITERAL
    | START WITH? NUMERIC_LITERAL | NOMAXVALUE | MAXVALUE NUMERIC_LITERAL | NOMINVALUE | MINVALUE NUMERIC_LITERAL
    | NOCYCLE | CYCLE | CACHE NUMERIC_LITERAL | NOCACHE | ORDER | NOORDER
    )*
    ;

createSynonym
    : CREATE (PUBLIC | PRIVATE)? SYNONYM (IF NOT EXISTS)? identifier FOR identifier
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
    : (DELETE | SELECT (OF identifier (COMMA identifier)*)?) ON tableName deleteAndSelectClause
    | UPDATE (OF identifier (COMMA identifier)*)? ON tableName updateClause
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
    : INSERT ON identifier (REFERENCING NEW AS? identifier)?
    | DELETE ON identifier (REFERENCING OLD AS? identifier)?
    | UPDATE ON identifier ((REFERENCING OLD AS? identifier (NEW AS? identifier)?)? | (REFERENCING NEW AS? identifier (OLD AS? identifier)?)?)
    ;

createTrustedContext
    : CREATE TRUSTED CONTEXT identifier (USER | BASED UPON CONNECTION USING SYSTEM AUTHID) identifier
    (ATTRIBUTES OPEN_PAR ADDRESS QUOTED_STRING (COMMA ADDRESS QUOTED_STRING)* CLOSE_PAR
    | WITH USE FOR authorizedUserClause (COMMA authorizedUserClause)*
    | NO DEFAULT ROLE | DEFAULT ROLE identifier | DISABLE | ENABLE
    )*
    ;

createUser
    : CREATE (DEFAULT USER WITH (ACCOUNT UNLOCK | ACCOUNT LOCK)? properties | USER identifier (PASSWORD QUOTED_STRING)? (ACCOUNT UNLOCK | ACCOUNT LOCK)? properties?)
    ;

properties
    : PROPERTIES (UID NUMERIC_LITERAL GROUP OPEN_PAR (NUMERIC_LITERAL | identifier) (COMMA (NUMERIC_LITERAL | identifier))* CLOSE_PAR
    | USER identifier (GROUP OPEN_PAR (NUMERIC_LITERAL | identifier) (COMMA (NUMERIC_LITERAL | identifier))* CLOSE_PAR)?)
    (HOME QUOTED_STRING)? (AUTHORIZATION OPEN_PAR (DBSA | DBSSO | AAO | BARGROUP) (COMMA (DBSA | DBSSO | AAO | BARGROUP))* CLOSE_PAR)?
    ;

createView
    : CREATE VIEW (IF NOT EXISTS)? viewName identifier? (OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR | OF TYPE dataTypeName)? AS
    selectStatement (WITH CHECK OPTION)?
    ;

createXadatasource
    : CREATE XADATASOURCE (IF NOT EXISTS)? identifier USING identifier
    ;

createXadatasourceType
    : CREATE XADATASOURCE TYPE (IF NOT EXISTS)? identifier OPEN_PAR purposeOptions (COMMA purposeOptions)* CLOSE_PAR
    ;

databaseStatement
    : DATABASE identifier EXCLUSIVE?
    ;

deallocateCollection
    : DEALLOCATE COLLECTION SYM1 identifier
    ;

deallocateDescriptor
    : DEALLOCATE DESCRIPTOR (identifier | QUOTED_STRING)
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
    : DELETE optimizerDirectives? FROM? (identifier (AS? identifier)? | ONLY OPEN_PAR identifier CLOSE_PAR | identifier)
    (WHERE condition | WHERE CURRENT OF identifier)?
    ;

optimizerDirectives //TODO: add more
    : (SYM9 | SYM4 | SYM10)  (SYM4 | SYM11)?
    ;

describeStatement
    : DESCRIBE OUTPUT? identifier (USING SQL DESCRIPTOR (identifier | QUOTED_STRING)
    | INTO SQL DESCRIPTOR (identifier | QUOTED_STRING) //TODO: sqlda_pointer
    )
    ;

describeInput
    : DESCRIBE INPUT identifier (USING SQL DESCRIPTOR (identifier | QUOTED_STRING)
    | INTO SQL DESCRIPTOR (identifier | QUOTED_STRING) //TODO: sqlda_pointer
    )
    ;

disconnectStatement
    : DISCONNECT (CURRENT | ALL | DEFAULT | identifier | QUOTED_STRING)
    ;

dropAccessMethod
    : DROP ACCESS_METHOD (IF EXISTS)? identifier RESTRICT
    ;

dropAggregate
    : DROP AGGREGATE (IF EXISTS)? identifier
    ;

dropCast
    : DROP CAST (IF EXISTS)? OPEN_PAR dataTypeName AS dataTypeName CLOSE_PAR
    ;

dropDatabase
    : DROP DATABASE (IF EXISTS)? identifier
    ;

dropFunction
    : DROP FUNCTION (IF EXISTS)? (identifier (OPEN_PAR dataTypeName (COMMA dataTypeName)* CLOSE_PAR)?
    | SPECIFIC FUNCTION (IF EXISTS)? identifier)
    ;

dropIndex
    : DROP INDEX (IF EXISTS)? identifier ONLINE?
    ;

dropOpclass
    : DROP OPCLASS (IF EXISTS)? identifier RESTRICT
    ;

dropProcedure
    : DROP (PROCEDURE (IF EXISTS)? identifier (OPEN_PAR dataTypeName (COMMA dataTypeName)* CLOSE_PAR)?
    | SPECIFIC PROCEDURE (IF EXISTS)? identifier
    )
    ;

//https://www.ibm.com/docs/en/informix-servers/14.10?topic=statements-drop-role-statement
dropRole
    : DROP ROLE (IF EXISTS)? roleName
    ;

dropRoutine
    : DROP (ROUTINE (IF EXISTS)? identifier (OPEN_PAR dataTypeName (COMMA dataTypeName)* CLOSE_PAR)?
    | SPECIFIC ROUTINE (IF EXISTS)? identifier
    )
    ;

dropRowType
    : DROP ROW TYPE (IF EXISTS)? identifier RESTRICT
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
    : DROP SYNONYM (IF EXISTS)? identifier
    ;

//https://www.ibm.com/docs/en/informix-servers/14.10?topic=statements-drop-table-statement
dropTable
    : DROP TABLE (IF EXISTS)? tableName (CASCADE | RESTRICT)?
    ;

dropTrigger
    : DROP TRIGGER (IF EXISTS)? identifier
    ;

dropTrustedContext
    : DROP TRUSTED CONTEXT identifier
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
    : DROP XADATASOURCE (IF EXISTS)? identifier RESTRICT
    ;

dropXadatasourceType
    : DROP XADATASOURCE TYPE (IF EXISTS)? identifier RESTRICT
    ;

executeStatement
    : EXECUTE intoClause? usingClause?
    ;

intoClause
    : INTO (identifier ((SYM1 | INDICATOR) identifier)?
    | SQL DESCRIPTOR (identifier | QUOTED_STRING)
    //| DESCRIPTOR sqlda_pointer TODO: sqlda_pointer
    )
    ;

usingClause
    : USING (identifier ((SYM1 | INDICATOR) identifier)?
    | SQL DESCRIPTOR (identifier | QUOTED_STRING)
    //| DESCRIPTOR sqlda_pointer TODO: sqlda_pointer
    )
    ;

executeFunction
    : EXECUTE FUNCTION identifier (OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR)? intoClause? (WITH TRIGGER REFERENCES)?
    ;

argument
    : (identifier ASSIGN) (expression | NULL | OPEN_PAR selectStatement CLOSE_PAR)
    ;

executeImmediate
    : EXECUTE IMMEDIATE (QUOTED_STRING | identifier | expression)
    ;

executeProcedure
    : EXECUTE PROCEDURE identifier (OPEN_PAR (argument (COMMA argument)*)? CLOSE_PAR)? (INTO identifier (COMMA identifier)*)? (WITH TRIGGER REFERENCES)?
    ;

fetchStatement
    : FETCH (NEXT | PRIOR | FIRST | LAST | CURRENT | ABSOLUTE PLUS? (NUMERIC_LITERAL | identifier) | RELATIVE (NUMERIC_LITERAL | identifier))? identifier
    (USING SQL DESCRIPTOR (identifier | QUOTED_STRING)
    | INTO (identifier | identifier ((INDICATOR | SYM1) identifier)?) (COMMA (identifier | identifier ((INDICATOR | SYM1) identifier)?))*)?
    ;

flushStatement
    : FLUSH identifier
    ;

freeStatement
    : FREE identifier
    ;

getDescriptor
    : GET DESCRIPTOR (identifier | QUOTED_STRING) (identifier ASSIGN COUNT | VALUE (identifier | NUMERIC_LITERAL) describedItemInformation (COMMA describedItemInformation)*)
    ;

describedItemInformation
    : identifier ASSIGN (DATA | IDATA | ILENGTH | INDICATOR | ITYPE | LENGTH | NAME | NULLABLE| PRECISION | SCALE
    | TYPE | EXTYPEID | EXTYPELENGTH | EXTYPENAME | EXTYPEOWNERLENGTH | EXTYPEOWNERNAME | SOURCEID | SOURCETYPE)
    ;

getDiagnostics
    : GET DIAGNOSTICS (statementClause | exceptionClause)
    ;

statementClause
    : identifier ASSIGN (identifier | NUMERIC_LITERAL)
    ;

exceptionClause
    : EXCEPTION (identifier | NUMERIC_LITERAL) identifier ASSIGN (identifier | NUMERIC_LITERAL) (COMMA identifier ASSIGN (identifier | NUMERIC_LITERAL))*
    ;

grantStatement
    : GRANT ((CONNECT | RESOURCE | DBA | DEFAULT ROLE rolename) TO (PUBLIC | STRING_LITERAL (COMMA STRING_LITERAL)*)
    | rolename toOptions | securityAdministrationOptions | accessToPropertiesClause
    | (tableLevelPrivileges | routineLevelPrivileges | languageLevelPrivileges | typeLevelPrivileges | sequenceLevelPrivileges) toOptions
    )
    ;

rolename
    : identifier
    | STRING_LITERAL
    ;

toOptions
    : TO (STRING_LITERAL (COMMA STRING_LITERAL)* (WITH GRANT OPTION)? | (STRING_LITERAL | PUBLIC) (COMMA (STRING_LITERAL | PUBLIC))*) (AS STRING_LITERAL)?
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
    (UID NUMERIC_LITERAL GROUP OPEN_PAR (NUMERIC_LITERAL | identifier) (COMMA (NUMERIC_LITERAL | identifier))* CLOSE_PAR
    | USER identifier (GROUP OPEN_PAR (NUMERIC_LITERAL | identifier) (COMMA (NUMERIC_LITERAL | identifier))* CLOSE_PAR)?)
    (HOME QUOTED_STRING)? (COMMA AUTHORIZATION OPEN_PAR QUOTED_STRING (COMMA QUOTED_STRING)* CLOSE_PAR)?
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
    : INSERT (INTO (tableName (OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR)? (valuesClause | executeRoutineClause | selectStatement)
    | external=identifier (OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR)? selectStatement?
    )
    | (AT position=NUMERIC_LITERAL)? INTO collectionDerivedTable fieldOptions
    )
    ;

executeRoutineClause
    : EXECUTE (PROCEDURE | FUNCTION) identifier OPEN_PAR argument (COMMA argument)* CLOSE_PAR
    ;

collectionDerivedTable
    : TABLE OPEN_PAR (collectionExpression=expression CLOSE_PAR (AS? alias=identifier)? (OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR)?
    | identifier CLOSE_PAR
    )
    ;

fieldOptions
    : (field=identifier (COMMA field=identifier)*)? (valuesClause | selectStatement)
    ;

valuesClause
    : VALUES OPEN_PAR (identifier (SYM1 identifier | SYM12 identifier)?
    | NULL | USER | QUOTED_STRING | NUMERIC_LITERAL | constExprssion | columnExpression
    | literalCollection | literalRow | expression
    ) CLOSE_PAR
    ;

loadStatement
    : LOAD FROM QUOTED_STRING (DELIMITER QUOTED_STRING)? INSERT INTO tableName (OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR)?
    ;

lockTable
    : LOCK TABLE tableName IN (EXCLUSIVE | SHARE) MODE
    ;

mergeStatement
    : MERGE (optimizerDirectives (COMMA optimizerDirectives)*)? INTO tableName (AS alias=identifier)?
    USING (tableName | selectStatement) (AS alias=identifier)? (OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR)? ON condition
    (WHEN NOT MATCHED THEN INSERT (OPEN_PAR identifier (COMMA identifier)* CLOSE_PAR)? valuesClause
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
    : PREPARE identifier FROM (expression | identifier | QUOTED_STRING)
    ;

putStatement
    : PUT identifier (FROM identifier (INDICATOR identifier | SYM1 identifier | SYM12 identifier)? (COMMA identifier (INDICATOR identifier | SYM1 identifier | SYM12 identifier)?)*
    | USING SQL DESCRIPTOR QUOTED_STRING)?
    ;

releaseSavepoint
    : RELEASE SAVEPOINT identifier
    ;

renameColumn
    : RENAME COLUMN tableName DOT old_column=identifier TO new_column=identifier
    ;

renameConstraint
    : RENAME CONSTRAINT identifier DOT old_constraint=identifier TO new_constraint=identifier
    ;

renameDatabase
    : RENAME DATABASE identifier DOT old_database=identifier TO new_database=identifier
    ;

renameIndex
    : RENAME INDEX identifier DOT old_index=identifier TO new_index=identifier
    ;

renameSecurity
    : RENAME SECURITY (POLICY | LABEL identifier | LABEL COMPONENT) old_name=identifier TO new_name=identifier
    ;

renameSequence
    : RENAME SEQUENCE identifier DOT old_sequence=identifier TO new_sequence=identifier
    ;

renameTable
    : RENAME TABLE identifier DOT old_table=identifier TO new_table=identifier
    ;

renameTrustedContext
    : RENAME TRUSTED CONTEXT old_name=identifier TO new_name=identifier
    ;

renameUser
    : RENAME USER old_name=identifier TO new_name=identifier
    ;

revokeStatement
    : REVOKE
    ;

revokeFragment
    :
    ;

rollbackWork
    : ROLLBACK WORK? (TO SAVEPOINT identifier?)?
    ;

saveExternalDirectives
    : SAVE EXTERNAL DIRECTIVES optimizerDirectives (COMMA optimizerDirectives)* (ACTIVE | INACTIVE | TEST ONLY)
    FOR selectStatement
    ;

savepointStatement
    : SAVEPOINT identifier UNIQUE?
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
    : (SKIP_ NUMERIC_LITERAL)? ((FIRST | LIMIT) NUMERIC_LITERAL)? (DISTINCT | ALL | UNIQUE)? selectList (COMMA selectList)*
    ;

selectList
    : (expression | column=identifier) (AS? displayLabel=identifier)?
    | tableName DOT (column=identifier (AS? identifier)? | STAR)
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
    : tableName (AS? alias=identifier)?
    | collectionDerivedTable
    | iterator
    ;

iterator
    : TABLE OPEN_PAR (FUNCTION | PROCEDURE)? identifier OPEN_PAR routineParameterList? CLOSE_PAR CLOSE_PAR
    (AS? tableName OPEN_PAR column=identifier (COMMA column=identifier)* CLOSE_PAR)?
    ;

ansiJoins
    : ((INNER | (LEFT | RIGHT | FULL) OUTER?)? JOIN ansiTables onCLause
    | CROSS JOIN ansiTables
    ) (AS? alias=identifier)?
    ;

onCLause
    : ON (tableName DOT column=identifier relationalOperator tableName DOT column=identifier | function | condition
    | OPEN_PAR selectStatement CLOSE_PAR | OPEN_PAR collectionSubquery CLOSE_PAR)
    ((AND | OR) (tableName DOT column=identifier relationalOperator tableName DOT column=identifier | function | condition
    | OPEN_PAR selectStatement CLOSE_PAR | OPEN_PAR collectionSubquery CLOSE_PAR))*
    ;

otherTables
    : external=identifier (AS? alias=identifier (OPEN_PAR derivedColumn=identifier (COMMA derivedColumn=identifier)* CLOSE_PAR)?)?
    | OPEN_PAR collectionSubquery CLOSE_PAR
    | OPEN_PAR selectStatement CLOSE_PAR
    | ONLY? tableName AS? alias=identifier OPEN_PAR derivedColumn=identifier (COMMA derivedColumn=identifier)* CLOSE_PAR
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
    : LATERAL OPEN_PAR selectStatement CLOSE_PAR AS? alias=identifier (OPEN_PAR columnAlias=identifier (COMMA columnAlias=identifier)* CLOSE_PAR)?
    ;

gridClause
    : GRID ALL? identifier
    ;

whereClause
    : WHERE (condition | (tableName DOT column=identifier relationalOperator tableName DOT column=identifier) | function
    | OPEN_PAR selectStatement CLOSE_PAR | OPEN_PAR collectionSubquery CLOSE_PAR | identifier)
    (logicalOperator (condition | (tableName DOT column=identifier relationalOperator tableName DOT column=identifier) | function
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
    : GROUP BY ((tableName DOT)? (identifier | expression) | colAlias=identifier | NUMERIC_LITERAL)
    (COMMA ((tableName DOT)? (identifier | expression) | colAlias=identifier | NUMERIC_LITERAL))*
    ;

havingClause
    : HAVING condition
    ;

orderByClause
    : ORDER SIBLINGS? BY ((tableName DOT column=identifier | expression) substring?
    | selectNumber=NUMERIC_LITERAL | displayLabel=identifier | caseExpression | ROWID //| olapWindowExpressions TODO
    ) (ASC | DESC)? (NULLS FIRST | NULLS LAST)?
    ;

substring
    : SYM2 NUMERIC_LITERAL COMMA NUMERIC_LITERAL SYM3
    ;

limitClause
    : LIMIT NUMERIC_LITERAL
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
    : SET AUTOFREE (ENABLED | DISABLED)? (FOR identifier)?
    ;

setCollation
    : SET (COLLATION QUOTED_STRING | NO COLLATION)
    ;

setLockMode
    : SET LOCK MODE TO (WAIT NUMERIC_LITERAL | NOWAIT)
    ;

setLog
    : SET BUFFERED? LOG
    ;

setOptimization
    : SET OPTIMIZATION (HIGH | LOW | FIRST_ROWS | ALL_ROWS | environmentOptions)
    ;

environmentOptions
    : ENVIROMENT (STAR_JOIN SYM6 (ENABLED | DISABLED | FORCED) SYM6
    | (FACT | AVOID_FACT | NON_DIM) (SYM6 tableName (COMMA tableName)* SYM6 | DEFAULT | SYM7))
    ;

setPdqpriority
    : SET PDQPRIORITY (DEFAULT | NUMERIC_LITERAL | LOW | OFF | HIGH)
    ;

setRole
    : SET ROLE (NULL | NONE | identifier | QUOTED_STRING | DEFAULT)
    ;

setSessionAuthorization
    : SET SESSION AUTHORIZATION TO QUOTED_STRING (USING QUOTED_STRING)?
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
    : SET USER PASSWORD OLD old_password=QUOTED_STRING NEW new_password=QUOTED_STRING
    ;

startViolationsTable
    : START VIOLATIONS TABLE FOR tableName (USING identifier COMMA identifier)? (MAX ROWS NUMERIC_LITERAL)?
    ;

stopViolationsTable
    : STOP VIOLATIONS TABLE FOR tableName
    ;

truncateStatement
    : TRUNCATE TABLE? tableName (DROP STORAGE | REUSE STORAGE KEEP STATISTICS)?
    ;

unloadStatement
    : UNLOAD TO filename=QUOTED_STRING (DELIMITER delimiter=QUOTED_STRING)? (selectStatement | identifier)
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
    : (FOR TABLE (ONLY? tableName (OPEN_PAR column=identifier (COMMA column=identifier)* CLOSE_PAR)?)?)?
    ;

resolutionClause
    : resolutionClauseMedium
    | resolutionClauseHigh
    ;

resolutionClauseMedium
    : (SAMPLING SIZE NUMERIC_LITERAL) (RESOLUTION NUMERIC_LITERAL NUMERIC_LITERAL?)? (DISTRIBUTIONS ONLY)?
    | (SAMPLING SIZE NUMERIC_LITERAL)? (RESOLUTION NUMERIC_LITERAL NUMERIC_LITERAL?) (DISTRIBUTIONS ONLY)?
    | (SAMPLING SIZE NUMERIC_LITERAL)? (RESOLUTION NUMERIC_LITERAL NUMERIC_LITERAL?)? (DISTRIBUTIONS ONLY)
    ;

resolutionClauseHigh
    : (RESOLUTION NUMERIC_LITERAL) (DISTRIBUTIONS ONLY)?
    | (RESOLUTION NUMERIC_LITERAL)? (DISTRIBUTIONS ONLY)
    ;

routineStatistics
    : FOR ((ROUTINE | PROCEDURE | FUNCTION) (identifier (OPEN_PAR routineParameterList? CLOSE_PAR)?)?
    | SPECIFIC (ROUTINE | PROCEDURE | FUNCTION) identifier
    )
    ;

updateTarget
    : tableName (AS? alias=identifier)?
    | ONLY OPEN_PAR tableName CLOSE_PAR
    ;

whereOptions
    : (WHERE condition
    | WHERE CURRENT OF identifier
    )?
    ;

wheneverStatement
    : WHENEVER (SQLERROR | ERROR | NOT FOUND | SQLWARNING) (CONTINUE | STOP | CALL identifier | (GOTO | GO TO) SYM1? identifier)
    ;

constExprssion
    : QUOTED_STRING
    | NUMERIC_LITERAL
    | USER
    | CURRENT_USER
    | CURRENT_ROLE
    | DEFAULT_ROLE
    | anyName
    | TODAY
    | (CURRENT | SYSDATE) NUMERIC_LITERAL?
    | literalDatetime
    | literalInterval
    | NUMERIC_LITERAL UNITS timeUnit
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
    : QUOTED_STRING
    | literalDatetime
    | NUMERIC_LITERAL
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
    : QUOTED_STRING
    | NUMERIC_LITERAL
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
    : tableName DOT (column=identifier (SYM2 NUMERIC_LITERAL COMMA NUMERIC_LITERAL SYM3)? | ROWID | rowColumn=identifier (DOT STAR | (DOT identifier)+)?)
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
    | identifier NOT? (LIKE | MATCHES) identifier (ESCAPE QUOTED_STRING)?
    ;

inCondition
    : expression NOT? IN (OPEN_PAR (NUMERIC_LITERAL | literalDatetime | identifier | literalInterval | USER | CURRENT_USER | TODAY | CURRENT datetimeField? | literalRow) CLOSE_PAR
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
    : functionName=identifier OPEN_PAR
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
    | (BIGSERIAL | SERIAL | SERIAL8) (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?
    | (MONEY | DECIMAL | DEC | NUMERIC) (OPEN_PAR NUMERIC_LITERAL (COMMA NUMERIC_LITERAL)? CLOSE_PAR)
    ;

approximateNumericDataType
    : (DECIMAL | DEC | NUMERIC) (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?
    | (FLOAT | DOUBLE PRECISION) (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?
    | SMALLFLOAT
    | REAL
    ;

characterDataType
    : (CHAR | CHARACTER | NCHAR) OPEN_PAR NUMERIC_LITERAL CLOSE_PAR
    | (NVARCHAR | VARCHAR | CHARACTER VARYING) OPEN_PAR NUMERIC_LITERAL (COMMA NUMERIC_LITERAL)? CLOSE_PAR
    | LVARCHAR OPEN_PAR NUMERIC_LITERAL CLOSE_PAR
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

tableName
    : identifier
    ;

viewName
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

datetimeField
    : YEAR TO (YEAR | MONTH | DAY | HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | MONTH TO (MONTH | DAY | HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | DAY TO (DAY | HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | HOUR TO (HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | MINUTE TO (MINUTE | SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | SECOND TO (SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | FRACTION TO FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?
    ;

intervalField
    : DAY (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)? TO (DAY | HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | HOUR (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)? TO (HOUR | MINUTE | SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | MINUTE (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)? TO (MINUTE | SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | SECOND (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)? TO (SECOND | FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?)
    | FRACTION TO FRACTION (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)?
    | YEAR (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)? TO (YEAR | MONTH)
    | MONTH (OPEN_PAR NUMERIC_LITERAL CLOSE_PAR)? TO MONTH
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
    | AVG
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
    | CLOB
    | CLOBDIR
    | CLOSE
    | CLUSTER
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
    | DATE
    | DATETIME
    | DAY
    | DBDATE
    | DBMONEY
    | DBSA
    | DBSECADM
    | DBSSO
    | DEALLOCATE
    | DEC
    | DECIMAL
    | DECLARE
    | DEFAULT
    | DEFAULT_ROLE
    | DEFERRABLE
    | DEFERRED
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
    | DISABLE
    | DISABLED
    | DISCARD
    | DISCONNECT
    | DISK
    | DISTINCT
    | DISTRIBUTIONS
    | DOCUMENT
    | DOUBLE
    | DROP
    | EACH
    | ELIF
    | ELSE
    | ENABLE
    | ENABLED
    | END
    | ENVIROMENT
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
    | EXTENT
    | EXTERNAL
    | EXTYPEID
    | EXTYPELENGTH
    | EXTYPENAME
    | EXTYPEOWNERLENGTH
    | EXTYPEOWNERNAME
    | FACT
    | FAIL
    | FETCH
    | FILLFACTOR
    | FILTERING
    | FINAL
    | FIRST_ROWS
    | FIXED
    | FLOAT
    | FLUSH
    | FORCE
    | FORCED
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
    | GROUP
    | HANDLESNULLS
    | HASH
    | HAVING
    | HEADINGS
    | HIGH
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
    | IGNORE
    | ILENGTH
    | IMMEDIATE
    | IMPLICIT
    | IN
    | INACTIVE
    | INCREMENT
    | INDEX
    | INDEXED
    | INDEXES
    | INDICATOR
    | INFO
    | INFORMIX
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
    | UNITS
    | UNLOAD
    | UNLOCK
    | UPDATE
    | UPDATING
    | UPON
    | USAGE
    | USE
    | USER
    | USING
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