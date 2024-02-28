// ignore_for_file: public_member_api_docs, sort_constructors_first, non_constant_identifier_names

class HomeEntity {
  double TotalClientes;
  double TotalCartoes;
  String Nome;
  double Total;
  int Porcentagem;
  HomeEntity({
    required this.TotalClientes,
    required this.TotalCartoes,
    required this.Nome,
    required this.Total,
    required this.Porcentagem,
  });

  HomeEntity copyWith({
    double? TotalClientes,
    double? TotalCartoes,
    String? Nome,
    double? Total,
    int? Porcentagem,
  }) {
    return HomeEntity(
      TotalClientes: TotalClientes ?? this.TotalClientes,
      TotalCartoes: TotalCartoes ?? this.TotalCartoes,
      Nome: Nome ?? this.Nome,
      Total: Total ?? this.Total,
      Porcentagem: Porcentagem ?? this.Porcentagem,
    );
  }
}
