abstract class IHomeDatasource {
  Future<List<dynamic>> listRankings(String cliente_id);
  Future<List<dynamic>> listRankingsCartoes(String cartao_id);
  Future<double> totalClientes();
  Future<double> totalCartoes(String cartao_id);
}
