// Bootstrap our dependencies.
const grpc = require('grpc');
const pl = require('@grpc/proto-loader');
const request = require('request-promise')

const AUTH_ENDPOINT = 'auth.airmap.com'
const API_ENDPOINT = 'api.airmap.com:443'

// Configure with your client id and client secret
const CLIENT_ID = '{CLIENT_ID}'
const CLIENT_SECRET = '{CLIENT_SECRET}'

// Load the gRPC package definitions for the traffic service.
const packageDefinition = pl.loadSync(
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
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const tracking = protoDescriptor.tracking

async function authenticate() {

  const options = {
    uri: `https://${AUTH_ENDPOINT}/realms/airmap/protocol/openid-connect/token`,
    method: 'POST',
    form: {
      grant_type: "client_credentials",
      client_id: CLIENT_ID,
      client_secret: CLIENT_SECRET
    }
  }

  try {
    const response = await request(options);
    const body = JSON.parse(response)
    return body.access_token
  } catch (error) {
    throw error
  }
}

function connectProcessor(token) {

  const collector = new tracking.Collector(API_ENDPOINT, grpc.credentials.createSsl());

  const metadata = new grpc.Metadata();
  metadata.add("authorization", "Bearer " + token);

  const client = collector.ConnectProcessor(metadata);

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

  client.on('end', function () {
    console.log('client connection to collector ended')
  })

  client.on('error', function (e) {
    console.log('client connection to collector entered error state: ', e)
  })

}

// Login using the client credentials and connect to the collector with a token
authenticate()
  .then(connectProcessor)
  .catch(console.log)
