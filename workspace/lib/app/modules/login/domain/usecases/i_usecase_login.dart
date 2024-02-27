import 'package:workspace/app/modules/login/domain/entities/login_entity.dart';

abstract class ILoginUsecase {
  Future<String> login(LoginEntity login);
  Future<bool> logout();
}
