import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';
import 'package:workspace/app/modules/cards/domain/exception/card_exception.dart';
import 'package:workspace/app/modules/cards/domain/repositories/i_card_repositories.dart';
import 'package:workspace/app/modules/cards/domain/usecases/update/i_update_usecase.dart';

class UpdateCardUsecaseImpl implements IUpdateCardUsecases {
  final ICardRepository _repository;

  UpdateCardUsecaseImpl({required ICardRepository repository}) : _repository = repository;
  @override
  Future<bool> Update(CardEntity card) async {
    try {
      return await _repository.Update(card);
    } on CardException {
      rethrow;
    }
  }
  
}