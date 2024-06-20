import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';

abstract class ICardDatasource {
   Future<Map<String, dynamic>> Find(int limit, int amount, String? search);
   Future<bool> Create(Map<String,dynamic> json);
   Future<bool> Update(Map<String,dynamic> json);
   Future<bool> Delete(String id);
}