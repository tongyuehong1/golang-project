package logger

type LogReporterInterface interface {
	ReportError(error)
}

type LogReport struct {
	provider LogReporterInterface
}

var GlobalLogReporter *LogReport

func (manager *LogReport) Error(err error) {
	manager.provider.ReportError(err)
}
