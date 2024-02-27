// ignore_for_file: public_member_api_docs, sort_constructors_first

import 'package:workspace/app/shared/exceptions/config.dart';

class LoginException implements CustomException {
  @override
  final String message;

  @override
  final String? hint;

  LoginException(
    this.message,
    this.hint,
  );
}
