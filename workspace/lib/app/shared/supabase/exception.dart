// ignore_for_file: public_member_api_docs, sort_constructors_first
class NetworkException implements Exception {
  final String message;
  final String? hint;
  final String? code;
  final dynamic data;

  NetworkException(
    this.message,
    this.hint,
    this.code,
    this.data,
  );
}
