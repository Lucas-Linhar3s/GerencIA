import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';

abstract class ICreateCardUsecase {
   Future<bool> Create(CardEntity card);
}