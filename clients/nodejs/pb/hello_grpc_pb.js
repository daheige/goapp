// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var hello_pb = require('./hello_pb.js');
// var google_api_annotations_pb = require('./google/api/annotations_pb.js');

function serialize_App_Grpc_Hello_HelloReply(arg) {
  if (!(arg instanceof hello_pb.HelloReply)) {
    throw new Error('Expected argument of type App.Grpc.Hello.HelloReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_HelloReply(buffer_arg) {
  return hello_pb.HelloReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_App_Grpc_Hello_HelloReq(arg) {
  if (!(arg instanceof hello_pb.HelloReq)) {
    throw new Error('Expected argument of type App.Grpc.Hello.HelloReq');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_HelloReq(buffer_arg) {
  return hello_pb.HelloReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_App_Grpc_Hello_InfoReply(arg) {
  if (!(arg instanceof hello_pb.InfoReply)) {
    throw new Error('Expected argument of type App.Grpc.Hello.InfoReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_InfoReply(buffer_arg) {
  return hello_pb.InfoReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_App_Grpc_Hello_InfoReq(arg) {
  if (!(arg instanceof hello_pb.InfoReq)) {
    throw new Error('Expected argument of type App.Grpc.Hello.InfoReq');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_InfoReq(buffer_arg) {
  return hello_pb.InfoReq.deserializeBinary(new Uint8Array(buffer_arg));
}


// service 定义开放调用的服务
var GreeterServiceService = exports.GreeterServiceService = {
  sayHello: {
    path: '/App.Grpc.Hello.GreeterService/SayHello',
    requestStream: false,
    responseStream: false,
    requestType: hello_pb.HelloReq,
    responseType: hello_pb.HelloReply,
    requestSerialize: serialize_App_Grpc_Hello_HelloReq,
    requestDeserialize: deserialize_App_Grpc_Hello_HelloReq,
    responseSerialize: serialize_App_Grpc_Hello_HelloReply,
    responseDeserialize: deserialize_App_Grpc_Hello_HelloReply,
  },
  info: {
    path: '/App.Grpc.Hello.GreeterService/Info',
    requestStream: false,
    responseStream: false,
    requestType: hello_pb.InfoReq,
    responseType: hello_pb.InfoReply,
    requestSerialize: serialize_App_Grpc_Hello_InfoReq,
    requestDeserialize: deserialize_App_Grpc_Hello_InfoReq,
    responseSerialize: serialize_App_Grpc_Hello_InfoReply,
    responseDeserialize: deserialize_App_Grpc_Hello_InfoReply,
  },
};

exports.GreeterServiceClient = grpc.makeGenericClientConstructor(GreeterServiceService);
