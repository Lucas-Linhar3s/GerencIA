class CardEntity {
  final String? Id;
  final String Name;
  final double Limit;

  CardEntity({this.Id, required this.Name, required this.Limit});


   CardEntity copyWith({
    String? Id,
    String? Name,
    double? Limit,
  }) {
    return CardEntity(
      Id: Id ?? this.Id,
      Name: Name ?? this.Name,
      Limit: Limit ?? this.Limit,
    );
  }
}