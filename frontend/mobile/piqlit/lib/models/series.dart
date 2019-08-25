import './media.dart';

class Series extends Media {
  int seasonCount;

  Series(this.seasonCount, String title, String artist, DateTime date)
      : super(title, artist, date);
}
