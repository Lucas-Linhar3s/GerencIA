import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';

abstract class IUpdateCardUsecases {
   Future<bool> Update(CardEntity card);
}