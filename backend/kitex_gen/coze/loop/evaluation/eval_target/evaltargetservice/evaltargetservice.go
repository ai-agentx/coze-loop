// Code generated by Kitex v0.13.1. DO NOT EDIT.

package evaltargetservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	eval_target "github.com/coze-dev/coze-loop/backend/kitex_gen/coze/loop/evaluation/eval_target"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateEvalTarget": kitex.NewMethodInfo(
		createEvalTargetHandler,
		newEvalTargetServiceCreateEvalTargetArgs,
		newEvalTargetServiceCreateEvalTargetResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"BatchGetEvalTargetsBySource": kitex.NewMethodInfo(
		batchGetEvalTargetsBySourceHandler,
		newEvalTargetServiceBatchGetEvalTargetsBySourceArgs,
		newEvalTargetServiceBatchGetEvalTargetsBySourceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetEvalTargetVersion": kitex.NewMethodInfo(
		getEvalTargetVersionHandler,
		newEvalTargetServiceGetEvalTargetVersionArgs,
		newEvalTargetServiceGetEvalTargetVersionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"BatchGetEvalTargetVersions": kitex.NewMethodInfo(
		batchGetEvalTargetVersionsHandler,
		newEvalTargetServiceBatchGetEvalTargetVersionsArgs,
		newEvalTargetServiceBatchGetEvalTargetVersionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListSourceEvalTargets": kitex.NewMethodInfo(
		listSourceEvalTargetsHandler,
		newEvalTargetServiceListSourceEvalTargetsArgs,
		newEvalTargetServiceListSourceEvalTargetsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListSourceEvalTargetVersions": kitex.NewMethodInfo(
		listSourceEvalTargetVersionsHandler,
		newEvalTargetServiceListSourceEvalTargetVersionsArgs,
		newEvalTargetServiceListSourceEvalTargetVersionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ExecuteEvalTarget": kitex.NewMethodInfo(
		executeEvalTargetHandler,
		newEvalTargetServiceExecuteEvalTargetArgs,
		newEvalTargetServiceExecuteEvalTargetResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetEvalTargetRecord": kitex.NewMethodInfo(
		getEvalTargetRecordHandler,
		newEvalTargetServiceGetEvalTargetRecordArgs,
		newEvalTargetServiceGetEvalTargetRecordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"BatchGetEvalTargetRecords": kitex.NewMethodInfo(
		batchGetEvalTargetRecordsHandler,
		newEvalTargetServiceBatchGetEvalTargetRecordsArgs,
		newEvalTargetServiceBatchGetEvalTargetRecordsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	evalTargetServiceServiceInfo = NewServiceInfo()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return evalTargetServiceServiceInfo
}

// NewServiceInfo creates a new ServiceInfo
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo()
}

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "EvalTargetService"
	handlerType := (*eval_target.EvalTargetService)(nil)
	extra := map[string]interface{}{
		"PackageName": "eval_target",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         serviceMethods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.13.1",
		Extra:           extra,
	}
	return svcInfo
}

func createEvalTargetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_target.EvalTargetServiceCreateEvalTargetArgs)
	realResult := result.(*eval_target.EvalTargetServiceCreateEvalTargetResult)
	success, err := handler.(eval_target.EvalTargetService).CreateEvalTarget(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvalTargetServiceCreateEvalTargetArgs() interface{} {
	return eval_target.NewEvalTargetServiceCreateEvalTargetArgs()
}

func newEvalTargetServiceCreateEvalTargetResult() interface{} {
	return eval_target.NewEvalTargetServiceCreateEvalTargetResult()
}

func batchGetEvalTargetsBySourceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_target.EvalTargetServiceBatchGetEvalTargetsBySourceArgs)
	realResult := result.(*eval_target.EvalTargetServiceBatchGetEvalTargetsBySourceResult)
	success, err := handler.(eval_target.EvalTargetService).BatchGetEvalTargetsBySource(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvalTargetServiceBatchGetEvalTargetsBySourceArgs() interface{} {
	return eval_target.NewEvalTargetServiceBatchGetEvalTargetsBySourceArgs()
}

func newEvalTargetServiceBatchGetEvalTargetsBySourceResult() interface{} {
	return eval_target.NewEvalTargetServiceBatchGetEvalTargetsBySourceResult()
}

func getEvalTargetVersionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_target.EvalTargetServiceGetEvalTargetVersionArgs)
	realResult := result.(*eval_target.EvalTargetServiceGetEvalTargetVersionResult)
	success, err := handler.(eval_target.EvalTargetService).GetEvalTargetVersion(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvalTargetServiceGetEvalTargetVersionArgs() interface{} {
	return eval_target.NewEvalTargetServiceGetEvalTargetVersionArgs()
}

func newEvalTargetServiceGetEvalTargetVersionResult() interface{} {
	return eval_target.NewEvalTargetServiceGetEvalTargetVersionResult()
}

func batchGetEvalTargetVersionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_target.EvalTargetServiceBatchGetEvalTargetVersionsArgs)
	realResult := result.(*eval_target.EvalTargetServiceBatchGetEvalTargetVersionsResult)
	success, err := handler.(eval_target.EvalTargetService).BatchGetEvalTargetVersions(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvalTargetServiceBatchGetEvalTargetVersionsArgs() interface{} {
	return eval_target.NewEvalTargetServiceBatchGetEvalTargetVersionsArgs()
}

func newEvalTargetServiceBatchGetEvalTargetVersionsResult() interface{} {
	return eval_target.NewEvalTargetServiceBatchGetEvalTargetVersionsResult()
}

func listSourceEvalTargetsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_target.EvalTargetServiceListSourceEvalTargetsArgs)
	realResult := result.(*eval_target.EvalTargetServiceListSourceEvalTargetsResult)
	success, err := handler.(eval_target.EvalTargetService).ListSourceEvalTargets(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvalTargetServiceListSourceEvalTargetsArgs() interface{} {
	return eval_target.NewEvalTargetServiceListSourceEvalTargetsArgs()
}

func newEvalTargetServiceListSourceEvalTargetsResult() interface{} {
	return eval_target.NewEvalTargetServiceListSourceEvalTargetsResult()
}

func listSourceEvalTargetVersionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_target.EvalTargetServiceListSourceEvalTargetVersionsArgs)
	realResult := result.(*eval_target.EvalTargetServiceListSourceEvalTargetVersionsResult)
	success, err := handler.(eval_target.EvalTargetService).ListSourceEvalTargetVersions(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvalTargetServiceListSourceEvalTargetVersionsArgs() interface{} {
	return eval_target.NewEvalTargetServiceListSourceEvalTargetVersionsArgs()
}

func newEvalTargetServiceListSourceEvalTargetVersionsResult() interface{} {
	return eval_target.NewEvalTargetServiceListSourceEvalTargetVersionsResult()
}

func executeEvalTargetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_target.EvalTargetServiceExecuteEvalTargetArgs)
	realResult := result.(*eval_target.EvalTargetServiceExecuteEvalTargetResult)
	success, err := handler.(eval_target.EvalTargetService).ExecuteEvalTarget(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvalTargetServiceExecuteEvalTargetArgs() interface{} {
	return eval_target.NewEvalTargetServiceExecuteEvalTargetArgs()
}

func newEvalTargetServiceExecuteEvalTargetResult() interface{} {
	return eval_target.NewEvalTargetServiceExecuteEvalTargetResult()
}

func getEvalTargetRecordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_target.EvalTargetServiceGetEvalTargetRecordArgs)
	realResult := result.(*eval_target.EvalTargetServiceGetEvalTargetRecordResult)
	success, err := handler.(eval_target.EvalTargetService).GetEvalTargetRecord(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvalTargetServiceGetEvalTargetRecordArgs() interface{} {
	return eval_target.NewEvalTargetServiceGetEvalTargetRecordArgs()
}

func newEvalTargetServiceGetEvalTargetRecordResult() interface{} {
	return eval_target.NewEvalTargetServiceGetEvalTargetRecordResult()
}

func batchGetEvalTargetRecordsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_target.EvalTargetServiceBatchGetEvalTargetRecordsArgs)
	realResult := result.(*eval_target.EvalTargetServiceBatchGetEvalTargetRecordsResult)
	success, err := handler.(eval_target.EvalTargetService).BatchGetEvalTargetRecords(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvalTargetServiceBatchGetEvalTargetRecordsArgs() interface{} {
	return eval_target.NewEvalTargetServiceBatchGetEvalTargetRecordsArgs()
}

func newEvalTargetServiceBatchGetEvalTargetRecordsResult() interface{} {
	return eval_target.NewEvalTargetServiceBatchGetEvalTargetRecordsResult()
}

type kClient struct {
	c  client.Client
	sc client.Streaming
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c:  c,
		sc: c.(client.Streaming),
	}
}

func (p *kClient) CreateEvalTarget(ctx context.Context, request *eval_target.CreateEvalTargetRequest) (r *eval_target.CreateEvalTargetResponse, err error) {
	var _args eval_target.EvalTargetServiceCreateEvalTargetArgs
	_args.Request = request
	var _result eval_target.EvalTargetServiceCreateEvalTargetResult
	if err = p.c.Call(ctx, "CreateEvalTarget", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchGetEvalTargetsBySource(ctx context.Context, request *eval_target.BatchGetEvalTargetsBySourceRequest) (r *eval_target.BatchGetEvalTargetsBySourceResponse, err error) {
	var _args eval_target.EvalTargetServiceBatchGetEvalTargetsBySourceArgs
	_args.Request = request
	var _result eval_target.EvalTargetServiceBatchGetEvalTargetsBySourceResult
	if err = p.c.Call(ctx, "BatchGetEvalTargetsBySource", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetEvalTargetVersion(ctx context.Context, request *eval_target.GetEvalTargetVersionRequest) (r *eval_target.GetEvalTargetVersionResponse, err error) {
	var _args eval_target.EvalTargetServiceGetEvalTargetVersionArgs
	_args.Request = request
	var _result eval_target.EvalTargetServiceGetEvalTargetVersionResult
	if err = p.c.Call(ctx, "GetEvalTargetVersion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchGetEvalTargetVersions(ctx context.Context, request *eval_target.BatchGetEvalTargetVersionsRequest) (r *eval_target.BatchGetEvalTargetVersionsResponse, err error) {
	var _args eval_target.EvalTargetServiceBatchGetEvalTargetVersionsArgs
	_args.Request = request
	var _result eval_target.EvalTargetServiceBatchGetEvalTargetVersionsResult
	if err = p.c.Call(ctx, "BatchGetEvalTargetVersions", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListSourceEvalTargets(ctx context.Context, request *eval_target.ListSourceEvalTargetsRequest) (r *eval_target.ListSourceEvalTargetsResponse, err error) {
	var _args eval_target.EvalTargetServiceListSourceEvalTargetsArgs
	_args.Request = request
	var _result eval_target.EvalTargetServiceListSourceEvalTargetsResult
	if err = p.c.Call(ctx, "ListSourceEvalTargets", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListSourceEvalTargetVersions(ctx context.Context, request *eval_target.ListSourceEvalTargetVersionsRequest) (r *eval_target.ListSourceEvalTargetVersionsResponse, err error) {
	var _args eval_target.EvalTargetServiceListSourceEvalTargetVersionsArgs
	_args.Request = request
	var _result eval_target.EvalTargetServiceListSourceEvalTargetVersionsResult
	if err = p.c.Call(ctx, "ListSourceEvalTargetVersions", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ExecuteEvalTarget(ctx context.Context, request *eval_target.ExecuteEvalTargetRequest) (r *eval_target.ExecuteEvalTargetResponse, err error) {
	var _args eval_target.EvalTargetServiceExecuteEvalTargetArgs
	_args.Request = request
	var _result eval_target.EvalTargetServiceExecuteEvalTargetResult
	if err = p.c.Call(ctx, "ExecuteEvalTarget", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetEvalTargetRecord(ctx context.Context, request *eval_target.GetEvalTargetRecordRequest) (r *eval_target.GetEvalTargetRecordResponse, err error) {
	var _args eval_target.EvalTargetServiceGetEvalTargetRecordArgs
	_args.Request = request
	var _result eval_target.EvalTargetServiceGetEvalTargetRecordResult
	if err = p.c.Call(ctx, "GetEvalTargetRecord", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchGetEvalTargetRecords(ctx context.Context, request *eval_target.BatchGetEvalTargetRecordsRequest) (r *eval_target.BatchGetEvalTargetRecordsResponse, err error) {
	var _args eval_target.EvalTargetServiceBatchGetEvalTargetRecordsArgs
	_args.Request = request
	var _result eval_target.EvalTargetServiceBatchGetEvalTargetRecordsResult
	if err = p.c.Call(ctx, "BatchGetEvalTargetRecords", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
