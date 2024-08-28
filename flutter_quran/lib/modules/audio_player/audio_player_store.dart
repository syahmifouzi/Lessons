import 'package:audioplayers/audioplayers.dart';
import 'package:flutter/material.dart';

enum AudioPlayerButtonState { initial, playing, pausing, stopped }

class RecordingFile {
  String name;
  String path;
  Duration? duration;

  RecordingFile({required this.name, required this.path});
}

class AudioPlayerStore extends ChangeNotifier {
  AudioPlayerButtonState buttonState = AudioPlayerButtonState.initial;
  RecordingFile recordingFile = RecordingFile(name: "", path: "");

  void setButtonState(AudioPlayerButtonState newButtonState) {
    buttonState = newButtonState;
    notifyListeners();
  }

  void setRecordingFile(RecordingFile newRecordingFile) {
    recordingFile = newRecordingFile;
    notifyListeners();
  }

  Future<void> getDuration() async {
    final player = AudioPlayer();
    await player.setSource(DeviceFileSource(recordingFile.path));
    recordingFile.duration = await player.getDuration();
    notifyListeners();
    print(recordingFile.duration);
    // await player.getDuration();
    // await player.play();
    // await player.stop();
  }
}
