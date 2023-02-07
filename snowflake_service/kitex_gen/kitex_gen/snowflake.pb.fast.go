// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package kitex_gen

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *NewIDRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *NewIDResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewIDResponse[number], err)
}

func (x *NewIDResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ID, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *NewIDRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *NewIDResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *NewIDResponse) fastWriteField1(buf []byte) (offset int) {
	if x.ID == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.ID)
	return offset
}

func (x *NewIDRequest) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *NewIDResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *NewIDResponse) sizeField1() (n int) {
	if x.ID == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.ID)
	return n
}

var fieldIDToName_NewIDRequest = map[int32]string{}

var fieldIDToName_NewIDResponse = map[int32]string{
	1: "ID",
}