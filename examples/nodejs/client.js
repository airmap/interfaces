var grpc = require('grpc')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var measurements_pb = require('./measurements_pb.js')
var traffic_pb = require('./traffic_pb.js')
var trafficService = require('./traffic_grpc_pb.js')

// Create a client for connecting to the collector.
// Please replace with the appropriate URL and appropriate credentials.
var client = new trafficService.TrafficProviderClient('13.77.142.211:7080', grpc.credentials.createInsecure())
var source = client.registerProvider()

source.on('data', function(ack) {
  console.log('received ack from collector: ', ack)
})

source.on('end', function() {
  console.log('source connection to collector ended')
})

source.on('error', function(e) {
  console.log('source connection to collector entered error state: ', e)
})

var update = new traffic_pb.Traffic.Update.FromProvider()

var submitted = new google_protobuf_timestamp_pb.Timestamp()
var observation = new traffic_pb.Traffic.Observation()
var sensor = new traffic_pb.Traffic.Sensor()
var identity = new traffic_pb.Traffic.Identity()

var observedTime = new google_protobuf_timestamp_pb.Timestamp()
var ttl = new google_protobuf_duration_pb.Duration()
var positions = new measurements_pb.Position()
var coordinate = new measurements_pb.Coordinate2D()


submitted.setSeconds(1542145471)
update.setSubmitted(submitted)


update.setObservationsList(observation)

source.write(update)

/*
source.write({
    submitted: "2018-11-13T21:44:31Z",
    received: "2018-11-13T21:44:32Z",
    observations: [
        {
            identities: [
                {
                    provider_id: {
                        as_string: 'some.provider.id'
                    },
                    track_id: {
                        as_string: 'some.track.id'
                    },
                    callsign: {
                        as_string: "some.callsign"
                    },
                    registration: {
                        as_string: "some.registration"
                    },
                    icao: {
                        address: {
                            as_string: "some.address"
                        },
                        aircraft_type: {
                            as_string: "some.type"
                        }
                    },
                    manufacturer: {
                        make: "some.make",
                        model: "some.model",
                        serial_number: "some.serial.number"
                    },
                    network_adress: {
                        as_string: "some.address"
                    }

                }
            ],
            observed: "2018-11-13T21:44:30Z",
            ttl: {
                seconds: 1,
                nanos: 1
            },
            position: {
                coordinate: {
                    latitude: {
                        value: 33.98635,
                    },
                    longitude: {
                        value: -118.47639
                    }
                },
                altitude: {
                    height: {
                        value: 50
                    },
                    reference: "ELLIPSOID"
                }
            },
            course: {   
                angle: {
                    value: 0
                }
            },
            velocity: {
                x: {
                    value: 0
                },
                y: {
                    value: 3.2
                },
                z: {
                    value: 0.1
                }
            },
            orientation: {
                yaw: {
                    value: 0
                },
                pitch: {
                    value: 0
                },
                roll: {
                    value: 0
                }
            }
        }
    ]
})
*/
