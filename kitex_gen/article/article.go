// Code generated by thriftgo (0.2.4). DO NOT EDIT.

package article

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type BaseResponse struct {
	StatusCode int64  `thrift:"status_code,1" frugal:"1,default,i64" json:"status_code"`
	StatusMsg  string `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
}

func NewBaseResponse() *BaseResponse {
	return &BaseResponse{}
}

func (p *BaseResponse) InitDefault() {
	*p = BaseResponse{}
}

func (p *BaseResponse) GetStatusCode() (v int64) {
	return p.StatusCode
}

func (p *BaseResponse) GetStatusMsg() (v string) {
	return p.StatusMsg
}
func (p *BaseResponse) SetStatusCode(val int64) {
	p.StatusCode = val
}
func (p *BaseResponse) SetStatusMsg(val string) {
	p.StatusMsg = val
}

var fieldIDToName_BaseResponse = map[int16]string{
	1: "status_code",
	2: "status_msg",
}

func (p *BaseResponse) Read(iprot thrift.TProtocol) (err error) {

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
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BaseResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *BaseResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.StatusCode = v
	}
	return nil
}

func (p *BaseResponse) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.StatusMsg = v
	}
	return nil
}

func (p *BaseResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("BaseResponse"); err != nil {
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

func (p *BaseResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status_code", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.StatusCode); err != nil {
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

func (p *BaseResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status_msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.StatusMsg); err != nil {
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

func (p *BaseResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaseResponse(%+v)", *p)
}

func (p *BaseResponse) DeepEqual(ano *BaseResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.StatusCode) {
		return false
	}
	if !p.Field2DeepEqual(ano.StatusMsg) {
		return false
	}
	return true
}

func (p *BaseResponse) Field1DeepEqual(src int64) bool {

	if p.StatusCode != src {
		return false
	}
	return true
}
func (p *BaseResponse) Field2DeepEqual(src string) bool {

	if strings.Compare(p.StatusMsg, src) != 0 {
		return false
	}
	return true
}

type PostArticleRequest struct {
	Title   string `thrift:"title,1" frugal:"1,default,string" json:"title"`
	Content string `thrift:"content,2" frugal:"2,default,string" json:"content"`
}

func NewPostArticleRequest() *PostArticleRequest {
	return &PostArticleRequest{}
}

func (p *PostArticleRequest) InitDefault() {
	*p = PostArticleRequest{}
}

func (p *PostArticleRequest) GetTitle() (v string) {
	return p.Title
}

func (p *PostArticleRequest) GetContent() (v string) {
	return p.Content
}
func (p *PostArticleRequest) SetTitle(val string) {
	p.Title = val
}
func (p *PostArticleRequest) SetContent(val string) {
	p.Content = val
}

var fieldIDToName_PostArticleRequest = map[int16]string{
	1: "title",
	2: "content",
}

func (p *PostArticleRequest) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PostArticleRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PostArticleRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Title = v
	}
	return nil
}

func (p *PostArticleRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Content = v
	}
	return nil
}

func (p *PostArticleRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PostArticleRequest"); err != nil {
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

func (p *PostArticleRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("title", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Title); err != nil {
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

func (p *PostArticleRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("content", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Content); err != nil {
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

func (p *PostArticleRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PostArticleRequest(%+v)", *p)
}

func (p *PostArticleRequest) DeepEqual(ano *PostArticleRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Title) {
		return false
	}
	if !p.Field2DeepEqual(ano.Content) {
		return false
	}
	return true
}

func (p *PostArticleRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Title, src) != 0 {
		return false
	}
	return true
}
func (p *PostArticleRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Content, src) != 0 {
		return false
	}
	return true
}

type PostArticleResponse struct {
	Base *BaseResponse `thrift:"base,1" frugal:"1,default,BaseResponse" json:"base"`
}

func NewPostArticleResponse() *PostArticleResponse {
	return &PostArticleResponse{}
}

func (p *PostArticleResponse) InitDefault() {
	*p = PostArticleResponse{}
}

var PostArticleResponse_Base_DEFAULT *BaseResponse

func (p *PostArticleResponse) GetBase() (v *BaseResponse) {
	if !p.IsSetBase() {
		return PostArticleResponse_Base_DEFAULT
	}
	return p.Base
}
func (p *PostArticleResponse) SetBase(val *BaseResponse) {
	p.Base = val
}

var fieldIDToName_PostArticleResponse = map[int16]string{
	1: "base",
}

func (p *PostArticleResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *PostArticleResponse) Read(iprot thrift.TProtocol) (err error) {

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
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PostArticleResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PostArticleResponse) ReadField1(iprot thrift.TProtocol) error {
	p.Base = NewBaseResponse()
	if err := p.Base.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *PostArticleResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PostArticleResponse"); err != nil {
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

func (p *PostArticleResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("base", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Base.Write(oprot); err != nil {
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

func (p *PostArticleResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PostArticleResponse(%+v)", *p)
}

func (p *PostArticleResponse) DeepEqual(ano *PostArticleResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *PostArticleResponse) Field1DeepEqual(src *BaseResponse) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

type ArticleService interface {
	PostArticle(ctx context.Context, req *PostArticleResponse) (r *PostArticleResponse, err error)
}

type ArticleServiceClient struct {
	c thrift.TClient
}

func NewArticleServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ArticleServiceClient {
	return &ArticleServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewArticleServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ArticleServiceClient {
	return &ArticleServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewArticleServiceClient(c thrift.TClient) *ArticleServiceClient {
	return &ArticleServiceClient{
		c: c,
	}
}

func (p *ArticleServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *ArticleServiceClient) PostArticle(ctx context.Context, req *PostArticleResponse) (r *PostArticleResponse, err error) {
	var _args ArticleServicePostArticleArgs
	_args.Req = req
	var _result ArticleServicePostArticleResult
	if err = p.Client_().Call(ctx, "PostArticle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type ArticleServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ArticleService
}

func (p *ArticleServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ArticleServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ArticleServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewArticleServiceProcessor(handler ArticleService) *ArticleServiceProcessor {
	self := &ArticleServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("PostArticle", &articleServiceProcessorPostArticle{handler: handler})
	return self
}
func (p *ArticleServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type articleServiceProcessorPostArticle struct {
	handler ArticleService
}

func (p *articleServiceProcessorPostArticle) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ArticleServicePostArticleArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("PostArticle", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := ArticleServicePostArticleResult{}
	var retval *PostArticleResponse
	if retval, err2 = p.handler.PostArticle(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing PostArticle: "+err2.Error())
		oprot.WriteMessageBegin("PostArticle", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("PostArticle", thrift.REPLY, seqId); err2 != nil {
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

type ArticleServicePostArticleArgs struct {
	Req *PostArticleResponse `thrift:"req,1" frugal:"1,default,PostArticleResponse" json:"req"`
}

func NewArticleServicePostArticleArgs() *ArticleServicePostArticleArgs {
	return &ArticleServicePostArticleArgs{}
}

func (p *ArticleServicePostArticleArgs) InitDefault() {
	*p = ArticleServicePostArticleArgs{}
}

var ArticleServicePostArticleArgs_Req_DEFAULT *PostArticleResponse

func (p *ArticleServicePostArticleArgs) GetReq() (v *PostArticleResponse) {
	if !p.IsSetReq() {
		return ArticleServicePostArticleArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *ArticleServicePostArticleArgs) SetReq(val *PostArticleResponse) {
	p.Req = val
}

var fieldIDToName_ArticleServicePostArticleArgs = map[int16]string{
	1: "req",
}

func (p *ArticleServicePostArticleArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ArticleServicePostArticleArgs) Read(iprot thrift.TProtocol) (err error) {

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
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ArticleServicePostArticleArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ArticleServicePostArticleArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewPostArticleResponse()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *ArticleServicePostArticleArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PostArticle_args"); err != nil {
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

func (p *ArticleServicePostArticleArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *ArticleServicePostArticleArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ArticleServicePostArticleArgs(%+v)", *p)
}

func (p *ArticleServicePostArticleArgs) DeepEqual(ano *ArticleServicePostArticleArgs) bool {
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

func (p *ArticleServicePostArticleArgs) Field1DeepEqual(src *PostArticleResponse) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type ArticleServicePostArticleResult struct {
	Success *PostArticleResponse `thrift:"success,0,optional" frugal:"0,optional,PostArticleResponse" json:"success,omitempty"`
}

func NewArticleServicePostArticleResult() *ArticleServicePostArticleResult {
	return &ArticleServicePostArticleResult{}
}

func (p *ArticleServicePostArticleResult) InitDefault() {
	*p = ArticleServicePostArticleResult{}
}

var ArticleServicePostArticleResult_Success_DEFAULT *PostArticleResponse

func (p *ArticleServicePostArticleResult) GetSuccess() (v *PostArticleResponse) {
	if !p.IsSetSuccess() {
		return ArticleServicePostArticleResult_Success_DEFAULT
	}
	return p.Success
}
func (p *ArticleServicePostArticleResult) SetSuccess(x interface{}) {
	p.Success = x.(*PostArticleResponse)
}

var fieldIDToName_ArticleServicePostArticleResult = map[int16]string{
	0: "success",
}

func (p *ArticleServicePostArticleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ArticleServicePostArticleResult) Read(iprot thrift.TProtocol) (err error) {

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
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ArticleServicePostArticleResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ArticleServicePostArticleResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewPostArticleResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *ArticleServicePostArticleResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PostArticle_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
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

func (p *ArticleServicePostArticleResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *ArticleServicePostArticleResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ArticleServicePostArticleResult(%+v)", *p)
}

func (p *ArticleServicePostArticleResult) DeepEqual(ano *ArticleServicePostArticleResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *ArticleServicePostArticleResult) Field0DeepEqual(src *PostArticleResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}