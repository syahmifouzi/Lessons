import 'package:flutter/material.dart';
import 'package:just_audio/just_audio.dart';

class AudioStore extends ChangeNotifier {
  AudioDbDoc audioDetails = AudioDbDoc("Empty", "Empty", "Empty");
  late AudioPlayer _audioPlayer;
  Duration sliderTotalDuration = Duration.zero;
  Duration sliderBufferedPosition = Duration.zero;
  Duration sliderPosition = Duration.zero;

  void setinit() async {
    _audioPlayer = AudioPlayer();
    String tempUrl =
        "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3";
    await _audioPlayer.setUrl(tempUrl);
    // await _audioPlayer.setUrl(audioDetails.url);

    _audioPlayer.playerStateStream.listen((playerState) {
      final isPlaying = playerState.playing;
      final processingState = playerState.processingState;
      if (processingState == ProcessingState.loading ||
          processingState == ProcessingState.buffering) {
        // button state loading
      } else if (!isPlaying) {
        // button state pause
      } else if (processingState != ProcessingState.completed) {
        // button state playing
      } else {
        _audioPlayer.seek(Duration.zero);
        _audioPlayer.pause();
      }
    });

    _audioPlayer.positionStream.listen((position) {
      // position at slider primary value
      setPosition(position);
    });

    _audioPlayer.bufferedPositionStream.listen((bufferedPosition) {
      // bufferedPosition at slider secondary value
      setBufferedPosition(bufferedPosition);
    });

    _audioPlayer.durationStream.listen((totalDuration) {
      // totalDuration ?? Duration.zero;
      setTotalDuration(totalDuration ?? Duration.zero);
    });
  }

  void setdispose() {
    setStop();
    _audioPlayer.dispose();
  }

  void setAudio(AudioDbDoc x) {
    audioDetails = x;
    notifyListeners();
  }

  void setPlay() {
    _audioPlayer.play();
  }

  void setStop() {
    _audioPlayer.stop();
    setSeekMs(0);
  }

  void setPosition(Duration x) {
    sliderPosition = x;
    notifyListeners();
  }

  void setBufferedPosition(Duration x) {
    sliderBufferedPosition = x;
    notifyListeners();
  }

  void setTotalDuration(Duration x) {
    sliderTotalDuration = x;
    notifyListeners();
  }

  void setSeekMs(double x) {
    Duration position = Duration(milliseconds: int.parse(x.round().toString()));
    _audioPlayer.seek(position);
    // notifyListeners();
  }

  double getTotalDurationMs() {
    return double.parse(sliderTotalDuration.inMilliseconds.toString());
  }

  double getPositionMs() {
    return double.parse(sliderPosition.inMilliseconds.toString());
  }

  double getBufferedPositionMs() {
    return double.parse(sliderBufferedPosition.inMilliseconds.toString());
  }

  String getPositionLabel() {
    String x = sliderPosition.toString();
    List<String> splitter = x.split(":");
    List<String> msSplitter = splitter[2].split(".");
    String ms = msSplitter[0];
    if (sliderPosition.inHours >= 1) {
      return "${splitter[0]}:${splitter[1]}:$ms";
    }
    return "${splitter[1]}:$ms";
  }

  String getTotalDurationLabel() {
    String x = sliderTotalDuration.toString();
    List<String> splitter = x.split(":");
    List<String> msSplitter = splitter[2].split(".");
    String ms = msSplitter[0];
    if (sliderPosition.inHours >= 1) {
      return "${splitter[0]}:${splitter[1]}:$ms";
    }
    return "${splitter[1]}:$ms";
  }
}

class AudioDbDoc {
  AudioDbDoc(this.id, this.title, this.url);
  final String id;
  String title;
  String url;

  factory AudioDbDoc.fromJson(String id1, Map<String, dynamic> data) {
    final id = id1;
    final title = data["title"];
    if (title is! String) {
      throw FormatException(
          'Invalid JSON: required "title" field of type String in $data');
    }
    final url = data["url"];
    if (url is! String) {
      throw FormatException(
          'Invalid JSON: required "subtitle" field of type String in $data');
    }

    return AudioDbDoc(id, title, url);
  }
}
