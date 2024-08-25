import 'dart:io';

import 'package:audioplayers/audioplayers.dart';
import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';
// import 'package:permission_handler/permission_handler.dart';
import 'package:record/record.dart';

class VoiceRecordingStore extends ChangeNotifier {
  final record = AudioRecorder();

  Future<String> get _localPath async {
    final directory = await getApplicationDocumentsDirectory();

    return directory.path;
  }

  Future<List<FileSystemEntity>> listFiles() async {
    final directory = await getApplicationDocumentsDirectory();
    List<FileSystemEntity> temp = directory.listSync();
    print(temp);
    print(temp.length);
    // print(temp[2]);
    // var x = temp[2];
    // print(x.uri);
    return temp;
  }

  // Future<File> _localFile(String title) async {
  //   final path = await _localPath;
  //   return File('$path/$title.m4a');
  // }

  // void checkPermissions() async {
  //   Map<Permission, PermissionStatus> permissions =
  //       await [Permission.microphone].request();

  //   bool permissionsGranted = permissions[Permission.microphone]!.isGranted;

  //   print(permissionsGranted);
  // }

  void start() async {
    // Check and request permission if needed
    if (await record.hasPermission()) {
      final path = await _localPath;
      // Start recording to file
      await record.start(const RecordConfig(), path: '$path/myFile.m4a');
    }
  }

  void stop() async {
    // Stop recording...
    final path = await record.stop();
    print('MY PRINT: $path');
    // record.dispose(); // As always, don't forget this one.
  }

  Future<void> playAudio() async {
    final directory = await getApplicationDocumentsDirectory();
    List<FileSystemEntity> temp = directory.listSync();
    var x = temp[3];
    print(x.path);
    print(x.uri);
    final player = AudioPlayer();
    await player.play(DeviceFileSource(x.path));
    // await player.play();
    // await player.stop();
  }
}
