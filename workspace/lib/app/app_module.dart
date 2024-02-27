import 'package:flutter_modular/flutter_modular.dart';
import 'package:supabase_flutter/supabase_flutter.dart';
import 'package:workspace/app/modules/login/login_module.dart';
import 'package:workspace/app/shared/supabase/init.dart';
import 'package:workspace/app/shared/widget/page_not_found.dart';

class AppModule extends Module {
  @override
  void binds(i) {
    i.add<ISupabase>(SupabaseInit.new);
    i.addInstance(Supabase.instance);
  }

  @override
  void routes(r) {
    r.module(Modular.initialRoute, module: LoginModule());
    r.wildcard(
        child: (context) => PageNotFound(),
        transition: TransitionType.noTransition);
  }
}
