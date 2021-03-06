package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKingActions(t *testing.T) {
	ts := []struct {
		name    string
		board   board
		actions []action
		color   color
		xy      xy
	}{
		{
			name: "white king: all actions",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞♜",
					"♟♟♟♟♟♟♟♟",
					"        ",
					"        ",
					"  ♙♔    ",
					"        ",
					"♙♙  ♙♙♙♙",
					"♖♘♗♕ ♗♘♖",
				},
				turn: "White",
			},
			color: colorWhite,
			xy:    xy{3, 4},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{3, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 4}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{3, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 5}},
			},
		},
		{
			name: "white king: only movements that escape check",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞♜",
					"♟♟♟ ♟♟♟♟",
					"        ",
					"        ",
					"  ♙♔    ",
					"        ",
					"♙♙  ♙♙♙♙",
					"♖♘♗♕ ♗♘♖",
				},
				turn: "White",
			},
			color: colorWhite,
			xy:    xy{3, 4},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 4}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 5}},
			},
		},
		{
			name: "white king: only movements that don't put king in check by rook",
			board: board{
				board: []string{
					" ♞♝♛♚♝♞ ",
					" ♟♟♟♟♟♟ ",
					"  ♜ ♜   ",
					"        ",
					"  ♙♔    ",
					"        ",
					"♙♙  ♙♙♙♙",
					"♖♘♗♕ ♗♘♖",
				},
				turn: "White",
			},
			color: colorWhite,
			xy:    xy{3, 4},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{3, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{3, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 5}},
			},
		},
		{
			name: "white king: only movements that don't put king in check by bishop",
			board: board{
				board: []string{
					"♜♞ ♛♚ ♞♜",
					"♟ ♟♟♟♟ ♟",
					"  ♝ ♝   ",
					"        ",
					"  ♙♔    ",
					"        ",
					"♙♙  ♙♙♙♙",
					"♖♘♗♕ ♗♘♖",
				},
				turn: "White",
			},
			color: colorWhite,
			xy:    xy{3, 4},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{3, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 5}},
			},
		},
		{
			name: "white king: only movements that don't put king in check by knight",
			board: board{
				board: []string{
					"♜ ♝♛♚♝♞♜",
					"♟♟♟♟♟♟♟♟",
					"   ♞    ",
					"        ",
					"  ♙♔    ",
					"        ",
					"♙♙  ♙♙♙♙",
					"♖♘♗♕ ♗♘♖",
				},
				turn: "White",
			},
			color: colorWhite,
			xy:    xy{3, 4},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{3, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{3, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 5}},
			},
		},
		{
			name: "white king: can capture opponent piece threatening it",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞♜",
					"♟♟ ♟♟♟♟♟",
					"        ",
					"  ♟     ",
					"  ♙♔    ",
					"        ",
					"♙♙  ♙♙♙♙",
					"♖♘♗♕ ♗♘♖",
				},
				turn: "White",
			},
			color: colorWhite,
			xy:    xy{3, 4},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 3}, isCapture: true, capturedPiece: piece{pieceType: piecePawn, owner: colorBlack, xy: xy{2, 3}}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{3, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 3}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 4}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{2, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{3, 5}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{3, 4}}, toXY: xy{4, 5}},
			},
		},
		{
			name: "black king: can castle kingside or move right",
			board: board{
				board: []string{
					"♜♞♝♛♚  ♜",
					"♟♟♟♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙♙♙♙",
					"♖♘♗♕♔♗♘♖",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorBlack,
			xy:    xy{4, 0},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{5, 0}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{6, 0}, isCastle: true, isKingsideCastle: true},
			},
		},
		{
			name: "black king: can't castle kingside due to a square in check",
			board: board{
				board: []string{
					"♜♞♝♛♚  ♜",
					"♟♟♟♟♟ ♟♟",
					"        ",
					"        ",
					"        ",
					"     ♖  ",
					"♙♙♙♙♙♙♙ ",
					"♖♘♗♕♔♗♘ ",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  false,
			},
			color:   colorBlack,
			xy:      xy{4, 0},
			actions: []action{},
		},
		{
			name: "black king: can't castle kingside due to another square in check",
			board: board{
				board: []string{
					"♜♞♝♛♚  ♜",
					"♟♟♟♟♟♟ ♟",
					"        ",
					"        ",
					"        ",
					"      ♖ ",
					"♙♙♙♙♙♙♙ ",
					"♖♘♗♕♔♗♘ ",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  false,
			},
			color: colorBlack,
			xy:    xy{4, 0},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{5, 0}},
			},
		},
		{
			name: "black king: can't castle kingside due to being in check",
			board: board{
				board: []string{
					"♜♞♝♛♚  ♜",
					"♟♟♟♟ ♟♟♟",
					"        ",
					"        ",
					"        ",
					"    ♖   ",
					"♙♙♙♙♙♙♙ ",
					"♖♘♗♕♔♗♘ ",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  false,
			},
			color: colorBlack,
			xy:    xy{4, 0},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{5, 0}},
			},
		},
		{
			name: "black king: can castle queenside or move left",
			board: board{
				board: []string{
					"♜   ♚♝♞♜",
					"♟♟♟♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙♙♙♙",
					"♖♘♗♕♔♗♘♖",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorBlack,
			xy:    xy{4, 0},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{3, 0}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{2, 0}, isCastle: true, isQueensideCastle: true},
			},
		},
		{
			name: "black king: can't castle queenside because there's a square in check",
			board: board{
				board: []string{
					"♜   ♚♝♞♜",
					"♟ ♟♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					" ♖      ",
					" ♙♙♙♙♙♙♙",
					" ♘♗♕♔♗♘♖",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: false,
				canWhiteKingsideCastle:  true,
			},
			color: colorBlack,
			xy:    xy{4, 0},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{3, 0}},
			},
		},
		{
			name: "black king: can't castle queenside because there's another square in check",
			board: board{
				board: []string{
					"♜   ♚♝♞♜",
					"♟♟ ♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"  ♖     ",
					" ♙♙♙♙♙♙♙",
					" ♘♗♕♔♗♘♖",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: false,
				canWhiteKingsideCastle:  true,
			},
			color: colorBlack,
			xy:    xy{4, 0},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{3, 0}},
			},
		},
		{
			name: "black king: can't castle queenside because there's a third square in check",
			board: board{
				board: []string{
					"♜   ♚♝♞♜",
					"♟♟  ♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"   ♖    ",
					" ♙♙♙♙♙♙♙",
					" ♘♗♕♔♗♘♖",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: false,
				canWhiteKingsideCastle:  true,
			},
			color:   colorBlack,
			xy:      xy{4, 0},
			actions: []action{},
		},
		{
			name: "black king: can't castle queenside because the king is in check",
			board: board{
				board: []string{
					"♜   ♚♝♞♜",
					"♟♟♟♟ ♟♟♟",
					"        ",
					"        ",
					"        ",
					"    ♖   ",
					" ♙♙♙♙♙♙♙",
					" ♘♗♕♔♗♘♖",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: false,
				canWhiteKingsideCastle:  true,
			},
			color: colorBlack,
			xy:    xy{4, 0},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{3, 0}},
			},
		},
		{
			name: "black king: can't castle kingside because it's not available anymore",
			board: board{
				board: []string{
					"♜♞♝♛♚  ♜",
					"♟♟♟♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙♙♙♙",
					"♖♘♗♕♔♗♘♖",
				},
				turn:                    "Black",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  false,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorBlack,
			xy:    xy{4, 0},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{5, 0}},
			},
		},
		{
			name: "black king: can't castle queenside because it's not available anymore",
			board: board{
				board: []string{
					"♜   ♚♝♞♜",
					"♟♟♟♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙♙♙♙",
					"♖♘♗♕♔♗♘♖",
				},
				turn:                    "Black",
				canBlackQueensideCastle: false,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorBlack,
			xy:    xy{4, 0},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorBlack, xy: xy{4, 0}}, toXY: xy{3, 0}},
			},
		},
		{
			name: "white king: can castle kingside or move right",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞♜",
					"♟♟♟♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙♙♙♙",
					"♖♘♗♕♔  ♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorWhite,
			xy:    xy{4, 7},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{5, 7}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{6, 7}, isCastle: true, isKingsideCastle: true},
			},
		},
		{
			name: "white king: can't castle kingside because there's a square in check",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞ ",
					"♟♟♟♟♟♟♟ ",
					"      ♜ ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙♙ ♙",
					"♖♘♗♕♔  ♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  false,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorWhite,
			xy:    xy{4, 7},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{5, 7}},
			},
		},
		{
			name: "white king: can't castle kingside because there's another square in check",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞ ",
					"♟♟♟♟♟♟♟ ",
					"     ♜  ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙ ♙♙",
					"♖♘♗♕♔  ♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  false,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color:   colorWhite,
			xy:      xy{4, 7},
			actions: []action{},
		},
		{
			name: "white king: can't castle kingside because the king is in check",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞ ",
					"♟♟♟♟♟♟♟ ",
					"    ♜   ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙ ♙♙♙",
					"♖♘♗♕♔  ♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  false,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorWhite,
			xy:    xy{4, 7},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{5, 7}},
			},
		},
		{
			name: "white king: can castle queenside or move left",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞♜",
					"♟♟♟♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙♙♙♙",
					"♖   ♔♗♘♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorWhite,
			xy:    xy{4, 7},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{3, 7}},
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{2, 7}, isCastle: true, isQueensideCastle: true},
			},
		},
		{
			name: "white king: can't castle queenside because there's a square in check",
			board: board{
				board: []string{
					" ♞♝♛♚♝♞♜",
					" ♟♟♟♟♟♟♟",
					" ♜      ",
					"        ",
					"        ",
					"        ",
					"♙ ♙♙♙♙♙♙",
					"♖   ♔♗♘♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: false,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorWhite,
			xy:    xy{4, 7},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{3, 7}},
			},
		},
		{
			name: "white king: can't castle queenside because there's another square in check",
			board: board{
				board: []string{
					" ♞♝♛♚♝♞♜",
					" ♟♟♟♟♟♟♟",
					"  ♜     ",
					"        ",
					"        ",
					"        ",
					"♙♙ ♙♙♙♙♙",
					"♖   ♔♗♘♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: false,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorWhite,
			xy:    xy{4, 7},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{3, 7}},
			},
		},
		{
			name: "white king: can't castle queenside because there's a third square in check",
			board: board{
				board: []string{
					" ♞♝♛♚♝♞♜",
					" ♟♟♟♟♟♟♟",
					"   ♜    ",
					"        ",
					"        ",
					"        ",
					"♙♙♙ ♙♙♙♙",
					"♖   ♔♗♘♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: false,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color:   colorWhite,
			xy:      xy{4, 7},
			actions: []action{},
		},
		{
			name: "white king: can't castle queenside because there's a third square in check",
			board: board{
				board: []string{
					" ♞♝♛♚♝♞♜",
					" ♟♟♟♟♟♟♟",
					"    ♜   ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙ ♙♙♙",
					"♖   ♔♗♘♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: false,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  true,
			},
			color: colorWhite,
			xy:    xy{4, 7},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{3, 7}},
			},
		},
		{
			name: "white king: can't castle kingside because it's not available anymore",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞♜",
					"♟♟♟♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙♙♙♙",
					"♖♘♗♕♔  ♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: true,
				canWhiteKingsideCastle:  false,
			},
			color: colorWhite,
			xy:    xy{4, 7},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{5, 7}},
			},
		},
		{
			name: "white king: can't castle queenside because it's not available anymore",
			board: board{
				board: []string{
					"♜♞♝♛♚♝♞♜",
					"♟♟♟♟♟♟♟♟",
					"        ",
					"        ",
					"        ",
					"        ",
					"♙♙♙♙♙♙♙♙",
					"♖   ♔♗♘♖",
				},
				turn:                    "White",
				canBlackQueensideCastle: true,
				canBlackKingsideCastle:  true,
				canWhiteQueensideCastle: false,
				canWhiteKingsideCastle:  true,
			},
			color: colorWhite,
			xy:    xy{4, 7},
			actions: []action{
				{fromPiece: piece{pieceType: pieceKing, owner: colorWhite, xy: xy{4, 7}}, toXY: xy{3, 7}},
			},
		},
	}
	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			g, err := newGameFromBoard(tc.board)
			require.NoError(t, err)
			assert.ElementsMatch(t, tc.actions, g.pieces[tc.color][tc.xy].calculateAllActions(g))
		})
	}
}
