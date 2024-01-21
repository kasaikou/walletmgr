import service_internal_pb2 as service__internal__pb2
from _typeshed import Incomplete

class InternalAppServiceStub:
    def __init__(self, channel) -> None:
        """Constructor.

                Args:
                    channel: A grpc.Channel.
        """

class InternalAppServiceServicer:
    def ListTransactions(self, request, context):
        """Missing associated documentation comment in .proto file."""
def add_InternalAppServiceServicer_to_server(servicer, server): ...

class InternalAppService:
    @staticmethod
    def ListTransactions(request, target, options: tuple = ..., channel_credentials: Incomplete | None = ..., call_credentials: Incomplete | None = ..., insecure: bool = ..., compression: Incomplete | None = ..., wait_for_ready: Incomplete | None = ..., timeout: Incomplete | None = ..., metadata: Incomplete | None = ...): ...
