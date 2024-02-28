// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_mobx/flutter_mobx.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:workspace/app/modules/login/presenter/store/login_store.dart';
import 'package:workspace/app/shared/widget/custom_text_form_field.dart';

class CardLogin extends StatefulWidget {
  LoginStore login;
  CardLogin({
    Key? key,
    required this.login,
  }) : super(key: key);

  @override
  State<CardLogin> createState() => _CardLoginState();
}

class _CardLoginState extends State<CardLogin> {
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();

  @override
  void setState(VoidCallback fn) {
    widget.login.setupValidations();
    super.setState(fn);
  }

  @override
  void dispose() {
    widget.login.dispose();
    super.dispose();
  }

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
            Observer(
              builder: (_) {
                return CustomSearch(
                  searchController: emailController,
                  preffix: Icons.email_outlined,
                  hintTextSearch: "E-mail",
                  onChanged: widget.login.setEmail,
                );
              },
            ),
            SizedBox(
              height: 10,
            ),
            Observer(
              builder: (_) {
                return CustomSearch(
                  errorText: widget.login.error.senha,
                  eyeIcon: widget.login.obscureText
                      ? Icons.remove_red_eye_outlined
                      : Icons.visibility_off_outlined,
                  iconOnPressed: () {
                    widget.login.setObscureText(
                      !widget.login.obscureText,
                    );
                  },
                  searchController: passwordController,
                  obscure: widget.login.obscureText,
                  preffix: Icons.lock_outlined,
                  hintTextSearch: "Senha",
                  onChanged: widget.login.setSenha,
                );
              },
            ),
            SizedBox(
              height: 30,
            ),
            Container(
                width: screenWidth < 600 ? 460 : 500,
                child: Observer(
                  builder: (_) {
                    return ElevatedButton(
                      onPressed: () async {
                        bool validate = widget.login.validateAll();
                        if (validate) {
                          var success = await widget.login.Login();
                          Timer.periodic(Duration(seconds: 5), (timer) {
                            widget.login.errorMessage = '';
                            timer.cancel();
                          });
                          if (success) {
                            Modular.to.pushReplacementNamed("/home/");
                          }
                        }
                      },
                      child: Text(
                        "Entrar",
                        style: TextStyle(fontSize: 20, color: Colors.white),
                      ),
                      style: ButtonStyle(
                        elevation: MaterialStatePropertyAll(10),
                        backgroundColor:
                            MaterialStatePropertyAll(Color(0xff7e2aea)),
                        padding: MaterialStatePropertyAll(EdgeInsets.all(20)),
                      ),
                    );
                  },
                ))
          ],
        ),
      ),
    );
  }
}
