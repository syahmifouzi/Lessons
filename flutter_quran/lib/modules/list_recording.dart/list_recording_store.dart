import 'dart:io';

import 'package:audioplayers/audioplayers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/audio_player/audio_player_store.dart';
import 'package:path_provider/path_provider.dart';

class ListRecordingStore extends ChangeNotifier {
  List<RecordingFile> recordingFileList = [];

  Future<void> updateListOfRecording() async {
    final directory = await getApplicationDocumentsDirectory();
    List<FileSystemEntity> temp = directory.listSync();
    var recordingList = temp.where((x) => x.path.contains('.m4a'));
    if (recordingList.isEmpty) return;
    recordingFileList.clear();
    for (var recording in recordingList) {
      String path = recording.path;
      // Extract the filename with extension
      String filenameWithExtension = path.split('/').last;
      // Remove the extension
      String filename = filenameWithExtension.split('.').first;
      recordingFileList.add(RecordingFile(name: filename, path: path));
    }
    notifyListeners();
  }

  Future<void> playAudio(RecordingFile recordingFile) async {
    final player = AudioPlayer();
    await player.play(DeviceFileSource(recordingFile.path));
    // await player.getDuration();
    // await player.play();
    // await player.stop();
  }
}
