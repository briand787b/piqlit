import 'package:flutter/material.dart';

import '../../models/server.dart';
import '../../models/media.dart';

class ServersPage extends StatelessWidget {
  final servers = [
    Server('192.168.1.2', List<Media>()),
    Server('192.168.1.3', [
      Media('backdoor sluts 9', 'eric cartman', DateTime.now()),
      Media('Jurassic Park', 'Michael Crichton', DateTime.now()),
      
    ]),
    Server('192.168.1.2', List<Media>()),
    Server('192.168.1.3', [
      Media('backdoor sluts 9', 'eric cartman', DateTime.now()),
      Media('Jurassic Park', 'Michael Crichton', DateTime.now()),
      
    ]),
    Server('192.168.1.2', List<Media>()),
    Server('192.168.1.3', [
      Media('backdoor sluts 9', 'eric cartman', DateTime.now()),
      Media('Jurassic Park', 'Michael Crichton', DateTime.now()),
      
    ]),
    Server('192.168.1.2', List<Media>()),
    Server('192.168.1.3', [
      Media('backdoor sluts 9', 'eric cartman', DateTime.now()),
      Media('Jurassic Park', 'Michael Crichton', DateTime.now()),
    ]),
    Server('192.168.1.2', List<Media>()),
    Server('192.168.1.3', [
      Media('backdoor sluts 9', 'eric cartman', DateTime.now()),
      Media('Jurassic Park', 'Michael Crichton', DateTime.now()),
    ]),
  ];

  void addServer() {
    print('"addServer" has been called');
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            Text(
              'Servers',
              textScaleFactor: 2,
              textAlign: TextAlign.start,
              style: TextStyle(fontWeight: FontWeight.w500),
            ),
            Text('Hello World!')
          ],
        ),
        Container(
          decoration: BoxDecoration(color: Colors.amber),
          child: SingleChildScrollView(
            controller: ScrollController(),
            child: Column(
              children: servers.map((Server s) {
                return Container(
                  child: Text(
                    'This is a card',
                    style: TextStyle(
                      fontSize: 50,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                );
              }).toList(),
            ),
          ),
        ),
        // ListView(

        // ),
      ],
    );
  }
}
