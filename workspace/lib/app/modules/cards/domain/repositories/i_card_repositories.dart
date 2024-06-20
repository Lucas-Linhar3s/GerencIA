import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';

abstract class ICardRepository {
   Future<List<CardEntity>> Find(int limit, int amount, String? search);
   Future<bool> Create(CardEntity card);
   Future<bool> Update(CardEntity card);
   Future<bool> Delete(String id);
}