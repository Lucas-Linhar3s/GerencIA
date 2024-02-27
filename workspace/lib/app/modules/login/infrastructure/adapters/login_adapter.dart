import 'package:workspace/app/modules/login/domain/entities/login_entity.dart';

class LoginAdapter {
  LoginAdapter._();

  static LoginEntity fromMap(Map<String, dynamic> map) {
    return LoginEntity(
      Email: map['email'],
      Password: map['password'],
    );
  }

  static Map<String, dynamic> toMap(LoginEntity model) {
    return {
      'email': model.Email,
      'password': model.Password,
    };
  }
}
