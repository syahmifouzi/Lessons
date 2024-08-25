import 'package:flutter/material.dart';

enum RecordingButtonState { recording, pausing, stopped }

class RecordingButtonStore extends ChangeNotifier {
  RecordingButtonState buttonState = RecordingButtonState.stopped;

  void setButtonState(RecordingButtonState newButtonState) {
    buttonState = newButtonState;
    notifyListeners();
  }
}
