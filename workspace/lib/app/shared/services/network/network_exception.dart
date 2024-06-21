class NetworkException implements Exception {
  final String message;
  final Map<String, dynamic>? logger;
  final int? code;
  final dynamic data;

  NetworkException(
    this.message,
    this.logger,
    this.code,
    this.data,
  );
}
