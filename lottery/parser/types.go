package parser

type Parser interface {
	Ssq(byte []byte) (ParseResult, error)
	Dlt(byte []byte) (ParseResult, error)
	Fc3d(byte []byte) (ParseResult, error)
	Q3(byte []byte) (ParseResult, error)
	// ...
}

type ParseResult struct {
	codeCountMap map[string]int
}