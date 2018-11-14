/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_descriptor_pb = require('google-protobuf/google/protobuf/descriptor_pb.js');
var units_pb = require('./units_pb.js');
goog.exportSymbol('proto.airmap.defaultCelsius', null, global);
goog.exportSymbol('proto.airmap.defaultDegrees', null, global);
goog.exportSymbol('proto.airmap.defaultMeters', null, global);
goog.exportSymbol('proto.airmap.defaultMetersPerSecond', null, global);
goog.exportSymbol('proto.airmap.defaultPascals', null, global);
goog.exportSymbol('proto.airmap.format', null, global);
goog.exportSymbol('proto.airmap.maxCelsius', null, global);
goog.exportSymbol('proto.airmap.maxDegrees', null, global);
goog.exportSymbol('proto.airmap.maxMeters', null, global);
goog.exportSymbol('proto.airmap.maxMetersPerSecond', null, global);
goog.exportSymbol('proto.airmap.maxPascals', null, global);
goog.exportSymbol('proto.airmap.minCelsius', null, global);
goog.exportSymbol('proto.airmap.minDegrees', null, global);
goog.exportSymbol('proto.airmap.minMeters', null, global);
goog.exportSymbol('proto.airmap.minMetersPerSecond', null, global);
goog.exportSymbol('proto.airmap.minPascals', null, global);
goog.exportSymbol('proto.airmap.nullable', null, global);

/**
 * A tuple of {field number, class constructor} for the extension
 * field named `nullable`.
 * @type {!jspb.ExtensionFieldInfo.<boolean>}
 */
proto.airmap.nullable = new jspb.ExtensionFieldInfo(
    50000,
    {nullable: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50000] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.nullable,
    jspb.BinaryReader.prototype.readBool,
    jspb.BinaryWriter.prototype.writeBool,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50000] = proto.airmap.nullable;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `format`.
 * @type {!jspb.ExtensionFieldInfo.<string>}
 */
proto.airmap.format = new jspb.ExtensionFieldInfo(
    50001,
    {format: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50001] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.format,
    jspb.BinaryReader.prototype.readString,
    jspb.BinaryWriter.prototype.writeString,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50001] = proto.airmap.format;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `defaultDegrees`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Degrees>}
 */
proto.airmap.defaultDegrees = new jspb.ExtensionFieldInfo(
    50002,
    {defaultDegrees: 0},
    units_pb.Degrees,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Degrees.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50002] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.defaultDegrees,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Degrees.serializeBinaryToWriter,
    units_pb.Degrees.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50002] = proto.airmap.defaultDegrees;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `minDegrees`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Degrees>}
 */
proto.airmap.minDegrees = new jspb.ExtensionFieldInfo(
    50003,
    {minDegrees: 0},
    units_pb.Degrees,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Degrees.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50003] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.minDegrees,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Degrees.serializeBinaryToWriter,
    units_pb.Degrees.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50003] = proto.airmap.minDegrees;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `maxDegrees`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Degrees>}
 */
proto.airmap.maxDegrees = new jspb.ExtensionFieldInfo(
    50004,
    {maxDegrees: 0},
    units_pb.Degrees,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Degrees.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50004] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.maxDegrees,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Degrees.serializeBinaryToWriter,
    units_pb.Degrees.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50004] = proto.airmap.maxDegrees;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `defaultMeters`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Meters>}
 */
proto.airmap.defaultMeters = new jspb.ExtensionFieldInfo(
    50005,
    {defaultMeters: 0},
    units_pb.Meters,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Meters.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50005] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.defaultMeters,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Meters.serializeBinaryToWriter,
    units_pb.Meters.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50005] = proto.airmap.defaultMeters;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `minMeters`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Meters>}
 */
proto.airmap.minMeters = new jspb.ExtensionFieldInfo(
    50006,
    {minMeters: 0},
    units_pb.Meters,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Meters.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50006] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.minMeters,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Meters.serializeBinaryToWriter,
    units_pb.Meters.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50006] = proto.airmap.minMeters;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `maxMeters`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Meters>}
 */
proto.airmap.maxMeters = new jspb.ExtensionFieldInfo(
    50007,
    {maxMeters: 0},
    units_pb.Meters,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Meters.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50007] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.maxMeters,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Meters.serializeBinaryToWriter,
    units_pb.Meters.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50007] = proto.airmap.maxMeters;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `defaultMetersPerSecond`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.MetersPerSecond>}
 */
proto.airmap.defaultMetersPerSecond = new jspb.ExtensionFieldInfo(
    50008,
    {defaultMetersPerSecond: 0},
    units_pb.MetersPerSecond,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.MetersPerSecond.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50008] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.defaultMetersPerSecond,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.MetersPerSecond.serializeBinaryToWriter,
    units_pb.MetersPerSecond.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50008] = proto.airmap.defaultMetersPerSecond;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `minMetersPerSecond`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.MetersPerSecond>}
 */
proto.airmap.minMetersPerSecond = new jspb.ExtensionFieldInfo(
    50009,
    {minMetersPerSecond: 0},
    units_pb.MetersPerSecond,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.MetersPerSecond.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50009] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.minMetersPerSecond,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.MetersPerSecond.serializeBinaryToWriter,
    units_pb.MetersPerSecond.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50009] = proto.airmap.minMetersPerSecond;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `maxMetersPerSecond`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.MetersPerSecond>}
 */
proto.airmap.maxMetersPerSecond = new jspb.ExtensionFieldInfo(
    50010,
    {maxMetersPerSecond: 0},
    units_pb.MetersPerSecond,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.MetersPerSecond.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50010] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.maxMetersPerSecond,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.MetersPerSecond.serializeBinaryToWriter,
    units_pb.MetersPerSecond.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50010] = proto.airmap.maxMetersPerSecond;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `defaultPascals`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Pascal>}
 */
proto.airmap.defaultPascals = new jspb.ExtensionFieldInfo(
    50011,
    {defaultPascals: 0},
    units_pb.Pascal,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Pascal.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50011] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.defaultPascals,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Pascal.serializeBinaryToWriter,
    units_pb.Pascal.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50011] = proto.airmap.defaultPascals;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `minPascals`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Pascal>}
 */
proto.airmap.minPascals = new jspb.ExtensionFieldInfo(
    50012,
    {minPascals: 0},
    units_pb.Pascal,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Pascal.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50012] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.minPascals,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Pascal.serializeBinaryToWriter,
    units_pb.Pascal.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50012] = proto.airmap.minPascals;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `maxPascals`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Pascal>}
 */
proto.airmap.maxPascals = new jspb.ExtensionFieldInfo(
    50013,
    {maxPascals: 0},
    units_pb.Pascal,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Pascal.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50013] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.maxPascals,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Pascal.serializeBinaryToWriter,
    units_pb.Pascal.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50013] = proto.airmap.maxPascals;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `defaultCelsius`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Celsius>}
 */
proto.airmap.defaultCelsius = new jspb.ExtensionFieldInfo(
    50014,
    {defaultCelsius: 0},
    units_pb.Celsius,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Celsius.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50014] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.defaultCelsius,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Celsius.serializeBinaryToWriter,
    units_pb.Celsius.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50014] = proto.airmap.defaultCelsius;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `minCelsius`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Celsius>}
 */
proto.airmap.minCelsius = new jspb.ExtensionFieldInfo(
    50015,
    {minCelsius: 0},
    units_pb.Celsius,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Celsius.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50015] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.minCelsius,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Celsius.serializeBinaryToWriter,
    units_pb.Celsius.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50015] = proto.airmap.minCelsius;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `maxCelsius`.
 * @type {!jspb.ExtensionFieldInfo.<!proto.airmap.Celsius>}
 */
proto.airmap.maxCelsius = new jspb.ExtensionFieldInfo(
    50016,
    {maxCelsius: 0},
    units_pb.Celsius,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         units_pb.Celsius.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50016] = new jspb.ExtensionFieldBinaryInfo(
    proto.airmap.maxCelsius,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    units_pb.Celsius.serializeBinaryToWriter,
    units_pb.Celsius.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50016] = proto.airmap.maxCelsius;

goog.object.extend(exports, proto.airmap);
