import 'dart:async';

import 'package:flutter/material.dart';

class StopwatchStore extends ChangeNotifier {
  Stopwatch stopwatch = Stopwatch();
  Timer? timer;

  void start() {
    stopwatch.start();

    // timer = Timer.periodic(const Duration(milliseconds: 30), updateTimer);
    timer = Timer.periodic(const Duration(seconds: 1), updateTimer);
  }

  void stop() {
    stopwatch.stop();
    timer?.cancel();
  }

  void reset() {
    stopwatch.reset();
    notifyListeners();
  }

  void updateTimer(Timer timer) {
    notifyListeners();
  }

  String timeNow() {
    final duration = stopwatch.elapsed;
    String twoDigits(int n) => n.toString().padLeft(2, "0");
    final hours = twoDigits(duration.inHours);
    final minutes = twoDigits(duration.inMinutes.remainder(60));
    final seconds = twoDigits(duration.inSeconds.remainder(60));
    // final milliseconds = (duration.inMilliseconds.remainder(1000) ~/ 10)
    //     .toString()
    //     .padLeft(2, "0");
    return "$hours:$minutes:$seconds";
  }
}
