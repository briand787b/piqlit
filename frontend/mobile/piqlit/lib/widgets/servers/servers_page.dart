import 'package:flutter/material.dart';

import '../../models/server.dart';

class ServersPage extends StatelessWidget {
  final servers = [
    Server('192.168.1.2', 'Downstairs TV', [], true, true),
    Server('192.168.1.22', 'Repo', [], false, true),
    Server('192.168.1.4', 'iMac', [], true, true),
    Server('192.168.1.8', 'Brian\'s iPad', [], true, true),
    Server('192.168.1.15', 'Tia\'s iPad', [], true, true),
    Server('192.168.1.15', 'Top', [], true, false),
  ];

  void addServer() {
    print('"addServer" has been called');
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          Container(
            // color: Colors.grey,
            padding: EdgeInsets.all(20),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                Text(
                  'Servers',
                  textScaleFactor: 2,
                  textAlign: TextAlign.start,
                  style: TextStyle(fontWeight: FontWeight.w500),
                ),
                RaisedButton(
                  child: Icon(
                    Icons.add,
                    color: Colors.blue,
                  ),
                  onPressed: () => addServer,
                ),
              ],
            ),
          ),
          Container(
            decoration: BoxDecoration(color: Colors.grey),
            child: Column(
              children: servers.map((Server s) {
                return Container(
                  color: s.reachable ? Colors.green : Colors.red,
                  child: Card(
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      children: <Widget>[
                        Column(
                          children: <Widget>[
                            Text(
                              s.name,
                              style: TextStyle(
                                fontSize: 20,
                                fontWeight: FontWeight.w500,
                              ),
                            ),
                            Text('IP Address: ${s.ipAddr}'),
                          ],
                        ),
                        Icon( s.reachable ? Icons.check : Icons.clear),
                      ],
                    ),
                  ),
                );
              }).toList(),
            ),
          ),
        ],
      ),
    );
  }
}
