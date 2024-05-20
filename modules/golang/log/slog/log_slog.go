// Code generated by define_gene; DO NOT EDIT.
package slog

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/json"
	_ "github.com/ZenLiuCN/engine/modules/golang/fmt"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engine/modules/golang/log"
	_ "github.com/ZenLiuCN/engine/modules/golang/time"
	"log/slog"
)

var (
	//go:embed log_slog.d.ts
	LogSlogDefine   []byte
	LogSlogDeclared = map[string]any{
		"int":            slog.Int,
		"LevelWarn":      slog.LevelWarn,
		"With":           slog.With,
		"KindBool":       slog.KindBool,
		"time":           slog.Time,
		"debugContext":   slog.DebugContext,
		"groupValue":     slog.GroupValue,
		"KindDuration":   slog.KindDuration,
		"KindLogValuer":  slog.KindLogValuer,
		"log":            slog.Log,
		"MessageKey":     slog.MessageKey,
		"warn":           slog.Warn,
		"info":           slog.Info,
		"KindInt64":      slog.KindInt64,
		"logAttrs":       slog.LogAttrs,
		"warnContext":    slog.WarnContext,
		"float64":        slog.Float64,
		"intValue":       slog.IntValue,
		"KindUint64":     slog.KindUint64,
		"KindString":     slog.KindString,
		"LevelDebug":     slog.LevelDebug,
		"setDefault":     slog.SetDefault,
		"SourceKey":      slog.SourceKey,
		"uint64Value":    slog.Uint64Value,
		"TimeKey":        slog.TimeKey,
		"any":            slog.Any,
		"duration":       slog.Duration,
		"durationValue":  slog.DurationValue,
		"int64Value":     slog.Int64Value,
		"KindFloat64":    slog.KindFloat64,
		"KindGroup":      slog.KindGroup,
		"LevelKey":       slog.LevelKey,
		"errorContext":   slog.ErrorContext,
		"int64":          slog.Int64,
		"newLogLogger":   slog.NewLogLogger,
		"newTextHandler": slog.NewTextHandler,
		"stringValue":    slog.StringValue,
		"Default":        slog.Default,
		"timeValue":      slog.TimeValue,
		"anyValue":       slog.AnyValue,
		"KindAny":        slog.KindAny,
		"newJSONHandler": slog.NewJSONHandler,
		"newRecord":      slog.NewRecord,
		"string":         slog.String,
		"group":          slog.Group,
		"KindTime":       slog.KindTime,
		"New":            slog.New,
		"error":          slog.Error,
		"float64Value":   slog.Float64Value,
		"LevelError":     slog.LevelError,
		"bool":           slog.Bool,
		"boolValue":      slog.BoolValue,
		"debug":          slog.Debug,
		"LevelInfo":      slog.LevelInfo,
		"infoContext":    slog.InfoContext,
		"uint64":         slog.Uint64,

		"emptyLevelVar":          engine.Empty[slog.LevelVar],
		"emptyRefLevelVar":       engine.EmptyRefer[slog.LevelVar],
		"refOfLevelVar":          engine.ReferOf[slog.LevelVar],
		"unRefLevelVar":          engine.UnRefer[slog.LevelVar],
		"emptyLogger":            engine.Empty[slog.Logger],
		"emptyRefLogger":         engine.EmptyRefer[slog.Logger],
		"refOfLogger":            engine.ReferOf[slog.Logger],
		"unRefLogger":            engine.UnRefer[slog.Logger],
		"emptySource":            engine.Empty[slog.Source],
		"emptyRefSource":         engine.EmptyRefer[slog.Source],
		"refOfSource":            engine.ReferOf[slog.Source],
		"unRefSource":            engine.UnRefer[slog.Source],
		"emptyValue":             engine.Empty[slog.Value],
		"emptyRefValue":          engine.EmptyRefer[slog.Value],
		"refOfValue":             engine.ReferOf[slog.Value],
		"unRefValue":             engine.UnRefer[slog.Value],
		"emptyHandlerOptions":    engine.Empty[slog.HandlerOptions],
		"emptyRefHandlerOptions": engine.EmptyRefer[slog.HandlerOptions],
		"refOfHandlerOptions":    engine.ReferOf[slog.HandlerOptions],
		"unRefHandlerOptions":    engine.UnRefer[slog.HandlerOptions],
		"emptyAttr":              engine.Empty[slog.Attr],
		"emptyRefAttr":           engine.EmptyRefer[slog.Attr],
		"refOfAttr":              engine.ReferOf[slog.Attr],
		"unRefAttr":              engine.UnRefer[slog.Attr]}
)

func init() {
	engine.RegisterModule(LogSlogModule{})
}

type LogSlogModule struct{}

func (S LogSlogModule) Identity() string {
	return "golang/log/slog"
}
func (S LogSlogModule) TypeDefine() []byte {
	return LogSlogDefine
}
func (S LogSlogModule) Exports() map[string]any {
	return LogSlogDeclared
}