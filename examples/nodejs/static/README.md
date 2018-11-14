This is an example of how to generate static code using protoc and the node grpc protoc plugin. This example assumes that `protoc` and `grpc_node_plugin` are present on the machine.

To generate the `_pb.js` files, execute the following commands beginning in this directory:

```sh
cd ../../..
npm install -g grpc-tools
grpc_tools_node_protoc -I ./grpc --js_out=import_style=commonjs,binary:examples/nodejs/static --grpc_out=./examples/nodejs/static --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` grpc/extensions.proto grpc/measurements.proto grpc/traffic.proto grpc/units.proto
```

The example client can then be run by:
```sh
node examples/nodejs/static/traffic_client.js
```