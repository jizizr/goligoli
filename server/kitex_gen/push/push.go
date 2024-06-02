// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package push

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
)

type PushBulletRequest struct {
	Bullet *base.Bullet `thrift:"bullet,1" frugal:"1,default,base.Bullet" json:"bullet"`
}

func NewPushBulletRequest() *PushBulletRequest {
	return &PushBulletRequest{}
}

func (p *PushBulletRequest) InitDefault() {
	*p = PushBulletRequest{}
}

var PushBulletRequest_Bullet_DEFAULT *base.Bullet

func (p *PushBulletRequest) GetBullet() (v *base.Bullet) {
	if !p.IsSetBullet() {
		return PushBulletRequest_Bullet_DEFAULT
	}
	return p.Bullet
}
func (p *PushBulletRequest) SetBullet(val *base.Bullet) {
	p.Bullet = val
}

var fieldIDToName_PushBulletRequest = map[int16]string{
	1: "bullet",
}

func (p *PushBulletRequest) IsSetBullet() bool {
	return p.Bullet != nil
}

func (p *PushBulletRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PushBulletRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PushBulletRequest) ReadField1(iprot thrift.TProtocol) error {
	_field := base.NewBullet()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Bullet = _field
	return nil
}

func (p *PushBulletRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PushBulletRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PushBulletRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("bullet", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Bullet.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *PushBulletRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PushBulletRequest(%+v)", *p)

}

func (p *PushBulletRequest) DeepEqual(ano *PushBulletRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Bullet) {
		return false
	}
	return true
}

func (p *PushBulletRequest) Field1DeepEqual(src *base.Bullet) bool {

	if !p.Bullet.DeepEqual(src) {
		return false
	}
	return true
}

type PushService interface {
	PushBullet(ctx context.Context, req *PushBulletRequest) (err error)
}

type PushServiceClient struct {
	c thrift.TClient
}

func NewPushServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PushServiceClient {
	return &PushServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewPushServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PushServiceClient {
	return &PushServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewPushServiceClient(c thrift.TClient) *PushServiceClient {
	return &PushServiceClient{
		c: c,
	}
}

func (p *PushServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *PushServiceClient) PushBullet(ctx context.Context, req *PushBulletRequest) (err error) {
	var _args PushServicePushBulletArgs
	_args.Req = req
	var _result PushServicePushBulletResult
	if err = p.Client_().Call(ctx, "PushBullet", &_args, &_result); err != nil {
		return
	}
	return nil
}

type PushServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      PushService
}

func (p *PushServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *PushServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *PushServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewPushServiceProcessor(handler PushService) *PushServiceProcessor {
	self := &PushServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("PushBullet", &pushServiceProcessorPushBullet{handler: handler})
	return self
}
func (p *PushServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type pushServiceProcessorPushBullet struct {
	handler PushService
}

func (p *pushServiceProcessorPushBullet) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := PushServicePushBulletArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("PushBullet", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := PushServicePushBulletResult{}
	if err2 = p.handler.PushBullet(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing PushBullet: "+err2.Error())
		oprot.WriteMessageBegin("PushBullet", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("PushBullet", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type PushServicePushBulletArgs struct {
	Req *PushBulletRequest `thrift:"req,1" frugal:"1,default,PushBulletRequest" json:"req"`
}

func NewPushServicePushBulletArgs() *PushServicePushBulletArgs {
	return &PushServicePushBulletArgs{}
}

func (p *PushServicePushBulletArgs) InitDefault() {
	*p = PushServicePushBulletArgs{}
}

var PushServicePushBulletArgs_Req_DEFAULT *PushBulletRequest

func (p *PushServicePushBulletArgs) GetReq() (v *PushBulletRequest) {
	if !p.IsSetReq() {
		return PushServicePushBulletArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *PushServicePushBulletArgs) SetReq(val *PushBulletRequest) {
	p.Req = val
}

var fieldIDToName_PushServicePushBulletArgs = map[int16]string{
	1: "req",
}

func (p *PushServicePushBulletArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PushServicePushBulletArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PushServicePushBulletArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PushServicePushBulletArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewPushBulletRequest()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *PushServicePushBulletArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PushBullet_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PushServicePushBulletArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *PushServicePushBulletArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PushServicePushBulletArgs(%+v)", *p)

}

func (p *PushServicePushBulletArgs) DeepEqual(ano *PushServicePushBulletArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *PushServicePushBulletArgs) Field1DeepEqual(src *PushBulletRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type PushServicePushBulletResult struct {
}

func NewPushServicePushBulletResult() *PushServicePushBulletResult {
	return &PushServicePushBulletResult{}
}

func (p *PushServicePushBulletResult) InitDefault() {
	*p = PushServicePushBulletResult{}
}

var fieldIDToName_PushServicePushBulletResult = map[int16]string{}

func (p *PushServicePushBulletResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PushServicePushBulletResult) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("PushBullet_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PushServicePushBulletResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PushServicePushBulletResult(%+v)", *p)

}

func (p *PushServicePushBulletResult) DeepEqual(ano *PushServicePushBulletResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}
