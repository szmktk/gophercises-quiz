package main

type QuizConfigurator interface {
	ParseUserInput() (fileName string, timeLimit int, shuffle bool)
}

type QuestionReader interface {
	ReadQuestions(fileName string) []Question
}

type QuestionScrambler interface {
	Scramble(questions []Question)
}

type QuizRunner interface {
	MainLoop(questions []Question, timeLimit int)
}
