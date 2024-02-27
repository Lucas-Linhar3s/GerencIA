import 'package:workspace/app/modules/login/domain/entities/login_entity.dart';
import 'package:workspace/app/modules/login/domain/exceptions/login_exception.dart';
import 'package:workspace/app/modules/login/domain/repositories/i_login_repository.dart';
import 'package:workspace/app/modules/login/domain/usecases/i_usecase_login.dart';

class LoginUsecaseImpl implements ILoginUsecase {
  final ILoginRepository _repository;

  LoginUsecaseImpl(this._repository);

  @override
  Future<String> login(LoginEntity login) {
    try {
      return _repository.login(login);
    } on LoginException {
      rethrow;
    }
  }

  @override
  Future<bool> logout() async {
    try {
      await _repository.logout();
      return true;
    } on LoginException {
      rethrow;
    }
  }
}
