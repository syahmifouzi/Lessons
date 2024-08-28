import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/audio_player/audio_player_store.dart';
import 'package:provider/provider.dart';

class AudioPlayerView extends StatelessWidget {
  const AudioPlayerView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Row(mainAxisAlignment: MainAxisAlignment.end, children: [
            Consumer<AudioPlayerStore>(builder: (context, cart, child) {
              final duration = cart.recordingFile.duration;
              return Text("0:0/$duration");
            }),
            const SizedBox(width: 8),
          ]),
          Slider(value: 0.0, onChanged: (value) {}),
          const AudioButtonState(),
        ],
      ),
    );
  }
}

class AudioButtonState extends StatelessWidget {
  const AudioButtonState({super.key});

  @override
  Widget build(BuildContext context) {
    return Consumer<AudioPlayerStore>(builder: (context, cart, child) {
      switch (cart.buttonState) {
        case AudioPlayerButtonState.initial:
          return const ButtonInInitialState();
        case AudioPlayerButtonState.playing:
          return const ButtonInPlayingState();
        case AudioPlayerButtonState.pausing:
          return const ButtonInPausingState();
        default:
          return const ButtonInInitialState();
      }
    });
  }
}

class ButtonInInitialState extends StatelessWidget {
  const ButtonInInitialState({super.key});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        onPressed: () {
          Provider.of<AudioPlayerStore>(context, listen: false)
              .setButtonState(AudioPlayerButtonState.playing);
        },
        child: const Text("Play"));
  }
}

class ButtonInPlayingState extends StatelessWidget {
  const ButtonInPlayingState({super.key});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        onPressed: () {
          Provider.of<AudioPlayerStore>(context, listen: false)
              .setButtonState(AudioPlayerButtonState.pausing);
        },
        child: const Text("Pause"));
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
            onPressed: () {
              Provider.of<AudioPlayerStore>(context, listen: false)
                  .setButtonState(AudioPlayerButtonState.playing);
            },
            child: const Text("Resume")),
        ElevatedButton(
            onPressed: () {
              Provider.of<AudioPlayerStore>(context, listen: false)
                  .setButtonState(AudioPlayerButtonState.initial);
            },
            child: const Text("Stop")),
      ],
    );
  }
}
