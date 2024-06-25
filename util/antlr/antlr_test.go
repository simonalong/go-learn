package test

import (
	"fmt"
	"testing"

	//"github.com/antlr/antlr4/runtime/Go/antlr"
	//"github.com/antlr/antlr4/runtime/!go/antlr"
	antlr "github.com/antlr4-go/antlr/v4"
	"go-learn/util/antlr/parser"
	//"go-hocon/parser"
)

type TreeShapeListener struct {
	*parser.BaseJSONListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func TestName(t *testing.T) {
	input, _ := antlr.NewFileStream("/Users/zhouzhenyong/project/private-go/go-learn/util/antlr/test.json")
	lexer := parser.NewJSONLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJSONParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Json()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
