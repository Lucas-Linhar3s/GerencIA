import 'package:workspace/app/modules/home/domain/exceptions/home_exception.dart';
import 'package:workspace/app/modules/home/infrastructure/datasource/home_datasource.dart';
import 'package:workspace/app/shared/supabase/exception.dart';
import 'package:workspace/app/shared/supabase/init.dart';

class HomeDatasourceImpl implements IHomeDatasource {
  final ISupabase _supabase;

  HomeDatasourceImpl(this._supabase);
  @override
  Future<List> listRankings(String cliente_id) async {
    try {
      final response = await _supabase.selectWhere("dividas",
          "SUM(d.valor) as valor,  clientes(nome)", "cliente_id", cliente_id);
      return response;
    } on NetworkException catch (e) {
      throw HomeException(e.message, e.hint);
    }
  }

  @override
  Future<List> listRankingsCartoes(String cartao_id) async {
    try {
      final response = await _supabase.selectWhere("dividas",
          "SUM(d.valor) as valor, cartoes(nome)", "cartao_id", cartao_id);
      return response;
    } on NetworkException catch (e) {
      throw HomeException(e.message, e.hint);
    }
  }

  @override
  Future<double> totalCartoes(String cartao_id) async {
    try {
      final response = await _supabase.selectWhere(
          "dividas", "SUM(valor) as valor", "cartao_id", cartao_id);
      return response.first["valor"];
    } on NetworkException catch (e) {
      throw HomeException(e.message, e.hint);
    }
  }

  @override
  Future<double> totalClientes() async {
    try {
      final response = await _supabase.select("dividas", "SUM(valor) as valor");
      return response.first["valor"];
    } on NetworkException catch (e) {
      throw HomeException(e.message, e.hint);
    }
  }
}
