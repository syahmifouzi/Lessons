import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_sound/flutter_sound.dart';
import 'package:path_provider/path_provider.dart';

enum AudioPlayerButtonState { initial, playing, pausing, stopped }

class RecordingFile {
  String name;
  String path;
  Duration? duration;

  RecordingFile({required this.name, required this.path});
}

class AudioSlider {
  double value;
  double min;
  double max;

  AudioSlider({this.value = 0, this.min = 0, this.max = 0});
}

class AudioPlayerStore extends ChangeNotifier {
  AudioPlayerButtonState _buttonState = AudioPlayerButtonState.initial;
  final List<RecordingFile> _recordingFileList = [];
  int _recordingFileIndex = -1;
  FlutterSoundPlayer player = FlutterSoundPlayer();
  Duration _currentPlaybackPosition = const Duration();
  AudioSlider audioSlider = AudioSlider();
  RecordingFile _recordingFileSelected = RecordingFile(name: "", path: "");
  // PlayerState _lastPlayerState

  AudioPlayerStore() {
    _initialize();
  }

  void _initialize() async {
    await player.openPlayer();
  }

  RecordingFile get recordingFileSelected => _recordingFileSelected;

  set recordingFileSelected(RecordingFile newFile) {
    _recordingFileSelected = newFile;
    notifyListeners();
  }

  List<RecordingFile> get recordingFileList =>
      _recordingFileList.reversed.toList();

  Future<void> updateListOfRecording() async {
    final directory = await getApplicationDocumentsDirectory();
    List<FileSystemEntity> temp = directory.listSync();
    var recordingList = temp.where((x) => x.path.contains('.m4a'));
    if (recordingList.isEmpty) return;
    _recordingFileList.clear();
    for (var recording in recordingList) {
      String path = recording.path;
      // Extract the filename with extension
      String filenameWithExtension = path.split('/').last;
      // Remove the extension
      String filename = filenameWithExtension.split('.').first;
      _recordingFileList.add(RecordingFile(name: filename, path: path));
    }
    notifyListeners();
  }

  void updateCurrentFile(Duration? duration) {
    final item = _recordingFileList
        .firstWhere((item) => item.path == recordingFileSelected.path);
    item.duration = duration;
  }

  Future<void> deleteCurrentFile() async {
    if (recordingFileIndex < 0 ||
        recordingFileIndex > _recordingFileList.length - 1) {
      return Future.error("No file selected.");
    }
    final index = _recordingFileList
        .indexWhere((item) => item.path == recordingFileSelected.path);
    _recordingFileList.removeAt(index);
    notifyListeners();
  }

  Future<void> initializeSingleAudioView() async {
    if (recordingFileIndex < 0 ||
        recordingFileIndex > _recordingFileList.length - 1) {
      return Future.error("No file selected.");
    }
    final recordingFileSelectedTemp = recordingFileList[recordingFileIndex];
    // await player.setSource(DeviceFileSource(recordingFileSelectedTemp.path));
    // final duration = await player.getDuration();
    // if (duration == null) {
    //   return Future.error("File has no duration.");
    // }
    // recordingFileSelectedTemp.duration = duration;
    // sliderMax = duration.inSeconds.toDouble();
    // Bypass calling setState during build error
    await Future.delayed(Duration.zero);
    recordingFileSelected = recordingFileSelectedTemp;
    // updateCurrentFile(recordingFileSelected.duration);
  }

  AudioPlayerButtonState get buttonState => _buttonState;

  set buttonState(AudioPlayerButtonState newState) {
    _buttonState = newState;
    notifyListeners();
  }

  int get recordingFileIndex => _recordingFileIndex;

  set recordingFileIndex(int newRecordingFileIndex) {
    _recordingFileIndex = newRecordingFileIndex;
    notifyListeners();
  }

  set currentPlaybackPosition(Duration newPosition) {
    _currentPlaybackPosition = newPosition;
    notifyListeners();
  }

  Duration get currentPlaybackPosition => _currentPlaybackPosition;

  void _whenFinished() {
    buttonState = AudioPlayerButtonState.initial;
  }

  Future<void> play() async {
    final file = recordingFileSelected;
    final duration = await player.startPlayer(
        fromURI: file.path, codec: Codec.aacMP4, whenFinished: _whenFinished);
    if (duration == null) {
      return Future.error("File has no duration.");
    }
    sliderMax = duration.inSeconds.toDouble();
    updateCurrentFile(duration);
    buttonState = AudioPlayerButtonState.playing;
  }

  void pause() async {
    await player.pausePlayer();
    buttonState = AudioPlayerButtonState.pausing;
  }

  void resume() async {
    await player.resumePlayer();
    buttonState = AudioPlayerButtonState.playing;
  }

  Future<void> delete() async {
    if (recordingFileIndex < 0 ||
        recordingFileIndex > _recordingFileList.length - 1) {
      return Future.error("No file selected.");
    }
    if (player.playerState != PlayerState.isStopped) {
      await stop();
    }
    final file = File(recordingFileSelected.path);
    if (!(await file.exists())) {
      return Future.error("File not found. Maybe wrong path.");
    }
    await file.delete();
    try {
      await deleteCurrentFile();
    } catch (e) {
      return Future.error(e);
    }
  }

  Future<void> stop() async {
    await player.stopPlayer();
    await player.seekToPlayer(const Duration());
    // await player.release();

    // function with notifyListener() must be called last inside any function
    // await must be called to prevent setState() during build
    buttonState = AudioPlayerButtonState.initial;
    sliderValue = 0;
  }

  Stream<PlaybackDisposition> get onPositionChanged async* {
    // TODO: check if the player is playing to handle this null value
    await player.setSubscriptionDuration(Durations.medium1);
    yield* player.onProgress!;
  }

  Stream<PlayerState> get onPlayerStateChanged async* {
    Duration interval = Durations.medium1;
    Stream<PlayerState> stream = Stream.periodic(interval, (x) {
      PlayerState playerState = player.playerState;
      return playerState;
    });
    yield* stream;
  }

  double get sliderMin {
    return audioSlider.min;
  }

  double get sliderMax {
    return audioSlider.max;
  }

  set sliderMax(double newValue) {
    audioSlider.max = newValue;
    notifyListeners();
  }

  double get sliderValue {
    return audioSlider.value;
  }

  set sliderValue(double newValue) {
    audioSlider.value = newValue;
    notifyListeners();
  }

  void sliderOnChanged(double value) async {
    await player.seekToPlayer(Duration(seconds: value.toInt()));
    sliderValue = value;
  }

  void sliderOnChangeStart(double value) => pause();

  void sliderOnChangeEnd(double value) => resume();
}
