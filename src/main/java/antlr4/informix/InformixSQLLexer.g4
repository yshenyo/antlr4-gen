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


lexer grammar InformixSQLLexer;

options { caseInsensitive = true; }

SCOL:      ';';
DOT:       '.';
OPEN_PAR:  '(';
CLOSE_PAR: ')';
COMMA:     ',';
ASSIGN:    '=';
STAR:      '*';
PLUS:      '+';
MINUS:     '-';
TILDE:     '~';
PIPE2:     '||';
DIV:       '/';
MOD:       '%';
LT2:       '<<';
GT2:       '>>';
AMP:       '&';
PIPE:      '|';
LT:        '<';
LT_EQ:     '<=';
GT:        '>';
GT_EQ:     '>=';
EQ:        '==';
NOT_EQ1:   '!=';
NOT_EQ2:   '<>';

SYM1:      ':';
SYM2:      '[';
SYM3:      ']';
SYM4:      '{';
SYM5:      '}';
SYM6:      '\'';
SYM7:      '\'\'';
SYM8:      '::';
SYM9:      '--';
SYM10:     '/*';
SYM11:     '*/';
SYM12:     '$';
SYM13:     '#';

// Interval type Keywords
AAO:               'AAO';
ABORT:             'ABORT';
ABSOLUTE:          'ABSOLUTE';
ACCELERATE:        'ACCELERATE';
ACCESS:            'ACCESS';
ACCESS_METHOD:     'ACCESS_METHOD';
ACCOUNT:           'ACCOUNT';
ACTION:            'ACTION';
ACTIVE:            'ACTIVE';
ADD:               'ADD';
ADDRESS:           'ADDRESS';
AGGREGATE:         'AGGREGATE';
AFTER:             'AFTER';
ALIGNMENT:         'ALIGNMENT';
ALL:               'ALL';
ALLOCATE:          'ALLOCATE';
ALL_ROWS:          'ALL_ROWS';
ALTER:             'ALTER';
ANALYZE:           'ANALYZE';
AND:               'AND';
ANSI:              'ANSI';
ANY:               'ANY';
APPEND:            'APPEND';
ARRAY:             'ARRAY';
AS:                'AS';
ASC:               'ASC';
AT:                'AT';
ATTACH:            'ATTACH';
ATTRIBUTES:        'ATTRIBUTES';
AUDIT:             'AUDIT';
AUTHENTICATION:    'AUTHENTICATION';
AUTHID:            'AUTHID';
AUTHORIZATION:     'AUTHORIZATION';
AUTHORIZED:        'AUTHORIZED';
AUTO:              'AUTO';
AUTOFREE:          'AUTOFREE';
AUTOINCREMENT:     'AUTOINCREMENT';
AUTOLOCATE:        'AUTOLOCATE';
AUTO_READAHEAD:    'AUTO_READAHEAD';
AUTO_STAT_MODE:    'AUTO_STAT_MODE';
AVG:               'AVG';
AVOID_EXECUTE:     'AVOID_EXECUTE';
AVOID_FACT:        'AVOID_FACT';
BARGROUP:          'BARGROUP';
BASED:             'BASED';
BEFORE:            'BEFORE';
BEGIN:             'BEGIN';
BETWEEN:           'BETWEEN';
BIGINT:            'BIGINT';
BIGSERIAL:         'BIGSERIAL';
BLOB:              'BLOB';
BLOBDIR:           'BLOBDIR';
BOUND_IMPL_PDQ:    'BOUND_IMPL_PDQ';
BOOLEAN:           'BOOLEAN';
BSON:              'BSON';
BSON_GET:          'BSON_GET';
BSON_VALUE_BIGINT: 'BSON_VALUE_BIGINT';
BSON_VALUE_BOOLEAN:'BSON_VALUE_BOOLEAN';
BSON_VALUE_DATE:   'BSON_VALUE_DATE';
BSON_VALUE_DOUBLE: 'BSON_VALUE_DOUBLE';
BSON_VALUE_INT:    'BSON_VALUE_INT';
BSON_VALUE_LVARCHAR:'BSON_VALUE_LVARCHAR';
BSON_VALUE_OBJECTID:'BSON_VALUE_OBJECTID';
BSON_VALUE_TIMESTAMP:'BSON_VALUE_TIMESTAMP';
BSON_VALUE_VARCHAR:'BSON_VALUE_VARCHAR';
BUCKETS:           'BUCKETS';
BUFFERED:          'BUFFERED';
BY:                'BY';
BYTE:              'BYTE';
C:                 'C';
CACHE:             'CACHE';
CALL:              'CALL';
CANNOTHASH:        'CANNOTHASH';
CASCADE:           'CASCADE';
CASE:              'CASE';
CAST:              'CAST';
CHAR:              'CHAR';
CHARACTER:         'CHARACTER';
CHECK:             'CHECK';
CLASS:             'CLASS';
CLEANUP:           'CLEANUP';
CLOB:              'CLOB';
CLOBDIR:           'CLOBDIR';
CLOSE:             'CLOSE';
CLUSTER:           'CLUSTER';
CLUSTER_TXN_SCOPE: 'CLUSTER_TXN_SCOPE';
COLLATE:           'COLLATE';
COLLATION:         'COLLATION';
COLLECTION:        'COLLECTION';
COLUMN:            'COLUMN';
COLUMNS:           'COLUMNS';
COMBINE:           'COMBINE';
COMMIT:            'COMMIT';
COMMITTED:         'COMMITTED';
COMPONENT:         'COMPONENT';
COMPONENTS:        'COMPONENTS';
COMPRESSED:        'COMPRESSED';
CONCURRENT:        'CONCURRENT';
CONFLICT:          'CONFLICT';
CONNECT:           'CONNECT';
CONNECTION:        'CONNECTION';
CONSTRAINT:        'CONSTRAINT';
CONSTRAINTS:       'CONSTRAINTS';
CONTEXT:           'CONTEXT';
COSTFUNC:          'COSTFUNC';
CONTINUE:          'CONTINUE';
COUNT:             'COUNT';
CRCOLS:            'CRCOLS';
CREATE:            'CREATE';
CROSS:             'CROSS';
CURRENT_DATE:      'CURRENT_DATE';
CURRENT_ROLE:      'CURRENT_ROLE';
CURRENT_TIME:      'CURRENT_TIME';
CURRENT_TIMESTAMP: 'CURRENT_TIMESTAMP';
CURRENT_USER:      'CURRENT_USER';
CURRVAL:           'CURRVAL';
CURSOR:            'CURSOR';
CYCLE:             'CYCLE';
DATA:              'DATA';
DATABASE:          'DATABASE';
DATAFILES:         'DATAFILES';
DATASKIP:          'DATASKIP';
DATE:              'DATE';
DATETIME:          'DATETIME';
DAY:               'DAY';
DBA:               'DBA';
DBDATE:            'DBDATE';
DBMONEY:           'DBMONEY';
DBSA:              'DBSA';
DBSECADM:          'DBSECADM';
DBSSO:             'DBSSO';
DEALLOCATE:        'DEALLOCATE';
DEBUG:             'DEBUG';
DEC:               'DEC';
DECIMAL:           'DECIMAL';
DECLARE:           'DECLARE';
DEFAULT:           'DEFAULT';
DEFAULTESCCHAR:    'DEFAULTESCCHAR';
DEFAULT_ROLE:      'DEFAULT_ROLE';
DEFERRABLE:        'DEFERRABLE';
DEFERRED:          'DEFERRED';
DEFERRED_PREPARE:  'DEFERRED_PREPARE';
DEFINE:            'DEFINE';
DELETE:            'DELETE';
DELETING:          'DELETING';
DELIMITED:         'DELIMITED';
DELIMITER:         'DELIMITER';
DELUXE:            'DELUXE';
DESC:              'DESC';
DESCRIBE:          'DESCRIBE';
DESCRIPTOR:        'DESCRIPTOR';
DETACH:            'DETACH';
DIAGNOSTICS:       'DIAGNOSTICS';
DIRECTIVES:        'DIRECTIVES';
DIRTY:             'DIRTY';
DISABLE:           'DISABLE';
DISABLED:          'DISABLED';
DISCARD:           'DISCARD';
DISCONNECT:        'DISCONNECT';
DISK:              'DISK';
DISTINCT:          'DISTINCT';
DISTRIBUTIONS:     'DISTRIBUTIONS';
DOCUMENT:          'DOCUMENT';
DORMANT:           'DORMANT';
DOUBLE:            'DOUBLE';
DROP:              'DROP';
EACH:              'EACH';
ELIF:              'ELIF';
ELSE:              'ELSE';
ENABLE:            'ENABLE';
ENABLED:           'ENABLED';
ENCRYPTION:        'ENCRYPTION';
END:               'END';
ENVIRONMENT:       'ENVIRONMENT';
ERKEY:             'ERKEY';
ERROR:             'ERROR';
ESCAPE:            'ESCAPE';
EXCEPT:            'EXCEPT';
EXCEPTION:         'EXCEPTION';
EXCLUSIVE:         'EXCLUSIVE';
EXECUTE:           'EXECUTE';
EXEMPTION:         'EXEMPTION';
EXISTS:            'EXISTS';
EXIT:              'EXIT';
EXPLAIN:           'EXPLAIN';
EXPLICIT:          'EXPLICIT';
EXPRESS:           'EXPRESS';
EXPRESSION:        'EXPRESSION';
EXTDIRECTIVES:     'EXTDIRECTIVES';
EXTENT:            'EXTENT';
EXTERNAL:          'EXTERNAL';
EXTYPEID:          'EXTYPEID';
EXTYPELENGTH:      'EXTYPELENGTH';
EXTYPENAME:        'EXTYPENAME';
EXTYPEOWNERLENGTH: 'EXTYPEOWNERLENGTH';
EXTYPEOWNERNAME:   'EXTYPEOWNERNAME';
FACT:              'FACT';
FAIL:              'FAIL';
FALLBACK:          'FALLBACK';
FETCH:             'FETCH';
FILE:              'FILE';
FILLFACTOR:        'FILLFACTOR';
FILTERING:         'FILTERING';
FINAL:             'FINAL';
FIRST_ROWS:        'FIRST_ROWS';
FIXED:             'FIXED';
FLOAT:             'FLOAT';
FLUSH:             'FLUSH';
FORCE:             'FORCE';
FORCED:            'FORCED';
FORCE_DDL_EXEC:    'FORCE_DDL_EXEC';
FOUND:             'FOUND';
FOR:               'FOR';
FOREACH:           'FOREACH';
FOREIGN:           'FOREIGN';
FORMAT:            'FORMAT';
FRACTION:          'FRACTION';
FRAGMENT:          'FRAGMENT';
FRAGMENTS:         'FRAGMENTS';
FREE:              'FREE';
FROM:              'FROM';
FULL:              'FULL';
FUNCTION:          'FUNCTION';
GET:               'GET';
GLOB:              'GLOB';
GO:                'GO';
GOTO:              'GOTO';
GRANT:             'GRANT';
GRID:              'GRID';
GRID_NODE_SKIP:    'GRID_NODE_SKIP';
GROUP:             'GROUP';
HANDLESNULLS:      'HANDLESNULLS';
HASH:              'HASH';
HAVING:            'HAVING';
HDR_TXN_SCOPE:     'HDR_TXN_SCOPE';
HEADINGS:          'HEADINGS';
HIGH:              'HIGH';
HINT:              'HINT';
HOLD:              'HOLD';
HOME:              'HOME';
HOUR:              'HOUR';
IDATA:             'IDATA';
IDSLBACREADARRAY:  'IDSLBACREADARRAY';
IDSLBACREADTREE:   'IDSLBACREADTREE';
IDSLBACREADSET:    'IDSLBACREADSET';
IDSLBACRULES:      'IDSLBACRULES';
IDSLBACWRITEARRAY: 'IDSLBACWRITEARRAY';
IDSLBACWRITESET:   'IDSLBACWRITESET';
IDSLBACWRITETREE:  'IDSLBACWRITETREE';
IDSSECURITYLABEL:  'IDSSECURITYLABEL';
IF:                'IF';
IFX_AUTO_REPREPARE: 'IFX_AUTO_REPREPARE';
IFX_BATCHEDREAD_INDEX: 'IFX_BATCHEDREAD_INDEX';
IFX_BATCHEDREAD_TABLE: 'IFX_BATCHEDREAD_TABLE';
IFX_SESSION_LIMIT_LOCKS: 'IFX_SESSION_LIMIT_LOCKS';
IGNORE:            'IGNORE';
ILENGTH:           'ILENGTH';
IMMEDIATE:         'IMMEDIATE';
IMPLICIT:          'IMPLICIT';
IMPLICIT_PDQ:      'IMPLICIT_PDQ';
IN:                'IN';
INACTIVE:          'INACTIVE';
INCREMENT:         'INCREMENT';
INDEX:             'INDEX';
INDEXED:           'INDEXED';
INDEXES:           'INDEXES';
INDICATOR:         'INDICATOR';
INFO:              'INFO';
INFORMIX:          'INFORMIX';
INFORMIXCONRETRY:  'INFORMIXCONRETRY';
INFORMIXCONTIME:   'INFORMIXCONTIME';
INIT:              'INIT';
INITIALLY:         'INITIALLY';
INNER:             'INNER';
INOUT:             'INOUT';
INPUT:             'INPUT';
INSENSITIVE:       'INSENSITIVE';
INSERT:            'INSERT';
INSERTING:         'INSERTING';
INSTEAD:           'INSTEAD';
INT:               'INT';
INT8:              'INT8';
INTEGER:           'INTEGER';
INTEG:             'INTEG';
INTERNAL:          'INTERNAL';
INTERNALLENGTH:    'INTERNALLENGTH';
INTERSECT:         'INTERSECT';
INTERVAL:          'INTERVAL';
INTO:              'INTO';
ISOLATION:         'ISOLATION';
ITER:              'ITER';
ITERATOR:          'ITERATOR';
ITEM:              'ITEM';
ITYPE:             'ITYPE';
IS:                'IS';
ISNULL:            'ISNULL';
JAVA:              'JAVA';
JOIN:              'JOIN';
JSON:              'JSON';
KEEP:              'KEEP';
KEY:               'KEY';
LABEL:             'LABEL';
LATERAL:           'LATERAL';
LANGUAGE:          'LANGUAGE';
LEFT:              'LEFT';
LENGTH:            'LENGTH';
LET:               'LET';
LEVEL:             'LEVEL';
LIKE:              'LIKE';
LIMIT:             'LIMIT';
LINK:              'LINK';
LIST:              'LIST';
LISTING:           'LISTING';
LOAD:              'LOAD';
LOCK:              'LOCK';
LOCKS:             'LOCKS';
LOG:               'LOG';
LOOP:              'LOOP';
LOW:               'LOW';
LVARCHAR:          'LVARCHAR';
MATCH:             'MATCH';
MATCHED:           'MATCHED';
MATCHES:           'MATCHES';
MAX:               'MAX';
MAXERRORS:         'MAXERRORS';
MAXLEN:            'MAXLEN';
MAXVALUE:          'MAXVALUE';
MEDIUM:            'MEDIUM';
MERGE:             'MERGE';
MIN:               'MIN';
MINUTE:            'MINUTE';
MINVALUE:          'MINVALUE';
MODE:              'MODE';
MODERATE:          'MODERATE';
MODIFY:            'MODIFY';
MONEY:             'MONEY';
MONTH:             'MONTH';
MULTISET:          'MULTISET';
NAME:              'NAME';
NATURAL:           'NATURAL';
NCHAR:             'NCHAR';
NEGATOR:           'NEGATOR';
NEXT:              'NEXT';
NEXTVAL:           'NEXTVAL';
NEW:               'NEW';
NLSCASE:           'NLSCASE';
NO:                'NO';
NOCACHE:           'NOCACHE';
NOCYCLE:           'NOCYCLE';
NOMAXVALUE:        'NOMAXVALUE';
NOMINVALUE:        'NOMINVALUE';
NON_DIM:           'NON_DIM';
NONE:              'NONE';
NOORDER:           'NOORDER';
NOT:               'NOT';
NOTNULL:           'NOTNULL';
NOVALIDATE:        'NOVALIDATE';
NOWAIT:            'NOWAIT';
NULL:              'NULL';
NULLABLE:          'NULLABLE';
NUMERIC:           'NUMERIC';
NUMROWS:           'NUMROWS';
NVARCHAR:          'NVARCHAR';
OF:                'OF';
OFF:               'OFF';
OFFSET:            'OFFSET';
OLD:               'OLD';
ON:                'ON';
ONLINE:            'ONLINE';
ONLY:              'ONLY';
OPAQUE:            'OPAQUE';
OPCLASS:           'OPCLASS';
OPEN:              'OPEN';
OPTCOMPIND:        'OPTCOMPIND';
OPTIMIZATION:      'OPTIMIZATION';
OPTION:            'OPTION';
OR:                'OR';
ORDER:             'ORDER';
OUT:               'OUT';
OUTER:             'OUTER';
OUTPUT:            'OUTPUT';
OVERRIDE:          'OVERRIDE';
PAGE:              'PAGE';
PARALLELIZABLE:    'PARALLELIZABLE';
PARAMETER:         'PARAMETER';
PASSEDBYVALUE:     'PASSEDBYVALUE';
PASSWORD:          'PASSWORD';
PDQPRIORITY:       'PDQPRIORITY';
PERCALL_COST:      'PERCALL_COST';
PLAN:              'PLAN';
POLICY:            'POLICY';
PRAGMA:            'PRAGMA';
PRECISION:         'PRECISION';
PREPARE:           'PREPARE';
PRIMARY:           'PRIMARY';
PRIOR:             'PRIOR';
PRIVATE:           'PRIVATE';
PRIVILEGE:         'PRIVILEGE';
PRIVILEGES:        'PRIVILEGES';
PROBE:             'PROBE';
PROCEDURE:         'PROCEDURE';
PROPERTIES:        'PROPERTIES';
PUBLIC:            'PUBLIC';
PUT:               'PUT';
QUERY:             'QUERY';
RAISE:             'RAISE';
RAW:               'RAW';
READ:              'READ';
REAL:              'REAL';
RECORDEND:         'RECORDEND';
RECURSIVE:         'RECURSIVE';
REFERENCING:       'REFERENCING';
REFERENCE:         'REFERENCE';
REFERENCES:        'REFERENCES';
REGEXP:            'REGEXP';
REINDEX:           'REINDEX';
REJECTFILE:        'REJECTFILE';
RELATIVE:          'RELATIVE';
RELEASE:           'RELEASE';
REMAINDER:         'REMAINDER';
RENAME:            'RENAME';
REOPTIMIZATION:    'REOPTIMIZATION';
REPEATABLE:        'REPEATABLE';
REPLACE:           'REPLACE';
REPLCHECK:         'REPLCHECK';
REPLICATION:       'REPLICATION';
RESOLUTION:        'RESOLUTION';
RESOURCE:          'RESOURCE';
RESTART:           'RESTART';
RESTRICT:          'RESTRICT';
RESUME:            'RESUME';
RETAIN:            'RETAIN';
RETAINUPDATELOCKS: 'RETAINUPDATELOCKS';
RETURN:            'RETURN';
RETURNING:         'RETURNING';
RETURNS:           'RETURNS';
REUSE:             'REUSE';
REVOKE:            'REVOKE';
RIGHT:             'RIGHT';
ROBIN:             'ROBIN';
ROLLBACK:          'ROLLBACK';
ROOT:              'ROOT';
ROUND:             'ROUND';
ROUTINE:           'ROUTINE';
ROW:               'ROW';
ROWID:             'ROWID';
ROWIDS:            'ROWIDS';
ROLE:              'ROLE';
ROLLING:           'ROLLING';
ROWS:              'ROWS';
RULE:              'RULE';
SAMEAS:            'SAMEAS';
SAMPLING:          'SAMPLING';
SAVE:              'SAVE';
SAVEPOINT:         'SAVEPOINT';
SCALE:             'SCALE';
SCROLL:            'SCROLL';
SECOND:            'SECOND';
SECONDARY:         'SECONDARY';
SECURED:           'SECURED';
SECURITY:          'SECURITY';
SELCONST:          'SELCONST';
SELECT:            'SELECT';
SELECT_GRID:       'SELECT_GRID';
SELECT_GRID_ALL:   'SELECT_GRID_ALL';
SELECTING:         'SELECTING';
SELFUNC:           'SELFUNC';
SENSITIVE:         'SENSITIVE';
SEQUENCE:          'SEQUENCE';
SERIAL:            'SERIAL';
SERIAL8:           'SERIAL8';
SERIALIZABLE:      'SERIALIZABLE';
SESSION:           'SESSION';
SET:               'SET';
SETSESSIONAUTH:    'SETSESSIONAUTH';
SCHEMA:            'SCHEMA';
SHARE:             'SHARE';
SIBLINGS:          'SIBLINGS';
SIZE:              'SIZE';
SKIP_:              'SKIP';
SMALLFLOAT:        'SMALLFLOAT';
SMALLINT:          'SMALLINT';
SOURCEID:          'SOURCEID';
SOURCETYPE:        'SOURCETYPE';
SPECIFIC:          'SPECIFIC';
SPL:               'SPL';
SQL:               'SQL';
SQLERROR:          'SQLERROR';
SQLWARNING:        'SQLWARNING';
STABILITY:         'STABILITY';
STACK:             'STACK';
STANDARD:          'STANDARD';
STAR_JOIN:         'STAR_JOIN';
START:             'START';
STATCHANGE:        'STATCHANGE';
STATEMENT:         'STATEMENT';
STATISTICS:        'STATISTICS';
STATLEVEL:         'STATLEVEL';
STATUS:            'STATUS';
STDEV:             'STDEV';
STEP:              'STEP';
STOP:              'STOP';
STORAGE:           'STORAGE';
STORE:             'STORE';
STRATEGIES:        'STRATEGIES';
STYLE:             'STYLE';
SUM:               'SUM';
SUPPORT:           'SUPPORT';
SYNONYM:           'SYNONYM';
SYSDATE:           'SYSDATE';
SYSTEM:            'SYSTEM';
TABLE:             'TABLE';
TABLES:            'TABLES';
TRACE:             'TRACE';
TEMP:              'TEMP';
TEMPORARY:         'TEMPORARY';
TEST:              'TEST';
TEXT:              'TEXT';
THEN:              'THEN';
TIME:              'TIME';
TO:                'TO';
TODAY:             'TODAY';
TRANSACTION:       'TRANSACTION';
TRANSITION:        'TRANSITION';
TREE:              'TREE';
TRIGGER:           'TRIGGER';
TRIGGERS:          'TRIGGERS';
TRUNCATE:          'TRUNCATE';
TRUSTED:           'TRUSTED';
TYPE:              'TYPE';
UID:               'UID';
UNCOMMITTED:       'UNCOMMITTED';
UNDER:             'UNDER';
UNION:             'UNION';
UNIQUE:            'UNIQUE';
UNIQUECHECK:       'UNIQUECHECK';
UNITS:             'UNITS';
UNLOAD:            'UNLOAD';
UNLOCK:            'UNLOCK';
UPDATE:            'UPDATE';
UPDATING:          'UPDATING';
UPON:              'UPON';
USAGE:             'USAGE';
USE:               'USE';
USE_DWA:           'USE_DWA';
USELASTCOMMITTED:  'USELASTCOMMITTED';
USER:              'USER';
USE_SHARDING:      'USE_SHARDING';
USING:             'USING';
USTLOW_SAMPLE:     'USTLOW_SAMPLE';
VACUUM:            'VACUUM';
VALUE:             'VALUE';
VALUES:            'VALUES';
VARCHAR:           'VARCHAR';
VARIABLE:          'VARIABLE';
VARIANCE:          'VARIANCE';
VARIANT:           'VARIANT';
VARYING:           'VARYING';
VERCOLS:           'VERCOLS';
VIEW:              'VIEW';
VIOLATIONS:        'VIOLATIONS';
VIRTUAL:           'VIRTUAL';
WAIT:              'WAIT';
WARNING:           'WARNING';
WHEN:              'WHEN';
WHENEVER:          'WHENEVER';
WHERE:             'WHERE';
WHILE:             'WHILE';
WITH:              'WITH';
WITHOUT:           'WITHOUT';
WORK:              'WORK';
WRITE:             'WRITE';
WRITEDOWN:         'WRITEDOWN';
WRITEUP:           'WRITEUP';
XADATASOURCE:      'XADATASOURCE';
YEAR:              'YEAR';
FIRST_VALUE:       'FIRST_VALUE';
OVER:              'OVER';
PARTITION:         'PARTITION';
RANGE:             'RANGE';
PRECEDING:         'PRECEDING';
UNBOUNDED:         'UNBOUNDED';
CURRENT:           'CURRENT';
FOLLOWING:         'FOLLOWING';
CUME_DIST:         'CUME_DIST';
DENSE_RANK:        'DENSE_RANK';
LAG:               'LAG';
LAST_VALUE:        'LAST_VALUE';
LEAD:              'LEAD';
NTH_VALUE:         'NTH_VALUE';
NTILE:             'NTILE';
PERCENT_RANK:      'PERCENT_RANK';
RANK:              'RANK';
ROW_NUMBER:        'ROW_NUMBER';
GENERATED:         'GENERATED';
ALWAYS:            'ALWAYS';
STORED:            'STORED';
TRUE:              'TRUE';
FALSE:             'FALSE';
WINDOW:            'WINDOW';
NULLS:             'NULLS';
FIRST:             'FIRST';
LAST:              'LAST';
FILTER:            'FILTER';
GROUPS:            'GROUPS';
EXCLUDE:           'EXCLUDE';
TIES:              'TIES';
OTHERS:            'OTHERS';
DO:                'DO';
NOTHING:           'NOTHING';

IDENTIFIER:
    '"' (~'"' | '""')* '"'
    | '`' (~'`' | '``')* '`'
    | '[' ~']'* ']'
    | [A-Z_] [A-Z_0-9]*
; // TODO check: needs more chars in set

NUMERIC_LITERAL: (DIGIT+ ('.' DIGIT*)? | '.' DIGIT+) ('E' [-+]? DIGIT+)? | '0x' HEX_DIGIT+;

BIND_PARAMETER: '?' DIGIT* | [:@$] IDENTIFIER;

STRING_LITERAL: '\'' (~'\'' | '\'\'')* '\'';

QUOTED_STRING: '\'' (~'\'' | '\'\'')* '\'' | '"' (~'"' | '""')* '"';

BLOB_LITERAL: 'X' STRING_LITERAL;

SINGLE_LINE_COMMENT: '--' ~[\r\n]* ('\r'? '\n' | EOF) -> channel(HIDDEN);

MULTILINE_COMMENT: '/*' .*? '*/' -> channel(HIDDEN);

SPACES: [ \u000B\t\r\n] -> channel(HIDDEN);

UNEXPECTED_CHAR: .;

fragment HEX_DIGIT: [0-9A-F];
fragment DIGIT:     [0-9];

QUOTA : '\'' | '"' | '`';