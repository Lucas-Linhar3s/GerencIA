import 'package:flutter/material.dart';
import 'package:workspace/app/shared/widget/custom_text_form_field.dart';

class CardLogin extends StatefulWidget {
  const CardLogin({super.key});

  @override
  State<CardLogin> createState() => _CardLoginState();
}

class _CardLoginState extends State<CardLogin> {
  @override
  Widget build(BuildContext context) {
    final screenSize = MediaQuery.of(context).size;
    final screenWidth = screenSize.width;
    return Center(
      child: Container(
        color: Colors.white,
        width: 800,
        height: 600,
        child: Column(
          children: [
            Image.asset(
              "assets/images/logo.png",
              height: 150,
              width: 300,
            ),
            Text(
              "Login",
              style: TextStyle(fontSize: 30, fontWeight: FontWeight.bold),
            ),
            SizedBox(
              height: 30,
            ),
            CustomSearch(
              preffix: Icons.email_outlined,
              hintTextSearch: "E-mail",
              onChanged: (p0) {},
            ),
            SizedBox(
              height: 10,
            ),
            CustomSearch(
              obscure: true,
              preffix: Icons.lock_outlined,
              hintTextSearch: "Senha",
              onChanged: (p0) {},
            ),
            SizedBox(
              height: 30,
            ),
            Container(
              width: screenWidth < 600 ? 460 : 500,
              child: ElevatedButton(
                onPressed: () {},
                child: Text(
                  "Entrar",
                  style: TextStyle(fontSize: 20, color: Colors.white),
                ),
                style: ButtonStyle(
                  elevation: MaterialStatePropertyAll(10),
                  backgroundColor: MaterialStatePropertyAll(Color(0xff7e2aea)),
                  padding: MaterialStatePropertyAll(EdgeInsets.all(20)),
                ),
              ),
            )
          ],
        ),
      ),
    );
  }
}
