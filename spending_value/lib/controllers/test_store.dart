import 'dart:isolate';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:record/record.dart';

class TestStore extends ChangeNotifier {
  // https://blog.codemagic.io/understanding-flutter-isolates/
  // https://blog.logrocket.com/multithreading-flutter-using-dart-isolates/

  // ReceivePort? port;
  SendPort? postport;
  // ReceivePort? responseport;
  Isolate? isolateAudioRecorder;

  void spawn() async {
    print('isolateAudioRecorder is spawned');
    ReceivePort port = ReceivePort();
    // final rootToken = RootIsolateToken.instance!;
    final audioRecorder = AudioRecorder();
    isolateAudioRecorder = await Isolate.spawn<List<dynamic>>(
        setupFn, [port.sendPort, "rootToken", audioRecorder]);
    // Link to #1
    postport = await port.first;
    ReceivePort responseport = ReceivePort();
    postport!.send(responseport.sendPort);
    await for (var msg in responseport) {
      print(msg);
    }
  }

  void kill() {
    if (isolateAudioRecorder == null) {
      print('isolateAudioRecorder is null');
      return;
    }
    isolateAudioRecorder!.kill(priority: Isolate.immediate);
    print('isolateAudioRecorder is killed');
  }

  void sendmsg(String msg) async {
    if (postport == null) {
      return;
    }
    postport!.send(msg);
  }

  static void setupFn(List<dynamic> args) async {
    SendPort sendPort = args[0];
    // RootIsolateToken token = args[1];
    AudioRecorder audioRecorder = args[2];
    // BackgroundIsolateBinaryMessenger.ensureInitialized(token);

    ReceivePort receivePort = ReceivePort();
    // Link to #1
    sendPort.send(receivePort.sendPort);
    // SendPort thispostport = await receivePort.first;
    SendPort? thispostport;
    await for (var msg in receivePort) {
      if (msg is SendPort) {
        thispostport = msg;
        continue;
      }
      switch (msg) {
        case "stop":
          if (thispostport != null) {
            thispostport.send("Ok");
          } else {
            print('thispostport is null');
          }
          break;
        default:
          if (thispostport != null) {
            thispostport.send(msg);
          } else {
            print('thispostport is null');
          }
      }
    }
  }
}
