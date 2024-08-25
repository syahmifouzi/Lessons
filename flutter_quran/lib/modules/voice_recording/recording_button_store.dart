import 'package:flutter/material.dart';

enum RecordingButtonState { initial, recording, pausing, stopped }

class RecordingButtonStore extends ChangeNotifier {
  RecordingButtonState buttonState = RecordingButtonState.initial;

  void setButtonState(RecordingButtonState newButtonState) {
    buttonState = newButtonState;
    notifyListeners();
  }
}
