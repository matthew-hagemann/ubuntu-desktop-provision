// Mocks generated by Mockito 5.4.2 from annotations
// in ubuntu_utils/test/proxy_asset_bundle_test.dart.
// Do not manually edit this file.

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:async' as _i2;
import 'dart:typed_data' as _i4;
import 'dart:ui' as _i5;

import 'package:flutter/src/services/asset_bundle.dart' as _i3;
import 'package:mockito/mockito.dart' as _i1;
import 'package:mockito/src/dummies.dart' as _i6;

import 'proxy_asset_bundle_test.dart' as _i7;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: subtype_of_sealed_class

class _FakeFuture_0<T1> extends _i1.SmartFake implements _i2.Future<T1> {
  _FakeFuture_0(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

/// A class which mocks [AssetBundle].
///
/// See the documentation for Mockito's code generation for more information.
class MockAssetBundle extends _i1.Mock implements _i3.AssetBundle {
  MockAssetBundle() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i2.Future<_i4.ByteData> load(String? key) => (super.noSuchMethod(
        Invocation.method(
          #load,
          [key],
        ),
        returnValue: _i2.Future<_i4.ByteData>.value(_i4.ByteData(0)),
      ) as _i2.Future<_i4.ByteData>);

  @override
  _i2.Future<_i5.ImmutableBuffer> loadBuffer(String? key) =>
      (super.noSuchMethod(
        Invocation.method(
          #loadBuffer,
          [key],
        ),
        returnValue: _i2.Future<_i5.ImmutableBuffer>.value(
            _i6.dummyValue<_i5.ImmutableBuffer>(
          this,
          Invocation.method(
            #loadBuffer,
            [key],
          ),
        )),
      ) as _i2.Future<_i5.ImmutableBuffer>);

  @override
  _i2.Future<String> loadString(
    String? key, {
    bool? cache = true,
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #loadString,
          [key],
          {#cache: cache},
        ),
        returnValue: _i2.Future<String>.value(''),
      ) as _i2.Future<String>);

  @override
  _i2.Future<T> loadStructuredData<T>(
    String? key,
    _i2.Future<T> Function(String)? parser,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #loadStructuredData,
          [
            key,
            parser,
          ],
        ),
        returnValue: _i7.loadMockData<T>(
          key,
          parser,
        ),
      ) as _i2.Future<T>);

  @override
  _i2.Future<T> loadStructuredBinaryData<T>(
    String? key,
    _i2.FutureOr<T> Function(_i4.ByteData)? parser,
  ) =>
      (super.noSuchMethod(
        Invocation.method(
          #loadStructuredBinaryData,
          [
            key,
            parser,
          ],
        ),
        returnValue: _i6.ifNotNull(
              _i6.dummyValueOrNull<T>(
                this,
                Invocation.method(
                  #loadStructuredBinaryData,
                  [
                    key,
                    parser,
                  ],
                ),
              ),
              (T v) => _i2.Future<T>.value(v),
            ) ??
            _FakeFuture_0<T>(
              this,
              Invocation.method(
                #loadStructuredBinaryData,
                [
                  key,
                  parser,
                ],
              ),
            ),
      ) as _i2.Future<T>);

  @override
  void evict(String? key) => super.noSuchMethod(
        Invocation.method(
          #evict,
          [key],
        ),
        returnValueForMissingStub: null,
      );

  @override
  void clear() => super.noSuchMethod(
        Invocation.method(
          #clear,
          [],
        ),
        returnValueForMissingStub: null,
      );
}
