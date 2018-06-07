// Path to the .proto that we are going to use for this example.
var PROTO_PATH = __dirname + '/../../../airmap/grpc/traffic.proto';

// Bootstrap our dependencies.
var ggrpc = require('grpc');
var pl = require('@grpc/proto-loader');

// Load the gRPC interface definitions.
var proto = pl.loadSync(PROTO_PATH, {});
var pkg = ggrpc.loadPackageDefinition(proto);

// Create a client for connecting to the collector.
// Please replace with the appropriate URL and appropriate credentials.
var client = new pkg.grpc.TrafficCollector('localhost:9090', ggrpc.credentials.createInsecure());
var source = client.registerSource();

source.on('data', function(ack) {
  console.log('received ack from collector', ack);
});

source.on('end', function() {
  console.log('source connection to collector ended');
});

source.on('error', function() {
  console.log('source connection to collector entered error state');
});

source.write({
  source_id: {
    as_string: "some.id"
  },
  traffic_id: {
    as_string: "some.track.id"
  },
  submitted: {
    seconds: 42,
    nanos: 42
  },
  observed: {
    seconds: 42,
    nanos: 42
  },
  received: {
    seconds: 42,
    nanos: 42
  },
  ttl: {
    seconds: 42,
    nanos: 0
  },
  icao_description: {
    address: {
      as_string: "43A8F9"
    },
    type: {
      as_string: "prop"
    }
  },
  callsign: "some.callsign",
  registration: "some.registration",
  position: {
    latitude: {
      value: 42.42
    },
    longitude: {
      value: 42.42
    },
    altitude: {
      value: 42.42
    },
    reference: 'ellipsoid'
  },
  speed_over_ground: {
    value: 42.42
  },
  vertical_speed: {
    value: 42.42
  },
  course: {
    value: 42.42
  },
  velocity: {
    value: 42.42
  }
});
// We have successfully registered as a source of traffic data now.
// Call source.write(...) for all the incoming observations you 
// want to dispatch to the collector.
//
// We will now register as a processor of telemetry data, too.
var params = {}
var processor = client.registerProcessor(params);

processor.on('data', function(update) {
  console.log('received a new observation from collector');
});

processor.on('end', function() {
  console.log('processor connection to collector ended');
});

processor.on('error', function() {
  console.log('processor connection to collector entered error state');
});



