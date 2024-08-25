import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/voice_recording/recording_button_store.dart';
import 'package:flutter_quran/modules/debug_print/debug_print_store.dart';
import 'package:provider/provider.dart';

class RecordingButton extends StatelessWidget {
  const RecordingButton({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Consumer<RecordingButtonStore>(builder: (context, cart, child) {
          switch (cart.buttonState) {
            case RecordingButtonState.initial:
              return const Text("Start New Recording");
            case RecordingButtonState.recording:
              return const Text("Now Recording");
            case RecordingButtonState.pausing:
              return const Text("Paused");
            default:
              return const Text("Saving");
          }
        }),
        Consumer<RecordingButtonStore>(builder: (context, cart, child) {
          switch (cart.buttonState) {
            case RecordingButtonState.initial:
              return const ButtonInInitialState();
            case RecordingButtonState.recording:
              return const ButtonInRecordingState();
            case RecordingButtonState.pausing:
              return const ButtonInPausingState();
            default:
              return const StopRecordButton();
          }
        }),
      ],
    );
  }
}

class ButtonInInitialState extends StatelessWidget {
  const ButtonInInitialState({super.key});

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

class ButtonInRecordingState extends StatelessWidget {
  const ButtonInRecordingState({super.key});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        onPressed: () =>
            Provider.of<RecordingButtonStore>(context, listen: false)
                .setButtonState(RecordingButtonState.pausing),
        child: const Text("Pause Recording"));
  }
}

class ButtonInPausingState extends StatelessWidget {
  const ButtonInPausingState({super.key});

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceAround,
      children: [
        ElevatedButton(
            onPressed: () =>
                Provider.of<RecordingButtonStore>(context, listen: false)
                    .setButtonState(RecordingButtonState.recording),
            child: const Text("Resume Recording")),
        ElevatedButton(
            onPressed: () =>
                Provider.of<RecordingButtonStore>(context, listen: false)
                    .setButtonState(RecordingButtonState.stopped),
            child: const Text("Stop Recording")),
      ],
    );
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
