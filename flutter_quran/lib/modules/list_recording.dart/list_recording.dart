import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/audio_player/audio_player_store.dart';
import 'package:flutter_quran/modules/audio_player/audio_player_view.dart';
import 'package:provider/provider.dart';

class ListRecording extends StatelessWidget {
  const ListRecording({super.key});

  @override
  Widget build(BuildContext context) {
    Provider.of<AudioPlayerStore>(context, listen: false)
        .updateListOfRecording();
    return Consumer<AudioPlayerStore>(builder: (context, cart, child) {
      List<RecordingFile> recordingFileList = cart.recordingFileList;
      return ListView.builder(
          itemCount: recordingFileList.length,
          itemBuilder: (context, index) {
            return ListTile(
                title: Text(recordingFileList[index].name),
                onTap: () async {
                  Provider.of<AudioPlayerStore>(context, listen: false)
                      .recordingFileIndex = index;
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
