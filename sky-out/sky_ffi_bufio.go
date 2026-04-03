package main

import (
	"bufio"
	_ffi_pkg "bufio"
	_ffi_fmt "fmt"
	io "io"
	_ffi_reflect "reflect"
)

var _ = _ffi_fmt.Sprintf
var _ = _ffi_reflect.TypeOf

func Sky_bufio_NewReadWriter(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := func() *_ffi_pkg.Reader {
			if arg0 == nil {
				return nil
			}
			return arg0.(*_ffi_pkg.Reader)
		}()
		_arg1 := func() *_ffi_pkg.Writer {
			if arg1 == nil {
				return nil
			}
			return arg1.(*_ffi_pkg.Writer)
		}()
		_val := _ffi_pkg.NewReadWriter(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_bufio_NewReader(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := arg0.(io.Reader)
		_val := _ffi_pkg.NewReader(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_bufio_NewReaderSize(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := arg0.(io.Reader)
		_arg1 := sky_asInt(arg1)
		_val := _ffi_pkg.NewReaderSize(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_bufio_NewScanner(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := arg0.(io.Reader)
		_val := _ffi_pkg.NewScanner(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_bufio_NewWriter(arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := arg0.(io.Writer)
		_val := _ffi_pkg.NewWriter(_arg0)
		return SkyOk(_val)
	}()
}

func Sky_bufio_NewWriterSize(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := arg0.(io.Writer)
		_arg1 := sky_asInt(arg1)
		_val := _ffi_pkg.NewWriterSize(_arg0, _arg1)
		return SkyOk(_val)
	}()
}

func Sky_bufio_ScanBytes(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asBool(arg1)
		_r0, _r1, _err := _ffi_pkg.ScanBytes(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_bufio_ScanLines(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asBool(arg1)
		_r0, _r1, _err := _ffi_pkg.ScanLines(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_bufio_ScanRunes(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asBool(arg1)
		_r0, _r1, _err := _ffi_pkg.ScanRunes(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_bufio_ScanWords(arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asBool(arg1)
		_r0, _r1, _err := _ffi_pkg.ScanWords(_arg0, _arg1)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_bufio_ReadWriterAvailable(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_val := _receiver.Available()
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterAvailableBuffer(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_val := _receiver.AvailableBuffer()
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterDiscard(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := sky_asInt(arg0)
		_val, _err := _receiver.Discard(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterFlush(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_err := _receiver.Flush()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ReadWriterPeek(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := sky_asInt(arg0)
		_val, _err := _receiver.Peek(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterRead(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := sky_asBytes(arg0)
		_val, _err := _receiver.Read(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterReadByte(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_val, _err := _receiver.ReadByte()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterReadBytes(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := byte(sky_asInt(arg0))
		_val, _err := _receiver.ReadBytes(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterReadFrom(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := arg0.(io.Reader)
		_val, _err := _receiver.ReadFrom(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterReadLine(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_r0, _r1, _err := _receiver.ReadLine()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_bufio_ReadWriterReadRune(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_r0, _r1, _err := _receiver.ReadRune()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_bufio_ReadWriterReadSlice(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := byte(sky_asInt(arg0))
		_val, _err := _receiver.ReadSlice(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterReadString(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := byte(sky_asInt(arg0))
		_val, _err := _receiver.ReadString(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterUnreadByte(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_err := _receiver.UnreadByte()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ReadWriterUnreadRune(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_err := _receiver.UnreadRune()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ReadWriterWrite(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := sky_asBytes(arg0)
		_val, _err := _receiver.Write(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterWriteByte(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := byte(sky_asInt(arg0))
		_err := _receiver.WriteByte(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ReadWriterWriteRune(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := rune(sky_asInt(arg0))
		_val, _err := _receiver.WriteRune(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterWriteString(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.WriteString(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReadWriterWriteTo(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.ReadWriter {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.ReadWriter)
		}()
		_arg0 := arg0.(io.Writer)
		_val, _err := _receiver.WriteTo(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderBuffered(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_val := _receiver.Buffered()
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderDiscard(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_arg0 := sky_asInt(arg0)
		_val, _err := _receiver.Discard(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderPeek(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_arg0 := sky_asInt(arg0)
		_val, _err := _receiver.Peek(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderRead(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_arg0 := sky_asBytes(arg0)
		_val, _err := _receiver.Read(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderReadByte(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_val, _err := _receiver.ReadByte()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderReadBytes(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_arg0 := byte(sky_asInt(arg0))
		_val, _err := _receiver.ReadBytes(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderReadLine(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_r0, _r1, _err := _receiver.ReadLine()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_bufio_ReaderReadRune(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_r0, _r1, _err := _receiver.ReadRune()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(SkyTuple2{V0: _r0, V1: _r1})
	}()
}

func Sky_bufio_ReaderReadSlice(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_arg0 := byte(sky_asInt(arg0))
		_val, _err := _receiver.ReadSlice(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderReadString(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_arg0 := byte(sky_asInt(arg0))
		_val, _err := _receiver.ReadString(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderReset(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_arg0 := arg0.(io.Reader)
		_receiver.Reset(_arg0)
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ReaderSize(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_val := _receiver.Size()
		return SkyOk(_val)
	}()
}

func Sky_bufio_ReaderUnreadByte(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_err := _receiver.UnreadByte()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ReaderUnreadRune(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_err := _receiver.UnreadRune()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ReaderWriteTo(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Reader {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Reader)
		}()
		_arg0 := arg0.(io.Writer)
		_val, _err := _receiver.WriteTo(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_ScannerBuffer(receiver any, arg0 any, arg1 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Scanner {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Scanner)
		}()
		_arg0 := sky_asBytes(arg0)
		_arg1 := sky_asInt(arg1)
		_receiver.Buffer(_arg0, _arg1)
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ScannerBytes(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Scanner {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Scanner)
		}()
		_val := _receiver.Bytes()
		return SkyOk(_val)
	}()
}

func Sky_bufio_ScannerErr(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Scanner {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Scanner)
		}()
		_err := _receiver.Err()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ScannerScan(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Scanner {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Scanner)
		}()
		_val := _receiver.Scan()
		return SkyOk(_val)
	}()
}

func Sky_bufio_ScannerSplit(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Scanner {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Scanner)
		}()
		_arg0 := func() _ffi_pkg.SplitFunc {
			if v, ok := arg0.(_ffi_pkg.SplitFunc); ok {
				return v
			}
			var zero _ffi_pkg.SplitFunc
			return zero
		}()
		_receiver.Split(_arg0)
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_ScannerText(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Scanner {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Scanner)
		}()
		_val := _receiver.Text()
		return SkyOk(_val)
	}()
}

func Sky_bufio_WriterAvailable(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_val := _receiver.Available()
		return SkyOk(_val)
	}()
}

func Sky_bufio_WriterAvailableBuffer(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_val := _receiver.AvailableBuffer()
		return SkyOk(_val)
	}()
}

func Sky_bufio_WriterBuffered(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_val := _receiver.Buffered()
		return SkyOk(_val)
	}()
}

func Sky_bufio_WriterFlush(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_err := _receiver.Flush()
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_WriterReadFrom(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_arg0 := arg0.(io.Reader)
		_val, _err := _receiver.ReadFrom(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_WriterReset(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_arg0 := arg0.(io.Writer)
		_receiver.Reset(_arg0)
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_WriterSize(receiver any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_val := _receiver.Size()
		return SkyOk(_val)
	}()
}

func Sky_bufio_WriterWrite(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_arg0 := sky_asBytes(arg0)
		_val, _err := _receiver.Write(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_WriterWriteByte(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_arg0 := byte(sky_asInt(arg0))
		_err := _receiver.WriteByte(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(struct{}{})
	}()
}

func Sky_bufio_WriterWriteRune(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_arg0 := rune(sky_asInt(arg0))
		_val, _err := _receiver.WriteRune(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_WriterWriteString(receiver any, arg0 any) any {
	return func() (ret any) {
		defer func() {
			if r := recover(); r != nil {
				ret = SkyErr(_ffi_fmt.Sprintf("FFI panic: %v", r))
			}
		}()
		_receiver := func() *_ffi_pkg.Writer {
			if receiver == nil {
				return nil
			}
			return receiver.(*_ffi_pkg.Writer)
		}()
		_arg0 := sky_asString(arg0)
		_val, _err := _receiver.WriteString(_arg0)
		if _err != nil {
			return SkyErr(_err.Error())
		}
		return SkyOk(_val)
	}()
}

func Sky_bufio_FIELD_ReadWriter_Reader(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Reader")
		if f.IsValid() {
			if f.IsNil() {
				return SkyNothing()
			}
			return SkyJust(f.Interface())
		}
	}
	return SkyNothing()
}

func Sky_bufio_SET_ReadWriter_Reader(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Reader")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val))
	}
	return receiver
}

func Sky_bufio_FIELD_ReadWriter_Writer(receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == _ffi_reflect.Struct {
		f := v.FieldByName("Writer")
		if f.IsValid() {
			if f.IsNil() {
				return SkyNothing()
			}
			return SkyJust(f.Interface())
		}
	}
	return SkyNothing()
}

func Sky_bufio_SET_ReadWriter_Writer(val any, receiver any) any {
	v := _ffi_reflect.ValueOf(receiver)
	for v.Kind() == _ffi_reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != _ffi_reflect.Struct {
		return receiver
	}
	f := v.FieldByName("Writer")
	if !f.IsValid() || !f.CanSet() {
		return receiver
	}
	if val != nil {
		f.Set(_ffi_reflect.ValueOf(val))
	}
	return receiver
}

func Sky_bufio_ErrAdvanceTooFar() any {
	return bufio.ErrAdvanceTooFar
}

func Sky_bufio_ErrBadReadCount() any {
	return bufio.ErrBadReadCount
}

func Sky_bufio_ErrBufferFull() any {
	return bufio.ErrBufferFull
}

func Sky_bufio_ErrFinalToken() any {
	return bufio.ErrFinalToken
}

func Sky_bufio_ErrInvalidUnreadByte() any {
	return bufio.ErrInvalidUnreadByte
}

func Sky_bufio_ErrInvalidUnreadRune() any {
	return bufio.ErrInvalidUnreadRune
}

func Sky_bufio_ErrNegativeAdvance() any {
	return bufio.ErrNegativeAdvance
}

func Sky_bufio_ErrNegativeCount() any {
	return bufio.ErrNegativeCount
}

func Sky_bufio_ErrTooLong() any {
	return bufio.ErrTooLong
}

func Sky_bufio_MaxScanTokenSize() any {
	return bufio.MaxScanTokenSize
}
