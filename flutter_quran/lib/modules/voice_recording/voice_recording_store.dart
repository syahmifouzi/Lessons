import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:path_provider/path_provider.dart';
import 'package:record/record.dart';

class VoiceRecordingStore extends ChangeNotifier {
  final record = AudioRecorder();

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

  void start() async {
    // Check and request permission if needed
    if (await record.hasPermission()) {
      final path = await _localPath;
      DateTime now = DateTime.now();
      String formattedDate = DateFormat('yyMMdd_HHmmss').format(now);
      // Start recording to file
      await record.start(const RecordConfig(),
          path: '$path/$formattedDate.m4a');
    }
  }

  void pause() async {
    await record.pause();
  }

  void resume() async {
    await record.resume();
  }

  void stop() async {
    // Stop recording...
    // final path = await record.stop();
    await record.stop();
    // record.dispose(); // As always, don't forget this one.
  }
}
