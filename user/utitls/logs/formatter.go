package logs

import (
	"bytes"
	"fmt"
	sequences "github.com/konsorten/go-windows-terminal-sequences"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sort"
	"sync"
	"syscall"
	"time"
)

const defaultTimestampFormat  = "2006-01-02 15:03:04"

var (
	baseTimestamp time.Time
	emptyFieldMap FieldMap
)

type fieldKey string

// FieldMap allows customization of the key names for default fields.
type FieldMap map[fieldKey]string

//type Fields map[string]interface{}

func init() {
	baseTimestamp = time.Now()
}

// MyTextFormatter formats logs into text
type MyTextFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool

	// Force disabling colors.
	DisableColors bool

	// Override coloring based on CLICOLOR and CLICOLOR_FORCE. - https://bixense.com/clicolors/
	EnvironmentOverrideColors bool

	// Disable timestamp logging. useful when output is redirected to logging
	// system that already adds timestamps.
	DisableTimestamp bool

	// Enable logging the full timestamp when a TTY is attached instead of just
	// the time passed since beginning of execution.
	FullTimestamp bool

	// TimestampFormat to use for display when a full timestamp is printed
	TimestampFormat string

	// The fields are sorted by default for a consistent output. For applications
	// that log extremely frequently and don't use the JSON formatter this may not
	// be desired.
	DisableSorting bool

	// The keys sorting function, when uninitialized it uses sort.Strings.
	SortingFunc func([]string)

	// Disables the truncation of the level text to 4 characters.
	DisableLevelTruncation bool

	// QuoteEmptyFields will wrap empty fields in quotes if true
	QuoteEmptyFields bool

	// Whether the logger's out is to a terminal
	isTerminal bool

	// FieldMap allows users to customize the names of keys for default fields.
	// As an example:
	// formatter := &MyTextFormatter{
	//     FieldMap: FieldMap{
	//         logrus.FieldKeyTime:  "@timestamp",
	//         logrus.FieldKeyLevel: "@level",
	//         logrus.FieldKeyMsg:   "@message"}}
	FieldMap FieldMap

	terminalInitOnce sync.Once
}

func (f *MyTextFormatter) init(entry *logrus.Entry) {
	if entry.Logger != nil {
		f.isTerminal = checkIfTerminal(entry.Logger.Out)

		if f.isTerminal {
			initTerminal(entry.Logger.Out)
		}
	}
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		var mode uint32
		err := syscall.GetConsoleMode(syscall.Handle(v.Fd()), &mode)
		return err == nil
	default:
		return false
	}
}

func initTerminal(w io.Writer) {
	switch v := w.(type) {
	case *os.File:
		sequences.EnableVirtualTerminalProcessing(syscall.Handle(v.Fd()), true)
	}
}

func (f *MyTextFormatter) isColored() bool {
	isColored := f.ForceColors || f.isTerminal

	if f.EnvironmentOverrideColors {
		if force, ok := os.LookupEnv("CLICOLOR_FORCE"); ok && force != "0" {
			isColored = true
		} else if ok && force == "0" {
			isColored = false
		} else if os.Getenv("CLICOLOR") == "0" {
			isColored = false
		}
	}

	return isColored && !f.DisableColors
}

// Format renders a single log entry
func (f *MyTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	prefixFieldClashes(entry.Data, f.FieldMap, entry.HasCaller())

	keys := make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}

	fixedKeys := make([]string, 0, 4+len(entry.Data))
	if !f.DisableTimestamp {
		fixedKeys = append(fixedKeys, f.FieldMap.resolve(logrus.FieldKeyTime))
	}
	fixedKeys = append(fixedKeys, f.FieldMap.resolve(logrus.FieldKeyLevel))
	if entry.Message != "" {
		fixedKeys = append(fixedKeys, f.FieldMap.resolve(logrus.FieldKeyMsg))
	}

	if entry.HasCaller() {
		fixedKeys = append(fixedKeys,
			f.FieldMap.resolve(logrus.FieldKeyFunc), f.FieldMap.resolve(logrus.FieldKeyFile))
	}

	if !f.DisableSorting {
		if f.SortingFunc == nil {
			sort.Strings(keys)
			fixedKeys = append(fixedKeys, keys...)
		} else {
			if !f.isColored() {
				fixedKeys = append(fixedKeys, keys...)
				f.SortingFunc(fixedKeys)
			} else {
				f.SortingFunc(keys)
			}
		}
	} else {
		fixedKeys = append(fixedKeys, keys...)
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.terminalInitOnce.Do(func() { f.init(entry) })

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	for _, key := range fixedKeys {
		var value interface{}
		switch {
		case key == f.FieldMap.resolve(logrus.FieldKeyTime):
			value = entry.Time.Format(timestampFormat)
		case key == f.FieldMap.resolve(logrus.FieldKeyLevel):
			value = entry.Level.String()
		case key == f.FieldMap.resolve(logrus.FieldKeyMsg):
			value = entry.Message
		case key == f.FieldMap.resolve(logrus.FieldKeyFunc) && entry.HasCaller():
			value = entry.Caller.Function
		case key == f.FieldMap.resolve(logrus.FieldKeyFile) && entry.HasCaller():
			value = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		default:
			value = entry.Data[key]
		}
		f.appendKeyValue(b, key, value)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func prefixFieldClashes(data map[string]interface{}, fieldMap FieldMap, reportCaller bool) {
	timeKey := fieldMap.resolve(logrus.FieldKeyTime)
	if t, ok := data[timeKey]; ok {
		data["fields."+timeKey] = t
		delete(data, timeKey)
	}

	msgKey := fieldMap.resolve(logrus.FieldKeyMsg)
	if m, ok := data[msgKey]; ok {
		data["fields."+msgKey] = m
		delete(data, msgKey)
	}

	levelKey := fieldMap.resolve(logrus.FieldKeyLevel)
	if l, ok := data[levelKey]; ok {
		data["fields."+levelKey] = l
		delete(data, levelKey)
	}

	logrusErrKey := fieldMap.resolve(logrus.FieldKeyLogrusError)
	if l, ok := data[logrusErrKey]; ok {
		data["fields."+logrusErrKey] = l
		delete(data, logrusErrKey)
	}

	// If reportCaller is not set, 'func' will not conflict.
	if reportCaller {
		funcKey := fieldMap.resolve(logrus.FieldKeyFunc)
		if l, ok := data[funcKey]; ok {
			data["fields."+funcKey] = l
		}
		fileKey := fieldMap.resolve(logrus.FieldKeyFile)
		if l, ok := data[fileKey]; ok {
			data["fields."+fileKey] = l
		}
	}
}

func (f FieldMap) resolve(key fieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}

	return string(key)
}


func (f *MyTextFormatter) needsQuoting(text string) bool {
	if f.QuoteEmptyFields && len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}
	return false
}

func (f *MyTextFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)
}

func (f *MyTextFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(stringVal)
}
