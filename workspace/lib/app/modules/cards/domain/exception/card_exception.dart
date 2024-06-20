class CardException implements Exception {
  final String message;

  final StackTrace? stackTrace;

  CardException(this.message, this.stackTrace);
}