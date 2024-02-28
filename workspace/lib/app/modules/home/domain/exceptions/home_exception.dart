import 'package:workspace/app/shared/exceptions/config.dart';

class HomeException implements CustomException {
  @override
  final String message;

  @override
  final String? hint;

  HomeException(
    this.message,
    this.hint,
  );
}
