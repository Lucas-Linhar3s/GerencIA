// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'package:workspace/app/modules/home/domain/entities/home_entity.dart';
import 'package:workspace/app/modules/home/domain/exceptions/home_exception.dart';
import 'package:workspace/app/modules/home/domain/repositories/home_repository.dart';
import 'package:workspace/app/modules/home/domain/usecases/i_home_usecase.dart';

class HomeUsecaseImpl implements IHomeUsecase {
  final IHomeRepository _repository;
  HomeUsecaseImpl(this._repository);
  @override
  Future<List<HomeEntity>> listRankings(String cliente_id) {
    try {
      return _repository.listRankings(cliente_id);
    } on HomeException {
      rethrow;
    }
  }

  @override
  Future<List<HomeEntity>> listRankingsCartoes(String cartao_id) {
    try {
      return _repository.listRankingsCartoes(cartao_id);
    } on HomeException {
      rethrow;
    }
  }

  @override
  Future<double> totalCartoes(String cartao_id) {
    try {
      return _repository.totalCartoes(cartao_id);
    } on HomeException {
      rethrow;
    }
  }

  @override
  Future<double> totalClientes() {
    try {
      return _repository.totalClientes();
    } on HomeException {
      rethrow;
    }
  }
}
