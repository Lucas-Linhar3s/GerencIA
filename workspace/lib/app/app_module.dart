import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:workspace/app/modules/login/login_module.dart';
import 'package:workspace/app/shared/widget/page_not_found.dart';

class AppModule extends Module {
  @override
  void binds(i) {}

  @override
  void routes(r) {
    r.module("/login", module: LoginModule());
    r.child(
      Modular.initialRoute,
      child: (context) => Scaffold(
        backgroundColor: Colors.red,
      ),
      children: [],
      guards: [],
      transition: TransitionType.noTransition,
    );
    r.wildcard(
        child: (context) => PageNotFound(),
        transition: TransitionType.noTransition);
  }
}
