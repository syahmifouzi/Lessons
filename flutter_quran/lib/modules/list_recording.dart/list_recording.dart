import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/audio_player/audio_player_store.dart';
import 'package:flutter_quran/modules/audio_player/audio_player_view.dart';
import 'package:flutter_quran/modules/list_recording.dart/list_recording_store.dart';
import 'package:provider/provider.dart';

class ListRecording extends StatelessWidget {
  const ListRecording({super.key});

  @override
  Widget build(BuildContext context) {
    Provider.of<ListRecordingStore>(context, listen: false)
        .updateListOfRecording();
    return Consumer<ListRecordingStore>(builder: (context, cart, child) {
      List<RecordingFile> recordingFileList = cart.recordingFileList;
      return ListView.builder(
          itemCount: recordingFileList.length,
          itemBuilder: (context, index) {
            return ListTile(
                title: Text(recordingFileList[index].name),
                onTap: () {
                  Provider.of<AudioPlayerStore>(context, listen: false)
                      .setRecordingFile(recordingFileList[index]);
                  Provider.of<AudioPlayerStore>(context, listen: false)
                      .getDuration();
                  Navigator.push(
                      context,
                      MaterialPageRoute(
                          builder: (context) => const AudioPlayerView()));
                  // onTap: () => cart.playAudio(recordingFileList[index]),
                });
          });
    });
  }
}
