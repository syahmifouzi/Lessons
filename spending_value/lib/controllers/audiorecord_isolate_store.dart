import 'dart:io';
import 'dart:isolate';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_storage/firebase_storage.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:just_audio/just_audio.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:record/record.dart';
import 'package:spending_value/main.dart';

class AudiorecordIsolateStore extends ChangeNotifier {
  // https://blog.codemagic.io/understanding-flutter-isolates/
  // https://blog.logrocket.com/multithreading-flutter-using-dart-isolates/

  // ReceivePort? port;
  SendPort? postport;
  // ReceivePort? responseport;
  Isolate? isolateAudioRecorder;

  RecordingState recordingState = RecordingState.standby;
  DateTime? recordingDate;
  String? recordedPath;

  String progressText = "Standby";

  void startRecording() async {
    Map<Permission, PermissionStatus> permissions =
        await [Permission.storage, Permission.microphone].request();

    bool permissionsGranted = permissions[Permission.storage]!.isGranted &&
        permissions[Permission.microphone]!.isGranted;

    if (!permissionsGranted) {
      print('Permission not granted');
      return;
    }
    recordedPath = null;
    spawn();
  }

  void stopRecording() {
    sendmsg("stop");
  }

  void spawn() async {
    print('isolateAudioRecorder is spawned');
    final rootToken = RootIsolateToken.instance!;
    final appFolder = await getTemporaryDirectory();
    ReceivePort port = ReceivePort();
    isolateAudioRecorder = await Isolate.spawn<List<dynamic>>(
        setupFn, [port.sendPort, rootToken, appFolder.path]);
    // Link to #1
    postport = await port.first;
    ReceivePort responseport = ReceivePort();
    postport!.send(responseport.sendPort);
    await for (var msg in responseport) {
      print(msg);
      // decodeMsg(msg);
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

  void decodeMsg(dynamic msg) {
    if (msg["cmd"] == "start") {
      startRecorderState();
    } else if (msg["cmd"] == "stop") {
      stopRecorderState(msg["path"], msg["datetime"]);
      kill();
    }
  }

  void startRecorderState() {
    recordingState = RecordingState.recording;
    progressText = "Now Recording...";
    notifyListeners();
  }

  void stopRecorderState(String recordedPathL, DateTime datetimeL) {
    recordingState = RecordingState.done;
    recordingDate = datetimeL;
    recordedPath = recordedPathL;
    progressText = "Recording Stopped";
    notifyListeners();
  }

  Future<int> saveRecording(
      String title, String ayatStart, String ayatEnd, String surah) async {
    if (recordingState != RecordingState.done) {
      stopRecording();
    }
    AudioPlayer audioPlayer = AudioPlayer();
    final duration = await audioPlayer.setFilePath(recordedPath!);
    final db = FirebaseFirestore.instance;
    final newIdRef = db.collection("audio").doc();
    final storageRef = FirebaseStorage.instance.ref();
    final audioStorageRef = storageRef.child("audio/${newIdRef.id}.m4a");
    File file = File(recordedPath!);
    int hasError = 0;
    try {
      await audioStorageRef.putFile(file);
      String downloadUrl = await audioStorageRef.getDownloadURL();
      final audioDB = {
        "title": title,
        "url": downloadUrl,
        "timestamp": recordingDate,
        "ayatStart": ayatStart,
        "ayatEnd": ayatEnd,
        "status": "Pending",
        "surah": surah,
        "duration": duration.toString()
      };
      newIdRef.set(audioDB);
    } on FirebaseException catch (e) {
      hasError = 1;
      print(e);
    }
    return hasError;
    // if (hasError) {
    //   printSnack("Failed save");
    // }
    // printSnack("Success save");
    // navigateBack();
  }
}

enum RecordingState { standby, recording, done }
