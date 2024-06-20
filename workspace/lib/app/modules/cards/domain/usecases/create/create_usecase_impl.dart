import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';
import 'package:workspace/app/modules/cards/domain/exception/card_exception.dart';
import 'package:workspace/app/modules/cards/domain/repositories/i_card_repositories.dart';
import 'package:workspace/app/modules/cards/domain/usecases/create/i_create_usecase.dart';

class CreateUsecaseImpl implements ICreateCardUsecase {
  final ICardRepository _repository;

  CreateUsecaseImpl({required ICardRepository repository}) : _repository = repository;
  @override
  Future<bool> Create(CardEntity card) async {
    try {
      return await _repository.Create(card);
    } on CardException {
      rethrow;
    }
  }
  
}