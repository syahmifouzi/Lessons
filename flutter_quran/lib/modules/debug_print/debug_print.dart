import 'package:flutter/material.dart';
import 'package:flutter_quran/modules/debug_print/debug_print_store.dart';
import 'package:provider/provider.dart';

class DebugprintView extends StatelessWidget {
  const DebugprintView({super.key});

  @override
  Widget build(BuildContext context) {
    return Consumer<DebugPrintStore>(builder: (context, cart, child) {
      List<String> messages = cart.showMessages();
      return Column(
        children: [
          Row(
            children: [
              ElevatedButton(
                  onPressed: () =>
                      Provider.of<DebugPrintStore>(context, listen: false)
                          .toggleReverse(),
                  child: const Text("Reverse Order")),
              ElevatedButton(
                  onPressed: () =>
                      Provider.of<DebugPrintStore>(context, listen: false)
                          .clearAll(),
                  child: const Text("Clear")),
            ],
          ),
          Expanded(
            child: ListView.builder(
                itemCount: messages.length,
                itemBuilder: (context, index) {
                  return ListTile(title: Text(messages[index]));
                }),
          ),
        ],
      );
    });
  }
}
