import 'package:flutter/material.dart';
import 'package:workspace/app/modules/login/presenter/ui/widgets/card_login.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  @override
  Widget build(BuildContext context) {
    final width = MediaQuery.sizeOf(context).width;
    final height = MediaQuery.sizeOf(context).height;
    return Scaffold(
      body: Container(
        width: width,
        height: height,
        decoration: BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.topLeft,
            end: Alignment.topRight,
            colors: [Color(0xff7e2aea), Color(0xff9e60ec)],
          ),
        ),
        child: Container(
          width: width,
          height: height,
          child: SingleChildScrollView(
            child: Stack(
              children: [
                Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    SizedBox(
                      height: height < 800 ? 100 : 150,
                    ),
                    CardLogin(),
                    SizedBox(
                      height: height < 800 ? 0 : 110,
                    ),
                  ],
                ),
                // Align(
                //   alignment: AlignmentDirectional.topEnd,
                //   child: Container(
                //     width: 400,
                //     height: 100,
                //     child: AlertNotification(
                //       title: 'Falha na autenticação',
                //       body: 'Email ou senha inválidos',
                //       type: AlertNotificationType.success,
                //       showLeadingStroke: true,
                //     ),
                //   ),
                // ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
