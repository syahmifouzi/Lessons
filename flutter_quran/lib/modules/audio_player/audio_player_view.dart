import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/audio_player/audio_player_store.dart';
import 'package:flutter_sound/flutter_sound.dart';
import 'package:provider/provider.dart';

class AudioPlayerView extends StatelessWidget {
  const AudioPlayerView({super.key});

  Future<int> _initialize(BuildContext context) async {
    try {
      await Provider.of<AudioPlayerStore>(context, listen: false)
          .initializeSingleAudioView();
    } catch (e) {
      Future.error(e);
    }
    return 0;
  }

  Future<void> _delayedStopCall(BuildContext context) async {
    try {
      await Provider.of<AudioPlayerStore>(context, listen: false).stop();
    } catch (e) {
      Future.error(e);
    }
  }

  Widget _errorPage() {
    return const Center(child: Text("Error"));
  }

  Widget _playerPositionChangedHandler(
      {required Stream<PlaybackDisposition> streamDuration,
      required String duration}) {
    return StreamBuilder(
        stream: streamDuration,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            String currentPosition = snapshot.data!.position
                .toString()
                .split('.')
                .first
                .padLeft(8, "0");
            return Text("$currentPosition/$duration");
          }
          return Text("00:00:00/$duration");
        });
  }

  Widget _playerStateHandler({
    required Stream<PlayerState> streamState,
    required Stream<PlaybackDisposition> streamDuration,
    required String duration,
  }) {
    return StreamBuilder(
        stream: streamState,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            switch (snapshot.data) {
              case PlayerState.isStopped:
                return Text("00:00:00/$duration");
              // case PlayerState.completed:
              //   return FutureBuilder(
              //       future: _delayedStopCall(context),
              //       builder: (context, snapshot) {
              //         return Text("00:00:00/$duration");
              //       });
              default:
                return _playerPositionChangedHandler(
                    streamDuration: streamDuration, duration: duration);
            }
          }
          return const Text("00:00:00/00:00:00");
        });
  }

  Widget _audioSliderHandler({
    required Stream<PlaybackDisposition> streamDuration,
    required double max,
    required void Function(double) onChanged,
    required void Function(double) onChangeStart,
    required void Function(double) onChangeEnd,
    required AudioPlayerButtonState buttonState,
  }) {
    switch (buttonState) {
      case AudioPlayerButtonState.initial:
        return const Slider(value: 0, onChanged: null);
      default:
        return StreamBuilder(
            stream: streamDuration,
            builder: (context, snapshot) {
              if (snapshot.hasData) {
                double currentPosition =
                    snapshot.data!.position.inSeconds.toDouble();
                return Slider(
                  value: currentPosition,
                  onChanged: onChanged,
                  max: max,
                  onChangeStart: onChangeStart,
                  onChangeEnd: onChangeEnd,
                );
              }
              return const Slider(value: 0, onChanged: null);
            });
    }
  }

  Future<void> _showErrorDialog(context, String error) async {
    return showDialog(
        context: context,
        barrierDismissible: true,
        builder: (context) {
          return AlertDialog(
            title: const Text("Error"),
            content: Text(error),
            actions: [
              TextButton(
                  onPressed: () => Navigator.of(context).pop(),
                  child: const Text("OK"))
            ],
          );
        });
  }

  Future<int> _deleteLogic(context) async {
    final result = await showDialog<String>(
        context: context,
        barrierDismissible: true,
        builder: (context) {
          return AlertDialog(
            title: const Text("About to delete this file!"),
            content: const Text("Are you sure?"),
            actions: [
              TextButton(
                  onPressed: () async {
                    try {
                      await Provider.of<AudioPlayerStore>(context,
                              listen: false)
                          .delete();
                      if (context.mounted) {
                        Navigator.of(context).pop("0");
                      }
                    } catch (e) {
                      if (context.mounted) {
                        Navigator.of(context).pop(e.toString());
                      }
                    }
                  },
                  child: const Text("Delete")),
              ElevatedButton(
                  onPressed: () => Navigator.of(context).pop("1"),
                  child: const Text("Cancel"))
            ],
          );
        });

    final resultString = result ?? "1";

    if (int.tryParse(resultString) == null) {
      return Future.error(resultString);
    }
    return Future.value(int.tryParse(resultString));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        actions: [
          ElevatedButton(
              onPressed: () async {
                try {
                  final result = await _deleteLogic(context);
                  if (result == 0 && context.mounted) {
                    Navigator.of(context).pop();
                  }
                } catch (e) {
                  if (context.mounted) {
                    _showErrorDialog(context, e.toString());
                  }
                }
              },
              child: const Text("Delete"))
        ],
      ),
      body: FutureBuilder(
        future: _initialize(context),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const CircularProgressIndicator();
          } else if (snapshot.hasError) {
            return _errorPage();
          }
          return Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Row(mainAxisAlignment: MainAxisAlignment.end, children: [
                Consumer<AudioPlayerStore>(builder: (context, cart, child) {
                  final recordingFile = cart.recordingFileSelected;
                  final duration = recordingFile.duration
                      .toString()
                      .split('.')
                      .first
                      .padLeft(8, "0");
                  final streamState = cart.onPlayerStateChanged;
                  final streamDuration = cart.onPositionChanged;
                  return _playerStateHandler(
                    streamDuration: streamDuration,
                    streamState: streamState,
                    duration: duration,
                  );
                }),
                const SizedBox(width: 8),
              ]),
              Consumer<AudioPlayerStore>(builder: (context, cart, child) {
                final streamDuration = cart.onPositionChanged;
                final max = cart.sliderMax;
                final onChanged = cart.sliderOnChanged;
                final onChangeStart = cart.sliderOnChangeStart;
                final onChangeEnd = cart.sliderOnChangeEnd;
                final buttonState = cart.buttonState;
                return _audioSliderHandler(
                    streamDuration: streamDuration,
                    max: max,
                    onChanged: onChanged,
                    onChangeStart: onChangeStart,
                    onChangeEnd: onChangeEnd,
                    buttonState: buttonState);
              }),
              const AudioButtonState(),
            ],
          );
        },
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
          Provider.of<AudioPlayerStore>(context, listen: false).play();
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
          Provider.of<AudioPlayerStore>(context, listen: false).pause();
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
              Provider.of<AudioPlayerStore>(context, listen: false).resume();
            },
            child: const Text("Resume")),
        ElevatedButton(
            onPressed: () {
              Provider.of<AudioPlayerStore>(context, listen: false).stop();
            },
            child: const Text("Stop")),
      ],
    );
  }
}
