// Bootstrap our dependencies.
var grpc = require('grpc');
var pl = require('@grpc/proto-loader');


// Provider Id is a unique identifier assigned to each traffic provider
const PROVIDER_ID = 'demo'

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
// Please replace with the appropriate URL and appropriate credentials.
var client = new tracking.Collector(ENDPOINT, grpc.credentials.createSsl());
var source = client.ConnectProvider();

source.on('data', function (response) {
  if (response.ack) {
    console.log('received ack from collector: ', response.ack.count)
  } else {
    console.log('received status from collector: ', response.status.level)
  }

})

source.on('end', function() {
  console.log('source connection to collector ended')
})

source.on('error', function(e) {
  console.log('source connection to collector entered error state: ', e)
})



function dispatchUpdates(lat, lng) {

  // the current time in seconds
  var now = new Date().getTime()
  var seconds = Math.floor(now / 1000)
  var nanos = (now - seconds*1000) * 1000

  var observed = {
    seconds: seconds,
    nanos: nanos
  }

  var position = {
    absolute: {
      coordinate: {
        latitude: {
          value: lat
        },
        longitude: {
          value: lng
        }
      },
      altitude: {
        height: {
          value: 50
        },
        reference: {
          value: "REFERENCE_ELLIPSOID"
        }
      }
    }
  }

  var batch = {
    tracks: [
      {
        observed: observed,
        position: position,
        sensor: {
          primary_radar: {}
        },
        identities: [
          {
            provider_id: {
              as_string: PROVIDER_ID
            }
          },
          {
            track_id: {
              as_string: "001"
            }
          }
        ]
      }
    ]
}
  
  var update = {
    batch: batch
  } 

  source.write(update)
}

function main() {
  var lat = 23.3575
  var lng = -109.823

  setInterval(function() {

    lng += 0.001
    lat += 0.001

    console.log(lat)
    console.log(lng)

    dispatchUpdates(lat, lng)

  }, 1000)
}

main()
