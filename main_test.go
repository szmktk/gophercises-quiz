package main

import "testing"

type MockConfigurator struct {
	FileName  string
	TimeLimit int
	Shuffle   bool
	Called    bool
}

func (m *MockConfigurator) ParseUserInput() (string, int, bool) {
	m.Called = true
	return m.FileName, m.TimeLimit, m.Shuffle
}

type MockReader struct {
	Questions []Question
	Called    bool
}

func (m *MockReader) ReadQuestions(fileName string) []Question {
	m.Called = true
	return m.Questions
}

type MockScrambler struct {
	Called bool
}

func (m *MockScrambler) Scramble(data []Question) {
	m.Called = true
}

type MockRunner struct {
	Called bool
}

func (m *MockRunner) MainLoop(questions []Question, timeLimit int) {
	m.Called = true
}

func Test_runQuizApp(t *testing.T) {
	// given
	mockConfigurator := &MockConfigurator{
		FileName:  "test.csv",
		TimeLimit: 30,
		Shuffle:   true,
	}
	mockReader := &MockReader{
		Questions: []Question{
			{question: "5+5", answer: "10"},
			{question: "3+4", answer: "7"},
		},
	}
	mockScrambler := &MockScrambler{}
	mockRunner := &MockRunner{}

	// when
	runQuizApp(mockConfigurator, mockReader, mockScrambler, mockRunner)

	// then
	if !mockConfigurator.Called {
		t.Errorf("Expected ParseUserInput to be called")
	}
	if !mockReader.Called {
		t.Errorf("Expected ReadQuestions to be called")
	}

	if !mockScrambler.Called {
		t.Errorf("Expected Scramble to be called")
	}

	if !mockRunner.Called {
		t.Errorf("Expected MainLoop to be called")
	}

}
