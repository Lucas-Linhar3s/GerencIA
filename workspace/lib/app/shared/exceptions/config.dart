// ignore_for_file: public_member_api_docs, sort_constructors_first
abstract class CustomException implements Exception {
  final String message;
  final String? hint;

  CustomException(this.message, this.hint);

  @override
  String toString() => 'CustomException(message: $message, hint: $hint)';
}
