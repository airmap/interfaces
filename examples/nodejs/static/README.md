This is an example of how to generate static code using protoc and the node grpc protoc plugin. This example assumes that `protoc` and `grpc_node_plugin` are present on the machine.

To generate the `_pb.js` files, execute the following commands:

```sh
cd examples/nodejs/static/
npm install -g grpc-tools
grpc_tools_node_protoc -I ../../../grpc --js_out=import_style=commonjs,binary:. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` ../../../grpc/*.proto
```

The example client can then be run by:
```sh
cd examples/nodejs/static/
node traffic_client.js
```