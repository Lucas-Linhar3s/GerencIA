import 'package:workspace/app/modules/home/domain/entities/home_entity.dart';
import 'package:workspace/app/modules/home/domain/exceptions/home_exception.dart';
import 'package:workspace/app/modules/home/domain/repositories/home_repository.dart';
import 'package:workspace/app/modules/home/infrastructure/adapters/home_adapter.dart';
import 'package:workspace/app/modules/home/infrastructure/datasource/home_datasource.dart';

class HomeRepositoryImpl implements IHomeRepository {
  final IHomeDatasource _datasource;

  HomeRepositoryImpl(this._datasource);
  @override
  Future<List<HomeEntity>> listRankings(String cliente_id) async {
    try {
      final result = await _datasource.listRankings(cliente_id);
      return HomeAdapter.listRankings(result);
    } on HomeException {
      rethrow;
    }
  }

  @override
  Future<List<HomeEntity>> listRankingsCartoes(String cartao_id) async {
    try {
      final result = await _datasource.listRankingsCartoes(cartao_id);
      return HomeAdapter.listRankings(result);
    } on HomeException {
      rethrow;
    }
  }

  @override
  Future<double> totalCartoes(String cartao_id) {
    try {
      return _datasource.totalCartoes(cartao_id);
    } on HomeException {
      rethrow;
    }
  }

  @override
  Future<double> totalClientes() {
    try {
      return _datasource.totalClientes();
    } on HomeException {
      rethrow;
    }
  }
}
