import 'package:supabase_flutter/supabase_flutter.dart';
import 'package:workspace/app/shared/supabase/exception.dart';

abstract class ISupabase {
  Future<AuthResponse> login(String email, String password);
  Future logout();
  PostgrestFilterBuilder<PostgrestList> select(String table, String columns);
  PostgrestFilterBuilder<List<Map<String, dynamic>>> selectWhere(
      String table, String columns, String columnWhere, String valueWhere);
  PostgrestFilterBuilder<dynamic> create(String table, Object values);
  PostgrestFilterBuilder<dynamic> update(
      String table, String id, Map<String, dynamic> values);
  PostgrestFilterBuilder<dynamic> delete(String table, String id);
}

class SupabaseInit implements ISupabase {
  final Supabase _supabase;

  SupabaseInit(this._supabase);

  @override
  Future<AuthResponse> login(String email, String password) async {
    try {
      final response = await _supabase.client.auth
          .signInWithPassword(password: password, email: email);
      return response;
    } on AuthException catch (e) {
      throw NetworkException(
        e.message,
        "",
        e.statusCode,
        "",
      );
    } catch (e) {
      throw NetworkException(
        e.toString(),
        "",
        "",
        "",
      );
    }
  }

  @override
  PostgrestFilterBuilder<PostgrestList> select(String table, String columns) {
    try {
      final se = _supabase.client.from(table).select(columns);
      return se;
    } on PostgrestException catch (e) {
      throw NetworkException(
        e.message,
        e.hint,
        e.code,
        e.details,
      );
    } catch (e) {
      throw NetworkException(
        e.toString(),
        "",
        "",
        "",
      );
    }
  }

  @override
  PostgrestFilterBuilder<dynamic> create(String table, Object values) {
    try {
      return _supabase.client.from(table).insert(values);
    } on PostgrestException catch (e) {
      throw NetworkException(
        e.message,
        e.hint,
        e.code,
        e.details,
      );
    } catch (e) {
      throw NetworkException(
        e.toString(),
        "",
        "",
        "",
      );
    }
  }

  @override
  PostgrestFilterBuilder<dynamic> delete(String table, String id) {
    try {
      return _supabase.client.from(table).delete().eq("id", id);
    } on PostgrestException catch (e) {
      throw NetworkException(
        e.message,
        e.hint,
        e.code,
        e.details,
      );
    } catch (e) {
      throw NetworkException(
        e.toString(),
        "",
        "",
        "",
      );
    }
  }

  @override
  Future logout() async {
    try {
      await _supabase.client.auth.signOut();
    } on AuthException catch (e) {
      throw NetworkException(
        e.message,
        "",
        e.statusCode,
        "",
      );
    } catch (e) {
      throw NetworkException(
        e.toString(),
        "",
        "",
        "",
      );
    }
  }

  @override
  PostgrestFilterBuilder update(
      String table, String id, Map<String, dynamic> values) {
    try {
      return _supabase.client.from(table).update(values).eq("id", id);
    } on AuthException catch (e) {
      throw NetworkException(
        e.message,
        "",
        e.statusCode,
        "",
      );
    } catch (e) {
      throw NetworkException(
        e.toString(),
        "",
        "",
        "",
      );
    }
  }

  @override
  PostgrestFilterBuilder<List<Map<String, dynamic>>> selectWhere(
      String table, String columns, String columnWhere, String valueWhere) {
    try {
      final se = _supabase.client
          .from(table)
          .select(columns)
          .eq(columnWhere, valueWhere);
      return se;
    } on PostgrestException catch (e) {
      throw NetworkException(
        e.message,
        e.hint,
        e.code,
        e.details,
      );
    } catch (e) {
      throw NetworkException(
        e.toString(),
        "",
        "",
        "",
      );
    }
  }
}
