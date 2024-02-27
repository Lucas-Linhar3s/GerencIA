// ignore_for_file: public_member_api_docs, sort_constructors_first
class LoginEntity {
  String Email;
  String Password;
  LoginEntity({
    required this.Email,
    required this.Password,
  });

  LoginEntity copyWith({
    String? Email,
    String? Password,
  }) {
    return LoginEntity(
      Email: Email ?? this.Email,
      Password: Password ?? this.Password,
    );
  }
}
