import 'package:flutter/material.dart';
import 'package:flutter_sound/flutter_sound.dart';
import 'package:intl/intl.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:wakelock_plus/wakelock_plus.dart';

class VoiceRecordingStore extends ChangeNotifier {
  final record = FlutterSoundRecorder();

// Declare Constructor
  VoiceRecordingStore() {
    _initialize();
  }

  void _initialize() async {
    await record.openRecorder();
  }

  Future<String> get _localPath async {
    final directory = await getApplicationDocumentsDirectory();

    return directory.path;
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

  Future<void> start() async {
    // Check and request permission if needed
    if (await Permission.microphone.request().isGranted) {
      // await record.openRecorder();
      final path = await _localPath;
      DateTime now = DateTime.now();
      String formattedDate = DateFormat('yyMMdd_HHmmss').format(now);
      // Enable Wakelock
      await WakelockPlus.enable();
      // Start recording to file
      await record.startRecorder(
          toFile: '$path/$formattedDate.m4a', codec: Codec.aacMP4);
      // If codec not supported, this line above will stuck
    } else {
      return Future.error("Permission not granted!");
    }
  }

  void pause() async {
    await record.pauseRecorder();
  }

  void resume() async {
    await record.resumeRecorder();
  }

  void stop() async {
    // Stop recording...
    // final path = await record.stop();
    await record.stopRecorder();
    // Disable Wakelock
    await WakelockPlus.disable();
    // await record.closeRecorder();
    // record.dispose(); // As always, don't forget this one.
  }
}
