import './media.dart';

class Server {
  String ipAddr;
  String name;
  List<Media> mediaList;
  bool hasVideo;
  bool reachable;

  Server(
    this.ipAddr,
    this.name,
    this.mediaList,
    this.hasVideo,
    this.reachable,
  );
}
