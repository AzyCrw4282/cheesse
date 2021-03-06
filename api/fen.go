package api

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	errFENRegexDoesNotMatch              = errors.New("FEN string does not match FEN regexp")
	errFENRankLargerThan8Squares         = errors.New("FEN string has a rank larger than 8 squares")
	errFENDuplicateKing                  = errors.New("FEN string has more than one king of the same color")
	errFENKingMissing                    = errors.New("FEN string is lacking one of the kings")
	errFENImpossibleEnPassant            = errors.New("impossible en passant target square, since there's no pawn of the right color next to it")
	errFENImpossibleBlackCastle          = errors.New("impossible for black to castle since king has moved")
	errFENImpossibleBlackQueensideCastle = errors.New("impossible for black to queenside castle since rook has moved")
	errFENImpossibleBlackKingsideCastle  = errors.New("impossible for black to kingside castle since rook has moved")
	errFENImpossibleWhiteCastle          = errors.New("impossible for white to castle since king has moved")
	errFENImpossibleWhiteQueensideCastle = errors.New("impossible for white to queenside castle since rook has moved")
	errFENImpossibleWhiteKingsideCastle  = errors.New("impossible for white to kingside castle since rook has moved")
	errFENPawnInImpossibleRank           = errors.New("impossible rank for pawn")
	errFENBlackHasMoreThan16Pieces       = errors.New("black has more than 16 pieces")
	errFENWhiteHasMoreThan16Pieces       = errors.New("white has more than 16 pieces")
	// TODO check if King is in checkmate that couldn't have been reached
	// TODO don't allow more than 8 pawns of any color
	// TODO check if both are in check
)

func newGameFromFEN(s string) (game, error) {
	rxFEN := regexp.MustCompile(`^([1-8rnbqkpRNBQKP]{1,8})\/([1-8rnbqkpRNBQKP]{1,8})\/([1-8rnbqkpRNBQKP]{1,8})\/([1-8rnbqkpRNBQKP]{1,8})\/([1-8rnbqkpRNBQKP]{1,8})\/([1-8rnbqkpRNBQKP]{1,8})\/([1-8rnbqkpRNBQKP]{1,8})\/([1-8rnbqkpRNBQKP]{1,8}) ([wb]) ([KQkq]{0,4}|-) ([a-h][36]|-) ([0-9]{1,3}) ([0-9]{1,3})$`)
	matches := rxFEN.FindAllStringSubmatch(s, -1)
	fmt.Print("here",matches)
	if matches == nil {
		return game{}, errFENRegexDoesNotMatch
	}

	fullMoveNumber := atoi(matches[0][13]) // The regex cannot pass a non-number here
	halfMoveClock := atoi(matches[0][12])  // The regex cannot pass a non-number here

	// moveNumber calculation
	moveNumber := 0
	turn := matches[0][9]
	switch turn { // The regex cannot pass a different case here
	case "w":
		moveNumber = 2 * (fullMoveNumber - 1)
	case "b":
		moveNumber = 2*(fullMoveNumber-1) + 1
	}

	// En passant calculation
	isLastMoveEnPassant := false
	enPassantTargetSquare := xy{}
	if matches[0][11] != "-" { // The regex cannot pass a different string here
		isLastMoveEnPassant = true
		enPassantTargetSquare = xy{int(matches[0][11][0] - 'a'), int('8' - matches[0][11][1])}
	}

	// Castling calculation
	castlingMap := map[byte]bool{}
	for i := 0; i < len(matches[0][10]); i++ {
		castlingMap[matches[0][10][i]] = true // matched returns 2d array vals and the next one indexes within that elem
	}
	canWhiteKingsideCastle := castlingMap['K']
	canWhiteQueensideCastle := castlingMap['Q']
	canWhiteCastle := canWhiteKingsideCastle || canWhiteQueensideCastle
	canBlackKingsideCastle := castlingMap['k']
	canBlackQueensideCastle := castlingMap['q']
	canBlackCastle := canBlackKingsideCastle || canBlackQueensideCastle

	// Pieces and kings calculation of hashmap use
	pieceTypeMap := map[byte]pieceType{'Q': pieceQueen, 'K': pieceKing, 'B': pieceBishop, 'N': pieceKnight, 'R': pieceRook, 'P': piecePawn}
	pieces := []map[xy]piece{{}, {}}//an empty slice of arrays of maps, mainly used for the two users
	kings := []piece{{}, {}}//slice of piece struct
	for y, row := range []string{matches[0][1], matches[0][2], matches[0][3], matches[0][4], matches[0][5], matches[0][6], matches[0][7], matches[0][8]} {
		x := 0
		for i := 0; i < len(row); i++ {
			b := row[i]
			if x >= 8 {
				return game{}, errFENRankLargerThan8Squares
			}
			switch b {
			case 'Q', 'K', 'B', 'N', 'R', 'P':
				pieces[colorWhite][xy{x, y}] = piece{pieceType: pieceTypeMap[b], owner: colorWhite, xy: xy{x, y}}
				x++
			case 'q', 'k', 'b', 'n', 'r', 'p':
				pieces[colorBlack][xy{x, y}] = piece{pieceType: pieceTypeMap[b-'a'+'A'], owner: colorBlack, xy: xy{x, y}}
				x++
			case '1', '2', '3', '4', '5', '6', '7', '8':
				x += int(b - '0')
			}
			switch {
			case b == 'k' && kings[colorBlack].pieceType == pieceKing, b == 'K' && kings[colorWhite].pieceType == pieceKing:
				return game{}, errFENDuplicateKing
			case b == 'K':
				kings[colorWhite] = pieces[colorWhite][xy{x - 1, y}]
			case b == 'k':
				kings[colorBlack] = pieces[colorBlack][xy{x - 1, y}]
			case (b == 'p' || b == 'P') && (y == 0 || y == 7):
				return game{}, errFENPawnInImpossibleRank
			}
		}
	}
	if kings[colorBlack].pieceType == pieceNone || kings[colorWhite].pieceType == pieceNone {
		return game{}, errFENKingMissing
	}
	if len(pieces[colorBlack]) > 16 {
		return game{}, errFENBlackHasMoreThan16Pieces
	}
	if len(pieces[colorWhite]) > 16 {
		return game{}, errFENWhiteHasMoreThan16Pieces
	}

	// En passant validation
	if isLastMoveEnPassant && turn == "b" && pieces[colorWhite][enPassantTargetSquare.add(xy{0, -1})].pieceType != piecePawn {
		return game{}, errFENImpossibleEnPassant
	}
	if isLastMoveEnPassant && turn == "w" && pieces[colorBlack][enPassantTargetSquare.add(xy{0, 1})].pieceType != piecePawn {
		return game{}, errFENImpossibleEnPassant
	}

	// Castling validation
	if canBlackCastle && !kings[colorBlack].xy.eq(xy{4, 0}) {
		return game{}, errFENImpossibleBlackCastle
	}
	if canBlackQueensideCastle && pieces[colorBlack][xy{0, 0}].pieceType != pieceRook {
		return game{}, errFENImpossibleBlackQueensideCastle
	}
	if canBlackKingsideCastle && pieces[colorBlack][xy{7, 0}].pieceType != pieceRook {
		return game{}, errFENImpossibleBlackKingsideCastle
	}
	if canWhiteCastle && !kings[colorWhite].xy.eq(xy{4, 7}) {
		return game{}, errFENImpossibleWhiteCastle
	}
	if canWhiteQueensideCastle && pieces[colorWhite][xy{0, 7}].pieceType != pieceRook {
		return game{}, errFENImpossibleWhiteQueensideCastle
	}
	if canWhiteKingsideCastle && pieces[colorWhite][xy{7, 7}].pieceType != pieceRook {
		return game{}, errFENImpossibleWhiteKingsideCastle
	}

	game := game{
		canWhiteCastle:          canWhiteCastle,
		canWhiteKingsideCastle:  canWhiteKingsideCastle,
		canWhiteQueensideCastle: canWhiteQueensideCastle,
		canBlackCastle:          canBlackCastle,
		canBlackKingsideCastle:  canBlackKingsideCastle,
		canBlackQueensideCastle: canBlackQueensideCastle,
		halfMoveClock:           halfMoveClock,
		fullMoveNumber:          fullMoveNumber,
		isLastMoveEnPassant:     isLastMoveEnPassant,
		enPassantTargetSquare:   enPassantTargetSquare,
		moveNumber:              moveNumber,
		pieces:                  pieces,
		kings:                   kings,
	}

	return game.calculateCriticalFlags(), nil
}

// This assumes that Atoi can't fail because the regex capture cannot return a non-number
func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func (g game) toFEN() string {
	var sb strings.Builder
	pieceTypeMap := map[pieceType]byte{pieceQueen: 'Q', pieceKing: 'K', pieceBishop: 'B', pieceKnight: 'N', pieceRook: 'R', piecePawn: 'P'}
	for y := 0; y < 8; y++ {
		count := 0
		for x := 0; x < 8; x++ {
			bp, bExists := g.pieces[colorBlack][xy{x, y}]
			wp, wExists := g.pieces[colorWhite][xy{x, y}]
			switch {
			case !bExists && !wExists:
				count++
			case bExists:
				if count > 0 {
					sb.WriteString(fmt.Sprintf("%v", count))
				}
				count = 0
				sb.WriteString(strings.ToLower(string(pieceTypeMap[bp.pieceType])))
			case wExists:
				if count > 0 {
					sb.WriteString(fmt.Sprintf("%v", count))
				}
				count = 0
				sb.WriteByte(pieceTypeMap[wp.pieceType])
			}
		}
		if count > 0 {
			sb.WriteString(fmt.Sprintf("%v", count))
		}
		if y < 7 {
			sb.WriteByte('/')
		}
	}

	turn := "b"
	if g.turn() == colorWhite {
		turn = "w"
	}

	var castlingSB strings.Builder
	if g.canWhiteKingsideCastle {
		castlingSB.WriteByte('K')
	}
	if g.canWhiteQueensideCastle {
		castlingSB.WriteByte('Q')
	}
	if g.canBlackKingsideCastle {
		castlingSB.WriteByte('k')
	}
	if g.canBlackQueensideCastle {
		castlingSB.WriteByte('q')
	}
	castling := castlingSB.String()
	if castling == "" {
		castling = "-"
	}

	enPassant := "-"
	if g.isLastMoveEnPassant {
		enPassant = fmt.Sprintf("%v%v", string("abcdefgh"[g.enPassantTargetSquare.x]), 8-g.enPassantTargetSquare.y)
	}

	sb.WriteString(fmt.Sprintf(" %v %v %v %v %v", turn, castling, enPassant, g.halfMoveClock, g.fullMoveNumber))

	return sb.String()
}
