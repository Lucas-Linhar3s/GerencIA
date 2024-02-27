import 'package:workspace/app/modules/login/domain/entities/login_entity.dart';

abstract class ILoginRepository {
  Future<String> login(LoginEntity login);
  Future<bool> logout();
}
