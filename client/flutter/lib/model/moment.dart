import 'package:json_annotation/json_annotation.dart';

import 'user.dart';

part 'moment.g.dart';

@JsonSerializable()
class Moment {
  Moment({required this.id, required this.createdAt,
    required this.content,required this.imageUrl,
  required this.mood,required this.tags,required this.user});
  int id;
  DateTime createdAt;
  String content;
  String imageUrl;
  Mood mood;
  List<Tag> tags;

  User user;

  factory Moment.fromJson(Map<String, dynamic> json) => _$MomentFromJson(json);

  Map<String, dynamic> toJson() => _$MomentToJson(this);
}

@JsonSerializable()
class Mood {
  Mood();
  String name;
  String description;
  String expressionUrl;
  int status;

  factory Mood.fromJson(Map<String, dynamic> json) => _$MoodFromJson(json);

  Map<String, dynamic> toJson() => _$MoodToJson(this);
}

@JsonSerializable()
class Tag {
  Tag();
  String name;
  String description;
  int status;

  factory Tag.fromJson(Map<String, dynamic> json) => _$TagFromJson(json);

  Map<String, dynamic> toJson() => _$TagToJson(this);
}

@JsonSerializable()
class Category {
  Category();
  int id;
  String name;
  int parentId;
  int sequence;
  int status;

  factory Category.fromJson(Map<String, dynamic> json) => _$CategoryFromJson(json);

  Map<String, dynamic> toJson() => _$CategoryToJson(this);
}

@JsonSerializable()
class MomentComment {
  MomentComment();
  int id;
  DateTime createdAt;

  factory MomentComment.fromJson(Map<String, dynamic> json) => _$MomentCommentFromJson(json);

  Map<String, dynamic> toJson() => _$MomentCommentToJson(this);
}
