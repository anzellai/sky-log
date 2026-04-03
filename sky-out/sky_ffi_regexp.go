package main

import (
	_ffi_fmt "fmt"
	_ffi_reflect "reflect"
	_ffi_pkg "regexp"
)

var _ = _ffi_fmt.Sprintf
var _ = _ffi_reflect.TypeOf

func Sky_regexp_Compile(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.Compile(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_regexp_CompilePOSIX(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _ffi_pkg.CompilePOSIX(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_regexp_Match(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asBytes(arg1)
		_val, _err := _ffi_pkg.Match(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_regexp_MatchString(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_val, _err := _ffi_pkg.MatchString(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_regexp_MustCompile(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val := _ffi_pkg.MustCompile(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_MustCompilePOSIX(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val := _ffi_pkg.MustCompilePOSIX(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_QuoteMeta(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asString(arg0)
		_val := _ffi_pkg.QuoteMeta(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpAppendText(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_val, _err := _receiver.AppendText(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpCopy(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_val := _receiver.Copy()
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpExpand(receiver any, arg0 any, arg1 any, arg2 any, arg3 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asBytes(arg1)
		_arg2 := sky_asBytes(arg2)
		_arg3 := func() []int {
			lst := sky_asList(arg3)
			out := make([]int, len(lst))
			for i, v := range lst {
				if cv, ok := v.(int); ok {
					out[i] = cv
				}
			}
			return out
		}()
		_val := _receiver.Expand(_arg0, _arg1, _arg2, _arg3)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpExpandString(receiver any, arg0 any, arg1 any, arg2 any, arg3 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asString(arg1)
		_arg2 := sky_asString(arg2)
		_arg3 := func() []int {
			lst := sky_asList(arg3)
			out := make([]int, len(lst))
			for i, v := range lst {
				if cv, ok := v.(int); ok {
					out[i] = cv
				}
			}
			return out
		}()
		_val := _receiver.ExpandString(_arg0, _arg1, _arg2, _arg3)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFind(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_val := _receiver.Find(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindAll(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asInt(arg1)
		_val := _receiver.FindAll(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindAllString(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asInt(arg1)
		_val := _receiver.FindAllString(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindAllSubmatch(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asInt(arg1)
		_val := _receiver.FindAllSubmatch(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindIndex(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_val := _receiver.FindIndex(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindString(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_val := _receiver.FindString(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindStringIndex(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_val := _receiver.FindStringIndex(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindStringSubmatch(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_val := _receiver.FindStringSubmatch(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindStringSubmatchIndex(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_val := _receiver.FindStringSubmatchIndex(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindSubmatch(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_val := _receiver.FindSubmatch(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpFindSubmatchIndex(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_val := _receiver.FindSubmatchIndex(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpLiteralPrefix(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_r0, _r1 := _receiver.LiteralPrefix()
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_regexp_RegexpLongest(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_receiver.Longest()
		return SkyOk(struct{}{})
	}()
}

func Sky_regexp_RegexpMarshalText(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_val, _err := _receiver.MarshalText()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpMatch(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_val := _receiver.Match(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpMatchString(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_val := _receiver.MatchString(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpNumSubexp(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_val := _receiver.NumSubexp()
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpReplaceAll(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asBytes(arg1)
		_val := _receiver.ReplaceAll(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpReplaceAllLiteral(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asBytes(arg1)
		_val := _receiver.ReplaceAllLiteral(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpReplaceAllLiteralString(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_val := _receiver.ReplaceAllLiteralString(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpReplaceAllString(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asString(arg1)
		_val := _receiver.ReplaceAllString(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpSplit(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_arg1 := sky_asInt(arg1)
		_val := _receiver.Split(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpString(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_val := _receiver.String()
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpSubexpIndex(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asString(arg0)
		_val := _receiver.SubexpIndex(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpSubexpNames(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_val := _receiver.SubexpNames()
		return SkyOk(_val)
	}()
}

func Sky_regexp_RegexpUnmarshalText(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Regexp {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Regexp)
		}()
		_arg0 := sky_asBytes(arg0)
		_err := _receiver.UnmarshalText(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_regexp_NEW_Regexp(_ any) any {
	return &_ffi_pkg.Regexp{}
}
