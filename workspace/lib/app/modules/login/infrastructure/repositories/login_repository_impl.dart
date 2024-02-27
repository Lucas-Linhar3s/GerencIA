import 'package:workspace/app/modules/login/domain/entities/login_entity.dart';
import 'package:workspace/app/modules/login/domain/exceptions/login_exception.dart';
import 'package:workspace/app/modules/login/domain/repositories/i_login_repository.dart';
import 'package:workspace/app/modules/login/infrastructure/datasource/i_login_datasource.dart';

class LoginRepositoryImpl implements ILoginRepository {
  final ILoginDatasource _datasource;

  LoginRepositoryImpl(this._datasource);

  @override
  Future<String> login(LoginEntity login) {
    try {
      return _datasource.login(login);
    } on LoginException {
      rethrow;
    }
  }

  @override
  Future<bool> logout() async {
    try {
      await _datasource.logout();
      return true;
    } on LoginException {
      rethrow;
    }
  }
}
