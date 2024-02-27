import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:supabase_flutter/supabase_flutter.dart';

import 'app/app_module.dart';
import 'app/app_widget.dart';

void main() {
  Supabase.initialize(
    url: 'https://nohewrlfkhparfzzeyjc.supabase.co',
    anonKey:
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im5vaGV3cmxma2hwYXJmenpleWpjIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDg5NjUyOTQsImV4cCI6MjAyNDU0MTI5NH0.8FZfHsbhO-6rxp-3lwU7DMdnaUCinyTA0CQGcC_ZIhU",
  );

  runApp(ModularApp(module: AppModule(), child: AppWidget()));
}
