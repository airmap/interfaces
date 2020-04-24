# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from tracking import collector_pb2 as tracking_dot_collector__pb2


class CollectorStub(object):
  """Collector service enables the exchange of surveillance updates with a collector service.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ConnectProvider = channel.stream_stream(
        '/tracking.Collector/ConnectProvider',
        request_serializer=tracking_dot_collector__pb2.Update.FromProvider.SerializeToString,
        response_deserializer=tracking_dot_collector__pb2.Update.ToProvider.FromString,
        )
    self.ConnectProcessor = channel.stream_stream(
        '/tracking.Collector/ConnectProcessor',
        request_serializer=tracking_dot_collector__pb2.Update.FromProcessor.SerializeToString,
        response_deserializer=tracking_dot_collector__pb2.Update.ToProcessor.FromString,
        )


class CollectorServicer(object):
  """Collector service enables the exchange of surveillance updates with a collector service.
  """

  def ConnectProvider(self, request_iterator, context):
    """ConnectProvider connects a stream of updates from a provider to a collector.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ConnectProcessor(self, request_iterator, context):
    """ConnectProcessor connects a stream of updates from a collector to a processor.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_CollectorServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ConnectProvider': grpc.stream_stream_rpc_method_handler(
          servicer.ConnectProvider,
          request_deserializer=tracking_dot_collector__pb2.Update.FromProvider.FromString,
          response_serializer=tracking_dot_collector__pb2.Update.ToProvider.SerializeToString,
      ),
      'ConnectProcessor': grpc.stream_stream_rpc_method_handler(
          servicer.ConnectProcessor,
          request_deserializer=tracking_dot_collector__pb2.Update.FromProcessor.FromString,
          response_serializer=tracking_dot_collector__pb2.Update.ToProcessor.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'tracking.Collector', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))