import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';
import 'package:workspace/app/modules/cards/domain/exception/card_exception.dart';
import 'package:workspace/app/modules/cards/domain/repositories/i_card_repositories.dart';
import 'package:workspace/app/modules/cards/infrastructure/adapters/card_adapeters.dart';
import 'package:workspace/app/modules/cards/infrastructure/datasource/i_card_datasource.dart';

class CardRepositoryImpl implements ICardRepository {
  final ICardDatasource _datasource;

  CardRepositoryImpl({required ICardDatasource datasource}) : _datasource = datasource;
  @override
  Future<bool> Create(CardEntity card) async {
    try {
      final json = CardAdapters.toJson(card);
      await _datasource.Create(json);
      return true;
    } on CardException {
      rethrow;
    }
  }

  @override
  Future<bool> Delete(String id) async {
    try {
      await _datasource.Delete(id);
      return true;
    } on CardException {
      rethrow;
    }
  }

  @override
  Future<List<CardEntity>> Find(int limit, int amount, String? search) async {
    try {
       final response = await _datasource.Find(limit, amount, search);
      final result = CardAdapters.fromListToJsonCards(response);
      return result;
    } on CardException {
      rethrow;
    }
  }

  @override
  Future<bool> Update(CardEntity card) async {
     try {
      final json = CardAdapters.toJson(card);
      await _datasource.Create(json);
      return true;
    } on CardException {
      rethrow;
    }
  }
  
}