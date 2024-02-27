import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

class PageNotFound extends StatefulWidget {
  const PageNotFound({super.key});

  @override
  State<PageNotFound> createState() => _PageNotFoundState();
}

class _PageNotFoundState extends State<PageNotFound> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color(0xffEDF3FB),
      body: SingleChildScrollView(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Image.asset(
              "assets/images/404Error.png",
              alignment: AlignmentDirectional.center,
              width: MediaQuery.sizeOf(context).width,
              height: MediaQuery.sizeOf(context).height,
            ),
          ],
        ),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerFloat,
      floatingActionButton: ElevatedButton(
        onPressed: () {
          Navigator.pushNamedAndRemoveUntil(context, "/", (route) => false);
        },
        child: Text("Voltar para o início",
            style: TextStyle(fontSize: 20, color: Colors.white)),
        style: ButtonStyle(
            elevation: MaterialStatePropertyAll(10),
            padding: MaterialStatePropertyAll(EdgeInsets.all(20)),
            backgroundColor: MaterialStateProperty.all(Color(0xff6515DD))),
      ),
    );
  }
}
