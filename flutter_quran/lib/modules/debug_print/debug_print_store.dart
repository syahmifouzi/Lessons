import 'package:flutter/material.dart';

// enum DebugPrintPage { recording }

class DebugPrintStore extends ChangeNotifier {
  List<String> messages = [];
  bool isReversed = false;

  void print(String newMessage) {
    DateTime now = DateTime.now();
    messages.add("$now: $newMessage");
    notifyListeners();
  }

  void clearAll() {
    messages.clear();
    notifyListeners();
  }

  void toggleReverse() {
    isReversed = !isReversed;
    notifyListeners();
  }

  List<String> showMessages() {
    List<String> toShowMessages =
        isReversed ? messages.reversed.toList() : messages;
    return toShowMessages;
  }
}
