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
		// This is the token stack we should end up with
		// when we read the input
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
			{token.EOF, ""},
		}
	
		lex := New(input)
	
		testLoop(lex, testSet, t)
	})

	t.Run("OneCharTokenSet", func(t *testing.T) {
		input := `!-/*5;
		5 < 10 > 5;
		`
		// This is the token stack we should end up with
		// when we read the input
		testSet := []Test{
			{token.BANG, "!"},
			{token.MINUS, "-"},
			{token.SLASH, "/"},
			{token.ASTERISK, "*"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.GT, ">"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
		}
	
		lex := New(input)
	
		testLoop(lex, testSet, t)
	})

	t.Run("ExtKeywordTokenSet", func(t *testing.T) {
		input := `if (5 < 10) {
			return true;
		} else {
			return false;
		}
		`
		// This is the token stack we should end up with
		// when we read the input
		testSet := []Test{
			{token.IF, "if"},
			{token.LPAREN, "("},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.TRUE, "true"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.ELSE, "else"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.FALSE, "false"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		}
	
		lex := New(input)
	
		testLoop(lex, testSet, t)
	})
}
