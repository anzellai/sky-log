package main

import (
	_ffi_fmt "fmt"
	io "io"
	"io/fs"
	"os"
	_ffi_pkg "os"
	_ffi_reflect "reflect"
	"syscall"
	"time"
)

var _ = _ffi_fmt.Sprintf
var _ = _ffi_reflect.TypeOf

func Sky_os_Chdir(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_err := _ffi_pkg.Chdir(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_Chmod(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := func() _ffi_pkg.FileMode {
			if v, ok := arg1.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_err := _ffi_pkg.Chmod(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_Chown(arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asInt(arg1)
		_arg2 := sky_asInt(arg2)
		_err := _ffi_pkg.Chown(_arg0, _arg1, _arg2)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_Clearenv() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_ffi_pkg.Clearenv()
		return SkyOk(struct{}{})
	}()
}

func Sky_os_Create(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.Create(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_CreateTemp(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_val, _err := _ffi_pkg.CreateTemp(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Environ() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val := _ffi_pkg.Environ()
		return SkyOk(_val)
	}()
}

func Sky_os_Executable() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val, _err := _ffi_pkg.Executable()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Exit(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asInt(arg0)
		_ffi_pkg.Exit(_arg0)
		return SkyOk(struct{}{})
	}()
}

func Sky_os_ExpandEnv(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val := _ffi_pkg.ExpandEnv(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_os_FindProcess(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asInt(arg0)
		_val, _err := _ffi_pkg.FindProcess(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Getegid() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val := _ffi_pkg.Getegid()
		return SkyOk(_val)
	}()
}

func Sky_os_Getenv(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val := _ffi_pkg.Getenv(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_os_Geteuid() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val := _ffi_pkg.Geteuid()
		return SkyOk(_val)
	}()
}

func Sky_os_Getgid() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val := _ffi_pkg.Getgid()
		return SkyOk(_val)
	}()
}

func Sky_os_Getgroups() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val, _err := _ffi_pkg.Getgroups()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Getpagesize() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val := _ffi_pkg.Getpagesize()
		return SkyOk(_val)
	}()
}

func Sky_os_Getpid() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val := _ffi_pkg.Getpid()
		return SkyOk(_val)
	}()
}

func Sky_os_Getppid() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val := _ffi_pkg.Getppid()
		return SkyOk(_val)
	}()
}

func Sky_os_Getuid() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val := _ffi_pkg.Getuid()
		return SkyOk(_val)
	}()
}

func Sky_os_Getwd() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val, _err := _ffi_pkg.Getwd()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Hostname() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val, _err := _ffi_pkg.Hostname()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_IsExist(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asError(arg0)
		_val := _ffi_pkg.IsExist(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_os_IsNotExist(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asError(arg0)
		_val := _ffi_pkg.IsNotExist(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_os_IsPathSeparator(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := byte(sky_asInt(arg0))
		_val := _ffi_pkg.IsPathSeparator(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_os_IsPermission(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asError(arg0)
		_val := _ffi_pkg.IsPermission(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_os_IsTimeout(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asError(arg0)
		_val := _ffi_pkg.IsTimeout(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_os_Lchown(arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asInt(arg1)
		_arg2 := sky_asInt(arg2)
		_err := _ffi_pkg.Lchown(_arg0, _arg1, _arg2)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_Link(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_err := _ffi_pkg.Link(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_LookupEnv(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_r0, _r1 := _ffi_pkg.LookupEnv(_arg0)
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_os_Lstat(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.Lstat(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Mkdir(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := func() _ffi_pkg.FileMode {
			if v, ok := arg1.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_err := _ffi_pkg.Mkdir(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_MkdirAll(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := func() _ffi_pkg.FileMode {
			if v, ok := arg1.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_err := _ffi_pkg.MkdirAll(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_MkdirTemp(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_val, _err := _ffi_pkg.MkdirTemp(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_NewSyscallError(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asError(arg1)
		_err := _ffi_pkg.NewSyscallError(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_Open(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.Open(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_OpenFile(arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asInt(arg1)
		_arg2 := func() _ffi_pkg.FileMode {
			if v, ok := arg2.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_val, _err := _ffi_pkg.OpenFile(_arg0, _arg1, _arg2)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_OpenInRoot(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_val, _err := _ffi_pkg.OpenInRoot(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_OpenRoot(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.OpenRoot(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Pipe() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_r0, _r1, _err := _ffi_pkg.Pipe()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_os_ReadDir(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.ReadDir(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_ReadFile(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.ReadFile(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Readlink(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.Readlink(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Remove(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_err := _ffi_pkg.Remove(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RemoveAll(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_err := _ffi_pkg.RemoveAll(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_Rename(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_err := _ffi_pkg.Rename(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_SameFile(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := func() _ffi_pkg.FileInfo {
			if v, ok := arg0.(_ffi_pkg.FileInfo); ok {
				return v
			}
			var zero _ffi_pkg.FileInfo
			return zero
		}()
		_arg1 := func() _ffi_pkg.FileInfo {
			if v, ok := arg1.(_ffi_pkg.FileInfo); ok {
				return v
			}
			var zero _ffi_pkg.FileInfo
			return zero
		}()
		_val := _ffi_pkg.SameFile(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_os_Setenv(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_err := _ffi_pkg.Setenv(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_StartProcess(arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asStringSlice(arg1)
		_arg2 := func() *_ffi_pkg.ProcAttr {
			if arg2 == nil {
				return nil
			}
			return arg2.(*_ffi_pkg.ProcAttr)
		}()
		_val, _err := _ffi_pkg.StartProcess(_arg0, _arg1, _arg2)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Stat(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.Stat(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_Symlink(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_err := _ffi_pkg.Symlink(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_TempDir() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val := _ffi_pkg.TempDir()
		return SkyOk(_val)
	}()
}

func Sky_os_Truncate(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asInt64(arg1)
		_err := _ffi_pkg.Truncate(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_Unsetenv(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_err := _ffi_pkg.Unsetenv(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_UserCacheDir() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val, _err := _ffi_pkg.UserCacheDir()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_UserConfigDir() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val, _err := _ffi_pkg.UserConfigDir()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_UserHomeDir() any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_val, _err := _ffi_pkg.UserHomeDir()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_WriteFile(arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asBytes(arg1)
		_arg2 := func() _ffi_pkg.FileMode {
			if v, ok := arg2.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_err := _ffi_pkg.WriteFile(_arg0, _arg1, _arg2)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_FileChdir(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_err := _receiver.Chdir()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_FileChmod(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := func() _ffi_pkg.FileMode {
			if v, ok := arg0.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_err := _receiver.Chmod(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_FileChown(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asInt(arg0)
		_arg1 := sky_asInt(arg1)
		_err := _receiver.Chown(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_FileClose(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_err := _receiver.Close()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_FileName(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_val := _receiver.Name()
		return SkyOk(_val)
	}()
}

func Sky_os_FileRead(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asBytes(arg0)
		_val, _err := _receiver.Read(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileReadAt(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asInt64(arg1)
		_val, _err := _receiver.ReadAt(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileReadDir(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asInt(arg0)
		_val, _err := _receiver.ReadDir(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileReadFrom(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := arg0.(io.Reader)
		_val, _err := _receiver.ReadFrom(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileReaddir(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asInt(arg0)
		_val, _err := _receiver.Readdir(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileReaddirnames(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asInt(arg0)
		_val, _err := _receiver.Readdirnames(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileSeek(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asInt64(arg0)
		_arg1 := sky_asInt(arg1)
		_val, _err := _receiver.Seek(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileStat(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_val, _err := _receiver.Stat()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileSync(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_err := _receiver.Sync()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_FileTruncate(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asInt64(arg0)
		_err := _receiver.Truncate(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_FileWrite(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asBytes(arg0)
		_val, _err := _receiver.Write(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileWriteAt(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asInt64(arg1)
		_val, _err := _receiver.WriteAt(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileWriteString(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.WriteString(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_FileWriteTo(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.File {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.File)
		}()
		_arg0 := arg0.(io.Writer)
		_val, _err := _receiver.WriteTo(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_LinkErrorError(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.LinkError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.LinkError)
		}()
		_val := _receiver.Error()
		return SkyOk(_val)
	}()
}

func Sky_os_LinkErrorUnwrap(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.LinkError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.LinkError)
		}()
		_err := _receiver.Unwrap()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_ProcessKill(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Process {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Process)
		}()
		_err := _receiver.Kill()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_ProcessRelease(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Process {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Process)
		}()
		_err := _receiver.Release()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_ProcessSignal(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Process {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Process)
		}()
		_arg0 := func() _ffi_pkg.Signal {
			if v, ok := arg0.(_ffi_pkg.Signal); ok {
				return v
			}
			var zero _ffi_pkg.Signal
			return zero
		}()
		_err := _receiver.Signal(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_ProcessWait(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Process {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Process)
		}()
		_val, _err := _receiver.Wait()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_ProcessStateExitCode(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ProcessState {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ProcessState)
		}()
		_val := _receiver.ExitCode()
		return SkyOk(_val)
	}()
}

func Sky_os_ProcessStateExited(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ProcessState {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ProcessState)
		}()
		_val := _receiver.Exited()
		return SkyOk(_val)
	}()
}

func Sky_os_ProcessStatePid(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ProcessState {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ProcessState)
		}()
		_val := _receiver.Pid()
		return SkyOk(_val)
	}()
}

func Sky_os_ProcessStateString(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ProcessState {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ProcessState)
		}()
		_val := _receiver.String()
		return SkyOk(_val)
	}()
}

func Sky_os_ProcessStateSuccess(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ProcessState {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ProcessState)
		}()
		_val := _receiver.Success()
		return SkyOk(_val)
	}()
}

func Sky_os_ProcessStateSys(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ProcessState {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ProcessState)
		}()
		_val := _receiver.Sys()
		return SkyOk(_val)
	}()
}

func Sky_os_ProcessStateSysUsage(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ProcessState {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ProcessState)
		}()
		_val := _receiver.SysUsage()
		return SkyOk(_val)
	}()
}

func Sky_os_RootChmod(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := func() _ffi_pkg.FileMode {
			if v, ok := arg1.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_err := _receiver.Chmod(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootChown(receiver any, arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asInt(arg1)
		_arg2 := sky_asInt(arg2)
		_err := _receiver.Chown(_arg0, _arg1, _arg2)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootClose(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_err := _receiver.Close()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootCreate(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.Create(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_RootLchown(receiver any, arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asInt(arg1)
		_arg2 := sky_asInt(arg2)
		_err := _receiver.Lchown(_arg0, _arg1, _arg2)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootLink(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_err := _receiver.Link(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootLstat(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.Lstat(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_RootMkdir(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := func() _ffi_pkg.FileMode {
			if v, ok := arg1.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_err := _receiver.Mkdir(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootMkdirAll(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := func() _ffi_pkg.FileMode {
			if v, ok := arg1.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_err := _receiver.MkdirAll(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootName(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_val := _receiver.Name()
		return SkyOk(_val)
	}()
}

func Sky_os_RootOpen(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.Open(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_RootOpenFile(receiver any, arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asInt(arg1)
		_arg2 := func() _ffi_pkg.FileMode {
			if v, ok := arg2.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_val, _err := _receiver.OpenFile(_arg0, _arg1, _arg2)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_RootOpenRoot(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.OpenRoot(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_RootReadFile(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.ReadFile(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_RootReadlink(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.Readlink(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_RootRemove(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_err := _receiver.Remove(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootRemoveAll(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_err := _receiver.RemoveAll(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootRename(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_err := _receiver.Rename(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootStat(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.Stat(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_RootSymlink(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_err := _receiver.Symlink(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_RootWriteFile(receiver any, arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Root {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Root)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asBytes(arg1)
		_arg2 := func() _ffi_pkg.FileMode {
			if v, ok := arg2.(_ffi_pkg.FileMode); ok {
				return v
			}
			var zero _ffi_pkg.FileMode
			return zero
		}()
		_err := _receiver.WriteFile(_arg0, _arg1, _arg2)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_SignalSignal(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := receiver.(_ffi_pkg.Signal)
		_receiver.Signal()
		return SkyOk(struct{}{})
	}()
}

func Sky_os_SignalString(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := receiver.(_ffi_pkg.Signal)
		_val := _receiver.String()
		return SkyOk(_val)
	}()
}

func Sky_os_SyscallErrorError(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.SyscallError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.SyscallError)
		}()
		_val := _receiver.Error()
		return SkyOk(_val)
	}()
}

func Sky_os_SyscallErrorTimeout(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.SyscallError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.SyscallError)
		}()
		_val := _receiver.Timeout()
		return SkyOk(_val)
	}()
}

func Sky_os_SyscallErrorUnwrap(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.SyscallError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.SyscallError)
		}()
		_err := _receiver.Unwrap()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_Args(_ any) any { return _ffi_pkg.Args }

func Sky_os_Interrupt(_ any) any { return _ffi_pkg.Interrupt }

func Sky_os_Kill(_ any) any { return _ffi_pkg.Kill }

func Sky_os_Stderr(_ any) any { return _ffi_pkg.Stderr }

func Sky_os_Stdin(_ any) any { return _ffi_pkg.Stdin }

func Sky_os_Stdout(_ any) any { return _ffi_pkg.Stdout }

func Sky_os_O_APPEND(_ any) any { return _ffi_pkg.O_APPEND }

func Sky_os_O_CREATE(_ any) any { return _ffi_pkg.O_CREATE }

func Sky_os_O_EXCL(_ any) any { return _ffi_pkg.O_EXCL }

func Sky_os_O_RDONLY(_ any) any { return _ffi_pkg.O_RDONLY }

func Sky_os_O_RDWR(_ any) any { return _ffi_pkg.O_RDWR }

func Sky_os_O_SYNC(_ any) any { return _ffi_pkg.O_SYNC }

func Sky_os_O_TRUNC(_ any) any { return _ffi_pkg.O_TRUNC }

func Sky_os_O_WRONLY(_ any) any { return _ffi_pkg.O_WRONLY }

func Sky_os_SEEK_CUR(_ any) any { return _ffi_pkg.SEEK_CUR }

func Sky_os_SEEK_END(_ any) any { return _ffi_pkg.SEEK_END }

func Sky_os_SEEK_SET(_ any) any { return _ffi_pkg.SEEK_SET }

func Sky_os_Chtimes(arg0 any, arg1 any, arg2 any) SkyResult {
	_arg0 := arg0.(string)
	_arg1 := arg1.(time.Time)
	_arg2 := arg2.(time.Time)
	err := os.Chtimes(_arg0, _arg1, _arg2)
	if err != nil {
		return SkyErr(err)
	}
	return SkyOk(struct{}{})
}

func Sky_os_CopyFS(arg0 any, arg1 any) SkyResult {
	_arg0 := arg0.(string)
	_arg1 := arg1.(fs.FS)
	err := os.CopyFS(_arg0, _arg1)
	if err != nil {
		return SkyErr(err)
	}
	return SkyOk(struct{}{})
}

func Sky_os_DirFS(arg0 any) fs.FS {
	_arg0 := arg0.(string)
	return os.DirFS(_arg0)
}

func Sky_os_Expand(arg0 any, arg1 any) string {
	_arg0 := arg0.(string)
	_skyFn1 := arg1.(func(any) any)
	_arg1 := func(p0 string) string {
		return _skyFn1(p0).(string)
	}
	return os.Expand(_arg0, _arg1)
}

func Sky_os_NewFile(arg0 any, arg1 any) *os.File {
	_arg0 := arg0.(uintptr)
	_arg1 := arg1.(string)
	return os.NewFile(_arg0, _arg1)
}

func Sky_os_ErrClosed() any {
	return os.ErrClosed
}

func Sky_os_ErrDeadlineExceeded() any {
	return os.ErrDeadlineExceeded
}

func Sky_os_ErrExist() any {
	return os.ErrExist
}

func Sky_os_ErrInvalid() any {
	return os.ErrInvalid
}

func Sky_os_ErrNoDeadline() any {
	return os.ErrNoDeadline
}

func Sky_os_ErrNoHandle() any {
	return os.ErrNoHandle
}

func Sky_os_ErrNotExist() any {
	return os.ErrNotExist
}

func Sky_os_ErrPermission() any {
	return os.ErrPermission
}

func Sky_os_ErrProcessDone() any {
	return os.ErrProcessDone
}

func Sky_os_DevNull() any {
	return os.DevNull
}

func Sky_os_ModeAppend() any {
	return os.ModeAppend
}

func Sky_os_ModeCharDevice() any {
	return os.ModeCharDevice
}

func Sky_os_ModeDevice() any {
	return os.ModeDevice
}

func Sky_os_ModeDir() any {
	return os.ModeDir
}

func Sky_os_ModeExclusive() any {
	return os.ModeExclusive
}

func Sky_os_ModeIrregular() any {
	return os.ModeIrregular
}

func Sky_os_ModeNamedPipe() any {
	return os.ModeNamedPipe
}

func Sky_os_ModePerm() any {
	return os.ModePerm
}

func Sky_os_ModeSetgid() any {
	return os.ModeSetgid
}

func Sky_os_ModeSetuid() any {
	return os.ModeSetuid
}

func Sky_os_ModeSocket() any {
	return os.ModeSocket
}

func Sky_os_ModeSticky() any {
	return os.ModeSticky
}

func Sky_os_ModeSymlink() any {
	return os.ModeSymlink
}

func Sky_os_ModeTemporary() any {
	return os.ModeTemporary
}

func Sky_os_ModeType() any {
	return os.ModeType
}

func Sky_os_PathListSeparator() any {
	return os.PathListSeparator
}

func Sky_os_PathSeparator() any {
	return os.PathSeparator
}

func Sky_os_FileFd(this any) uintptr {
	var _this *os.File
	if _p, ok := this.(*os.File); ok {
		_this = _p
	} else {
		_v := this.(os.File)
		_this = &_v
	}

	return _this.Fd()
}

func Sky_os_FileSetDeadline(this any, arg0 any) SkyResult {
	var _this *os.File
	if _p, ok := this.(*os.File); ok {
		_this = _p
	} else {
		_v := this.(os.File)
		_this = &_v
	}
	_arg0 := arg0.(time.Time)
	err := _this.SetDeadline(_arg0)
	if err != nil {
		return SkyErr(err)
	}
	return SkyOk(struct{}{})
}

func Sky_os_FileSetReadDeadline(this any, arg0 any) SkyResult {
	var _this *os.File
	if _p, ok := this.(*os.File); ok {
		_this = _p
	} else {
		_v := this.(os.File)
		_this = &_v
	}
	_arg0 := arg0.(time.Time)
	err := _this.SetReadDeadline(_arg0)
	if err != nil {
		return SkyErr(err)
	}
	return SkyOk(struct{}{})
}

func Sky_os_FileSetWriteDeadline(this any, arg0 any) SkyResult {
	var _this *os.File
	if _p, ok := this.(*os.File); ok {
		_this = _p
	} else {
		_v := this.(os.File)
		_this = &_v
	}
	_arg0 := arg0.(time.Time)
	err := _this.SetWriteDeadline(_arg0)
	if err != nil {
		return SkyErr(err)
	}
	return SkyOk(struct{}{})
}

func Sky_os_FileSyscallConn(this any) SkyResult {
	var _this *os.File
	if _p, ok := this.(*os.File); ok {
		_this = _p
	} else {
		_v := this.(os.File)
		_this = &_v
	}

	res, err := _this.SyscallConn()
	if err != nil {
		return SkyErr(err)
	}
	return SkyOk(res)
}

func Sky_os_LinkErrorOp(this any) string {
	var _this *os.LinkError
	if _p, ok := this.(*os.LinkError); ok {
		_this = _p
	} else {
		_v := this.(os.LinkError)
		_this = &_v
	}

	return _this.Op
}

func Sky_os_LinkErrorOld(this any) string {
	var _this *os.LinkError
	if _p, ok := this.(*os.LinkError); ok {
		_this = _p
	} else {
		_v := this.(os.LinkError)
		_this = &_v
	}

	return _this.Old
}

func Sky_os_LinkErrorNew(this any) string {
	var _this *os.LinkError
	if _p, ok := this.(*os.LinkError); ok {
		_this = _p
	} else {
		_v := this.(os.LinkError)
		_this = &_v
	}

	return _this.New
}

func Sky_os_LinkErrorErr(this any) error {
	var _this *os.LinkError
	if _p, ok := this.(*os.LinkError); ok {
		_this = _p
	} else {
		_v := this.(os.LinkError)
		_this = &_v
	}

	return _this.Err
}

func Sky_os_ProcAttrDir(this any) string {
	var _this *os.ProcAttr
	if _p, ok := this.(*os.ProcAttr); ok {
		_this = _p
	} else {
		_v := this.(os.ProcAttr)
		_this = &_v
	}

	return _this.Dir
}

func Sky_os_ProcAttrEnv(this any) any {
	var _this *os.ProcAttr
	if _p, ok := this.(*os.ProcAttr); ok {
		_this = _p
	} else {
		_v := this.(os.ProcAttr)
		_this = &_v
	}

	_val := _this.Env
	_result := make([]any, len(_val))
	for _i, _v := range _val {
		_result[_i] = _v
	}
	return _result
}

func Sky_os_ProcAttrFiles(this any) any {
	var _this *os.ProcAttr
	if _p, ok := this.(*os.ProcAttr); ok {
		_this = _p
	} else {
		_v := this.(os.ProcAttr)
		_this = &_v
	}

	_val := _this.Files
	_result := make([]any, len(_val))
	for _i, _v := range _val {
		_result[_i] = _v
	}
	return _result
}

func Sky_os_ProcAttrSys(this any) *syscall.SysProcAttr {
	var _this *os.ProcAttr
	if _p, ok := this.(*os.ProcAttr); ok {
		_this = _p
	} else {
		_v := this.(os.ProcAttr)
		_this = &_v
	}

	return _this.Sys
}

func Sky_os_ProcessWithHandle(this any, arg0 any) SkyResult {
	var _this *os.Process
	if _p, ok := this.(*os.Process); ok {
		_this = _p
	} else {
		_v := this.(os.Process)
		_this = &_v
	}
	_skyFn0 := arg0.(func(any) any)
	_arg0 := func(p0 uintptr) {
		_skyFn0(p0)
	}
	err := _this.WithHandle(_arg0)
	if err != nil {
		return SkyErr(err)
	}
	return SkyOk(struct{}{})
}

func Sky_os_ProcessPid(this any) int {
	var _this *os.Process
	if _p, ok := this.(*os.Process); ok {
		_this = _p
	} else {
		_v := this.(os.Process)
		_this = &_v
	}

	return _this.Pid
}

func Sky_os_ProcessStateSystemTime(this any) time.Duration {
	var _this *os.ProcessState
	if _p, ok := this.(*os.ProcessState); ok {
		_this = _p
	} else {
		_v := this.(os.ProcessState)
		_this = &_v
	}

	return _this.SystemTime()
}

func Sky_os_ProcessStateUserTime(this any) time.Duration {
	var _this *os.ProcessState
	if _p, ok := this.(*os.ProcessState); ok {
		_this = _p
	} else {
		_v := this.(os.ProcessState)
		_this = &_v
	}

	return _this.UserTime()
}

func Sky_os_RootChtimes(this any, arg0 any, arg1 any, arg2 any) SkyResult {
	var _this *os.Root
	if _p, ok := this.(*os.Root); ok {
		_this = _p
	} else {
		_v := this.(os.Root)
		_this = &_v
	}
	_arg0 := arg0.(string)
	_arg1 := arg1.(time.Time)
	_arg2 := arg2.(time.Time)
	err := _this.Chtimes(_arg0, _arg1, _arg2)
	if err != nil {
		return SkyErr(err)
	}
	return SkyOk(struct{}{})
}

func Sky_os_RootFS(this any) fs.FS {
	var _this *os.Root
	if _p, ok := this.(*os.Root); ok {
		_this = _p
	} else {
		_v := this.(os.Root)
		_this = &_v
	}

	return _this.FS()
}

func Sky_os_SyscallErrorSyscall(this any) string {
	var _this *os.SyscallError
	if _p, ok := this.(*os.SyscallError); ok {
		_this = _p
	} else {
		_v := this.(os.SyscallError)
		_this = &_v
	}

	return _this.Syscall
}

func Sky_os_SyscallErrorErr(this any) error {
	var _this *os.SyscallError
	if _p, ok := this.(*os.SyscallError); ok {
		_this = _p
	} else {
		_v := this.(os.SyscallError)
		_this = &_v
	}

	return _this.Err
}
