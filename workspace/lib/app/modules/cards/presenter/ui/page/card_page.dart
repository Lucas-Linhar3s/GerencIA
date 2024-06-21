import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:workspace/app/modules/cards/domain/exception/card_exception.dart';
import 'package:workspace/app/modules/cards/domain/usecases/find/i_find_card_usecases.dart';

class CardPage extends StatefulWidget {
  const CardPage({super.key});

  @override
  State<CardPage> createState() => _CardPageState();
}

class _CardPageState extends State<CardPage> {
  final findUsecase = Modular.get<IFindCardUsecases>();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      floatingActionButton: FloatingActionButton(
        onPressed: () async {
          try {
            final result = await findUsecase.Find(10, 0, "");
            print(result);
          } on CardException catch (e) {
            print(e.message);
          }
        },
      ),
    );
  }
}
