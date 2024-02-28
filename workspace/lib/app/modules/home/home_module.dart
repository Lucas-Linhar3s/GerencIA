import 'package:flutter_modular/flutter_modular.dart';
import 'package:supabase_flutter/supabase_flutter.dart';
import 'package:workspace/app/modules/home/domain/repositories/home_repository.dart';
import 'package:workspace/app/modules/home/domain/usecases/home_usecase_impl.dart';
import 'package:workspace/app/modules/home/domain/usecases/i_home_usecase.dart';
import 'package:workspace/app/modules/home/external/datasource/home_datasource_impl.dart';
import 'package:workspace/app/modules/home/infrastructure/datasource/home_datasource.dart';
import 'package:workspace/app/modules/home/infrastructure/repositories/home_repository_impl.dart';
import 'package:workspace/app/shared/supabase/init.dart';

class HomeModule extends Module {
  @override
  void binds(Injector i) {
    i.add<IHomeDatasource>(HomeDatasourceImpl.new);
    i.add<IHomeRepository>(HomeRepositoryImpl.new);
    i.add<IHomeUsecase>(HomeUsecaseImpl.new);
    i.add<ISupabase>(SupabaseInit.new);
    i.addInstance(Supabase.instance);
  }

  @override
  void routes(RouteManager r) {}
}
