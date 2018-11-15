var grpc = require('grpc')
var traffic_grpc = require('./traffic_grpc_pb.js')
var traffic_pb = require('./traffic_pb.js')
var measurements_pb = require('./measurements_pb.js')
var units_pb = require('./units_pb.js')
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');


// Provider Id is a unique identifier assigned to each traffic provider
const PROVIDER_ID = 'demo'

// Endpoint points to sandbox environment
const ENDPOINT = '13.77.181.58:7080'


// Create a client for connecting to the collector.
var client = new traffic_grpc.TrafficProviderClient(ENDPOINT, grpc.credentials.createInsecure())
var source = client.registerProvider()


source.on('data', function (response) {
  console.log('received ack from collector: ', response.getAck().getCount())
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

  var sensor = new traffic_pb.Traffic.Sensor()
  sensor.setPrimaryRadar(new traffic_pb.Traffic.Sensor.PrimaryRadar())
  
  var observed = new google_protobuf_timestamp_pb.Timestamp()
  observed.setSeconds(seconds - 1)
  observed.setNanos(nanos)
  
  var latitude = new units_pb.Degrees()
  var longitude = new units_pb.Degrees()
  latitude.setValue(lat)
  longitude.setValue(lng)
  
  var coordinate = new measurements_pb.Coordinate2D()
  coordinate.setLatitude(latitude)
  coordinate.setLongitude(longitude)
  
  var height = new units_pb.Meters()
  height.setValue(50)

  var altitude = new measurements_pb.Altitude()
  altitude.setHeight(height)
  altitude.setReference(measurements_pb.Altitude.Reference.SURFACE)
  
  var position = new measurements_pb.Position()
  position.setCoordinate(coordinate)
  position.setAltitude(altitude)

  // required
  var provider = new traffic_pb.Traffic.Identity.ProviderId()
  provider.setAsString(PROVIDER_ID)
  var providerIdentity = new traffic_pb.Traffic.Identity()
  providerIdentity.setProviderId(provider)

  // required; unique per track
  var track = new traffic_pb.Traffic.Identity.TrackId()
  track.setAsString("123")
  var trackIdentity = new traffic_pb.Traffic.Identity()
  trackIdentity.setTrackId(track)

  var callsign = new traffic_pb.Traffic.Identity.Callsign()
  callsign.setAsString("AIRMAP1")
  var callsignIdentity = new traffic_pb.Traffic.Identity()
  callsignIdentity.setCallsign(callsign)
  
  var observation = new traffic_pb.Traffic.Observation()
  observation.setObserved(observed)
  observation.setPosition(position)
  observation.setIdentitiesList([trackIdentity, callsignIdentity, providerIdentity])
  observation.setSensor(sensor)

  var submitted = new google_protobuf_timestamp_pb.Timestamp()
  submitted.setSeconds(seconds)
  submitted.setNanos(nanos)

  var observations = [observation]
  
  var update = new traffic_pb.Traffic.Update.FromProvider()
  update.setSubmitted(submitted)
  update.setObservationsList(observations)
  
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
