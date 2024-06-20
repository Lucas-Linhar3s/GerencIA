import 'package:dio/dio.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:workspace/app/modules/cards/card_module.dart';
import 'package:workspace/app/modules/cards/domain/repositories/i_card_repositories.dart';
import 'package:workspace/app/modules/cards/domain/usecases/create/create_usecase_impl.dart';
import 'package:workspace/app/modules/cards/domain/usecases/create/i_create_usecase.dart';
import 'package:workspace/app/modules/cards/domain/usecases/delete/delete_usecase_impl.dart';
import 'package:workspace/app/modules/cards/domain/usecases/delete/i_delete_usecase.dart';
import 'package:workspace/app/modules/cards/domain/usecases/find/find_usecase_impl.dart';
import 'package:workspace/app/modules/cards/domain/usecases/find/i_find_card_usecases.dart';
import 'package:workspace/app/modules/cards/domain/usecases/update/i_update_usecase.dart';
import 'package:workspace/app/modules/cards/domain/usecases/update/update_usecase_impl.dart';
import 'package:workspace/app/modules/cards/external/datasource/card_datasource_impl.dart';
import 'package:workspace/app/modules/cards/infrastructure/datasource/i_card_datasource.dart';
import 'package:workspace/app/modules/cards/infrastructure/repositories/card_repository_impl.dart';
import 'package:workspace/app/modules/cards/presenter/ui/page/card_page.dart';
import 'package:workspace/app/shared/services/network/dio/dio_network_service.dart';
import 'package:workspace/app/shared/services/network/network_service.dart';

class AppModule extends Module {
  @override
  void binds(i) {
    i.add<NetworkService>(DioNetworkService.new);
    i.addInstance(Dio(BaseOptions(
      baseUrl: "https://gerencia-v05a.onrender.com",
    )));
    i.addInstance(CardModule().binds(i));
  }

  @override
  void routes(r) {
    r.module(Modular.initialRoute, module: CardModule());
  }
}
