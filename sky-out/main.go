package main

import (
	"bufio"
	"context"
	encoding_json "encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var skyVersion = "dev"

type SkyTuple2 struct{ V0, V1 any }

type SkyTuple3 struct{ V0, V1, V2 any }

type SkyResult struct {
	Tag               int
	SkyName           string
	OkValue, ErrValue any
}

type SkyMaybe struct {
	Tag       int
	SkyName   string
	JustValue any
}

var stdinReader *bufio.Reader
var sky_jsonDecoder_string = func(v any) any {
	if s, ok := v.(string); ok {
		return SkyOk(s)
	}
	return SkyErr("expected string")
}
var sky_jsonDecoder_int = func(v any) any {
	switch n := v.(type) {
	case float64:
		return SkyOk(int(n))
	case int:
		return SkyOk(n)
	}
	return SkyErr("expected int")
}
var sky_jsonDecoder_float = func(v any) any {
	if f, ok := v.(float64); ok {
		return SkyOk(f)
	}
	return SkyErr("expected float")
}
var sky_jsonDecoder_bool = func(v any) any {
	if b, ok := v.(bool); ok {
		return SkyOk(b)
	}
	return SkyErr("expected bool")
}
var sky_liveAppImpl = func(config any) any { return config }
var LogMsg = Log_Webhook_LogMsg
var BuildRules = Log_Webhook_BuildRules
var MatchesFilter = Log_Webhook_MatchesFilter
var BuildPayload = Log_Webhook_BuildPayload
var Send = Log_Webhook_Send
var ProcessNewEntries = Log_Webhook_ProcessNewEntries
var File = Os_File
var LinkError = Os_LinkError
var ProcAttr = Os_ProcAttr
var Process = Os_Process
var ProcessState = Os_ProcessState
var Root = Os_Root
var Signal = Os_Signal
var SyscallError = Os_SyscallError
var fileWriteString = Os_FileWriteString
var FileWriteString = Os_FileWriteString
var stderr = Os_Stderr
var Stderr = Os_Stderr
var stdin = Os_Stdin
var Stdin = Os_Stdin
var command = Os_Exec_Command
var Command = Os_Exec_Command
var commandContext = Os_Exec_CommandContext
var CommandContext = Os_Exec_CommandContext
var lookPath = Os_Exec_LookPath
var LookPath = Os_Exec_LookPath
var Cmd = Os_Exec_Cmd
var Error = Os_Exec_Error
var ExitError = Os_Exec_ExitError
var cmdPath = Os_Exec_CmdPath
var CmdPath = Os_Exec_CmdPath
var cmdArgs = Os_Exec_CmdArgs
var CmdArgs = Os_Exec_CmdArgs
var cmdEnv = Os_Exec_CmdEnv
var CmdEnv = Os_Exec_CmdEnv
var cmdDir = Os_Exec_CmdDir
var CmdDir = Os_Exec_CmdDir
var cmdStdin = Os_Exec_CmdStdin
var CmdStdin = Os_Exec_CmdStdin
var cmdStdout = Os_Exec_CmdStdout
var CmdStdout = Os_Exec_CmdStdout
var cmdStderr = Os_Exec_CmdStderr
var CmdStderr = Os_Exec_CmdStderr
var cmdExtraFiles = Os_Exec_CmdExtraFiles
var CmdExtraFiles = Os_Exec_CmdExtraFiles
var cmdSysProcAttr = Os_Exec_CmdSysProcAttr
var CmdSysProcAttr = Os_Exec_CmdSysProcAttr
var cmdProcess = Os_Exec_CmdProcess
var CmdProcess = Os_Exec_CmdProcess
var cmdProcessState = Os_Exec_CmdProcessState
var CmdProcessState = Os_Exec_CmdProcessState
var cmdErr = Os_Exec_CmdErr
var CmdErr = Os_Exec_CmdErr
var cmdWaitDelay = Os_Exec_CmdWaitDelay
var CmdWaitDelay = Os_Exec_CmdWaitDelay
var errorName = Os_Exec_ErrorName
var ErrorName = Os_Exec_ErrorName
var errorErr = Os_Exec_ErrorErr
var ErrorErr = Os_Exec_ErrorErr
var exitErrorProcessState = Os_Exec_ExitErrorProcessState
var ExitErrorProcessState = Os_Exec_ExitErrorProcessState
var exitErrorStderr = Os_Exec_ExitErrorStderr
var ExitErrorStderr = Os_Exec_ExitErrorStderr
var cmdCombinedOutput = Os_Exec_CmdCombinedOutput
var CmdCombinedOutput = Os_Exec_CmdCombinedOutput
var cmdEnviron = Os_Exec_CmdEnviron
var CmdEnviron = Os_Exec_CmdEnviron
var cmdOutput = Os_Exec_CmdOutput
var CmdOutput = Os_Exec_CmdOutput
var cmdRun = Os_Exec_CmdRun
var CmdRun = Os_Exec_CmdRun
var cmdStart = Os_Exec_CmdStart
var CmdStart = Os_Exec_CmdStart
var cmdStderrPipe = Os_Exec_CmdStderrPipe
var CmdStderrPipe = Os_Exec_CmdStderrPipe
var cmdStdinPipe = Os_Exec_CmdStdinPipe
var CmdStdinPipe = Os_Exec_CmdStdinPipe
var cmdStdoutPipe = Os_Exec_CmdStdoutPipe
var CmdStdoutPipe = Os_Exec_CmdStdoutPipe
var cmdString = Os_Exec_CmdString
var CmdString = Os_Exec_CmdString
var cmdWait = Os_Exec_CmdWait
var CmdWait = Os_Exec_CmdWait
var errorError = Os_Exec_ErrorError
var ErrorError = Os_Exec_ErrorError
var errorUnwrap = Os_Exec_ErrorUnwrap
var ErrorUnwrap = Os_Exec_ErrorUnwrap
var exitErrorError = Os_Exec_ExitErrorError
var ExitErrorError = Os_Exec_ExitErrorError
var exitErrorExitCode = Os_Exec_ExitErrorExitCode
var ExitErrorExitCode = Os_Exec_ExitErrorExitCode
var exitErrorExited = Os_Exec_ExitErrorExited
var ExitErrorExited = Os_Exec_ExitErrorExited
var exitErrorPid = Os_Exec_ExitErrorPid
var ExitErrorPid = Os_Exec_ExitErrorPid
var exitErrorString = Os_Exec_ExitErrorString
var ExitErrorString = Os_Exec_ExitErrorString
var exitErrorSuccess = Os_Exec_ExitErrorSuccess
var ExitErrorSuccess = Os_Exec_ExitErrorSuccess
var exitErrorSys = Os_Exec_ExitErrorSys
var ExitErrorSys = Os_Exec_ExitErrorSys
var exitErrorSysUsage = Os_Exec_ExitErrorSysUsage
var ExitErrorSysUsage = Os_Exec_ExitErrorSysUsage
var compile = Regexp_Compile
var Compile = Regexp_Compile
var compilePOSIX = Regexp_CompilePOSIX
var CompilePOSIX = Regexp_CompilePOSIX
var match = Regexp_Match
var Match = Regexp_Match
var matchString = Regexp_MatchString
var MatchString = Regexp_MatchString
var mustCompile = Regexp_MustCompile
var MustCompile = Regexp_MustCompile
var mustCompilePOSIX = Regexp_MustCompilePOSIX
var MustCompilePOSIX = Regexp_MustCompilePOSIX
var quoteMeta = Regexp_QuoteMeta
var QuoteMeta = Regexp_QuoteMeta
var Regexp = Regexp_Regexp
var regexpAppendText = Regexp_RegexpAppendText
var RegexpAppendText = Regexp_RegexpAppendText
var regexpCopy = Regexp_RegexpCopy
var RegexpCopy = Regexp_RegexpCopy
var regexpFind = Regexp_RegexpFind
var RegexpFind = Regexp_RegexpFind
var regexpFindAllString = Regexp_RegexpFindAllString
var RegexpFindAllString = Regexp_RegexpFindAllString
var regexpFindString = Regexp_RegexpFindString
var RegexpFindString = Regexp_RegexpFindString
var regexpFindStringSubmatch = Regexp_RegexpFindStringSubmatch
var RegexpFindStringSubmatch = Regexp_RegexpFindStringSubmatch
var regexpLiteralPrefix = Regexp_RegexpLiteralPrefix
var RegexpLiteralPrefix = Regexp_RegexpLiteralPrefix
var regexpLongest = Regexp_RegexpLongest
var RegexpLongest = Regexp_RegexpLongest
var regexpMarshalText = Regexp_RegexpMarshalText
var RegexpMarshalText = Regexp_RegexpMarshalText
var regexpMatch = Regexp_RegexpMatch
var RegexpMatch = Regexp_RegexpMatch
var regexpMatchString = Regexp_RegexpMatchString
var RegexpMatchString = Regexp_RegexpMatchString
var regexpNumSubexp = Regexp_RegexpNumSubexp
var RegexpNumSubexp = Regexp_RegexpNumSubexp
var regexpReplaceAll = Regexp_RegexpReplaceAll
var RegexpReplaceAll = Regexp_RegexpReplaceAll
var regexpReplaceAllLiteral = Regexp_RegexpReplaceAllLiteral
var RegexpReplaceAllLiteral = Regexp_RegexpReplaceAllLiteral
var regexpReplaceAllLiteralString = Regexp_RegexpReplaceAllLiteralString
var RegexpReplaceAllLiteralString = Regexp_RegexpReplaceAllLiteralString
var regexpReplaceAllString = Regexp_RegexpReplaceAllString
var RegexpReplaceAllString = Regexp_RegexpReplaceAllString
var regexpSplit = Regexp_RegexpSplit
var RegexpSplit = Regexp_RegexpSplit
var regexpString = Regexp_RegexpString
var RegexpString = Regexp_RegexpString
var regexpSubexpIndex = Regexp_RegexpSubexpIndex
var RegexpSubexpIndex = Regexp_RegexpSubexpIndex
var regexpSubexpNames = Regexp_RegexpSubexpNames
var RegexpSubexpNames = Regexp_RegexpSubexpNames
var regexpUnmarshalText = Regexp_RegexpUnmarshalText
var RegexpUnmarshalText = Regexp_RegexpUnmarshalText
var Debug = Log_Entry_Debug
var Info = Log_Entry_Info
var Warn = Log_Entry_Warn
var levelFromString = Log_Entry_LevelFromString
var LevelFromString = Log_Entry_LevelFromString
var levelToString = Log_Entry_LevelToString
var LevelToString = Log_Entry_LevelToString
var levelToInt = Log_Entry_LevelToInt
var LevelToInt = Log_Entry_LevelToInt
var levelDecoder = Log_Entry_LevelDecoder()
var LevelDecoder = Log_Entry_LevelDecoder()
var entryDecoder = Log_Entry_EntryDecoder()
var EntryDecoder = Log_Entry_EntryDecoder()
var decodeEntry = Log_Entry_DecodeEntry
var DecodeEntry = Log_Entry_DecodeEntry
var NoSection = Log_Config_NoSection
var InSource = Log_Config_InSource
var InWebhook = Log_Config_InWebhook
var emptySource = Log_Config_EmptySource()
var EmptySource = Log_Config_EmptySource()
var emptyWebhook = Log_Config_EmptyWebhook()
var EmptyWebhook = Log_Config_EmptyWebhook()
var parseConfig = Log_Config_ParseConfig
var ParseConfig = Log_Config_ParseConfig
var finalizeState = Log_Config_FinalizeState
var FinalizeState = Log_Config_FinalizeState
var parseLines = Log_Config_ParseLines
var ParseLines = Log_Config_ParseLines
var parseSourceLine = Log_Config_ParseSourceLine
var ParseSourceLine = Log_Config_ParseSourceLine
var parseWebhookLine = Log_Config_ParseWebhookLine
var ParseWebhookLine = Log_Config_ParseWebhookLine
var extractValue = Log_Config_ExtractValue
var ExtractValue = Log_Config_ExtractValue
var stripQuotes = Log_Config_StripQuotes
var StripQuotes = Log_Config_StripQuotes
var Log_Entry_Debug = map[string]any{"Tag": 0, "SkyName": "Debug"}
var Log_Entry_Info = map[string]any{"Tag": 1, "SkyName": "Info"}
var Log_Entry_Warn = map[string]any{"Tag": 2, "SkyName": "Warn"}
var Log_Entry_Error = map[string]any{"Tag": 3, "SkyName": "Error"}
var Log_Config_NoSection = map[string]any{"Tag": 0, "SkyName": "NoSection"}
var Os_File = map[string]any{"Tag": 0, "SkyName": "File"}
var Os_LinkError = map[string]any{"Tag": 0, "SkyName": "LinkError"}
var Os_ProcAttr = map[string]any{"Tag": 0, "SkyName": "ProcAttr"}
var Os_Process = map[string]any{"Tag": 0, "SkyName": "Process"}
var Os_ProcessState = map[string]any{"Tag": 0, "SkyName": "ProcessState"}
var Os_Root = map[string]any{"Tag": 0, "SkyName": "Root"}
var Os_Signal = map[string]any{"Tag": 0, "SkyName": "Signal"}
var Os_SyscallError = map[string]any{"Tag": 0, "SkyName": "SyscallError"}
var Os_Exec_Cmd = map[string]any{"Tag": 0, "SkyName": "Cmd"}
var Os_Exec_Error = map[string]any{"Tag": 0, "SkyName": "Error"}
var Os_Exec_ExitError = map[string]any{"Tag": 0, "SkyName": "ExitError"}
var Bufio_ReadWriter = map[string]any{"Tag": 0, "SkyName": "ReadWriter"}
var Bufio_Reader = map[string]any{"Tag": 0, "SkyName": "Reader"}
var Bufio_Scanner = map[string]any{"Tag": 0, "SkyName": "Scanner"}
var Bufio_SplitFunc = map[string]any{"Tag": 0, "SkyName": "SplitFunc"}
var Bufio_Writer = map[string]any{"Tag": 0, "SkyName": "Writer"}
var Regexp_Regexp = map[string]any{"Tag": 0, "SkyName": "Regexp"}
var LogPage = map[string]any{"Tag": 0, "SkyName": "LogPage"}
var FileMode = map[string]any{"Tag": 0, "SkyName": "FileMode"}
var StreamMode = map[string]any{"Tag": 1, "SkyName": "StreamMode"}
var Tick = map[string]any{"Tag": 0, "SkyName": "Tick"}
var ToggleAutoScroll = map[string]any{"Tag": 5, "SkyName": "ToggleAutoScroll"}
var ToggleTheme = map[string]any{"Tag": 6, "SkyName": "ToggleTheme"}
var ClearLogs = map[string]any{"Tag": 7, "SkyName": "ClearLogs"}
var newReadWriter = Bufio_NewReadWriter
var NewReadWriter = Bufio_NewReadWriter
var newReader = Bufio_NewReader
var NewReader = Bufio_NewReader
var newReaderSize = Bufio_NewReaderSize
var NewReaderSize = Bufio_NewReaderSize
var newScanner = Bufio_NewScanner
var NewScanner = Bufio_NewScanner
var newWriter = Bufio_NewWriter
var NewWriter = Bufio_NewWriter
var newWriterSize = Bufio_NewWriterSize
var NewWriterSize = Bufio_NewWriterSize
var ReadWriter = Bufio_ReadWriter
var Reader = Bufio_Reader
var Scanner = Bufio_Scanner
var SplitFunc = Bufio_SplitFunc
var Writer = Bufio_Writer
var readWriterReader = Bufio_ReadWriterReader
var ReadWriterReader = Bufio_ReadWriterReader
var readWriterWriter = Bufio_ReadWriterWriter
var ReadWriterWriter = Bufio_ReadWriterWriter
var readWriterAvailable = Bufio_ReadWriterAvailable
var ReadWriterAvailable = Bufio_ReadWriterAvailable
var readWriterAvailableBuffer = Bufio_ReadWriterAvailableBuffer
var ReadWriterAvailableBuffer = Bufio_ReadWriterAvailableBuffer
var readWriterDiscard = Bufio_ReadWriterDiscard
var ReadWriterDiscard = Bufio_ReadWriterDiscard
var readWriterFlush = Bufio_ReadWriterFlush
var ReadWriterFlush = Bufio_ReadWriterFlush
var readWriterPeek = Bufio_ReadWriterPeek
var ReadWriterPeek = Bufio_ReadWriterPeek
var readWriterRead = Bufio_ReadWriterRead
var ReadWriterRead = Bufio_ReadWriterRead
var readWriterReadFrom = Bufio_ReadWriterReadFrom
var ReadWriterReadFrom = Bufio_ReadWriterReadFrom
var readWriterUnreadByte = Bufio_ReadWriterUnreadByte
var ReadWriterUnreadByte = Bufio_ReadWriterUnreadByte
var readWriterUnreadRune = Bufio_ReadWriterUnreadRune
var ReadWriterUnreadRune = Bufio_ReadWriterUnreadRune
var readWriterWrite = Bufio_ReadWriterWrite
var ReadWriterWrite = Bufio_ReadWriterWrite
var readWriterWriteString = Bufio_ReadWriterWriteString
var ReadWriterWriteString = Bufio_ReadWriterWriteString
var readWriterWriteTo = Bufio_ReadWriterWriteTo
var ReadWriterWriteTo = Bufio_ReadWriterWriteTo
var readerBuffered = Bufio_ReaderBuffered
var ReaderBuffered = Bufio_ReaderBuffered
var readerDiscard = Bufio_ReaderDiscard
var ReaderDiscard = Bufio_ReaderDiscard
var readerPeek = Bufio_ReaderPeek
var ReaderPeek = Bufio_ReaderPeek
var readerRead = Bufio_ReaderRead
var ReaderRead = Bufio_ReaderRead
var readerReset = Bufio_ReaderReset
var ReaderReset = Bufio_ReaderReset
var readerSize = Bufio_ReaderSize
var ReaderSize = Bufio_ReaderSize
var readerUnreadByte = Bufio_ReaderUnreadByte
var ReaderUnreadByte = Bufio_ReaderUnreadByte
var readerUnreadRune = Bufio_ReaderUnreadRune
var ReaderUnreadRune = Bufio_ReaderUnreadRune
var readerWriteTo = Bufio_ReaderWriteTo
var ReaderWriteTo = Bufio_ReaderWriteTo
var scannerBuffer = Bufio_ScannerBuffer
var ScannerBuffer = Bufio_ScannerBuffer
var scannerBytes = Bufio_ScannerBytes
var ScannerBytes = Bufio_ScannerBytes
var scannerErr = Bufio_ScannerErr
var ScannerErr = Bufio_ScannerErr
var scannerScan = Bufio_ScannerScan
var ScannerScan = Bufio_ScannerScan
var scannerSplit = Bufio_ScannerSplit
var ScannerSplit = Bufio_ScannerSplit
var scannerText = Bufio_ScannerText
var ScannerText = Bufio_ScannerText
var writerAvailable = Bufio_WriterAvailable
var WriterAvailable = Bufio_WriterAvailable
var writerAvailableBuffer = Bufio_WriterAvailableBuffer
var WriterAvailableBuffer = Bufio_WriterAvailableBuffer
var writerBuffered = Bufio_WriterBuffered
var WriterBuffered = Bufio_WriterBuffered
var writerFlush = Bufio_WriterFlush
var WriterFlush = Bufio_WriterFlush
var writerReadFrom = Bufio_WriterReadFrom
var WriterReadFrom = Bufio_WriterReadFrom
var writerReset = Bufio_WriterReset
var WriterReset = Bufio_WriterReset
var writerSize = Bufio_WriterSize
var WriterSize = Bufio_WriterSize
var writerWrite = Bufio_WriterWrite
var WriterWrite = Bufio_WriterWrite
var writerWriteString = Bufio_WriterWriteString
var WriterWriteString = Bufio_WriterWriteString
var logMsg = Log_Webhook_LogMsg
var buildRules = Log_Webhook_BuildRules
var matchesFilter = Log_Webhook_MatchesFilter
var buildPayload = Log_Webhook_BuildPayload
var send = Log_Webhook_Send
var processNewEntries = Log_Webhook_ProcessNewEntries

func SkyOk(v any) SkyResult { return SkyResult{Tag: 0, SkyName: "Ok", OkValue: v} }

func SkyErr(v any) SkyResult { return SkyResult{Tag: 1, SkyName: "Err", ErrValue: v} }

func SkyJust(v any) SkyMaybe { return SkyMaybe{Tag: 0, SkyName: "Just", JustValue: v} }

func SkyNothing() SkyMaybe { return SkyMaybe{Tag: 1, SkyName: "Nothing"} }

func sky_asInt(v any) int {
	switch x := v.(type) {
	case int:
		return x
	case float64:
		return int(x)
	default:
		return 0
	}
}

func sky_asFloat(v any) float64 {
	switch x := v.(type) {
	case float64:
		return x
	case int:
		return float64(x)
	default:
		return 0
	}
}

func sky_asString(v any) string {
	switch x := v.(type) {
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func sky_asBool(v any) bool {
	if b, ok := v.(bool); ok {
		return b
	}
	return false
}

func sky_asList(v any) []any {
	if l, ok := v.([]any); ok {
		return l
	}
	return []any{}
}

func sky_asBytes(v any) []byte {
	if b, ok := v.([]byte); ok {
		return b
	}
	if s, ok := v.(string); ok {
		return []byte(s)
	}
	return nil
}

func sky_asStringSlice(v any) []string {
	items := sky_asList(v)
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = sky_asString(item)
	}
	return result
}

func sky_asContext(v any) context.Context {
	if c, ok := v.(context.Context); ok {
		return c
	}
	return context.Background()
}

func sky_numBinop(op string, a, b any) any {
	af, aIsF := a.(float64)
	bf, bIsF := b.(float64)
	if aIsF || bIsF {
		if !aIsF {
			af = sky_asFloat(a)
		}
		if !bIsF {
			bf = sky_asFloat(b)
		}
		switch op {
		case "+":
			return af + bf
		case "-":
			return af - bf
		case "*":
			return af * bf
		case "%":
			return int(af) % int(bf)
		}
		return af + bf
	}
	ai, bi := sky_asInt(a), sky_asInt(b)
	switch op {
	case "+":
		return ai + bi
	case "-":
		return ai - bi
	case "*":
		return ai * bi
	case "%":
		return ai % bi
	}
	return ai + bi
}

func sky_numCompare(op string, a, b any) bool {
	af, aIsF := a.(float64)
	bf, bIsF := b.(float64)
	if aIsF || bIsF {
		if !aIsF {
			af = sky_asFloat(a)
		}
		if !bIsF {
			bf = sky_asFloat(b)
		}
		switch op {
		case "<":
			return af < bf
		case "<=":
			return af <= bf
		case ">":
			return af > bf
		case ">=":
			return af >= bf
		}
		return false
	}
	ai, bi := sky_asInt(a), sky_asInt(b)
	switch op {
	case "<":
		return ai < bi
	case "<=":
		return ai <= bi
	case ">":
		return ai > bi
	case ">=":
		return ai >= bi
	}
	return false
}

func sky_asMap(v any) map[string]any {
	if m, ok := v.(map[string]any); ok {
		return m
	}
	return map[string]any{}
}

func sky_equal(a, b any) bool {
	switch av := a.(type) {
	case string:
		if bv, ok := b.(string); ok {
			return av == bv
		}
	case int:
		if bv, ok := b.(int); ok {
			return av == bv
		}
	case bool:
		if bv, ok := b.(bool); ok {
			return av == bv
		}
	case float64:
		if bv, ok := b.(float64); ok {
			return av == bv
		}
	}
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func sky_isAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}
	return true
}

func sky_concat(a, b any) any {
	if la, ok := a.([]any); ok {
		if lb, ok := b.([]any); ok {
			return append(la, lb...)
		}
	}
	return sky_asString(a) + sky_asString(b)
}

func sky_stringFromInt(v any) any { return strconv.Itoa(sky_asInt(v)) }

func sky_stringToUpper(v any) any { return strings.ToUpper(sky_asString(v)) }

func sky_stringToLower(v any) any { return strings.ToLower(sky_asString(v)) }

func sky_stringLength(v any) any {
	s := sky_asString(v)
	if sky_isAscii(s) {
		return len(s)
	}
	return len([]rune(s))
}

func sky_stringTrim(v any) any { return strings.TrimSpace(sky_asString(v)) }

func sky_stringContains(sub any) any {
	return func(s any) any { return strings.Contains(sky_asString(s), sky_asString(sub)) }
}

func sky_stringStartsWith(prefix any) any {
	return func(s any) any { return strings.HasPrefix(sky_asString(s), sky_asString(prefix)) }
}

func sky_stringEndsWith(suffix any) any {
	return func(s any) any { return strings.HasSuffix(sky_asString(s), sky_asString(suffix)) }
}

func sky_stringSplit(sep any) any {
	return func(s any) any {
		parts := strings.Split(sky_asString(s), sky_asString(sep))
		result := make([]any, len(parts))
		for i, p := range parts {
			result[i] = p
		}
		return result
	}
}

func sky_stringIsEmpty(v any) any { return sky_asString(v) == "" }

func sky_stringSlice(start any) any {
	return func(end any) any {
		return func(s any) any {
			str := sky_asString(s)
			a := sky_asInt(start)
			b := sky_asInt(end)
			if a < 0 {
				a = 0
			}
			if sky_isAscii(str) {
				if b > len(str) {
					b = len(str)
				}
				if a > b {
					return ""
				}
				return str[a:b]
			}
			runes := []rune(str)
			if b > len(runes) {
				b = len(runes)
			}
			if a > b {
				return ""
			}
			return string(runes[a:b])
		}
	}
}

func sky_stringJoin(sep any) any {
	return func(list any) any {
		parts := sky_asList(list)
		ss := make([]string, len(parts))
		for i, p := range parts {
			ss[i] = sky_asString(p)
		}
		return strings.Join(ss, sky_asString(sep))
	}
}

func sky_listMap(fn any) any {
	return func(list any) any {
		items := sky_asList(list)
		result := make([]any, len(items))
		for i, item := range items {
			result[i] = fn.(func(any) any)(item)
		}
		return result
	}
}

func sky_listFilter(fn any) any {
	return func(list any) any {
		items := sky_asList(list)
		var result []any
		for _, item := range items {
			if sky_asBool(fn.(func(any) any)(item)) {
				result = append(result, item)
			}
		}
		return result
	}
}

func sky_listFoldl(fn any) any {
	return func(init any) any {
		return func(list any) any {
			acc := init
			for _, item := range sky_asList(list) {
				acc = fn.(func(any) any)(item).(func(any) any)(acc)
			}
			return acc
		}
	}
}

func sky_listLength(list any) any { return len(sky_asList(list)) }

func sky_listReverse(list any) any {
	items := sky_asList(list)
	result := make([]any, len(items))
	for i, item := range items {
		result[len(items)-1-i] = item
	}
	return result
}

func sky_listIsEmpty(list any) any { return len(sky_asList(list)) == 0 }

func sky_listAppend(a any) any {
	return func(b any) any { return append(sky_asList(a), sky_asList(b)...) }
}

func sky_listFilterMap(fn any) any {
	return func(list any) any {
		var result []any
		for _, item := range sky_asList(list) {
			r := fn.(func(any) any)(item)
			if m, ok := r.(SkyMaybe); ok && m.Tag == 0 {
				result = append(result, m.JustValue)
			}
		}
		if result == nil {
			return []any{}
		}
		return result
	}
}

func sky_listDrop(n any) any {
	return func(list any) any {
		items := sky_asList(list)
		c := sky_asInt(n)
		if c >= len(items) {
			return []any{}
		}
		return items[c:]
	}
}

func sky_listMember(item any) any {
	return func(list any) any {
		for _, x := range sky_asList(list) {
			if sky_equal(x, item) {
				return true
			}
		}
		return false
	}
}

func sky_recordUpdate(base any, updates any) any {
	m := sky_asMap(base)
	result := make(map[string]any)
	for k, v := range m {
		result[k] = v
	}
	for k, v := range sky_asMap(updates) {
		result[k] = v
	}
	return result
}

func sky_println(args ...any) any {
	ss := make([]any, len(args))
	for i, a := range args {
		ss[i] = sky_asString(a)
	}
	fmt.Println(ss...)
	return struct{}{}
}

func sky_asSkyResult(v any) SkyResult {
	if r, ok := v.(SkyResult); ok {
		return r
	}
	return SkyResult{}
}

func sky_asSkyMaybe(v any) SkyMaybe {
	if m, ok := v.(SkyMaybe); ok {
		return m
	}
	return SkyMaybe{Tag: 1, SkyName: "Nothing"}
}

func sky_asTuple2(v any) SkyTuple2 {
	if t, ok := v.(SkyTuple2); ok {
		return t
	}
	return SkyTuple2{}
}

func sky_not(v any) any { return !sky_asBool(v) }

func sky_fileRead(path any) any {
	data, err := os.ReadFile(sky_asString(path))
	if err != nil {
		return SkyErr(err.Error())
	}
	return SkyOk(string(data))
}

func sky_processExit(code any) any { os.Exit(sky_asInt(code)); return struct{}{} }

func sky_processGetArgs(u any) any {
	args := make([]any, len(os.Args))
	for i, a := range os.Args {
		args[i] = a
	}
	return args
}

func sky_processGetArg(n any) any {
	idx := sky_asInt(n)
	if idx < len(os.Args) {
		return SkyJust(os.Args[idx])
	}
	return SkyNothing()
}

func sky_dictEmpty() any { return map[string]any{} }

func sky_dictInsert(k any) any {
	return func(v any) any {
		return func(d any) any {
			m := sky_asMap(d)
			result := make(map[string]any, len(m)+1)
			for key, val := range m {
				result[key] = val
			}
			result[sky_asString(k)] = v
			return result
		}
	}
}

func sky_dictGet(k any) any {
	return func(d any) any {
		m := sky_asMap(d)
		if v, ok := m[sky_asString(k)]; ok {
			return SkyJust(v)
		}
		return SkyNothing()
	}
}

func sky_identity(v any) any { return v }

func sky_js(v any) any {
	if s, ok := v.(string); ok && s == "nil" {
		return nil
	}
	return v
}

func sky_call(f any, arg any) any {
	if fn, ok := f.(func(any) any); ok {
		return fn(arg)
	}
	if s, ok := f.(string); ok {
		if args, ok := arg.([]any); ok {
			parts := make([]string, len(args))
			for i, a := range args {
				parts[i] = sky_asString(a)
			}
			return s + "(" + strings.Join(parts, ", ") + ")"
		}
		return s + " " + sky_asString(arg)
	}
	panic(fmt.Sprintf("sky_call: cannot call %T", f))
}

func sky_runTask(task any) any {
	if t, ok := task.(func() any); ok {
		var result any
		func() {
			defer func() {
				if r := recover(); r != nil {
					result = SkyErr(fmt.Sprintf("panic: %v", r))
				}
			}()
			result = t()
		}()
		return result
	}
	if r, ok := task.(SkyResult); ok {
		return r
	}
	return SkyOk(task)
}

func sky_runMainTask(result any) {
	if _, ok := result.(func() any); ok {
		r := sky_runTask(result)
		if sky_asSkyResult(r).Tag == 1 {
			fmt.Fprintln(os.Stderr, sky_asSkyResult(r).ErrValue)
			os.Exit(1)
		}
	}
}

func sky_resultWithDefault(def any) any {
	return func(r any) any {
		res := sky_asSkyResult(r)
		if res.Tag == 0 {
			return res.OkValue
		}
		return def
	}
}

func sky_listTake(n any) any {
	return func(list any) any {
		items := sky_asList(list)
		c := sky_asInt(n)
		if c >= len(items) {
			return list
		}
		return items[:c]
	}
}

func sky_listAny(fn any) any {
	return func(list any) any {
		for _, item := range sky_asList(list) {
			if sky_asBool(fn.(func(any) any)(item)) {
				return true
			}
		}
		return false
	}
}

func sky_stringLines(s any) any {
	parts := strings.Split(sky_asString(s), "\n")
	result := make([]any, len(parts))
	for i, p := range parts {
		result[i] = p
	}
	return result
}

func sky_jsonEncString(v any) any { return sky_asString(v) }

func sky_jsonEncInt(v any) any { return sky_asInt(v) }

func sky_jsonEncObject(pairs any) any {
	m := make(map[string]any)
	for _, p := range sky_asList(pairs) {
		t := sky_asTuple2(p)
		m[sky_asString(t.V0)] = t.V1
	}
	return m
}

func sky_jsonEncode(indent any) any {
	return func(value any) any {
		var b []byte
		var err error
		n := sky_asInt(indent)
		if n > 0 {
			b, err = encoding_json.MarshalIndent(value, "", strings.Repeat(" ", n))
		} else {
			b, err = encoding_json.Marshal(value)
		}
		if err != nil {
			return "null"
		}
		return string(b)
	}
}

func sky_jsonDecString(decoder any) any {
	return func(jsonStr any) any {
		var v any
		if err := encoding_json.Unmarshal([]byte(sky_asString(jsonStr)), &v); err != nil {
			return SkyErr(err.Error())
		}
		return decoder.(func(any) any)(v)
	}
}

func sky_jsonDecMap(fn any) any {
	return func(decoder any) any {
		return func(v any) any {
			r := decoder.(func(any) any)(v)
			res := sky_asSkyResult(r)
			if res.Tag == 1 {
				return r
			}
			return SkyOk(fn.(func(any) any)(res.OkValue))
		}
	}
}

func sky_jsonDecSucceed(v any) any { return func(_ any) any { return SkyOk(v) } }

func sky_jsonPipeRequired(key any) any {
	return func(decoder any) any {
		return func(pipeline any) any {
			return func(v any) any {
				pr := pipeline.(func(any) any)(v)
				pres := sky_asSkyResult(pr)
				if pres.Tag == 1 {
					return pr
				}
				m, ok := v.(map[string]any)
				if !ok {
					return SkyErr("expected object")
				}
				val, exists := m[sky_asString(key)]
				if !exists {
					return SkyErr("field '" + sky_asString(key) + "' required")
				}
				fr := decoder.(func(any) any)(val)
				fres := sky_asSkyResult(fr)
				if fres.Tag == 1 {
					return fr
				}
				return SkyOk(pres.OkValue.(func(any) any)(fres.OkValue))
			}
		}
	}
}

func sky_cmdNone() any { return []any{} }

func sky_timeEvery(interval any) any {
	return func(msg any) any { return map[string]any{"SkyName": "SubTimer", "V0": interval, "V1": msg} }
}

func sky_htmlEl(tag any) any {
	return func(attrs any) any {
		return func(children any) any {
			return map[string]any{"tag": tag, "attrs": sky_flattenAttrs(attrs), "children": children, "text": ""}
		}
	}
}

func sky_flattenAttrs(attrs any) []any {
	var out []any
	for _, a := range sky_asList(attrs) {
		if list, ok := a.([]any); ok {
			out = append(out, list...)
		} else {
			out = append(out, a)
		}
	}
	return out
}

func sky_htmlVoid(tag any) any {
	return func(attrs any) any {
		return map[string]any{"tag": tag, "attrs": attrs, "children": []any{}, "text": ""}
	}
}

func sky_htmlText(s any) any {
	return map[string]any{"tag": "", "attrs": []any{}, "children": []any{}, "text": sky_asString(s)}
}

func sky_attrSimple(key any) any {
	return func(v any) any { return SkyTuple2{sky_asString(key), sky_asString(v)} }
}

func sky_evtHandler(evtType any) any {
	return func(msg any) any { return sky_msgAttrs(sky_asString(evtType), msg) }
}

func sky_msgAttrs(evtType string, msg any) any {
	name := sky_msgName(msg)
	args := sky_msgArgs(msg)
	if len(args) == 0 {
		return SkyTuple2{"sky-" + evtType, name}
	}
	b, _ := encoding_json.Marshal(args)
	return []any{SkyTuple2{"sky-" + evtType, name}, SkyTuple2{"sky-args", string(b)}}
}

func sky_msgName(msg any) string {
	if m, ok := msg.(map[string]any); ok {
		if name, exists := m["SkyName"]; exists {
			return sky_asString(name)
		}
	}
	if fn, ok := msg.(func(any) any); ok {
		result := fn(nil)
		if m2, ok2 := result.(map[string]any); ok2 {
			if name, exists := m2["SkyName"]; exists {
				return sky_asString(name)
			}
		}
	}
	return fmt.Sprintf("%v", msg)
}

func sky_msgArgs(msg any) []any {
	m, ok := msg.(map[string]any)
	if !ok {
		return nil
	}
	var args []any
	for i := 0; ; i++ {
		v, exists := m[fmt.Sprintf("V%d", i)]
		if !exists {
			break
		}
		if vm, ok := v.(map[string]any); ok {
			if sn, ok := vm["SkyName"]; ok {
				args = append(args, sn)
				continue
			}
		}
		args = append(args, v)
	}
	return args
}

func sky_cssStyles(props any) any {
	var parts []string
	for _, p := range sky_asList(props) {
		parts = append(parts, sky_asString(p))
	}
	return strings.Join(parts, "; ")
}

func sky_liveRoute(path any) any {
	return func(page any) any { return map[string]any{"path": path, "page": page} }
}

func sky_liveApp(config any) any { return sky_liveAppImpl(config) }

func Log_Webhook_LogMsg(msg any) any {
	return func() any { Os_FileWriteString(Os_Stderr(struct{}{}), sky_concat(msg, "\n")); return struct{}{} }()
}

func Log_Webhook_BuildRules(config any) any {
	return sky_listFilterMap
}

func Log_Webhook_MatchesFilter(pattern any, entry any) any {
	return func() any {
		if sky_asBool(sky_stringIsEmpty(pattern)) {
			return false
		}
		return func() any {
			return func() any {
				__subject := Regexp_Compile(sky_concat("(?i)", pattern))
				if sky_asSkyResult(__subject).SkyName == "Ok" {
					re := sky_asSkyResult(__subject).OkValue
					_ = re
					return Regexp_RegexpMatchString(re, sky_asMap(entry)["message"])
				}
				if sky_asSkyResult(__subject).SkyName == "Err" {
					return sky_call(sky_stringContains(sky_stringToLower(pattern)), sky_stringToLower(sky_asMap(entry)["message"]))
				}
				panic("non-exhaustive case expression")
			}()
		}()
	}()
}

func Log_Webhook_BuildPayload(entry any) any {
	return sky_call(sky_jsonEncode(0), sky_jsonEncObject([]any{SkyTuple2{V0: "source", V1: sky_jsonEncString(sky_asMap(entry)["source"])}, SkyTuple2{V0: "level", V1: sky_jsonEncString(levelToString(sky_asMap(entry)["level"]))}, SkyTuple2{V0: "scope", V1: sky_jsonEncString(sky_asMap(entry)["scope"])}, SkyTuple2{V0: "message", V1: sky_jsonEncString(sky_asMap(entry)["message"])}, SkyTuple2{V0: "timestamp", V1: sky_jsonEncInt(sky_asMap(entry)["timestamp"])}}))
}

func Log_Webhook_Send(url any, entry any) any {
	return func() any {
		payload := Log_Webhook_BuildPayload(entry)
		_ = payload
		cmd := Os_Exec_Command("curl", []any{"-s", "-X", "POST", "-H", "Content-Type: application/json", "-d", payload, url})
		_ = cmd
		Os_Exec_CmdRun(cmd)
		return struct{}{}
	}()
}

func Log_Webhook_ProcessNewEntries(rules any, entries any) any {
	return func() any {
		if sky_asBool(sky_listIsEmpty(rules)) {
			return struct{}{}
		}
		return sky_call(sky_call(sky_listFoldl(func(entry any) any {
			return func(_ any) any {
				return sky_call(sky_call(sky_listFoldl(func(rule any) any {
					return func(_ any) any {
						return func() any {
							if sky_asBool(sky_asBool(sky_equal(sky_asMap(rule)["sourceName"], sky_asMap(entry)["source"])) && sky_asBool(Log_Webhook_MatchesFilter(sky_asMap(rule)["filter"], entry))) {
								return func() any {
									Log_Webhook_LogMsg(sky_concat("[webhook] filter matched: source=", sky_concat(sky_asMap(entry)["source"], sky_concat(" filter=", sky_concat(sky_asMap(rule)["filter"], sky_concat(" message=", sky_asMap(entry)["message"]))))))
									Log_Webhook_Send(sky_asMap(rule)["url"], entry)
									Log_Webhook_LogMsg(sky_concat("[webhook] fired: url=", sky_asMap(rule)["url"]))
									return struct{}{}
								}()
							}
							return struct{}{}
						}()
					}
				}), struct{}{}), rules)
			}
		}), struct{}{}), entries)
	}()
}

func Log_Entry_LevelFromString(str any) any {
	return func() any {
		return func() any {
			__subject := sky_stringToLower(str)
			if sky_asString(__subject) == "debug" {
				return map[string]any{"Tag": 2, "SkyName": "Debug"}
			}
			if sky_asString(__subject) == "info" {
				return map[string]any{"Tag": 0, "SkyName": "Info"}
			}
			if sky_asString(__subject) == "warn" {
				return map[string]any{"Tag": 1, "SkyName": "Warn"}
			}
			if sky_asString(__subject) == "error" {
				return map[string]any{"Tag": 3, "SkyName": "Error"}
			}
			if true {
				return map[string]any{"Tag": 0, "SkyName": "Info"}
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func Log_Entry_LevelToString(level any) any {
	return func() any {
		return func() any {
			__subject := level
			__sky_tag := sky_asMap(__subject)["SkyName"]
			if __sky_tag == "Debug" {
				return "DEBUG"
			}
			if __sky_tag == "Info" {
				return "INFO"
			}
			if __sky_tag == "Warn" {
				return "WARN"
			}
			if __sky_tag == "Error" {
				return "ERROR"
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func Log_Entry_LevelToInt(level any) any {
	return func() any {
		return func() any {
			__subject := level
			__sky_tag := sky_asMap(__subject)["SkyName"]
			if __sky_tag == "Debug" {
				return 0
			}
			if __sky_tag == "Info" {
				return 1
			}
			if __sky_tag == "Warn" {
				return 2
			}
			if __sky_tag == "Error" {
				return 3
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func Log_Entry_LevelDecoder() any {
	return sky_call(sky_jsonDecMap(Log_Entry_LevelFromString), sky_jsonDecoder_string)
}

func Log_Entry_EntryDecoder() any {
	return sky_call(sky_call(sky_jsonPipeRequired("message"), sky_jsonDecoder_string), sky_call(sky_call(sky_jsonPipeRequired("scope"), sky_jsonDecoder_string), sky_call(sky_call(sky_jsonPipeRequired("level"), Log_Entry_LevelDecoder()), sky_call(sky_call(sky_jsonPipeRequired("timestamp"), sky_jsonDecoder_int), sky_jsonDecSucceed(func(t any) any {
		return func(l any) any {
			return func(s any) any {
				return func(m any) any {
					return map[string]any{"timestamp": t, "level": l, "scope": s, "message": m, "source": ""}
				}
			}
		}
	})))))
}

func Log_Entry_DecodeEntry(line any, source any) any {
	return func() any {
		return func() any {
			__subject := sky_call(sky_jsonDecString(Log_Entry_EntryDecoder()), line)
			if sky_asSkyResult(__subject).SkyName == "Ok" {
				entry := sky_asSkyResult(__subject).OkValue
				_ = entry
				return sky_recordUpdate(entry, map[string]any{"source": source})
			}
			if sky_asSkyResult(__subject).SkyName == "Err" {
				return map[string]any{"timestamp": 0, "level": map[string]any{"Tag": 0, "SkyName": "Info"}, "scope": "raw", "message": line, "source": source}
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func Log_Config_InSource(v0 any) any {
	return map[string]any{"Tag": 1, "SkyName": "InSource", "V0": v0}
}

func Log_Config_InWebhook(v0 any) any {
	return map[string]any{"Tag": 2, "SkyName": "InWebhook", "V0": v0}
}

func Log_Config_EmptySource() any {
	return map[string]any{"name": "", "command": "", "filter": "", "webhookUrl": ""}
}

func Log_Config_EmptyWebhook() any {
	return map[string]any{"url": "", "filter": ""}
}

func Log_Config_ParseConfig(path any) any {
	return func() any {
		return func() any {
			__subject := sky_fileRead(path)
			if sky_asSkyResult(__subject).SkyName == "Ok" {
				content := sky_asSkyResult(__subject).OkValue
				_ = content
				return sky_call(func(__pa0 any) any {
					return Log_Config_ParseLines(map[string]any{"sources": []any{}, "webhook": Log_Config_EmptyWebhook(), "section": map[string]any{"Tag": 0, "SkyName": "NoSection"}}, __pa0)
				}, sky_call(sky_stringLines, content))
			}
			if sky_asSkyResult(__subject).SkyName == "Err" {
				return map[string]any{"sources": []any{}, "webhook": Log_Config_EmptyWebhook()}
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func Log_Config_FinalizeState(state any) any {
	return func() any {
		return func() any {
			__subject := sky_asMap(state)["section"]
			__sky_tag := sky_asMap(__subject)["SkyName"]
			if __sky_tag == "InSource" {
				s := sky_asMap(__subject)["V0"]
				_ = s
				return map[string]any{"sources": sky_call(sky_listAppend(sky_asMap(state)["sources"]), []any{s}), "webhook": sky_asMap(state)["webhook"]}
			}
			if __sky_tag == "InWebhook" {
				w := sky_asMap(__subject)["V0"]
				_ = w
				return map[string]any{"sources": sky_asMap(state)["sources"], "webhook": w}
			}
			if __sky_tag == "NoSection" {
				return map[string]any{"sources": sky_asMap(state)["sources"], "webhook": sky_asMap(state)["webhook"]}
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func Log_Config_ParseLines(state any, lines any) any {
	return func() any {
		return func() any {
			__subject := lines
			if len(sky_asList(__subject)) == 0 {
				return Log_Config_FinalizeState(state)
			}
			if len(sky_asList(__subject)) > 0 {
				line := sky_asList(__subject)[0]
				_ = line
				rest := sky_asList(__subject)[1:]
				_ = rest
				return func() any {
					trimmed := sky_stringTrim(line)
					_ = trimmed
					return func() any {
						if sky_asBool(sky_equal(trimmed, "[[source]]")) {
							return func() any {
								newState := func() any {
									return func() any {
										__subject := sky_asMap(state)["section"]
										__sky_tag := sky_asMap(__subject)["SkyName"]
										if __sky_tag == "InSource" {
											s := sky_asMap(__subject)["V0"]
											_ = s
											return sky_recordUpdate(state, map[string]any{"sources": sky_call(sky_listAppend(sky_asMap(state)["sources"]), []any{s}), "section": InSource(Log_Config_EmptySource())})
										}
										if __sky_tag == "InWebhook" {
											w := sky_asMap(__subject)["V0"]
											_ = w
											return sky_recordUpdate(state, map[string]any{"webhook": w, "section": InSource(Log_Config_EmptySource())})
										}
										if __sky_tag == "NoSection" {
											return sky_recordUpdate(state, map[string]any{"section": InSource(Log_Config_EmptySource())})
										}
										panic("non-exhaustive case expression")
									}()
								}()
								_ = newState
								return Log_Config_ParseLines(newState, rest)
							}()
						}
						if sky_asBool(sky_equal(trimmed, "[webhook]")) {
							return func() any {
								newState := func() any {
									return func() any {
										__subject := sky_asMap(state)["section"]
										__sky_tag := sky_asMap(__subject)["SkyName"]
										if __sky_tag == "InSource" {
											s := sky_asMap(__subject)["V0"]
											_ = s
											return sky_recordUpdate(state, map[string]any{"sources": sky_call(sky_listAppend(sky_asMap(state)["sources"]), []any{s}), "section": InWebhook(sky_asMap(state)["webhook"])})
										}
										if true {
											return sky_recordUpdate(state, map[string]any{"section": InWebhook(sky_asMap(state)["webhook"])})
										}
										panic("non-exhaustive case expression")
									}()
								}()
								_ = newState
								return Log_Config_ParseLines(newState, rest)
							}()
						}
						return func() any {
							newState := func() any {
								return func() any {
									__subject := sky_asMap(state)["section"]
									__sky_tag := sky_asMap(__subject)["SkyName"]
									if __sky_tag == "InSource" {
										s := sky_asMap(__subject)["V0"]
										_ = s
										return sky_recordUpdate(state, map[string]any{"section": InSource(Log_Config_ParseSourceLine(trimmed, s))})
									}
									if __sky_tag == "InWebhook" {
										w := sky_asMap(__subject)["V0"]
										_ = w
										return sky_recordUpdate(state, map[string]any{"section": InWebhook(Log_Config_ParseWebhookLine(trimmed, w))})
									}
									if __sky_tag == "NoSection" {
										return state
									}
									panic("non-exhaustive case expression")
								}()
							}()
							_ = newState
							return Log_Config_ParseLines(newState, rest)
						}()
					}()
				}()
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func Log_Config_ParseSourceLine(line any, source any) any {
	return func() any {
		if sky_asBool(sky_call(sky_stringStartsWith("name"), line)) {
			return sky_recordUpdate(source, map[string]any{"name": Log_Config_ExtractValue(line)})
		}
		if sky_asBool(sky_call(sky_stringStartsWith("command"), line)) {
			return sky_recordUpdate(source, map[string]any{"command": Log_Config_ExtractValue(line)})
		}
		if sky_asBool(sky_call(sky_stringStartsWith("filter"), line)) {
			return sky_recordUpdate(source, map[string]any{"filter": Log_Config_ExtractValue(line)})
		}
		if sky_asBool(sky_call(sky_stringStartsWith("webhook_url"), line)) {
			return sky_recordUpdate(source, map[string]any{"webhookUrl": Log_Config_ExtractValue(line)})
		}
		return source
	}()
}

func Log_Config_ParseWebhookLine(line any, webhook any) any {
	return func() any {
		if sky_asBool(sky_call(sky_stringStartsWith("url"), line)) {
			return sky_recordUpdate(webhook, map[string]any{"url": Log_Config_ExtractValue(line)})
		}
		if sky_asBool(sky_call(sky_stringStartsWith("filter"), line)) {
			return sky_recordUpdate(webhook, map[string]any{"filter": Log_Config_ExtractValue(line)})
		}
		return webhook
	}()
}

func Log_Config_ExtractValue(line any) any {
	return func() any {
		return func() any {
			__subject := sky_call(sky_stringSplit("="), line)
			if len(sky_asList(__subject)) == 0 {
				return ""
			}
			if len(sky_asList(__subject)) > 0 {
				rest := sky_asList(__subject)[1:]
				_ = rest
				return sky_call(Log_Config_StripQuotes, sky_call(sky_stringTrim, sky_call(sky_stringJoin("="), rest)))
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func Log_Config_StripQuotes(s any) any {
	return func() any {
		len_ := sky_stringLength(s)
		_ = len_
		return func() any {
			if sky_asBool(sky_asBool(sky_call(sky_stringStartsWith("\""), s)) && sky_asBool(sky_call(sky_stringEndsWith("\""), s))) {
				return sky_call(sky_call(sky_stringSlice(1), sky_numBinop("-", len_, 1)), s)
			}
			return s
		}()
	}()
}

func Os_FileWriteString(a0 any, a1 any) any { return Sky_os_FileWriteString(a0, a1) }

func Os_Stderr(a0 any) any { return Sky_os_Stderr(a0) }

func Os_Stdin(a0 any) any { return Sky_os_Stdin(a0) }

func Os_Exec_Command(a0 any, a1 any) any { return Sky_os_exec_Command(a0, a1) }

func Os_Exec_CommandContext(a0 any, a1 any, a2 any) any {
	return Sky_os_exec_CommandContext(a0, a1, a2)
}

func Os_Exec_LookPath(a0 any) any { return Sky_os_exec_LookPath(a0) }

func Os_Exec_CmdPath(a0 any) any { return Sky_os_exec_FIELD_Cmd_Path(a0) }

func Os_Exec_CmdArgs(a0 any) any { return Sky_os_exec_FIELD_Cmd_Args(a0) }

func Os_Exec_CmdEnv(a0 any) any { return Sky_os_exec_FIELD_Cmd_Env(a0) }

func Os_Exec_CmdDir(a0 any) any { return Sky_os_exec_FIELD_Cmd_Dir(a0) }

func Os_Exec_CmdStdin(a0 any) any { return Sky_os_exec_FIELD_Cmd_Stdin(a0) }

func Os_Exec_CmdStdout(a0 any) any { return Sky_os_exec_FIELD_Cmd_Stdout(a0) }

func Os_Exec_CmdStderr(a0 any) any { return Sky_os_exec_FIELD_Cmd_Stderr(a0) }

func Os_Exec_CmdExtraFiles(a0 any) any { return Sky_os_exec_FIELD_Cmd_ExtraFiles(a0) }

func Os_Exec_CmdSysProcAttr(a0 any) any { return Sky_os_exec_FIELD_Cmd_SysProcAttr(a0) }

func Os_Exec_CmdProcess(a0 any) any { return Sky_os_exec_FIELD_Cmd_Process(a0) }

func Os_Exec_CmdProcessState(a0 any) any { return Sky_os_exec_FIELD_Cmd_ProcessState(a0) }

func Os_Exec_CmdErr(a0 any) any { return Sky_os_exec_FIELD_Cmd_Err(a0) }

func Os_Exec_CmdWaitDelay(a0 any) any { return Sky_os_exec_FIELD_Cmd_WaitDelay(a0) }

func Os_Exec_ErrorName(a0 any) any { return Sky_os_exec_FIELD_Error_Name(a0) }

func Os_Exec_ErrorErr(a0 any) any { return Sky_os_exec_FIELD_Error_Err(a0) }

func Os_Exec_ExitErrorProcessState(a0 any) any { return Sky_os_exec_FIELD_ExitError_ProcessState(a0) }

func Os_Exec_ExitErrorStderr(a0 any) any { return Sky_os_exec_FIELD_ExitError_Stderr(a0) }

func Os_Exec_CmdCombinedOutput(a0 any) any { return Sky_os_exec_CmdCombinedOutput(a0) }

func Os_Exec_CmdEnviron(a0 any) any { return Sky_os_exec_CmdEnviron(a0) }

func Os_Exec_CmdOutput(a0 any) any { return Sky_os_exec_CmdOutput(a0) }

func Os_Exec_CmdRun(a0 any) any { return Sky_os_exec_CmdRun(a0) }

func Os_Exec_CmdStart(a0 any) any { return Sky_os_exec_CmdStart(a0) }

func Os_Exec_CmdStderrPipe(a0 any) any { return Sky_os_exec_CmdStderrPipe(a0) }

func Os_Exec_CmdStdinPipe(a0 any) any { return Sky_os_exec_CmdStdinPipe(a0) }

func Os_Exec_CmdStdoutPipe(a0 any) any { return Sky_os_exec_CmdStdoutPipe(a0) }

func Os_Exec_CmdString(a0 any) any { return Sky_os_exec_CmdString(a0) }

func Os_Exec_CmdWait(a0 any) any { return Sky_os_exec_CmdWait(a0) }

func Os_Exec_ErrorError(a0 any) any { return Sky_os_exec_ErrorError(a0) }

func Os_Exec_ErrorUnwrap(a0 any) any { return Sky_os_exec_ErrorUnwrap(a0) }

func Os_Exec_ExitErrorError(a0 any) any { return Sky_os_exec_ExitErrorError(a0) }

func Os_Exec_ExitErrorExitCode(a0 any) any { return Sky_os_exec_ExitErrorExitCode(a0) }

func Os_Exec_ExitErrorExited(a0 any) any { return Sky_os_exec_ExitErrorExited(a0) }

func Os_Exec_ExitErrorPid(a0 any) any { return Sky_os_exec_ExitErrorPid(a0) }

func Os_Exec_ExitErrorString(a0 any) any { return Sky_os_exec_ExitErrorString(a0) }

func Os_Exec_ExitErrorSuccess(a0 any) any { return Sky_os_exec_ExitErrorSuccess(a0) }

func Os_Exec_ExitErrorSys(a0 any) any { return Sky_os_exec_ExitErrorSys(a0) }

func Os_Exec_ExitErrorSysUsage(a0 any) any { return Sky_os_exec_ExitErrorSysUsage(a0) }

func Bufio_NewReadWriter(a0 any, a1 any) any { return Sky_bufio_NewReadWriter(a0, a1) }

func Bufio_NewReader(a0 any) any { return Sky_bufio_NewReader(a0) }

func Bufio_NewReaderSize(a0 any, a1 any) any { return Sky_bufio_NewReaderSize(a0, a1) }

func Bufio_NewScanner(a0 any) any { return Sky_bufio_NewScanner(a0) }

func Bufio_NewWriter(a0 any) any { return Sky_bufio_NewWriter(a0) }

func Bufio_NewWriterSize(a0 any, a1 any) any { return Sky_bufio_NewWriterSize(a0, a1) }

func Bufio_ReadWriterReader(a0 any) any { return Sky_bufio_FIELD_ReadWriter_Reader(a0) }

func Bufio_ReadWriterWriter(a0 any) any { return Sky_bufio_FIELD_ReadWriter_Writer(a0) }

func Bufio_ReadWriterAvailable(a0 any) any { return Sky_bufio_ReadWriterAvailable(a0) }

func Bufio_ReadWriterAvailableBuffer(a0 any) any { return Sky_bufio_ReadWriterAvailableBuffer(a0) }

func Bufio_ReadWriterDiscard(a0 any, a1 any) any { return Sky_bufio_ReadWriterDiscard(a0, a1) }

func Bufio_ReadWriterFlush(a0 any) any { return Sky_bufio_ReadWriterFlush(a0) }

func Bufio_ReadWriterPeek(a0 any, a1 any) any { return Sky_bufio_ReadWriterPeek(a0, a1) }

func Bufio_ReadWriterRead(a0 any, a1 any) any { return Sky_bufio_ReadWriterRead(a0, a1) }

func Bufio_ReadWriterReadFrom(a0 any, a1 any) any { return Sky_bufio_ReadWriterReadFrom(a0, a1) }

func Bufio_ReadWriterUnreadByte(a0 any) any { return Sky_bufio_ReadWriterUnreadByte(a0) }

func Bufio_ReadWriterUnreadRune(a0 any) any { return Sky_bufio_ReadWriterUnreadRune(a0) }

func Bufio_ReadWriterWrite(a0 any, a1 any) any { return Sky_bufio_ReadWriterWrite(a0, a1) }

func Bufio_ReadWriterWriteString(a0 any, a1 any) any { return Sky_bufio_ReadWriterWriteString(a0, a1) }

func Bufio_ReadWriterWriteTo(a0 any, a1 any) any { return Sky_bufio_ReadWriterWriteTo(a0, a1) }

func Bufio_ReaderBuffered(a0 any) any { return Sky_bufio_ReaderBuffered(a0) }

func Bufio_ReaderDiscard(a0 any, a1 any) any { return Sky_bufio_ReaderDiscard(a0, a1) }

func Bufio_ReaderPeek(a0 any, a1 any) any { return Sky_bufio_ReaderPeek(a0, a1) }

func Bufio_ReaderRead(a0 any, a1 any) any { return Sky_bufio_ReaderRead(a0, a1) }

func Bufio_ReaderReset(a0 any, a1 any) any { return Sky_bufio_ReaderReset(a0, a1) }

func Bufio_ReaderSize(a0 any) any { return Sky_bufio_ReaderSize(a0) }

func Bufio_ReaderUnreadByte(a0 any) any { return Sky_bufio_ReaderUnreadByte(a0) }

func Bufio_ReaderUnreadRune(a0 any) any { return Sky_bufio_ReaderUnreadRune(a0) }

func Bufio_ReaderWriteTo(a0 any, a1 any) any { return Sky_bufio_ReaderWriteTo(a0, a1) }

func Bufio_ScannerBuffer(a0 any, a1 any, a2 any) any { return Sky_bufio_ScannerBuffer(a0, a1, a2) }

func Bufio_ScannerBytes(a0 any) any { return Sky_bufio_ScannerBytes(a0) }

func Bufio_ScannerErr(a0 any) any { return Sky_bufio_ScannerErr(a0) }

func Bufio_ScannerScan(a0 any) any { return Sky_bufio_ScannerScan(a0) }

func Bufio_ScannerSplit(a0 any, a1 any) any { return Sky_bufio_ScannerSplit(a0, a1) }

func Bufio_ScannerText(a0 any) any { return Sky_bufio_ScannerText(a0) }

func Bufio_WriterAvailable(a0 any) any { return Sky_bufio_WriterAvailable(a0) }

func Bufio_WriterAvailableBuffer(a0 any) any { return Sky_bufio_WriterAvailableBuffer(a0) }

func Bufio_WriterBuffered(a0 any) any { return Sky_bufio_WriterBuffered(a0) }

func Bufio_WriterFlush(a0 any) any { return Sky_bufio_WriterFlush(a0) }

func Bufio_WriterReadFrom(a0 any, a1 any) any { return Sky_bufio_WriterReadFrom(a0, a1) }

func Bufio_WriterReset(a0 any, a1 any) any { return Sky_bufio_WriterReset(a0, a1) }

func Bufio_WriterSize(a0 any) any { return Sky_bufio_WriterSize(a0) }

func Bufio_WriterWrite(a0 any, a1 any) any { return Sky_bufio_WriterWrite(a0, a1) }

func Bufio_WriterWriteString(a0 any, a1 any) any { return Sky_bufio_WriterWriteString(a0, a1) }

func Regexp_Compile(a0 any) any { return Sky_regexp_Compile(a0) }

func Regexp_CompilePOSIX(a0 any) any { return Sky_regexp_CompilePOSIX(a0) }

func Regexp_Match(a0 any, a1 any) any { return Sky_regexp_Match(a0, a1) }

func Regexp_MatchString(a0 any, a1 any) any { return Sky_regexp_MatchString(a0, a1) }

func Regexp_MustCompile(a0 any) any { return Sky_regexp_MustCompile(a0) }

func Regexp_MustCompilePOSIX(a0 any) any { return Sky_regexp_MustCompilePOSIX(a0) }

func Regexp_QuoteMeta(a0 any) any { return Sky_regexp_QuoteMeta(a0) }

func Regexp_RegexpAppendText(a0 any, a1 any) any { return Sky_regexp_RegexpAppendText(a0, a1) }

func Regexp_RegexpCopy(a0 any) any { return Sky_regexp_RegexpCopy(a0) }

func Regexp_RegexpFind(a0 any, a1 any) any { return Sky_regexp_RegexpFind(a0, a1) }

func Regexp_RegexpFindAllString(a0 any, a1 any, a2 any) any {
	return Sky_regexp_RegexpFindAllString(a0, a1, a2)
}

func Regexp_RegexpFindString(a0 any, a1 any) any { return Sky_regexp_RegexpFindString(a0, a1) }

func Regexp_RegexpFindStringSubmatch(a0 any, a1 any) any {
	return Sky_regexp_RegexpFindStringSubmatch(a0, a1)
}

func Regexp_RegexpLiteralPrefix(a0 any) any { return Sky_regexp_RegexpLiteralPrefix(a0) }

func Regexp_RegexpLongest(a0 any) any { return Sky_regexp_RegexpLongest(a0) }

func Regexp_RegexpMarshalText(a0 any) any { return Sky_regexp_RegexpMarshalText(a0) }

func Regexp_RegexpMatch(a0 any, a1 any) any { return Sky_regexp_RegexpMatch(a0, a1) }

func Regexp_RegexpMatchString(a0 any, a1 any) any { return Sky_regexp_RegexpMatchString(a0, a1) }

func Regexp_RegexpNumSubexp(a0 any) any { return Sky_regexp_RegexpNumSubexp(a0) }

func Regexp_RegexpReplaceAll(a0 any, a1 any, a2 any) any {
	return Sky_regexp_RegexpReplaceAll(a0, a1, a2)
}

func Regexp_RegexpReplaceAllLiteral(a0 any, a1 any, a2 any) any {
	return Sky_regexp_RegexpReplaceAllLiteral(a0, a1, a2)
}

func Regexp_RegexpReplaceAllLiteralString(a0 any, a1 any, a2 any) any {
	return Sky_regexp_RegexpReplaceAllLiteralString(a0, a1, a2)
}

func Regexp_RegexpReplaceAllString(a0 any, a1 any, a2 any) any {
	return Sky_regexp_RegexpReplaceAllString(a0, a1, a2)
}

func Regexp_RegexpSplit(a0 any, a1 any, a2 any) any { return Sky_regexp_RegexpSplit(a0, a1, a2) }

func Regexp_RegexpString(a0 any) any { return Sky_regexp_RegexpString(a0) }

func Regexp_RegexpSubexpIndex(a0 any, a1 any) any { return Sky_regexp_RegexpSubexpIndex(a0, a1) }

func Regexp_RegexpSubexpNames(a0 any) any { return Sky_regexp_RegexpSubexpNames(a0) }

func Regexp_RegexpUnmarshalText(a0 any, a1 any) any { return Sky_regexp_RegexpUnmarshalText(a0, a1) }

func SetLevelFilter(v0 any) any {
	return map[string]any{"Tag": 1, "SkyName": "SetLevelFilter", "V0": v0}
}

func SetScopeFilter(v0 any) any {
	return map[string]any{"Tag": 2, "SkyName": "SetScopeFilter", "V0": v0}
}

func SetSearchFilter(v0 any) any {
	return map[string]any{"Tag": 3, "SkyName": "SetSearchFilter", "V0": v0}
}

func SetSourceFilter(v0 any) any {
	return map[string]any{"Tag": 4, "SkyName": "SetSourceFilter", "V0": v0}
}

func Navigate(v0 any) any {
	return map[string]any{"Tag": 8, "SkyName": "Navigate", "V0": v0}
}

func helpText() any {
	return "sky-log — real-time log viewer\n\nUsage:\n  sky-log [options] [file ...]\n  command | sky-log\n\nOptions:\n  -h, --help    Show this help message\n\nExamples:\n  sky-log app.log                  Watch a log file\n  sky-log app.log server.log       Watch multiple files\n  tail -f /var/log/syslog | sky-log  Pipe stdin\n  sky-log                          Use sources from sky-log.toml\n"
}

func readFileEntries(source any) any {
	return func() any {
		return func() any {
			__subject := sky_fileRead(sky_asMap(source)["path"])
			if sky_asSkyResult(__subject).SkyName == "Ok" {
				content := sky_asSkyResult(__subject).OkValue
				_ = content
				return sky_call(sky_listMap(func(l any) any { return Log_Entry_DecodeEntry(l, sky_asMap(source)["label"]) }), sky_call(sky_listFilter(func(l any) any { return sky_not(sky_stringIsEmpty(l)) }), sky_call(sky_stringLines, content)))
			}
			if sky_asSkyResult(__subject).SkyName == "Err" {
				return []any{}
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func collectFileEntries(watched any, counts any, existing any) any {
	return sky_call(sky_call(sky_listFoldl(func(source any) any {
		return func(acc any) any {
			return func() any {
				seen := func() any {
					return func() any {
						__subject := sky_call(sky_dictGet(sky_asMap(source)["path"]), sky_asMap(acc)["counts"])
						if sky_asSkyMaybe(__subject).SkyName == "Just" {
							n := sky_asSkyMaybe(__subject).JustValue
							_ = n
							return n
						}
						if sky_asSkyMaybe(__subject).SkyName == "Nothing" {
							return 0
						}
						panic("non-exhaustive case expression")
					}()
				}()
				_ = seen
				allEntries := readFileEntries(source)
				_ = allEntries
				total := sky_listLength(allEntries)
				_ = total
				newOnes := sky_call(sky_listDrop(seen), allEntries)
				_ = newOnes
				return map[string]any{"entries": sky_call(sky_listAppend(sky_asMap(acc)["entries"]), newOnes), "counts": sky_call(sky_call(sky_dictInsert(sky_asMap(source)["path"]), total), sky_asMap(acc)["counts"]), "newEntries": sky_call(sky_listAppend(sky_asMap(acc)["newEntries"]), newOnes)}
			}()
		}
	}), map[string]any{"entries": existing, "counts": counts, "newEntries": []any{}}), watched)
}

func scanLine(source any) any {
	return func() any {
		if sky_asBool(sky_call(sky_resultWithDefault(false), Bufio_ScannerScan(sky_asMap(source)["scanner"]))) {
			return func() any {
				text := sky_call(sky_resultWithDefault(""), Bufio_ScannerText(sky_asMap(source)["scanner"]))
				_ = text
				return SkyJust(Log_Entry_DecodeEntry(text, sky_asMap(source)["label"]))
			}()
		}
		return SkyNothing()
	}()
}

func collectStreamEntries(scanners any) any {
	return sky_call(sky_listFilterMap(scanLine), scanners)
}

func initStdinScanner() any {
	return func() any {
		return func() any {
			__subject := Bufio_NewScanner(sky_identity(Os_Stdin(struct{}{})))
			if sky_asSkyResult(__subject).SkyName == "Ok" {
				scanner := sky_asSkyResult(__subject).OkValue
				_ = scanner
				return SkyJust(map[string]any{"scanner": scanner, "label": "stdin"})
			}
			if sky_asSkyResult(__subject).SkyName == "Err" {
				return SkyNothing()
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func initCommandScanner(source any) any {
	return func() any {
		return func() any {
			__subject := Os_Exec_Command("sh", []any{"-c", sky_asMap(source)["command"]})
			if sky_asSkyResult(__subject).SkyName == "Ok" {
				cmd := sky_asSkyResult(__subject).OkValue
				_ = cmd
				return func() any {
					return func() any {
						__subject := Os_Exec_CmdStdoutPipe(cmd)
						if sky_asSkyResult(__subject).SkyName == "Ok" {
							pipe := sky_asSkyResult(__subject).OkValue
							_ = pipe
							return func() any {
								return func() any {
									__subject := Os_Exec_CmdStart(cmd)
									if sky_asSkyResult(__subject).SkyName == "Ok" {
										return func() any {
											return func() any {
												__subject := Bufio_NewScanner(sky_identity(pipe))
												if sky_asSkyResult(__subject).SkyName == "Ok" {
													scanner := sky_asSkyResult(__subject).OkValue
													_ = scanner
													return SkyJust(map[string]any{"scanner": scanner, "label": sky_asMap(source)["name"]})
												}
												if sky_asSkyResult(__subject).SkyName == "Err" {
													return SkyNothing()
												}
												panic("non-exhaustive case expression")
											}()
										}()
									}
									if sky_asSkyResult(__subject).SkyName == "Err" {
										return SkyNothing()
									}
									panic("non-exhaustive case expression")
								}()
							}()
						}
						if sky_asSkyResult(__subject).SkyName == "Err" {
							return SkyNothing()
						}
						panic("non-exhaustive case expression")
					}()
				}()
			}
			if sky_asSkyResult(__subject).SkyName == "Err" {
				return SkyNothing()
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func resolveMode(args any, sources any) any {
	return func() any {
		if sky_asBool(sky_not(sky_listIsEmpty(args))) {
			return map[string]any{"watched": sky_call(sky_listMap(func(f any) any { return map[string]any{"path": f, "label": f} }), args), "scanners": []any{}, "mode": map[string]any{"Tag": 0, "SkyName": "FileMode"}}
		}
		return func() any {
			cmdScanners := sky_call(sky_listFilterMap(initCommandScanner), sources)
			_ = cmdScanners
			return func() any {
				if sky_asBool(sky_not(sky_listIsEmpty(cmdScanners))) {
					return map[string]any{"watched": []any{}, "scanners": cmdScanners, "mode": map[string]any{"Tag": 1, "SkyName": "StreamMode"}}
				}
				return map[string]any{"watched": []any{}, "scanners": sky_call(sky_listFilterMap(sky_identity), []any{initStdinScanner()}), "mode": map[string]any{"Tag": 1, "SkyName": "StreamMode"}}
			}()
		}()
	}()
}

func init_(_ any) any {
	return func() any {
		rawArgs := sky_call(sky_listDrop(1), sky_processGetArgs(struct{}{}))
		_ = rawArgs
		func() any {
			if sky_asBool(sky_call(sky_listAny(func(a any) any { return sky_asBool(sky_equal(a, "--help")) || sky_asBool(sky_equal(a, "-h")) }), rawArgs)) {
				return func() any { sky_println("", helpText()); return sky_processExit(0) }()
			}
			return struct{}{}
		}()
		args := sky_call(sky_listFilter(func(a any) any { return sky_not(sky_call(sky_stringStartsWith("-"), a)) }), rawArgs)
		_ = args
		config := Log_Config_ParseConfig("sky-log.toml")
		_ = config
		resolved := resolveMode(args, sky_asMap(config)["sources"])
		_ = resolved
		webhookRules := Log_Webhook_BuildRules(config)
		_ = webhookRules
		result := collectFileEntries(sky_asMap(resolved)["watched"], sky_dictEmpty(), []any{})
		_ = result
		return SkyTuple2{V0: map[string]any{"entries": sky_asMap(result)["entries"], "levelFilter": "", "scopeFilter": "", "searchFilter": "", "sourceFilter": "", "autoScroll": true, "theme": "dark", "watched": sky_asMap(resolved)["watched"], "scanners": sky_asMap(resolved)["scanners"], "fileCounts": sky_asMap(result)["counts"], "sourceMode": sky_asMap(resolved)["mode"], "webhookRules": webhookRules}, V1: sky_cmdNone()}
	}()
}

func update(msg any, model any) any {
	return func() any {
		return func() any {
			__subject := msg
			__sky_tag := sky_asMap(__subject)["SkyName"]
			if __sky_tag == "Tick" {
				return func() any {
					return func() any {
						__subject := sky_asMap(model)["sourceMode"]
						__sky_tag := sky_asMap(__subject)["SkyName"]
						if __sky_tag == "FileMode" {
							return func() any {
								result := collectFileEntries(sky_asMap(model)["watched"], sky_asMap(model)["fileCounts"], sky_asMap(model)["entries"])
								_ = result
								Log_Webhook_ProcessNewEntries(sky_asMap(model)["webhookRules"], sky_asMap(result)["newEntries"])
								return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"entries": sky_asMap(result)["entries"], "fileCounts": sky_asMap(result)["counts"]}), V1: sky_cmdNone()}
							}()
						}
						if __sky_tag == "StreamMode" {
							return func() any {
								newEntries := collectStreamEntries(sky_asMap(model)["scanners"])
								_ = newEntries
								Log_Webhook_ProcessNewEntries(sky_asMap(model)["webhookRules"], newEntries)
								return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"entries": sky_call(sky_listAppend(sky_asMap(model)["entries"]), newEntries)}), V1: sky_cmdNone()}
							}()
						}
						panic("non-exhaustive case expression")
					}()
				}()
			}
			if __sky_tag == "SetLevelFilter" {
				value := sky_asMap(__subject)["V0"]
				_ = value
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"levelFilter": value}), V1: sky_cmdNone()}
			}
			if __sky_tag == "SetScopeFilter" {
				value := sky_asMap(__subject)["V0"]
				_ = value
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"scopeFilter": value}), V1: sky_cmdNone()}
			}
			if __sky_tag == "SetSearchFilter" {
				value := sky_asMap(__subject)["V0"]
				_ = value
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"searchFilter": value}), V1: sky_cmdNone()}
			}
			if __sky_tag == "SetSourceFilter" {
				value := sky_asMap(__subject)["V0"]
				_ = value
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"sourceFilter": value}), V1: sky_cmdNone()}
			}
			if __sky_tag == "ToggleAutoScroll" {
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"autoScroll": sky_not(sky_asMap(model)["autoScroll"])}), V1: sky_cmdNone()}
			}
			if __sky_tag == "ToggleTheme" {
				return func() any {
					newTheme := func() any {
						if sky_asBool(sky_equal(sky_asMap(model)["theme"], "dark")) {
							return "light"
						}
						return "dark"
					}()
					_ = newTheme
					return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"theme": newTheme}), V1: sky_cmdNone()}
				}()
			}
			if __sky_tag == "ClearLogs" {
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"entries": []any{}}), V1: sky_cmdNone()}
			}
			if __sky_tag == "Navigate" {
				return SkyTuple2{V0: model, V1: sky_cmdNone()}
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func subscriptions(_ any) any {
	return sky_call(sky_timeEvery(200), map[string]any{"Tag": 0, "SkyName": "Tick"})
}

func filterEntry(levelStr any, scopeStr any, searchStr any, sourceStr any, entry any) any {
	return func() any {
		levelOk := func() any {
			if sky_asBool(sky_stringIsEmpty(levelStr)) {
				return true
			}
			return sky_numCompare(">=", Log_Entry_LevelToInt(sky_asMap(entry)["level"]), Log_Entry_LevelToInt(Log_Entry_LevelFromString(levelStr)))
		}()
		_ = levelOk
		scopeOk := func() any {
			if sky_asBool(sky_stringIsEmpty(scopeStr)) {
				return true
			}
			return sky_call(sky_stringContains(sky_stringToLower(scopeStr)), sky_stringToLower(sky_asMap(entry)["scope"]))
		}()
		_ = scopeOk
		searchOk := func() any {
			if sky_asBool(sky_stringIsEmpty(searchStr)) {
				return true
			}
			return func() any {
				return func() any {
					__subject := Regexp_Compile(sky_concat("(?i)", searchStr))
					if sky_asSkyResult(__subject).SkyName == "Ok" {
						re := sky_asSkyResult(__subject).OkValue
						_ = re
						return Regexp_RegexpMatchString(re, sky_asMap(entry)["message"])
					}
					if sky_asSkyResult(__subject).SkyName == "Err" {
						return sky_call(sky_stringContains(sky_stringToLower(searchStr)), sky_stringToLower(sky_asMap(entry)["message"]))
					}
					panic("non-exhaustive case expression")
				}()
			}()
		}()
		_ = searchOk
		sourceOk := func() any {
			if sky_asBool(sky_stringIsEmpty(sourceStr)) {
				return true
			}
			return sky_equal(sky_asMap(entry)["source"], sourceStr)
		}()
		_ = sourceOk
		return sky_asBool(levelOk) && sky_asBool(sky_asBool(scopeOk) && sky_asBool(sky_asBool(searchOk) && sky_asBool(sourceOk)))
	}()
}

func levelClass(level any) any {
	return func() any {
		return func() any {
			__subject := level
			__sky_tag := sky_asMap(__subject)["SkyName"]
			if __sky_tag == "Debug" {
				return "log-entry l-debug"
			}
			if __sky_tag == "Info" {
				return "log-entry l-info"
			}
			if __sky_tag == "Warn" {
				return "log-entry l-warn"
			}
			if __sky_tag == "Error" {
				return "log-entry l-error"
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func padRight(targetLen any, pad any, str any) any {
	return func() any {
		if sky_asBool(sky_numCompare(">=", sky_stringLength(str), targetLen)) {
			return str
		}
		return padRight(targetLen, pad, sky_concat(str, pad))
	}()
}

func formatTimestamp(ts any) any {
	return func() any {
		secs := func() any {
			if sky_asBool(sky_numCompare(">", ts, 9999999999)) {
				return sky_asInt(ts) / sky_asInt(1000)
			}
			return ts
		}()
		_ = secs
		totalSeconds := sky_numBinop("-", secs, sky_numBinop("*", sky_asInt(secs)/sky_asInt(86400), 86400))
		_ = totalSeconds
		h := sky_asInt(totalSeconds) / sky_asInt(3600)
		_ = h
		remainder := sky_numBinop("-", totalSeconds, sky_numBinop("*", h, 3600))
		_ = remainder
		m := sky_asInt(remainder) / sky_asInt(60)
		_ = m
		s := sky_numBinop("-", remainder, sky_numBinop("*", m, 60))
		_ = s
		return sky_concat(padZero(h), sky_concat(":", sky_concat(padZero(m), sky_concat(":", padZero(s)))))
	}()
}

func padZero(n any) any {
	return func() any {
		if sky_asBool(sky_numCompare("<", n, 10)) {
			return sky_concat("0", sky_stringFromInt(n))
		}
		return sky_stringFromInt(n)
	}()
}

func viewEntry(entry any) any {
	return sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), levelClass(sky_asMap(entry)["level"]))}), []any{sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "ts")}), []any{sky_htmlText(formatTimestamp(sky_asMap(entry)["timestamp"]))}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "level")}), []any{sky_htmlText(sky_stringToUpper(padRight(5, " ", Log_Entry_LevelToString(sky_asMap(entry)["level"]))))}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "source")}), []any{sky_htmlText(sky_asMap(entry)["source"])}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "scope")}), []any{sky_htmlText(sky_asMap(entry)["scope"])}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "msg")}), []any{sky_htmlText(sky_asMap(entry)["message"])})})
}

func viewEntries(entries any) any {
	return func() any {
		if sky_asBool(sky_listIsEmpty(entries)) {
			return sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "empty")}), []any{sky_htmlText("Waiting for logs...")})
		}
		return sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("id"), "log-entries")}), sky_call(sky_listMap(viewEntry), sky_call(sky_listTake(5000), sky_listReverse(sky_listReverse(entries)))))
	}()
}

func statusText(model any) any {
	return func() any {
		return func() any {
			__subject := sky_asMap(model)["sourceMode"]
			__sky_tag := sky_asMap(__subject)["SkyName"]
			if __sky_tag == "FileMode" {
				return func() any {
					labels := sky_call(sky_listMap(func(w any) any { return sky_asMap(w)["label"] }), sky_asMap(model)["watched"])
					_ = labels
					return sky_concat("Watching ", sky_concat(sky_stringFromInt(sky_listLength(sky_asMap(model)["watched"])), sky_concat(" source(s): ", sky_call(sky_stringJoin(", "), labels))))
				}()
			}
			if __sky_tag == "StreamMode" {
				return func() any {
					labels := sky_call(sky_listMap(func(s any) any { return sky_asMap(s)["label"] }), sky_asMap(model)["scanners"])
					_ = labels
					return sky_concat("Streaming: ", sky_call(sky_stringJoin(", "), labels))
				}()
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

func uniqueSources(entries any) any {
	return sky_call(sky_call(sky_listFoldl(func(s any) any {
		return func(acc any) any {
			return func() any {
				if sky_asBool(sky_call(sky_listMember(s), acc)) {
					return acc
				}
				return sky_call(sky_listAppend(acc), []any{s})
			}()
		}
	}), []any{}), sky_call(sky_listFilter(func(s any) any { return sky_not(sky_stringIsEmpty(s)) }), sky_call(sky_listMap(func(e any) any { return sky_asMap(e)["source"] }), entries)))
}

func view(model any) any {
	return func() any {
		filtered := sky_call(sky_listFilter(func(__pa0 any) any {
			return filterEntry(sky_asMap(model)["levelFilter"], sky_asMap(model)["scopeFilter"], sky_asMap(model)["searchFilter"], sky_asMap(model)["sourceFilter"], __pa0)
		}), sky_asMap(model)["entries"])
		_ = filtered
		sources := uniqueSources(sky_asMap(model)["entries"])
		_ = sources
		totalCount := sky_listLength(sky_asMap(model)["entries"])
		_ = totalCount
		filteredCount := sky_listLength(filtered)
		_ = filteredCount
		scrollLabel := func() any {
			if sky_asBool(sky_asMap(model)["autoScroll"]) {
				return "Auto-scroll"
			}
			return "Paused"
		}()
		_ = scrollLabel
		scrollClass := func() any {
			if sky_asBool(sky_asMap(model)["autoScroll"]) {
				return "btn active"
			}
			return "btn"
		}()
		_ = scrollClass
		themeIcon := func() any {
			if sky_asBool(sky_equal(sky_asMap(model)["theme"], "dark")) {
				return "☼"
			}
			return "☾"
		}()
		_ = themeIcon
		return sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), sky_concat("root ", sky_asMap(model)["theme"]))}), []any{sky_cssStyles, sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "toolbar")}), []any{sky_call(sky_call(sky_htmlEl("h1"), []any{}), []any{sky_htmlText("sky"), sky_call(sky_call(sky_htmlEl("span"), []any{}), []any{sky_htmlText("-log")})}), sky_call(sky_call(sky_htmlEl("button"), []any{sky_call(sky_attrSimple("class"), "btn"), sky_call(sky_evtHandler("click"), map[string]any{"Tag": 3, "SkyName": "ToggleTheme"})}), []any{sky_htmlText(themeIcon)}), sky_call(sky_call(sky_htmlEl("select"), []any{sky_call(sky_attrSimple("id"), "level-filter"), sky_call(sky_evtHandler("change"), SetLevelFilter)}), []any{sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "")}), []any{sky_htmlText("Level")}), sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "debug")}), []any{sky_htmlText("DEBUG")}), sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "info")}), []any{sky_htmlText("INFO")}), sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "warn")}), []any{sky_htmlText("WARN")}), sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "error")}), []any{sky_htmlText("ERROR")})}), sky_call(sky_htmlVoid("input"), []any{sky_call(sky_attrSimple("id"), "scope-filter"), sky_call(sky_attrSimple("type"), "text"), sky_call(sky_attrSimple("placeholder"), "scope"), sky_call(sky_evtHandler("input"), SetScopeFilter)}), sky_call(sky_htmlVoid("input"), []any{sky_call(sky_attrSimple("id"), "search-filter"), sky_call(sky_attrSimple("type"), "text"), sky_call(sky_attrSimple("placeholder"), "regex search"), sky_call(sky_evtHandler("input"), SetSearchFilter)}), sky_call(sky_call(sky_htmlEl("select"), []any{sky_call(sky_attrSimple("id"), "source-filter"), sky_call(sky_evtHandler("change"), SetSourceFilter)}), sky_call(sky_listAppend([]any{sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "")}), []any{sky_htmlText("Source")})}), sky_call(sky_listMap(func(s any) any {
			return sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), s)}), []any{sky_htmlText(s)})
		}), sources))), sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "spacer")}), []any{}), sky_call(sky_call(sky_htmlEl("button"), []any{sky_call(sky_attrSimple("class"), scrollClass), sky_call(sky_evtHandler("click"), map[string]any{"Tag": 7, "SkyName": "ToggleAutoScroll"})}), []any{sky_htmlText(scrollLabel)}), sky_call(sky_call(sky_htmlEl("button"), []any{sky_call(sky_attrSimple("class"), "btn danger"), sky_call(sky_evtHandler("click"), map[string]any{"Tag": 1, "SkyName": "ClearLogs"})}), []any{sky_htmlText("Clear")})}), viewEntries(filtered), sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "status")}), []any{sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "dot")}), []any{}), sky_htmlText(statusText(model)), sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "spacer")}), []any{}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "stats")}), []any{sky_htmlText(sky_concat(sky_stringFromInt(filteredCount), sky_concat("/", sky_stringFromInt(totalCount))))})})})
	}()
}

func main() {
	sky_runMainTask(sky_liveApp(map[string]any{"init": init_, "update": update, "view": view, "subscriptions": subscriptions, "routes": []any{sky_call(sky_liveRoute("/"), map[string]any{"Tag": 0, "SkyName": "LogPage"})}, "notFound": map[string]any{"Tag": 0, "SkyName": "LogPage"}}))
}
