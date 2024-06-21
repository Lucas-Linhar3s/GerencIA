import 'package:workspace/app/modules/cards/domain/entities/card_entity.dart';

class CardAdapters extends CardEntity {
  CardAdapters({required super.Id, required super.Name, required super.Limit});

  static CardEntity fromJson(Map<String, dynamic> json) {
    return CardAdapters(
      Id: json['id'],
      Name: json['nome'],
      Limit: json['limite'],
    );
  }

  static Map<String, dynamic> toJson(CardEntity card) {
    return {
      'id': card.Id,
      'nome': card.Name,
      'limite': card.Limit,
    };
  }

  static List<CardEntity> fromListToJsonCards(json) {
    final castedList = List.castFrom<dynamic, Map<String, dynamic>>(json);
    final parsedList =
        castedList.map((cardsMap) => CardAdapters.fromJson(cardsMap)).toList();
    return parsedList;
  }
}
