import 'package:flutter_modular/flutter_modular.dart';
import 'package:supabase_flutter/supabase_flutter.dart';
import 'package:workspace/app/modules/login/domain/repositories/i_login_repository.dart';
import 'package:workspace/app/modules/login/domain/usecases/i_usecase_login.dart';
import 'package:workspace/app/modules/login/domain/usecases/usecase_login_impl.dart';
import 'package:workspace/app/modules/login/external/datasource/login_datasource_impl.dart';
import 'package:workspace/app/modules/login/infrastructure/datasource/i_login_datasource.dart';
import 'package:workspace/app/modules/login/infrastructure/repositories/login_repository_impl.dart';
import 'package:workspace/app/modules/login/presenter/store/login_store.dart';
import 'package:workspace/app/modules/login/presenter/ui/pages/login_page.dart';
import 'package:workspace/app/shared/supabase/init.dart';

class LoginModule extends Module {
  @override
  void binds(Injector i) {
    i.add<ILoginRepository>(LoginRepositoryImpl.new);
    i.add<ILoginUsecase>(LoginUsecaseImpl.new);
    i.add<ILoginDatasource>(LoginDatasourceImpl.new);
    i.add(LoginStore.new);
    i.add<ISupabase>(SupabaseInit.new);
    i.addInstance(Supabase.instance);
  }

  @override
  void routes(RouteManager r) {
    r.child("/", child: ((context) => LoginPage()));
  }
}
