# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ids/ids.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='ids/ids.proto',
  package='ids',
  syntax='proto3',
  serialized_options=_b('Z\'github.com/airmap/interfaces/src/go/ids'),
  serialized_pb=_b('\n\rids/ids.proto\x12\x03ids\"\x1e\n\tOperation\x12\x11\n\tas_string\x18\x01 \x01(\t\"\x18\n\x03USP\x12\x11\n\tas_string\x18\x01 \x01(\tB)Z\'github.com/airmap/interfaces/src/go/idsb\x06proto3')
)




_OPERATION = _descriptor.Descriptor(
  name='Operation',
  full_name='ids.Operation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='as_string', full_name='ids.Operation.as_string', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=22,
  serialized_end=52,
)


_USP = _descriptor.Descriptor(
  name='USP',
  full_name='ids.USP',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='as_string', full_name='ids.USP.as_string', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=54,
  serialized_end=78,
)

DESCRIPTOR.message_types_by_name['Operation'] = _OPERATION
DESCRIPTOR.message_types_by_name['USP'] = _USP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Operation = _reflection.GeneratedProtocolMessageType('Operation', (_message.Message,), {
  'DESCRIPTOR' : _OPERATION,
  '__module__' : 'ids.ids_pb2'
  # @@protoc_insertion_point(class_scope:ids.Operation)
  })
_sym_db.RegisterMessage(Operation)

USP = _reflection.GeneratedProtocolMessageType('USP', (_message.Message,), {
  'DESCRIPTOR' : _USP,
  '__module__' : 'ids.ids_pb2'
  # @@protoc_insertion_point(class_scope:ids.USP)
  })
_sym_db.RegisterMessage(USP)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
