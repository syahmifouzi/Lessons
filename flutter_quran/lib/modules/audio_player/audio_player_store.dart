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

  Future<int> getDuration() async {
    final player = AudioPlayer();
    await player.setSource(DeviceFileSource(recordingFile.path));
    final duration = await player.getDuration();
    return 1;
    if (duration == null) {
      return 1;
    }
    recordingFile.duration = duration;
    notifyListeners();
    print(recordingFile.duration);
    // await player.getDuration();
    // await player.play();
    // await player.stop();
    return 0;
  }
}
