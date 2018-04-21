package lexer

import (
	"github.com/nicolasjhampton/simple_interpreter/internal/token"
	"testing"
)

// TestTable: Go uses TestTables
// if NextToken took inputs, those would be in the table as well
type Test struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func testLoop(lex *Lexer, testSet []Test, t *testing.T) {
	for i, test := range testSet {
		token := lex.NextToken()

		if token.Type != test.expectedType {
			t.Fatalf("tests[%d]  - tokentype wrong. Expected %q, got %q",
				i, test.expectedType, token.Type)
		}

		if token.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d]  - literal wrong. Expected %q, got %q",
				i, test.expectedLiteral, token.Literal)
		}
	}
}

func TestNextToken(t *testing.T) { 

	t.Run("BasicTokenSet", func(t *testing.T) {
		input := `=+(){},;`

		testSet := []Test{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}
	
		lex := New(input)
	
		testLoop(lex, testSet, t)
	})

	t.Run("RunnableTokenSet", func(t *testing.T) {
		input := `let five = 5;
		let ten = 10;
		
		let add = fn(x, y) {
		  x + y;
		};
		
		let result = add(five, ten);
		`

		testSet := []Test{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""}
		}
	
		lex := New(input)
	
		testLoop(lex, testSet, t)
	})
}
