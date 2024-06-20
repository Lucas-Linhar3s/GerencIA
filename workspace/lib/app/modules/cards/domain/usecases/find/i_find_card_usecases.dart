import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';

abstract class IFindCardUsecases {
   Future<List<CardEntity>> Find(int limit, int amount, String? search);
}