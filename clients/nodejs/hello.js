let services = require('./pb/hello_grpc_pb.js');
let messages = require('./pb/hello_pb.js');
let grpc = require('grpc');

let request = new messages.HelloReq();
request.setId(1);

let client = new services.GreeterServiceClient(
    'localhost:50051',
    // 'localhost:50050', //nginx grpc pass port
    grpc.credentials.createInsecure()
);

client.sayHello(request, function(err, data) {
    if (err) {
        console.error(err);
        return;
    }

    console.log(data);
    console.log(data.getMessage());
    console.log(data.getName());
});

let req2 = new messages.InfoReq();
req2.setName("daheige");
client.info(req2, function(err, data) {
    if (err) {
        console.error(err);
        return;
    }

    console.log(data);
    console.log(data.getMessage());
    console.log(data.getAddress());
});

/**
 % node hello.js
 {
  wrappers_: null,
  messageId_: undefined,
  arrayIndexOffset_: -1,
  array: [ 'username: xiaoming', 'call ok' ],
  pivot_: 1.7976931348623157e+308,
  convertedPrimitiveFields_: {}
}
 call ok
 username: xiaoming

 {
  wrappers_: null,
  messageId_: undefined,
  arrayIndexOffset_: -1,
  array: [ 'shenzhen', 'ok' ],
  pivot_: 1.7976931348623157e+308,
  convertedPrimitiveFields_: {}
}
 ok
 shenzhen
 */
