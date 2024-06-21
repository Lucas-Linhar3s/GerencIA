import 'dart:convert';

import 'package:crypto/crypto.dart';
import 'package:dio/dio.dart';
import 'package:uuid/v4.dart';

class ConfigDio {
  // Variáveis privadas
  static final String _appKey = "5f1f5d5f-6b9e-4f6f-a6b6-5d5f6f6f6f6f";
  static final String _nonce = UuidV4().generate();
  static final String _timestamp =
      DateTime.now().millisecondsSinceEpoch.toString();
  static final String _appVersion = "1.0.0";
  static final String _sign = _generateMd5("AppKey" +
          _appKey +
          "AppVersion" +
          _appVersion +
          "Nonce" +
          _nonce +
          "Timestamp" +
          _timestamp +
          _appKey)
      .toUpperCase();

  // Método privado para gerar MD5
  static String _generateMd5(String input) {
    var bytes = utf8.encode(input);
    var digest = md5.convert(bytes);
    return digest.toString();
  }

  static Dio get() {
    return Dio(BaseOptions(
      baseUrl: "http://localhost:8000",
      headers: {
        "Nonce": _nonce,
        "Timestamp": _timestamp,
        "Sign": _sign,
        "App-Version": _appVersion,
      },
    ));
  }
}
