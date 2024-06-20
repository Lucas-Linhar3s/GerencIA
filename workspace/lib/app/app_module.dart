import 'package:dio/dio.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:workspace/app/modules/cards/card_module.dart';
import 'package:workspace/app/shared/services/network/dio/dio_network_service.dart';
import 'package:workspace/app/shared/services/network/network_service.dart';

class AppModule extends Module {
  @override
  void binds(i) {
    i.add(CardModule.new);
    i.add<NetworkService>(DioNetworkService.new);
    i.addInstance(Dio(BaseOptions(
      baseUrl: "http://localhost:8000",
    )));
  }

  @override
  void routes(r) {

  }
}