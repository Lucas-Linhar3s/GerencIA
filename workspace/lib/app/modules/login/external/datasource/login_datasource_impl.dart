import 'package:workspace/app/modules/login/domain/entities/login_entity.dart';
import 'package:workspace/app/modules/login/domain/exceptions/login_exception.dart';
import 'package:workspace/app/modules/login/infrastructure/datasource/i_login_datasource.dart';
import 'package:workspace/app/shared/supabase/exception.dart';
import 'package:workspace/app/shared/supabase/init.dart';

class LoginDatasourceImpl implements ILoginDatasource {
  final ISupabase _supabase;

  LoginDatasourceImpl(this._supabase);
  @override
  Future<String> login(LoginEntity login) async {
    try {
      final response = await _supabase.login(
        login.Email,
        login.Password,
      );

      return response.session!.accessToken;
    } on NetworkException catch (e) {
      throw LoginException(e.message, e.hint);
    }
  }

  @override
  Future logout() async {
    try {
      await _supabase.logout();
      return true;
    } on NetworkException catch (e) {
      throw LoginException(e.message, e.hint);
    }
  }
}
