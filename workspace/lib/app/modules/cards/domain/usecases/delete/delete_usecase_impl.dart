import 'package:workspace/app/modules/cards/domain/exception/card_exception.dart';
import 'package:workspace/app/modules/cards/domain/repositories/i_card_repositories.dart';
import 'package:workspace/app/modules/cards/domain/usecases/delete/i_delete_usecase.dart';

class DeleteCardUsecaseImpl implements IDeleteCardUsecases {
  final ICardRepository _repository;

  DeleteCardUsecaseImpl({required ICardRepository repository}) : _repository = repository;
  @override
  Future<bool> Delete(String id) async {
    try {
      return await _repository.Delete(id);
    } on CardException {
      rethrow;
    }
  }
  
}