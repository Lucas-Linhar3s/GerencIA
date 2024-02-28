import 'package:flutter_modular/flutter_modular.dart';
import 'package:mobx/mobx.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:workspace/app/modules/login/domain/entities/login_entity.dart';
import 'package:workspace/app/modules/login/domain/exceptions/login_exception.dart';
import 'package:workspace/app/modules/login/domain/usecases/i_usecase_login.dart';
part 'login_store.g.dart';

class LoginStore = _LoginStoreBase with _$LoginStore;

abstract class _LoginStoreBase with Store {
  final loginUsecase = Modular.get<ILoginUsecase>();

  @observable
  var loginModel = LoginEntity(
    Email: '',
    Password: '',
  );

  @observable
  String errorMessage = '';

  @observable
  bool obscureText = true;
  @action
  bool setObscureText(bool value) {
    obscureText = value;
    return obscureText;
  }

  @observable
  var token = '';

  @observable
  var email = '';
  @action
  void setEmail(String value) {
    email = value;
    loginModel.Email = email;
  }

  @observable
  var senha = '';
  @action
  void setSenha(String value) {
    senha = value;
    loginModel.Password = senha;
  }

  @action
  bool validateEmail(String value) {
    if (value.isEmpty) {
      error.email = 'E-mail é obrigatório';
      return false;
    } else if (RegExp(
                r"^[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+")
            .hasMatch(email) ==
        false) {
      error.email = 'Esse e-mail não é válido';
      return false;
    }
    error.email = null;
    return true;
  }

  @action
  bool validatePassword(String value) {
    if (value.isEmpty) {
      error.senha = 'Senha é obrigatório';
      return false;
    } else if (value.length < 6) {
      error.senha = 'Senha muito curta';
      return false;
    }
    error.senha = null;
    return true;
  }

  @action
  Future<bool> Login() async {
    var prefis = await SharedPreferences.getInstance();
    try {
      var result = await loginUsecase.login(loginModel);
      return prefis.setString("jwt", result);
    } on LoginException catch (e) {
      errorMessage = e.message;
      return false;
    }
  }

  @action
  Future<bool> Logout() async {
    var prefis = await SharedPreferences.getInstance();

    try {
      await loginUsecase.logout();
      prefis.remove("jwt");
      return true;
    } on LoginException catch (e) {
      errorMessage = e.message;
      return false;
    }
  }

  @action
  Future<String> getToken() async {
    final prefs = await SharedPreferences.getInstance();
    String stk = prefs.getString('jwt') ?? '';
    return stk;
  }

  Future<bool> isLoggedIn() async {
    String token = await getToken();
    if (token.isEmpty) {
      return false;
    } else {
      return true;
    }
  }

  final FormErrorState error = FormErrorState();

  late List<ReactionDisposer> _disposers;

  void setupValidations() {
    _disposers = [
      reaction((_) => email, validateEmail),
      reaction((_) => senha, validatePassword),
    ];
  }

  void dispose() {
    for (final d in _disposers) {
      d();
    }
  }

  bool validateAll() {
    if (validateEmail(email) && validatePassword(senha)) {
      return true;
    } else {
      return false;
    }
  }
}

class FormErrorState = _FormErrorState with _$FormErrorState;

abstract class _FormErrorState with Store {
  @observable
  String? email;

  @observable
  String? senha;

  @computed
  bool get hasErrors => email != null || senha != null;
}
