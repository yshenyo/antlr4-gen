package com.thinice.bbc.core.parser.antlr.informix;
// Generated from InformixSQLParser.g4 by ANTLR 4.12.0
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class InformixSQLParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.12.0", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		SCOL=1, DOT=2, OPEN_PAR=3, CLOSE_PAR=4, COMMA=5, ASSIGN=6, STAR=7, PLUS=8, 
		MINUS=9, TILDE=10, PIPE2=11, DIV=12, MOD=13, LT2=14, GT2=15, AMP=16, PIPE=17, 
		LT=18, LT_EQ=19, GT=20, GT_EQ=21, EQ=22, NOT_EQ1=23, NOT_EQ2=24, ABORT=25, 
		ACTION=26, ADD=27, AFTER=28, ALL=29, ALTER=30, ANALYZE=31, AND=32, AS=33, 
		ASC=34, ATTACH=35, AUTOINCREMENT=36, BEFORE=37, BEGIN=38, BETWEEN=39, 
		BY=40, CASCADE=41, CASE=42, CAST=43, CHECK=44, COLLATE=45, COLUMN=46, 
		COMMIT=47, CONFLICT=48, CONSTRAINT=49, CREATE=50, CROSS=51, CURRENT_DATE=52, 
		CURRENT_TIME=53, CURRENT_TIMESTAMP=54, DATABASE=55, DEFAULT=56, DEFERRABLE=57, 
		DEFERRED=58, DELETE=59, DESC=60, DETACH=61, DISTINCT=62, DROP=63, EACH=64, 
		ELSE=65, END=66, ESCAPE=67, EXCEPT=68, EXCLUSIVE=69, EXISTS=70, EXPLAIN=71, 
		FAIL=72, FOR=73, FOREIGN=74, FROM=75, FULL=76, GLOB=77, GROUP=78, HAVING=79, 
		IF=80, IGNORE=81, IMMEDIATE=82, IN=83, INDEX=84, INDEXED=85, INITIALLY=86, 
		INNER=87, INSERT=88, INSTEAD=89, INTERSECT=90, INTO=91, IS=92, ISNULL=93, 
		JOIN=94, KEY=95, LEFT=96, LIKE=97, LIMIT=98, MATCH=99, NATURAL=100, NO=101, 
		NOT=102, NOTNULL=103, NULL=104, OF=105, OFFSET=106, ON=107, OR=108, ORDER=109, 
		OUTER=110, PLAN=111, PRAGMA=112, PRIMARY=113, QUERY=114, RAISE=115, RECURSIVE=116, 
		REFERENCES=117, REGEXP=118, REINDEX=119, RELEASE=120, RENAME=121, REPLACE=122, 
		RESTRICT=123, RETURNING=124, RIGHT=125, ROLLBACK=126, ROW=127, ROLE=128, 
		ROWS=129, SAVEPOINT=130, SELECT=131, SET=132, TABLE=133, TEMP=134, TEMPORARY=135, 
		THEN=136, TO=137, TRANSACTION=138, TRIGGER=139, UNION=140, UNIQUE=141, 
		UPDATE=142, USING=143, VACUUM=144, VALUES=145, VIEW=146, VIRTUAL=147, 
		WHEN=148, WHERE=149, WITH=150, WITHOUT=151, FIRST_VALUE=152, OVER=153, 
		PARTITION=154, RANGE=155, PRECEDING=156, UNBOUNDED=157, CURRENT=158, FOLLOWING=159, 
		CUME_DIST=160, DENSE_RANK=161, LAG=162, LAST_VALUE=163, LEAD=164, NTH_VALUE=165, 
		NTILE=166, PERCENT_RANK=167, RANK=168, ROW_NUMBER=169, GENERATED=170, 
		ALWAYS=171, STORED=172, TRUE=173, FALSE=174, WINDOW=175, NULLS=176, FIRST=177, 
		LAST=178, FILTER=179, GROUPS=180, EXCLUDE=181, TIES=182, OTHERS=183, DO=184, 
		NOTHING=185, IDENTIFIER=186, NUMERIC_LITERAL=187, BIND_PARAMETER=188, 
		STRING_LITERAL=189, BLOB_LITERAL=190, SINGLE_LINE_COMMENT=191, MULTILINE_COMMENT=192, 
		SPACES=193, UNEXPECTED_CHAR=194;
	public static final int
		RULE_sqlScript = 0, RULE_unitStatement = 1, RULE_createRole = 2, RULE_dropRole = 3, 
		RULE_keyword = 4, RULE_roleName = 5, RULE_anyName = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"sqlScript", "unitStatement", "createRole", "dropRole", "keyword", "roleName", 
			"anyName"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'.'", "'('", "')'", "','", "'='", "'*'", "'+'", "'-'", 
			"'~'", "'||'", "'/'", "'%'", "'<<'", "'>>'", "'&'", "'|'", "'<'", "'<='", 
			"'>'", "'>='", "'=='", "'!='", "'<>'", "'ABORT'", "'ACTION'", "'ADD'", 
			"'AFTER'", "'ALL'", "'ALTER'", "'ANALYZE'", "'AND'", "'AS'", "'ASC'", 
			"'ATTACH'", "'AUTOINCREMENT'", "'BEFORE'", "'BEGIN'", "'BETWEEN'", "'BY'", 
			"'CASCADE'", "'CASE'", "'CAST'", "'CHECK'", "'COLLATE'", "'COLUMN'", 
			"'COMMIT'", "'CONFLICT'", "'CONSTRAINT'", "'CREATE'", "'CROSS'", "'CURRENT_DATE'", 
			"'CURRENT_TIME'", "'CURRENT_TIMESTAMP'", "'DATABASE'", "'DEFAULT'", "'DEFERRABLE'", 
			"'DEFERRED'", "'DELETE'", "'DESC'", "'DETACH'", "'DISTINCT'", "'DROP'", 
			"'EACH'", "'ELSE'", "'END'", "'ESCAPE'", "'EXCEPT'", "'EXCLUSIVE'", "'EXISTS'", 
			"'EXPLAIN'", "'FAIL'", "'FOR'", "'FOREIGN'", "'FROM'", "'FULL'", "'GLOB'", 
			"'GROUP'", "'HAVING'", "'IF'", "'IGNORE'", "'IMMEDIATE'", "'IN'", "'INDEX'", 
			"'INDEXED'", "'INITIALLY'", "'INNER'", "'INSERT'", "'INSTEAD'", "'INTERSECT'", 
			"'INTO'", "'IS'", "'ISNULL'", "'JOIN'", "'KEY'", "'LEFT'", "'LIKE'", 
			"'LIMIT'", "'MATCH'", "'NATURAL'", "'NO'", "'NOT'", "'NOTNULL'", "'NULL'", 
			"'OF'", "'OFFSET'", "'ON'", "'OR'", "'ORDER'", "'OUTER'", "'PLAN'", "'PRAGMA'", 
			"'PRIMARY'", "'QUERY'", "'RAISE'", "'RECURSIVE'", "'REFERENCES'", "'REGEXP'", 
			"'REINDEX'", "'RELEASE'", "'RENAME'", "'REPLACE'", "'RESTRICT'", "'RETURNING'", 
			"'RIGHT'", "'ROLLBACK'", "'ROW'", "'ROLE'", "'ROWS'", "'SAVEPOINT'", 
			"'SELECT'", "'SET'", "'TABLE'", "'TEMP'", "'TEMPORARY'", "'THEN'", "'TO'", 
			"'TRANSACTION'", "'TRIGGER'", "'UNION'", "'UNIQUE'", "'UPDATE'", "'USING'", 
			"'VACUUM'", "'VALUES'", "'VIEW'", "'VIRTUAL'", "'WHEN'", "'WHERE'", "'WITH'", 
			"'WITHOUT'", "'FIRST_VALUE'", "'OVER'", "'PARTITION'", "'RANGE'", "'PRECEDING'", 
			"'UNBOUNDED'", "'CURRENT'", "'FOLLOWING'", "'CUME_DIST'", "'DENSE_RANK'", 
			"'LAG'", "'LAST_VALUE'", "'LEAD'", "'NTH_VALUE'", "'NTILE'", "'PERCENT_RANK'", 
			"'RANK'", "'ROW_NUMBER'", "'GENERATED'", "'ALWAYS'", "'STORED'", "'TRUE'", 
			"'FALSE'", "'WINDOW'", "'NULLS'", "'FIRST'", "'LAST'", "'FILTER'", "'GROUPS'", 
			"'EXCLUDE'", "'TIES'", "'OTHERS'", "'DO'", "'NOTHING'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "SCOL", "DOT", "OPEN_PAR", "CLOSE_PAR", "COMMA", "ASSIGN", "STAR", 
			"PLUS", "MINUS", "TILDE", "PIPE2", "DIV", "MOD", "LT2", "GT2", "AMP", 
			"PIPE", "LT", "LT_EQ", "GT", "GT_EQ", "EQ", "NOT_EQ1", "NOT_EQ2", "ABORT", 
			"ACTION", "ADD", "AFTER", "ALL", "ALTER", "ANALYZE", "AND", "AS", "ASC", 
			"ATTACH", "AUTOINCREMENT", "BEFORE", "BEGIN", "BETWEEN", "BY", "CASCADE", 
			"CASE", "CAST", "CHECK", "COLLATE", "COLUMN", "COMMIT", "CONFLICT", "CONSTRAINT", 
			"CREATE", "CROSS", "CURRENT_DATE", "CURRENT_TIME", "CURRENT_TIMESTAMP", 
			"DATABASE", "DEFAULT", "DEFERRABLE", "DEFERRED", "DELETE", "DESC", "DETACH", 
			"DISTINCT", "DROP", "EACH", "ELSE", "END", "ESCAPE", "EXCEPT", "EXCLUSIVE", 
			"EXISTS", "EXPLAIN", "FAIL", "FOR", "FOREIGN", "FROM", "FULL", "GLOB", 
			"GROUP", "HAVING", "IF", "IGNORE", "IMMEDIATE", "IN", "INDEX", "INDEXED", 
			"INITIALLY", "INNER", "INSERT", "INSTEAD", "INTERSECT", "INTO", "IS", 
			"ISNULL", "JOIN", "KEY", "LEFT", "LIKE", "LIMIT", "MATCH", "NATURAL", 
			"NO", "NOT", "NOTNULL", "NULL", "OF", "OFFSET", "ON", "OR", "ORDER", 
			"OUTER", "PLAN", "PRAGMA", "PRIMARY", "QUERY", "RAISE", "RECURSIVE", 
			"REFERENCES", "REGEXP", "REINDEX", "RELEASE", "RENAME", "REPLACE", "RESTRICT", 
			"RETURNING", "RIGHT", "ROLLBACK", "ROW", "ROLE", "ROWS", "SAVEPOINT", 
			"SELECT", "SET", "TABLE", "TEMP", "TEMPORARY", "THEN", "TO", "TRANSACTION", 
			"TRIGGER", "UNION", "UNIQUE", "UPDATE", "USING", "VACUUM", "VALUES", 
			"VIEW", "VIRTUAL", "WHEN", "WHERE", "WITH", "WITHOUT", "FIRST_VALUE", 
			"OVER", "PARTITION", "RANGE", "PRECEDING", "UNBOUNDED", "CURRENT", "FOLLOWING", 
			"CUME_DIST", "DENSE_RANK", "LAG", "LAST_VALUE", "LEAD", "NTH_VALUE", 
			"NTILE", "PERCENT_RANK", "RANK", "ROW_NUMBER", "GENERATED", "ALWAYS", 
			"STORED", "TRUE", "FALSE", "WINDOW", "NULLS", "FIRST", "LAST", "FILTER", 
			"GROUPS", "EXCLUDE", "TIES", "OTHERS", "DO", "NOTHING", "IDENTIFIER", 
			"NUMERIC_LITERAL", "BIND_PARAMETER", "STRING_LITERAL", "BLOB_LITERAL", 
			"SINGLE_LINE_COMMENT", "MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "InformixSQLParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public InformixSQLParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SqlScriptContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(InformixSQLParser.EOF, 0); }
		public List<UnitStatementContext> unitStatement() {
			return getRuleContexts(UnitStatementContext.class);
		}
		public UnitStatementContext unitStatement(int i) {
			return getRuleContext(UnitStatementContext.class,i);
		}
		public SqlScriptContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sqlScript; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).enterSqlScript(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).exitSqlScript(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof InformixSQLParserVisitor ) return ((InformixSQLParserVisitor<? extends T>)visitor).visitSqlScript(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SqlScriptContext sqlScript() throws RecognitionException {
		SqlScriptContext _localctx = new SqlScriptContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_sqlScript);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(17);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CREATE || _la==DROP) {
				{
				{
				setState(14);
				unitStatement();
				}
				}
				setState(19);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(20);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UnitStatementContext extends ParserRuleContext {
		public TerminalNode SCOL() { return getToken(InformixSQLParser.SCOL, 0); }
		public CreateRoleContext createRole() {
			return getRuleContext(CreateRoleContext.class,0);
		}
		public DropRoleContext dropRole() {
			return getRuleContext(DropRoleContext.class,0);
		}
		public UnitStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unitStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).enterUnitStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).exitUnitStatement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof InformixSQLParserVisitor ) return ((InformixSQLParserVisitor<? extends T>)visitor).visitUnitStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final UnitStatementContext unitStatement() throws RecognitionException {
		UnitStatementContext _localctx = new UnitStatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_unitStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CREATE:
				{
				setState(22);
				createRole();
				}
				break;
			case DROP:
				{
				setState(23);
				dropRole();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(26);
			match(SCOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CreateRoleContext extends ParserRuleContext {
		public TerminalNode CREATE() { return getToken(InformixSQLParser.CREATE, 0); }
		public TerminalNode ROLE() { return getToken(InformixSQLParser.ROLE, 0); }
		public RoleNameContext roleName() {
			return getRuleContext(RoleNameContext.class,0);
		}
		public TerminalNode IF() { return getToken(InformixSQLParser.IF, 0); }
		public TerminalNode NOT() { return getToken(InformixSQLParser.NOT, 0); }
		public TerminalNode EXISTS() { return getToken(InformixSQLParser.EXISTS, 0); }
		public CreateRoleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_createRole; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).enterCreateRole(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).exitCreateRole(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof InformixSQLParserVisitor ) return ((InformixSQLParserVisitor<? extends T>)visitor).visitCreateRole(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CreateRoleContext createRole() throws RecognitionException {
		CreateRoleContext _localctx = new CreateRoleContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_createRole);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(28);
			match(CREATE);
			setState(29);
			match(ROLE);
			setState(33);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(30);
				match(IF);
				setState(31);
				match(NOT);
				setState(32);
				match(EXISTS);
				}
				break;
			}
			setState(35);
			roleName();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DropRoleContext extends ParserRuleContext {
		public TerminalNode DROP() { return getToken(InformixSQLParser.DROP, 0); }
		public TerminalNode ROLE() { return getToken(InformixSQLParser.ROLE, 0); }
		public RoleNameContext roleName() {
			return getRuleContext(RoleNameContext.class,0);
		}
		public TerminalNode IF() { return getToken(InformixSQLParser.IF, 0); }
		public TerminalNode EXISTS() { return getToken(InformixSQLParser.EXISTS, 0); }
		public DropRoleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dropRole; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).enterDropRole(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).exitDropRole(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof InformixSQLParserVisitor ) return ((InformixSQLParserVisitor<? extends T>)visitor).visitDropRole(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DropRoleContext dropRole() throws RecognitionException {
		DropRoleContext _localctx = new DropRoleContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_dropRole);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(37);
			match(DROP);
			setState(38);
			match(ROLE);
			setState(41);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(39);
				match(IF);
				setState(40);
				match(EXISTS);
				}
				break;
			}
			setState(43);
			roleName();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class KeywordContext extends ParserRuleContext {
		public TerminalNode ABORT() { return getToken(InformixSQLParser.ABORT, 0); }
		public TerminalNode ACTION() { return getToken(InformixSQLParser.ACTION, 0); }
		public TerminalNode ADD() { return getToken(InformixSQLParser.ADD, 0); }
		public TerminalNode AFTER() { return getToken(InformixSQLParser.AFTER, 0); }
		public TerminalNode ALL() { return getToken(InformixSQLParser.ALL, 0); }
		public TerminalNode ALTER() { return getToken(InformixSQLParser.ALTER, 0); }
		public TerminalNode ANALYZE() { return getToken(InformixSQLParser.ANALYZE, 0); }
		public TerminalNode AND() { return getToken(InformixSQLParser.AND, 0); }
		public TerminalNode AS() { return getToken(InformixSQLParser.AS, 0); }
		public TerminalNode ASC() { return getToken(InformixSQLParser.ASC, 0); }
		public TerminalNode ATTACH() { return getToken(InformixSQLParser.ATTACH, 0); }
		public TerminalNode AUTOINCREMENT() { return getToken(InformixSQLParser.AUTOINCREMENT, 0); }
		public TerminalNode BEFORE() { return getToken(InformixSQLParser.BEFORE, 0); }
		public TerminalNode BEGIN() { return getToken(InformixSQLParser.BEGIN, 0); }
		public TerminalNode BETWEEN() { return getToken(InformixSQLParser.BETWEEN, 0); }
		public TerminalNode BY() { return getToken(InformixSQLParser.BY, 0); }
		public TerminalNode CASCADE() { return getToken(InformixSQLParser.CASCADE, 0); }
		public TerminalNode CASE() { return getToken(InformixSQLParser.CASE, 0); }
		public TerminalNode CAST() { return getToken(InformixSQLParser.CAST, 0); }
		public TerminalNode CHECK() { return getToken(InformixSQLParser.CHECK, 0); }
		public TerminalNode COLLATE() { return getToken(InformixSQLParser.COLLATE, 0); }
		public TerminalNode COLUMN() { return getToken(InformixSQLParser.COLUMN, 0); }
		public TerminalNode COMMIT() { return getToken(InformixSQLParser.COMMIT, 0); }
		public TerminalNode CONFLICT() { return getToken(InformixSQLParser.CONFLICT, 0); }
		public TerminalNode CONSTRAINT() { return getToken(InformixSQLParser.CONSTRAINT, 0); }
		public TerminalNode CREATE() { return getToken(InformixSQLParser.CREATE, 0); }
		public TerminalNode CROSS() { return getToken(InformixSQLParser.CROSS, 0); }
		public TerminalNode CURRENT_DATE() { return getToken(InformixSQLParser.CURRENT_DATE, 0); }
		public TerminalNode CURRENT_TIME() { return getToken(InformixSQLParser.CURRENT_TIME, 0); }
		public TerminalNode CURRENT_TIMESTAMP() { return getToken(InformixSQLParser.CURRENT_TIMESTAMP, 0); }
		public TerminalNode DATABASE() { return getToken(InformixSQLParser.DATABASE, 0); }
		public TerminalNode DEFAULT() { return getToken(InformixSQLParser.DEFAULT, 0); }
		public TerminalNode DEFERRABLE() { return getToken(InformixSQLParser.DEFERRABLE, 0); }
		public TerminalNode DEFERRED() { return getToken(InformixSQLParser.DEFERRED, 0); }
		public TerminalNode DELETE() { return getToken(InformixSQLParser.DELETE, 0); }
		public TerminalNode DESC() { return getToken(InformixSQLParser.DESC, 0); }
		public TerminalNode DETACH() { return getToken(InformixSQLParser.DETACH, 0); }
		public TerminalNode DISTINCT() { return getToken(InformixSQLParser.DISTINCT, 0); }
		public TerminalNode DROP() { return getToken(InformixSQLParser.DROP, 0); }
		public TerminalNode EACH() { return getToken(InformixSQLParser.EACH, 0); }
		public TerminalNode ELSE() { return getToken(InformixSQLParser.ELSE, 0); }
		public TerminalNode END() { return getToken(InformixSQLParser.END, 0); }
		public TerminalNode ESCAPE() { return getToken(InformixSQLParser.ESCAPE, 0); }
		public TerminalNode EXCEPT() { return getToken(InformixSQLParser.EXCEPT, 0); }
		public TerminalNode EXCLUSIVE() { return getToken(InformixSQLParser.EXCLUSIVE, 0); }
		public TerminalNode EXISTS() { return getToken(InformixSQLParser.EXISTS, 0); }
		public TerminalNode EXPLAIN() { return getToken(InformixSQLParser.EXPLAIN, 0); }
		public TerminalNode FAIL() { return getToken(InformixSQLParser.FAIL, 0); }
		public TerminalNode FOR() { return getToken(InformixSQLParser.FOR, 0); }
		public TerminalNode FOREIGN() { return getToken(InformixSQLParser.FOREIGN, 0); }
		public TerminalNode FROM() { return getToken(InformixSQLParser.FROM, 0); }
		public TerminalNode FULL() { return getToken(InformixSQLParser.FULL, 0); }
		public TerminalNode GLOB() { return getToken(InformixSQLParser.GLOB, 0); }
		public TerminalNode GROUP() { return getToken(InformixSQLParser.GROUP, 0); }
		public TerminalNode HAVING() { return getToken(InformixSQLParser.HAVING, 0); }
		public TerminalNode IF() { return getToken(InformixSQLParser.IF, 0); }
		public TerminalNode IGNORE() { return getToken(InformixSQLParser.IGNORE, 0); }
		public TerminalNode IMMEDIATE() { return getToken(InformixSQLParser.IMMEDIATE, 0); }
		public TerminalNode IN() { return getToken(InformixSQLParser.IN, 0); }
		public TerminalNode INDEX() { return getToken(InformixSQLParser.INDEX, 0); }
		public TerminalNode INDEXED() { return getToken(InformixSQLParser.INDEXED, 0); }
		public TerminalNode INITIALLY() { return getToken(InformixSQLParser.INITIALLY, 0); }
		public TerminalNode INNER() { return getToken(InformixSQLParser.INNER, 0); }
		public TerminalNode INSERT() { return getToken(InformixSQLParser.INSERT, 0); }
		public TerminalNode INSTEAD() { return getToken(InformixSQLParser.INSTEAD, 0); }
		public TerminalNode INTERSECT() { return getToken(InformixSQLParser.INTERSECT, 0); }
		public TerminalNode INTO() { return getToken(InformixSQLParser.INTO, 0); }
		public TerminalNode IS() { return getToken(InformixSQLParser.IS, 0); }
		public TerminalNode ISNULL() { return getToken(InformixSQLParser.ISNULL, 0); }
		public TerminalNode JOIN() { return getToken(InformixSQLParser.JOIN, 0); }
		public TerminalNode KEY() { return getToken(InformixSQLParser.KEY, 0); }
		public TerminalNode LEFT() { return getToken(InformixSQLParser.LEFT, 0); }
		public TerminalNode LIKE() { return getToken(InformixSQLParser.LIKE, 0); }
		public TerminalNode LIMIT() { return getToken(InformixSQLParser.LIMIT, 0); }
		public TerminalNode MATCH() { return getToken(InformixSQLParser.MATCH, 0); }
		public TerminalNode NATURAL() { return getToken(InformixSQLParser.NATURAL, 0); }
		public TerminalNode NO() { return getToken(InformixSQLParser.NO, 0); }
		public TerminalNode NOT() { return getToken(InformixSQLParser.NOT, 0); }
		public TerminalNode NOTNULL() { return getToken(InformixSQLParser.NOTNULL, 0); }
		public TerminalNode NULL() { return getToken(InformixSQLParser.NULL, 0); }
		public TerminalNode OF() { return getToken(InformixSQLParser.OF, 0); }
		public TerminalNode OFFSET() { return getToken(InformixSQLParser.OFFSET, 0); }
		public TerminalNode ON() { return getToken(InformixSQLParser.ON, 0); }
		public TerminalNode OR() { return getToken(InformixSQLParser.OR, 0); }
		public TerminalNode ORDER() { return getToken(InformixSQLParser.ORDER, 0); }
		public TerminalNode OUTER() { return getToken(InformixSQLParser.OUTER, 0); }
		public TerminalNode PLAN() { return getToken(InformixSQLParser.PLAN, 0); }
		public TerminalNode PRAGMA() { return getToken(InformixSQLParser.PRAGMA, 0); }
		public TerminalNode PRIMARY() { return getToken(InformixSQLParser.PRIMARY, 0); }
		public TerminalNode QUERY() { return getToken(InformixSQLParser.QUERY, 0); }
		public TerminalNode RAISE() { return getToken(InformixSQLParser.RAISE, 0); }
		public TerminalNode RECURSIVE() { return getToken(InformixSQLParser.RECURSIVE, 0); }
		public TerminalNode REFERENCES() { return getToken(InformixSQLParser.REFERENCES, 0); }
		public TerminalNode REGEXP() { return getToken(InformixSQLParser.REGEXP, 0); }
		public TerminalNode REINDEX() { return getToken(InformixSQLParser.REINDEX, 0); }
		public TerminalNode RELEASE() { return getToken(InformixSQLParser.RELEASE, 0); }
		public TerminalNode RENAME() { return getToken(InformixSQLParser.RENAME, 0); }
		public TerminalNode REPLACE() { return getToken(InformixSQLParser.REPLACE, 0); }
		public TerminalNode RESTRICT() { return getToken(InformixSQLParser.RESTRICT, 0); }
		public TerminalNode RIGHT() { return getToken(InformixSQLParser.RIGHT, 0); }
		public TerminalNode ROLLBACK() { return getToken(InformixSQLParser.ROLLBACK, 0); }
		public TerminalNode ROW() { return getToken(InformixSQLParser.ROW, 0); }
		public TerminalNode ROWS() { return getToken(InformixSQLParser.ROWS, 0); }
		public TerminalNode SAVEPOINT() { return getToken(InformixSQLParser.SAVEPOINT, 0); }
		public TerminalNode SELECT() { return getToken(InformixSQLParser.SELECT, 0); }
		public TerminalNode SET() { return getToken(InformixSQLParser.SET, 0); }
		public TerminalNode TABLE() { return getToken(InformixSQLParser.TABLE, 0); }
		public TerminalNode TEMP() { return getToken(InformixSQLParser.TEMP, 0); }
		public TerminalNode TEMPORARY() { return getToken(InformixSQLParser.TEMPORARY, 0); }
		public TerminalNode THEN() { return getToken(InformixSQLParser.THEN, 0); }
		public TerminalNode TO() { return getToken(InformixSQLParser.TO, 0); }
		public TerminalNode TRANSACTION() { return getToken(InformixSQLParser.TRANSACTION, 0); }
		public TerminalNode TRIGGER() { return getToken(InformixSQLParser.TRIGGER, 0); }
		public TerminalNode UNION() { return getToken(InformixSQLParser.UNION, 0); }
		public TerminalNode UNIQUE() { return getToken(InformixSQLParser.UNIQUE, 0); }
		public TerminalNode UPDATE() { return getToken(InformixSQLParser.UPDATE, 0); }
		public TerminalNode USING() { return getToken(InformixSQLParser.USING, 0); }
		public TerminalNode VACUUM() { return getToken(InformixSQLParser.VACUUM, 0); }
		public TerminalNode VALUES() { return getToken(InformixSQLParser.VALUES, 0); }
		public TerminalNode VIEW() { return getToken(InformixSQLParser.VIEW, 0); }
		public TerminalNode VIRTUAL() { return getToken(InformixSQLParser.VIRTUAL, 0); }
		public TerminalNode WHEN() { return getToken(InformixSQLParser.WHEN, 0); }
		public TerminalNode WHERE() { return getToken(InformixSQLParser.WHERE, 0); }
		public TerminalNode WITH() { return getToken(InformixSQLParser.WITH, 0); }
		public TerminalNode WITHOUT() { return getToken(InformixSQLParser.WITHOUT, 0); }
		public TerminalNode FIRST_VALUE() { return getToken(InformixSQLParser.FIRST_VALUE, 0); }
		public TerminalNode OVER() { return getToken(InformixSQLParser.OVER, 0); }
		public TerminalNode PARTITION() { return getToken(InformixSQLParser.PARTITION, 0); }
		public TerminalNode RANGE() { return getToken(InformixSQLParser.RANGE, 0); }
		public TerminalNode PRECEDING() { return getToken(InformixSQLParser.PRECEDING, 0); }
		public TerminalNode UNBOUNDED() { return getToken(InformixSQLParser.UNBOUNDED, 0); }
		public TerminalNode CURRENT() { return getToken(InformixSQLParser.CURRENT, 0); }
		public TerminalNode FOLLOWING() { return getToken(InformixSQLParser.FOLLOWING, 0); }
		public TerminalNode CUME_DIST() { return getToken(InformixSQLParser.CUME_DIST, 0); }
		public TerminalNode DENSE_RANK() { return getToken(InformixSQLParser.DENSE_RANK, 0); }
		public TerminalNode LAG() { return getToken(InformixSQLParser.LAG, 0); }
		public TerminalNode LAST_VALUE() { return getToken(InformixSQLParser.LAST_VALUE, 0); }
		public TerminalNode LEAD() { return getToken(InformixSQLParser.LEAD, 0); }
		public TerminalNode NTH_VALUE() { return getToken(InformixSQLParser.NTH_VALUE, 0); }
		public TerminalNode NTILE() { return getToken(InformixSQLParser.NTILE, 0); }
		public TerminalNode PERCENT_RANK() { return getToken(InformixSQLParser.PERCENT_RANK, 0); }
		public TerminalNode RANK() { return getToken(InformixSQLParser.RANK, 0); }
		public TerminalNode ROW_NUMBER() { return getToken(InformixSQLParser.ROW_NUMBER, 0); }
		public TerminalNode GENERATED() { return getToken(InformixSQLParser.GENERATED, 0); }
		public TerminalNode ALWAYS() { return getToken(InformixSQLParser.ALWAYS, 0); }
		public TerminalNode STORED() { return getToken(InformixSQLParser.STORED, 0); }
		public TerminalNode TRUE() { return getToken(InformixSQLParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(InformixSQLParser.FALSE, 0); }
		public TerminalNode WINDOW() { return getToken(InformixSQLParser.WINDOW, 0); }
		public TerminalNode NULLS() { return getToken(InformixSQLParser.NULLS, 0); }
		public TerminalNode FIRST() { return getToken(InformixSQLParser.FIRST, 0); }
		public TerminalNode LAST() { return getToken(InformixSQLParser.LAST, 0); }
		public TerminalNode FILTER() { return getToken(InformixSQLParser.FILTER, 0); }
		public TerminalNode GROUPS() { return getToken(InformixSQLParser.GROUPS, 0); }
		public TerminalNode EXCLUDE() { return getToken(InformixSQLParser.EXCLUDE, 0); }
		public KeywordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).enterKeyword(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).exitKeyword(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof InformixSQLParserVisitor ) return ((InformixSQLParserVisitor<? extends T>)visitor).visitKeyword(this);
			else return visitor.visitChildren(this);
		}
	}

	public final KeywordContext keyword() throws RecognitionException {
		KeywordContext _localctx = new KeywordContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_keyword);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(45);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & -33554432L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & -1152921504606846977L) != 0) || ((((_la - 129)) & ~0x3f) == 0 && ((1L << (_la - 129)) & 9007199254740991L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RoleNameContext extends ParserRuleContext {
		public AnyNameContext anyName() {
			return getRuleContext(AnyNameContext.class,0);
		}
		public RoleNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_roleName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).enterRoleName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).exitRoleName(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof InformixSQLParserVisitor ) return ((InformixSQLParserVisitor<? extends T>)visitor).visitRoleName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final RoleNameContext roleName() throws RecognitionException {
		RoleNameContext _localctx = new RoleNameContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_roleName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			anyName();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AnyNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(InformixSQLParser.IDENTIFIER, 0); }
		public KeywordContext keyword() {
			return getRuleContext(KeywordContext.class,0);
		}
		public TerminalNode STRING_LITERAL() { return getToken(InformixSQLParser.STRING_LITERAL, 0); }
		public TerminalNode OPEN_PAR() { return getToken(InformixSQLParser.OPEN_PAR, 0); }
		public AnyNameContext anyName() {
			return getRuleContext(AnyNameContext.class,0);
		}
		public TerminalNode CLOSE_PAR() { return getToken(InformixSQLParser.CLOSE_PAR, 0); }
		public AnyNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_anyName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).enterAnyName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof InformixSQLParserListener ) ((InformixSQLParserListener)listener).exitAnyName(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof InformixSQLParserVisitor ) return ((InformixSQLParserVisitor<? extends T>)visitor).visitAnyName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AnyNameContext anyName() throws RecognitionException {
		AnyNameContext _localctx = new AnyNameContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_anyName);
		try {
			setState(56);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(49);
				match(IDENTIFIER);
				}
				break;
			case ABORT:
			case ACTION:
			case ADD:
			case AFTER:
			case ALL:
			case ALTER:
			case ANALYZE:
			case AND:
			case AS:
			case ASC:
			case ATTACH:
			case AUTOINCREMENT:
			case BEFORE:
			case BEGIN:
			case BETWEEN:
			case BY:
			case CASCADE:
			case CASE:
			case CAST:
			case CHECK:
			case COLLATE:
			case COLUMN:
			case COMMIT:
			case CONFLICT:
			case CONSTRAINT:
			case CREATE:
			case CROSS:
			case CURRENT_DATE:
			case CURRENT_TIME:
			case CURRENT_TIMESTAMP:
			case DATABASE:
			case DEFAULT:
			case DEFERRABLE:
			case DEFERRED:
			case DELETE:
			case DESC:
			case DETACH:
			case DISTINCT:
			case DROP:
			case EACH:
			case ELSE:
			case END:
			case ESCAPE:
			case EXCEPT:
			case EXCLUSIVE:
			case EXISTS:
			case EXPLAIN:
			case FAIL:
			case FOR:
			case FOREIGN:
			case FROM:
			case FULL:
			case GLOB:
			case GROUP:
			case HAVING:
			case IF:
			case IGNORE:
			case IMMEDIATE:
			case IN:
			case INDEX:
			case INDEXED:
			case INITIALLY:
			case INNER:
			case INSERT:
			case INSTEAD:
			case INTERSECT:
			case INTO:
			case IS:
			case ISNULL:
			case JOIN:
			case KEY:
			case LEFT:
			case LIKE:
			case LIMIT:
			case MATCH:
			case NATURAL:
			case NO:
			case NOT:
			case NOTNULL:
			case NULL:
			case OF:
			case OFFSET:
			case ON:
			case OR:
			case ORDER:
			case OUTER:
			case PLAN:
			case PRAGMA:
			case PRIMARY:
			case QUERY:
			case RAISE:
			case RECURSIVE:
			case REFERENCES:
			case REGEXP:
			case REINDEX:
			case RELEASE:
			case RENAME:
			case REPLACE:
			case RESTRICT:
			case RIGHT:
			case ROLLBACK:
			case ROW:
			case ROWS:
			case SAVEPOINT:
			case SELECT:
			case SET:
			case TABLE:
			case TEMP:
			case TEMPORARY:
			case THEN:
			case TO:
			case TRANSACTION:
			case TRIGGER:
			case UNION:
			case UNIQUE:
			case UPDATE:
			case USING:
			case VACUUM:
			case VALUES:
			case VIEW:
			case VIRTUAL:
			case WHEN:
			case WHERE:
			case WITH:
			case WITHOUT:
			case FIRST_VALUE:
			case OVER:
			case PARTITION:
			case RANGE:
			case PRECEDING:
			case UNBOUNDED:
			case CURRENT:
			case FOLLOWING:
			case CUME_DIST:
			case DENSE_RANK:
			case LAG:
			case LAST_VALUE:
			case LEAD:
			case NTH_VALUE:
			case NTILE:
			case PERCENT_RANK:
			case RANK:
			case ROW_NUMBER:
			case GENERATED:
			case ALWAYS:
			case STORED:
			case TRUE:
			case FALSE:
			case WINDOW:
			case NULLS:
			case FIRST:
			case LAST:
			case FILTER:
			case GROUPS:
			case EXCLUDE:
				enterOuterAlt(_localctx, 2);
				{
				setState(50);
				keyword();
				}
				break;
			case STRING_LITERAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(51);
				match(STRING_LITERAL);
				}
				break;
			case OPEN_PAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(52);
				match(OPEN_PAR);
				setState(53);
				anyName();
				setState(54);
				match(CLOSE_PAR);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u00c2;\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0001\u0000\u0005\u0000\u0010"+
		"\b\u0000\n\u0000\f\u0000\u0013\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001"+
		"\u0001\u0001\u0003\u0001\u0019\b\u0001\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002\"\b\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0003\u0003*\b\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u00069\b\u0006\u0001\u0006"+
		"\u0000\u0000\u0007\u0000\u0002\u0004\u0006\b\n\f\u0000\u0001\u0003\u0000"+
		"\u0019{}\u007f\u0081\u00b5:\u0000\u0011\u0001\u0000\u0000\u0000\u0002"+
		"\u0018\u0001\u0000\u0000\u0000\u0004\u001c\u0001\u0000\u0000\u0000\u0006"+
		"%\u0001\u0000\u0000\u0000\b-\u0001\u0000\u0000\u0000\n/\u0001\u0000\u0000"+
		"\u0000\f8\u0001\u0000\u0000\u0000\u000e\u0010\u0003\u0002\u0001\u0000"+
		"\u000f\u000e\u0001\u0000\u0000\u0000\u0010\u0013\u0001\u0000\u0000\u0000"+
		"\u0011\u000f\u0001\u0000\u0000\u0000\u0011\u0012\u0001\u0000\u0000\u0000"+
		"\u0012\u0014\u0001\u0000\u0000\u0000\u0013\u0011\u0001\u0000\u0000\u0000"+
		"\u0014\u0015\u0005\u0000\u0000\u0001\u0015\u0001\u0001\u0000\u0000\u0000"+
		"\u0016\u0019\u0003\u0004\u0002\u0000\u0017\u0019\u0003\u0006\u0003\u0000"+
		"\u0018\u0016\u0001\u0000\u0000\u0000\u0018\u0017\u0001\u0000\u0000\u0000"+
		"\u0019\u001a\u0001\u0000\u0000\u0000\u001a\u001b\u0005\u0001\u0000\u0000"+
		"\u001b\u0003\u0001\u0000\u0000\u0000\u001c\u001d\u00052\u0000\u0000\u001d"+
		"!\u0005\u0080\u0000\u0000\u001e\u001f\u0005P\u0000\u0000\u001f \u0005"+
		"f\u0000\u0000 \"\u0005F\u0000\u0000!\u001e\u0001\u0000\u0000\u0000!\""+
		"\u0001\u0000\u0000\u0000\"#\u0001\u0000\u0000\u0000#$\u0003\n\u0005\u0000"+
		"$\u0005\u0001\u0000\u0000\u0000%&\u0005?\u0000\u0000&)\u0005\u0080\u0000"+
		"\u0000\'(\u0005P\u0000\u0000(*\u0005F\u0000\u0000)\'\u0001\u0000\u0000"+
		"\u0000)*\u0001\u0000\u0000\u0000*+\u0001\u0000\u0000\u0000+,\u0003\n\u0005"+
		"\u0000,\u0007\u0001\u0000\u0000\u0000-.\u0007\u0000\u0000\u0000.\t\u0001"+
		"\u0000\u0000\u0000/0\u0003\f\u0006\u00000\u000b\u0001\u0000\u0000\u0000"+
		"19\u0005\u00ba\u0000\u000029\u0003\b\u0004\u000039\u0005\u00bd\u0000\u0000"+
		"45\u0005\u0003\u0000\u000056\u0003\f\u0006\u000067\u0005\u0004\u0000\u0000"+
		"79\u0001\u0000\u0000\u000081\u0001\u0000\u0000\u000082\u0001\u0000\u0000"+
		"\u000083\u0001\u0000\u0000\u000084\u0001\u0000\u0000\u00009\r\u0001\u0000"+
		"\u0000\u0000\u0005\u0011\u0018!)8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}