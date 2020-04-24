# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tracking/sensors.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import extensions_pb2 as extensions__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='tracking/sensors.proto',
  package='tracking',
  syntax='proto3',
  serialized_options=_b('Z,github.com/airmap/interfaces/src/go/tracking'),
  serialized_pb=_b('\n\x16tracking/sensors.proto\x12\x08tracking\x1a\x10\x65xtensions.proto\"\xf8\x03\n\x06Sensor\x12\'\n\x05\x66used\x18\x01 \x01(\x0b\x32\x16.tracking.Sensor.FusedH\x00\x12\x36\n\rprimary_radar\x18\x02 \x01(\x0b\x32\x1d.tracking.Sensor.PrimaryRadarH\x00\x12:\n\x0fsecondary_radar\x18\x03 \x01(\x0b\x32\x1f.tracking.Sensor.SecondaryRadarH\x00\x12%\n\x04\x61\x64sb\x18\x04 \x01(\x0b\x32\x15.tracking.Sensor.AdsbH\x00\x12%\n\x04gnss\x18\x05 \x01(\x0b\x32\x15.tracking.Sensor.GnssH\x00\x12:\n\x0fradio_frequency\x18\x06 \x01(\x0b\x32\x1f.tracking.Sensor.RadioFrequencyH\x00\x12:\n\x0f\x65lectro_optical\x18\x07 \x01(\x0b\x32\x1f.tracking.Sensor.ElectroOpticalH\x00\x1a*\n\x05\x46used\x12!\n\x07sensors\x18\x01 \x03(\x0b\x32\x10.tracking.Sensor\x1a\x0e\n\x0cPrimaryRadar\x1a\x10\n\x0eSecondaryRadar\x1a\x06\n\x04\x41\x64sb\x1a\x10\n\x0eRadioFrequency\x1a\x06\n\x04Gnss\x1a\x10\n\x0e\x45lectroOpticalB\t\n\x07\x64\x65tailsB.Z,github.com/airmap/interfaces/src/go/trackingb\x06proto3')
  ,
  dependencies=[extensions__pb2.DESCRIPTOR,])




_SENSOR_FUSED = _descriptor.Descriptor(
  name='Fused',
  full_name='tracking.Sensor.Fused',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='sensors', full_name='tracking.Sensor.Fused.sensors', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=420,
  serialized_end=462,
)

_SENSOR_PRIMARYRADAR = _descriptor.Descriptor(
  name='PrimaryRadar',
  full_name='tracking.Sensor.PrimaryRadar',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=464,
  serialized_end=478,
)

_SENSOR_SECONDARYRADAR = _descriptor.Descriptor(
  name='SecondaryRadar',
  full_name='tracking.Sensor.SecondaryRadar',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=480,
  serialized_end=496,
)

_SENSOR_ADSB = _descriptor.Descriptor(
  name='Adsb',
  full_name='tracking.Sensor.Adsb',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=498,
  serialized_end=504,
)

_SENSOR_RADIOFREQUENCY = _descriptor.Descriptor(
  name='RadioFrequency',
  full_name='tracking.Sensor.RadioFrequency',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=506,
  serialized_end=522,
)

_SENSOR_GNSS = _descriptor.Descriptor(
  name='Gnss',
  full_name='tracking.Sensor.Gnss',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=524,
  serialized_end=530,
)

_SENSOR_ELECTROOPTICAL = _descriptor.Descriptor(
  name='ElectroOptical',
  full_name='tracking.Sensor.ElectroOptical',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=532,
  serialized_end=548,
)

_SENSOR = _descriptor.Descriptor(
  name='Sensor',
  full_name='tracking.Sensor',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='fused', full_name='tracking.Sensor.fused', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='primary_radar', full_name='tracking.Sensor.primary_radar', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='secondary_radar', full_name='tracking.Sensor.secondary_radar', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='adsb', full_name='tracking.Sensor.adsb', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='gnss', full_name='tracking.Sensor.gnss', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='radio_frequency', full_name='tracking.Sensor.radio_frequency', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='electro_optical', full_name='tracking.Sensor.electro_optical', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SENSOR_FUSED, _SENSOR_PRIMARYRADAR, _SENSOR_SECONDARYRADAR, _SENSOR_ADSB, _SENSOR_RADIOFREQUENCY, _SENSOR_GNSS, _SENSOR_ELECTROOPTICAL, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='details', full_name='tracking.Sensor.details',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=55,
  serialized_end=559,
)

_SENSOR_FUSED.fields_by_name['sensors'].message_type = _SENSOR
_SENSOR_FUSED.containing_type = _SENSOR
_SENSOR_PRIMARYRADAR.containing_type = _SENSOR
_SENSOR_SECONDARYRADAR.containing_type = _SENSOR
_SENSOR_ADSB.containing_type = _SENSOR
_SENSOR_RADIOFREQUENCY.containing_type = _SENSOR
_SENSOR_GNSS.containing_type = _SENSOR
_SENSOR_ELECTROOPTICAL.containing_type = _SENSOR
_SENSOR.fields_by_name['fused'].message_type = _SENSOR_FUSED
_SENSOR.fields_by_name['primary_radar'].message_type = _SENSOR_PRIMARYRADAR
_SENSOR.fields_by_name['secondary_radar'].message_type = _SENSOR_SECONDARYRADAR
_SENSOR.fields_by_name['adsb'].message_type = _SENSOR_ADSB
_SENSOR.fields_by_name['gnss'].message_type = _SENSOR_GNSS
_SENSOR.fields_by_name['radio_frequency'].message_type = _SENSOR_RADIOFREQUENCY
_SENSOR.fields_by_name['electro_optical'].message_type = _SENSOR_ELECTROOPTICAL
_SENSOR.oneofs_by_name['details'].fields.append(
  _SENSOR.fields_by_name['fused'])
_SENSOR.fields_by_name['fused'].containing_oneof = _SENSOR.oneofs_by_name['details']
_SENSOR.oneofs_by_name['details'].fields.append(
  _SENSOR.fields_by_name['primary_radar'])
_SENSOR.fields_by_name['primary_radar'].containing_oneof = _SENSOR.oneofs_by_name['details']
_SENSOR.oneofs_by_name['details'].fields.append(
  _SENSOR.fields_by_name['secondary_radar'])
_SENSOR.fields_by_name['secondary_radar'].containing_oneof = _SENSOR.oneofs_by_name['details']
_SENSOR.oneofs_by_name['details'].fields.append(
  _SENSOR.fields_by_name['adsb'])
_SENSOR.fields_by_name['adsb'].containing_oneof = _SENSOR.oneofs_by_name['details']
_SENSOR.oneofs_by_name['details'].fields.append(
  _SENSOR.fields_by_name['gnss'])
_SENSOR.fields_by_name['gnss'].containing_oneof = _SENSOR.oneofs_by_name['details']
_SENSOR.oneofs_by_name['details'].fields.append(
  _SENSOR.fields_by_name['radio_frequency'])
_SENSOR.fields_by_name['radio_frequency'].containing_oneof = _SENSOR.oneofs_by_name['details']
_SENSOR.oneofs_by_name['details'].fields.append(
  _SENSOR.fields_by_name['electro_optical'])
_SENSOR.fields_by_name['electro_optical'].containing_oneof = _SENSOR.oneofs_by_name['details']
DESCRIPTOR.message_types_by_name['Sensor'] = _SENSOR
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Sensor = _reflection.GeneratedProtocolMessageType('Sensor', (_message.Message,), {

  'Fused' : _reflection.GeneratedProtocolMessageType('Fused', (_message.Message,), {
    'DESCRIPTOR' : _SENSOR_FUSED,
    '__module__' : 'tracking.sensors_pb2'
    # @@protoc_insertion_point(class_scope:tracking.Sensor.Fused)
    })
  ,

  'PrimaryRadar' : _reflection.GeneratedProtocolMessageType('PrimaryRadar', (_message.Message,), {
    'DESCRIPTOR' : _SENSOR_PRIMARYRADAR,
    '__module__' : 'tracking.sensors_pb2'
    # @@protoc_insertion_point(class_scope:tracking.Sensor.PrimaryRadar)
    })
  ,

  'SecondaryRadar' : _reflection.GeneratedProtocolMessageType('SecondaryRadar', (_message.Message,), {
    'DESCRIPTOR' : _SENSOR_SECONDARYRADAR,
    '__module__' : 'tracking.sensors_pb2'
    # @@protoc_insertion_point(class_scope:tracking.Sensor.SecondaryRadar)
    })
  ,

  'Adsb' : _reflection.GeneratedProtocolMessageType('Adsb', (_message.Message,), {
    'DESCRIPTOR' : _SENSOR_ADSB,
    '__module__' : 'tracking.sensors_pb2'
    # @@protoc_insertion_point(class_scope:tracking.Sensor.Adsb)
    })
  ,

  'RadioFrequency' : _reflection.GeneratedProtocolMessageType('RadioFrequency', (_message.Message,), {
    'DESCRIPTOR' : _SENSOR_RADIOFREQUENCY,
    '__module__' : 'tracking.sensors_pb2'
    # @@protoc_insertion_point(class_scope:tracking.Sensor.RadioFrequency)
    })
  ,

  'Gnss' : _reflection.GeneratedProtocolMessageType('Gnss', (_message.Message,), {
    'DESCRIPTOR' : _SENSOR_GNSS,
    '__module__' : 'tracking.sensors_pb2'
    # @@protoc_insertion_point(class_scope:tracking.Sensor.Gnss)
    })
  ,

  'ElectroOptical' : _reflection.GeneratedProtocolMessageType('ElectroOptical', (_message.Message,), {
    'DESCRIPTOR' : _SENSOR_ELECTROOPTICAL,
    '__module__' : 'tracking.sensors_pb2'
    # @@protoc_insertion_point(class_scope:tracking.Sensor.ElectroOptical)
    })
  ,
  'DESCRIPTOR' : _SENSOR,
  '__module__' : 'tracking.sensors_pb2'
  # @@protoc_insertion_point(class_scope:tracking.Sensor)
  })
_sym_db.RegisterMessage(Sensor)
_sym_db.RegisterMessage(Sensor.Fused)
_sym_db.RegisterMessage(Sensor.PrimaryRadar)
_sym_db.RegisterMessage(Sensor.SecondaryRadar)
_sym_db.RegisterMessage(Sensor.Adsb)
_sym_db.RegisterMessage(Sensor.RadioFrequency)
_sym_db.RegisterMessage(Sensor.Gnss)
_sym_db.RegisterMessage(Sensor.ElectroOptical)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
