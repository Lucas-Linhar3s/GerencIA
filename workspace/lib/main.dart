import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:workspace/app/modules/cards/domain/usecases/find/i_find_card_usecases.dart';

import 'app/app_module.dart';
import 'app/app_widget.dart';

void main() {
  runApp(ModularApp(module: AppModule(), child: AppWidget()));
  final findUsecase = Modular.get<IFindCardUsecases>();
  print(findUsecase.Find(1, 1, null));
}

