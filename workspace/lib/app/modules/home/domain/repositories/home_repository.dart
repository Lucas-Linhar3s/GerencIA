import 'package:workspace/app/modules/home/domain/entities/home_entity.dart';

abstract class IHomeRepository {
  Future<List<HomeEntity>> listRankings(String cliente_id);
  Future<List<HomeEntity>> listRankingsCartoes(String cartao_id);
  Future<double> totalClientes();
  Future<double> totalCartoes(String cartao_id);
}
