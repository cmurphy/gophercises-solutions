// Code generated by "stringer -type=Suit,Rank"; DO NOT EDIT.

package deck

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Spade-0]
	_ = x[Diamond-1]
	_ = x[Club-2]
	_ = x[Heart-3]
	_ = x[Joker-4]
}

const _Suit_name = "SpadeDiamondClubHeartJoker"

var _Suit_index = [...]uint8{0, 5, 12, 16, 21, 26}

func (i Suit) String() string {
	if i >= Suit(len(_Suit_index)-1) {
		return "Suit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Suit_name[_Suit_index[i]:_Suit_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ace-0]
	_ = x[Two-1]
	_ = x[Three-2]
	_ = x[Four-3]
	_ = x[Five-4]
	_ = x[Six-5]
	_ = x[Seven-6]
	_ = x[Eight-7]
	_ = x[Nine-8]
	_ = x[Ten-9]
	_ = x[Jack-10]
	_ = x[Queen-11]
	_ = x[King-12]
}

const _Rank_name = "AceTwoThreeFourFiveSixSevenEightNineTenJackQueenKing"

var _Rank_index = [...]uint8{0, 3, 6, 11, 15, 19, 22, 27, 32, 36, 39, 43, 48, 52}

func (i Rank) String() string {
	if i >= Rank(len(_Rank_index)-1) {
		return "Rank(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Rank_name[_Rank_index[i]:_Rank_index[i+1]]
}
