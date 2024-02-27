import 'package:workspace/app/modules/login/domain/entities/login_entity.dart';

abstract class ILoginDatasource {
  Future<String> login(LoginEntity login);
  Future logout();
}
