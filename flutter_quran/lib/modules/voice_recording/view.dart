import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/voice_recording/recording_button.dart';

class VoiceRecording extends StatefulWidget {
  const VoiceRecording({super.key});

  @override
  State<VoiceRecording> createState() => _VoiceRecordingState();
}

class _VoiceRecordingState extends State<VoiceRecording> {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(children: [RecordingButton()]),
    );
  }
}
