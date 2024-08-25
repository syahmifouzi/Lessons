import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/voice_recording/recording_button_store.dart';
import 'package:flutter_quran/modules/debug_print/debug_print_store.dart';
import 'package:provider/provider.dart';

class RecordingButton extends StatelessWidget {
  const RecordingButton({super.key});

  @override
  Widget build(BuildContext context) {
    return Consumer<RecordingButtonStore>(builder: (context, cart, child) {
      switch (cart.buttonState) {
        case RecordingButtonState.stopped:
          return StartRecordButton();
        default:
          return StopRecordButton();
      }
    });
  }
}

class StartRecordButton extends StatelessWidget {
  const StartRecordButton({super.key});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        onPressed: () {
          Provider.of<RecordingButtonStore>(context, listen: false)
              .setButtonState(RecordingButtonState.recording);
          Provider.of<DebugPrintStore>(context, listen: false).print(
              "recording_button.dart: StartRecordButton: Started Recording");
        },
        child: const Text("Start Recording"));
  }
}

class PauseRecordButton extends StatelessWidget {
  const PauseRecordButton({super.key});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        onPressed: () {}, child: const Text("Pause Recording"));
  }
}

class ResumeRecordButton extends StatelessWidget {
  const ResumeRecordButton({super.key});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        onPressed: () {}, child: const Text("Resume Recording"));
  }
}

class StopRecordButton extends StatelessWidget {
  const StopRecordButton({super.key});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        onPressed: () {
          Provider.of<RecordingButtonStore>(context, listen: false)
              .setButtonState(RecordingButtonState.stopped);
          Provider.of<DebugPrintStore>(context, listen: false).print(
              "recording_button.dart: StopRecordButton: Stopped Recording");
        },
        child: const Text("Stop Recording"));
  }
}
