package product

import (
	"fmt"
	"io"

	"github.com/mattermost/logr/v2"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
	"github.com/sirupsen/logrus"
)

// LogrusHook is a logrus.Hook for emitting plugin logs through the RPC API for inclusion in the
// server logs.
//
// To configure the default Logrus logger for use with plugin logging, simply invoke:
//
//   pluginapi.ConfigureLogrus(logrus.StandardLogger(), pluginAPIClient)
//
// Alternatively, construct your own logger to pass to pluginapi.ConfigureLogrus.
type LogrusHook struct {
	log mlog.LoggerIFace
}

// NewLogrusHook creates a new instance of LogrusHook.
func NewLogrusHook(log mlog.LoggerIFace) *LogrusHook {
	return &LogrusHook{
		log: log,
	}
}

// Levels allows LogrusHook to process any log level.
func (lh *LogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire proxies logrus entries through the plugin API at the appropriate level.
func (lh *LogrusHook) Fire(entry *logrus.Entry) error {
	fields := []logr.Field{}
	for key, value := range entry.Data {
		fields = append(fields,
			logr.Field{
				Key:       key,
				Interface: value,
				//TODO: add type check and more fields
			})
	}

	if entry.Caller != nil {
		fields = append(fields,
			logr.Field{
				Key:    "plugin_caller",
				String: fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line),
			})
	}

	switch entry.Level {
	case logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel:
		lh.log.Error(entry.Message, fields...)
	case logrus.WarnLevel:
		lh.log.Warn(entry.Message, fields...)
	case logrus.InfoLevel:
		lh.log.Info(entry.Message, fields...)
	case logrus.DebugLevel, logrus.TraceLevel:
		lh.log.Debug(entry.Message, fields...)
	}

	return nil
}

// ConfigureLogrus configures the given logrus logger with a hook to proxy through the RPC API,
// discarding the default output to avoid duplicating the events across the standard STDOUT proxy.
func ConfigureLogrus(logger *logrus.Logger, log mlog.LoggerIFace) {
	hook := NewLogrusHook(log)
	logger.Hooks.Add(hook)
	logger.SetOutput(io.Discard)
	logrus.SetReportCaller(true)

	// By default, log everything to the server, and let it decide what gets through.
	logrus.SetLevel(logrus.TraceLevel)
}
