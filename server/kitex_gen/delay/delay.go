// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package delay

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type DelayTaskRequest struct {
	Id      int64 `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	EndTime int64 `thrift:"end_time,2" frugal:"2,default,i64" json:"end_time"`
}

func NewDelayTaskRequest() *DelayTaskRequest {
	return &DelayTaskRequest{}
}

func (p *DelayTaskRequest) InitDefault() {
	*p = DelayTaskRequest{}
}

func (p *DelayTaskRequest) GetId() (v int64) {
	return p.Id
}

func (p *DelayTaskRequest) GetEndTime() (v int64) {
	return p.EndTime
}
func (p *DelayTaskRequest) SetId(val int64) {
	p.Id = val
}
func (p *DelayTaskRequest) SetEndTime(val int64) {
	p.EndTime = val
}

var fieldIDToName_DelayTaskRequest = map[int16]string{
	1: "id",
	2: "end_time",
}

func (p *DelayTaskRequest) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField2(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DelayTaskRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DelayTaskRequest) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Id = _field
	return nil
}
func (p *DelayTaskRequest) ReadField2(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.EndTime = _field
	return nil
}

func (p *DelayTaskRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("DelayTaskRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
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

func (p *DelayTaskRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
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

func (p *DelayTaskRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("end_time", thrift.I64, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.EndTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *DelayTaskRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DelayTaskRequest(%+v)", *p)

}

func (p *DelayTaskRequest) DeepEqual(ano *DelayTaskRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.EndTime) {
		return false
	}
	return true
}

func (p *DelayTaskRequest) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *DelayTaskRequest) Field2DeepEqual(src int64) bool {

	if p.EndTime != src {
		return false
	}
	return true
}

type DelayTaskService interface {
	DelayTask(ctx context.Context, req *DelayTaskRequest) (err error)
}

type DelayTaskServiceClient struct {
	c thrift.TClient
}

func NewDelayTaskServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *DelayTaskServiceClient {
	return &DelayTaskServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewDelayTaskServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *DelayTaskServiceClient {
	return &DelayTaskServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewDelayTaskServiceClient(c thrift.TClient) *DelayTaskServiceClient {
	return &DelayTaskServiceClient{
		c: c,
	}
}

func (p *DelayTaskServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *DelayTaskServiceClient) DelayTask(ctx context.Context, req *DelayTaskRequest) (err error) {
	var _args DelayTaskServiceDelayTaskArgs
	_args.Req = req
	var _result DelayTaskServiceDelayTaskResult
	if err = p.Client_().Call(ctx, "delayTask", &_args, &_result); err != nil {
		return
	}
	return nil
}

type DelayTaskServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      DelayTaskService
}

func (p *DelayTaskServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *DelayTaskServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *DelayTaskServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewDelayTaskServiceProcessor(handler DelayTaskService) *DelayTaskServiceProcessor {
	self := &DelayTaskServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("delayTask", &delayTaskServiceProcessorDelayTask{handler: handler})
	return self
}
func (p *DelayTaskServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type delayTaskServiceProcessorDelayTask struct {
	handler DelayTaskService
}

func (p *delayTaskServiceProcessorDelayTask) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DelayTaskServiceDelayTaskArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("delayTask", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := DelayTaskServiceDelayTaskResult{}
	if err2 = p.handler.DelayTask(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing delayTask: "+err2.Error())
		oprot.WriteMessageBegin("delayTask", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("delayTask", thrift.REPLY, seqId); err2 != nil {
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

type DelayTaskServiceDelayTaskArgs struct {
	Req *DelayTaskRequest `thrift:"req,1" frugal:"1,default,DelayTaskRequest" json:"req"`
}

func NewDelayTaskServiceDelayTaskArgs() *DelayTaskServiceDelayTaskArgs {
	return &DelayTaskServiceDelayTaskArgs{}
}

func (p *DelayTaskServiceDelayTaskArgs) InitDefault() {
	*p = DelayTaskServiceDelayTaskArgs{}
}

var DelayTaskServiceDelayTaskArgs_Req_DEFAULT *DelayTaskRequest

func (p *DelayTaskServiceDelayTaskArgs) GetReq() (v *DelayTaskRequest) {
	if !p.IsSetReq() {
		return DelayTaskServiceDelayTaskArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *DelayTaskServiceDelayTaskArgs) SetReq(val *DelayTaskRequest) {
	p.Req = val
}

var fieldIDToName_DelayTaskServiceDelayTaskArgs = map[int16]string{
	1: "req",
}

func (p *DelayTaskServiceDelayTaskArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DelayTaskServiceDelayTaskArgs) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DelayTaskServiceDelayTaskArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DelayTaskServiceDelayTaskArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewDelayTaskRequest()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *DelayTaskServiceDelayTaskArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("delayTask_args"); err != nil {
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

func (p *DelayTaskServiceDelayTaskArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *DelayTaskServiceDelayTaskArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DelayTaskServiceDelayTaskArgs(%+v)", *p)

}

func (p *DelayTaskServiceDelayTaskArgs) DeepEqual(ano *DelayTaskServiceDelayTaskArgs) bool {
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

func (p *DelayTaskServiceDelayTaskArgs) Field1DeepEqual(src *DelayTaskRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type DelayTaskServiceDelayTaskResult struct {
}

func NewDelayTaskServiceDelayTaskResult() *DelayTaskServiceDelayTaskResult {
	return &DelayTaskServiceDelayTaskResult{}
}

func (p *DelayTaskServiceDelayTaskResult) InitDefault() {
	*p = DelayTaskServiceDelayTaskResult{}
}

var fieldIDToName_DelayTaskServiceDelayTaskResult = map[int16]string{}

func (p *DelayTaskServiceDelayTaskResult) Read(iprot thrift.TProtocol) (err error) {

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

func (p *DelayTaskServiceDelayTaskResult) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("delayTask_result"); err != nil {
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

func (p *DelayTaskServiceDelayTaskResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DelayTaskServiceDelayTaskResult(%+v)", *p)

}

func (p *DelayTaskServiceDelayTaskResult) DeepEqual(ano *DelayTaskServiceDelayTaskResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}
