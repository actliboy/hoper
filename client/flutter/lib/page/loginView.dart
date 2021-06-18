import 'package:app/model/state/user.dart';
import 'package:app/model/user.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class LoginView extends StatelessWidget {
  const LoginView({Key? key}) : super(key: key);

  User login(String account,password){
    final api = "/user/login";
    print('$account,$password,$api');
    return User()..name="贾一饼";
  }

  @override
  Widget build(BuildContext context) {
    var _account = '';
    var _password = '';
    final _formKey = GlobalKey<FormState>();
    return Scaffold(
        resizeToAvoidBottomInset: false,
        body: Center(
          child: Container(
            padding: EdgeInsets.all(60.0),
            child: Form(
              key: _formKey,
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    'Welcome',
                    style: Theme.of(context).textTheme.headline2,
                  ),
                  TextFormField(
                    decoration: InputDecoration(
                      hintText: '邮箱/手机',
                    ),
                    onSaved: (value) {
                      _account = value!;
                    },
                  ),
                  TextFormField(
                    decoration: InputDecoration(
                      hintText: '密码',
                    ),
                    onSaved: (value) {
                      _password = value!;
                    },
                    obscureText: true,
                  ),
                  SizedBox(
                    height: 24,
                  ),
                  ElevatedButton(
                    style: ButtonStyle(
                        foregroundColor:ButtonStyleButton.allOrNull<Color>(Colors.yellow)
                    ),
                    child: Text('登录'),
                    onPressed: () {
                      var _state = _formKey.currentState;
                      _state!.save();
                      print(_account);
                      final user =  login(_account, _password);
                      final AuthState userState =  Get.find();
                      userState.user.value = user;
                      Navigator.pop(context);
                    },
                  ),
                ],
              ),
            ),
          ),
        ),
    );
  }
}
