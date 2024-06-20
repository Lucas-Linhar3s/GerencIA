import 'package:flutter_modular/flutter_modular.dart';
import 'package:workspace/app/app_module.dart';
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

class CardModule extends Module {
  @override
  void binds(i) {
    i.add<ICardRepository>(CardRepositoryImpl.new);
    i.add<IFindCardUsecases>(FindCardUsecaseImpl.new);
    i.add<IUpdateCardUsecases>(UpdateCardUsecaseImpl.new);
    i.add<ICreateCardUsecase>(CreateUsecaseImpl.new);
    i.add<IDeleteCardUsecases>(DeleteCardUsecaseImpl.new);
    i.add<ICardDatasource>(CardDatasourceImpl.new);
  }

  @override
  void routes(r) {
    r.child(Modular.initialRoute, child: (ctx) => CardPage());
  }
}
