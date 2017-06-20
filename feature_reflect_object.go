package jsoniter

import (
	"fmt"
	"io"
	"reflect"
	"unsafe"
)

func encoderOfStruct(cfg *frozenConfig, typ reflect.Type) (ValEncoder, error) {
	structEncoder_ := &structEncoder{}
	fields := map[string]*structFieldEncoder{}
	structDescriptor, err := describeStruct(cfg, typ)
	if err != nil {
		return nil, err
	}
	for _, binding := range structDescriptor.Fields {
		for _, fieldName := range binding.ToNames {
			fields[fieldName] = &structFieldEncoder{binding.Field, fieldName, binding.Encoder, binding.ShouldOmitEmpty}
		}
	}
	if len(fields) == 0 {
		return &emptyStructEncoder{}, nil
	}
	for _, field := range fields {
		structEncoder_.fields = append(structEncoder_.fields, field)
	}
	return structEncoder_, nil
}

func decoderOfStruct(cfg *frozenConfig, typ reflect.Type) (ValDecoder, error) {
	fields := map[string]*structFieldDecoder{}
	structDescriptor, err := describeStruct(cfg, typ)
	if err != nil {
		return nil, err
	}
	for _, binding := range structDescriptor.Fields {
		for _, fieldName := range binding.FromNames {
			fields[fieldName] = &structFieldDecoder{binding.Field, binding.Decoder}
		}
	}
	return createStructDecoder(typ, fields)
}

func createStructDecoder(typ reflect.Type, fields map[string]*structFieldDecoder) (ValDecoder, error) {
	knownHash := map[int32]struct{}{
		0: {},
	}
	switch len(fields) {
	case 0:
		return &skipDecoder{typ}, nil
	case 1:
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			return &oneFieldStructDecoder{typ, fieldHash, fieldDecoder}, nil
		}
	case 2:
		var fieldHash1 int32
		var fieldHash2 int32
		var fieldDecoder1 *structFieldDecoder
		var fieldDecoder2 *structFieldDecoder
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			if fieldHash1 == 0 {
				fieldHash1 = fieldHash
				fieldDecoder1 = fieldDecoder
			} else {
				fieldHash2 = fieldHash
				fieldDecoder2 = fieldDecoder
			}
		}
		return &twoFieldsStructDecoder{typ, fieldHash1, fieldDecoder1, fieldHash2, fieldDecoder2}, nil
	case 3:
		var fieldName1 int32
		var fieldName2 int32
		var fieldName3 int32
		var fieldDecoder1 *structFieldDecoder
		var fieldDecoder2 *structFieldDecoder
		var fieldDecoder3 *structFieldDecoder
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			if fieldName1 == 0 {
				fieldName1 = fieldHash
				fieldDecoder1 = fieldDecoder
			} else if fieldName2 == 0 {
				fieldName2 = fieldHash
				fieldDecoder2 = fieldDecoder
			} else {
				fieldName3 = fieldHash
				fieldDecoder3 = fieldDecoder
			}
		}
		return &threeFieldsStructDecoder{typ,
										 fieldName1, fieldDecoder1, fieldName2, fieldDecoder2, fieldName3, fieldDecoder3}, nil
	case 4:
		var fieldName1 int32
		var fieldName2 int32
		var fieldName3 int32
		var fieldName4 int32
		var fieldDecoder1 *structFieldDecoder
		var fieldDecoder2 *structFieldDecoder
		var fieldDecoder3 *structFieldDecoder
		var fieldDecoder4 *structFieldDecoder
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			if fieldName1 == 0 {
				fieldName1 = fieldHash
				fieldDecoder1 = fieldDecoder
			} else if fieldName2 == 0 {
				fieldName2 = fieldHash
				fieldDecoder2 = fieldDecoder
			} else if fieldName3 == 0 {
				fieldName3 = fieldHash
				fieldDecoder3 = fieldDecoder
			} else {
				fieldName4 = fieldHash
				fieldDecoder4 = fieldDecoder
			}
		}
		return &fourFieldsStructDecoder{typ,
										fieldName1, fieldDecoder1, fieldName2, fieldDecoder2, fieldName3, fieldDecoder3,
										fieldName4, fieldDecoder4}, nil
	case 5:
		var fieldName1 int32
		var fieldName2 int32
		var fieldName3 int32
		var fieldName4 int32
		var fieldName5 int32
		var fieldDecoder1 *structFieldDecoder
		var fieldDecoder2 *structFieldDecoder
		var fieldDecoder3 *structFieldDecoder
		var fieldDecoder4 *structFieldDecoder
		var fieldDecoder5 *structFieldDecoder
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			if fieldName1 == 0 {
				fieldName1 = fieldHash
				fieldDecoder1 = fieldDecoder
			} else if fieldName2 == 0 {
				fieldName2 = fieldHash
				fieldDecoder2 = fieldDecoder
			} else if fieldName3 == 0 {
				fieldName3 = fieldHash
				fieldDecoder3 = fieldDecoder
			} else if fieldName4 == 0 {
				fieldName4 = fieldHash
				fieldDecoder4 = fieldDecoder
			} else {
				fieldName5 = fieldHash
				fieldDecoder5 = fieldDecoder
			}
		}
		return &fiveFieldsStructDecoder{typ,
										fieldName1, fieldDecoder1, fieldName2, fieldDecoder2, fieldName3, fieldDecoder3,
										fieldName4, fieldDecoder4, fieldName5, fieldDecoder5}, nil
	case 6:
		var fieldName1 int32
		var fieldName2 int32
		var fieldName3 int32
		var fieldName4 int32
		var fieldName5 int32
		var fieldName6 int32
		var fieldDecoder1 *structFieldDecoder
		var fieldDecoder2 *structFieldDecoder
		var fieldDecoder3 *structFieldDecoder
		var fieldDecoder4 *structFieldDecoder
		var fieldDecoder5 *structFieldDecoder
		var fieldDecoder6 *structFieldDecoder
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			if fieldName1 == 0 {
				fieldName1 = fieldHash
				fieldDecoder1 = fieldDecoder
			} else if fieldName2 == 0 {
				fieldName2 = fieldHash
				fieldDecoder2 = fieldDecoder
			} else if fieldName3 == 0 {
				fieldName3 = fieldHash
				fieldDecoder3 = fieldDecoder
			} else if fieldName4 == 0 {
				fieldName4 = fieldHash
				fieldDecoder4 = fieldDecoder
			} else if fieldName5 == 0 {
				fieldName5 = fieldHash
				fieldDecoder5 = fieldDecoder
			} else {
				fieldName6 = fieldHash
				fieldDecoder6 = fieldDecoder
			}
		}
		return &sixFieldsStructDecoder{typ,
									   fieldName1, fieldDecoder1, fieldName2, fieldDecoder2, fieldName3, fieldDecoder3,
									   fieldName4, fieldDecoder4, fieldName5, fieldDecoder5, fieldName6, fieldDecoder6}, nil
	case 7:
		var fieldName1 int32
		var fieldName2 int32
		var fieldName3 int32
		var fieldName4 int32
		var fieldName5 int32
		var fieldName6 int32
		var fieldName7 int32
		var fieldDecoder1 *structFieldDecoder
		var fieldDecoder2 *structFieldDecoder
		var fieldDecoder3 *structFieldDecoder
		var fieldDecoder4 *structFieldDecoder
		var fieldDecoder5 *structFieldDecoder
		var fieldDecoder6 *structFieldDecoder
		var fieldDecoder7 *structFieldDecoder
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			if fieldName1 == 0 {
				fieldName1 = fieldHash
				fieldDecoder1 = fieldDecoder
			} else if fieldName2 == 0 {
				fieldName2 = fieldHash
				fieldDecoder2 = fieldDecoder
			} else if fieldName3 == 0 {
				fieldName3 = fieldHash
				fieldDecoder3 = fieldDecoder
			} else if fieldName4 == 0 {
				fieldName4 = fieldHash
				fieldDecoder4 = fieldDecoder
			} else if fieldName5 == 0 {
				fieldName5 = fieldHash
				fieldDecoder5 = fieldDecoder
			} else if fieldName6 == 0 {
				fieldName6 = fieldHash
				fieldDecoder6 = fieldDecoder
			} else {
				fieldName7 = fieldHash
				fieldDecoder7 = fieldDecoder
			}
		}
		return &sevenFieldsStructDecoder{typ,
										 fieldName1, fieldDecoder1, fieldName2, fieldDecoder2, fieldName3, fieldDecoder3,
										 fieldName4, fieldDecoder4, fieldName5, fieldDecoder5, fieldName6, fieldDecoder6,
										 fieldName7, fieldDecoder7}, nil
	case 8:
		var fieldName1 int32
		var fieldName2 int32
		var fieldName3 int32
		var fieldName4 int32
		var fieldName5 int32
		var fieldName6 int32
		var fieldName7 int32
		var fieldName8 int32
		var fieldDecoder1 *structFieldDecoder
		var fieldDecoder2 *structFieldDecoder
		var fieldDecoder3 *structFieldDecoder
		var fieldDecoder4 *structFieldDecoder
		var fieldDecoder5 *structFieldDecoder
		var fieldDecoder6 *structFieldDecoder
		var fieldDecoder7 *structFieldDecoder
		var fieldDecoder8 *structFieldDecoder
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			if fieldName1 == 0 {
				fieldName1 = fieldHash
				fieldDecoder1 = fieldDecoder
			} else if fieldName2 == 0 {
				fieldName2 = fieldHash
				fieldDecoder2 = fieldDecoder
			} else if fieldName3 == 0 {
				fieldName3 = fieldHash
				fieldDecoder3 = fieldDecoder
			} else if fieldName4 == 0 {
				fieldName4 = fieldHash
				fieldDecoder4 = fieldDecoder
			} else if fieldName5 == 0 {
				fieldName5 = fieldHash
				fieldDecoder5 = fieldDecoder
			} else if fieldName6 == 0 {
				fieldName6 = fieldHash
				fieldDecoder6 = fieldDecoder
			} else if fieldName7 == 0 {
				fieldName7 = fieldHash
				fieldDecoder7 = fieldDecoder
			} else {
				fieldName8 = fieldHash
				fieldDecoder8 = fieldDecoder
			}
		}
		return &eightFieldsStructDecoder{typ,
										 fieldName1, fieldDecoder1, fieldName2, fieldDecoder2, fieldName3, fieldDecoder3,
										 fieldName4, fieldDecoder4, fieldName5, fieldDecoder5, fieldName6, fieldDecoder6,
										 fieldName7, fieldDecoder7, fieldName8, fieldDecoder8}, nil
	case 9:
		var fieldName1 int32
		var fieldName2 int32
		var fieldName3 int32
		var fieldName4 int32
		var fieldName5 int32
		var fieldName6 int32
		var fieldName7 int32
		var fieldName8 int32
		var fieldName9 int32
		var fieldDecoder1 *structFieldDecoder
		var fieldDecoder2 *structFieldDecoder
		var fieldDecoder3 *structFieldDecoder
		var fieldDecoder4 *structFieldDecoder
		var fieldDecoder5 *structFieldDecoder
		var fieldDecoder6 *structFieldDecoder
		var fieldDecoder7 *structFieldDecoder
		var fieldDecoder8 *structFieldDecoder
		var fieldDecoder9 *structFieldDecoder
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			if fieldName1 == 0 {
				fieldName1 = fieldHash
				fieldDecoder1 = fieldDecoder
			} else if fieldName2 == 0 {
				fieldName2 = fieldHash
				fieldDecoder2 = fieldDecoder
			} else if fieldName3 == 0 {
				fieldName3 = fieldHash
				fieldDecoder3 = fieldDecoder
			} else if fieldName4 == 0 {
				fieldName4 = fieldHash
				fieldDecoder4 = fieldDecoder
			} else if fieldName5 == 0 {
				fieldName5 = fieldHash
				fieldDecoder5 = fieldDecoder
			} else if fieldName6 == 0 {
				fieldName6 = fieldHash
				fieldDecoder6 = fieldDecoder
			} else if fieldName7 == 0 {
				fieldName7 = fieldHash
				fieldDecoder7 = fieldDecoder
			} else if fieldName8 == 0 {
				fieldName8 = fieldHash
				fieldDecoder8 = fieldDecoder
			} else {
				fieldName9 = fieldHash
				fieldDecoder9 = fieldDecoder
			}
		}
		return &nineFieldsStructDecoder{typ,
										fieldName1, fieldDecoder1, fieldName2, fieldDecoder2, fieldName3, fieldDecoder3,
										fieldName4, fieldDecoder4, fieldName5, fieldDecoder5, fieldName6, fieldDecoder6,
										fieldName7, fieldDecoder7, fieldName8, fieldDecoder8, fieldName9, fieldDecoder9}, nil
	case 10:
		var fieldName1 int32
		var fieldName2 int32
		var fieldName3 int32
		var fieldName4 int32
		var fieldName5 int32
		var fieldName6 int32
		var fieldName7 int32
		var fieldName8 int32
		var fieldName9 int32
		var fieldName10 int32
		var fieldDecoder1 *structFieldDecoder
		var fieldDecoder2 *structFieldDecoder
		var fieldDecoder3 *structFieldDecoder
		var fieldDecoder4 *structFieldDecoder
		var fieldDecoder5 *structFieldDecoder
		var fieldDecoder6 *structFieldDecoder
		var fieldDecoder7 *structFieldDecoder
		var fieldDecoder8 *structFieldDecoder
		var fieldDecoder9 *structFieldDecoder
		var fieldDecoder10 *structFieldDecoder
		for fieldName, fieldDecoder := range fields {
			fieldHash := calcHash(fieldName)
			_, known := knownHash[fieldHash]
			if known {
				return &generalStructDecoder{typ, fields}, nil
			} else {
				knownHash[fieldHash] = struct{}{}
			}
			if fieldName1 == 0 {
				fieldName1 = fieldHash
				fieldDecoder1 = fieldDecoder
			} else if fieldName2 == 0 {
				fieldName2 = fieldHash
				fieldDecoder2 = fieldDecoder
			} else if fieldName3 == 0 {
				fieldName3 = fieldHash
				fieldDecoder3 = fieldDecoder
			} else if fieldName4 == 0 {
				fieldName4 = fieldHash
				fieldDecoder4 = fieldDecoder
			} else if fieldName5 == 0 {
				fieldName5 = fieldHash
				fieldDecoder5 = fieldDecoder
			} else if fieldName6 == 0 {
				fieldName6 = fieldHash
				fieldDecoder6 = fieldDecoder
			} else if fieldName7 == 0 {
				fieldName7 = fieldHash
				fieldDecoder7 = fieldDecoder
			} else if fieldName8 == 0 {
				fieldName8 = fieldHash
				fieldDecoder8 = fieldDecoder
			} else if fieldName9 == 0 {
				fieldName9 = fieldHash
				fieldDecoder9 = fieldDecoder
			} else {
				fieldName10 = fieldHash
				fieldDecoder10 = fieldDecoder
			}
		}
		return &tenFieldsStructDecoder{typ,
									   fieldName1, fieldDecoder1, fieldName2, fieldDecoder2, fieldName3, fieldDecoder3,
									   fieldName4, fieldDecoder4, fieldName5, fieldDecoder5, fieldName6, fieldDecoder6,
									   fieldName7, fieldDecoder7, fieldName8, fieldDecoder8, fieldName9, fieldDecoder9,
									   fieldName10, fieldDecoder10}, nil
	}
	return &generalStructDecoder{typ, fields}, nil
}

type generalStructDecoder struct {
	typ    reflect.Type
	fields map[string]*structFieldDecoder
}

func (decoder *generalStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	fieldBytes := iter.readObjectFieldAsBytes()
	field := *(*string)(unsafe.Pointer(&fieldBytes))
	fieldDecoder := decoder.fields[field]
	if fieldDecoder == nil {
		iter.Skip()
	} else {
		fieldDecoder.decode(ptr, iter)
	}
	for iter.nextToken() == ',' {
		fieldBytes = iter.readObjectFieldAsBytes()
		field = *(*string)(unsafe.Pointer(&fieldBytes))
		fieldDecoder = decoder.fields[field]
		if fieldDecoder == nil {
			iter.Skip()
		} else {
			fieldDecoder.decode(ptr, iter)
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type skipDecoder struct {
	typ reflect.Type
}

func (decoder *skipDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	iter.Skip()
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type oneFieldStructDecoder struct {
	typ          reflect.Type
	fieldHash    int32
	fieldDecoder *structFieldDecoder
}

func (decoder *oneFieldStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		if iter.readFieldHash() == decoder.fieldHash {
			decoder.fieldDecoder.decode(ptr, iter)
		} else {
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type twoFieldsStructDecoder struct {
	typ           reflect.Type
	fieldHash1    int32
	fieldDecoder1 *structFieldDecoder
	fieldHash2    int32
	fieldDecoder2 *structFieldDecoder
}

func (decoder *twoFieldsStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		switch iter.readFieldHash() {
		case decoder.fieldHash1:
			decoder.fieldDecoder1.decode(ptr, iter)
		case decoder.fieldHash2:
			decoder.fieldDecoder2.decode(ptr, iter)
		default:
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type threeFieldsStructDecoder struct {
	typ           reflect.Type
	fieldHash1    int32
	fieldDecoder1 *structFieldDecoder
	fieldHash2    int32
	fieldDecoder2 *structFieldDecoder
	fieldHash3    int32
	fieldDecoder3 *structFieldDecoder
}

func (decoder *threeFieldsStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		switch iter.readFieldHash() {
		case decoder.fieldHash1:
			decoder.fieldDecoder1.decode(ptr, iter)
		case decoder.fieldHash2:
			decoder.fieldDecoder2.decode(ptr, iter)
		case decoder.fieldHash3:
			decoder.fieldDecoder3.decode(ptr, iter)
		default:
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type fourFieldsStructDecoder struct {
	typ           reflect.Type
	fieldHash1    int32
	fieldDecoder1 *structFieldDecoder
	fieldHash2    int32
	fieldDecoder2 *structFieldDecoder
	fieldHash3    int32
	fieldDecoder3 *structFieldDecoder
	fieldHash4    int32
	fieldDecoder4 *structFieldDecoder
}

func (decoder *fourFieldsStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		switch iter.readFieldHash() {
		case decoder.fieldHash1:
			decoder.fieldDecoder1.decode(ptr, iter)
		case decoder.fieldHash2:
			decoder.fieldDecoder2.decode(ptr, iter)
		case decoder.fieldHash3:
			decoder.fieldDecoder3.decode(ptr, iter)
		case decoder.fieldHash4:
			decoder.fieldDecoder4.decode(ptr, iter)
		default:
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type fiveFieldsStructDecoder struct {
	typ           reflect.Type
	fieldHash1    int32
	fieldDecoder1 *structFieldDecoder
	fieldHash2    int32
	fieldDecoder2 *structFieldDecoder
	fieldHash3    int32
	fieldDecoder3 *structFieldDecoder
	fieldHash4    int32
	fieldDecoder4 *structFieldDecoder
	fieldHash5    int32
	fieldDecoder5 *structFieldDecoder
}

func (decoder *fiveFieldsStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		switch iter.readFieldHash() {
		case decoder.fieldHash1:
			decoder.fieldDecoder1.decode(ptr, iter)
		case decoder.fieldHash2:
			decoder.fieldDecoder2.decode(ptr, iter)
		case decoder.fieldHash3:
			decoder.fieldDecoder3.decode(ptr, iter)
		case decoder.fieldHash4:
			decoder.fieldDecoder4.decode(ptr, iter)
		case decoder.fieldHash5:
			decoder.fieldDecoder5.decode(ptr, iter)
		default:
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type sixFieldsStructDecoder struct {
	typ           reflect.Type
	fieldHash1    int32
	fieldDecoder1 *structFieldDecoder
	fieldHash2    int32
	fieldDecoder2 *structFieldDecoder
	fieldHash3    int32
	fieldDecoder3 *structFieldDecoder
	fieldHash4    int32
	fieldDecoder4 *structFieldDecoder
	fieldHash5    int32
	fieldDecoder5 *structFieldDecoder
	fieldHash6    int32
	fieldDecoder6 *structFieldDecoder
}

func (decoder *sixFieldsStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		switch iter.readFieldHash() {
		case decoder.fieldHash1:
			decoder.fieldDecoder1.decode(ptr, iter)
		case decoder.fieldHash2:
			decoder.fieldDecoder2.decode(ptr, iter)
		case decoder.fieldHash3:
			decoder.fieldDecoder3.decode(ptr, iter)
		case decoder.fieldHash4:
			decoder.fieldDecoder4.decode(ptr, iter)
		case decoder.fieldHash5:
			decoder.fieldDecoder5.decode(ptr, iter)
		case decoder.fieldHash6:
			decoder.fieldDecoder6.decode(ptr, iter)
		default:
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type sevenFieldsStructDecoder struct {
	typ           reflect.Type
	fieldHash1    int32
	fieldDecoder1 *structFieldDecoder
	fieldHash2    int32
	fieldDecoder2 *structFieldDecoder
	fieldHash3    int32
	fieldDecoder3 *structFieldDecoder
	fieldHash4    int32
	fieldDecoder4 *structFieldDecoder
	fieldHash5    int32
	fieldDecoder5 *structFieldDecoder
	fieldHash6    int32
	fieldDecoder6 *structFieldDecoder
	fieldHash7    int32
	fieldDecoder7 *structFieldDecoder
}

func (decoder *sevenFieldsStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		switch iter.readFieldHash() {
		case decoder.fieldHash1:
			decoder.fieldDecoder1.decode(ptr, iter)
		case decoder.fieldHash2:
			decoder.fieldDecoder2.decode(ptr, iter)
		case decoder.fieldHash3:
			decoder.fieldDecoder3.decode(ptr, iter)
		case decoder.fieldHash4:
			decoder.fieldDecoder4.decode(ptr, iter)
		case decoder.fieldHash5:
			decoder.fieldDecoder5.decode(ptr, iter)
		case decoder.fieldHash6:
			decoder.fieldDecoder6.decode(ptr, iter)
		case decoder.fieldHash7:
			decoder.fieldDecoder7.decode(ptr, iter)
		default:
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type eightFieldsStructDecoder struct {
	typ           reflect.Type
	fieldHash1    int32
	fieldDecoder1 *structFieldDecoder
	fieldHash2    int32
	fieldDecoder2 *structFieldDecoder
	fieldHash3    int32
	fieldDecoder3 *structFieldDecoder
	fieldHash4    int32
	fieldDecoder4 *structFieldDecoder
	fieldHash5    int32
	fieldDecoder5 *structFieldDecoder
	fieldHash6    int32
	fieldDecoder6 *structFieldDecoder
	fieldHash7    int32
	fieldDecoder7 *structFieldDecoder
	fieldHash8    int32
	fieldDecoder8 *structFieldDecoder
}

func (decoder *eightFieldsStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		switch iter.readFieldHash() {
		case decoder.fieldHash1:
			decoder.fieldDecoder1.decode(ptr, iter)
		case decoder.fieldHash2:
			decoder.fieldDecoder2.decode(ptr, iter)
		case decoder.fieldHash3:
			decoder.fieldDecoder3.decode(ptr, iter)
		case decoder.fieldHash4:
			decoder.fieldDecoder4.decode(ptr, iter)
		case decoder.fieldHash5:
			decoder.fieldDecoder5.decode(ptr, iter)
		case decoder.fieldHash6:
			decoder.fieldDecoder6.decode(ptr, iter)
		case decoder.fieldHash7:
			decoder.fieldDecoder7.decode(ptr, iter)
		case decoder.fieldHash8:
			decoder.fieldDecoder8.decode(ptr, iter)
		default:
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type nineFieldsStructDecoder struct {
	typ           reflect.Type
	fieldHash1    int32
	fieldDecoder1 *structFieldDecoder
	fieldHash2    int32
	fieldDecoder2 *structFieldDecoder
	fieldHash3    int32
	fieldDecoder3 *structFieldDecoder
	fieldHash4    int32
	fieldDecoder4 *structFieldDecoder
	fieldHash5    int32
	fieldDecoder5 *structFieldDecoder
	fieldHash6    int32
	fieldDecoder6 *structFieldDecoder
	fieldHash7    int32
	fieldDecoder7 *structFieldDecoder
	fieldHash8    int32
	fieldDecoder8 *structFieldDecoder
	fieldHash9    int32
	fieldDecoder9 *structFieldDecoder
}

func (decoder *nineFieldsStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		switch iter.readFieldHash() {
		case decoder.fieldHash1:
			decoder.fieldDecoder1.decode(ptr, iter)
		case decoder.fieldHash2:
			decoder.fieldDecoder2.decode(ptr, iter)
		case decoder.fieldHash3:
			decoder.fieldDecoder3.decode(ptr, iter)
		case decoder.fieldHash4:
			decoder.fieldDecoder4.decode(ptr, iter)
		case decoder.fieldHash5:
			decoder.fieldDecoder5.decode(ptr, iter)
		case decoder.fieldHash6:
			decoder.fieldDecoder6.decode(ptr, iter)
		case decoder.fieldHash7:
			decoder.fieldDecoder7.decode(ptr, iter)
		case decoder.fieldHash8:
			decoder.fieldDecoder8.decode(ptr, iter)
		case decoder.fieldHash9:
			decoder.fieldDecoder9.decode(ptr, iter)
		default:
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type tenFieldsStructDecoder struct {
	typ            reflect.Type
	fieldHash1     int32
	fieldDecoder1  *structFieldDecoder
	fieldHash2     int32
	fieldDecoder2  *structFieldDecoder
	fieldHash3     int32
	fieldDecoder3  *structFieldDecoder
	fieldHash4     int32
	fieldDecoder4  *structFieldDecoder
	fieldHash5     int32
	fieldDecoder5  *structFieldDecoder
	fieldHash6     int32
	fieldDecoder6  *structFieldDecoder
	fieldHash7     int32
	fieldDecoder7  *structFieldDecoder
	fieldHash8     int32
	fieldDecoder8  *structFieldDecoder
	fieldHash9     int32
	fieldDecoder9  *structFieldDecoder
	fieldHash10    int32
	fieldDecoder10 *structFieldDecoder
}

func (decoder *tenFieldsStructDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.readObjectStart() {
		return
	}
	for {
		switch iter.readFieldHash() {
		case decoder.fieldHash1:
			decoder.fieldDecoder1.decode(ptr, iter)
		case decoder.fieldHash2:
			decoder.fieldDecoder2.decode(ptr, iter)
		case decoder.fieldHash3:
			decoder.fieldDecoder3.decode(ptr, iter)
		case decoder.fieldHash4:
			decoder.fieldDecoder4.decode(ptr, iter)
		case decoder.fieldHash5:
			decoder.fieldDecoder5.decode(ptr, iter)
		case decoder.fieldHash6:
			decoder.fieldDecoder6.decode(ptr, iter)
		case decoder.fieldHash7:
			decoder.fieldDecoder7.decode(ptr, iter)
		case decoder.fieldHash8:
			decoder.fieldDecoder8.decode(ptr, iter)
		case decoder.fieldHash9:
			decoder.fieldDecoder9.decode(ptr, iter)
		case decoder.fieldHash10:
			decoder.fieldDecoder10.decode(ptr, iter)
		default:
			iter.Skip()
		}
		if iter.nextToken() != ',' {
			break
		}
	}
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%v: %s", decoder.typ, iter.Error.Error())
	}
}

type structFieldDecoder struct {
	field        *reflect.StructField
	fieldDecoder ValDecoder
}

func (decoder *structFieldDecoder) decode(ptr unsafe.Pointer, iter *Iterator) {
	fieldPtr := uintptr(ptr) + decoder.field.Offset
	decoder.fieldDecoder.decode(unsafe.Pointer(fieldPtr), iter)
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("%s: %s", decoder.field.Name, iter.Error.Error())
	}
}

type structFieldEncoder struct {
	field        *reflect.StructField
	fieldName    string
	fieldEncoder ValEncoder
	omitempty    bool
}

func (encoder *structFieldEncoder) encode(ptr unsafe.Pointer, stream *Stream) {
	fieldPtr := uintptr(ptr) + encoder.field.Offset
	stream.WriteObjectField(encoder.fieldName)
	encoder.fieldEncoder.encode(unsafe.Pointer(fieldPtr), stream)
	if stream.Error != nil && stream.Error != io.EOF {
		stream.Error = fmt.Errorf("%s: %s", encoder.field.Name, stream.Error.Error())
	}
}

func (encoder *structFieldEncoder) encodeInterface(val interface{}, stream *Stream) {
	writeToStream(val, stream, encoder)
}

func (encoder *structFieldEncoder) isEmpty(ptr unsafe.Pointer) bool {
	fieldPtr := uintptr(ptr) + encoder.field.Offset
	return encoder.fieldEncoder.isEmpty(unsafe.Pointer(fieldPtr))
}

type structEncoder struct {
	fields []*structFieldEncoder
}

func (encoder *structEncoder) encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteObjectStart()
	isNotFirst := false
	for _, field := range encoder.fields {
		if field.omitempty && field.isEmpty(ptr) {
			continue
		}
		if isNotFirst {
			stream.WriteMore()
		}
		field.encode(ptr, stream)
		isNotFirst = true
	}
	stream.WriteObjectEnd()
}

func (encoder *structEncoder) encodeInterface(val interface{}, stream *Stream) {
	var encoderToUse ValEncoder
	encoderToUse = encoder
	if len(encoder.fields) == 1 {
		firstEncoder := encoder.fields[0].fieldEncoder
		firstEncoderName := reflect.TypeOf(firstEncoder).String()
		// interface{} has inline optimization for this case
		if firstEncoderName == "*jsoniter.optionalEncoder" {
			encoderToUse = &structEncoder{
				fields: []*structFieldEncoder{{
					field:        encoder.fields[0].field,
					fieldName:    encoder.fields[0].fieldName,
					fieldEncoder: firstEncoder.(*optionalEncoder).valueEncoder,
					omitempty:    encoder.fields[0].omitempty,
				}},
			}
		}
	}
	writeToStream(val, stream, encoderToUse)
}

func (encoder *structEncoder) isEmpty(ptr unsafe.Pointer) bool {
	for _, field := range encoder.fields {
		if !field.isEmpty(ptr) {
			return false
		}
	}
	return true
}

type emptyStructEncoder struct {
}

func (encoder *emptyStructEncoder) encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteEmptyObject()
}

func (encoder *emptyStructEncoder) encodeInterface(val interface{}, stream *Stream) {
	writeToStream(val, stream, encoder)
}

func (encoder *emptyStructEncoder) isEmpty(ptr unsafe.Pointer) bool {
	return true
}
