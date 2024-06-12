package lsp

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PositionsTestSuite struct {
	suite.Suite
}

func (s *PositionsTestSuite) Test_position_byte_index_for_utf16_offset_counting() {
	input := "This is the first line\n" + // 22 bytes
		"This is the second line\n" + // 23 bytes
		"This is the third line\n" + // 22 bytes
		"Fourth line𒀃𐐀here" // 15 utf-16 code points before target char, 19 bytes.
	position := Position{
		Line: 3,
		// "h" in "here"
		Character: 15,
	}
	s.Equal(89, position.IndexIn(input, PositionEncodingKindUTF16))
}

func (s *PositionsTestSuite) Test_position_byte_index_for_utf8_offset_counting() {
	input := "This is the first line\n" + // 22 bytes
		"This is the second line\n" + // 23 bytes
		"This is the third line\n" + // 22 bytes
		"Fourth line𒀃𐐀here" // 19 utf-8 code points before target char, 19 bytes.
	position := Position{
		Line: 3,
		// "h" in "here"
		Character: 19,
	}

	s.Equal(89, position.IndexIn(input, PositionEncodingKindUTF8))
}

func (s *PositionsTestSuite) Test_position_byte_index_for_utf32_offset_counting() {
	input := "This is the first line\n" + // 22 bytes
		"This is the second line\n" + // 23 bytes
		"This is the third line\n" + // 22 bytes
		"Fourth line𒀃𐐀here" // 13 utf-32 code points before target char, 19 bytes.
	position := Position{
		Line: 3,
		// "h" in "here"
		Character: 13,
	}

	s.Equal(89, position.IndexIn(input, PositionEncodingKindUTF32))
}

func (s *PositionsTestSuite) Test_position_byte_index_returns_0_index_for_char_beyond_end_of_line() {
	input := "This is a line\n" +
		"This is a second line\n" +
		"This is a third line\n" +
		"Fourth line𒀃𐐀here"
	position := Position{
		Line: 3,
		// char 150 is beyond the end of the line
		Character: 150,
	}
	s.Equal(0, position.IndexIn(input, PositionEncodingKindUTF16))
}

func (s *PositionsTestSuite) Test_get_position_for_end_of_line_for_current_position() {
	input := "First line\n" + // 10 bytes
		"Second line\n" + // 11 bytes
		"Third line\n" + // 10 bytes
		"Fourth line𒀃𐐀here\n" // 15 utf-16 code points before target char, 19 bytes.

	position := Position{
		Line: 3,
		// "h" in "here"
		Character: 15,
	}
	eolPos := position.EndOfLineIn(input, PositionEncodingKindUTF16)
	s.Equal(
		Position{
			Line:      3,
			Character: 19,
		},
		eolPos,
	)
}

func (s *PositionsTestSuite) Test_range_positions_byte_index() {
	input := "This is the first line\n" + // 22 bytes
		"This is the second line\n" + // 23 bytes
		"This is the third line\n" + // 22 bytes
		"Fourth line𒀃𐐀here" // 15 utf-16 code points before target char, 19 bytes.
	start := Position{
		Line: 3,
		// "h" in "here"
		Character: 15,
	}
	end := Position{
		Line: 3,
		// last "e" in "here"
		Character: 18,
	}
	rnge := Range{
		Start: start,
		End:   end,
	}
	startIdx, endIdx := rnge.IndexesIn(input, PositionEncodingKindUTF16)
	s.Equal(89, startIdx)
	s.Equal(92, endIdx)
}

func TestPositionsTestSuite(t *testing.T) {
	suite.Run(t, new(PositionsTestSuite))
}
