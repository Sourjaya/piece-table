package piece_table

type Piece struct {
	source bool
	start  int
	length int
}
type Pieces struct {
	pieces []Piece
}
type Piece_Table struct {
	file []rune
	add  []rune
}
