package lia_cli

const (
	OK                       = 0
	Generic                  = 1
	BotExists                = 2
	FailedToFindRepo         = 3
	BotDownloadFailed        = 4
	OsCallFailed             = 5
	ReplayViewerFailed       = 6
	FailedGettingBotLang     = 7
	PreparingBotFailed       = 8
	CopyingRunScriptFailed   = 9
	FailedToReadConfig       = 10
	FailedToGetEnvironment   = 11
	GameGeneratorFailed      = 12
	FailedToGetLiaJson       = 13
	FailedToGenerateGame     = 14
	ZippingBotFailed         = 15
	FailedToGetLatestRelease = 16
	FailedToReadReleaseFile  = 17
)