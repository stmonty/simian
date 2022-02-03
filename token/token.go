package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string

}

const (

    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    //Identifiers and Literals
    IDENT = "IDENT"
    INT   = "INT"

    // Operators
    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"
    LT       = "<"
    GT       = ">"

    //Delimiters
    SEMICOLON = ";"
    COMMA     = ","
    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"

    //Keywords
    FUNCTION = "FUNC"
    LET      = "LET"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
