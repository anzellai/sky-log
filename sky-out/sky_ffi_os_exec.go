package main

import (
	_ffi_fmt "fmt"
	"io"
	_ffi_pkg "os/exec"
	exec "os/exec"
	_ffi_reflect "reflect"
	"time"
)

var _ = _ffi_fmt.Sprintf
var _ = _ffi_reflect.TypeOf

func Sky_os_exec_Command(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asStringSlice(arg1)
		_val := _ffi_pkg.Command(_arg0, _arg1...)
		return SkyOk(_val)
	}()
}

func Sky_os_exec_CommandContext(arg0 any, arg1 any, arg2 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asContext(arg0)
		_arg1 := sky_asString(arg1)
		_arg2 := sky_asStringSlice(arg2)
		_val := _ffi_pkg.CommandContext(_arg0, _arg1, _arg2...)
		return SkyOk(_val)
	}()
}

func Sky_os_exec_LookPath(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.LookPath(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_exec_CmdCombinedOutput(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_val, _err := _receiver.CombinedOutput()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_exec_CmdEnviron(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_val := _receiver.Environ()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_CmdOutput(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_val, _err := _receiver.Output()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_exec_CmdRun(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_err := _receiver.Run()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_exec_CmdStart(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_err := _receiver.Start()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_exec_CmdStderrPipe(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_val, _err := _receiver.StderrPipe()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_exec_CmdStdinPipe(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_val, _err := _receiver.StdinPipe()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_exec_CmdStdoutPipe(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_val, _err := _receiver.StdoutPipe()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_os_exec_CmdString(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_val := _receiver.String()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_CmdWait(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Cmd {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Cmd)
		}()
		_err := _receiver.Wait()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_exec_ErrorError(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Error {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Error)
		}()
		_val := _receiver.Error()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_ErrorUnwrap(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Error {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Error)
		}()
		_err := _receiver.Unwrap()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_os_exec_ExitErrorError(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ExitError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ExitError)
		}()
		_val := _receiver.Error()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_ExitErrorExitCode(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ExitError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ExitError)
		}()
		_val := _receiver.ExitCode()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_ExitErrorExited(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ExitError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ExitError)
		}()
		_val := _receiver.Exited()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_ExitErrorPid(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ExitError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ExitError)
		}()
		_val := _receiver.Pid()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_ExitErrorString(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ExitError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ExitError)
		}()
		_val := _receiver.String()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_ExitErrorSuccess(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ExitError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ExitError)
		}()
		_val := _receiver.Success()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_ExitErrorSys(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ExitError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ExitError)
		}()
		_val := _receiver.Sys()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_ExitErrorSysUsage(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ExitError {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ExitError)
		}()
		_val := _receiver.SysUsage()
		return SkyOk(_val)
	}()
}

func Sky_os_exec_NEW_Cmd(_ any) any {
	return &_ffi_pkg.Cmd{}
}

func Sky_os_exec_FIELD_Cmd_Path(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Path")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_Path(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Path")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	f.Set(_ffi_reflect.ValueOf(sky_asString(val)))
	return receiver
}

func Sky_os_exec_FIELD_Cmd_Args(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Args")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_Args(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Args")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	f.Set(_ffi_reflect.ValueOf(sky_asStringSlice(val)))
	return receiver
}

func Sky_os_exec_FIELD_Cmd_Env(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Env")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_Env(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Env")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	f.Set(_ffi_reflect.ValueOf(sky_asStringSlice(val)))
	return receiver
}

func Sky_os_exec_FIELD_Cmd_Dir(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Dir")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_Dir(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Dir")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	f.Set(_ffi_reflect.ValueOf(sky_asString(val)))
	return receiver
}

func Sky_os_exec_FIELD_Cmd_Stdin(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Stdin")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_Stdin(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Stdin")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val).Convert(f.Type()))
	}
	return receiver
}

func Sky_os_exec_FIELD_Cmd_Stdout(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Stdout")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_Stdout(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Stdout")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val).Convert(f.Type()))
	}
	return receiver
}

func Sky_os_exec_FIELD_Cmd_Stderr(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Stderr")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_Stderr(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Stderr")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val).Convert(f.Type()))
	}
	return receiver
}

func Sky_os_exec_FIELD_Cmd_ExtraFiles(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("ExtraFiles")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_ExtraFiles(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("ExtraFiles")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	items := sky_asList(val)
	slice := _ffi_reflect.MakeSlice(f.Type(), len(items), len(items))
	for i, item := range items {
		if item != nil {
			slice.Index(i).Set(_ffi_reflect.ValueOf(item).Elem())
		}
	}
	f.Set(slice)
	return receiver
}

func Sky_os_exec_FIELD_Cmd_SysProcAttr(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("SysProcAttr")
		if f.IsValid() {
			if f.IsNil() {
				return SkyNothing()
			}
			return SkyJust(f.Interface())
		}
	}
	return SkyNothing()
}

func Sky_os_exec_SET_Cmd_SysProcAttr(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("SysProcAttr")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val))
	}
	return receiver
}

func Sky_os_exec_FIELD_Cmd_Process(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Process")
		if f.IsValid() {
			if f.IsNil() {
				return SkyNothing()
			}
			return SkyJust(f.Interface())
		}
	}
	return SkyNothing()
}

func Sky_os_exec_SET_Cmd_Process(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Process")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val))
	}
	return receiver
}

func Sky_os_exec_FIELD_Cmd_ProcessState(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("ProcessState")
		if f.IsValid() {
			if f.IsNil() {
				return SkyNothing()
			}
			return SkyJust(f.Interface())
		}
	}
	return SkyNothing()
}

func Sky_os_exec_SET_Cmd_ProcessState(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("ProcessState")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val))
	}
	return receiver
}

func Sky_os_exec_FIELD_Cmd_Err(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Err")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_Err(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Err")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val).Convert(f.Type()))
	}
	return receiver
}

func Sky_os_exec_FIELD_Cmd_WaitDelay(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("WaitDelay")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Cmd_WaitDelay(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("WaitDelay")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val).Convert(f.Type()))
	}
	return receiver
}

func Sky_os_exec_NEW_Error(_ any) any {
	return &_ffi_pkg.Error{}
}

func Sky_os_exec_FIELD_Error_Name(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Name")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Error_Name(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Name")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	f.Set(_ffi_reflect.ValueOf(sky_asString(val)))
	return receiver
}

func Sky_os_exec_FIELD_Error_Err(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Err")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_Error_Err(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Err")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val).Convert(f.Type()))
	}
	return receiver
}

func Sky_os_exec_NEW_ExitError(_ any) any {
	return &_ffi_pkg.ExitError{}
}

func Sky_os_exec_FIELD_ExitError_ProcessState(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("ProcessState")
		if f.IsValid() {
			if f.IsNil() {
				return SkyNothing()
			}
			return SkyJust(f.Interface())
		}
	}
	return SkyNothing()
}

func Sky_os_exec_SET_ExitError_ProcessState(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("ProcessState")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val))
	}
	return receiver
}

func Sky_os_exec_FIELD_ExitError_Stderr(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Stderr")
		if f.IsValid() {
			return f.Interface()
		}
	}
	return nil
}

func Sky_os_exec_SET_ExitError_Stderr(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Stderr")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	items := sky_asList(val)
	slice := _ffi_reflect.MakeSlice(f.Type(), len(items), len(items))
	for i, item := range items {
		if item != nil {
			slice.Index(i).Set(_ffi_reflect.ValueOf(item).Elem())
		}
	}
	f.Set(slice)
	return receiver
}

func Sky_os_exec_ErrDot() any {
	return exec.ErrDot
}

func Sky_os_exec_ErrNotFound() any {
	return exec.ErrNotFound
}

func Sky_os_exec_ErrWaitDelay() any {
	return exec.ErrWaitDelay
}

func Sky_os_exec_CmdEnv(this any) any {
	_this := this.(*exec.Cmd)

	_val := _this.Env
	_result := make([]any, len(_val))
	for _i, _v := range _val {
		_result[_i] = _v
	}
	return _result
}

func Sky_os_exec_CmdStdin(this any) io.Reader {
	_this := this.(*exec.Cmd)

	return _this.Stdin
}

func Sky_os_exec_CmdStdout(this any) io.Writer {
	_this := this.(*exec.Cmd)

	return _this.Stdout
}

func Sky_os_exec_CmdStderr(this any) io.Writer {
	_this := this.(*exec.Cmd)

	return _this.Stderr
}

func Sky_os_exec_CmdCancel(this any) func() error {
	_this := this.(*exec.Cmd)

	return _this.Cancel
}

func Sky_os_exec_ErrorErr(this any) error {
	_this := this.(*exec.Error)

	return _this.Err
}

func Sky_os_exec_ExitErrorSystemTime(this any) time.Duration {
	_this := this.(*exec.ExitError)

	return _this.SystemTime()
}

func Sky_os_exec_ExitErrorUserTime(this any) time.Duration {
	_this := this.(*exec.ExitError)

	return _this.UserTime()
}
