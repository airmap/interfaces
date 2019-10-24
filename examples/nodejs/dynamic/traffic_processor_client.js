// Bootstrap our dependencies.
var grpc = require('grpc');
var pl = require('@grpc/proto-loader');


// Endpoint points to sandbox environment
const ENDPOINT = 'stage.api.airmap.com:443'


// Load the gRPC package definitions for the traffic service.
var packageDefinition = pl.loadSync(
  [
      'collector.proto',
      'identity.proto',
      'track.proto',
      'emitter.proto',
      'sensors.proto',
      'status.proto',
      'ack.proto',
      'measurements.proto',
      'units.proto',
      'ids.proto'
  ],
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,

    includeDirs: [
     __dirname + '/../../../grpc/tracking',
     __dirname + '/../../../grpc/system',
     __dirname + '/../../../grpc/measurements',
     __dirname + '/../../../grpc/units',
     __dirname + '/../../../grpc/ids',
     __dirname + "/node_modules/google-proto-files"
    ]
});
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
var tracking = protoDescriptor.tracking

// Create a client for connecting to the collector.
var collector = new tracking.Collector(ENDPOINT, grpc.credentials.createSsl());
var client = collector.ConnectProcessor();

client.on('data', function (response) {
  switch (response.details) {
    case "batch":
      for (let i in response.batch.tracks) {
        let track = response.batch.tracks[i]
        console.log(track.identities)
        console.log(track.position.absolute)
      }
      break
    case "status":
      console.log('received status from collector: ', response.status.level, response.status.message)
      break
  }
})

client.on('end', function() {
  console.log('client connection to collector ended')
})

client.on('error', function(e) {
  console.log('client connection to collector entered error state: ', e)
})
