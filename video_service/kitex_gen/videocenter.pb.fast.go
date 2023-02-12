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

func (x *BasicResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BasicResponse[number], err)
}

func (x *BasicResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BasicResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishRequest[number], err)
}

func (x *PublishRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PublishRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PlayUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishResponse[number], err)
}

func (x *PublishResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PublishResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Uuid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteRequest[number], err)
}

func (x *DeleteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uuid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FeedRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FeedRequest[number], err)
}

func (x *FeedRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LastTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FeedResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FeedResponse[number], err)
}

func (x *FeedResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Videos = append(x.Videos, &v)
	return offset, nil
}

func (x *Video) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Video[number], err)
}

func (x *Video) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uuid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PlayUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.FavoriteCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.CommentCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.CreateTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BasicResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *BasicResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *BasicResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *PublishRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *PublishRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *PublishRequest) fastWriteField2(buf []byte) (offset int) {
	if x.PlayUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.PlayUrl)
	return offset
}

func (x *PublishRequest) fastWriteField3(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.CoverUrl)
	return offset
}

func (x *PublishRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.Title)
	return offset
}

func (x *PublishResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *PublishResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *PublishResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *PublishResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Uuid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.Uuid)
	return offset
}

func (x *DeleteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Uuid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Uuid)
	return offset
}

func (x *FeedRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *FeedRequest) fastWriteField1(buf []byte) (offset int) {
	if x.LastTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.LastTime)
	return offset
}

func (x *FeedResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *FeedResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Videos == nil {
		return offset
	}
	for i := range x.Videos {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.Videos[i])
	}
	return offset
}

func (x *Video) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *Video) fastWriteField1(buf []byte) (offset int) {
	if x.Uuid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Uuid)
	return offset
}

func (x *Video) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.UserId)
	return offset
}

func (x *Video) fastWriteField3(buf []byte) (offset int) {
	if x.PlayUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.PlayUrl)
	return offset
}

func (x *Video) fastWriteField4(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.CoverUrl)
	return offset
}

func (x *Video) fastWriteField5(buf []byte) (offset int) {
	if x.FavoriteCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.FavoriteCount)
	return offset
}

func (x *Video) fastWriteField6(buf []byte) (offset int) {
	if x.CommentCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.CommentCount)
	return offset
}

func (x *Video) fastWriteField7(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.Title)
	return offset
}

func (x *Video) fastWriteField8(buf []byte) (offset int) {
	if x.CreateTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 8, x.CreateTime)
	return offset
}

func (x *BasicResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *BasicResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.StatusCode)
	return n
}

func (x *BasicResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *PublishRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *PublishRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *PublishRequest) sizeField2() (n int) {
	if x.PlayUrl == "" {
		return n
	}
	n += fastpb.SizeString(2, x.PlayUrl)
	return n
}

func (x *PublishRequest) sizeField3() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.CoverUrl)
	return n
}

func (x *PublishRequest) sizeField4() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(4, x.Title)
	return n
}

func (x *PublishResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *PublishResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.StatusCode)
	return n
}

func (x *PublishResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *PublishResponse) sizeField3() (n int) {
	if x.Uuid == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.Uuid)
	return n
}

func (x *DeleteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteRequest) sizeField1() (n int) {
	if x.Uuid == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Uuid)
	return n
}

func (x *FeedRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *FeedRequest) sizeField1() (n int) {
	if x.LastTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.LastTime)
	return n
}

func (x *FeedResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *FeedResponse) sizeField1() (n int) {
	if x.Videos == nil {
		return n
	}
	for i := range x.Videos {
		n += fastpb.SizeMessage(1, x.Videos[i])
	}
	return n
}

func (x *Video) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *Video) sizeField1() (n int) {
	if x.Uuid == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Uuid)
	return n
}

func (x *Video) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.UserId)
	return n
}

func (x *Video) sizeField3() (n int) {
	if x.PlayUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.PlayUrl)
	return n
}

func (x *Video) sizeField4() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.CoverUrl)
	return n
}

func (x *Video) sizeField5() (n int) {
	if x.FavoriteCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.FavoriteCount)
	return n
}

func (x *Video) sizeField6() (n int) {
	if x.CommentCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.CommentCount)
	return n
}

func (x *Video) sizeField7() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(7, x.Title)
	return n
}

func (x *Video) sizeField8() (n int) {
	if x.CreateTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(8, x.CreateTime)
	return n
}

var fieldIDToName_BasicResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var fieldIDToName_PublishRequest = map[int32]string{
	1: "UserId",
	2: "PlayUrl",
	3: "CoverUrl",
	4: "Title",
}

var fieldIDToName_PublishResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "Uuid",
}

var fieldIDToName_DeleteRequest = map[int32]string{
	1: "Uuid",
}

var fieldIDToName_FeedRequest = map[int32]string{
	1: "LastTime",
}

var fieldIDToName_FeedResponse = map[int32]string{
	1: "Videos",
}

var fieldIDToName_Video = map[int32]string{
	1: "Uuid",
	2: "UserId",
	3: "PlayUrl",
	4: "CoverUrl",
	5: "FavoriteCount",
	6: "CommentCount",
	7: "Title",
	8: "CreateTime",
}
