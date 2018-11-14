// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2018 AirMap, Inc.  All rights reserved.
//
'use strict';
var grpc = require('grpc');
var traffic_pb = require('./traffic_pb.js');
var google_protobuf_any_pb = require('google-protobuf/google/protobuf/any_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var extensions_pb = require('./extensions_pb.js');
var measurements_pb = require('./measurements_pb.js');

function serialize_airmap_Traffic_Update_FromProvider(arg) {
  if (!(arg instanceof traffic_pb.Traffic.Update.FromProvider)) {
    throw new Error('Expected argument of type airmap.Traffic.Update.FromProvider');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_airmap_Traffic_Update_FromProvider(buffer_arg) {
  return traffic_pb.Traffic.Update.FromProvider.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_airmap_Traffic_Update_ToProvider(arg) {
  if (!(arg instanceof traffic_pb.Traffic.Update.ToProvider)) {
    throw new Error('Expected argument of type airmap.Traffic.Update.ToProvider');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_airmap_Traffic_Update_ToProvider(buffer_arg) {
  return traffic_pb.Traffic.Update.ToProvider.deserializeBinary(new Uint8Array(buffer_arg));
}


// TrafficProvider expose a service that enables the exchange of traffic updates with a traffic service.
var TrafficProviderService = exports.TrafficProviderService = {
  // RegisterProvider registers a provider of traffic updates.
  registerProvider: {
    path: '/airmap.TrafficProvider/RegisterProvider',
    requestStream: true,
    responseStream: true,
    requestType: traffic_pb.Traffic.Update.FromProvider,
    responseType: traffic_pb.Traffic.Update.ToProvider,
    requestSerialize: serialize_airmap_Traffic_Update_FromProvider,
    requestDeserialize: deserialize_airmap_Traffic_Update_FromProvider,
    responseSerialize: serialize_airmap_Traffic_Update_ToProvider,
    responseDeserialize: deserialize_airmap_Traffic_Update_ToProvider,
  },
};

exports.TrafficProviderClient = grpc.makeGenericClientConstructor(TrafficProviderService);
