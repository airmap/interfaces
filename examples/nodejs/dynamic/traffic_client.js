// Bootstrap our dependencies.
var grpc = require('grpc');
var pl = require('@grpc/proto-loader');


// Provider Id is a unique identifier assigned to each traffic provider
const PROVIDER_ID = 'demo'

// Endpoint points to sandbox environment
const ENDPOINT = '13.77.181.58:7080'


// Load the gRPC package definitions for the traffic service.
var packageDefinition = pl.loadSync('traffic.proto', {
  includeDirs: [
    __dirname + '/../../../grpc',
    __dirname + "/node_modules/google-proto-files"
  ] 
});
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
var airmap = protoDescriptor.airmap

// Create a client for connecting to the collector.
// Please replace with the appropriate URL and appropriate credentials.
var client = new airmap.TrafficProvider(ENDPOINT, grpc.credentials.createInsecure());
var source = client.registerProvider();



source.on('data', function (response) {
  console.log('received ack from collector: ', response.ack.count.toInt())
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

  var sensor = {
    primaryRadar: {}
  }

  var observed = {
    seconds: seconds,
    nanos: nanos
  }

  var position = {
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
        value: "SURFACE"
      }
    }
  }

  // required
  var providerIdentity = {
    providerId: {
      asString: PROVIDER_ID,
    }
  }

  // required; unique per track
  var trackIdentity = {
    trackId: {
      asString: "123"
    }
  }

  var callsignIdentity = {
    callsign: {
      asString: "AIRMAP1"
    }
  }
    
  var observation = {
    sensor: sensor,
    observed: observed,
    position: position,
    identities: [providerIdentity, trackIdentity, callsignIdentity],
  }

  var submitted = {
    seconds: seconds,
    nanos: nanos
  }
  
  var update = {
    submitted: submitted,
    observations: [observation]
  } 
  
  source.write(update)
}

function main() {
  var lat = 23.3575
  var lng = -109.823

  setInterval(function() {

    lng += 0.001
    lat += 0.001

    dispatchUpdates(lat, lng)

  }, 1000)
}

main()
