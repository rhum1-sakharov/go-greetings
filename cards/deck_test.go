package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const FILE_DECK_TESTING = "file_deck_testing"

type TestSuite struct {
	suite.Suite
	d deck
}

func (suite *TestSuite) SetupTest() {
	suite.d = newDeck()
}

func (s *TestSuite) AfterTest(_, _ string) {
	os.Remove(FILE_DECK_TESTING)
}

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) Test_whenNewDeck_thenReturn16CardsDeck() {
	assert.Equal(t.T(), 16, len(t.d), "deck should return 16 cards")
}

func (t *TestSuite) Test_whenNewDeck_thenFirstCardIsAnAceOfSpades() {
	assert.Equal(t.T(), "Ace of Spades", t.d[0], "first card should be an Ace of Spaces")
}

func (t *TestSuite) Test_whenNewDeck_thenLastCardIsAFourOfClubs() {
	assert.Equal(t.T(), "Four of Clubs", t.d[len(t.d)-1], "last card should be a Four of Clubs")
}

func (t *TestSuite) Test_whenSaveToFile_thenFileIsCreated() {
	t.d.saveToFile(FILE_DECK_TESTING)
	assert.FileExists(t.T(), FILE_DECK_TESTING)
}

func (t *TestSuite) Test_givenFileDeck_whenNewDeckFromFile_thenDeckHas16Cards() {
	// given
	t.d.saveToFile(FILE_DECK_TESTING)

	// when
	t.d = newDeckFromFile(FILE_DECK_TESTING)

	// then
	assert.Equal(t.T(), 16, len(t.d), "deck should have 16 cards")
}
