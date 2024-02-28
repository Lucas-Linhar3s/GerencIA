import 'package:workspace/app/modules/home/domain/entities/home_entity.dart';

class HomeAdapter {
  HomeAdapter._();

  static List<HomeEntity> listRankings(List<dynamic> list) {
    return list
        .map((e) => HomeEntity(
              TotalClientes: e["total_clientes"] ?? 0,
              TotalCartoes: e["total_cartoes"] ?? 0,
              Nome: e["nome"] ?? '',
              Total: e["total"] ?? 0,
              Porcentagem: e["porcentagem"] ?? 0,
            ))
        .toList();
  }

  static Map<String, dynamic> toMap(HomeEntity model) {
    return {
      'total_clientes': model.TotalClientes,
      'total_cartoes': model.TotalCartoes,
      'nome': model.Nome,
      'total': model.Total,
      'porcentagem': model.Porcentagem,
    };
  }

  static HomeEntity fromMap(Map<String, dynamic> map) {
    return HomeEntity(
      TotalClientes: map['total_clientes'],
      TotalCartoes: map['total_cartoes'],
      Nome: map['nome'],
      Total: map['total'],
      Porcentagem: map['porcentagem'],
    );
  }
}
