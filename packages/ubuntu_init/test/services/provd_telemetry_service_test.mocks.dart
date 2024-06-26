// Mocks generated by Mockito 5.4.4 from annotations
// in ubuntu_init/test/services/provd_telemetry_service_test.dart.
// Do not manually edit this file.

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:async' as _i4;

import 'package:mockito/mockito.dart' as _i1;
import 'package:mockito/src/dummies.dart' as _i5;
import 'package:provd_client/src/generated/telemetry.pbgrpc.dart' as _i2;
import 'package:provd_client/src/provd_telemetry_client.dart' as _i3;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: deprecated_member_use
// ignore_for_file: deprecated_member_use_from_same_package
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: subtype_of_sealed_class

class _FakeSendResponse_SendResponseType_0 extends _i1.SmartFake
    implements _i2.SendResponse_SendResponseType {
  _FakeSendResponse_SendResponseType_0(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

/// A class which mocks [ProvdTelemetryClient].
///
/// See the documentation for Mockito's code generation for more information.
class MockProvdTelemetryClient extends _i1.Mock
    implements _i3.ProvdTelemetryClient {
  MockProvdTelemetryClient() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i4.Future<String> collect() => (super.noSuchMethod(
        Invocation.method(
          #collect,
          [],
        ),
        returnValue: _i4.Future<String>.value(_i5.dummyValue<String>(
          this,
          Invocation.method(
            #collect,
            [],
          ),
        )),
      ) as _i4.Future<String>);

  @override
  _i4.Future<_i2.SendResponse_SendResponseType> collectAndSend() =>
      (super.noSuchMethod(
        Invocation.method(
          #collectAndSend,
          [],
        ),
        returnValue: _i4.Future<_i2.SendResponse_SendResponseType>.value(
            _FakeSendResponse_SendResponseType_0(
          this,
          Invocation.method(
            #collectAndSend,
            [],
          ),
        )),
      ) as _i4.Future<_i2.SendResponse_SendResponseType>);

  @override
  _i4.Future<_i2.SendResponse_SendResponseType> sendDecline() =>
      (super.noSuchMethod(
        Invocation.method(
          #sendDecline,
          [],
        ),
        returnValue: _i4.Future<_i2.SendResponse_SendResponseType>.value(
            _FakeSendResponse_SendResponseType_0(
          this,
          Invocation.method(
            #sendDecline,
            [],
          ),
        )),
      ) as _i4.Future<_i2.SendResponse_SendResponseType>);
}
