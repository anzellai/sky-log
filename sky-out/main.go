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

type SkyADT struct {
	Tag     int
	SkyName string
	Fields  []any
}

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
var fileWriteString = Os_FileWriteString
var FileWriteString = Os_FileWriteString
var stderr = Os_Stderr
var Stderr = Os_Stderr
var stdin = Os_Stdin
var Stdin = Os_Stdin
var chdir = Os_Chdir
var Chdir = Os_Chdir
var chmod = Os_Chmod
var Chmod = Os_Chmod
var chown = Os_Chown
var Chown = Os_Chown
var chtimes = Os_Chtimes
var Chtimes = Os_Chtimes
var clearenv = Os_Clearenv()
var Clearenv = Os_Clearenv()
var copyFS = Os_CopyFS
var CopyFS = Os_CopyFS
var create = Os_Create
var Create = Os_Create
var createTemp = Os_CreateTemp
var CreateTemp = Os_CreateTemp
var dirFS = Os_DirFS
var DirFS = Os_DirFS
var environ = Os_Environ()
var Environ = Os_Environ()
var executable = Os_Executable()
var Executable = Os_Executable()
var exit = Os_Exit
var Exit = Os_Exit
var expand = Os_Expand
var Expand = Os_Expand
var expandEnv = Os_ExpandEnv
var ExpandEnv = Os_ExpandEnv
var findProcess = Os_FindProcess
var FindProcess = Os_FindProcess
var getegid = Os_Getegid()
var Getegid = Os_Getegid()
var getenv = Os_Getenv
var Getenv = Os_Getenv
var geteuid = Os_Geteuid()
var Geteuid = Os_Geteuid()
var getgid = Os_Getgid()
var Getgid = Os_Getgid()
var getgroups = Os_Getgroups()
var Getgroups = Os_Getgroups()
var getpagesize = Os_Getpagesize()
var Getpagesize = Os_Getpagesize()
var getpid = Os_Getpid()
var Getpid = Os_Getpid()
var getppid = Os_Getppid()
var Getppid = Os_Getppid()
var getuid = Os_Getuid()
var Getuid = Os_Getuid()
var getwd = Os_Getwd()
var Getwd = Os_Getwd()
var hostname = Os_Hostname()
var Hostname = Os_Hostname()
var isExist = Os_IsExist
var IsExist = Os_IsExist
var isNotExist = Os_IsNotExist
var IsNotExist = Os_IsNotExist
var isPathSeparator = Os_IsPathSeparator
var IsPathSeparator = Os_IsPathSeparator
var isPermission = Os_IsPermission
var IsPermission = Os_IsPermission
var isTimeout = Os_IsTimeout
var IsTimeout = Os_IsTimeout
var lchown = Os_Lchown
var Lchown = Os_Lchown
var link = Os_Link
var Link = Os_Link
var lookupEnv = Os_LookupEnv
var LookupEnv = Os_LookupEnv
var lstat = Os_Lstat
var Lstat = Os_Lstat
var mkdir = Os_Mkdir
var Mkdir = Os_Mkdir
var mkdirAll = Os_MkdirAll
var MkdirAll = Os_MkdirAll
var mkdirTemp = Os_MkdirTemp
var MkdirTemp = Os_MkdirTemp
var newFile = Os_NewFile
var NewFile = Os_NewFile
var newSyscallError = Os_NewSyscallError
var NewSyscallError = Os_NewSyscallError
var open = Os_Open
var Open = Os_Open
var openFile = Os_OpenFile
var OpenFile = Os_OpenFile
var openInRoot = Os_OpenInRoot
var OpenInRoot = Os_OpenInRoot
var openRoot = Os_OpenRoot
var OpenRoot = Os_OpenRoot
var pipe = Os_Pipe()
var Pipe = Os_Pipe()
var readDir = Os_ReadDir
var ReadDir = Os_ReadDir
var readFile = Os_ReadFile
var ReadFile = Os_ReadFile
var readlink = Os_Readlink
var Readlink = Os_Readlink
var remove = Os_Remove
var Remove = Os_Remove
var removeAll = Os_RemoveAll
var RemoveAll = Os_RemoveAll
var rename = Os_Rename
var Rename = Os_Rename
var sameFile = Os_SameFile
var SameFile = Os_SameFile
var setenv = Os_Setenv
var Setenv = Os_Setenv
var startProcess = Os_StartProcess
var StartProcess = Os_StartProcess
var stat = Os_Stat
var Stat = Os_Stat
var symlink = Os_Symlink
var Symlink = Os_Symlink
var tempDir = Os_TempDir()
var TempDir = Os_TempDir()
var truncate = Os_Truncate
var Truncate = Os_Truncate
var unsetenv = Os_Unsetenv
var Unsetenv = Os_Unsetenv
var userCacheDir = Os_UserCacheDir()
var UserCacheDir = Os_UserCacheDir()
var userConfigDir = Os_UserConfigDir()
var UserConfigDir = Os_UserConfigDir()
var userHomeDir = Os_UserHomeDir()
var UserHomeDir = Os_UserHomeDir()
var writeFile = Os_WriteFile
var WriteFile = Os_WriteFile
var args = Os_Args()
var Args = Os_Args()
var errClosed = Os_ErrClosed()
var ErrClosed = Os_ErrClosed()
var errDeadlineExceeded = Os_ErrDeadlineExceeded()
var ErrDeadlineExceeded = Os_ErrDeadlineExceeded()
var errExist = Os_ErrExist()
var ErrExist = Os_ErrExist()
var errInvalid = Os_ErrInvalid()
var ErrInvalid = Os_ErrInvalid()
var errNoDeadline = Os_ErrNoDeadline()
var ErrNoDeadline = Os_ErrNoDeadline()
var errNoHandle = Os_ErrNoHandle()
var ErrNoHandle = Os_ErrNoHandle()
var errNotExist = Os_ErrNotExist()
var ErrNotExist = Os_ErrNotExist()
var errPermission = Os_ErrPermission()
var ErrPermission = Os_ErrPermission()
var errProcessDone = Os_ErrProcessDone()
var ErrProcessDone = Os_ErrProcessDone()
var interrupt = Os_Interrupt()
var Interrupt = Os_Interrupt()
var kill = Os_Kill()
var Kill = Os_Kill()
var stdout = Os_Stdout()
var Stdout = Os_Stdout()
var devNull = Os_DevNull()
var DevNull = Os_DevNull()
var modeAppend = Os_ModeAppend()
var ModeAppend = Os_ModeAppend()
var modeCharDevice = Os_ModeCharDevice()
var ModeCharDevice = Os_ModeCharDevice()
var modeDevice = Os_ModeDevice()
var ModeDevice = Os_ModeDevice()
var modeDir = Os_ModeDir()
var ModeDir = Os_ModeDir()
var modeExclusive = Os_ModeExclusive()
var ModeExclusive = Os_ModeExclusive()
var modeIrregular = Os_ModeIrregular()
var ModeIrregular = Os_ModeIrregular()
var modeNamedPipe = Os_ModeNamedPipe()
var ModeNamedPipe = Os_ModeNamedPipe()
var modePerm = Os_ModePerm()
var ModePerm = Os_ModePerm()
var modeSetgid = Os_ModeSetgid()
var ModeSetgid = Os_ModeSetgid()
var modeSetuid = Os_ModeSetuid()
var ModeSetuid = Os_ModeSetuid()
var modeSocket = Os_ModeSocket()
var ModeSocket = Os_ModeSocket()
var modeSticky = Os_ModeSticky()
var ModeSticky = Os_ModeSticky()
var modeSymlink = Os_ModeSymlink()
var ModeSymlink = Os_ModeSymlink()
var modeTemporary = Os_ModeTemporary()
var ModeTemporary = Os_ModeTemporary()
var modeType = Os_ModeType()
var ModeType = Os_ModeType()
var o_APPEND = Os_O_APPEND()
var O_APPEND = Os_O_APPEND()
var o_CREATE = Os_O_CREATE()
var O_CREATE = Os_O_CREATE()
var o_EXCL = Os_O_EXCL()
var O_EXCL = Os_O_EXCL()
var o_RDONLY = Os_O_RDONLY()
var O_RDONLY = Os_O_RDONLY()
var o_RDWR = Os_O_RDWR()
var O_RDWR = Os_O_RDWR()
var o_SYNC = Os_O_SYNC()
var O_SYNC = Os_O_SYNC()
var o_TRUNC = Os_O_TRUNC()
var O_TRUNC = Os_O_TRUNC()
var o_WRONLY = Os_O_WRONLY()
var O_WRONLY = Os_O_WRONLY()
var pathListSeparator = Os_PathListSeparator()
var PathListSeparator = Os_PathListSeparator()
var pathSeparator = Os_PathSeparator()
var PathSeparator = Os_PathSeparator()
var sEEK_CUR = Os_SEEK_CUR()
var SEEK_CUR = Os_SEEK_CUR()
var sEEK_END = Os_SEEK_END()
var SEEK_END = Os_SEEK_END()
var sEEK_SET = Os_SEEK_SET()
var SEEK_SET = Os_SEEK_SET()
var fileChdir = Os_FileChdir
var FileChdir = Os_FileChdir
var fileChmod = Os_FileChmod
var FileChmod = Os_FileChmod
var fileChown = Os_FileChown
var FileChown = Os_FileChown
var fileClose = Os_FileClose
var FileClose = Os_FileClose
var fileFd = Os_FileFd
var FileFd = Os_FileFd
var fileName = Os_FileName
var FileName = Os_FileName
var fileRead = Os_FileRead
var FileRead = Os_FileRead
var fileReadAt = Os_FileReadAt
var FileReadAt = Os_FileReadAt
var fileReadDir = Os_FileReadDir
var FileReadDir = Os_FileReadDir
var fileReadFrom = Os_FileReadFrom
var FileReadFrom = Os_FileReadFrom
var fileReaddir = Os_FileReaddir
var FileReaddir = Os_FileReaddir
var fileReaddirnames = Os_FileReaddirnames
var FileReaddirnames = Os_FileReaddirnames
var fileSeek = Os_FileSeek
var FileSeek = Os_FileSeek
var fileSetDeadline = Os_FileSetDeadline
var FileSetDeadline = Os_FileSetDeadline
var fileSetReadDeadline = Os_FileSetReadDeadline
var FileSetReadDeadline = Os_FileSetReadDeadline
var fileSetWriteDeadline = Os_FileSetWriteDeadline
var FileSetWriteDeadline = Os_FileSetWriteDeadline
var fileStat = Os_FileStat
var FileStat = Os_FileStat
var fileSync = Os_FileSync
var FileSync = Os_FileSync
var fileSyscallConn = Os_FileSyscallConn
var FileSyscallConn = Os_FileSyscallConn
var fileTruncate = Os_FileTruncate
var FileTruncate = Os_FileTruncate
var fileWrite = Os_FileWrite
var FileWrite = Os_FileWrite
var fileWriteAt = Os_FileWriteAt
var FileWriteAt = Os_FileWriteAt
var fileWriteTo = Os_FileWriteTo
var FileWriteTo = Os_FileWriteTo
var linkErrorError = Os_LinkErrorError
var LinkErrorError = Os_LinkErrorError
var linkErrorUnwrap = Os_LinkErrorUnwrap
var LinkErrorUnwrap = Os_LinkErrorUnwrap
var linkErrorOp = Os_LinkErrorOp
var LinkErrorOp = Os_LinkErrorOp
var linkErrorOld = Os_LinkErrorOld
var LinkErrorOld = Os_LinkErrorOld
var linkErrorNew = Os_LinkErrorNew
var LinkErrorNew = Os_LinkErrorNew
var linkErrorErr = Os_LinkErrorErr
var LinkErrorErr = Os_LinkErrorErr
var procAttrDir = Os_ProcAttrDir
var ProcAttrDir = Os_ProcAttrDir
var procAttrEnv = Os_ProcAttrEnv
var ProcAttrEnv = Os_ProcAttrEnv
var procAttrFiles = Os_ProcAttrFiles
var ProcAttrFiles = Os_ProcAttrFiles
var procAttrSys = Os_ProcAttrSys
var ProcAttrSys = Os_ProcAttrSys
var processKill = Os_ProcessKill
var ProcessKill = Os_ProcessKill
var processRelease = Os_ProcessRelease
var ProcessRelease = Os_ProcessRelease
var processSignal = Os_ProcessSignal
var ProcessSignal = Os_ProcessSignal
var processWait = Os_ProcessWait
var ProcessWait = Os_ProcessWait
var processWithHandle = Os_ProcessWithHandle
var ProcessWithHandle = Os_ProcessWithHandle
var processPid = Os_ProcessPid
var ProcessPid = Os_ProcessPid
var processStateExitCode = Os_ProcessStateExitCode
var ProcessStateExitCode = Os_ProcessStateExitCode
var processStateExited = Os_ProcessStateExited
var ProcessStateExited = Os_ProcessStateExited
var processStatePid = Os_ProcessStatePid
var ProcessStatePid = Os_ProcessStatePid
var processStateString = Os_ProcessStateString
var ProcessStateString = Os_ProcessStateString
var processStateSuccess = Os_ProcessStateSuccess
var ProcessStateSuccess = Os_ProcessStateSuccess
var processStateSys = Os_ProcessStateSys
var ProcessStateSys = Os_ProcessStateSys
var processStateSysUsage = Os_ProcessStateSysUsage
var ProcessStateSysUsage = Os_ProcessStateSysUsage
var processStateSystemTime = Os_ProcessStateSystemTime
var ProcessStateSystemTime = Os_ProcessStateSystemTime
var processStateUserTime = Os_ProcessStateUserTime
var ProcessStateUserTime = Os_ProcessStateUserTime
var rootChmod = Os_RootChmod
var RootChmod = Os_RootChmod
var rootChown = Os_RootChown
var RootChown = Os_RootChown
var rootChtimes = Os_RootChtimes
var RootChtimes = Os_RootChtimes
var rootClose = Os_RootClose
var RootClose = Os_RootClose
var rootCreate = Os_RootCreate
var RootCreate = Os_RootCreate
var rootFS = Os_RootFS
var RootFS = Os_RootFS
var rootLchown = Os_RootLchown
var RootLchown = Os_RootLchown
var rootLink = Os_RootLink
var RootLink = Os_RootLink
var rootLstat = Os_RootLstat
var RootLstat = Os_RootLstat
var rootMkdir = Os_RootMkdir
var RootMkdir = Os_RootMkdir
var rootMkdirAll = Os_RootMkdirAll
var RootMkdirAll = Os_RootMkdirAll
var rootName = Os_RootName
var RootName = Os_RootName
var rootOpen = Os_RootOpen
var RootOpen = Os_RootOpen
var rootOpenFile = Os_RootOpenFile
var RootOpenFile = Os_RootOpenFile
var rootOpenRoot = Os_RootOpenRoot
var RootOpenRoot = Os_RootOpenRoot
var rootReadFile = Os_RootReadFile
var RootReadFile = Os_RootReadFile
var rootReadlink = Os_RootReadlink
var RootReadlink = Os_RootReadlink
var rootRemove = Os_RootRemove
var RootRemove = Os_RootRemove
var rootRemoveAll = Os_RootRemoveAll
var RootRemoveAll = Os_RootRemoveAll
var rootRename = Os_RootRename
var RootRename = Os_RootRename
var rootStat = Os_RootStat
var RootStat = Os_RootStat
var rootSymlink = Os_RootSymlink
var RootSymlink = Os_RootSymlink
var rootWriteFile = Os_RootWriteFile
var RootWriteFile = Os_RootWriteFile
var signalSignal = Os_SignalSignal
var SignalSignal = Os_SignalSignal
var signalString = Os_SignalString
var SignalString = Os_SignalString
var syscallErrorError = Os_SyscallErrorError
var SyscallErrorError = Os_SyscallErrorError
var syscallErrorTimeout = Os_SyscallErrorTimeout
var SyscallErrorTimeout = Os_SyscallErrorTimeout
var syscallErrorUnwrap = Os_SyscallErrorUnwrap
var SyscallErrorUnwrap = Os_SyscallErrorUnwrap
var syscallErrorSyscall = Os_SyscallErrorSyscall
var SyscallErrorSyscall = Os_SyscallErrorSyscall
var syscallErrorErr = Os_SyscallErrorErr
var SyscallErrorErr = Os_SyscallErrorErr
var command = Os_Exec_Command
var Command = Os_Exec_Command
var commandContext = Os_Exec_CommandContext
var CommandContext = Os_Exec_CommandContext
var lookPath = Os_Exec_LookPath
var LookPath = Os_Exec_LookPath
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
var newCmd = Os_Exec_NewCmd
var NewCmd = Os_Exec_NewCmd
var cmdPath = Os_Exec_CmdPath
var CmdPath = Os_Exec_CmdPath
var cmdSetPath = Os_Exec_CmdSetPath
var CmdSetPath = Os_Exec_CmdSetPath
var cmdArgs = Os_Exec_CmdArgs
var CmdArgs = Os_Exec_CmdArgs
var cmdSetArgs = Os_Exec_CmdSetArgs
var CmdSetArgs = Os_Exec_CmdSetArgs
var cmdEnv = Os_Exec_CmdEnv
var CmdEnv = Os_Exec_CmdEnv
var cmdSetEnv = Os_Exec_CmdSetEnv
var CmdSetEnv = Os_Exec_CmdSetEnv
var cmdDir = Os_Exec_CmdDir
var CmdDir = Os_Exec_CmdDir
var cmdSetDir = Os_Exec_CmdSetDir
var CmdSetDir = Os_Exec_CmdSetDir
var cmdStdin = Os_Exec_CmdStdin
var CmdStdin = Os_Exec_CmdStdin
var cmdSetStdin = Os_Exec_CmdSetStdin
var CmdSetStdin = Os_Exec_CmdSetStdin
var cmdStdout = Os_Exec_CmdStdout
var CmdStdout = Os_Exec_CmdStdout
var cmdSetStdout = Os_Exec_CmdSetStdout
var CmdSetStdout = Os_Exec_CmdSetStdout
var cmdStderr = Os_Exec_CmdStderr
var CmdStderr = Os_Exec_CmdStderr
var cmdSetStderr = Os_Exec_CmdSetStderr
var CmdSetStderr = Os_Exec_CmdSetStderr
var cmdExtraFiles = Os_Exec_CmdExtraFiles
var CmdExtraFiles = Os_Exec_CmdExtraFiles
var cmdSetExtraFiles = Os_Exec_CmdSetExtraFiles
var CmdSetExtraFiles = Os_Exec_CmdSetExtraFiles
var cmdSysProcAttr = Os_Exec_CmdSysProcAttr
var CmdSysProcAttr = Os_Exec_CmdSysProcAttr
var cmdSetSysProcAttr = Os_Exec_CmdSetSysProcAttr
var CmdSetSysProcAttr = Os_Exec_CmdSetSysProcAttr
var cmdProcess = Os_Exec_CmdProcess
var CmdProcess = Os_Exec_CmdProcess
var cmdSetProcess = Os_Exec_CmdSetProcess
var CmdSetProcess = Os_Exec_CmdSetProcess
var cmdProcessState = Os_Exec_CmdProcessState
var CmdProcessState = Os_Exec_CmdProcessState
var cmdSetProcessState = Os_Exec_CmdSetProcessState
var CmdSetProcessState = Os_Exec_CmdSetProcessState
var cmdErr = Os_Exec_CmdErr
var CmdErr = Os_Exec_CmdErr
var cmdSetErr = Os_Exec_CmdSetErr
var CmdSetErr = Os_Exec_CmdSetErr
var cmdWaitDelay = Os_Exec_CmdWaitDelay
var CmdWaitDelay = Os_Exec_CmdWaitDelay
var cmdSetWaitDelay = Os_Exec_CmdSetWaitDelay
var CmdSetWaitDelay = Os_Exec_CmdSetWaitDelay
var newError = Os_Exec_NewError
var NewError = Os_Exec_NewError
var errorName = Os_Exec_ErrorName
var ErrorName = Os_Exec_ErrorName
var errorSetName = Os_Exec_ErrorSetName
var ErrorSetName = Os_Exec_ErrorSetName
var errorErr = Os_Exec_ErrorErr
var ErrorErr = Os_Exec_ErrorErr
var errorSetErr = Os_Exec_ErrorSetErr
var ErrorSetErr = Os_Exec_ErrorSetErr
var newExitError = Os_Exec_NewExitError
var NewExitError = Os_Exec_NewExitError
var exitErrorProcessState = Os_Exec_ExitErrorProcessState
var ExitErrorProcessState = Os_Exec_ExitErrorProcessState
var exitErrorSetProcessState = Os_Exec_ExitErrorSetProcessState
var ExitErrorSetProcessState = Os_Exec_ExitErrorSetProcessState
var exitErrorStderr = Os_Exec_ExitErrorStderr
var ExitErrorStderr = Os_Exec_ExitErrorStderr
var exitErrorSetStderr = Os_Exec_ExitErrorSetStderr
var ExitErrorSetStderr = Os_Exec_ExitErrorSetStderr
var errDot = Os_Exec_ErrDot()
var ErrDot = Os_Exec_ErrDot()
var errNotFound = Os_Exec_ErrNotFound()
var ErrNotFound = Os_Exec_ErrNotFound()
var errWaitDelay = Os_Exec_ErrWaitDelay()
var ErrWaitDelay = Os_Exec_ErrWaitDelay()
var cmdCancel = Os_Exec_CmdCancel
var CmdCancel = Os_Exec_CmdCancel
var exitErrorSystemTime = Os_Exec_ExitErrorSystemTime
var ExitErrorSystemTime = Os_Exec_ExitErrorSystemTime
var exitErrorUserTime = Os_Exec_ExitErrorUserTime
var ExitErrorUserTime = Os_Exec_ExitErrorUserTime
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
var regexpAppendText = Regexp_RegexpAppendText
var RegexpAppendText = Regexp_RegexpAppendText
var regexpCopy = Regexp_RegexpCopy
var RegexpCopy = Regexp_RegexpCopy
var regexpExpand = Regexp_RegexpExpand
var RegexpExpand = Regexp_RegexpExpand
var regexpExpandString = Regexp_RegexpExpandString
var RegexpExpandString = Regexp_RegexpExpandString
var regexpFind = Regexp_RegexpFind
var RegexpFind = Regexp_RegexpFind
var regexpFindAll = Regexp_RegexpFindAll
var RegexpFindAll = Regexp_RegexpFindAll
var regexpFindAllString = Regexp_RegexpFindAllString
var RegexpFindAllString = Regexp_RegexpFindAllString
var regexpFindAllSubmatch = Regexp_RegexpFindAllSubmatch
var RegexpFindAllSubmatch = Regexp_RegexpFindAllSubmatch
var regexpFindIndex = Regexp_RegexpFindIndex
var RegexpFindIndex = Regexp_RegexpFindIndex
var regexpFindString = Regexp_RegexpFindString
var RegexpFindString = Regexp_RegexpFindString
var regexpFindStringIndex = Regexp_RegexpFindStringIndex
var RegexpFindStringIndex = Regexp_RegexpFindStringIndex
var regexpFindStringSubmatch = Regexp_RegexpFindStringSubmatch
var RegexpFindStringSubmatch = Regexp_RegexpFindStringSubmatch
var regexpFindStringSubmatchIndex = Regexp_RegexpFindStringSubmatchIndex
var RegexpFindStringSubmatchIndex = Regexp_RegexpFindStringSubmatchIndex
var regexpFindSubmatch = Regexp_RegexpFindSubmatch
var RegexpFindSubmatch = Regexp_RegexpFindSubmatch
var regexpFindSubmatchIndex = Regexp_RegexpFindSubmatchIndex
var RegexpFindSubmatchIndex = Regexp_RegexpFindSubmatchIndex
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
var newRegexp = Regexp_NewRegexp
var NewRegexp = Regexp_NewRegexp
var Debug = Log_Entry_Debug
var Info = Log_Entry_Info
var Warn = Log_Entry_Warn
var ErrorLevel = Log_Entry_ErrorLevel
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
var parseRawLine = Log_Entry_ParseRawLine
var ParseRawLine = Log_Entry_ParseRawLine
var extractTimestamp = Log_Entry_ExtractTimestamp
var ExtractTimestamp = Log_Entry_ExtractTimestamp
var detectLevel = Log_Entry_DetectLevel
var DetectLevel = Log_Entry_DetectLevel
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
var Log_Entry_Debug = SkyADT{Tag: 0, SkyName: "Debug"}
var Log_Entry_Info = SkyADT{Tag: 1, SkyName: "Info"}
var Log_Entry_Warn = SkyADT{Tag: 2, SkyName: "Warn"}
var Log_Entry_ErrorLevel = SkyADT{Tag: 3, SkyName: "ErrorLevel"}
var Log_Config_NoSection = SkyADT{Tag: 0, SkyName: "NoSection"}
var Os_File = SkyADT{Tag: 0, SkyName: "File"}
var Os_LinkError = SkyADT{Tag: 0, SkyName: "LinkError"}
var Os_ProcAttr = SkyADT{Tag: 0, SkyName: "ProcAttr"}
var Os_Process = SkyADT{Tag: 0, SkyName: "Process"}
var Os_ProcessState = SkyADT{Tag: 0, SkyName: "ProcessState"}
var Os_Root = SkyADT{Tag: 0, SkyName: "Root"}
var Os_Signal = SkyADT{Tag: 0, SkyName: "Signal"}
var Os_SyscallError = SkyADT{Tag: 0, SkyName: "SyscallError"}
var Os_Exec_Cmd = SkyADT{Tag: 0, SkyName: "Cmd"}
var Os_Exec_Error = SkyADT{Tag: 0, SkyName: "Error"}
var Os_Exec_ExitError = SkyADT{Tag: 0, SkyName: "ExitError"}
var Bufio_ReadWriter = SkyADT{Tag: 0, SkyName: "ReadWriter"}
var Bufio_Reader = SkyADT{Tag: 0, SkyName: "Reader"}
var Bufio_Scanner = SkyADT{Tag: 0, SkyName: "Scanner"}
var Bufio_Writer = SkyADT{Tag: 0, SkyName: "Writer"}
var Regexp_Regexp = SkyADT{Tag: 0, SkyName: "Regexp"}
var LogPage = SkyADT{Tag: 0, SkyName: "LogPage"}
var FileMode = SkyADT{Tag: 0, SkyName: "FileMode"}
var StreamMode = SkyADT{Tag: 1, SkyName: "StreamMode"}
var Tick = SkyADT{Tag: 0, SkyName: "Tick"}
var ToggleAutoScroll = SkyADT{Tag: 5, SkyName: "ToggleAutoScroll"}
var ToggleTheme = SkyADT{Tag: 6, SkyName: "ToggleTheme"}
var ClearLogs = SkyADT{Tag: 7, SkyName: "ClearLogs"}
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
var scanBytes = Bufio_ScanBytes
var ScanBytes = Bufio_ScanBytes
var scanLines = Bufio_ScanLines
var ScanLines = Bufio_ScanLines
var scanRunes = Bufio_ScanRunes
var ScanRunes = Bufio_ScanRunes
var scanWords = Bufio_ScanWords
var ScanWords = Bufio_ScanWords
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
var readWriterReadByte = Bufio_ReadWriterReadByte
var ReadWriterReadByte = Bufio_ReadWriterReadByte
var readWriterReadBytes = Bufio_ReadWriterReadBytes
var ReadWriterReadBytes = Bufio_ReadWriterReadBytes
var readWriterReadFrom = Bufio_ReadWriterReadFrom
var ReadWriterReadFrom = Bufio_ReadWriterReadFrom
var readWriterReadLine = Bufio_ReadWriterReadLine
var ReadWriterReadLine = Bufio_ReadWriterReadLine
var readWriterReadRune = Bufio_ReadWriterReadRune
var ReadWriterReadRune = Bufio_ReadWriterReadRune
var readWriterReadSlice = Bufio_ReadWriterReadSlice
var ReadWriterReadSlice = Bufio_ReadWriterReadSlice
var readWriterReadString = Bufio_ReadWriterReadString
var ReadWriterReadString = Bufio_ReadWriterReadString
var readWriterUnreadByte = Bufio_ReadWriterUnreadByte
var ReadWriterUnreadByte = Bufio_ReadWriterUnreadByte
var readWriterUnreadRune = Bufio_ReadWriterUnreadRune
var ReadWriterUnreadRune = Bufio_ReadWriterUnreadRune
var readWriterWrite = Bufio_ReadWriterWrite
var ReadWriterWrite = Bufio_ReadWriterWrite
var readWriterWriteByte = Bufio_ReadWriterWriteByte
var ReadWriterWriteByte = Bufio_ReadWriterWriteByte
var readWriterWriteRune = Bufio_ReadWriterWriteRune
var ReadWriterWriteRune = Bufio_ReadWriterWriteRune
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
var readerReadByte = Bufio_ReaderReadByte
var ReaderReadByte = Bufio_ReaderReadByte
var readerReadBytes = Bufio_ReaderReadBytes
var ReaderReadBytes = Bufio_ReaderReadBytes
var readerReadLine = Bufio_ReaderReadLine
var ReaderReadLine = Bufio_ReaderReadLine
var readerReadRune = Bufio_ReaderReadRune
var ReaderReadRune = Bufio_ReaderReadRune
var readerReadSlice = Bufio_ReaderReadSlice
var ReaderReadSlice = Bufio_ReaderReadSlice
var readerReadString = Bufio_ReaderReadString
var ReaderReadString = Bufio_ReaderReadString
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
var writerWriteByte = Bufio_WriterWriteByte
var WriterWriteByte = Bufio_WriterWriteByte
var writerWriteRune = Bufio_WriterWriteRune
var WriterWriteRune = Bufio_WriterWriteRune
var writerWriteString = Bufio_WriterWriteString
var WriterWriteString = Bufio_WriterWriteString
var readWriterReader = Bufio_ReadWriterReader
var ReadWriterReader = Bufio_ReadWriterReader
var readWriterSetReader = Bufio_ReadWriterSetReader
var ReadWriterSetReader = Bufio_ReadWriterSetReader
var readWriterWriter = Bufio_ReadWriterWriter
var ReadWriterWriter = Bufio_ReadWriterWriter
var readWriterSetWriter = Bufio_ReadWriterSetWriter
var ReadWriterSetWriter = Bufio_ReadWriterSetWriter
var errAdvanceTooFar = Bufio_ErrAdvanceTooFar()
var ErrAdvanceTooFar = Bufio_ErrAdvanceTooFar()
var errBadReadCount = Bufio_ErrBadReadCount()
var ErrBadReadCount = Bufio_ErrBadReadCount()
var errBufferFull = Bufio_ErrBufferFull()
var ErrBufferFull = Bufio_ErrBufferFull()
var errFinalToken = Bufio_ErrFinalToken()
var ErrFinalToken = Bufio_ErrFinalToken()
var errInvalidUnreadByte = Bufio_ErrInvalidUnreadByte()
var ErrInvalidUnreadByte = Bufio_ErrInvalidUnreadByte()
var errInvalidUnreadRune = Bufio_ErrInvalidUnreadRune()
var ErrInvalidUnreadRune = Bufio_ErrInvalidUnreadRune()
var errNegativeAdvance = Bufio_ErrNegativeAdvance()
var ErrNegativeAdvance = Bufio_ErrNegativeAdvance()
var errNegativeCount = Bufio_ErrNegativeCount()
var ErrNegativeCount = Bufio_ErrNegativeCount()
var errTooLong = Bufio_ErrTooLong()
var ErrTooLong = Bufio_ErrTooLong()
var maxScanTokenSize = Bufio_MaxScanTokenSize()
var MaxScanTokenSize = Bufio_MaxScanTokenSize()
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

func sky_adtTag(v any) int {
	if a, ok := v.(SkyADT); ok {
		return a.Tag
	}
	if m, ok := v.(map[string]any); ok {
		return sky_asInt(m["Tag"])
	}
	return -1
}

func sky_adtField(v any, idx int) any {
	if a, ok := v.(SkyADT); ok {
		if idx < len(a.Fields) {
			return a.Fields[idx]
		}
		return nil
	}
	if m, ok := v.(map[string]any); ok {
		return m[fmt.Sprintf("V%d", idx)]
	}
	return nil
}

func sky_getSkyName(v any) string {
	if a, ok := v.(SkyADT); ok {
		return a.SkyName
	}
	if m, ok := v.(map[string]any); ok {
		if s, ok := m["SkyName"].(string); ok {
			return s
		}
		return ""
	}
	return ""
}

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

func sky_asError(v any) error {
	if e, ok := v.(error); ok {
		return e
	}
	return fmt.Errorf("%v", v)
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

func sky_asInt64(v any) int64 { return int64(sky_asInt(v)) }

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

func sky_callZeroOrNil(f any) any {
	if fn, ok := f.(func() any); ok {
		return fn()
	}
	if fn, ok := f.(func(any) any); ok {
		return fn(nil)
	}
	return f
}

func sky_asMap(v any) map[string]any {
	if m, ok := v.(map[string]any); ok {
		return m
	}
	if m2, ok := v.(map[any]any); ok {
		r := make(map[string]any, len(m2))
		for k, val := range m2 {
			r[fmt.Sprintf("%v", k)] = val
		}
		return r
	}
	if a, ok := v.(SkyADT); ok {
		m := map[string]any{"Tag": a.Tag, "SkyName": a.SkyName}
		for i, f := range a.Fields {
			m[fmt.Sprintf("V%d", i)] = f
		}
		return m
	}
	if t, ok := v.(SkyTuple2); ok {
		return map[string]any{"V0": t.V0, "V1": t.V1}
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

func sky_stringToInt(s any) any {
	n, err := strconv.Atoi(strings.TrimSpace(sky_asString(s)))
	if err != nil {
		return SkyNothing()
	}
	return SkyJust(n)
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

func sky_listHead(list any) any {
	items := sky_asList(list)
	if len(items) > 0 {
		return SkyJust(items[0])
	}
	return SkyNothing()
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

func sky_maybeWithDefault(def any) any {
	return func(m any) any {
		mb := sky_asSkyMaybe(m)
		if mb.Tag == 0 {
			return mb.JustValue
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

func sky_htmlStyleNode(attrs any) any {
	return func(css any) any {
		return map[string]any{"tag": "style", "attrs": attrs, "children": []any{map[string]any{"tag": "", "attrs": []any{}, "children": []any{}, "text": sky_asString(css)}}, "text": ""}
	}
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
	if a, ok := msg.(SkyADT); ok {
		return a.SkyName
	}
	if m, ok := msg.(map[string]any); ok {
		if name, exists := m["SkyName"]; exists {
			return sky_asString(name)
		}
	}
	if fn, ok := msg.(func(any) any); ok {
		result := fn(nil)
		if a2, ok2 := result.(SkyADT); ok2 {
			return a2.SkyName
		}
		if m2, ok2 := result.(map[string]any); ok2 {
			if name, exists := m2["SkyName"]; exists {
				return sky_asString(name)
			}
		}
	}
	return fmt.Sprintf("%v", msg)
}

func sky_msgArgs(msg any) []any {
	if a, ok := msg.(SkyADT); ok {
		var args []any
		for _, v := range a.Fields {
			if v == nil {
				break
			}
			if va, ok := v.(SkyADT); ok {
				args = append(args, va.SkyName)
				continue
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
		if va, ok := v.(SkyADT); ok {
			args = append(args, va.SkyName)
			continue
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

func sky_cssStylesheet(rules any) any {
	var sb strings.Builder
	for _, r := range sky_asList(rules) {
		sb.WriteString(sky_asString(r))
		sb.WriteString("\n")
	}
	return sb.String()
}

func sky_cssRule(selector any) any {
	return func(props any) any {
		var sb strings.Builder
		sb.WriteString(sky_asString(selector))
		sb.WriteString(" { ")
		for _, p := range sky_asList(props) {
			sb.WriteString(sky_asString(p))
			sb.WriteString("; ")
		}
		sb.WriteString("}")
		return sb.String()
	}
}

func sky_cssProp(key any) any {
	return func(val any) any { return sky_asString(key) + ": " + sky_asString(val) }
}

func sky_cssPx(n any) any { return fmt.Sprintf("%dpx", sky_asInt(n)) }

func sky_cssPct(n any) any { return fmt.Sprintf("%.0f%%", sky_asFloat(n)) }

func sky_cssHex(s any) any { return "#" + sky_asString(s) }

func sky_cssVal(v any) any { return func(_ any) any { return sky_asString(v) } }

func sky_cssStyles(props any) any {
	var parts []string
	for _, p := range sky_asList(props) {
		parts = append(parts, sky_asString(p))
	}
	return strings.Join(parts, "; ")
}

func sky_cssPadding2(v any) any {
	return func(h any) any { return "padding: " + sky_asString(v) + " " + sky_asString(h) }
}

func sky_cssPropFn(prop any) any {
	return func(val any) any { return sky_asString(prop) + ": " + sky_asString(val) }
}

func sky_liveRoute(path any) any {
	return func(page any) any { return map[string]any{"path": path, "page": page} }
}

func sky_liveApp(config any) any { return sky_liveAppImpl(config) }

// sky:type logMsg : String -> Unit

func Log_Webhook_LogMsg(msg any) any {
	return func() any { Os_FileWriteString(Os_Stderr(struct{}{}), sky_concat(msg, "\n")); return struct{}{} }()
}

// sky:type buildRules : any -> any

func Log_Webhook_BuildRules(config any) any {
	return sky_call(sky_listFilterMap(func(source any) any {
		return func() any {
			filter := func() any {
				if sky_asBool(sky_stringIsEmpty(sky_asMap(source)["filter"])) {
					return sky_asMap(sky_asMap(config)["webhook"])["filter"]
				}
				return sky_asMap(source)["filter"]
			}()
			_ = filter
			url := func() any {
				if sky_asBool(sky_stringIsEmpty(sky_asMap(source)["webhookUrl"])) {
					return sky_asMap(sky_asMap(config)["webhook"])["url"]
				}
				return sky_asMap(source)["webhookUrl"]
			}()
			_ = url
			return func() any {
				if sky_asBool(sky_asBool(sky_stringIsEmpty(filter)) || sky_asBool(sky_stringIsEmpty(url))) {
					return SkyNothing()
				}
				return SkyJust(map[string]any{"filter": filter, "url": url, "sourceName": sky_asMap(source)["name"]})
			}()
		}()
	}), sky_asMap(config)["sources"])
}

// sky:type matchesFilter : String -> any -> Bool

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

// sky:type buildPayload : any -> any

func Log_Webhook_BuildPayload(entry any) any {
	return sky_call(sky_jsonEncode(0), sky_jsonEncObject([]any{SkyTuple2{V0: "source", V1: sky_jsonEncString(sky_asMap(entry)["source"])}, SkyTuple2{V0: "level", V1: sky_jsonEncString(levelToString(sky_asMap(entry)["level"]))}, SkyTuple2{V0: "scope", V1: sky_jsonEncString(sky_asMap(entry)["scope"])}, SkyTuple2{V0: "message", V1: sky_jsonEncString(sky_asMap(entry)["message"])}, SkyTuple2{V0: "timestamp", V1: sky_jsonEncInt(sky_asMap(entry)["timestamp"])}}))
}

// sky:type send : String -> any -> Unit

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

// sky:type processNewEntries : any -> any -> Unit

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

// sky:type levelFromString : any -> Level

func Log_Entry_LevelFromString(str any) any {
	return func() any {
		return func() any {
			__subject := sky_stringToLower(str)
			if sky_asString(__subject) == "debug" {
				return SkyADT{Tag: 0, SkyName: "Debug"}
			}
			if sky_asString(__subject) == "info" {
				return SkyADT{Tag: 1, SkyName: "Info"}
			}
			if sky_asString(__subject) == "warn" {
				return SkyADT{Tag: 2, SkyName: "Warn"}
			}
			if sky_asString(__subject) == "error" {
				return SkyADT{Tag: 3, SkyName: "ErrorLevel"}
			}
			if true {
				return SkyADT{Tag: 1, SkyName: "Info"}
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

// sky:type levelToString : Level -> String

func Log_Entry_LevelToString(level any) any {
	return func() any {
		return func() any {
			__subject := level
			__sky_tag := sky_adtTag(__subject)
			__sky_name := sky_getSkyName(__subject)
			_ = __sky_name
			if __sky_name == "Debug" || (__sky_name == "" && __sky_tag == 0) {
				return "DEBUG"
			}
			if __sky_name == "Info" || (__sky_name == "" && __sky_tag == 1) {
				return "INFO"
			}
			if __sky_name == "Warn" || (__sky_name == "" && __sky_tag == 2) {
				return "WARN"
			}
			if __sky_name == "ErrorLevel" || (__sky_name == "" && __sky_tag == 3) {
				return "ERROR"
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

// sky:type levelToInt : Level -> Int

func Log_Entry_LevelToInt(level any) any {
	return func() any {
		return func() any {
			__subject := level
			__sky_tag := sky_adtTag(__subject)
			__sky_name := sky_getSkyName(__subject)
			_ = __sky_name
			if __sky_name == "Debug" || (__sky_name == "" && __sky_tag == 0) {
				return 0
			}
			if __sky_name == "Info" || (__sky_name == "" && __sky_tag == 1) {
				return 1
			}
			if __sky_name == "Warn" || (__sky_name == "" && __sky_tag == 2) {
				return 2
			}
			if __sky_name == "ErrorLevel" || (__sky_name == "" && __sky_tag == 3) {
				return 3
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

// sky:type levelDecoder : any

func Log_Entry_LevelDecoder() any {
	return sky_call(sky_jsonDecMap(Log_Entry_LevelFromString), sky_jsonDecoder_string)
}

// sky:type entryDecoder : any

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

// sky:type decodeEntry : any -> any -> any

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
				return Log_Entry_ParseRawLine(line, source)
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

// sky:type parseRawLine : any -> any -> { level : t171 , scope : t168 , source : t168 , message : t167 , timestamp : t170 }

func Log_Entry_ParseRawLine(line any, source any) any {
	return func() any {
		ts := Log_Entry_ExtractTimestamp(line)
		_ = ts
		level := Log_Entry_DetectLevel(line)
		_ = level
		scope := func() any {
			if sky_asBool(sky_call(sky_stringContains("|"), line)) {
				return sky_stringTrim(func() any {
					return func() any {
						__subject := sky_call(sky_stringSplit("|"), line)
						if len(sky_asList(__subject)) > 0 {
							rest := sky_asList(__subject)[1:]
							_ = rest
							return ""
						}
						if true {
							return ""
						}
						panic("non-exhaustive case expression")
					}()
				}())
			}
			return "raw"
		}()
		_ = scope
		return map[string]any{"timestamp": ts, "level": level, "scope": source, "message": line, "source": source}
	}()
}

// sky:type extractTimestamp : any -> Int

func Log_Entry_ExtractTimestamp(line any) any {
	return func() any {
		prefix := sky_call(sky_call(sky_stringSlice(0), 19), line)
		_ = prefix
		hasIsoDate := sky_asBool(sky_numCompare(">=", sky_stringLength(prefix), 19)) && sky_asBool(sky_asBool(sky_call(sky_stringContains("-"), sky_call(sky_call(sky_stringSlice(0), 10), prefix))) && sky_asBool(sky_call(sky_stringContains(":"), sky_call(sky_call(sky_stringSlice(11), 19), prefix))))
		_ = hasIsoDate
		return func() any {
			if sky_asBool(hasIsoDate) {
				return func() any {
					timePart := sky_call(sky_call(sky_stringSlice(11), 19), prefix)
					_ = timePart
					parts := sky_call(sky_stringSplit(":"), timePart)
					_ = parts
					h := sky_call(sky_maybeWithDefault(0), sky_stringToInt(sky_call(sky_maybeWithDefault("0"), sky_listHead(parts))))
					_ = h
					m := sky_call(sky_maybeWithDefault(0), sky_stringToInt(sky_call(sky_maybeWithDefault("0"), sky_listHead(sky_call(sky_listDrop(1), parts)))))
					_ = m
					s := sky_call(sky_maybeWithDefault(0), sky_stringToInt(sky_call(sky_maybeWithDefault("0"), sky_listHead(sky_call(sky_listDrop(2), parts)))))
					_ = s
					return sky_numBinop("+", sky_numBinop("+", sky_numBinop("*", h, 3600), sky_numBinop("*", m, 60)), s)
				}()
			}
			return 0
		}()
	}()
}

// sky:type detectLevel : any -> Level

func Log_Entry_DetectLevel(line any) any {
	return func() any {
		lower := sky_stringToLower(line)
		_ = lower
		return func() any {
			if sky_asBool(sky_asBool(sky_call(sky_stringContains("error"), lower)) || sky_asBool(sky_call(sky_stringContains("err"), lower))) {
				return SkyADT{Tag: 3, SkyName: "ErrorLevel"}
			}
			if sky_asBool(sky_call(sky_stringContains("warn"), lower)) {
				return SkyADT{Tag: 2, SkyName: "Warn"}
			}
			if sky_asBool(sky_call(sky_stringContains("debug"), lower)) {
				return SkyADT{Tag: 0, SkyName: "Debug"}
			}
			return SkyADT{Tag: 1, SkyName: "Info"}
		}()
	}()
}

func Log_Config_InSource(v0 any) any {
	return SkyADT{Tag: 1, SkyName: "InSource", Fields: []any{v0}}
}

func Log_Config_InWebhook(v0 any) any {
	return SkyADT{Tag: 2, SkyName: "InWebhook", Fields: []any{v0}}
}

// sky:type emptySource : { command : String , filter : String , webhookUrl : String , name : String }

func Log_Config_EmptySource() any {
	return map[string]any{"name": "", "command": "", "filter": "", "webhookUrl": ""}
}

// sky:type emptyWebhook : { filter : String , url : String }

func Log_Config_EmptyWebhook() any {
	return map[string]any{"url": "", "filter": ""}
}

// sky:type parseConfig : any -> { sources : List elem , webhook : { filter : String , url : String } }

func Log_Config_ParseConfig(path any) any {
	return func() any {
		return func() any {
			__subject := sky_fileRead(path)
			if sky_asSkyResult(__subject).SkyName == "Ok" {
				content := sky_asSkyResult(__subject).OkValue
				_ = content
				return sky_call(func(__pa0 any) any {
					return Log_Config_ParseLines(map[string]any{"sources": []any{}, "webhook": Log_Config_EmptyWebhook(), "section": SkyADT{Tag: 0, SkyName: "NoSection"}}, __pa0)
				}, sky_call(sky_stringLines, content))
			}
			if sky_asSkyResult(__subject).SkyName == "Err" {
				return map[string]any{"sources": []any{}, "webhook": Log_Config_EmptyWebhook()}
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

// sky:type finalizeState : any -> { sources : t137 , webhook : WebhookConfig }

func Log_Config_FinalizeState(state any) any {
	return func() any {
		return func() any {
			__subject := sky_asMap(state)["section"]
			__sky_tag := sky_adtTag(__subject)
			__sky_name := sky_getSkyName(__subject)
			_ = __sky_name
			if __sky_name == "InSource" || (__sky_name == "" && __sky_tag == 1) {
				s := sky_adtField(__subject, 0)
				_ = s
				return map[string]any{"sources": sky_call(sky_listAppend(sky_asMap(state)["sources"]), []any{s}), "webhook": sky_asMap(state)["webhook"]}
			}
			if __sky_name == "InWebhook" || (__sky_name == "" && __sky_tag == 2) {
				w := sky_adtField(__subject, 0)
				_ = w
				return map[string]any{"sources": sky_asMap(state)["sources"], "webhook": w}
			}
			if __sky_name == "NoSection" || (__sky_name == "" && __sky_tag == 0) {
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
										__sky_tag := sky_adtTag(__subject)
										__sky_name := sky_getSkyName(__subject)
										_ = __sky_name
										if __sky_name == "InSource" || (__sky_name == "" && __sky_tag == 1) {
											s := sky_adtField(__subject, 0)
											_ = s
											return sky_recordUpdate(state, map[string]any{"sources": sky_call(sky_listAppend(sky_asMap(state)["sources"]), []any{s}), "section": InSource(Log_Config_EmptySource())})
										}
										if __sky_name == "InWebhook" || (__sky_name == "" && __sky_tag == 2) {
											w := sky_adtField(__subject, 0)
											_ = w
											return sky_recordUpdate(state, map[string]any{"webhook": w, "section": InSource(Log_Config_EmptySource())})
										}
										if __sky_name == "NoSection" || (__sky_name == "" && __sky_tag == 0) {
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
										__sky_tag := sky_adtTag(__subject)
										__sky_name := sky_getSkyName(__subject)
										_ = __sky_name
										if __sky_name == "InSource" || (__sky_name == "" && __sky_tag == 1) {
											s := sky_adtField(__subject, 0)
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
									__sky_tag := sky_adtTag(__subject)
									__sky_name := sky_getSkyName(__subject)
									_ = __sky_name
									if __sky_name == "InSource" || (__sky_name == "" && __sky_tag == 1) {
										s := sky_adtField(__subject, 0)
										_ = s
										return sky_recordUpdate(state, map[string]any{"section": InSource(Log_Config_ParseSourceLine(trimmed, s))})
									}
									if __sky_name == "InWebhook" || (__sky_name == "" && __sky_tag == 2) {
										w := sky_adtField(__subject, 0)
										_ = w
										return sky_recordUpdate(state, map[string]any{"section": InWebhook(Log_Config_ParseWebhookLine(trimmed, w))})
									}
									if __sky_name == "NoSection" || (__sky_name == "" && __sky_tag == 0) {
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

// sky:type parseSourceLine : any -> any -> any

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

// sky:type parseWebhookLine : any -> any -> any

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

// sky:type extractValue : any -> String

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

// sky:type stripQuotes : any -> any

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

func Os_Chdir(a0 any) any { return Sky_os_Chdir(a0) }

func Os_Chmod(a0 any, a1 any) any { return Sky_os_Chmod(a0, a1) }

func Os_Chown(a0 any, a1 any, a2 any) any { return Sky_os_Chown(a0, a1, a2) }

func Os_Chtimes(a0 any, a1 any, a2 any) any { return Sky_os_Chtimes(a0, a1, a2) }

func Os_Clearenv(_ ...any) any { return sky_callZeroOrNil(Sky_os_Clearenv) }

func Os_CopyFS(a0 any, a1 any) any { return Sky_os_CopyFS(a0, a1) }

func Os_Create(a0 any) any { return Sky_os_Create(a0) }

func Os_CreateTemp(a0 any, a1 any) any { return Sky_os_CreateTemp(a0, a1) }

func Os_DirFS(a0 any) any { return Sky_os_DirFS(a0) }

func Os_Environ(_ ...any) any { return sky_callZeroOrNil(Sky_os_Environ) }

func Os_Executable(_ ...any) any { return sky_callZeroOrNil(Sky_os_Executable) }

func Os_Exit(a0 any) any { return Sky_os_Exit(a0) }

func Os_Expand(a0 any, a1 any) any { return Sky_os_Expand(a0, a1) }

func Os_ExpandEnv(a0 any) any { return Sky_os_ExpandEnv(a0) }

func Os_FindProcess(a0 any) any { return Sky_os_FindProcess(a0) }

func Os_Getegid(_ ...any) any { return sky_callZeroOrNil(Sky_os_Getegid) }

func Os_Getenv(a0 any) any { return Sky_os_Getenv(a0) }

func Os_Geteuid(_ ...any) any { return sky_callZeroOrNil(Sky_os_Geteuid) }

func Os_Getgid(_ ...any) any { return sky_callZeroOrNil(Sky_os_Getgid) }

func Os_Getgroups(_ ...any) any { return sky_callZeroOrNil(Sky_os_Getgroups) }

func Os_Getpagesize(_ ...any) any { return sky_callZeroOrNil(Sky_os_Getpagesize) }

func Os_Getpid(_ ...any) any { return sky_callZeroOrNil(Sky_os_Getpid) }

func Os_Getppid(_ ...any) any { return sky_callZeroOrNil(Sky_os_Getppid) }

func Os_Getuid(_ ...any) any { return sky_callZeroOrNil(Sky_os_Getuid) }

func Os_Getwd(_ ...any) any { return sky_callZeroOrNil(Sky_os_Getwd) }

func Os_Hostname(_ ...any) any { return sky_callZeroOrNil(Sky_os_Hostname) }

func Os_IsExist(a0 any) any { return Sky_os_IsExist(a0) }

func Os_IsNotExist(a0 any) any { return Sky_os_IsNotExist(a0) }

func Os_IsPathSeparator(a0 any) any { return Sky_os_IsPathSeparator(a0) }

func Os_IsPermission(a0 any) any { return Sky_os_IsPermission(a0) }

func Os_IsTimeout(a0 any) any { return Sky_os_IsTimeout(a0) }

func Os_Lchown(a0 any, a1 any, a2 any) any { return Sky_os_Lchown(a0, a1, a2) }

func Os_Link(a0 any, a1 any) any { return Sky_os_Link(a0, a1) }

func Os_LookupEnv(a0 any) any { return Sky_os_LookupEnv(a0) }

func Os_Lstat(a0 any) any { return Sky_os_Lstat(a0) }

func Os_Mkdir(a0 any, a1 any) any { return Sky_os_Mkdir(a0, a1) }

func Os_MkdirAll(a0 any, a1 any) any { return Sky_os_MkdirAll(a0, a1) }

func Os_MkdirTemp(a0 any, a1 any) any { return Sky_os_MkdirTemp(a0, a1) }

func Os_NewFile(a0 any, a1 any) any { return Sky_os_NewFile(a0, a1) }

func Os_NewSyscallError(a0 any, a1 any) any { return Sky_os_NewSyscallError(a0, a1) }

func Os_Open(a0 any) any { return Sky_os_Open(a0) }

func Os_OpenFile(a0 any, a1 any, a2 any) any { return Sky_os_OpenFile(a0, a1, a2) }

func Os_OpenInRoot(a0 any, a1 any) any { return Sky_os_OpenInRoot(a0, a1) }

func Os_OpenRoot(a0 any) any { return Sky_os_OpenRoot(a0) }

func Os_Pipe(_ ...any) any { return sky_callZeroOrNil(Sky_os_Pipe) }

func Os_ReadDir(a0 any) any { return Sky_os_ReadDir(a0) }

func Os_ReadFile(a0 any) any { return Sky_os_ReadFile(a0) }

func Os_Readlink(a0 any) any { return Sky_os_Readlink(a0) }

func Os_Remove(a0 any) any { return Sky_os_Remove(a0) }

func Os_RemoveAll(a0 any) any { return Sky_os_RemoveAll(a0) }

func Os_Rename(a0 any, a1 any) any { return Sky_os_Rename(a0, a1) }

func Os_SameFile(a0 any, a1 any) any { return Sky_os_SameFile(a0, a1) }

func Os_Setenv(a0 any, a1 any) any { return Sky_os_Setenv(a0, a1) }

func Os_StartProcess(a0 any, a1 any, a2 any) any { return Sky_os_StartProcess(a0, a1, a2) }

func Os_Stat(a0 any) any { return Sky_os_Stat(a0) }

func Os_Symlink(a0 any, a1 any) any { return Sky_os_Symlink(a0, a1) }

func Os_TempDir(_ ...any) any { return sky_callZeroOrNil(Sky_os_TempDir) }

func Os_Truncate(a0 any, a1 any) any { return Sky_os_Truncate(a0, a1) }

func Os_Unsetenv(a0 any) any { return Sky_os_Unsetenv(a0) }

func Os_UserCacheDir(_ ...any) any { return sky_callZeroOrNil(Sky_os_UserCacheDir) }

func Os_UserConfigDir(_ ...any) any { return sky_callZeroOrNil(Sky_os_UserConfigDir) }

func Os_UserHomeDir(_ ...any) any { return sky_callZeroOrNil(Sky_os_UserHomeDir) }

func Os_WriteFile(a0 any, a1 any, a2 any) any { return Sky_os_WriteFile(a0, a1, a2) }

func Os_Args(_ ...any) any { return sky_callZeroOrNil(Sky_os_Args) }

func Os_ErrClosed(_ ...any) any { return sky_callZeroOrNil(Sky_os_ErrClosed) }

func Os_ErrDeadlineExceeded(_ ...any) any { return sky_callZeroOrNil(Sky_os_ErrDeadlineExceeded) }

func Os_ErrExist(_ ...any) any { return sky_callZeroOrNil(Sky_os_ErrExist) }

func Os_ErrInvalid(_ ...any) any { return sky_callZeroOrNil(Sky_os_ErrInvalid) }

func Os_ErrNoDeadline(_ ...any) any { return sky_callZeroOrNil(Sky_os_ErrNoDeadline) }

func Os_ErrNoHandle(_ ...any) any { return sky_callZeroOrNil(Sky_os_ErrNoHandle) }

func Os_ErrNotExist(_ ...any) any { return sky_callZeroOrNil(Sky_os_ErrNotExist) }

func Os_ErrPermission(_ ...any) any { return sky_callZeroOrNil(Sky_os_ErrPermission) }

func Os_ErrProcessDone(_ ...any) any { return sky_callZeroOrNil(Sky_os_ErrProcessDone) }

func Os_Interrupt(_ ...any) any { return sky_callZeroOrNil(Sky_os_Interrupt) }

func Os_Kill(_ ...any) any { return sky_callZeroOrNil(Sky_os_Kill) }

func Os_Stdout(_ ...any) any { return sky_callZeroOrNil(Sky_os_Stdout) }

func Os_DevNull(_ ...any) any { return sky_callZeroOrNil(Sky_os_DevNull) }

func Os_ModeAppend(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeAppend) }

func Os_ModeCharDevice(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeCharDevice) }

func Os_ModeDevice(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeDevice) }

func Os_ModeDir(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeDir) }

func Os_ModeExclusive(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeExclusive) }

func Os_ModeIrregular(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeIrregular) }

func Os_ModeNamedPipe(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeNamedPipe) }

func Os_ModePerm(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModePerm) }

func Os_ModeSetgid(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeSetgid) }

func Os_ModeSetuid(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeSetuid) }

func Os_ModeSocket(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeSocket) }

func Os_ModeSticky(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeSticky) }

func Os_ModeSymlink(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeSymlink) }

func Os_ModeTemporary(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeTemporary) }

func Os_ModeType(_ ...any) any { return sky_callZeroOrNil(Sky_os_ModeType) }

func Os_O_APPEND(_ ...any) any { return sky_callZeroOrNil(Sky_os_O_APPEND) }

func Os_O_CREATE(_ ...any) any { return sky_callZeroOrNil(Sky_os_O_CREATE) }

func Os_O_EXCL(_ ...any) any { return sky_callZeroOrNil(Sky_os_O_EXCL) }

func Os_O_RDONLY(_ ...any) any { return sky_callZeroOrNil(Sky_os_O_RDONLY) }

func Os_O_RDWR(_ ...any) any { return sky_callZeroOrNil(Sky_os_O_RDWR) }

func Os_O_SYNC(_ ...any) any { return sky_callZeroOrNil(Sky_os_O_SYNC) }

func Os_O_TRUNC(_ ...any) any { return sky_callZeroOrNil(Sky_os_O_TRUNC) }

func Os_O_WRONLY(_ ...any) any { return sky_callZeroOrNil(Sky_os_O_WRONLY) }

func Os_PathListSeparator(_ ...any) any { return sky_callZeroOrNil(Sky_os_PathListSeparator) }

func Os_PathSeparator(_ ...any) any { return sky_callZeroOrNil(Sky_os_PathSeparator) }

func Os_SEEK_CUR(_ ...any) any { return sky_callZeroOrNil(Sky_os_SEEK_CUR) }

func Os_SEEK_END(_ ...any) any { return sky_callZeroOrNil(Sky_os_SEEK_END) }

func Os_SEEK_SET(_ ...any) any { return sky_callZeroOrNil(Sky_os_SEEK_SET) }

func Os_FileChdir(a0 any) any { return Sky_os_FileChdir(a0) }

func Os_FileChmod(a0 any, a1 any) any { return Sky_os_FileChmod(a0, a1) }

func Os_FileChown(a0 any, a1 any, a2 any) any { return Sky_os_FileChown(a0, a1, a2) }

func Os_FileClose(a0 any) any { return Sky_os_FileClose(a0) }

func Os_FileFd(a0 any) any { return Sky_os_FileFd(a0) }

func Os_FileName(a0 any) any { return Sky_os_FileName(a0) }

func Os_FileRead(a0 any, a1 any) any { return Sky_os_FileRead(a0, a1) }

func Os_FileReadAt(a0 any, a1 any, a2 any) any { return Sky_os_FileReadAt(a0, a1, a2) }

func Os_FileReadDir(a0 any, a1 any) any { return Sky_os_FileReadDir(a0, a1) }

func Os_FileReadFrom(a0 any, a1 any) any { return Sky_os_FileReadFrom(a0, a1) }

func Os_FileReaddir(a0 any, a1 any) any { return Sky_os_FileReaddir(a0, a1) }

func Os_FileReaddirnames(a0 any, a1 any) any { return Sky_os_FileReaddirnames(a0, a1) }

func Os_FileSeek(a0 any, a1 any, a2 any) any { return Sky_os_FileSeek(a0, a1, a2) }

func Os_FileSetDeadline(a0 any, a1 any) any { return Sky_os_FileSetDeadline(a0, a1) }

func Os_FileSetReadDeadline(a0 any, a1 any) any { return Sky_os_FileSetReadDeadline(a0, a1) }

func Os_FileSetWriteDeadline(a0 any, a1 any) any { return Sky_os_FileSetWriteDeadline(a0, a1) }

func Os_FileStat(a0 any) any { return Sky_os_FileStat(a0) }

func Os_FileSync(a0 any) any { return Sky_os_FileSync(a0) }

func Os_FileSyscallConn(a0 any) any { return Sky_os_FileSyscallConn(a0) }

func Os_FileTruncate(a0 any, a1 any) any { return Sky_os_FileTruncate(a0, a1) }

func Os_FileWrite(a0 any, a1 any) any { return Sky_os_FileWrite(a0, a1) }

func Os_FileWriteAt(a0 any, a1 any, a2 any) any { return Sky_os_FileWriteAt(a0, a1, a2) }

func Os_FileWriteTo(a0 any, a1 any) any { return Sky_os_FileWriteTo(a0, a1) }

func Os_LinkErrorError(a0 any) any { return Sky_os_LinkErrorError(a0) }

func Os_LinkErrorUnwrap(a0 any) any { return Sky_os_LinkErrorUnwrap(a0) }

func Os_LinkErrorOp(a0 any) any { return Sky_os_LinkErrorOp(a0) }

func Os_LinkErrorOld(a0 any) any { return Sky_os_LinkErrorOld(a0) }

func Os_LinkErrorNew(a0 any) any { return Sky_os_LinkErrorNew(a0) }

func Os_LinkErrorErr(a0 any) any { return Sky_os_LinkErrorErr(a0) }

func Os_ProcAttrDir(a0 any) any { return Sky_os_ProcAttrDir(a0) }

func Os_ProcAttrEnv(a0 any) any { return Sky_os_ProcAttrEnv(a0) }

func Os_ProcAttrFiles(a0 any) any { return Sky_os_ProcAttrFiles(a0) }

func Os_ProcAttrSys(a0 any) any { return Sky_os_ProcAttrSys(a0) }

func Os_ProcessKill(a0 any) any { return Sky_os_ProcessKill(a0) }

func Os_ProcessRelease(a0 any) any { return Sky_os_ProcessRelease(a0) }

func Os_ProcessSignal(a0 any, a1 any) any { return Sky_os_ProcessSignal(a0, a1) }

func Os_ProcessWait(a0 any) any { return Sky_os_ProcessWait(a0) }

func Os_ProcessWithHandle(a0 any, a1 any) any { return Sky_os_ProcessWithHandle(a0, a1) }

func Os_ProcessPid(a0 any) any { return Sky_os_ProcessPid(a0) }

func Os_ProcessStateExitCode(a0 any) any { return Sky_os_ProcessStateExitCode(a0) }

func Os_ProcessStateExited(a0 any) any { return Sky_os_ProcessStateExited(a0) }

func Os_ProcessStatePid(a0 any) any { return Sky_os_ProcessStatePid(a0) }

func Os_ProcessStateString(a0 any) any { return Sky_os_ProcessStateString(a0) }

func Os_ProcessStateSuccess(a0 any) any { return Sky_os_ProcessStateSuccess(a0) }

func Os_ProcessStateSys(a0 any) any { return Sky_os_ProcessStateSys(a0) }

func Os_ProcessStateSysUsage(a0 any) any { return Sky_os_ProcessStateSysUsage(a0) }

func Os_ProcessStateSystemTime(a0 any) any { return Sky_os_ProcessStateSystemTime(a0) }

func Os_ProcessStateUserTime(a0 any) any { return Sky_os_ProcessStateUserTime(a0) }

func Os_RootChmod(a0 any, a1 any, a2 any) any { return Sky_os_RootChmod(a0, a1, a2) }

func Os_RootChown(a0 any, a1 any, a2 any, a3 any) any { return Sky_os_RootChown(a0, a1, a2, a3) }

func Os_RootChtimes(a0 any, a1 any, a2 any, a3 any) any { return Sky_os_RootChtimes(a0, a1, a2, a3) }

func Os_RootClose(a0 any) any { return Sky_os_RootClose(a0) }

func Os_RootCreate(a0 any, a1 any) any { return Sky_os_RootCreate(a0, a1) }

func Os_RootFS(a0 any) any { return Sky_os_RootFS(a0) }

func Os_RootLchown(a0 any, a1 any, a2 any, a3 any) any { return Sky_os_RootLchown(a0, a1, a2, a3) }

func Os_RootLink(a0 any, a1 any, a2 any) any { return Sky_os_RootLink(a0, a1, a2) }

func Os_RootLstat(a0 any, a1 any) any { return Sky_os_RootLstat(a0, a1) }

func Os_RootMkdir(a0 any, a1 any, a2 any) any { return Sky_os_RootMkdir(a0, a1, a2) }

func Os_RootMkdirAll(a0 any, a1 any, a2 any) any { return Sky_os_RootMkdirAll(a0, a1, a2) }

func Os_RootName(a0 any) any { return Sky_os_RootName(a0) }

func Os_RootOpen(a0 any, a1 any) any { return Sky_os_RootOpen(a0, a1) }

func Os_RootOpenFile(a0 any, a1 any, a2 any, a3 any) any { return Sky_os_RootOpenFile(a0, a1, a2, a3) }

func Os_RootOpenRoot(a0 any, a1 any) any { return Sky_os_RootOpenRoot(a0, a1) }

func Os_RootReadFile(a0 any, a1 any) any { return Sky_os_RootReadFile(a0, a1) }

func Os_RootReadlink(a0 any, a1 any) any { return Sky_os_RootReadlink(a0, a1) }

func Os_RootRemove(a0 any, a1 any) any { return Sky_os_RootRemove(a0, a1) }

func Os_RootRemoveAll(a0 any, a1 any) any { return Sky_os_RootRemoveAll(a0, a1) }

func Os_RootRename(a0 any, a1 any, a2 any) any { return Sky_os_RootRename(a0, a1, a2) }

func Os_RootStat(a0 any, a1 any) any { return Sky_os_RootStat(a0, a1) }

func Os_RootSymlink(a0 any, a1 any, a2 any) any { return Sky_os_RootSymlink(a0, a1, a2) }

func Os_RootWriteFile(a0 any, a1 any, a2 any, a3 any) any {
	return Sky_os_RootWriteFile(a0, a1, a2, a3)
}

func Os_SignalSignal(a0 any) any { return Sky_os_SignalSignal(a0) }

func Os_SignalString(a0 any) any { return Sky_os_SignalString(a0) }

func Os_SyscallErrorError(a0 any) any { return Sky_os_SyscallErrorError(a0) }

func Os_SyscallErrorTimeout(a0 any) any { return Sky_os_SyscallErrorTimeout(a0) }

func Os_SyscallErrorUnwrap(a0 any) any { return Sky_os_SyscallErrorUnwrap(a0) }

func Os_SyscallErrorSyscall(a0 any) any { return Sky_os_SyscallErrorSyscall(a0) }

func Os_SyscallErrorErr(a0 any) any { return Sky_os_SyscallErrorErr(a0) }

func Os_Exec_Command(a0 any, a1 any) any { return Sky_os_exec_Command(a0, a1) }

func Os_Exec_CommandContext(a0 any, a1 any, a2 any) any {
	return Sky_os_exec_CommandContext(a0, a1, a2)
}

func Os_Exec_LookPath(a0 any) any { return Sky_os_exec_LookPath(a0) }

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

func Os_Exec_NewCmd(a0 any) any { return Sky_os_exec_NEW_Cmd(a0) }

func Os_Exec_CmdPath(a0 any) any { return Sky_os_exec_FIELD_Cmd_Path(a0) }

func Os_Exec_CmdSetPath(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_Path(a0, a1) }
}

func Os_Exec_CmdArgs(a0 any) any { return Sky_os_exec_FIELD_Cmd_Args(a0) }

func Os_Exec_CmdSetArgs(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_Args(a0, a1) }
}

func Os_Exec_CmdEnv(a0 any) any { return Sky_os_exec_FIELD_Cmd_Env(a0) }

func Os_Exec_CmdSetEnv(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_Env(a0, a1) }
}

func Os_Exec_CmdDir(a0 any) any { return Sky_os_exec_FIELD_Cmd_Dir(a0) }

func Os_Exec_CmdSetDir(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_Dir(a0, a1) }
}

func Os_Exec_CmdStdin(a0 any) any { return Sky_os_exec_FIELD_Cmd_Stdin(a0) }

func Os_Exec_CmdSetStdin(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_Stdin(a0, a1) }
}

func Os_Exec_CmdStdout(a0 any) any { return Sky_os_exec_FIELD_Cmd_Stdout(a0) }

func Os_Exec_CmdSetStdout(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_Stdout(a0, a1) }
}

func Os_Exec_CmdStderr(a0 any) any { return Sky_os_exec_FIELD_Cmd_Stderr(a0) }

func Os_Exec_CmdSetStderr(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_Stderr(a0, a1) }
}

func Os_Exec_CmdExtraFiles(a0 any) any { return Sky_os_exec_FIELD_Cmd_ExtraFiles(a0) }

func Os_Exec_CmdSetExtraFiles(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_ExtraFiles(a0, a1) }
}

func Os_Exec_CmdSysProcAttr(a0 any) any { return Sky_os_exec_FIELD_Cmd_SysProcAttr(a0) }

func Os_Exec_CmdSetSysProcAttr(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_SysProcAttr(a0, a1) }
}

func Os_Exec_CmdProcess(a0 any) any { return Sky_os_exec_FIELD_Cmd_Process(a0) }

func Os_Exec_CmdSetProcess(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_Process(a0, a1) }
}

func Os_Exec_CmdProcessState(a0 any) any { return Sky_os_exec_FIELD_Cmd_ProcessState(a0) }

func Os_Exec_CmdSetProcessState(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_ProcessState(a0, a1) }
}

func Os_Exec_CmdErr(a0 any) any { return Sky_os_exec_FIELD_Cmd_Err(a0) }

func Os_Exec_CmdSetErr(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_Err(a0, a1) }
}

func Os_Exec_CmdWaitDelay(a0 any) any { return Sky_os_exec_FIELD_Cmd_WaitDelay(a0) }

func Os_Exec_CmdSetWaitDelay(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Cmd_WaitDelay(a0, a1) }
}

func Os_Exec_NewError(a0 any) any { return Sky_os_exec_NEW_Error(a0) }

func Os_Exec_ErrorName(a0 any) any { return Sky_os_exec_FIELD_Error_Name(a0) }

func Os_Exec_ErrorSetName(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Error_Name(a0, a1) }
}

func Os_Exec_ErrorErr(a0 any) any { return Sky_os_exec_FIELD_Error_Err(a0) }

func Os_Exec_ErrorSetErr(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_Error_Err(a0, a1) }
}

func Os_Exec_NewExitError(a0 any) any { return Sky_os_exec_NEW_ExitError(a0) }

func Os_Exec_ExitErrorProcessState(a0 any) any { return Sky_os_exec_FIELD_ExitError_ProcessState(a0) }

func Os_Exec_ExitErrorSetProcessState(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_ExitError_ProcessState(a0, a1) }
}

func Os_Exec_ExitErrorStderr(a0 any) any { return Sky_os_exec_FIELD_ExitError_Stderr(a0) }

func Os_Exec_ExitErrorSetStderr(a0 any) any {
	return func(a1 any) any { return Sky_os_exec_SET_ExitError_Stderr(a0, a1) }
}

func Os_Exec_ErrDot(_ ...any) any { return sky_callZeroOrNil(Sky_os_exec_ErrDot) }

func Os_Exec_ErrNotFound(_ ...any) any { return sky_callZeroOrNil(Sky_os_exec_ErrNotFound) }

func Os_Exec_ErrWaitDelay(_ ...any) any { return sky_callZeroOrNil(Sky_os_exec_ErrWaitDelay) }

func Os_Exec_CmdCancel(a0 any) any { return Sky_os_exec_CmdCancel(a0) }

func Os_Exec_ExitErrorSystemTime(a0 any) any { return Sky_os_exec_ExitErrorSystemTime(a0) }

func Os_Exec_ExitErrorUserTime(a0 any) any { return Sky_os_exec_ExitErrorUserTime(a0) }

func Bufio_NewReadWriter(a0 any, a1 any) any { return Sky_bufio_NewReadWriter(a0, a1) }

func Bufio_NewReader(a0 any) any { return Sky_bufio_NewReader(a0) }

func Bufio_NewReaderSize(a0 any, a1 any) any { return Sky_bufio_NewReaderSize(a0, a1) }

func Bufio_NewScanner(a0 any) any { return Sky_bufio_NewScanner(a0) }

func Bufio_NewWriter(a0 any) any { return Sky_bufio_NewWriter(a0) }

func Bufio_NewWriterSize(a0 any, a1 any) any { return Sky_bufio_NewWriterSize(a0, a1) }

func Bufio_ScanBytes(a0 any, a1 any) any { return Sky_bufio_ScanBytes(a0, a1) }

func Bufio_ScanLines(a0 any, a1 any) any { return Sky_bufio_ScanLines(a0, a1) }

func Bufio_ScanRunes(a0 any, a1 any) any { return Sky_bufio_ScanRunes(a0, a1) }

func Bufio_ScanWords(a0 any, a1 any) any { return Sky_bufio_ScanWords(a0, a1) }

func Bufio_ReadWriterAvailable(a0 any) any { return Sky_bufio_ReadWriterAvailable(a0) }

func Bufio_ReadWriterAvailableBuffer(a0 any) any { return Sky_bufio_ReadWriterAvailableBuffer(a0) }

func Bufio_ReadWriterDiscard(a0 any, a1 any) any { return Sky_bufio_ReadWriterDiscard(a0, a1) }

func Bufio_ReadWriterFlush(a0 any) any { return Sky_bufio_ReadWriterFlush(a0) }

func Bufio_ReadWriterPeek(a0 any, a1 any) any { return Sky_bufio_ReadWriterPeek(a0, a1) }

func Bufio_ReadWriterRead(a0 any, a1 any) any { return Sky_bufio_ReadWriterRead(a0, a1) }

func Bufio_ReadWriterReadByte(a0 any) any { return Sky_bufio_ReadWriterReadByte(a0) }

func Bufio_ReadWriterReadBytes(a0 any, a1 any) any { return Sky_bufio_ReadWriterReadBytes(a0, a1) }

func Bufio_ReadWriterReadFrom(a0 any, a1 any) any { return Sky_bufio_ReadWriterReadFrom(a0, a1) }

func Bufio_ReadWriterReadLine(a0 any) any { return Sky_bufio_ReadWriterReadLine(a0) }

func Bufio_ReadWriterReadRune(a0 any) any { return Sky_bufio_ReadWriterReadRune(a0) }

func Bufio_ReadWriterReadSlice(a0 any, a1 any) any { return Sky_bufio_ReadWriterReadSlice(a0, a1) }

func Bufio_ReadWriterReadString(a0 any, a1 any) any { return Sky_bufio_ReadWriterReadString(a0, a1) }

func Bufio_ReadWriterUnreadByte(a0 any) any { return Sky_bufio_ReadWriterUnreadByte(a0) }

func Bufio_ReadWriterUnreadRune(a0 any) any { return Sky_bufio_ReadWriterUnreadRune(a0) }

func Bufio_ReadWriterWrite(a0 any, a1 any) any { return Sky_bufio_ReadWriterWrite(a0, a1) }

func Bufio_ReadWriterWriteByte(a0 any, a1 any) any { return Sky_bufio_ReadWriterWriteByte(a0, a1) }

func Bufio_ReadWriterWriteRune(a0 any, a1 any) any { return Sky_bufio_ReadWriterWriteRune(a0, a1) }

func Bufio_ReadWriterWriteString(a0 any, a1 any) any { return Sky_bufio_ReadWriterWriteString(a0, a1) }

func Bufio_ReadWriterWriteTo(a0 any, a1 any) any { return Sky_bufio_ReadWriterWriteTo(a0, a1) }

func Bufio_ReaderBuffered(a0 any) any { return Sky_bufio_ReaderBuffered(a0) }

func Bufio_ReaderDiscard(a0 any, a1 any) any { return Sky_bufio_ReaderDiscard(a0, a1) }

func Bufio_ReaderPeek(a0 any, a1 any) any { return Sky_bufio_ReaderPeek(a0, a1) }

func Bufio_ReaderRead(a0 any, a1 any) any { return Sky_bufio_ReaderRead(a0, a1) }

func Bufio_ReaderReadByte(a0 any) any { return Sky_bufio_ReaderReadByte(a0) }

func Bufio_ReaderReadBytes(a0 any, a1 any) any { return Sky_bufio_ReaderReadBytes(a0, a1) }

func Bufio_ReaderReadLine(a0 any) any { return Sky_bufio_ReaderReadLine(a0) }

func Bufio_ReaderReadRune(a0 any) any { return Sky_bufio_ReaderReadRune(a0) }

func Bufio_ReaderReadSlice(a0 any, a1 any) any { return Sky_bufio_ReaderReadSlice(a0, a1) }

func Bufio_ReaderReadString(a0 any, a1 any) any { return Sky_bufio_ReaderReadString(a0, a1) }

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

func Bufio_WriterWriteByte(a0 any, a1 any) any { return Sky_bufio_WriterWriteByte(a0, a1) }

func Bufio_WriterWriteRune(a0 any, a1 any) any { return Sky_bufio_WriterWriteRune(a0, a1) }

func Bufio_WriterWriteString(a0 any, a1 any) any { return Sky_bufio_WriterWriteString(a0, a1) }

func Bufio_ReadWriterReader(a0 any) any { return Sky_bufio_FIELD_ReadWriter_Reader(a0) }

func Bufio_ReadWriterSetReader(a0 any) any {
	return func(a1 any) any { return Sky_bufio_SET_ReadWriter_Reader(a0, a1) }
}

func Bufio_ReadWriterWriter(a0 any) any { return Sky_bufio_FIELD_ReadWriter_Writer(a0) }

func Bufio_ReadWriterSetWriter(a0 any) any {
	return func(a1 any) any { return Sky_bufio_SET_ReadWriter_Writer(a0, a1) }
}

func Bufio_ErrAdvanceTooFar(_ ...any) any { return sky_callZeroOrNil(Sky_bufio_ErrAdvanceTooFar) }

func Bufio_ErrBadReadCount(_ ...any) any { return sky_callZeroOrNil(Sky_bufio_ErrBadReadCount) }

func Bufio_ErrBufferFull(_ ...any) any { return sky_callZeroOrNil(Sky_bufio_ErrBufferFull) }

func Bufio_ErrFinalToken(_ ...any) any { return sky_callZeroOrNil(Sky_bufio_ErrFinalToken) }

func Bufio_ErrInvalidUnreadByte(_ ...any) any {
	return sky_callZeroOrNil(Sky_bufio_ErrInvalidUnreadByte)
}

func Bufio_ErrInvalidUnreadRune(_ ...any) any {
	return sky_callZeroOrNil(Sky_bufio_ErrInvalidUnreadRune)
}

func Bufio_ErrNegativeAdvance(_ ...any) any { return sky_callZeroOrNil(Sky_bufio_ErrNegativeAdvance) }

func Bufio_ErrNegativeCount(_ ...any) any { return sky_callZeroOrNil(Sky_bufio_ErrNegativeCount) }

func Bufio_ErrTooLong(_ ...any) any { return sky_callZeroOrNil(Sky_bufio_ErrTooLong) }

func Bufio_MaxScanTokenSize(_ ...any) any { return sky_callZeroOrNil(Sky_bufio_MaxScanTokenSize) }

func Regexp_Compile(a0 any) any { return Sky_regexp_Compile(a0) }

func Regexp_CompilePOSIX(a0 any) any { return Sky_regexp_CompilePOSIX(a0) }

func Regexp_Match(a0 any, a1 any) any { return Sky_regexp_Match(a0, a1) }

func Regexp_MatchString(a0 any, a1 any) any { return Sky_regexp_MatchString(a0, a1) }

func Regexp_MustCompile(a0 any) any { return Sky_regexp_MustCompile(a0) }

func Regexp_MustCompilePOSIX(a0 any) any { return Sky_regexp_MustCompilePOSIX(a0) }

func Regexp_QuoteMeta(a0 any) any { return Sky_regexp_QuoteMeta(a0) }

func Regexp_RegexpAppendText(a0 any, a1 any) any { return Sky_regexp_RegexpAppendText(a0, a1) }

func Regexp_RegexpCopy(a0 any) any { return Sky_regexp_RegexpCopy(a0) }

func Regexp_RegexpExpand(a0 any, a1 any, a2 any, a3 any, a4 any) any {
	return Sky_regexp_RegexpExpand(a0, a1, a2, a3, a4)
}

func Regexp_RegexpExpandString(a0 any, a1 any, a2 any, a3 any, a4 any) any {
	return Sky_regexp_RegexpExpandString(a0, a1, a2, a3, a4)
}

func Regexp_RegexpFind(a0 any, a1 any) any { return Sky_regexp_RegexpFind(a0, a1) }

func Regexp_RegexpFindAll(a0 any, a1 any, a2 any) any { return Sky_regexp_RegexpFindAll(a0, a1, a2) }

func Regexp_RegexpFindAllString(a0 any, a1 any, a2 any) any {
	return Sky_regexp_RegexpFindAllString(a0, a1, a2)
}

func Regexp_RegexpFindAllSubmatch(a0 any, a1 any, a2 any) any {
	return Sky_regexp_RegexpFindAllSubmatch(a0, a1, a2)
}

func Regexp_RegexpFindIndex(a0 any, a1 any) any { return Sky_regexp_RegexpFindIndex(a0, a1) }

func Regexp_RegexpFindString(a0 any, a1 any) any { return Sky_regexp_RegexpFindString(a0, a1) }

func Regexp_RegexpFindStringIndex(a0 any, a1 any) any {
	return Sky_regexp_RegexpFindStringIndex(a0, a1)
}

func Regexp_RegexpFindStringSubmatch(a0 any, a1 any) any {
	return Sky_regexp_RegexpFindStringSubmatch(a0, a1)
}

func Regexp_RegexpFindStringSubmatchIndex(a0 any, a1 any) any {
	return Sky_regexp_RegexpFindStringSubmatchIndex(a0, a1)
}

func Regexp_RegexpFindSubmatch(a0 any, a1 any) any { return Sky_regexp_RegexpFindSubmatch(a0, a1) }

func Regexp_RegexpFindSubmatchIndex(a0 any, a1 any) any {
	return Sky_regexp_RegexpFindSubmatchIndex(a0, a1)
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

func Regexp_NewRegexp(a0 any) any { return Sky_regexp_NEW_Regexp(a0) }

func SetLevelFilter(v0 any) any {
	return SkyADT{Tag: 1, SkyName: "SetLevelFilter", Fields: []any{v0}}
}

func SetScopeFilter(v0 any) any {
	return SkyADT{Tag: 2, SkyName: "SetScopeFilter", Fields: []any{v0}}
}

func SetSearchFilter(v0 any) any {
	return SkyADT{Tag: 3, SkyName: "SetSearchFilter", Fields: []any{v0}}
}

func SetSourceFilter(v0 any) any {
	return SkyADT{Tag: 4, SkyName: "SetSourceFilter", Fields: []any{v0}}
}

func Navigate(v0 any) any {
	return SkyADT{Tag: 8, SkyName: "Navigate", Fields: []any{v0}}
}

// sky:type helpText : String

func helpText() any {
	return "sky-log — real-time log viewer\n\nUsage:\n  sky-log [options] [file ...]\n  command | sky-log\n\nOptions:\n  -h, --help    Show this help message\n\nExamples:\n  sky-log app.log                  Watch a log file\n  sky-log app.log server.log       Watch multiple files\n  tail -f /var/log/syslog | sky-log  Pipe stdin\n  sky-log                          Use sources from sky-log.toml\n"
}

// sky:type readFileEntries : any -> List LogEntry

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

// sky:type collectFileEntries : List any -> Dict any Int -> List any -> { entries : List t175 , newEntries : List elem , counts : Dict t183 Int }

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

// sky:type scanLine : any -> Maybe LogEntry

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

// sky:type collectStreamEntries : any -> any

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

// sky:type resolveMode : List any -> any -> { scanners : List elem , mode : SourceMode , watched : List { path : t252 , label : t252 } }

func resolveMode(args any, sources any) any {
	return func() any {
		if sky_asBool(sky_not(sky_listIsEmpty(args))) {
			return map[string]any{"watched": sky_call(sky_listMap(func(f any) any { return map[string]any{"path": f, "label": f} }), args), "scanners": []any{}, "mode": SkyADT{Tag: 0, SkyName: "FileMode"}}
		}
		return func() any {
			cmdScanners := sky_call(sky_listFilterMap(initCommandScanner), sources)
			_ = cmdScanners
			return func() any {
				if sky_asBool(sky_not(sky_listIsEmpty(cmdScanners))) {
					return map[string]any{"watched": []any{}, "scanners": cmdScanners, "mode": SkyADT{Tag: 1, SkyName: "StreamMode"}}
				}
				return map[string]any{"watched": []any{}, "scanners": sky_call(sky_listFilterMap(sky_identity), []any{initStdinScanner()}), "mode": SkyADT{Tag: 1, SkyName: "StreamMode"}}
			}()
		}()
	}()
}

// sky:type init : any -> ( { sourceFilter : String , theme : String , watched : List { label : String , path : String } , scopeFilter : String , scanners : List t323 , searchFilter : String , sourceMode : SourceMode , webhookRules : List WebhookRule , levelFilter : String , autoScroll : Bool , fileCounts : Dict t326 Int , entries : List t319 } , any )

func init_(_ any) any {
	return func() any {
		rawArgs := sky_call(sky_listDrop(1), sky_processGetArgs(struct{}{}))
		_ = rawArgs
		func() any {
			if sky_asBool(sky_call(sky_listAny(func(a any) any { return sky_asBool(sky_equal(a, "--help")) || sky_asBool(sky_equal(a, "-h")) }), rawArgs)) {
				return func() any { sky_println(helpText()); return sky_processExit(0) }()
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

// sky:type update : Msg -> any -> ( any , any )

func update(msg any, model any) any {
	return func() any {
		return func() any {
			__subject := msg
			__sky_tag := sky_adtTag(__subject)
			__sky_name := sky_getSkyName(__subject)
			_ = __sky_name
			if __sky_name == "Tick" || (__sky_name == "" && __sky_tag == 0) {
				return func() any {
					return func() any {
						__subject := sky_asMap(model)["sourceMode"]
						__sky_tag := sky_adtTag(__subject)
						__sky_name := sky_getSkyName(__subject)
						_ = __sky_name
						if __sky_name == "FileMode" || (__sky_name == "" && __sky_tag == 0) {
							return func() any {
								result := collectFileEntries(sky_asMap(model)["watched"], sky_asMap(model)["fileCounts"], sky_asMap(model)["entries"])
								_ = result
								Log_Webhook_ProcessNewEntries(sky_asMap(model)["webhookRules"], sky_asMap(result)["newEntries"])
								return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"entries": sky_asMap(result)["entries"], "fileCounts": sky_asMap(result)["counts"]}), V1: sky_cmdNone()}
							}()
						}
						if __sky_name == "StreamMode" || (__sky_name == "" && __sky_tag == 1) {
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
			if __sky_name == "SetLevelFilter" || (__sky_name == "" && __sky_tag == 1) {
				value := sky_adtField(__subject, 0)
				_ = value
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"levelFilter": value}), V1: sky_cmdNone()}
			}
			if __sky_name == "SetScopeFilter" || (__sky_name == "" && __sky_tag == 2) {
				value := sky_adtField(__subject, 0)
				_ = value
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"scopeFilter": value}), V1: sky_cmdNone()}
			}
			if __sky_name == "SetSearchFilter" || (__sky_name == "" && __sky_tag == 3) {
				value := sky_adtField(__subject, 0)
				_ = value
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"searchFilter": value}), V1: sky_cmdNone()}
			}
			if __sky_name == "SetSourceFilter" || (__sky_name == "" && __sky_tag == 4) {
				value := sky_adtField(__subject, 0)
				_ = value
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"sourceFilter": value}), V1: sky_cmdNone()}
			}
			if __sky_name == "ToggleAutoScroll" || (__sky_name == "" && __sky_tag == 5) {
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"autoScroll": sky_not(sky_asMap(model)["autoScroll"])}), V1: sky_cmdNone()}
			}
			if __sky_name == "ToggleTheme" || (__sky_name == "" && __sky_tag == 6) {
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
			if __sky_name == "ClearLogs" || (__sky_name == "" && __sky_tag == 7) {
				return SkyTuple2{V0: sky_recordUpdate(model, map[string]any{"entries": []any{}}), V1: sky_cmdNone()}
			}
			if __sky_name == "Navigate" || (__sky_name == "" && __sky_tag == 8) {
				return SkyTuple2{V0: model, V1: sky_cmdNone()}
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

// sky:type subscriptions : any -> any

func subscriptions(_ any) any {
	return sky_call(sky_timeEvery(200), SkyADT{Tag: 0, SkyName: "Tick"})
}

// sky:type filterEntry : String -> String -> String -> String -> any -> Bool

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
						return sky_call(sky_resultWithDefault(false), Regexp_RegexpMatchString(re, sky_asMap(entry)["message"]))
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

// sky:type levelClass : Level -> String

func levelClass(level any) any {
	return func() any {
		return func() any {
			__subject := level
			__sky_tag := sky_adtTag(__subject)
			__sky_name := sky_getSkyName(__subject)
			_ = __sky_name
			if __sky_name == "Debug" || (__sky_name == "" && __sky_tag == 0) {
				return "log-entry l-debug"
			}
			if __sky_name == "Info" || (__sky_name == "" && __sky_tag == 1) {
				return "log-entry l-info"
			}
			if __sky_name == "Warn" || (__sky_name == "" && __sky_tag == 2) {
				return "log-entry l-warn"
			}
			if __sky_name == "ErrorLevel" || (__sky_name == "" && __sky_tag == 3) {
				return "log-entry l-error"
			}
			panic("non-exhaustive case expression")
		}()
	}()
}

// sky:type padRight : Int -> String -> String -> String

func padRight(targetLen any, pad any, str any) any {
	return func() any {
		if sky_asBool(sky_numCompare(">=", sky_stringLength(str), targetLen)) {
			return str
		}
		return padRight(targetLen, pad, sky_concat(str, pad))
	}()
}

// sky:type formatTimestamp : Int -> String

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

// sky:type padZero : Int -> String

func padZero(n any) any {
	return func() any {
		if sky_asBool(sky_numCompare("<", n, 10)) {
			return sky_concat("0", sky_stringFromInt(n))
		}
		return sky_stringFromInt(n)
	}()
}

// sky:type viewEntry : any -> any

func viewEntry(entry any) any {
	return sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), levelClass(sky_asMap(entry)["level"]))}), []any{sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "ts")}), []any{sky_htmlText(formatTimestamp(sky_asMap(entry)["timestamp"]))}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "level")}), []any{sky_htmlText(sky_stringToUpper(padRight(5, " ", Log_Entry_LevelToString(sky_asMap(entry)["level"]))))}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "source")}), []any{sky_htmlText(sky_asMap(entry)["source"])}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "scope")}), []any{sky_htmlText(sky_asMap(entry)["scope"])}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "msg")}), []any{sky_htmlText(sky_asMap(entry)["message"])})})
}

// sky:type viewEntries : List any -> any

func viewEntries(entries any) any {
	return func() any {
		if sky_asBool(sky_listIsEmpty(entries)) {
			return sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "empty")}), []any{sky_htmlText("Waiting for logs...")})
		}
		return sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("id"), "log-entries")}), sky_call(sky_listMap(viewEntry), sky_call(sky_listTake(5000), entries)))
	}()
}

// sky:type statusText : any -> String

func statusText(model any) any {
	return func() any {
		return func() any {
			__subject := sky_asMap(model)["sourceMode"]
			__sky_tag := sky_adtTag(__subject)
			__sky_name := sky_getSkyName(__subject)
			_ = __sky_name
			if __sky_name == "FileMode" || (__sky_name == "" && __sky_tag == 0) {
				return func() any {
					labels := sky_call(sky_listMap(func(w any) any { return sky_asMap(w)["label"] }), sky_asMap(model)["watched"])
					_ = labels
					return sky_concat("Watching ", sky_concat(sky_stringFromInt(sky_listLength(sky_asMap(model)["watched"])), sky_concat(" source(s): ", sky_call(sky_stringJoin(", "), labels))))
				}()
			}
			if __sky_name == "StreamMode" || (__sky_name == "" && __sky_tag == 1) {
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

// sky:type styles : any

func styles() any {
	return sky_call(sky_htmlStyleNode([]any{}), sky_cssStylesheet([]any{sky_call(sky_cssRule("*"), []any{sky_call(sky_cssPropFn("margin"), sky_call(sky_cssVal("0"), struct{}{})), sky_call(sky_cssPropFn("padding"), sky_call(sky_cssVal("0"), struct{}{})), sky_call(sky_cssProp("box-sizing"), "border-box")}), sky_call(sky_cssRule("html, body"), []any{sky_call(sky_cssPropFn("height"), sky_cssPct(100)), sky_call(sky_cssPropFn("overflow"), "hidden")}), sky_call(sky_cssRule(".root.dark"), []any{sky_call(sky_cssProp("--bg"), "#0a0e14"), sky_call(sky_cssProp("--fg"), "#c5cdd8"), sky_call(sky_cssProp("--muted"), "#64748b"), sky_call(sky_cssProp("--dimmed"), "#475569"), sky_call(sky_cssProp("--input-bg"), "#0a0e14"), sky_call(sky_cssProp("--input-border"), "rgba(255,255,255,0.12)"), sky_call(sky_cssProp("--border"), "rgba(255,255,255,0.06)"), sky_call(sky_cssProp("--hover"), "rgba(255,255,255,0.02)"), sky_call(sky_cssProp("--scroll-thumb"), "rgba(255,255,255,0.08)"), sky_call(sky_cssProp("--scroll-hover"), "rgba(255,255,255,0.15)"), sky_call(sky_cssProp("--btn-border"), "rgba(255,255,255,0.08)"), sky_call(sky_cssProp("--btn-hover-fg"), "#c5cdd8"), sky_call(sky_cssProp("--btn-hover-border"), "rgba(255,255,255,0.15)")}), sky_call(sky_cssRule(".root.light"), []any{sky_call(sky_cssProp("--bg"), "#ffffff"), sky_call(sky_cssProp("--fg"), "#1e293b"), sky_call(sky_cssProp("--muted"), "#94a3b8"), sky_call(sky_cssProp("--dimmed"), "#94a3b8"), sky_call(sky_cssProp("--input-bg"), "#f8fafc"), sky_call(sky_cssProp("--input-border"), "#d1d5db"), sky_call(sky_cssProp("--border"), "#e5e7eb"), sky_call(sky_cssProp("--hover"), "rgba(0,0,0,0.02)"), sky_call(sky_cssProp("--scroll-thumb"), "rgba(0,0,0,0.12)"), sky_call(sky_cssProp("--scroll-hover"), "rgba(0,0,0,0.2)"), sky_call(sky_cssProp("--btn-border"), "#d1d5db"), sky_call(sky_cssProp("--btn-hover-fg"), "#1e293b"), sky_call(sky_cssProp("--btn-hover-border"), "#9ca3af")}), sky_call(sky_cssRule("body"), []any{sky_call(sky_cssPropFn("background-color"), sky_cssHex("#0a0e14")), sky_call(sky_cssPropFn("color"), sky_cssHex("#c5cdd8")), sky_call(sky_cssPropFn("font-family"), "'SF Mono', 'JetBrains Mono', ui-monospace, Menlo, monospace"), sky_call(sky_cssPropFn("font-size"), sky_cssPx(11)), sky_call(sky_cssProp("-webkit-font-smoothing"), "antialiased")}), sky_call(sky_cssRule(".root"), []any{sky_call(sky_cssPropFn("position"), "fixed"), sky_call(sky_cssPropFn("top"), sky_call(sky_cssVal("0"), struct{}{})), sky_call(sky_cssProp("left"), "0"), sky_call(sky_cssProp("right"), "0"), sky_call(sky_cssPropFn("bottom"), sky_call(sky_cssVal("0"), struct{}{})), sky_call(sky_cssPropFn("display"), "flex"), sky_call(sky_cssPropFn("flex-direction"), "column"), sky_call(sky_cssPropFn("overflow"), "hidden"), sky_call(sky_cssPropFn("background-color"), "var(--bg, #0a0e14)"), sky_call(sky_cssPropFn("color"), "var(--fg, #c5cdd8)")}), sky_call(sky_cssRule(".toolbar"), []any{sky_call(sky_cssPropFn("display"), "flex"), sky_call(sky_cssPropFn("align-items"), "center"), sky_call(sky_cssPropFn("gap"), sky_cssPx(8)), sky_call(sky_cssPadding2(sky_cssPx(6)), sky_cssPx(12)), sky_call(sky_cssProp("border-bottom"), "1px solid var(--border)"), sky_call(sky_cssProp("flex-shrink"), "0")}), sky_call(sky_cssRule(".toolbar h1"), []any{sky_call(sky_cssPropFn("font-size"), sky_cssPx(11)), sky_call(sky_cssPropFn("font-weight"), "500"), sky_call(sky_cssPropFn("color"), "var(--muted)"), sky_call(sky_cssProp("margin-right"), "4px")}), sky_call(sky_cssRule(".toolbar h1 span"), []any{sky_call(sky_cssPropFn("color"), "var(--fg)"), sky_call(sky_cssPropFn("font-weight"), "600")}), sky_call(sky_cssRule(".toolbar select, .toolbar input"), []any{sky_call(sky_cssPropFn("background-color"), "var(--input-bg)"), sky_call(sky_cssPropFn("color"), "var(--fg)"), sky_call(sky_cssProp("border"), "1px solid var(--input-border)"), sky_call(sky_cssPropFn("border-radius"), sky_cssPx(4)), sky_call(sky_cssPadding2(sky_cssPx(3)), sky_cssPx(6)), sky_call(sky_cssPropFn("font-family"), "inherit"), sky_call(sky_cssPropFn("font-size"), sky_cssPx(11)), sky_call(sky_cssProp("outline"), "none"), sky_call(sky_cssPropFn("transition"), "border-color 0.15s")}), sky_call(sky_cssRule(".toolbar select:focus, .toolbar input:focus"), []any{sky_call(sky_cssPropFn("border-color"), sky_cssHex("#60a5fa"))}), sky_call(sky_cssRule(".toolbar input"), []any{sky_call(sky_cssPropFn("width"), sky_cssPx(120))}), sky_call(sky_cssRule(".spacer"), []any{sky_call(sky_cssPropFn("flex"), "1")}), sky_call(sky_cssRule(".stats"), []any{sky_call(sky_cssPropFn("color"), "var(--dimmed)"), sky_call(sky_cssPropFn("font-size"), sky_cssPx(11))}), sky_call(sky_cssRule(".btn"), []any{sky_call(sky_cssPropFn("background-color"), "transparent"), sky_call(sky_cssPropFn("color"), "var(--muted)"), sky_call(sky_cssProp("border"), "1px solid var(--btn-border)"), sky_call(sky_cssPropFn("border-radius"), sky_cssPx(4)), sky_call(sky_cssPadding2(sky_cssPx(3)), sky_cssPx(8)), sky_call(sky_cssPropFn("font-family"), "inherit"), sky_call(sky_cssPropFn("font-size"), sky_cssPx(11)), sky_call(sky_cssPropFn("cursor"), "pointer"), sky_call(sky_cssPropFn("transition"), "all 0.15s"), sky_call(sky_cssProp("user-select"), "none")}), sky_call(sky_cssRule(".btn:hover"), []any{sky_call(sky_cssPropFn("color"), "var(--btn-hover-fg)"), sky_call(sky_cssProp("border-color"), "var(--btn-hover-border)")}), sky_call(sky_cssRule(".btn.active"), []any{sky_call(sky_cssPropFn("color"), sky_cssHex("#60a5fa")), sky_call(sky_cssPropFn("border-color"), "rgba(96,165,250,0.3)")}), sky_call(sky_cssRule(".btn.danger:hover"), []any{sky_call(sky_cssPropFn("color"), sky_cssHex("#f87171")), sky_call(sky_cssPropFn("border-color"), "rgba(248,113,113,0.3)")}), sky_call(sky_cssRule("#log-entries"), []any{sky_call(sky_cssPropFn("flex"), "1"), sky_call(sky_cssProp("min-height"), "0"), sky_call(sky_cssProp("overflow-y"), "scroll")}), sky_call(sky_cssRule("#log-entries::-webkit-scrollbar"), []any{sky_call(sky_cssPropFn("width"), sky_cssPx(6))}), sky_call(sky_cssRule("#log-entries::-webkit-scrollbar-track"), []any{sky_call(sky_cssPropFn("background-color"), "transparent")}), sky_call(sky_cssRule("#log-entries::-webkit-scrollbar-thumb"), []any{sky_call(sky_cssProp("background-color"), "var(--scroll-thumb)"), sky_call(sky_cssPropFn("border-radius"), sky_cssPx(3))}), sky_call(sky_cssRule("#log-entries::-webkit-scrollbar-thumb:hover"), []any{sky_call(sky_cssProp("background-color"), "var(--scroll-hover)")}), sky_call(sky_cssRule(".log-entry"), []any{sky_call(sky_cssPropFn("display"), "flex"), sky_call(sky_cssPropFn("align-items"), "baseline"), sky_call(sky_cssPadding2(sky_cssPx(0)), sky_cssPx(12)), sky_call(sky_cssProp("line-height"), "18px"), sky_call(sky_cssPropFn("gap"), sky_cssPx(8)), sky_call(sky_cssPropFn("border-left"), "2px solid transparent")}), sky_call(sky_cssRule(".log-entry:hover"), []any{sky_call(sky_cssPropFn("background-color"), "var(--hover)")}), sky_call(sky_cssRule(".l-debug .level"), []any{sky_call(sky_cssPropFn("color"), sky_cssHex("#64748b"))}), sky_call(sky_cssRule(".l-info .level"), []any{sky_call(sky_cssPropFn("color"), sky_cssHex("#60a5fa"))}), sky_call(sky_cssRule(".l-warn .level"), []any{sky_call(sky_cssPropFn("color"), sky_cssHex("#fbbf24"))}), sky_call(sky_cssRule(".l-error .level"), []any{sky_call(sky_cssPropFn("color"), sky_cssHex("#f87171"))}), sky_call(sky_cssRule(".log-entry.l-error"), []any{sky_call(sky_cssPropFn("border-left"), "2px solid rgba(248,113,113,0.5)"), sky_call(sky_cssPropFn("background-color"), "rgba(248,113,113,0.03)")}), sky_call(sky_cssRule(".log-entry.l-warn"), []any{sky_call(sky_cssPropFn("border-left"), "2px solid rgba(251,191,36,0.4)"), sky_call(sky_cssPropFn("background-color"), "rgba(251,191,36,0.02)")}), sky_call(sky_cssRule(".ts"), []any{sky_call(sky_cssPropFn("color"), "var(--dimmed)"), sky_call(sky_cssProp("white-space"), "nowrap"), sky_call(sky_cssPropFn("min-width"), sky_cssPx(56)), sky_call(sky_cssProp("font-variant-numeric"), "tabular-nums")}), sky_call(sky_cssRule(".level"), []any{sky_call(sky_cssPropFn("font-weight"), "600"), sky_call(sky_cssProp("white-space"), "nowrap"), sky_call(sky_cssPropFn("min-width"), sky_cssPx(40)), sky_call(sky_cssPropFn("text-transform"), "uppercase"), sky_call(sky_cssPropFn("letter-spacing"), "0.3px")}), sky_call(sky_cssRule(".source"), []any{sky_call(sky_cssProp("white-space"), "nowrap"), sky_call(sky_cssPropFn("max-width"), sky_cssPx(100)), sky_call(sky_cssPropFn("overflow"), "hidden"), sky_call(sky_cssProp("text-overflow"), "ellipsis"), sky_call(sky_cssPropFn("color"), sky_cssHex("#a78bfa")), sky_call(sky_cssPropFn("background-color"), "rgba(167,139,250,0.08)"), sky_call(sky_cssPadding2(sky_cssPx(0)), sky_cssPx(4)), sky_call(sky_cssPropFn("border-radius"), sky_cssPx(3))}), sky_call(sky_cssRule(".scope"), []any{sky_call(sky_cssPropFn("color"), sky_cssHex("#34d399")), sky_call(sky_cssProp("white-space"), "nowrap"), sky_call(sky_cssPropFn("max-width"), sky_cssPx(120)), sky_call(sky_cssPropFn("overflow"), "hidden"), sky_call(sky_cssProp("text-overflow"), "ellipsis")}), sky_call(sky_cssRule(".msg"), []any{sky_call(sky_cssPropFn("color"), "var(--fg)"), sky_call(sky_cssProp("word-break"), "break-word"), sky_call(sky_cssPropFn("flex"), "1")}), sky_call(sky_cssRule(".empty"), []any{sky_call(sky_cssPropFn("display"), "flex"), sky_call(sky_cssPropFn("align-items"), "center"), sky_call(sky_cssPropFn("justify-content"), "center"), sky_call(sky_cssPropFn("height"), sky_cssPct(100)), sky_call(sky_cssPropFn("color"), "var(--dimmed)")}), sky_call(sky_cssRule(".status"), []any{sky_call(sky_cssPropFn("display"), "flex"), sky_call(sky_cssPropFn("align-items"), "center"), sky_call(sky_cssPropFn("gap"), sky_cssPx(6)), sky_call(sky_cssPadding2(sky_cssPx(4)), sky_cssPx(12)), sky_call(sky_cssProp("border-top"), "1px solid var(--border)"), sky_call(sky_cssPropFn("font-size"), sky_cssPx(10)), sky_call(sky_cssPropFn("color"), "var(--dimmed)"), sky_call(sky_cssProp("flex-shrink"), "0")}), sky_call(sky_cssRule(".dot"), []any{sky_call(sky_cssPropFn("width"), sky_cssPx(5)), sky_call(sky_cssPropFn("height"), sky_cssPx(5)), sky_call(sky_cssPropFn("border-radius"), sky_cssPct(50)), sky_call(sky_cssPropFn("background-color"), sky_cssHex("#34d399")), sky_call(sky_cssProp("box-shadow"), "0 0 4px rgba(52,211,153,0.4)")})}))
}

// sky:type uniqueSources : List any -> List String

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

// sky:type view : any -> any

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
		return sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), sky_concat("root ", sky_asMap(model)["theme"]))}), []any{sky_cssStyles, sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "toolbar")}), []any{sky_call(sky_call(sky_htmlEl("h1"), []any{}), []any{sky_htmlText("sky"), sky_call(sky_call(sky_htmlEl("span"), []any{}), []any{sky_htmlText("-log")})}), sky_call(sky_call(sky_htmlEl("button"), []any{sky_call(sky_attrSimple("class"), "btn"), sky_call(sky_evtHandler("click"), SkyADT{Tag: 6, SkyName: "ToggleTheme"})}), []any{sky_htmlText(themeIcon)}), sky_call(sky_call(sky_htmlEl("select"), []any{sky_call(sky_attrSimple("id"), "level-filter"), sky_call(sky_evtHandler("change"), SetLevelFilter)}), []any{sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "")}), []any{sky_htmlText("Level")}), sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "debug")}), []any{sky_htmlText("DEBUG")}), sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "info")}), []any{sky_htmlText("INFO")}), sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "warn")}), []any{sky_htmlText("WARN")}), sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "error")}), []any{sky_htmlText("ERROR")})}), sky_call(sky_htmlVoid("input"), []any{sky_call(sky_attrSimple("id"), "scope-filter"), sky_call(sky_attrSimple("type"), "text"), sky_call(sky_attrSimple("placeholder"), "scope"), sky_call(sky_evtHandler("input"), SetScopeFilter)}), sky_call(sky_htmlVoid("input"), []any{sky_call(sky_attrSimple("id"), "search-filter"), sky_call(sky_attrSimple("type"), "text"), sky_call(sky_attrSimple("placeholder"), "regex search"), sky_call(sky_evtHandler("input"), SetSearchFilter)}), sky_call(sky_call(sky_htmlEl("select"), []any{sky_call(sky_attrSimple("id"), "source-filter"), sky_call(sky_evtHandler("change"), SetSourceFilter)}), sky_call(sky_listAppend([]any{sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), "")}), []any{sky_htmlText("Source")})}), sky_call(sky_listMap(func(s any) any {
			return sky_call(sky_call(sky_htmlEl("option"), []any{sky_call(sky_attrSimple("value"), s)}), []any{sky_htmlText(s)})
		}), sources))), sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "spacer")}), []any{}), sky_call(sky_call(sky_htmlEl("button"), []any{sky_call(sky_attrSimple("class"), scrollClass), sky_call(sky_evtHandler("click"), SkyADT{Tag: 5, SkyName: "ToggleAutoScroll"})}), []any{sky_htmlText(scrollLabel)}), sky_call(sky_call(sky_htmlEl("button"), []any{sky_call(sky_attrSimple("class"), "btn danger"), sky_call(sky_evtHandler("click"), SkyADT{Tag: 7, SkyName: "ClearLogs"})}), []any{sky_htmlText("Clear")})}), viewEntries(filtered), sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "status")}), []any{sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "dot")}), []any{}), sky_htmlText(statusText(model)), sky_call(sky_call(sky_htmlEl("div"), []any{sky_call(sky_attrSimple("class"), "spacer")}), []any{}), sky_call(sky_call(sky_htmlEl("span"), []any{sky_call(sky_attrSimple("class"), "stats")}), []any{sky_htmlText(sky_concat(sky_stringFromInt(filteredCount), sky_concat("/", sky_stringFromInt(totalCount))))})})})
	}()
}

// sky:type main : any

func main() {
	sky_runMainTask(sky_liveApp(map[string]any{"init": init_, "update": update, "view": view, "subscriptions": subscriptions, "routes": []any{sky_call(sky_liveRoute("/"), SkyADT{Tag: 0, SkyName: "LogPage"})}, "notFound": SkyADT{Tag: 0, SkyName: "LogPage"}}))
}
