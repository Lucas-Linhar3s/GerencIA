import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';
import 'package:workspace/app/modules/cards/domain/exception/card_exception.dart';
import 'package:workspace/app/modules/cards/domain/repositories/i_card_repositories.dart';
import 'package:workspace/app/modules/cards/domain/usecases/find/i_find_card_usecases.dart';

class FindCardUsecaseImpl implements IFindCardUsecases {
  final ICardRepository _repository;

  FindCardUsecaseImpl({required ICardRepository repository}) : _repository = repository;
  @override
  Future<List<CardEntity>> Find(int limit, int amount, String? search) async {
    try {
      return await _repository.Find(limit, amount, search);
    } on CardException {
      rethrow;
    }
  }
  
}