package parser

import (
	"fmt"
	"strconv"

	"github.com/dzeban/monkey/ast"
	"github.com/dzeban/monkey/lexer"
	"github.com/dzeban/monkey/token"
)

// Parser parses the program by building AST
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// Precedence levels
const (
	_ int = iota
	LOWEST
	EQUALS  // ==
	LTGT    // < >
	SUM     // +
	PRODUCT // *
	PREFIX  // -x or !x
	CALL    // f(x)
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NEQ:      EQUALS,
	token.LT:       LTGT,
	token.GT:       LTGT,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

// New creates new parser with a given Lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.setupParseFns()

	// Read 2 token to set curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) setupParseFns() {
	p.prefixParseFns = map[token.TokenType]prefixParseFn{
		token.IDENT:    p.parseIdentifier,
		token.INT:      p.parseIntegerLiteral,
		token.BANG:     p.parsePrefixExpression,
		token.MINUS:    p.parsePrefixExpression,
		token.TRUE:     p.parseBoolean,
		token.FALSE:    p.parseBoolean,
		token.LPAREN:   p.parseGroupExpression,
		token.IF:       p.parseIfExpression,
		token.FUNCTION: p.parseFunctionExpression,
	}

	p.infixParseFns = map[token.TokenType]infixParseFn{
		token.PLUS:     p.parseInfixExpression,
		token.MINUS:    p.parseInfixExpression,
		token.SLASH:    p.parseInfixExpression,
		token.ASTERISK: p.parseInfixExpression,
		token.EQ:       p.parseInfixExpression,
		token.NEQ:      p.parseInfixExpression,
		token.LT:       p.parseInfixExpression,
		token.GT:       p.parseInfixExpression,
		token.LPAREN:   p.parseCallExpression,
	}
}

// Errors returns the list of parsing errors
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected token %s, got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram builds AST
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Ident = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token:      p.curToken,
		Expression: p.parseExpression(LOWEST),
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		msg := fmt.Sprintf("no prefix parse function for %s found", p.curToken.Type)
		p.errors = append(p.errors, msg)
		return nil
	}

	leftExpr := prefix()
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExpr
		}
		p.nextToken()
		leftExpr = infix(leftExpr)
	}

	return leftExpr
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	val, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("couldn't parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = val
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	expr.Right = p.parseExpression(PREFIX)
	return expr
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expr.Right = p.parseExpression(precedence)
	return expr
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE),
	}
}

func (p *Parser) parseGroupExpression() ast.Expression {
	// Skip LPAREN token
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	expr := &ast.IfExpression{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()

	expr.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expr.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		expr.Alternative = p.parseBlockStatement()
	}

	return expr
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	bs := &ast.BlockStatement{
		Token:      p.curToken, // token.LBRACE
		Statements: []ast.Statement{},
	}
	p.nextToken() // skip LBRACE

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			bs.Statements = append(bs.Statements, stmt)
		}
		p.nextToken() // why?
	}

	return bs
}

func (p *Parser) parseFunctionExpression() ast.Expression {
	fn := &ast.FunctionExpression{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	fn.Params = p.parseFunctionParams()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	fn.Body = p.parseBlockStatement()

	return fn
}

func (p *Parser) parseFunctionParams() []*ast.Identifier {
	var params []*ast.Identifier

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return params
	}

	p.nextToken()

	param := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	params = append(params, param)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // skip comma
		p.nextToken() // advance to the next param

		param := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		params = append(params, param)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return params
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	expr := &ast.CallExpression{Token: p.curToken, Function: function}
	expr.Arguments = p.parseCallArguments()
	return expr
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}
	p.nextToken()

	args = append(args, p.parseExpression(LOWEST))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}
