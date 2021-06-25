// Bootstrap our dependencies.
const grpc = require('grpc')
const auth = require('../auth')
const proto = require('../proto')
const fs = require('fs')

// Configure your env with your client id and client secret
const API_ENDPOINT = process.env.API_ENDPOINT || 'api.airmap.com:443'
const CLIENT_ID = process.env.CLIENT_ID
const CLIENT_SECRET = process.env.CLIENT_SECRET

var items_interval
var zeros_counter

function nop() {}  // Just a no-op callback

function pub_metric() {
  const d = new Date()
  fs.appendFile('log.log', `${d.toISOString()}, ${items_interval} items\n`, nop)
  fs.writeFile('metric', items_interval, nop)
  // track how many times this metric reports a zero value. If 5x in a row
  // then we restart everything
  if (items_interval == 0) {
    zeros_counter++
    if (zeros_counter == 5) {
      fs.appendFile('log.log', 'FATAL: 5 zero values in a row, restarting processor!\n', nop)
      login()
    }
  } else {
    zeros_counter = 0
  }
  items_interval = 0
  setTimeout(pub_metric, 60000)
}

function connectProcessor(token) {

  const tracking = proto.descriptors.tracking
  const collector = new tracking.Collector(API_ENDPOINT, grpc.credentials.createSsl())

  const metadata = new grpc.Metadata()
  metadata.add("authorization", "Bearer " + token)

  const client = collector.ConnectProcessor(metadata)

  items_interval = 0
  zeros_counter = 0
  setTimeout(pub_metric, 60000)

  client.on('data', function (response) {
    switch (response.details) {
      case "batch":
        items_interval += Object.keys(response.batch.tracks).length
        break
      case "status":
        fs.appendFile('log.log',
          `received status from collector: ${response.status.level}, ${response.status.message}\n`,
          nop)
        break
    }
  })

  client.on('end', function () {
    fs.appendFile('log.log','client connection to collector ended\n',nop)
    login()
  })

  client.on('error', function (error) {
    // throw error
    fs.appendFile('log.log',
      `hit error event from client, code:${error.code}, details:${error.details}\ntrying to restart\n`,
      nop)
    login()
  })
}

function login() {
  // Login using the client credentials,
  // and connect to the collector with a token
  auth.authenticateServiceAccount(CLIENT_ID, CLIENT_SECRET)
    .then(connectProcessor)
    .catch(console.log)
}

login()
