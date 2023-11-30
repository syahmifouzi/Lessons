import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:just_audio/just_audio.dart';

class AudioStore extends ChangeNotifier {
  AudioDbDoc audioDetails =
      AudioDbDoc("", "", "", "", DateTime.now(), "", "", "0", "0");
  late AudioPlayer _audioPlayer;
  Duration sliderTotalDuration = Duration.zero;
  Duration sliderBufferedPosition = Duration.zero;
  Duration sliderPosition = Duration.zero;

  void setinit() async {
    _audioPlayer = AudioPlayer();
    // String tempUrl =
    // "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3";
    // await _audioPlayer.setUrl(tempUrl);
    await _audioPlayer.setUrl(audioDetails.url);

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

  void setPause() {
    _audioPlayer.pause();
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
  AudioDbDoc(this.id, this.title, this.url, this.duration, this.datetime,
      this.surah, this.status, this.ayatStart, this.ayatEnd);
  final String id;
  String title;
  String url;
  String duration;
  String surah;
  String status;
  String ayatStart;
  String ayatEnd;
  DateTime datetime;

  String getDate() {
    return DateFormat("dd-MM-yyyy").format(datetime);
  }

  String getTime() {
    return DateFormat("HH:mm").format(datetime);
  }

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
          'Invalid JSON: required "url" field of type String in $data');
    }
    final duration = data["duration"];
    if (duration is! String) {
      throw FormatException(
          'Invalid JSON: required "duration" field of type String in $data');
    }
    final status = data["status"];
    if (status is! String) {
      throw FormatException(
          'Invalid JSON: required "status" field of type String in $data');
    }
    final surah = data["surah"];
    if (surah is! String) {
      throw FormatException(
          'Invalid JSON: required "surah" field of type String in $data');
    }
    final ayatStart = data["ayatStart"];
    if (ayatStart is! String) {
      throw FormatException(
          'Invalid JSON: required "ayatStart" field of type String in $data');
    }
    final ayatEnd = data["ayatEnd"];
    if (ayatEnd is! String) {
      throw FormatException(
          'Invalid JSON: required "ayatEnd" field of type String in $data');
    }
    final timestamp = data["timestamp"];
    if (timestamp is! Timestamp) {
      throw FormatException(
          'Invalid JSON: required "datetime" field of type Timestamp in $data');
    }
    final datetime = timestamp.toDate();

    return AudioDbDoc(
        id, title, url, duration, datetime, surah, status, ayatStart, ayatEnd);
  }
}
