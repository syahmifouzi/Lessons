import 'package:flutter/material.dart';

class SplashStore extends ChangeNotifier {
  String uid = "";
  String username = "";

  void setUsername(String x) {
    username = x;
    notifyListeners();
  }

  void setuid(String x) {
    uid = x;
  }
}
