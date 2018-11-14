var grpc = require('grpc')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var measurements_pb = require('./measurements_pb.js')
var traffic_pb = require('./traffic_pb.js')
var trafficService = require('./traffic_grpc_pb.js')
var units_pb = require('./units_pb.js')

// Create a client for connecting to the collector.
// Please replace with the appropriate URL and appropriate credentials.
var client = new trafficService.TrafficProviderClient('localhost:7080', grpc.credentials.createInsecure())
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



function dispatchUpdate() {
    var update = new traffic_pb.Traffic.Update.FromProvider()
    var submitted = new google_protobuf_timestamp_pb.Timestamp()
    var observation = new traffic_pb.Traffic.Observation()
    var sensor = new traffic_pb.Traffic.Sensor()
    var identity = new traffic_pb.Traffic.Identity()
    var observedTime = new google_protobuf_timestamp_pb.Timestamp()
    var ttl = new google_protobuf_duration_pb.Duration()
    var position = new measurements_pb.Position()
    var coordinate = new measurements_pb.Coordinate2D()
    var course = new measurements_pb.Course()
    var velocity = new measurements_pb.Velocity()
    var orientation = new measurements_pb.Orientation()

    var seconds = Math.round(new Date().getTime()/1000)

    identity.setProviderId(new traffic_pb.Traffic.Identity.ProviderId('provider.id'))

    submitted.setSeconds(seconds+1)
    update.setSubmitted(submitted)

    observedTime.setSeconds(seconds)
    ttl.setSeconds(1)

    coordinate.setLatitude(new units_pb.Degrees(33.98635))
    coordinate.setLongitude(new units_pb.Degrees(-118.47639))

    position.setCoordinate(coordinate)
    position.setAltitude(new units_pb.Meters(50))

    course.setAngle(0)

    velocity.setX(new units_pb.MetersPerSecond(0))
    velocity.setY(new units_pb.MetersPerSecond(3.2))
    velocity.setZ(new units_pb.MetersPerSecond(0.1))

    orientation.setYaw(new units_pb.Degrees(42))
    orientation.setPitch(new units_pb.Degrees(42))
    orientation.setRoll(new units_pb.Degrees(42))

    observation.setSensor(sensor)
    observation.setIdentitiesList(identity)
    observation.setObserved(observedTime)
    observation.setTtl(ttl)
    observation.setPosition(position)
    observation.setCourse(course)
    observation.setVelocity(velocity)
    observation.setOrientation(orientation)

    update.setObservationsList(observation)
    source.write(update)
}

setInterval(function() {
    dispatchUpdate()
}, 1000)

